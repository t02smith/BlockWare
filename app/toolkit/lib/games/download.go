package games

import (
	"errors"
	"log"

	"github.com/spf13/viper"
)

type Download struct {

	// The root hash of the game being downloaded
	GameRootHash [32]byte

	// progress of each file being downloaded
	Progress map[[32]byte]*FileProgress
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

	// generate dummy files

	dir := viper.GetString("games.installFolder")
	if len(dir) == 0 {
		return nil, errors.New("game install folder not found")
	}

	log.Printf("Generating dummy files for %s-%s", game.Title, game.Version)
	err := game.data.CreateDummyFiles(dir, game.Title)
	if err != nil {
		return nil, err
	}

	//

	return nil, nil
}
