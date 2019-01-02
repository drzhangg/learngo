package main

import "fmt"

const (
	A,B = iota,iota << 10
	C,D
)

func main() {
	fmt.Println(A,B,C,D)
}
