package main

type InMemoryPlayerStore struct {
	scores map[string]int
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{scores: map[string]int{}}
}

func (s *InMemoryPlayerStore) GetPlayerScore(player string) int {
	score := s.scores[player]
	return score
}

func (s *InMemoryPlayerStore) RecordWin(player string) {
	s.scores[player] += 1
}

func (st *InMemoryPlayerStore) GetLeague() League {
	var players League
	for k, v := range st.scores {
		player := Player{Name: k, Wins: v}
		players = append(players, player)
	}
	return players
}
