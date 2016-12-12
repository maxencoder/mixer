// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sqlparser

import (
	"errors"
	"fmt"
	"strings"
)

func Parse(sql string) (Command, error) {
	tokenizer := NewStringTokenizer(sql)
	if yyParse(tokenizer) != 0 {
		return nil, errors.New(tokenizer.LastError)
	}
	return tokenizer.ParseTree, nil
}

type AdmNode interface {
	Format(buf *TrackedBuffer)
}

func String(node AdmNode) string {
	buf := NewTrackedBuffer(nil)
	buf.Myprintf("%v", node)
	return buf.String()
}

type Command interface {
	iCommand()
	AdmNode
}

func (*AddRoute) iCommand()   {}
func (*AlterRoute) iCommand() {}
func (*Delete) iCommand()     {}
func (*Show) iCommand()       {}

type AddRoute struct {
	Name  string
	Route Route
}

func (n *AddRoute) Format(buf *TrackedBuffer) {
	buf.Myprintf("add route %s %v", n.Name, n.Route)
}

type AlterRoute struct {
	Name  string
	Route Route
}

func (n *AlterRoute) Format(buf *TrackedBuffer) {
	buf.Myprintf("alter route %s %v", n.Name, n.Route)
}

type Delete struct {
	Name string
}

func (n *Delete) Format(buf *TrackedBuffer) {
	buf.Myprintf("delete route %s", n.Name)
}

type Show struct {
}

func (n *Show) Format(buf *TrackedBuffer) {
	buf.Myprintf("show routes")
}

type Route interface {
	iRoute()
}

func (*HashRoute) iRoute()  {}
func (*RangeRoute) iRoute() {}

type HashRoute struct {
	Type   string
	Routes []string
}

func (n *HashRoute) Format(buf *TrackedBuffer) {
	buf.Myprintf("hash (type %s, route %s)",
		n.Type, strings.Join(n.Routes, ", "))
}

// HashRoute.Type
const (
	ModuloStr = "modulo"
)

type RangeRoute struct {
	Ranges KeyRangeRoutes
}

func (n *RangeRoute) Format(buf *TrackedBuffer) {
	buf.Myprintf("range (%v)", n.Ranges)
}

type KeyRangeRoutes []KeyRangeRoute

func (ns KeyRangeRoutes) Format(buf *TrackedBuffer) {
	var prefix string
	for _, n := range ns {
		buf.Myprintf("%s%v", prefix, n)
		prefix = ", "
	}
}

type KeyRangeRoute struct {
	Start, End RangeNum
	Route      string
}

func (n KeyRangeRoute) Format(buf *TrackedBuffer) {
	buf.Myprintf("%v to %v route %s", n.Start, n.End, n.Route)
}

type RangeNum struct {
	Num int64
	Inf bool
}

func (n RangeNum) Format(buf *TrackedBuffer) {
	if n.Inf {
		buf.Myprintf("inf")
	} else {
		buf.Myprintf(fmt.Sprintf("%d", n.Num))
	}
}

type TableIdent string

func (node TableIdent) Format(buf *TrackedBuffer) {
	name := string(node)
	if _, ok := keywords[strings.ToLower(name)]; ok {
		buf.Myprintf("`%s`", name)
		return
	}
	buf.Myprintf("%s", name)
}
