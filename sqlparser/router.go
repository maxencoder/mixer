// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sqlparser

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/maxencoder/mixer/db"
	"github.com/maxencoder/mixer/router"
)

const (
	EID_NODE = iota
	VALUE_NODE
	LIST_NODE
	OTHER_NODE
)

type ExecPlan struct {
	sql string

	Stmt Statement

	Node string

	IsMirror bool

	Conn *db.SqlConn
}

func (p *ExecPlan) Sql() string {
	if p.sql != "" {
		return p.sql
	}
	return String(p.Stmt)
}

type AnalysisPlan struct {
	router *router.TableRouter

	table string

	isInsert bool
	isSelect bool

	criteria SQLNode

	bindVars map[string]interface{}

	insertKeyPos int
}

/*
	Limitation:

	where, eg, key name is id:

		where id = 1
		where id in (1, 2, 3)

	todo:
		where id > 1
		where id >= 1
		where id < 1
		where id <= 1
		where id between 1 and 10
		where id >= 1 and id < 10
*/

func RouteStmt(stmt Statement, sql string, r *router.Router, bindVars map[string]interface{}) (execPlans []*ExecPlan, err error) {
	defer handleError(&err)

	plan := getAnalysisPlan(stmt)

	plan.bindVars = bindVars

	tr, err := r.GetTableRouter(plan.table)
	if err != nil {
		return nil, err
	}

	// TODO: optimize routing when there's only one shard node (default)

	ke := plan.keyExprFromPlan()

	switch ke := ke.(type) {
	case *router.KeyList:
		keyRoutes := tr.FindForKeys(ke.Keys)

		for node, keys := range keyRoutes.SplitByNode() {
			newStmt := stmt
			newSql := sql

			if plan.isInsert {
				newStmt, err = FilterStmt(stmt, keys, plan)
				if err != nil {
					return nil, err
				}
				// insert sql is transformed
				newSql = String(newStmt)
			}

			execPlans = append(execPlans,
				&ExecPlan{newSql, newStmt, node, false, nil})
		}

		for node, keys := range keyRoutes.SplitByMirrorNode() {
			newStmt := stmt
			newSql := sql

			if plan.isInsert {
				newStmt, err = FilterStmt(stmt, keys, plan)
				if err != nil {
					return nil, err
				}
				// insert sql is transformed
				newSql = String(newStmt)
			}

			execPlans = append(execPlans,
				&ExecPlan{newSql, newStmt, node, true, nil})
		}
	case *router.KeyUnknown:
		if plan.isInsert {
			panic(errors.New("unable to route unknown keys for Insert"))
		}
		for _, node := range tr.FullList() {
			execPlans = append(execPlans, &ExecPlan{sql, stmt, node, false, nil})
		}
		for _, node := range tr.FullMirrorList(plan.isSelect) {
			execPlans = append(execPlans, &ExecPlan{sql, stmt, node, true, nil})
		}
	default:
		panic(fmt.Errorf("unsupported KeyExpr type: %T", ke))
	}

	if err := plan.checkUpdateExprs(stmt, execPlans); err != nil {
		return nil, err
	}

	return execPlans, nil
}

func FilterStmt(stmt Statement, keys []router.Key, plan *AnalysisPlan) (result Statement, err error) {
	defer handleError(&err)

	insert, ok := stmt.(*Insert)
	if !ok {
		return stmt, nil
	}

	new := &Insert{
		Comments: insert.Comments,
		Ignore:   insert.Ignore,
		Table:    insert.Table,
		Columns:  insert.Columns,
		OnDup:    insert.OnDup,
	}

	var filtered Values
	vals := insert.Rows.(Values)

	for i := 0; i < len(vals); i++ {
		tuple := vals[i].(ValTuple)
		val := tuple[plan.insertKeyPos]

		if contains(keys, plan.getKey(val)) {
			filtered = append(filtered, tuple)
		}
	}

	new.Rows = filtered

	return new, nil
}

func (plan *AnalysisPlan) getKeyExpr(expr BoolExpr) router.KeyExpr {
	switch criteria := expr.(type) {
	case *ComparisonExpr:
		switch criteria.Operator {
		case "=":
			var k router.Key
			if plan.routingAnalyzeValue(criteria.Left) == EID_NODE {
				k = plan.getKey(criteria.Right)
			} else {
				k = plan.getKey(criteria.Left)
			}
			return &router.KeyList{Keys: []router.Key{k}}
		case "in":
			return plan.getKeyList(criteria.Right)
		case "<=>", "!=":
			return &router.KeyUnknown{}
		case "<", ">", "<=", ">=":
			return &router.KeyUnknown{}
		case "not in":
			return &router.KeyUnknown{}
		}
	case *RangeCond:
		return &router.KeyUnknown{}
	default:
		return &router.KeyUnknown{}
	}

	return &router.KeyUnknown{}
}

