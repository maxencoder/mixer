package proxy

import (
	"fmt"
	"log"

	"github.com/maxencoder/mixer/adminparser"

	"github.com/maxencoder/mixer/sqlparser"
	. "github.com/siddontang/go-mysql/mysql"
)

func (c *Conn) handleAdmin(cmd adminparser.Command, sql string) (*Result, error) {
	log.Printf("%#v\n", cmd)

	switch cmd := cmd.(type) {
	case *adminparser.AddRoute:
		return c.addRoute(cmd)
	case *adminparser.AddDbRouter:
		return c.addDbRouter(cmd)

		/*
		   func (*AddDbRouter) iCommand()       {}
		   func (*AddTableRouter) iCommand()    {}
		   func (*AlterRoute) iCommand()        {}
		   func (*AlterDbRouter) iCommand()     {}
		   func (*AlterTableRouter) iCommand()  {}
		   func (*DeleteRoute) iCommand()       {}
		   func (*DeleteDbRouter) iCommand()    {}
		   func (*DeleteTableRouter) iCommand() {}
		   func (*Show) iCommand()              {}
		*/
	}

	return nil, nil
}

func (c *Conn) handleToAdmin(adminparser *sqlparser.Admin) (*Result, error) {
	c.isAdminMode = true

	return nil, nil
}

func (c *Conn) addRoute(cmd *adminparser.AddRoute) (*Result, error) {
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
