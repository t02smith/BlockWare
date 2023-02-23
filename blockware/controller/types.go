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

	// blockchain related
	IPFSId   string         `json:"IPFSId"`
	Price    *big.Int       `json:"price"`
	Uploader common.Address `json:"uploader"`

	Download *ControllerDownload `json:"download"`
}

type ControllerDownload struct {
	Progress    map[string]ControllerFileProgress
	TotalBlocks int
}

type ControllerFileProgress struct {
	AbsolutePath    string
	BlocksRemaining []string
}