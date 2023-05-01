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
	lib := peer.Peer().Library()
	ds := lib.GetDownloads()

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

	if g.GetDownload().Paused {
		return 3
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

// start all in progress downloads
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

// start a download for a given game
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

	fmt.Println(g)
	if g.Download == nil {
		c.controllerErrorf("download not started for game %s", gh)
		return
	}

	g.Download.ContinueDownload(gameHash, lib.DownloadManager.RequestDownload)
	runtime.EventsEmit(c.ctx, "update-downloads")
	runtime.EventsEmit(c.ctx, "update-owned-games")
	fmt.Println(g)
}

// cancel an in progress download
func (c *Controller) CancelDownload(gh string) {
	gameHash, err := hashStringToByte32(gh)
	if err != nil {
		return
	}

	lib := peer.Peer().Library()
	g := lib.GetOwnedGame(gameHash)
	if g == nil {
		c.controllerErrorf("game %s doesn't exist", gh)
		return
	}

	if g.Download == nil {
		c.controllerErrorf("download not started for game %s", gh)
		return
	}

	err = lib.Uninstall(gameHash)
	if err != nil {
		c.controllerErrorf("Error uninstalling game %s", err)
		return
	}

	runtime.EventsEmit(c.ctx, "update-downloads")
	runtime.EventsEmit(c.ctx, "update-owned-games")
}

// pause an in progress download
func (c *Controller) PauseDownload(gh string) {
	gameHash, err := hashStringToByte32(gh)
	if err != nil {
		return
	}

	lib := peer.Peer().Library()
	g := lib.GetOwnedGame(gameHash)
	if g == nil {
		c.controllerErrorf("game %s doesn't exist", gh)
		return
	}

	if g.Download == nil {
		c.controllerErrorf("download not started for game %s", gh)
		return
	}

	g.Download.Paused = true
	runtime.EventsEmit(c.ctx, "update-downloads")
	runtime.EventsEmit(c.ctx, "update-owned-games")
}
