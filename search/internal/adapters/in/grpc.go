package grpc

import (
	"context"
	"search/internal/application/domain"
	"search/internal/ports/in"
	searchpb "search/proto/golang"
)

type SearchHandler struct {
	searchpb.UnimplementedSearchServiceServer
	usecase in.SearchUseCase
}

func NewSearchHandler(usecase in.SearchUseCase) *SearchHandler {
	return &SearchHandler{usecase: usecase}
}

func (h *SearchHandler) SearchGame(ctx context.Context, req *searchpb.SearchGameRequest) (*searchpb.SearchGameResponse, error) {

	// Extract request to Game.
	queryGame := domain.Game{
		HomeTeam: req.GameRequest.HomeTeam,
		AwayTeam: req.GameRequest.AwayTeam,
		GameDate: req.GameRequest.GameDate,
	}

	results, err := h.usecase.Search(ctx, queryGame)
	if err != nil {
		return nil, err
	}

	// TODO: Construct the response
	resp := &searchpb.SearchGameResponse{}
	for _, g := range results {
		resp.Games = append(resp.Games, &searchpb.Game{
			HomeTeam: g.HomeTeam,
			AwayTeam: g.AwayTeam,
			GameDate: g.GameDate,
		})
	}

	return resp, nil

}
