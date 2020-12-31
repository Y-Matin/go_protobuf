package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"main/service"
)

func main() {

	// 只有 服务端证书
	/*	creds, err := credentials.NewClientTLSFromFile("keys/server.crt", "yeds")
		if err != nil {
			log.Fatal(err)
		}*/

	// 服务端、客户端证书都有
	cert, _ := tls.LoadX509KeyPair("cert/client.pem", "cert/client.key")
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("cert/ca.pem")
	certPool.AppendCertsFromPEM(ca)
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   "localhost",
		RootCAs:      certPool,
	})

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
