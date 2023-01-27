package games

import (
	"bufio"
	"encoding/gob"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

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

func DeserializeDownload(gameHash []byte) (*Download, error) {
	log.Printf("Deserialized download %x", gameHash)
	dir := viper.GetString("games.tracker.directory")
	if len(dir) == 0 {
		return nil, errors.New("tracker directory not found")
	}

	f, err := os.Open(filepath.Join(dir, fmt.Sprintf("%x", gameHash)))
	if err != nil {
		return nil, err
	}
	defer f.Close()

	decoder := gob.NewDecoder(bufio.NewReader(f))
	d := &Download{}

	err = decoder.Decode(d)
	if err != nil && err != io.EOF {
		return nil, err
	}

	return d, nil
}

func (d *Download) Serialise() error {
	log.Printf("Serialising download %x", d.GameRootHash)
	dir := viper.GetString("games.tracker.directory")
	if len(dir) == 0 {
		return errors.New("tracker directory not found")
	}

	f, err := os.Create(filepath.Join(dir, fmt.Sprintf("%x", d.GameRootHash)))
	if err != nil {
		return err
	}
	defer f.Close()

	writer := bufio.NewWriter(f)
	encoder := gob.NewEncoder(writer)
	err = encoder.Encode(d)
	if err != nil {
		return err
	}

	writer.Flush()

	log.Printf("%x serialised successfully", d.GameRootHash)
	return nil
}

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

	err = d.Serialise()
	if err != nil {
		log.Printf("Error saving download to file: %s", err)
		return nil, err
	}

	return d, nil
}
