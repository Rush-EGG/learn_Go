package main

import "fmt"

type Person struct {
	name string "姓名"
	age int "年龄"
}

func test1(){
	p1 := Person{"卢小喷", 8}
	p2 := p1

	p1.name = "卢大喷"

	fmt.Println(p1)
	fmt.Println(p2)

	// 结构体指针
	pp1 := &Person{"lzt", 23}
	pp2 := pp1

	fmt.Println(pp1, pp2)

	// 直接修改 或者用(*pp1).*** = ***
	pp1.name = "pig"
	fmt.Println(pp1, pp2)
}
