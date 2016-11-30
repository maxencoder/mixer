// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package router

import (
	"fmt"
	"strconv"

	"github.com/maxencoder/mixer/config"
	"github.com/maxencoder/mixer/hack"
	"github.com/maxencoder/mixer/node"
)

func NumValue(value interface{}) int64 {
	switch val := value.(type) {
	case int:
		return int64(val)
	case uint64:
		return int64(val)
	case int64:
		return int64(val)
	case string:
		if v, err := strconv.ParseInt(val, 10, 64); err != nil {
			panic(fmt.Errorf("invalid num format %s", v))
		} else {
			return v
		}
	case []byte:
		if v, err := strconv.ParseInt(hack.String(val), 10, 64); err != nil {
			panic(fmt.Errorf("invalid num format %s", v))
		} else {
			return v
		}
	}
	panic(fmt.Errorf("Unexpected key variable type %T", value))
}

type Shard interface {
	FindForKey(key interface{}) int
}

type RangeShard interface {
	Shard
	EqualStart(key interface{}, index int) bool
	EqualStop(key interface{}, index int) bool
}

type HashShard struct {
	ShardNum int
}

func (s *HashShard) FindForKey(key interface{}) int {
	h := NumValue(key)

	return int(modulo(h, int64(s.ShardNum)))
}

type NumRangeShard struct {
	Shards []NumKeyRange
}

func (s *NumRangeShard) FindForKey(key interface{}) int {
	v := NumValue(key)
	for i, r := range s.Shards {
		if r.Contains(v) {
			return i
		}
	}
	panic(fmt.Errorf("Unexpected key %v, not in range", key))
}

func (s *NumRangeShard) EqualStart(key interface{}, index int) bool {
	v := NumValue(key)
	return s.Shards[index].Start == v
}
func (s *NumRangeShard) EqualStop(key interface{}, index int) bool {
	v := NumValue(key)
	return s.Shards[index].End == v
}

type HashLookupShard struct {
	ShardNum int
	Lookup   Lookup
}

type Lookup struct {
	config.LookupConfig
}

func (s *HashLookupShard) FindForKey(key interface{}) int {
	conn, err := node.GetNode(s.Lookup.Node).GetSelectConn()
	if err != nil {
		panic(err)
	}

	k := strconv.FormatInt(NumValue(key), 10)
	sql := fmt.Sprintf(s.Lookup.Query, k)
	r, err := conn.Execute(sql)
	if err != nil {
		panic(fmt.Errorf("Failed to get shard_id: %v", err))
	}
	if len(r.Values) != 1 {
		panic(fmt.Errorf("Failed to find hotel_id for room_id: %s", k))
	}
	h := NumValue(r.Values[0][0])

	return int(h % int64(s.ShardNum))
}

type DefaultShard struct {
}

func (s *DefaultShard) FindForKey(key interface{}) int {
	return 0
}
