package main

import (
	"context"
	"fmt"

	"github.com/alexdcox/grpc-with-tls/schema"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func client() {
	fmt.Println("Starting grpc client")

	grpcClient, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		logrus.Fatalf("%+v\n", errors.WithStack(err))
	}

	greeterClient := schema.NewGreeterClient(grpcClient)

	response, err := greeterClient.SayHello(context.Background(), &schema.HelloRequest{Name: "Alphi"})
	if err != nil {
		logrus.Fatalf("%+v\n", errors.WithStack(err))
	}

	fmt.Println("Client received response: ", response.Message)
}
