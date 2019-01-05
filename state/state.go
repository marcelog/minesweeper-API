package state

import (
	"github.com/marcelog/minesweeper-API/user"
)

// State represents our server/game state (i.e: our "poor man's" persistence layer).
type State struct {
	Users []*user.User
}

// New creates a new state from scratch.
func New() *State {
	return &State{
		Users: []*user.User{},
	}
}
