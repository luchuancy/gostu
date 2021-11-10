package main

import "fmt"

func main()  {
	var m1 map[string]int   //nil
	m1 = make(map[string]int, 10)   // 估算容量   避免动态扩容
	m1["理想"] = 18
	fmt.Println(m1)
	// ok 接受布尔值
	value, ok := m1["李想"]
	if !ok {
		fmt.Println("无")
	} else {
		fmt.Println(value)
	}
	fmt.Println(m1["李想"])   // 输出0值
	fmt.Println("-------------------------")
	// map 遍历
	for v, k := range m1 {
		fmt.Println(v)
		fmt.Println(k)
	}
	// map 删除
	delete(m1, "理想")
	fmt.Println(m1)
}
