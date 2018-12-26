package main

import "fmt"

func fire() {
	fmt.Println("fire in the ho")
}

func main() {

	var f func()

	f = fire

	f()
}
