package service

import (
	"context"
	"fmt"
)

type ProdService struct {
}

func (p *ProdService) GetProdStock(c context.Context, req *ProdRequest) (*ProdResponse, error) {
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

func (p *ProdService) QueryProdStock(ctx context.Context, q *QueryProd) (*ProdList, error) {
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

func (p ProdService) GetProdInfo(ctx context.Context, in *ProdRequest) (*ProdModel, error) {
	return &ProdModel{ProdId: in.ProdId, ProdName: "苹果", ProdPrice: 12.3}, nil
}

func (p *ProdService) GetOderStatus(ctx context.Context, order *OrderRequest) (*OrderStatus, error) {
	fmt.Print(order.OrderMain)
	return &OrderStatus{OrderMsg: "已完成", OrderStatus: "finish"}, nil

}
