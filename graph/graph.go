package main

import (
	"errors"
	"fmt"
)

type Node struct {
	Key   string
	Value int
	Edges []*Edge
}

type Edge struct {
	Node  *Node
	Value int
}
type Graph struct {
	Nodes []*Node
}

func (g *Graph) Add(key string) error {
	n := g.InGraph(key)
	if n == nil {
		n = &Node{
			Key: key,
		}
		g.Nodes = append(g.Nodes, n)
		return nil
	}
	return errors.New("The node already exists")
}

func (g *Graph) Remove(key string) error {
	n := g.InGraph(key)
	if n != nil {
		var mainRemoved bool
		for e := 0; e <= len(g.Nodes)-1; e++ {
			if g.Nodes[e].Key == n.Key && !mainRemoved {
				g.Nodes = append(g.Nodes[:e], g.Nodes[e+1:]...)
				mainRemoved = false
				e -= 1
			} else {
				for l := range g.Nodes[e].Edges {
					if g.Nodes[e].Edges[l].Node.Key == n.Key {
						g.Nodes[e].Edges = append(g.Nodes[e].Edges[:l], g.Nodes[e].Edges[l+1:]...)
						break
					}
				}
			}
		}
		return nil
	}
	return errors.New("The node doesn't exist")
}

func (g *Graph) InGraph(node string) *Node {
	for n := range g.Nodes {
		if node == g.Nodes[n].Key {
			return g.Nodes[n]
		}
	}
	return nil
}

func (g *Graph) Adjacent(nodeOne string, nodeTwo string) bool {
	n1 := g.InGraph(nodeOne)
	n2 := g.InGraph(nodeTwo)
	if n1 != nil && n2 != nil {
		for n := range n1.Edges {
			if n2.Key == n1.Edges[n].Node.Key {
				return true
			}
		}
	}
	return false
}

func (g *Graph) Neighbors(node string) ([]*Node, error) {
	n := g.InGraph(node)
	var nodes []*Node
	if n != nil {
		for neigh := range n.Edges {
			nodes = append(nodes, n.Edges[neigh].Node)
		}
		return nodes, nil
	}
	return nodes, errors.New("The node doesn't exist")
}

func (g *Graph) Connect(nodeOne string, nodeTwo string, value int) error {
	n1 := g.InGraph(nodeOne)
	n2 := g.InGraph(nodeTwo)
	adj := g.Adjacent(nodeOne, nodeTwo)
	if n1 != nil && n2 != nil && !adj {
		edge1 := &Edge{
			Value: value,
			Node:  n2,
		}
		edge2 := &Edge{
			Value: value,
			Node:  n1,
		}
		n1.Edges = append(n1.Edges, edge1)
		n2.Edges = append(n2.Edges, edge2)
		return nil
	}
	return errors.New("The two nodes do not exist or are already connected")
}

func (g *Graph) Disconnect(nodeOne string, nodeTwo string) error {
	n1 := g.InGraph(nodeOne)
	n2 := g.InGraph(nodeTwo)
	adj := g.Adjacent(nodeOne, nodeTwo)
	if n1 != nil && n2 != nil && adj {
		for e := range g.Nodes {
			if n1.Key == g.Nodes[e].Key || n2.Key == g.Nodes[e].Key {
				for l := range g.Nodes[e].Edges {
					if g.Nodes[e].Edges[l].Node.Key == n2.Key {
						g.Nodes[e].Edges = append(g.Nodes[e].Edges[:l], g.Nodes[e].Edges[l+1:]...)
						break
					} else if g.Nodes[e].Edges[l].Node.Key == n1.Key {
						break
					}
				}
			}
		}
		return nil
	}
	return errors.New("The two nodes do not exist or are not connected")
}

func (g *Graph) SetValue(node string, value int) error {
	n := g.InGraph(node)
	if n != nil {
		n.Value = value
		return nil
	}
	return errors.New("The node doesn't exist")
}

func (g *Graph) GetValue(node string) (int, error) {
	n := g.InGraph(node)
	if n != nil {
		return n.Value, nil
	}
	return 1, errors.New("The node doesn't exist")
}

func (g *Graph) GetEdgeValue(nodeOne string, nodeTwo string) (int, error) {
	n1 := g.InGraph(nodeOne)
	n2 := g.InGraph(nodeTwo)
	adj := g.Adjacent(nodeOne, nodeTwo)
	if n1 != nil && n2 != nil && adj {
		for e := range n1.Edges {
			if n1.Edges[e].Node.Key == n2.Key {
				return n1.Edges[e].Value, nil
			}
		}
		return 1, errors.New("Unexpected error")
	}
	return 1, errors.New("The two nodes do not exist or are not connected")
}

func main() {
	g := new(Graph)
	g.Add("cazzo")
	g.Add("pino")
	g.Add("fuckme")
	//fmt.Println(g.InGraph("pino"))
	//fmt.Println(g.Neighbors("cazzo"))
	g.Connect("cazzo", "pino", 3)
	g.Connect("fuckme", "pino", 2)
	g.Connect("fuckme", "cazzo", 5)
	//fmt.Println(g.Neighbors("cazzo")[0].Key)
	//fmt.Println(g.Neighbors("pino")[0].Value)
	//fmt.Println(g.InGraph("pin"))
	//fmt.Println(g.InGraph("cazzo"))
	fmt.Println(g.InGraph("cazzo"))
	fmt.Println(g.InGraph("pino"))
	//g.Remove("cazzo")
	fmt.Println(g.InGraph("pino"))
	fmt.Println(g.InGraph("cazzo"))
	//g.Disconnect("cazzo", "pino")
	//fmt.Println(g.InGraph("cazzo"))
	//fmt.Println(g.InGraph("pino"))
	//fmt.Println(g.GetValue("cazzo"))
	g.SetValue("cazzo", 2)
	//fmt.Println(g.GetValue("cazzo"))
	fmt.Println(g.GetEdgeValue("cazzo", "pino"))
	//fmt.Println(g.InGraph("cazzo"))
	//fmt.Println(g.Neighbors("pino")[0].Value)
}
