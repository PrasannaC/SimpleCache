package main

import (
	"SimpleCache/Cache"
	"testing"
)

func TestCache(t *testing.T) {
	var cache Cache.ICache
	cache = Cache.CreateCache(3, 10)
	cache.Put("zzzz", "test data")
	cache.Put("zzzzzzzzz", 10)
	cache.Get("zzzzzzzzz")
	cache.Put("zzzz", "testing data")
	cache.Get("zzzz")
}

func Benchmark(t *testing.B) {
	var cache Cache.ICache
	cache = Cache.CreateCache(30, 10)

	for i := 0; i < t.N; i++ {
		cache.Put("zzzz", "this is a test")
		cache.Put("zzzzzzzzz", i)
	}
}

func Benchmark2(t *testing.B) {
	var cache Cache.ICache
	cache = Cache.CreateCache(3, 10)
	for i := 0; i < t.N; i++ {
		cache.Put("random", i)
	}
}
