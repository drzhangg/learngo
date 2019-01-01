package main

import (
	"fmt"
	"regexp"
)

func main() {

	buf := "abc adc a2c aac 111 q3c a1c rac"

	//reg1 := regexp.MustCompile(`a.c`)
	reg1 := regexp.MustCompile("a[0-9]c")
	if reg1 == nil {
		fmt.Println("regex err")
		return
	}

	//-1表示返回所有值
	result1 := reg1.FindAllStringSubmatch(buf,-1)
	fmt.Println(result1)
}
