package main

import (
	"context"
	"fmt"
	"github.com/bufbuild/connect-go"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	v1 "grpc-connect-go/gen/hello/v1"
	"grpc-connect-go/gen/hello/v1/hellov1connect"
	"log"
	"net/http"
)

type HelloServer struct{}

func main() {
	hello := &HelloServer{}
	mux := http.NewServeMux()
	path, handler := helloV1connect.NewHelloServiceHandler(hello)
	mux.Handle(path, handler)
	log.Println("path", path)
	if err := http.ListenAndServe(":3000", h2c.NewHandler(mux, &http2.Server{})); err != nil {
		panic(err)
	}
}

func (h *HelloServer) Hello(ctx context.Context, req *connect.Request[v1.HelloRequest]) (*connect.Response[v1.HelloResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&v1.HelloResponse{
		Hello: fmt.Sprintf("Hello, %s!", req.Msg.Name),
	})
	res.Header().Set("Hello-Version", "v1")
	return res, nil
}

func (h *HelloServer) HelloStream(ctx context.Context, req *connect.Request[v1.HelloRequest], res *connect.ServerStream[v1.HelloResponse]) error {
	log.Println("req", req.Msg.GetName())
	return res.Send(&v1.HelloResponse{
		Hello: "HelloStream: " + req.Msg.GetName(),
	})
}
