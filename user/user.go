package user

import (
	"encoding/json"
	"fmt"
)

// User represents a user (player).
type User struct {
	ID     int    `json:"id"`
	APIKey string `json:"api_key"`
}

// New creates a new user.
func New(id int) *User {
	return &User{
		ID:     id,
		APIKey: fmt.Sprintf("apikey_%d", id),
	}
}

// JSON serializes this user as a json string.
func (u *User) JSON() string {
	j, _ := json.Marshal(u)
	return string(j)
}
