package service

import (
	"context"
	"errors"
	"search/internal/application/domain"
	"search/internal/ports/in"
	"search/internal/ports/out"
)

type SearchService struct {
	repo out.SearchRepository
}

func NewSearchService(repo out.SearchRepository) in.SearchUseCase {
	return &SearchService{repo: repo}
}

func (s *SearchService) Search(ctx context.Context, queryGame domain.Game) ([]domain.Game, error) {
	// bussiness rule: Reject if all fields are null
	if queryGame.HomeTeam == "" && queryGame.AwayTeam == "" && queryGame.GameDate == "" {
		return nil, errors.New("At least one fields should be provided")
	}

	return s.repo.Search(ctx, queryGame)
}
