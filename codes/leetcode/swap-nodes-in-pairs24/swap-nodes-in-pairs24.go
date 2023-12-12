package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) addNode(data int) {
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
	if l == nil {
		fmt.Println("list is nil")
	}
	for l != nil {
		fmt.Printf("->%d", l.Val)
		l = l.Next
	}
	return

}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy

	for head != nil && head.Next != nil {
		pre.Next = head.Next
		next := head.Next.Next
		head.Next.Next = head
		head.Next = next

		pre = head
		head = next

	}

	return dummy.Next

}

func main() {

	//var list *ListNode
	list := &ListNode{}
	list.addNode(1)
	list.addNode(2)
	list.addNode(3)
	list.addNode(4)

	list.Traverse()
	fmt.Println()
	list = swapPairs(list)
	list.Traverse()
}
