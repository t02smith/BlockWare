package controller

import (
	"fmt"
	"math/big"
	"time"

	"github.com/spf13/viper"
	"github.com/t02smith/part-iii-project/toolkit/model/ethereum"
	"github.com/t02smith/part-iii-project/toolkit/model/games"
	"github.com/t02smith/part-iii-project/toolkit/model/net"
	"github.com/t02smith/part-iii-project/toolkit/util"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// get a list of owned games
func (a *Controller) GetOwnedGames() []*ControllerGame {
	gs := net.Peer().Library().GetOwnedGames()
	out := []*ControllerGame{}

	for _, g := range gs {
		out = append(out, &ControllerGame{
			Title:           g.Title,
			Version:         g.Version,
			ReleaseDate:     g.ReleaseDate,
			Developer:       g.Developer,
			RootHash:        fmt.Sprintf("%x", g.RootHash),
			PreviousVersion: fmt.Sprintf("%x", g.PreviousVersion),
			IPFSId:          g.HashTreeIPFSAddress,
			Price:           g.Price,
			Uploader:        g.Uploader,
			Download:        downloadToAppDownload(g.Download),
		})
	}

	return out
}

// get a list of games from the eth store
func (c *Controller) GetStoreGames() []*ControllerGame {
	lib := net.Peer().Library()
	err := ethereum.FillLibraryBlockchainGames(lib)
	if err != nil {
		c.controllerErrorf("Error getting games from ETH: %s", err)
		return []*ControllerGame{}
	}

	gs := lib.GetBlockchainGames()
	out := []*ControllerGame{}

	for _, g := range gs {
		out = append(out, &ControllerGame{
			Title:           g.Title,
			Version:         g.Version,
			ReleaseDate:     g.ReleaseDate,
			Developer:       g.Developer,
			RootHash:        fmt.Sprintf("%x", g.RootHash),
			PreviousVersion: fmt.Sprintf("%x", g.PreviousVersion),
			IPFSId:          g.HashTreeIPFSAddress,
			Price:           g.Price,
			Uploader:        g.Uploader,
			Download:        downloadToAppDownload(g.Download),
		})
	}

	return out
}

// upload a new game
func (c *Controller) UploadGame(title, version, dev, rootDir string, shardSize, price, workerCount uint) {
	release := time.Now().String()
	progress := make(chan int)

	go func() {
		fileCount, current := <-progress, 0
		runtime.EventsEmit(c.ctx, "file-count", fileCount)

		for current < fileCount {
			<-progress
			current++
			runtime.EventsEmit(c.ctx, "file-progress", current)
		}
	}()

	viper.Set("meta.hashes.workerCount", workerCount)
	g, err := games.CreateGame(title, version, release, dev, rootDir, big.NewInt(int64(price)), shardSize, progress)
	if err != nil {
		c.controllerErrorf("Error creating game %s", err)
		return
	}
	close(progress)

	err = ethereum.Upload(g)
	if err != nil {
		c.controllerErrorf("Error uploading game %s", err)
		return
	}

	net.Peer().Library().AddOwnedGame(g)
	err = games.OutputAllGameDataToFile(g)
	if err != nil {
		c.controllerErrorf("Error saving game to file %s", err)
		return
	}
}

// find an uploaded game given its root hash
func (c *Controller) GetGameFromStoreByRootHash(rh string) *ControllerGame {
	gh, err := hashStringToByte32(rh)
	if err != nil {
		util.Logger.Errorf("Error parsing hash %s", err)
	}

	lib := net.Peer().Library()
	g := lib.GetBlockchainGame(gh)
	if g == nil {
		c.controllerErrorf("Game %s not found", rh)
		return nil
	}

	return &ControllerGame{
		Title:           g.Title,
		Version:         g.Version,
		ReleaseDate:     g.ReleaseDate,
		Developer:       g.Developer,
		RootHash:        fmt.Sprintf("%x", g.RootHash),
		PreviousVersion: fmt.Sprintf("%x", g.PreviousVersion),
		IPFSId:          g.HashTreeIPFSAddress,
		Price:           g.Price,
		Uploader:        g.Uploader,
		Download:        downloadToAppDownload(g.Download),
	}
}

func (c *Controller) PurchaseGame(rh string) {
	gh, err := hashStringToByte32(rh)
	if err != nil {
		c.controllerErrorf("Error parsing hash %s", err)
		return
	}

	lib := net.Peer().Library()
	err = ethereum.Purchase(lib, gh)
	if err != nil {
		c.controllerErrorf("Error purchasing game %s", err)
		return
	}

	runtime.EventsEmit(c.ctx, "update-owned-games")
}
