package main

import (
	"fmt"
	"net"
	"strings"
)

func HandlerConn(conn net.Conn) {
	defer conn.Close()

	//获取客户端的网络地址信息
	addr := conn.RemoteAddr().String()
	fmt.Println(addr," connect successful")

	buf := make([]byte,1024 * 2)

	for{
		//读取用户数据
		n,err := conn.Read(buf)
		if err != nil {
			fmt.Println("err = ",err)
			return
		}
		fmt.Printf("[%s]: %s\n",addr,string(buf[:n]))
		fmt.Println("len(buf):",len(string(buf[:n])))

		if "exit" == string(buf[:n-1]) {
			return
		}

		//把数据转换为大写，再给用户发送
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}

}

func main() {
	//监听
	listener ,err := net.Listen("tcp","127.0.0.1:8000")
	if err != nil {
		fmt.Println("err = ",err)
		return
	}

	defer listener.Close()

	//接收多个用户请求
	for{
		conn,err := listener.Accept()
		if err != nil {
			fmt.Println("err :",err)
			return
		}

		go HandlerConn(conn)
	}

}
