package main

import (
	"fmt"
	"os"
)

func main() {
	str := `magnet:?xt=urn:sha1:YNCKHTQCWBTRNJIV4WNAE52SJUQCZO5C`

	file, err := os.Create("./cili.txt")
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < 20000; i++ {
		n, err := file.WriteString(str)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(n)
	}

}
