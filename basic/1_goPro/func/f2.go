package main

import "fmt"

func main() {

	c, d := 10, 20
	fmt.Printf("main &c = %p &d = %p\n", &c, &d)

	// fmt.Println(fn_val(c, d))

	fmt.Println(fn_ptr(&c, &d))
	fmt.Printf("main c = %v, d = %v\n", c, d)
}

func fn_val(c, d int) (int, int) {

	// 内存中复制了一份并返回，不影响原来的变量
	fmt.Printf("val inner &c = %p &d = %p\n", &c, &d)
	c += 100
	d += 100
	return c, d
}


// 传递的是指针
func fn_ptr(c, d *int) (*int, *int) {
	fmt.Printf("ptr inner &c = %p &d = %p\n", c, d)
	*c *= 10
	*d *= 20
	return c, d
}
