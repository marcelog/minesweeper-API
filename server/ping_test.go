package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

var port = 8000

func newServer() (*Server, int, string) {
	p := port
	port++

	s := New(&Args{
		Address: "127.0.0.1",
		Port:    p,
	})
	url := fmt.Sprintf("http://127.0.0.1:%d", p)
	s.Run()
	return s, p, url
}

func runRequest(endpoint string) (string, error) {
	s, _, url := newServer()
	res, err := http.Get(fmt.Sprintf("%s/%s", url, endpoint))
	if err != nil {
		s.Stop()
		return "", err
	}

	byteValue, err := ioutil.ReadAll(res.Body)
	if err != nil {
		s.Stop()
		return "", err
	}
	res.Body.Close()

	result := string(byteValue)
	s.Stop()
	return result, nil
}

func TestPing(t *testing.T) {
	result, err := runRequest("ping")
	if err != nil {
		t.Fatal("Unexpected error:", err.Error())
	}

	if result != "PONG" {
		t.Fatal("Unexpected result:", result)
	}
}
