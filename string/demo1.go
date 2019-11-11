package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {

	str := "430100000000"
	utf8.RuneCountInString(str)
	fmt.Println(str[5:6])
	sss := strings.Replace(str, str[5:6], "1", )
	fmt.Println(str)
	fmt.Println(sss)
}
