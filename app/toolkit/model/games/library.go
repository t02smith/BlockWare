package games

import (
	"errors"
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/t02smith/part-iii-project/toolkit/model"
)

/*

A library represents an abstraction over the difference between
a game and a download. A user will have a collection of downloaded
and partially downloaded games and can in theory share shards from
both of them.

*/

// stores user's game information on owned and downloading games
type Library struct {

	// a user's owned games
	games map[[32]byte]*Game
}

// create a new library
func NewLibrary() *Library {
	return &Library{
		games: make(map[[32]byte]*Game),
	}
}

func (l *Library) createDownload(g *Game) error {
	model.Logger.Infof("Creating download for %s:%x", g.Title, g.RootHash)
	if _, ok := l.games[g.RootHash]; !ok {
		return errors.New("game not found in library, cannot add download")
	}

	err := setupDownload(g)
	if err != nil {
		return err
	}

	model.Logger.Infof("Download created for %s:%x", g.Title, g.RootHash)
	return nil
}

// get a game and its download if they exist
func (l *Library) GetGame(rootHash [32]byte) *Game {
	if _, ok := l.games[rootHash]; !ok {
		return nil
	}

	return l.games[rootHash]
}

func (l *Library) GetGames() []*Game {
	gs := []*Game{}
	for _, g := range l.games {
		gs = append(gs, g)
	}
	return gs
}

//

// add a game to the library
func (l *Library) AddGame(g *Game) {
	l.games[g.RootHash] = g
}

func (l *Library) OutputGamesTable() {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Title", "Version", "Release"})

	counter := 1
	for _, g := range l.games {
		t.AppendRow(table.Row{fmt.Sprint(counter), g.Title, g.Version, g.ReleaseDate})
		counter++
	}

	t.Render()
}
