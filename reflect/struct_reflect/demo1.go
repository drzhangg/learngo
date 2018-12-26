package main

import (
	"fmt"
	"reflect"
)

type Stu1 struct {

}

func main() {

	var v *Stu1
	typeV := reflect.TypeOf(v)
	fmt.Println(typeV)

	fmt.Println(typeV.Kind())

	fmt.Println(typeV.Elem())

	fmt.Println(typeV.Elem().Kind())
}
