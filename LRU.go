package main

import (
	"sync"
)

type LRU struct {
	miss  int
	cache map[interface{}]*Node
	queue Queue
	lock  sync.Mutex
}

type lruValue struct {
	key, value interface{}
}

func NewLRUCache(cap int) *LRU {
	head, tail := &Node{}, &Node{}
	head.next, tail.prev = tail, head
	return &LRU{cache: make(map[interface{}]*Node, cap), queue: Queue{head: head, tail: tail, capacity: cap}}
}

func (lru *LRU) Get(key interface{}) interface{} {
	lru.lock.Lock()
	defer lru.lock.Unlock()
	if node, ok := lru.cache[key]; ok {
		lru.queue.moveToHead(node)
		return node.Value.(*lruValue).value
	}
	lru.miss++
	return nil
}

func (lru *LRU) Put(key, value interface{}) {
	lru.lock.Lock()
	defer lru.lock.Unlock()
	if node, ok := lru.cache[key]; ok {
		node.Value.(*lruValue).value = value
		lru.queue.moveToHead(node)
		return
	}
	node := &Node{Value: &lruValue{key: key, value: value}}
	if len(lru.cache) == lru.queue.capacity {
		lru.evict()
	}
	lru.queue.addToHead(node)
	lru.cache[key] = node
}

func (lru *LRU) evict() {
	delete(lru.cache, (lru.queue.tail.prev).Value.(*lruValue).key)
	lru.queue.removeNode(lru.queue.tail.prev)
}

