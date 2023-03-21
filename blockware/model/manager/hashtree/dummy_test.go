package hash

import (
	"bufio"
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*

function: CreateDummyFiles
purpose: create skeleton files from a given hash tree

? Test cases
success
	| #1 => fileSize % shardSize == 0
	| #2 => fileSize % shardSize != 0
	| #3 => empty file

failure
	| illegal arguments
			| #1 => root directory doesn't exist
			| #2 => title matches file in root directory

*/

func TestCreateDummyFiles(t *testing.T) {
	tmpFile := "../../test/data/tmp/skeleton.txt"

	smoke := t.Run("function setupFile", func(t *testing.T) {

		t.Run("success", func(t *testing.T) {

			inputs := []struct {
				name       string
				shardSize  uint
				shardCount int
				fileSize   int
			}{
				{"fileSize % shardSize == 0", 64, 100, 6400},
				{"fileSize % shardSize != 0", 64, 101, 6452},
				{"empty file", 0, 0, 0},
			}

			for _, in := range inputs {
				t.Run(in.name, func(t *testing.T) {
					err := setupFile(tmpFile, in.shardSize, in.shardCount, in.fileSize)
					assert.Nil(t, err)

					t.Cleanup(func() {
						os.Remove(tmpFile)
					})

					file, err := os.Stat(tmpFile)
					assert.Nil(t, err)

					assert.Equal(t, in.fileSize, int(file.Size()))
				})
			}

		})

		t.Run("illegal arguments", func(t *testing.T) {

			// TODO

		})
	})

	if !smoke {
		t.FailNow()
	}
}

// Inserting shards

func TestInsertData(t *testing.T) {
	tmpFile := "../../../test/data/tmp/skeleton.txt"
	err := setupFile(tmpFile, 64, 100, 6400)
	if err != nil {
		t.Fatalf("error setting up dummy file %s", err)
	}

	t.Run("success", func(t *testing.T) {
		newShard := make([]byte, 64)
		for i := range newShard {
			newShard[i] = 255
		}

		err = InsertData("../../../test/data/tmp/skeleton.txt", 64, 12, newShard)
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
