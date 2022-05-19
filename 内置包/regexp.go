package main

import (
	"fmt"
	"regexp"
)

func main(){
	value := "seafood"
	flag1, _ := regexp.MatchString("foo.*", value)

	fmt.Println(flag1)

	//根据字符串查找
	reg1 := regexp.MustCompile(`\d{11}`)
	ret1 := reg1.FindString("15118975623adf11147859652")
	ret2 := reg1.FindAllString("15114789523pp25611478911", -1)
	fmt.Println(ret1, ret2)
}
