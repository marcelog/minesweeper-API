package server

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCantCreateGameWithInvalidParams(t *testing.T) {
	s, _, url := newServer(t)

	u := s.State.AddUser()

	// Bad width
	result, res, _ := runPost(t, url, "games", authHeader(u), `{"width": 65, "height": 8, "mines": 1}`)
	if res.StatusCode != 400 {
		t.Fatal("Unexpected status code:", res.StatusCode)
	}
	assertErrorMessage(t, result, "width too high")
	s.Stop()
}

func TestCantCreateGameWithBadJSON(t *testing.T) {
	s, _, url := newServer(t)

	u := s.State.AddUser()

	_, res, _ := runPost(t, url, "games", authHeader(u), "{")
	if res.StatusCode != 400 {
		t.Fatal("Unexpected status code:", res.StatusCode)
	}

	s.Stop()
}

func TestCreateGame(t *testing.T) {
	s, _, url := newServer(t)

	u := s.State.AddUser()

	result, res, err := runPost(t, url, "games", authHeader(u), `{"width": 8, "height": 8, "mines": 1}`)

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

	expected := `{"id":1,"width":8,"height":8,"mines":1,"state":"started","board":[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]}`

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

	s.Stop()
}

func TestErrorFlaggingInvalidGameInt(t *testing.T) {
	s, _, url := newServer(t)

	u := s.State.AddUser()
	result, res, _ := runPost(t, url, "games/blah/cells/3/flag", authHeader(u), "{}")
	if res.StatusCode != 400 {
		t.Fatal("Unexpected status code:", res.StatusCode)
	}
	assertErrorMessage(t, result, "invalid game id")
	s.Stop()
}

func TestErrorFlaggingInvalidCellInt(t *testing.T) {
	s, _, url := newServer(t)

	u := s.State.AddUser()
	result, res, _ := runPost(t, url, "games/3/cells/blah/flag", authHeader(u), "{}")
	if res.StatusCode != 400 {
		t.Fatal("Unexpected status code:", res.StatusCode)
	}
	assertErrorMessage(t, result, "invalid cell id")
	s.Stop()
}

func TestErrorFlaggingUnknownGame(t *testing.T) {
	var body map[string]interface{}
	s, _, url := newServer(t)

	// Create a game with a user, try to use it from a different user.
	u1 := s.State.AddUser()
	u2 := s.State.AddUser()

	result, _, _ := runPost(t, url, "games", authHeader(u1), `{"width": 8, "height": 8, "mines": 1}`)
	_ = json.Unmarshal([]byte(result), &body)
	gameID := int(body["id"].(float64))

	_, res, _ := runPost(t, url, fmt.Sprintf("games/%d/cells/3/flag", gameID), authHeader(u2), "{}")
	if res.StatusCode != 404 {
		t.Fatal("Unexpected status code:", res.StatusCode)
	}
	s.Stop()
}

func TestCanFlag(t *testing.T) {
	var body map[string]interface{}
	s, _, url := newServer(t)

	// Create a game with a user, try to use it from a different user.
	u1 := s.State.AddUser()
	result, _, _ := runPost(t, url, "games", authHeader(u1), `{"width": 8, "height": 8, "mines": 1}`)

	_ = json.Unmarshal([]byte(result), &body)
	gameID := int(body["id"].(float64))

	_, res, _ := runPost(t, url, fmt.Sprintf("games/%d/cells/3/flag", gameID), authHeader(u1), "{}")
	if res.StatusCode != 200 {
		t.Fatal("Unexpected status code:", res.StatusCode)
	}
	s.Stop()
}

func TestErrorUnflaggingInvalidGameInt(t *testing.T) {
	s, _, url := newServer(t)

	u := s.State.AddUser()
	result, res, _ := runDelete(t, url, "games/blah/cells/3/flag", authHeader(u), "{}")
	if res.StatusCode != 400 {
		t.Fatal("Unexpected status code:", res.StatusCode)
	}
	assertErrorMessage(t, result, "invalid game id")
	s.Stop()
}

func TestErrorUnflaggingInvalidCellInt(t *testing.T) {
	s, _, url := newServer(t)

	u := s.State.AddUser()
	result, res, _ := runDelete(t, url, "games/3/cells/blah/flag", authHeader(u), "{}")
	if res.StatusCode != 400 {
		t.Fatal("Unexpected status code:", res.StatusCode)
	}
	assertErrorMessage(t, result, "invalid cell id")
	s.Stop()
}

func TestErrorUnflaggingUnknownGame(t *testing.T) {
	var body map[string]interface{}
	s, _, url := newServer(t)

	// Create a game with a user, try to use it from a different user.
	u1 := s.State.AddUser()
	u2 := s.State.AddUser()

	result, _, _ := runPost(t, url, "games", authHeader(u1), `{"width": 8, "height": 8, "mines": 1}`)
	_ = json.Unmarshal([]byte(result), &body)
	gameID := int(body["id"].(float64))

	_, res, _ := runDelete(t, url, fmt.Sprintf("games/%d/cells/3/flag", gameID), authHeader(u2), "{}")
	if res.StatusCode != 404 {
		t.Fatal("Unexpected status code:", res.StatusCode)
	}
	s.Stop()
}

func TestCanUnflag(t *testing.T) {
	var body map[string]interface{}
	s, _, url := newServer(t)

	// Create a game with a user, try to use it from a different user.
	u1 := s.State.AddUser()

	result, _, _ := runPost(t, url, "games", authHeader(u1), `{"width": 8, "height": 8, "mines": 1}`)
	_ = json.Unmarshal([]byte(result), &body)
	gameID := int(body["id"].(float64))

	_, res, _ := runDelete(t, url, fmt.Sprintf("games/%d/cells/3/flag", gameID), authHeader(u1), "{}")
	if res.StatusCode != 200 {
		t.Fatal("Unexpected status code:", res.StatusCode)
	}
	s.Stop()
}
