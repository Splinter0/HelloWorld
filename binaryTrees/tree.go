package main

import "fmt"

//Constructors needed

//Node : struct
type Node struct {
	Left  *Node
	Right *Node
	Value int
}

//Tree : struct
type Tree struct {
	Root *Node
}

func treeConts() *Tree {
	t := new(Tree)
	t.Root = nil
	return t
}

//Functions
func printOut(root *Node) {
	if root != nil {
		printOut(root.Left)
		fmt.Println(root.Value)
		printOut(root.Right)
	}
}
func exists(root *Node, value int) bool {
	if root != nil {
		if root.Value != value {
			if root.Value > value {
				return exists(root.Left, value)
			} else {
				return exists(root.Right, value)
			}
		} else {
			return true
		}
	} else {
		return false
	}
}
func addToTree(root *Node, value int) *Node {
	if root == nil {
		n := &Node{
			Value: value,
		}
		return n
	}
	if exists(root, value) {
		return root
	}
	if root.Value < value {
		root.Right = addToTree(root.Right, value)
	} else {
		root.Left = addToTree(root.Left, value)
	}
	return root
}

//Methods
func (t *Tree) put(value int) {
	t.Root = addToTree(t.Root, value)
}
func (t *Tree) write() {
	printOut(t.Root)
}

func main() {
	t := treeConts()
	t.put(3)
	t.put(5)
	t.put(2)
	t.put(7)
	t.write()
}
