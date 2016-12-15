package proxy

import (
	"fmt"
	"net"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/maxencoder/log"
	"github.com/maxencoder/mixer/conf"
	"github.com/maxencoder/mixer/db"
	"github.com/maxencoder/mixer/node"
	"github.com/maxencoder/mixer/sqlparser"
	. "github.com/siddontang/go-mysql/mysql"
	"github.com/siddontang/go-mysql/server"
)

//client <-> proxy
type Conn struct {
	sync.Mutex

	c *server.Conn

	server *Server

	connectionId uint32

	status uint16

	charset string

	db string

	txConns map[*node.Node]*db.SqlConn

	closed bool

	lastInsertId int64
	affectedRows int64

	isAdminMode bool
}

var baseConnID uint32 = 10000

func (s *Server) newConn(co net.Conn) (c *Conn, err error) {
	c = new(Conn)

	c.connectionId = atomic.AddUint32(&baseConnID, 1)

	c.server = s

	c.status = SERVER_STATUS_AUTOCOMMIT

	c.txConns = make(map[*node.Node]*db.SqlConn)

	c.closed = false

	c.charset = DEFAULT_CHARSET

	c.c, err = server.NewConn(co, s.user, s.password, c)

	return
}

func (c *Conn) Close() error {
	if c.closed {
		return nil
	}

	if c.c != nil && !c.c.Closed() {
		c.c.Close()
	}

	c.rollback()

	c.closed = true

	return nil
}

func (c *Conn) Run() {
	defer func() {
		r := recover()
		if r != nil {
			err := r.(error)

			const size = 4096
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]

			log.Error("%v, %s", err, buf)
		}

		c.Close()
	}()

	for {
		err := c.c.HandleCommand()
		if err != nil {
			log.Error("dispatch error: %s", err.Error())
			return
		}

		if c.c.Closed() {
			return
		}
	}
}

//
//  server.Handler interface
//

func (c *Conn) UseDB(db string) error {
	if s := c.server.conf.GetSchema(db); s == nil {
		return NewDefaultError(ER_BAD_DB_ERROR, db)
	} else {
		c.db = db
	}
	return nil
}

func (c *Conn) HandleQuery(sql string) (*Result, error) {
	return c.handleQuery(sql)
}

func (c *Conn) HandleFieldList(table string, fieldWildcard string) ([]*Field, error) {
	if c.schema == nil {
		return nil, NewDefaultError(ER_NO_DB_ERROR)
	}

	// TODO
	var nodeName string
	//nodeName := c.schema.router.GetTableRouter(table).DefaultNode()

	n := c.server.getNode(nodeName)

	co, err := n.GetMasterConn()
	if err != nil {
		return nil, err
	}
	defer co.Close()

	if err = co.UseDB(c.db); err != nil {
		return nil, err
	}

	var fs []*Field

	if fs, err = co.FieldList(table, fieldWildcard); err != nil {
		return nil, err
	}

	return fs, nil
}

func (c *Conn) HandleStmtPrepare(query string) (params int, columns int, context interface{}, err error) {
	return c.handleStmtPrepare(query)
}

func (c *Conn) HandleStmtExecute(context interface{}, sql string, args []interface{}) (r *Result, err error) {
	s := context.(sqlparser.Statement)

	switch stmt := s.(type) {
	case *sqlparser.Select:
		r, err = c.handleSelect(stmt, sql, args)
	case *sqlparser.Insert:
		r, err = c.handleExec(s, sql, args)
	case *sqlparser.Update:
		r, err = c.handleExec(s, sql, args)
	case *sqlparser.Delete:
		r, err = c.handleExec(s, sql, args)
	case *sqlparser.Replace:
		r, err = c.handleExec(s, sql, args)
	default:
		err = fmt.Errorf("command %T not supported now", stmt)
	}

	return
}

func (c *Conn) HandleStmtClose(context interface{}) error {
	return nil
}

func (c *Conn) handleStmtPrepare(sql string) (int, int, interface{}, error) {
	if c.schema == nil {
		return 0, 0, nil, NewDefaultError(ER_NO_DB_ERROR)
	}

	sql = strings.TrimRight(sql, ";")

	p, err := sqlparser.Parse(sql)
	if err != nil {
		return 0, 0, nil, fmt.Errorf(`parse sql "%s" error`, sql)
	}

	var tableName string
	switch t := p.(type) {
	case *sqlparser.Select:
		tableName = nstring(t.From)
	case *sqlparser.Insert:
		tableName = nstring(t.Table)
	case *sqlparser.Update:
		tableName = nstring(t.Table)
	case *sqlparser.Delete:
		tableName = nstring(t.Table)
	case *sqlparser.Replace:
		panic("unreachable")
	default:
		return 0, 0, nil, fmt.Errorf(`unsupport prepare sql "%s"`, sql)
	}

	// XXX: not implemented
	_ = tableName
	return 0, 0, nil, NewDefaultError(ER_NO_DB_ERROR)

	/*
		r := c.schema.router.GetTableRouter(tableName)

		n := c.server.getNode(r.Nodes[0])

		var params, columns int
		if co, err := n.GetMasterConn(); err != nil {
			return 0, 0, nil, fmt.Errorf("prepare error %s", err)
		} else {
			defer co.Close()

			if err = co.UseDB(c.schema.db); err != nil {
				return 0, 0, nil, fmt.Errorf("parepre error %s", err)
			}

			if t, err := co.Prepare(sql); err != nil {
				return 0, 0, nil, fmt.Errorf("parepre error %s", err)
			} else {
				params = t.ParamNum()
				columns = t.ColumnNum()
			}
		}

		return params, columns, p, nil
	*/
}

func (c *Conn) schema() *conf.Schema {
	return c.server.conf.GetSchema(c.db)
}
