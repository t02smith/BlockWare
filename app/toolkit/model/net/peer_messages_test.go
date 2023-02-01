package net

import (
	"bytes"
	"crypto/sha256"
	"fmt"
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

		mockPeer.SetResponse("LIBRARY\n", fmt.Sprintf("GAMES;%s;\n", s))
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

	t.Run("BLOCK", func(t *testing.T) {
		// TODO

		t.Run("block exists", func(t *testing.T) {

		})

		t.Run("block doesn't exist", func(t *testing.T) {

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
