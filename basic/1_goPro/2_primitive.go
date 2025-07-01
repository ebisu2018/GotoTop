package main

import (
	"fmt"
)

func main2() {

	// 和计算机操作系统一致，默认64位为int64
	// 但是int和int64本质是不同类型
	var x int
	fmt.Println(x)

	var a int = 10
	fmt.Println(a)

	// 短格式，声明并赋值，不用var关键字
	b := a


	// 使用_占位符，表示不使用该变量，用_接受
	_ = b

	var (
		c int
		d float64
		g bool = true
	)


	// 暂时不使用
	_, _ = c, g

	d = 1.58786898
	fmt.Printf("d = %f\n", d)
	fmt.Printf("d = %.2f\n", d)
	fmt.Printf("d = %g\n", d)
	fmt.Printf("d = %e\n", d)
	fmt.Printf("d = %[1]f, a = %[2]d, d = %[1]e, a = %[2]d\n", d, a)

	// 进制，默认十进制输出
	e := 0o10
	f := 0xff
	fmt.Printf("e = %d, f = %d\n", e, f)

	fmt.Println(1_000_000_000)

	// 小数则默认float64
	k := 35.
	fmt.Printf("%f\n", k)
}