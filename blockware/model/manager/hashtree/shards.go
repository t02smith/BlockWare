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

type ShardLocation struct {
	File         *HashTreeFile
	AbsolutePath string
	Offset       int
}

// GetShard locate a given shard within a hash tree by looking through all files
func (ht *HashTree) GetShard(hash [32]byte) (bool, []byte, error) {
	util.Logger.Debugf("Looking for shard %x in %s", hash, ht.RootDirLocation)
	locations := ht.FindShard(hash)
	if len(locations) == 0 {
		return false, nil, nil
	}

	// ? choose first found
	shard := locations[0]

	util.Logger.Infof("Shard found at %s - piece %d", shard.AbsolutePath, shard.Offset)
	data, err := ht.readShard(filepath.Join(ht.RootDirLocation, shard.AbsolutePath), shard.Offset)
	if err != nil {
		return false, nil, err
	}

	return true, data, nil
}

// LOCATING SHARDS

// looks for a given shard in the entire hash tree
// return <hashFound> <file struct> <relative filename> <shard position in file>
func (ht *HashTree) FindShard(hash [32]byte) []*ShardLocation {
	return ht.RootDir.findShard(hash)
}

// auxillary function for FindShard => looks in particular directory
func (htd *HashTreeDir) findShard(hash [32]byte) []*ShardLocation {
	var shardLocations []*ShardLocation
	for _, f := range htd.Files {
		for i, h := range f.Hashes {
			if bytes.Equal(hash[:], h[:]) {
				shardLocations = append(shardLocations, &ShardLocation{
					File:         f,
					AbsolutePath: filepath.Join(htd.Dirname, f.Filename),
					Offset:       i,
				})
			}
		}
	}

	for _, sd := range htd.Subdirs {
		locations := sd.findShard(hash)
		for i := range locations {
			locations[i].AbsolutePath = filepath.Join(htd.Dirname, locations[i].AbsolutePath)
		}

		shardLocations = append(shardLocations, locations...)

	}

	return shardLocations

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
