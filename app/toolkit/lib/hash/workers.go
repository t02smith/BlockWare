package hash

import (
	"log"
	"sync"
)

/**

A worker pool for sharding and hashing files

*/

// starts a worker pool with N workers awaiting instructions
func hasherPool(capacity int, fileCount int, shardSize uint) (*sync.WaitGroup, chan *HashTreeFile, chan error) {

	files := make(chan *HashTreeFile, fileCount)
	errors := make(chan error, fileCount)

	var wg sync.WaitGroup
	wg.Add(fileCount)

	for w := 1; w <= capacity; w++ {
		go worker(w, &wg, shardSize, files, errors)
	}

	return &wg, files, errors
}

func worker(id int, wg *sync.WaitGroup, shardSize uint, files <-chan *HashTreeFile, errors chan<- error) {

	for f := range files {
		log.Printf("WORKER %d: Sharding file %s\n", id, f.AbsoluteFilename)
		err := f.shardFile(shardSize)
		if err != nil {
			errors <- err
		}
		wg.Done()
	}
}