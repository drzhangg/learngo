package main

import (
	"fmt"
	"net"
)

//处理和客户端的通讯
func process(conn net.Conn) {
	defer conn.Close()

	//循环接收客户端的通讯
	for {
		buf := make([]byte, 8096)
		fmt.Println("开始读取客户端发送的数据。。。")
		n,err := conn.Read(buf[:4])	//读取数据长度为4
		if n != 4 || err != nil {
			fmt.Println("conn read failed, err = ",err)
			return
		}
		fmt.Println("读到的buf = ",buf[:4])
	}
}

func main() {

	fmt.Println("服务器在8889端口监听。。。")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net listen failed,err = ", err)
		return
	}

	//一旦监听成功,就等待客户端来连接服务器
	for {
		fmt.Println("等待客户端来连接服务器。。。")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen accept failed,err = ", err)
		}

		go process(conn)
	}
}
