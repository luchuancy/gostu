package main

import "fmt"

//声明变量 函数 等
var name string
var age int
var isOk bool
// 批量声明
var (
	a string
	b int
	c bool
	d float32
)

const (
	pi = 3.1415
	e = 2.7182
)
const (
	n1 = 100
	n2
	n3
)    //n1 n2 n3  都是100

const (
	s1 = iota //0
	s2        //1
	s3        //2
	s4        //3
)

const (
	_  = iota
	KB = 1 << (10 * iota)   // 1 <<左移10位 1024
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main()  {
	name = "Cyul"
	age = 18
	isOk = true
	var s0 string = "haha"

	// 匿名变量用下划线 _  如x, _ = fun(t)

	fmt.Println(age)  //后面换行
	fmt.Print(isOk)     //打印
	fmt.Printf("name:%s\n", name)   //%s 占位符
	fmt.Println(s0)
	fmt.Println(s3)
}