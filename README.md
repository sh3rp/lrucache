# lrucache
LRUCache written in Golang
# Usage
    cache := NewLRUCache(2)
	  cache.Put(1, 1)
	  cache.Put(2, 2)
	  num := cache.Get(1)
	  assert.Equal(t, num, 1)
	  cache.Put(3, 3)
	  num = cache.Get(3)
	  assert.Equal(t, num, 3)
    assert.Nil(t, cache.Get(2))
