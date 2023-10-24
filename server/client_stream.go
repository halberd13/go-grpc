package main

import (
	"io"
	"log"

	pb "inixindo.id/grpc/proto"
)

// type ProductIdList struct {
// 	string ProductId
// }

func (s *ProductServer) GetProductClientStreaming(stream pb.ProductService_GetProductClientStreamingServer) error {
	var products []*pb.Product
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.ProductList{Product: products})
		}

		if err != nil {
			return err
		}

		log.Printf("Got request for client %v", req.ProductId)
		// products = append(products, req.ProductId)
	}
}
