package games

import (
	"crypto/sha256"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/t02smith/part-iii-project/toolkit/test/testutil"
)

func TestSetupDownload(t *testing.T) {
	testutil.ShortTest(t)
	gamesTestSetup()
	defer gamesTestTeardown()

	_, err := setupTestDownload()
	if err != nil {
		t.Error(err)
	}
}

func TestFindBlock(t *testing.T) {
	testutil.ShortTest(t)

	t.Cleanup(gamesTestTeardown)

	t.Run("success", func(t *testing.T) {
		gamesTestSetup()
		// TODO
	})

}

func TestContinueDownload(t *testing.T) {

	// * SETUP
	gameHash := sha256.Sum256([]byte("game"))
	fileHash := sha256.Sum256([]byte("file"))
	shardHash := sha256.Sum256([]byte("shard"))

	channel := make(chan DownloadRequest)

	download := Download{
		Progress:    make(map[[32]byte]FileProgress),
		TotalBlocks: 1,
	}

	fileProgress := FileProgress{
		AbsolutePath:    "test-file.x",
		BlocksRemaining: make(map[[32]byte]uint),
	}
	fileProgress.BlocksRemaining[shardHash] = 0

	download.Progress[fileHash] = fileProgress

	// * TESTT

	t.Run("success", func(t *testing.T) {
		download.ContinueDownload(gameHash, channel)

		request := <-channel
		assert.Equal(t, gameHash, request.GameHash, "correct game hash in request")
		assert.Equal(t, shardHash, request.BlockHash, "correct block hash in request")

	})

	close(channel)
}
