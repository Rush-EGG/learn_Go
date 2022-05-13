package main

import "fmt"

func test2(){
	var name string
	var age int

	fmt.Print("请输入：")
	//Scanln和Scan的区别在于：如果要求输入两个值，那么用户利用空格分割共输入了三个值
	//Scanln不会对第三个值进行截断，且会报错，expected newline，第二个值也不能赋值
	//但是Scan会截断且不会报错

	//同时，Scanln等待换行符，不管有没有输入足够多的值赋给变量
	count, err := fmt.Scanln(&name, &age)

	fmt.Println(count, err)
	fmt.Println(name, age)
}
