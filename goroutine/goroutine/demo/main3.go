package main

import "fmt"

func main() {
	var intChan chan int
	intChan = make(chan int,3)

	fmt.Printf("%v   %p\n",intChan,&intChan)

	intChan <- 23
	num := 13
	intChan <- num
	intChan <- 1
	//intChan <- 2

	<- intChan
	<- intChan
	<- intChan
}
