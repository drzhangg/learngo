package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for{
		buf := make([]byte,512)
		n,err := conn.Read(buf)
		if err != nil {
			fmt.Println("read err:",err)
			return
		}
		fmt.Println("read:",string(buf[:n]))
	}
}

func main() {
	fmt.Println("server start...")
	listen,err := net.Listen("tcp",":8000")
	if err != nil {
		fmt.Println("listen failed,err:",err)
		return
	}
	for{
		conn,err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed,err : ",err)
			continue
		}
		go process(conn)
	}
}
