package main

import "fmt"

func main() {
	a := func() {
		fmt.Println("hello world first")
	}
	a()
	fmt.Printf("%T",a)

	func(n string){
		fmt.Println("Welcome",n)
	}("Tom")
}
