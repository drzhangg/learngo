package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	myMap = make(map[int]int, 20)
	lock  sync.Mutex
)

func test(n int) {
	res := 1

	for i := 1; i <= n; i++ {
		res *= i
	}

	lock.Lock()
	myMap[n] = res
	lock.Unlock()
}

func main() {

	for i := 1; i < 20; i++ {
		go test(i)
	}

	time.Sleep(time.Second * 5)

	lock.Lock()
	for i, v := range myMap {
		fmt.Printf("map[%d] = %d\n", i, v)
	}
	lock.Unlock()
}
