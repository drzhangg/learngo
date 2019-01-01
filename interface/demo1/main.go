package main

import "fmt"

type Humaner interface {
	sayhi()
}

type Student struct {
	name string
	age int
}

func (s *Student) sayhi() {
	fmt.Printf("Studnet{%s,%d} sahyhi\n",s.name,s.age)
}

type Teacher struct {
	addr string
	group string
}

func (t *Teacher) sayhi() {
	fmt.Printf("Teacher{%s,%s} sayhi\n",t.addr,t.group)
}

type Mystr string

func (m Mystr) sayhi() {
	fmt.Printf("Mystr[%s] sayhi\n",m)
}



func main() {

	var i Humaner

	s := &Student{"tom",12}
	i = s
	i.sayhi()

	t := &Teacher{"hangzhou","go"}
	i = t
	i.sayhi()

	m := Mystr("hello")
	i = m
	i.sayhi()
}
