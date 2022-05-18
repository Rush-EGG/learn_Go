package main

import "fmt"

type Person struct {
	name string
	age int
}

type School struct {
	name string
	city string
}

func something(arg interface{}){
	switch tp := arg.(type) {
	case Person:
		fmt.Println(tp.name)
	case School:
		fmt.Println(tp.name)
	default:
		fmt.Println(tp)
	}
}

func test1(){
	something("卢小喷")
	something(School{name: "besti", city: "北京"})
}
