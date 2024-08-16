package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := NewInMemoryPlayerStore()
	srv := NewPlayerServer(store)
	name := "pepper"

	srv.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(name))
	srv.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(name))
	srv.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(name))

	t.Run("get score", func(t *testing.T) {
		resp := httptest.NewRecorder()
		srv.ServeHTTP(resp, newGetScoreRequest(name))
		assertResponseCode(t, resp.Code, http.StatusOK)
		assertResponseBody(t, resp.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		resp := httptest.NewRecorder()
		srv.ServeHTTP(resp, newGetLeagueRequest())
		assertResponseCode(t, resp.Code, http.StatusOK)
		got := getLeagueFromResponse(t, resp.Body)
		want := []Player{
			{"pepper", 3},
		}
		assertLeague(t, got, want)
	})
}
