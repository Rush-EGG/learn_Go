package main

import "fmt"

func test5() {
	v1 := []int{11, 22, 33, 44, 55, 66}

	insertIndex := 3 // 在索引3的位置加入99

	result := make([]int, 0, len(v1)+1)
	result = append(result, v1[:insertIndex]...)
	result = append(result, 99)
	result = append(result, v1[insertIndex:]...)

	fmt.Println(result)
}
