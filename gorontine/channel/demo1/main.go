package main

import (
	"fmt"
	"time"
)

var ch = make(chan int)

func Printer(str string){
	for _,data := range str{
		fmt.Printf("%c",data)
		time.Sleep(time.Second)
	}
	fmt.Printf("\n")
}

func person1() {
	Printer("hello")
	ch <- 444 		//给管道写数据，发送
}

func person2() {
	<- ch		//从管道取数据，接收，如果通道没有数据他就会阻塞
	Printer("goo")
}


func main() {
	//Printer("hello")
	go person1()
	go person2()

	for {

	}
}
