package main

import (
	"encoding/json"
	"fmt"
)

type Test struct {
	Name string
	Age int
}

var tests []Test

func main() {

	tests = append(tests,Test{"a",1})
	tests = append(tests,Test{"b",2})

	byte,_ := json.Marshal(tests)
	fmt.Println(string(byte))
}
