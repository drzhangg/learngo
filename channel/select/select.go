package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int{
	out := make(chan int)
	go func() {
		i:=0
		for{
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func worker(id int, c chan int) {
	for n := range c{
		time.Sleep(time.Second)
		fmt.Printf("working %d received %d\n",id,n)
	}
}

func createWorker(i int) chan <- int{
	c := make(chan int)
	go worker(i,c)
	return c
}

func main() {
	var c1,c2 = generator(),generator()
	var worker = createWorker(0)

	var values []int
	tm := time.After(10 * time.Second)
	//定时
	tick := time.Tick(time.Second)
	for{
		var activeWorker chan <- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}

		select {
		case n := <-c1:
			values = append(values,n)
		case n := <-c2:
			values = append(values,n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <-time.After(800 * time.Millisecond):
			fmt.Println("timeout")
		case <-tick:
			fmt.Println("queue len = ",len(values))
		case <- tm:
			fmt.Println("bye")
			return
		}
	}
}
