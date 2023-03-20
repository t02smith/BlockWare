package games

import (
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

// helper

func setupTestLibrary(t *testing.T) (*Library, *Game) {
	lib := NewLibrary()
	testGame, err := fetchTestGame()
	if err != nil {
		t.Fatal(err)
	}

	lib.AddOrUpdateOwnedGame(testGame)
	t.Cleanup(func() {
		lib.ClearOwnedGames()
		lib.Close()
	})

	return lib, testGame
}

// tests

/*

function: NewLibrary
purpose: create a new instance of a library

? Test cases
success
	#1 => base case

*/

func TestNewLibrary(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		l := NewLibrary()

		assert.NotNil(t, l.ownedGames)
		assert.NotNil(t, l.blockchainGames)
		assert.NotNil(t, l.DownloadManager)
	})
}

/*

function: CreateDownload
purpose: create a download from an existing owned game

? Test cases
success
	#1 => download created
					a - dummy files made
					b - download continued

failure
	illegal arguments
			#1 => game not owned
			#2 => download already started

	unexpected err
			#1 => download not started but files for download already exist

*/

func TestCreateDownload(t *testing.T) {
	lib, testGame := setupTestLibrary(t)

	t.Run("success", func(t *testing.T) {
		err := lib.CreateDownload(testGame)
		assert.Nil(t, err, err)
		t.Cleanup(func() {
			lib.Uninstall(testGame.RootHash)
		})

		assert.NotNil(t, testGame.Download)

		// check files have been created
		for _, f := range testGame.Download.Progress {
			_, err := os.Stat(f.AbsolutePath)
			assert.Nil(t, err)
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Run("illegal arguments", func(t *testing.T) {
			t.Run("game not owned", func(t *testing.T) {
				fakeGame := &Game{
					RootHash: sha256.Sum256([]byte("hello world")),
				}

				err := lib.CreateDownload(fakeGame)
				assert.NotNil(t, err)
				assert.Equal(t, fmt.Sprintf("game %x not found in library, cannot add download", fakeGame.RootHash), err.Error())
			})

			t.Run("download already started", func(t *testing.T) {
				err := lib.CreateDownload(testGame)
				if err != nil {
					t.Fatal(err)
				}

				t.Cleanup(func() {
					lib.Uninstall(testGame.RootHash)
				})

				err = lib.CreateDownload(testGame)
				assert.NotNil(t, err)
				assert.Equal(t, fmt.Sprintf("download already started for game %x", testGame.RootHash), err.Error())
			})
		})

		t.Run("unexpected err", func(t *testing.T) {
			t.Run("files with same name already exist", func(t *testing.T) {
				dir := filepath.Join(viper.GetString("games.installFolder"), testGame.Title)
				f, err := os.Create(dir)
				if err != nil {
					t.Fatal(err)
				}
				f.Close()
				t.Cleanup(func() {
					os.Remove(dir)
				})

				err = lib.CreateDownload(testGame)
				t.Cleanup(func() {
					lib.Uninstall(testGame.RootHash)
				})

				assert.NotNil(t, err)
				assert.Equal(t, fmt.Sprintf("folder %s already found, cannot start download", dir), err.Error())
			})
		})

	})
}

/*

function: GetOwnedGame
purpose: retrieve a user's owned game by its root hash or return null

? Test cases
success
	#1 => game found
	#2 => game not found

*/

func TestGetOwnedGame(t *testing.T) {
	lib, testGame := setupTestLibrary(t)

	t.Run("success", func(t *testing.T) {
		t.Run("game owned", func(t *testing.T) {
			res := lib.GetOwnedGame(testGame.RootHash)
			assert.NotNil(t, res)
			assert.True(t, testGame.Equals(res))
		})

		t.Run("game not owned", func(t *testing.T) {
			var empty [32]byte
			res := lib.GetOwnedGame(empty)
			assert.Nil(t, res)
		})
	})
}

/*

function: GetOwnedGames
purpose: get a list of a user's owned games (that are stored locally)

? Test cases
success
	#1 => no games
	#2 => one game
	#3 => many games

*/

func TestGetOwnedGames(t *testing.T) {
	lib, testGame := setupTestLibrary(t)

	t.Run("success", func(t *testing.T) {
		t.Run("no games", func(t *testing.T) {
			lib.ClearOwnedGames()
			t.Cleanup(func() {
				lib.AddOrUpdateOwnedGame(testGame)
			})

			res := lib.GetBlockchainGames()
			assert.Empty(t, res)
		})

		t.Run("one game", func(t *testing.T) {
			res := lib.GetOwnedGames()
			assert.Equal(t, 1, len(res))
			assert.True(t, testGame.Equals(res[0]))
		})

		t.Run("many games", func(t *testing.T) {
			games := []*Game{
				{
					RootHash: sha256.Sum256([]byte("hello world")),
				},
				{
					RootHash: sha256.Sum256([]byte("tom smith")),
				},
				{
					RootHash: sha256.Sum256([]byte("root hash")),
				},
			}

			for _, g := range games {
				lib.AddOrUpdateOwnedGame(g)
			}

			t.Cleanup(func() {
				lib.ClearOwnedGames()
				lib.AddOrUpdateOwnedGame(testGame)
			})

			res := lib.GetOwnedGames()
			assert.Equal(t, 4, len(res))

		})
	})
}

/*

function: AddOrUpdateOwnedGame
purpose: add a new game or overwrite an existing game

? Test cases
success
	#1 => new game
	#2 => overwrite existing game

*/

func TestAddOrUpdateOwnedGame(t *testing.T) {
	lib := NewLibrary()
	testGame, err := fetchTestGame()
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		lib.Close()
	})

	t.Run("success", func(t *testing.T) {
		t.Run("new game", func(t *testing.T) {
			lib.AddOrUpdateOwnedGame(testGame)
			t.Cleanup(func() {
				lib.ClearOwnedGames()
			})

			g := lib.GetOwnedGame(testGame.RootHash)
			assert.NotNil(t, g)
			assert.True(t, testGame.Equals(g))
		})

		t.Run("overwrite existing game", func(t *testing.T) {
			testGameII := *testGame
			testGameII.Title = "hello world"
			testGameII.Developer = "not tom smith"

			lib.AddOrUpdateOwnedGame(testGame)
			t.Cleanup(func() {
				lib.ClearOwnedGames()
			})

			lib.AddOrUpdateOwnedGame(&testGameII)

			res := lib.GetOwnedGame(testGame.RootHash)
			assert.NotNil(t, res)
			assert.False(t, testGame.Equals(res))
			assert.True(t, testGameII.Equals(res))
		})
	})
}

