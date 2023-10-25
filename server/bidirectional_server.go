package main

import (
	"io"
	"log"

	pb "inixindo.id/grpc/proto"
)

func (s *ProductServer) GetProductBidirectionalStreaming(stream pb.ProductService_GetProductBidirectionalStreamingServer) error {
	// var products []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return nil
		}

		log.Printf("Got request from client %v", req.ProductId)
		res := &pb.ProductIdResponse{ProductId: "->" + req.ProductId}
		if err := stream.Send(res); err != nil {
			return err
		}

		// products = append(products, "->>>"+req.ProductId)
	}

}
