package net

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/t02smith/part-iii-project/toolkit/lib/games"
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

func TestOnMessageLibrary(t *testing.T) {
	beforeEach()

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
		t.Error(err)
		return
	}

	mockPeer.SetResponse("LIBRARY\n", fmt.Sprintf("GAMES;%s;\n", s))
	mockPeerClient.SendString("LIBRARY\n")
	time.Sleep(25 * time.Millisecond)

	if pd, ok := testPeer.peers[mockPeerClient]; ok {

		if len(pd.Library) == 0 {
			t.Error("Games not stored")
			return
		}

		g := pd.Library[0]

		fmt.Println(g)

		if !g.Equals(fakeGame) {
			t.Error("Games not equal")
			return
		}
	} else {
		t.Error("Game not stored in peer's library")
	}
}

func TestOnMessageBlockWhenBlockExists(t *testing.T) {
	t.Skip()
	beforeEach()

	testHash := sha256.Sum256([]byte("1234567890"))
	message := fmt.Sprintf("BLOCK;%x;\n", testHash)

	data := make([]byte, 8192)
	for i := range data {
		data[i] = 1
	}

	mockPeer.SetResponse(message, "SEND_BLOCK;%x;\n")
	mockPeer.SetOnReceive(message, func() {
		mockPeer.SendBytes(data)
	})

	mockPeerClient.SendString(message)
	time.Sleep(25 * time.Millisecond)
}

func TestOnMessageBlockWhenBlockDoesntExist(t *testing.T) {

}

// util functions

func TestFetchBlock(t *testing.T) {
	beforeEach()

	// ! game doesn't exist
	_, err := fetchBlock([32]byte{}, [32]byte{})
	if err == nil {
		t.Error("Missing game not identified as not existing")
	}

	// ! game exists but block does not

	p := GetPeerInstance()
	g, err := fetchTestGame()
	if err != nil {
		t.Error(err)
		return
	}

	p.library.AddGame(g)

	_, err = fetchBlock(g.RootHash, [32]byte{})
	if err == nil {
		t.Error("Block should not have been identified")
	}

	// ! game and block exists and is fetched successfully
	ht, err := g.GetData()
	if err != nil {
		t.Error(err)
		return
	}

	hash := ht.RootDir.Files["architecture-diagram.png"].Hashes[1]
	data, err := fetchBlock(g.RootHash, hash)
	if err != nil {
		t.Error(err)
		return
	}

	dataHash := sha256.Sum256(data)
	if !bytes.Equal(dataHash[:], hash[:]) {
		t.Error("incorrect block fetched")
	}
}
