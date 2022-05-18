package main

import "fmt"

func main() {
	var funcList []func()

	for i := 0; i < 5; i++ {
		funcion := func(arg int) func() {
			return func() {
				fmt.Println(arg)
			}
		}(i)
		funcList = append(funcList, funcion)
	}

	funcList[0]()
	funcList[1]()
	funcList[2]()
}
