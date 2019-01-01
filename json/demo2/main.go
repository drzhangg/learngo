package main

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	Company string	`json:"company"`	//-表示该字段不会显示在控制台
	Subjects []string	`json:"subjects"`
	IsOK bool	`json:"isok"`
	Price float64	`json:"price"`
} 

func main() {
	buf := `{
 		"company": "itcast",
 		"subjects": [
 		 "go",
 		 "java"
 		],
 		"isok": true,
 		"price": 11.11
		}
		`

	var tmp IT
	err := json.Unmarshal([]byte(buf),&tmp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("tmp = %+v\n",tmp)

	type IT2 struct {
		Subjects []string `json:"subjects"`
	}

	var tmp2 IT2
	err = json.Unmarshal([]byte(buf),&tmp2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("tmp2 = %+v\n",tmp2)
}
