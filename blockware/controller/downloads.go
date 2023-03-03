package controller

import (
	"encoding/hex"
	"fmt"

	"github.com/t02smith/part-iii-project/toolkit/model/net/peer"
	"github.com/t02smith/part-iii-project/toolkit/util"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// get a list of downloads
func (a *Controller) GetDownloads() map[string]*ControllerDownload {
	ds := peer.Peer().Library().GetDownloads()
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

	lib := peer.Peer().Library()
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
	util.Logger.Infof("download for game %s in progress", gh)
	return 2
}

// listen for incoming download progress alerts
func (a *Controller) StartDownloadListener() {
	go func() {
		downloadChannel := peer.Peer().Library().DownloadProgress
		for progress := range downloadChannel {
			util.Logger.Infof("Download event received %x-%x", progress.GameHash, progress.BlockHash)
			runtime.EventsEmit(a.ctx, fmt.Sprintf("%x", progress.GameHash), fmt.Sprintf("%x", progress.BlockHash))
		}
	}()
}

// create a new download for a given game
func (c *Controller) CreateDownload(gh string) {
	util.Logger.Infof("Initiated download for %s", gh)

	gh_tmp, err := hex.DecodeString(gh)
	if err != nil {
		c.controllerErrorf("Error creating download %s", err)
		return
	}
	gameHash := [32]byte{}
	copy(gameHash[:], gh_tmp[:])

	lib := peer.Peer().Library()
	g := lib.GetOwnedGame(gameHash)

	// ? game exists
	if g == nil {
		c.controllerErrorf("game %s doesn't exist", gh)
		return
	}

	// ? game already has download
	if g.GetDownload() != nil {
		c.controllerErrorf("Download for %s already exists", gh)
		return
	}

	err = lib.CreateDownload(g)
	if err != nil {
		c.controllerErrorf("Error creating download for game %x: %s", gameHash, err)
		return
	}

}
