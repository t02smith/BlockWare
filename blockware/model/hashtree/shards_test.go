package hash

import (
	"bytes"
	"crypto/sha256"
	"testing"

	"github.com/t02smith/part-iii-project/toolkit/test/testutil"
)

func produceTestHashTree() (*HashTree, error) {
	ht, err := NewHashTree("../../test/data/testdir", 1024, nil)
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
		found, _, filename, offset := ht.FindShard(file.Hashes[0])

		if !found {
			t.Error("Existing shard not found")
		}

		if filename != file.Filename {
			t.Errorf("Incorrect filepath returned %s", filename)
		}

		if offset != 0 {
			t.Error("Incorrect offset")
		}
	})

	t.Run("invalid shard", func(t *testing.T) {
		found, _, _, _ := ht.FindShard([32]byte{})
		if found {
			t.Errorf("Shard does not exist but states it is found")
		}
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
