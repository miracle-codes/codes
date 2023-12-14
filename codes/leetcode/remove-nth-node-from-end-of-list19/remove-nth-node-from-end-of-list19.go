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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	fast := head
	cur := dummyHead

	for i := 1; i < n; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return dummyHead.Next
}

func main() {
	list := &ListNode{}
	list.addNode(1)
	list.addNode(2)
	list.addNode(3)
	list.addNode(4)
	list.addNode(5)
	list.Traverse()
	fmt.Println()
	list = removeNthFromEnd(list, 2)

	list.Traverse()
}
