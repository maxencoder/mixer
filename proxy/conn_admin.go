package proxy

import (
	"github.com/maxencoder/mixer/admin"

	"github.com/maxencoder/mixer/sqlparser"
	. "github.com/siddontang/go-mysql/mysql"
)

func (c *Conn) handleAdmin(admin admin.Command, sql string) (*Result, error) {
	return nil, nil
}

func (c *Conn) handleToAdmin(admin *sqlparser.Admin) (*Result, error) {
	c.isAdminMode = true

	return nil, nil
}
