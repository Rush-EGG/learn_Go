package main

import "fmt"

func test2() {
	info := make(map[string]string, 10)

	fmt.Println(info, len(info))

	info["n1"] = "1"
	info["n2"] = "2"

	fmt.Println(info, len(info))

}
