package main

import (
	"fmt"
)

func main() {
	switch_1()
	switch_2()
}

func switch_1() {
	color := "red"

	// 对静态值的判断，是一个固定的值
	switch color {
	case "red":
		fmt.Println("red")
	case "yellow", "green":  // 或的关系
		fmt.Println("yellow")
	default:
		fmt.Println("other")
	}
}


func switch_2()  {
	a := 10

	// 对表达式判断，switch后面什么都没有，在case中判断
	switch {
	case a > 10:
		fmt.Println("a is greater than 10")
		fallthrough
	case a == 10:
		fmt.Println("a is equal to 10")
	case a < 10:
		fmt.Println("a is less than 10")
	}
}