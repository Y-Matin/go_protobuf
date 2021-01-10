package service

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"time"
)

type ProtoServiceImpl struct {
}

func (p *ProtoServiceImpl) GetProdStock(c context.Context, req *ProdRequest) (*ProdResponse, error) {
	switch req.Areas {
	case ProdAreas_A:
		return &ProdResponse{ProdStock: 20}, nil
	case ProdAreas_B:
		return &ProdResponse{ProdStock: 30}, nil
	case ProdAreas_C:
		return &ProdResponse{ProdStock: 40}, nil
	}
	return nil, nil
}

func (p *ProtoServiceImpl) QueryProdStock(ctx context.Context, q *QueryProd) (*ProdList, error) {
	stocks := []*ProdResponse{
		{ProdStock: 20},
		{ProdStock: 30},
		{ProdStock: 40},
		{ProdStock: 50},
		{ProdStock: 60},
	}
	list := ProdList{Prods: stocks}
	return &list, nil
}

func (p ProtoServiceImpl) GetProdInfo(ctx context.Context, in *ProdRequest) (*ProdModel, error) {
	return &ProdModel{ProdId: in.ProdId, ProdName: "苹果", ProdPrice: 12.3}, nil
}

func (p *ProtoServiceImpl) GetOderStatus(ctx context.Context, order *OrderRequest) (*OrderStatus, error) {
	fmt.Println(order.OrderMain)
	return &OrderStatus{OrderMsg: "已完成", OrderStatus: "finish"}, nil
}

/** 普通方法 */
func (p *ProtoServiceImpl) GetUserScore(c context.Context, req *UserScoreRequest) (*UserScoreResponse, error) {
	size := len(req.Users)

	r := rand.New(rand.NewSource(time.Now().Unix()))

	result := make([]*UserInfo, size)
	for i := 0; i < size; i++ {
		temp := &UserInfo{
			UserId:    req.Users[i].UserId,
			UserScore: r.Int31n(100),
		}
		result[i] = temp
	}

	return &UserScoreResponse{
		Users: result,
	}, nil
}

// 服务端流
func (p *ProtoServiceImpl) GetUserScoreByServerStream(req *UserScoreRequest, stream UserService_GetUserScoreByServerStreamServer) error {
	size := len(req.Users)

	r := rand.New(rand.NewSource(time.Now().Unix()))

	result := make([]*UserInfo, 0)
	for i := 0; i < size; i++ {
		temp := &UserInfo{
			UserId:    req.Users[i].UserId,
			UserScore: r.Int31n(100),
		}
		result = append(result, temp)
		if i%2 == 1 {
			err := stream.Send(&UserScoreResponse{Users: result})
			if err != nil {
				return err
			}
			result = nil
			result = result[0:0]
		}

	}
	if len(result) != 0 {
		stream.Send(&UserScoreResponse{
			Users: result,
		})
	}
	return nil
}

// 客户端流
func (p *ProtoServiceImpl) GetUserScoreByClientStream(stream UserService_GetUserScoreByClientStreamServer) error {
	result := make([]*UserInfo, 0)
	for true {
		recv, err := stream.Recv()
		if err == io.EOF {
			// 流结束,返回整个结果
			return stream.SendAndClose(&UserScoreResponse{
				Users: result,
			})
		} else if err != nil {
			log.Fatal(err)
		}
		size := len(recv.Users)
		time.Sleep(time.Second)
		fmt.Println(recv.Users)

		r := rand.New(rand.NewSource(time.Now().Unix()))
		for i := 0; i < size; i++ {
			temp := &UserInfo{
				UserId:    recv.Users[i].UserId,
				UserScore: r.Int31n(100),
			}
			result = append(result, temp)
		}

	}
	return nil
}

//双向流
func (p *ProtoServiceImpl) GetUserScoreByBothStream(stream UserService_GetUserScoreByBothStreamServer) error {
	result := make([]*UserInfo, 0)
	for true {
		recv, err := stream.Recv()
		if err == io.EOF {
			// 流结束,返回整个结果
			return nil
		} else if err != nil {
			log.Fatal(err)
		}
		size := len(recv.Users)
		fmt.Println(recv.Users)

		r := rand.New(rand.NewSource(time.Now().Unix()))
		for i := 0; i < size; i++ {
			temp := &UserInfo{
				UserId:    recv.Users[i].UserId,
				UserScore: r.Int31n(100),
			}
			result = append(result, temp)
		}
		stream.Send(&UserScoreResponse{Users: result})
		result = (result)[0:0]

	}
	return nil
}
