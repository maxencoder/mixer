// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package router

import "math"

const (
	MaxKey = Key(math.MaxInt64)
	MinKey = Key(math.MinInt64)
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

// KeyUnknown means any key possible
type KeyUnknown struct{}

func (ku *KeyUnknown) Contains(key Key) bool {
	return true
}

func KeyRangesIntersect(first, second KeyRange) bool {
	return (first.End == MaxKey || second.Start < first.End) &&
		(second.End == MaxKey || first.Start < second.End)
}

func KeyExprAnd(k1, k2 KeyExpr) (r KeyExpr) {
	// TODO: do better on ranges?

	switch ke1 := k1.(type) {
	case *KeyUnknown:
		return k2
	case *KeyRange:
		return &KeyUnknown{}
	case *KeyList:
		switch ke2 := k2.(type) {
		case *KeyList:
			return &KeyList{Keys: interlist(ke1.Keys, ke2.Keys)}
		case *KeyRange:
			return &KeyUnknown{}
		case *KeyUnknown:
			return ke1
		}
	}
	return
}

func KeyExprOr(k1, k2 KeyExpr) (r KeyExpr) {
	// TODO: do better on ranges?

	switch ke1 := k1.(type) {
	case *KeyUnknown:
		return &KeyUnknown{}
	case *KeyRange:
		return &KeyUnknown{}
	case *KeyList:
		switch ke2 := k2.(type) {
		case *KeyList:
			return &KeyList{Keys: unionlist(ke1.Keys, ke2.Keys)}
		case *KeyRange:
			return &KeyUnknown{}
		case *KeyUnknown:
			return &KeyUnknown{}
		}
	}
	return r
}
