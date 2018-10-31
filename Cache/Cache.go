package Cache

import (
	"SimpleCache/Utils"
)

type ICache interface {
	Get(string) (bool, interface{})
	Size() uint64
	Put(string, interface{})
}

func CreateCache(n, size uint64) ICache {
	var c ICache = &Cache{
		n:      n,
		sets:   size,
		_cache: make(map[uint64]*CacheEntrySet, 0)}
	return c
}

type Cache struct {
	_cache map[uint64]*CacheEntrySet
	n      uint64
	sets   uint64
}

func (c *Cache) Put(key string, data interface{}) {
	bigHash := Utils.GetStringHash(key)
	//key to identify set number
	intKey := bigHash % c.sets

	if v, ok := c._cache[intKey]; !ok {
		c._cache[intKey] = NewCacheEntrySet(c.n)
		c._cache[intKey].Put(bigHash, data)
	} else {
		v.Put(bigHash, data)
	}
}

func (c *Cache) Get(key string) (bool, interface{}) {
	bigHash := Utils.GetStringHash(key)
	//key to identify set number
	intKey := bigHash % c.sets

	if v, ok := c._cache[intKey]; !ok {
		return false, nil
	} else {
		return v.Get(bigHash)
	}
}

func (c *Cache) Size() uint64 {
	var size uint64 = 0
	for _, v := range c._cache {
		size += v.GetSize()
	}
	return size
}
