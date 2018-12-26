package main

import "fmt"

//此通道只能写，不能读
func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i*i
	}
	close(out)
}

//此通道只能读，不能写
func consumer(in <-chan int) {
	for num := range in{
		fmt.Println("num = ",num)
	}
}

func main() {

	ch := make(chan int)

	go producer(ch)
	consumer(ch)
}
