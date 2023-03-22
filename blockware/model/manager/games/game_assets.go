package games

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/spf13/viper"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

/*
Game assets will be stored on IPFS and should be fetched as
necessary and will include things like artwork, description,
etc.

The assets should be stored in a .zip archive and use the
following naming scheme:

- "cover.png" = main piece of artwork for the game
- "background.png" = wide element on store page
- "description.md" = a description of the game to be displayed on its library page
*/

type GameAssets struct {
	// the address of the assets on IPFS
	Cid string

	// where the assets are stored locally
	AbsolutePath string
}

const (
	ASSET_COVER       string = "cover.png"
	ASSET_DESCRIPTION string = "description.md"
	ASSET_BACKGROUND  string = "background.png"
)

//

// upload game assets to IPFS
func (g *Game) UploadAssets() error {
	if g.Assets.AbsolutePath == "" {
		return fmt.Errorf("assets directory not specified")
	}

	util.Logger.Infof("Looking for asset folder at %s", g.Assets.AbsolutePath)

	// * upload to IPFS
	util.Logger.Infof("Uploading assets for %x to IPFS", g.RootHash)
	sh := shell.NewShell("localhost:5001")

	cid, err := sh.AddDir(g.Assets.AbsolutePath)
	if err != nil {
		return err
	}

	g.Assets.Cid = cid
	if err = g.CopyAssetsFolder(); err != nil {
		util.Logger.Warnf("error copying assets folder %s", err)
	}

	util.Logger.Infof("Uploaded assets for %x to IPFS", g.RootHash)
	return nil
}

// get assets off of IPFS
func (g *Game) DownloadAssets() error {
	if g.Assets == nil {
		return fmt.Errorf("asset folder not found for game %x", g.RootHash)
	}

	g.Assets.AbsolutePath = filepath.Join(viper.GetString("meta.directory"), "assets", fmt.Sprintf("%x", g.RootHash))

	// * get data from IPFS
	sh := shell.NewShell("localhost:5001")
	err := sh.Get(g.Assets.Cid, g.Assets.AbsolutePath)
	if err != nil {
		return err
	}

	return nil
}

func (g *Game) CopyAssetsFolder() error {
	util.Logger.Debugf("Copying asset folder %s", g.Assets.AbsolutePath)
	dirname := filepath.Join(viper.GetString("meta.directory"), "assets", fmt.Sprintf("%x", g.RootHash))
	err := os.Mkdir(dirname, 0755)
	if err != nil {
		return err
	}

	expectedFiles := []string{ASSET_COVER, ASSET_BACKGROUND, ASSET_DESCRIPTION}

	_, err = os.Stat(g.Assets.AbsolutePath)
	if err != nil {
		return err
	}

	for _, f := range expectedFiles {
		util.Logger.Debugf("copying %s", f)
		err = func() error {
			input, err := os.Open(filepath.Join(g.Assets.AbsolutePath, f))
			if err != nil {
				return err
			}
			defer input.Close()

			output, err := os.Create(filepath.Join(dirname, f))
			if err != nil {
				return err
			}
			defer output.Close()

			_, err = io.Copy(output, input)
			if err != nil {
				return err
			}

			return nil
		}()

		if err != nil {
			return err
		}

	}

	util.Logger.Debugf("copied asset folder %s", g.Assets.AbsolutePath)
	return nil
}
