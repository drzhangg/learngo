package main

import (
	"fmt"
	"reflect"
)

func main() {

	type dog struct {
		LegCount int
	}
	//获取dog实例地址的反射值对象
	valueOfDog := reflect.ValueOf(&dog{})

	valueOfDog = valueOfDog.Elem()

	//获取legCount字段的值
	vLegCount := valueOfDog.FieldByName("LegCount")

	vLegCount.SetInt(2)
	fmt.Println(vLegCount.Int())
}
