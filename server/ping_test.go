package server

import (
	"testing"
)

func TestPing(t *testing.T) {
	result, err := runRequest(t, "ping")
	if err != nil {
		t.Fatal("Unexpected error:", err.Error())
	}

	if result != "PONG" {
		t.Fatal("Unexpected result:", result)
	}
}
