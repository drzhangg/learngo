package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	b,err := ioutil.ReadFile("main.go")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)
	fmt.Println(string(b))
}
