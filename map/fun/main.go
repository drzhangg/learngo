package main

import "fmt"

func test(m map[int]string, key int) {
	delete(m,key)
}

func main() {
	m := map[int]string{1:"bob",2:"tom",3:"jerry"}

	fmt.Println(m)

	test(m,2)

	fmt.Println(m)
}
