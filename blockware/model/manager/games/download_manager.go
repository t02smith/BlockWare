package games

import (
	"github.com/t02smith/part-iii-project/toolkit/util"
)

const shardInserterCount uint8 = 10

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
		This request has failed at least once and needs to be queued again
	*/
	DeferredRequests chan DownloadRequest
}

// request for the data manager to insert a piece of data
type InsertShardRequest struct {
	FileHash  [32]byte
	BlockHash [32]byte
	Data      []byte
}

// creatae a new download manager
func NewDownloadManager() *DownloadManager {
	return &DownloadManager{
		RequestDownload:  make(chan DownloadRequest, 50),
		DeferredRequests: make(chan DownloadRequest, 50),
	}
}

// stop the download manager
func (d *DownloadManager) Close() {
	close(d.RequestDownload)
	close(d.DeferredRequests)
}

// workers

/*
Start a new shard inserter pool that will process requests sent by the
networking component to insert a piece of data into storage.

This will distribute insertion across N different nodes and is exclusive for
a given game
*/
func shardInserterPool(capacity int, game *Game) chan InsertShardRequest {
	util.Logger.Infof("Creating shard inserter pool")
	input := make(chan InsertShardRequest, 10)

	for w := 1; w <= capacity; w++ {
		go shardInserterWorker(w, game, input)
	}

	return input
}

func shardInserterWorker(id int, game *Game, input chan InsertShardRequest) {
	util.Logger.Debugf("INSERT WORKER %d: STARTED", id)
	for shard := range input {
		util.Logger.Debugf("INSERT WORKER %d: attempting to insert %x", id, shard.BlockHash)
		err := game.insertData(shard.FileHash, shard.BlockHash, shard.Data)
		if err != nil {
			util.Logger.Debugf("INSERT WORKER %d: error inserting shard %x => %s", id, shard.BlockHash, err)
		}
		util.Logger.Debugf("INSERT WORKER %d: Inserted %x", id, shard.BlockHash)
	}
	util.Logger.Debugf("INSERT WORKER %d: FINISHED", id)
}
