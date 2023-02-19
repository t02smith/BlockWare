package games

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"

	"github.com/spf13/viper"
	"github.com/t02smith/part-iii-project/toolkit/model/hash"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

type Download struct {

	// progress of each file being downloaded
	Progress map[[32]byte]FileProgress

	// total number of blocks to install
	TotalBlocks int
}

type FileProgress struct {

	// where the file is located in storage
	AbsolutePath string

	// block hash => position in file
	BlocksRemaining map[[32]byte]uint
}

// ? setup

// create a new download for an existing game
func (game *Game) setupDownload() error {

	// read hash data if isn't already loaded
	if game.data == nil {
		err := game.ReadHashData()
		if err != nil {
			return err
		}
	}

	// create download obj
	d := &Download{
		Progress:    make(map[[32]byte]FileProgress),
		TotalBlocks: 0,
	}

	// generate dummy files
	dir := viper.GetString("games.installFolder")
	if len(dir) == 0 {
		return errors.New("game install folder not found")
	}

	util.Logger.Infof("Generating dummy files for %s-%s", game.Title, game.Version)
	err := game.data.CreateDummyFiles(dir, game.Title, func(path string, htf *hash.HashTreeFile) {
		p := FileProgress{
			AbsolutePath:    path,
			BlocksRemaining: make(map[[32]byte]uint),
		}

		for i, b := range htf.Hashes {
			p.BlocksRemaining[b] = uint(i)
		}

		d.Progress[htf.RootHash] = p
	})

	if err != nil {
		return err
	}

	d.TotalBlocks = len(d.Progress)
	game.Download = d
	return nil
}

// ? downloading data

// insert data into a file for a given game download
func (d *Download) InsertData(fileHash, blockHash [32]byte, data []byte) error {
	util.Logger.Infof("Attempting to insert shard %x into %x with data %x", blockHash, fileHash, data)
	file, ok := d.Progress[fileHash]
	if !ok {
		return fmt.Errorf("file %x not in download queue", fileHash)
	}

	offset, ok := file.BlocksRemaining[blockHash]
	if !ok {
		return fmt.Errorf("block %x in File %x not in download queue", blockHash, fileHash)
	}

	dataHash := sha256.Sum256(data)
	if !bytes.Equal(blockHash[:], dataHash[:]) {
		return fmt.Errorf("block %x data for File %x given does not match expected content", blockHash, fileHash)
	}

	err := hash.InsertData(file.AbsolutePath, uint(len(data)), uint(offset), data)
	if err != nil {
		return err
	}

	delete(file.BlocksRemaining, blockHash)
	util.Logger.Infof("successfully inserted shard %x into %x with data %x", blockHash, fileHash, data)
	return nil
}

// start up an existing download
func (d *Download) ContinueDownload() error {

	// which peers have the desired game

	return nil
}

// find a new block from current peers
func (d *Download) FindBlockFromPeers(hash [32]byte) error {

	// which peers have the game

	// which peers have that block

	return nil
}

// ? misc functions

// get the total download progress as a percent
func (d *Download) GetProgress() float32 {
	return 1 - float32(len(d.Progress))/float32(d.TotalBlocks)
}

// whether any more blocks are still needed
func (d *Download) Finished() bool {
	return len(d.Progress) == 0
}
