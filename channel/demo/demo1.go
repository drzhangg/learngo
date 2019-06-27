package main

import "fmt"


var ch chan int = make(chan int)

func loop() {
	for i :=0;i<10;i++{
		fmt.Printf("%d\n",i)
	}
	ch <- 1
	fmt.Println(len(ch))
}

func main() {
	go loop()
	<- ch
	ch2 := make(chan int)
	ch2 <- 1
	<- ch2

	fmt.Println("is hi")

}
