package games

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/spf13/viper"
	"github.com/t02smith/part-iii-project/toolkit/model/hash"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

func (g *Game) UploadDataToIPFS() error {
	sh := shell.NewShell("localhost:5001")

	if g.data == nil {
		err := g.ReadHashData()
		if err != nil {
			return err
		}
	}

	json, err := json.Marshal(g.data)
	if err != nil {
		return err
	}

	util.Logger.Info("Uploading game %s data to IPFS", g.Title)
	reader := bytes.NewReader(json)
	cid, err := sh.Add(reader)
	if err != nil {
		return err
	}

	g.IPFSId = cid
	util.Logger.Infof("Uploaded game %s data to IPFS. CID = %s", g.Title, cid)
	return nil
}

// read game data from IPFS network
func (g *Game) ReadDataFromIPFS() error {
	if g.IPFSId == "" {
		return errors.New("ipfs id not stored")
	}

	util.Logger.Info("Reading game %s data from IPFS using cid=%s", g.Title, g.IPFSId)
	sh := shell.NewShell("localhost:5001")
	data, err := sh.Cat(g.IPFSId)
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
	err = g.data.OutputToFile(filepath.Join(viper.GetString("meta.hashes.directory"), fmt.Sprintf("%s-%s-%s.hash.json", g.Title, g.Version, g.Developer)))
	if err != nil {
		return err
	}

	util.Logger.Info("Read game %s data from IPFS", g.Title)
	return nil
}
