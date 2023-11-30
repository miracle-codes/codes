package main

import "fmt"

// 定义链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) addNode(val int) *ListNode {
	if head == nil {
		return &ListNode{
			val,
			nil,
		}
	}
	for head.Next != nil {
		head = head.Next
	}
	head.Next = &ListNode{
		val, nil,
	}
	return head

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	for head.Next != nil {
		if head.Next.Val == val {
			head.Next = head.Next.Next
		}
		head = head.Next
	}

	return head
}

func main() {
	mylist := &ListNode{}
	mylist.addNode(1)
	mylist.addNode(2)
	mylist.addNode(3)
	mylist.addNode(4)
	mylist.addNode(5)
	mylist.addNode(6)

	for mylist.Next != nil {
		fmt.Println(mylist.Next.Val)
		mylist = mylist.Next
	}

	//myList := &list.LinkedList{}
	//myList.Append(1)
	//myList.Append(2)
	//myList.Append(3)
	//myList.Append(4)
	//myList.Display()
	//res := removeElements(myList, 2)
	//res.Display()
	//myList.Append(1)
	//myList.Append(2)
	//myList.Append(3)
	//myList.Append(4)
	//myList.Append(5)
	//myList.Append(6)
	//myList = removeElements(myList,1)
	//myList.Display() // 输出: 1 -> 2 -> 3 -> nil
}

//func removeElements(head *list.LinkedList, val int) *list.LinkedList {
//	//dummyHead := &list.LinkedList{}
//	dummyHead := head
//	cur := dummyHead.Head
//	for cur != nil && cur.Next != nil {
//		if cur.Next.Val == val {
//			cur.Next = cur.Next.Next
//		} else {
//			cur = cur.Next
//		}
//	}
//	return dummyHead
//
//}
