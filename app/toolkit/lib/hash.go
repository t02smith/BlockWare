package lib

import (
	"bufio"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type hashTree struct {
	RootDir         hashTreeDir `json:"rootdir"`
	RootDirLocation string
	ShardSize       uint `json:"shardsize"`
}

type hashTreeDir struct {
	Dirname string         `json:"dirname"`
	Subdirs []hashTreeDir  `json:"subdirs"`
	Files   []hashTreeFile `json:"files"`
}

type hashTreeFile struct {
	Filename string     `json:"filename"`
	Hashes   [][32]byte `json:"hashes"`
}

func NewHashTree(rootDir string, shardSize uint) (*hashTree, error) {
	if shardSize == 0 {
		return nil, errors.New("shard size must be greater than 0")
	}

	if _, err := os.Stat(rootDir); os.IsNotExist(err) {
		return nil, err
	}

	return &hashTree{
		RootDir: hashTreeDir{
			Dirname: rootDir,
			Subdirs: []hashTreeDir{},
			Files:   []hashTreeFile{},
		},
		RootDirLocation: rootDir,
		ShardSize:       shardSize,
	}, nil
}

func (ht *hashTree) OutputToFile(filename string) error {
	fmt.Printf("outputting to file %s\n", filename)
	e, err := json.Marshal(ht)
	if err != nil {
		fmt.Println(err)
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString(string(e))
	writer.Flush()
	return nil
}

func (ht *hashTree) Hash() error {
	fmt.Printf("Starting hash on directory %s\n", ht.RootDirLocation)
	dir, err := ht.hashDir(ht.RootDirLocation, "")
	if err != nil {
		return err
	}

	ht.RootDir = *dir
	return nil
}

func (ht *hashTree) hashDir(currentDir string, directory string) (*hashTreeDir, error) {
	fmt.Printf("Hashing directory %s\n", directory)
	file, err := os.Open(filepath.Join(currentDir, directory))
	if err != nil {
		return nil, err
	}

	defer file.Close()
	list, err := file.Readdirnames(0)
	if err != nil {
		return nil, err
	}

	dir := &hashTreeDir{
		Dirname: directory,
		Subdirs: []hashTreeDir{},
		Files:   []hashTreeFile{},
	}

	for _, name := range list {
		f, err := os.Stat(filepath.Join(currentDir, directory, name))
		if err != nil {
			return nil, err
		}

		if f.IsDir() {
			subdir, err := ht.hashDir(filepath.Join(currentDir, directory), name)
			if err != nil {
				return nil, err
			}

			dir.Subdirs = append(dir.Subdirs, *subdir)
		} else {
			htf, err := ht.shardFile(filepath.Join(currentDir, directory), name)
			if err != nil {
				return nil, err
			}

			dir.Files = append(dir.Files, *htf)
		}
	}

	return dir, nil
}

func (ht *hashTree) shardFile(currentDirectory string, filename string) (*hashTreeFile, error) {
	file, err := os.Open(filepath.Join(currentDirectory, filename))
	if err != nil {
		return nil, err
	}

	htf := &hashTreeFile{
		Filename: filename,
		Hashes:   [][32]byte{},
	}

	buffer := make([]byte, ht.ShardSize)
	reader := bufio.NewReader(file)

	fmt.Printf("\tSharding file '%s'\n", filename)
	for {
		_, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		hash := sha256.Sum256(buffer)
		htf.Hashes = append(htf.Hashes, hash)
	}

	return htf, nil
}
