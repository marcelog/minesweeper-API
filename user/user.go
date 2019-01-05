package user

import (
	"encoding/json"
)

// User represents a user (player).
type User struct {
	ID     int
	APIKey string
}

// JSON serializes this user as a json string.
func (u *User) JSON() string {
	j, _ := json.Marshal(u)
	return string(j)
}
