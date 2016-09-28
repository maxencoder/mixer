package sqlparser

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

func (*Begin) IStatement()    {}
func (*Commit) IStatement()   {}
func (*Rollback) IStatement() {}

type Begin struct {
}

func (node *Begin) Format(buf *TrackedBuffer) {
	buf.Myprintf("begin")
}

type Commit struct {
}

func (node *Commit) Format(buf *TrackedBuffer) {
	buf.Myprintf("commit")
}

type Rollback struct {
}

func (node *Rollback) Format(buf *TrackedBuffer) {
	buf.Myprintf("rollback")
}

// Replace represents an REPLACE statement.
type Replace struct {
	Comments Comments
	Table    *TableName
	Columns  Columns
	Rows     InsertRows
}

func (node *Replace) Format(buf *TrackedBuffer) {
	buf.Myprintf("replace %vinto %v%v %v%v",
		node.Comments,
		node.Table, node.Columns, node.Rows)
}

func (*Replace) IStatement() {}

type SimpleSelect struct {
	Comments    Comments
	Distinct    string
	SelectExprs SelectExprs
}

func (node *SimpleSelect) Format(buf *TrackedBuffer) {
	buf.Myprintf("select %v%s%v", node.Comments, node.Distinct, node.SelectExprs)
}

func (*SimpleSelect) IStatement()       {}
func (*SimpleSelect) ISelectStatement() {}
func (*SimpleSelect) IInsertRows()      {}

type Admin struct {
	Name   SQLName
	Values ValExprs
}

func (*Admin) IStatement() {}

func (node *Admin) Format(buf *TrackedBuffer) {
	buf.Myprintf("admin %s(%v)", node.Name, node.Values)
}

type Show struct {
	Section     string
	Key         string
	From        ValExpr
	LikeOrWhere Expr
}

func (*Show) IStatement() {}

func (node *Show) Format(buf *TrackedBuffer) {
	buf.Myprintf("show %s %s %v %v", node.Section, node.Key, node.From, node.LikeOrWhere)
}

/*

// -- sql.y

// Transaction Tokens
%token <empty> BEGIN COMMIT ROLLBACK

// Charset Tokens
%token <empty> NAMES

// Replace
%token <empty> REPLACE

// Mixer admin
%token <empty> ADMIN

// Show
%token <empty> DATABASES TABLES PROXY


// - types
%type <statement> begin_statement commit_statement rollback_statement
%type <statement> replace_statement
%type <statement> show_statement
%type <statement> admin_statement

%type <valExpr> from_opt
%type <expr> like_or_where_opt

//after ^command:
| begin_statement
| commit_statement
| rollback_statement
| replace_statement
| show_statement
| admin_statement

// after ^insert_statement:

replace_statement:
  REPLACE comment_opt INTO dml_table_expression column_list_opt row_list
  {
    $$ = &Replace{Comments: Comments($2), Table: $4, Columns: $5, Rows: $6}
  }
| REPLACE comment_opt INTO dml_table_expression SET update_list
  {
    cols := make(Columns, 0, len($6))
    vals := make(ValTuple, 0, len($6))
    for _, col := range $6 {
      cols = append(cols, &NonStarExpr{Expr: col.Name})
      vals = append(vals, col.Expr)
    }
    $$ = &Replace{Comments: Comments($2), Table: $4, Columns: cols, Rows: Values{vals}}
  }


//
begin_statement:
  BEGIN
  {
    $$ = &Begin{}
  }

commit_statement:
  COMMIT
  {
    $$ = &Commit{}
  }

rollback_statement:
  ROLLBACK
  {
    $$ = &Rollback{}
  }

admin_statement:
  ADMIN sql_id '(' value_expression_list ')'
  {
    $$ = &Admin{Name : $2, Values : $4}
  }

show_statement:
  SHOW DATABASES like_or_where_opt
  {
    $$ = &Show{Section: "databases", LikeOrWhere: $3}
  }
| SHOW TABLES from_opt like_or_where_opt
  {
    $$ = &Show{Section: "tables", From: $3, LikeOrWhere: $4}
  }
| SHOW PROXY sql_id from_opt like_or_where_opt
  {
    $$ = &Show{Section: "proxy", Key: string($3), From: $4, LikeOrWhere: $5}
  }


//where_expression_opt:

like_or_where_opt:
  {
    $$ = nil
  }
| WHERE boolean_expression
  {
    $$ = $2
  }
| LIKE value_expression
  {
    $$ = $2
  }

from_opt:
  {
    $$ = nil
  }
| FROM value_expression
  {
    $$ = $2
  }


// token.go
//		in tokens:
 	//for mixer admin
 	"admin": ADMIN,
 	"proxy": PROXY,

*/
