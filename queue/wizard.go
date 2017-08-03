package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
func (q *LinkedQ) insert(value int, index int) {
	new := &Node{
		Value: value,
	}
	n := q.First
	for i := 0; i <= index-2; i++ {
		n = n.Next
	}
	new.Next = n.Next
	n.Next = new
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input order of the cards (split numbers by whitespaces) : ")
	userOrder, _ := reader.ReadString('\n')
	userOrder = strings.Split(userOrder, " ")
	q := newLinked()
	for x := range userOrder {
		q.enqueue(strconv.Atoi(x))
	}
	fmt.Println("Order : " + q.getQueue())
	var table []int
	for {
		if q.isEmpty() {
			break
		}
		n := q.unqueue()
		q.enqueue(n)
		n := q.unqueue()
		table = append(table, n)
	}
	fmt.Println("Deck is empty, table cards : ", table)

}
