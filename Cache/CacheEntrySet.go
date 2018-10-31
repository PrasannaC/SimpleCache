package Cache

import . "SimpleCache/DoubleLinkedList"

type CacheEntrySet struct {
	listSize uint64
	list     *DoubleLinkedList
}

func NewCacheEntrySet(n uint64) *CacheEntrySet {
	return &CacheEntrySet{listSize: n, list: new(DoubleLinkedList)}
}

func (c *CacheEntrySet) GetSize() uint64 {
	return c.list.Size
}

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

func (c *CacheEntrySet) Get(key uint64) (bool, interface{}) {
	for v := range c.list.Iterate() {
		var cNode = v.(CacheEntry)
		if cNode.key == key {
			return true, cNode.data
		}
	}
	return false, nil
}
