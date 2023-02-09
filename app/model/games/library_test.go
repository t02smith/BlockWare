package games

import (
	"bytes"
	"crypto/sha256"
	"testing"

	"github.com/t02smith/part-iii-project/toolkit/test/testutil"
)

func TestGetGame(t *testing.T) {
	l := NewLibrary()
	g, err := setupTestDownload()
	if err != nil {
		t.Fatalf("Error generating test download %s", err)
	}

	t.Run("game not found", func(t *testing.T) {
		g1 := l.GetGame(g.RootHash)
		if g1 != nil {
			t.Error("expected nil pointers for missing game & download")
		}
	})

	t.Run("success", func(t *testing.T) {
		l.AddGame(g)
		g1 := l.GetGame(g.RootHash)
		if g1 == nil {
			t.Error("Incorrect GetGame after adding single game")
		}
	})

	testutil.ClearTmp("../../")
	testutil.SetupTmp("../../")

}

func TestAddGame(t *testing.T) {
	l := NewLibrary()
	g, err := setupTestDownload()
	if err != nil {
		t.Fatalf("Error generating test download %s", err)
		return
	}

	t.Run("success", func(t *testing.T) {
		l.AddGame(g)
		if _, ok := l.Games[g.RootHash]; !ok {
			t.Errorf("Game not added")
		}
	})

	testutil.ClearTmp("../../")
	testutil.SetupTmp("../../")

}

func TestLibraryFindBlock(t *testing.T) {

	l := NewLibrary()
	g, err := setupTestDownload()
	if err != nil {
		t.Fatalf("Error generating test download %s", err)
	}
	l.AddGame(g)

	t.Run("success", func(t *testing.T) {
		shardHash := g.data.RootDir.Files["test.txt"].RootHash

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