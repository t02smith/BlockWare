package ethereum

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/t02smith/part-iii-project/toolkit/build/contracts/library"
	"github.com/t02smith/part-iii-project/toolkit/model/manager/games"
	"github.com/t02smith/part-iii-project/toolkit/model/net/peer"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

/*

This represents functions that are to interface with the Library smart contract.
Most users will connect to an already deployed version of the contract that
should be stored locally when officially deployed.

*/

var (
	lib_instance     *library.Library
	auth_instance    *bind.TransactOpts
	contract_address common.Address
)

// get the address of the currently in-use Library contract
func GetContractAddress() common.Address {
	return contract_address
}

// setup instance of smart contract to interact with blockchain
func DeployLibraryContract(privateKey string) (*bind.TransactOpts, *library.Library, error) {
	if eth_client == nil {
		return nil, nil, errors.New("you need to instantiate the Eth Client => run StartClient")
	}

	// run setup functions at most once
	util.Logger.Info("Starting library deployment")
	auth, err := generateAuthInstance(privateKey)
	if err != nil {
		return nil, nil, err
	}

	addr, _, instance, err := library.DeployLibrary(auth, eth_client)
	if err != nil {
		return nil, nil, err
	}
	contract_address = addr

	lib_instance = instance
	util.Logger.Info("Deployed library")

	err = ReadPreviousGameEvents()
	if err != nil {
		util.Logger.Errorf("Error reading previous games: %s", err)
	}

	err = watchNewGameEvent()
	if err != nil {
		util.Logger.Errorf("Error watching for new games: %s", err)
		return nil, nil, err
	}

	return auth_instance, lib_instance, nil
}

// Connect to an existing contract
func ConnectToLibraryInstance(address common.Address, privateKey string) error {
	util.Logger.Infof("Connecting to addr %s", address)
	lib, err := library.NewLibrary(address, eth_client)
	if err != nil {
		util.Logger.Errorf("Error connecting to library instance: %s", err)
	}

	contract_address = address
	lib_instance = lib
	util.Logger.Infof("Connected to %s", address)
	_, err = generateAuthInstance(privateKey)
	if err != nil {
		return err
	}

	return nil
}

// populate a library with games from the blockchain
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
	gs := []*games.Game{}

	gameSize, err := lib_instance.LibSize(nil)
	if err != nil {
		return nil, err
	}
	util.Logger.Infof("Found %d games. Fetching game metadata.", gameSize)

	one := big.NewInt(1)
	for i := big.NewInt(0); i.Cmp(gameSize) < 0; i.Add(i, one) {
		gameHash, err := lib_instance.GameHashes(nil, i)
		if err != nil {
			util.Logger.Warnf("Error getting game hash %s", err)
			continue
		}

		game, err := lib_instance.Games(nil, gameHash)
		if err != nil {
			util.Logger.Warnf("Error getting game %s", err)
			continue
		}

		gs = append(gs, &games.Game{
			Title:               game.Title,
			Version:             game.Version,
			ReleaseDate:         game.ReleaseDate,
			Developer:           game.Developer,
			RootHash:            game.RootHash,
			HashTreeIPFSAddress: game.HashTreeIPFSAddress,
			Uploader:            game.Uploader,
			Price:               game.Price,
			PreviousVersion:     game.PreviousVersion,
			Assets: &games.GameAssets{
				Cid: game.AssetsIPFSAddress,
			},
		})
	}

	util.Logger.Infof("Fetched %d games from ethereum", len(gs))
	return gs, nil
}

// upload a new game to ethereum as well as its data from IPFS
func uploadToEthereum(g *games.Game) error {

	// upload data to IPFS
	util.Logger.Infof("Uploading game data for %s to IPFS", g.Title)
	err := g.UploadHashTreeToIPFS()
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
	_, err = lib_instance.UploadGame(auth_instance, library.LibraryGameEntry{
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

// upload a new game to ethereum as well as its data from IPFS with some checks
func Upload(g *games.Game) error {
	if lib_instance == nil || auth_instance == nil {
		return errors.New("lib or auth instance are nil => run DeployLibraryContract first to initialise them")
	}

	// * upload
	return uploadToEthereum(g)
}

// Will listen for NewGame events emitted on ethereum
func watchNewGameEvent() error {
	newGameChannel := make(chan *library.LibraryNewGame)

	sub, err := lib_instance.WatchNewGame(&bind.WatchOpts{
		Start:   nil,
		Context: nil,
	}, newGameChannel)

	if err != nil {
		return err
	}

	util.Logger.Info("Watching for new games")
	go func() {
		p := peer.Peer()
		defer util.Logger.Info("Stopped watching for new games")
		defer sub.Unsubscribe()

		for {
			select {
			case err := <-sub.Err():
				if err != nil {
					util.Logger.Error(err)
				}
			case newGame := <-newGameChannel:
				util.Logger.Infof("New game received with hash %x. Adding to library.", newGame.Hash)
				p.Library().SetBlockchainGame(newGame.Hash, gameEntryToGame(&newGame.Game))
			}
		}

	}()

	return nil
}

// Will look at previous GameEntry events to fill store games
func ReadPreviousGameEvents() error {
	util.Logger.Info("Reading previous games from eth")
	newGameIterator, err := lib_instance.FilterNewGame(&bind.FilterOpts{
		End:     nil,
		Start:   1,
		Context: nil,
	})
	if err != nil {
		return err
	}

	lib := peer.Peer().Library()
	count := 0
	for newGameIterator.Next() {
		g := newGameIterator.Event.Game
		util.Logger.Infof("Found game %s", g.Title)
		lib.SetBlockchainGame(g.RootHash, gameEntryToGame(&g))
		count++
	}

	util.Logger.Infof("Finished reading previous games from eth. Found %d games.", count)
	return nil
}

// translates from the ethereum version to the locally used struct
func gameEntryToGame(game *library.LibraryGameEntry) *games.Game {
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
	}
}

// purchase a new game off the blockchain
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
		gx, err := lib_instance.Games(&bind.CallOpts{}, rootHash)
		if err != nil {
			return err
		}

		g = gameEntryToGame(&library.LibraryGameEntry{
			Title:               gx.Title,
			Version:             gx.Version,
			ReleaseDate:         gx.ReleaseDate,
			Developer:           gx.Developer,
			RootHash:            gx.RootHash,
			PreviousVersion:     gx.PreviousVersion,
			Price:               gx.Price,
			Uploader:            gx.Uploader,
			HashTreeIPFSAddress: gx.HashTreeIPFSAddress,
			AssetsIPFSAddress:   gx.AssetsIPFSAddress,
		})

		l.SetBlockchainGame(g.RootHash, g)
	}

	// * purchase the game
	util.Logger.Infof("Game %x found => proceeding to purchase", rootHash)
	txn, err := lib_instance.PurchaseGame(auth_instance, rootHash)
	if err != nil {
		return err
	}

	receipt, err := bind.WaitMined(context.Background(), eth_client, txn)
	if err != nil {
		return err
	}

	if receipt.Status == 0 {
		return fmt.Errorf("transaction failed %v", receipt.Logs)
	}

	// * store game
	err = l.AddOwnedGame(g)
	if err != nil {
		return err
	}

	// ? fetch hash data
	err = g.GetHashTreeFromIPFS()
	if err != nil {
		util.Logger.Warnf("Error getting hash tree from IPFS %s", err)
	}

	games.OutputAllGameDataToFile(g)

	util.Logger.Info("Game purchased successfully")
	return nil
}
