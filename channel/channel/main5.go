package main

import "fmt"

func main() {
	c := make(chan bool,100)
	for i := 0; i < 100; i++ {
		fmt.Println("外面的i地址：",&i)
		go func() {
			//fmt.Println("第一个",i)
			fmt.Println("1里面的i地址：",&i)
			c <- true
		}()

		go func(i int) {
			fmt.Println("2里面的i地址：",&i)
			//fmt.Println("第二个：",i)
			c <- true
		}(i)
	}

	for i := 0; i < 100; i++ {
		<- c
	}
}
