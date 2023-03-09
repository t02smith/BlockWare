package peer

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/t02smith/part-iii-project/toolkit/model/manager/games"
	hash "github.com/t02smith/part-iii-project/toolkit/model/manager/hashtree"
	"github.com/t02smith/part-iii-project/toolkit/test/testutil"
	testenv "github.com/t02smith/part-iii-project/toolkit/test/testutil/env"
)

/*

function: generateLIBRARY
purpose: generate a message requesting a users library

? test cases
success
	| #1 => function returns a constnat

*/

func TestGenerateLIBRARY(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		assert.Equal(t, "LIBRARY\n", generateLIBRARY(), "smh its a constant change the value")
	})
}

/*

function: handleLIBRARY
purpose: respond to an incoming library message with your games list

? test cases
success
	| #1 => user has no games
	| #2 => user has some games

failure
TODO

*/

func TestHandleLIBRARY(t *testing.T) {
	mp, _ := createMockPeer(t)

	t.Run("success", func(t *testing.T) {

		t.Run("no games", func(t *testing.T) {
			Peer().library.ClearOwnedGames()

			mp.SendStringAndWait(25, generateLIBRARY())
			res := mp.GetLastMessage()

			assert.Equal(t, "GAMES;\n", res, "invalid GAMES message received")
		})

		t.Run("some games", func(t *testing.T) {
			var rh [32]byte
			copy(rh[:], []byte("hello there"))

			Peer().library.AddOwnedGame(&games.Game{
				RootHash: rh,
			})

			mp.SendStringAndWait(25, generateLIBRARY())
			res := mp.GetLastMessage()

			assert.Equal(t, fmt.Sprintf("GAMES;%x\n", rh), res, "invalid GAMES message")
		})
	})

}

/*

function: generateGAMES
purpose: generates a message containing a series of games given to it

? test cases
success
	| #1 => no games
	| #2 => unique games
	| #3 => duplicate games

failure
TODO
*/

func TestGenerateGames(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		t.Run("no games", func(t *testing.T) {
			res := generateGAMES()
			assert.Equal(t, "GAMES;\n", res, "invalid message format")
		})

		// * setup

		var h1 [32]byte
		copy(h1[:], []byte("test"))

		var h2 [32]byte
		copy(h2[:], []byte("tester hash"))

		gameLS := []*games.Game{
			{
				Title:       "Test Game",
				Version:     "1.0.2",
				ReleaseDate: time.Now().String(),
				Developer:   "tcs1g20",
				RootHash:    h1,
			},
			{
				Title:       "Borderlands 3",
				Version:     "3.4.17",
				ReleaseDate: time.Now().String(),
				Developer:   "Gearbox",
				RootHash:    h2,
			},
		}

		checkOutput := func(res string, gs []*games.Game) bool {
			res = res[:len(res)-1]
			parts := strings.Split(res, ";")
			if parts[0] != "GAMES" {
				t.Error("Wrong command")
			}

			for i, g := range gs {
				if fmt.Sprintf("%x", g.RootHash) != parts[i+1] {
					return false
				}
			}
			return true
		}

		//

		t.Run("unique games", func(t *testing.T) {
			res := generateGAMES(gameLS...)
			assert.Truef(t, checkOutput(res, gameLS), "invalid message generated %s", res)
		})

		t.Run("duplicate games", func(t *testing.T) {
			gameLS = append(gameLS, gameLS[0])
			res := generateGAMES(gameLS...)
			assert.Truef(t, checkOutput(res, gameLS[:2]), "invalid message generated %s", res)
		})
	})

}

/*

function: handleGAMES
purpose: react to an incoming GAMES message by storing a peer's game data

? Test cases
success
	| #1 => no games
	| #2 => known games
	| #3 => filter invalid hashes

failure
TODO
*/

