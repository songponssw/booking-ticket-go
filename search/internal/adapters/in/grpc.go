package grpc

import (
	"context"
	"search/internal/application/domain"
	"search/internal/application/ports/in"
	searchpb "search/proto/golang"
)

type SearchHandler struct {
	searchpb.UnimplementedMatchSearchServiceServer
	usecase in.SearchUseCase
}

func NewSearchHandler(usecase in.SearchUseCase) *SearchHandler {
	return &SearchHandler{usecase: usecase}
}

func (h *SearchHandler) Search(ctx context.Context, req *searchpb.SearchRequest) (*searchpb.SearchResponse, error) {

	// Extract request to Game.
	queryGame := damain.Game{
		HomeTeam: req.GameRequest.HomeTeam,
		AwayTeam: req.GameRequest.AwayTeam,
		GameDate: req.GameRequest.GameDate,
	}

	result, err := h.usecase.Search(ctx, queryGame)
	if err != nil {
		return nil, err
	}

	// TODO: Construct the response
	resp := &searchpb.SearchResponse{}
	for _, g := range results {
		resp.Games = append(resp.Games, &searchpb.Game{
			HomeTeam: g.HomeTeam,
			AwayTeam: g.AwayTeam,
			GameDate: g.GameDate,
		})
	}

	return resp, nil

}
