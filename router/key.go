// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package router

import ()

const (
	MaxUint64 = ^uint64(0)            // 18446744073709551615
	MaxInt64  = int64(MaxUint64 >> 1) // 9223372036854775807
	MinInt64  = -MaxInt64 - 1         // -9223372036854775808

	MaxKey = Key(MaxInt64)
	MinKey = Key(MinInt64)
)

// routing key
type Key int64

func NewKey(k interface{}) Key {
	return Key(NumValue(k))
}

type KeyExpr interface {
	iKeyExpr()
}

type KeyList struct {
	Keys []Key
}

type KeyRange struct {
	Start Key
	End   Key
}

type KeyUnknown struct{}

func (k *KeyList) iKeyExpr()    {}
func (k *KeyRange) iKeyExpr()   {}
func (k *KeyUnknown) iKeyExpr() {}

func KeyRangesIntersect(first, second KeyRange) bool {
	return (first.End == MaxKey || second.Start < first.End) &&
		(second.End == MaxKey || first.Start < second.End)
}
