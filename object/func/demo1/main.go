package main

import "fmt"

func add1(a, b int) int {
	return a + b
}

type Long int

func (l Long) add2(other Long) Long {
	return l + other
}

func main() {

	var result int
	result = add1(1,2)
	fmt.Println("add1:",result)

	a := Long(3)
	r := a.add2(1)
	fmt.Println(r)
	
}
