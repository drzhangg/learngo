package main

import (
	"fmt"
)

//ch只写，quit只读
func fibonacci(ch chan<- int, quit <-chan bool) {
	x,y := 1,1

	for{
		select {
		case ch <- x:

			x,y = y,x+y
		case flag := <- quit:
			fmt.Println("flag = ",flag)
			return
		}
	}
}

func main() {
	ch := make(chan int)
	quit := make(chan bool)

	go func() {
		for i := 0; i < 8; i++ {
			num := <- ch
			fmt.Println(num)
		}
		quit <- true
	}()

	fibonacci(ch,quit)
}
