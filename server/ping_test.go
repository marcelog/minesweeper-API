package server

import (
	"testing"
)

func TestPing(t *testing.T) {
	s, _, url := newServer(t)

	result, res, err := runGet(t, url, "ping", map[string]string{})
	if err != nil {
		t.Fatal("Unexpected error:", err.Error())
	}

	if res.StatusCode != 200 {
		t.Fatal("Unexpected status code:", res.StatusCode)
	}

	if res.Header["Content-Type"][0] != "text/plain" {
		t.Fatal("Unexpected content type:", res.Header["Content-Type"])
	}

	if result != "PONG" {
		t.Fatal("Unexpected result:", result)
	}

	s.Stop()
}
