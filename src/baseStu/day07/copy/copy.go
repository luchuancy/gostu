package main

import "fmt"

func main()  {
	a := [...]int{0, 1, 2, 3, 4, 5, 6}
	b := a[0:3]
	c := make([]int, 2, 2)
	copy(c, b)
	fmt.Printf("b: %p\n", b)
	fmt.Printf("b[1]: %p\n", &b[1])
	fmt.Printf("c: %p\n", c)
	fmt.Printf("c[1]: %p\n", &c[1])
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	a[0] = 666
	fmt.Printf("%p\n", c)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}
