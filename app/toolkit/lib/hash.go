package lib

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Describes a directory in terms of its files hashes
type hashTree struct {

	// the directory object for the root folder
	RootDir *hashTreeDir `json:"rootdir"`

	// what size should each shard of data be in bytes
	ShardSize uint `json:"shardsize"`

	// the local physical location of the root directory
	RootDirLocation string
}

// Describes a directory with tracked files
type hashTreeDir struct {

	// the path relative to the root directory location
	Dirname string `json:"dirname"`

	// all subdirectories recursively stored
	Subdirs []*hashTreeDir `json:"subdirs"`

	// all files within this folder
	Files []*hashTreeFile `json:"files"`
}

// Describes a singular tracked files
type hashTreeFile struct {
	Filename string `json:"filename"`

	// A list of SHA256 hashes that represent each shard of data in the file
	Hashes [][32]byte `json:"hashes"`
}

type VerifyHashTreeConfig struct {
	IgnoreNewFilesAndDirs     bool
	ContinueAfterError        bool
	continueAfterErrorCounter uint
}

func NewHashTree(rootDir string, shardSize uint) (*hashTree, error) {
	if shardSize == 0 {
		return nil, errors.New("shard size must be greater than 0")
	}

	if _, err := os.Stat(rootDir); os.IsNotExist(err) {
		return nil, err
	}

	return &hashTree{
		RootDir:         nil,
		RootDirLocation: rootDir,
		ShardSize:       shardSize,
	}, nil
}

// IO

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

func ReadHashTreeFromFile(filename string) (*hashTree, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	ht := &hashTree{
		RootDirLocation: filename,
	}

	if err := json.Unmarshal(data, &ht); err != nil {
		return nil, err
	}

	return ht, nil
}

//

func (ht *hashTree) Hash() error {
	fmt.Printf("Starting hash on directory %s\n", ht.RootDirLocation)
	dir, err := ht.hashDir(ht.RootDirLocation, "")
	if err != nil {
		return err
	}

	ht.RootDir = dir
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
		Subdirs: []*hashTreeDir{},
		Files:   []*hashTreeFile{},
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

			dir.Subdirs = append(dir.Subdirs, subdir)
		} else {
			htf, err := ht.shardFile(filepath.Join(currentDir, directory), name)
			if err != nil {
				return nil, err
			}

			dir.Files = append(dir.Files, htf)
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

func (ht *hashTree) VerifyTree(config *VerifyHashTreeConfig, chosenDirectory string) (bool, error) {
	if ht.RootDir == nil {
		return false, errors.New("hash tree not found to compare given directory to")
	}

	fmt.Printf("Verifying directory %s\n", chosenDirectory)
	return ht.verifyDir(config, chosenDirectory, "", ht.RootDir)
}

func (ht *hashTree) verifyDir(config *VerifyHashTreeConfig, currentDir string, directoryBeingVerified string, htDir *hashTreeDir) (bool, error) {
	fmt.Printf("verifying directory %s/%s\n", currentDir, directoryBeingVerified)
	file, err := os.Open(filepath.Join(currentDir, directoryBeingVerified))
	if err != nil {
		return false, err
	}

	defer file.Close()
	list, err := file.Readdirnames(0)
	if err != nil {
		return false, err
	}

	for _, name := range list {
		f, err := os.Stat(filepath.Join(currentDir, directoryBeingVerified, name))
		if err != nil {
			return false, err
		}

		if f.IsDir() {

			// check the subdir should exist
			subdir := func() *hashTreeDir {
				for _, htd := range htDir.Subdirs {
					if htd.Dirname == name {
						return htd
					}
				}
				return nil
			}()

			if subdir == nil {
				if config.IgnoreNewFilesAndDirs {
					continue
				}

				fmt.Printf("Unexpected directory %s/%s/%s\n", currentDir, directoryBeingVerified, name)

				if !config.ContinueAfterError {
					return false, nil
				}
			}

			subdirRes, err := ht.verifyDir(config, filepath.Join(currentDir, directoryBeingVerified), name, subdir)
			if err != nil {
				return false, err
			}

			if !subdirRes {
				if !config.ContinueAfterError {
					return false, nil
				}
			}

		} else {

			// check file should exist
			fileExists := func() *hashTreeFile {
				for _, htf := range htDir.Files {
					if htf.Filename == name {
						return htf
					}
				}
				return nil
			}()

			if fileExists == nil {
				if config.IgnoreNewFilesAndDirs {
					continue
				}

				fmt.Printf("Unexpected file %s/%s/%s\n", currentDir, directoryBeingVerified, name)
				if !config.ContinueAfterError {
					return false, nil
				}
			}

			// compare hashes
			fileRes, err := ht.verifyFile(fileExists, filepath.Join(currentDir, directoryBeingVerified), name)
			if err != nil {
				return false, err
			}

			if !fileRes {
				if !config.ContinueAfterError {
					return false, nil
				}
			}
		}
	}

	return true, nil
}

func (ht *hashTree) verifyFile(htf *hashTreeFile, currentDirectory string, filename string) (bool, error) {
	file, err := os.Open(filepath.Join(currentDirectory, filename))
	if err != nil {
		return false, err
	}

	buffer := make([]byte, ht.ShardSize)
	reader := bufio.NewReader(file)
	counter := 0

	fmt.Printf("\tSharding file '%s'\n", filename)
	for {
		_, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}

		if err != nil {
			return false, err
		}

		hash := sha256.Sum256(buffer)
		if !bytes.Equal(hash[:], htf.Hashes[counter][:]) {
			fmt.Printf("Incorrect hash found in file %s/%s at block %d\n", currentDirectory, filename, counter)
			return false, nil
		}

		counter++
	}

	return true, nil
}
