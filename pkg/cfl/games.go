package cfl

type Type struct {
	ID    int `json:"event_type_id"`
	Name  string
	Title string
}
type Status struct {
	ID     int `json:"event_status_id"`
	Name   string
	Active bool `json:"is_active"`
}

type Linescore struct {
	Quarter interface{} `json:"quarter"`
	Score   int         `json:"score"`
}

type TeamGame struct {
	ID         int `json:"team_id"`
	Location   string
	Nickname   string
	Score      int
	Linescores []Linescore `json:"linescores"`
	Home       bool        `json:"is_at_home"`
	Winner     bool        `json:"is_winner"`
}

type Game struct {
	ID         int `json:"game_id"`
	Number     int `json:"game_number"`
	Week       int `json:"week"`
	Season     int
	Attendance int
	Type       Type     `json:"event_type"`
	Status     Status   `json:"event_status"`
	Team1      TeamGame `json:"team_1"`
	Team2      TeamGame `json:"team_2"`
}

type Season struct {
	Games []Game `json:"data"`
}
