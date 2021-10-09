package v2

import "testing"

func Test_lruCache_v2(t *testing.T) {
	lru := NewLRUCache(10)
	lru.Set(1, 2)
	v, ok := lru.Get(1)
	println(v, ok, 1)
	println(1)
}
