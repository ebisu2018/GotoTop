package main

import "fmt"

func main6() {

	// 使用二进制表示
	a := 200
	fmt.Printf("%b\n", a)

	// 判断第4位是否为1
	// 1 向左移3位，其他都是0
	c := 1 << 3
	r := a & c
	fmt.Printf("%b & %b = %v\n", a, c, r)

	binFunc(10)
}


func binFunc(a uint64)  {
	
	var c uint64 = 1 << 63

	for i := 0; i < 64; i++ {
		if a & c == c {
			fmt.Print("1")
		} else {
			fmt.Print("0")
		}
		c = c >> 1
	}
}