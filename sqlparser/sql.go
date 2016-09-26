//line sql.y:6
package sqlparser

import __yyfmt__ "fmt"

//line sql.y:6
import "bytes"

func SetParseTree(yylex interface{}, stmt Statement) {
	yylex.(*Tokenizer).ParseTree = stmt
}

func SetAllowComments(yylex interface{}, allow bool) {
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

func ForceEOF(yylex interface{}) {
	yylex.(*Tokenizer).ForceEOF = true
}

var (
	SHARE        = []byte("share")
	MODE         = []byte("mode")
	IF_BYTES     = []byte("if")
	VALUES_BYTES = []byte("values")
)

//line sql.y:44
type yySymType struct {
	yys         int
	empty       struct{}
	statement   Statement
	selStmt     SelectStatement
	byt         byte
	bytes       []byte
	bytes2      [][]byte
	str         string
	selectExprs SelectExprs
	selectExpr  SelectExpr
	columns     Columns
	colName     *ColName
	tableExprs  TableExprs
	tableExpr   TableExpr
	smTableExpr SimpleTableExpr
	tableName   *TableName
	indexHints  *IndexHints
	expr        Expr
	boolExpr    BoolExpr
	valExpr     ValExpr
	tuple       Tuple
	valExprs    ValExprs
	values      Values
	subquery    *Subquery
	caseExpr    *CaseExpr
	whens       []*When
	when        *When
	orderBy     OrderBy
	order       *Order
	limit       *Limit
	insRows     InsertRows
	updateExprs UpdateExprs
	updateExpr  *UpdateExpr
}

const LEX_ERROR = 57346
const SELECT = 57347
const INSERT = 57348
const UPDATE = 57349
const DELETE = 57350
const FROM = 57351
const WHERE = 57352
const GROUP = 57353
const HAVING = 57354
const ORDER = 57355
const BY = 57356
const LIMIT = 57357
const FOR = 57358
const ALL = 57359
const DISTINCT = 57360
const AS = 57361
const EXISTS = 57362
const IN = 57363
const IS = 57364
const LIKE = 57365
const BETWEEN = 57366
const NULL = 57367
const ASC = 57368
const DESC = 57369
const VALUES = 57370
const INTO = 57371
const DUPLICATE = 57372
const KEY = 57373
const DEFAULT = 57374
const SET = 57375
const LOCK = 57376
const ID = 57377
const STRING = 57378
const NUMBER = 57379
const VALUE_ARG = 57380
const COMMENT = 57381
const LE = 57382
const GE = 57383
const NE = 57384
const NULL_SAFE_EQUAL = 57385
const UNION = 57386
const MINUS = 57387
const EXCEPT = 57388
const INTERSECT = 57389
const JOIN = 57390
const STRAIGHT_JOIN = 57391
const LEFT = 57392
const RIGHT = 57393
const INNER = 57394
const OUTER = 57395
const CROSS = 57396
const NATURAL = 57397
const USE = 57398
const FORCE = 57399
const ON = 57400
const AND = 57401
const OR = 57402
const NOT = 57403
const UNARY = 57404
const CASE = 57405
const WHEN = 57406
const THEN = 57407
const ELSE = 57408
const END = 57409
const BEGIN = 57410
const COMMIT = 57411
const ROLLBACK = 57412
const NAMES = 57413
const REPLACE = 57414
const ADMIN = 57415
const SHOW = 57416
const DATABASES = 57417
const TABLES = 57418
const PROXY = 57419
const CREATE = 57420
const ALTER = 57421
const DROP = 57422
const RENAME = 57423
const TABLE = 57424
const INDEX = 57425
const VIEW = 57426
const TO = 57427
const IGNORE = 57428
const IF = 57429
const UNIQUE = 57430
const USING = 57431

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"LEX_ERROR",
	"SELECT",
	"INSERT",
	"UPDATE",
	"DELETE",
	"FROM",
	"WHERE",
	"GROUP",
	"HAVING",
	"ORDER",
	"BY",
	"LIMIT",
	"FOR",
	"ALL",
	"DISTINCT",
	"AS",
	"EXISTS",
	"IN",
	"IS",
	"LIKE",
	"BETWEEN",
	"NULL",
	"ASC",
	"DESC",
	"VALUES",
	"INTO",
	"DUPLICATE",
	"KEY",
	"DEFAULT",
	"SET",
	"LOCK",
	"ID",
	"STRING",
	"NUMBER",
	"VALUE_ARG",
	"COMMENT",
	"LE",
	"GE",
	"NE",
	"NULL_SAFE_EQUAL",
	"'('",
	"'='",
	"'<'",
	"'>'",
	"'~'",
	"UNION",
	"MINUS",
	"EXCEPT",
	"INTERSECT",
	"','",
	"JOIN",
	"STRAIGHT_JOIN",
	"LEFT",
	"RIGHT",
	"INNER",
	"OUTER",
	"CROSS",
	"NATURAL",
	"USE",
	"FORCE",
	"ON",
	"AND",
	"OR",
	"NOT",
	"'&'",
	"'|'",
	"'^'",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"'.'",
	"UNARY",
	"CASE",
	"WHEN",
	"THEN",
	"ELSE",
	"END",
	"BEGIN",
	"COMMIT",
	"ROLLBACK",
	"NAMES",
	"REPLACE",
	"ADMIN",
	"SHOW",
	"DATABASES",
	"TABLES",
	"PROXY",
	"CREATE",
	"ALTER",
	"DROP",
	"RENAME",
	"TABLE",
	"INDEX",
	"VIEW",
	"TO",
	"IGNORE",
	"IF",
	"UNIQUE",
	"USING",
	"')'",
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

const yyNprod = 217
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 656

var yyAct = [...]int{

	109, 317, 147, 384, 253, 186, 352, 309, 271, 106,
	137, 95, 118, 72, 70, 262, 264, 226, 201, 107,
	161, 162, 209, 96, 393, 196, 156, 187, 3, 84,
	88, 315, 224, 59, 77, 285, 286, 287, 288, 289,
	278, 290, 291, 44, 74, 46, 224, 79, 140, 47,
	81, 334, 336, 73, 85, 80, 49, 112, 50, 93,
	254, 61, 117, 363, 362, 123, 101, 34, 35, 36,
	37, 51, 99, 114, 115, 116, 254, 149, 254, 136,
	361, 113, 78, 254, 254, 121, 254, 144, 343, 263,
	335, 91, 139, 148, 243, 151, 75, 154, 223, 158,
	153, 52, 53, 54, 103, 161, 162, 188, 119, 120,
	97, 189, 56, 57, 58, 124, 135, 263, 296, 307,
	345, 100, 150, 254, 160, 192, 133, 195, 74, 128,
	358, 74, 215, 130, 205, 204, 244, 73, 130, 122,
	73, 104, 199, 310, 203, 206, 274, 71, 60, 310,
	67, 213, 154, 143, 216, 219, 174, 175, 176, 101,
	232, 205, 161, 162, 83, 360, 236, 104, 230, 241,
	242, 220, 245, 246, 247, 248, 249, 250, 251, 252,
	231, 104, 104, 222, 237, 152, 233, 190, 191, 359,
	193, 255, 256, 101, 101, 257, 328, 332, 74, 74,
	326, 329, 258, 260, 198, 327, 331, 73, 269, 330,
	275, 146, 267, 212, 214, 211, 202, 126, 270, 86,
	129, 276, 74, 198, 266, 184, 185, 281, 224, 202,
	155, 73, 369, 228, 104, 368, 280, 279, 145, 104,
	104, 230, 298, 299, 282, 347, 295, 132, 266, 221,
	169, 170, 171, 172, 173, 174, 175, 176, 297, 283,
	113, 229, 101, 302, 113, 304, 94, 197, 104, 104,
	113, 303, 130, 314, 156, 313, 125, 306, 113, 316,
	104, 60, 294, 234, 235, 312, 169, 170, 171, 172,
	173, 174, 175, 176, 75, 339, 230, 230, 293, 390,
	338, 324, 325, 341, 159, 18, 228, 366, 342, 337,
	344, 172, 173, 174, 175, 176, 74, 391, 349, 321,
	60, 350, 353, 320, 273, 348, 285, 286, 287, 288,
	289, 354, 290, 291, 218, 229, 217, 104, 200, 68,
	141, 104, 138, 364, 113, 134, 340, 131, 365, 169,
	170, 171, 172, 173, 174, 175, 176, 82, 127, 346,
	154, 228, 228, 375, 373, 18, 87, 367, 66, 301,
	397, 381, 353, 207, 142, 383, 382, 318, 385, 385,
	385, 74, 386, 387, 238, 308, 239, 240, 265, 392,
	73, 394, 395, 89, 398, 388, 64, 259, 399, 112,
	400, 62, 357, 319, 117, 272, 90, 123, 34, 35,
	36, 37, 356, 323, 99, 114, 115, 116, 202, 92,
	69, 396, 380, 113, 18, 39, 17, 121, 16, 15,
	14, 13, 12, 18, 104, 208, 104, 45, 38, 377,
	378, 379, 277, 210, 48, 76, 103, 268, 112, 389,
	119, 120, 97, 117, 370, 351, 123, 124, 40, 41,
	42, 43, 355, 75, 114, 115, 116, 322, 305, 55,
	194, 261, 113, 111, 112, 108, 121, 110, 374, 117,
	376, 122, 123, 311, 254, 18, 105, 163, 102, 75,
	114, 115, 116, 333, 227, 103, 284, 225, 113, 119,
	120, 98, 121, 292, 157, 117, 124, 63, 123, 18,
	19, 20, 21, 33, 65, 75, 114, 115, 116, 371,
	372, 103, 11, 10, 113, 119, 120, 9, 121, 8,
	122, 117, 124, 7, 123, 6, 5, 22, 4, 2,
	1, 75, 114, 115, 116, 0, 0, 0, 0, 0,
	113, 119, 120, 0, 121, 0, 122, 0, 124, 0,
	0, 169, 170, 171, 172, 173, 174, 175, 176, 0,
	0, 0, 0, 0, 0, 0, 0, 119, 120, 0,
	0, 0, 122, 0, 124, 0, 0, 27, 28, 29,
	0, 30, 32, 31, 0, 0, 0, 23, 24, 26,
	25, 164, 168, 166, 167, 0, 300, 0, 122, 169,
	170, 171, 172, 173, 174, 175, 176, 0, 0, 0,
	180, 181, 182, 183, 0, 177, 178, 179, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 165, 169, 170,
	171, 172, 173, 174, 175, 176,
}
var yyPact = [...]int{

	504, -1000, -1000, 359, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -54, -43, -26, 4, -1000, -1000, -1000,
	-1000, 22, 246, 419, 384, -1000, -1000, -1000, 378, -1000,
	339, 304, 411, 61, -68, -16, 246, -1000, -42, 246,
	-1000, 322, -73, 246, -73, 337, 383, 410, 246, 222,
	-1000, -1000, -1000, 37, -1000, 237, 304, 325, 53, 304,
	80, 312, -1000, 202, -1000, 50, 310, 49, 246, -1000,
	307, -1000, -52, 305, 354, 89, 246, 304, -1000, 454,
	506, 383, 506, 410, 506, 221, -1000, -1000, 285, 48,
	97, 580, -1000, 454, 428, -1000, -1000, -1000, 506, 220,
	220, -1000, 220, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 506, -1000, 234, 259, 303, 408,
	259, -1000, 506, 246, -1000, 353, -82, -1000, 119, -1000,
	301, -1000, -1000, 299, -1000, 216, 97, 580, 218, 480,
	-1000, 218, 383, -7, 218, 226, 37, -1000, -1000, 246,
	113, 454, 454, 506, 220, 363, 506, 506, 69, 506,
	506, 506, 506, 506, 506, 506, 506, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -45, -21, 18, 580, -1000,
	379, 37, -1000, 419, 10, 218, 360, 259, 259, 219,
	-1000, 392, 454, -1000, 218, -1000, -1000, -1000, 82, 246,
	-1000, -60, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	360, 259, -1000, -1000, 506, 206, 272, 263, 300, 42,
	-1000, -1000, -1000, -1000, -1000, -1000, 218, -1000, 220, 506,
	506, 218, 541, -1000, 344, 240, 240, 240, 83, 83,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -27, 37,
	-27, 38, -1000, 454, 79, 220, 359, 85, -22, -1000,
	392, 362, 389, 97, 288, -1000, -1000, 284, -1000, -1000,
	80, 218, 402, 226, 226, -1000, -1000, 146, 142, 155,
	152, 143, -11, -1000, 274, -19, 260, -1000, 218, 281,
	506, -1000, -1000, -27, -1000, 6, -1000, 506, 40, -1000,
	329, 192, -1000, -1000, -1000, 259, 362, -1000, 506, 506,
	-1000, -1000, 400, 388, 272, 66, -1000, 135, -1000, 111,
	-1000, -1000, -1000, -1000, -18, -34, -35, -1000, -1000, -1000,
	506, 218, -1000, -1000, 218, 506, 276, 220, -1000, -1000,
	182, 179, -1000, 493, -1000, 392, 454, 506, 454, -1000,
	-1000, 220, 220, 220, 218, 218, 415, -1000, 506, 506,
	-1000, -1000, -1000, 362, 97, 175, 97, 246, 246, 246,
	259, 218, -1000, 283, -29, -1000, -29, -29, 80, -1000,
	414, 349, -1000, 246, -1000, -1000, -1000, 246, -1000, 246,
	-1000,
}
var yyPgo = [...]int{

	0, 540, 539, 27, 538, 536, 535, 533, 529, 527,
	523, 522, 438, 514, 513, 507, 11, 23, 504, 503,
	501, 497, 17, 496, 494, 150, 493, 3, 18, 121,
	488, 487, 16, 486, 2, 19, 5, 483, 477, 12,
	475, 9, 473, 471, 15, 470, 468, 467, 462, 8,
	455, 6, 454, 1, 449, 25, 447, 7, 14, 13,
	164, 445, 444, 443, 442, 437, 435, 0, 10, 432,
	431, 430, 429, 428, 426, 91, 30, 425, 77, 4,
}
var yyR1 = [...]int{

	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 3, 3, 3,
	4, 4, 72, 72, 5, 6, 7, 7, 69, 70,
	71, 74, 73, 73, 73, 8, 8, 8, 9, 9,
	9, 10, 11, 11, 11, 77, 12, 13, 13, 14,
	14, 14, 14, 14, 15, 15, 16, 16, 17, 17,
	17, 20, 20, 18, 18, 18, 21, 21, 22, 22,
	22, 22, 19, 19, 19, 23, 23, 23, 23, 23,
	23, 23, 23, 23, 24, 24, 24, 25, 25, 26,
	26, 26, 26, 27, 27, 28, 28, 76, 76, 76,
	75, 75, 29, 29, 29, 29, 29, 30, 30, 30,
	30, 30, 30, 30, 30, 30, 30, 31, 31, 31,
	31, 31, 31, 31, 32, 32, 37, 37, 35, 35,
	39, 36, 36, 34, 34, 34, 34, 34, 34, 34,
	34, 34, 34, 34, 34, 34, 34, 34, 34, 34,
	38, 38, 40, 40, 40, 42, 45, 45, 43, 43,
	44, 46, 46, 41, 41, 33, 33, 33, 33, 47,
	47, 48, 48, 49, 49, 50, 50, 51, 52, 52,
	52, 53, 53, 53, 54, 54, 54, 55, 55, 56,
	56, 57, 57, 58, 58, 59, 60, 60, 61, 61,
	62, 62, 63, 63, 63, 63, 63, 64, 64, 65,
	65, 66, 66, 67, 78, 79, 68,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 4, 12, 3,
	7, 7, 6, 6, 8, 7, 3, 4, 1, 1,
	1, 5, 3, 4, 5, 5, 8, 4, 6, 7,
	4, 5, 4, 5, 5, 0, 2, 0, 2, 1,
	2, 1, 1, 1, 0, 1, 1, 3, 1, 2,
	3, 1, 1, 0, 1, 2, 1, 3, 3, 3,
	3, 5, 0, 1, 2, 1, 1, 2, 3, 2,
	3, 2, 2, 2, 1, 3, 1, 1, 3, 0,
	5, 5, 5, 1, 3, 0, 2, 0, 2, 2,
	0, 2, 1, 3, 3, 2, 3, 3, 3, 4,
	3, 4, 5, 6, 3, 4, 2, 1, 1, 1,
	1, 1, 1, 1, 2, 1, 1, 3, 3, 1,
	3, 1, 3, 1, 1, 1, 3, 3, 3, 3,
	3, 3, 3, 3, 2, 3, 4, 5, 4, 1,
	1, 1, 1, 1, 1, 5, 0, 1, 1, 2,
	4, 0, 2, 1, 3, 1, 1, 1, 1, 0,
	3, 0, 2, 0, 3, 1, 3, 2, 0, 1,
	1, 0, 2, 4, 0, 2, 4, 0, 3, 1,
	3, 0, 5, 1, 3, 3, 0, 2, 0, 3,
	0, 1, 1, 1, 1, 1, 1, 0, 1, 0,
	1, 0, 2, 1, 1, 1, 0,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	-10, -11, -69, -70, -71, -72, -73, -74, 5, 6,
	7, 8, 33, 93, 94, 96, 95, 83, 84, 85,
	87, 89, 88, -14, 49, 50, 51, 52, -12, -77,
	-12, -12, -12, -12, 97, -65, 99, 103, -62, 99,
	101, 97, 97, 98, 99, -12, 90, 91, 92, -67,
	35, -3, 17, -15, 18, -13, 29, -25, 35, 9,
	-58, 86, -59, -41, -67, 35, -61, 102, 98, -67,
	97, -67, 35, -60, 102, -67, -60, 29, -76, 10,
	23, -75, 9, -67, 44, -16, -17, 73, -20, 35,
	-29, -34, -30, 67, -78, -33, -41, -35, -40, -67,
	-38, -42, 20, 44, 36, 37, 38, 25, -39, 71,
	72, 48, 102, 28, 78, 39, -25, 33, 76, -25,
	53, 35, 45, 76, 35, 67, -67, -68, 35, -68,
	100, 35, 20, 64, -67, -25, -29, -34, -34, -78,
	-76, -34, -75, -36, -34, 9, 53, -18, -67, 19,
	76, 65, 66, -31, 21, 67, 23, 24, 22, 68,
	69, 70, 71, 72, 73, 74, 75, 45, 46, 47,
	40, 41, 42, 43, -29, -29, -36, -3, -34, -34,
	-78, -78, -39, -78, -45, -34, -55, 33, -78, -58,
	35, -28, 10, -59, -34, -67, -68, 20, -66, 104,
	-63, 96, 94, 32, 95, 13, 35, 35, 35, -68,
	-55, 33, -76, 105, 53, -21, -22, -24, -78, 35,
	-39, -17, -67, 73, -29, -29, -34, -35, 21, 23,
	24, -34, -34, 25, 67, -34, -34, -34, -34, -34,
	-34, -34, -34, -79, 105, -79, -79, -79, -16, 18,
	-16, -43, -44, 79, -32, 28, -3, -58, -56, -41,
	-28, -49, 13, -29, 64, -67, -68, -64, 100, -32,
	-58, -34, -28, 53, -23, 54, 55, 56, 57, 58,
	60, 61, -19, 35, 19, -22, 76, -35, -34, -34,
	65, 25, -79, -16, -79, -46, -44, 81, -29, -57,
	64, -37, -35, -57, -79, 53, -49, -53, 15, 14,
	35, 35, -47, 11, -22, -22, 54, 59, 54, 59,
	54, 54, 54, -26, 62, 101, 63, 35, -79, 35,
	65, -34, -79, 82, -34, 80, 30, 53, -41, -53,
	-34, -50, -51, -34, -68, -48, 12, 14, 64, 54,
	54, 98, 98, 98, -34, -34, 31, -35, 53, 53,
	-52, 26, 27, -49, -29, -36, -29, -78, -78, -78,
	7, -34, -51, -53, -27, -67, -27, -27, -58, -54,
	16, 34, -79, 53, -79, -79, 7, 21, -67, -67,
	-67,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 16, 45, 45,
	45, 45, 45, 209, 200, 0, 0, 28, 29, 30,
	45, 0, 0, 0, 49, 51, 52, 53, 54, 47,
	0, 0, 0, 0, 198, 0, 0, 210, 0, 0,
	201, 0, 196, 0, 196, 0, 97, 100, 0, 0,
	213, 19, 50, 0, 55, 46, 0, 0, 87, 0,
	26, 0, 193, 0, 163, 213, 0, 0, 0, 216,
	0, 216, 0, 0, 0, 0, 0, 0, 32, 0,
	0, 97, 0, 100, 0, 17, 56, 58, 63, 213,
	61, 62, 102, 0, 0, 133, 134, 135, 0, 163,
	0, 149, 0, 214, 165, 166, 167, 168, 129, 152,
	153, 154, 150, 151, 156, 48, 187, 0, 0, 95,
	0, 27, 0, 0, 216, 0, 211, 37, 0, 40,
	0, 42, 197, 0, 216, 187, 98, 0, 99, 0,
	33, 101, 97, 0, 131, 0, 0, 59, 64, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 117, 118, 119,
	120, 121, 122, 123, 105, 0, 0, 0, 131, 144,
	0, 0, 116, 0, 0, 157, 0, 0, 0, 95,
	88, 173, 0, 194, 195, 164, 35, 199, 0, 0,
	216, 207, 202, 203, 204, 205, 206, 41, 43, 44,
	0, 0, 34, 31, 0, 95, 66, 72, 0, 84,
	86, 57, 65, 60, 103, 104, 107, 108, 0, 0,
	0, 110, 0, 114, 0, 136, 137, 138, 139, 140,
	141, 142, 143, 106, 215, 128, 130, 145, 0, 0,
	0, 161, 158, 0, 191, 0, 125, 191, 0, 189,
	173, 181, 0, 96, 0, 212, 38, 0, 208, 22,
	23, 132, 169, 0, 0, 75, 76, 0, 0, 0,
	0, 0, 89, 73, 0, 0, 0, 109, 111, 0,
	0, 115, 146, 0, 148, 0, 159, 0, 0, 20,
	0, 124, 126, 21, 188, 0, 181, 25, 0, 0,
	216, 39, 171, 0, 67, 70, 77, 0, 79, 0,
	81, 82, 83, 68, 0, 0, 0, 74, 69, 85,
	0, 112, 147, 155, 162, 0, 0, 0, 190, 24,
	182, 174, 175, 178, 36, 173, 0, 0, 0, 78,
	80, 0, 0, 0, 113, 160, 0, 127, 0, 0,
	177, 179, 180, 181, 172, 170, 71, 0, 0, 0,
	0, 183, 176, 184, 0, 93, 0, 0, 192, 18,
	0, 0, 90, 0, 91, 92, 185, 0, 94, 0,
	186,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 75, 68, 3,
	44, 105, 73, 71, 53, 72, 76, 74, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	46, 45, 47, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 70, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 69, 3, 48,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 49, 50, 51, 52, 54, 55, 56, 57,
	58, 59, 60, 61, 62, 63, 64, 65, 66, 67,
	77, 78, 79, 80, 81, 82, 83, 84, 85, 86,
	87, 88, 89, 90, 91, 92, 93, 94, 95, 96,
	97, 98, 99, 100, 101, 102, 103, 104,
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
		//line sql.y:183
		{
			SetParseTree(yylex, yyDollar[1].statement)
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:189
		{
			yyVAL.statement = yyDollar[1].selStmt
		}
	case 17:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:209
		{
			yyVAL.selStmt = &SimpleSelect{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, SelectExprs: yyDollar[4].selectExprs}
		}
	case 18:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line sql.y:213
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, SelectExprs: yyDollar[4].selectExprs, From: yyDollar[6].tableExprs, Where: NewWhere(AST_WHERE, yyDollar[7].boolExpr), GroupBy: GroupBy(yyDollar[8].valExprs), Having: NewWhere(AST_HAVING, yyDollar[9].boolExpr), OrderBy: yyDollar[10].orderBy, Limit: yyDollar[11].limit, Lock: yyDollar[12].str}
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:217
		{
			yyVAL.selStmt = &Union{Type: yyDollar[2].str, Left: yyDollar[1].selStmt, Right: yyDollar[3].selStmt}
		}
	case 20:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:224
		{
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: yyDollar[5].columns, Rows: yyDollar[6].insRows, OnDup: OnDup(yyDollar[7].updateExprs)}
		}
	case 21:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:228
		{
			cols := make(Columns, 0, len(yyDollar[6].updateExprs))
			vals := make(ValTuple, 0, len(yyDollar[6].updateExprs))
			for _, col := range yyDollar[6].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: cols, Rows: Values{vals}, OnDup: OnDup(yyDollar[7].updateExprs)}
		}
	case 22:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:240
		{
			yyVAL.statement = &Replace{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: yyDollar[5].columns, Rows: yyDollar[6].insRows}
		}
	case 23:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:244
		{
			cols := make(Columns, 0, len(yyDollar[6].updateExprs))
			vals := make(ValTuple, 0, len(yyDollar[6].updateExprs))
			for _, col := range yyDollar[6].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Replace{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: cols, Rows: Values{vals}}
		}
	case 24:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:257
		{
			yyVAL.statement = &Update{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[3].tableName, Exprs: yyDollar[5].updateExprs, Where: NewWhere(AST_WHERE, yyDollar[6].boolExpr), OrderBy: yyDollar[7].orderBy, Limit: yyDollar[8].limit}
		}
	case 25:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:263
		{
			yyVAL.statement = &Delete{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Where: NewWhere(AST_WHERE, yyDollar[5].boolExpr), OrderBy: yyDollar[6].orderBy, Limit: yyDollar[7].limit}
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:269
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: yyDollar[3].updateExprs}
		}
	case 27:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:273
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: UpdateExprs{&UpdateExpr{Name: &ColName{Name: []byte("names")}, Expr: StrVal(yyDollar[4].bytes)}}}
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:279
		{
			yyVAL.statement = &Begin{}
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:285
		{
			yyVAL.statement = &Commit{}
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:291
		{
			yyVAL.statement = &Rollback{}
		}
	case 31:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:297
		{
			yyVAL.statement = &Admin{Name: yyDollar[2].bytes, Values: yyDollar[4].valExprs}
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:303
		{
			yyVAL.statement = &Show{Section: "databases", LikeOrWhere: yyDollar[3].expr}
		}
	case 33:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:307
		{
			yyVAL.statement = &Show{Section: "tables", From: yyDollar[3].valExpr, LikeOrWhere: yyDollar[4].expr}
		}
	case 34:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:311
		{
			yyVAL.statement = &Show{Section: "proxy", Key: string(yyDollar[3].bytes), From: yyDollar[4].valExpr, LikeOrWhere: yyDollar[5].expr}
		}
	case 35:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:317
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[4].bytes}
		}
	case 36:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:321
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[7].bytes, NewName: yyDollar[7].bytes}
		}
	case 37:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:326
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[3].bytes}
		}
	case 38:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:332
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[4].bytes, NewName: yyDollar[4].bytes}
		}
	case 39:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:336
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[4].bytes, NewName: yyDollar[7].bytes}
		}
	case 40:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:341
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[3].bytes, NewName: yyDollar[3].bytes}
		}
	case 41:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:347
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[3].bytes, NewName: yyDollar[5].bytes}
		}
	case 42:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:353
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 43:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:357
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[5].bytes, NewName: yyDollar[5].bytes}
		}
	case 44:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:362
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 45:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:367
		{
			SetAllowComments(yylex, true)
		}
	case 46:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:371
		{
			yyVAL.bytes2 = yyDollar[2].bytes2
			SetAllowComments(yylex, false)
		}
	case 47:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:377
		{
			yyVAL.bytes2 = nil
		}
	case 48:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:381
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[2].bytes)
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:387
		{
			yyVAL.str = AST_UNION
		}
	case 50:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:391
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:395
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:399
		{
			yyVAL.str = AST_EXCEPT
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:403
		{
			yyVAL.str = AST_INTERSECT
		}
	case 54:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:408
		{
			yyVAL.str = ""
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:412
		{
			yyVAL.str = AST_DISTINCT
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:418
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:422
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 58:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:428
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 59:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:432
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].bytes}
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:436
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyDollar[1].bytes}
		}
	case 61:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:442
		{
			yyVAL.expr = yyDollar[1].boolExpr
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:446
		{
			yyVAL.expr = yyDollar[1].valExpr
		}
	case 63:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:451
		{
			yyVAL.bytes = nil
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:455
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 65:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:459
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:465
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:469
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:475
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].smTableExpr, As: yyDollar[2].bytes, Hints: yyDollar[3].indexHints}
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:479
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyDollar[2].tableExpr}
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:483
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 71:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:487
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 72:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:492
		{
			yyVAL.bytes = nil
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:496
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 74:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:500
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 75:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:506
		{
			yyVAL.str = AST_JOIN
		}
	case 76:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:510
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 77:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:514
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:518
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 79:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:522
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:526
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 81:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:530
		{
			yyVAL.str = AST_JOIN
		}
	case 82:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:534
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:538
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 84:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:544
		{
			yyVAL.smTableExpr = &TableName{Name: yyDollar[1].bytes}
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:548
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:552
		{
			yyVAL.smTableExpr = yyDollar[1].subquery
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:558
		{
			yyVAL.tableName = &TableName{Name: yyDollar[1].bytes}
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:562
		{
			yyVAL.tableName = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 89:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:567
		{
			yyVAL.indexHints = nil
		}
	case 90:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:571
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyDollar[4].bytes2}
		}
	case 91:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:575
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyDollar[4].bytes2}
		}
	case 92:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:579
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyDollar[4].bytes2}
		}
	case 93:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:585
		{
			yyVAL.bytes2 = [][]byte{yyDollar[1].bytes}
		}
	case 94:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:589
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[3].bytes)
		}
	case 95:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:594
		{
			yyVAL.boolExpr = nil
		}
	case 96:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:598
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 97:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:603
		{
			yyVAL.expr = nil
		}
	case 98:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:607
		{
			yyVAL.expr = yyDollar[2].boolExpr
		}
	case 99:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:611
		{
			yyVAL.expr = yyDollar[2].valExpr
		}
	case 100:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:616
		{
			yyVAL.valExpr = nil
		}
	case 101:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:620
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 103:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:627
		{
			yyVAL.boolExpr = &AndExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 104:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:631
		{
			yyVAL.boolExpr = &OrExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 105:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:635
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyDollar[2].boolExpr}
		}
	case 106:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:639
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyDollar[2].boolExpr}
		}
	case 107:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:645
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 108:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:649
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_IN, Right: yyDollar[3].tuple}
		}
	case 109:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:653
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_IN, Right: yyDollar[4].tuple}
		}
	case 110:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:657
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_LIKE, Right: yyDollar[3].valExpr}
		}
	case 111:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:661
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_LIKE, Right: yyDollar[4].valExpr}
		}
	case 112:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:665
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_BETWEEN, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 113:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:669
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_NOT_BETWEEN, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 114:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:673
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyDollar[1].valExpr}
		}
	case 115:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:677
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyDollar[1].valExpr}
		}
	case 116:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:681
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 117:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:687
		{
			yyVAL.str = AST_EQ
		}
	case 118:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:691
		{
			yyVAL.str = AST_LT
		}
	case 119:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:695
		{
			yyVAL.str = AST_GT
		}
	case 120:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:699
		{
			yyVAL.str = AST_LE
		}
	case 121:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:703
		{
			yyVAL.str = AST_GE
		}
	case 122:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:707
		{
			yyVAL.str = AST_NE
		}
	case 123:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:711
		{
			yyVAL.str = AST_NSE
		}
	case 124:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:717
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:721
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 126:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:727
		{
			yyVAL.values = Values{yyDollar[1].tuple}
		}
	case 127:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:731
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].tuple)
		}
	case 128:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:737
		{
			yyVAL.tuple = ValTuple(yyDollar[2].valExprs)
		}
	case 129:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:741
		{
			yyVAL.tuple = yyDollar[1].subquery
		}
	case 130:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:747
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 131:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:753
		{
			yyVAL.valExprs = ValExprs{yyDollar[1].valExpr}
		}
	case 132:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:757
		{
			yyVAL.valExprs = append(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:763
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 134:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:767
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 135:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:771
		{
			yyVAL.valExpr = yyDollar[1].tuple
		}
	case 136:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:775
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITAND, Right: yyDollar[3].valExpr}
		}
	case 137:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:779
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITOR, Right: yyDollar[3].valExpr}
		}
	case 138:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:783
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITXOR, Right: yyDollar[3].valExpr}
		}
	case 139:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:787
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_PLUS, Right: yyDollar[3].valExpr}
		}
	case 140:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:791
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MINUS, Right: yyDollar[3].valExpr}
		}
	case 141:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:795
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MULT, Right: yyDollar[3].valExpr}
		}
	case 142:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:799
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_DIV, Right: yyDollar[3].valExpr}
		}
	case 143:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:803
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MOD, Right: yyDollar[3].valExpr}
		}
	case 144:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:807
		{
			if num, ok := yyDollar[2].valExpr.(NumVal); ok {
				switch yyDollar[1].byt {
				case '-':
					yyVAL.valExpr = append(NumVal("-"), num...)
				case '+':
					yyVAL.valExpr = num
				default:
					yyVAL.valExpr = &UnaryExpr{Operator: yyDollar[1].byt, Expr: yyDollar[2].valExpr}
				}
			} else {
				yyVAL.valExpr = &UnaryExpr{Operator: yyDollar[1].byt, Expr: yyDollar[2].valExpr}
			}
		}
	case 145:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:822
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes}
		}
	case 146:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:826
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 147:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:830
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 148:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:834
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 149:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:838
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 150:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:844
		{
			yyVAL.bytes = IF_BYTES
		}
	case 151:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:848
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 152:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:854
		{
			yyVAL.byt = AST_UPLUS
		}
	case 153:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:858
		{
			yyVAL.byt = AST_UMINUS
		}
	case 154:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:862
		{
			yyVAL.byt = AST_TILDA
		}
	case 155:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:868
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 156:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:873
		{
			yyVAL.valExpr = nil
		}
	case 157:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:877
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 158:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:883
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 159:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:887
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 160:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:893
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 161:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:898
		{
			yyVAL.valExpr = nil
		}
	case 162:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:902
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 163:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:908
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].bytes}
		}
	case 164:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:912
		{
			yyVAL.colName = &ColName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 165:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:918
		{
			yyVAL.valExpr = StrVal(yyDollar[1].bytes)
		}
	case 166:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:922
		{
			yyVAL.valExpr = NumVal(yyDollar[1].bytes)
		}
	case 167:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:926
		{
			yyVAL.valExpr = ValArg(yyDollar[1].bytes)
		}
	case 168:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:930
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 169:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:935
		{
			yyVAL.valExprs = nil
		}
	case 170:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:939
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 171:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:944
		{
			yyVAL.boolExpr = nil
		}
	case 172:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:948
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 173:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:953
		{
			yyVAL.orderBy = nil
		}
	case 174:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:957
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 175:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:963
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 176:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:967
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 177:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:973
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 178:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:978
		{
			yyVAL.str = AST_ASC
		}
	case 179:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:982
		{
			yyVAL.str = AST_ASC
		}
	case 180:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:986
		{
			yyVAL.str = AST_DESC
		}
	case 181:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:991
		{
			yyVAL.limit = nil
		}
	case 182:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:995
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr}
		}
	case 183:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:999
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 184:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1004
		{
			yyVAL.str = ""
		}
	case 185:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1008
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 186:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1012
		{
			if !bytes.Equal(yyDollar[3].bytes, SHARE) {
				yylex.Error("expecting share")
				return 1
			}
			if !bytes.Equal(yyDollar[4].bytes, MODE) {
				yylex.Error("expecting mode")
				return 1
			}
			yyVAL.str = AST_SHARE_MODE
		}
	case 187:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1025
		{
			yyVAL.columns = nil
		}
	case 188:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1029
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 189:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1035
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyDollar[1].colName}}
		}
	case 190:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1039
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyDollar[3].colName})
		}
	case 191:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1044
		{
			yyVAL.updateExprs = nil
		}
	case 192:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:1048
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 193:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1054
		{
			yyVAL.updateExprs = UpdateExprs{yyDollar[1].updateExpr}
		}
	case 194:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1058
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 195:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1064
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colName, Expr: yyDollar[3].valExpr}
		}
	case 196:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1069
		{
			yyVAL.empty = struct{}{}
		}
	case 197:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1071
		{
			yyVAL.empty = struct{}{}
		}
	case 198:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1074
		{
			yyVAL.empty = struct{}{}
		}
	case 199:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1076
		{
			yyVAL.empty = struct{}{}
		}
	case 200:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1079
		{
			yyVAL.empty = struct{}{}
		}
	case 201:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1081
		{
			yyVAL.empty = struct{}{}
		}
	case 202:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1085
		{
			yyVAL.empty = struct{}{}
		}
	case 203:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1087
		{
			yyVAL.empty = struct{}{}
		}
	case 204:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1089
		{
			yyVAL.empty = struct{}{}
		}
	case 205:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1091
		{
			yyVAL.empty = struct{}{}
		}
	case 206:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1093
		{
			yyVAL.empty = struct{}{}
		}
	case 207:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1096
		{
			yyVAL.empty = struct{}{}
		}
	case 208:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1098
		{
			yyVAL.empty = struct{}{}
		}
	case 209:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1101
		{
			yyVAL.empty = struct{}{}
		}
	case 210:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1103
		{
			yyVAL.empty = struct{}{}
		}
	case 211:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1106
		{
			yyVAL.empty = struct{}{}
		}
	case 212:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1108
		{
			yyVAL.empty = struct{}{}
		}
	case 213:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1112
		{
			yyVAL.bytes = bytes.ToLower(yyDollar[1].bytes)
		}
	case 214:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1118
		{
			if incNesting(yylex) {
				yylex.Error("max nesting level reached")
				return 1
			}
		}
	case 215:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1127
		{
			decNesting(yylex)
		}
	case 216:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1132
		{
			ForceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
