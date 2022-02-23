package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/alexdcox/grpc-with-tls/schema"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Greeter struct {
	schema.UnimplementedGreeterServer
}

func (g Greeter) SayHello(ctx context.Context, request *schema.HelloRequest) (response *schema.HelloReply, err error) {
	response = &schema.HelloReply{}
	response.Message = fmt.Sprintf("%s.%d", request.Name, time.Now().Unix())
	return
}

func server() {
	fmt.Println("Starting grpc server")

	grpcServer := grpc.NewServer()
	greeter := new(Greeter)
	schema.RegisterGreeterServer(grpcServer, greeter)

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", GRPCPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	err = grpcServer.Serve(lis)
	if err != nil {
	    logrus.Fatalf("%+v\n", errors.WithStack(err))
	}

	fmt.Println("Done")
}