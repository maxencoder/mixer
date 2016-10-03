package sqlparser

import "fmt"

// -- analyzer.go

// GetDBName parses the specified DML and returns the
// db name if it was used to qualify the table name.
// It returns an error if parsing fails or if the statement
// is not a DML.
func GetDBName(sql string) (string, error) {
	statement, err := Parse(sql)
	if err != nil {
		return "", err
	}
	switch stmt := statement.(type) {
	case *Insert:
		return string(stmt.Table.Qualifier), nil
	case *Update:
		return string(stmt.Table.Qualifier), nil
	case *Delete:
		return string(stmt.Table.Qualifier), nil
	}
	return "", fmt.Errorf("statement '%s' is not a dml", sql)
}

// -- ast.go

func (*Begin) iStatement()    {}
func (*Commit) iStatement()   {}
func (*Rollback) iStatement() {}

type Begin struct {
}

func (node *Begin) Format(buf *TrackedBuffer) {
	buf.Myprintf("begin")
}

// WalkSubtree walks the nodes of the subtree
func (node *Begin) WalkSubtree(visit Visit) error {
	return nil
}

type Commit struct {
}

func (node *Commit) Format(buf *TrackedBuffer) {
	buf.Myprintf("commit")
}

// WalkSubtree walks the nodes of the subtree
func (node *Commit) WalkSubtree(visit Visit) error {
	return nil
}

type Rollback struct {
}

func (node *Rollback) Format(buf *TrackedBuffer) {
	buf.Myprintf("rollback")
}

// WalkSubtree walks the nodes of the subtree
func (node *Rollback) WalkSubtree(visit Visit) error {
	return nil
}

type SimpleSelect struct {
	Comments    Comments
	Distinct    string
	SelectExprs SelectExprs
}

func (node *SimpleSelect) Format(buf *TrackedBuffer) {
	buf.Myprintf("select %v%s%v", node.Comments, node.Distinct, node.SelectExprs)
}

func (node *SimpleSelect) WalkSubtree(visit Visit) error {
	return nil
}

func (*SimpleSelect) iStatement()       {}
func (*SimpleSelect) iSelectStatement() {}
func (*SimpleSelect) iInsertRows()      {}

type Admin struct {
	Name   string
	Values ValExprs
}

func (*Admin) iStatement() {}

func (node *Admin) Format(buf *TrackedBuffer) {
	buf.Myprintf("admin %s(%v)", node.Name, node.Values)
}

// WalkSubtree walks the nodes of the subtree
func (node *Admin) WalkSubtree(visit Visit) error {
	if node == nil {
		return nil
	}
	return Walk(
		visit,
		node.Values,
	)
}

type Show struct {
	Section     string
	Key         string
	From        ValExpr
	LikeOrWhere Expr
}

func (*Show) iStatement() {}

func (node *Show) Format(buf *TrackedBuffer) {
	buf.Myprintf("show %s %s %v %v", node.Section, node.Key, node.From, node.LikeOrWhere)
}

// WalkSubtree walks the nodes of the subtree
func (node *Show) WalkSubtree(visit Visit) error {
	if node == nil {
		return nil
	}
	return Walk(
		visit,
		node.From,
		node.LikeOrWhere,
	)
}
