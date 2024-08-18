package player

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const jsonContentType = "application/json"

// PlayerStore stores score information about players
type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() []Player
}

// Player stores a name with a number of wins
type Player struct {
	Name string
	Wins int
}

//

// PlayerServer is a HTTP interface for player information
type PlayerServer struct {
	store PlayerStore
	http.Handler
}

// NewPlayerServer creates a PlayerServer with routing configured
func NewPlayerServer(store PlayerStore) *PlayerServer {
	ps := new(PlayerServer)
	ps.store = store
	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(ps.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(ps.playersHandler))
	ps.Handler = router
	return ps
}

// League Handler
func (ps *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonContentType)
	json.NewEncoder(w).Encode(ps.store.GetLeague())
}

// playersHandler
func (ps *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	switch r.Method {
	case http.MethodGet:
		ps.showScore(w, player)
	case http.MethodPost:
		ps.processWin(w, player)
	}
}

// showScore
func (ps *PlayerServer) showScore(w http.ResponseWriter, name string) {
	score := ps.store.GetPlayerScore(name)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Fprint(w, score)
}

// processWin
func (ps *PlayerServer) processWin(w http.ResponseWriter, name string) {
	ps.store.RecordWin(name)
	w.WriteHeader(http.StatusAccepted)
}
