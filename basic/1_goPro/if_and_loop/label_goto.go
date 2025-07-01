package main

import "fmt"

func main3() {
	label1()
}

func label1() {
	i := 4

	//定义标签
LABEL1:
	i += 3
	i *= 2
	fmt.Println(i)
	if i >= 100 {
		// 当大于等于100，强制退出函数
		return
	}

	// 使用标签
	goto LABEL1
}