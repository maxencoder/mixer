package proxy

import (
	"fmt"
	"strings"

	"github.com/maxencoder/mixer/sqlparser"
	. "github.com/siddontang/go-mysql/mysql"
)

func (c *Conn) handleAdmin(admin *sqlparser.Admin) (*Result, error) {
	name := string(admin.Name)

	var err error

	switch strings.ToLower(name) {
	case "upnode":
		err = c.adminUpNodeServer(admin.Values)
	case "downnode":
		err = c.adminDownNodeServer(admin.Values)
	default:
		return nil, fmt.Errorf("admin %s not supported now", name)
	}

	return nil, err
}

func (c *Conn) adminUpNodeServer(values sqlparser.ValExprs) error {
	if len(values) != 3 {
		return fmt.Errorf("upnode needs 3 args, not %d", len(values))
	}

	nodeName := nstring(values[0])
	sType := strings.ToLower(nstring(values[1]))
	addr := strings.ToLower(nstring(values[2]))

	switch sType {
	case Master:
		return c.server.UpMaster(nodeName, addr)
	case Slave:
		return c.server.UpSlave(nodeName, addr)
	default:
		return fmt.Errorf("invalid server type %s", sType)
	}
}

func (c *Conn) adminDownNodeServer(values sqlparser.ValExprs) error {
	if len(values) != 2 {
		return fmt.Errorf("upnode needs 2 args, not %d", len(values))
	}

	nodeName := nstring(values[0])
	sType := strings.ToLower(nstring(values[1]))

	switch sType {
	case Master:
		return c.server.DownMaster(nodeName)
	case Slave:
		return c.server.DownSlave(nodeName)
	default:
		return fmt.Errorf("invalid server type %s", sType)
	}
}
