package library

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/t02smith/part-iii-project/toolkit/build/contracts/library"
	"github.com/t02smith/part-iii-project/toolkit/model/manager/games"
	"github.com/t02smith/part-iii-project/toolkit/model/persistence/ethereum"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

/*

This represents functions that are to interface with the Library smart contract.
Most users will connect to an already deployed version of the contract that
should be stored locally when officially deployed.

*/

var (
	libInstance     *library.Library
	authInstance    *bind.TransactOpts
	contractAddress common.Address
)

// GetContractAddress get the address of the currently in-use Library contract
func GetContractAddress() common.Address {
	return contractAddress
}

// DeployLibraryContract setup instance of smart contract to interact with blockchain
func DeployLibraryContract(privateKey string) (*bind.TransactOpts, *library.Library, error) {
	if ethereum.Client() == nil {
		return nil, nil, errors.New("you need to instantiate the Eth Client => run StartClient")
	}

	// run setup functions at most once
	util.Logger.Info("Starting library deployment")
	auth, err := ethereum.GenerateAuthInstance(privateKey)
	if err != nil {
		return nil, nil, err
	}
	authInstance = auth

	addr, _, instance, err := library.DeployLibrary(auth, ethereum.Client())
	if err != nil {
		return nil, nil, err
	}
	contractAddress = addr

	libInstance = instance
	util.Logger.Info("Deployed library")

	return authInstance, libInstance, nil
}

// ConnectToLibraryInstance Connect to an existing contract
func ConnectToLibraryInstance(address common.Address, privateKey string) error {
	util.Logger.Infof("Connecting to addr %s", address)
	lib, err := library.NewLibrary(address, ethereum.Client())
	if err != nil {
		util.Logger.Errorf("Error connecting to library instance: %s", err)
	}

	contractAddress = address
	libInstance = lib
	util.Logger.Infof("Connected to %s", address)
	auth, err := ethereum.GenerateAuthInstance(privateKey)
	if err != nil {
		return err
	}
	authInstance = auth

	return nil
}

// FillLibraryBlockchainGames populate a library with games from the blockchain
func FillLibraryBlockchainGames(lib *games.Library) error {
	util.Logger.Info("Filling library with game metadata from ethereum")
	gs, err := fetchGamesFromEthereum()
	if err != nil {
		return err
	}

	util.Logger.Infof("Populating library with %d games", len(gs))
	for _, g := range gs {
		lib.SetBlockchainGame(g.RootHash, g)
		err := g.DownloadAssets()
		if err != nil {
			util.Logger.Warnf("Error fetching assets for game %x", g.RootHash)
		}
	}
	util.Logger.Infof("Library populated with %d games", len(gs))

	return nil
}

// fetch a list of games from ethereum
func fetchGamesFromEthereum() ([]*games.Game, error) {
	util.Logger.Info("Fetching games from ethereum")

	gameSet := make(map[[32]byte]*games.Game)

	gameSize, err := libInstance.LibSize(nil)
	if err != nil {
		return nil, err
	}
	util.Logger.Infof("Found %d games. Fetching game metadata.", gameSize)

	limit := big.NewInt(10)
	if limit.Cmp(gameSize) == 1 {
		limit = gameSize
	}

	one := big.NewInt(1)
	var empty [32]byte
	for i := big.NewInt(0); i.Cmp(limit) < 0; i.Add(i, one) {
		gameHash, err := libInstance.GameHashes(nil, i)
		if err != nil {
			util.Logger.Warnf("Error getting game hash %s", err)
			continue
		}

		game, err := libInstance.Games(nil, gameHash)
		if err != nil {
			util.Logger.Warnf("Error getting game %s", err)
			continue
		}

		if !bytes.Equal(game.NextVersion[:], empty[:]) {
			mostRecentVersion, err := libInstance.GetMostRecentVersion(nil, game.RootHash)
			if err != nil {
				util.Logger.Warnf("Error getting game %s", err)
				continue
			}

			if _, ok := gameSet[mostRecentVersion]; ok {
				continue
			}

			game, err = libInstance.Games(nil, mostRecentVersion)
			if err != nil {
				util.Logger.Warnf("Error getting game %s", err)
				continue
			}
		}

		gameSet[game.RootHash] = gameStructToGame(game)
	}

	util.Logger.Infof("Fetched %d games from ethereum", len(gameSet))
	var gs []*games.Game
	for _, g := range gameSet {
		gs = append(gs, g)
	}

	return gs, nil
}

// upload a new game to ethereum as well as its data from IPFS
func uploadToEthereum(g *games.Game) error {

	// upload data to IPFS
	util.Logger.Infof("Uploading game data for %s to IPFS", g.Title)
	err := g.UploadHashTree()
	if err != nil {
		return err
	}
	util.Logger.Infof("Successfully uploaded game data for %s to IPFS", g.Title)

	// upload assets
	util.Logger.Infof("Uploading assets to IPFS")
	err = g.UploadAssets()
	if err != nil {
		return err
	}
	util.Logger.Info("Uploaded assets to IPFS")

	// upload metadata to blockchain
	util.Logger.Infof("Uploading game metadata for %s to Ethereum", g.Title)
	util.Logger.Info("%x", g.PreviousVersion)
	_, err = libInstance.UploadGame(authInstance, library.LibraryGameEntry{
		Title:               g.Title,
		Version:             g.Version,
		ReleaseDate:         g.ReleaseDate,
		Developer:           g.Developer,
		RootHash:            g.RootHash,
		HashTreeIPFSAddress: g.HashTreeIPFSAddress,
		PreviousVersion:     g.PreviousVersion,
		Price:               g.Price,
		Uploader:            g.Uploader,
		AssetsIPFSAddress:   g.Assets.Cid,
	})

	if err != nil {
		return err
	}
	util.Logger.Infof("Successfully uploaded game metadata for %s to Ethereum", g.Title)

	return nil
}

