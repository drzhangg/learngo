package main

import (
	"fmt"
	"strconv"
)

func main() {

	str := "430100000000"
	ss,_ := strconv.Atoi(str)
	s := ss+999000
	fmt.Println(strconv.Itoa(s))

	fmt.Println(ss)
	fmt.Println(s)
	fmt.Println(ss + 999000)
}
