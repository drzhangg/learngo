package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTimer(2 *time.Second)
	fmt.Println("当前时间：",time.Now())

	t1 := <- t.C
	fmt.Println("t1 = ",t1)


}
