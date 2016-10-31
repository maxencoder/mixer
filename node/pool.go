package node

import (
	"fmt"
	"sync"

	"github.com/maxencoder/mixer/config"
)

// all nodes
var pool = &Pool{}

func GetNode(name string) *Node {
	return pool.GetNode(name)
}

func ParseNodes(cfg *config.Config) error {
	return pool.ParseNodes(cfg)
}

type Pool struct {
	sync.Mutex
	nodes map[string]*Node
}

func (p *Pool) GetNode(name string) *Node {
	p.Lock()
	defer p.Unlock()

	return p.nodes[name]
}

func (p *Pool) ParseNodes(cfg *config.Config) error {
	p.Lock()
	defer p.Unlock()

	p.nodes = make(map[string]*Node, len(cfg.Nodes))

	for _, v := range cfg.Nodes {
		if _, ok := p.nodes[v.Name]; ok {
			return fmt.Errorf("duplicate node [%s].", v.Name)
		}

		n, err := NewNode(v)
		if err != nil {
			return err
		}

		p.nodes[v.Name] = n
	}

	return nil
}
