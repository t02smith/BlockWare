package profiles

import (
	"errors"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
	"github.com/t02smith/part-iii-project/toolkit/model"
	"github.com/t02smith/part-iii-project/toolkit/model/ethereum"
	"github.com/t02smith/part-iii-project/toolkit/model/net"
	listenOnly "github.com/t02smith/part-iii-project/toolkit/test/profiles/listenOnly"
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
	None Profile = iota
	ListenOnly
)

// run a given profile by its ID number
func RunProfile(profileNumber Profile, contractAddr string) error {
	util.Logger.Infof("Profile number %d selected. Attempting to start profile", profileNumber)

	addr := common.HexToAddress(contractAddr)

	switch profileNumber {
	case ListenOnly:
		err := SetupProfile("./test/profiles/listenOnly", listenOnly.PrivateKey, addr, net.PeerConfig{
			ContinueDownloads: false,
			LoadPeersFromFile: false,
		})
		if err != nil {
			return err
		}
		listenOnly.Run()
	case None:
	default:
		return errors.New("unknown profile")
	}

	CloseProfile()
	return nil
}

// general setup needed for all peers
func SetupProfile(path, privateKey string, contractAddr common.Address, config net.PeerConfig) error {
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
	_, err = net.StartPeer(
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

	err = ethereum.ConnectToLibraryInstance(contractAddr, privateKey)
	if err != nil {
		return fmt.Errorf("error connecting to lib instance %s", err)
	}

	return nil
}

// tear down after profile runtime finished
func CloseProfile() {
	net.Peer().Close()
	ethereum.CloseEthClient()
}
