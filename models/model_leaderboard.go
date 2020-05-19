package models

// Leaderboard Information on Top Users
// GET /leaderboard current leaderboard
// GET /leaderboard/name your alias on the leaderboard
type Leaderboard struct {
	Name       string  `json:"name"`
	IsRealName bool    `json:"isRealName,omitempty"`
	Profit     float64 `json:"profit,omitempty"`
}
