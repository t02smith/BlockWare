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
	gs := net.Peer().GetLibrary().GetOwnedGames()
	out := []*ControllerGame{}

	for _, g := range gs {
		out = append(out, &ControllerGame{
			Title:           g.Title,
			Version:         g.Version,
			ReleaseDate:     g.ReleaseDate,
			Developer:       g.Developer,
			RootHash:        fmt.Sprintf("%x", g.RootHash),
			PreviousVersion: fmt.Sprintf("%x", g.PreviousVersion),
			IPFSId:          g.IPFSId,
			Price:           g.Price,
			Uploader:        g.Uploader,
			Download:        downloadToAppDownload(g.Download),
		})
	}

	return out
}

// get a list of games from the eth store
func (a *Controller) GetStoreGames() []*ControllerGame {
	lib := net.Peer().GetLibrary()
	err := ethereum.FillLibraryBlockchainGames(lib)
	if err != nil {
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
			IPFSId:          g.IPFSId,
			Price:           g.Price,
			Uploader:        g.Uploader,
			Download:        downloadToAppDownload(g.Download),
		})
	}

	return out
}

// upload a new game
func (a *Controller) UploadGame(title, version, dev, rootDir string, shardSize, price, workerCount uint) string {
	release := time.Now().String()
	progress := make(chan int)

	go func() {
		fileCount, current := <-progress, 0
		runtime.EventsEmit(a.ctx, "file-count", fileCount)

		for current < fileCount {
			<-progress
			current++
			runtime.EventsEmit(a.ctx, "file-progress", current)
		}
	}()

	viper.Set("meta.hashes.workerCount", workerCount)
	g, err := games.CreateGame(title, version, release, dev, rootDir, big.NewInt(int64(price)), shardSize, progress)
	if err != nil {
		util.Logger.Errorf("Error creating game %s", err)
		return err.Error()
	}
	close(progress)

	err = ethereum.Upload(g)
	if err != nil {
		util.Logger.Errorf("Error uploading game %s", err)
		return err.Error()
	}

	net.Peer().GetLibrary().AddOwnedGame(g)
	err = games.OutputAllGameDataToFile(g)
	if err != nil {
		util.Logger.Errorf("Error saving game to file %s", err)
		return err.Error()
	}

	return ""

}
