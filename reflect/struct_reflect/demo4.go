package main

import (
	"fmt"
	"reflect"
)

func fn(p1, p2 int) int {
	return p1+p2
}

func main() {
	valueFunc := reflect.ValueOf(fn)
	paramList := []reflect.Value{
		reflect.ValueOf(22),
		reflect.ValueOf(20),
	}
	resultList := valueFunc.Call(paramList)
	fmt.Println(resultList[0].Int())
}
