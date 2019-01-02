package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	var num int64 = 999

	str := strconv.FormatInt(num,10)
	fmt.Println(str)
	fmt.Println(reflect.TypeOf(str))

}
