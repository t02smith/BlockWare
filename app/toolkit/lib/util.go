package lib

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func SetupToolkitEnvironment() error {

	// check for .toolkit directory
	toolkitDir := viper.GetString("meta.directory")
	if len(toolkitDir) == 0 {
		toolkitDir = ".toolkit"
	}

	err := createDirectoryIfNotExist(toolkitDir)
	if err != nil {
		return err
	}

	// look for hash directory
	hashDir := viper.GetString("meta.hashes.directory")
	if len(toolkitDir) == 0 {
		hashDir = filepath.Join(toolkitDir, "hashes")
	}

	err = createDirectoryIfNotExist(hashDir)
	if err != nil {
		return err
	}

	return nil
}

func createDirectoryIfNotExist(dir string) error {
	_, err := os.Stat(dir)

	if err != nil && !os.IsNotExist(err) {
		return err
	}

	if os.IsNotExist(err) {
		log.Printf("Directory %s not found. Creating directory", dir)
		err = os.Mkdir(dir, 0775)

		if err != nil {
			return err
		}
	}

	return nil
}