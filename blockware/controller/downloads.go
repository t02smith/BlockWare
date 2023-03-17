package controller

import (
	"encoding/hex"

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
		return 0
	}

	// ? download finished
	if g.GetDownload().Finished() {
		return 1
	}

	// ? download in progress
	return 2
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

	runtime.EventsEmit(c.ctx, "update-downloads")
}

func (c *Controller) ContinueAllDownloads() {
	util.Logger.Infof("continuing all downloads")

	lib := peer.Peer().Library()
	for _, g := range lib.GetOwnedGames() {
		if g.Download == nil || g.Download.Finished() {
			continue
		}

		g.Download.ContinueDownload(g.RootHash, lib.DownloadManager.RequestDownload)
	}

	runtime.EventsEmit(c.ctx, "update-downloads")
}

func (c *Controller) ContinueDownload(gh string) {
	util.Logger.Infof("continuing download for %s", gh)

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

	if g.Download == nil {
		c.controllerErrorf("download not started for game %s", gh)
		return
	}

	g.Download.ContinueDownload(gameHash, lib.DownloadManager.RequestDownload)
	runtime.EventsEmit(c.ctx, "update-downloads")
}
