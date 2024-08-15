package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

//

type PlayerServer struct {
	store  PlayerStore
	router *http.ServeMux
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	ps := &PlayerServer{
		store:  store,
		router: http.NewServeMux(),
	}

	ps.router.Handle("/league", http.HandlerFunc(ps.leagueHandler))
	ps.router.Handle("/players/", http.HandlerFunc(ps.playersHandler))

	return ps
}

func (ps *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ps.router.ServeHTTP(w, r)
}

func (ps *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
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
