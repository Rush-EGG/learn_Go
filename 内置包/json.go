package main

import (
	"encoding/json"
	"fmt"
)

type Plant struct {
	//如果属性名首字母是小写的话将无法被json格式化
	//name string
	Name string `json:"name"`
	//subject string
	Subject string `json:"subject"`
}

func main() {
	i := []interface{}{
		"卢小喷",
		123,
		true,
		5.16,
		map[string]interface{}{
			"name": "besti",
			"city": "北京",
		},
		Plant{Name: "吊兰", Subject: "兰花科"},
	}

	//序列化
	res, _ := json.Marshal(i)
	fmt.Println(res)
	data := string(res)
	fmt.Println(data)

	//反序列化
	content := `["卢小喷", 123, true, 5.16,{"city":"北京", "name":"besti"},{"name":"吊兰", "subject":
	"兰花科"}]`
	var value []interface{}
	json.Unmarshal([]byte(content), &value)
	fmt.Println(value)
}
