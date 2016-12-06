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
	Contains(Key) bool
}

func (k *KeyList) iKeyExpr()    {}
func (k *KeyRange) iKeyExpr()   {}
func (k *KeyUnknown) iKeyExpr() {}

type KeyList struct {
	Keys []Key
}

func (kl *KeyList) Contains(key Key) bool {
	// TODO: optimize
	for _, k := range kl.Keys {
		if k == key {
			return true
		}
	}
	return false
}

type KeyRange struct {
	Start Key
	End   Key
}

func (kr *KeyRange) Contains(key Key) bool {
	return kr.Start <= key && key < kr.End
}

type KeyUnknown struct{}

func (ku *KeyUnknown) Contains(key Key) bool {
	return true
}

func KeyRangesIntersect(first, second KeyRange) bool {
	return (first.End == MaxKey || second.Start < first.End) &&
		(second.End == MaxKey || first.Start < second.End)
}

func KeyExprAnd(k1, k2 KeyExpr) KeyExpr {
	return k1

	/* TODO: implement
	switch ke1 := k1.(type) {
	case KeyUnknown, KeyRange:
		return k2
	case KeyList:
		switch ke2 := k2.(type) {
		case KeyList:
			return interKeyList(k1
		}
	}
	*/
}

func KeyExprOr(k1, k2 KeyExpr) KeyExpr {
	return k1

	// TODO: implement
}
