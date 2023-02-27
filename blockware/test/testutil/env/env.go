package testenv

import (
	"math/big"
	"path/filepath"
	"testing"
	"time"

	"github.com/t02smith/part-iii-project/toolkit/model/games"
)

// creates a test game and stores the data in the tmp folder
func CreateTestGame(t *testing.T, toRoot string) (*games.Game, error) {
	t.Helper()

	datetime := time.Date(2002, 1, 10, 0, 0, 0, 0, time.UTC).String()
	game, err := games.CreateGame("toolkit", "1.0.4", datetime, "google.com", filepath.Join(toRoot, "./test/data/testdir"), big.NewInt(0), 256, nil)
	if err != nil {
		return nil, err
	}

	return game, nil
}
