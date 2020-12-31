package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"main/service"
	"net"
)

func main() {
	creds, err := credentials.NewServerTLSFromFile("key/server.crt", "keys/server_no_passwd.key")
	if err != nil {
		log.Fatal(err)
	}

	rpcServer := grpc.NewServer(grpc.Creds(creds))
	service.RegisterProdServiceServer(rpcServer, new(service.ProdService))
	listen, _ := net.Listen("tcp", ":8081")
	rpcServer.Serve(listen)
}
