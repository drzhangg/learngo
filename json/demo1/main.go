package main

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	Company string	`json:"company"`	//-表示该字段不会显示在控制台
	Subjects []string	`json:"subjects"`
	IsOK bool	`json:",string"`
	Price float64	`json:",string"`
}

func main() {

	s := IT{"itcast",[]string{"go","java"},true,11.11}

	//buf,err := json.Marshal(s)
	buf, err := json.MarshalIndent(s,""," ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("buf = ",string(buf))
}
