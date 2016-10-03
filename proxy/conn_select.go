package proxy

import (
	"fmt"
	"strings"

	"github.com/maxencoder/mixer/hack"
	"github.com/maxencoder/mixer/sqlparser"
	. "github.com/siddontang/go-mysql/mysql"
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

	switch strings.ToLower(f.Name) {
	case "last_insert_id":
		r, err = c.buildSimpleSelectResult(c.lastInsertId, f.Name, expr.As.String())
	case "row_count":
		r, err = c.buildSimpleSelectResult(c.affectedRows, f.Name, expr.As.String())
	case "version":
		r, err = c.buildSimpleSelectResult(ServerVersion, f.Name, expr.As.String())
	case "connection_id":
		r, err = c.buildSimpleSelectResult(c.connectionId, f.Name, expr.As.String())
	case "database":
		if c.schema != nil {
			r, err = c.buildSimpleSelectResult(c.schema.db, f.Name, expr.As.String())
		} else {
			r, err = c.buildSimpleSelectResult("NULL", f.Name, expr.As.String())
		}
	default:
		return nil, fmt.Errorf("function %s not support", f.Name)
	}

	if err != nil {
		return nil, err
	}

	return &Result{Status: c.status, Resultset: r}, nil
}

func (c *Conn) buildSimpleSelectResult(value interface{}, name string, asName string) (*Resultset, error) {
	field := &Field{}

	field.Name = hack.Slice(name)

	if asName != "" {
		field.Name = hack.Slice(asName)
	}

	field.OrgName = hack.Slice(name)

	formatField(field, value)

	r := &Resultset{Fields: []*Field{field}}
	row, err := formatValue(value)
	if err != nil {
		return nil, err
	}
	r.RowDatas = append(r.RowDatas, PutLengthEncodedString(row))

	return r, nil
}
