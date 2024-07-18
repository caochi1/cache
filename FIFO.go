package main

import "sync"

type FIFO struct {
	miss  int
	cache map[interface{}]*Node
	queue Queue
	lock  sync.Mutex
}

type FIFOValue struct {
	key, value interface{}
}

func NewFIFOCache(cap int) *FIFO {
	head, tail := &Node{}, &Node{}
	head.next, tail.prev = tail, head
	return &FIFO{cache: make(map[interface{}]*Node, cap), queue: Queue{head: head, tail: tail, capacity: cap}}
}

func (fifo *FIFO) Get(key interface{}) interface{} {
	fifo.lock.Lock()
	defer fifo.lock.Unlock()
	if node, ok := fifo.cache[key]; ok {
		return node.Value.(*FIFOValue).value
	}
	fifo.miss++
	return nil
}

func (fifo *FIFO) Put(key, value interface{}) {
	fifo.lock.Lock()
	defer fifo.lock.Unlock()
	if node, ok := fifo.cache[key]; ok {
		node.Value.(*FIFOValue).value = value
		return
	}
	node := &Node{Value: &FIFOValue{key: key, value: value}}
	if len(fifo.cache) == fifo.queue.capacity {
		fifo.evict()
	}
	fifo.cache[key] = node
	fifo.queue.addToHead(node)
}

func (fifo *FIFO) evict() {
	delete(fifo.cache, (fifo.queue.tail.prev).Value.(*FIFOValue).key)
	fifo.queue.removeNode(fifo.queue.tail.prev)
}

