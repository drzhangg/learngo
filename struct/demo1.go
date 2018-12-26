package main

import "fmt"

//基础颜色
type BasicColor struct {
	R,G,B float32
}

//完整颜色定义
type Color struct {
	Basic BasicColor

	Alpha float32
}

func main()  {

	var c Color

	c.Basic.R = 1
	c.Basic.G = 1
	c.Basic.B = 0

	c.Alpha = 1

	fmt.Printf("%+v",c)
}
