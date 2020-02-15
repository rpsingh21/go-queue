package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// Node contain json document
type Node struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func (n *Node) String() string {
	doc, err := json.Marshal(n)
	if err != nil {
		panic(err)
	}
	return fmt.Sprint(string(doc))
}

// Queue type of custom queue
type Queue struct {
	nodes []*Node
	size  int
	head  int
	tail  int
	count int
}

// Push handel add element in queue(Circuler)
func (q *Queue) Push(n *Node) {
	if q.head == q.tail && q.count > 0 {
		nodes := make([]*Node, len(q.nodes)+q.size)
		copy(nodes, q.nodes[q.head:])
		copy(nodes[len(q.nodes)-q.head:], q.nodes[:q.head])
		q.head = 0
		q.tail = len(q.nodes)
		q.nodes = nodes
	}
	q.nodes[q.tail] = n
	q.tail = (q.tail + 1) % len(q.nodes)
	q.count++
}

// Pop elenmen for queue
func (q *Queue) Pop() *Node {
	if q.count == 0 {
		panic("Queue is empty")
	}
	node := q.nodes[q.head]
	q.head = (q.head + 1) % len(q.nodes)
	q.count--
	return node
}

// NewQueue return queue eleement
func NewQueue(size int) *Queue {
	return &Queue{
		nodes: make([]*Node, size),
		size:  size,
	}
}

func main() {
	q := NewQueue(10)
	pushDur, _ := time.ParseDuration("80ms")
	popDur, _ := time.ParseDuration("100ms")
	// pushChan := make(chan Node)
	// popChan := make(chan bool)

	// Loop for push value
	id := 1
	go func() {
		for true {
			q.Push(&Node{id, "testing"})
			// pushChan <- Node{id, "testing"}
			id++
			time.Sleep(pushDur)
		}
	}()

	// Loop for pop value
	go func() {
		for true {
			fmt.Println(q.Pop())
			// popChan <- true
			time.Sleep(popDur)
		}
	}()

	mainF, _ := time.ParseDuration("5s")
	time.Sleep(mainF)
}
