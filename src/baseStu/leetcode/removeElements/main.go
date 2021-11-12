//package main
//
//import "fmt"
//
//func main()  {
//	type ListNode struct {
//		Val  int
//		Next *ListNode
//	}
//	v = []int{1, 2}
//	ListNode{1,v}
//	fmt.Println(head)
//
//}
package main

import "fmt"
type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	a1 := []int{1, 2}
	// head := &ListNode
	// fmt.Println(head)
	fmt.Println(a1)
}