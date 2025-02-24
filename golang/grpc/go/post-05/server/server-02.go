package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"

	"github.com/liuliqiang/blog_codes/golang/grpc/go/post-05/interceptor"
	helloworld "github.com/liuliqiang/blog_codes/golang/grpc/go/proto-gens"
)

var (
	port *int
)

func main() {
	port = new(int)
	flag.IntVar(port, "port", 8080, "port to serve")
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				interceptor.UnaryServerDumpInterceptor()),
		),
	)
	helloworld.RegisterGreeterServer(grpcServer, &helloLiqiangIO{})
	grpcServer.Serve(lis)
}

// hello to https://liqiang.io
type helloLiqiangIO struct {
}

func (*helloLiqiangIO) SayHello(ctx context.Context, req *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{
		Message: fmt.Sprintf("Hello: %s, Welcome to https://liqiang.io.io", req.Name),
	}, nil
}