/*

function: FindAndRetrieveBlock
purpose: find and retrieve a given block given its game and block hash

? Test cases
success
	#1 => correct block received


failure
	illegal arguments
			#1 => block doesn't belong to game
			#2 => game not owned
*/

func TestFindAndRetrieveBlock(t *testing.T) {
	lib, testGame := setupTestLibrary(t)

	t.Run("success", func(t *testing.T) {
		hashTree, err := testGame.GetData()
		if err != nil {
			t.Fatal(err)
		}

		hash := hashTree.RootDir.Files["test.txt"].Hashes[0]

		found, data, err := lib.FindAndRetrieveBlock(testGame.RootHash, hash)
		assert.Nil(t, err)
		assert.Equal(t, hashTree.ShardSize, uint(len(data)))
		assert.Equal(t, hash, sha256.Sum256(data))
		assert.True(t, found)
	})

	t.Run("failure", func(t *testing.T) {
		fakeHash := sha256.Sum256([]byte("hi"))

		t.Run("illegal arguments", func(t *testing.T) {
			t.Run("block doesn't belong to game", func(t *testing.T) {
				found, data, err := lib.FindAndRetrieveBlock(testGame.RootHash, fakeHash)

				assert.Nil(t, err)
				assert.Nil(t, data)
				assert.False(t, found)
			})

			t.Run("game not owned", func(t *testing.T) {
				found, data, err := lib.FindAndRetrieveBlock(fakeHash, fakeHash)

				assert.Nil(t, err)
				assert.Nil(t, data)
				assert.False(t, found)
			})
		})
	})
}

/*

function: Uninstall
purpose: uninstall a downloaded/downloading game

Assumptions:
#1 SetupDownload works correctly and creates all dummy files

? Test cases
success
	#1 => download stopped and files removed

failure
	illegal arguments
			#1 => game not owned
			#2 => download not started

	unexpected err
			#1 => download started but files not found

*/

func TestUninstallGame(t *testing.T) {
	lib, testGame := setupTestLibrary(t)
	if testGame.Download != nil {
		// ! idk why this is needed and i can't be bothered to find out
		testGame.Download = nil
	}

	t.Run("success", func(t *testing.T) {
		t.Run("downloaded started and files found", func(t *testing.T) {
			err := lib.CreateDownload(testGame)
			if err != nil {
				t.Fatal(err)
			}

			location := testGame.Download.AbsolutePath
			err = lib.Uninstall(testGame.RootHash)
			assert.Nil(t, err)
			assert.Nil(t, testGame.Download)

			_, err = os.Stat(location)
			assert.NotNil(t, err)
			assert.ErrorIs(t, err, os.ErrNotExist)
		})
	})

	t.Run("failure", func(t *testing.T) {
		t.Run("illegal arguments", func(t *testing.T) {
			t.Run("game not owned", func(t *testing.T) {
				var empty [32]byte
				err := lib.Uninstall(empty)
				assert.NotNil(t, err)
				assert.Equal(t, fmt.Sprintf("game %x not owned", empty), err.Error())
			})

			t.Run("download not started", func(t *testing.T) {
				err := lib.Uninstall(testGame.RootHash)
				assert.NotNil(t, err)
				assert.Equal(t, fmt.Sprintf("Download not found for game %x", testGame.RootHash), err.Error())
			})

		})

		t.Run("unexpected err", func(t *testing.T) {
			t.Run("download started but files not found", func(t *testing.T) {
				err := lib.CreateDownload(testGame)
				assert.Nil(t, err)

				location := testGame.Download.AbsolutePath
				err = os.RemoveAll(location)
				if err != nil {
					t.Fatal(err)
				}

				err = lib.Uninstall(testGame.RootHash)
				assert.NotNil(t, err)
				assert.ErrorIs(t, err, os.ErrNotExist)
			})
		})
	})
}
