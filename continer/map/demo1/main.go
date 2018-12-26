package main

import "fmt"

func main() {

	m := map[string]string{
		"W":"forward",
		"A":"left",
		"S":"backward",
		"D":"right",
	}
	fmt.Println(m)

	for k,v := range m{
		fmt.Println(k,v)
	}
}
