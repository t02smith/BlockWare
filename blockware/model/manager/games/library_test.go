package games

import (
	"bytes"
	"crypto/sha256"
	"testing"

	"github.com/t02smith/part-iii-project/toolkit/test/testutil"
)

func TestGetGame(t *testing.T) {
	testutil.ShortTest(t)
	l := NewLibrary()
	g, err := setupTestDownload(t)
	if err != nil {
		t.Fatalf("Error generating test download %s", err)
	}

	t.Run("game not found", func(t *testing.T) {
		g1 := l.GetOwnedGame(g.RootHash)
		if g1 != nil {
			t.Error("expected nil pointers for missing game & download")
		}
	})

	t.Run("success", func(t *testing.T) {
		l.AddOwnedGame(g)
		g1 := l.GetOwnedGame(g.RootHash)
		if g1 == nil {
			t.Error("Incorrect GetGame after adding single game")
		}
	})

	testutil.ClearTmp("../../")
}

func TestAddGame(t *testing.T) {
	testutil.ShortTest(t)
	l := NewLibrary()
	g, err := setupTestDownload(t)
	if err != nil {
		t.Fatalf("Error generating test download %s", err)
		return
	}

	t.Run("success", func(t *testing.T) {
		l.AddOwnedGame(g)
		if _, ok := l.ownedGames[g.RootHash]; !ok {
			t.Errorf("Game not added")
		}
	})

	testutil.ClearTmp("../../")

}

func TestLibraryFindBlock(t *testing.T) {
	testutil.ShortTest(t)
	l := NewLibrary()
	g, err := setupTestDownload(t)
	if err != nil {
		t.Fatalf("Error generating test download %s", err)
	}
	l.AddOwnedGame(g)

	t.Run("success", func(t *testing.T) {
		shardHash := g.data.RootDir.Files["test.txt"].Hashes[0]

		found, data, err := l.FindBlock(g.RootHash, shardHash)
		if err != nil {
			t.Fatal(err)
		}

		if !found {
			t.Fatal("block should have been found")
		}

		foundHash := sha256.Sum256(data)
		if !bytes.Equal(foundHash[:], shardHash[:]) {
			t.Fatal("Incorrect block found")
		}
	})

	t.Run("not found", func(t *testing.T) {
		t.Run("game", func(t *testing.T) {
			found, _, err := l.FindBlock([32]byte{}, [32]byte{})
			if err != nil {
				t.Fatal(err)
			}

			if found {
				t.Fatal("Game should not have been found")
			}
		})

		t.Run("shard", func(t *testing.T) {
			found, _, err := l.FindBlock(g.RootHash, [32]byte{})
			if err != nil {
				t.Fatal(err)
			}

			if found {
				t.Fatal("shard should not have been found")
			}
		})

	})

}
