package main

import "fmt"

func main() {
	//append()添加元素和切片扩容
	// append()只能用原来的切片变量接收
	var numSlice []int
	for i := 0; i < 10; i++ {
		numSlice = append(numSlice, i)
		fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n",
			numSlice, len(numSlice), cap(numSlice), numSlice)
	}

	// copy()复制切片
	a := []int{1, 2, 3, 4, 5}
	b := a    // 赋值地址    引用类型
	c := make([]int, 5, 5)
	copy(c, a)     //使用copy()函数将切片a中的元素复制到切片c    值类型
	fmt.Println(a) //[1 2 3 4 5]
	fmt.Println(b) //[1 2 3 4 5]
	fmt.Println(c) //[1 2 3 4 5]
	b[0] = 1000
	fmt.Println(a) //[1000 2 3 4 5]
	fmt.Println(b) //[1000 2 3 4 5]
	fmt.Println(c) //[1 2 3 4 5]

	// 从切片中删除元素    a1 := []int{30, 31, 32, 33, 34, 35, 36, 37}  切片
	a1 := [...]int{30, 31, 32, 33, 34, 35, 36, 37}  // 数组
	s1 := a1[:]     // 切片
	// 要删除索引为2的元素
	s1 = append(s1[:2], s1[3:]...)   // ...表示拆开{33, 34, 35, 36, 37} >> {33}, {34}, {35}, {36}, {37}
									// 追加必须拆开
	fmt.Println(s1) //[30 31 33 34 35 36 37]
	fmt.Println(a1) //[30 31 33 34 35 36 37 37]
}
