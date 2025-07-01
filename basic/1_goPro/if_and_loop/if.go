package main

import "fmt"

func main1() {

	var a int = 10

	// 在if内设置局部变量，只在if内可见
	if b := 100; b > a {
		fmt.Printf("b %v greater than a\n", b)
	} else {
		fmt.Printf("a greater than %v\n", b)
	}
}