package games

import (
	"fmt"
	"sync"

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

	//
	DownloadManager *DownloadManager

	lock sync.Mutex
}

// NewLibrary create a new library
func NewLibrary() *Library {
	util.Logger.Info("Creating new library")
	return &Library{
		ownedGames:      make(map[[32]byte]*Game),
		blockchainGames: make(map[[32]byte]*Game),
		DownloadManager: NewDownloadManager(),
	}
}

// CreateDownload create a download for a given game
func (l *Library) CreateDownload(g *Game) error {
	util.Logger.Infof("Creating download for %s:%x", g.Title, g.RootHash)
	if _, ok := l.ownedGames[g.RootHash]; !ok {
		return fmt.Errorf("game %x not found in library, cannot add download", g.RootHash)
	}

	if g.Download != nil {
		return fmt.Errorf("download already started for game %x", g.RootHash)
	}

	err := g.SetupDownload()
	if err != nil {
		return err
	}

	util.Logger.Infof("Download created for %s:%x", g.Title, g.RootHash)
	g.OutputToFile()
	g.Download.ContinueDownload(g.RootHash, l.DownloadManager.RequestDownload)
	return nil
}

// GetOwnedGame get a game and its download if they exist
func (l *Library) GetOwnedGame(rootHash [32]byte) *Game {
	l.lock.Lock()
	defer l.lock.Unlock()

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

// AddOwnedGame add a game to the library
func (l *Library) AddOrUpdateOwnedGame(g *Game) {
	l.ownedGames[g.RootHash] = g
}

// FindAndRetrieveBlock find a given block within a game in a player's library
func (l *Library) FindAndRetrieveBlock(gameHash [32]byte, blockHash [32]byte) (bool, []byte, error) {
	g, ok := l.ownedGames[gameHash]
	if !ok {
		return false, nil, nil
	}

	return g.FetchShard(blockHash)
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

	l.lock.Lock()
	for hash, g := range l.ownedGames {
		if g.Download != nil {
			ds[hash] = g.Download
		}
	}
	l.lock.Unlock()

	return ds
}

// Close close down a current library instance
func (l *Library) Close() {
	l.StopDownloads()

	l.lock.Lock()
	for _, g := range l.ownedGames {
		err := g.OutputToFile()
		if err != nil {
			util.Logger.Warnf("Error outputting game %x to file: %s", g.RootHash, err)
		}
	}
	l.lock.Unlock()
}

// StopDownloads stop downloads from making requests
func (l *Library) StopDownloads() {
	util.Logger.Info("Stopping download requests")
	l.DownloadManager.Close()
}

// ContinueDownloads continue a libraries downloads
func (l *Library) ContinueDownloads() {
	util.Logger.Info("Continuing downloads")
	l.DownloadManager.RequestDownload = make(chan DownloadRequest)

	count := 0
	for _, g := range l.ownedGames {
		if g.Download == nil || g.Download.Finished() {
			continue
		}

		g.Download.ContinueDownload(g.RootHash, l.DownloadManager.RequestDownload)
		count++
	}
	util.Logger.Infof("Started %d downloads", count)
}

// ClearOwnedGames clear stored owned games from memory
func (l *Library) ClearOwnedGames() {
	util.Logger.Info("Flushing owned games from memory")
	l.ownedGames = make(map[[32]byte]*Game)
}

func (l *Library) Uninstall(gameHash [32]byte) error {
	game := l.GetOwnedGame(gameHash)
	if game == nil {
		return fmt.Errorf("game %x not owned", gameHash)
	}
	util.Logger.Infof("Uninstalling game %s", game.Title)

	err := game.CancelDownload()
	if err != nil {
		return err
	}

	util.Logger.Infof("Game %s uninstalled", game.Title)
	return nil
}

func (l *Library) Lock() {
	l.lock.Lock()
}

func (l *Library) Unlock() {
	l.lock.Unlock()
}
