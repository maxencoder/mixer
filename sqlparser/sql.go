//line sql.y:6
package sqlparser

import __yyfmt__ "fmt"

//line sql.y:6
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

//line sql.y:34
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
	colIdent    ColIdent
	colIdents   []ColIdent
	tableIdent  TableIdent
}

const LEX_ERROR = 57346
const UNION = 57347
const SELECT = 57348
const INSERT = 57349
const UPDATE = 57350
const DELETE = 57351
const FROM = 57352
const WHERE = 57353
const GROUP = 57354
const HAVING = 57355
const ORDER = 57356
const BY = 57357
const LIMIT = 57358
const FOR = 57359
const ALL = 57360
const DISTINCT = 57361
const AS = 57362
const EXISTS = 57363
const ASC = 57364
const DESC = 57365
const INTO = 57366
const DUPLICATE = 57367
const KEY = 57368
const DEFAULT = 57369
const SET = 57370
const LOCK = 57371
const VALUES = 57372
const LAST_INSERT_ID = 57373
const NEXT = 57374
const VALUE = 57375
const JOIN = 57376
const STRAIGHT_JOIN = 57377
const LEFT = 57378
const RIGHT = 57379
const INNER = 57380
const OUTER = 57381
const CROSS = 57382
const NATURAL = 57383
const USE = 57384
const FORCE = 57385
const ON = 57386
const ID = 57387
const STRING = 57388
const NUMBER = 57389
const VALUE_ARG = 57390
const LIST_ARG = 57391
const COMMENT = 57392
const NULL = 57393
const TRUE = 57394
const FALSE = 57395
const OR = 57396
const AND = 57397
const NOT = 57398
const BETWEEN = 57399
const CASE = 57400
const WHEN = 57401
const THEN = 57402
const ELSE = 57403
const LE = 57404
const GE = 57405
const NE = 57406
const NULL_SAFE_EQUAL = 57407
const IS = 57408
const LIKE = 57409
const REGEXP = 57410
const IN = 57411
const SHIFT_LEFT = 57412
const SHIFT_RIGHT = 57413
const UNARY = 57414
const END = 57415
const CREATE = 57416
const ALTER = 57417
const DROP = 57418
const RENAME = 57419
const ANALYZE = 57420
const TABLE = 57421
const INDEX = 57422
const VIEW = 57423
const TO = 57424
const IGNORE = 57425
const IF = 57426
const UNIQUE = 57427
const USING = 57428
const SHOW = 57429
const DESCRIBE = 57430
const EXPLAIN = 57431
const UNUSED = 57432
const BEGIN = 57433
const COMMIT = 57434
const ROLLBACK = 57435
const NAMES = 57436
const ADMIN = 57437
const DATABASES = 57438
const TABLES = 57439
const PROXY = 57440

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"LEX_ERROR",
	"UNION",
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
	"ASC",
	"DESC",
	"INTO",
	"DUPLICATE",
	"KEY",
	"DEFAULT",
	"SET",
	"LOCK",
	"VALUES",
	"LAST_INSERT_ID",
	"NEXT",
	"VALUE",
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
	"'('",
	"','",
	"')'",
	"ID",
	"STRING",
	"NUMBER",
	"VALUE_ARG",
	"LIST_ARG",
	"COMMENT",
	"NULL",
	"TRUE",
	"FALSE",
	"OR",
	"AND",
	"NOT",
	"BETWEEN",
	"CASE",
	"WHEN",
	"THEN",
	"ELSE",
	"'='",
	"'<'",
	"'>'",
	"LE",
	"GE",
	"NE",
	"NULL_SAFE_EQUAL",
	"IS",
	"LIKE",
	"REGEXP",
	"IN",
	"'|'",
	"'&'",
	"SHIFT_LEFT",
	"SHIFT_RIGHT",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"'^'",
	"'~'",
	"UNARY",
	"'.'",
	"END",
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
	"UNUSED",
	"BEGIN",
	"COMMIT",
	"ROLLBACK",
	"NAMES",
	"ADMIN",
	"DATABASES",
	"TABLES",
	"PROXY",
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
	-1, 143,
	45, 237,
	88, 237,
	-2, 236,
}

const yyNprod = 241
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 757

