package main

import "fmt"

type Person struct {
	name string
	age int
	sex byte
}

type Student struct {
	Person
	id int
	addr string
}


func main() {

	var s1 Student = Student{Person{"mike",19,'m'},1,"hangzhou"}
	fmt.Println("s1 : ",s1)

	s2 := Student{Person{"tom",20,'m'},2,"zhengzhou"}
	fmt.Printf("s2 : %v\n",s2)
	fmt.Printf("s2 : %+v\n",s2)
}
