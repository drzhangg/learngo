package main

import "fmt"

func main() {
	intChan := make(chan int,10)
	for i := 0; i < 10; i++ {
		intChan<-i
	}
	close(intChan)

	for{
		//for i := range intChan{
		//	fmt.Println(i)
		//}
		i,ok := <- intChan
		if !ok{
			fmt.Println("channel is close")
			return
		}
		fmt.Println(i)
	}
}
