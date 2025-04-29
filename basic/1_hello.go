package main

import "fmt"


// 大写是全局可见，小写是包内可见
func HelloWorld()  {
	fmt.Println("Hello World")
}

func main1()  {
	fmt.Println("Hello World")
}