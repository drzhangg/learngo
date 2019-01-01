package main

import "fmt"

func testa(x int) {
	fmt.Println("aaa")
}

func testb(x int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var a [10]int
	a[x] = 111
	fmt.Println(a[x])
}

func testc(x int) {
	fmt.Println("ccc")
}

func main() {

	testa(1)
	testb(1)
	testc(1)
}
