package main

import (
	"fmt"
	"strings"
)

//Max length for items
const m int = 100000

//Node object for the LinkedQ
type Node struct {
	Item *Item
	Next *Node
}

//LinkedQ object
type LinkedQ struct {
	First *Node
	Last  *Node
}

//HashTable object
type HashTable struct {
	Size int
	Item []*Items
}

//Items object (contains a list of empty slots for the Item)
type Items struct {
	Length int
	Li     *LinkedQ
}

//Item object
type Item struct {
	Key   string
	Value interface{}
}

//Constructor for LinkedQ
func NewLinked() *LinkedQ {
	q := new(LinkedQ)
	q.First = nil
	q.Last = nil
	return q
}

//Constructor of the HashTable
func newTable(size int) *HashTable {
	return &HashTable{
		Size: size,
		Item: make([]*Items, m),
	}
}

//Enqueue function
func (q *LinkedQ) enqueue(item *Item) {
	n := &Node{
		Item: item,
	}
	if q.First == nil {
		q.First = n
		q.Last = n
	} else {
		q.Last.Next = n
		q.Last = n
	}
}

//Dequeue function
func (q *LinkedQ) dequeue() *Node {
	n := q.First
	q.First = q.First.Next
	return n
}

//Hash function for the HashTable
func (t *HashTable) Hash(key string) int {
	sum := 0
	str := strings.ToUpper(key)
	for i := 0; i < len(str); i++ {
		sum += int(str[i])
	}
	return sum % m
}

//Add Item to HashTable
func (t *HashTable) Add(key string, value interface{}) {
	index := t.Hash(key)
	if t.Item[index] == nil {
		t.Item[index] = &Items{
			Length: 0,
			Li:     NewLinked(),
		}
		t.Item[index].Li.enqueue(&Item{
			Key:   key,
			Value: value,
		})
		t.Item[index].Length = 1
	} else {
		t.Item[index].Li.enqueue(&Item{
			Key:   key,
			Value: value,
		})
		t.Item[index].Length += 1
	}
}

//Get an Item using a key
func (t *HashTable) Get(key string) interface{} {
	index := t.Hash(key)
	l := t.Item[index].Li
	for s := l.First; s != nil; s = s.Next {
		o := s.Item
		if o.Key == key {
			return o.Value
		}
	}
	return nil
}

func main() {
	s := newTable(100)
	s.Add("test", 100)
	fmt.Println(s.Get("test"))
	s.Add("cube", 90)
	fmt.Println(s.Get("cube"))
}
