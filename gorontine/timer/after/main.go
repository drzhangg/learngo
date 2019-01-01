package main

import (
	"fmt"
	"time"
)

func main() {
	main1()
	main2()
	main3()
}

func main1() {
	timer := time.NewTimer(2 * time.Second)
	<- timer.C
	fmt.Println("main1时间到")
}

func main2() {
	time.Sleep(2 * time.Second)
	fmt.Println("main2时间到")
}

func main3() {
	time.After(2 * time.Second)
	fmt.Println("main3时间到")
}
