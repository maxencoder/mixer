package sqlparser

import (
	"fmt"
	"sort"
	"testing"

	"github.com/maxencoder/mixer/adminparser"
	"github.com/maxencoder/mixer/node"
	"github.com/maxencoder/mixer/router"
)

func init() {
	node.InitPool()

	for i := 0; i < 10; i++ {
		name := fmt.Sprintf("node%d", i)
		n, err := node.NewNode(
			name,
			"root",
			"",
			4,
			300,
			"127.0.0.1:3306",
			[]string{"127.0.0.1:3307"},
		)

		if err != nil {
			panic(err)
		}

		node.SetNode(name, n)
	}
}

func newTestDbRouter(t *testing.T) *router.Router {
	rt, err := router.NewRouter("db1", "node0")

	if err != nil {
		t.Fatal(err)
	}

	rt.NewModuloHashRoute("test1-rt", 10, []string{"node0", "node1", "node2", "node3", "node4", "node5", "node6", "node7", "node8", "node9"})

	_, err = rt.NewTableRouter("db1", "test1", "id", "test1-rt")

	if err != nil {
		t.Fatal(err)
	}

	// ranges: -10000-20000-
	krr1 := adminparser.KeyRangeRoute{
		Start: adminparser.RangeNum{Inf: true},
		End:   adminparser.RangeNum{Num: 10000},
		Route: "node0",
	}
	krr2 := adminparser.KeyRangeRoute{
		Start: adminparser.RangeNum{Num: 10000},
		End:   adminparser.RangeNum{Num: 20000},
		Route: "node1",
	}
	krr3 := adminparser.KeyRangeRoute{
		Start: adminparser.RangeNum{Num: 20000},
		End:   adminparser.RangeNum{Inf: true},
		Route: "node2",
	}

	rt.NewRangeRoute("test2-rt", &adminparser.RangeRoute{
		Ranges: []adminparser.KeyRangeRoute{krr1, krr2, krr3},
	})

	_, err = rt.NewTableRouter("db1", "test2", "id", "test2-rt")

	if err != nil {
		t.Fatal(err)
	}

	return rt
}

func checkSharding(t *testing.T, sql string, args []int, checkNodeIndex ...int) {
	r := newTestDbRouter(t)

	bindVars := make(map[string]interface{}, len(args))
	for i, v := range args {
		bindVars[fmt.Sprintf("v%d", i+1)] = v
	}

	stmt, err := Parse(sql)
	if err != nil {
		t.Fatal(err)
	}

	eps, err := RouteStmt(stmt, sql, r, bindVars)
	if err != nil {
		t.Fatal(sql, err)
	}

	var got, expect []string

	for _, i := range checkNodeIndex {
		expect = append(expect, fmt.Sprintf("node%d", i))
	}

	for _, ep := range eps {
		got = append(got, ep.Node)
	}

	sort.Strings(got)

	if !sliceEq(got, expect) {
		s := fmt.Sprintf("%v %v", got, checkNodeIndex)
		t.Fatal(sql, s)
	}
}

