package io

import (
	"bufio"
	"bytes"
	"errors"
	"log"
	"os"
	"path/filepath"
)

func (t *HashTree) GetShard(hash [32]byte) ([]byte, error) {
	log.Printf("Looking for shard %x in %s", hash, t.RootDirLocation)
	found, location, offset := t.findShard(hash)
	if !found {
		return nil, errors.New("shard not found")
	}

	log.Printf("Shard found at %s - piece %d", location, offset)
	return t.readShard(location, offset)
}

// LOCATING SHARDS

func (t *HashTree) findShard(hash [32]byte) (bool, string, int) {
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

func (t *HashTree) readShard(filename string, offset int) ([]byte, error) {

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

// DOWNLOADING A FILE

func setupFile(filename string, shardSize, shardCount uint) error {
	log.Printf("Creating empty file %s", filename)
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	emptyBuffer := make([]byte, shardSize)
	writer := bufio.NewWriter(file)

	for i := 0; i < int(shardCount); i++ {
		_, err := writer.Write(emptyBuffer)
		writer.Flush()
		if err != nil {
			return err
		}
	}

	log.Printf("%s created", filename)
	return nil
}

func insertData(filename string, shardSize, offset uint, data []byte) error {
	if len(data) != int(shardSize) {
		return errors.New("data should be the same length as the byte size")
	}

	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Seek(int64(offset*shardSize), 0)
	if err != nil {
		return err
	}

	log.Printf("Writing shard to %s:%d", filename, offset)
	writer := bufio.NewWriter(file)
	writer.Write(data)
	writer.Flush()

	return nil
}