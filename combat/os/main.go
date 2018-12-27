package main

import (
	"fmt"
	"os"
)

func main() {

	userFile := "main.go"
	fl,err := os.Open(userFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fl.Close()

	buf :=make([]byte,1024)
	for{
		n,_:= fl.Read(buf)
		if n == 0 {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}
