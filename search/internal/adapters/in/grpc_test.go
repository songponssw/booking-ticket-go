package grpc

import (
	"context"
	"testing"

	"search/internal/application/domain"
	// "search/internal/application/service"
	searchpb "search/proto/golang"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockSearchUseCase struct{ mock.Mock }

func (m *MockSearchUseCase) Search(ctx context.Context, query domain.Game) ([]domain.Game, error) {
	args := m.Called(ctx, query)
	return args.Get(0).([]domain.Game), args.Error(1)
}


// func TestGameHandler_Search(t *testing.T) {
// 	mockService := new(MockService)
// 	handler := NewSearchHandler(&service.SearchService{repo: mockService})
//
// 	req := &pb.SearchGameRequest{HomeTeam: "Lakers"}
// 	service.On("SearchGames", mock.Anything, domain.Game{HomeTeam: "Lakers"}).
// 		Return([]domain.Game{
// 			{HomeTeam: "Lakers", AwayTeam: "Heat"},
// 		}, nil)
//
// 	resp, err := handler.Search(context.Background(), req)
// 	assert.NoError(t, err)
// 	assert.Len(t, resp.Games, 1)
// 	assert.Equal(t, "Lakers", resp.Games[0].HomeTeam)
//
// 	service.AssertExpectations(t)
// }


func TestSearchHandler_Search_Success(t *testing.T) {
    // Arrange
    mockUsecase := new(MockSearchUseCase)
    handler := NewSearchHandler(mockUsecase)

    req := &searchpb.SearchGameRequest{
        GameRequest: &searchpb.Game{
            HomeTeam: "Lakers",
            AwayTeam: "",
            GameDate: "",
        },
    }

    expectedQuery := domain.Game{
        HomeTeam: "Lakers",
        AwayTeam: "",
        GameDate: "",
    }

    mockUsecase.
        On("Search", mock.Anything, expectedQuery).
        Return([]domain.Game{
            {
                HomeTeam: "Lakers",
                AwayTeam: "Heat",
                GameDate: "2025-12-25",
            },
        }, nil)

    // Act
    resp, err := handler.Search(context.Background(), req)

    // Assert
    assert.NoError(t, err)
    assert.NotNil(t, resp)
    assert.Len(t, resp.Games, 1)

    assert.Equal(t, "Lakers", resp.Games[0].HomeTeam)
    assert.Equal(t, "Heat", resp.Games[0].AwayTeam)
    assert.Equal(t, "2025-12-25", resp.Games[0].GameDate)

    mockUsecase.AssertExpectations(t)
}

