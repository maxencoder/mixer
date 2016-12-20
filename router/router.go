package router

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/maxencoder/mixer/adminparser"
	"github.com/maxencoder/mixer/node"
)

const (
	MIRROR_R = iota
	MIRROR_W
	MIRROR_RW
)

/*
	Router
		|-TableRouter
			|-Route
				|-Route...
*/

type Router struct {
	sync.RWMutex

	version uint64

	DB string

	tr map[string]*TableRouter // keyed by table in lower case

	rt map[string]Route

	DefaultNode        string
	DefaultTableRouter *TableRouter
}

func NewRouter(db string, defaultNode string) (*Router, error) {
	r := &Router{DB: db}

	r.rt = make(map[string]Route)
	r.tr = make(map[string]*TableRouter)

	if _, err := r.NewRouteRef(defaultNode); err != nil {
		return nil, err
	}

	r.DefaultNode = defaultNode
	r.DefaultTableRouter = r.NewDefaultTableRouter()

	return r, nil
}

func (r *Router) Clone() *Router {
	new := &Router{DB: r.DB}

	new.version = r.version + 1

	new.rt = make(map[string]Route)
	new.tr = make(map[string]*TableRouter)

	for k, v := range r.tr {
		new.tr[k] = v
	}
	for k, v := range r.rt {
		new.rt[k] = v
	}

	new.DefaultNode = r.DefaultNode
	new.DefaultTableRouter = r.DefaultTableRouter

	return new
}

func (r *Router) GetRoute(name string) (ro Route, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = err.(error)
		}
	}()

	ro = r.getRoute(name)

	return
}

func (r *Router) getRoute(name string) Route {
	r.RLock()
	defer r.RUnlock()

	ro, ok := r.rt[name]
	if !ok {
		panic(fmt.Errorf("route '%s' does not exist", name))
	}
	return ro
}

func (r *Router) routeExists(name string) bool {
	r.RLock()
	defer r.RUnlock()

	_, ok := r.rt[name]

	return ok
}

func (r *Router) SetRoute(name string, route Route) {
	r.Lock()
	defer r.Unlock()

	r.rt[name] = route
}

func (r *Router) DeleteRoute(name string) error {
	if _, err := r.GetRoute(name); err != nil {
		return err
	}

	r.Lock()
	defer r.Unlock()

	if ref := r.anythingLinkedToRoute(name); ref != "" {
		return fmt.Errorf("route '%s' is referred by '%s'", name, ref)
	}

	delete(r.rt, name)

	return nil
}

func (r *Router) NewTableRouter(db, table, key, route string) (*TableRouter, error) {
	ref, err := r.NewRouteRef(route)

	if err != nil {
		return nil, err
	}

	new := &TableRouter{
		DB:    db,
		Table: table,
		Key:   key,
		Route: ref,
	}

	r.SetTableRouter(table, new)

	return new, nil
}

func (r *Router) GetTableRouterOrDefault(name string) *TableRouter {
	r.RLock()
	defer r.RUnlock()

	tr, ok := r.tr[name]
	if !ok {
		return r.DefaultTableRouter
	}
	return tr
}

func (r *Router) GetTableRouter(name string) (*TableRouter, error) {
	r.RLock()
	defer r.RUnlock()

	tr, ok := r.tr[name]
	if !ok {
		return nil, fmt.Errorf("table router '%s' does not exist", name)
	}
	return tr, nil
}

func (r *Router) SetTableRouter(name string, router *TableRouter) {
	r.Lock()
	defer r.Unlock()

	r.tr[name] = router
}

func (r *Router) DeleteTableRouter(name string) error {
	if _, err := r.GetTableRouter(name); err != nil {
		return err
	}

	r.Lock()
	defer r.Unlock()

	delete(r.tr, name)

	return nil
}

func (r *Router) NewDefaultTableRouter() *TableRouter {
	return &TableRouter{
		DB:    r.DB,
		Table: "",
		Key:   "",

		IsDefault: true,

		Route: RouteRef{router: r, to: r.DefaultNode},
	}
}

func (r *Router) NewRoute(name string, rt adminparser.Route) (Route, error) {
	var new Route
	var err error

	switch rt := rt.(type) {
	case *adminparser.HashRoute:
		new, err = r.NewModuloHashRoute(name, len(rt.Routes), rt.Routes)
	case *adminparser.MirrorRoute:
		new, err = r.NewMirrorRoute(name, rt.Kind, rt.Main, rt.Mirrors)
	case *adminparser.RangeRoute:
		new, err = r.NewRangeRoute(name, rt)
	}

	if err != nil {
		return nil, err
	}

	return new, nil
}

