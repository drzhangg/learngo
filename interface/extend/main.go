package main

import "fmt"

type Humaner interface {
	sayhi()
}

type Person interface {
	Humaner
	sing(str string)
}

type Student struct {
	name string
	age int
}

func (s *Student) sayhi() {
	fmt.Printf("Student[%s,%d]\n",s.name,s.age)
}

func (s *Student) sing(str string) {
	fmt.Println("Student唱歌：",str)
}

func main() {

	var i Person
	s :=&Student{"tom",12}
	i = s
	i.sayhi()
	i.sing("国歌")
}
