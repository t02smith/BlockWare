package profiles

import (
	"errors"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
	"github.com/t02smith/part-iii-project/toolkit/model"
	"github.com/t02smith/part-iii-project/toolkit/model/net/peer"
	"github.com/t02smith/part-iii-project/toolkit/model/persistence/ethereum"
	"github.com/t02smith/part-iii-project/toolkit/model/persistence/ethereum/library"
	deployEth "github.com/t02smith/part-iii-project/toolkit/test/profiles/deploy"
	listenOnly "github.com/t02smith/part-iii-project/toolkit/test/profiles/listenOnly"
	listenOnlyUpload "github.com/t02smith/part-iii-project/toolkit/test/profiles/listenOnlyWithUpload"
	"github.com/t02smith/part-iii-project/toolkit/test/testutil"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

/*

Profiles are cmd based instances of this application
and will run a predefined set of steps to mimic a
typically peer whilst automating certain aspects.

Each profile has a specific purpose that is explained
in the file of each one and its own custom config file
to be parsed initially.

*/

type Profile uint8

const (
	None                  Profile = 0
	_listenOnlyWithUpload Profile = 1
	_listenOnly           Profile = 2
	_deploy               Profile = 3
)

// run a given profile by its ID number
func RunProfile(profileNumber Profile, contractAddr string) error {
	util.Logger.Infof("Profile number %d selected. Attempting to start profile", profileNumber)

	addr := common.HexToAddress(contractAddr)

	switch profileNumber {
	case _listenOnly:
		err := SetupProfile("./test/profiles/listenOnly", listenOnly.PrivateKey, addr, peer.Config{
			ContinueDownloads: false,
			LoadPeersFromFile: false,
			ServeAssets:       false,
		})
		if err != nil {
			return err
		}
		listenOnly.Run()

	case _listenOnlyWithUpload:
		err := SetupProfile("./test/profiles/listenOnlyWithUpload", listenOnlyUpload.PrivateKey, addr, peer.Config{
			ContinueDownloads: false,
			LoadPeersFromFile: false,
			ServeAssets:       false,
		})
		if err != nil {
			return err
		}
		listenOnlyUpload.Run()
	case _deploy:
		return deployEth.Run(testutil.Accounts[3][1])
	case None:
	default:
		return errors.New("unknown profile")
	}

	CloseProfile()
	return nil
}

// general setup needed for all peers
func SetupProfile(path, privateKey string, contractAddr common.Address, config peer.Config) error {
	err := os.Chdir(path)
	if err != nil {
		return err
	}

	// * setup config
	viper.AddConfigPath(".")
	viper.ReadInConfig()
	util.Logger.Debugf("config loaded: %s", viper.AllSettings())

	err = os.RemoveAll(".toolkit")
	if err != nil {
		return err
	}

	err = model.SetupToolkitEnvironment()
	if err != nil {
		return fmt.Errorf("error setting up toolkit env %s", err)
	}

	// * start peer
	_, err = peer.StartPeer(
		config,
		"localhost",
		viper.GetUint("net.port"),
		viper.GetString("games.installFolder"),
		viper.GetString("meta.directory"),
	)

	if err != nil {
		return fmt.Errorf("error starting peer %s", err)
	}

	// * start eth client
	_, _, err = ethereum.StartClient(viper.GetString("eth.address"))
	if err != nil {
		return fmt.Errorf("error starting eth client %s", err)
	}

	library.ConnectToLibraryInstance(contractAddr, privateKey)

	return nil
}

// tear down after profile runtime finished
func CloseProfile() {
	peer.Peer().Close()
	ethereum.CloseEthClient()
}
