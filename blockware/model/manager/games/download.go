package games

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"

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

type DownloadStage string

const (
	DS_NotStarted          DownloadStage = "Not started"
	DS_GeneratingDummies   DownloadStage = "Setting up download"
	DS_Downloading         DownloadStage = "Downloading data"
	DS_FinishedDownloading DownloadStage = "Finished"
	DS_Cancelled           DownloadStage = "Cancelled"
)

// Download A download manager for a game
type Download struct {

	// progress of each file being downloaded
	Progress     map[[32]byte]*FileProgress
	progressLock sync.Mutex

	// total number of blocks to install
	TotalBlocks int

	// absolute path of root directory
	AbsolutePath string

	//
	inserterPool chan InsertShardRequest

	// the stage the download is in
	Stage DownloadStage
}

// FileProgress the progress of a specific file's download
type FileProgress struct {

	// where the file is located in storage
	AbsolutePath string

	// block hash => position in file
	BlocksRemaining map[[32]byte][]uint

	Size int

	lock sync.Mutex
}

// DownloadRequest /*
type DownloadRequest struct {

	// details to uniquely identify the block
	GameHash  [32]byte
	BlockHash [32]byte
}

// ? setup

// create a new download for an existing game
func (g *Game) SetupDownload() error {

	// read hash data if isn't already loaded
	data, err := g.GetData()
	if err != nil {
		return err
	}

	// create download obj
	d := &Download{
		Progress:    make(map[[32]byte]*FileProgress),
		TotalBlocks: 0,
		Stage:       DS_GeneratingDummies,
	}
	g.Download = d

	// d.inserterPool = shardInserterPool(int(shardInserterCount), g)

	// generate dummy files
	dir := viper.GetString("games.installFolder")
	if len(dir) == 0 {
		return errors.New("game install folder not found")
	}

	d.AbsolutePath = filepath.Join(dir, g.Title)

	util.Logger.Infof("Generating dummy files for %s-%s", g.Title, g.Version)
	err = data.CreateDummyFiles(dir, g.Title, func(path string, htf *hash.HashTreeFile) {
		p := &FileProgress{
			AbsolutePath:    path,
			BlocksRemaining: make(map[[32]byte][]uint),
			Size:            htf.Size,
		}

		for i, b := range htf.Hashes {
			if offsets, ok := p.BlocksRemaining[b]; ok {
				p.BlocksRemaining[b] = append(offsets, uint(i))
			} else {
				p.BlocksRemaining[b] = []uint{uint(i)}
			}
		}

		d.TotalBlocks += len(p.BlocksRemaining)

		d.progressLock.Lock()
		d.Progress[htf.RootHash] = p
		d.progressLock.Unlock()
	})

	if err != nil {
		return err
	}

	d.Stage = DS_Downloading
	return nil
}

// CancelDownload cancel an existing download
func (g *Game) CancelDownload() error {
	if g.Download == nil {
		util.Logger.Warnf("Download not found for game %x", g.RootHash)
		return nil
	}

	g.Download.Stage = DS_Cancelled
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
func (g *Game) insertData(fileHash, blockHash [32]byte, data []byte) error {
	util.Logger.Debugf("Attempting to insert shard %x into %x", blockHash, fileHash)
	d := g.Download

	d.progressLock.Lock()
	file, ok := d.Progress[fileHash]
	d.progressLock.Unlock()

	if !ok {
		util.Logger.Debugf("file %x not in download queue", fileHash)
		return nil
	}

	file.lock.Lock()
	defer file.lock.Unlock()

	offsets, ok := file.BlocksRemaining[blockHash]
	if !ok {
		util.Logger.Warnf("block %x not in download queue", blockHash)
		return nil
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

	d.progressLock.Lock()
	delete(file.BlocksRemaining, blockHash)
	util.Logger.Debugf("successfully inserted shard %x into %x", blockHash, fileHash)
	d.progressLock.Unlock()

	if len(file.BlocksRemaining) == 0 {
		util.Logger.Debugf("Download complete for file %s", file.AbsolutePath)
		err := CleanFile(file.AbsolutePath, file.Size)
		if err != nil {
			util.Logger.Errorf("Error cleaning file %s: %s", file.AbsolutePath, err)
		}

		// verify data
		// data, err := g.GetData()
		// if err != nil {
		// 	return err
		// }

		// htf := data.GetFile(fileHash)
		// if htf == nil {
		// 	// ? shouldn't even make it this far but just in case
		// 	return nil
		// }

		// correct, incorrectBlocks, err := hash.VerifyFile(htf, file.AbsolutePath, data.ShardSize)
		// if err != nil {
		// 	return err
		// }

		// if correct {
		// 	return nil
		// }

		// util.Logger.Debugf("Found %d blocks in %s that are incorrect", len(incorrectBlocks), file.AbsolutePath)
		// for blockHash, offset := range incorrectBlocks {
		// 	file.BlocksRemaining[blockHash] = []uint{offset}
		// }
	}

	return nil
}

// ContinueDownload continue/start a download where the desired blocks are already known
// this function will make requests down the libraries download request
// channel
func (d *Download) ContinueDownload(gameHash [32]byte, newRequest chan DownloadRequest) {

	/*
		TODO
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

		d.progressLock.Lock()
		for _, file := range d.Progress {
			util.Logger.Debugf("Requesting file %s for game %x", file.AbsolutePath, gameHash)

			for shard := range file.BlocksRemaining {
				file.lock.Lock()
				util.Logger.Debugf("Requesting shard %x for file %s in game %x", shard, file.AbsolutePath, gameHash)
				newRequest <- DownloadRequest{
					GameHash:  gameHash,
					BlockHash: shard,
				}
				file.lock.Unlock()
			}
		}
		d.progressLock.Unlock()
	}()
}

// ? misc functions

// CleanFile /*
func CleanFile(path string, size int) error {
	util.Logger.Debugf("Cleaning file %s", path)

	// stat, err := os.Stat(path)
	// if err != nil {
	// 	return err
	// }

	file, err := os.OpenFile(path, os.O_RDWR, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	// reader := bufio.NewReader(file)
	// bytesToRemove, buf := 0, make([]byte, 1)
	// for {
	// 	file.Seek(int64(bytesToRemove)-1, io.SeekEnd)
	// 	_, err := reader.Read(buf)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	if buf[0] == 0x00 {
	// 		bytesToRemove--
	// 	} else {
	// 		break
	// 	}
	// }

	if err = file.Truncate(int64(size)); err != nil {
		return err
	}
	util.Logger.Debugf("Cleaned file %s", path)
	return nil
}

// GetProgress get the total download progress as a percent
func (d *Download) GetProgress() float32 {
	d.progressLock.Lock()
	length := len(d.Progress)
	d.progressLock.Unlock()

	return 1 - float32(length)/float32(d.TotalBlocks)
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

func (d *Download) GetProgressLock() *sync.Mutex {
	return &d.progressLock
}

func (d *Download) InserterPool() chan InsertShardRequest {
	return d.inserterPool
}
