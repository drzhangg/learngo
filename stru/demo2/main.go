package main

import "fmt"

type Student struct {
	id int
	name string
	sex byte
	age int
	addr string
}

func main() {
	var p *Student
	var s Student
	p = &s
	p.name="tom"

	fmt.Println(*s1)
}
