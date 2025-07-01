package main

import "fmt"

func main2() {

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}


	a, b := 10, 30
	for a < 15 && b > 15  {
		fmt.Printf("a = %v, b = %v\n", a, b)
		a++
		b--
	}

	// for _, v := range v {

	// }


	// 死循环，使用break中断
	c := 0
	for {
		c++
		fmt.Println("死循环")
		if c == 10 {
			break
		}
	}
}