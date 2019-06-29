package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func wgo(i int) {
	fmt.Println(i)
	wg.Done()
}

func main() {

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go wgo(i)
	}

	wg.Wait()
	
}
