package main

type Queue struct {
	head, tail *Node
	capacity   int
}

type Node struct {
	prev, next *Node
	Value      interface{}
}

type List interface {
	Get(key interface{}) interface{}
	Put(key, value interface{})
}

func (q *Queue) addToHead(node *Node) {
	node.prev = q.head
	node.next = q.head.next
	q.head.next, q.head.next.prev = node, node
}

func (q *Queue) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (q *Queue) moveToHead(node *Node) {
	q.removeNode(node)
	q.addToHead(node)
}
