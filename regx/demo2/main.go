package main

import (
	"fmt"
	"regexp"
)

func main() {

	buf := "23.13 1w.31 2.12 12312 0.12 asad 1.1111"

	reg := regexp.MustCompile(`\d+\.\d+`)
	if reg == nil {
		fmt.Println("mustCompile error")
		return
	}

	result := reg.FindAllString(buf,-1)
	fmt.Println(result)
}
