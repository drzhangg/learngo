package main

import (
	"context"
	"google.golang.org/grpc"
	"learngo/grpc-test/proto/hello"
	"log"
)

const (
	Address = "127.0.0.1:50052"
)

func main() {
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := hello.NewHelloClient(conn)

	req := &hello.HelloRequest{Name: "golang"}
	resp, err := client.SayHello(context.TODO(), req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp.Message)
}
