package controller

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/t02smith/part-iii-project/toolkit/model/ethereum"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

/*

This file represents the controller interface between
the frontend and backend. Methods on the App type can
be called from the frontend code.

*/

// ? App setup

type Controller struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewController() *Controller {
	return &Controller{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *Controller) Startup(ctx context.Context) {
	util.Logger.Info("Starting app context")
	a.ctx = ctx
}

// ? Interface functions

// deploy a new instance of the library contract
func (a *Controller) DeployLibraryInstance(privateKey string) string {
	_, _, err := ethereum.DeployLibraryContract(privateKey)
	if err != nil {
		return ""
	}

	return ethereum.GetContractAddress().Hex()
}

// connect to an existing library contract
func (a *Controller) JoinLibraryInstance(address, privateKey string) {
	addr := common.HexToAddress(address)
	err := ethereum.ConnectToLibraryInstance(addr, privateKey)
	if err != nil {
		util.Logger.Errorf("Error joining lib instance %s", err)
	}
}
