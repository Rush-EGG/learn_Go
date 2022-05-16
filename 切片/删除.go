package main

import "fmt"

func test4(){
	v1 := []int{11, 22, 33, 44, 55, 66, 77}

	deleteIndex := 3

	// 切片是没有原生的删除函数的，只能自己伪造
	result := append(v1[: deleteIndex], v1[deleteIndex+1:]...)

	fmt.Println(result)
	// 这样做的缺点就是会影响到v1中原有的数据
	fmt.Println(v1)
}
