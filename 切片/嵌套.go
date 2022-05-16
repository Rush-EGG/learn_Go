package main

import "fmt"

func main() {
	v1 := []int{11, 22, 33}

	v2 := [][]int{[]int{11}, []int{22, 33}}

	v3 := [][2]int{[2]int{11, 22}, [2]int{33, 44}}

	fmt.Println(v1, v2, v3)
}
