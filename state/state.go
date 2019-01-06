package state

import (
	"github.com/marcelog/minesweeper-API/game"
	"github.com/marcelog/minesweeper-API/user"
)

// State represents our server/game state (i.e: our "poor man's" persistence layer).
type State struct {
	Users      map[int]*user.User
	Games      map[int]*game.Game
	NextUserID int
	NextGameID int
}

// AddUser adds a user.
func (s *State) AddUser() *user.User {
	u := user.New(s.NextUserID)
	s.NextUserID++
	s.Users[u.ID] = u
	return u
}

// AddGame adds a game.
func (s *State) AddGame(owner *user.User, width int, height int, mines int) (*game.Game, error) {
	g, err := game.New(s.NextGameID, owner.ID, width, height, mines)
	if err != nil {
		return nil, err
	}
	s.NextGameID++
	s.Games[g.ID] = g
	return g, nil
}

// New creates a new state from scratch.
func New() *State {
	return &State{
		NextUserID: 1,
		NextGameID: 1,
		Users:      map[int]*user.User{},
		Games:      map[int]*game.Game{},
	}
}

// FindUserByAPIKey returns a user (if found) with the given api key.
func (s *State) FindUserByAPIKey(key string) *user.User {
	for _, user := range s.Users {
		if user.APIKey == key {
			return user
		}
	}
	return nil
}

// FindGame returns a game (if found) for the given owner and id.
func (s *State) FindGame(owner *user.User, id int) *game.Game {
	for _, game := range s.Games {
		if game.OwnerID == owner.ID && game.ID == id {
			return game
		}
	}
	return nil

}
