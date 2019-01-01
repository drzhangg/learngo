package main

import "fmt"

type Person struct {
	name string
	age int
	sex byte
}

func (p Person) printInfo() {
	fmt.Println("tmp = ",p)
}

func (p *Person) SetInfo(n string,s byte,a int) {

	p.name = n
	p.sex = s
	p.age = a
}

func main() {

	p := Person{"tom",'m',21}
	p.printInfo()

	var a Person
	(&a).SetInfo("bob",'s',23)
	a.printInfo()

}
