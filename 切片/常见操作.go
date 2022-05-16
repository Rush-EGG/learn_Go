package main

import "fmt"

func test3() {
	v1 := []int{11, 22, 33, 44}

	v2 := v1[1:3]
	v3 := v1[:3]
	v4 := v1[2:]

	fmt.Println(v1, v2, v3, v4)

	v1[1] = 100
	fmt.Println(v1, v2, v3, v4)

	v5 := append(v1, 100, 200, 300)
	// 非常迷的...
	v6 := append(v1, []int{999, 888}...)

	fmt.Println(v5, v6)
}
