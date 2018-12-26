package main

import (
	"fmt"
	"time"
)

func main() {
	intChan := make(chan int,10)

	for i := 0; i < 10; i++ {
		intChan<-i
	}
	close(intChan)
	for{
		i:= <- intChan
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}
