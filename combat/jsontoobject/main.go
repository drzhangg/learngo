package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Person struct {
	Name string	`json:"name"`
	Worker string	`json:"worker"`
}
var person,p1 Person

func main() {
	jsonStr := []byte(`[{"Name":"Student","Worker":"Study"},{"Name":"worker","Worker":"make money"}]`)

	var persons []Person
	err := json.Unmarshal(jsonStr,&persons)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("persons:",persons)



	for s,val := range persons{
		person = val
		p1 = persons[s]
		fmt.Println("person:",person)
		fmt.Println("p1",p1)
	}

	pp := reflect.ValueOf(&p1)
	field := pp.Elem().FieldByName("Name")
	fmt.Println("field:",field)




}