func (r *Router) NewNodeRoute(node string) (Route, error) {
	n := &NodeRoute{Node: node}

	if ex, _ := r.GetRoute(node); ex != nil {
		return nil, fmt.Errorf("route '%s' already exists", node)
	}

	r.SetRoute(node, n)

	return n, nil
}

func (r *Router) NewRouteRef(to string) (RouteRef, error) {
	if _, err := r.GetRoute(to); err != nil {
		if n := node.GetNode(to); n != nil {
			r.NewNodeRoute(to)

			return RouteRef{router: r, to: to}, nil
		}

		return RouteRef{}, err
	}

	return RouteRef{router: r, to: to}, nil
}

func (r *Router) NewModuloHashRoute(name string, N int, routes []string) (Route, error) {
	n := &ModuloHashRoute{name: name, router: r, N: N}

	for _, ro := range routes {
		rf, err := r.NewRouteRef(ro)

		if err != nil {
			return nil, err
		}

		n.Routes = append(n.Routes, rf)
	}

	if ex, _ := r.GetRoute(name); ex != nil {
		return nil, fmt.Errorf("route '%s' already exists", name)
	}

	r.SetRoute(name, n)

	return n, nil
}

func (r *Router) NewRangeRoute(name string, rt *adminparser.RangeRoute) (Route, error) {
	new := &RangeRoute{name: name, router: r}

	for _, ra := range rt.Ranges {
		rf, err := r.NewRouteRef(ra.Route)

		if err != nil {
			return nil, err
		}

		new.Ranges = append(new.Ranges, Range{
			KeyRange: KeyRange{
				astkey2key(ra.Start, true),
				astkey2key(ra.End, false)},
			Route: rf,
		})
	}

	if ex, _ := r.GetRoute(name); ex != nil {
		return nil, fmt.Errorf("route '%s' already exists", name)
	}

	r.SetRoute(name, new)

	return new, nil
}

func (r *Router) NewMirrorRoute(name string, kind string, main string, mirror []string) (Route, error) {
	mrf, err := r.NewRouteRef(main)

	if err != nil {
		return nil, err
	}

	n := &MirrorRoute{name: name, router: r, Main: mrf}

	switch kind {
	case adminparser.MirrorR:
		n.Kind = MIRROR_R
	case adminparser.MirrorW:
		n.Kind = MIRROR_W
	case adminparser.MirrorRW:
		n.Kind = MIRROR_RW
	default:
		return nil, fmt.Errorf("unknown mirror type %s", kind)
	}

	for _, ro := range mirror {
		rf, err := r.NewRouteRef(ro)

		if err != nil {
			return nil, err
		}

		n.Mirror = append(n.Mirror, rf)
	}

	if ex, _ := r.GetRoute(name); ex != nil {
		return nil, fmt.Errorf("route '%s' already exists", name)
	}

	r.SetRoute(name, n)

	return n, nil
}

/*
func (r *Router) NewLookupRoute(name string, query string) (Route, error) {

	if ex, _ := r.GetRoute(name); ex != nil {
		return nil, fmt.Errorf("route '%s' already exists", name)
	}

	r.SetRoute(name, n)

	return n, nil
}

*/

func (r *Router) anythingLinkedToRoute(name string) string {
	for n, rt := range r.rt {
		for _, rf := range rt.LinkedTo() {
			if rf.to == name {
				return n
			}
		}
	}
	for n, tr := range r.tr {
		if tr.Route.to == name {
			return n
		}
	}
	if r.DefaultNode == name {
		return "default router route"
	}

	return ""
}

func (r *Router) FullList() []string {
	var f []string

	for _, tr := range r.tr {
		f = append(f, tr.Route.FullList()...)
	}

	f = append(f, r.DefaultNode)

	return f
}

func (r *Router) ToAst() adminparser.AdmNode {
	return &adminparser.AddDbRouter{
		Db:      r.DB,
		Default: routeId(r.DefaultNode),
	}
}

func (r *Router) ToAstNodes() (cmds []adminparser.AdmNode) {
	cmds = append(cmds, r.ToAst())

	seen := make(map[string]bool)

	visit := func(route Route) {
		if _, ok := route.(*NodeRoute); ok {
			return
		}
		if seen[route.Name()] {
			return
		}
		seen[route.Name()] = true

		cmds = append(cmds, route.ToAst())
	}

	for _, tr := range r.tr {
		tr.Route.Route().PostOrder(visit)

		cmds = append(cmds, tr.ToAst())
	}

	// everything else defined but not referenced yet
	for _, rt := range r.rt {
		visit(rt)
	}

	return cmds
}

