package main

import "fmt"

func test2(){
	v1 := make([]int, 1, 4)

	v2 := append(v1, 99)

	fmt.Printf("v1:=%d, v2:=%d\n", v1, v2)
	fmt.Printf("v1的地址:%p，v2的地址:%p\n", &v1, &v2)

	// 向原切片中添加数据
	v3 := make([]int, 1, 3)
	v3 = append(v3, 2)
	fmt.Println(v3)

	v4 := []int{11, 22}
	fmt.Printf("v4的长度是%d，v4的容量是%d\n", len(v4), cap(v4))
	v5 := append(v4, 33)
	fmt.Printf("v5的长度是%d，v5的容量是%d\n", len(v5), cap(v5))
	v4[0] = 99
	fmt.Printf("修改后，v4变为%d\n", v4)
	fmt.Printf("但v5仍然是%d\n", v5)

	v6 := []int{11, 22, 33}
	fmt.Printf("添加前：v6的长度是%d，v6的容量是%d\n", len(v6), cap(v6))
	v6 = append(v6, 44)
	// 因为是自己往自己添加，所以扩容时是(3+1)*2
	fmt.Printf("添加后：v6的长度是%d，v6的容量是%d\n", len(v6), cap(v6))
}
