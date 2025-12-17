package out

import (
	"context"
	"search/internal/application/domain"
)

type SearchRepository interface {
	Search(ctx context.Context, query domain.Game)([]domain.Game, error)
}