type TableRouter struct {
	DB    string
	Table string
	Key   string

	IsDefault bool

	Route RouteRef
}

func (r *TableRouter) FullList() []string {
	l := r.Route.FullList()

	sort.Strings(l)

	return l
}

func (r *TableRouter) FullMirrorList(isSelect bool) []string {
	l := r.Route.FullMirrorList(isSelect)

	sort.Strings(l)

	return l
}

func (r *TableRouter) FindForKeys(ks []Key) *RoutingResult {
	return r.Route.Route().FindForKeys(ks)
}

func (r *TableRouter) ToAst() adminparser.AdmNode {
	return &adminparser.AddTableRouter{
		Db:    r.DB,
		Table: r.Table,
		Key:   r.Key,
		Route: routeId(r.Route.to),
	}
}

/*
	Result of routing:

	keyA -> node1, node2, etc	// node2 is a mirror node for keyA;
								// it receives a copy of request for keyA;
								// mirrored results are discarded by mixer
	keyB -> node2, node3, etc
*/

type RoutingResult struct {
	R map[Key]KeyResult
}

type KeyResult struct {
	Node   string
	Mirror []string
}

func NewRoutingResult() *RoutingResult {
	return &RoutingResult{R: make(map[Key]KeyResult)}
}

func (r *RoutingResult) Merge(q *RoutingResult) {
	for k, kr := range q.R {
		r.R[k] = mergeKeyResult(r.R[k], kr)
	}
}

func (r *RoutingResult) SplitByNode() map[string][]Key {
	m := make(map[string][]Key)

	for k, kr := range r.R {
		m[kr.Node] = append(m[kr.Node], k)
	}

	return m
}

func (r *RoutingResult) SplitByMirrorNode() map[string][]Key {
	m := make(map[string][]Key)

	for k, kr := range r.R {
		for _, mi := range kr.Mirror {
			m[mi] = append(m[mi], k)
		}
	}

	return m
}

/*
	Routes
*/

type Visit func(Route)

type Route interface {
	FindForKeys(keys []Key) *RoutingResult

	FullList() []string

	FullMirrorList(bool) []string

	LinkedTo() []RouteRef

	Name() string

	PostOrder(Visit)

	ToAst() adminparser.AdmNode
}

type RouteRef struct {
	router *Router

	to string
}

func (rf *RouteRef) Route() Route {
	return rf.router.getRoute(rf.to)
}

func (rf *RouteRef) FullList() []string {
	return rf.Route().FullList()
}

func (rf *RouteRef) FullMirrorList(isSelect bool) []string {
	return rf.Route().FullMirrorList(isSelect)
}

type NodeRoute struct {
	Node string
}

func (r *NodeRoute) FindForKeys(keys []Key) *RoutingResult {
	ret := NewRoutingResult()

	for _, k := range keys {
		ret.R[k] = KeyResult{Node: r.Node}
	}

	return ret
}

func (r *NodeRoute) LinkedTo() []RouteRef {
	return nil
}

func (r *NodeRoute) FullList() []string {
	return []string{r.Node}
}

func (r *NodeRoute) FullMirrorList(isSelect bool) []string {
	return []string{}
}

func (r *NodeRoute) Name() string {
	return r.Node
}

func (r *NodeRoute) PostOrder(v Visit) {
	v(r)
}

func (r *NodeRoute) ToAst() adminparser.AdmNode {
	return nil
}

type MirrorRoute struct {
	router *Router
	name   string

	Kind byte

	Main   RouteRef
	Mirror []RouteRef
}

func (r *MirrorRoute) FindForKeys(keys []Key) *RoutingResult {
	ret := r.Main.Route().FindForKeys(keys)

	mr := NewRoutingResult()

	for _, route := range r.Mirror {
		t := route.Route().FindForKeys(keys)

		// on mirror side everything becomes mirrored
		for k, kr := range t.R {
			t.R[k] = KeyResult{
				Node:   "",
				Mirror: append(kr.Mirror, kr.Node),
			}
		}

		mr.Merge(t)
	}

	ret.Merge(mr)

	return ret
}

func (r *MirrorRoute) LinkedTo() []RouteRef {
	return append([]RouteRef{r.Main}, r.Mirror...)
}

func (r *MirrorRoute) FullList() []string {
	return r.Main.FullList()
}

