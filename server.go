package main

import (
	"google.golang.org/grpc"
	"main/service"
	"net"
)

func main() {
	rpcServer := grpc.NewServer()
	service.RegisterProdServiceServer(rpcServer, new(service.ProdService))
	listen, _ := net.Listen("tcp", ":8081")
	rpcServer.Serve(listen)

}
