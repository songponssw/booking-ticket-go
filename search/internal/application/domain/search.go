package domain

type Game struct {
	// ID int
	HomeTeam string `json:"home_team"`
	AwayTeam string `json:"away_team"`
	GameDate string `json:"game_date"`
	// Arena string `"json:arena"`
}
