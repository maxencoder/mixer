// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package adminparser

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/maxencoder/mixer/sqltypes"
)

const eofChar = 0x100

// Tokenizer is the struct used to generate SQL
// tokens for the parser.
type Tokenizer struct {
	InStream      *strings.Reader
	AllowComments bool
	ForceEOF      bool
	lastChar      uint16
	Position      int
	lastToken     []byte
	LastError     string
	posVarIndex   int
	ParseTree     Command
	nesting       int
}

// NewStringTokenizer creates a new Tokenizer for the
// sql string.
func NewStringTokenizer(sql string) *Tokenizer {
	return &Tokenizer{InStream: strings.NewReader(sql)}
}

var keywords = map[string]int{
	"add":      ADD,
	"alter":    ALTER,
	"database": DATABASE,
	"default":  DEFAULT,
	"delete":   DELETE,
	"hash":     HASH,
	"inf":      INF,
	"key":      KEY,
	"modulo":   MODULO,
	"range":    RANGE,
	"route":    ROUTE,
	"router":   ROUTER,
	"routes":   ROUTES,
	"show":     SHOW,
	"table":    TABLE,
	"to":       TO,
	"type":     TYPE,
}

// Lex returns the next token form the Tokenizer.
// This function is used by go yacc.
func (tkn *Tokenizer) Lex(lval *yySymType) int {
	typ, val := tkn.Scan()
	for typ == COMMENT {
		if tkn.AllowComments {
			break
		}
		typ, val = tkn.Scan()
	}
	switch typ {
	case ID, STRING, NUMBER, LIST_ARG, COMMENT:
		lval.bytes = val
	}
	tkn.lastToken = val
	return typ
}

// Error is called by go yacc if there's a parsing error.
func (tkn *Tokenizer) Error(err string) {
	buf := &bytes.Buffer{}
	if tkn.lastToken != nil {
		fmt.Fprintf(buf, "%s at position %v near '%s'", err, tkn.Position, tkn.lastToken)
	} else {
		fmt.Fprintf(buf, "%s at position %v", err, tkn.Position)
	}
	tkn.LastError = buf.String()
}

// Scan scans the tokenizer for the next token and returns
// the token type and an optional value.
func (tkn *Tokenizer) Scan() (int, []byte) {
	if tkn.ForceEOF {
		return 0, nil
	}

	if tkn.lastChar == 0 {
		tkn.next()
	}
	tkn.skipBlank()
	switch ch := tkn.lastChar; {
	case isLetter(ch):
		return tkn.scanIdentifier()
	case isDigit(ch):
		return tkn.scanNumber(false, false)
	default:
		tkn.next()
		switch ch {
		case eofChar:
			return 0, nil
		case '=', ',', ';', '(', ')', '+', '*', '%', '&', '|', '^', '~':
			return int(ch), nil
		case '.':
			if isDigit(tkn.lastChar) {
				return tkn.scanNumber(true, false)
			}
			return int(ch), nil
		case '/':
			switch tkn.lastChar {
			case '/':
				tkn.next()
				return tkn.scanCommentType1("//")
			case '*':
				tkn.next()
				return tkn.scanCommentType2()
			default:
				return int(ch), nil
			}
		case '-':
			if tkn.lastChar == '-' {
				tkn.next()
				return tkn.scanCommentType1("--")
			}
			return tkn.scanNumber(false, true)
		case '\'', '"':
			return tkn.scanString(ch, STRING)
		case '`':
			return tkn.scanLiteralIdentifier()
		default:
			return LEX_ERROR, []byte{byte(ch)}
		}
	}
}

func (tkn *Tokenizer) skipBlank() {
	ch := tkn.lastChar
	for ch == ' ' || ch == '\n' || ch == '\r' || ch == '\t' {
		tkn.next()
		ch = tkn.lastChar
	}
}

func (tkn *Tokenizer) scanIdentifier() (int, []byte) {
	buffer := &bytes.Buffer{}
	buffer.WriteByte(byte(tkn.lastChar))
	for tkn.next(); isValidIdentifierChar(tkn.lastChar); tkn.next() {
		buffer.WriteByte(byte(tkn.lastChar))
	}
	lowered := bytes.ToLower(buffer.Bytes())
	loweredStr := string(lowered)
	if keywordID, found := keywords[loweredStr]; found {
		return keywordID, lowered
	}
	return ID, buffer.Bytes()
}

func (tkn *Tokenizer) scanLiteralIdentifier() (int, []byte) {
	buffer := &bytes.Buffer{}
	buffer.WriteByte(byte(tkn.lastChar))
	if !isLetter(tkn.lastChar) {
		return LEX_ERROR, buffer.Bytes()
	}
	for tkn.next(); isLetter(tkn.lastChar) || isDigit(tkn.lastChar); tkn.next() {
		buffer.WriteByte(byte(tkn.lastChar))
	}
	if tkn.lastChar != '`' {
		return LEX_ERROR, buffer.Bytes()
	}
	tkn.next()
	return ID, buffer.Bytes()
}

