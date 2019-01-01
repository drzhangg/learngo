package main

import (
	"fmt"
	"os"
)

func main() {

	os.Stdout.WriteString("asdas")

	var a int
	fmt.Println("a:")
	fmt.Scan(&a)
	fmt.Printf("a = ",a)
}