func (plan *AnalysisPlan) keyExprFromPlan() router.KeyExpr {
	if plan.criteria == nil {
		return &router.KeyUnknown{}
	}

	switch criteria := plan.criteria.(type) {
	case Values:
		return plan.getInsertKeyExpr(criteria)
	case BoolExpr:
		return plan.getKeyExprFromBoolExpr(criteria)
	default:
		return &router.KeyUnknown{}
	}
}

func (plan *AnalysisPlan) checkUpdateExprs(statement Statement, plans []*ExecPlan) error {
	if plan.router.IsDefault {
		return nil
	}
	if len(plans) == 1 {
		return nil
	}

	switch stmt := statement.(type) {
	case *Insert:
		if stmt.OnDup != nil {
			return plan.checkNotUpdatingKey(UpdateExprs(stmt.OnDup))
		}
	case *Update:
		return plan.checkNotUpdatingKey(stmt.Exprs)
	default:
		return nil
	}
	return nil
}

func (plan *AnalysisPlan) checkNotUpdatingKey(exprs UpdateExprs) error {
	for _, e := range exprs {
		if e.Name.Lowered() == plan.router.Key {
			return errors.New("routing key can not be in update expression")
		}
	}
	return nil
}

func getAnalysisPlan(statement Statement) (plan *AnalysisPlan) {
	plan = &AnalysisPlan{}
	var where *Where
	var whereRequired bool = true

	switch stmt := statement.(type) {
	case *Insert:
		if _, ok := stmt.Rows.(SelectStatement); ok {
			panic(errors.New("select in insert not allowed"))
		}

		plan.table = String(stmt.Table)
		plan.isInsert = true
		plan.criteria = plan.routingAnalyzeValues(stmt)

		return plan
	case *Select:
		plan.table = String(stmt.From[0])
		plan.isSelect = true
		where = stmt.Where
		whereRequired = false
	case *Update:
		plan.table = String(stmt.Table)
		where = stmt.Where
	case *Delete:
		plan.table = String(stmt.Table)
		where = stmt.Where
	}

	if where != nil {
		plan.criteria = where.Expr
	} else if whereRequired {
		panic(errors.New("statements with empty `where` clause not supported"))
	}

	return plan
}

func (plan *AnalysisPlan) routingAnalyzeValues(stmt *Insert) Values {
	var keyPos int = -1

	for i, c := range stmt.Columns {
		if c.Lowered() == plan.router.Key {
			keyPos = i
			break
		}
	}
	if keyPos == -1 {
		panic("failed to find sharding key in insert")
	}

	plan.insertKeyPos = keyPos

	vals := stmt.Rows.(Values)

	for i := 0; i < len(vals); i++ {
		switch tuple := vals[i].(type) {
		case ValTuple:
			result := plan.routingAnalyzeValue(tuple[keyPos])
			if result != VALUE_NODE {
				panic(errors.New("insert is too complex"))
			}
		default:
			// subquery
			panic(errors.New("insert is too complex"))
		}
	}
	return vals
}

func (plan *AnalysisPlan) getKeyExprFromBoolExpr(node BoolExpr) router.KeyExpr {
	switch node := node.(type) {
	case *AndExpr:
		left := plan.getKeyExprFromBoolExpr(node.Left)
		right := plan.getKeyExprFromBoolExpr(node.Right)

		return router.KeyExprAnd(left, right)
	case *OrExpr:
		left := plan.getKeyExprFromBoolExpr(node.Left)
		right := plan.getKeyExprFromBoolExpr(node.Right)
		return router.KeyExprOr(left, right)
	case *ParenBoolExpr:
		return plan.getKeyExprFromBoolExpr(node.Expr)
	case *ComparisonExpr:
		switch {
		case StringIn(node.Operator, "=", "<", ">", "<=", ">=", "<=>"):
			left := plan.routingAnalyzeValue(node.Left)
			right := plan.routingAnalyzeValue(node.Right)
			if (left == EID_NODE && right == VALUE_NODE) || (left == VALUE_NODE && right == EID_NODE) {
				return plan.getKeyExpr(node)
			}
		case StringIn(node.Operator, "in", "not in"):
			left := plan.routingAnalyzeValue(node.Left)
			right := plan.routingAnalyzeValue(node.Right)
			if left == EID_NODE && right == LIST_NODE {
				return plan.getKeyExpr(node)
			}
		}
	case *RangeCond:
		left := plan.routingAnalyzeValue(node.Left)
		from := plan.routingAnalyzeValue(node.From)
		to := plan.routingAnalyzeValue(node.To)
		if left == EID_NODE && from == VALUE_NODE && to == VALUE_NODE {
			return plan.getKeyExpr(node)
		}
	}
	return &router.KeyUnknown{}
}

