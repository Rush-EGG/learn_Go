package main

import "fmt"

func change(arg *[2]int){
	arg[0] = 666
}

func test1(){
	dataList := [2]int{11, 22}

	change(&dataList)

	fmt.Println(dataList)
}
