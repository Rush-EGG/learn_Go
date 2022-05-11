package main

import (
	"fmt"
	"learn_Go/api"
)

func main() {
	fmt.Println("我是猪")
	Add()

	api.BaiDu()
	api.Google()
//	在文件中编写功能时，首字母尽量大写
//  如果写的是小写的话，默认这个功能只能在包内部使用
}
