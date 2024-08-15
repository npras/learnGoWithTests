package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := &InMemoryPlayerStore{scores: map[string]int{}}
	server := NewPlayerServer(store)
	name := "pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(name))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(name))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(name))

	resp := httptest.NewRecorder()
	server.ServeHTTP(resp, newGetScoreRequest(name))

	assertResponseCode(t, resp.Code, http.StatusOK)
	assertResponseBody(t, resp.Body.String(), "3")
}
