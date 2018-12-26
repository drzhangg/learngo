package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	type ColorGroup struct {
		Id int
		Name string
		Group []string
	}

	cg := ColorGroup{
		Id:1,
		Name:"bob",
		Group:[]string{"red","yellow","blue"},
	}

	c,err := json.Marshal(cg)
	if err != nil {
		fmt.Println("error:",err)

	}

	fmt.Println(string(c))
	os.Stdout.Write(c)
	
}
