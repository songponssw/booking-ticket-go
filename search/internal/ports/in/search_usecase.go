package in

import (
	"context"
	"search/internal/application/domain"
)

type SearchUseCase interface {
	Search(ctx context.Context, queryGame domain.Game)([]domain.Game, error)
}
