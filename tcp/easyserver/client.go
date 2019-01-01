package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn,err := net.Dial("tcp","127.0.0.1:8000")
	if err != nil {
		fmt.Println("err = ",err)
		return
	}

	//接收服务器回复的任务，新数据
	go func() {
		//切片缓存
		buf := make([]byte,2048)

		for{
			n,err := conn.Write(buf)		//接收服务器的请求
			if err != nil {
				fmt.Println("err1 = ",err)
				return
			}
			fmt.Println(string(buf[:n]))	//打印接收到的内容，转换为字符串再打印
		}
	}()

	//从键盘输入内容，给服务器发送内容
	str := make([]byte,2048)
	for{
		n,err := os.Stdin.Read(str)
		if err != nil {
			fmt.Println("err2 = ",err)
			return
		}

		//把输入的内容给服务器发送
		conn.Write(str[:n])
	}
}
