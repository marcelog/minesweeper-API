package game

import (
	"encoding/json"
)

// Game represents a game.
type Game struct {
	ID      int `json:"id"`
	OwnerID int `json:"-"`
	Width   int `json:"width"`
	Height  int `json:"height"`
	Mines   int `json:"mines"`
}

// New creates a new game.
func New(id int, ownerID int, width int, height int, mines int) *Game {
	return &Game{
		ID:      id,
		OwnerID: ownerID,
		Width:   width,
		Height:  height,
		Mines:   mines,
	}
}

// JSON serializes this user as a json string.
func (g *Game) JSON() string {
	j, _ := json.Marshal(g)
	return string(j)
}
