package main

import "fmt"

func main(){
	name1 := [2]string{"卢小喷", "卢大喷"}
	name2 := name1

	fmt.Printf("%p\n", &name1)
	fmt.Printf("%p\n", &name2)

	name1[0] = "lzt"
	fmt.Println(name1, name2)
}
