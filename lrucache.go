package lrucache

import "fmt"

type LRUCache struct {
	Size  int
	head  *node
	tail  *node
	nodes map[interface{}]*node
}

type node struct {
	key   interface{}
	value interface{}
	prev  *node
	next  *node
}

func NewLRUCache(size int) *LRUCache {
	return &LRUCache{
		Size:  size,
		nodes: make(map[interface{}]*node),
	}
}

func (cache *LRUCache) Put(key interface{}, value interface{}) {
	if len(cache.nodes) >= cache.Size {
		// remove from the lookup
		staleKey := cache.tail.key
		delete(cache.nodes, staleKey)
		// make tail.prev = tail
		cache.tail.prev = cache.tail
		cache.tail.next = nil
		// add new value to head
		newNode := &node{
			key:   key,
			value: value,
		}
		cache.head.prev = newNode
		newNode.next = cache.head
		// add to hash table
		cache.nodes[key] = newNode
	} else {
		if cache.nodes[key] != nil {
			// move the node up to head if it exists
			updatedNode := cache.nodes[key]
			updatedNode.next = cache.head
			cache.head.prev = updatedNode
			cache.nodes[key] = updatedNode
		} else {
			// generate a new node
			newNode := &node{
				key:   key,
				value: value,
			}
			if cache.head != nil {
				cache.head.prev = newNode
				newNode.next = cache.head
			} else {
				cache.head = newNode
			}
			// add to hash table
			cache.nodes[key] = newNode
		}
	}
}

func (cache *LRUCache) Get(key interface{}) interface{} {
	if _, v := cache.nodes[key]; v {
		// first make the node.next = node.prev
		// next capture oldHead pointer
		// next make cache.head = node
		// next make cache.head.next = oldHead
		// next make cache.head.prev = nil
		prev := cache.nodes[key].prev
		next := cache.nodes[key].next
		if prev != nil && next != nil {
			next.prev = prev
		}
		if next != nil && prev != nil {
			prev.next = next
		}
		oldHead := cache.head
		cache.head = cache.nodes[key]
		cache.head.next = oldHead
		cache.head.prev = nil
		return cache.nodes[key].value
	} else {
		fmt.Println("shoes")
		return nil
	}
}
