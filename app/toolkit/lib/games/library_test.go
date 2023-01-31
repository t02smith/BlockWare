package games

import (
	"testing"

	"github.com/t02smith/part-iii-project/toolkit/test/testutil"
)

func TestGetGame(t *testing.T) {
	testutil.ClearTmp("../../")
	testutil.SetupTmp("../../")

	// setup
	l := NewLibrary()
	d, g, err := setupTestDownload()
	if err != nil {
		t.Errorf("Error generating test download %s", err)
		return
	}

	// no game or download
	g1, d1 := l.GetGame(g.RootHash)
	if g1 != nil || d1 != nil {
		t.Error("expected nil pointers for missing game & download")
	}

	// game but no download
	l.AddGame(g)
	g1, d1 = l.GetGame(g.RootHash)
	if g1 == nil || d1 != nil {
		t.Error("Incorrect GetGame after adding single game")
	}

	// game and download
	l.AddDownload(d)
	g1, d1 = l.GetGame(g.RootHash)
	if g1 == nil || d1 == nil {
		t.Error("Game or download missing from response")
	}
}

func TestAddGame(t *testing.T) {
	testutil.ClearTmp("../../")
	testutil.SetupTmp("../../")

	l := NewLibrary()
	_, g, err := setupTestDownload()
	if err != nil {
		t.Errorf("Error generating test download %s", err)
		return
	}

	l.AddGame(g)
	if _, ok := l.games[g.RootHash]; !ok {
		t.Errorf("Game not added")
	}

}

func TestAddDownload(t *testing.T) {
	testutil.ClearTmp("../../")
	testutil.SetupTmp("../../")

	l := NewLibrary()
	d, g, err := setupTestDownload()
	if err != nil {
		t.Errorf("Error generating test download %s", err)
		return
	}

	// add a game not in the games map
	err = l.AddDownload(d)
	if err == nil {
		t.Errorf("game was accepted despite not being in the library")
	}

	// add a game in the games map
	l.AddGame(g)
	err = l.AddDownload(d)
	if err != nil {
		t.Errorf("error adding game to downloads %s", err)
	}

	if _, ok := l.downloads[g.RootHash]; !ok {
		t.Errorf("Game not added to downloads")
	}
}
