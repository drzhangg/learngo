package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10000; i++ {
		go func(i int) {
			fmt.Printf("Hello from goroutine %d\n",i)
		}(i)
	}

	time.Sleep(time.Minute)
}

func test(str string,num int) (string, int, error) {
	return str,num,nil
}
