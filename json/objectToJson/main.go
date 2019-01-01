package main

import (
	"encoding/json"
	"fmt"
)

type Response1 struct {
	Page int
	Fruits []string
}

type Response2 struct {
	Page int	`json:"page"`
	Fruits []string		`json:"fruits"`
}

func main() {

	bolB,_ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB,_ := json.Marshal(123)
	fmt.Println(string(intB))

	fltb,_ := json.Marshal(3.14)
	fmt.Println(string(fltb))

	strB,_:= json.Marshal("hehehe")
	fmt.Println(string(strB))

	slcD := []string{"apple","orange","pear"}
	slcB,_ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple":5,"banana":7}
	mapB,_:= json.Marshal(mapD)
	fmt.Println(string(mapB))

	res1D := &Response1{
		Page:1,
		Fruits:[]string{"apple","pear","orange"},
	}
	res1B,_:= json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &Response2{
		Page:1,
		Fruits:[]string{
			"apple",
			"orange",
			"banana",
		},
	}
	res2B,_:= json.Marshal(res2D)
	fmt.Println(string(res2B))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var data map[string]interface{}

	if err := json.Unmarshal(byt, &data); err != nil {
		panic(err)
	}
	fmt.Println(data)

	num := data["num"].(float64)	//强转
	fmt.Println(num)

	strs := data["strs"].([]interface{})
	str1 := strs[1].(string)
	fmt.Println(str1)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := Response2{}
	json.Unmarshal([]byte(str),&res)
	fmt.Println(res)
	fmt.Println(res.Fruits[1])
}
