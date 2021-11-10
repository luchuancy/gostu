package main

import "fmt"
// 函数的传参是值类型
func f(x int) {     // 参数进来只是把值传进来
	x++
}
func f2(x *int) {
	(*x)++          //根据内存地址找到原变量  语法糖等价  x++
}
func f3(x int) int {    // 有返回值
	x++
	return x
}

func main() {
	a := 3
	f(a)
	fmt.Println(a)   // 3
	f2(&a)                // 内存地址
	fmt.Println(a)   // 4
	f3(a)
	fmt.Println(a)   // 4
}