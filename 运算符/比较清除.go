package main

import "fmt"

func main() {
	//	比较清除的运算符为&^

	r1 := 5 //	5 --> 00000101
	r2 := 99 //	99 --> 01100011
	r3 := r1 &^ r2 // r3 --> 00000100 = 4

	//	该运算以第一个运算数为基准，如果同一位置上均为1，则将该位置0
	//	否则照搬第一个运算数
	fmt.Println(r3)

}
