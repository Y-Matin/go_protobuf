package service

import "context"

type ProdService struct {
}

func (p *ProdService) GetProdStock(context.Context, *ProdRequest) (*ProdResponse, error) {
	return &ProdResponse{ProdStock: 20}, nil
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
