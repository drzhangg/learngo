package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1,41,23,35,2,10}
	sort.Ints(a)

	for i,v := range a{
		fmt.Println(i,v)
	}
}
