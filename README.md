# SimpleCache
A simple n-way set associative, LRU cache implemented in Go

## Usage
```Go
import "SimpleCache/Cache"

var cache Cache.ICache
// To create a cache with 10 sets of length 3 each.
cache = Cache.CreateCache(3, 10)
// The key must be a string
// Value can be any type
cache.Put("Key",value)
cachedValue = cache.Get("key")
```
