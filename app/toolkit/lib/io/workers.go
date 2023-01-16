package io

import "sync"

/**

A worker pool for sharding and hashing files

*/

// starts a worker pool with N workers awaiting instructions
func hasherPool(capacity int, fileCount int, shardSize uint) (*sync.WaitGroup, chan *hashTreeFile, chan error) {

	files := make(chan *hashTreeFile, fileCount)
	errors := make(chan error, fileCount)

	var wg sync.WaitGroup
	wg.Add(fileCount)

	for w := 1; w <= capacity; w++ {
		go worker(&wg, shardSize, files, errors)
	}

	return &wg, files, errors
}

func worker(wg *sync.WaitGroup, shardSize uint, files <-chan *hashTreeFile, errors chan<- error) {

	for f := range files {
		err := f.shardFile(shardSize)
		if err != nil {
			errors <- err
		}
		wg.Done()
	}
}
