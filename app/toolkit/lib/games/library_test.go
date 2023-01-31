package games

import (
	"testing"

	"github.com/t02smith/part-iii-project/toolkit/test/testutil"
)

func TestGetGame(t *testing.T) {
	l := NewLibrary()
	d, g, err := setupTestDownload()
	if err != nil {
		t.Fatalf("Error generating test download %s", err)
	}

	t.Run("no game or download", func(t *testing.T) {
		g1, d1 := l.GetGame(g.RootHash)
		if g1 != nil || d1 != nil {
			t.Error("expected nil pointers for missing game & download")
		}
	})

	t.Run("game but no download", func(t *testing.T) {
		l.AddGame(g)
		g1, d1 := l.GetGame(g.RootHash)
		if g1 == nil || d1 != nil {
			t.Error("Incorrect GetGame after adding single game")
		}
	})

	t.Run("game and download", func(t *testing.T) {
		l.AddDownload(d)
		g1, d1 := l.GetGame(g.RootHash)
		if g1 == nil || d1 == nil {
			t.Error("Game or download missing from response")
		}
	})

	testutil.ClearTmp("../../")
	testutil.SetupTmp("../../")

}

func TestAddGame(t *testing.T) {
	l := NewLibrary()
	_, g, err := setupTestDownload()
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

func TestAddDownload(t *testing.T) {

	l := NewLibrary()
	d, g, err := setupTestDownload()
	if err != nil {
		t.Fatalf("Error generating test download %s", err)
	}

	t.Run("add a download not in the games list", func(t *testing.T) {
		err = l.AddDownload(d)
		if err == nil {
			t.Errorf("game was accepted despite not being in the library")
		}
	})

	t.Run("success", func(t *testing.T) {
		l.AddGame(g)
		err = l.AddDownload(d)
		if err != nil {
			t.Errorf("error adding game to downloads %s", err)
		}

		if _, ok := l.downloads[g.RootHash]; !ok {
			t.Errorf("Game not added to downloads")
		}
	})

	testutil.ClearTmp("../../")
	testutil.SetupTmp("../../")
}
