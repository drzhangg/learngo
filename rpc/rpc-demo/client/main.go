package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("rpc dial error:", err)
	}

	var reply string
	//err = client.Call("HelloService.Hello", "this is client", &reply)
	err = client.Call(HelloServiceName, "this is client", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
