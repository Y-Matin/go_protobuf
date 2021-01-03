package main

import (
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"main/service"
	"net"
)

func main() {
	/*	creds, err := credentials.NewServerTLSFromFile("key/server.crt", "keys/server_no_passwd.key")
		if err != nil {
			log.Fatal(err)
		}
	*/
	cert, _ := tls.LoadX509KeyPair("cert/server.pem", "cert/server.key")
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("cert/ca.pem")
	certPool.AppendCertsFromPEM(ca)
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	})

	rpcServer := grpc.NewServer(grpc.Creds(creds))
	service.RegisterProdServiceServer(rpcServer, new(service.ProdService))
	listen, _ := net.Listen("tcp", ":8081")
	rpcServer.Serve(listen)
}
