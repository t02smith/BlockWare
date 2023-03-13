package games

import "github.com/t02smith/part-iii-project/toolkit/util"

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

	/*
		Once a block has been downloaded a message will be sent down
		this channel.
		This can be used to signal to the UI that a download has been
		completed
	*/
	DownloadProgress chan DownloadRequest
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
		DownloadProgress: make(chan DownloadRequest, 25),
	}
}

func (d *DownloadManager) Close() {
	close(d.DeferredRequests)
	close(d.DownloadProgress)
	close(d.RequestDownload)
}

// workers

func shardInserterPool(capacity int, download *Download) chan InsertShardRequest {
	util.Logger.Infof("Creating shard inserted pool")
	input := make(chan InsertShardRequest, 25)

	for w := 1; w <= capacity; w++ {
		go shardInserterWorker(w, download, input)
	}

	return input
}

func shardInserterWorker(id int, download *Download, input chan InsertShardRequest) {
	util.Logger.Debugf("INSERT WORKER %d: STARTED", id)
	for shard := range input {
		util.Logger.Debugf("INSERT WORKER %d: attempting to insert %x", id, shard.BlockHash)
		err := download.insertData(shard.FileHash, shard.BlockHash, shard.Data)
		if err != nil {
			util.Logger.Errorf("INSERT WORKER %d: error inserted shard %x => %s", id, shard.BlockHash, err)
		}
		util.Logger.Debugf("INSERT WORKER %d: Inserted %x", id, shard.BlockHash)
	}
	util.Logger.Debugf("INSERT WORKER %d: FINISHED", id)

}
