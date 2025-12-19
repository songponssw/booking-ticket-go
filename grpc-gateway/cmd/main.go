package main

import (
	"context"
	"log"
	"net/http"
	// "time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	// gatewaypb "grpc-gateway/proto/golang" // generated package
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// gRPC-Gateway mux
	mux := runtime.NewServeMux()

	// gRPC client connection to Search service
	conn, err := grpc.DialContext(
		ctx,
		"localhost:50051", // Search gRPC server
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// ✅ Register REST handlers for BACKEND service (SearchService)
	err = searchpb.RegisterSearchServiceHandler(
		ctx,
		mux,
		conn,
	)
	if err != nil {
		log.Fatalf("failed to register gateway: %v", err)
	}

	// HTTP server
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("🚀 REST Gateway listening on :8080")
	log.Println("➡  Forwarding to Search gRPC at :50051")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
