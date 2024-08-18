package player

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
