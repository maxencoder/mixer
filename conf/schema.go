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
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.schemas[db]
}

func (c *Conf) NewSchema(db string, defaultNode string) error {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if n := node.GetNode(defaultNode); n == nil {
		return fmt.Errorf("node %s does not exist", defaultNode)
	}

	r, err := router.NewRouter(db, defaultNode)

	if err != nil {
		return err
	}

	c.schemas[db] = &Schema{DB: db, Router: r}

	return nil
}
