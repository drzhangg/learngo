package main

import (
	"fmt"
	"reflect"
)

type Stu struct {
	Name string
	Sn string
}

func (this *Stu) SetName() {
}

func (this *Stu) SetSn() {
}

func main() {

	v := Stu{}
	typeV := reflect.TypeOf(v)
	fmt.Println(typeV.NumField())

	for i, c := 0, typeV.NumField(); i < c; i++ {
		fmt.Println(typeV.Field(i))
	}

	vp := &v
	typeVP := reflect.TypeOf(vp)
	fmt.Println(typeVP.NumMethod())

	for i,c := 0,typeVP.NumMethod();i<c;i++{
		fmt.Println(typeVP.Method(i))
	}
}
