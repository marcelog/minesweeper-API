package server

import (
	"testing"
)

func TestCreateUser(t *testing.T) {
	s, _, url := newServer(t)

	result, res, err := runPost(t, url, "users", map[string]string{}, "{}")

	if err != nil {
		t.Fatal("Unexpected error:", err.Error())
	}

	// Assert we get the right json payload & status code.
	if res.StatusCode != 201 {
		t.Fatal("Unexpected status code:", res.StatusCode)
	}

	if res.Header["Content-Type"][0] != "application/json" {
		t.Fatal("Unexpected content type:", res.Header["Content-Type"])
	}

	expected := "{\"id\":1,\"api_key\":\"apikey_1\"}"

	if result != expected {
		t.Fatal("Unexpected result:", result)
	}

	// Assert that the server now knows about this new user.
	if len(s.State.Users) != 1 {
		t.Fatal("Unexpected number of users:", len(s.State.Users))
	}

	if s.State.Users[1].JSON() != expected {
		t.Fatal("Unexpected state:", s.State.Users[1])
	}
}
