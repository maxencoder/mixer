package proxy

import (
	"fmt"
	"log"

	"github.com/maxencoder/mixer/adminparser"
	"github.com/maxencoder/mixer/router"
	"github.com/maxencoder/mixer/sqlparser"
	. "github.com/siddontang/go-mysql/mysql"
)

var astring = adminparser.String

func (c *Conn) handleAdmin(cmd adminparser.Command, sql string) (*Result, error) {
	log.Printf("%#v\n", cmd)

	switch cmd := cmd.(type) {
	case *adminparser.AddRoute:
		return c.addRoute(cmd)
	case *adminparser.AddDbRouter:
		return c.addDbRouter(cmd)
	case *adminparser.AddTableRouter:
		return c.addTableRouter(cmd)
	case *adminparser.Show:
		return c.adminShow(cmd)
	default:
		return nil, fmt.Errorf("unknown command")

		/*
		   func (*AlterRoute) iCommand()        {}
		   func (*AlterDbRouter) iCommand()     {}
		   func (*AlterTableRouter) iCommand()  {}
		   func (*DeleteRoute) iCommand()       {}
		   func (*DeleteDbRouter) iCommand()    {}
		   func (*DeleteTableRouter) iCommand() {}
		*/
	}

	return nil, nil
}

func (c *Conn) handleConfigChange(cmd adminparser.Command, sql string) (*Result, error) {
	// TODO
	return nil, nil
}

func (c *Conn) handleToAdmin(adminparser *sqlparser.Admin) (*Result, error) {
	c.isAdminMode = true

	return nil, nil
}

func (c *Conn) addRoute(cmd *adminparser.AddRoute) (*Result, error) {
	if c.db == "" {
		return nil, fmt.Errorf("database is not selected; run 'use db' first")
	}

	schema := c.server.conf.GetSchema(c.db)

	if schema == nil {
		return nil, fmt.Errorf("router for database %s undefined", c.db)
	}

	if _, err := schema.Router.NewRoute(cmd.Name, cmd.Route); err != nil {
		return nil, err
	}

	return nil, nil
}

func (c *Conn) addDbRouter(cmd *adminparser.AddDbRouter) (*Result, error) {
	if c.server.conf.GetSchema(cmd.Db) != nil {
		return nil, fmt.Errorf("router for %s is already defined", cmd.Db)
	}

	if err := c.server.conf.NewSchema(cmd.Db, string(cmd.Default)); err != nil {
		return nil, err
	}

	return nil, nil
}

func (c *Conn) addTableRouter(cmd *adminparser.AddTableRouter) (*Result, error) {
	schema := c.server.conf.GetSchema(cmd.Db)

	if schema == nil {
		return nil, fmt.Errorf("router for database %s undefined", cmd.Db)
	}

	r := schema.Router

	if tr, _ := r.GetTableRouter(cmd.Table); tr != nil {
		return nil, fmt.Errorf("table router %s already exists", cmd.Table)
	}

	if route, _ := r.GetRoute(string(cmd.Route)); route == nil {
		return nil, fmt.Errorf("route %s does not exist", cmd.Route)
	}

	ref, err := r.NewRouteRef(string(cmd.Route))

	if err != nil {
		return nil, err
	}

	new := &router.TableRouter{
		DB:    cmd.Db,
		Table: cmd.Table,
		Key:   cmd.Key,
		Route: ref,
	}

	r.SetTableRouter(cmd.Table, new)

	return nil, nil
}

func (c *Conn) adminShow(cmd *adminparser.Show) (*Result, error) {
	var err error
	var r *Resultset

	r, err = c.adminShowRoutes(cmd)

	if err != nil {
		return nil, err
	}

	return &Result{Status: c.status, Resultset: r}, nil
}

func (c *Conn) adminShowRoutes(cmd *adminparser.Show) (*Resultset, error) {
	var names []string = []string{"Command"}
	var rows [][]string

	for _, db := range c.server.conf.Schemas() {
		schema := c.server.conf.GetSchema(db)

		cmds := schema.Router.ToAstNodes()

		for _, cmd := range cmds {
			rows = append(rows, []string{astring(cmd)})
		}
	}

	return c.buildResultset(names, strings2_to_interfaces2(rows, 1))
}

func strings2_to_interfaces2(rows [][]string, nrColumns int) [][]interface{} {
	var values [][]interface{} = make([][]interface{}, len(rows))

	for i := range rows {
		values[i] = make([]interface{}, nrColumns)
		for j := range rows[i] {
			values[i][j] = rows[i][j]
		}
	}

	return values
}
