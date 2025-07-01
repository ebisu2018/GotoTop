package main

import "fmt"

func main3() {

	// 已知excel最大列为XFD，获取是多少列

	// 用%d输出是UTF8的编码
	fmt.Printf("A = %d, Z = %d\n", 'A', 'Z')

	// 获取进制单位，默认excel是26进制
	var base int = 'Z' - 'A' + 1
	fmt.Printf("%d 进制\n", base)

	var total int
	total += 'D' - 'A' + 1
	total += ('F' - 'A' + 1) * base
	total += ('X' - 'A' + 1) * base * base
	fmt.Printf("excel 最大有 %d 列\n", total)
}