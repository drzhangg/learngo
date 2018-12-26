package main

import (
	"fmt"
	"time"
)

func running() {

	var times int

	//无限循环
	for {
		times++
		fmt.Println("tick",times)

		//延时1秒
		time.Sleep(time.Second)
	}
}

func main() {
	go running()

	//接收命令行输入，不做任何事情
	var input string
	fmt.Scanln(&input)
}
