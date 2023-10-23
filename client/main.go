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
	conn, err := grpc.Dial("10.159.10.78"+port,
		grpc.WithTransportCredentials(
			insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Didnt connect : %v", err)
	}

	defer conn.Close()

	client := pb.NewProductServiceClient(conn)
	CallGetProduct(client)
}
