package main

import "fmt"

func main(){
	var name, address, action string

	fmt.Print("姓名：")
	fmt.Scanln(&name)

	fmt.Print("住址：")
	fmt.Scanln(&address)

	fmt.Print("行为：")
	fmt.Scanln(&action)

	result := fmt.Sprintf("我叫%s，我在%s做%s", name, address, action)
	fmt.Println(result)
}
