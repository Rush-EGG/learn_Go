package main

import "fmt"

func changeData(data string){
	data = "哈哈哈"
}

func changePData(ptr *string){
	*ptr = "哈哈哈"
}

func test3(){
	name := "卢小喷"

	changeData(name)
	fmt.Println(name)

	changePData(&name)
	fmt.Println(name)
}
