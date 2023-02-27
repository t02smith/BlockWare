package games

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/t02smith/part-iii-project/toolkit/model"
	hash "github.com/t02smith/part-iii-project/toolkit/model/hashtree"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

/**

Hash data for each game will be uploaded to IPFS so
that it can be easily accessed by anyone, anywhere.

*/

// Upload game data to IPFS
func (g *Game) UploadHashTreeToIPFS() error {
	sh := shell.NewShell("localhost:5001")
	data, err := g.GetData()
	if err != nil {
		return err
	}

	json, err := json.Marshal(data)
	if err != nil {
		return err
	}

	util.Logger.Info("Uploading game %s data to IPFS", g.Title)
	reader := bytes.NewReader(json)
	cid, err := sh.Add(reader)
	if err != nil {
		return err
	}

	g.HashTreeIPFSAddress = cid
	util.Logger.Infof("Uploaded game %s data to IPFS. CID = %s", g.Title, cid)
	return nil
}

// read game data from IPFS network
func (g *Game) GetHashTreeFromIPFS() error {
	if g.HashTreeIPFSAddress == "" {
		return errors.New("ipfs id not stored")
	}

	util.Logger.Info("Reading game %s data from IPFS using cid=%s", g.Title, g.HashTreeIPFSAddress)
	sh := shell.NewShell("localhost:5001")
	data, err := sh.Cat(g.HashTreeIPFSAddress)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(data)

	g.data = &hash.HashTree{}
	err = json.Unmarshal(buf.Bytes(), g.data)
	if err != nil {
		return err
	}

	util.Logger.Info("Read game %s data from IPFS", g.Title)
	return nil
}

func (ga *GameAssets) FetchGameAssets(outputLocation string) error {
	if ga.IPFSId == "" {
		return fmt.Errorf("asset location not specified")
	}

	if outputLocation == "" {
		outputLocation = "."
	}

	// * fetch from ipfs
	sh := shell.NewShell("localhost:5001")
	util.Logger.Infof("Fetching assets from IPFS address %s", ga.IPFSId)

	data, err := sh.Cat(ga.IPFSId)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(data)

	util.Logger.Infof("Writing archive to file")
	archive, err := os.Create(filepath.Join(outputLocation, fmt.Sprintf("%s.zip", ga.IPFSId)))
	if err != nil {
		return err
	}

	_, err = archive.Write(buf.Bytes())
	if err != nil {
		archive.Close()
		return err
	}

	archive.Close()
	util.Logger.Info("Archive written to file")

	// * unzip archive
	ga.absolutePath = filepath.Join(outputLocation, ga.IPFSId)
	err = model.UnzipArchive(filepath.Join(outputLocation, fmt.Sprintf("%s.zip", ga.IPFSId)), ga.absolutePath)
	if err != nil {
		return err
	}

	// * tear down
	err = os.Remove(filepath.Join(outputLocation, fmt.Sprintf("%s.zip", ga.IPFSId)))
	if err != nil {
		util.Logger.Warnf("Error removing archive: %s", err)
	}

	return nil
}
