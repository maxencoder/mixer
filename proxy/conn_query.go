package proxy

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/maxencoder/mixer/db"
	"github.com/maxencoder/mixer/hack"
	"github.com/maxencoder/mixer/sqlparser"
	. "github.com/siddontang/go-mysql/mysql"
)

func (c *Conn) handleQuery(sql string) (r *Result, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("execute %s error %v", sql, e)
			return
		}
	}()

	sql = strings.TrimRight(sql, ";")

	var stmt sqlparser.Statement

	stmt, err = sqlparser.Parse(sql)
	if err != nil {
		return c.handleUnknown(sql, nil)
	}

	switch v := stmt.(type) {
	case *sqlparser.Select:
		r, err = c.handleSelect(v, sql, nil)
	case *sqlparser.Insert:
		r, err = c.handleExec(stmt, sql, nil)
	case *sqlparser.Update:
		r, err = c.handleExec(stmt, sql, nil)
	case *sqlparser.Delete:
		r, err = c.handleExec(stmt, sql, nil)
	case *sqlparser.Replace:
		r, err = c.handleExec(stmt, sql, nil)
	case *sqlparser.Set:
		err = c.handleSet(v)
	case *sqlparser.Begin:
		err = c.handleBegin()
	case *sqlparser.Commit:
		err = c.handleCommit()
	case *sqlparser.Rollback:
		err = c.handleRollback()
	case *sqlparser.SimpleSelect:
		r, err = c.handleSimpleSelect(sql, v)
	case *sqlparser.Show:
		r, err = c.handleShow(sql, v)
	case *sqlparser.Admin:
		r, err = c.handleAdmin(v)
	default:
		err = fmt.Errorf("statement %T not support now", stmt)
	}

	return
}

func (c *Conn) getShardList(stmt sqlparser.Statement, bindVars map[string]interface{}) ([]*Node, error) {
	if c.schema == nil {
		return nil, NewDefaultError(ER_NO_DB_ERROR)
	}

	ns, err := sqlparser.GetStmtShardList(stmt, c.schema.rule, bindVars)
	if err != nil {
		return nil, err
	}

	if len(ns) == 0 {
		return nil, nil
	}

	n := make([]*Node, 0, len(ns))
	for _, name := range ns {
		n = append(n, c.server.getNode(name))
	}
	return n, nil
}

func (c *Conn) getConn(n *Node, isSelect bool) (co *db.SqlConn, err error) {
	if !c.needBeginTx() {
		if isSelect {
			co, err = n.getSelectConn()
		} else {
			co, err = n.getMasterConn()
		}
		if err != nil {
			return
		}
	} else {
		var ok bool
		c.Lock()
		co, ok = c.txConns[n]
		c.Unlock()

		if !ok {
			if co, err = n.getMasterConn(); err != nil {
				return
			}

			if err = co.Begin(); err != nil {
				return
			}

			c.Lock()
			c.txConns[n] = co
			c.Unlock()
		}
	}

	//todo, set conn charset, etc...
	if err = co.UseDB(c.schema.db); err != nil {
		return
	}

	if err = co.SetCharset(c.charset); err != nil {
		return
	}

	return
}

func (c *Conn) getDefaultConn(isSelect bool) (*db.SqlConn, error) {
	n := c.schema.nodes["node1"]

	var co *db.SqlConn
	var err error
	co, err = c.getConn(n, isSelect)
	if err != nil {
		return nil, err
	}

	return co, err
}

func (c *Conn) getShardConns(isSelect bool, stmt sqlparser.Statement, bindVars map[string]interface{}) ([]*db.SqlConn, error) {
	nodes, err := c.getShardList(stmt, bindVars)
	if err != nil {
		return nil, err
	} else if nodes == nil {
		return nil, nil
	}

	conns := make([]*db.SqlConn, 0, len(nodes))

	var co *db.SqlConn
	for _, n := range nodes {
		co, err = c.getConn(n, isSelect)
		if err != nil {
			break
		}

		conns = append(conns, co)
	}

	return conns, err
}

