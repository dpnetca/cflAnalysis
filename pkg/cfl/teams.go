package cfl

type Team struct {
	ID   int    `json:"team_id"`
	Name string `json:"full_name"`
}
type Teams struct {
	Team []Team `json:"data"`
}
