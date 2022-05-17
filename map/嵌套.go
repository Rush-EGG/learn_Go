package main

import "fmt"

func main(){
	n := make(map[string][2]map[string]string)

	n["1"] = [2]map[string]string{map[string]string{"name": "pig"}, map[string]string{"age": "1"}}

	fmt.Println(n)
}
