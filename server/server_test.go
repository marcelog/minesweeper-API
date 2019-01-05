package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

var port = 8000

func newServer(t *testing.T) (*Server, int, string) {
	p := port
	port++

	s := New(&Args{
		Address: "127.0.0.1",
		Port:    p,
	})
	url := fmt.Sprintf("http://127.0.0.1:%d", p)
	err := s.Run()
	if err != nil {
		t.Fatal("Unexpected error:", err.Error())
	}

	return s, p, url
}

func runRequest(t *testing.T, endpoint string) (string, *http.Response, error) {
	s, _, url := newServer(t)
	res, err := http.Get(fmt.Sprintf("%s/%s", url, endpoint))
	if err != nil {
		t.Fatal("Unexpected error:", err.Error())
	}

	byteValue, err := ioutil.ReadAll(res.Body)
	if err != nil {
		s.Stop()
		return "", res, err
	}
	res.Body.Close()

	result := string(byteValue)
	s.Stop()
	return result, res, nil
}

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
	s, _, url := newServer(t)

	res, err := http.Get(fmt.Sprintf(url))
	if err != nil {
		t.Fatal("Unexpected error:", err.Error())
	}
	res.Body.Close()

	if res.StatusCode != 404 {
		t.Fatal("Unexpected status code:", res.StatusCode)
	}

	s.Stop()
	_, err = http.Get(url)
	if err == nil {
		t.Fatal("Expected an error")
	}
}
