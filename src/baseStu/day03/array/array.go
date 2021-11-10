package main

import "fmt"

func main()  {
	// 数组初始化
	// var a []int  nil没有内存空间
 	var a [3]int
	var b = [4]int{}
	var c = [...]int{1, 2, 3}
	d := [...]int{1: 1, 3: 5}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	// 数组遍历
	var a1 = [...]string{"北京", "上海", "深圳"}
	// 方法1：for循环遍历
	for i := 0; i < len(a); i++ {
		fmt.Println(a1[i])
	}

	// 方法2：for range遍历
	for index, value := range a1 {
		fmt.Println(index, value)
	}

}
