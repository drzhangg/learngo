package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	file,err := os.Open("./demo2.go")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	for{
		str,err := reader.ReadString('\n')
		if err == io.EOF{
			return
		}
		fmt.Print(str)
	}
}
