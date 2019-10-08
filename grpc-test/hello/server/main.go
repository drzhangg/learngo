package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
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

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, grpc.Errorf(codes.Unauthenticated, "无Token认证信息")
	}

	var (
		appid  string
		appkey string
	)

	if val, ok := md["appid"]; ok {
		appid = val[0]
	}
	if val, ok := md["appkey"]; ok {
		appkey = val[0]
	}

	if appid != "101010" || appkey != "i am key" {
		return nil,grpc.Errorf(codes.Unauthenticated, "Token认证信息无效: appid=%s, appkey=%s", appid, appkey)
	}
	
	resp := new(pb.HelloResponse)
	resp.Message = fmt.Sprintf("Hello %s.\nToken info: appid=%s,appkey=%s", in.Name, appid, appkey)
	return resp, nil
}

func main() {

	lis, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatal(err)
	}

	//TLS认证
	creds, err := credentials.NewServerTLSFromFile("../learngo/grpc-test/hello/keys/server.pem", "../learngo/grpc-test/hello/keys/server.key")
	if err != nil {
		log.Fatal(err)
	}

	//实例化grpc server，并开启TLS认证
	s := grpc.NewServer(grpc.Creds(creds))
	//注册helloService
	pb.RegisterHelloServer(s, helloService{})

	log.Println("Listen on " + Address + " with TLS")

	s.Serve(lis)

}
