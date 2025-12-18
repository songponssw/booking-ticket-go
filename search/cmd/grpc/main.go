package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "search/proto/golang"
	grpcAdapter "search/internal/adapters/in"
	postgresAdapter "search/internal/adapters/out"
	appService "search/internal/application/service"
	infraPostgres "search/internal/infra"
)

func main() {
	db, err := infraPostgres.NewDB(
		"postgres://user:pass@db/mydb?sslmode=disable",
	)
	if err != nil {
		log.Fatal(err)
	}

	repo := postgresAdapter.NewSearchRepository(db)
	usecase := appService.NewSearchService(repo)
	handler := grpcAdapter.NewSearchHandler(usecase)

	lis, _ := net.Listen("tcp", ":50051")
	server := grpc.NewServer()

	pb.RegisterSearchServiceServer(server, handler)

	log.Println("gRPC Search Service running on :50051")
	server.Serve(lis)
}
