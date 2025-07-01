package main

import "fmt"

// 外部变量对内部作用域穿透可见
// 内部作用域对外不可见

// 外部定义的为全局作用域，向内穿透，局部作用域都可见
var G = "global var"

func fun1() {
	fmt.Printf("G = %v\n", G)
}

func main5() {
	// 大括号为定义一个作用域，变量常量仅在{}内部可见
	{
		const (
			PI = 3.14
			G  = "golang" // 就近原则，和全局同名，则优先使用自己的
		)
		fmt.Printf("PI = %v, G = %v\n", PI, G)
	}
	{
		const (
			PI = 5.18
		)
		fmt.Printf("PI = %v, G = %v\n", PI, G)
	}
	fun1()
}
