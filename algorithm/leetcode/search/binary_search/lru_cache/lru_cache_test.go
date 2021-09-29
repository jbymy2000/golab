package lru_cache

import (
	"testing"
)

func Test_lruCache(t *testing.T) {
	lru := NewLRUCache(10)
	lru.Put(1, 2)
	v, _ := lru.Get(1)
	println(v)
}
