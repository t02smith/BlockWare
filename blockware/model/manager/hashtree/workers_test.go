package hash

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

/*

function: worker
purpose: Start a worker to accept files from a given channel

? Test cases
success
	| #1 sharded successfully

*/

func TestWorker(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(1)

		files := make(chan *HashTreeFile, 1)
		errors := make(chan error, 1)

		file := &HashTreeFile{
			Filename:         "chip8.c",
			AbsoluteFilename: "../../test/data/testdir/subdir/chip8.c",
		}
		files <- file

		go worker(1, &wg, 256, files, errors, nil)
		time.Sleep(25 * time.Millisecond)
		close(files)

		assert.NotNil(t, file.Hashes, "hash tree not generated")
		assert.NotNil(t, file.RootHash, "root hash not calculated")
	})

}

/*

function: hasherPool
purpose: Create a set of workers

? Test cases
TODO

*/
