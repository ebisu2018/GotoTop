package main

import "fmt"

func main() {
	fn1()
	fn2(100)
	
	a := 10
	b := 20

	var c int
	c = fn3(a, b)

	fmt.Printf("c = %v\n", c)
	fmt.Printf("a = %v, b = %v\n", a, b)

	fmt.Println(fn4(a, b))

	fmt.Println(fn5(&a, &b))
	fmt.Printf("a = %v, b = %v\n", a, b)
	
}

func fn1() {
	fmt.Println("无参无返回值函数")
	// return
	// fmt.Println("不执行")
}


// 有参数无返回值函数
func fn2(a int)  {
	fmt.Printf("a = %v\n", a)
}

func fn3(a, b int) int {
	a = a + b
	return a
}


func fn4(a, b int) (c int) {
	c = a * b
	return c
}


func fn5(a, b *int) int {
	*a = *a + * b
	return *a
}
