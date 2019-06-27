package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	go func() {
		for i := 0; i <= 10; i++ {
			fmt.Println("hello,goroutine +" + strconv.Itoa(i))
			time.Sleep(time.Second)
		}
	}()

	for i := 0; i <= 10; i++ {
		fmt.Println("hello,main +" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
