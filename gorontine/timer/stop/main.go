package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(3 * time.Second)

main1()
	go func() {
		<- timer.C
		fmt.Println("子协程可以打印了")
		timer.Stop()
	}()
	//timer.Stop()

	for{
		//timer.Stop()
	}
}

func main1() {
	timer := time.NewTimer(3 * time.Second)
	timer.Reset(time.Second)
	fmt.Print("猜猜")
}
