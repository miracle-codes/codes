package main

import (
	"container/list"
	"fmt"
)

func main() {
	//通过 list.New 创建列表
	listHaiCoder := list.New()
	listHaiCoder.PushBack("Hello")
	listHaiCoder.PushBack("HaiCoder")

	for i := listHaiCoder.Front(); i != nil; i = i.Next() {
		fmt.Println("Element =", i.Value)
	}
}
