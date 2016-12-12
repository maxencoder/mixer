//line admin.y:6
package admin

import __yyfmt__ "fmt"

//line admin.y:6
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

//line admin.y:38
type yySymType struct {
	yys        int
	empty      struct{}
	command    Command
	byt        byte
	bytes      []byte
	strings    []string
	bytes2     [][]byte
	str        string
	tableIdent TableIdent
	routeID    RouteID

	hashRoute      *HashRoute
	keyRangeRoute  KeyRangeRoute
	rangeRouteList []KeyRangeRoute
	rangeRoute     *RangeRoute
	rangeNum       RangeNum
	tableRouterDef TableRouterDef
}

const LEX_ERROR = 57346
const ADD = 57347
const ALTER = 57348
const DELETE = 57349
const SHOW = 57350
const DEFAULT = 57351
const DATABASE = 57352
const HASH = 57353
const INF = 57354
const KEY = 57355
const MODULO = 57356
const TO = 57357
const RANGE = 57358
const TABLE = 57359
const TYPE = 57360
const ROUTE = 57361
const ROUTES = 57362
const ROUTER = 57363
const ID = 57364
const STRING = 57365
const NUMBER = 57366
const LIST_ARG = 57367
const COMMENT = 57368

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"LEX_ERROR",
	"ADD",
	"ALTER",
	"DELETE",
	"SHOW",
	"DEFAULT",
	"DATABASE",
	"HASH",
	"INF",
	"KEY",
	"MODULO",
	"TO",
	"RANGE",
	"TABLE",
	"TYPE",
	"ROUTE",
	"ROUTES",
	"ROUTER",
	"'('",
	"','",
	"')'",
	"ID",
	"STRING",
	"NUMBER",
	"LIST_ARG",
	"COMMENT",
	"'.'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 34
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 93

