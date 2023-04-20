package library

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/t02smith/part-iii-project/toolkit/build/contracts/library"
	"github.com/t02smith/part-iii-project/toolkit/model/persistence/ethereum"
)

/*

Tests for "library.sol" smart contract
! make sure to restart the ganache server with "make ganache"
! before every rerun otherwise you will get a txn error about
! an incorrect nonce

*/

// utility

// compare two games => one uploaded and one fetched (hence the struct)
func assertGamesEqual(t *testing.T, g1 library.LibraryGameEntry, g2 struct {
	Title               string
	Version             string
	ReleaseDate         string
	Developer           string
	RootHash            [32]byte
	PreviousVersion     [32]byte
	NextVersion         [32]byte
	Price               *big.Int
	Uploader            common.Address
	HashTreeIPFSAddress string
	AssetsIPFSAddress   string
}) {
	assert.True(t, bytes.Equal(g1.RootHash[:], g2.RootHash[:]))
	assert.True(t, bytes.Equal(g1.PreviousVersion[:], g2.PreviousVersion[:]))
	assert.True(t, bytes.Equal(g1.NextVersion[:], g2.NextVersion[:]))
	assert.Equal(t, g1.Title, g2.Title)
	assert.Equal(t, g1.Version, g2.Version)
	assert.Equal(t, g1.ReleaseDate, g2.ReleaseDate)
	assert.Equal(t, g1.Developer, g2.Developer)
	assert.Zero(t, g1.Price.Cmp(g2.Price))
	assert.Equal(t, g1.Uploader, g2.Uploader)
	assert.Equal(t, g1.AssetsIPFSAddress, g2.AssetsIPFSAddress)
	assert.Equal(t, g1.HashTreeIPFSAddress, g2.HashTreeIPFSAddress)
}

var empty [32]byte

var gameI library.LibraryGameEntry = library.LibraryGameEntry{
	Title:               "game I",
	Version:             "1.7",
	ReleaseDate:         time.Date(2002, time.January, 10, 0, 0, 0, 0, time.UTC).String(),
	Developer:           "tcs1g20",
	Price:               big.NewInt(1),
	Uploader:            common.Address{},
	RootHash:            sha256.Sum256([]byte("hello")),
	AssetsIPFSAddress:   "hello world",
	HashTreeIPFSAddress: "testing",
	PreviousVersion:     empty,
	NextVersion:         empty,
}

var gameII library.LibraryGameEntry = library.LibraryGameEntry{
	Title:               "game I",
	Version:             "2.1",
	ReleaseDate:         time.Date(2002, time.January, 10, 0, 0, 0, 0, time.UTC).String(),
	Developer:           "tcs1g20",
	Price:               big.NewInt(1),
	Uploader:            common.Address{},
	RootHash:            sha256.Sum256([]byte("hello world")),
	AssetsIPFSAddress:   "hello world",
	HashTreeIPFSAddress: "testing",
	PreviousVersion:     gameI.RootHash,
	NextVersion:         empty,
}

// tests

func TestLibrarySmartContract(t *testing.T) {
	ops, lib := deployContract(t)
	/*

		Smart Contract function: library.uploadGame
		Purpose: upload a game's metadata to eth

		? Test cases
		success
			for each check:
				a - game metadata present in games mapping
				b - hash present in gameHashes arr
				c - ?? event emitted

			#1 => new game uploaded
			#2 => game update uploaded

		failure
			with update
				#1 => no previous version found
				#2 => uploader not original game owner
	*/
	smoke := t.Run("Upload Game", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {

			table := []struct {
				name      string
				game      library.LibraryGameEntry
				libSize   int64
				hashIndex int64
			}{
				{"new game", gameI, 1, 0},
				{"update game", gameII, 2, 1},
			}

			for _, in := range table {
				t.Run(in.name, func(t *testing.T) {
					in.game.Uploader = ethereum.Address()

					_, err := lib.UploadGame(ops, in.game)
					if err != nil {
						t.Fatal(err)
					}

					g, err := lib.Games(nil, in.game.RootHash)
					assert.Nil(t, err)

					assertGamesEqual(t, in.game, g)

					libSize, err := lib.LibSize(nil)
					assert.Nil(t, err)
					assert.Zero(t, big.NewInt(in.libSize).Cmp(libSize))

					hashArr, err := lib.GameHashes(nil, big.NewInt(in.hashIndex))
					assert.Nil(t, err)
					assert.True(t, bytes.Equal(in.game.RootHash[:], hashArr[:]))
				})
			}

		})

		t.Run("update made changes", func(t *testing.T) {
			g, err := lib.Games(nil, gameI.RootHash)
			assert.Nil(t, err)

			assert.True(t, bytes.Equal(g.NextVersion[:], gameII.RootHash[:]), gameI)
		})
	})

	if !smoke {
		t.Fatal("games not uploaded successfully => aborting")
	}

	/*

		function: purchaseGame
		purpose: purchase a new game on the smart contract

		? Test case
		success
			#1 => purchase successful

		failure
			#1 => game not found
			#2 => user already owns game

	*/
	t.Run("Purchase Game", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			t.Run("purchase successful", func(t *testing.T) {
				authInstance.Value = gameI.Price
				_, err := lib.PurchaseGame(ops, gameI.RootHash)
				assert.Nil(t, err)

				purchased, err := lib.HasPurchased(nil, gameI.RootHash, ethereum.Address())
				assert.Nil(t, err)
				assert.True(t, purchased)
			})

			t.Run("tangential ownership", func(t *testing.T) {
				purchased, err := lib.HasPurchased(nil, gameII.RootHash, ethereum.Address())
				assert.Nil(t, err)
				assert.True(t, purchased)
			})
		})

		t.Run("failure", func(t *testing.T) {
			t.Run("game not found", func(t *testing.T) {
				_, err := lib.PurchaseGame(ops, sha256.Sum256([]byte("fake game")))
				assert.Nil(t, err)

				purchased, err := lib.HasPurchased(nil, sha256.Sum256([]byte("fake game")), ethereum.Address())
				assert.Nil(t, err)
				assert.False(t, purchased, purchased)
			})

			t.Run("game already owned", func(t *testing.T) {
				_, err := lib.PurchaseGame(ops, gameI.RootHash)
				assert.Nil(t, err)

				purchased, err := lib.HasPurchased(nil, gameI.RootHash, ethereum.Address())
				assert.Nil(t, err)
				assert.True(t, purchased)
			})
		})
	})

	t.Run("get most recent version", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			gHash, err := lib.GetMostRecentVersion(nil, gameI.RootHash)
			assert.Nil(t, err)
			fmt.Println(gHash)
			assert.True(t, bytes.Equal(gameII.RootHash[:], gHash[:]))
		})
	})
}
