package main

import (
	"context"
	"log"
	"time"

	pb "inixindo.id/grpc/proto"
)

func CallGetProduct(client pb.ProductServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.GetProduct(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("Couldn't get product :%v", err)
	}
	log.Printf("The Product : %v", res)
}
