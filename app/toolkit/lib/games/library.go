package games

import (
	"errors"

	"github.com/t02smith/part-iii-project/toolkit/lib"
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

	// downloads is a subset of games
	// some subset of a users games will be current downloads
	downloads map[[32]byte]*Download
}

// create a new library
func NewLibrary() *Library {
	return &Library{
		games:     make(map[[32]byte]*Game),
		downloads: make(map[[32]byte]*Download),
	}
}

func (l *Library) createDownload(g *Game) error {
	lib.Logger.Infof("Creating download for %s:%x", g.Title, g.RootHash)
	if _, ok := l.games[g.RootHash]; !ok {
		return errors.New("game not found in library, cannot add download")
	}

	if _, ok := l.downloads[g.RootHash]; !ok {
		return nil
	}

	d, err := setupDownload(g)
	if err != nil {
		return err
	}

	l.AddDownload(d)
	lib.Logger.Infof("Download created for %s:%x", g.Title, g.RootHash)
	return nil
}

func (l *Library) addGame(g *Game, download bool) error {
	return nil
}

// get a game and its download if they exist
func (l *Library) GetGame(rootHash [32]byte) (*Game, *Download) {
	if _, ok := l.games[rootHash]; !ok {
		return nil, nil
	}

	if _, ok := l.downloads[rootHash]; !ok {
		return l.games[rootHash], nil
	}

	return l.games[rootHash], l.downloads[rootHash]
}

func (l *Library) IsDownload(g *Game) bool {
	_, ok := l.downloads[g.RootHash]
	return ok
}

//

// add a game to the library
func (l *Library) AddGame(g *Game) {
	l.games[g.RootHash] = g
}

// add a new download
func (l *Library) AddDownload(d *Download) error {
	if _, ok := l.games[d.GameRootHash]; !ok {
		return errors.New("the given game is not in your library and can't be downloaded")
	}

	l.downloads[d.GameRootHash] = d
	return nil
}
