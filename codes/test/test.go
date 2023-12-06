package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type Linelink struct {
	Head *Node
}

func (node *Linelink) AddNode(data int) {
	newNode := &Node{Data: data}
	if node.Head == nil {
		node.Head = newNode
	} else {
		current := node.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}

}

func (l *Linelink) Traverse() {
	current := l.Head

	for current != nil {
		fmt.Printf("->%d", current.Data)
		current = current.Next
	}

}

func main() {
	Node := &Linelink{}
	Node.AddNode(1)
	Node.AddNode(2)
	Node.AddNode(2)
	Node.AddNode(2)
	Node.AddNode(2)
	Node.AddNode(2)
	Node.Traverse()

}
