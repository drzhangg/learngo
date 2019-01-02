package main

import (
	"fmt"
)

func main() {
	s := "abc汉字！"

	for i := 0; i < len(s); i++ {
		fmt.Printf("%c",s[i])
	}

	fmt.Println()

	for _,r := range s{
		fmt.Printf("%c",r)
	}
}
