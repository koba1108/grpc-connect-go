// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: hello/v1/hello.proto

package helloV1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "grpc-connect-go/gen/hello/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// HelloServiceName is the fully-qualified name of the HelloService service.
	HelloServiceName = "hello.v1.HelloService"
)

// HelloServiceClient is a client for the hello.v1.HelloService service.
type HelloServiceClient interface {
	Hello(context.Context, *connect_go.Request[v1.HelloRequest]) (*connect_go.Response[v1.HelloResponse], error)
	HelloStream(context.Context, *connect_go.Request[v1.HelloRequest]) (*connect_go.ServerStreamForClient[v1.HelloResponse], error)
}

// NewHelloServiceClient constructs a client for the hello.v1.HelloService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewHelloServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) HelloServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &helloServiceClient{
		hello: connect_go.NewClient[v1.HelloRequest, v1.HelloResponse](
			httpClient,
			baseURL+"/hello.v1.HelloService/Hello",
			opts...,
		),
		helloStream: connect_go.NewClient[v1.HelloRequest, v1.HelloResponse](
			httpClient,
			baseURL+"/hello.v1.HelloService/HelloStream",
			opts...,
		),
	}
}

// helloServiceClient implements HelloServiceClient.
type helloServiceClient struct {
	hello       *connect_go.Client[v1.HelloRequest, v1.HelloResponse]
	helloStream *connect_go.Client[v1.HelloRequest, v1.HelloResponse]
}

// Hello calls hello.v1.HelloService.Hello.
func (c *helloServiceClient) Hello(ctx context.Context, req *connect_go.Request[v1.HelloRequest]) (*connect_go.Response[v1.HelloResponse], error) {
	return c.hello.CallUnary(ctx, req)
}

// HelloStream calls hello.v1.HelloService.HelloStream.
func (c *helloServiceClient) HelloStream(ctx context.Context, req *connect_go.Request[v1.HelloRequest]) (*connect_go.ServerStreamForClient[v1.HelloResponse], error) {
	return c.helloStream.CallServerStream(ctx, req)
}

// HelloServiceHandler is an implementation of the hello.v1.HelloService service.
type HelloServiceHandler interface {
	Hello(context.Context, *connect_go.Request[v1.HelloRequest]) (*connect_go.Response[v1.HelloResponse], error)
	HelloStream(context.Context, *connect_go.Request[v1.HelloRequest], *connect_go.ServerStream[v1.HelloResponse]) error
}

// NewHelloServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewHelloServiceHandler(svc HelloServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/hello.v1.HelloService/Hello", connect_go.NewUnaryHandler(
		"/hello.v1.HelloService/Hello",
		svc.Hello,
		opts...,
	))
	mux.Handle("/hello.v1.HelloService/HelloStream", connect_go.NewServerStreamHandler(
		"/hello.v1.HelloService/HelloStream",
		svc.HelloStream,
		opts...,
	))
	return "/hello.v1.HelloService/", mux
}

// UnimplementedHelloServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedHelloServiceHandler struct{}

func (UnimplementedHelloServiceHandler) Hello(context.Context, *connect_go.Request[v1.HelloRequest]) (*connect_go.Response[v1.HelloResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("hello.v1.HelloService.Hello is not implemented"))
}

func (UnimplementedHelloServiceHandler) HelloStream(context.Context, *connect_go.Request[v1.HelloRequest], *connect_go.ServerStream[v1.HelloResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("hello.v1.HelloService.HelloStream is not implemented"))
}
