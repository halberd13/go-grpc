package main

import (
	"context"
	"io"
	"log"

	pb "inixindo.id/grpc/proto"
)

func CallGetProductServerStream(client pb.ProductServiceClient, products *pb.ProductIdList) {
	log.Printf("Streaming Started..!")
	stream, err := client.GetProductServerStreaming(context.Background(), products)
	if err != nil {
		log.Fatalf("Couldnt sent Product ID List")
	}

	for {
		productId, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while streaming: %v", err)
		}
		log.Println(productId)
	}
	log.Printf("Streaming Finished..!")
}
