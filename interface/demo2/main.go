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

func WhoSayHi(i Humaner) {
	i.sayhi()
}



func main() {

	//var i Humaner

	s := &Student{"tom",12}
	t := &Teacher{"hangzhou","go"}
	m := Mystr("hello")

	WhoSayHi(s)
	WhoSayHi(t)
	WhoSayHi(&m)

	x:= make([]Humaner,3)
	x[0] = s
	x[1] = t
	x[2] = m

	for _,i := range x{
		i.sayhi()
	}

}
