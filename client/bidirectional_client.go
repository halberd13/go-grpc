package main

import (
	"context"
	"io"
	"log"

	pb "inixindo.id/grpc/proto"
)

func CallGetProductBidirectionalStreaming(client pb.ProductServiceClient, products *pb.ProductIdList) {
	log.Println("Client Streaming Started")
	stream, err := client.GetProductBidirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("Streaming Error: %v", err)
	}

	waitc := make(chan struct{})
	go func() {
		for {
			productId, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Couldnt send Product Id List: %v", err)
			}

			log.Printf("Product ID From server : %v\n", productId)
		}
		close(waitc)
	}()

	//kirim data products ke server
	for _, productId := range products.ProductIds {
		req := &pb.ProductIdRequest{
			ProductId: productId,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending %v", err)
		}
	}

	stream.CloseSend()
	<-waitc
	log.Printf("Bidirectional Streaming finished: %v", err)

}
