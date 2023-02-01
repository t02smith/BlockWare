package games

import (
	"testing"

	"github.com/t02smith/part-iii-project/toolkit/test/testutil"
)

func TestGetGame(t *testing.T) {
	l := NewLibrary()
	g, err := setupTestDownload()
	if err != nil {
		t.Fatalf("Error generating test download %s", err)
	}

	t.Run("game not found", func(t *testing.T) {
		g1 := l.GetGame(g.RootHash)
		if g1 != nil {
			t.Error("expected nil pointers for missing game & download")
		}
	})

	t.Run("success", func(t *testing.T) {
		l.AddGame(g)
		g1 := l.GetGame(g.RootHash)
		if g1 == nil {
			t.Error("Incorrect GetGame after adding single game")
		}
	})

	testutil.ClearTmp("../../")
	testutil.SetupTmp("../../")

}

func TestAddGame(t *testing.T) {
	l := NewLibrary()
	g, err := setupTestDownload()
	if err != nil {
		t.Fatalf("Error generating test download %s", err)
		return
	}

	t.Run("success", func(t *testing.T) {
		l.AddGame(g)
		if _, ok := l.games[g.RootHash]; !ok {
			t.Errorf("Game not added")
		}
	})

	testutil.ClearTmp("../../")
	testutil.SetupTmp("../../")

}
