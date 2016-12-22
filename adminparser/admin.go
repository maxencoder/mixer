//line admin.y:6
package adminparser

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
	mirrorRoute    *MirrorRoute
	rangeRoute     *RangeRoute
	keyRangeRoute  KeyRangeRoute
	rangeRouteList []KeyRangeRoute
	rangeNum       RangeNum
	tableRouterDef TableRouterDef
}

const LEX_ERROR = 57346
const ADMIN = 57347
const ADD = 57348
const ALTER = 57349
const DELETE = 57350
const SHOW = 57351
const DEFAULT = 57352
const DATABASE = 57353
const FROM = 57354
const HASH = 57355
const INF = 57356
const KEY = 57357
const KIND = 57358
const MIRROR = 57359
const MIRRORS = 57360
const MODULO = 57361
const TO = 57362
const RANGE = 57363
const TABLE = 57364
const TYPE = 57365
const ROUTE = 57366
const ROUTES = 57367
const ROUTER = 57368
const ID = 57369
const STRING = 57370
const NUMBER = 57371
const LIST_ARG = 57372
const COMMENT = 57373

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"LEX_ERROR",
	"ADMIN",
	"ADD",
	"ALTER",
	"DELETE",
	"SHOW",
	"DEFAULT",
	"DATABASE",
	"FROM",
	"HASH",
	"INF",
	"KEY",
	"KIND",
	"MIRROR",
	"MIRRORS",
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

const yyNprod = 38
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 114

