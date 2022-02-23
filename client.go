package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"

	"github.com/alexdcox/grpc-with-tls/schema"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func client() {
	fmt.Println("Starting grpc client with tls")

	pemServerCA, err := ioutil.ReadFile("certs/ca.cert")
	if err != nil {
	    logrus.Fatalf("%+v\n", errors.WithStack(err))
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		logrus.Fatalf("failed to add to cert pool\n")
	}

	config := &tls.Config{
		RootCAs:      certPool,
	}

	tlsCredentials := credentials.NewTLS(config)

	grpcClient, err := grpc.Dial("localhost:1234", grpc.WithTransportCredentials(tlsCredentials))
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
