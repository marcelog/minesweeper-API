package user

import (
	"encoding/json"
	"fmt"
)

var id = 1

// User represents a user (player).
type User struct {
	ID     int    `json:"id"`
	APIKey string `json:"api_key"`
}

// New creates a new user.
func New() *User {
	newID := id
	id++
	return &User{
		ID:     newID,
		APIKey: fmt.Sprintf("apikey_%d", newID),
	}
}

// JSON serializes this user as a json string.
func (u *User) JSON() string {
	j, _ := json.Marshal(u)
	return string(j)
}
