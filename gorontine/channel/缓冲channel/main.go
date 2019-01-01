package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int,3)
	fmt.Printf("len(ch) = %d,cap(ch) = %d",len(ch),cap(ch))

	go func() {

		for i := 0; i < 3; i++ {
			ch <- i
			fmt.Printf("子协程[%d]:len(ch) = %d,cap(ch) = %d\n",i,len(ch),cap(ch))
			time.Sleep(time.Second)

		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 3; i++ {
		num := <- ch
		fmt.Println("num = ",num)
	}
}
