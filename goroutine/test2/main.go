package main

import (
	"fmt"
	"time"
)

func main() {

	exit := make(chan int)

	fmt.Println("start")

	time.AfterFunc(3 * time.Second, func() {

		fmt.Println("one second after")
		exit <- 2
	})

	<-exit
	fmt.Println(&exit)
}
