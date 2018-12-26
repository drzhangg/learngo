package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(time.Millisecond * 500)
	stopper := time.NewTimer(2 * time.Second)

	var i int

	for{
		select {
		case <-stopper.C:
			fmt.Println("stop")
		goto StopHere
		case <-ticker.C:
			i++
			fmt.Println("tick:",i)
		}
		 b := true && false
		 fmt.Println(b)
	}

	StopHere:
		fmt.Println("done")
}
