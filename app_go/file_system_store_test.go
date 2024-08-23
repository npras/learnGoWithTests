package main

import (
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {

	t.Run("#GetLeague_is_sorted", func(t *testing.T) {
		dataFile, removeFileFn := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 32}]`)
		defer removeFileFn()
		store, _ := NewFileSystemStore(dataFile)
		got := store.GetLeague()
		want := []Player{
			{"Chris", 32},
			{"Cleo", 10},
		}
		assertLeague(t, got, want)
		// re-reading to test the seeker
		got = store.GetLeague()
		assertLeague(t, got, want)
	})

	t.Run("#GetPlayerScore", func(t *testing.T) {
		dataFile, removeFileFn := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 32}]`)
		defer removeFileFn()
		store, _ := NewFileSystemStore(dataFile)
		got := store.GetPlayerScore("Chris")
		assertInt(t, got, 32)
	})

	t.Run("#RecordWin_existing_player", func(t *testing.T) {
		dataFile, removeFileFn := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 30}]`)
		defer removeFileFn()
		store, _ := NewFileSystemStore(dataFile)
		store.RecordWin("Chris")
		store.RecordWin("Chris")
		got := store.GetPlayerScore("Chris")
		assertInt(t, got, 32)
	})

	t.Run("#RecordWin_new_player", func(t *testing.T) {
		dataFile, removeFileFn := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 30}]`)
		defer removeFileFn()
		store, _ := NewFileSystemStore(dataFile)
		store.RecordWin("Leo")
		assertInt(t, store.GetPlayerScore("Leo"), 1)
		want := []Player{
			{"Chris", 30},
			{"Cleo", 10},
			{"Leo", 1},
		}
		assertLeague(t, store.GetLeague(), want)
	})

}

// assertions

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didn't expect an error but got one: %v", err)
	}
}

// helpers
func createTempFile(t testing.TB, initialData string) (*os.File, func()) {
	t.Helper()
	tempfile, err := os.CreateTemp("", "db")
	if err != nil {
		t.Fatalf("couldn't create tmp file. error: %v", err)
	}
	tempfile.Write([]byte(initialData))
	removeFileFn := func() {
		tempfile.Close()
		os.Remove(tempfile.Name())
	}
	return tempfile, removeFileFn
}
