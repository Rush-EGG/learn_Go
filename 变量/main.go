package main

import "fmt"

func main(){
	name := "lzt"
	nickname := name
	fmt.Println(name, &name)
	fmt.Println(nickname, &nickname)

	name = "lxp"
	fmt.Println(name, &name)
	fmt.Println(nickname, &nickname)
}
