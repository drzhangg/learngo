package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "123 2323"
	//strings.Split(str)
	str1 := strings.Trim(str," ")
	fmt.Println(str1)
}
