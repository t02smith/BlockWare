package controller

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type ControllerGame struct {

	// game metadata
	Title           string `json:"title"`
	Version         string `json:"version"`
	ReleaseDate     string `json:"release"`
	Developer       string `json:"dev"`
	RootHash        string `json:"rootHash"`
	PreviousVersion string `json:"previousVersion"`

	AssetsFolder string

	// blockchain related
	IPFSId   string         `json:"IPFSId"`
	Price    *big.Int       `json:"price"`
	Uploader common.Address `json:"uploader"`

	Download *ControllerDownload `json:"download"`
	IsOwner  bool
}

type ControllerDownload struct {
	Name        string
	Version     string
	TotalBlocks int
	Stage       string
	ElapsedTime string
	Finished    bool
	BlocksLeft  int
	Paused      bool
}

type ControllerFileProgress struct {
	AbsolutePath    string
	BlocksRemaining []string
}

type ControllerPeerData struct {
	Hostname   string
	Port       uint
	Server     string
	Library    []string
	Validated  bool
	Reputation string
}
