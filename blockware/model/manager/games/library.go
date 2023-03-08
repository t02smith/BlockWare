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

// Library stores user's game information on owned and downloading games
type Library struct {

	// a user's owned games (includes downloads)
	ownedGames map[[32]byte]*Game

	// games present on the blockchain
	blockchainGames map[[32]byte]*Game

	/**
	Download manager threads will send requests down this channel
	to prompt a peer listener to attempt to download the block
	*/
	RequestDownload chan DownloadRequest

	/**
	Once a block has been downloaded a message will be sent down
	this channel.
	This can be used to signal to the UI that a download has been
	completed
	*/
	DownloadProgress chan DownloadRequest
}

// NewLibrary create a new library
func NewLibrary() *Library {
	util.Logger.Info("Creating new library")
	return &Library{
		ownedGames:       make(map[[32]byte]*Game),
		blockchainGames:  make(map[[32]byte]*Game),
		RequestDownload:  make(chan DownloadRequest),
		DownloadProgress: make(chan DownloadRequest),
	}
}

// CreateDownload create a download for a given game
func (l *Library) CreateDownload(g *Game) error {
	util.Logger.Infof("Creating download for %s:%x", g.Title, g.RootHash)
	if _, ok := l.ownedGames[g.RootHash]; !ok {
		return errors.New("game not found in library, cannot add download")
	}

	err := g.SetupDownload()
	if err != nil {
		return err
	}

	util.Logger.Infof("Download created for %s:%x", g.Title, g.RootHash)
	g.Download.ContinueDownload(g.RootHash, l.RequestDownload)
	return nil
}

// GetOwnedGame get a game and its download if they exist
func (l *Library) GetOwnedGame(rootHash [32]byte) *Game {
	if _, ok := l.ownedGames[rootHash]; !ok {
		return nil
	}

	return l.ownedGames[rootHash]
}

// GetOwnedGames get all games stored in the library
func (l *Library) GetOwnedGames() []*Game {
	gs := []*Game{}
	for _, g := range l.ownedGames {
		gs = append(gs, g)
	}
	return gs
}

//

// AddOwnedGame add a game to the library
func (l *Library) AddOwnedGame(g *Game) error {
	util.Logger.Debug(g)
	l.ownedGames[g.RootHash] = g
	return nil
}

// OutputGamesTable output a table representation of the games list to the console
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

// FindBlock find a given block within a game in a player's library
func (l *Library) FindBlock(gameHash [32]byte, hash [32]byte) (bool, []byte, error) {
	g, ok := l.ownedGames[gameHash]
	if !ok {
		return false, nil, nil
	}

	return g.FetchShard(hash)
}

// SetBlockchainGame store a details from a game on the blockchain store
func (l *Library) SetBlockchainGame(rootHash [32]byte, game *Game) {
	l.blockchainGames[rootHash] = game
}

// GetBlockchainGame get a game and its download if they exist
func (l *Library) GetBlockchainGame(rootHash [32]byte) *Game {
	if _, ok := l.blockchainGames[rootHash]; !ok {
		return nil
	}

	return l.blockchainGames[rootHash]
}

// GetBlockchainGames get all games stored in the library
func (l *Library) GetBlockchainGames() []*Game {
	gs := []*Game{}
	for _, g := range l.blockchainGames {
		gs = append(gs, g)
	}
	return gs
}

// GetDownloads get the games being downloaded
func (l *Library) GetDownloads() map[[32]byte]*Download {
	ds := make(map[[32]byte]*Download)

	for hash, g := range l.ownedGames {
		if g.Download != nil {
			ds[hash] = g.Download
		}
	}

	return ds
}

// Close close down a current library instance
func (l *Library) Close() {
	close(l.DownloadProgress)
	l.StopDownloads()

	metaDir := viper.GetString("meta.directory")
	for _, g := range l.ownedGames {
		err := g.OutputToFile(filepath.Join(metaDir, "games", fmt.Sprintf("%x", g.RootHash)))
		if err != nil {
			util.Logger.Warnf("Error outputting game %x to file: %s", g.RootHash, err)
		}
	}
}

// StopDownloads stop downloads from making requests
func (l *Library) StopDownloads() {
	util.Logger.Info("Stopping download requests")
	close(l.RequestDownload)
	l.RequestDownload = nil
}

// ContinueDownloads continue a libraries downloads
func (l *Library) ContinueDownloads() {
	util.Logger.Info("Continuing downloads")
	l.RequestDownload = make(chan DownloadRequest)

	count := 0
	for _, g := range l.ownedGames {
		if g.Download.Finished() {
			continue
		}

		g.Download.ContinueDownload(g.RootHash, l.RequestDownload)
		count++
	}
	util.Logger.Infof("Started %d downloads", count)
}

// ClearOwnedGames clear stored owned games from memory
func (l *Library) ClearOwnedGames() {
	util.Logger.Info("Flushing owned games from memory")
	l.ownedGames = make(map[[32]byte]*Game)
}
