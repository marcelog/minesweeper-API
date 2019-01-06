package server

import (
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

	result, res, _ = runPost(t, url, "games", authHeader(u), `{"width": 7, "height": 8, "mines": 1}`)
	if res.StatusCode != 400 {
		t.Fatal("Unexpected status code:", res.StatusCode)
	}
	assertErrorMessage(t, result, "width too low")

	// Bad height
	result, res, _ = runPost(t, url, "games", authHeader(u), `{"width": 8, "height": 65, "mines": 1}`)
	if res.StatusCode != 400 {
		t.Fatal("Unexpected status code:", res.StatusCode)
	}
	assertErrorMessage(t, result, "height too high")

	result, res, _ = runPost(t, url, "games", authHeader(u), `{"width": 8, "height": 7, "mines": 1}`)
	if res.StatusCode != 400 {
		t.Fatal("Unexpected status code:", res.StatusCode)
	}
	assertErrorMessage(t, result, "height too low")

	// Bad # of mines
	result, res, _ = runPost(t, url, "games", authHeader(u), `{"width": 8, "height": 8, "mines": 0}`)
	if res.StatusCode != 400 {
		t.Fatal("Unexpected status code:", res.StatusCode)
	}
	assertErrorMessage(t, result, "mine number too low")

	result, res, _ = runPost(t, url, "games", authHeader(u), `{"width": 8, "height": 8, "mines": 33}`)
	if res.StatusCode != 400 {
		t.Fatal("Unexpected status code:", res.StatusCode)
	}
	assertErrorMessage(t, result, "mine number too high")

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
