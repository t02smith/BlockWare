package games

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"
	"time"

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

type DownloadRequest struct {
	GameHash  [32]byte
	BlockHash [32]byte
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

// continue/start a download where the desired blocks are already known
// this function will make requests down the libraries download request
// channel
func (d *Download) ContinueDownload(gameHash [32]byte, newRequest chan DownloadRequest) {

	// TODO
	/**
	Currently, this function will go through files sequentially and request
	individual blocks at a time and will block until that block is received

	Just for simplicity :)
	*/

	util.Logger.Infof("Continuing download for game %x", gameHash)
	go func() {
		if d.Finished() {
			return
		}

		for _, file := range d.Progress {
			util.Logger.Infof("Requesting file %s for game %x", file.AbsolutePath, gameHash)

			shards := [][32]byte{}
			for sh := range file.BlocksRemaining {
				shards = append(shards, sh)
			}

			for _, shard := range shards {
				util.Logger.Infof("Requesting shard %x for file %s in game %x", shard, file.AbsolutePath, gameHash)
				newRequest <- DownloadRequest{
					GameHash:  gameHash,
					BlockHash: shard,
				}

				for {
					if _, ok := d.Progress[shard]; ok {
						time.Sleep(time.Second)
					} else {
						util.Logger.Infof("Shard %x downloaded for game %x", shard, gameHash)
						break
					}
				}
			}

		}
	}()

	if d.Finished() {
		util.Logger.Infof("Download complete for game %x", gameHash)
	} else {
		util.Logger.Infof("Download paused for game %x", gameHash)
	}
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
