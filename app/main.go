package main

import (
	"log"
	"net/http"
)

type InMemoryStore struct {
	scores map[string]int
}

func (ps *InMemoryStore) GetPlayerScore(player string) int {
	score := ps.scores[player]
	return score
}

func main() {
	store := &InMemoryStore{
		map[string]int{
			"Pepper": 20,
			"Kyle":   27,
		},
	}
	srv := &PlayerServer{store: store}
	log.Fatal(http.ListenAndServe(":5000", srv))
}
