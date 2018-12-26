package main

import (
	"gopkg.in/yaml.v2"
	"log"
)

func main() {
	// map数据
	data := map[string]interface{}{
		"field1": "小韩说课",
		"field2": map[string]interface{}{
			"field3": "value",
			"field4": []int{42, 1024},
		},
	}
	text, err := yaml.Marshal(&data)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
