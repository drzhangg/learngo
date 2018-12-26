package main

import "fmt"

type BasicColor1 struct {
	R,G,B float32
}

type Color1 struct {
	BasicColor1
	Alpha float32
}

func main() {
	var c Color1

	c.R = 1
	c.G=1
	c.B=0

	c.Alpha=1
	fmt.Printf("%+v",c)
}
