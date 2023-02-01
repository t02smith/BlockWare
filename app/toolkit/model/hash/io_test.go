package hash

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"os"
	"path/filepath"
	"testing"
)

func produceTestHashTree() (*HashTree, error) {
	ht, err := NewHashTree("../../test/data/testdir", 1024)
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
	ht, err := produceTestHashTree()
	if err != nil {
		t.Fatalf("error making test hash tree %s", err)
	}

	t.Run("valid shard", func(t *testing.T) {
		file := ht.RootDir.Files["test.txt"]
		found, filename, offset := ht.FindShard(file.Hashes[0])

		if !found {
			t.Error("Existing shard not found")
		}

		if filename != filepath.Join("../../test/data/testdir", file.Filename) {
			t.Errorf("Incorrect filepath returned %s", filename)
		}

		if offset != 0 {
			t.Error("Incorrect offset")
		}
	})

	t.Run("invalid shard", func(t *testing.T) {
		found, _, _ := ht.FindShard([32]byte{})
		if found {
			t.Errorf("Shard does not exist but states it is found")
		}
	})

}

// Test getting a shard

func TestGetShard(t *testing.T) {
	ht, err := produceTestHashTree()
	if err != nil {
		t.Fatalf("error making test hash tree %s", err)
	}

	file := ht.RootDir.Files["test.txt"]

	t.Run("shard gotten successfully", func(t *testing.T) {
		data, err := ht.GetShard(file.Hashes[0])
		if err != nil {
			t.Fatalf("failed to get shard %s", err)
		}

		hash := sha256.Sum256(data)
		if !bytes.Equal(file.Hashes[0][:], hash[:]) {
			t.Error("Incorrect shard fetched")
		}
	})

	t.Run("invalid shard given", func(t *testing.T) {
		_, err := ht.GetShard([32]byte{})
		if err == nil {
			t.Fatalf("invalid shard not detected")
		}

	})
}

// Creating empty file

func TestCreateDummyFiles(t *testing.T) {
	tmpFile := "../../test/data/tmp/skeleton.txt"

	smoke := t.Run("function setupFile", func(t *testing.T) {

		t.Run("success", func(t *testing.T) {
			err := setupFile(tmpFile, 64, 100)
			if err != nil {
				t.Errorf("%s", err)
			}

			file, err := os.Stat(tmpFile)
			if err != nil {
				t.Errorf("%s", err)
			}

			if file.Size() != 6400 {
				t.Errorf("Incorrect filesize.\nExpected: %d\nGot: %d", 6400, file.Size())
			}

			err = os.Remove(tmpFile)
			if err != nil {
				t.Errorf("Error removing tmp file")
			}
		})

		t.Run("illegal arguments", func(t *testing.T) {

			// TODO

		})
	})

	if !smoke {
		t.FailNow()
	}

	smoke = t.Run("function createDummyFilesFromDirectory", func(t *testing.T) {
		// TODO
	})

	if !smoke {
		t.FailNow()
	}

	t.Run("function CreateDummyFiles", func(t *testing.T) {
		// TODO
	})

}

// Inserting shards

func TestInsertData(t *testing.T) {
	tmpFile := "../../test/data/tmp/skeleton.txt"
	err := setupFile(tmpFile, 64, 100)
	if err != nil {
		t.Fatalf("error setting up dummy file %s", err)
	}

	t.Run("success", func(t *testing.T) {
		newShard := make([]byte, 64)
		for i := range newShard {
			newShard[i] = 255
		}

		err = insertData("../../test/data/tmp/skeleton.txt", 64, 12, newShard)
		if err != nil {
			t.Errorf("%s", err)
		}

		// ? check the changes occurred
		file, err := os.Open(tmpFile)
		if err != nil {
			t.Fatalf("error opening tmp file %s", err)
		}
		defer file.Close()

		buffer := make([]byte, 64)
		reader := bufio.NewReader(file)

		file.Seek(64*12, 0)
		reader.Read(buffer)

		if !bytes.Equal(buffer, newShard) {
			t.Error("Shard not written correctly to file")
		}
	})

	t.Run("illegal arguments", func(t *testing.T) {
		// TODO
	})

	err = os.Remove(tmpFile)
	if err != nil {
		t.Errorf("Error removing tmp file %s", err)
	}
}
