package proxy

import (
	"fmt"
	"strings"

	"github.com/maxencoder/mixer/node"
	"github.com/maxencoder/mixer/sqlparser"
	. "github.com/siddontang/go-mysql/mysql"
)

func (c *Conn) handleAdmin(admin *sqlparser.Admin) (*Result, error) {
	c.isAdminMode = true

	return nil, err
}
