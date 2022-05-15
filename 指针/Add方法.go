package main

import (
	"fmt"
	"math/big"
)

func main(){
	p1 := new(big.Int)
	p2 := new(big.Int)

	p1.SetString("12454961649879", 10)
	p2.SetString("15", 8)

	var result1 big.Int
	result1.Add(p1, p2)

	fmt.Println(result1.String())

	var i1 big.Int
	var i2 big.Int

	i1.SetString("1235998746", 10)
	i2.SetString("1489563287", 10)

	var result2 big.Int
	// Add方法传入两个指针
	result2.Add(&i1, &i2)

	fmt.Println(result2.String())
}
