//line sql.y:6
package sqlparser

import __yyfmt__ "fmt"

//line sql.y:6
import "strings"

func setParseTree(yylex interface{}, stmt Statement) {
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

//line sql.y:36
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
	colTuple    ColTuple
	valExprs    ValExprs
	values      Values
	rowTuple    RowTuple
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
	sqlID       SQLName
	sqlIDs      []SQLName
	tableID     TableID
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
const KEYRANGE = 57377
const ID = 57378
const STRING = 57379
const NUMBER = 57380
const VALUE_ARG = 57381
const LIST_ARG = 57382
const COMMENT = 57383
const LE = 57384
const GE = 57385
const NE = 57386
const NULL_SAFE_EQUAL = 57387
const UNION = 57388
const MINUS = 57389
const EXCEPT = 57390
const INTERSECT = 57391
const JOIN = 57392
const STRAIGHT_JOIN = 57393
const LEFT = 57394
const RIGHT = 57395
const INNER = 57396
const OUTER = 57397
const CROSS = 57398
const NATURAL = 57399
const USE = 57400
const FORCE = 57401
const ON = 57402
const OR = 57403
const AND = 57404
const NOT = 57405
const UNARY = 57406
const CASE = 57407
const WHEN = 57408
const THEN = 57409
const ELSE = 57410
const END = 57411
const BEGIN = 57412
const COMMIT = 57413
const ROLLBACK = 57414
const NAMES = 57415
const REPLACE = 57416
const ADMIN = 57417
const DATABASES = 57418
const TABLES = 57419
const PROXY = 57420
const CREATE = 57421
const ALTER = 57422
const DROP = 57423
const RENAME = 57424
const ANALYZE = 57425
const TABLE = 57426
const INDEX = 57427
const VIEW = 57428
const TO = 57429
const IGNORE = 57430
const IF = 57431
const UNIQUE = 57432
const USING = 57433
const SHOW = 57434
const DESCRIBE = 57435
const EXPLAIN = 57436

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
	"KEYRANGE",
	"ID",
	"STRING",
	"NUMBER",
	"VALUE_ARG",
	"LIST_ARG",
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
	"OR",
	"AND",
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
	"DATABASES",
	"TABLES",
	"PROXY",
	"CREATE",
	"ALTER",
	"DROP",
	"RENAME",
	"ANALYZE",
	"TABLE",
	"INDEX",
	"VIEW",
	"TO",
	"IGNORE",
	"IF",
	"UNIQUE",
	"USING",
	"SHOW",
	"DESCRIBE",
	"EXPLAIN",
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
	-1, 85,
	78, 222,
	-2, 221,
}

const yyNprod = 226
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 786

