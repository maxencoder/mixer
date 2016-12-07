package proxy

import (
	"github.com/maxencoder/mixer/node"
	"github.com/maxencoder/mixer/router"
)

type Schema struct {
	db string

	nodes map[string]*node.Node

	router *router.Router
}

func (s *Server) getSchema(db string) *Schema {
	return s.schemas[db]
}
