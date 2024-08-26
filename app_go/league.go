package main

import (
	"encoding/json"
	"fmt"
	"io"
	"sort"
)

type League []Player

func (l League) Sort() League {
	sort.Slice(l, func(i, j int) bool {
		return l[i].Wins > l[j].Wins
	})
	return l
}

func (l League) Find(name string) (*Player, int) {
	for i, player := range l {
		if player.Name == name {
			return &l[i], i
		}
	}
	return nil, 0
}

// NewLeague
func NewLeague(rdr io.Reader) (League, error) {
	var league League
	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		err = fmt.Errorf("unable to parse reader: %q, err: '%v'", rdr, err)
	}
	return league, err
}
