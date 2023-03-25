package games

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*

function: SetupDownload
purpose: create a new download of an owned game

assumptions:
#1 game hash data can be found and gotten
#2 dummy files are created successfully
#3 the user owns the game

? Test cases
success
	#1 => User creates a download of an owned game

failure
	illegal arguments
			#1 => download already exists

*/

func TestSetupDownload(t *testing.T) {
	gamesTestSetup()
	t.Cleanup(func() {
		gamesTestTeardown()
	})

	g := fetchTestGame(t)

	t.Run("success", func(t *testing.T) {
		t.Cleanup(func() {
			g.Download = nil
		})

		err := g.SetupDownload()
		t.Cleanup(func() {
			g.CancelDownload()
		})

		assert.Nil(t, err)

		assert.NotNil(t, g.Download)

		d := g.Download
		assert.Equal(t, filepath.Join("../../../test/data/tmp", "toolkit"), d.AbsolutePath)
		assert.Equal(t, 3, len(d.Progress))

		queuedBlocks := 0
		for _, f := range d.Progress {
			for _, blocks := range f.BlocksRemaining {
				queuedBlocks += len(blocks)
			}
		}

		assert.Equal(t, 236, queuedBlocks)
	})

	t.Run("failure", func(t *testing.T) {
		t.Run("illegal arguments", func(t *testing.T) {
			t.Run("download already exists", func(t *testing.T) {
				g.Download = &Download{}
				t.Cleanup(func() {
					g.Download = nil
				})

				err := g.SetupDownload()
				assert.NotNil(t, err)
				assert.Equal(t, fmt.Sprintf("download for game %x already exists", g.RootHash), err.Error())
			})
		})
	})
}

/*

function: CancelDownload
purpose: cancel an existing download and remove the contents

? Test cases
success
	#1 => download removed and files removed from storage

failure:
	illegal arguments
			#1 => download not active
			#2 => data not found at given path

*/

func TestCancelDownload(t *testing.T) {
	gamesTestSetup()
	t.Cleanup(func() {
		gamesTestTeardown()
	})

	t.Run("success", func(t *testing.T) {

		// setup
		g := setupTestDownload(t)

		path := g.Download.AbsolutePath

		//
		err := g.CancelDownload()

		assert.Nil(t, err)
		assert.Nil(t, g.Download)

		_, err = os.Stat(path)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, os.ErrNotExist)

	})

	t.Run("failure", func(t *testing.T) {
		t.Run("illegal arguments", func(t *testing.T) {
			g := setupTestDownload(t)
			path := g.Download.AbsolutePath

			t.Run("download not active", func(t *testing.T) {
				d := g.Download
				g.Download = nil
				t.Cleanup(func() {
					g.Download = d
				})

				err := g.CancelDownload()
				assert.NotNil(t, err)

				_, err = os.Stat(path)
				assert.Nil(t, err)
			})

			t.Run("data not found at path", func(t *testing.T) {
				g.Download.AbsolutePath = "./fake/path/oops"
				t.Cleanup(func() {
					g.Download.AbsolutePath = path
				})

				err := g.CancelDownload()
				assert.NotNil(t, err)
				assert.ErrorIs(t, err, os.ErrNotExist)
			})
		})
	})
}

/*

function: insertData
purpose: find, check and insert a block of data into a file

? Test cases
success
	#1 => all good :)

failure
	illegal arguments
			#1 => file not found in download
			#2 => block not found in download
			incorrect data
					#1 => contents
					#2 => length

*/

func TestInsertData(t *testing.T) {
	gamesTestSetup()
	t.Cleanup(func() {
		gamesTestTeardown()
	})

	g := setupTestDownload(t)

	t.Run("success", func(t *testing.T) {
		// setup
		data := make([]byte, g.data.ShardSize)
		for i := range data {
			data[i] = 255
		}

		hash := sha256.Sum256(data)
		file := g.data.RootDir.Subdirs["subdir"].Files["chip8.c"]

		g.Download.Progress[file.RootHash].BlocksRemaining[hash] = []uint{0}

		//
		err := g.insertData(file.RootHash, hash, data)
		assert.Nil(t, err)

		// read inserted data
		f, err := os.Open("../../../test/data/tmp/toolkit/subdir/chip8.c")
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()

		buffer := make([]byte, g.data.ShardSize)
		reader := bufio.NewReader(f)

		reader.Read(buffer)
		assert.Equal(t, data, buffer)
	})

	t.Run("failure", func(t *testing.T) {
		t.Run("illegal arguments", func(t *testing.T) {
			t.Run("file not in download queue", func(t *testing.T) {
				fakeFileHash := sha256.Sum256([]byte("hello"))
				err := g.insertData(fakeFileHash, [32]byte{}, []byte{})
				assert.NotNil(t, err)
				assert.Equal(t, fmt.Sprintf("file %x not in download queue", fakeFileHash), err.Error())
			})

			file := g.data.RootDir.Files["test.txt"]

			t.Run("block not in download queue", func(t *testing.T) {
				fakeBlockHash := sha256.Sum256([]byte("hello"))
				err := g.insertData(file.RootHash, fakeBlockHash, []byte{})
				assert.NotNil(t, err)
				assert.Equal(t, fmt.Sprintf("block %x not in download queue", fakeBlockHash), err.Error())
			})

			block := file.Hashes[0]

			t.Run("incorrect data", func(t *testing.T) {
				t.Run("contents", func(t *testing.T) {
					err := g.insertData(file.RootHash, block, make([]byte, g.data.ShardSize))
					assert.NotNil(t, err)
					assert.Equal(t, fmt.Sprintf("block %x data does not match expected content", block), err.Error())
				})

				t.Run("length", func(t *testing.T) {
					err := g.insertData(file.RootHash, block, make([]byte, 16))
					assert.NotNil(t, err)
					assert.Equal(t, fmt.Sprintf("data is not correct length. Got %d, expected %d", 16, g.data.ShardSize), err.Error())
				})
			})
		})
	})
}

