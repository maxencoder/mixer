package proxy

import (
	"errors"
	"fmt"
	"runtime/debug"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/maxencoder/log"
	"github.com/maxencoder/mixer/adminparser"
	"github.com/maxencoder/mixer/db"
	"github.com/maxencoder/mixer/hack"
	"github.com/maxencoder/mixer/node"
	"github.com/maxencoder/mixer/sqlparser"
	. "github.com/siddontang/go-mysql/mysql"
)

func (c *Conn) handleQuery(sql string) (r *Result, err error) {
	defer func() {
		if e := recover(); e != nil {
			debug.PrintStack()
			err = fmt.Errorf("execute %s error %v", sql, e)
			return
		}
	}()

	sql = strings.TrimRight(sql, ";")

	if c.isAdminMode {
		cmd, err := adminparser.Parse(sql)
		if err != nil {
			log.Info("failed to parse command: %s /* %s */", sql, err)
			return nil, err
		}

		return c.handleAdmin(cmd, sql)
	}

	var stmt sqlparser.Statement

	stmt, err = sqlparser.Parse(sql)
	if err != nil {
		//log.Printf("failed to parse query: %s /* %s */", sql, err)
		return nil, err
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
		//err = c.handleSet(v)
		err = fmt.Errorf("statement %T is not supported", stmt)
	case *sqlparser.Begin:
		//err = c.handleBegin()
		err = fmt.Errorf("statement %T is not supported", stmt)
	case *sqlparser.Commit:
		//err = c.handleCommit()
		err = fmt.Errorf("statement %T is not supported", stmt)
	case *sqlparser.Rollback:
		//err = c.handleRollback()
		err = fmt.Errorf("statement %T is not supported", stmt)
	case *sqlparser.SimpleSelect:
		r, err = c.handleSimpleSelect(sql, v)
	case *sqlparser.Show:
		r, err = c.handleShow(sql, v)
	case *sqlparser.Admin:
		r, err = c.handleToAdmin(v)
	default:
		err = fmt.Errorf("statement %T is not supported", stmt)
	}

	return
}