func (tkn *Tokenizer) scanMantissa(base int, buffer *bytes.Buffer) {
	for digitVal(tkn.lastChar) < base {
		tkn.consumeNext(buffer)
	}
}

func (tkn *Tokenizer) scanNumber(seenDecimalPoint, seenMinus bool) (int, []byte) {
	buffer := &bytes.Buffer{}
	if seenDecimalPoint {
		buffer.WriteByte('.')
		tkn.scanMantissa(10, buffer)
		goto exponent
	}
	if seenMinus {
		buffer.WriteByte('-')
	}

	if tkn.lastChar == '0' {
		// int or float
		tkn.consumeNext(buffer)
		if tkn.lastChar == 'x' || tkn.lastChar == 'X' {
			// hexadecimal int
			tkn.consumeNext(buffer)
			tkn.scanMantissa(16, buffer)
		} else {
			// octal int or float
			seenDecimalDigit := false
			tkn.scanMantissa(8, buffer)
			if tkn.lastChar == '8' || tkn.lastChar == '9' {
				// illegal octal int or float
				seenDecimalDigit = true
				tkn.scanMantissa(10, buffer)
			}
			if tkn.lastChar == '.' || tkn.lastChar == 'e' || tkn.lastChar == 'E' {
				goto fraction
			}
			// octal int
			if seenDecimalDigit {
				return LEX_ERROR, buffer.Bytes()
			}
		}
		goto exit
	}

	// decimal int or float
	tkn.scanMantissa(10, buffer)

fraction:
	if tkn.lastChar == '.' {
		tkn.consumeNext(buffer)
		tkn.scanMantissa(10, buffer)
	}

exponent:
	if tkn.lastChar == 'e' || tkn.lastChar == 'E' {
		tkn.consumeNext(buffer)
		if tkn.lastChar == '+' || tkn.lastChar == '-' {
			tkn.consumeNext(buffer)
		}
		tkn.scanMantissa(10, buffer)
	}

exit:
	return NUMBER, buffer.Bytes()
}

func (tkn *Tokenizer) scanString(delim uint16, typ int) (int, []byte) {
	buffer := &bytes.Buffer{}
	for {
		ch := tkn.lastChar
		tkn.next()
		if ch == delim {
			if tkn.lastChar == delim {
				tkn.next()
			} else {
				break
			}
		} else if ch == '\\' {
			if tkn.lastChar == eofChar {
				return LEX_ERROR, buffer.Bytes()
			}
			if decodedChar := sqltypes.SqlDecodeMap[byte(tkn.lastChar)]; decodedChar == sqltypes.DONTESCAPE {
				ch = tkn.lastChar
			} else {
				ch = uint16(decodedChar)
			}
			tkn.next()
		}
		if ch == eofChar {
			return LEX_ERROR, buffer.Bytes()
		}
		buffer.WriteByte(byte(ch))
	}
	return typ, buffer.Bytes()
}

func (tkn *Tokenizer) scanCommentType1(prefix string) (int, []byte) {
	buffer := &bytes.Buffer{}
	buffer.WriteString(prefix)
	for tkn.lastChar != eofChar {
		if tkn.lastChar == '\n' {
			tkn.consumeNext(buffer)
			break
		}
		tkn.consumeNext(buffer)
	}
	return COMMENT, buffer.Bytes()
}

func (tkn *Tokenizer) scanCommentType2() (int, []byte) {
	buffer := &bytes.Buffer{}
	buffer.WriteString("/*")
	for {
		if tkn.lastChar == '*' {
			tkn.consumeNext(buffer)
			if tkn.lastChar == '/' {
				tkn.consumeNext(buffer)
				break
			}
			continue
		}
		if tkn.lastChar == eofChar {
			return LEX_ERROR, buffer.Bytes()
		}
		tkn.consumeNext(buffer)
	}
	return COMMENT, buffer.Bytes()
}

func (tkn *Tokenizer) consumeNext(buffer *bytes.Buffer) {
	if tkn.lastChar == eofChar {
		// This should never happen.
		panic("unexpected EOF")
	}
	buffer.WriteByte(byte(tkn.lastChar))
	tkn.next()
}

func (tkn *Tokenizer) next() {
	if ch, err := tkn.InStream.ReadByte(); err != nil {
		// Only EOF is possible.
		tkn.lastChar = eofChar
	} else {
		tkn.lastChar = uint16(ch)
	}
	tkn.Position++
}

func isLetter(ch uint16) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch == '@'
}

func digitVal(ch uint16) int {
	switch {
	case '0' <= ch && ch <= '9':
		return int(ch) - '0'
	case 'a' <= ch && ch <= 'f':
		return int(ch) - 'a' + 10
	case 'A' <= ch && ch <= 'F':
		return int(ch) - 'A' + 10
	}
	return 16 // larger than any legal digit val
}

func isDigit(ch uint16) bool {
	return '0' <= ch && ch <= '9'
}

func isValidIdentifierChar(ch uint16) bool {
	return isLetter(ch) || isDigit(ch) || ch == '-' || ch == '_'
}
