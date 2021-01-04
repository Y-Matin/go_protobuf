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
	/*	  //1.根据id 获取对应商品的 库存
	prodResponse, err := client.GetProdStock(context.Background(), &service.ProdRequest{ProdId: 10,Areas: service.ProdAreas_C})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(prodResponse.ProdStock)*/

	/*//  2.查询 固定数目的商品，返回一组商品的库存
	prodList, err := client.QueryProdStock(context.Background(), &service.QueryProd{Size: 1})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(prodList)*/

	// 3. 根据某个id，查询商品具体信息（model）
	info, err := client.GetProdInfo(context.Background(), &service.ProdRequest{ProdId: 10})
	fmt.Print(info)

	// 4. 使用时间格式（timestamp）

}
