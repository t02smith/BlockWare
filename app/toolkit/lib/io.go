package lib

import (
	"bufio"
	"bytes"
	"errors"
	"os"
	"path/filepath"
)

func (t *hashTree) GetShard(hash [32]byte) ([]byte, error) {
	found, location, offset := t.findShard(hash)
	if !found {
		return nil, errors.New("shard not found")
	}

	return t.readShard(location, offset)
}

// LOCATING SHARDS

func (t *hashTree) findShard(hash [32]byte) (bool, string, int) {
	found, path, offset := t.RootDir.findShard(hash)
	return found, filepath.Join(t.RootDirLocation, path), offset
}

func (td *hashTreeDir) findShard(hash [32]byte) (bool, string, int) {

	for _, f := range td.Files {
		for i, h := range f.Hashes {
			if bytes.Equal(hash[:], h[:]) {
				return true, f.Filename, i
			}
		}
	}

	for _, sd := range td.Subdirs {
		found, filename, offset := sd.findShard(hash)
		if found {
			return true, filepath.Join(td.Dirname, filename), offset
		}

	}

	return false, "", -1

}

// READING SHARD

func (t *hashTree) readShard(filename string, offset int) ([]byte, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	buffer := make([]byte, t.ShardSize)
	reader := bufio.NewReader(file)

	_, err = file.Seek(int64(offset*int(t.ShardSize)), 0)
	if err != nil {
		return nil, err
	}

	_, err = reader.Read(buffer)
	if err != nil {
		return nil, err
	}

	return buffer, nil

}
