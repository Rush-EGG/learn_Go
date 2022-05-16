package main

import "fmt"

func test1(){
	var nums []int

	var data = []int{11, 22, 33}
	// 或是简写为 data := []int{11, 22, 33}

	// make只用于 切片、字典、channel
	// make函数的第二、三个参数分别表示已指定的长度和总共预留的长度
	var users = make([]int, 2, 5)

	// 返回切片指针
	var pSlice = new([]int)

	// 返回指向空的切片的指针
	var pSliceNil *[]int

	fmt.Println(nums, data, users)
	fmt.Printf("%p %d\n", pSlice, *pSlice)
	fmt.Printf("%p\n", pSliceNil)
	fmt.Println(pSliceNil)
}
