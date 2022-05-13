package main

import "fmt"

func main() {
f1:
	for i := 1; i < 5; i++ {
		for j := 0; j < 2; j++ {
			if i == 2 && j == 1 {
				continue f1
			}
			fmt.Print("我是猪 ")
		}
		fmt.Println(i)
	}
}
