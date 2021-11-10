package main

import "fmt"

func main()  {
	//s := "陈年印象"
	//for i, v := range s {
	//	fmt.Printf("%d %c\n", i, v)
	//}
	finger := 3
	switch finger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效的输入！")
	}
	// 每个switch只能有一个default分支。   case  只对值进行判断
	// case 1, 3, 5, 7, 9:
	// case age < 25:

	// 	case s == "a":
	//		fmt.Println("a")
	//		fallthrough
	//	case s == "b":
	//		fmt.Println("b")
	// fallthrough语法可以执行满足条件的case的下一个case
}
