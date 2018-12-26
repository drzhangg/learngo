package main

import (
	"fmt"
	"reflect"
)

func main() {

	//声明整型变量a并赋值
	var a int = 1024

	//获取变量a的反射值对象
	valueOfA := reflect.ValueOf(&a)

	valueOfA = valueOfA.Elem()

	valueOfA.SetInt(1)

	fmt.Println(a)
}