var yyAct = [...]int{

	23, 69, 63, 34, 62, 59, 44, 50, 33, 48,
	46, 65, 43, 24, 73, 70, 70, 27, 32, 89,
	30, 78, 29, 77, 45, 31, 64, 28, 26, 37,
	25, 22, 41, 21, 20, 38, 83, 82, 42, 81,
	47, 60, 49, 51, 17, 47, 54, 51, 55, 53,
	52, 18, 58, 19, 39, 14, 74, 67, 11, 40,
	66, 68, 15, 72, 16, 12, 35, 13, 75, 76,
	71, 36, 57, 7, 8, 9, 10, 80, 79, 56,
	1, 85, 84, 86, 87, 61, 88, 6, 5, 4,
	90, 3, 2,
}
var yyPact = [...]int{

	68, -1000, -1000, -1000, -1000, -1000, -1000, 48, 45, 34,
	14, 12, 10, -12, 9, 7, -12, 6, 1, -12,
	-1000, -7, -7, 55, -1000, -7, -7, 43, -7, -7,
	-1000, 2, -1000, 2, -21, 2, 2, 2, 2, 2,
	2, -1000, -1000, -1000, 70, -1000, -1000, 59, -7, 23,
	-1000, -1, -1000, -1000, 23, -1000, -12, -7, -1000, -8,
	56, -9, -1000, 41, -1000, -1000, -8, -8, 0, -1000,
	-1000, -2, -1000, -1, -1, -1000, -1000, 20, 18, -1000,
	17, -12, -12, -12, -8, -4, -1000, -1000, -1000, -12,
	-1000,
}
var yyPgo = [...]int{

	0, 92, 91, 89, 88, 87, 8, 5, 7, 85,
	4, 12, 10, 2, 81, 0, 3, 80, 6, 1,
}
var yyR1 = [...]int{

	0, 17, 1, 1, 1, 1, 2, 2, 2, 2,
	3, 3, 3, 3, 4, 4, 4, 5, 11, 12,
	7, 8, 9, 9, 10, 13, 13, 14, 14, 6,
	15, 16, 18, 19,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 5, 5, 7, 5,
	5, 5, 7, 5, 4, 4, 3, 2, 4, 7,
	5, 3, 1, 3, 5, 1, 1, 1, 3, 3,
	1, 1, 1, 1,
}
var yyChk = [...]int{

	-1000, -17, -1, -2, -3, -4, -5, 5, 6, 7,
	8, 10, 17, 19, 10, 17, 19, 10, 17, 19,
	20, 21, 21, -15, 25, 21, 21, -15, 21, 21,
	-15, -16, 25, -6, -16, 11, 16, -16, -6, 11,
	16, -16, -6, -11, -18, 22, -12, -18, 30, -18,
	-8, -18, -11, -12, -18, -8, 9, 13, -16, -7,
	18, -9, -10, -13, 27, 12, -7, -15, -16, -19,
	24, 14, -19, 23, 15, -19, -19, 23, 23, -10,
	-13, 19, 19, 19, -15, -14, -15, -15, -19, 23,
	-15,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	17, 0, 0, 0, 30, 0, 0, 0, 0, 0,
	16, 0, 31, 0, 0, 0, 0, 0, 0, 0,
	0, 14, 15, 6, 0, 32, 7, 0, 0, 0,
	9, 0, 10, 11, 0, 13, 0, 0, 29, 0,
	0, 0, 22, 0, 25, 26, 0, 0, 0, 8,
	33, 0, 21, 0, 0, 12, 18, 0, 0, 23,
	0, 0, 0, 0, 0, 20, 27, 24, 19, 0,
	28,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	22, 24, 3, 3, 23, 3, 30,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	25, 26, 27, 28, 29,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line admin.y:89
		{
			setParseTree(yylex, yyDollar[1].command)
		}
	case 6:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line admin.y:101
		{
			yyVAL.command = &AddDbRouter{Db: yyDollar[4].str, Default: yyDollar[5].routeID}
		}
	case 7:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line admin.y:105
		{
			yyVAL.command = &AddTableRouter{Db: yyDollar[4].tableIdent.Db, Table: yyDollar[4].tableIdent.Table, Key: yyDollar[5].tableRouterDef.Key, Route: yyDollar[5].tableRouterDef.Route}
		}
	case 8:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line admin.y:109
		{
			yyVAL.command = &AddRoute{Name: yyDollar[3].str, Route: yyDollar[6].hashRoute}
		}
	case 9:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line admin.y:113
		{
			yyVAL.command = &AddRoute{Name: yyDollar[3].str, Route: yyDollar[5].rangeRoute}
		}
	case 10:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line admin.y:119
		{
			yyVAL.command = &AlterDbRouter{Db: yyDollar[4].str, Default: yyDollar[5].routeID}
		}
	case 11:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line admin.y:123
		{
			yyVAL.command = &AlterTableRouter{Db: yyDollar[4].tableIdent.Db, Table: yyDollar[4].tableIdent.Table, Key: yyDollar[5].tableRouterDef.Key, Route: yyDollar[5].tableRouterDef.Route}
		}
	case 12:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line admin.y:127
		{
			yyVAL.command = &AlterRoute{Name: yyDollar[3].str, Route: yyDollar[6].hashRoute}
		}
	case 13:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line admin.y:131
		{
			yyVAL.command = &AlterRoute{Name: yyDollar[3].str, Route: yyDollar[5].rangeRoute}
		}
	case 14:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line admin.y:137
		{
			yyVAL.command = &DeleteDbRouter{Db: yyDollar[4].str}
		}
	case 15:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line admin.y:141
		{
			yyVAL.command = &DeleteTableRouter{Db: yyDollar[4].tableIdent.Db, Table: yyDollar[4].tableIdent.Table}
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line admin.y:145
		{
			yyVAL.command = &DeleteRoute{Name: yyDollar[3].str}
		}
	case 17:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line admin.y:151
		{
			yyVAL.command = &Show{}
		}
	case 18:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line admin.y:157
		{
			yyVAL.routeID = RouteID(yyDollar[3].str)
		}
	case 19:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line admin.y:163
		{
			yyVAL.tableRouterDef = TableRouterDef{Key: yyDollar[3].str, Route: RouteID(yyDollar[6].str)}
		}
	case 20:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line admin.y:169
		{
			yyVAL.hashRoute = &HashRoute{Type: ModuloStr, Routes: yyDollar[5].strings}
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line admin.y:175
		{
			yyVAL.rangeRoute = &RangeRoute{Ranges: yyDollar[2].rangeRouteList}
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line admin.y:181
		{
			yyVAL.rangeRouteList = []KeyRangeRoute{yyDollar[1].keyRangeRoute}
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line admin.y:185
		{
			yyVAL.rangeRouteList = append(yyVAL.rangeRouteList, yyDollar[3].keyRangeRoute)
		}
	case 24:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line admin.y:191
		{
			yyVAL.keyRangeRoute = KeyRangeRoute{Start: yyDollar[1].rangeNum, End: yyDollar[3].rangeNum, Route: yyDollar[5].str}
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line admin.y:197
		{
			n, err := strconv.ParseInt(string(yyDollar[1].bytes), 10, 64)
			if err != nil {
				yylex.Error("expecting int")
				return 1
			}
			yyVAL.rangeNum = RangeNum{Num: n}
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line admin.y:206
		{
			yyVAL.rangeNum = RangeNum{Inf: true}
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line admin.y:212
		{
			yyVAL.strings = []string{yyDollar[1].str}
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line admin.y:216
		{
			yyVAL.strings = append(yyVAL.strings, yyDollar[3].str)
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line admin.y:222
		{
			yyVAL.tableIdent = TableIdent{Db: yyDollar[1].str, Table: yyDollar[3].str}
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line admin.y:228
		{
			yyVAL.str = string(yyDollar[1].bytes)
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line admin.y:234
		{
			yyVAL.str = string(yyDollar[1].bytes)
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line admin.y:240
		{
			if incNesting(yylex) {
				yylex.Error("max nesting level reached")
				return 1
			}
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line admin.y:249
		{
			decNesting(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
