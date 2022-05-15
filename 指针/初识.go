package main

import "fmt"

func test1(){
	var p1 *int // 指向nil
	p2 := new(int)

	fmt.Printf("in point, p1 = %p\n", p1)
	fmt.Printf("in numb, p1 = %d\n", p1)
	fmt.Printf("in point, p2 = %p", p2)
}
