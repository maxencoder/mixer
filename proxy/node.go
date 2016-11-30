package proxy

import (
	"fmt"

	"github.com/maxencoder/mixer/node"
)

func (s *Server) getNode(name string) *node.Node {
	return node.GetNode(name)
}

func (s *Server) getNodes(names []string) []*node.Node {
	n := make([]*node.Node, 0, len(names))
	for _, name := range names {
		n = append(n, s.getNode(name))
	}
	return n
}

func (s *Server) UpMaster(node string, addr string) error {
	n := s.getNode(node)
	if n == nil {
		return fmt.Errorf("invalid node %s", node)
	}

	return n.UpMaster(addr)
}

func (s *Server) UpSlave(node string, addr string) error {
	n := s.getNode(node)
	if n == nil {
		return fmt.Errorf("invalid node %s", node)
	}

	return n.UpSlave(addr)
}
func (s *Server) DownMaster(node string) error {
	n := s.getNode(node)
	if n == nil {
		return fmt.Errorf("invalid node %s", node)
	}
	return n.DownMaster()
}

func (s *Server) DownSlave(node string) error {
	n := s.getNode(node)
	if n == nil {
		return fmt.Errorf("invalid node [%s].", node)
	}
	return n.DownSlave()
}
