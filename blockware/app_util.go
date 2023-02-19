package main

import (
	"fmt"

	"github.com/t02smith/part-iii-project/toolkit/model/games"
)

func downloadToGameDownloads(ds map[[32]byte]*games.Download) map[string]*AppDownload {
	out := make(map[string]*AppDownload)

	for hash, d := range ds {
		out[fmt.Sprintf("%x", hash)] = downloadToAppDownload(d)
	}

	return out
}

func downloadToAppDownload(d *games.Download) *AppDownload {
	if d == nil {
		return nil
	}

	x := &AppDownload{
		TotalBlocks: d.TotalBlocks,
		Progress:    make(map[string]AppFileProgress),
	}

	for fHash, f := range d.Progress {
		fProgress := &AppFileProgress{
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
