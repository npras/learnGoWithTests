package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// stubs and spys

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (ps *StubPlayerStore) GetPlayerScore(player string) int {
	score := ps.scores[player]
	return score
}

func (ps *StubPlayerStore) RecordWin(player string) {
	ps.winCalls = append(ps.winCalls, player)
}

// test cases

func TestGETPlayers(t *testing.T) {
	store := &StubPlayerStore{
		scores: map[string]int{
			"Pepper": 20,
			"Kyle":   27,
		},
	}
	srv := NewPlayerServer(store)

	t.Run("returns  Pep's score", func(t *testing.T) {
		req := newGetScoreRequest("Pepper")
		respWriter := httptest.NewRecorder()

		srv.ServeHTTP(respWriter, req)

		assertResponseCode(t, respWriter.Code, http.StatusOK)
		assertResponseBody(t, respWriter.Body.String(), "20")
	})

	t.Run("returns  Kyle's score", func(t *testing.T) {
		req := newGetScoreRequest("Kyle")
		respWriter := httptest.NewRecorder()

		srv.ServeHTTP(respWriter, req)

		assertResponseCode(t, respWriter.Code, http.StatusOK)
		assertResponseBody(t, respWriter.Body.String(), "27")
	})

	t.Run("returns  404 on missing players", func(t *testing.T) {
		req := newGetScoreRequest("doesnotexist")
		respWriter := httptest.NewRecorder()

		srv.ServeHTTP(respWriter, req)

		assertResponseCode(t, respWriter.Code, http.StatusNotFound)
		assertResponseBody(t, respWriter.Body.String(), "")
	})
}

func TestStoreWins(t *testing.T) {
	store := &StubPlayerStore{
		map[string]int{},
		nil,
	}
	srv := NewPlayerServer(store)

	t.Run("records wins when POSTed", func(t *testing.T) {
		name := "aplayer"
		req := newPostWinRequest(name)
		respWriter := httptest.NewRecorder()

		srv.ServeHTTP(respWriter, req)

		assertResponseCode(t, respWriter.Code, http.StatusAccepted)

		if got := len(store.winCalls); got != 1 {
			t.Fatalf("got: %d, want: %d", got, 1)
		}

		if got := store.winCalls[0]; got != name {
			t.Errorf("got: %q, want: %q", got, name)
		}
	})
}

func TestLeague(t *testing.T) {
	store := &StubPlayerStore{}
	srv := NewPlayerServer(store)
	t.Run("displays league info", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/league", nil)
		resp := httptest.NewRecorder()
		srv.ServeHTTP(resp, req)
		assertResponseCode(t, resp.Code, http.StatusOK)
	})
}

// helpers

func newPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/players/"+name, nil)
	return req
}

func assertResponseCode(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got: %d, want: %d", got, want)
	}
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got: %q, want: %q", got, want)
	}
}
