package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
	"github.com/t02smith/part-iii-project/toolkit/model/ethereum"
	"github.com/t02smith/part-iii-project/toolkit/model/games"
	"github.com/t02smith/part-iii-project/toolkit/model/net"
	"github.com/t02smith/part-iii-project/toolkit/util"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// utility types

type AppGame struct {

	// game metadata
	Title           string `json:"title"`
	Version         string `json:"version"`
	ReleaseDate     string `json:"release"`
	Developer       string `json:"dev"`
	RootHash        string `json:"rootHash"`
	PreviousVersion string `json:"previousVersion"`

	// blockchain related
	IPFSId   string         `json:"IPFSId"`
	Price    *big.Int       `json:"price"`
	Uploader common.Address `json:"uploader"`

	Download *AppDownload `json:"download"`
}

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
	util.Logger.Info("Starting app context")
	a.ctx = ctx
}

// ? Interface functions

// get a list of owned games
func (a *App) GetOwnedGames() []*AppGame {
	gs := net.GetPeerInstance().GetLibrary().GetOwnedGames()
	out := []*AppGame{}

	for _, g := range gs {
		out = append(out, &AppGame{
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
func (a *App) GetAllGames() {
	ethereum.ReadPreviousGameEvents()
}

// get a list of downloads
func (a *App) GetDownloads() map[string]*AppDownload {
	ds := net.GetPeerInstance().GetLibrary().GetDownloads()
	return downloadToGameDownloads(ds)
}

func (a *App) IsDownloading(gh string) int {
	gh_tmp, err := hex.DecodeString(gh)
	if err != nil {
		util.Logger.Warnf("Error decoding hash string %s", err)
		return -1
	}

	gameHash := [32]byte{}
	copy(gameHash[:], gh_tmp[:])

	lib := net.GetPeerInstance().GetLibrary()
	g := lib.GetOwnedGame(gameHash)

	// ? game exists
	if g == nil {
		util.Logger.Warnf("Game %s doesn't exist", gh)
		return -1
	}

	// ? download not started
	if g.GetDownload() == nil {
		util.Logger.Info("download for game %s not started", gh)
		return 0
	}

	// ? download finished
	if g.GetDownload().Finished() {
		util.Logger.Info("download for game %s finished", gh)

		return 1
	}

	// ? download in progress
	util.Logger.Info("download for game %s in progress", gh)
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

func (a *App) CreateDownload(gh string) bool {
	util.Logger.Infof("Initiated download for %s", gh)

	gh_tmp, err := hex.DecodeString(gh)
	if err != nil {
		util.Logger.Error("Error creating download %s", err)
		return false
	}
	gameHash := [32]byte{}
	copy(gameHash[:], gh_tmp[:])

	lib := net.GetPeerInstance().GetLibrary()
	g := lib.GetOwnedGame(gameHash)

	// ? game exists
	if g == nil {
		util.Logger.Warnf("game %s doesn't exist", gh)
		return false
	}

	// ? game already has download
	if g.GetDownload() != nil {
		util.Logger.Warnf("Download for %s already exists", gh)
		return false
	}

	err = lib.CreateDownload(g)
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

	err = games.OutputAllGameDataToFile(g)
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
