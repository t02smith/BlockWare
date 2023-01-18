package net

import (
	"strings"
	"testing"
	"time"

	"github.com/t02smith/part-iii-project/toolkit/lib/games"
)

func TestGameListToMessage(t *testing.T) {
	games := []*games.Game{
		{
			Title:       "Test Game",
			Version:     "1.0.2",
			ReleaseDate: time.Now().String(),
			Developer:   "tcs1g20",
			RootHash:    []byte("test"),
		},
		{
			Title:       "Borderlands 3",
			Version:     "3.4.17",
			ReleaseDate: time.Now().String(),
			Developer:   "Gearbox",
			RootHash:    []byte("tester hash"),
		},
	}

	res, err := gameListToMessage(games)
	if err != nil {
		t.Error(err)
		return
	}

	parts := strings.Split(res, ";")
	if parts[0] != "GAMES" {
		t.Error("Wrong command")
	}

	for i, g := range games {
		if serialised, err := g.Serialise(); err != nil || serialised != parts[i+1] {
			t.Errorf("Incorrect serialised game in pos %d", i)
		}
	}

}
