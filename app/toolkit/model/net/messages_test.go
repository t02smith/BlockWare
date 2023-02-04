package net

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

	"github.com/t02smith/part-iii-project/toolkit/model/games"
	"github.com/t02smith/part-iii-project/toolkit/test/testutil"
)

func TestGameListToMessage(t *testing.T) {

	var h1 [32]byte
	copy(h1[:], []byte("test"))

	var h2 [32]byte
	copy(h2[:], []byte("tester hash"))

	games := []*games.Game{
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

	res, err := gameListToMessage(games)
	if err != nil {
		t.Error(err)
		return
	}

	res = res[:len(res)-1]
	parts := strings.Split(res, ";")
	if parts[0] != "GAMES" {
		t.Error("Wrong command")
	}

	for i, g := range games {
		if serialised, err := g.Serialise(); err != nil || serialised != parts[i+1] {
			t.Errorf("Incorrect serialised game in pos %d", i)
		}
	}
}

func TestOnMessage(t *testing.T) {

	mp, err := testutil.StartMockPeer(7887)
	if err != nil {
		t.Fatalf("Failed to start mock peer %s", err)
	}

	t.Cleanup(func() {
		mp.Clear()
		testutil.ClearTmp("../../")
		testutil.SetupTmp("../../")
	})

	t.Run("LIBRARY", func(t *testing.T) {
		var h [32]byte
		copy(h[:], []byte("hello there"))

		datetime := time.Date(2002, 01, 10, 0, 0, 0, 0, time.Local).String()
		fakeGame := &games.Game{
			Title:       "fake game",
			Version:     "2.7.5",
			ReleaseDate: datetime,
			Developer:   "t02smith.com",
			RootHash:    h,
		}

		s, err := fakeGame.Serialise()
		if err != nil {
			t.Fatal(err)
		}

		mockPeer.SetResponse("LIBRARY\n", fmt.Sprintf("GAMES;%s\n", s))
		mockPeerClient.SendString("LIBRARY\n")
		time.Sleep(25 * time.Millisecond)

		if pd, ok := testPeer.peers[mockPeerClient]; ok {

			if len(pd.Library) == 0 {
				t.Fatal("Games not stored")
			}

			g := pd.Library[0]
			if !g.Equals(fakeGame) {
				t.Fatal("Games not equal")
			}

			return
		}

		t.Fatal("Game not stored in peer's library")
	})

	// SETUP TEST GAME
	l := games.NewLibrary()
	g, err := fetchTestGame()
	if err != nil {
		t.Fatalf("Error generating test download %s", err)
	}
	l.AddGame(g)

	t.Run("BLOCK", func(t *testing.T) {

		gData, _ := g.GetData()
		blockHash := gData.RootDir.Files["architecture-diagram.png"].Hashes[1]

		t.Run("block exists", func(t *testing.T) {
			mockPeer.SendStringAndWait(25, "BLOCK;%x;%x\n", g.RootHash, blockHash)

			msg := mockPeer.GetLastMessage()
			msg = msg[:len(msg)-1]

			res := strings.Split(msg, ";")
			if len(res) != 4 {
				t.Fatalf("incorrect response received '%s'", msg)
			}

			data, err := hex.DecodeString(res[3])
			if err != nil {
				t.Fatal(err)
			}

			hash := sha256.Sum256(data)
			if !bytes.Equal(hash[:], blockHash[:]) {
				t.Fatalf("block's hashes do not match\ngot: %x\nexpected: %x", hash, blockHash)
			}
		})

		t.Run("block doesn't exist", func(t *testing.T) {
			mockPeer.SendStringAndWait(25, "BLOCK;%x;%x\n", g.RootHash, [32]byte{})

			msg := mockPeer.GetLastMessage()
			if msg != fmt.Sprintf("ERROR;Block %x not found\n", [32]byte{}) {
				t.Fatalf("Expected an error message from peer")
			}
		})

	})

	t.Run("SEND_BLOCK", func(t *testing.T) {

		t.Run("success", func(t *testing.T) {

			err := l.CreateDownload(g)
			if err != nil {
				t.Fatal(err)
			}

			// * SETUP
			gameData, err := g.GetData()
			if err != nil {
				t.Fatal(err)
			}

			t.Run("single shard", func(t *testing.T) {
				shard := gameData.RootDir.Files["architecture-diagram.png"].Hashes[2]
				found, data, err := g.FetchShard(shard)
				if err != nil {
					t.Fatal(err)
				}

				if !found {
					t.Fatal("Test shard not found")
				}

				mockPeer.SetResponse(
					fmt.Sprintf("BLOCK;%x;%x\n", g.RootHash, shard),
					fmt.Sprintf("SEND_BLOCK;%x;%x;%x\n", g.RootHash, shard, data),
				)

				mockPeerClient.SendStringf("BLOCK;%x;%x\n", g.RootHash, shard)
				time.Sleep(25 * time.Millisecond)

				// ? was the shard inserted
				buffer := make([]byte, gameData.ShardSize)

				f, err := os.Open("../../test/data/tmp/toolkit/architecture-diagram.png")
				if err != nil {
					t.Errorf("Unable to verify whether shard was inserted %s", err)
				}
				defer f.Close()

				_, err = f.Seek(int64(gameData.ShardSize)*int64(2), 0)
				if err != nil {
					t.Errorf("Unable to verify whether shard was inserted %s", err)
				}

				reader := bufio.NewReader(f)
				reader.Read(buffer)
				if !bytes.Equal(buffer, data) {
					t.Fatalf("incorrect shard inserted\ngot %x\nexpected %x", buffer, data)
				}
			})

			t.Run("entire file", func(t *testing.T) {
				// TODO bug where last shard of file is not being hashed correctly

				file := gameData.RootDir.Subdirs["subdir"].Files["chip8.c"]
				buffer := make([]byte, gameData.ShardSize)

				for i, shard := range file.Hashes {
					found, data, err := g.FetchShard(shard)
					if err != nil {
						t.Fatal(err)
					}

					if !found {
						t.Fatal("Test shard not found")
					}

					mockPeer.SetResponse(
						fmt.Sprintf("BLOCK;%x;%x\n", g.RootHash, shard),
						fmt.Sprintf("SEND_BLOCK;%x;%x;%x\n", g.RootHash, shard, data),
					)

					mockPeerClient.SendStringf("BLOCK;%x;%x\n", g.RootHash, shard)
					time.Sleep(10 * time.Millisecond)

					// ? was the shard inserted
					f, err := os.Open("../../test/data/tmp/toolkit/subdir/chip8.c")
					if err != nil {
						t.Errorf("Unable to verify whether shard was inserted %s", err)
					}

					_, err = f.Seek(int64(gameData.ShardSize)*int64(i), 0)
					if err != nil {
						t.Errorf("Unable to verify whether shard was inserted %s", err)
					}

					reader := bufio.NewReader(f)
					reader.Read(buffer)
					if !bytes.Equal(buffer, data) {
						t.Errorf("incorrect shard inserted\ngot %x\nexpected %x", buffer, data)
					}

					f.Close()
				}
			})
		})

		t.Run("failure", func(t *testing.T) {

			t.Run("sending incorrect data", func(t *testing.T) {

			})

		})

	})

	mp.Close()
}

// util functions

func TestFetchBlock(t *testing.T) {

	t.Run("game doesn't exist", func(t *testing.T) {
		_, err := fetchBlock([32]byte{}, [32]byte{})
		if err == nil {
			t.Error("Missing game not identified as not existing")
		}
	})

	p := GetPeerInstance()
	g, err := fetchTestGame()
	if err != nil {
		t.Fatal(err)
	}

	p.library.AddGame(g)

	t.Run("game exists but block does not", func(t *testing.T) {
		_, err = fetchBlock(g.RootHash, [32]byte{})
		if err == nil {
			t.Error("Block should not have been identified")
		}
	})

	t.Run("success", func(t *testing.T) {
		ht, err := g.GetData()
		if err != nil {
			t.Fatal(err)
		}

		hash := ht.RootDir.Files["architecture-diagram.png"].Hashes[1]
		data, err := fetchBlock(g.RootHash, hash)
		if err != nil {
			t.Fatal(err)
		}

		dataHash := sha256.Sum256(data)
		if !bytes.Equal(dataHash[:], hash[:]) {
			t.Fatal("incorrect block fetched")
		}
	})

}
