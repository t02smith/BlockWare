package controller

import (
	"fmt"

	"github.com/t02smith/part-iii-project/toolkit/model/games"
)

func downloadToGameDownloads(ds map[[32]byte]*games.Download) map[string]*ControllerDownload {
	out := make(map[string]*ControllerDownload)

	for hash, d := range ds {
		out[fmt.Sprintf("%x", hash)] = downloadToAppDownload(d)
	}

	return out
}

func downloadToAppDownload(d *games.Download) *ControllerDownload {
	if d == nil {
		return nil
	}

	x := &ControllerDownload{
		TotalBlocks: d.TotalBlocks,
		Progress:    make(map[string]ControllerFileProgress),
	}

	for fHash, f := range d.Progress {
		fProgress := &ControllerFileProgress{
			AbsolutePath:    f.AbsolutePath,
			BlocksRemaining: []string{},
		}

		for b := range f.BlocksRemaining {
			fProgress.BlocksRemaining = append(fProgress.BlocksRemaining, fmt.Sprintf("%x", b))
		}

		x.Progress[fmt.Sprintf("%x", fHash)] = *fProgress
	}

	return x
}
