package proxy

import (
	"fmt"
	"strings"

	"github.com/maxencoder/mixer/sqlparser"
	. "github.com/siddontang/go-mysql/mysql"
)

var nstring = sqlparser.String

func (c *Conn) handleSet(stmt *sqlparser.Set) error {
	if len(stmt.Exprs) != 1 {
		return fmt.Errorf("must set one item once, not %s", nstring(stmt))
	}

	k := string(stmt.Exprs[0].Name.String())

	switch strings.ToUpper(k) {
	case `AUTOCOMMIT`:
		return c.handleSetAutoCommit(stmt.Exprs[0].Expr)
	case `NAMES`:
		return c.handleSetNames(stmt.Exprs[0].Expr)
	default:
		return fmt.Errorf("set %s is not supported now", k)
	}
}

func (c *Conn) handleSetAutoCommit(val sqlparser.ValExpr) error {
	value, ok := val.(sqlparser.NumVal)
	if !ok {
		return fmt.Errorf("set autocommit error")
	}
	switch value[0] {
	case '1':
		c.status |= SERVER_STATUS_AUTOCOMMIT
	case '0':
		c.status &= ^SERVER_STATUS_AUTOCOMMIT
	default:
		return fmt.Errorf("invalid autocommit flag %s", value)
	}

	return nil
}

func (c *Conn) handleSetNames(val sqlparser.ValExpr) error {
	value, ok := val.(sqlparser.StrVal)
	if !ok {
		return fmt.Errorf("set names charset error")
	}

	charset := strings.ToLower(string(value))
	c.charset = charset

	return nil
}
