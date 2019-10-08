package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "learngo/grpc-test/proto/hello"
	"log"
	"net"
)

const (
	Address = "127.0.0.1:50052"
)

type helloService struct {
}

var HelloService = helloService{}

func (h helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	resp := new(pb.HelloResponse)
	resp.Message = fmt.Sprintf("Hello %s", in.Name)
	return resp, nil
}

func main() {

	lis, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterHelloServer(s, helloService{})

	log.Println("Listen on " + Address)

	s.Serve(lis)

}
