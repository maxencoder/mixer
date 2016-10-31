package proxy

import (
	"github.com/maxencoder/mixer/db"
	"github.com/maxencoder/mixer/node"
	. "github.com/siddontang/go-mysql/mysql"
)

func (c *Conn) isInTransaction() bool {
	return c.status&SERVER_STATUS_IN_TRANS > 0
}

func (c *Conn) isAutoCommit() bool {
	return c.status&SERVER_STATUS_AUTOCOMMIT > 0
}

func (c *Conn) handleBegin() error {
	c.status |= SERVER_STATUS_IN_TRANS

	return nil
}

func (c *Conn) handleCommit() (err error) {
	c.status &= ^SERVER_STATUS_IN_TRANS

	for _, co := range c.txConns {
		if e := co.Commit(); e != nil {
			err = e
		}
		co.Close()
	}

	c.txConns = map[*node.Node]*db.SqlConn{}

	return
}

func (c *Conn) handleRollback() error {
	return c.rollback()
}

func (c *Conn) rollback() (err error) {
	c.status &= ^SERVER_STATUS_IN_TRANS

	for _, co := range c.txConns {
		if e := co.Rollback(); e != nil {
			err = e
		}
		co.Close()
	}

	c.txConns = map[*node.Node]*db.SqlConn{}

	return
}

//if status is in_trans, need
//else if status is not autocommit, need
//else no need
func (c *Conn) needBeginTx() bool {
	return c.isInTransaction() || !c.isAutoCommit()
}
