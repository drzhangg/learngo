package main

import (
	"fmt"
	"reflect"
)

type order struct {
	ordId int
	customerId int
}

func createQuery(q interface{}) {
	t := reflect.TypeOf(q)
	v := reflect.ValueOf(q)
	k := t.Kind()
	fmt.Println("Type ",t)
	fmt.Println("Value ",v)
	fmt.Println("kind ",k)
}

func main() {
	o := order{
		ordId:1234,
		customerId:567,
	}
	createQuery(o)
}
