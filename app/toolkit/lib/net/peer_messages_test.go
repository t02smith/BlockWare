package net

import (
	"crypto/sha256"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/t02smith/part-iii-project/toolkit/lib/games"
)

func TestGameListToMessage(t *testing.T) {
	games := []*games.Game{
		{
			Title:       "Test Game",
			Version:     "1.0.2",
			ReleaseDate: time.Now().String(),
			Developer:   "tcs1g20",
			RootHash:    []byte("test"),
		},
		{
			Title:       "Borderlands 3",
			Version:     "3.4.17",
			ReleaseDate: time.Now().String(),
			Developer:   "Gearbox",
			RootHash:    []byte("tester hash"),
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

	datetime := time.Date(2002, 01, 10, 0, 0, 0, 0, time.Local).String()
	fakeGame := &games.Game{
		Title:       "fake game",
		Version:     "2.7.5",
		ReleaseDate: datetime,
		Developer:   "t02smith.com",
		RootHash:    []byte("hello there"),
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

		if !games.GamesAreEqual(fakeGame, g) {
			t.Error("Games not equal")
			return
		}
	} else {
		t.Error("Game not stored in peer's library")
	}
}

func TestOnMessageBlock(t *testing.T) {
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
