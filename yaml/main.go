package main

import (
	"gopkg.in/yaml.v2"
	"log"
)

func main() {
	yamlContent:= `
		field1:小韩说课
		field2:
			field3: value
			field4: [42,0124]
	`

	//存储解析数据
	result := make(map[string]interface{})
	err := yaml.Unmarshal([]byte(yamlContent),&result)
	if err != nil {
		log.Fatalf("error:%v",err)
	}
}
