package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file,err := os.Open("./file/dir/demo3.go")
	if err != nil{
		fmt.Println(err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	for{
		//缓冲区
		buf := make([]byte,1024)
		n,err := reader.Read(buf)

		if n == 0 && err == io.EOF {
			break
		}
		fmt.Println(string(buf[:n]))
	}
}
