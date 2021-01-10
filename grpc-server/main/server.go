package main

import (
	"google.golang.org/grpc"
	"main/helper"
	"main/service"
	"net"
)

func main() {
	rpcServer := grpc.NewServer(grpc.Creds(helper.GetServerCred()))
	service.RegisterProdServiceServer(rpcServer, new(service.ProtoServiceImpl))
	service.RegisterOrderServiceServer(rpcServer, new(service.ProtoServiceImpl))
	service.RegisterUserServiceServer(rpcServer, new(service.ProtoServiceImpl))
	listen, _ := net.Listen("tcp", ":8081")
	_ = rpcServer.Serve(listen)
}
