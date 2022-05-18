package main

import (
	"fmt"
	"reflect"
)

type newPerson struct {
	name string "姓名"
	age int "年龄"
}

func test2(){
	p := newPerson{}
	pType := reflect.TypeOf(p)

	fieldNum := pType.NumField()
	for index := 0; index < fieldNum; index ++{
		field := pType.Field(index)
		fmt.Println(field.Name, field.Tag)
	}
}
