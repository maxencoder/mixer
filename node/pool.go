package node

import (
	"fmt"
	"sync"

	"github.com/maxencoder/mixer/config"
)

// all nodes
var pool = &Pool{}

func FullList() []string {
	return pool.FullList()
}

func GetNode(name string) *Node {
	return pool.GetNode(name)
}

func SetNode(name string, n *Node) {
	pool.SetNode(name, n)
}

func InitPool() {
	pool.nodes = make(map[string]*Node)
}

func ParseNodes(cfg *config.Config) error {
	return pool.ParseNodes(cfg)
}

type Pool struct {
	sync.Mutex
	nodes map[string]*Node
}

func (p *Pool) FullList() []string {
	p.Lock()
	defer p.Unlock()

	var l []string

	for n, _ := range p.nodes {
		l = append(l, n)
	}

	return l
}

func (p *Pool) GetNode(name string) *Node {
	p.Lock()
	defer p.Unlock()

	return p.nodes[name]
}

func (p *Pool) SetNode(name string, n *Node) {
	p.Lock()
	defer p.Unlock()

	p.nodes[name] = n
}

func (p *Pool) ParseNodes(cfg *config.Config) error {
	p.Lock()
	defer p.Unlock()

	p.nodes = make(map[string]*Node, len(cfg.Nodes))

	for _, v := range cfg.Nodes {
		if _, ok := p.nodes[v.Name]; ok {
			return fmt.Errorf("duplicate node [%s].", v.Name)
		}

		n, err := NewNodeFromConfig(v)

		if err != nil {
			return err
		}

		p.nodes[v.Name] = n
	}

	return nil
}