// Upload upload a new game to ethereum as well as its data from IPFS with some checks
func Upload(g *games.Game) error {
	if libInstance == nil || authInstance == nil {
		return errors.New("lib or auth instance are nil => run DeployLibraryContract first to initialise them")
	}

	// * upload
	return uploadToEthereum(g)
}

// Purchase purchase a new game off the blockchain
func Purchase(l *games.Library, rootHash [32]byte) error {
	var g *games.Game
	util.Logger.Infof("Attempting to purchase game %x", rootHash)

	// ? do they already own the game
	if game := l.GetOwnedGame(rootHash); game != nil {
		return fmt.Errorf("game %x already purchased", rootHash)
	}

	// ? does the game exist
	util.Logger.Infof("Looking for game %x", rootHash)
	if g = l.GetBlockchainGame(rootHash); g == nil {

		// ! doesn't exist locally => check blockchain again
		util.Logger.Infof("Game not found locally, looking for game %x on eth", rootHash)
		gx, err := libInstance.Games(&bind.CallOpts{}, rootHash)
		if err != nil {
			return err
		}

		g = gameStructToGame(gx)
		l.SetBlockchainGame(g.RootHash, g)
	}

	// * purchase the game
	util.Logger.Infof("Game %x found => proceeding to purchase", rootHash)

	// add value to message for purchase
	authInstance.Value = g.Price
	defer func() { authInstance.Value = big.NewInt(0) }()
	txn, err := libInstance.PurchaseGame(authInstance, rootHash)
	if err != nil {
		return err
	}

	receipt, err := bind.WaitMined(context.Background(), ethereum.Client(), txn)
	if err != nil {
		return err
	}

	if receipt.Status == 0 {
		return fmt.Errorf("transaction failed %v", receipt.Logs)
	}

	// * store game
	l.AddOrUpdateOwnedGame(g)

	// ? fetch hash data
	if err := g.DownloadHashTree(); err != nil {
		util.Logger.Warnf("Error getting hash tree from IPFS %s", err)
	}

	games.OutputAllGameDataToFile(g)
	util.Logger.Info("Game purchased successfully")
	return nil
}

// check whether a user has purchased a game on the blockchain
func HasPurchased(gameHash [32]byte, addr common.Address) (bool, error) {
	return libInstance.HasPurchased(nil, gameHash, addr)
}

// fetch an owned game from blockchain
func FetchOwnedGame(l *games.Library, gameHash [32]byte) error {
	util.Logger.Infof("Fetching game data for %x", gameHash)
	purchased, err := HasPurchased(gameHash, ethereum.Address())
	if err != nil {
		return err
	}

	if !purchased {
		util.Logger.Warnf("User does not own game %x", gameHash)
		return nil
	}

	game, err := libInstance.Games(nil, gameHash)
	if err != nil {
		return err
	}

	localGame := gameStructToGame(game)

	err = localGame.DownloadAllData()
	if err != nil {
		util.Logger.Warnf("error downloading game data %s", err)
	}

	l.AddOrUpdateOwnedGame(localGame)
	util.Logger.Infof("Fetched game data for %x", gameHash)
	return nil
}

// check for updates for all games in a library
func CheckForGameUpdates(lib *games.Library) error {
	var updates []*games.Game

	util.Logger.Debug("Checking for game updates")
	for _, game := range lib.GetOwnedGames() {
		g, err := GetMostRecentVersion(game.RootHash)
		if err != nil {
			return err
		}

		if g != nil {
			updates = append(updates, g)
		}
	}

	util.Logger.Infof("Found %d updates", len(updates))
	for _, update := range updates {
		lib.AddOrUpdateOwnedGame(update)
	}
	util.Logger.Debugf("%d game updates added successfully", len(updates))
	return nil
}

// Get the most recent version of a given game or nil otherwise
func GetMostRecentVersion(gameHash [32]byte) (*games.Game, error) {
	mostRecentVersion, err := libInstance.GetMostRecentVersion(nil, gameHash)
	if err != nil {
		return nil, err
	}

	empty := [32]byte{}
	if bytes.Equal(mostRecentVersion[:], empty[:]) {
		return nil, nil
	}

	game, err := libInstance.Games(nil, mostRecentVersion)
	if err != nil {
		return nil, err
	}

	return gameStructToGame(game), nil
}

func gameStructToGame(game struct {
	Title               string
	Version             string
	ReleaseDate         string
	Developer           string
	RootHash            [32]byte
	PreviousVersion     [32]byte
	NextVersion         [32]byte
	Price               *big.Int
	Uploader            common.Address
	HashTreeIPFSAddress string
	AssetsIPFSAddress   string
}) *games.Game {
	return &games.Game{
		Title:               game.Title,
		Version:             game.Version,
		ReleaseDate:         game.ReleaseDate,
		Developer:           game.Developer,
		RootHash:            game.RootHash,
		HashTreeIPFSAddress: game.HashTreeIPFSAddress,
		Assets: &games.GameAssets{
			Cid: game.AssetsIPFSAddress,
		},
		Uploader:        game.Uploader,
		Price:           game.Price,
		PreviousVersion: game.PreviousVersion,
		NextVersion:     game.NextVersion,
	}
}
