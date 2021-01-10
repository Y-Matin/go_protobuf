package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io"
	"io/ioutil"
	"log"
	"main/service"
	"math/rand"
	"time"
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

	//queryStockById(conn)

	//queryListByIds(conn)

	//queryObjectListByIds(conn)

	//UseTimestamp(conn)

	//batchQueryByNormal(conn)

	//batchQueryByServerStream(conn)

	//batchQueryByClientStream(conn)

	batchQueryByBothStream(conn)
}

//1.根据id 获取对应商品的 库存
func queryStockById(conn *grpc.ClientConn) {
	client := service.NewProdServiceClient(conn)
	prodResponse, err := client.GetProdStock(context.Background(), &service.ProdRequest{ProdId: 10, Areas: service.ProdAreas_C})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(prodResponse.ProdStock)
}

//  2.查询 固定数目的商品，返回一组商品的库存
func queryListByIds(conn *grpc.ClientConn) {
	client := service.NewProdServiceClient(conn)
	prodList, err := client.QueryProdStock(context.Background(), &service.QueryProd{Size: 1})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(prodList)
}

// 3. 根据某个id，查询商品具体信息（model）
func queryObjectListByIds(conn *grpc.ClientConn) {
	client := service.NewProdServiceClient(conn)
	info, _ := client.GetProdInfo(context.Background(), &service.ProdRequest{ProdId: 10})
	fmt.Println(info)
}

// 4. 使用时间格式（timestamp）
func UseTimestamp(conn *grpc.ClientConn) {
	t := timestamp.Timestamp{Seconds: time.Now().Unix()}
	orderClient := service.NewOrderServiceClient(conn)
	status, err := orderClient.GetOderStatus(context.Background(), &service.OrderRequest{
		OrderMain: &service.OrderMain{
			OrderId:    1,
			OrderPrice: 0,
			CreateTime: &t,
			Detail: []*service.OrderDetail{
				{
					OrderId:  10,
					DetailId: 11,
					ProdId:   12,
					ProdNum:  1,
				},
				{
					OrderId:  100,
					DetailId: 110,
					ProdId:   120,
					ProdNum:  10,
				},
			},
		},
	},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(status)
}

// 5. not流模式 ： 批量查询 传统方式
func batchQueryByNormal(conn *grpc.ClientConn) {
	orderServiceClient := service.NewUserServiceClient(conn)
	size := 10
	users := make([]*service.UserInfo, size)
	rand.Seed(time.Now().Unix())
	for i := 0; i < size; i++ {
		users[i] = &service.UserInfo{
			UserId: rand.Int31n(50),
		}
	}
	request := service.UserScoreRequest{Users: users}
	response, err := orderServiceClient.GetUserScore(context.Background(), &request)
	if err != nil {
		log.Fatal(err)
	}
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

// 6. 服务端流模式，
func batchQueryByServerStream(conn *grpc.ClientConn) {
	orderServiceClient := service.NewUserServiceClient(conn)
	size := 20
	users := make([]*service.UserInfo, size)
	rand.Seed(time.Now().Unix())
	for i := 0; i < size; i++ {
		users[i] = &service.UserInfo{
			UserId: rand.Int31n(50),
		}
	}
	request := service.UserScoreRequest{Users: users}
	stream, err := orderServiceClient.GetUserScoreByServerStream(context.Background(), &request)
	if err != nil {
		log.Fatal(err)
	}
	// 循环读取server端返回的流，读到数据直接处理，知道读到流的结束。
	for true {
		recv, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		// here can use goroutine to do the next step
		time.Sleep(time.Second * 2)
		fmt.Println(recv.Users)
	}
}

// 客户端流
func batchQueryByClientStream(conn *grpc.ClientConn) {
	orderServiceClient := service.NewUserServiceClient(conn)
	stream, _ := orderServiceClient.GetUserScoreByClientStream(context.Background())
	size := 20
	rand.Seed(time.Now().Unix())
	temp := make([]*service.UserInfo, 0)
	for i := 0; i < size; i++ {
		user := &service.UserInfo{
			UserId: rand.Int31n(50),
		}
		temp = append(temp, user)
		if i%2 == 1 {
			stream.Send(&service.UserScoreRequest{Users: temp})
			//clear silence
			temp = (temp)[0:0]
		}

	}
	if len(temp) != 0 {
		stream.Send(&service.UserScoreRequest{Users: temp})
	}
	recv, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(recv.Users)

}

func batchQueryByBothStream(conn *grpc.ClientConn) {
	client := service.NewUserServiceClient(conn)
	stream, err := client.GetUserScoreByBothStream(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	size := 20
	rand.Seed(time.Now().Unix())
	temp := make([]*service.UserInfo, 0)
	for i := 0; i < size; i++ {
		user := &service.UserInfo{
			UserId: rand.Int31n(50),
		}
		temp = append(temp, user)
		if i%2 == 1 {
			stream.Send(&service.UserScoreRequest{Users: temp})
			//clear silence
			temp = (temp)[0:0]
			recv, err := stream.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(recv.Users)
		}

	}
	stream.CloseSend()

	for true {
		recv, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(recv.Users)
	}
}
