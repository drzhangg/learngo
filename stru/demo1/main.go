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

	var s1 Student = Student{1,"tom",'m',21,"hangzhou "}
	fmt.Println(s1)

	s2 := Student{name:"tom"}
	fmt.Println(s2)
}
