package main

import "fmt"

func simple() func(a,b int) int {
	//fmt.Println(a(60,7))
	f := func(a,b int) int {
		return a + b
	}
	return f
}

func main() {
	s := simple()
	fmt.Println(s(61,8))
}
