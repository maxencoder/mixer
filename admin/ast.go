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

func (*AddRoute) iCommand()          {}
func (*AddDbRouter) iCommand()       {}
func (*AddTableRouter) iCommand()    {}
func (*AlterRoute) iCommand()        {}
func (*AlterDbRouter) iCommand()     {}
func (*AlterTableRouter) iCommand()  {}
func (*DeleteRoute) iCommand()       {}
func (*DeleteDbRouter) iCommand()    {}
func (*DeleteTableRouter) iCommand() {}
func (*Show) iCommand()              {}

type AddRoute struct {
	Name  string
	Route Route
}

func (n *AddRoute) Format(buf *TrackedBuffer) {
	buf.Myprintf("add route %s %v", n.Name, n.Route)
}

type RouteID string

func (n RouteID) Format(buf *TrackedBuffer) {
	buf.Myprintf("%s", string(n))
}

type AddDbRouter struct {
	Db      string
	Default RouteID
}

func (n *AddDbRouter) Format(buf *TrackedBuffer) {
	buf.Myprintf("add database router %s (default %v)", n.Db, n.Default)
}

type AddTableRouter struct {
	Db    string
	Table string
	Key   string
	Route RouteID
}

func (n *AddTableRouter) Format(buf *TrackedBuffer) {
	buf.Myprintf("add table router %s.%s (key %s, route %v)",
		n.Db, n.Table, n.Key, n.Route)
}

type AlterRoute struct {
	Name  string
	Route Route
}

func (n *AlterRoute) Format(buf *TrackedBuffer) {
	buf.Myprintf("alter route %s %v", n.Name, n.Route)
}

type AlterDbRouter struct {
	Db      string
	Default RouteID
}

func (n *AlterDbRouter) Format(buf *TrackedBuffer) {
	buf.Myprintf("alter database router %s (default %v)", n.Db, n.Default)
}

type AlterTableRouter struct {
	Db    string
	Table string
	Key   string
	Route RouteID
}

func (n *AlterTableRouter) Format(buf *TrackedBuffer) {
	buf.Myprintf("alter table router %s.%s (key %s, route %v)",
		n.Db, n.Table, n.Key, n.Route)
}

type DeleteRoute struct {
	Name string
}

func (n *DeleteRoute) Format(buf *TrackedBuffer) {
	buf.Myprintf("delete route %s", n.Name)
}

type DeleteDbRouter struct {
	Db string
}

func (n *DeleteDbRouter) Format(buf *TrackedBuffer) {
	buf.Myprintf("delete database router %s", n.Db)
}

type DeleteTableRouter struct {
	Db    string
	Table string
}

func (n *DeleteTableRouter) Format(buf *TrackedBuffer) {
	buf.Myprintf("delete table router %s.%s", n.Db, n.Table)
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

type TableIdent struct {
	Db    string
	Table string
}

type TableRouterDef struct {
	Key   string
	Route RouteID
}
