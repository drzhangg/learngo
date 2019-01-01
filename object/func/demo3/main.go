package main

import "fmt"

type Person struct {
	name string
	sex byte
	age int
}

func (p Person) SetInfoValue() {
	fmt.Printf("SetInfoValue:%p ,%v\n",&p,p)
}

func (p *Person) SetInfoPointer() {
	fmt.Printf("SetInfoPinter: %p, %v\n",p ,p)
}

func main() {

	p := Person{"tom",'m',12}
	fmt.Printf("main:%p,%v\n",&p,p)

	p.SetInfoValue()

	p.SetInfoPointer()

	f := (*Person).SetInfoPointer
	f(&p)

	f2 := (Person).SetInfoValue
	f2(p)
	
}