func TestConditionSharding(t *testing.T) {
	var sql string

	sql = "select * from test1 where id = 5"
	checkSharding(t, sql, nil, 5)

	sql = "select * from test1 where id != 5"
	checkSharding(t, sql, nil, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)

	sql = "select * from test1 where id in (5, 6)"
	checkSharding(t, sql, nil, 5, 6)

	sql = "select * from test1 where id > 5"
	checkSharding(t, sql, nil, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)

	sql = "select * from test1 where id in (5, 6) and id in (5, 6, 7)"
	checkSharding(t, sql, nil, 5, 6)

	sql = "select * from test1 where id in (5, 6) or id in (5, 6, 7,8)"
	checkSharding(t, sql, nil, 5, 6, 7, 8)

	sql = "select * from test1 where id not in (5, 6) or id in (5, 6, 7,8)"
	checkSharding(t, sql, nil, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)

	sql = "select * from test1 where id not in (5, 6)"
	checkSharding(t, sql, nil, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)

	sql = "select * from test1 where id in (5, 6) or (id in (5, 6, 7,8) and id in (1,5,7))"
	checkSharding(t, sql, nil, 5, 6, 7)

	sql = "select * from test2 where id = 10000"
	checkSharding(t, sql, nil, 1)

	sql = "select * from test2 where id between 10000 and 100000"
	checkSharding(t, sql, nil, 0, 1, 2)

	sql = "select * from test2 where id not between 1000 and 100000"
	checkSharding(t, sql, nil, 0, 1, 2)

	sql = "select * from test2 where id not between 10000 and 100000"
	checkSharding(t, sql, nil, 0, 1, 2)

	sql = "select * from test2 where id > 10000"
	checkSharding(t, sql, nil, 0, 1, 2)

	sql = "select * from test2 where id >= 10000"
	checkSharding(t, sql, nil, 0, 1, 2)

	sql = "select * from test2 where id <= 10000"
	checkSharding(t, sql, nil, 0, 1, 2)

	sql = "select * from test2 where id < 10000"
	checkSharding(t, sql, nil, 0, 1, 2)

	sql = "select * from test2 where  10000 < id"
	checkSharding(t, sql, nil, 0, 1, 2)

	sql = "select * from test2 where  10000 <= id"
	checkSharding(t, sql, nil, 0, 1, 2)

	sql = "select * from test2 where  10000 > id"
	checkSharding(t, sql, nil, 0, 1, 2)

	sql = "select * from test2 where  10000 >= id"
	checkSharding(t, sql, nil, 0, 1, 2)

	sql = "select * from test2 where id >= 10000 and id <= 100000"
	checkSharding(t, sql, nil, 0, 1, 2)

	sql = "select * from test2 where (id >= 10000 and id <= 100000) or id < 100"
	checkSharding(t, sql, nil, 0, 1, 2)

	sql = "select * from test2 where (id >= 10000 and id <= 100000) or (id < 100 and name > 100000)"
	checkSharding(t, sql, nil, 0, 1, 2)

	sql = "select * from test2 where id in (1, 10000)"
	checkSharding(t, sql, nil, 0, 1)

	sql = "select * from test2 where id in (-1, 1, 2, 3)"
	checkSharding(t, sql, nil, 0)

	sql = "select * from test2 where id not in (1, 10000)"
	checkSharding(t, sql, nil, 0, 1, 2)

	sql = "select * from test2 where id in (1000, 10000)"
	checkSharding(t, sql, nil, 0, 1)

	sql = "select * from test2 where id > -1"
	checkSharding(t, sql, nil, 0, 1, 2)

	sql = "select * from test2 where id > -1 and id < 11000"
	checkSharding(t, sql, nil, 0, 1, 2)
}

func TestConditionVarArgSharding(t *testing.T) {
	var sql string

	sql = "select * from test1 where id = ?"
	checkSharding(t, sql, []int{5}, 5)

	sql = "select * from test1 where id in (?, ?)"
	checkSharding(t, sql, []int{5, 6}, 5, 6)

	sql = "select * from test1 where id > ?"
	checkSharding(t, sql, []int{5}, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)

	sql = "select * from test1 where id in (?, ?) and id in (?, ?, ?)"
	checkSharding(t, sql, []int{5, 6, 5, 6, 7}, 5, 6)

	sql = "select * from test1 where id in (?, ?) or id in (?, ?,?,?)"
	checkSharding(t, sql, []int{5, 6, 5, 6, 7, 8}, 5, 6, 7, 8)

	sql = "select * from test1 where id not in (?, ?) or id in (?, ?, ?,?)"
	checkSharding(t, sql, []int{5, 6, 5, 6, 7, 8}, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)

	sql = "select * from test1 where id not in (?, ?)"
	checkSharding(t, sql, []int{5, 6}, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)

	sql = "select * from test1 where id in (?, ?) or (id in (?, ?, ?,?) and id in (?,?,?))"
	checkSharding(t, sql, []int{5, 6, 5, 6, 7, 8, 1, 5, 7}, 5, 6, 7)

	sql = "select * from test2 where id = ?"
	checkSharding(t, sql, []int{10000}, 1)

	sql = "select * from test2 where id between ? and ?"
	checkSharding(t, sql, []int{10000, 100000}, 0, 1, 2)

	sql = "select * from test2 where id not between ? and ?"
	checkSharding(t, sql, []int{10000, 100000}, 0, 1, 2)

	sql = "select * from test2 where id not between ? and ?"
	checkSharding(t, sql, []int{10000, 100000}, 0, 1, 2)

	sql = "select * from test2 where id > ?"
	checkSharding(t, sql, []int{10000}, 0, 1, 2)

	sql = "select * from test2 where id >= ?"
	checkSharding(t, sql, []int{10000}, 0, 1, 2)

	sql = "select * from test2 where id <= ?"
	checkSharding(t, sql, []int{10000}, 0, 1, 2)

	sql = "select * from test2 where id < ?"
	checkSharding(t, sql, []int{10000}, 0, 1, 2)

	sql = "select * from test2 where  ? < id"
	checkSharding(t, sql, []int{10000}, 0, 1, 2)

	sql = "select * from test2 where  ? <= id"
	checkSharding(t, sql, []int{10000}, 0, 1, 2)

	sql = "select * from test2 where  ? > id"
	checkSharding(t, sql, []int{10000}, 0, 1, 2)

	sql = "select * from test2 where  ? >= id"
	checkSharding(t, sql, []int{10000}, 0, 1, 2)

	sql = "select * from test2 where id >= ? and id <= ?"
	checkSharding(t, sql, []int{10000, 100000}, 0, 1, 2)

	sql = "select * from test2 where (id >= ? and id <= ?) or id < ?"
	checkSharding(t, sql, []int{10000, 100000, 100}, 0, 1, 2)

	sql = "select * from test2 where (id >= ? and id <= ?) or (id < ? and name > ?)"
	checkSharding(t, sql, []int{10000, 100000, 100, 100000}, 0, 1, 2)

	sql = "select * from test2 where id in (?, ?)"
	checkSharding(t, sql, []int{1, 10000}, 0, 1)

	sql = "select * from test2 where id not in (?, ?)"
	checkSharding(t, sql, []int{1, 10000}, 0, 1, 2)

	sql = "select * from test2 where id in (?, ?)"
	checkSharding(t, sql, []int{1000, 10000}, 0, 1)

	sql = "select * from test2 where id > ?"
	checkSharding(t, sql, []int{-1}, 0, 1, 2)

	sql = "select * from test2 where id > ? and id < ?"
	checkSharding(t, sql, []int{-1, 11000}, 0, 1, 2)
}

