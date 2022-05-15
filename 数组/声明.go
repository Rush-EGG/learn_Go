package main

import "fmt"

func test1() {
	// 先声明再赋值
	var numbers [3]int
	numbers[0] = 999
	numbers[1] = 888
	numbers[2] = 777

	// 声明+赋值
	var names = [2]string{"卢小喷", "卢大喷"}
	fmt.Println(names)

	// 声明+赋值+指定位置
	var ages = [3]int{0: 13, 1: 23, 2: 99}
	fmt.Println(ages)

	// 省略个数
	var newNames = [...]string{"1", "2", "3"}
	var newAges = [...]int{1, 3, 5, 7, 9, 11}
	fmt.Println(newNames, newAges)

	// 指针类型的数组，但不会开辟内存，即pnumbers = nil
	var pnumbers *[3]int
	fmt.Println(pnumbers)

	//声明数组并初始化，返回的是指针类型的数组
	newNumbers := new([3]int)
	fmt.Println(newNumbers)
	fmt.Println(*newNumbers)

}
