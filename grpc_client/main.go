package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"main/service"
)

func main() {
	creds, err := credentials.NewClientTLSFromFile("keys/server.crt", "yeds")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := service.NewProdServiceClient(conn)
	prodResponse, err := client.GetProdStock(context.Background(), &service.ProdRequest{ProdId: 10})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(prodResponse.ProdStock)

}
