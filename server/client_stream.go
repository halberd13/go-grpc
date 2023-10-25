package main

import (
	"io"
	"log"

	pb "inixindo.id/grpc/proto"
)

func (s *ProductServer) GetProductClientStreaming(stream pb.ProductService_GetProductClientStreamingServer) error {
	var products []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.ProductIdList{
				ProductIds: products,
			})
		}

		if err != nil {
			return err
		}

		log.Printf("Got request from client %v", req.ProductId)
		products = append(products, "->"+req.ProductId)
	}

}
