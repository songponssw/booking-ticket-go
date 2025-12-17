package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	searchpb "search/proto/golang"
	"search/internal/application"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	searchpb.RegisterMatchSearchServiceServer(
		grpcServer,
		&application.SearchService{},
	)

	log.Println("gRPC server listening on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

