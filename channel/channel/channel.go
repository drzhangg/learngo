package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	for{
		n,ok := <- c
		if !ok{
			break
		}
		fmt.Printf("working %d received %c\n",id, n)
	}
}

func createWorker(i int) chan <- int{
	c := make(chan int)
	go worker(i,c)

	return c
}

func chanDemo() {
	 var channels [10]chan<- int

	for i := 0; i < 10; i++  {
		channels[i] =  createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i]  <- 'a' + i
	}

	 //go worker(0,c)
	 //
	 //c <- 1
	 //c <- 2

	 time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	c := make(chan int,3)
	go worker(0,c)

	c <- 'a'
	c <- 'd'
	c <- 'c'
	c <- 'b'
	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int,3)
	go worker(0,c)
	c <- 'a'
	c <- 'd'
	c <- 'c'
	c <- 'b'
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	//chanDemo()
	//bufferedChannel()
	channelClose()
}
