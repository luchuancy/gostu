package main

import (
	"fmt"
)
// 浮点数， 布尔值
func main() {
	// math.MaxFloat32   // float32 最大值
	f1 := 1.2345
	fmt.Printf("%T\n", f1) // go默认float64
	f2 := float32(1.234)
	fmt.Printf("%T\n", f2)   // f1 和 f2  不相等

	// 布尔类型变量的默认值为false。

}