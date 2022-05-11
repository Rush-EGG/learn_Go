package main

import "fmt"

// 全局变量，只能用这种形式命名
var school string = "DKY"
// 或是var school = "DKY"
// 但是不认school := "DKY"

func main() {
	//	内置函数
	print("好吃")
	fmt.Println("好吃")
	println("不好吃")

	// 变量
	var a int = 6
	var s string = "我是猪"
	var b bool = true
	fmt.Println(a, s, b)

	var name string

	fmt.Scanf("%s", &name)
	fmt.Println(name)

	if name == "🐖" {
		fmt.Println("输入正确")
	} else {
		fmt.Println("输入错误")
	}

	var (
		name1  = "卢小🐖"
		a1     = 18
		salary = "1w+"
	)
	fmt.Println(name1, a1, salary)
}
