package main

import "fmt"

func test1(){
	// 创建一个键是string类型，值也是string类型的map，但其中不存在任何键值对
	// userInfo := map[string]string{}
	userInfo := map[string]string{"name": "卢小喷", "age": "8"}

	fmt.Println(userInfo["name"])

	userInfo["age"] = "9"
	// 新增键值对
	userInfo["email"] = "woshizhu.com"

	fmt.Println(userInfo)
}
