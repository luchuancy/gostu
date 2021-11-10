package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}
//& 取地址
func makeListNode(arr []int) *ListNode {
	// 处理空数组
	if len(arr) == 0 {
		return nil
	}
	// 定义一个哑结点
	dummyHead := &ListNode {
		Next: nil,
	}
	// 定义指针，遍历数组，加入链表
	temp := dummyHead
	for _, arrEle := range arr {
		temp.Next = &ListNode {
			Val: arrEle,
		}
 		temp = temp.Next
	}
	return dummyHead.Next
}


func main() {
	var head *ListNode
	array := []int{1, 2, 3, 7}
	head = makeListNode(array)
	fmt.Println(head) // &{1 0xc00003c240}
	fmt.Println(head.Next) // &{2 0xc00003c250}
	fmt.Println(head.Next.Next) // &{3 0xc00003c260}
	fmt.Println(head.Next.Next.Next) // &{7 <nil>}
}