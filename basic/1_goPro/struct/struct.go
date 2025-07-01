package main

import (
	"fmt"
)

// 结构体声明
// 结构体是一个类型
type Person struct {
	Name string // 成员变量，没有逗号
	Age  int
}


// 结构体方法
func (p *Person) hello()  {

	// 通过指针取值赋值
	fmt.Printf("hello, %v\n", p.Name)
}

func main1() {

	// 结构体赋值，使用var可以类型推断
	var p1 = Person{
		Name: "p1",
		Age:  20, // 必须结尾有逗号
	}

	fmt.Println(p1)

	// 可以不使用字段名，不推荐，表述不清
	p2 := Person{"p2", 30}
	fmt.Println(p2)

	// 取地址
	p3 := &p2
	fmt.Printf("%T\n", p3)

	// 通过*获取值
	fmt.Printf("%v\n", *p3)

	// p3和p2地址相同，会连带修改p2
	p3.Name = "p3"
	fmt.Println(p2)
	fmt.Printf("%v\n", *p3)


	// 指针类型
	p4 := new(Person)
	fmt.Printf("%T\n", p4)

	// 通过地址可以赋值, Go语言做了优化，可以通过指针赋值
	p4.Name = "p4"
	fmt.Println(*p4)

	// 通过地址也可以调用方法
	p4.hello()


	// 匿名结构体
	var u struct {
		Name string
	}
	u.Name = "anoy"
	fmt.Printf("匿名结构体 %v\n", u)
}