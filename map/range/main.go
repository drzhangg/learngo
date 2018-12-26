package main

import "fmt"

func main() {

	m := map[int]string{1:"tom",2:"jerry",3:"bob"}

	for k,v := range m{
		fmt.Printf("%d ====> %s\n",k,v)
	}

	value,ok := m[1]

	if ok {
		fmt.Println("m[1] = ",value)
	}else {
		fmt.Println("key不存在")
	}
	
}
