package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const jsonContentType = "application/json"

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() []Player
}

type Player struct {
	Name string
	Wins int
}

//

type PlayerServer struct {
	store PlayerStore
	http.Handler
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	ps := new(PlayerServer)
	ps.store = store
	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(ps.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(ps.playersHandler))
	ps.Handler = router
	return ps
}

func (ps *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonContentType)
	json.NewEncoder(w).Encode(ps.store.GetLeague())
}

func (ps *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	switch r.Method {
	case http.MethodGet:
		ps.showScore(w, player)
	case http.MethodPost:
		ps.processWin(w, player)
	}
}

func (ps *PlayerServer) showScore(w http.ResponseWriter, name string) {
	score := ps.store.GetPlayerScore(name)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Fprint(w, score)
}

func (ps *PlayerServer) processWin(w http.ResponseWriter, name string) {
	ps.store.RecordWin(name)
	w.WriteHeader(http.StatusAccepted)
}
