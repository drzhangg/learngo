package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan string)

	defer fmt.Println("主协程调用完毕")

	go func() {

		defer fmt.Println("子协程调用完毕")

		for i := 0; i < 2; i++ {
			fmt.Println("子协程 i = ",i)
			time.Sleep(time.Second)
		}

		ch <- "我是子协程，完成了工作"
	}()

	str := <- ch		//没有传值，阻塞造成死锁
	fmt.Println("str = ",str)
}
