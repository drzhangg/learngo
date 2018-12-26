package main

import "fmt"

func main() {

	ch := make(chan int,10)
	ch <- 11
	ch <- 12

	close(ch)

	for x:= range ch{
		fmt.Println(x)
	}

	x,ok := <- ch
	fmt.Println(x,ok)
}
