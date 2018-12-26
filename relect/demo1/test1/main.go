package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int

	typeOfA := reflect.TypeOf(a)
	fmt.Println(typeOfA.Name(),typeOfA.Kind())
}
