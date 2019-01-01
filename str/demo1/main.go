package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(strings.Contains("hellogo","hello"))
	fmt.Println(strings.Contains("hellogo","abvc"))

	s:= []string{"abc","hello","mike","go"}
	buf := strings.Join(s,"@")
	fmt.Println("buf = ",buf)

	fmt.Println(strings.Index("adchellogo","go"))
	fmt.Println(strings.Index("adchellogo","hi"))

	buf = strings.Repeat("go",4)
	fmt.Println("buf = ",buf)

	//split
	buf = "hello@asd@hehe@tom"
	fmt.Println(strings.Split(buf,"@"))

	//trim
	buf = "    are     u   ok?"
	fmt.Println(strings.Trim(buf," "))

	s1 := strings.Fields(buf)
	for i,data := range s1{
		fmt.Println(i,",",data)
	}
}