var yyAct = [...]int{

	120, 335, 157, 405, 268, 198, 287, 117, 118, 80,
	116, 238, 327, 278, 372, 249, 213, 280, 81, 39,
	40, 41, 42, 107, 106, 173, 172, 414, 197, 3,
	167, 236, 98, 208, 269, 221, 76, 333, 68, 301,
	302, 303, 304, 305, 236, 306, 307, 94, 49, 83,
	51, 111, 89, 101, 52, 91, 82, 352, 354, 87,
	54, 90, 55, 294, 103, 151, 383, 70, 269, 130,
	382, 381, 123, 88, 363, 112, 60, 129, 269, 56,
	135, 279, 269, 325, 279, 269, 269, 124, 85, 126,
	127, 128, 269, 269, 173, 172, 353, 155, 125, 235,
	312, 171, 133, 158, 64, 161, 144, 140, 165, 365,
	169, 164, 138, 57, 58, 59, 141, 146, 199, 69,
	258, 114, 200, 173, 172, 131, 132, 108, 61, 62,
	63, 172, 136, 142, 160, 378, 65, 66, 328, 207,
	83, 163, 290, 83, 328, 217, 216, 82, 154, 211,
	82, 156, 185, 186, 187, 346, 134, 162, 245, 380,
	347, 215, 165, 344, 259, 379, 195, 196, 345, 350,
	112, 244, 217, 349, 93, 348, 214, 248, 214, 142,
	256, 257, 159, 260, 261, 262, 263, 264, 265, 266,
	267, 243, 166, 203, 148, 232, 150, 233, 39, 40,
	41, 42, 270, 271, 112, 112, 272, 236, 390, 367,
	83, 83, 322, 125, 252, 105, 276, 82, 285, 283,
	125, 299, 291, 142, 246, 247, 273, 275, 286, 183,
	184, 185, 186, 187, 96, 83, 242, 282, 167, 297,
	20, 143, 82, 137, 296, 251, 69, 310, 234, 78,
	218, 295, 311, 165, 298, 115, 313, 315, 316, 125,
	231, 125, 282, 209, 78, 85, 289, 78, 147, 314,
	170, 78, 129, 95, 411, 139, 125, 112, 319, 387,
	321, 125, 115, 366, 126, 127, 128, 69, 104, 332,
	330, 324, 412, 334, 20, 75, 331, 115, 115, 320,
	318, 99, 418, 201, 202, 219, 204, 205, 153, 71,
	242, 342, 343, 73, 100, 336, 356, 281, 358, 377,
	360, 210, 337, 251, 227, 361, 288, 292, 364, 43,
	376, 326, 341, 362, 83, 253, 369, 254, 255, 370,
	373, 368, 214, 225, 102, 79, 210, 228, 417, 240,
	115, 45, 46, 47, 48, 115, 115, 401, 250, 20,
	44, 19, 384, 18, 17, 16, 67, 385, 386, 242,
	242, 301, 302, 303, 304, 305, 388, 306, 307, 15,
	165, 14, 394, 396, 115, 115, 308, 84, 168, 220,
	50, 293, 402, 373, 222, 53, 404, 115, 86, 406,
	406, 406, 83, 407, 408, 403, 224, 226, 223, 82,
	413, 409, 415, 416, 284, 419, 410, 391, 371, 420,
	375, 421, 340, 240, 323, 206, 277, 122, 395, 119,
	397, 121, 329, 174, 77, 113, 250, 351, 239, 300,
	237, 109, 72, 374, 92, 38, 74, 359, 97, 180,
	181, 182, 183, 184, 185, 186, 187, 115, 13, 12,
	110, 11, 115, 77, 392, 393, 10, 77, 9, 8,
	274, 7, 123, 6, 145, 5, 4, 129, 149, 2,
	135, 152, 240, 240, 1, 0, 0, 124, 85, 126,
	127, 128, 77, 0, 20, 21, 22, 23, 125, 0,
	0, 0, 133, 0, 0, 0, 0, 0, 180, 181,
	182, 183, 184, 185, 186, 187, 0, 0, 0, 0,
	0, 114, 24, 0, 0, 131, 132, 108, 212, 389,
	0, 0, 136, 0, 0, 0, 0, 0, 0, 229,
	0, 0, 230, 0, 180, 181, 182, 183, 184, 185,
	186, 187, 0, 0, 241, 110, 134, 0, 0, 115,
	0, 115, 269, 0, 398, 399, 400, 0, 0, 0,
	0, 0, 0, 0, 33, 34, 35, 0, 36, 37,
	20, 0, 0, 25, 26, 28, 27, 29, 0, 110,
	110, 0, 0, 0, 0, 123, 30, 31, 32, 0,
	129, 0, 0, 135, 0, 0, 0, 0, 0, 0,
	124, 85, 126, 127, 128, 123, 0, 0, 0, 0,
	129, 125, 0, 135, 0, 133, 0, 309, 241, 0,
	124, 85, 126, 127, 128, 0, 0, 0, 0, 0,
	0, 125, 0, 0, 114, 133, 0, 0, 131, 132,
	0, 20, 0, 0, 0, 136, 0, 0, 0, 0,
	0, 0, 110, 0, 114, 0, 0, 0, 131, 132,
	0, 129, 0, 0, 135, 136, 0, 0, 338, 134,
	0, 339, 85, 126, 127, 128, 0, 241, 241, 0,
	129, 0, 125, 135, 0, 0, 133, 0, 355, 134,
	357, 85, 126, 127, 128, 0, 175, 179, 177, 178,
	0, 125, 0, 0, 0, 133, 0, 0, 0, 131,
	132, 0, 0, 0, 0, 0, 136, 191, 192, 193,
	194, 0, 188, 189, 190, 0, 0, 0, 131, 132,
	0, 0, 0, 0, 0, 136, 0, 0, 0, 0,
	134, 0, 0, 0, 176, 180, 181, 182, 183, 184,
	185, 186, 187, 0, 0, 0, 0, 0, 317, 134,
	180, 181, 182, 183, 184, 185, 186, 187, 180, 181,
	182, 183, 184, 185, 186, 187,
}
var yyPact = [...]int{

	489, -1000, -1000, 147, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -51, -41, -20, 14, -23,
	37, -1000, -1000, -1000, -1000, -1000, -1000, 210, 354, 292,
	-1000, -1000, -1000, 295, -1000, 266, 231, 336, 229, -45,
	-27, 210, -1000, -38, 210, -1000, 231, -57, 237, -57,
	231, 291, 335, 210, -1000, -1000, -1000, 259, 169, -1000,
	-1000, -1000, 52, -1000, 202, 231, 242, 29, -1000, 231,
	124, -1000, 194, -1000, 28, -1000, 231, 48, 232, -1000,
	231, -1000, -37, 231, 288, 82, 210, -1000, -1000, 595,
	665, 291, 665, 335, 231, 665, 183, -1000, -1000, 251,
	23, 56, 685, -1000, 595, 575, -1000, -1000, -1000, 665,
	167, 167, -1000, 167, 167, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 665, -1000, 230, 229,
	231, 332, 229, 665, 210, -1000, 285, -71, -1000, 311,
	-1000, 231, -1000, -1000, 231, -1000, 56, 685, 708, 646,
	-1000, 708, 291, 215, -11, 708, 213, 52, -1000, -1000,
	210, 83, 595, 595, 665, 174, 314, 665, 665, 95,
	665, 665, 665, 665, 665, 665, 665, 665, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -42, -32, -24, 685,
	-1000, 452, 52, -1000, 354, 247, 3, 708, 289, 229,
	229, 168, -1000, 313, 595, -1000, 708, -1000, -1000, -1000,
	76, 210, -1000, -39, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 289, 229, -1000, 665, 166, 315, 228,
	235, 22, -1000, -1000, -1000, -1000, -1000, 63, 708, -1000,
	646, -1000, -1000, 174, 665, 665, 708, 700, -1000, 275,
	156, 156, 156, 77, 77, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -25, 52, -25, 157, 0, -1000, 595,
	72, 167, 147, 78, -18, -1000, 313, 300, 308, 56,
	231, -1000, -1000, 231, -1000, -1000, 124, 708, 321, 213,
	213, -1000, -1000, 107, 99, 119, 117, 113, -7, -1000,
	231, -17, 231, -24, -1000, 708, 379, 665, -1000, -1000,
	-25, -1000, 247, -10, -1000, 665, 27, -1000, 253, 154,
	-1000, -1000, -1000, 229, 300, -1000, 665, 665, -1000, -1000,
	318, 305, 315, 69, -1000, 109, -1000, 103, -1000, -1000,
	-1000, -1000, -29, -30, -34, -1000, -1000, -1000, -1000, 665,
	708, -1000, -76, -1000, 708, 665, 248, 167, -1000, -1000,
	474, 153, -1000, 438, -1000, 313, 595, 665, 595, -1000,
	-1000, 167, 167, 167, 708, -1000, 708, 350, -1000, 665,
	665, -1000, -1000, -1000, 300, 56, 152, 56, 210, 210,
	210, 229, 708, -1000, 258, -28, -1000, -28, -28, 124,
	-1000, 341, 281, -1000, 210, -1000, -1000, -1000, 210, -1000,
	210, -1000,
}
var yyPgo = [...]int{

	0, 484, 479, 28, 476, 475, 473, 471, 469, 468,
	466, 461, 459, 458, 329, 446, 445, 442, 24, 23,
	441, 440, 11, 439, 438, 36, 437, 3, 16, 51,
	435, 433, 17, 10, 2, 15, 5, 432, 8, 431,
	69, 429, 7, 427, 426, 13, 425, 424, 422, 420,
	6, 418, 14, 417, 1, 416, 33, 414, 12, 9,
	18, 174, 398, 395, 394, 391, 390, 389, 0, 388,
	387, 386, 104, 381, 379, 365, 364, 363, 361, 53,
	32, 360, 182, 4,
}
var yyR1 = [...]int{

	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 3,
	3, 4, 4, 76, 76, 5, 6, 7, 73, 74,
	75, 78, 77, 77, 77, 8, 8, 8, 9, 9,
	9, 10, 11, 11, 11, 12, 13, 13, 13, 81,
	14, 15, 15, 16, 16, 16, 16, 16, 17, 17,
	18, 18, 19, 19, 19, 20, 20, 69, 69, 69,
	21, 21, 22, 22, 22, 22, 71, 71, 71, 23,
	23, 23, 23, 23, 23, 23, 23, 23, 24, 24,
	24, 25, 25, 26, 26, 26, 26, 27, 27, 28,
	28, 80, 80, 80, 79, 79, 29, 29, 29, 29,
	29, 30, 30, 30, 30, 30, 30, 30, 30, 30,
	30, 30, 31, 31, 31, 31, 31, 31, 31, 35,
	35, 35, 40, 36, 36, 34, 34, 34, 34, 34,
	34, 34, 34, 34, 34, 34, 34, 34, 34, 34,
	34, 34, 39, 39, 41, 41, 41, 43, 46, 46,
	44, 44, 45, 47, 47, 42, 42, 33, 33, 33,
	33, 48, 48, 49, 49, 50, 50, 51, 51, 52,
	53, 53, 53, 54, 54, 54, 55, 55, 55, 56,
	56, 57, 57, 58, 58, 32, 32, 37, 37, 38,
	38, 59, 59, 60, 61, 61, 62, 62, 63, 63,
	64, 64, 64, 64, 64, 65, 65, 66, 66, 67,
	67, 68, 70, 82, 83, 72,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 12,
	3, 7, 7, 6, 6, 8, 7, 3, 1, 1,
	1, 5, 3, 4, 5, 5, 8, 4, 6, 7,
	4, 5, 4, 5, 5, 3, 2, 2, 2, 0,
	2, 0, 2, 1, 2, 1, 1, 1, 0, 1,
	1, 3, 1, 2, 3, 1, 1, 0, 1, 2,
	1, 3, 3, 3, 3, 5, 0, 1, 2, 1,
	1, 2, 3, 2, 3, 2, 2, 2, 1, 3,
	1, 1, 3, 0, 5, 5, 5, 1, 3, 0,
	2, 0, 2, 2, 0, 2, 1, 3, 3, 2,
	3, 3, 3, 4, 3, 4, 5, 6, 3, 4,
	2, 6, 1, 1, 1, 1, 1, 1, 1, 3,
	1, 1, 3, 1, 3, 1, 1, 1, 3, 3,
	3, 3, 3, 3, 3, 3, 2, 3, 4, 5,
	4, 1, 1, 1, 1, 1, 1, 5, 0, 1,
	1, 2, 4, 0, 2, 1, 3, 1, 1, 1,
	1, 0, 3, 0, 2, 0, 3, 1, 3, 2,
	0, 1, 1, 0, 2, 4, 0, 2, 4, 0,
	3, 1, 3, 0, 5, 2, 1, 1, 3, 3,
	1, 1, 3, 3, 0, 2, 0, 3, 0, 1,
	1, 1, 1, 1, 1, 0, 1, 0, 1, 0,
	2, 1, 1, 1, 1, 0,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	-10, -11, -12, -13, -73, -74, -75, -76, -77, -78,
	5, 6, 7, 8, 33, 94, 95, 97, 96, 98,
	107, 108, 109, 85, 86, 87, 89, 90, -16, 51,
	52, 53, 54, -14, -81, -14, -14, -14, -14, 99,
	-66, 101, 105, -63, 101, 103, 99, 99, 100, 101,
	99, 91, 92, 93, -72, -72, -72, -14, -68, 36,
	-3, 17, -17, 18, -15, 29, -25, -70, 36, 9,
	-59, -60, -42, -68, -70, 36, -62, 104, 100, -68,
	99, -68, -70, -61, 104, 36, -61, -70, -80, 10,
	23, -79, 9, -68, 29, 46, -18, -19, 75, -20,
	-70, -29, -34, -30, 69, -82, -33, -42, -38, -41,
	-68, -39, -43, 20, 35, 46, 37, 38, 39, 25,
	-40, 73, 74, 50, 104, 28, 80, 41, -25, 33,
	78, -25, 55, 47, 78, -70, 69, 36, -72, -70,
	-72, 102, -70, 20, 66, -68, -29, -34, -34, -82,
	-80, -34, -79, -25, -36, -34, 9, 55, -69, -68,
	19, 78, 68, 67, -31, 21, 69, 23, 24, 22,
	70, 71, 72, 73, 74, 75, 76, 77, 47, 48,
	49, 42, 43, 44, 45, -29, -29, -3, -36, -34,
	-34, -82, -82, -40, -82, -82, -46, -34, -56, 33,
	-82, -59, -70, -28, 10, -60, -34, -68, -72, 20,
	-67, 106, -64, 97, 95, 32, 96, 13, 36, -70,
	-70, -72, -80, -56, 33, 110, 55, -21, -22, -24,
	-82, -70, -40, -19, -68, 75, -29, -29, -34, -35,
	-82, -40, 40, 21, 23, 24, -34, -34, 25, 69,
	-34, -34, -34, -34, -34, -34, -34, -34, -83, 110,
	-83, -83, -83, -18, 18, -18, -33, -44, -45, 81,
	-32, 28, -3, -59, -57, -42, -28, -50, 13, -29,
	66, -68, -72, -65, 102, -32, -59, -34, -28, 55,
	-23, 56, 57, 58, 59, 60, 62, 63, -71, -70,
	19, -22, 78, -36, -35, -34, -34, 68, 25, -83,
	-18, -83, 55, -47, -45, 83, -29, -58, 66, -37,
	-38, -58, -83, 55, -50, -54, 15, 14, -70, -70,
	-48, 11, -22, -22, 56, 61, 56, 61, 56, 56,
	56, -26, 64, 103, 65, -70, -83, -70, -83, 68,
	-34, -83, -33, 84, -34, 82, 30, 55, -42, -54,
	-34, -51, -52, -34, -72, -49, 12, 14, 66, 56,
	56, 100, 100, 100, -34, -83, -34, 31, -38, 55,
	55, -53, 26, 27, -50, -29, -36, -29, -82, -82,
	-82, 7, -34, -52, -54, -27, -68, -27, -27, -59,
	-55, 16, 34, -83, 55, -83, -83, 7, 21, -68,
	-68, -68,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 16, 17, 18,
	49, 49, 49, 49, 49, 217, 208, 0, 0, 0,
	225, 225, 225, 28, 29, 30, 49, 0, 0, 53,
	55, 56, 57, 58, 51, 0, 0, 0, 0, 206,
	0, 0, 218, 0, 0, 209, 0, 204, 0, 204,
	0, 101, 104, 0, 46, 47, 48, 0, 0, 221,
	20, 54, 0, 59, 50, 0, 0, 91, 222, 0,
	27, 201, 0, 165, 0, -2, 0, 0, 0, 225,
	0, 225, 0, 0, 0, 0, 0, 45, 32, 0,
	0, 101, 0, 104, 0, 0, 0, 60, 62, 67,
	0, 65, 66, 106, 0, 0, 135, 136, 137, 0,
	165, 0, 151, 0, 0, 223, 167, 168, 169, 170,
	200, 154, 155, 156, 152, 153, 158, 52, 189, 0,
	0, 99, 0, 0, 0, 225, 0, 219, 37, 0,
	40, 0, 42, 205, 0, 225, 102, 0, 103, 0,
	33, 105, 101, 189, 0, 133, 0, 0, 63, 68,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 122, 123,
	124, 125, 126, 127, 128, 109, 0, 0, 0, 133,
	146, 0, 0, 120, 0, 0, 0, 159, 0, 0,
	0, 99, 92, 175, 0, 202, 203, 166, 35, 207,
	0, 0, 225, 215, 210, 211, 212, 213, 214, 41,
	43, 44, 34, 0, 0, 31, 0, 99, 70, 76,
	0, 88, 90, 61, 69, 64, 107, 108, 111, 112,
	0, 130, 131, 0, 0, 0, 114, 0, 118, 0,
	138, 139, 140, 141, 142, 143, 144, 145, 110, 224,
	132, 199, 147, 0, 0, 0, 0, 163, 160, 0,
	193, 0, 196, 193, 0, 191, 175, 183, 0, 100,
	0, 220, 38, 0, 216, 23, 24, 134, 171, 0,
	0, 79, 80, 0, 0, 0, 0, 0, 93, 77,
	0, 0, 0, 0, 113, 115, 0, 0, 119, 148,
	0, 150, 0, 0, 161, 0, 0, 21, 0, 195,
	197, 22, 190, 0, 183, 26, 0, 0, 225, 39,
	173, 0, 71, 74, 81, 0, 83, 0, 85, 86,
	87, 72, 0, 0, 0, 78, 73, 89, 129, 0,
	116, 149, 0, 157, 164, 0, 0, 0, 192, 25,
	184, 176, 177, 180, 36, 175, 0, 0, 0, 82,
	84, 0, 0, 0, 117, 121, 162, 0, 198, 0,
	0, 179, 181, 182, 183, 174, 172, 75, 0, 0,
	0, 0, 185, 178, 186, 0, 97, 0, 0, 194,
	19, 0, 0, 94, 0, 95, 96, 187, 0, 98,
	0, 188,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 77, 70, 3,
	46, 110, 75, 73, 55, 74, 78, 76, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	48, 47, 49, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 72, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 71, 3, 50,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 51, 52, 53, 54, 56, 57,
	58, 59, 60, 61, 62, 63, 64, 65, 66, 67,
	68, 69, 79, 80, 81, 82, 83, 84, 85, 86,
	87, 88, 89, 90, 91, 92, 93, 94, 95, 96,
	97, 98, 99, 100, 101, 102, 103, 104, 105, 106,
	107, 108, 109,
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
		//line sql.y:182
		{
			setParseTree(yylex, yyDollar[1].statement)
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:188
		{
			yyVAL.statement = yyDollar[1].selStmt
		}
	case 19:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line sql.y:210
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, SelectExprs: yyDollar[4].selectExprs, From: yyDollar[6].tableExprs, Where: NewWhere(AST_WHERE, yyDollar[7].boolExpr), GroupBy: GroupBy(yyDollar[8].valExprs), Having: NewWhere(AST_HAVING, yyDollar[9].boolExpr), OrderBy: yyDollar[10].orderBy, Limit: yyDollar[11].limit, Lock: yyDollar[12].str}
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:214
		{
			yyVAL.selStmt = &Union{Type: yyDollar[2].str, Left: yyDollar[1].selStmt, Right: yyDollar[3].selStmt}
		}
	case 21:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:220
		{
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: yyDollar[5].columns, Rows: yyDollar[6].insRows, OnDup: OnDup(yyDollar[7].updateExprs)}
		}
	case 22:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:224
		{
			cols := make(Columns, 0, len(yyDollar[6].updateExprs))
			vals := make(ValTuple, 0, len(yyDollar[6].updateExprs))
			for _, col := range yyDollar[6].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: cols, Rows: Values{vals}, OnDup: OnDup(yyDollar[7].updateExprs)}
		}
	case 23:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:236
		{
			yyVAL.statement = &Replace{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: yyDollar[5].columns, Rows: yyDollar[6].insRows}
		}
	case 24:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:240
		{
			cols := make(Columns, 0, len(yyDollar[6].updateExprs))
			vals := make(ValTuple, 0, len(yyDollar[6].updateExprs))
			for _, col := range yyDollar[6].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Replace{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: cols, Rows: Values{vals}}
		}
	case 25:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:253
		{
			yyVAL.statement = &Update{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[3].tableName, Exprs: yyDollar[5].updateExprs, Where: NewWhere(AST_WHERE, yyDollar[6].boolExpr), OrderBy: yyDollar[7].orderBy, Limit: yyDollar[8].limit}
		}
	case 26:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:259
		{
			yyVAL.statement = &Delete{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Where: NewWhere(AST_WHERE, yyDollar[5].boolExpr), OrderBy: yyDollar[6].orderBy, Limit: yyDollar[7].limit}
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:265
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: yyDollar[3].updateExprs}
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:271
		{
			yyVAL.statement = &Begin{}
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:277
		{
			yyVAL.statement = &Commit{}
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:283
		{
			yyVAL.statement = &Rollback{}
		}
	case 31:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:289
		{
			yyVAL.statement = &Admin{Name: yyDollar[2].sqlID, Values: yyDollar[4].valExprs}
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:295
		{
			yyVAL.statement = &Show{Section: "databases", LikeOrWhere: yyDollar[3].expr}
		}
	case 33:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:299
		{
			yyVAL.statement = &Show{Section: "tables", From: yyDollar[3].valExpr, LikeOrWhere: yyDollar[4].expr}
		}
	case 34:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:303
		{
			yyVAL.statement = &Show{Section: "proxy", Key: string(yyDollar[3].sqlID), From: yyDollar[4].valExpr, LikeOrWhere: yyDollar[5].expr}
		}
	case 35:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:309
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[4].tableID}
		}
	case 36:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:313
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[7].tableID, NewName: yyDollar[7].tableID}
		}
	case 37:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:318
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: TableID(yyDollar[3].sqlID)}
		}
	case 38:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:324
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[4].tableID, NewName: yyDollar[4].tableID}
		}
	case 39:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:328
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[4].tableID, NewName: yyDollar[7].tableID}
		}
	case 40:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:333
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: TableID(yyDollar[3].sqlID), NewName: TableID(yyDollar[3].sqlID)}
		}
	case 41:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:339
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[3].tableID, NewName: yyDollar[5].tableID}
		}
	case 42:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:345
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].tableID}
		}
	case 43:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:349
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[5].tableID, NewName: yyDollar[5].tableID}
		}
	case 44:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:354
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: TableID(yyDollar[4].sqlID)}
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:360
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[3].tableID, NewName: yyDollar[3].tableID}
		}
	case 46:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:366
		{
			yyVAL.statement = &Other{}
		}
	case 47:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:370
		{
			yyVAL.statement = &Other{}
		}
	case 48:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:374
		{
			yyVAL.statement = &Other{}
		}
	case 49:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:379
		{
			setAllowComments(yylex, true)
		}
	case 50:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:383
		{
			yyVAL.bytes2 = yyDollar[2].bytes2
			setAllowComments(yylex, false)
		}
	case 51:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:389
		{
			yyVAL.bytes2 = nil
		}
	case 52:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:393
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[2].bytes)
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:399
		{
			yyVAL.str = AST_UNION
		}
	case 54:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:403
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:407
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:411
		{
			yyVAL.str = AST_EXCEPT
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:415
		{
			yyVAL.str = AST_INTERSECT
		}
	case 58:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:420
		{
			yyVAL.str = ""
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:424
		{
			yyVAL.str = AST_DISTINCT
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:430
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:434
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:440
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 63:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:444
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].sqlID}
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:448
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyDollar[1].tableID}
		}
	case 65:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:454
		{
			yyVAL.expr = yyDollar[1].boolExpr
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:458
		{
			yyVAL.expr = yyDollar[1].valExpr
		}
	case 67:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:463
		{
			yyVAL.sqlID = ""
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:467
		{
			yyVAL.sqlID = yyDollar[1].sqlID
		}
	case 69:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:471
		{
			yyVAL.sqlID = yyDollar[2].sqlID
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:477
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:481
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:487
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].smTableExpr, As: yyDollar[2].tableID, Hints: yyDollar[3].indexHints}
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:491
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyDollar[2].tableExpr}
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:495
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 75:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:499
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 76:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:504
		{
			yyVAL.tableID = ""
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:508
		{
			yyVAL.tableID = yyDollar[1].tableID
		}
	case 78:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:512
		{
			yyVAL.tableID = yyDollar[2].tableID
		}
	case 79:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:518
		{
			yyVAL.str = AST_JOIN
		}
	case 80:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:522
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 81:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:526
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:530
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:534
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:538
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 85:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:542
		{
			yyVAL.str = AST_JOIN
		}
	case 86:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:546
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 87:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:550
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 88:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:556
		{
			yyVAL.smTableExpr = &TableName{Name: yyDollar[1].tableID}
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:560
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyDollar[1].tableID, Name: yyDollar[3].tableID}
		}
	case 90:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:564
		{
			yyVAL.smTableExpr = yyDollar[1].subquery
		}
	case 91:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:570
		{
			yyVAL.tableName = &TableName{Name: yyDollar[1].tableID}
		}
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:574
		{
			yyVAL.tableName = &TableName{Qualifier: yyDollar[1].tableID, Name: yyDollar[3].tableID}
		}
	case 93:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:579
		{
			yyVAL.indexHints = nil
		}
	case 94:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:583
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyDollar[4].sqlIDs}
		}
	case 95:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:587
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyDollar[4].sqlIDs}
		}
	case 96:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:591
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyDollar[4].sqlIDs}
		}
	case 97:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:597
		{
			yyVAL.sqlIDs = []SQLName{yyDollar[1].sqlID}
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:601
		{
			yyVAL.sqlIDs = append(yyDollar[1].sqlIDs, yyDollar[3].sqlID)
		}
	case 99:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:606
		{
			yyVAL.boolExpr = nil
		}
	case 100:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:610
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 101:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:615
		{
			yyVAL.expr = nil
		}
	case 102:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:619
		{
			yyVAL.expr = yyDollar[2].boolExpr
		}
	case 103:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:623
		{
			yyVAL.expr = yyDollar[2].valExpr
		}
	case 104:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:628
		{
			yyVAL.valExpr = nil
		}
	case 105:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:632
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 107:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:639
		{
			yyVAL.boolExpr = &AndExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 108:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:643
		{
			yyVAL.boolExpr = &OrExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 109:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:647
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyDollar[2].boolExpr}
		}
	case 110:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:651
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyDollar[2].boolExpr}
		}
	case 111:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:657
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 112:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:661
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_IN, Right: yyDollar[3].colTuple}
		}
	case 113:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:665
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_IN, Right: yyDollar[4].colTuple}
		}
	case 114:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:669
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_LIKE, Right: yyDollar[3].valExpr}
		}
	case 115:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:673
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_LIKE, Right: yyDollar[4].valExpr}
		}
	case 116:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:677
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_BETWEEN, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 117:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:681
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_NOT_BETWEEN, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 118:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:685
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyDollar[1].valExpr}
		}
	case 119:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:689
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyDollar[1].valExpr}
		}
	case 120:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:693
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 121:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:697
		{
			yyVAL.boolExpr = &KeyrangeExpr{Start: yyDollar[3].valExpr, End: yyDollar[5].valExpr}
		}
	case 122:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:703
		{
			yyVAL.str = AST_EQ
		}
	case 123:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:707
		{
			yyVAL.str = AST_LT
		}
	case 124:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:711
		{
			yyVAL.str = AST_GT
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:715
		{
			yyVAL.str = AST_LE
		}
	case 126:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:719
		{
			yyVAL.str = AST_GE
		}
	case 127:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:723
		{
			yyVAL.str = AST_NE
		}
	case 128:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:727
		{
			yyVAL.str = AST_NSE
		}
	case 129:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:733
		{
			yyVAL.colTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 130:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:737
		{
			yyVAL.colTuple = yyDollar[1].subquery
		}
	case 131:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:741
		{
			yyVAL.colTuple = ListArg(yyDollar[1].bytes)
		}
	case 132:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:747
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:753
		{
			yyVAL.valExprs = ValExprs{yyDollar[1].valExpr}
		}
	case 134:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:757
		{
			yyVAL.valExprs = append(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 135:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:763
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 136:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:767
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 137:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:771
		{
			yyVAL.valExpr = yyDollar[1].rowTuple
		}
	case 138:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:775
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITAND, Right: yyDollar[3].valExpr}
		}
	case 139:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:779
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITOR, Right: yyDollar[3].valExpr}
		}
	case 140:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:783
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITXOR, Right: yyDollar[3].valExpr}
		}
	case 141:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:787
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_PLUS, Right: yyDollar[3].valExpr}
		}
	case 142:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:791
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MINUS, Right: yyDollar[3].valExpr}
		}
	case 143:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:795
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MULT, Right: yyDollar[3].valExpr}
		}
	case 144:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:799
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_DIV, Right: yyDollar[3].valExpr}
		}
	case 145:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:803
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MOD, Right: yyDollar[3].valExpr}
		}
	case 146:
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
	case 147:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:822
		{
			yyVAL.valExpr = &FuncExpr{Name: string(yyDollar[1].sqlID)}
		}
	case 148:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:826
		{
			yyVAL.valExpr = &FuncExpr{Name: string(yyDollar[1].sqlID), Exprs: yyDollar[3].selectExprs}
		}
	case 149:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:830
		{
			yyVAL.valExpr = &FuncExpr{Name: string(yyDollar[1].sqlID), Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 150:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:834
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].str, Exprs: yyDollar[3].selectExprs}
		}
	case 151:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:838
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 152:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:844
		{
			yyVAL.str = "if"
		}
	case 153:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:848
		{
			yyVAL.str = "values"
		}
	case 154:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:854
		{
			yyVAL.byt = AST_UPLUS
		}
	case 155:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:858
		{
			yyVAL.byt = AST_UMINUS
		}
	case 156:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:862
		{
			yyVAL.byt = AST_TILDA
		}
	case 157:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:868
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 158:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:873
		{
			yyVAL.valExpr = nil
		}
	case 159:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:877
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 160:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:883
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 161:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:887
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 162:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:893
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 163:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:898
		{
			yyVAL.valExpr = nil
		}
	case 164:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:902
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 165:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:908
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].sqlID}
		}
	case 166:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:912
		{
			yyVAL.colName = &ColName{Qualifier: yyDollar[1].tableID, Name: yyDollar[3].sqlID}
		}
	case 167:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:918
		{
			yyVAL.valExpr = StrVal(yyDollar[1].bytes)
		}
	case 168:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:922
		{
			yyVAL.valExpr = NumVal(yyDollar[1].bytes)
		}
	case 169:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:926
		{
			yyVAL.valExpr = ValArg(yyDollar[1].bytes)
		}
	case 170:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:930
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 171:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:935
		{
			yyVAL.valExprs = nil
		}
	case 172:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:939
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 173:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:944
		{
			yyVAL.boolExpr = nil
		}
	case 174:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:948
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 175:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:953
		{
			yyVAL.orderBy = nil
		}
	case 176:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:957
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 177:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:963
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 178:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:967
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 179:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:973
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 180:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:978
		{
			yyVAL.str = AST_ASC
		}
	case 181:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:982
		{
			yyVAL.str = AST_ASC
		}
	case 182:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:986
		{
			yyVAL.str = AST_DESC
		}
	case 183:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:991
		{
			yyVAL.limit = nil
		}
	case 184:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:995
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr}
		}
	case 185:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:999
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 186:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1004
		{
			yyVAL.str = ""
		}
	case 187:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1008
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 188:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1012
		{
			if yyDollar[3].sqlID != "share" {
				yylex.Error("expecting share")
				return 1
			}
			if yyDollar[4].sqlID != "mode" {
				yylex.Error("expecting mode")
				return 1
			}
			yyVAL.str = AST_SHARE_MODE
		}
	case 189:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1025
		{
			yyVAL.columns = nil
		}
	case 190:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1029
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 191:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1035
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyDollar[1].colName}}
		}
	case 192:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1039
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyDollar[3].colName})
		}
	case 193:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1044
		{
			yyVAL.updateExprs = nil
		}
	case 194:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:1048
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 195:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1054
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 196:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1058
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 197:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1064
		{
			yyVAL.values = Values{yyDollar[1].rowTuple}
		}
	case 198:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1068
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].rowTuple)
		}
	case 199:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1074
		{
			yyVAL.rowTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 200:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1078
		{
			yyVAL.rowTuple = yyDollar[1].subquery
		}
	case 201:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1084
		{
			yyVAL.updateExprs = UpdateExprs{yyDollar[1].updateExpr}
		}
	case 202:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1088
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 203:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1094
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colName, Expr: yyDollar[3].valExpr}
		}
	case 204:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1099
		{
			yyVAL.empty = struct{}{}
		}
	case 205:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1101
		{
			yyVAL.empty = struct{}{}
		}
	case 206:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1104
		{
			yyVAL.empty = struct{}{}
		}
	case 207:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1106
		{
			yyVAL.empty = struct{}{}
		}
	case 208:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1109
		{
			yyVAL.empty = struct{}{}
		}
	case 209:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1111
		{
			yyVAL.empty = struct{}{}
		}
	case 210:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1115
		{
			yyVAL.empty = struct{}{}
		}
	case 211:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1117
		{
			yyVAL.empty = struct{}{}
		}
	case 212:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1119
		{
			yyVAL.empty = struct{}{}
		}
	case 213:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1121
		{
			yyVAL.empty = struct{}{}
		}
	case 214:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1123
		{
			yyVAL.empty = struct{}{}
		}
	case 215:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1126
		{
			yyVAL.empty = struct{}{}
		}
	case 216:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1128
		{
			yyVAL.empty = struct{}{}
		}
	case 217:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1131
		{
			yyVAL.empty = struct{}{}
		}
	case 218:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1133
		{
			yyVAL.empty = struct{}{}
		}
	case 219:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1136
		{
			yyVAL.empty = struct{}{}
		}
	case 220:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1138
		{
			yyVAL.empty = struct{}{}
		}
	case 221:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1142
		{
			yyVAL.sqlID = SQLName(strings.ToLower(string(yyDollar[1].bytes)))
		}
	case 222:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1148
		{
			yyVAL.tableID = TableID(yyDollar[1].bytes)
		}
	case 223:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1154
		{
			if incNesting(yylex) {
				yylex.Error("max nesting level reached")
				return 1
			}
		}
	case 224:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1163
		{
			decNesting(yylex)
		}
	case 225:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1168
		{
			forceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
