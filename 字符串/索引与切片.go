package main

import "fmt"

func main(){
	var name string = "卢小喷"

	v1 := name[0:3]
	fmt.Println(v1)  // 卢

	// 取每一个字节
	for i:=0; i < len(name); i++{
		fmt.Println(i, name[i])
	}

	// 取每一个字符
	for index, item := range name{
		fmt.Println(index, item, string(item))
	}

	var newName = "a卢k小4喷7"

	dataList := []rune(newName)
	for i:=0; i < len(dataList); i++{
		fmt.Println(dataList[i], string(dataList[i]))
	}
}
