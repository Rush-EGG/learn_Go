package main

import "fmt"

type IMessage interface {
	send() bool
}

type Email struct {
	email string
	content string
}

func (p *Email)send() bool{
	fmt.Println("发送邮件", p.email, p.content)
	return true
}

type Wechat struct {
	vid string
	content string
}

func (p *Wechat)send() bool{
	fmt.Println("发送微信", p.vid, p.content)
	return true
}

func doSth(objectList []IMessage){
	for _, item := range objectList{
		_ = item.send()
		//fmt.Println(result)
	}
}

func main(){
	messageObjectList := []IMessage{
		&Email{"59@163.com", "注册啦"},
		&Wechat{"lzt", "注册啦"},
	}
	doSth(messageObjectList)
}
