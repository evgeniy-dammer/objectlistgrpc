package handlers

import (
	"context"

	productservice "github.com/evgeniy-dammer/objectsgrpc/productservice/proto"
)

type ProductServiceServer struct {
	productservice.UnimplementedProductServiceServer
}

func (*ProductServiceServer) FindAll(ctx context.Context, in *productservice.FindAllRequest) (*productservice.FindAllResponse, error) {
	return &productservice.FindAllResponse{
		Products: []*productservice.Product{
			{
				Id:       "p01",
				Name:     "name 1",
				Price:    4.5,
				Quantity: 4,
				Status:   true,
			},
			{
				Id:       "p02",
				Name:     "name 2",
				Price:    22,
				Quantity: 3,
				Status:   false,
			},
			{
				Id:       "p03",
				Name:     "name 3",
				Price:    27,
				Quantity: 3,
				Status:   true,
			},
		},
	}, nil
}
