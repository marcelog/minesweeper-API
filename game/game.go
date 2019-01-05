package game

import (
	"encoding/json"
)

// Game represents a game.
type Game struct {
	ID      int `json:"id"`
	OwnerID int `json:"owner_id"`
}

// New creates a new game.
func New(id int, ownerID int) *Game {
	return &Game{
		ID:      id,
		OwnerID: ownerID,
	}
}

// JSON serializes this user as a json string.
func (g *Game) JSON() string {
	j, _ := json.Marshal(g)
	return string(j)
}
