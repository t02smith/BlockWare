package hash

import (
	"bufio"
	"bytes"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

/*

function: CreateDummyFiles
purpose: create a set of dummy files from a given directory

? Test cases
success
	#1 => all files created

failure
	#1 => root dir not found
	#2

*/

func TestCreateDummyFiles(t *testing.T) {
	ht, err := NewHashTree("../../../test/data/testdir", 1024, nil)
	if err != nil {
		t.Fatalf("NewHashTree failed on valid input")
	}

	if err := ht.Hash(); err != nil {
		t.Fatal(err)
	}

	t.Run("success", func(t *testing.T) {
		counter := 0
		var counterLock sync.Mutex

		fs := ht.ListFiles()

		err := ht.CreateDummyFiles("../../../test/data/tmp", "dummyTest", func(s string, file *HashTreeFile) {
			counterLock.Lock()
			counter++
			counterLock.Unlock()
		})
		t.Cleanup(func() {
			os.RemoveAll("../../../test/data/tmp/dummyTest")
		})

		assert.Nil(t, err)
		assert.Equal(t, len(fs), counter)

		for _, f := range fs {
			stat, err := os.Stat(filepath.Join("../../../test/data/tmp/dummyTest", f.RelativeFilename))
			assert.Nil(t, err)
			assert.Equal(t, f.Size, int(stat.Size()))
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Run("root dir not found", func(t *testing.T) {
			err := ht.CreateDummyFiles("./not/found/dont/make", "fake", func(s string, file *HashTreeFile) {
			})

			assert.NotNil(t, err)
			assert.ErrorIs(t, err, os.ErrNotExist)
		})
	})
}

/*

function: pushFilesToCreateDummies
purpose: push all a hash trees files down a channel

? Test cases
success
	#1 => all files pushed

*/

func TestPushFilesToCreateDummies(t *testing.T) {
	ht, err := NewHashTree("../../../test/data/testdir", 1024, nil)
	if err != nil {
		t.Fatalf("NewHashTree failed on valid input")
	}

	if err := ht.Hash(); err != nil {
		t.Fatal(err)
	}

	t.Run("success", func(t *testing.T) {
		fs := ht.ListFiles()

		toCreate := make(chan *HashTreeFile, len(fs))
		t.Cleanup(func() {
			close(toCreate)
		})

		pushFilesToCreateDummies(ht, toCreate)

		for i := 0; i < len(fs); i++ {
			f, found := <-toCreate, false
			for _, file := range fs {
				if file.Equals(f) {
					found = true
					break
				}
			}

			if !found {
				t.Errorf("File %s not expected", f.AbsoluteFilename)
			}
		}
	})
}

/*

function: dummyFileCreatorWorker
purpose: create files based upon input from channel

? Test cases
success
	#1 => files created

*/

func TestDummyFileCreatorWorker(t *testing.T) {
	htf := &HashTreeFile{
		Filename:         "test.txt",
		AbsoluteFilename: "../../../test/data/testdir/test.txt",
		RelativeFilename: "./test.txt",
	}

	if err := htf.shardFile(4); err != nil {
		t.Fatal(err)
	}

	t.Run("success", func(t *testing.T) {
		onCreateCalled := false
		var wg sync.WaitGroup
		wg.Add(1)
		toCreate := make(chan *HashTreeFile, 10)

		toCreate <- htf
		go dummyFileCreatorWorker("../../../test/data/tmp", 4, &wg, toCreate, func(s string, file *HashTreeFile) {
			onCreateCalled = true
		})

		time.Sleep(25 * time.Millisecond)
		close(toCreate)

		t.Cleanup(func() {
			os.Remove("../../../test/data/tmp/test.txt")
		})

		assert.True(t, onCreateCalled)
		stat, err := os.Stat("../../../test/data/tmp/test.txt")
		assert.Nil(t, err)
		assert.Equal(t, htf.Size, int(stat.Size()))

		wg.Wait()
	})
}

/*

function: setupFile
purpose: create skeleton file from given hashes

? Test cases
success
	| #1 => fileSize % shardSize == 0
	| #2 => fileSize % shardSize != 0
	| #3 => empty file

failure
	| illegal arguments
			| #1 => root directory doesn't exist
			| #2 => title matches file in root directory

*/

func TestSetupFile(t *testing.T) {
	tmpFile := "../../test/data/tmp/skeleton.txt"

	smoke := t.Run("function setupFile", func(t *testing.T) {

		t.Run("success", func(t *testing.T) {

			inputs := []struct {
				name       string
				shardSize  uint
				shardCount int
				fileSize   int
			}{
				{"fileSize % shardSize == 0", 64, 100, 6400},
				{"fileSize % shardSize != 0", 64, 101, 6452},
				{"empty file", 0, 0, 0},
			}

			for _, in := range inputs {
				t.Run(in.name, func(t *testing.T) {
					err := setupFile(tmpFile, in.shardSize, in.shardCount, in.fileSize)
					assert.Nil(t, err)

					t.Cleanup(func() {
						os.Remove(tmpFile)
					})

					file, err := os.Stat(tmpFile)
					assert.Nil(t, err)

					assert.Equal(t, in.fileSize, int(file.Size()))
				})
			}

		})

		t.Run("illegal arguments", func(t *testing.T) {

			// TODO

		})
	})

	if !smoke {
		t.FailNow()
	}
}

// Inserting shards

func TestInsertData(t *testing.T) {
	tmpFile := "../../../test/data/tmp/skeleton.txt"
	err := setupFile(tmpFile, 64, 100, 6400)
	if err != nil {
		t.Fatalf("error setting up dummy file %s", err)
	}

	t.Run("success", func(t *testing.T) {
		newShard := make([]byte, 64)
		for i := range newShard {
			newShard[i] = 255
		}

		err = InsertData("../../../test/data/tmp/skeleton.txt", 64, 12, newShard)
		if err != nil {
			t.Errorf("%s", err)
		}

		// ? check the changes occurred
		file, err := os.Open(tmpFile)
		if err != nil {
			t.Fatalf("error opening tmp file %s", err)
		}
		defer file.Close()

		buffer := make([]byte, 64)
		reader := bufio.NewReader(file)

		file.Seek(64*12, 0)
		reader.Read(buffer)

		if !bytes.Equal(buffer, newShard) {
			t.Error("Shard not written correctly to file")
		}
	})

	t.Run("illegal arguments", func(t *testing.T) {
		// TODO
	})

	err = os.Remove(tmpFile)
	if err != nil {
		t.Errorf("Error removing tmp file %s", err)
	}
}
