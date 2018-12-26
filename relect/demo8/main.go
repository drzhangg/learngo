package main

import (
	"fmt"
	"reflect"
)

func main() {

	var a int

	//取变量a的反射类型对象
	typeOfA := reflect.TypeOf(a)

	//根据反射类型对象创建类型实例
	aIns := reflect.New(typeOfA)

	//输出value的类型和种类
	fmt.Println(aIns.Type(),aIns.Kind())
}
