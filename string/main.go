package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "hello world"
	res0 := strings.HasPrefix(str,"http://")
	res1 := strings.HasPrefix(str,"hello")
	fmt.Printf("res0 is %v\n",res0)
	fmt.Printf("res1 is %v\n",res1)
	hasSuffix(str)
	count(str)
}

func hasSuffix(str string){
	res0 := strings.HasSuffix(str,"http://")
	res1 := strings.HasSuffix(str,"world")
	fmt.Printf("res0 is %v\n",res0)
	fmt.Printf("res1 is %v\n",res1)
}

func count(str string) {
	c1 := strings.Count(str,"l")
	c2 := strings.Count(str,"i")
	fmt.Printf("c1 is %v\n",c1)
	fmt.Printf("c2 is %v\n",c2)
}
