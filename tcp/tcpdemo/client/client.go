package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	conn, err := net.Dial("tcp", "0.0.0.0:8080")
	if err != nil {
		fmt.Println("dial failed, err= ", err)
		return
	}
	defer conn.Close()

	//客户端可以发送单行数据，然后就退出
	reader := bufio.NewReader(os.Stdin) //os.Stdin	代表标准输入[终端]

	for {
		//从终端读取一行用户输入，并准备发给服务器
		line, err := reader.ReadString('\n') //通过换行符判断一行数据
		if err != nil {
			fmt.Println("read string err = ", err)
		}

		line = strings.Trim(line, " \r\n")
		if line == "exit" {
			fmt.Println("客户端退出。。。")
			break
		}

		//将line发给服务器
		_, err = conn.Write([]byte(line + "\n"))
		if err != nil {
			fmt.Println("conn write err = ", err)
		}

	}
}