var yyAct = [...]int{

	141, 309, 125, 425, 75, 190, 359, 130, 240, 294,
	334, 349, 165, 289, 56, 260, 292, 189, 3, 153,
	250, 152, 293, 172, 255, 91, 87, 296, 57, 58,
	59, 80, 44, 76, 46, 385, 387, 49, 47, 50,
	50, 52, 53, 54, 77, 60, 61, 82, 247, 114,
	84, 284, 410, 126, 63, 146, 94, 409, 408, 81,
	178, 83, 55, 51, 367, 325, 104, 201, 100, 71,
	157, 143, 127, 176, 127, 438, 186, 127, 107, 256,
	143, 137, 138, 139, 290, 109, 140, 123, 124, 92,
	118, 121, 386, 144, 180, 86, 145, 111, 148, 113,
	151, 158, 105, 150, 77, 301, 101, 77, 163, 142,
	168, 78, 131, 132, 154, 234, 19, 222, 133, 290,
	147, 329, 73, 169, 127, 191, 204, 205, 206, 201,
	162, 263, 135, 183, 218, 219, 220, 175, 177, 174,
	167, 202, 203, 204, 205, 206, 201, 225, 122, 151,
	89, 93, 179, 149, 122, 127, 232, 143, 73, 185,
	184, 345, 110, 119, 199, 207, 208, 202, 203, 204,
	205, 206, 201, 244, 186, 226, 239, 122, 122, 185,
	184, 127, 217, 184, 73, 369, 96, 235, 245, 73,
	221, 223, 187, 188, 186, 259, 88, 186, 268, 269,
	270, 62, 272, 273, 274, 275, 276, 277, 278, 279,
	280, 281, 221, 106, 257, 258, 434, 256, 238, 271,
	332, 256, 122, 285, 158, 233, 158, 251, 253, 254,
	127, 291, 252, 158, 300, 285, 216, 242, 77, 307,
	122, 122, 305, 283, 267, 287, 282, 228, 308, 261,
	299, 256, 395, 78, 304, 248, 249, 265, 266, 264,
	392, 185, 184, 97, 151, 230, 256, 317, 319, 320,
	321, 314, 315, 316, 228, 256, 186, 122, 404, 122,
	318, 355, 256, 228, 227, 298, 122, 158, 207, 208,
	202, 203, 204, 205, 206, 201, 134, 229, 350, 37,
	106, 166, 328, 262, 237, 331, 324, 350, 323, 166,
	356, 352, 326, 357, 360, 348, 243, 353, 117, 48,
	261, 127, 407, 380, 391, 364, 366, 361, 381, 378,
	406, 354, 368, 230, 379, 99, 332, 377, 72, 297,
	122, 256, 362, 373, 106, 375, 122, 376, 85, 365,
	68, 19, 90, 383, 298, 372, 393, 374, 431, 411,
	70, 330, 382, 67, 340, 341, 399, 103, 102, 170,
	432, 72, 400, 390, 262, 303, 108, 116, 347, 161,
	112, 64, 65, 115, 310, 403, 160, 311, 298, 298,
	298, 298, 241, 402, 371, 156, 166, 413, 360, 72,
	412, 164, 414, 95, 74, 437, 151, 423, 297, 417,
	415, 181, 19, 37, 182, 39, 1, 424, 18, 17,
	16, 426, 426, 426, 77, 427, 428, 15, 429, 14,
	336, 339, 340, 341, 337, 439, 338, 342, 346, 440,
	405, 441, 297, 297, 297, 297, 343, 394, 231, 171,
	433, 45, 435, 436, 397, 398, 72, 246, 122, 173,
	122, 122, 79, 126, 420, 421, 422, 19, 20, 21,
	22, 159, 306, 416, 236, 418, 419, 200, 199, 207,
	208, 202, 203, 204, 205, 206, 201, 127, 430, 23,
	143, 137, 138, 139, 396, 358, 140, 123, 124, 401,
	370, 121, 327, 144, 224, 288, 136, 129, 200, 199,
	207, 208, 202, 203, 204, 205, 206, 201, 156, 286,
	156, 351, 131, 132, 154, 128, 72, 156, 133, 302,
	192, 286, 120, 384, 19, 38, 335, 333, 295, 155,
	312, 98, 135, 313, 66, 36, 69, 13, 12, 126,
	11, 24, 25, 27, 26, 28, 40, 41, 42, 43,
	10, 9, 8, 7, 29, 30, 31, 6, 32, 33,
	34, 5, 35, 127, 4, 2, 143, 137, 138, 139,
	363, 156, 140, 123, 124, 0, 0, 121, 0, 144,
	0, 0, 0, 344, 0, 72, 0, 0, 200, 199,
	207, 208, 202, 203, 204, 205, 206, 201, 131, 132,
	0, 126, 0, 19, 133, 0, 0, 0, 127, 0,
	0, 143, 137, 138, 139, 0, 0, 140, 135, 72,
	72, 72, 72, 0, 144, 127, 0, 0, 143, 137,
	138, 139, 388, 389, 140, 123, 124, 0, 0, 121,
	0, 144, 127, 131, 132, 143, 137, 138, 139, 133,
	0, 140, 0, 0, 0, 0, 0, 0, 144, 0,
	131, 132, 0, 135, 0, 0, 133, 200, 199, 207,
	208, 202, 203, 204, 205, 206, 201, 131, 132, 0,
	135, 0, 0, 133, 336, 339, 340, 341, 337, 0,
	338, 342, 194, 197, 0, 0, 0, 135, 209, 210,
	211, 212, 213, 214, 215, 198, 195, 196, 193, 200,
	199, 207, 208, 202, 203, 204, 205, 206, 201, 322,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 200, 199, 207,
	208, 202, 203, 204, 205, 206, 201,
}
var yyPact = [...]int{

	461, -1000, -1000, 408, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -63, -60, -32, -54, -33, -84,
	-1000, -1000, -1000, -1000, -1000, 153, 406, 363, 331, -1000,
	-59, 74, 394, 63, -69, -37, 63, -1000, -34, 63,
	-1000, 74, -74, 148, -74, 74, -1000, 78, 393, 138,
	-1000, -1000, 218, -1000, -1000, -1000, 300, 63, -1000, 53,
	344, 339, -22, -1000, 74, 167, -1000, 13, -1000, 74,
	26, 114, -1000, 74, -1000, -49, 74, 356, 274, 63,
	-1000, -1000, 590, 573, 78, 573, 393, 573, 442, -1000,
	369, -1000, 74, 63, 74, 385, 63, 573, -1000, 348,
	-79, -1000, 46, -1000, 74, -1000, -1000, 74, -1000, 102,
	-1000, 590, 528, -1000, -1000, 643, 185, -1000, -1000, -1000,
	-1000, 573, 573, 573, 29, 185, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 573, 601, 607, -1000, 601, 78,
	237, 601, 287, -1000, -1000, 205, 27, 102, 643, 74,
	-1000, -1000, 276, 298, -1000, 378, 590, -1000, 601, -1000,
	-1000, 272, 63, -1000, -50, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 590, 590, 173, 4, 204, 294,
	228, 643, 573, 79, 184, 573, 573, 573, 173, 573,
	573, 573, 573, 573, 573, 573, 573, 573, 573, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 406, -1000, -1000,
	-1000, 32, 109, 442, 22, 601, -1000, -1000, 573, 136,
	442, -1000, -1000, 63, 23, -1000, 345, 63, 63, 378,
	368, 372, 102, 74, -1000, -1000, 74, -1000, 4, 125,
	-1000, -1000, 217, -1000, -1000, -1000, -1000, -1000, -1000, 601,
	-1000, 607, -1000, -1000, 79, 573, 573, 573, 601, 601,
	671, -1000, 210, 87, -1000, 44, 44, -18, -18, -18,
	61, 61, -1000, 219, 442, -1000, -23, 219, 57, -1000,
	590, 601, 290, 660, -1000, -1000, 141, 358, 110, -1000,
	-1000, -1000, 263, 185, 408, 254, 235, -1000, 368, -1000,
	573, 573, -1000, -1000, -1000, -1000, -1000, 228, -1000, 601,
	601, 522, 573, -1000, 219, 63, -1000, -25, -1000, 573,
	122, 382, 136, 136, 136, 136, -1000, 313, 303, -1000,
	295, 289, 328, -7, -1000, 74, 74, -1000, 174, -1000,
	299, 214, -1000, -1000, -1000, 63, -1000, 401, 206, -1000,
	432, -1000, -1000, 573, 601, -1000, -1000, -1000, 601, 573,
	380, 370, 660, 234, 396, -1000, -1000, -1000, -1000, 296,
	-1000, 288, -1000, -1000, -1000, -38, -39, -44, -1000, -1000,
	-1000, 333, 185, -1000, 573, 573, -1000, -1000, -1000, 601,
	601, 378, 590, 573, 590, 590, -1000, -1000, 185, 185,
	185, 399, -1000, 601, -1000, 368, 102, 201, 102, 102,
	63, 63, 63, 63, 341, 170, -1000, 170, 170, 167,
	-1000, 397, 0, -1000, 63, -1000, -1000, -1000, 63, -1000,
	63, -1000,
}
var yyPgo = [...]int{

	0, 575, 17, 574, 571, 567, 563, 562, 561, 560,
	550, 548, 547, 535, 546, 545, 544, 541, 21, 19,
	539, 16, 22, 9, 538, 537, 10, 536, 27, 533,
	3, 12, 70, 532, 530, 529, 525, 2, 20, 15,
	5, 521, 7, 109, 507, 506, 505, 13, 504, 502,
	500, 499, 8, 495, 6, 494, 1, 488, 474, 472,
	11, 4, 33, 471, 319, 95, 462, 459, 457, 451,
	449, 0, 448, 296, 446, 438, 14, 429, 427, 420,
	419, 418, 56, 25, 416, 415, 55, 24,
}
var yyR1 = [...]int{

	0, 84, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 2, 2,
	2, 3, 3, 4, 5, 6, 7, 7, 7, 8,
	8, 8, 9, 10, 10, 10, 11, 12, 12, 12,
	77, 78, 79, 81, 80, 80, 80, 85, 13, 14,
	14, 15, 15, 15, 16, 16, 17, 17, 18, 18,
	19, 19, 19, 20, 20, 72, 72, 72, 21, 21,
	22, 22, 23, 23, 23, 24, 24, 24, 24, 75,
	75, 74, 74, 74, 25, 25, 25, 25, 26, 26,
	26, 26, 27, 27, 28, 28, 29, 29, 29, 29,
	30, 30, 31, 31, 83, 83, 83, 82, 82, 32,
	32, 32, 32, 32, 32, 33, 33, 33, 33, 33,
	33, 33, 33, 33, 33, 33, 33, 33, 38, 38,
	38, 38, 38, 38, 34, 34, 34, 34, 34, 34,
	34, 39, 39, 39, 43, 40, 40, 37, 37, 37,
	37, 37, 37, 37, 37, 37, 37, 37, 37, 37,
	37, 37, 37, 37, 37, 37, 37, 37, 45, 48,
	48, 46, 46, 47, 49, 49, 44, 44, 44, 36,
	36, 36, 36, 50, 50, 51, 51, 52, 52, 53,
	53, 54, 55, 55, 55, 56, 56, 56, 57, 57,
	57, 58, 58, 59, 59, 60, 60, 35, 35, 41,
	41, 42, 42, 61, 61, 62, 63, 63, 65, 65,
	66, 66, 64, 64, 67, 67, 67, 67, 67, 67,
	68, 68, 69, 69, 70, 70, 71, 73, 86, 87,
	76,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 13, 6,
	3, 8, 8, 8, 7, 3, 5, 8, 4, 6,
	7, 4, 5, 4, 5, 5, 3, 2, 2, 2,
	1, 1, 1, 5, 3, 4, 5, 0, 2, 0,
	2, 1, 2, 2, 0, 1, 0, 1, 1, 3,
	1, 2, 3, 1, 1, 0, 1, 2, 1, 3,
	1, 1, 3, 3, 3, 3, 5, 5, 3, 0,
	1, 0, 1, 2, 1, 2, 2, 1, 2, 3,
	2, 3, 2, 2, 1, 3, 0, 5, 5, 5,
	1, 3, 0, 2, 0, 2, 2, 0, 2, 1,
	3, 3, 2, 3, 3, 1, 1, 3, 3, 4,
	3, 4, 3, 4, 5, 6, 3, 2, 1, 2,
	1, 2, 1, 2, 1, 1, 1, 1, 1, 1,
	1, 3, 1, 1, 3, 1, 3, 1, 1, 1,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	2, 2, 2, 3, 4, 5, 4, 1, 5, 0,
	1, 1, 2, 4, 0, 2, 1, 3, 5, 1,
	1, 1, 1, 0, 3, 0, 2, 0, 3, 1,
	3, 2, 0, 1, 1, 0, 2, 4, 0, 2,
	4, 0, 3, 1, 3, 0, 5, 2, 1, 1,
	3, 3, 1, 1, 3, 3, 1, 1, 0, 2,
	0, 3, 0, 1, 1, 1, 1, 1, 1, 1,
	0, 1, 0, 1, 0, 2, 1, 1, 1, 1,
	0,
}
var yyChk = [...]int{

	-1000, -84, -1, -2, -3, -4, -5, -6, -7, -8,
	-9, -10, -11, -12, -77, -78, -79, -80, -81, 6,
	7, 8, 9, 28, 90, 91, 93, 92, 94, 103,
	104, 105, 107, 108, 109, 111, -15, 5, -13, -85,
	-13, -13, -13, -13, 95, -69, 97, 101, -64, 97,
	99, 95, 95, 96, 97, 95, -76, 112, 113, 114,
	-76, -76, 48, -2, 18, 19, -16, 32, 19, -14,
	-64, -28, -73, 48, 10, -61, -62, -71, 48, -66,
	100, 96, -71, 95, -71, -73, -65, 100, 48, -65,
	-73, -83, 11, 73, -82, 10, 48, 45, -17, 35,
	-71, 53, 24, 28, 88, -28, 46, 65, -73, 59,
	48, -76, -73, -76, 98, -73, 21, 44, -71, -32,
	-33, 59, -86, 55, 56, -37, 21, 45, -36, -44,
	-42, 80, 81, 86, -73, 100, -45, 49, 50, 51,
	54, -71, -43, 48, 61, -37, -86, -83, -37, -82,
	-40, -37, -18, -19, 82, -20, -73, -32, -37, -63,
	17, 10, -28, -61, -73, -31, 11, -62, -37, -76,
	21, -70, 102, -67, 93, 91, 27, 92, 14, 106,
	48, -73, -73, -76, 58, 57, 72, -32, -32, -2,
	-40, -37, -34, 75, 59, 73, 74, 60, 72, 77,
	76, 85, 80, 81, 82, 83, 84, 78, 79, 65,
	66, 67, 68, 69, 70, 71, -43, -86, -37, -37,
	-37, -86, 88, -86, -48, -37, -83, 47, 46, 10,
	46, -72, -71, 20, 88, -28, -58, 28, -86, -31,
	-52, 14, -32, 44, -71, -76, -68, 98, -32, -32,
	-38, 54, 59, 55, 56, -87, 47, -87, -87, -37,
	-39, -86, -43, 52, 75, 73, 74, 60, -37, -37,
	-37, -38, -37, -37, -37, -37, -37, -37, -37, -37,
	-37, -37, -87, -18, 19, -71, -73, -18, -46, -47,
	62, -37, -21, -22, -23, -24, -28, -43, -86, -19,
	-71, 82, -35, 30, -2, -61, -59, -71, -52, -56,
	16, 15, -73, -73, 54, 55, 56, -40, -39, -37,
	-37, -37, 58, -87, -18, 88, -87, -49, -47, 64,
	-32, -31, 46, -25, -26, -27, 34, 38, 40, 35,
	36, 37, 41, -74, -73, 20, -75, 20, -21, -60,
	44, -41, -42, -60, -87, 46, -56, -37, -53, -54,
	-37, -76, -87, 58, -37, -87, -71, 89, -37, 63,
	-50, 12, -22, -23, -22, -23, 34, 34, 34, 39,
	34, 39, 34, -26, -29, 42, 99, 43, -73, -73,
	-87, 25, 46, -71, 46, 46, -55, 22, 23, -37,
	-37, -51, 13, 15, 44, 44, 34, 34, 96, 96,
	96, 26, -42, -37, -54, -52, -32, -40, -32, -32,
	-86, -86, -86, 8, -56, -30, -71, -30, -30, -61,
	-57, 17, 29, -87, 46, -87, -87, 8, 75, -71,
	-71, -71,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 16, 17, 47,
	47, 47, 47, 47, 232, 222, 0, 0, 0, 240,
	240, 240, 40, 41, 42, 0, 0, 51, 54, 49,
	222, 0, 0, 0, 220, 0, 0, 233, 0, 0,
	223, 0, 218, 0, 218, 0, 37, 104, 107, 0,
	38, 39, 0, 20, 52, 53, 56, 0, 55, 48,
	0, 0, 94, 237, 0, 25, 213, 0, 236, 0,
	0, 0, 240, 0, 240, 0, 0, 0, 0, 0,
	36, 44, 0, 0, 104, 0, 107, 0, 0, 57,
	0, 50, 0, 0, 0, 102, 0, 0, 240, 0,
	234, 28, 0, 31, 0, 33, 219, 0, 240, 105,
	109, 0, 0, 115, 116, 0, 0, 238, 147, 148,
	149, 0, 0, 0, 0, 0, 167, 179, 180, 181,
	182, 176, 212, -2, 169, 106, 0, 45, 108, 104,
	0, 145, 0, 58, 60, 65, 0, 63, 64, 0,
	216, 217, 201, 102, 95, 187, 0, 214, 215, 26,
	221, 0, 0, 240, 230, 224, 225, 226, 227, 228,
	229, 32, 34, 35, 0, 0, 0, 112, 0, 0,
	0, 145, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 134,
	135, 136, 137, 138, 139, 140, 127, 0, 160, 161,
	162, 0, 0, 0, 0, 170, 46, 43, 0, 0,
	0, 61, 66, 0, 0, 19, 0, 0, 0, 187,
	195, 0, 103, 0, 235, 29, 0, 231, 110, 111,
	114, 128, 0, 130, 132, 113, 239, 144, 211, 117,
	118, 0, 142, 143, 0, 0, 0, 0, 120, 122,
	0, 126, 150, 151, 152, 153, 154, 155, 156, 157,
	158, 159, 163, 0, 0, 177, 0, 0, 174, 171,
	0, 146, 102, 68, 70, 71, 81, 79, 0, 59,
	67, 62, 205, 0, 208, 205, 0, 203, 195, 24,
	0, 0, 240, 30, 129, 131, 133, 0, 119, 121,
	123, 0, 0, 164, 0, 0, 166, 0, 172, 0,
	0, 183, 0, 0, 0, 0, 84, 0, 0, 87,
	0, 0, 0, 96, 82, 0, 0, 80, 0, 21,
	0, 207, 209, 22, 202, 0, 23, 196, 188, 189,
	192, 27, 141, 0, 124, 165, 178, 168, 175, 0,
	185, 0, 69, 75, 0, 78, 85, 86, 88, 0,
	90, 0, 92, 93, 72, 0, 0, 0, 83, 73,
	74, 0, 0, 204, 0, 0, 191, 193, 194, 125,
	173, 187, 0, 0, 0, 0, 89, 91, 0, 0,
	0, 0, 210, 197, 190, 195, 186, 184, 76, 77,
	0, 0, 0, 0, 198, 0, 100, 0, 0, 206,
	18, 0, 0, 97, 0, 98, 99, 199, 0, 101,
	0, 200,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 84, 77, 3,
	45, 47, 82, 80, 46, 81, 88, 83, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	66, 65, 67, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 85, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 76, 3, 86,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 48, 49, 50, 51, 52, 53, 54,
	55, 56, 57, 58, 59, 60, 61, 62, 63, 64,
	68, 69, 70, 71, 72, 73, 74, 75, 78, 79,
	87, 89, 90, 91, 92, 93, 94, 95, 96, 97,
	98, 99, 100, 101, 102, 103, 104, 105, 106, 107,
	108, 109, 110, 111, 112, 113, 114,
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
		//line sql.y:192
		{
			setParseTree(yylex, yyDollar[1].statement)
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:198
		{
			yyVAL.statement = yyDollar[1].selStmt
		}
	case 18:
		yyDollar = yyS[yypt-13 : yypt+1]
		//line sql.y:219
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, Hints: yyDollar[4].str, SelectExprs: yyDollar[5].selectExprs, From: yyDollar[7].tableExprs, Where: NewWhere(WhereStr, yyDollar[8].boolExpr), GroupBy: GroupBy(yyDollar[9].valExprs), Having: NewWhere(HavingStr, yyDollar[10].boolExpr), OrderBy: yyDollar[11].orderBy, Limit: yyDollar[12].limit, Lock: yyDollar[13].str}
		}
	case 19:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:223
		{
			if yyDollar[4].colIdent.Lowered() != "value" {
				yylex.Error("expecting value after next")
				return 1
			}
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), SelectExprs: SelectExprs{Nextval{}}, From: TableExprs{&AliasedTableExpr{Expr: yyDollar[6].tableName}}}
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:231
		{
			yyVAL.selStmt = &Union{Type: yyDollar[2].str, Left: yyDollar[1].selStmt, Right: yyDollar[3].selStmt}
		}
	case 21:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:237
		{
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Ignore: yyDollar[3].str, Table: yyDollar[5].tableName, Columns: yyDollar[6].columns, Rows: yyDollar[7].insRows, OnDup: OnDup(yyDollar[8].updateExprs)}
		}
	case 22:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:241
		{
			cols := make(Columns, 0, len(yyDollar[7].updateExprs))
			vals := make(ValTuple, 0, len(yyDollar[7].updateExprs))
			for _, updateList := range yyDollar[7].updateExprs {
				cols = append(cols, updateList.Name)
				vals = append(vals, updateList.Expr)
			}
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Ignore: yyDollar[3].str, Table: yyDollar[5].tableName, Columns: cols, Rows: Values{vals}, OnDup: OnDup(yyDollar[8].updateExprs)}
		}
	case 23:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:253
		{
			yyVAL.statement = &Update{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[3].tableName, Exprs: yyDollar[5].updateExprs, Where: NewWhere(WhereStr, yyDollar[6].boolExpr), OrderBy: yyDollar[7].orderBy, Limit: yyDollar[8].limit}
		}
	case 24:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:259
		{
			yyVAL.statement = &Delete{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Where: NewWhere(WhereStr, yyDollar[5].boolExpr), OrderBy: yyDollar[6].orderBy, Limit: yyDollar[7].limit}
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:265
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: yyDollar[3].updateExprs}
		}
	case 26:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:271
		{
			yyVAL.statement = &DDL{Action: CreateStr, NewName: yyDollar[4].tableIdent}
		}
	case 27:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:275
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[7].tableIdent, NewName: yyDollar[7].tableIdent}
		}
	case 28:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:280
		{
			yyVAL.statement = &DDL{Action: CreateStr, NewName: TableIdent(yyDollar[3].colIdent.Lowered())}
		}
	case 29:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:286
		{
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[4].tableIdent, NewName: yyDollar[4].tableIdent}
		}
	case 30:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:290
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: RenameStr, Table: yyDollar[4].tableIdent, NewName: yyDollar[7].tableIdent}
		}
	case 31:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:295
		{
			yyVAL.statement = &DDL{Action: AlterStr, Table: TableIdent(yyDollar[3].colIdent.Lowered()), NewName: TableIdent(yyDollar[3].colIdent.Lowered())}
		}
	case 32:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:301
		{
			yyVAL.statement = &DDL{Action: RenameStr, Table: yyDollar[3].tableIdent, NewName: yyDollar[5].tableIdent}
		}
	case 33:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:307
		{
			var exists bool
			if yyDollar[3].byt != 0 {
				exists = true
			}
			yyVAL.statement = &DDL{Action: DropStr, Table: yyDollar[4].tableIdent, IfExists: exists}
		}
	case 34:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:315
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[5].tableIdent, NewName: yyDollar[5].tableIdent}
		}
	case 35:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:320
		{
			var exists bool
			if yyDollar[3].byt != 0 {
				exists = true
			}
			yyVAL.statement = &DDL{Action: DropStr, Table: TableIdent(yyDollar[4].colIdent.Lowered()), IfExists: exists}
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:330
		{
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[3].tableIdent, NewName: yyDollar[3].tableIdent}
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:336
		{
			yyVAL.statement = &Other{}
		}
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:340
		{
			yyVAL.statement = &Other{}
		}
	case 39:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:344
		{
			yyVAL.statement = &Other{}
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:350
		{
			yyVAL.statement = &Begin{}
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:356
		{
			yyVAL.statement = &Commit{}
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:362
		{
			yyVAL.statement = &Rollback{}
		}
	case 43:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:368
		{
			yyVAL.statement = &Admin{Name: string(yyDollar[2].bytes), Values: yyDollar[4].valExprs}
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:374
		{
			yyVAL.statement = &Show{Section: "databases", LikeOrWhere: yyDollar[3].expr}
		}
	case 45:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:378
		{
			yyVAL.statement = &Show{Section: "tables", From: yyDollar[3].valExpr, LikeOrWhere: yyDollar[4].expr}
		}
	case 46:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:382
		{
			yyVAL.statement = &Show{Section: "proxy", Key: string(yyDollar[3].bytes), From: yyDollar[4].valExpr, LikeOrWhere: yyDollar[5].expr}
		}
	case 47:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:387
		{
			setAllowComments(yylex, true)
		}
	case 48:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:391
		{
			yyVAL.bytes2 = yyDollar[2].bytes2
			setAllowComments(yylex, false)
		}
	case 49:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:397
		{
			yyVAL.bytes2 = nil
		}
	case 50:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:401
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[2].bytes)
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:407
		{
			yyVAL.str = UnionStr
		}
	case 52:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:411
		{
			yyVAL.str = UnionAllStr
		}
	case 53:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:415
		{
			yyVAL.str = UnionDistinctStr
		}
	case 54:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:420
		{
			yyVAL.str = ""
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:424
		{
			yyVAL.str = DistinctStr
		}
	case 56:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:429
		{
			yyVAL.str = ""
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:433
		{
			yyVAL.str = StraightJoinHint
		}
	case 58:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:439
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:443
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:449
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:453
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].colIdent}
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:457
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyDollar[1].tableIdent}
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:463
		{
			yyVAL.expr = yyDollar[1].boolExpr
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:467
		{
			yyVAL.expr = yyDollar[1].valExpr
		}
	case 65:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:472
		{
			yyVAL.colIdent = ColIdent{}
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:476
		{
			yyVAL.colIdent = yyDollar[1].colIdent
		}
	case 67:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:480
		{
			yyVAL.colIdent = yyDollar[2].colIdent
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:486
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:490
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:500
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].tableName, As: yyDollar[2].tableIdent, Hints: yyDollar[3].indexHints}
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:504
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].subquery, As: yyDollar[3].tableIdent}
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:508
		{
			yyVAL.tableExpr = &ParenTableExpr{Exprs: yyDollar[2].tableExprs}
		}
	case 75:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:521
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 76:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:525
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 77:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:529
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:533
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 79:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:538
		{
			yyVAL.empty = struct{}{}
		}
	case 80:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:540
		{
			yyVAL.empty = struct{}{}
		}
	case 81:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:543
		{
			yyVAL.tableIdent = ""
		}
	case 82:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:547
		{
			yyVAL.tableIdent = yyDollar[1].tableIdent
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:551
		{
			yyVAL.tableIdent = yyDollar[2].tableIdent
		}
	case 84:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:557
		{
			yyVAL.str = JoinStr
		}
	case 85:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:561
		{
			yyVAL.str = JoinStr
		}
	case 86:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:565
		{
			yyVAL.str = JoinStr
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:569
		{
			yyVAL.str = StraightJoinStr
		}
	case 88:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:575
		{
			yyVAL.str = LeftJoinStr
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:579
		{
			yyVAL.str = LeftJoinStr
		}
	case 90:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:583
		{
			yyVAL.str = RightJoinStr
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:587
		{
			yyVAL.str = RightJoinStr
		}
	case 92:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:593
		{
			yyVAL.str = NaturalJoinStr
		}
	case 93:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:597
		{
			if yyDollar[2].str == LeftJoinStr {
				yyVAL.str = NaturalLeftJoinStr
			} else {
				yyVAL.str = NaturalRightJoinStr
			}
		}
	case 94:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:607
		{
			yyVAL.tableName = &TableName{Name: yyDollar[1].tableIdent}
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:611
		{
			yyVAL.tableName = &TableName{Qualifier: yyDollar[1].tableIdent, Name: yyDollar[3].tableIdent}
		}
	case 96:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:616
		{
			yyVAL.indexHints = nil
		}
	case 97:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:620
		{
			yyVAL.indexHints = &IndexHints{Type: UseStr, Indexes: yyDollar[4].colIdents}
		}
	case 98:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:624
		{
			yyVAL.indexHints = &IndexHints{Type: IgnoreStr, Indexes: yyDollar[4].colIdents}
		}
	case 99:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:628
		{
			yyVAL.indexHints = &IndexHints{Type: ForceStr, Indexes: yyDollar[4].colIdents}
		}
	case 100:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:634
		{
			yyVAL.colIdents = []ColIdent{yyDollar[1].colIdent}
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:638
		{
			yyVAL.colIdents = append(yyDollar[1].colIdents, yyDollar[3].colIdent)
		}
	case 102:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:643
		{
			yyVAL.boolExpr = nil
		}
	case 103:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:647
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 104:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:652
		{
			yyVAL.expr = nil
		}
	case 105:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:656
		{
			yyVAL.expr = yyDollar[2].boolExpr
		}
	case 106:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:660
		{
			yyVAL.expr = yyDollar[2].valExpr
		}
	case 107:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:665
		{
			yyVAL.valExpr = nil
		}
	case 108:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:669
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 110:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:676
		{
			yyVAL.boolExpr = &AndExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 111:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:680
		{
			yyVAL.boolExpr = &OrExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 112:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:684
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyDollar[2].boolExpr}
		}
	case 113:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:688
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyDollar[2].boolExpr}
		}
	case 114:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:692
		{
			yyVAL.boolExpr = &IsExpr{Operator: yyDollar[3].str, Expr: yyDollar[1].boolExpr}
		}
	case 115:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:698
		{
			yyVAL.boolExpr = BoolVal(true)
		}
	case 116:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:702
		{
			yyVAL.boolExpr = BoolVal(false)
		}
	case 117:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:706
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 118:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:710
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: InStr, Right: yyDollar[3].colTuple}
		}
	case 119:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:714
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: NotInStr, Right: yyDollar[4].colTuple}
		}
	case 120:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:718
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: LikeStr, Right: yyDollar[3].valExpr}
		}
	case 121:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:722
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: NotLikeStr, Right: yyDollar[4].valExpr}
		}
	case 122:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:726
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: RegexpStr, Right: yyDollar[3].valExpr}
		}
	case 123:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:730
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: NotRegexpStr, Right: yyDollar[4].valExpr}
		}
	case 124:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:734
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: BetweenStr, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 125:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:738
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: NotBetweenStr, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 126:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:742
		{
			yyVAL.boolExpr = &IsExpr{Operator: yyDollar[3].str, Expr: yyDollar[1].valExpr}
		}
	case 127:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:746
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 128:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:752
		{
			yyVAL.str = IsNullStr
		}
	case 129:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:756
		{
			yyVAL.str = IsNotNullStr
		}
	case 130:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:760
		{
			yyVAL.str = IsTrueStr
		}
	case 131:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:764
		{
			yyVAL.str = IsNotTrueStr
		}
	case 132:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:768
		{
			yyVAL.str = IsFalseStr
		}
	case 133:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:772
		{
			yyVAL.str = IsNotFalseStr
		}
	case 134:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:778
		{
			yyVAL.str = EqualStr
		}
	case 135:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:782
		{
			yyVAL.str = LessThanStr
		}
	case 136:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:786
		{
			yyVAL.str = GreaterThanStr
		}
	case 137:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:790
		{
			yyVAL.str = LessEqualStr
		}
	case 138:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:794
		{
			yyVAL.str = GreaterEqualStr
		}
	case 139:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:798
		{
			yyVAL.str = NotEqualStr
		}
	case 140:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:802
		{
			yyVAL.str = NullSafeEqualStr
		}
	case 141:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:808
		{
			yyVAL.colTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 142:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:812
		{
			yyVAL.colTuple = yyDollar[1].subquery
		}
	case 143:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:816
		{
			yyVAL.colTuple = ListArg(yyDollar[1].bytes)
		}
	case 144:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:822
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 145:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:828
		{
			yyVAL.valExprs = ValExprs{yyDollar[1].valExpr}
		}
	case 146:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:832
		{
			yyVAL.valExprs = append(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 147:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:838
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 148:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:842
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 149:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:846
		{
			yyVAL.valExpr = yyDollar[1].rowTuple
		}
	case 150:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:850
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: BitAndStr, Right: yyDollar[3].valExpr}
		}
	case 151:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:854
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: BitOrStr, Right: yyDollar[3].valExpr}
		}
	case 152:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:858
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: BitXorStr, Right: yyDollar[3].valExpr}
		}
	case 153:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:862
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: PlusStr, Right: yyDollar[3].valExpr}
		}
	case 154:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:866
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: MinusStr, Right: yyDollar[3].valExpr}
		}
	case 155:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:870
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: MultStr, Right: yyDollar[3].valExpr}
		}
	case 156:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:874
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: DivStr, Right: yyDollar[3].valExpr}
		}
	case 157:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:878
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ModStr, Right: yyDollar[3].valExpr}
		}
	case 158:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:882
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ShiftLeftStr, Right: yyDollar[3].valExpr}
		}
	case 159:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:886
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ShiftRightStr, Right: yyDollar[3].valExpr}
		}
	case 160:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:890
		{
			if num, ok := yyDollar[2].valExpr.(NumVal); ok {
				yyVAL.valExpr = num
			} else {
				yyVAL.valExpr = &UnaryExpr{Operator: UPlusStr, Expr: yyDollar[2].valExpr}
			}
		}
	case 161:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:898
		{
			if num, ok := yyDollar[2].valExpr.(NumVal); ok {
				// Handle double negative
				if num[0] == '-' {
					yyVAL.valExpr = num[1:]
				} else {
					yyVAL.valExpr = append(NumVal("-"), num...)
				}
			} else {
				yyVAL.valExpr = &UnaryExpr{Operator: UMinusStr, Expr: yyDollar[2].valExpr}
			}
		}
	case 162:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:911
		{
			yyVAL.valExpr = &UnaryExpr{Operator: TildaStr, Expr: yyDollar[2].valExpr}
		}
	case 163:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:915
		{
			yyVAL.valExpr = &FuncExpr{Name: string(yyDollar[1].tableIdent)}
		}
	case 164:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:919
		{
			yyVAL.valExpr = &FuncExpr{Name: string(yyDollar[1].tableIdent), Exprs: yyDollar[3].selectExprs}
		}
	case 165:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:923
		{
			yyVAL.valExpr = &FuncExpr{Name: string(yyDollar[1].tableIdent), Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 166:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:927
		{
			yyVAL.valExpr = &FuncExpr{Name: "if", Exprs: yyDollar[3].selectExprs}
		}
	case 167:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:931
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 168:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:937
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 169:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:942
		{
			yyVAL.valExpr = nil
		}
	case 170:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:946
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 171:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:952
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 172:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:956
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 173:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:962
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 174:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:967
		{
			yyVAL.valExpr = nil
		}
	case 175:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:971
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 176:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:977
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].colIdent}
		}
	case 177:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:981
		{
			yyVAL.colName = &ColName{Qualifier: &TableName{Name: yyDollar[1].tableIdent}, Name: yyDollar[3].colIdent}
		}
	case 178:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:985
		{
			yyVAL.colName = &ColName{Qualifier: &TableName{Qualifier: yyDollar[1].tableIdent, Name: yyDollar[3].tableIdent}, Name: yyDollar[5].colIdent}
		}
	case 179:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:991
		{
			yyVAL.valExpr = StrVal(yyDollar[1].bytes)
		}
	case 180:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:995
		{
			yyVAL.valExpr = NumVal(yyDollar[1].bytes)
		}
	case 181:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:999
		{
			yyVAL.valExpr = ValArg(yyDollar[1].bytes)
		}
	case 182:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1003
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 183:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1008
		{
			yyVAL.valExprs = nil
		}
	case 184:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1012
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 185:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1017
		{
			yyVAL.boolExpr = nil
		}
	case 186:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1021
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 187:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1026
		{
			yyVAL.orderBy = nil
		}
	case 188:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1030
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 189:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1036
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 190:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1040
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 191:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1046
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 192:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1051
		{
			yyVAL.str = AscScr
		}
	case 193:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1055
		{
			yyVAL.str = AscScr
		}
	case 194:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1059
		{
			yyVAL.str = DescScr
		}
	case 195:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1064
		{
			yyVAL.limit = nil
		}
	case 196:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1068
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr}
		}
	case 197:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1072
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 198:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1077
		{
			yyVAL.str = ""
		}
	case 199:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1081
		{
			yyVAL.str = ForUpdateStr
		}
	case 200:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1085
		{
			if yyDollar[3].colIdent.Lowered() != "share" {
				yylex.Error("expecting share")
				return 1
			}
			if yyDollar[4].colIdent.Lowered() != "mode" {
				yylex.Error("expecting mode")
				return 1
			}
			yyVAL.str = ShareModeStr
		}
	case 201:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1098
		{
			yyVAL.columns = nil
		}
	case 202:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1102
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 203:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1108
		{
			yyVAL.columns = Columns{yyDollar[1].colIdent}
		}
	case 204:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1112
		{
			yyVAL.columns = append(yyVAL.columns, yyDollar[3].colIdent)
		}
	case 205:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1117
		{
			yyVAL.updateExprs = nil
		}
	case 206:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:1121
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 207:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1127
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 208:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1131
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 209:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1137
		{
			yyVAL.values = Values{yyDollar[1].rowTuple}
		}
	case 210:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1141
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].rowTuple)
		}
	case 211:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1147
		{
			yyVAL.rowTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 212:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1151
		{
			yyVAL.rowTuple = yyDollar[1].subquery
		}
	case 213:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1157
		{
			yyVAL.updateExprs = UpdateExprs{yyDollar[1].updateExpr}
		}
	case 214:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1161
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 215:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1167
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colIdent, Expr: yyDollar[3].valExpr}
		}
	case 218:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1176
		{
			yyVAL.byt = 0
		}
	case 219:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1178
		{
			yyVAL.byt = 1
		}
	case 220:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1181
		{
			yyVAL.empty = struct{}{}
		}
	case 221:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1183
		{
			yyVAL.empty = struct{}{}
		}
	case 222:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1186
		{
			yyVAL.str = ""
		}
	case 223:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1188
		{
			yyVAL.str = IgnoreStr
		}
	case 224:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1192
		{
			yyVAL.empty = struct{}{}
		}
	case 225:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1194
		{
			yyVAL.empty = struct{}{}
		}
	case 226:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1196
		{
			yyVAL.empty = struct{}{}
		}
	case 227:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1198
		{
			yyVAL.empty = struct{}{}
		}
	case 228:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1200
		{
			yyVAL.empty = struct{}{}
		}
	case 229:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1202
		{
			yyVAL.empty = struct{}{}
		}
	case 230:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1205
		{
			yyVAL.empty = struct{}{}
		}
	case 231:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1207
		{
			yyVAL.empty = struct{}{}
		}
	case 232:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1210
		{
			yyVAL.empty = struct{}{}
		}
	case 233:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1212
		{
			yyVAL.empty = struct{}{}
		}
	case 234:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1215
		{
			yyVAL.empty = struct{}{}
		}
	case 235:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1217
		{
			yyVAL.empty = struct{}{}
		}
	case 236:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1221
		{
			yyVAL.colIdent = NewColIdent(string(yyDollar[1].bytes))
		}
	case 237:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1227
		{
			yyVAL.tableIdent = TableIdent(yyDollar[1].bytes)
		}
	case 238:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1233
		{
			if incNesting(yylex) {
				yylex.Error("max nesting level reached")
				return 1
			}
		}
	case 239:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1242
		{
			decNesting(yylex)
		}
	case 240:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1247
		{
			forceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
