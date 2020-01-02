package main

import "fmt"

func test(argList []int) {
	fmt.Println(argList)
}

func main() {
	var a = []int{1,2,3}
	test(a)
	test([]int{1})
	test([]int{1,2})
}
