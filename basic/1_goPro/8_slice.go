package main

/*
slice 是一个包含array地址，len，cap的结构体
*/

import "fmt"

func main8() {
	// slice1()
	expansionCap()
}

func slice1() {

	s1 := []int{0, 0, 0, 0, 0}
	fmt.Printf("s1 %v len %v cap %v\n", s1, len(s1), cap(s1))

	// s2 := make([]int, 10)
	s2 := make([]int, 3, 5)
	fmt.Printf("s2 %v len %v cap %v\n", s2, len(s2), cap(s2))

	// slice自身地址是slice的header的地址
	fmt.Printf("s2 address %p, \n", &s2)

	// slice用p打印就是数组的地址，和首元素地址相同
	fmt.Printf("array address %p %p\n", s2, &s2[0])

	s3 := [][]int{
		{1},
		{2, 5},
	}
	fmt.Printf("s3 = %v\n", s3)
}


// 扩容逻辑，之前是2倍的扩，之后是1.5倍扩容
func expansionCap()  {
	
	s := make([]int, 0, 3)
	prevCap := cap(s)
	for i := 0; i < 100; i++ {
		s = append(s, i)
		currCap := cap(s)
		if currCap > prevCap {
			fmt.Printf("从%d扩容到%d\n", prevCap, currCap)
			prevCap = currCap
		}
	}
}