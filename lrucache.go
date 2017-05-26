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
			cache.setNewHead(cache.nodes[key])
			cache.nodes[key].value = value
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
		cache.setNewHead(cache.nodes[key])
		return cache.nodes[key].value
	} else {
		return nil
	}
}

func (cache *LRUCache) setNewHead(node *node) {
	if node.prev != nil {
		prev := node.prev
		prev.next = nil
		node.prev = nil
		node.next = cache.head
		cache.head.prev = node
		if prev.next == nil {
			cache.tail = prev
		}
	}
	if node.next != nil {
		next := node.next
		next.prev = node.prev
	}
}

func (cache *LRUCache) debug() {
	head := cache.head

	for head != nil {
		fmt.Printf("%v\n", head)
		head = head.next
	}
}
