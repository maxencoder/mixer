package router

import (
	"reflect"
	"testing"

	"github.com/maxencoder/mixer/adminparser"
	"github.com/maxencoder/mixer/config"
	"github.com/maxencoder/mixer/node"
)

var testConfigData = []byte(`
addr : 127.0.0.1:4000
user : root
password : 

nodes :
- 
    name : node1 
    down_after_noalive : 300
    idle_conns : 16
    user: root
    password:
    master : 127.0.0.1:3306
    slaves : 
      - 127.0.0.1:3306
- 
    name : node2
    down_after_noalive : 300
    idle_conns : 16
    user: root
    password:
    master : 127.0.0.1:3307
    slaves : 
      - 127.0.0.1:3307
- 
    name : node3 
    down_after_noalive : 300
    idle_conns : 16
    user: root
    password:
    master : 127.0.0.1:3308
    slaves : 
      - 127.0.0.1:3308
`)

func init() {
	cfg, err := config.ParseConfigData(testConfigData)
	if err != nil {
		panic(err)
	}

	if err := node.ParseNodes(cfg); err != nil {
		panic(err)
	}
}

func TestRouter(t *testing.T) {
	rt, err := NewRouter("mixer", "node1")

	if err != nil {
		t.Fatal(err)
	}

	rt.NewModuloHashRoute("mixer_test_hash-rt", 2, []string{"node2", "node3"})

	trHash, err := rt.NewTableRouter("mixer", "mixer_test_shard_hash", "id", "mixer_test_hash-rt")

	if err != nil {
		t.Fatal(err)
	}

	krr1 := adminparser.KeyRangeRoute{
		Start: adminparser.RangeNum{Inf: true},
		End:   adminparser.RangeNum{Num: 10000},
		Route: "node2",
	}
	krr2 := adminparser.KeyRangeRoute{
		Start: adminparser.RangeNum{Num: 10000},
		End:   adminparser.RangeNum{Inf: true},
		Route: "node3",
	}

	rt.NewRangeRoute("mixer_test_range-rt", &adminparser.RangeRoute{
		Ranges: []adminparser.KeyRangeRoute{krr1, krr2},
	})

	trRange, err := rt.NewTableRouter("mixer", "mixer_test_shard_range", "id", "mixer_test_range-rt")

	if err != nil {
		t.Fatal(err)
	}

	if rt.DefaultNode != "node1" {
		t.Fatal("default rule parse not correct.")
	}

	if !sliceEq(trHash.FullList(), []string{"node2", "node3"}) {
		t.Fatal("hash nodes not correct.")
	}

	if !sliceEq(trRange.FullList(), []string{"node2", "node3"}) {
		t.Fatal("range nodes not correct.")
	}

	r1 := trHash.FindForKeys([]Key{Key(-1), Key(11), Key(12)})

	r1expect := &RoutingResult{
		R: map[Key]KeyResult{
			-1: KeyResult{Node: "node3", Mirror: []string(nil)},
			11: KeyResult{Node: "node3", Mirror: []string(nil)},
			12: KeyResult{Node: "node2", Mirror: []string(nil)},
		}}

	if !reflect.DeepEqual(r1, r1expect) {
		t.Fatal("wrong routing result for hash")
	}

	r2 := trRange.FindForKeys([]Key{Key(-1), Key(11), Key(100000)})

	r2expect := &RoutingResult{
		R: map[Key]KeyResult{
			-1:     KeyResult{Node: "node2", Mirror: []string(nil)},
			11:     KeyResult{Node: "node2", Mirror: []string(nil)},
			100000: KeyResult{Node: "node3", Mirror: []string(nil)},
		}}

	if !reflect.DeepEqual(r2, r2expect) {
		t.Fatal("wrong routing result for range")
	}

	dr := rt.GetTableRouterOrDefault("mixer_default_table")

	if !dr.IsDefault {
		t.Fatal("expected default router")
	}
}

func TestRouter2(t *testing.T) {
	r, err := NewRouter("db1", "node1")

	if err != nil {
		t.Fatal(err)
	}

	r.NewModuloHashRoute("h12", 2, []string{"node1", "node2"})

	m1, err := r.NewMirrorRoute("m1", "r", "node1", []string{"h12", "node3"})

	if err != nil {
		t.Fatal(err)
	}

	if err := r.DeleteRoute("node2"); err == nil {
		t.Fatal("should not be possible: 'node2' is used")
	}

	if !sliceEq(m1.FullList(), []string{"node1"}) {
		t.Fatal("fulllist not correct.")
	}

	if !sliceEq(m1.FullMirrorList(true), []string{"node1", "node2", "node3"}) {
		t.Fatal("fulllist not correct.")
	}

	if !sliceEq(m1.FullMirrorList(false), []string{}) {
		t.Fatal("fulllist not correct.")
	}

	r1 := m1.FindForKeys([]Key{Key(-2), Key(-1), Key(0)})

	r1expect := &RoutingResult{
		R: map[Key]KeyResult{
			-2: KeyResult{Node: "node1", Mirror: []string{"node1", "node3"}},
			-1: KeyResult{Node: "node1", Mirror: []string{"node2", "node3"}},
			0:  KeyResult{Node: "node1", Mirror: []string{"node1", "node3"}},
		}}

	if !reflect.DeepEqual(r1, r1expect) {
		t.Fatal("wrong routing result for hash")
	}
}

func sliceEq(a, b []string) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
