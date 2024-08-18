package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

// stubs and spys

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   []Player
}

func (st *StubPlayerStore) GetPlayerScore(player string) int {
	score := st.scores[player]
	return score
}

func (st *StubPlayerStore) RecordWin(player string) {
	st.winCalls = append(st.winCalls, player)
}

func (st *StubPlayerStore) GetLeague() []Player {
	return st.league
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
	t.Run("it returns the league table as JSON", func(t *testing.T) {
		wantedLeague := []Player{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}
		store := &StubPlayerStore{scores: nil, winCalls: nil, league: wantedLeague}
		srv := NewPlayerServer(store)

		req := newGetLeagueRequest()
		resp := httptest.NewRecorder()
		srv.ServeHTTP(resp, req)

		got := getLeagueFromResponse(t, resp.Body)
		fmt.Println("=============")
		fmt.Println(json.MarshalIndent(wantedLeague, "", "		"))
		assertResponseCode(t, resp.Code, http.StatusOK)
		assertLeague(t, got, wantedLeague)
		assertContentType(t, resp, jsonContentType)
	})
}

// helpers

func newGetLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}

func newPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/players/"+name, nil)
	return req
}

func getLeagueFromResponse(t testing.TB, body io.Reader) []Player {
	t.Helper()
	var got []Player
	err := json.NewDecoder(body).Decode(&got)

	if err != nil {
		t.Fatalf("unable to parse resp_body: %q, err: '%v'", body, err)
	}
	return got
}

// test helpers

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

func assertLeague(t testing.TB, got, want []Player) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertContentType(t testing.TB, resp *httptest.ResponseRecorder, want string) {
	t.Helper()
	respHeaders := resp.Result().Header
	if respHeaders.Get("content-type") != "application/json" {
		t.Errorf("response did not have content-type of application/json, got %v", respHeaders)
	}
}
