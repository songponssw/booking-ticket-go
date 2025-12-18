package postgres

import (
	"context"
	"database/sql"

	"search/internal/ports/out"
	"search/internal/application/domain"
)

type SearchRepository struct {
	db *sql.DB
}

func NewSearchRepository(db *sql.DB) out.SearchRepository {
	return &SearchRepository{db: db}
}

// func buildWhereClause(){
//
// }


func (r *SearchRepository) Search(
	ctx context.Context,
	query domain.Game,
) ([]domain.Game, error) {

	rows, err := r.db.QueryContext(
		ctx,
		`SELECT hometeamname, awayteamname  
		 FROM leagueschedule25_26
		 WHERE hometeamname ILIKE '%' || $1 || '%'`,
		query.HomeTeam,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var games []domain.Game
	for rows.Next() {
		var g domain.Game
                if err := rows.Scan(&g.HomeTeam, &g.AwayTeam); err != nil {
                        return nil, err
                }
		games = append(games, g)
	}

	return games, nil
}
