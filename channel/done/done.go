package main

import (
	"fmt"
	"sync"
)

func doWorker(id int, w worker) {
	for n := range w.in {
		fmt.Printf("Worker %d received %c\n", id, n)
		w.done()
	}
}

type worker struct {
	in   chan int
	done func()
}

func createWorker(id int,wg *sync.WaitGroup) worker {
	w := worker{
		in:make(chan int),
		done:func() {
			wg.Done()
		},
	}
	go doWorker(id, w)
	return w
}

func chanDemo() {
	var wg sync.WaitGroup

	var channels [10]worker
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i,&wg)
	}

	wg.Add(20)
	for i,channel := range channels{
		channel.in <- 'a' + i
	}
	for i,channel := range channels{
		channel.in <- 'a' + i
	}
	wg.Wait()
}

func main() {
	chanDemo()
}
