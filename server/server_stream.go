package main

import (
	"log"
	"time"

	pb "inixindo.id/grpc/proto"
)

func (s *ProductServer) GetProductServerStreaming(req *pb.ProductIdList,
	stream pb.ProductService_GetProductServerStreamingServer) error {
	log.Printf("Got request with productId: %v", req.ProductIds)
	for _, productId := range req.ProductIds {
		res := &pb.ProductIdResponse{
			ProductId: "->" + productId,
		}

		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}
