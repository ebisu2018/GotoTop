package main

import "fmt"

func main4() {

	// 常量用const表示
	// 常量必须声明并赋值
	// 常量一般用大写字母
	// 可以不写类型，让编译器类型推断

	// 即使不用也不会报错，和var不同点
	const (
		PI = 3.14
		E  = 2.75
	)

	// 枚举
	const (
		USA   = 1
		CHINA = 2
		Italy = 3
	)

	fmt.Println(USA, CHINA, Italy)


	// iota = 0，依次加1
	const (
		a = iota
		b
		c = 8
		d = iota
		_ // iota直接跳过
		e
	)
	fmt.Println(a, b, c, d, e)

	// 下面的常量会继承上面的表达式
	const (
		NOT_USE = 1 << (10 * iota)
		KB  // 1 * 2 ^ 10
		MB  // 1 * 2 ^ 10 ^ 10
		GB  // 1 * 2 ^ 10 ^ 10 ^ 10
	)
	fmt.Println(KB, MB, GB)
}
