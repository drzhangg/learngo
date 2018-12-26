package main

import "fmt"

func printer(c chan int)  {

	for  {

		//从channel中获取一个数据
		data := <- c

		//将0视为数据结束
		if data == 0 {
			break
		}

		fmt.Println(data)
	}

	c <- 0
}

func main() {

	c:=make(chan int)

	go printer(c)

	for i := 1; i < 10; i++ {
		c <- i
	}

	c <-0
	<-c

}
