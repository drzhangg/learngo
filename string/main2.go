package main

import "fmt"

func main() {

	var f [2][3]int = [...][3]int{{1,2,3},{4,5,6}}

	for k1,v1 := range f{
		for k2,v2 := range v1{
			fmt.Printf("(%d,%d)=%d",k1,k2,v2)
		}
		fmt.Println()
	}
}
