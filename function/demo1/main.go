package main

import "fmt"

type Person struct {
	name string
	sex byte
	age int
}

func (p Person) SetInfoValue(n string,s byte,a int) {

	p.name = n
	p.sex = s
	p.age = a
	fmt.Println("p = ",p)
	fmt.Printf("setValue &p:%p\n",&p)
}

func (p *Person) SetInfoPointer(n string, s byte, a int) {
	p.name = n
	p.sex = s
	p.age = a
	fmt.Printf("setPointer p:%p\n",p)
}

func main() {
	s1 := Person{"tom",'m',12}
	fmt.Printf("&s1 : %p\n",&s1)
	s1.SetInfoValue("bob",'s',13)
	s1.SetInfoPointer("jerry",'s',14)
}
