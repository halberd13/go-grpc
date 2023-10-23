package main

import (
	"context"

	pb "inixindo.id/grpc/proto"
)

func (s *ProductServer) GetProduct(ctx context.Context, req *pb.NoParam) (*pb.Product, error) {
	return &pb.Product{
		ProductId: "1",
		Name:      "Nokia 3310",
		Sku:       5,
	}, nil
}
