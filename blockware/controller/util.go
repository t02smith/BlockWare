package controller

import (
	"encoding/hex"
	"fmt"
	"time"

	"github.com/spf13/viper"
	"github.com/t02smith/part-iii-project/toolkit/model/manager/games"
	"github.com/t02smith/part-iii-project/toolkit/model/net/peer"
	"github.com/t02smith/part-iii-project/toolkit/util"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

/*

Error messages
All error messages are emitted as events so that
the frontend can show them on screen for the user
to help debug

*/

// log and emit an error
func (c *Controller) controllerError(err string) {
	util.Logger.Error(err)
	runtime.EventsEmit(c.ctx, "error", err)
}

// Controller.controllerError fmt.sprintf wrapper
func (c *Controller) controllerErrorf(err string, args ...any) {
	c.controllerError(fmt.Sprintf(err, args...))
}

// open a dialog to select a file and return its absolute path
func (c *Controller) SelectTxtFile() string {
	res, err := runtime.OpenFileDialog(c.ctx, runtime.OpenDialogOptions{
		Filters: []runtime.FileFilter{{Pattern: "*.txt"}},
	})

	if err != nil {
		c.controllerError("Error opening folder")
		return ""
	}

	return res
}

// open a dialog to select a folder and returns its absolute path
func (c *Controller) SelectFolder() string {
	res, err := runtime.OpenDirectoryDialog(c.ctx, runtime.OpenDialogOptions{})

	if err != nil {
		c.controllerError("Error opening folder")
		return ""
	}

	return res
}

// get the toolkit directory
func (c *Controller) GetDirectory() string {
	return viper.GetString("meta.directory")
}

func (c *Controller) GetContractAddress() string {
	return viper.GetString("contract.address")
}

// misc

func downloadToGameDownloads(ds map[[32]byte]*games.Download) map[string]*ControllerDownload {
	out := make(map[string]*ControllerDownload)

	for hash, d := range ds {
		g := peer.Peer().Library().GetOwnedGame(hash)

		out[fmt.Sprintf("%x", hash)] = downloadToAppDownload(d, g.Title)
	}

	return out
}

func downloadToAppDownload(d *games.Download, name string) *ControllerDownload {
	if d == nil {
		return nil
	}

	finished := d.Finished()
	var diff time.Duration
	if finished && d.EndTime != nil {
		diff = d.EndTime.Sub(d.StartTime)
	} else {
		diff = time.Since(d.StartTime)
	}

	x := &ControllerDownload{
		TotalBlocks: d.TotalBlocks,
		Name:        name,
		Stage:       string(d.Stage),
		ElapsedTime: fmt.Sprintf("%02d:%02d:%02d", int(diff.Hours()), int(diff.Minutes())%60, int(diff.Seconds())%60),
		Finished:    finished,
		BlocksLeft:  0,
	}

	lock := d.GetProgressLock()
	lock.Lock()
	for _, f := range d.Progress {
		x.BlocksLeft += len(f.BlocksRemaining)
	}
	lock.Unlock()

	return x
}

// converts a hex string to a 32 byte array
func hashStringToByte32(hash string) ([32]byte, error) {
	gh_tmp, err := hex.DecodeString(hash)
	if err != nil {
		return [32]byte{}, err
	}
	gameHash := [32]byte{}
	copy(gameHash[:], gh_tmp[:])
	return gameHash, nil
}
