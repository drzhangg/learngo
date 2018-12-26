package main

import "fmt"

func iMap(s []int, f func(int) int) []int {
	var r []int
	for _,v := range s{
		r = append(r,f(v))
	}
	return r
}

func main() {
	a := []int{5,6,4,7,8}
	r := iMap(a, func(n int) int {
		return n * 5
	})
	fmt.Println(r)
}
