package main

import "time"

func main() {

	intCh := make(chan int,1)

	close(intCh)

	<- intCh

	time.Sleep(time.Second * 5)

	//
	//str := "abcdefg"
	//
	//str1 := str[len(str) -2 :]
	//fmt.Println(str1)
}
