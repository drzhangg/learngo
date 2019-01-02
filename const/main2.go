package main

import "fmt"

type Color int

const(
	Black Color = iota
	Red
	Blue
)

func test(c Color) {
	fmt.Println(c)
}

func main() {

	c := Black
	test(c)
	test(2)
}
