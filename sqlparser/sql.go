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
	-1, 126,
	45, 238,
	88, 238,
	-2, 237,
}

const yyNprod = 242
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 778

var yyAct = [...]int{

	123, 311, 146, 427, 75, 188, 242, 361, 112, 336,
	351, 296, 294, 155, 56, 295, 281, 92, 255, 265,
	187, 3, 79, 274, 125, 167, 194, 88, 298, 154,
	57, 58, 59, 81, 76, 387, 389, 50, 19, 20,
	21, 22, 271, 135, 77, 60, 61, 83, 84, 49,
	85, 50, 44, 55, 46, 159, 412, 63, 47, 95,
	23, 52, 53, 54, 411, 410, 82, 51, 101, 365,
	71, 128, 128, 200, 174, 175, 176, 177, 178, 173,
	316, 109, 173, 105, 78, 126, 198, 176, 177, 178,
	173, 139, 388, 440, 208, 207, 206, 148, 132, 150,
	134, 153, 160, 106, 152, 77, 93, 202, 77, 165,
	208, 170, 108, 149, 236, 185, 181, 182, 183, 303,
	266, 130, 24, 25, 27, 26, 28, 153, 102, 266,
	190, 320, 164, 87, 124, 29, 30, 31, 79, 32,
	33, 34, 169, 35, 191, 258, 211, 288, 73, 140,
	197, 199, 196, 126, 205, 207, 206, 151, 234, 206,
	286, 287, 285, 128, 128, 201, 73, 37, 94, 228,
	208, 284, 226, 208, 245, 246, 247, 248, 249, 250,
	251, 252, 253, 254, 131, 347, 259, 160, 90, 160,
	237, 241, 275, 277, 278, 268, 97, 276, 209, 210,
	436, 258, 207, 206, 235, 89, 262, 263, 367, 258,
	269, 399, 400, 73, 256, 280, 261, 208, 289, 290,
	291, 324, 325, 326, 244, 334, 258, 107, 143, 279,
	62, 230, 79, 293, 143, 160, 302, 259, 283, 397,
	77, 309, 292, 352, 307, 107, 301, 231, 310, 230,
	258, 184, 186, 232, 258, 394, 299, 357, 258, 306,
	160, 128, 272, 273, 239, 172, 171, 179, 180, 174,
	175, 176, 177, 178, 173, 314, 168, 143, 143, 98,
	317, 128, 227, 232, 319, 153, 406, 315, 327, 329,
	330, 331, 352, 184, 338, 341, 342, 343, 339, 240,
	340, 344, 328, 143, 407, 230, 229, 267, 168, 100,
	283, 334, 358, 350, 354, 359, 362, 364, 355, 143,
	333, 143, 321, 366, 138, 299, 19, 356, 382, 384,
	48, 342, 343, 383, 363, 371, 409, 368, 408, 380,
	379, 143, 143, 107, 381, 378, 369, 375, 282, 377,
	374, 257, 376, 147, 385, 104, 413, 103, 395, 299,
	299, 299, 299, 68, 393, 128, 300, 143, 73, 392,
	401, 70, 19, 402, 192, 137, 67, 128, 433, 258,
	126, 119, 120, 121, 349, 312, 122, 144, 145, 405,
	434, 142, 143, 127, 64, 65, 305, 163, 116, 415,
	362, 143, 243, 414, 162, 416, 313, 404, 153, 373,
	417, 419, 113, 114, 156, 168, 96, 74, 115, 426,
	282, 439, 425, 428, 428, 428, 77, 429, 430, 19,
	431, 37, 117, 39, 38, 300, 1, 441, 18, 17,
	72, 442, 16, 443, 15, 14, 435, 348, 437, 438,
	86, 345, 233, 193, 91, 40, 41, 42, 43, 45,
	418, 147, 420, 421, 270, 195, 80, 161, 308, 300,
	300, 300, 300, 72, 238, 432, 398, 360, 403, 129,
	372, 318, 189, 133, 264, 128, 136, 118, 126, 119,
	120, 121, 111, 353, 122, 144, 145, 110, 158, 142,
	304, 127, 72, 212, 166, 179, 180, 174, 175, 176,
	177, 178, 173, 19, 141, 386, 337, 335, 297, 157,
	113, 114, 156, 99, 66, 36, 115, 69, 147, 338,
	341, 342, 343, 339, 203, 340, 344, 204, 13, 143,
	117, 143, 143, 12, 11, 422, 423, 424, 10, 9,
	8, 370, 128, 7, 6, 126, 119, 120, 121, 5,
	72, 122, 144, 145, 4, 2, 142, 0, 127, 172,
	171, 179, 180, 174, 175, 176, 177, 178, 173, 0,
	0, 147, 0, 158, 260, 158, 0, 113, 114, 0,
	0, 0, 0, 115, 0, 0, 0, 0, 0, 0,
	19, 0, 0, 0, 0, 128, 0, 117, 126, 119,
	120, 121, 0, 0, 122, 144, 145, 0, 0, 142,
	0, 127, 0, 0, 0, 0, 0, 0, 0, 0,
	72, 158, 0, 0, 0, 260, 0, 0, 0, 128,
	113, 114, 126, 119, 120, 121, 115, 0, 122, 0,
	0, 0, 0, 0, 396, 127, 158, 0, 0, 0,
	117, 0, 0, 0, 0, 0, 322, 128, 0, 323,
	126, 119, 120, 121, 113, 114, 122, 0, 0, 0,
	115, 0, 0, 127, 172, 171, 179, 180, 174, 175,
	176, 177, 178, 173, 117, 0, 0, 346, 0, 72,
	0, 0, 113, 114, 0, 332, 0, 0, 115, 172,
	171, 179, 180, 174, 175, 176, 177, 178, 173, 0,
	0, 0, 117, 172, 171, 179, 180, 174, 175, 176,
	177, 178, 173, 72, 72, 72, 72, 0, 0, 0,
	0, 0, 214, 217, 0, 0, 390, 391, 219, 220,
	221, 222, 223, 224, 225, 218, 215, 216, 213, 172,
	171, 179, 180, 174, 175, 176, 177, 178, 173, 171,
	179, 180, 174, 175, 176, 177, 178, 173,
}
var yyPact = [...]int{

	32, -1000, -1000, 426, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -43, -48, -28, -34, -42, -82,
	-1000, -1000, -1000, -1000, -1000, 182, 423, 376, 344, -1000,
	-62, 100, 407, -26, -67, -30, 90, -1000, -47, 90,
	-1000, 100, -73, 157, -73, 100, -1000, 95, 406, 148,
	-1000, -1000, 234, -1000, -1000, -1000, 274, 90, -1000, 75,
	333, 327, -5, -1000, 100, 181, -1000, 47, 622, -1000,
	100, 62, 136, -1000, 100, -1000, -55, 100, 354, 280,
	90, -1000, -1000, 560, 622, 95, 622, 406, 622, 440,
	-1000, 387, -1000, 100, -26, 100, 404, -26, 622, 633,
	-1000, -1000, -1000, 622, 622, 622, 27, 216, -1000, -1000,
	-1000, -1000, -1000, -1000, 594, -1000, -1000, 622, -1000, -1000,
	353, -76, -1000, 59, -1000, 100, -1000, -1000, 100, -1000,
	38, -1000, 560, 507, -1000, -1000, 683, 216, 633, -1000,
	633, 95, 259, 633, 237, -1000, -1000, 184, 26, 38,
	683, 100, -1000, -1000, 236, 297, -1000, 388, 560, -1000,
	633, 622, 622, 622, 622, 622, 622, 622, 622, 622,
	622, -1000, -1000, -1000, 332, 105, 440, 162, 203, 58,
	633, -1000, -1000, 263, 90, -1000, -56, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 560, 560, 138, 22,
	98, 683, 622, 119, 87, 622, 622, 622, 138, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 423, -1000, -1000,
	622, 118, 440, -1000, -1000, 90, 37, -1000, 366, -26,
	90, 388, 369, 391, 38, 427, 692, -1000, 5, 5,
	-3, -3, -3, -6, -6, -1000, 207, 440, -1000, -1000,
	-8, 207, -1000, -1000, 67, -1000, 560, 100, -1000, -1000,
	100, -1000, 22, 101, -1000, -1000, 167, -1000, -1000, -1000,
	633, -1000, 594, -1000, -1000, 119, 622, 622, 622, 633,
	633, 647, -1000, 633, 265, 495, -1000, -1000, 165, 364,
	320, -1000, -1000, -1000, 248, 216, 426, 199, 211, -1000,
	369, -1000, 622, 622, -1000, 207, 90, -1000, -20, -1000,
	622, 145, -1000, -1000, -1000, -1000, -1000, 203, -1000, 633,
	633, 493, 622, 397, 118, 118, 118, 118, -1000, 311,
	306, -1000, 305, 294, 295, -7, -1000, 100, 100, -1000,
	179, -1000, 339, 209, -1000, -1000, -1000, 90, -1000, 608,
	193, -1000, 189, -1000, -1000, -1000, 633, 622, -1000, -1000,
	622, 633, 394, 374, 495, 242, 260, -1000, -1000, -1000,
	-1000, 304, -1000, 302, -1000, -1000, -1000, -31, -32, -40,
	-1000, -1000, -1000, 330, 216, -1000, 622, 622, -1000, -1000,
	-1000, 633, 633, 388, 560, 622, 560, 560, -1000, -1000,
	216, 216, 216, 414, -1000, 633, -1000, 369, 38, 185,
	38, 38, 90, 90, 90, -26, 361, 154, -1000, 154,
	154, 181, -1000, 413, 18, -1000, 90, -1000, -1000, -1000,
	90, -1000, 90, -1000,
}
var yyPgo = [...]int{

	0, 565, 20, 564, 559, 554, 553, 550, 549, 548,
	544, 543, 538, 434, 527, 525, 524, 523, 29, 13,
	519, 12, 15, 11, 518, 517, 9, 516, 28, 515,
	3, 25, 55, 514, 503, 500, 497, 2, 23, 16,
	5, 493, 8, 24, 492, 487, 484, 19, 482, 481,
	480, 478, 6, 477, 7, 476, 1, 475, 474, 468,
	10, 4, 34, 467, 330, 133, 466, 465, 464, 459,
	453, 0, 452, 398, 451, 447, 14, 445, 444, 442,
	439, 438, 59, 17, 436, 433, 134, 18,
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
	41, 42, 42, 61, 61, 62, 62, 63, 63, 65,
	65, 66, 66, 64, 64, 67, 67, 67, 67, 67,
	67, 68, 68, 69, 69, 70, 70, 71, 73, 86,
	87, 76,
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
	3, 3, 1, 1, 3, 3, 2, 1, 1, 0,
	2, 0, 3, 0, 1, 1, 1, 1, 1, 1,
	1, 0, 1, 0, 1, 0, 2, 1, 1, 1,
	1, 0,
}
var yyChk = [...]int{

	-1000, -84, -1, -2, -3, -4, -5, -6, -7, -8,
	-9, -10, -11, -12, -77, -78, -79, -80, -81, 6,
	7, 8, 9, 28, 90, 91, 93, 92, 94, 103,
	104, 105, 107, 108, 109, 111, -15, 5, -13, -85,
	-13, -13, -13, -13, 95, -69, 97, 101, -64, 97,
	99, 95, 95, 96, 97, 95, -76, 112, 113, 114,
	-76, -76, 48, -2, 18, 19, -16, 32, 19, -14,
	-64, -28, -73, 48, 10, -61, -62, -71, 110, 48,
	-66, 100, 96, -71, 95, -71, -73, -65, 100, 48,
	-65, -73, -83, 11, 73, -82, 10, 48, 45, -17,
	35, -71, 53, 24, 28, 88, -28, 46, 65, -37,
	-36, -44, -42, 80, 81, 86, -73, 100, -45, 49,
	50, 51, 54, -71, -86, -43, 48, 61, 45, -73,
	59, 48, -76, -73, -76, 98, -73, 21, 44, -71,
	-32, -33, 59, -86, 55, 56, -37, 21, -37, -83,
	-37, -82, -40, -37, -18, -19, 82, -20, -73, -32,
	-37, -63, 17, 10, -28, -61, -73, -31, 11, -62,
	-37, 77, 76, 85, 80, 81, 82, 83, 84, 78,
	79, -37, -37, -37, -86, 88, -86, -2, -40, -48,
	-37, -76, 21, -70, 102, -67, 93, 91, 27, 92,
	14, 106, 48, -73, -73, -76, 58, 57, 72, -32,
	-32, -37, -34, 75, 59, 73, 74, 60, 72, 65,
	66, 67, 68, 69, 70, 71, -43, -86, -83, 47,
	46, 10, 46, -72, -71, 20, 88, -28, -58, 28,
	-86, -31, -52, 14, -32, -37, -37, -37, -37, -37,
	-37, -37, -37, -37, -37, -87, -18, 19, 47, -71,
	-73, -18, -87, -87, -46, -47, 62, 44, -71, -76,
	-68, 98, -32, -32, -38, 54, 59, 55, 56, -87,
	-37, -39, -86, -43, 52, 75, 73, 74, 60, -37,
	-37, -37, -38, -37, -21, -22, -23, -24, -28, -43,
	-86, -19, -71, 82, -35, 30, -2, -61, -59, -71,
	-52, -56, 16, 15, -87, -18, 88, -87, -49, -47,
	64, -32, -73, -73, 54, 55, 56, -40, -39, -37,
	-37, -37, 58, -31, 46, -25, -26, -27, 34, 38,
	40, 35, 36, 37, 41, -74, -73, 20, -75, 20,
	-21, -60, 44, -41, -42, -60, -87, 46, -56, -37,
	-53, -54, -37, -87, -71, 89, -37, 63, -76, -87,
	58, -37, -50, 12, -22, -23, -22, -23, 34, 34,
	34, 39, 34, 39, 34, -26, -29, 42, 99, 43,
	-73, -73, -87, 25, 46, -71, 46, 46, -55, 22,
	23, -37, -37, -51, 13, 15, 44, 44, 34, 34,
	96, 96, 96, 26, -42, -37, -54, -52, -32, -40,
	-32, -32, -86, -86, -86, 8, -56, -30, -71, -30,
	-30, -61, -57, 17, 29, -87, 46, -87, -87, 8,
	75, -71, -71, -71,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 16, 17, 47,
	47, 47, 47, 47, 233, 223, 0, 0, 0, 241,
	241, 241, 40, 41, 42, 0, 0, 51, 54, 49,
	223, 0, 0, 0, 221, 0, 0, 234, 0, 0,
	224, 0, 219, 0, 219, 0, 37, 104, 107, 0,
	38, 39, 0, 20, 52, 53, 56, 0, 55, 48,
	0, 0, 94, 238, 0, 25, 213, 0, 0, 237,
	0, 0, 0, 241, 0, 241, 0, 0, 0, 0,
	0, 36, 44, 0, 0, 104, 0, 107, 0, 0,
	57, 0, 50, 0, 0, 0, 102, 0, 0, 216,
	147, 148, 149, 0, 0, 0, 0, 0, 167, 179,
	180, 181, 182, 176, 0, 212, -2, 169, 239, 241,
	0, 235, 28, 0, 31, 0, 33, 220, 0, 241,
	105, 109, 0, 0, 115, 116, 0, 0, 106, 45,
	108, 104, 0, 145, 0, 58, 60, 65, 0, 63,
	64, 0, 217, 218, 201, 102, 95, 187, 0, 214,
	215, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 160, 161, 162, 0, 0, 0, 0, 0, 0,
	170, 26, 222, 0, 0, 241, 231, 225, 226, 227,
	228, 229, 230, 32, 34, 35, 0, 0, 0, 112,
	0, 145, 0, 0, 0, 0, 0, 0, 0, 134,
	135, 136, 137, 138, 139, 140, 127, 0, 46, 43,
	0, 0, 0, 61, 66, 0, 0, 19, 0, 0,
	0, 187, 195, 0, 103, 150, 151, 152, 153, 154,
	155, 156, 157, 158, 159, 163, 0, 0, 240, 177,
	0, 0, 144, 211, 174, 171, 0, 0, 236, 29,
	0, 232, 110, 111, 114, 128, 0, 130, 132, 113,
	117, 118, 0, 142, 143, 0, 0, 0, 0, 120,
	122, 0, 126, 146, 102, 68, 70, 71, 81, 79,
	0, 59, 67, 62, 205, 0, 208, 205, 0, 203,
	195, 24, 0, 0, 164, 0, 0, 166, 0, 172,
	0, 0, 241, 30, 129, 131, 133, 0, 119, 121,
	123, 0, 0, 183, 0, 0, 0, 0, 84, 0,
	0, 87, 0, 0, 0, 96, 82, 0, 0, 80,
	0, 21, 0, 207, 209, 22, 202, 0, 23, 196,
	188, 189, 192, 165, 178, 168, 175, 0, 27, 141,
	0, 124, 185, 0, 69, 75, 0, 78, 85, 86,
	88, 0, 90, 0, 92, 93, 72, 0, 0, 0,
	83, 73, 74, 0, 0, 204, 0, 0, 191, 193,
	194, 173, 125, 187, 0, 0, 0, 0, 89, 91,
	0, 0, 0, 0, 210, 197, 190, 195, 186, 184,
	76, 77, 0, 0, 0, 0, 198, 0, 100, 0,
	0, 206, 18, 0, 0, 97, 0, 98, 99, 199,
	0, 101, 0, 200,
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
		//line sql.y:191
		{
			setParseTree(yylex, yyDollar[1].statement)
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:197
		{
			yyVAL.statement = yyDollar[1].selStmt
		}
	case 18:
		yyDollar = yyS[yypt-13 : yypt+1]
		//line sql.y:218
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, Hints: yyDollar[4].str, SelectExprs: yyDollar[5].selectExprs, From: yyDollar[7].tableExprs, Where: NewWhere(WhereStr, yyDollar[8].boolExpr), GroupBy: GroupBy(yyDollar[9].valExprs), Having: NewWhere(HavingStr, yyDollar[10].boolExpr), OrderBy: yyDollar[11].orderBy, Limit: yyDollar[12].limit, Lock: yyDollar[13].str}
		}
	case 19:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:222
		{
			if yyDollar[4].colIdent.Lowered() != "value" {
				yylex.Error("expecting value after next")
				return 1
			}
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), SelectExprs: SelectExprs{Nextval{}}, From: TableExprs{&AliasedTableExpr{Expr: yyDollar[6].tableName}}}
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:230
		{
			yyVAL.selStmt = &Union{Type: yyDollar[2].str, Left: yyDollar[1].selStmt, Right: yyDollar[3].selStmt}
		}
	case 21:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:236
		{
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Ignore: yyDollar[3].str, Table: yyDollar[5].tableName, Columns: yyDollar[6].columns, Rows: yyDollar[7].insRows, OnDup: OnDup(yyDollar[8].updateExprs)}
		}
	case 22:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:240
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
		//line sql.y:252
		{
			yyVAL.statement = &Update{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[3].tableName, Exprs: yyDollar[5].updateExprs, Where: NewWhere(WhereStr, yyDollar[6].boolExpr), OrderBy: yyDollar[7].orderBy, Limit: yyDollar[8].limit}
		}
	case 24:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:258
		{
			yyVAL.statement = &Delete{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Where: NewWhere(WhereStr, yyDollar[5].boolExpr), OrderBy: yyDollar[6].orderBy, Limit: yyDollar[7].limit}
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:264
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: yyDollar[3].updateExprs}
		}
	case 26:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:270
		{
			yyVAL.statement = &DDL{Action: CreateStr, NewName: yyDollar[4].tableIdent}
		}
	case 27:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:274
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[7].tableIdent, NewName: yyDollar[7].tableIdent}
		}
	case 28:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:279
		{
			yyVAL.statement = &DDL{Action: CreateStr, NewName: TableIdent(yyDollar[3].colIdent.Lowered())}
		}
	case 29:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:285
		{
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[4].tableIdent, NewName: yyDollar[4].tableIdent}
		}
	case 30:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:289
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: RenameStr, Table: yyDollar[4].tableIdent, NewName: yyDollar[7].tableIdent}
		}
	case 31:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:294
		{
			yyVAL.statement = &DDL{Action: AlterStr, Table: TableIdent(yyDollar[3].colIdent.Lowered()), NewName: TableIdent(yyDollar[3].colIdent.Lowered())}
		}
	case 32:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:300
		{
			yyVAL.statement = &DDL{Action: RenameStr, Table: yyDollar[3].tableIdent, NewName: yyDollar[5].tableIdent}
		}
	case 33:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:306
		{
			var exists bool
			if yyDollar[3].byt != 0 {
				exists = true
			}
			yyVAL.statement = &DDL{Action: DropStr, Table: yyDollar[4].tableIdent, IfExists: exists}
		}
	case 34:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:314
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[5].tableIdent, NewName: yyDollar[5].tableIdent}
		}
	case 35:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:319
		{
			var exists bool
			if yyDollar[3].byt != 0 {
				exists = true
			}
			yyVAL.statement = &DDL{Action: DropStr, Table: TableIdent(yyDollar[4].colIdent.Lowered()), IfExists: exists}
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:329
		{
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[3].tableIdent, NewName: yyDollar[3].tableIdent}
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:335
		{
			yyVAL.statement = &Other{}
		}
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:339
		{
			yyVAL.statement = &Other{}
		}
	case 39:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:343
		{
			yyVAL.statement = &Other{}
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:349
		{
			yyVAL.statement = &Begin{}
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:355
		{
			yyVAL.statement = &Commit{}
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:361
		{
			yyVAL.statement = &Rollback{}
		}
	case 43:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:367
		{
			yyVAL.statement = &Admin{Name: string(yyDollar[2].bytes), Values: yyDollar[4].valExprs}
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:373
		{
			yyVAL.statement = &Show{Section: "databases", LikeOrWhere: yyDollar[3].expr}
		}
	case 45:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:377
		{
			yyVAL.statement = &Show{Section: "tables", From: yyDollar[3].valExpr, LikeOrWhere: yyDollar[4].expr}
		}
	case 46:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:381
		{
			yyVAL.statement = &Show{Section: "proxy", Key: string(yyDollar[3].bytes), From: yyDollar[4].valExpr, LikeOrWhere: yyDollar[5].expr}
		}
	case 47:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:386
		{
			setAllowComments(yylex, true)
		}
	case 48:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:390
		{
			yyVAL.bytes2 = yyDollar[2].bytes2
			setAllowComments(yylex, false)
		}
	case 49:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:396
		{
			yyVAL.bytes2 = nil
		}
	case 50:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:400
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[2].bytes)
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:406
		{
			yyVAL.str = UnionStr
		}
	case 52:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:410
		{
			yyVAL.str = UnionAllStr
		}
	case 53:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:414
		{
			yyVAL.str = UnionDistinctStr
		}
	case 54:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:419
		{
			yyVAL.str = ""
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:423
		{
			yyVAL.str = DistinctStr
		}
	case 56:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:428
		{
			yyVAL.str = ""
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:432
		{
			yyVAL.str = StraightJoinHint
		}
	case 58:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:438
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:442
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:448
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:452
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].colIdent}
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:456
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyDollar[1].tableIdent}
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:462
		{
			yyVAL.expr = yyDollar[1].boolExpr
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:466
		{
			yyVAL.expr = yyDollar[1].valExpr
		}
	case 65:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:471
		{
			yyVAL.colIdent = ColIdent{}
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:475
		{
			yyVAL.colIdent = yyDollar[1].colIdent
		}
	case 67:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:479
		{
			yyVAL.colIdent = yyDollar[2].colIdent
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:485
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:489
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:499
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].tableName, As: yyDollar[2].tableIdent, Hints: yyDollar[3].indexHints}
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:503
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].subquery, As: yyDollar[3].tableIdent}
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:507
		{
			yyVAL.tableExpr = &ParenTableExpr{Exprs: yyDollar[2].tableExprs}
		}
	case 75:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:520
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 76:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:524
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 77:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:528
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:532
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 79:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:537
		{
			yyVAL.empty = struct{}{}
		}
	case 80:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:539
		{
			yyVAL.empty = struct{}{}
		}
	case 81:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:542
		{
			yyVAL.tableIdent = ""
		}
	case 82:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:546
		{
			yyVAL.tableIdent = yyDollar[1].tableIdent
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:550
		{
			yyVAL.tableIdent = yyDollar[2].tableIdent
		}
	case 84:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:556
		{
			yyVAL.str = JoinStr
		}
	case 85:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:560
		{
			yyVAL.str = JoinStr
		}
	case 86:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:564
		{
			yyVAL.str = JoinStr
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:568
		{
			yyVAL.str = StraightJoinStr
		}
	case 88:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:574
		{
			yyVAL.str = LeftJoinStr
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:578
		{
			yyVAL.str = LeftJoinStr
		}
	case 90:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:582
		{
			yyVAL.str = RightJoinStr
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:586
		{
			yyVAL.str = RightJoinStr
		}
	case 92:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:592
		{
			yyVAL.str = NaturalJoinStr
		}
	case 93:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:596
		{
			if yyDollar[2].str == LeftJoinStr {
				yyVAL.str = NaturalLeftJoinStr
			} else {
				yyVAL.str = NaturalRightJoinStr
			}
		}
	case 94:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:606
		{
			yyVAL.tableName = &TableName{Name: yyDollar[1].tableIdent}
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:610
		{
			yyVAL.tableName = &TableName{Qualifier: yyDollar[1].tableIdent, Name: yyDollar[3].tableIdent}
		}
	case 96:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:615
		{
			yyVAL.indexHints = nil
		}
	case 97:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:619
		{
			yyVAL.indexHints = &IndexHints{Type: UseStr, Indexes: yyDollar[4].colIdents}
		}
	case 98:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:623
		{
			yyVAL.indexHints = &IndexHints{Type: IgnoreStr, Indexes: yyDollar[4].colIdents}
		}
	case 99:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:627
		{
			yyVAL.indexHints = &IndexHints{Type: ForceStr, Indexes: yyDollar[4].colIdents}
		}
	case 100:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:633
		{
			yyVAL.colIdents = []ColIdent{yyDollar[1].colIdent}
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:637
		{
			yyVAL.colIdents = append(yyDollar[1].colIdents, yyDollar[3].colIdent)
		}
	case 102:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:642
		{
			yyVAL.boolExpr = nil
		}
	case 103:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:646
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 104:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:651
		{
			yyVAL.expr = nil
		}
	case 105:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:655
		{
			yyVAL.expr = yyDollar[2].boolExpr
		}
	case 106:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:659
		{
			yyVAL.expr = yyDollar[2].valExpr
		}
	case 107:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:664
		{
			yyVAL.valExpr = nil
		}
	case 108:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:668
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 110:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:675
		{
			yyVAL.boolExpr = &AndExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 111:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:679
		{
			yyVAL.boolExpr = &OrExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 112:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:683
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyDollar[2].boolExpr}
		}
	case 113:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:687
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyDollar[2].boolExpr}
		}
	case 114:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:691
		{
			yyVAL.boolExpr = &IsExpr{Operator: yyDollar[3].str, Expr: yyDollar[1].boolExpr}
		}
	case 115:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:697
		{
			yyVAL.boolExpr = BoolVal(true)
		}
	case 116:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:701
		{
			yyVAL.boolExpr = BoolVal(false)
		}
	case 117:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:705
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 118:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:709
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: InStr, Right: yyDollar[3].colTuple}
		}
	case 119:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:713
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: NotInStr, Right: yyDollar[4].colTuple}
		}
	case 120:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:717
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: LikeStr, Right: yyDollar[3].valExpr}
		}
	case 121:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:721
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: NotLikeStr, Right: yyDollar[4].valExpr}
		}
	case 122:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:725
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: RegexpStr, Right: yyDollar[3].valExpr}
		}
	case 123:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:729
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: NotRegexpStr, Right: yyDollar[4].valExpr}
		}
	case 124:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:733
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: BetweenStr, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 125:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:737
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: NotBetweenStr, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 126:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:741
		{
			yyVAL.boolExpr = &IsExpr{Operator: yyDollar[3].str, Expr: yyDollar[1].valExpr}
		}
	case 127:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:745
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 128:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:751
		{
			yyVAL.str = IsNullStr
		}
	case 129:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:755
		{
			yyVAL.str = IsNotNullStr
		}
	case 130:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:759
		{
			yyVAL.str = IsTrueStr
		}
	case 131:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:763
		{
			yyVAL.str = IsNotTrueStr
		}
	case 132:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:767
		{
			yyVAL.str = IsFalseStr
		}
	case 133:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:771
		{
			yyVAL.str = IsNotFalseStr
		}
	case 134:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:777
		{
			yyVAL.str = EqualStr
		}
	case 135:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:781
		{
			yyVAL.str = LessThanStr
		}
	case 136:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:785
		{
			yyVAL.str = GreaterThanStr
		}
	case 137:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:789
		{
			yyVAL.str = LessEqualStr
		}
	case 138:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:793
		{
			yyVAL.str = GreaterEqualStr
		}
	case 139:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:797
		{
			yyVAL.str = NotEqualStr
		}
	case 140:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:801
		{
			yyVAL.str = NullSafeEqualStr
		}
	case 141:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:807
		{
			yyVAL.colTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 142:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:811
		{
			yyVAL.colTuple = yyDollar[1].subquery
		}
	case 143:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:815
		{
			yyVAL.colTuple = ListArg(yyDollar[1].bytes)
		}
	case 144:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:821
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 145:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:827
		{
			yyVAL.valExprs = ValExprs{yyDollar[1].valExpr}
		}
	case 146:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:831
		{
			yyVAL.valExprs = append(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 147:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:837
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 148:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:841
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 149:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:845
		{
			yyVAL.valExpr = yyDollar[1].rowTuple
		}
	case 150:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:849
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: BitAndStr, Right: yyDollar[3].valExpr}
		}
	case 151:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:853
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: BitOrStr, Right: yyDollar[3].valExpr}
		}
	case 152:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:857
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: BitXorStr, Right: yyDollar[3].valExpr}
		}
	case 153:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:861
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: PlusStr, Right: yyDollar[3].valExpr}
		}
	case 154:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:865
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: MinusStr, Right: yyDollar[3].valExpr}
		}
	case 155:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:869
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: MultStr, Right: yyDollar[3].valExpr}
		}
	case 156:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:873
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: DivStr, Right: yyDollar[3].valExpr}
		}
	case 157:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:877
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ModStr, Right: yyDollar[3].valExpr}
		}
	case 158:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:881
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ShiftLeftStr, Right: yyDollar[3].valExpr}
		}
	case 159:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:885
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ShiftRightStr, Right: yyDollar[3].valExpr}
		}
	case 160:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:889
		{
			if num, ok := yyDollar[2].valExpr.(NumVal); ok {
				yyVAL.valExpr = num
			} else {
				yyVAL.valExpr = &UnaryExpr{Operator: UPlusStr, Expr: yyDollar[2].valExpr}
			}
		}
	case 161:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:897
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
		//line sql.y:910
		{
			yyVAL.valExpr = &UnaryExpr{Operator: TildaStr, Expr: yyDollar[2].valExpr}
		}
	case 163:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:914
		{
			yyVAL.valExpr = &FuncExpr{Name: string(yyDollar[1].tableIdent)}
		}
	case 164:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:918
		{
			yyVAL.valExpr = &FuncExpr{Name: string(yyDollar[1].tableIdent), Exprs: yyDollar[3].selectExprs}
		}
	case 165:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:922
		{
			yyVAL.valExpr = &FuncExpr{Name: string(yyDollar[1].tableIdent), Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 166:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:926
		{
			yyVAL.valExpr = &FuncExpr{Name: "if", Exprs: yyDollar[3].selectExprs}
		}
	case 167:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:930
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 168:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:936
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 169:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:941
		{
			yyVAL.valExpr = nil
		}
	case 170:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:945
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 171:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:951
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 172:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:955
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 173:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:961
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 174:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:966
		{
			yyVAL.valExpr = nil
		}
	case 175:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:970
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 176:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:976
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].colIdent}
		}
	case 177:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:980
		{
			yyVAL.colName = &ColName{Qualifier: &TableName{Name: yyDollar[1].tableIdent}, Name: yyDollar[3].colIdent}
		}
	case 178:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:984
		{
			yyVAL.colName = &ColName{Qualifier: &TableName{Qualifier: yyDollar[1].tableIdent, Name: yyDollar[3].tableIdent}, Name: yyDollar[5].colIdent}
		}
	case 179:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:990
		{
			yyVAL.valExpr = StrVal(yyDollar[1].bytes)
		}
	case 180:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:994
		{
			yyVAL.valExpr = NumVal(yyDollar[1].bytes)
		}
	case 181:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:998
		{
			yyVAL.valExpr = ValArg(yyDollar[1].bytes)
		}
	case 182:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1002
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 183:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1007
		{
			yyVAL.valExprs = nil
		}
	case 184:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1011
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 185:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1016
		{
			yyVAL.boolExpr = nil
		}
	case 186:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1020
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 187:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1025
		{
			yyVAL.orderBy = nil
		}
	case 188:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1029
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 189:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1035
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 190:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1039
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 191:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1045
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 192:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1050
		{
			yyVAL.str = AscScr
		}
	case 193:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1054
		{
			yyVAL.str = AscScr
		}
	case 194:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1058
		{
			yyVAL.str = DescScr
		}
	case 195:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1063
		{
			yyVAL.limit = nil
		}
	case 196:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1067
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr}
		}
	case 197:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1071
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 198:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1076
		{
			yyVAL.str = ""
		}
	case 199:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1080
		{
			yyVAL.str = ForUpdateStr
		}
	case 200:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1084
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
		//line sql.y:1097
		{
			yyVAL.columns = nil
		}
	case 202:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1101
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 203:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1107
		{
			yyVAL.columns = Columns{yyDollar[1].colIdent}
		}
	case 204:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1111
		{
			yyVAL.columns = append(yyVAL.columns, yyDollar[3].colIdent)
		}
	case 205:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1116
		{
			yyVAL.updateExprs = nil
		}
	case 206:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:1120
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 207:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1126
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 208:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1130
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 209:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1136
		{
			yyVAL.values = Values{yyDollar[1].rowTuple}
		}
	case 210:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1140
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].rowTuple)
		}
	case 211:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1146
		{
			yyVAL.rowTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 212:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1150
		{
			yyVAL.rowTuple = yyDollar[1].subquery
		}
	case 213:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1156
		{
			yyVAL.updateExprs = UpdateExprs{yyDollar[1].updateExpr}
		}
	case 214:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1160
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 215:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1166
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colIdent, Expr: yyDollar[3].valExpr}
		}
	case 216:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1170
		{
			yyVAL.updateExpr = &UpdateExpr{Name: NewColIdent("names"), Expr: yyDollar[2].valExpr}
		}
	case 219:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1179
		{
			yyVAL.byt = 0
		}
	case 220:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1181
		{
			yyVAL.byt = 1
		}
	case 221:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1184
		{
			yyVAL.empty = struct{}{}
		}
	case 222:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1186
		{
			yyVAL.empty = struct{}{}
		}
	case 223:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1189
		{
			yyVAL.str = ""
		}
	case 224:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1191
		{
			yyVAL.str = IgnoreStr
		}
	case 225:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1195
		{
			yyVAL.empty = struct{}{}
		}
	case 226:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1197
		{
			yyVAL.empty = struct{}{}
		}
	case 227:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1199
		{
			yyVAL.empty = struct{}{}
		}
	case 228:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1201
		{
			yyVAL.empty = struct{}{}
		}
	case 229:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1203
		{
			yyVAL.empty = struct{}{}
		}
	case 230:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1205
		{
			yyVAL.empty = struct{}{}
		}
	case 231:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1208
		{
			yyVAL.empty = struct{}{}
		}
	case 232:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1210
		{
			yyVAL.empty = struct{}{}
		}
	case 233:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1213
		{
			yyVAL.empty = struct{}{}
		}
	case 234:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1215
		{
			yyVAL.empty = struct{}{}
		}
	case 235:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1218
		{
			yyVAL.empty = struct{}{}
		}
	case 236:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1220
		{
			yyVAL.empty = struct{}{}
		}
	case 237:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1224
		{
			yyVAL.colIdent = NewColIdent(string(yyDollar[1].bytes))
		}
	case 238:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1230
		{
			yyVAL.tableIdent = TableIdent(yyDollar[1].bytes)
		}
	case 239:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1236
		{
			if incNesting(yylex) {
				yylex.Error("max nesting level reached")
				return 1
			}
		}
	case 240:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1245
		{
			decNesting(yylex)
		}
	case 241:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1250
		{
			forceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
