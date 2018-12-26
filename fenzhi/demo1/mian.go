package main

import "fmt"

func printType(v interface{}) {

	switch v.(type) {
	case int:
		fmt.Printf("%v is int\n",v)
	case string:
		fmt.Printf("%v is string\n",v)
	case bool:
		fmt.Printf("%v is bool\n",v)
	}
}

func main() {
	printType(1222)
	printType("bob")
	printType(true)
}
