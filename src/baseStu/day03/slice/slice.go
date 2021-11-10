package main

import "fmt"

func main()  {
	a := [...]int{1, 3, 3, 4, 5, 5, 2, 3, 4}
	s := a[3:6]
	fmt.Println(s)
	fmt.Println(a[:4])
	// 切片指向底层数组
	// 切片的长度就是切片数组长度
	// 切片的容量就是底层数组从第一个元素往后数
	// 切片可以扩容
	// 切片是引用类型   可以理解为切片的是存储位置，而不是具体的值  底层数组改变切片也改变
	fmt.Printf("len(s):%d cap(s):%d", len(s), cap(s))
	// make 制造切片
	s1 := make([]int, 5, 10)
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d", s1, len(s1), cap(s1))
}
