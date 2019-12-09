package main

import (
	"fmt"
	"strings"
)

func main() {
	s := `30m13s`
	ss := strings.Split(s,"h")
	fmt.Println(ss)
}
