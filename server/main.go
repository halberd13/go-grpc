package main

import (
	"log"
	"net"

	pb "inixindo.id/grpc/proto"

	"google.golang.org/grpc"
)

const (
	port = ":9090"
)

type ProductServer struct {
	pb.ProductServiceServer
}

func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start server %v", err)
	}

	// create a new gRPC Server
	grpcServer := grpc.NewServer()

	// register the ProductService
	pb.RegisterProductServiceServer(grpcServer, &ProductServer{})
	log.Println("Server Running")
	log.Printf("Server started at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
	log.Println("Server Shutdown")

}
