package main

import (
	"fmt"
	"time"
)

func main1() {
	fn1()
	fn2(100)
	
	a := 10
	b := 20

	var c int
	c = fn3(a, b)

	fmt.Printf("c = %v\n", c)
	fmt.Printf("a = %v, b = %v\n", a, b)

	fmt.Println(fn4(a, b))
	fmt.Println(fn5(10, 8))

	fmt.Println(multiReturn())
}


// 无参数无返回值
func fn1() {
	fmt.Println("无参无返回值函数")
	// return 函数已经返回
	// fmt.Println("不执行")
}


// 有参数无返回值函数
func fn2(a int)  {
	fmt.Printf("a = %v\n", a)
}


// 参数类型一致，写一个
func fn3(a, b int) int {
	a = a + b
	return a
}


// 有返回值，并且已定义好返回值变量直接返回
func fn4(a, b int) (c int) {
	c = a * b
	return 
}


// 只定义返回类型，变量需要定义并返回
// 只返回一个，不需要小括号
func fn5(a, b int) int {
	c := a / b
	return c
}


// 多个返回值，需要包含小括号
func multiReturn() (int, int)  {
	now := time.Now()
	return now.Hour(), now.Minute()
}