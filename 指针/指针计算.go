package main

import (
	"fmt"
	"unsafe"
)

func main(){
	dataList := [3]int8{11, 22, 33}

	// 指向第一个元素的指针
	firstDataPtr := &dataList[0]
	//var firstDataPtr *int8 = &dataList[0]

	// 转成Pointer类型
	ptr := unsafe.Pointer(firstDataPtr)

	// 转成uintptr类型才能进行计算
	targetAddress := uintptr(ptr) + 1

	// 转成Pointer类型
	newPtr := unsafe.Pointer(targetAddress)

	// 得到指向下一个位置的元素的指针
	value := (*int8)(newPtr)

	fmt.Println("最终结果为：", *value)
}
