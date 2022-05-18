package main

import "fmt"

type School struct {
	name string "名称"
	city string "城市"
}

func main() {
	var allSchool = []School{}
	var name, city, flag string
	for {
		fmt.Println("请输入学校名称:")
		fmt.Scanln(&name)
		fmt.Println("请输入学校所在城市:")
		fmt.Scanln(&city)

		allSchool = append(allSchool, School{name: name, city: city})

		fmt.Println("按n退出")
		fmt.Scanln(&flag)
		if flag == "n" {
			break
		}
	}

	allSchool = append(allSchool, School{name: "猪", city: "BJ"})

	fmt.Println(allSchool)
}
