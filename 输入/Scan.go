package main

import "fmt"

func test1() {
	var name string
	var age int

	fmt.Println("请输入：")
	//当使用Scan时，会返回两个值：count是输入了几个值，err是错误信息
	count, err := fmt.Scan(&name, &age)
	fmt.Println(count, err)
	if err == nil {
		fmt.Println(name, age)
	}else{
		fmt.Println(err)
	}
}
