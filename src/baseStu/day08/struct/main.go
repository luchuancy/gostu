package main

import "fmt"

type person struct {
	name   string
	gender string
}

// go语言传参是拷贝
func f(x person) {
	x.gender = "女"
}

func f2(x *person) {
	(*x).gender = "女" //根据内存地址找到原变量
}

func main() {
	var p person
	p.name = "Cyul"
	p.gender = "男"
	f(p)
	fmt.Println(p.gender) // 男
	f2(&p)                // 内存地址
	fmt.Println(p.gender) // 女

	var p2 = new(person) // new 指针
	p2.name = "Cyul"
	fmt.Printf("%T\n", p2) // *main.person
	fmt.Printf("%p\n", p2) // 0xc0000a43c0
}
