package main

import "fmt"

// 声明类型
type myInt int

//声明只有该类型能用的方法
func (i myInt) doSth(a1 int, a2 int) int{
	return a1 + a2 + int(i)
}

func main(){
	var i = myInt(5)

	fmt.Println(i.doSth(1, 2))

}
