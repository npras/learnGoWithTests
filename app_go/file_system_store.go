package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type FileSystemStore struct {
	league League
	db     *json.Encoder
}

// funcs

func NewFileSystemStore(f *os.File) (*FileSystemStore, error) {
	err := initialiseFile(f)
	if err != nil {
		return nil, fmt.Errorf("problem initialising player db file, %v", err)
	}

	league, err := NewLeague(f)
	if err != nil {
		return nil, fmt.Errorf("problem loading PlayerStore from file %s, %v", f.Name(), err)
	}

	s := new(FileSystemStore)
	s.league = league
	s.db = json.NewEncoder(&tape{f})
	return s, nil
}

// methods

func (s *FileSystemStore) GetPlayerScore(name string) int {
	player := s.league.Find(name)
	if player != nil {
		return player.Wins
	}
	return 0
}

func (s *FileSystemStore) RecordWin(name string) {
	player := s.league.Find(name)
	if player == nil {
		s.league = append(s.league, Player{name, 1})
	} else {
		player.Wins++
	}
	s.db.Encode(s.league)
}

func (s *FileSystemStore) GetLeague() League {
	return s.league.Sort()
}

// helpers

func initialiseFile(f *os.File) error {
	stat, err := f.Stat()
	if err != nil {
		return fmt.Errorf("problem getting file info from file %s, %v", f.Name(), err)
	}
	if stat.Size() == 0 {
		f.WriteString("[]")
	}
	f.Seek(0, io.SeekStart)
	return nil
}
