package server

import (
	"testing"
)

func TestPing(t *testing.T) {
	result, res, err := runRequest(t, "ping")
	if err != nil {
		t.Fatal("Unexpected error:", err.Error())
	}

	if result != "PONG" {
		t.Fatal("Unexpected result:", result)
	}

	if res.Header["Content-Type"][0] != "text/plain" {
		t.Fatal("Unexpected content type:", res.Header["Content-Type"])
	}
}