/*

function: CleanFile
purpose: remove the trailing NULL bytes from a file

? Test cases
success
	#1 trailing null bytes removed

failure
	illegal arguments
			#1 => file not found
			#2 => truncate size too large

*/

func TestCleanFile(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		filename := "../../../test/data/tmp/trunacte.txt"
		t.Cleanup(func() {
			os.Remove(filename)
		})

		f, err := os.Create(filename)
		if err != nil {
			t.Fatal(err)
		}

		data := []byte("hello world")
		writer := bufio.NewWriter(f)

		for i := 0; i < 10; i++ {
			writer.Write(data)
			writer.Flush()
		}

		f.Close()

		stat, err := os.Stat(filename)
		assert.Nil(t, err)
		assert.Equal(t, int64(len(data)*10), stat.Size(), "initial size")

		//
		err = CleanFile(filename, len(data)*4)
		assert.Nil(t, err)

		//
		stat, err = os.Stat(filename)
		assert.Nil(t, err)
		assert.Equal(t, int64(len(data)*4), stat.Size(), "final size")
	})

	t.Run("failure", func(t *testing.T) {
		t.Run("illegal arguments", func(t *testing.T) {
			t.Run("file not found", func(t *testing.T) {
				err := CleanFile("./fake/file/not/here.txt", 100)
				assert.NotNil(t, err)
				assert.ErrorIs(t, err, os.ErrNotExist)
			})

			t.Run("file too small", func(t *testing.T) {
				err := CleanFile("../../../test/data/testfiles/TRUNCATE.txt", 100000)
				assert.NotNil(t, err)
				assert.Equal(t, fmt.Sprintf("file is too small to be truncated to %d", 100000), err.Error())
			})
		})
	})
}

/*
TODO
*/

func TestContinueDownload(t *testing.T) {

	// * SETUP
	gameHash := sha256.Sum256([]byte("game"))
	fileHash := sha256.Sum256([]byte("file"))
	shardHash := sha256.Sum256([]byte("shard"))

	channel := make(chan DownloadRequest)

	download := Download{
		Progress:    make(map[[32]byte]*FileProgress),
		TotalBlocks: 1,
	}

	fileProgress := &FileProgress{
		AbsolutePath:    "test-file.x",
		BlocksRemaining: make(map[[32]byte][]uint),
	}
	fileProgress.BlocksRemaining[shardHash] = []uint{0}

	download.Progress[fileHash] = fileProgress

	// * TESTS

	t.Run("success", func(t *testing.T) {
		download.ContinueDownload(gameHash, channel)

		request := <-channel

		t.Run("request received", func(t *testing.T) {
			assert.Equal(t, gameHash, request.GameHash, "correct game hash in request")
			assert.Equal(t, shardHash, request.BlockHash, "correct block hash in request")
		})

	})

	close(channel)
}

/*

function: Game.completeFile
purpose: clean and verify a file after downloading all data

? Test cases
success
	#1 => file finished && file cleaned and verified correctly
	#2 => file not finished

failure
	#1 => file not valid
	#2 => file not found

*/

func TestCompleteFile(t *testing.T) {
	g := setupTestDownload(t)

	data, err := g.GetData()
	if err != nil {
		t.Fatal(err)
	}

	fileHash := data.RootDir.Files["test.txt"].RootHash
	file := g.Download.Progress[fileHash]

	t.Run("success", func(t *testing.T) {
		t.Run("download finished", func(t *testing.T) {
			oldBR := file.BlocksRemaining
			file.BlocksRemaining = make(map[[32]byte][]uint)

			oldPath := file.AbsolutePath
			file.AbsolutePath = "../../../test/data/testdir/test.txt"

			t.Cleanup(func() {
				file.AbsolutePath = oldPath
				file.BlocksRemaining = oldBR
			})

			err := g.completeFile(fileHash, file)
			assert.Nil(t, err, err)
		})
	})

	t.Run("failure", func(t *testing.T) {
		t.Run("file not found", func(t *testing.T) {
			oldPath := file.AbsolutePath
			file.AbsolutePath = "../../../test/data/testdir/fake/file"

			t.Cleanup(func() {
				file.AbsolutePath = oldPath
			})

			err := g.completeFile(fileHash, file)
			assert.NotNil(t, err)
			assert.ErrorIs(t, err, os.ErrNotExist)
		})

		t.Run("file not valid", func(t *testing.T) {
			oldBR := file.BlocksRemaining
			file.BlocksRemaining = make(map[[32]byte][]uint)

			oldPath := file.AbsolutePath
			file.AbsolutePath = "../../../test/data/testfiles/bad-verify.txt"

			t.Cleanup(func() {
				file.AbsolutePath = oldPath
				file.BlocksRemaining = oldBR
			})

			// will fail verification and add all wrong blocks back to progress
			err := g.completeFile(fileHash, file)
			assert.Nil(t, err)

			assert.NotZero(t, file.BlocksRemaining)
			assert.Equal(t, len(oldBR), len(file.BlocksRemaining))
		})
	})
}
