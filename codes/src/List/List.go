package List

import "fmt"

type Node struct {
	Val int
	Next *Node
}

// 定义链表结构
type LinkedList struct {
	Head *Node
}

// 向链表末尾添加新节点
func (list *LinkedList) Append(Val int) {
	newNode := &Node{Val: Val, Next: nil}

	if list.Head == nil {
		list.Head = newNode
	} else {
		current := list.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
}

// 遍历并打印链表
func (list *LinkedList) Display() {
	current := list.Head
	for current != nil {
		fmt.Printf("%d -> ", current.Val)
		current = current.Next
	}
	fmt.Println("nil")
}