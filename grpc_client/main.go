package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"main/service"
)

func main() {



	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := service.NewProdServiceClient(conn)
	prodResponse, err := client.GetProdStock(context.Background(), &service.ProdRequest{ProdId: 20})
	if err!= nil {
		log.Fatal(err)
	}
	fmt.Print(prodResponse.ProdStock)

}
