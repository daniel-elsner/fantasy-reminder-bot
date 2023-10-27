package models

type Game struct {
	GameID      string `json:"gameId"`
	GameTimeUTC string `json:"gameTimeUTC"`
	GameTimeET  string `json:"gameTimeET"`
	HomeTeam    Team   `json:"homeTeam"`
	AwayTeam    Team   `json:"awayTeam"`
}