func TestHandleGAMES(t *testing.T) {
	mp, it := createMockPeer(t)
	data := Peer().GetPeer(it)
	if data == nil {
		t.Fatalf("mock peer not created properly")
	}

	// * test game
	var h1 [32]byte
	copy(h1[:], []byte("test"))
	game := &games.Game{
		Title:       "Test Game",
		Version:     "1.0.2",
		ReleaseDate: time.Now().String(),
		Developer:   "tcs1g20",
		RootHash:    h1,
	}

	t.Run("success", func(t *testing.T) {

		t.Run("no games", func(t *testing.T) {
			mp.SendStringAndWait(25, generateGAMES())
			assert.Zero(t, len(data.Library), "no games expected to be known about")
		})

		t.Run("known games", func(t *testing.T) {
			t.Cleanup(func() {
				Peer().library.ClearOwnedGames()
				data.Library = make(map[[32]byte]ownership)
			})

			Peer().library.AddOwnedGame(game)
			mp.SendStringAndWait(25, generateGAMES(game))

			assert.Equal(t, 1, len(data.Library), "Game not recognised")

			hasGame, ok := data.Library[game.RootHash]
			assert.True(t, ok, "game not found in peers collection")
			assert.Equal(t, unknown, hasGame, "game ownership not stored correctly => should be unknown initially")

		})

		t.Run("invalid hash filtered", func(t *testing.T) {
			mp.SendStringAndWait(25, "GAMES;127382231039019392738\n")
			assert.Zero(t, len(data.Library), "game should not have been added")
		})

	})
}

/*

function: generateBLOCK
purpose: create a request message for a block

? Test cases
success
	| #1 => base case

*/

func TestGenerateBLOCK(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		gh, rh := sha256.Sum256([]byte("game hash")), sha256.Sum256([]byte("block hash"))

		res := generateBLOCK(gh, rh)
		assert.Equal(t, fmt.Sprintf("BLOCK;%x;%x\n", gh, rh), res, "incorrect message receieved")

	})
}

/*

function: handleBLOCK
purpose: respond to an incoming block message by attempting to find and send the data

? Test cases
success
	| #1 correct block data sent successfully

failure
	| illegal arguments
			| #1 => invalid game hash
			| #2 => invalid block hash
			| #3 => user doesn't own game

*/

func TestHandleBLOCK(t *testing.T) {
	mp, _ := createMockPeer(t)

	game := testenv.CreateTestGame(t, "../../../")
	err := Peer().library.AddOwnedGame(game)
	if err != nil {
		t.Fatalf("Error adding game to library: %s", err)
	}

	t.Cleanup(func() {
		Peer().library.ClearOwnedGames()

	})

	t.Run("failure", func(t *testing.T) {
		t.Run("illegal arguments", func(t *testing.T) {
			t.Run("game hash", func(t *testing.T) {
				mp.SendStringAndWait(50, "BLOCK;2612732183721371;%x\n", sha256.Sum256([]byte("test")))

				res := mp.GetLastMessage()
				assert.Equal(t,
					"ERROR;error reading game hash on BLOCK cmd: invalid hash length for hash 2612732183721371\n",
					res,
					"invalid err message")
			})

			t.Run("block hash", func(t *testing.T) {
				mp.SendStringAndWait(50, "BLOCK;%x;7236173612213\n", game.RootHash)

				res := mp.GetLastMessage()
				assert.Equal(t,
					"ERROR;error reading shard hash on BLOCK cmd: invalid hash length for hash 7236173612213\n",
					res,
					"invalid err message")
			})

			t.Run("user doesn't own game", func(t *testing.T) {
				Peer().library.ClearOwnedGames()
				t.Cleanup(func() {
					err = Peer().library.AddOwnedGame(game)
					if err != nil {
						t.Fatalf("Error during cleanup adding game to library: %s", err)
					}
				})

				mp.SendStringAndWait(50, "BLOCK;%x;%x\n", game.RootHash, game.RootHash)
				res := mp.GetLastMessage()
				assert.Equal(t,
					fmt.Sprintf("ERROR;block %x not found\n", game.RootHash),
					res,
					"invalid err message")
			})
		})
	})

	t.Run("success", func(t *testing.T) {
		gameData, err := game.GetData()
		if err != nil {
			t.Fatal(err)
		}

		t.Run("correct block received", func(t *testing.T) {
			blockHash := gameData.RootDir.Files["architecture-diagram.png"].Hashes[1]

			mp.SendStringAndWait(50, generateBLOCK(game.RootHash, blockHash))
			msg := mp.GetLastMessage()
			msg = msg[:len(msg)-1]

			res := strings.Split(msg, ";")
			assert.Equal(t, 4, len(res), "invalid number of sections in response")

			data, err := hex.DecodeString(res[3])
			assert.Nil(t, err, err)

			hash := sha256.Sum256(data)
			assert.Equal(t, blockHash[:], hash[:], "hashes do not match => incorrect data sent")
		})

	})

}

