package hash

import (
	"sync"

	"github.com/t02smith/part-iii-project/toolkit/util"
)

/*

We can shard many files concurrently and this is achieved using a set of
worker processes that each handle an individual file.

Files are sent through a channel and are greedily by any free worker until
there are no more files to be processed. All files in the hash tree are
discovered before this is started so the ordering of files does not matter.

*/

// starts a worker pool with N workers awaiting instructions
func hasherPool(capacity int, fileCount int, shardSize uint, progress chan int) (*sync.WaitGroup, chan *HashTreeFile, chan error) {
	util.Logger.Infof("Creating hasher pool with %d workers fir %d files", capacity, fileCount)
	files := make(chan *HashTreeFile, fileCount)
	errors := make(chan error, fileCount)

	var wg sync.WaitGroup
	wg.Add(fileCount)

	for w := 1; w <= capacity; w++ {
		go worker(w, &wg, shardSize, files, errors, progress)
	}

	util.Logger.Info("Hasher pool generated")
	return &wg, files, errors
}

// start a new worker thread to greedily take files sent down a channel
func worker(id int, wg *sync.WaitGroup, shardSize uint, files <-chan *HashTreeFile, errors chan<- error, progress chan int) {
	for f := range files {
		util.Logger.Infof("WORKER %d: Sharding file %s", id, f.AbsoluteFilename)
		err := f.shardFile(shardSize)
		if err != nil {
			util.Logger.Errorf("WORKER %d: Error sharding file %s: %s", id, f.AbsoluteFilename, err)
			// errors <- err
		}
		wg.Done()
		util.Logger.Infof("WORKER %d: Completed sharding %s", id, f.AbsoluteFilename)

		if progress != nil {
			progress <- 1
		}
	}

	util.Logger.Infof("WORKER %d: FINISHED", id)
}