func (c *Conn) executeInShard(conns []*db.SqlConn, sql string, args []interface{}) ([]*Result, error) {
	var wg sync.WaitGroup
	wg.Add(len(conns))

	rs := make([]interface{}, len(conns))

	f := func(rs []interface{}, i int, co *db.SqlConn) {
		r, err := co.Execute(sql, args...)
		if err != nil {
			rs[i] = err
		} else {
			rs[i] = r
		}

		wg.Done()
	}

	for i, co := range conns {
		go f(rs, i, co)
	}

	wg.Wait()

	var err error
	r := make([]*Result, len(conns))
	for i, v := range rs {
		if e, ok := v.(error); ok {
			err = e
			break
		}
		r[i] = rs[i].(*Result)
	}

	return r, err
}

func (c *Conn) closeShardConns(conns []*db.SqlConn, rollback bool) {
	if c.isInTransaction() {
		return
	}

	for _, co := range conns {
		if rollback {
			co.Rollback()
		}

		co.Close()
	}
}

func (c *Conn) newEmptyResult(stmt *sqlparser.Select) *Result {
	rs := c.newEmptyResultset(stmt)

	r := new(Result)

	r.Resultset = rs
	r.Status = c.status

	return r
}

func (c *Conn) newEmptyResultset(stmt *sqlparser.Select) *Resultset {
	r := new(Resultset)
	r.Fields = make([]*Field, len(stmt.SelectExprs))

	for i, expr := range stmt.SelectExprs {
		r.Fields[i] = &Field{}
		switch e := expr.(type) {
		case *sqlparser.StarExpr:
			r.Fields[i].Name = []byte("*")
		case *sqlparser.NonStarExpr:
			if e.As != nil {
				r.Fields[i].Name = e.As
				r.Fields[i].OrgName = hack.Slice(nstring(e.Expr))
			} else {
				r.Fields[i].Name = hack.Slice(nstring(e.Expr))
			}
		default:
			r.Fields[i].Name = hack.Slice(nstring(e))
		}
	}

	r.Values = make([][]interface{}, 0)
	r.RowDatas = make([]RowData, 0)

	return r
}

func makeBindVars(args []interface{}) map[string]interface{} {
	bindVars := make(map[string]interface{}, len(args))

	for i, v := range args {
		bindVars[fmt.Sprintf("v%d", i+1)] = v
	}

	return bindVars
}

func (c *Conn) handleUnknown(sql string, args []interface{}) (*Result, error) {
	co, err := c.getDefaultConn(true)
	if err != nil {
		return nil, err
	}

	var r *Result

	r, err = co.Execute(sql, args...)
	co.Close()

	return r, err
}

func (c *Conn) handleSelect(stmt *sqlparser.Select, sql string, args []interface{}) (*Result, error) {
	bindVars := makeBindVars(args)

	conns, err := c.getShardConns(true, stmt, bindVars)
	if err != nil {
		return nil, err
	} else if conns == nil {
		return c.newEmptyResult(stmt), nil
	}

	var rs []*Result

	rs, err = c.executeInShard(conns, sql, args)

	c.closeShardConns(conns, false)

	if err != nil {
		return nil, err
	}

	return c.mergeSelectResult(rs, stmt)
}

func (c *Conn) beginShardConns(conns []*db.SqlConn) error {
	if c.isInTransaction() {
		return nil
	}

	for _, co := range conns {
		if err := co.Begin(); err != nil {
			return err
		}
	}

	return nil
}

func (c *Conn) commitShardConns(conns []*db.SqlConn) error {
	if c.isInTransaction() {
		return nil
	}

	for _, co := range conns {
		if err := co.Commit(); err != nil {
			return err
		}
	}

	return nil
}

