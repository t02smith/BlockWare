package controller

import (
	"encoding/hex"
	"fmt"

	"github.com/t02smith/part-iii-project/toolkit/model/net"
	"github.com/t02smith/part-iii-project/toolkit/util"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// get a list of downloads
func (a *Controller) GetDownloads() map[string]*ControllerDownload {
	ds := net.Peer().GetLibrary().GetDownloads()
	return downloadToGameDownloads(ds)
}

// is the given game downloading
func (a *Controller) IsDownloading(gh string) int {
	gh_tmp, err := hex.DecodeString(gh)
	if err != nil {
		util.Logger.Warnf("Error decoding hash string %s", err)
		return -1
	}

	gameHash := [32]byte{}
	copy(gameHash[:], gh_tmp[:])

	lib := net.Peer().GetLibrary()
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
func (a *Controller) StartDownloadListener() {
	go func() {
		downloadChannel := net.Peer().GetLibrary().DownloadProgress
		for progress := range downloadChannel {
			util.Logger.Infof("Download event received %x-%x", progress.GameHash, progress.BlockHash)
			runtime.EventsEmit(a.ctx, fmt.Sprintf("%x", progress.GameHash), fmt.Sprintf("%x", progress.BlockHash))
		}
	}()
}

// create a new download for a given game
func (a *Controller) CreateDownload(gh string) bool {
	util.Logger.Infof("Initiated download for %s", gh)

	gh_tmp, err := hex.DecodeString(gh)
	if err != nil {
		util.Logger.Error("Error creating download %s", err)
		return false
	}
	gameHash := [32]byte{}
	copy(gameHash[:], gh_tmp[:])

	lib := net.Peer().GetLibrary()
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
