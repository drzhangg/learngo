package main

import "fmt"

func main() {

	var p *int		//p是*int类型，指向int类型

	p = new(int)

	*p = 123
	fmt.Println("*p = ",*p)
}
