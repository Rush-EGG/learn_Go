package main

import (
	"fmt"
	"strings"
)

func test2(){
	var builder strings.Builder
	builder.WriteString("hhh,")
	builder.WriteString("吃大便")
	value := builder.String()

	fmt.Println(value)
}
