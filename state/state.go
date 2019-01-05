package state

import (
	"github.com/marcelog/minesweeper-API/user"
)

// State represents our server/game state (i.e: our "poor man's" persistence layer).
type State struct {
	Users      map[int]*user.User
	NextUserID int
}

// AddUser adds a user.
func (s *State) AddUser() *user.User {
	u := user.New(s.NextUserID)
	s.NextUserID++
	s.Users[u.ID] = u
	return u
}

// New creates a new state from scratch.
func New() *State {
	return &State{
		NextUserID: 1,
		Users:      map[int]*user.User{},
	}
}

// FindByAPIKey returns a user (if found) with the given api key.
func (s *State) FindByAPIKey(key string) *user.User {
	for _, user := range s.Users {
		if user.APIKey == key {
			return user
		}
	}
	return nil
}
