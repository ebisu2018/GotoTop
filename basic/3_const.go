package main

import "fmt"

func main4() {

	// 常量必须声明并赋值
	// 常量用大写字母
	const g int = 10
	fmt.Println(g)

	// 即使不用也不会报错
	const (
		PI = 3.14
		E  = 2.75
	)

	// 枚举写法
	const (
		USA   = 1
		CHINA = 2
		Italy = 3
	)


	// iota = 0，依次加1
	const (
		a = iota
		b
		c = 8
		d = iota
		_
		e
	)
	fmt.Println(a, b, c, d, e)

	// 下面的常量会继承上面的公式
	const (
		NOT_USE = 1 << (10 * iota)
		KB  // 1 * 2 ^ 10
		MB  // 1 * 2 ^ 10 ^ 10
		GB  // 1 * 2 ^ 10 ^ 10 ^ 10
	)
	fmt.Println(KB, MB, GB)
}
