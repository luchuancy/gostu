package main

import (
	"errors"
	"fmt"
	"os"
)

// singleQueue
// 数组模拟队列queue，双指针
// 队列使用数组或者链表，先进先出

type Queue struct {
	maxSize int
	array [5]int //数组 模拟队列
	front int //指向队列首, 不含队首
	rear int // 指向队列尾
}

// 添加数据到队列， 方法 
func (this *Queue) AddQueue(val int) (err error) {

	// 先判断队列已满
	if this.rear == this.maxSize - 1 {  // rear指向含有最后元素的队列尾部
		return errors.New("queue full")
	}
	this.rear++  // rear后移
	this.array[this.rear] = val
	return
}

// 从队列取出数据
func (this *Queue) GetQueue() (val int, err error) {
	// 先判断队列空
	if this.rear == this.front { // 队空
		return -1, errors.New("queue empty")
	}
	this.front++ 
	val = this.array[this.front]
	return val, err
}

// 显示队列, 找到队首，遍历至队尾
func (this *Queue) ShowQueue() {
	fmt.Println("队列当前的情况是：")
	for i := this.front + 1; i <= this.rear; i++ {
		fmt.Printf("array[%d]=%d\t", i, this.array[i])
	}
	fmt.Println()
}

func main() {
	//创建队列
	queue := &Queue{
		maxSize: 5,
		front: -1,
		rear: -1,
	}
	var key string 
	var val int 
	for {
		fmt.Println("1. 输入add 表示添加数据到队列")
		fmt.Println("1. 输入get 表示从队列获取数据")
		fmt.Println("1. 输入show 表示显示队列")
		fmt.Println("1. 输入exit 表示退出队列")
   
		fmt.Scanln(&key)
		switch key {
			case "add":
				fmt.Println("输入你要入的队列数")
				fmt.Scanln(&val)
				err := queue.AddQueue(val) 
				if err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Println("加入队列ok")
				}
			case "get":
				val, err := queue.GetQueue()
				if err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Println("从队列取出一个数：", val)
				}
			case "show":
				queue.ShowQueue()
			case "exit":
				os.Exit(0)
		}
	}
}
