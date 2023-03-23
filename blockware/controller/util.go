package controller

import (
	"encoding/hex"
	"fmt"
	"time"

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
	if d == nil || d.Finished() {
		return nil
	}

	diff := time.Since(d.StartTime)
	x := &ControllerDownload{
		TotalBlocks: d.TotalBlocks,
		Progress:    make(map[string]ControllerFileProgress),
		Name:        name,
		Stage:       string(d.Stage),
		ElapsedTime: fmt.Sprintf("%d:%d:%d", int(diff.Hours()), int(diff.Minutes()), int(diff.Seconds())),
	}

	lock := d.GetProgressLock()
	lock.Lock()
	for fHash, f := range d.Progress {
		fProgress := &ControllerFileProgress{
			AbsolutePath:    f.AbsolutePath,
			BlocksRemaining: []string{},
		}

		if len(f.BlocksRemaining) > 0 {
			for b := range f.BlocksRemaining {
				fProgress.BlocksRemaining = append(fProgress.BlocksRemaining, fmt.Sprintf("%x", b))
			}
		}

		x.Progress[fmt.Sprintf("%x", fHash)] = *fProgress
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
