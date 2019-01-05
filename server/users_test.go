package server

import (
	"testing"
)

func TestCreateUser(t *testing.T) {
	server, result, res, err := runPost(t, "users", "{}")
	if err != nil {
		t.Fatal("Unexpected error:", err.Error())
	}

	if res.Header["Content-Type"][0] != "application/json" {
		t.Fatal("Unexpected content type:", res.Header["Content-Type"])
	}

	expected := "{\"id\":1,\"api_key\":\"apikey_1\"}"

	// Assert we get the right json payload & status code.
	if res.StatusCode != 201 {
		t.Fatal("Unexpected status code:", res.StatusCode)
	}

	if result != expected {
		t.Fatal("Unexpected result:", result)
	}

	// Assert that the server now knows about this new user.
	if len(server.State.Users) != 1 {
		t.Fatal("Unexpected number of users:", len(server.State.Users))
	}

	if server.State.Users[1].JSON() != expected {
		t.Fatal("Unexpected state:", server.State.Users[1])
	}
}
