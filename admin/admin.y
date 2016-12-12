// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

%{
package sqlparser

import (
  "strconv"
)

func setParseTree(yylex interface{}, stmt Command) {
  yylex.(*Tokenizer).ParseTree = stmt
}

func setAllowComments(yylex interface{}, allow bool) {
  yylex.(*Tokenizer).AllowComments = allow
}

func incNesting(yylex interface{}) bool {
  yylex.(*Tokenizer).nesting++
  if yylex.(*Tokenizer).nesting == 200 {
    return true
  }
  return false
}

func decNesting(yylex interface{}) {
  yylex.(*Tokenizer).nesting--
}

func forceEOF(yylex interface{}) {
  yylex.(*Tokenizer).ForceEOF = true
}

%}

%union {
  empty       struct{}
  command     Command
  byt         byte
  bytes       []byte
  strings     []string
  bytes2      [][]byte
  str         string
  tableIdent  TableIdent

  hashRoute         *HashRoute
  keyRangeRoute     KeyRangeRoute
  rangeRouteList    []KeyRangeRoute
  rangeRoute        *RangeRoute
  numberInf         RangeNum
}

%token LEX_ERROR
%token <empty> ADD ALTER DELETE SHOW
%token <empty> DEFAULT
%token <empty> HASH RANGE
%token <empty> INF MODULO TO
%token <empty> TYPE ROUTE ROUTES

%token <empty> '(' ',' ')'
%token <bytes> ID STRING NUMBER LIST_ARG COMMENT

%type <command> command
%type <command> add_command alter_command delete_command show_command
//%type <tableName> table_name
//%type <tableIdent> table_id
%type <hashRoute> hash_route_def
%type <rangeRoute> range_route
%type <rangeRouteList> range_route_list
%type <keyRangeRoute> key_range_route
%type <numberInf> number_inf
%type <strings> route_list
%type <str> route_id

%start any_command

%%

any_command:
  command
  {
    setParseTree(yylex, $1)
  }

command:
  add_command
| alter_command
| delete_command
| show_command

add_command:
  ADD ROUTE route_id HASH openb hash_route_def closeb
  {
    $$ = &AddRoute{Name: $3, Route: $6}
  }
| ADD ROUTE route_id RANGE range_route
  {
    $$ = &AddRoute{Name: $3, Route: $5}
  }

alter_command:
  ALTER ROUTE route_id HASH openb hash_route_def closeb
  {
    $$ = &AlterRoute{Name: $3, Route: $6}
  }
| ALTER ROUTE route_id RANGE range_route
  {
    $$ = &AlterRoute{Name: $3, Route: $5}
  }

delete_command:
  DELETE ROUTE route_id
  {
    $$ = &Delete{Name: $3}
  }

show_command:
  SHOW ROUTES
  {
    $$ = &Show{}
  }

hash_route_def:
  TYPE MODULO ',' ROUTE route_list
  {
    $$ = &HashRoute{Type: ModuloStr, Routes: $5}
  }

range_route:
  openb range_route_list closeb
  {
    $$ = &RangeRoute{Ranges: $2}
  }

range_route_list:
  key_range_route
  {
    $$ = []KeyRangeRoute{$1}
  }
| range_route_list ',' key_range_route
  {
    $$ = append($$, $3)
  }

key_range_route:
  number_inf TO number_inf ROUTE route_id
  {
    $$ = KeyRangeRoute{Start: $1, End: $3, Route: $5}
  }

number_inf:
  NUMBER
  {
    n, err := strconv.ParseInt(string($1), 10, 64)
    if err != nil {
      yylex.Error("expecting int")
      return 1
    }
    $$ = RangeNum{Num: n}
  }
| INF
  {
    $$ = RangeNum{Inf: true}
  }

route_list:
  route_id
  {
    $$ = []string{$1}
  }
| route_list ',' route_id
  {
    $$ = append($$, $3)
  }

route_id:
  ID
  {
    $$ = string($1)
  }

openb:
  '('
  {
    if incNesting(yylex) {
      yylex.Error("max nesting level reached")
      return 1
    }
  }

closeb:
  ')'
  {
    decNesting(yylex)
  }
