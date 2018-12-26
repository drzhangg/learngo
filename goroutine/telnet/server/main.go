package server

import (
	"fmt"
	"net"
)

//服务逻辑，传入地址和退出的通道
func server(address string, exitChan chan int) {

	//根据给定地址进行侦听
	l,err := net.Listen("tcp",address)
	if err != nil {
		fmt.Println(err.Error())
		exitChan <- 1
	}

	fmt.Println("listen:",address)

	defer l.Close()

	for{

		//新连接没有到来时，Accept是阻塞的
		conn,err := l.Accept()

		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		go handleSession(conn,exitChan)
	}
}
