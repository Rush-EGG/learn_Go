package main

import "fmt"

func test1(){
	var f1 float32
	f1 = 3.14
	f2 := 99.9
	f3 := float64(f1) + f2

	fmt.Println(f1, f2, f3)

	f4 := 0.1
	f5 := 0.2
	result := f4 + f5
	fmt.Println(result)
}
