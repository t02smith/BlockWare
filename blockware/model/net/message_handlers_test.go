package net

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/t02smith/part-iii-project/toolkit/model/games"
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

? test cases
success
	| #1 => no games
	| #2 => some games
	| #3 => some games with unknowns

failure
	| illegal arguments
			| #1 => invalid root hash

*/

func TestHandleGAMES(t *testing.T) {
	mp, _ := createMockPeer(t)
	_ = mp

	t.Run("success", func(t *testing.T) {
		t.Run("no games", func(t *testing.T) {

		})

		t.Run("known games", func(t *testing.T) {

		})

		t.Run("unknown games", func(t *testing.T) {

		})
	})
}
