package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()

	//创建一个死循环，一直读取客户端连接
	for {
		buf := make([]byte, 1024)

		fmt.Printf("服务器在等待客户端%s 发送信息\n", conn.RemoteAddr().String())
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("客户端退出")
			return
		}
		fmt.Println("client data:", string(buf[:n]))
	}
}

func main() {

	fmt.Println("服务端开始监听。。。")
	listen, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer listen.Close()

	for {
		fmt.Println("等待客户端连接。。。")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed ,err = ", err)
		} else {
			fmt.Printf("accept success,conn = %v,ip = %v\n", conn, conn.RemoteAddr().String())
		}

		//开启协程，接收连接
		go process(conn)
	}
}
