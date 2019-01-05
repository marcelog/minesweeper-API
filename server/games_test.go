package server

import (
	"testing"
)

func TestCreateGame(t *testing.T) {
	s, _, url := newServer(t)

	u := s.State.AddUser()

	result, res, err := runPost(t, url, "games", authHeader(u), "{}")

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

	expected := "{\"id\":1}"

	if result != expected {
		t.Fatal("Unexpected result:", result)
	}

	// Assert that the server now knows about this new game.
	if len(s.State.Games) != 1 {
		t.Fatal("Unexpected number of games:", len(s.State.Games))
	}

	if s.State.Games[1].JSON() != expected {
		t.Fatal("Unexpected state:", s.State.Games[1])
	}

	if s.State.Games[1].OwnerID != 1 {
		t.Fatal("Unexpected owner:", s.State.Games[1])
	}
}
