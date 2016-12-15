package conf

import (
	"fmt"

	"github.com/maxencoder/mixer/node"
	"github.com/maxencoder/mixer/router"
)

type Schema struct {
	DB string

	Router *router.Router
}

func (c *Conf) GetSchema(db string) *Schema {
	c.RLock()
	defer c.RUnlock()

	return c.schemas[db]
}

func (c *Conf) NewSchema(db string, defaultNode string) error {
	if n := node.GetNode(defaultNode); n == nil {
		return fmt.Errorf("node %s does not exist", defaultNode)
	}

	r := router.NewRouter(db, defaultNode)

	c.schemas[db] = &Schema{DB: db, Router: r}

	return nil
}
