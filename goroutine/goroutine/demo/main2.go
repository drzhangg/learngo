package main

import "fmt"

func main() {
	go func(){
		for i:=1;i<1000000;i++{
			fmt.Println(i)
		}
	}()

	fmt.Println("----main----")
}
