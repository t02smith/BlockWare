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

	// a user's owned games (includes downloads)
	ownedGames map[[32]byte]*Game

	// games present on the blockchain
	blockchainGames map[[32]byte]*Game

	// used to send messages about download progress
	DownloadProgress chan *struct {
		GameHash  [32]byte
		BlockHash [32]byte
	}
}

// create a new library
func NewLibrary() *Library {
	return &Library{
		ownedGames:      make(map[[32]byte]*Game),
		blockchainGames: make(map[[32]byte]*Game),
		DownloadProgress: make(chan *struct {
			GameHash  [32]byte
			BlockHash [32]byte
		}),
	}
}

// create a download for a given game
func (l *Library) CreateDownload(g *Game) error {
	util.Logger.Infof("Creating download for %s:%x", g.Title, g.RootHash)
	if _, ok := l.ownedGames[g.RootHash]; !ok {
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
func (l *Library) GetOwnedGame(rootHash [32]byte) *Game {
	if _, ok := l.ownedGames[rootHash]; !ok {
		return nil
	}

	return l.ownedGames[rootHash]
}

// get all games stored in the library
func (l *Library) GetOwnedGames() []*Game {
	gs := []*Game{}
	for _, g := range l.ownedGames {
		gs = append(gs, g)
	}
	return gs
}

//

// add a game to the library
func (l *Library) AddOwnedGame(g *Game) error {
	l.ownedGames[g.RootHash] = g

	// fetch download if it exists
	d, err := DeserializeDownload(g.RootHash)
	if err != nil {
		return err
	}

	g.download = d
	return nil
}

// output a table representation of the games list to the console
func (l *Library) OutputGamesTable() {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Title", "Version", "Release"})

	counter := 1
	for _, g := range l.ownedGames {
		t.AppendRow(table.Row{fmt.Sprint(counter), g.Title, g.Version, g.ReleaseDate})
		counter++
	}

	t.Render()
}

func (l *Library) FindBlock(gameHash [32]byte, hash [32]byte) (bool, []byte, error) {
	g, ok := l.ownedGames[gameHash]
	if !ok {
		return false, nil, nil
	}

	return g.FetchShard(hash)
}

//

func (l *Library) SetBlockchainGame(rootHash [32]byte, game *Game) {
	l.blockchainGames[rootHash] = game
}

// get a game and its download if they exist
func (l *Library) GetBlockchainGame(rootHash [32]byte) *Game {
	if _, ok := l.blockchainGames[rootHash]; !ok {
		return nil
	}

	return l.blockchainGames[rootHash]
}

// get all games stored in the library
func (l *Library) GetBlockchainGames() []*Game {
	gs := []*Game{}
	for _, g := range l.blockchainGames {
		gs = append(gs, g)
	}
	return gs
}

// get the games being downloaded
func (l *Library) GetDownloads() map[[32]byte]*Download {
	ds := make(map[[32]byte]*Download)

	for hash, g := range l.ownedGames {
		if g.download != nil {
			ds[hash] = g.download
		}
	}

	return ds
}

func (l *Library) Close() {
	close(l.DownloadProgress)

	// serialise downloads
	for _, g := range l.ownedGames {
		if g.download == nil {
			continue
		}

		err := g.download.Serialise(fmt.Sprintf("%x", g.RootHash))
		if err != nil {
			util.Logger.Error(err)
		}
	}
}
