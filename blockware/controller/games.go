package controller

import (
	"bytes"
	"fmt"
	"math/big"
	"sort"
	"strings"
	"time"

	"github.com/spf13/viper"
	"github.com/t02smith/part-iii-project/toolkit/model/manager/games"
	"github.com/t02smith/part-iii-project/toolkit/model/net/peer"
	"github.com/t02smith/part-iii-project/toolkit/model/persistence/ethereum"
	"github.com/t02smith/part-iii-project/toolkit/model/persistence/ethereum/library"
	"github.com/t02smith/part-iii-project/toolkit/util"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// get a list of owned games
func (a *Controller) GetOwnedGames() []*ControllerGame {
	gs := peer.Peer().Library().GetOwnedGames()
	addr := ethereum.Address()
	out := []*ControllerGame{}

	for _, g := range gs {
		var rd string
		_rd, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", strings.Split(g.ReleaseDate, " m=")[0])
		if err != nil {
			rd = g.ReleaseDate
		} else {
			rd = fmt.Sprintf("%d %s %d", _rd.Day(), _rd.Month(), _rd.Year())
		}

		out = append(out, &ControllerGame{
			Title:           g.Title,
			Version:         g.Version,
			ReleaseDate:     rd,
			Developer:       g.Developer,
			RootHash:        fmt.Sprintf("%x", g.RootHash),
			PreviousVersion: fmt.Sprintf("%x", g.PreviousVersion),
			IPFSId:          g.HashTreeIPFSAddress,
			Price:           g.Price,
			Uploader:        g.Uploader,
			Download:        downloadToAppDownload(g.Download, g.Title, g.Version),
			AssetsFolder:    g.Assets.AbsolutePath,
			IsOwner:         bytes.Equal(addr.Bytes(), g.Uploader.Bytes()),
		})
	}

	sort.Slice(out, func(i, j int) bool {
		return gs[i].ReleaseDate >= gs[j].ReleaseDate
	})

	sort.Slice(out, func(i, j int) bool {
		return out[i].Title < out[j].Title
	})

	return out
}

// get a list of games from the eth store
func (c *Controller) GetStoreGames() []*ControllerGame {
	lib := peer.Peer().Library()
	err := library.FillLibraryBlockchainGames(lib)
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
			Download:        downloadToAppDownload(g.Download, g.Title, g.Version),
			AssetsFolder:    g.Assets.AbsolutePath,
		})
	}

	return out
}

// upload a new game
func (c *Controller) UploadGame(title, version, dev, rootDir string, shardSize, price, workerCount uint, assetsDir, previousVersion string) {
	release := time.Now().String()
	progress := make(chan int)

	var prevVersion [32]byte

	if previousVersion != "" {
		prevVersionBytes, err := hashStringToByte32(previousVersion)
		if err != nil {
			c.controllerErrorf("Invalid previous version")
		}
		prevVersion = prevVersionBytes
	}

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
	g, err := games.CreateGame(games.NewGame{
		Title:           title,
		Version:         version,
		ReleaseDate:     release,
		Developer:       dev,
		RootDir:         rootDir,
		Price:           big.NewInt(int64(price)),
		ShardSize:       shardSize,
		AssetsDir:       assetsDir,
		PreviousVersion: prevVersion,
	}, progress)
	if err != nil {
		c.controllerErrorf("Error creating game %s", err)
		return
	}

	err = library.UploadToEthereum(g)
	if err != nil {
		c.controllerErrorf("Error uploading game %s", err)
		return
	}

	peer.Peer().Library().AddOrUpdateOwnedGame(g)
	err = games.OutputAllGameDataToFile(g)
	if err != nil {
		c.controllerErrorf("Error saving game to file %s", err)
		return
	}

	close(progress)
}

// find an uploaded game given its root hash
func (c *Controller) GetGameFromStoreByRootHash(rh string) *ControllerGame {
	gh, err := hashStringToByte32(rh)
	if err != nil {
		util.Logger.Errorf("Error parsing hash %s", err)
	}

	lib := peer.Peer().Library()
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
		Download:        downloadToAppDownload(g.Download, g.Title, g.Version),
		AssetsFolder:    g.Assets.AbsolutePath,
	}
}

// purchase a new game by its hash
func (c *Controller) PurchaseGame(rh string) {
	gh, err := hashStringToByte32(rh)
	if err != nil {
		c.controllerErrorf("Error parsing hash %s", err)
		return
	}

	lib := peer.Peer().Library()
	err = library.Purchase(lib, gh)
	if err != nil {
		c.controllerErrorf("Error purchasing game %s", err)
		return
	}

	runtime.EventsEmit(c.ctx, "update-owned-games")
}

// fetch an owned game from blockchain
func (c *Controller) FetchOwnedGame(gh string) {
	hash, err := hashStringToByte32(gh)
	if err != nil {
		c.controllerErrorf("Error fetching owned game %s", err)
		return
	}

	err = library.FetchOwnedGame(peer.Peer().Library(), hash)
	if err != nil {
		c.controllerErrorf("Error fetching owned game %s", err)
		return
	}

	runtime.EventsEmit(c.ctx, "update-owned-games")
}

// uninstall an owned game
func (c *Controller) UninstallGame(rh string) {
	gh, err := hashStringToByte32(rh)
	if err != nil {
		c.controllerErrorf("Error parsing hash %s", err)
		return
	}

	lib := peer.Peer().Library()
	err = lib.Uninstall(gh)
	if err != nil {
		c.controllerErrorf("Error uninstalling game %s", err)
		return
	}

	runtime.EventsEmit(c.ctx, "update-downloads")
}

// checks for updates to owned games
func (c *Controller) CheckForUpdates() {
	lib := peer.Peer().Library()
	if err := library.CheckForGameUpdates(lib); err != nil {
		c.controllerError("Error checking for updates")
		util.Logger.Error(err)
	}
}
