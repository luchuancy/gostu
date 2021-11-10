package main

import "fmt"

// 创建单链表结构体
type ListNode struct {
	Val int
	Next *ListNode
}

// 单链表初始化
func main()  {
	head := &ListNode{
		Next: nil,
	}
	head = &ListNode{
		Next: nil,
	}
	fmt.Println(&head.Next)
	fmt.Println(&head.Next)
}