func (plan *AnalysisPlan) routingAnalyzeValue(valExpr ValExpr) int {
	switch node := valExpr.(type) {
	case *ColName:
		if node.Name.Lowered() == plan.router.Key {
			return EID_NODE
		}
	case ValTuple:
		for _, n := range node {
			if plan.routingAnalyzeValue(n) != VALUE_NODE {
				return OTHER_NODE
			}
		}
		return LIST_NODE
	case StrVal, NumVal, ValArg:
		return VALUE_NODE
	}
	return OTHER_NODE
}

func (plan *AnalysisPlan) getKeyList(valExpr ValExpr) router.KeyExpr {
	l := &router.KeyList{Keys: []router.Key{}}

	switch node := valExpr.(type) {
	case ValTuple:
		for _, n := range node {
			l.Keys = append(l.Keys, plan.getKey(n))
		}
	}

	return l
}

func (plan *AnalysisPlan) getInsertKeyExpr(vals Values) router.KeyExpr {
	ks := make([]router.Key, 0)

	for i := 0; i < len(vals); i++ {
		value_expression := vals[i].(ValTuple)[plan.insertKeyPos]
		ks = append(ks, plan.getKey(value_expression))
	}

	return &router.KeyList{Keys: ks}
}

func (plan *AnalysisPlan) getKey(valExpr ValExpr) router.Key {
	value := plan.getBoundValue(valExpr)
	return router.NewKey(value)
}

func (plan *AnalysisPlan) getBoundValue(valExpr ValExpr) interface{} {
	switch node := valExpr.(type) {
	case ValTuple:
		if len(node) != 1 {
			panic(errors.New("tuples not allowed as insert values"))
		}
		// TODO: Change parser to create single value tuples into non-tuples.
		return plan.getBoundValue(node[0])
	case StrVal:
		return string(node)
	case NumVal:
		val, err := strconv.ParseInt(string(node), 10, 64)
		if err != nil {
			panic(fmt.Errorf("%s", err.Error()))
		}
		return val
	case ValArg:
		return plan.bindVars[string(node[1:])]
	}
	panic("Unexpected token")
}

func makeList(start, end int) []int {
	list := make([]int, end-start)
	for i := start; i < end; i++ {
		list[i-start] = i
	}
	return list
}

// l1 & l2
func interList(l1 []int, l2 []int) []int {
	if len(l1) == 0 || len(l2) == 0 {
		return []int{}
	}

	l3 := make([]int, 0, len(l1)+len(l2))
	var i = 0
	var j = 0
	for i < len(l1) && j < len(l2) {
		if l1[i] == l2[j] {
			l3 = append(l3, l1[i])
			i++
			j++
		} else if l1[i] < l2[j] {
			i++
		} else {
			j++
		}
	}

	return l3
}

// l1 | l2
func unionList(l1 []int, l2 []int) []int {
	if len(l1) == 0 {
		return l2
	} else if len(l2) == 0 {
		return l1
	}

	l3 := make([]int, 0, len(l1)+len(l2))

	var i = 0
	var j = 0
	for i < len(l1) && j < len(l2) {
		if l1[i] < l2[j] {
			l3 = append(l3, l1[i])
			i++
		} else if l1[i] > l2[j] {
			l3 = append(l3, l2[j])
			j++
		} else {
			l3 = append(l3, l1[i])
			i++
			j++
		}
	}

	if i != len(l1) {
		l3 = append(l3, l1[i:]...)
	} else if j != len(l2) {
		l3 = append(l3, l2[j:]...)
	}

	return l3
}

// l1 - l2
func differentList(l1 []int, l2 []int) []int {
	if len(l1) == 0 {
		return []int{}
	} else if len(l2) == 0 {
		return l1
	}

	l3 := make([]int, 0, len(l1))

	var i = 0
	var j = 0
	for i < len(l1) && j < len(l2) {
		if l1[i] < l2[j] {
			l3 = append(l3, l1[i])
			i++
		} else if l1[i] > l2[j] {
			j++
		} else {
			i++
			j++
		}
	}

	if i != len(l1) {
		l3 = append(l3, l1[i:]...)
	}

	return l3
}

func contains(keys []router.Key, key router.Key) bool {
	for _, k := range keys {
		if k == key {
			return true
		}
	}
	return false
}

func handleError(err *error) {
	if x := recover(); x != nil {
		*err = x.(error)
	}
}
