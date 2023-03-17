package games

import (
	"bytes"
	"crypto/sha256"
	"fmt"

	hash "github.com/t02smith/part-iii-project/toolkit/model/manager/hashtree"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

const shardInserterCount uint8 = 1

/*
	 <RequestDownload>  <--  [peer joins]
					|                     |
		[block found] - no - <DeferredRequest>
					|
				 yes
					|

[download block]

	|

<DownloadProgress>
*/
type DownloadManager struct {
	/*
		Download manager threads will send requests down this channel
		to prompt a peer listener to attempt to download the block
	*/
	RequestDownload chan DownloadRequest

	/*

	 */
	DeferredRequests chan DownloadRequest
}

type InsertShardRequest struct {
	FileHash  [32]byte
	BlockHash [32]byte
	Data      []byte
}

func NewDownloadManager() *DownloadManager {
	return &DownloadManager{
		RequestDownload:  make(chan DownloadRequest, 25),
		DeferredRequests: make(chan DownloadRequest, 25),
	}
}

func (d *DownloadManager) Close() {
	// close(d.DeferredRequests)
	// close(d.RequestDownload)
}

// workers

// func shardInserterPool(capacity int, game *Game) chan InsertShardRequest {
// 	util.Logger.Infof("Creating shard inserter pool")
// 	input := make(chan InsertShardRequest, 1)

// 	for w := 1; w <= capacity; w++ {
// 		go shardInserterWorker(w, game, input)
// 	}

// 	return input
// }

// func shardInserterWorker(id int, game *Game, input chan InsertShardRequest) {
// 	util.Logger.Debugf("INSERT WORKER %d: STARTED", id)
// 	for shard := range input {
// 		util.Logger.Debugf("INSERT WORKER %d: attempting to insert %x", id, shard.BlockHash)
// 		err := game.insertData(shard.FileHash, shard.BlockHash, shard.Data)
// 		if err != nil {
// 			util.Logger.Errorf("INSERT WORKER %d: error inserting shard %x => %s", id, shard.BlockHash, err)
// 		}
// 		util.Logger.Debugf("INSERT WORKER %d: Inserted %x", id, shard.BlockHash)
// 	}
// 	util.Logger.Debugf("INSERT WORKER %d: FINISHED", id)
// }

//

func (d *Download) InsertData(fileHash, blockHash [32]byte, data []byte) error {
	util.Logger.Debugf("Attempting to insert shard %x into %x", blockHash, fileHash)
	file, ok := d.Progress[fileHash]
	if !ok {
		util.Logger.Debugf("file %x not in download queue", fileHash)
		return nil
	}

	offsets, ok := file.BlocksRemaining[blockHash]
	if !ok {
		util.Logger.Debugf("block %x not in download queue", blockHash)
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
	}

	return nil
}