func TestValueSharding(t *testing.T) {
	var sql string

	sql = "insert into test1 (id) values (5)"
	checkSharding(t, sql, nil, 5)

	sql = "insert into test2 (id) values (10000)"
	checkSharding(t, sql, nil, 1)

	sql = "insert into test2 (id) values (20000)"
	checkSharding(t, sql, nil, 2)

	sql = "insert into test2 (id) values (200000)"
	checkSharding(t, sql, nil, 2)
}

func TestValueVarArgSharding(t *testing.T) {
	var sql string

	sql = "insert into test1 (id) values (?)"
	checkSharding(t, sql, []int{5}, 5)

	sql = "insert into test2 (id) values (?)"
	checkSharding(t, sql, []int{10000}, 1)

	sql = "insert into test2 (id) values (?)"
	checkSharding(t, sql, []int{20000}, 2)

	sql = "insert into test2 (id) values (?)"
	checkSharding(t, sql, []int{200000}, 2)
}

func TestBadUpdateExpr(t *testing.T) {
	var sql string

	r := newTestDbRouter(t)

	sql = "insert into test1 (id) values (5) on duplicate key update id = 10"

	if _, err := routeSql(sql, r, nil); err == nil {
		t.Fatal("must err")
	}

	sql = "update test1 set id = 10 where id = 5"

	if _, err := routeSql(sql, r, nil); err == nil {
		t.Fatal("must err")
	}

	sql = "insert into test1 (id, k) values (5, 55), (6, 66)"

	if _, err := routeSql(sql, r, nil); err == nil {
		t.Fatal("must err")
	}

	sql = "insert into test1 (id, k) select * from atable"

	if _, err := routeSql(sql, r, nil); err == nil {
		t.Fatal("must err")
	}
}

func TestFilterStmt(t *testing.T) {
	var sql, expect string
	var plan *AnalysisPlan
	var ks []router.Key

	sql = "insert into A (k, v) values (1, 2), (2, 3)"
	expect = "insert into A(k, v) values (1, 2)"
	ks = []router.Key{router.Key(1)}
	plan = &AnalysisPlan{insertKeyPos: 0}

	checkFilterStmt(t, sql, expect, ks, plan)

	sql = `insert into A (k, v) values ("1", 2), (2, 3)`
	expect = "insert into A(k, v) values ('1', 2)"
	ks = []router.Key{router.Key(5), router.Key(1)}

	checkFilterStmt(t, sql, expect, ks, plan)

	sql = `insert into A (k, v) values ("1", 2, 3.0), (2, 3, 4.3)`
	expect = "insert into A(k, v) values ('1', 2, 3.0), (2, 3, 4.3)"
	ks = []router.Key{router.Key(2), router.Key(1)}

	checkFilterStmt(t, sql, expect, ks, plan)
}

func checkFilterStmt(t *testing.T, sql, expect string, keys []router.Key, plan *AnalysisPlan) {
	stmt, err := Parse(sql)
	if err != nil {
		t.Fatal(err)
	}

	filtered, err := FilterStmt(stmt, keys, plan)
	if err != nil {
		t.Fatal(err)
	}

	if String(filtered) != expect {
		t.Fatalf("expected: %s, got: %s\n", expect, String(filtered))
	}
}

func routeSql(sql string, r *router.Router, bindVars map[string]interface{}) ([]*ExecPlan, error) {
	stmt, err := Parse(sql)

	if err != nil {
		return nil, err
	}

	eps, err := RouteStmt(stmt, sql, r, bindVars)

	if err != nil {
		return nil, err
	}

	return eps, nil
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
