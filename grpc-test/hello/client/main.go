package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"learngo/grpc-test/proto/hello"
	"log"
)

const (
	Address = "127.0.0.1:50052"

	OpenTLS = true
)

type customerCredential struct {
}

func (c customerCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "101010",
		"appkey": "i am key",
	}, nil
}

func (c customerCredential) RequireTransportSecurity() bool {
	return OpenTLS
}

func main() {

	var err error
	var opts []grpc.DialOption

	if OpenTLS {
		creds, err := credentials.NewClientTLSFromFile("../learngo/grpc-test/hello/keys/server.pem", "zhang")
		if err != nil {
			log.Fatal(err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	opts = append(opts, grpc.WithPerRPCCredentials(new(customerCredential)))

	conn, err := grpc.Dial(Address, opts...)
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