func (c *Conn) handleExec(stmt sqlparser.Statement, sql string, args []interface{}) (*Result, error) {
	bindVars := makeBindVars(args)

	conns, err := c.getShardConns(false, stmt, bindVars)
	if err != nil {
		return nil, err
	} else if conns == nil {
		return nil, nil
	}

	var rs []*Result

	if len(conns) == 1 {
		rs, err = c.executeInShard(conns, sql, args)
	} else {
		//for multi nodes, 2PC simple, begin, exec, commit
		//if commit error, data maybe corrupt
		for {
			if err = c.beginShardConns(conns); err != nil {
				break
			}

			if rs, err = c.executeInShard(conns, sql, args); err != nil {
				break
			}

			err = c.commitShardConns(conns)
			break
		}
	}

	c.closeShardConns(conns, err != nil)

	if err != nil {
		return nil, err
	}

	return c.mergeExecResult(rs), nil
}

func (c *Conn) mergeExecResult(rs []*Result) *Result {
	r := new(Result)

	for _, v := range rs {
		r.Status |= v.Status
		r.AffectedRows += v.AffectedRows
		if r.InsertId == 0 {
			r.InsertId = v.InsertId
		} else if r.InsertId > v.InsertId {
			//last insert id is first gen id for multi row inserted
			//see http://dev.mysql.com/doc/refman/5.6/en/information-functions.html#function_last-insert-id
			r.InsertId = v.InsertId
		}
	}

	if r.InsertId > 0 {
		c.lastInsertId = int64(r.InsertId)
	}

	c.affectedRows = int64(r.AffectedRows)

	return r
}

func (c *Conn) mergeSelectResult(rs []*Result, stmt *sqlparser.Select) (*Result, error) {
	r := rs[0]
	s := r.Resultset

	r.Status |= c.status

	for i := 1; i < len(rs); i++ {
		r.Status |= rs[i].Status

		//check fields equal

		for j := range rs[i].Values {
			s.Values = append(s.Values, rs[i].Values[j])
			s.RowDatas = append(s.RowDatas, rs[i].RowDatas[j])
		}
	}

	//to do order by, group by, limit offset
	c.sortSelectResult(s, stmt)
	//to do, add log here, sort may error because order by key not exist in resultset fields

	if err := c.limitSelectResult(s, stmt); err != nil {
		return nil, err
	}

	return r, nil
}

func (c *Conn) sortSelectResult(r *Resultset, stmt *sqlparser.Select) error {
	if stmt.OrderBy == nil {
		return nil
	}

	sk := make([]SortKey, len(stmt.OrderBy))

	for i, o := range stmt.OrderBy {
		sk[i].Name = nstring(o.Expr)
		sk[i].Direction = o.Direction
	}

	s, err := newResultsetSorter(r, sk)
	if err != nil {
		return err
	}
	sort.Sort(s)
	return nil
}

func (c *Conn) limitSelectResult(r *Resultset, stmt *sqlparser.Select) error {
	if stmt.Limit == nil {
		return nil
	}

	var offset, count int64
	var err error
	if stmt.Limit.Offset == nil {
		offset = 0
	} else {
		if o, ok := stmt.Limit.Offset.(sqlparser.NumVal); !ok {
			return fmt.Errorf("invalid select limit %s", nstring(stmt.Limit))
		} else {
			if offset, err = strconv.ParseInt(hack.String([]byte(o)), 10, 64); err != nil {
				return err
			}
		}
	}

	if o, ok := stmt.Limit.Rowcount.(sqlparser.NumVal); !ok {
		return fmt.Errorf("invalid limit %s", nstring(stmt.Limit))
	} else {
		if count, err = strconv.ParseInt(hack.String([]byte(o)), 10, 64); err != nil {
			return err
		} else if count < 0 {
			return fmt.Errorf("invalid limit %s", nstring(stmt.Limit))
		}
	}

	if offset+count > int64(len(r.Values)) {
		count = int64(len(r.Values)) - offset
	}

	r.Values = r.Values[offset : offset+count]
	r.RowDatas = r.RowDatas[offset : offset+count]

	return nil
}