/*

function: generateSEND_BLOCK
purpose: create a SEND_BLOCK message to send data to other users

? Test cases
success
	| #1 => base case (function is just a string formatting)

*/

func TestGenerateSEND_BLOCK(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		data := []byte("sajhdasgudijsaoidhyqwgasjhdyasgdhasbidhauiwhdushuidajhuigdwhudhuihasudhwudhaws")
		gh, sh := sha256.Sum256([]byte("fake-game")), sha256.Sum256(data)

		res := generateSEND_BLOCK(gh, sh, data)
		assert.Equal(t,
			fmt.Sprintf("SEND_BLOCK;%x;%x;%x\n", gh, sh, data),
			res,
			"incorrect SEND_BLOCK message generated",
		)
	})
}

/*

function: handleSEND_BLOCK
purpose: receive and insert a block of data from storage

? Test cases
success
	| #1 => single shard
	| #2 => whole file

failure
	| illegal arguments
			| #1 => invalid game hash
			| #2 => invalid shard hash
			| invalid data
					| #1 => wrong length
					| #2 => does not match hash
	| unexpected data
			| #1 => download not started
			| #2 => block not needed

*/

func TestHandleSEND_BLOCK(t *testing.T) {
	mp, _ := createMockPeer(t)

	game := testenv.CreateTestGame(t, "../../../")
	err := Peer().library.AddOwnedGame(game)
	if err != nil {
		t.Fatalf("Error adding game to library: %s", err)
	}

	t.Cleanup(func() {
		Peer().library.ClearOwnedGames()
	})

	t.Run("failure", func(t *testing.T) {
		t.Run("illegal arguments", func(t *testing.T) {
			t.Run("invalid game hash", func(t *testing.T) {
				mp.SendStringAndWait(25,
					"SEND_BLOCK;12313123124123;%x;%x\n", sha256.Sum256([]byte("hello")), []byte("shashduiashid"))

				res := mp.GetLastMessage()
				assert.Equal(t,
					"ERROR;error reading game hash on BLOCK cmd: invalid hash length for hash 12313123124123\n",
					res,
					"invalid err message",
				)
			})

			t.Run("invalid shard hash", func(t *testing.T) {
				mp.SendStringAndWait(25, "SEND_BLOCK;%x;29381209382;%x\n", game.RootHash, []byte("shashduiashid"))

				res := mp.GetLastMessage()
				assert.Equal(t,
					"ERROR;error reading shard hash on BLOCK cmd: invalid hash length for hash 29381209382\n",
					res,
					"invalid err message",
				)
			})

			t.Run("invalid data", func(t *testing.T) {
				testenv.SetupTestDownload(t, game, "../../../")

				gData, err := game.GetData()
				if err != nil {
					t.Fatal(err)
				}

				blockHash := gData.RootDir.Files["architecture-diagram.png"].Hashes[1]

				t.Run("invalid length", func(t *testing.T) {
					mp.SendStringAndWait(25, generateSEND_BLOCK(game.RootHash, blockHash, []byte("123")))

					res := mp.GetLastMessage()
					assert.Equal(t,
						fmt.Sprintf("ERROR;incorrect data length. expected %d, got %d\n", gData.ShardSize, len([]byte("123"))),
						res,
						"invalid err msessage",
					)
				})

				t.Run("doesn't match hash", func(t *testing.T) {
					data := make([]byte, gData.ShardSize)
					copy(data[:], []byte("hello world"))

					mp.SendStringAndWait(25, generateSEND_BLOCK(game.RootHash, blockHash, data))

					res := mp.GetLastMessage()
					assert.Equal(t,
						fmt.Sprintf("ERROR;block %x data does not match expected content\n", blockHash),
						res,
						"invalid err message",
					)
				})
			})
		})

		t.Run("unexpected data", func(t *testing.T) {
			gData, err := game.GetData()
			if err != nil {
				t.Fatal(err)
			}

			blockHash := gData.RootDir.Files["architecture-diagram.png"].Hashes[1]

			t.Run("download not started", func(t *testing.T) {
				mp.SendStringAndWait(25, generateSEND_BLOCK(game.RootHash, blockHash, []byte("hello")))

				res := mp.GetLastMessage()
				assert.Equal(t,
					fmt.Sprintf("ERROR;download not found for %x\n", game.RootHash),
					res,
					"invalid err message",
				)
			})

			testenv.SetupTestDownload(t, game, "../../../")

			t.Run("block not needed", func(t *testing.T) {
				fileHash := gData.RootDir.Files["architecture-diagram.png"].RootHash
				delete(game.Download.Progress[fileHash].BlocksRemaining, blockHash)

				data := make([]byte, gData.ShardSize)
				mp.SendStringAndWait(25, generateSEND_BLOCK(game.RootHash, blockHash, data))

				res := mp.GetLastMessage()
				assert.Equal(t,
					fmt.Sprintf("ERROR;block %x not in download queue\n", blockHash),
					res,
					"invalid err message",
				)
			})
		})
	})

	t.Run("success", func(t *testing.T) {
		game := testenv.CreateTestGame(t, "../../../")
		testenv.SetupTestDownload(t, game, "../../../")

		Peer().library.AddOwnedGame(game)
		t.Cleanup(func() {
			Peer().library.ClearOwnedGames()
		})

		go func() {
			for range Peer().library.DownloadManager.DownloadProgress {
				continue
			}
		}()

		gData, err := game.GetData()
		if err != nil {
			t.Fatal(err)
		}

		smoke := t.Run("single shard", func(t *testing.T) {
			shardHash := gData.RootDir.Files["architecture-diagram.png"].Hashes[1]
			buffer := make([]byte, gData.ShardSize)

			_sendBlock(t, "../../../test/data/tmp/toolkit/architecture-diagram.png", mp, game, gData, shardHash, 1, buffer)

		})

		if !smoke {
			t.FailNow()
		}

		t.Run("whole file", func(t *testing.T) {
			file := gData.RootDir.Subdirs["subdir"].Files["chip8.c"]
			buffer := make([]byte, gData.ShardSize)

			for i, shard := range file.Hashes {
				_sendBlock(t, "../../../test/data/tmp/toolkit/subdir/chip8.c", mp, game, gData, shard, uint(i), buffer)
				time.Sleep(25 * time.Millisecond)
			}
		})
	})
}

