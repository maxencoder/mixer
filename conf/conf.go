package conf

import (
	"sort"
	"sync"
)

type Conf struct {
	mu sync.RWMutex

	version uint64

	schemas map[string]*Schema
}

func NewConf() *Conf {
	c := &Conf{}

	c.schemas = make(map[string]*Schema)

	return c
}

func (c *Conf) Schemas() (s []string) {
	for k, _ := range c.schemas {
		s = append(s, k)
	}

	sort.Strings(s)

	return s
}

func (c *Conf) rLock() {
	c.mu.RLock()
}

func (c *Conf) rUnlock() {
	c.version++

	c.mu.RUnlock()
}
