package profiles

import (
	"errors"
	"flag"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
	"github.com/t02smith/part-iii-project/toolkit/model/manager/games"
	"github.com/t02smith/part-iii-project/toolkit/model/net/peer"
	"github.com/t02smith/part-iii-project/toolkit/model/persistence/ethereum"
	"github.com/t02smith/part-iii-project/toolkit/model/persistence/ethereum/library"
	model "github.com/t02smith/part-iii-project/toolkit/model/util"
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
	_selfish              Profile = 3
	_sender               Profile = 4
	_unreliable           Profile = 5
)

func ProfileCLI() bool {
	// * FLAGS
	profile := flag.Uint("profile", 0, "Run the application as a profile. (default off | see profiles.go for details)")
	contractAddr := flag.String("contract", "", "The address of the deployed contract")
	showDebugLogs := flag.Bool("debug", false, "whether debug logs should be displayed")
	port := flag.Uint("port", 6051, "what port should the profile run on")
	deploy := flag.Bool("deploy", false, "deploy a new instance of the contract")

	flag.Parse()
	util.InitLogger(*showDebugLogs)

	// * setup config
	viper.AddConfigPath(".")
	viper.ReadInConfig()
	util.Logger.Debugf("config loaded: %s", viper.AllSettings())

	if *deploy {
		if err := DeployContract(); err != nil {
			util.Logger.Fatal("Failed to deploy contract")
		}
		util.Logger.Info("Contract deployed")
	}

	// ? run a profile
	if *profile != 0 {
		viper.Set("net.port", *port)
		viper.Set("meta.directory", fmt.Sprintf(".toolkit-%d", *port))
		viper.Set("games.installFolder", "./test/profiles/downloads")

		util.Logger.Infof("profile %d selected", *profile)
		err := RunProfile(Profile(*profile), *contractAddr)
		if err != nil {
			util.Logger.Fatalf("Error running profile %d: %s", *profile, err)
		}

		return true
	}

	return false
}

// run a given profile by its ID number
func RunProfile(profileNumber Profile, contractAddr string) error {
	util.Logger.Infof("Profile number %d selected. Attempting to start profile", profileNumber)

	addr := common.HexToAddress(contractAddr)

	switch profileNumber {
	case _listenOnly:
		err := SetupProfile("./test/profiles/", listenOnlyPrivateKey, addr, peer.Config{
			ContinueDownloads:  false,
			LoadPeersFromFile:  false,
			ServeAssets:        false,
			SkipValidation:     false,
			MaxConnections:     50,
			TrackContributions: false,
		})
		if err != nil {
			return err
		}
		listenOnlyRun()

	case _listenOnlyWithUpload:
		err := SetupProfile("./test/profiles/", listenOnlyUploadPrivateKey, addr, peer.Config{
			ContinueDownloads:  false,
			LoadPeersFromFile:  false,
			ServeAssets:        false,
			SkipValidation:     false,
			MaxConnections:     50,
			TrackContributions: false,
		})
		if err != nil {
			return err
		}
		listenOnlyUploadRun()
	case _selfish:
		if err := SetupProfile("./test/profiles/", selfishPrivateKey, addr, peer.Config{
			ContinueDownloads:  false,
			LoadPeersFromFile:  false,
			ServeAssets:        false,
			SkipValidation:     false,
			MaxConnections:     50,
			TrackContributions: false,
		}); err != nil {
			return err
		}
		selfishRun()
	case _sender:
		if err := SetupProfile("./test/profiles/", senderPrivateKey, addr, peer.Config{
			ContinueDownloads:  false,
			LoadPeersFromFile:  false,
			ServeAssets:        false,
			SkipValidation:     false,
			MaxConnections:     50,
			TrackContributions: false,
		}); err != nil {
			return err
		}
		senderRun()
	case _unreliable:
		if err := SetupProfile("./test/profiles/", unrelaiblePrivateKey, addr, peer.Config{
			ContinueDownloads:  false,
			LoadPeersFromFile:  false,
			ServeAssets:        false,
			SkipValidation:     false,
			MaxConnections:     50,
			TrackContributions: false,
		}); err != nil {
			return err
		}
		unreliableRun(50)
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

	err = os.RemoveAll(viper.GetString("meta.directory"))
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

// deploy a new instance of the contract
func DeployContract() error {
	client, _, err := ethereum.StartClient("http://localhost:8545")
	if err != nil {
		return fmt.Errorf("error starting eth client %s", err)
	}

	_, _, err = library.DeployLibraryContract("af9668cd6ebc3ba4c0e5036c284e128ed66e18ba9e4ed87b2c0c6d9642f2b879")
	if err != nil {
		return fmt.Errorf("error connecting to lib instance %s", err)
	}

	client.Close()
	return nil
}

func SetupGame() error {
	p := peer.Peer()

	g1, err := games.CreateGame(games.NewGame{
		Title:       "t02smith.github.io",
		Version:     "4.7.1",
		ReleaseDate: time.Date(2002, time.January, 10, 0, 0, 0, 0, time.UTC).String(),
		Developer:   "tcs1g20",
		RootDir:     "./games/t02smith.github.io",
		Price:       big.NewInt(150),
		ShardSize:   4194304,
		AssetsDir:   "../data/assets"},
		nil,
	)

	if err != nil {
		return err
	}

	p.Library().AddOrUpdateOwnedGame(g1)
	err = games.OutputAllGameDataToFile(g1)
	if err != nil {
		return err
	}

	return err
}
