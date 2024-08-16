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

func (st *InMemoryPlayerStore) GetLeague() []Player {
	return nil
}
