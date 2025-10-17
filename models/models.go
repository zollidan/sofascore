package models

// ============ MODELS ============

type Game struct {
	ID             int    `json:"id"`
	StartTimestamp int64  `json:"startTimestamp"`
	HomeTeamName   string
	AwayTeamName   string
	HomeScore      int
	AwayScore      int
	HasScore       bool
	StatusType     string
	TournamentName string
	Round          int
}