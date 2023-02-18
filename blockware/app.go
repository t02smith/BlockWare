package main

import (
	"context"
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

// utility types

type AppDownload struct {
	Progress    map[string]AppFileProgress
	TotalBlocks int
}

type AppFileProgress struct {
	AbsolutePath    string
	BlocksRemaining []string
}

// ? App setup

type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// ? Interface functions

// get a list of owned games
func (a *App) GetOwnedGames() []*games.Game {
	return net.GetPeerInstance().GetLibrary().GetOwnedGames()
}

// get a list of games from the eth store
func (a *App) GetAllGames() {
	ethereum.ReadPreviousGameEvents()
}

// get a list of downloads
func (a *App) GetDownloads() map[string]*AppDownload {
	ds := net.GetPeerInstance().GetLibrary().GetDownloads()
	out := make(map[string]*AppDownload)

	for hash, d := range ds {
		x := &AppDownload{
			TotalBlocks: d.TotalBlocks,
			Progress:    make(map[string]AppFileProgress),
		}

		for fHash, f := range d.Progress {
			fProgress := &AppFileProgress{
				AbsolutePath:    f.AbsolutePath,
				BlocksRemaining: []string{},
			}

			for b := range f.BlocksRemaining {
				fProgress.BlocksRemaining = append(fProgress.BlocksRemaining, fmt.Sprintf("%x", b))
			}

			x.Progress[fmt.Sprintf("%x", fHash)] = *fProgress
		}

		out[fmt.Sprintf("%x", hash)] = x
	}

	return out
}

func (a *App) IsDownloading(gameHash [32]byte) int {
	lib := net.GetPeerInstance().GetLibrary()
	g := lib.GetOwnedGame(gameHash)

	// ? game exists
	if g == nil {
		return -1
	}

	// ? download not started
	if g.GetDownload() == nil {
		return 0
	}

	// ? download finished
	if g.GetDownload().Finished() {
		return 1
	}

	// ? download in progress
	return 2
}

// listen for incoming download progress alerts
func (a *App) StartDownloadListener() {
	go func() {
		downloadChannel := net.GetPeerInstance().GetLibrary().DownloadProgress
		for progress := range downloadChannel {
			util.Logger.Infof("Download event received %x-%x", progress.GameHash, progress.BlockHash)
			runtime.EventsEmit(a.ctx, fmt.Sprintf("%x", progress.GameHash), fmt.Sprintf("%x", progress.BlockHash))
		}
	}()
}

func (a *App) CreateDownload(gameHash [32]byte) bool {
	util.Logger.Infof("Initiated download for %x", gameHash)
	lib := net.GetPeerInstance().GetLibrary()
	g := lib.GetOwnedGame(gameHash)

	// ? game exists
	if g == nil {
		return false
	}

	// ? game doesn't already have download
	if g.GetDownload() != nil {
		return false
	}

	err := lib.CreateDownload(g)
	if err != nil {
		util.Logger.Errorf("Error creating download for game %x: %s", gameHash, err)
		return false
	}

	return true
}

// upload a new game
func (a *App) UploadGame(title, version, dev, rootDir string, shardSize, price, workerCount uint) string {
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

	err = games.OutputToFile(g)
	if err != nil {
		util.Logger.Errorf("Error saving game to file %s", err)
		return err.Error()
	}

	return ""

}

// deploy a new instance of the library contract
func (a *App) DeployLibraryInstance(privateKey string) string {
	_, _, err := ethereum.DeployLibraryContract(privateKey)
	if err == nil {
		return ""
	}

	return err.Error()
}
