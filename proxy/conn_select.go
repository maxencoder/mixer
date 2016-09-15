package proxy

import (
	"fmt"
	"github.com/maxencoder/mixer/sqlparser"
	. "github.com/siddontang/go-mysql/mysql"
	"strings"
)

func (c *Conn) handleSimpleSelect(sql string, stmt *sqlparser.SimpleSelect) (*Result, error) {
	if len(stmt.SelectExprs) != 1 {
		return nil, fmt.Errorf("support select one informaction function, %s", sql)
	}

	expr, ok := stmt.SelectExprs[0].(*sqlparser.NonStarExpr)
	if !ok {
		return nil, fmt.Errorf("support select informaction function, %s", sql)
	}

	var f *sqlparser.FuncExpr
	f, ok = expr.Expr.(*sqlparser.FuncExpr)
	if !ok {
		return nil, fmt.Errorf("support select informaction function, %s", sql)
	}

	var r *Resultset
	var err error

	switch strings.ToLower(string(f.Name)) {
	case "last_insert_id":
		r, err = c.buildSimpleSelectResult(c.lastInsertId, f.Name, expr.As)
	case "row_count":
		r, err = c.buildSimpleSelectResult(c.affectedRows, f.Name, expr.As)
	case "version":
		r, err = c.buildSimpleSelectResult(ServerVersion, f.Name, expr.As)
	case "connection_id":
		r, err = c.buildSimpleSelectResult(c.connectionId, f.Name, expr.As)
	case "database":
		if c.schema != nil {
			r, err = c.buildSimpleSelectResult(c.schema.db, f.Name, expr.As)
		} else {
			r, err = c.buildSimpleSelectResult("NULL", f.Name, expr.As)
		}
	default:
		return nil, fmt.Errorf("function %s not support", f.Name)
	}

	if err != nil {
		return nil, err
	}

	return &Result{Status: c.status, Resultset: r}, nil
}

func (c *Conn) buildSimpleSelectResult(value interface{}, name []byte, asName []byte) (*Resultset, error) {
	field := &Field{}

	field.Name = name

	if asName != nil {
		field.Name = asName
	}

	field.OrgName = name

	formatField(field, value)

	r := &Resultset{Fields: []*Field{field}}
	row, err := formatValue(value)
	if err != nil {
		return nil, err
	}
	r.RowDatas = append(r.RowDatas, PutLengthEncodedString(row))

	return r, nil
}
