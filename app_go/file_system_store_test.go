package player

import (
	"strings"
	"testing"
)

func TestFileSystemStore(t *testing.T) {

	t.Run("#GetLeague", func(t *testing.T) {
		database := strings.NewReader(`[{"Name": "Cleo", "Wins": 10},{"Name": "Chris", "Wins": 32}]`)
		store := FileSystemStore{database: database}
		got := store.GetLeague()
		want := []Player{
			{"Cleo", 10},
			{"Chris", 32},
		}
		assertLeague(t, got, want)
		// re-reading to test the seeker
		got = store.GetLeague()
		assertLeague(t, got, want)
	})

	t.Run("#GetPlayerScore", func(t *testing.T) {
		database := strings.NewReader(`[{"Name": "Cleo", "Wins": 10},{"Name": "Chris", "Wins": 32}]`)
		store := FileSystemStore{database: database}
		got := store.GetPlayerScore("Chris")
		assertInt(t, got, 32)
	})

	t.Run("#RecordWin", func(t *testing.T) {
		database := strings.NewReader(`[{"Name": "Cleo", "Wins": 10},{"Name": "Chris", "Wins": 30}]`)
		store := FileSystemStore{database: database}
		store.RecordWin("Chris")
		store.RecordWin("Chris")
		store.RecordWin("Chris")
		got := store.GetPlayerScore("Chris")
		assertInt(t, got, 33)
	})

}
