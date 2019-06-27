package main

import (
	"encoding/json"
	"fmt"
)

type Result struct {
	Status int
}


func main() {
	var data = []byte(`{"status":200}`)
	result := &Result{}
	if err := json.Unmarshal(data, result); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