func (c *Conn) getConn(n *node.Node, isSelect bool) (co *db.SqlConn, err error) {
	if !c.needBeginTx() {
		if isSelect {
			co, err = n.GetSelectConn()
		} else {
			co, err = n.GetMasterConn()
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
			if co, err = n.GetMasterConn(); err != nil {
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

	if err = co.UseDB(c.db); err != nil {
		return
	}

	// TODO: do we need to restore conn state? conn is shared.
	/*
		if err = co.SetCharset(c.charset); err != nil {
			return
		}
	*/

	return
}

func (c *Conn) getDefaultConn(isSelect bool) (co *db.SqlConn, err error) {
	schema := c.schema()
	if schema == nil {
		return nil, NewDefaultError(ER_NO_DB_ERROR)
	}

	node := schema.Router.DefaultNode

	n := c.server.getNode(node)

	if co, err = c.getConn(n, isSelect); err != nil {
		return nil, err
	}

	return co, err
}

func (c *Conn) getConns(plans []*sqlparser.ExecPlan, isSelect bool) error {
	var err error

	for _, p := range plans {
		p.Conn, err = c.getConn(node.GetNode(p.Node), isSelect)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Conn) executePlans(plans []*sqlparser.ExecPlan) ([]*Result, error) {
	var wg sync.WaitGroup
	wg.Add(len(plans))

	var args []interface{} // TODO: send params properly

	rs := make([]interface{}, len(plans))

	f := func(rs []interface{}, i int, p *sqlparser.ExecPlan) {
		r, err := p.Conn.Execute(p.Sql(), args...)
		if err != nil {
			rs[i] = err
		} else {
			rs[i] = r
		}

		wg.Done()
	}

	for i, p := range plans {
		go f(rs, i, p)
	}

	wg.Wait()

	var err error
	r := make([]*Result, 0, len(plans))
	for i, v := range rs {
		if e, ok := v.(error); ok {
			if p := plans[i]; p.IsMirror {
				log.Warn("exec on mirrored conn to %s failed: %s", p.Node, e)
			} else {
				err = e
				break
			}
		}

		// don't need results from mirrored conns
		if plans[i].IsMirror {
			continue
		}

		r = append(r, rs[i].(*Result))
	}

	return r, err
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

func (c *Conn) closePlanConns(plans []*sqlparser.ExecPlan, rollback bool) {
	if c.isInTransaction() {
		return
	}

	for _, p := range plans {
		if rollback {
			p.Conn.Rollback()
		}

		p.Conn.Close()
	}
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
			if !e.As.EqualString("") {
				r.Fields[i].Name = hack.Slice(e.As.String())
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

func (c *Conn) handleUnparsedSelect(sql string, args []interface{}) (*Result, error) {
	if c.schema == nil {
		return nil, NewDefaultError(ER_NO_DB_ERROR)
	}

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

	schema := c.schema()
	if c.schema == nil {
		return nil, NewDefaultError(ER_NO_DB_ERROR)
	}

	plans, err := sqlparser.RouteStmt(stmt, sql, schema.Router, bindVars)
	if err != nil {
		return nil, err
	}
	if plans == nil {
		return c.newEmptyResult(stmt), nil
	}

	if len(plans) > 1 {
		if err := c.canSelectInManyShards(stmt); err != nil {
			return nil, err
		}
	}

	if err := c.getConns(plans, true); err != nil {
		return nil, err
	}

	var rs []*Result

	rs, err = c.executePlans(plans)

	c.closePlanConns(plans, false)

	if err != nil {
		return nil, err
	}

	return c.mergeSelectResult(rs, stmt)
}

func (c *Conn) beginPlanConns(plans []*sqlparser.ExecPlan) error {
	if c.isInTransaction() {
		return nil
	}

	for _, p := range plans {
		if err := p.Conn.Begin(); err != nil {
			return err
		}
	}

	return nil
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

func (c *Conn) commitPlanConns(plans []*sqlparser.ExecPlan) error {
	if c.isInTransaction() {
		return nil
	}

	for _, p := range plans {
		if err := p.Conn.Commit(); err != nil {
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

	schema := c.schema()
	if c.schema == nil {
		return nil, NewDefaultError(ER_NO_DB_ERROR)
	}

	plans, err := sqlparser.RouteStmt(stmt, sql, schema.Router, bindVars)
	if err != nil {
		return nil, err
	}
	if plans == nil {
		return nil, nil
	}

	if err := c.getConns(plans, true); err != nil {
		return nil, err
	}

	var rs []*Result

	if len(plans) == 1 {
		rs, err = c.executePlans(plans)
	} else {
		//for multi nodes, 2PC simple, begin, exec, commit
		//if commit error, data maybe corrupt
		for {
			if err = c.beginPlanConns(plans); err != nil {
				break
			}

			if rs, err = c.executePlans(plans); err != nil {
				break
			}

			err = c.commitPlanConns(plans)
			break
		}
	}

	c.closePlanConns(plans, err != nil)

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
	if err := c.sortSelectResult(s, stmt); err != nil {
		return nil, err
	}

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

	s, err := NewResultSetSorter(r, sk)
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

func (c *Conn) canSelectInManyShards(stmt *sqlparser.Select) error {
	if stmt.Distinct != "" {
		return errors.New("Distict not supported in multi-shard queries")
	}

	if stmt.SelectExprs == nil {
		return errors.New("SelectExprs not defined")
	}

	for _, se := range stmt.SelectExprs {
		switch se := se.(type) {
		case *sqlparser.StarExpr:
			continue
		case *sqlparser.NonStarExpr:
			switch e := se.Expr.(type) {
			case *sqlparser.ColName:
				continue
			case sqlparser.NumVal:
				continue
			case sqlparser.StrVal:
				continue
			case sqlparser.BoolVal:
				continue
			default:
				return fmt.Errorf("unexpected SelectExpr %T", e)
			}
		default:
			return fmt.Errorf("unexpected SelectExpr %T", se)
		}
	}
	return nil
}
