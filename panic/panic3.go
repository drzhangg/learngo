package main

import (
	"fmt"
	"runtime/debug"
)

func r() {
	if r := recover(); r != nil {
		fmt.Println("Recovered",r)
		debug.PrintStack()
	}
}

func a1()  {
	defer r()
	n := []int{5,7,4}
	fmt.Println(n[3])
	fmt.Println("normally returned from a")
}

func main() {
	a1()
	fmt.Println("normally returned from main")
}