var yyAct = [...]int{

	97, 96, 76, 68, 37, 67, 64, 48, 54, 52,
	36, 50, 70, 47, 27, 83, 26, 80, 77, 30,
	35, 77, 33, 101, 102, 90, 22, 87, 86, 34,
	69, 49, 32, 41, 31, 29, 45, 28, 25, 24,
	42, 94, 93, 46, 51, 19, 53, 55, 56, 92,
	51, 59, 55, 60, 58, 57, 20, 63, 21, 16,
	13, 91, 74, 65, 81, 78, 73, 75, 104, 79,
	17, 14, 18, 15, 82, 38, 84, 85, 62, 40,
	43, 72, 61, 39, 23, 89, 88, 1, 44, 66,
	71, 7, 95, 6, 98, 99, 5, 4, 100, 3,
	2, 0, 103, 0, 0, 0, 105, 8, 9, 10,
	11, 0, 0, 12,
}
var yyPact = [...]int{

	101, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 49, 48,
	34, 1, 79, 13, 12, -16, 11, 9, -16, 8,
	6, -16, -1000, -1000, -10, -10, 62, -1000, -10, -10,
	67, -10, -10, -1000, 4, -1000, 4, -26, 4, 4,
	4, 4, 4, 4, 4, -1000, -1000, -1000, 72, -1000,
	-1000, 63, -10, 40, -1000, -2, 65, -1000, -1000, 40,
	-1000, -16, -10, -1000, -8, 46, -11, -1000, 44, -1000,
	-1000, -8, -15, -8, -8, 0, -1000, -1000, -1, -1000,
	-2, -2, -1000, -3, -1000, -1000, 37, 25, -1000, 18,
	17, -16, -16, -16, -16, -8, -5, -1000, -1000, -4,
	-1000, -16, 50, -1000, -16, -5,
}
var yyPgo = [...]int{

	0, 100, 99, 97, 96, 93, 91, 10, 6, 90,
	8, 89, 5, 13, 11, 3, 1, 0, 4, 87,
	7, 2,
}
var yyR1 = [...]int{

	0, 19, 1, 1, 1, 1, 1, 2, 2, 2,
	2, 2, 3, 3, 3, 3, 4, 4, 4, 5,
	6, 13, 14, 8, 9, 10, 11, 11, 12, 15,
	15, 16, 16, 7, 17, 18, 20, 21,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 5, 5, 7,
	5, 7, 5, 5, 7, 5, 4, 4, 3, 2,
	2, 4, 7, 5, 8, 3, 1, 3, 5, 1,
	1, 1, 3, 3, 1, 1, 1, 1,
}
var yyChk = [...]int{

	-1000, -19, -1, -2, -3, -4, -5, -6, 6, 7,
	8, 9, 12, 11, 22, 24, 11, 22, 24, 11,
	22, 24, 25, 5, 26, 26, -17, 30, 26, 26,
	-17, 26, 26, -17, -18, 30, -7, -18, 13, 21,
	17, -18, -7, 13, 21, -18, -7, -13, -20, 27,
	-14, -20, 35, -20, -10, -20, -20, -13, -14, -20,
	-10, 10, 15, -18, -8, 23, -11, -12, -15, 32,
	14, -9, 16, -8, -17, -18, -21, 29, 19, -21,
	28, 20, -21, 30, -21, -21, 28, 28, -12, -15,
	28, 24, 24, 24, 24, -17, -16, -17, -17, -17,
	-21, 28, 28, -17, 18, -16,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 19, 20, 0, 0, 0, 34, 0, 0,
	0, 0, 0, 18, 0, 35, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 16, 17, 7, 0, 36,
	8, 0, 0, 0, 10, 0, 0, 12, 13, 0,
	15, 0, 0, 33, 0, 0, 0, 26, 0, 29,
	30, 0, 0, 0, 0, 0, 9, 37, 0, 25,
	0, 0, 11, 0, 14, 21, 0, 0, 27, 0,
	0, 0, 0, 0, 0, 0, 23, 31, 28, 0,
	22, 0, 0, 32, 0, 24,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	27, 29, 3, 3, 28, 3, 35,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 30, 31, 32, 33, 34,
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
		//line admin.y:92
		{
			setParseTree(yylex, yyDollar[1].command)
		}
	case 7:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line admin.y:105
		{
			yyVAL.command = &AddDbRouter{Db: yyDollar[4].str, Default: yyDollar[5].routeID}
		}
	case 8:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line admin.y:109
		{
			yyVAL.command = &AddTableRouter{Db: yyDollar[4].tableIdent.Db, Table: yyDollar[4].tableIdent.Table, Key: yyDollar[5].tableRouterDef.Key, Route: yyDollar[5].tableRouterDef.Route}
		}
	case 9:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line admin.y:113
		{
			yyVAL.command = &AddRoute{Name: yyDollar[3].str, Route: yyDollar[6].hashRoute}
		}
	case 10:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line admin.y:117
		{
			yyVAL.command = &AddRoute{Name: yyDollar[3].str, Route: yyDollar[5].rangeRoute}
		}
	case 11:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line admin.y:121
		{
			yyVAL.command = &AddRoute{Name: yyDollar[3].str, Route: yyDollar[6].mirrorRoute}
		}
	case 12:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line admin.y:127
		{
			yyVAL.command = &AlterDbRouter{Db: yyDollar[4].str, Default: yyDollar[5].routeID}
		}
	case 13:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line admin.y:131
		{
			yyVAL.command = &AlterTableRouter{Db: yyDollar[4].tableIdent.Db, Table: yyDollar[4].tableIdent.Table, Key: yyDollar[5].tableRouterDef.Key, Route: yyDollar[5].tableRouterDef.Route}
		}
	case 14:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line admin.y:135
		{
			yyVAL.command = &AlterRoute{Name: yyDollar[3].str, Route: yyDollar[6].hashRoute}
		}
	case 15:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line admin.y:139
		{
			yyVAL.command = &AlterRoute{Name: yyDollar[3].str, Route: yyDollar[5].rangeRoute}
		}
	case 16:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line admin.y:145
		{
			yyVAL.command = &DeleteDbRouter{Db: yyDollar[4].str}
		}
	case 17:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line admin.y:149
		{
			yyVAL.command = &DeleteTableRouter{Db: yyDollar[4].tableIdent.Db, Table: yyDollar[4].tableIdent.Table}
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line admin.y:153
		{
			yyVAL.command = &DeleteRoute{Name: yyDollar[3].str}
		}
	case 19:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line admin.y:159
		{
			yyVAL.command = &Show{}
		}
	case 20:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line admin.y:165
		{
			yyVAL.command = &FromAdmin{}
		}
	case 21:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line admin.y:171
		{
			yyVAL.routeID = RouteID(yyDollar[3].str)
		}
	case 22:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line admin.y:177
		{
			yyVAL.tableRouterDef = TableRouterDef{Key: yyDollar[3].str, Route: RouteID(yyDollar[6].str)}
		}
	case 23:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line admin.y:183
		{
			yyVAL.hashRoute = &HashRoute{Type: ModuloStr, Routes: yyDollar[5].strings}
		}
	case 24:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line admin.y:189
		{
			yyVAL.mirrorRoute = &MirrorRoute{Kind: string(yyDollar[2].bytes), Main: yyDollar[5].str, Mirrors: yyDollar[8].strings}
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line admin.y:195
		{
			yyVAL.rangeRoute = &RangeRoute{Ranges: yyDollar[2].rangeRouteList}
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line admin.y:201
		{
			yyVAL.rangeRouteList = []KeyRangeRoute{yyDollar[1].keyRangeRoute}
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line admin.y:205
		{
			yyVAL.rangeRouteList = append(yyVAL.rangeRouteList, yyDollar[3].keyRangeRoute)
		}
	case 28:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line admin.y:211
		{
			yyVAL.keyRangeRoute = KeyRangeRoute{Start: yyDollar[1].rangeNum, End: yyDollar[3].rangeNum, Route: yyDollar[5].str}
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line admin.y:217
		{
			n, err := strconv.ParseInt(string(yyDollar[1].bytes), 10, 64)
			if err != nil {
				yylex.Error("expecting int")
				return 1
			}
			yyVAL.rangeNum = RangeNum{Num: n}
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line admin.y:226
		{
			yyVAL.rangeNum = RangeNum{Inf: true}
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line admin.y:232
		{
			yyVAL.strings = []string{yyDollar[1].str}
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line admin.y:236
		{
			yyVAL.strings = append(yyVAL.strings, yyDollar[3].str)
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line admin.y:242
		{
			yyVAL.tableIdent = TableIdent{Db: yyDollar[1].str, Table: yyDollar[3].str}
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line admin.y:248
		{
			yyVAL.str = string(yyDollar[1].bytes)
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line admin.y:254
		{
			yyVAL.str = string(yyDollar[1].bytes)
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line admin.y:260
		{
			if incNesting(yylex) {
				yylex.Error("max nesting level reached")
				return 1
			}
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line admin.y:269
		{
			decNesting(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
