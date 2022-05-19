package main

import (
	"fmt"
	"time"
)

func test2(){
	currentDate := time.Now()

	//格式化处理：将Time类型时间转为字符串   %Y  %m %d %H %M %Ss
	dateString := currentDate.Format("2006_01_02_15_04_05")

	fmt.Println(dateString)
}
