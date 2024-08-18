package player

import (
	"io"
)

type FileSystemStore struct {
	database io.ReadWriteSeeker
}

func NewFileSystemStore() *FileSystemStore {
	return &FileSystemStore{}
}

func (s *FileSystemStore) GetPlayerScore(name string) int {
	var wins int
	for _, player := range s.GetLeague() {
		if player.Name == name {
			wins = player.Wins
			break
		}
	}
	return wins
}

func (s *FileSystemStore) RecordWin(player string) {
}

func (s *FileSystemStore) GetLeague() []Player {
	s.database.Seek(0, io.SeekStart)
	league, _ := NewLeague(s.database)
	return league
}
