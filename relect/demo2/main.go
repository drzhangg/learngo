package main

import (
	"fmt"
	"reflect"
)

func main() {
	
	//声明一个空结构体
	type cat struct {
	}

	ins := &cat{}

	typeOfCat := reflect.TypeOf(ins)

	fmt.Printf("name:'%v' kind:'%v'\n",typeOfCat.Name(),typeOfCat.Kind())

	typeOfCat = typeOfCat.Elem()

	fmt.Printf("element name: '%v', element kind: '%v'\n", typeOfCat.Name(), typeOfCat.Kind())
}
