package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonBuf := `{
 		"company": "itcast",
 		"subjects": [
 		 "go",
 		 "java"
 		],
 		"isok": true,
 		"price": 11.11
		}
		`

	m := make(map[string]interface{},4)

	err := json.Unmarshal([]byte(jsonBuf),&m)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("m = %+v\n",m)

	//var str string
	for key,value := range m{
		//fmt.Printf("%s ==> %v\n",key,value)
		switch data := value.(type) {
		case string:
			fmt.Printf("map[%s]的类型为string，值为%v\n",key,data)
		case bool:
			fmt.Printf("map[%s]的类型为bool，值为%v\n",key,data)
		case float64:
			fmt.Printf("map[%s]的类型为float64,值为%v\n",key,data)
		case []interface{}:
			fmt.Printf("map[%s]的类型为[]interface,值为%v\n",key,data)
		}
	}
}
