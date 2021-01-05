package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	"main/helper"
	"main/service"
	"net/http"
)

func main() {
	mux := runtime.NewServeMux()
	options := []grpc.DialOption{grpc.WithTransportCredentials(helper.GetClientCred())}
	err := service.RegisterProdServiceHandlerFromEndpoint(context.Background(), mux, "localhost:8081", options)
	if err != nil {
		log.Fatal(err)
	}
	err = service.RegisterOrderServiceHandlerFromEndpoint(context.Background(), mux, "localhost:8081", options)
	if err != nil {
		log.Fatal(err)
	}
	httpserver := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	_ = httpserver.ListenAndServe()
	// http.ListenAndServe(":8081", mux)

}
