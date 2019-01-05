package state

import (
	"github.com/marcelog/minesweeper-API/user"
)

// State represents our server/game state (i.e: our "poor man's" persistence layer).
type State struct {
	Users map[int]*user.User
}

// AddUser adds a user.
func (s *State) AddUser(u *user.User) {
	s.Users[u.ID] = u
}

// New creates a new state from scratch.
func New() *State {
	return &State{
		Users: map[int]*user.User{},
	}
}
