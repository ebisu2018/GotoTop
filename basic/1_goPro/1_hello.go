// package名称可以和目录一致，也可以不一致
package main

import "fmt"


// 大写是全局可见，小写是包内可见
func HelloWorld()  {
	fmt.Println("Hello World")
}

// 包中有且只有一个main函数
// package为main，并且main()是程序的入口
func main1()  {
	fmt.Println("Hello World")
}