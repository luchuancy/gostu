package main

import "fmt"

func main()  {
	n :=18
	// 取地址
	p := &n
	// 根据地址取值
	m := *p
	fmt.Println(p)
	fmt.Printf("%T\n", p)
	fmt.Println(m)
	fmt.Printf("%T\n", m)
	fmt.Println("-----------------------------")
	//var a *int   // nil
	//*a = 100
	//fmt.Println(*a)
	var a = new(int)        // 定义新指针
	fmt.Println(a)
	*a = 100
	fmt.Println(*a)
}
