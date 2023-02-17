package main

import (
	"context"
	"math/big"
	"time"

	"github.com/spf13/viper"
	"github.com/t02smith/part-iii-project/toolkit/model/ethereum"
	"github.com/t02smith/part-iii-project/toolkit/model/games"
	"github.com/t02smith/part-iii-project/toolkit/model/net"
	"github.com/t02smith/part-iii-project/toolkit/util"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

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

func (a *App) GetOwnedGames() []*games.Game {
	return net.GetPeerInstance().GetLibrary().GetOwnedGames()
}

func (a *App) GetAllGames() {
	ethereum.ReadPreviousGameEvents()
}

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

func (a *App) DeployLibraryInstance(privateKey string) string {
	_, _, err := ethereum.DeployLibraryContract(privateKey)
	if err == nil {
		return ""
	}

	return err.Error()
}
