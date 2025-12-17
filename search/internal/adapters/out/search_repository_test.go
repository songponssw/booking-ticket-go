package postgres

import (
    "context"
    "testing"

    "github.com/DATA-DOG/go-sqlmock"
    "github.com/stretchr/testify/assert"

    "search/internal/application/domain"
)

func TestSearchRepository_Search(t *testing.T) {
    // 1. Create mock DB
    db, mock, err := sqlmock.New()
    assert.NoError(t, err)
    defer db.Close()

    repo := NewSearchRepository(db)

    // 2. Table-driven test cases
    tests := []struct {
        name       string
        query      domain.Game
        mockRows   *sqlmock.Rows
        wantResult []domain.Game
    }{
        {
            name:  "Search Lakers",
            query: domain.Game{
                HomeTeam: "Lakers",
                AwayTeam: "",
		GameDate: "",
            },
            mockRows: sqlmock.NewRows([]string{"hometeamname", "awayteamname"}).
                AddRow("Lakers", "Heat").
                AddRow("Lakers", "Celtics"),
            wantResult: []domain.Game{
                {HomeTeam: "Lakers", AwayTeam: "Heat"},
                {HomeTeam: "Lakers", AwayTeam: "Celtics"},
            },
        },
        {
            name:  "Search Celtics",
            query: domain.Game{
                HomeTeam: "Celtics",
                AwayTeam: "",
		GameDate: "",
            },
            mockRows: sqlmock.NewRows([]string{"hometeamname", "awayteamname"}).
                AddRow("Celtics", "Bulls"),
            wantResult: []domain.Game{
                {HomeTeam: "Celtics", AwayTeam: "Bulls"},
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            mock.ExpectQuery(`SELECT hometeamname, awayteamname`).
                WithArgs(tt.query.HomeTeam).
                WillReturnRows(tt.mockRows)

            // Call repository
            result, err := repo.Search(context.Background(), tt.query)
            assert.NoError(t, err)
            assert.Equal(t, tt.wantResult, result)

            // Ensure all expectations were met
            assert.NoError(t, mock.ExpectationsWereMet())
        })
    }
}

