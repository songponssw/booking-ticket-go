package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	searchpb "search/proto/golang"
)

func main() {
	conn, err := grpc.Dial(
		"localhost:50051",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := searchpb.NewMatchSearchServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.GetMatch(ctx, &searchpb.GetMatchRequest{
		GameRequest: &searchpb.Game{
			HomeTeam: "Arsenal",
			AwayTeam: "Chelsea",
			Gameday:  "2025-01-01",
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Response: %+v\n", resp.Games)
}

