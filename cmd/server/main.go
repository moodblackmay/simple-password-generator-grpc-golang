package main

import (
	"golang-grpc-password-generator/pkg/api"
	"golang-grpc-password-generator/pkg/generator"
	"google.golang.org/grpc"
	"net"
)

func main() {
	server := grpc.NewServer()
	generatorSrv := &generator.Generator{}

	api.RegisterGeneratorServer(server, generatorSrv)

	l, err := net.Listen("tcp", "127.0.0.1:8000")

	if err != nil {
		panic(err)
	}

	if err := server.Serve(l); err != nil {
		panic(err)
	}

}