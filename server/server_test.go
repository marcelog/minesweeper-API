package server

import (
	"net/http"
	"testing"
)

func TestErrorOnListen(t *testing.T) {
	s := New(&Args{
		Address: "127.0.0.1",
		Port:    65536,
	})
	err := s.Run()
	if err == nil {
		t.Fatal("Expected an error")
	}

	if err.Error() != "listen tcp4: address 65536: invalid port" {
		t.Fatal("Unexpected error:", err.Error())
	}
}

func TestListenServerAndShutdown(t *testing.T) {
	s := New(&Args{
		Address: "127.0.0.1",
		Port:    10000,
	})
	err := s.Run()
	if err != nil {
		t.Fatal("Unexpected error:", err.Error())
	}

	res, err := http.Get("http://127.0.0.1:10000/unknown_endpoint")
	if err != nil {
		t.Fatal("Unexpected error:", err.Error())
	}
	res.Body.Close()

	if res.StatusCode != 404 {
		t.Fatal("Unexpected status code:", res.StatusCode)
	}

	s.Stop()
	_, err = http.Get("http://127.0.0.1:10000/unknown_endpoint")
	if err == nil {
		t.Fatal("Expected an error")
	}
}
