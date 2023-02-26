package games

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
	"github.com/t02smith/part-iii-project/toolkit/model/hash"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

/*

This file is responsible for choosing blocks from each download
and putting a request through to the downloader thread for them.

Each game may have a download struct as part of it that can be
in three possible states:

1. Download not started => value is nil
2. Download is started but not finished => Download.Progress has at least 1 value
3. Download is finished => Download.Progress is empty

Starting a download will:
1. Use the HashTree to generate a set of dummy files that match the names
   and directory structure of the game. By filling these files with empty
	 data we can test whether the user has sufficient storage.
	 TODO remove data if they don't
2. Each file will be allocated a FileProgress struct that will be used to
   track its blocks. This will alllow us to easily prioritise completing a
	 whole file first.
3. When donwloads are activated, it will send requests down a given channel
   that will be used to actually process the download
	 TODO limit how many requests can be sent at once i.e. wait for a batch
	 TODO   to be downloaded first

*/

// A download manager for a game
type Download struct {

	// progress of each file being downloaded
	Progress map[[32]byte]FileProgress

	// total number of blocks to install
	TotalBlocks int
}

// the progress of a specific file's download
type FileProgress struct {

	// where the file is located in storage
	AbsolutePath string

	// block hash => position in file
	BlocksRemaining map[[32]byte]uint
}

/*
Stores the details about a given block to attempt
to download it.
*/
type DownloadRequest struct {

	// details to uniquely identify the block
	GameHash  [32]byte
	BlockHash [32]byte

	/* how many attempts have already been made to find
	this block. After a certain number, this block will
	either be timed out or cancelled for this session
	*/
	Attempts uint8
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

	if len(file.BlocksRemaining) == 0 {
		err = CleanFile(file.AbsolutePath)
		if err != nil {
			util.Logger.Errorf("Error cleaning file %s: %s", file.AbsolutePath, err)
		}
	}

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
			util.Logger.Info("Download finished")
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
					Attempts:  0,
				}

			}

			time.Sleep(500 * time.Millisecond)

		}
	}()

	if d.Finished() {
		util.Logger.Infof("Download complete for game %x", gameHash)
	} else {
		util.Logger.Infof("Download paused for game %x", gameHash)
	}
}

// ? misc functions

/*
It is extremely likely that files won't be multiples of
the shard size so this will result in a trail of bytes
that are unexepcted.

This function will remove any trailing bytes from the
end of a file
*/
func CleanFile(path string) error {
	util.Logger.Infof("Cleaning file %s", path)
	file, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	bytesToRemove, buf := 0, make([]byte, 1)
	for {
		file.Seek(int64(bytesToRemove), 0)
		_, err := reader.Read(buf)
		if err != nil {
			return err
		}

		if buf[0] != 0x00 {
			bytesToRemove++
		} else {
			break
		}
	}

	if err = file.Truncate(int64(bytesToRemove)); err != nil {
		return err
	}
	util.Logger.Infof("Cleaned file %s", path)
	return nil
}

// get the total download progress as a percent
func (d *Download) GetProgress() float32 {
	return 1 - float32(len(d.Progress))/float32(d.TotalBlocks)
}

// whether any more blocks are still needed
func (d *Download) Finished() bool {
	return len(d.Progress) == 0
}
