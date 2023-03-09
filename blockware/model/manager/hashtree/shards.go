package hash

import (
	"bufio"
	"bytes"
	"os"
	"path/filepath"

	"github.com/t02smith/part-iii-project/toolkit/util"
)

/*

A shard is stored as a SHA256 hash of a sequential list of bytes.
Hashing allows us to easily generate a unique fingerprint of each
shard in a given file.

*/

// GetShard locate a given shard within a hash tree by looking through all files
func (ht *HashTree) GetShard(hash [32]byte) (bool, []byte, error) {
	util.Logger.Infof("Looking for shard %x in %s", hash, ht.RootDirLocation)
	found, _, location, offset := ht.FindShard(hash)
	if !found {
		return false, nil, nil
	}

	util.Logger.Infof("Shard found at %s - piece %d", location, offset)
	data, err := ht.readShard(filepath.Join(ht.RootDirLocation, location), offset)
	if err != nil {
		return false, nil, err
	}

	return true, data, nil
}

// LOCATING SHARDS

// looks for a given shard in the entire hash tree
// return <hashFound> <file struct> <relative filename> <shard position in file>
func (ht *HashTree) FindShard(hash [32]byte) (bool, *HashTreeFile, string, int) {
	return ht.RootDir.findShard(hash)
}

// auxillary function for FindShard => looks in particular directory
func (htd *HashTreeDir) findShard(hash [32]byte) (bool, *HashTreeFile, string, int) {
	for _, f := range htd.Files {
		for i, h := range f.Hashes {
			if bytes.Equal(hash[:], h[:]) {
				return true, f, filepath.Join(htd.Dirname, f.Filename), i
			}
		}
	}

	for _, sd := range htd.Subdirs {
		found, htf, filename, offset := sd.findShard(hash)
		if found {
			return true, htf, filepath.Join(htd.Dirname, filename), offset
		}

	}

	return false, nil, "", -1

}

// READING SHARD

// Reads a shard from a given file
func (ht *HashTree) readShard(filename string, offset int) ([]byte, error) {
	fileState, err := os.Stat(filename)
	if err != nil {
		return nil, err
	}

	if fileState.Size() == 0 {
		return make([]byte, ht.ShardSize), nil
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	buffer := make([]byte, ht.ShardSize)
	reader := bufio.NewReader(file)

	_, err = file.Seek(int64(offset*int(ht.ShardSize)), 0)
	if err != nil {
		return nil, err
	}

	_, err = reader.Read(buffer)
	if err != nil {
		return nil, err
	}

	return buffer, nil

}