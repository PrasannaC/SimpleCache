package Cache

import . "SimpleCache/DoubleLinkedList"

/*
	Structure of the set of cache entries
*/
type CacheEntrySet struct {
	listSize uint64
	list     *DoubleLinkedList
}

/*
	Factory function for creating a cache entry set.
	Accepts: n - number of sets of entries in every entry set
*/
func NewCacheEntrySet(n uint64) *CacheEntrySet {
	return &CacheEntrySet{listSize: n, list: new(DoubleLinkedList)}
}

/*
	Returns the size of the current entry set
*/
func (c *CacheEntrySet) GetSize() uint64 {
	return c.list.Size
}

/*
	Put: Accepts a key-value pair to be stored in the current entry set.
*/
func (c *CacheEntrySet) Put(key uint64, data interface{}) {
	cNode := CacheEntry{key: key, data: data}
	if ok, oldVal := c.Get(cNode.key); ok {
		oldCNode := CacheEntry{key: cNode.key, data: oldVal}
		c.list.Remove(oldCNode)
		c.list.AddHead(cNode)
	} else {
		if c.list.Size == c.listSize {
			c.list.RemoveTail()
			c.list.AddHead(cNode)
		} else {
			c.list.AddHead(cNode)
		}
	}
}

/*
	Get: Accepts a key and retrieves the cache entry data if present.
	Returns True,Data if data is present else False,Nil
*/
func (c *CacheEntrySet) Get(key uint64) (bool, interface{}) {
	for v := range c.list.Iterate() {
		var cNode = v.(CacheEntry)
		if cNode.key == key {
			return true, cNode.data
		}
	}
	return false, nil
}
