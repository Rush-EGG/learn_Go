package main

import (
	"fmt"
)

func test1() {
	age := 66
	var temp int
	value := 3
	for {
		fmt.Println("请猜年龄：")
		fmt.Scanln(&temp)
		if temp != age{
			value --
			if value == 0{
				fmt.Println("你无了！")
				break
			}
			message := fmt.Sprintf("错了，还剩%d次机会", value)
			fmt.Println(message)
		}else{
			fmt.Println("恭喜，猜对了！")
			break
		}
	}
}
