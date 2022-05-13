package main

import "fmt"

func test3(){
	var name string
	var age int

	//不完善的占位符，Go全责
	fmt.Println("请输入：")
	fmt.Scanf("我叫%s，我今年%d岁", &name, &age)

	fmt.Println(name, age)
}
