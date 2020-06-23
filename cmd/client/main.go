package main

import (
	"context"
	"fmt"
	"golang-grpc-password-generator/pkg/api"
	"google.golang.org/grpc"
)

func main() {
	var length int64
	var num bool
	var up bool

	fmt.Print("Password Length: ")
	_, _ = fmt.Scan(&length)
	fmt.Print("Include numbers (0 or 1): ")
	_, _ = fmt.Scan(&num)
	fmt.Print("Include upper chars (0 or 1): ")
	_, _ = fmt.Scan(&up)

	conn, err := grpc.Dial("127.0.0.1:8000", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	request := &api.GeneratorRequest{Length: length, Num: num, Up: up}

	client := api.NewGeneratorClient(conn)

	response, _ := client.Generator(context.Background(), request)

	fmt.Println(response)

}