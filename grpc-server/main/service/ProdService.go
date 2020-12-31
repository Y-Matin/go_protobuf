package service

import "context"

type ProdService struct {
}

func (p *ProdService) GetProdStock(context.Context, *ProdRequest) (*ProdResponse, error) {
	return &ProdResponse{ProdStock: 20}, nil
}
