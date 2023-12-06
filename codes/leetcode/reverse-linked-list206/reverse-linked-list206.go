package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) AddNode(data int) {
	newNode := &ListNode{Val: data}
	if l == nil {
		l = newNode
		return
	}
	for l.Next != nil {
		l = l.Next
	}
	l.Next = newNode

	return
}

func (l *ListNode) Traverse() {
	for l != nil {
		fmt.Printf("->%d", l.Val)
		l = l.Next
	}
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	current := head.Next
	var pre *ListNode
	for current != nil {
		next := current.Next
		current.Next = pre
		pre = current
		current = next

	}
	return pre

}

func main() {
	list := ListNode{}
	list.AddNode(1)
	list.AddNode(2)
	list.AddNode(3)
	list.AddNode(4)
	list.Traverse()

	res := reverseList(&list)
	fmt.Println()
	//遍历量表
	//list.Traverse()

	res.Traverse()
}
