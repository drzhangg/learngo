package main

import "fmt"

func main() {

	var userChan chan interface{}	//chan里面放interface类型
	userChan = make(chan interface{},10)

	var readOnlyChan <- chan int		//只读chan
	var writeOnlyChan chan <- int

	//chan放取数据
	userChan <- "nick"
	name := <- userChan
	name,ok := <- userChan

	//关闭chan
	intChan1 := make(chan int,1)
	intChan1 <- 9
	close(intChan1)

	//range chan
	intChan := make(chan int,20)
	for i:=0;i<20;i++{
		intChan <- i
	}
	close(intChan)

	for v:= range intChan{
		fmt.Println(v)
	}

	useChan := make(chan interface{})
	go func() {
		useChan <- "nick"
	}()
	name := <- useChan
}
