package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"time"

	"github.com/alexdcox/grpc-with-tls/schema"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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
	fmt.Println("Starting grpc server with tls")

	serverCert, err := tls.LoadX509KeyPair("certs/dcrwallet-rpc.cert", "certs/dcrwallet-rpc.key")
	if err != nil {
		logrus.Fatalf("%+v\n", errors.WithStack(err))
	}

	pemClientCA, err := ioutil.ReadFile("certs/dcrwallet-clients.pem")
	if err != nil {
		logrus.Fatalf("%+v\n", errors.WithStack(err))
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemClientCA) {
		logrus.Fatalf("failed to load certificate pool blah blah yada yada")
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	}

	tlsCredentials := credentials.NewTLS(config)

	grpcServer := grpc.NewServer(grpc.Creds(tlsCredentials))
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
