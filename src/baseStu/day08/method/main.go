package main
// 方法
import "fmt"

type dog struct {
	name string
}

type person struct {
	name string
	age int
}

//函数
func newDog(name string) dog {
	return dog{
		name: name,
	}
}

func newPerson(name string, age int) person {
	return person {
		name: name,
		age: age,
	}	
}


// 方法作用于特定类型的函数
// 接受者表示的是调用方法的具体类型变量，用类型首字母小写
func (d dog) wang() {
	fmt.Printf("%s: 汪汪", d.name)
}

// 值接收者，接受后是另一个变量
func (p person) guonian() {
	p.age++
}
func (p *person) zhenguonian() {
	p.age++ 
}

func main() {
	d1 := newDog("Cyul")
	d1.wang()
	p1 := newPerson("cy", 18)
	fmt.Println(p1.age)
	p1.guonian()
	fmt.Println(p1.age)
	p1.zhenguonian()  // 实质是传地址， 语法糖
	fmt.Println(p1.age)
}
