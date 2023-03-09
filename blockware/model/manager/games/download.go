package games

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
	hash "github.com/t02smith/part-iii-project/toolkit/model/manager/hashtree"
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

// Download A download manager for a game
type Download struct {

	// progress of each file being downloaded
	Progress map[[32]byte]FileProgress

	// total number of blocks to install
	TotalBlocks int

	// absolute path of root directory
	AbsolutePath string
}

// FileProgress the progress of a specific file's download
type FileProgress struct {

	// where the file is located in storage
	AbsolutePath string

	// block hash => position in file
	BlocksRemaining map[[32]byte][]uint
}

// DownloadRequest /*
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

// SetupDownload create a new download for an existing game
func (g *Game) SetupDownload() error {

	// read hash data if isn't already loaded
	data, err := g.GetData()
	if err != nil {
		return err
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

	d.AbsolutePath = filepath.Join(dir, g.Title)

	util.Logger.Infof("Generating dummy files for %s-%s", g.Title, g.Version)
	err = data.CreateDummyFiles(dir, g.Title, func(path string, htf *hash.HashTreeFile) {
		p := FileProgress{
			AbsolutePath:    path,
			BlocksRemaining: make(map[[32]byte][]uint),
		}

		for i, b := range htf.Hashes {
			if offsets, ok := p.BlocksRemaining[b]; ok {
				offsets = append(offsets, uint(i))
			} else {
				p.BlocksRemaining[b] = []uint{uint(i)}
			}
		}

		d.TotalBlocks += len(p.BlocksRemaining)
		d.Progress[htf.RootHash] = p
	})

	if err != nil {
		return err
	}

	g.Download = d
	return nil
}

// CancelDownload cancel an existing download
func (g *Game) CancelDownload() error {
	if g.Download == nil {
		util.Logger.Warnf("Download not found for game %x", g.RootHash)
		return nil
	}

	util.Logger.Infof("Cancelling download for game %x", g.RootHash)

	// * remove dummy files
	util.Logger.Infof("Clearing dummy files from game %x", g.RootHash)

	err := os.RemoveAll(g.Download.AbsolutePath)
	if err != nil {
		return err
	}

	util.Logger.Infof("Download cancelled for game %x", g.RootHash)
	g.Download = nil
	return nil
}

// ? downloading data

// InsertData insert data into a file for a given game download
func (d *Download) InsertData(fileHash, blockHash [32]byte, data []byte) error {
	util.Logger.Infof("Attempting to insert shard %x into %x with data %x", blockHash, fileHash, data)
	file, ok := d.Progress[fileHash]
	if !ok {
		return fmt.Errorf("file %x not in download queue", fileHash)
	}

	offsets, ok := file.BlocksRemaining[blockHash]
	if !ok {
		return fmt.Errorf("block %x not in download queue", blockHash)
	}

	dataHash := sha256.Sum256(data)
	if !bytes.Equal(blockHash[:], dataHash[:]) {
		return fmt.Errorf("block %x data does not match expected content", blockHash)
	}

	for _, offset := range offsets {
		err := hash.InsertData(file.AbsolutePath, uint(len(data)), uint(offset), data)
		if err != nil {
			return err
		}
	}

	delete(file.BlocksRemaining, blockHash)
	util.Logger.Infof("successfully inserted shard %x into %x with data %x", blockHash, fileHash, data)

	if len(file.BlocksRemaining) == 0 {
		util.Logger.Infof("Download complete for file %s", file.AbsolutePath)
		err := CleanFile(file.AbsolutePath)
		if err != nil {
			util.Logger.Errorf("Error cleaning file %s: %s", file.AbsolutePath, err)
		}
	}

	return nil
}

// ContinueDownload continue/start a download where the desired blocks are already known
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

			for shard := range file.BlocksRemaining {
				util.Logger.Infof("Requesting shard %x for file %s in game %x", shard, file.AbsolutePath, gameHash)
				newRequest <- DownloadRequest{
					GameHash:  gameHash,
					BlockHash: shard,
					Attempts:  0,
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

// CleanFile /*
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

// GetProgress get the total download progress as a percent
func (d *Download) GetProgress() float32 {
	return 1 - float32(len(d.Progress))/float32(d.TotalBlocks)
}

// Finished whether any more blocks are still needed
func (d *Download) Finished() bool {
	for _, d := range d.Progress {
		if len(d.BlocksRemaining) > 0 {
			return false
		}
	}

	return true
}
