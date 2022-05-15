package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

func main() {
	var v1 = decimal.NewFromFloat(0.00000019)
	//fmt.Println(v1)
	var v2 = decimal.NewFromFloat(0.29)

	var v3 = v1.Add(v2)

	fmt.Println(v1, v2, v3)

	var price = decimal.NewFromFloat(3.4685)
	var data1 = price.Round(1) // 四舍五入保留一位小数
	var data2 = price.Truncate(1) // 保留小数点后一位

	fmt.Println(price, data1, data2)
}
