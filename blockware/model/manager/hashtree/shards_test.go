package hash

import (
	"bytes"
	"crypto/sha256"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/t02smith/part-iii-project/toolkit/test/testutil"
)

func produceTestHashTree() (*HashTree, error) {
	ht, err := NewHashTree("../../../test/data/testdir", 1024, nil)
	if err != nil {
		return nil, err
	}

	err = ht.Hash()
	if err != nil {
		return nil, err
	}

	return ht, nil
}

func TestFindShard(t *testing.T) {
	testutil.ShortTest(t)

	ht, err := produceTestHashTree()
	if err != nil {
		t.Fatalf("error making test hash tree %s", err)
	}

	t.Run("valid shard", func(t *testing.T) {
		file := ht.RootDir.Files["test.txt"]
		locations := ht.FindShard(file.Hashes[0])

		assert.NotEmpty(t, locations)
		assert.Equal(t, file.Filename, locations[0].AbsolutePath, "was %s", locations[0].AbsolutePath)
		assert.Zero(t, locations[0].Offset)
	})

	t.Run("invalid shard", func(t *testing.T) {
		locations := ht.FindShard([32]byte{})
		assert.Empty(t, locations)
	})

}

// Test getting a shard

func TestGetShard(t *testing.T) {
	testutil.ShortTest(t)

	ht, err := produceTestHashTree()
	if err != nil {
		t.Fatalf("error making test hash tree %s", err)
	}

	file := ht.RootDir.Files["test.txt"]

	t.Run("shard gotten successfully", func(t *testing.T) {
		found, data, err := ht.GetShard(file.Hashes[0])
		if err != nil {
			t.Fatalf("failed to get shard %s", err)
		}

		if !found {
			t.Error("found should be true")
		}

		hash := sha256.Sum256(data)
		if !bytes.Equal(file.Hashes[0][:], hash[:]) {
			t.Error("Incorrect shard fetched")
		}
	})

	t.Run("invalid shard given", func(t *testing.T) {

		found, _, _ := ht.GetShard([32]byte{})
		if found {
			t.Fatalf("invalid shard not detected")
		}

	})
}
