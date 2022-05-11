package main

import "fmt"

func main(){
	//const (
	//	v1 = 0
	//	v2 = 1
	//	v3 = 2
	//  v4 = 3
	//)
	const (
		// iota从0开始计数
		v1 = iota + 2 // 2
		_ // 3
		v2 // 4
		v3 // 5
		v4 // 6
	)

	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)
	fmt.Println(v4)
}
