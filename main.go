//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative schema/schema.proto

package main

import (
	"fmt"
	"os"
)

const GRPCPort = 1234

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: grpc-with-tls [server|client]")
	}
	switch os.Args[1] {
	case "client":
		client()
	case "server":
		server()
	}
}
