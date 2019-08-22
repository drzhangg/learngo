package main

import (
	"fmt"
	"strings"
)

func main() {
	//var a int = 90
	//
	//x := float64(a) * 1.1
	//s :=strconv.FormatFloat(x,'f',-1,64)
	//fmt.Println(s)


	str1 := "abc"
	//strings.Replace("d",str1,)
	str2 := strings.Split(str1,"c")

	str3 := str2[0] + "d"
	fmt.Println(str2)
	fmt.Println(str3)
}
