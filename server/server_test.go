package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/marcelog/minesweeper-API/user"
)

var port = 8000

func assertErrorMessage(t *testing.T, result string, expected string) {
	m, _ := json.Marshal(map[string]string{"message": expected})
	if result != string(m) {
		t.Fatal("Unexpected result:", result)
	}
}

func authHeader(u *user.User) map[string]string {
	return map[string]string{"X-API-Key": u.APIKey}
}

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

func runGet(t *testing.T, baseURL string, endpoint string, headers map[string]string) (string, *http.Response, error) {
	return runRequest(t, baseURL, "GET", endpoint, headers, "")
}

func runPost(t *testing.T, baseURL string, endpoint string, headers map[string]string, body string) (string, *http.Response, error) {
	return runRequest(t, baseURL, "POST", endpoint, headers, body)
}

func runRequest(t *testing.T, baseURL string, method string, endpoint string, headers map[string]string, body string) (string, *http.Response, error) {
	var res *http.Response
	var err error

	reqURL := fmt.Sprintf("%s/%s", baseURL, endpoint)

	client := &http.Client{}
	req, err := http.NewRequest(method, reqURL, strings.NewReader(body))

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	if method == "post" {
		req.Header.Add("Content-Type", "application/json")
	}

	res, err = client.Do(req)
	if err != nil {
		t.Fatal("Unexpected error:", err.Error())
	}

	byteValue, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", res, err
	}
	res.Body.Close()

	result := string(byteValue)
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

func TestAuthErrorIfMissingAuthHeader(t *testing.T) {
	s, _, url := newServer(t)

	_, res, _ := runPost(t, url, "games", map[string]string{}, "{}")
	if res.StatusCode != 401 {
		t.Fatal("Unexpected status code:", res.StatusCode)
	}

	s.Stop()
}

func TestAuthErrorIfUserNotfound(t *testing.T) {
	s, _, url := newServer(t)

	_, res, _ := runPost(t, url, "games", map[string]string{"X-API-Key": "whatever"}, "{}")
	if res.StatusCode != 401 {
		t.Fatal("Unexpected status code:", res.StatusCode)
	}

	s.Stop()
}

func TestAuth(t *testing.T) {
	s, _, url := newServer(t)

	u := s.State.AddUser()

	_, res, _ := runPost(t, url, "games", authHeader(u), `{"width": 8, "height": 8, "mines": 1}`)

	if res.StatusCode != 201 {
		t.Fatal("Unexpected status code:", res.StatusCode)
	}

	s.Stop()
}
