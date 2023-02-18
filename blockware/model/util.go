package model

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
	u "github.com/t02smith/part-iii-project/toolkit/util"
)

func SetupToolkitEnvironment() error {

	// check for .toolkit directory
	toolkitDir := viper.GetString("meta.directory")
	if len(toolkitDir) == 0 {
		toolkitDir = ".toolkit"
	}

	err := CreateDirectoryIfNotExist(toolkitDir)
	if err != nil {
		return err
	}

	// look for hash directory
	hashDir := filepath.Join(toolkitDir, "hashes")
	err = CreateDirectoryIfNotExist(hashDir)
	if err != nil {
		return err
	}

	return nil
}

func CreateDirectoryIfNotExist(dir string) error {
	_, err := os.Stat(dir)

	if err != nil && !os.IsNotExist(err) {
		return err
	}

	if os.IsNotExist(err) {
		u.Logger.Infof("Directory %s not found. Creating directory", dir)
		err = os.Mkdir(dir, 0775)

		if err != nil {
			return err
		}
	}

	return nil
}
