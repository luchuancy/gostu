package main

import "fmt"

func main() {
	s1 := "big"
	// 强制类型转换
	byteS1 := []byte(s1)    // 英文
	byteS1[0] = 'p'          // 字符
	fmt.Println(string(byteS1))

	s2 := "白萝卜"
	runeS2 := []rune(s2)    // 中文等其他语言
	runeS2[0] = '红'
	fmt.Println(string(runeS2))
}
