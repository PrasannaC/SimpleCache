package Cache

import "SimpleCache/Utils"

/*
	The Cache interface.
*/
type ICache interface {
	Get(string) (bool, interface{})
	Size() uint64
	Put(string, interface{})
}

/*
	Method to create a concrete instance of Cache
*/
func CreateCache(n, size uint64) ICache {
	var c ICache = &Cache{
		n:      n,
		sets:   size,
		_cache: make(map[uint64]*CacheEntrySet, 0)}
	return c
}

/*
	Structure of the cache object
*/
type Cache struct {
	_cache map[uint64]*CacheEntrySet
	n      uint64
	sets   uint64
}

/*
	Put: Accepts a key-value pair and adds to cache.
*/
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

/*
	Get: Accepts a key and retrieves the cache entry data if present.
	Returns True,Data if data is present else False,Nil
*/
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

/*
	Size: Returns the size of the cache at any given moment
*/
func (c *Cache) Size() uint64 {
	var size uint64 = 0
	for _, v := range c._cache {
		size += v.GetSize()
	}
	return size
}
