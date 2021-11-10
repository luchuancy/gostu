package main

// 稀疏数组
import (
	"fmt"
)

type ValNode struct {
	row int
	col int
	val int
}

func main() {
	// 1. 创建原始数据
	var chessMap [11][11]int
	chessMap[1][2] = 1 // 黑子
	chessMap[2][3] = 2 // 白子
	// 2. 输出原始数据
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
	// 3. 转成稀疏数组
	var sparseArr []ValNode
	// 标准的稀疏数组 记录元素的二维数组规模 行列及默认值
	valNode := ValNode{
		row: 11,
		col: 11,
		val: 0,
	}
	sparseArr = append(sparseArr, valNode)

	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				// 创建一个val值结点
				valNode := ValNode{
					row: i,
					col: j,
					val: v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}
	// 输出稀疏数组
	fmt.Println("稀疏数组为：", sparseArr)
	for i, valNode := range sparseArr {
		fmt.Printf("%d: %d %d %d\n", i, valNode.row, valNode.col, valNode.val)
	}
	// 将稀疏数组存盘

	// 恢复原始数组
	// 打开存盘文件，恢复数组
	// 直接使用稀疏数组恢复原始数组
	var chessMap2 [11][11]int
	// 遍历sparseArr
	for i, valNode := range sparseArr {
		if i != 0 { // 跳过第一行记录的数据
			chessMap2[valNode.row][valNode.col] = valNode.val
		}
	}
	fmt.Println("恢复后的原始数据：", chessMap2)
	for _, v := range chessMap2 {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
}
