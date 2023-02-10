package games

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/viper"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

/*

A library represents an abstraction over the difference between
a game and a download. A user will have a collection of downloaded
and partially downloaded games and can in theory share shards from
both of them.

*/

// stores user's game information on owned and downloading games
type Library struct {

	// a user's owned Games
	Games map[[32]byte]*Game
}

// create a new library
func NewLibrary() *Library {
	return &Library{
		Games: make(map[[32]byte]*Game),
	}
}

func (l *Library) CreateDownload(g *Game) error {
	util.Logger.Infof("Creating download for %s:%x", g.Title, g.RootHash)
	if _, ok := l.Games[g.RootHash]; !ok {
		return errors.New("game not found in library, cannot add download")
	}

	err := g.setupDownload()
	if err != nil {
		return err
	}

	g.download.Serialise(filepath.Join(viper.GetString("games.installFolder"), fmt.Sprintf("%x", g.RootHash)))
	util.Logger.Infof("Download created for %s:%x", g.Title, g.RootHash)
	return nil
}

// get a game and its download if they exist
func (l *Library) GetGame(rootHash [32]byte) *Game {
	if _, ok := l.Games[rootHash]; !ok {
		return nil
	}

	return l.Games[rootHash]
}

func (l *Library) GetGames() []*Game {
	gs := []*Game{}
	for _, g := range l.Games {
		gs = append(gs, g)
	}
	return gs
}

//

// add a game to the library
func (l *Library) AddGame(g *Game) {
	l.Games[g.RootHash] = g
}

func (l *Library) OutputGamesTable() {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Title", "Version", "Release"})

	counter := 1
	for _, g := range l.Games {
		t.AppendRow(table.Row{fmt.Sprint(counter), g.Title, g.Version, g.ReleaseDate})
		counter++
	}

	t.Render()
}

func (l *Library) FindBlock(gameHash [32]byte, hash [32]byte) (bool, []byte, error) {
	g, ok := l.Games[gameHash]
	if !ok {
		return false, nil, nil
	}

	return g.FetchShard(hash)
}
