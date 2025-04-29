package main

import "fmt"

func main6() {
	a := 200
	fmt.Printf("%b\n", a)

	// 判断第4位是否为1
	c := 1 << 3

	r := a & c
	fmt.Printf("%b & %b = %v\n", a, c, r)
	binFunc(200)
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