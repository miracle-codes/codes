package main


import list "../../src/List"
// 定义链表节点
type ListNode struct {
    Val int
    Next *ListNode
}

func main() {
    myList := &list.Node{}

    //myList.Append(1)
    //myList.Append(2)
    //myList.Append(3)
    //myList.Append(4)
    //myList.Append(5)
    //myList.Append(6)
    //myList = removeElements(myList,1)
    //myList.Display() // 输出: 1 -> 2 -> 3 -> nil
}

func removeElements(head *list.Node, val int) *ListNode {
    dummyHead := &list.Node{}
    dummyHead = head
    cur := dummyHead
    for cur != nil && cur.Next != nil {
        if cur.Next.Val == val {
            cur.Next = cur.Next.Next
        } else {
            cur = cur.Next
        }
    }
    return dummyHead.Next


}