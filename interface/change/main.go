package main

import "fmt"

type Humaner interface {
	sayhi()
}
type Person interface {
	Humaner
	sing(lrc string)
}

type Student struct {
	name string
	age int
}

func (tmp *Student) sayhi() {
	fmt.Printf("Student[%s,%d] sayhi:",tmp.name,tmp.age)
}

func (tmp *Student) sing(lrc string) {
	fmt.Printf("Student sing:",lrc)
}



func main() {
	
}
