package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func test1(){
	var name string = "卢小喷"

	// 卢 = 229 141 162
	fmt.Println(name[0], strconv.FormatInt(int64(name[0]), 2))
	fmt.Println(name[1], strconv.FormatInt(int64(name[1]), 2))
	fmt.Println(name[2], strconv.FormatInt(int64(name[2]), 2))

	// 小 = 229 176 143
	fmt.Println(name[3], strconv.FormatInt(int64(name[3]), 2))
	fmt.Println(name[4], strconv.FormatInt(int64(name[4]), 2))
	fmt.Println(name[5], strconv.FormatInt(int64(name[5]), 2))

	// 把字符串转换为一个“字节集合”
	byteSet := []byte(name)
	fmt.Println(byteSet)

	// 把字节的集合转为字符串
	byteList := []byte{229, 141, 162, 229, 176, 143, 229, 150, 183}
	targetString := string(byteList)
	fmt.Println(targetString)

	// 将字符串转换为unicode的码点的集合
	tempSet := []rune(name)
	fmt.Println(tempSet) // [21346 23567 21943] // 三个十进制的数
	fmt.Println(strconv.FormatInt(int64(tempSet[0]), 16))

	// 把rune集合转为字符串
	runeList := []rune{21346, 23567, 21943}
	targetName := string(runeList)
	fmt.Println(targetName)

	// 字符长度
	runeLength := utf8.RuneCountInString(name)
	fmt.Printf("字符长度:%d\n", runeLength)
	// 字节长度
	fmt.Printf("字节长度:%d",len(name))
}
