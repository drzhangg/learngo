package main

import "fmt"

func fn() int {
	n := 0
	defer func() {
		n++
	}()
	return n
}

func main() {
	fmt.Println("fn()=",fn())
}
