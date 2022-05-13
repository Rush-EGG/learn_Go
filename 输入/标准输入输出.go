package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	reader := bufio.NewReader(os.Stdin)

	//line:从stdin中读取一行的数据（字节集合，可以转为字符串）
	//reader默认一次能读4096个字节
	//	如果能一次性读完，那么isPrefix = false
	//	如果读不完，那么只读一部分，且isPrefix = true
	//err是错误信息
	line, isPrefix, err := reader.ReadLine()
	data := string(line)
	fmt.Println(data)

	_ = isPrefix
	_ = err
}
