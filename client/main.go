package main

import (
	"log"

	pb "inixindo.id/grpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":9090"
)

func main() {
	conn, err := grpc.Dial("localhost"+port,
		grpc.WithTransportCredentials(
			insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Didnt connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewProductServiceClient(conn)

	products := &pb.ProductIdList{
		ProductIds: []string{"1", "2", "3", "4", "5", "6"},
	}

	// CallGetProduct(client)
	// CallGetProductClientStream(client, products)
	// CallGetProductServerStream(client, products)
	CallGetProductBidirectionalStreaming(client, products)

}
