package main

import (
	"context"
	"github.com/bufbuild/connect-go"
	v1 "grpc-connect-go/gen/hello/v1"
	"grpc-connect-go/gen/hello/v1/hellov1connect"
	"log"
	"net/http"
)

const serverURL = "http://localhost:3000"

func main() {
	client := helloV1connect.NewHelloServiceClient(http.DefaultClient, serverURL, connect.WithGRPCWeb())
	res, err := client.Hello(context.Background(), connect.NewRequest(&v1.HelloRequest{
		Name: "Jane",
	}))
	if err != nil {
		panic(err)
	}
	log.Println(res.Msg.Hello)
}
