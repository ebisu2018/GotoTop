package main

import "fmt"

func main() {

	// 官方不承认继承，只是其他语言的叫法
	type user struct {
		name string
		age int
	}

	// 有名字的成员变量，使用指针类型
	type login struct {
		user     *user
		password string
	}

	
	type account struct {
		*user // 匿名成员
		name string // 如果和结构体成员重名，默认指当前自己的变量
		region string
	}

	u := user{
		name: "u1",
		age: 35,
	}
	
	l := login{
		user: &u,
	}

	fmt.Printf("%v, %v\n",l.user.name, l.user.age)

	a := account{
		user: &u,
		name: "account-name",
	}

	// 因为重名，需要显示调用匿名成员名字
	fmt.Printf("%v\n",a.name)
	fmt.Printf("%v\n", a.user.name)

	// 匿名成员直接调用
	fmt.Printf("%v\n", a.age)

}
