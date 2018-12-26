package main

import (
	json2 "encoding/json"
	"fmt"
	"reflect"
)

type Stu2 struct {
	Name string `json:"j_name" bson:"b_name"`
	Sn string `json:"j_sn" bson:"b_sn"`
}

func main() {
	var v Stu2
	typeV := reflect.TypeOf(v)
	for i, c := 0, typeV.NumField(); i < c; i++ {
		fmt.Println(typeV.Field(i).Tag.Get("json"),typeV.Field(i).Tag.Get("bson"))
	}

	 v1 := Stu2{
	 	"Tom",
	 	"kang-111",
	 }
	 json,err := json2.Marshal(v1)
	 fmt.Println(string(json),err)
}
