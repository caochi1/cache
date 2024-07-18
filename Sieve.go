package main

import "sync"

type Sieve struct {
	miss  int
	cache map[interface{}]*Node
	queue Queue
	hand  *Node
	lock  sync.Mutex
}

type sieveValue struct {
	k, v    interface{}
	visited bool
}

func NewSieveCache(cap int) *Sieve {
	head, tail := &Node{Value: &sieveValue{}}, &Node{}
	head.next, tail.prev = tail, head
	return &Sieve{
		cache: make(map[interface{}]*Node, cap),
		queue: Queue{head: head, tail: tail, capacity: cap},
		hand:  head,
	}
}

func (s *Sieve) Get(key interface{}) interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()
	if node, exist := s.cache[key]; exist {
		node.Value.(*sieveValue).visited = true
		return node.Value.(*sieveValue).v
	}
	s.miss++
	return nil
}

func (s *Sieve) Put(key, value interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if node, exist := s.cache[key]; exist {
		node.Value.(*sieveValue).v = value
		return
	}
	if len(s.cache) == s.queue.capacity {
		s.evict()
	}
	node := &Node{Value: &sieveValue{k: key, v: value}}
	s.cache[key] = node
	s.queue.addToHead(node)
}

func (s *Sieve) evict() {
	for {
		if s.hand == s.queue.head {
			s.hand = s.queue.tail.prev
		}
		sv := s.hand.Value.(*sieveValue)
		if sv.visited {
			sv.visited = false
			s.hand = s.hand.prev
		} else {
			delete(s.cache, sv.k)
			s.queue.removeNode(s.hand)
			s.hand = s.hand.prev
			return
		}
	}
}
