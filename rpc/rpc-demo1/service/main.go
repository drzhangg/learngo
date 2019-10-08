package main

import (
	"fmt"
	"net"
	"net/rpc"
)

const HelloServiceName = "HelloService"

type HelloServiceInterface =  interface {
	Hello(request string, reply *string) error
}

func Hello(request string, reply *string) error {
	*reply = "Hello " + request
	return nil
}

func main() {

	rpc.RegisterName(HelloServiceName, new(HelloServiceInterface))

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println(err)
	}

	conn, err := lis.Accept()
	if err != nil {
		fmt.Println(err)
	}

	rpc.ServeConn(conn)

}
