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

	client := searchpb.NewSearchServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.SearchGame(ctx, &searchpb.SearchGameRequest{
		GameRequest: &searchpb.Game{
			HomeTeam: "Nuggets",
			AwayTeam: "",
			GameDate:  "",
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Response: %+v\n", resp.Games)
}

