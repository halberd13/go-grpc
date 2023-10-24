package main

import (
	"context"
	"io"
	"log"

	pb "inixindo.id/grpc/proto"
)

func CallGetProductClientStream(client pb.ProductServiceClient, products *pb.ProductIdList) {
	log.Println("Client Streaming Started")
	stream, err := client.GetProductServerStreaming(context.Background(), products)
	if err != nil {
		log.Fatalf("Couldnt sent productIdList: %v", err)
	}
	for {
		productId, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Couldnt send Product Id List: %v", err)
		}

		log.Println(productId)
	}
	log.Println("Streaming finished")
}