// test whether a user stores a required block after being sent it
func _sendBlock(t *testing.T, filename string, mp *testutil.MockPeer, g *games.Game, gData *hash.HashTree, blockHash [32]byte, offset uint, buffer []byte) {
	t.Helper()

	found, data, err := g.FetchShard(blockHash)
	if err != nil {
		t.Fatal(err)
	}

	if !found {
		t.Fatalf("block %x not found", blockHash)
	}

	mp.SendStringAndWait(25, generateSEND_BLOCK(g.RootHash, blockHash, data))

	// ? was the shard inserted
	f, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Unable to verify whether shard was inserted %s", err)
	}
	defer f.Close()

	_, err = f.Seek(int64(gData.ShardSize)*int64(offset), 0)
	if err != nil {
		t.Fatalf("Unable to verify whether shard was inserted %s", err)
	}

	// ? clear buffer
	for i := range buffer {
		buffer[i] = 0x00
	}

	reader := bufio.NewReader(f)
	_, err = reader.Read(buffer)
	if err != nil {
		t.Fatalf("error reading shard %s", err)
	}

	if !bytes.Equal(buffer, data) {
		t.Fatalf("incorrect shard inserted\ngot %x\nexpected %x", buffer, data)
	}
}

/*

function: generateVALIDATE_REQ
purpose: build a structured message from id verification

? Test cases
success
	| #1 base case

*/

func TestGnerateVALIDATE_REQ(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		t.Run("base case", func(t *testing.T) {
			msg := []byte("hello world - blockware")
			res := generateVALIDATE_REQ(msg)

			assert.Equal(t, fmt.Sprintf("VALIDATE_REQ;%x\n", msg), res, "invalid message")
		})
	})
}
