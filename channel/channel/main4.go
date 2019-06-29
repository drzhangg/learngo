package main

import "fmt"

func getNum(numChan chan int) {
	for i := 0; i < 100; i++ {
		numChan <- i
	}
	close(numChan)
}

func readChan(numChan chan int, resChan chan map[int]int, exitChan chan bool) {
	for {
		num, ok := <-numChan
		if !ok {
			break
		}
		m := map[int]int{num: (1 + num) * num / 2}
		resChan <- m

	}
	exitChan <- true

}

func main() {

	numChan := make(chan int, 100)
	resChan := make(chan map[int]int, 100)
	exitChan := make(chan bool, 8)

	go getNum(numChan)

	for i := 0; i < 8; i++ {
		go readChan(numChan, resChan, exitChan)
	}

	go func() {
		for i := 0; i < 8; i++ {
			<-exitChan
		}
		close(resChan)
	}()

	for {
		v, ok := <-resChan
		if !ok {
			break
		}
		fmt.Println(v)
	}
}
