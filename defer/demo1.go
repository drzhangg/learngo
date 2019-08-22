package main

import "fmt"

func main() {
	println(DeferFunc1(1))
	println(DeferFunc2(1))
	println(DeferFunc3(1))
}

func DeferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	fmt.Println("func1")
	return t
}

func DeferFunc2(i int)  int {
	t := i
	defer func() {
		t += 3
	}()
	fmt.Println("func2")
	return t
}

func DeferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	fmt.Println("func3")
	return 2
}