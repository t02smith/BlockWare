package games

import (
	"errors"
	"log"

	"github.com/spf13/viper"
	"github.com/t02smith/part-iii-project/toolkit/lib/hash"
)

type Download struct {

	// The root hash of the game being downloaded
	GameRootHash []byte

	// progress of each file being downloaded
	Progress map[[32]byte]FileProgress
}

type FileProgress struct {

	// where the file is located in storage
	AbsolutePath string

	// block hash => position in file
	BlocksRemaining map[[32]byte]int
}

// serialisation

// setup

func SetupDownload(game *Game) (*Download, error) {

	// read hash data if isn't already loaded
	if game.data == nil {
		err := game.ReadHashData()
		if err != nil {
			return nil, err
		}
	}

	// create download obj
	d := &Download{
		GameRootHash: game.RootHash,
		Progress:     make(map[[32]byte]FileProgress),
	}

	// generate dummy files
	dir := viper.GetString("games.installFolder")
	if len(dir) == 0 {
		return nil, errors.New("game install folder not found")
	}

	log.Printf("Generating dummy files for %s-%s", game.Title, game.Version)
	err := game.data.CreateDummyFiles(dir, game.Title, func(path string, htf *hash.HashTreeFile) {
		p := FileProgress{
			AbsolutePath:    path,
			BlocksRemaining: make(map[[32]byte]int),
		}

		for i, b := range htf.Hashes {
			p.BlocksRemaining[b] = i
		}

		d.Progress[htf.RootHash] = p
	})

	if err != nil {
		return nil, err
	}

	return d, nil
}
