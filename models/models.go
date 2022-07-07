package models

type Hotels struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Rooms    int    `json:"rooms"`
	Location string `json:"location"`
}
