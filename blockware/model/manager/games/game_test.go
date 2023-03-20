package games

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/t02smith/part-iii-project/toolkit/test/testutil"
)

/*

function: CreateGame
purpose: create a new game

? Test cases
success
	#1 => base case

failure
	illegal arguments
			#1 => root directory not found
			#2 => release date invalid
			#3 => invalid shard size
			#4 =>

*/

func TestCreateGame(t *testing.T) {

	t.Run("illegal arguments", func(t *testing.T) {
		datetime := time.Date(2002, 1, 10, 0, 0, 0, 0, time.UTC).String()

		table := []struct {
			name string
			game NewGame
		}{
			{"root directory", NewGame{"test-game", "1.1.1", datetime, "tcs1g20", "./fake/root/dir", big.NewInt(16), 256, "../../../test/data/assets"}},
			{"release date", NewGame{"test-game", "1.1.2", "not a real time", "tcs1g20", ".", big.NewInt(16), 256, "../../../test/data/assets"}},
			{"shard size", NewGame{"test-game", "1.1.1", datetime, "google.com", ".", big.NewInt(16), 0, "../../../test/data/assets"}},
			{"price", NewGame{"test-game", "1.1.1", datetime, "google.com", ".", big.NewInt(-1), 256, "../../../test/data/assets"}},
			{"assets directory", NewGame{"test-game", "1.1.1", datetime, "google.com", ".", big.NewInt(16), 256, "./test/data/assets"}},
		}

		for _, x := range table {
			t.Run(x.name, func(t *testing.T) {
				_, err := CreateGame(x.game, nil)
				assert.NotNilf(t, err, "Failed to detect illegal argument: %s", x.name)
			})
		}
	})

	t.Run("success", func(t *testing.T) {
		datetime := time.Date(2002, 1, 10, 0, 0, 0, 0, time.UTC).String()

		game, err := CreateGame(NewGame{
			"test-game", "1.1.1", datetime, "Tom Smith",
			"../../../test/data/testdir", big.NewInt(10), 64,
			"../../../test/data/assets"}, nil)
		assert.Nil(t, err, err)

		assert.Equal(t, "test-game", game.Title)
		assert.Equal(t, "1.1.1", game.Version)
		assert.Equal(t, datetime, game.ReleaseDate)
		assert.Equal(t, "Tom Smith", game.Developer)
		assert.Zero(t, big.NewInt(10).Cmp(game.Price))
		assert.Equal(t, "../../../test/data/assets", game.Assets.AbsolutePath)
		assert.NotNil(t, game.data)
		assert.NotNil(t, game.RootHash)

	})
}

/*

function: GetData
purpose: Get a game's hash tree. If it isn't present then get from file

Assumptions
#1 the hash file is always present => g.readHashData works every time

? Test cases
success
	#1 => Game data already present
	#2 => Game data read from file

failure
	NONE

*/

func TestGetData(t *testing.T) {
	g, err := fetchTestGame()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("success", func(t *testing.T) {
		tree := g.data

		t.Run("already loaded", func(t *testing.T) {
			data, err := g.GetData()
			assert.Nil(t, err)
			assert.True(t, tree.Equals(data))
		})

		t.Run("from file", func(t *testing.T) {
			g.data = nil
			t.Cleanup(func() {
				g.data = tree
			})

			data, err := g.GetData()
			assert.Nil(t, err)
			assert.True(t, tree.Equals(data))
		})
	})
}

/*

function: readHashData
purpose: read a game's hash tree from a file

? Test cases
success
	#1 => hash tree read from file

failure
	illegal arguments
			#1 file not found

*/

func TestReadHashData(t *testing.T) {
	g, err := fetchTestGame()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("success", func(t *testing.T) {
		tree := g.data
		g.data = nil
		t.Cleanup(func() {
			g.data = tree
		})

		err := g.readHashData()
		assert.Nil(t, err, err)
		assert.True(t, tree.Equals(g.data))
	})

	t.Run("failure", func(t *testing.T) {
		t.Run("illegal arguments", func(t *testing.T) {
			t.Run("hash file not found", func(t *testing.T) {
				old := g.RootHash
				g.RootHash = sha256.Sum256([]byte("hello world"))
				t.Cleanup(func() {
					g.RootHash = old
				})

				err := g.readHashData()
				assert.NotNil(t, err, err)
				assert.ErrorIs(t, err, os.ErrNotExist)
			})
		})
	})
}

/*

function: OutputAllGameDataToFile
purpose: Wrapper function for Game.OutputToFile and HashTree.OutputToFile

? Test cases
NONE => see Game.OutputToFile and HashTree.OutputToFile
*/

/*

function: Game.OutputToFile
purpose: Output game metadata to a file

? Test cases
success
	#1 => game data written

failure
	NONE => apart from unexpected IO errors

*/

func TestOutputToFile(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		g, err := fetchTestGame()
		if err != nil {
			t.Fatal(err)
		}

		viper.Set("meta.directory", "../../../test/data/tmp")
		os.Mkdir("../../../test/data/tmp/games", 0644)
		t.Cleanup(func() {
			viper.Set("meta.directory", "../../../test/data/.toolkit")
			testutil.ClearTmp("../../../")
		})

		err = g.OutputToFile()
		assert.Nil(t, err, err)

		stat, err := os.Stat(fmt.Sprintf("../../../test/data/tmp/games/%x", g.RootHash))
		assert.Nil(t, err)
		assert.NotZero(t, stat.Size())

	})
}
