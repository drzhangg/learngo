package main

import "fmt"

func main() {
	c := make(chan int,10)

	for i := 0; i < 10; i++ {
		c <- i
	}
	//close(c)
	for v := range c{
		fmt.Println(v)
		if len(c) == 0 {
			break
		}
	}
}
