package main

import "fmt"

func writeChan(numChan chan int) {
	for i := 1; i <= 100; i++ {
		numChan <- i
	}
	close(numChan)
}

//读取数据
func resultChan(numChan chan int, resChan chan map[int]int, exitChan chan bool) {
	m := make(map[int]int)
	res := 0
	for {
		v, ok := <-numChan
		if !ok {
			break
		} else {
			//numChan每次进来是一个数，把进来的那个数存成key，和前面相加的和存成value
			//v是每次进来取的数
			res += v
			m = map[int]int{v: res}
			resChan <- m

		}
	}
	exitChan <- true

}

func main() {

	numChan := make(chan int, 100)
	exitChan := make(chan bool, 1)
	resChan := make(chan map[int]int, 100)

	go writeChan(numChan)

	go resultChan(numChan, resChan, exitChan)

	go func() {
		<-exitChan

	}()

	for {
		v, ok := <-resChan
		if !ok {
			break
		}
		fmt.Println(v)
	}

}
