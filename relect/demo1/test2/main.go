package main

import (
	"fmt"
	"reflect"
)

type Enum int

const (
	Zero Enum = 0
)

func main() {
	type cat struct {

	}

	typeOfCat := reflect.TypeOf(cat{})

	fmt.Println(typeOfCat.Name(),typeOfCat.Kind())

	typeOfA := reflect.TypeOf(Zero)

	fmt.Println(typeOfA.Name(),typeOfA.Kind())
}
