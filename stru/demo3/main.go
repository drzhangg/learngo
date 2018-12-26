package main

import "fmt"

type Student struct {
	id int
	name string
	sex byte
	age int
	addr string
}

func test2(s *Student) {
	s.id = 12
	fmt.Println("test2:",s)
}

func test01(s Student) {
	s.id = 22
	fmt.Println("test01:",s)
}

func main() {
	s := Student{1,"tom",'m',22,"hangzhou"}

	test01(s)
	fmt.Println("main before test2:",s)
	test2(&s)
	fmt.Println("main after test2:",s)
}
