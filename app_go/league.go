package player

import (
	"encoding/json"
	"fmt"
	"io"
)

// NewLeague
func NewLeague(rdr io.Reader) ([]Player, error) {
	var league []Player
	err := json.NewDecoder(rdr).Decode(&league)

	if err != nil {
		err = fmt.Errorf("unable to parse reader: %q, err: '%v'", rdr, err)
	}
	return league, err
}
