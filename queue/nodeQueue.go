package main

import (
	"fmt"
	"strconv"
)

//Node : struct
type Node struct {
	Value int
	Next  *Node
}

//LinkedQ : struct
type LinkedQ struct {
	First *Node
	Last  *Node
}

func newLinked() *LinkedQ {
	q := new(LinkedQ)
	q.First = nil
	q.Last = nil
	return q
}

func (q *LinkedQ) enqueue(value int) {
	n := &Node{
		Value: value,
	}
	//fmt.Println("test")
	if q.First == nil {
		q.First = n
		q.Last = n
	} else {
		q.Last.Next = n
		q.Last = n
	}
}
func (q *LinkedQ) unqueue() int {
	n := q.First
	q.First = q.First.Next
	return n.Value
}
func (q *LinkedQ) isEmpty() bool {
	return q.First == nil
}
func (q *LinkedQ) getQueue() string {
	var queue string
	n := q.First
	for {
		queue = queue + strconv.Itoa(n.Value) + " "
		n = n.Next
		if n == nil {
			break
		}
	}
	return queue
}
func main() {
	q := newLinked()
	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(7)
	fmt.Println(q.getQueue())
	q.unqueue()
	fmt.Println(q.getQueue())
}
