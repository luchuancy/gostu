package main
// 变量类型
import "fmt"

func main()  {
	// 十进制
	var i1 = 101
	fmt.Printf("%d\n", i1)
	fmt.Printf("%o\n", i1)   // 10转8
	fmt.Printf("%x\n", i1)   // 10转16
	fmt.Printf("%b\n", i1)   // 10转2
	fmt.Printf("%v\n", i1)   // 查看值，万能
	fmt.Printf("%#v\n", i1)   // 加引号

	fmt.Printf("%s\n", i1)   // 字符串

	// 多行字符串  反引号````
	s1 := `第一行
第二行
第三行
`
	fmt.Println(s1)


	// 八进制
	i2 := 077
	fmt.Printf("%d\n", i2)
	// 十六进制   内存地址0-f
	i3 := 0x344
	fmt.Printf("%d\n", i3)
	fmt.Printf("%T\n", i3)   // 查看类型

	//声明int8类型
	i4 := int8(9)
	fmt.Printf("%T\n", i4)
}