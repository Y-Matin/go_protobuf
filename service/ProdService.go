package service

import "context"

type ProdService struct {
}

func (p *ProdService) GetProdStock(context.Context, *ProdRequest) (*ProdResponse, error) {
	return new(ProdResponse), nil
}