func (r *MirrorRoute) FullMirrorList(isSelect bool) []string {
	l := make([]string, 0)

	if isSelect {
		if r.Kind == MIRROR_R || r.Kind == MIRROR_RW {
			for _, m := range r.Mirror {
				l = append(l, m.FullList()...)
				l = append(l, m.FullMirrorList(isSelect)...)
			}
		}
	} else {
		if r.Kind == MIRROR_W || r.Kind == MIRROR_RW {
			for _, m := range r.Mirror {
				l = append(l, m.FullList()...)
				l = append(l, m.FullMirrorList(isSelect)...)
			}
		}
	}

	return l
}

func (r *MirrorRoute) Name() string {
	return r.name
}

func (r *MirrorRoute) PostOrder(v Visit) {
	r.Main.Route().PostOrder(v)

	for _, m := range r.Mirror {
		m.Route().PostOrder(v)
	}

	v(r)
}

func (r *MirrorRoute) ToAst() adminparser.AdmNode {
	n := &adminparser.MirrorRoute{
		Main: r.Main.to,
	}

	switch r.Kind {
	case MIRROR_R:
		n.Kind = adminparser.MirrorR
	case MIRROR_W:
		n.Kind = adminparser.MirrorW
	case MIRROR_RW:
		n.Kind = adminparser.MirrorRW
	}

	for _, m := range r.Mirror {
		n.Mirrors = append(n.Mirrors, m.to)
	}

	return &adminparser.AddRoute{
		Name:  r.name,
		Route: n,
	}
}

type ModuloHashRoute struct {
	router *Router
	name   string

	N      int
	Routes []RouteRef
}

func (r *ModuloHashRoute) FindForKeys(keys []Key) *RoutingResult {
	buckets := make(map[int][]Key)

	for _, k := range keys {
		n := int(modulo(int64(k), int64(r.N)))
		buckets[n] = append(buckets[n], k)
	}

	ret := NewRoutingResult()

	for n, ks := range buckets {
		ret.Merge(r.Routes[n].Route().FindForKeys(ks))
	}

	return ret
}

func (r *ModuloHashRoute) LinkedTo() []RouteRef {
	return r.Routes
}

func (r *ModuloHashRoute) FullList() (l []string) {
	for _, ro := range r.Routes {
		l = append(l, ro.FullList()...)
	}
	return
}

func (r *ModuloHashRoute) FullMirrorList(isSelect bool) (l []string) {
	for _, ro := range r.Routes {
		l = append(l, ro.FullMirrorList(isSelect)...)
	}
	return
}

func (r *ModuloHashRoute) Name() string {
	return r.name
}

func (r *ModuloHashRoute) PostOrder(v Visit) {
	for _, n := range r.Routes {
		n.Route().PostOrder(v)
	}

	v(r)
}

func (r *ModuloHashRoute) ToAst() adminparser.AdmNode {
	hr := &adminparser.HashRoute{
		Type:   adminparser.ModuloStr,
		Routes: make([]string, len(r.Routes)),
	}
	for i, ro := range r.Routes {
		hr.Routes[i] = ro.to
	}

	return &adminparser.AddRoute{
		Name:  r.name,
		Route: hr,
	}
}

type RangeRoute struct {
	router *Router
	name   string

	Ranges []Range
}

type Range struct {
	KeyRange
	Route RouteRef
}

func (r *RangeRoute) FindForKeys(keys []Key) *RoutingResult {
	buckets := make(map[KeyRange][]Key)

	for _, k := range keys {
		for _, ra := range r.Ranges {
			if ra.Contains(k) {
				buckets[ra.KeyRange] = append(buckets[ra.KeyRange], k)
			}
		}
	}

	ret := NewRoutingResult()

	for _, ra := range r.Ranges {
		ret.Merge(ra.Route.Route().FindForKeys(buckets[ra.KeyRange]))
	}

	return ret
}

func (r *RangeRoute) LinkedTo() []RouteRef {
	rf := make([]RouteRef, len(r.Ranges))

	for i, ra := range r.Ranges {
		rf[i] = ra.Route
	}

	return rf
}

func (r *RangeRoute) FullList() (l []string) {
	for _, ra := range r.Ranges {
		l = append(l, ra.Route.FullList()...)
	}
	return
}

func (r *RangeRoute) FullMirrorList(isSelect bool) (l []string) {
	for _, ra := range r.Ranges {
		l = append(l, ra.Route.FullMirrorList(isSelect)...)
	}
	return
}

