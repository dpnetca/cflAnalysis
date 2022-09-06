package cfl

type TeamStanding struct {
	DivisionID        int    `json:"division_id"`
	DivisionName      string `json:"division_name"`
	Place             int    `json:"place"`
	Flags             string `json:"flags"`
	ID                int    `json:"team_id"`
	Letter            string `json:"letter"`
	Abbreviation      string `json:"abbreviation"`
	Location          string `json:"location"`
	Nickname          string `json:"nickname"`
	FullName          string `json:"full_name"`
	GamesPlayed       int    `json:"games_played"`
	Wins              int    `json:"wins"`
	Losses            int    `json:"losses"`
	Ties              int    `json:"ties"`
	Points            int    `json:"points"`
	WinningPercentage int    `json:"winning_percentage"`
	PointsFor         int    `json:"points_for"`
	PointsAgainst     int    `json:"points_against"`
	HomeWins          int    `json:"home_wins"`
	HomeLosses        int    `json:"home_losses"`
	HomeTies          int    `json:"home_ties"`
	AwayWins          int    `json:"away_wins"`
	AwayLosses        int    `json:"away_losses"`
	AwayTies          int    `json:"away_ties"`
	DivisionWins      int    `json:"division_wins"`
	DivisionLosses    int    `json:"division_losses"`
	DivisionTies      int    `json:"division_ties"`
	Season            int    `json:"season"`
}

type Division struct {
	DivisionID   int            `json:"division_id"`
	DivisionName string         `json:"division_name"`
	DivisionSlug string         `json:"division_slug"`
	Standings    []TeamStanding `json:"standings"`
}

type Divisions struct {
	West Division `json:"west"`
	East Division `json:"east"`
}

type Standings struct {
	Data struct {
		Divisions Divisions `json:"divisions"`
	} `json:"data"`
}
