package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores map[string]int
}

func (ps *StubPlayerStore) GetPlayerScore(player string) int {
	score := ps.scores[player]
	return score
}

func TestGETPlayers(t *testing.T) {
	store := &StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Kyle":   27,
		},
	}
	srv := &PlayerServer{store: store}

	t.Run("returns  Pep's score", func(t *testing.T) {
		req := newGetScoreRequest("Pepper")
		respWriter := httptest.NewRecorder()

		srv.ServeHTTP(respWriter, req)

		assertResponseBody(t, respWriter.Body.String(), "20")
		assertResponseCode(t, respWriter.Code, http.StatusOK)
	})

	t.Run("returns  Kyle's score", func(t *testing.T) {
		req := newGetScoreRequest("Kyle")
		respWriter := httptest.NewRecorder()

		srv.ServeHTTP(respWriter, req)

		assertResponseBody(t, respWriter.Body.String(), "27")
		assertResponseCode(t, respWriter.Code, http.StatusOK)
	})

	t.Run("returns  404 on missing players", func(t *testing.T) {
		req := newGetScoreRequest("doesnotexist")
		respWriter := httptest.NewRecorder()

		srv.ServeHTTP(respWriter, req)

		assertResponseCode(t, respWriter.Code, http.StatusNotFound)
	})
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