func (r *RangeRoute) Name() string {
	return r.name
}

func (r *RangeRoute) PostOrder(v Visit) {
	for _, ra := range r.Ranges {
		ra.Route.Route().PostOrder(v)
	}

	v(r)
}

func (r *RangeRoute) ToAst() adminparser.AdmNode {
	krr := make([]adminparser.KeyRangeRoute, len(r.Ranges))

	for i, ra := range r.Ranges {
		krr[i] = adminparser.KeyRangeRoute{
			Start: key2astkey(ra.Start),
			End:   key2astkey(ra.End),
			Route: ra.Route.Route().Name(),
		}
	}

	return &adminparser.AddRoute{
		Name: r.name,
		Route: &adminparser.RangeRoute{
			Ranges: krr,
		}}
}

type LookupRoute struct {
	router *Router

	node   string
	query  string
	Routes []RouteRef
}

func (r *LookupRoute) FindForKeys(keys []Key) *RoutingResult {
	buckets, err := r.fetchShardKeysMap(keys)
	if err != nil {
		panic(err)
	}

	ret := NewRoutingResult()

	for n, ks := range buckets {
		ret.Merge(r.Routes[n].Route().FindForKeys(ks))
	}

	return ret
}

func (r *LookupRoute) fetchShardKeysMap(keys []Key) (map[int][]Key, error) {
	if keys == nil {
		return nil, errors.New("fetch called for empty key list")
	}

	conn, err := node.GetNode(r.node).GetSelectConn()
	if err != nil {
		panic(err)
	}

	var ks = make([]string, len(keys))

	for i, k := range keys {
		ks[i] = strconv.FormatInt(int64(k), 10)
	}

	csv := strings.Join(ks, ", ")

	sql := fmt.Sprintf(r.query, csv)

	rset, err := conn.Execute(sql)

	if err != nil {
		return nil, fmt.Errorf("failed to get shard_id: %v", err)
	}
	if len(rset.Values) != len(keys) {
		return nil, fmt.Errorf("failed to get shard_id for some keys; expected %s got %s results", len(keys), len(rset.Values))
	}

	var buckets map[int][]Key

	for _, row := range rset.Values {
		key, shard := Key(NumValue(row[0])), int(NumValue(row[1]))
		buckets[shard] = append(buckets[shard], key)
	}

	return buckets, nil
}

func (r *LookupRoute) LinkedTo() []RouteRef {
	return r.Routes
}

func (r *LookupRoute) FullList() (l []string) {
	for _, ro := range r.Routes {
		l = append(l, ro.FullList()...)
	}
	return
}

/*
... *** *** ***
*/

func includeNode(nodes []string, node string) bool {
	for _, n := range nodes {
		if n == node {
			return true
		}
	}
	return false
}

func mergeKeyResult(q, r KeyResult) KeyResult {
	t := KeyResult{}

	t.Node = q.Node

	if q.Node != "" && r.Node != "" {
		panic("can only have one primary route for a key")
	}

	if r.Node != "" {
		t.Node = r.Node
	}

	t.Mirror = append(q.Mirror, r.Mirror...)

	return t
}

func interlist(l1, l2 []Key) (r []Key) {
	has2 := make(map[Key]bool)

	for _, k := range l2 {
		has2[k] = true
	}

	for _, k := range l1 {
		if has2[k] {
			r = append(r, k)
		}
	}

	return r
}

func unionlist(l1, l2 []Key) (r []Key) {
	seen := make(map[Key]bool)

	for _, k := range append(l1, l2...) {
		if ok := seen[k]; !ok {
			r = append(r, k)
		}
		seen[k] = true
	}

	return r
}

func dedupe(in ...string) (out []string) {
	seen := make(map[string]bool)

	for _, s := range in {
		if ok := seen[s]; !ok {
			out = append(out, s)
		}
		seen[s] = true
	}

	return out
}

func modulo(i, n int64) int64 {
	m := i % n
	if m < 0 {
		m = m + n
	}
	return m
}

func routeId(s string) adminparser.RouteID {
	return adminparser.RouteID(s)
}

func key2astkey(k Key) adminparser.RangeNum {
	if k == MaxKey || k == MinKey {
		return adminparser.RangeNum{Inf: true}
	}

	return adminparser.RangeNum{Num: int64(k)}
}

func astkey2key(k adminparser.RangeNum, isStart bool) Key {
	if k.Inf {
		if isStart {
			return MinKey
		} else {
			return MaxKey
		}
	}

	return Key(k.Num)
}
