package main

import "fmt"

func main7() {
	// arr1d()
	arr2d()
}

func arr1d() {

	var arr1 = [5]int{}
	var arr2 = [5]int{2, 6}
	var arr3 = [5]int{2: 10, 4: 100}
	var arr4 = [...]int{3, 2, 8, 7, 9, 10, 18}
	var arr5 = [...]struct {
		name string
		age  int
	}{{"James", 18}, {"Tyler", 24}}

	fmt.Printf("arr1 = %v\n", arr1)
	fmt.Printf("arr2 = %v\n", arr2)
	fmt.Printf("arr3 = %#v\n", arr3)
	fmt.Printf("arr4 = %#v\n", arr4)
	fmt.Printf("arr5 = %#v\n", arr5)

	fmt.Println(len(arr4))
	fmt.Println(cap(arr4))

	// 遍历
	for i, val := range arr5 {
		fmt.Printf("%v - %v\n", i, val)
	}

	for i := 0; i < len(arr5); i++ {
		fmt.Printf("%v\n", arr5[i])
	}
}


func arr2d()  {
	arr1 := [5][3]int{{1}, {2, 3}}
	arr2 := [...][3]int{{2}, {5, 7}}

	fmt.Printf("arr1 = %v\n", arr1)
	fmt.Printf("arr2 = %v\n", arr2)

	arr3 := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("arr3 address %p\n", &arr3)
	updateByValue(arr3)
	updateByPointer(&arr3)
	fmt.Println(arr3)
}


// 相当于参数复制了一份新的
func updateByValue(arr [5]int)  {
	fmt.Printf("copy arr value %p\n", &arr)
	arr[0] = 123

	// 不影响外部数组
	fmt.Println(arr)
}

// 传进来的是地址
func updateByPointer(arr *[5]int)  {
	fmt.Printf("copy pointer %p\n", arr)
	arr[0] = 777
	fmt.Printf("func arr %v\n", arr)
}