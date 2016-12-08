package router

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/maxencoder/mixer/node"
)

const (
	MIRROR_RO = iota
	MIRROR_RW
	MIRROR_ALL
)

/*
	Router
		|-TableRouter
			|-Route
				|-Route...
*/

type Router struct {
	sync.Mutex

	DB string

	tr map[string]*TableRouter // keyed by table in lower case

	rt map[string]Route

	DefaultNode        string
	DefaultTableRouter *TableRouter
}

func NewRouter(db string, defaultNode string) *Router {
	r := &Router{DB: db}

	r.rt = make(map[string]Route)
	r.tr = make(map[string]*TableRouter)

	_, _ = r.NewNodeRoute(defaultNode)
	r.DefaultNode = defaultNode
	r.DefaultTableRouter = r.NewDefaultTableRouter()

	return r
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
	r.Lock()
	defer r.Unlock()

	ro, ok := r.rt[name]
	if !ok {
		panic(fmt.Errorf("route '%s' does not exist", name))
	}
	return ro
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

func (r *Router) GetTableRouterOrDefault(name string) *TableRouter {
	r.Lock()
	defer r.Unlock()

	tr, ok := r.tr[name]
	if !ok {
		return r.DefaultTableRouter
	}
	return tr
}

func (r *Router) GetTableRouter(name string) (*TableRouter, error) {
	r.Lock()
	defer r.Unlock()

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

func (r *Router) NewNodeRoute(node string) (Route, error) {
	n := &NodeRoute{router: r, Node: node}

	// TODO: check node exists

	if ex, _ := r.GetRoute(node); ex != nil {
		return nil, fmt.Errorf("route '%s' already exists", node)
	}

	r.SetRoute(node, n)

	return n, nil
}

func (r *Router) routeRef(to string) (RouteRef, error) {
	if _, err := r.GetRoute(to); err != nil {
		return RouteRef{}, err
	}

	return RouteRef{router: r, to: to}, nil
}

func (r *Router) NewModuloHashRoute(name string, N int, routes []string) (Route, error) {
	n := &ModuloHashRoute{router: r, N: N}

	for _, ro := range routes {
		rf, err := r.routeRef(ro)

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

func (r *Router) NewMirrorRoute(name string, main string, mirror []string) (Route, error) {
	mrf, err := r.routeRef(main)

	if err != nil {
		return nil, err
	}

	n := &MirrorRoute{router: r, Main: mrf}

	for _, ro := range mirror {
		rf, err := r.routeRef(ro)

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

type Route interface {
	FindForKeys(keys []Key) *RoutingResult

	FullList() []string

	FullMirrorList(bool) []string

	LinkedTo() []RouteRef
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
	router *Router

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

type MirrorRoute struct {
	router *Router

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
		for _, kr := range t.R {
			kr.Mirror = append(kr.Mirror, kr.Node)
			kr.Node = ""
		}

		mr.Merge(t)
	}

	ret.Merge(mr)

	return ret
}

func (r *MirrorRoute) LinkedTo() []RouteRef {
	return append(r.Mirror, r.Main)
}

func (r *MirrorRoute) FullList() []string {
	return r.Main.FullList()
}

func (r *MirrorRoute) FullMirrorList(isSelect bool) []string {
	l := make([]string, 0)

	if isSelect {
		if r.Kind == MIRROR_RO || r.Kind == MIRROR_ALL {
			for _, m := range r.Mirror {
				l = append(l, m.FullList()...)
				l = append(l, m.FullMirrorList(isSelect)...)
			}
		}
	} else {
		if r.Kind == MIRROR_RW || r.Kind == MIRROR_ALL {
			for _, m := range r.Mirror {
				l = append(l, m.FullList()...)
				l = append(l, m.FullMirrorList(isSelect)...)
			}
		}
	}

	return l
}

type ModuloHashRoute struct {
	router *Router

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

	var ks []string
	for _, k := range keys {
		ks = append(ks, strconv.FormatInt(int64(k), 10))
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
