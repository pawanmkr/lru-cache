package main

import (
	"fmt"
)

const SIZE = 5

type Node struct {
	val   string
	left  *Node
	right *Node
}

type Queue struct {
	head   *Node
	tail   *Node
	length int
}

type Hash map[string]*Node

type Cache struct {
	Queue Queue
	Hash  Hash
}

func NewCache() Cache {
	return Cache{Queue: NewQueue(), Hash: Hash{}}
}

func NewQueue() Queue {
	head := &Node{}
	tail := &Node{}
	head.right = tail
	tail.left = head
	return Queue{head: head, tail: tail}
}

func (c *Cache) Check(str string) {
	node := &Node{}
	if val, ok := c.Hash[str]; ok {
		node = c.Remove(val)
	} else {
		node = &Node{val: str}
	}
	c.Add(node)
	c.Hash[str] = node
}

func (c *Cache) Remove(n *Node) *Node {
	fmt.Printf("Removing... %v\n", n.val)
	prev := n.left
	nxt := n.right

	prev.right = nxt
	nxt.left = prev
	c.Queue.length -= 1
	delete(c.Hash, n.val)
	return n
}

func (c *Cache) Add(n *Node) {
	fmt.Printf("Adding... %v\n", n.val)
	tmp := c.Queue.head.right
	c.Queue.head.right = n
	n.right = tmp
	n.left = c.Queue.head.right
	c.Queue.length++

	if c.Queue.length > SIZE {
		c.Remove(c.Queue.tail.left)
	}
}

func (c *Cache) Display() {
	c.Queue.Display()
}

func (q *Queue) Display() {
	n := q.head
	fmt.Printf("\n")
	for i := 0; i < q.length; i++ {
		fmt.Printf("%v <---> ", n.right.val)
		n = n.right
	}
}

func main() {
	fmt.Println("Starting LRU Cache...")
	cache := NewCache()
	for _, word := range []string{"guava", "berries", "slice", "coke", "fanta"} {
		cache.Check(word)
	}
	cache.Display()
}
