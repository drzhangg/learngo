package main

import "fmt"

func main() {

	intChan := make(chan int, 100)

	for i := 0; i < 100; i++ {
		intChan <- i * 2
	}

	close(intChan)

	for v := range intChan{
		fmt.Println(v)
	}
}
