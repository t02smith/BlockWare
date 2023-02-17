package ethereum

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/t02smith/part-iii-project/toolkit/build/contracts/library"
	"github.com/t02smith/part-iii-project/toolkit/model/games"
	"github.com/t02smith/part-iii-project/toolkit/model/net"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

var (
	lib_instance  *library.Library
	auth_instance *bind.TransactOpts
	lib_auth_once sync.Once
)

// setup instance of smart contract to interact with blockchain
func DeployLibraryContract(privateKey string) (*bind.TransactOpts, *library.Library, error) {
	if eth_client == nil {
		return nil, nil, errors.New("you need to instantiate the Eth Client => run StartClient")
	}

	// run setup functions at most once
	lib_auth_once.Do(func() {
		util.Logger.Info("Starting library deployment")
		privKeyECDSA, err := crypto.HexToECDSA(privateKey)
		if err != nil {
			util.Logger.Panic(err)
		}

		pubKeyECDSA, ok := privKeyECDSA.Public().(*ecdsa.PublicKey)
		if !ok {
			util.Logger.Panic("public key of incorrect type")
		}

		fromAddress := crypto.PubkeyToAddress(*pubKeyECDSA)
		nonce, err := eth_client.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			util.Logger.Panic(err)
		}

		util.Logger.Info("Getting chain ID")
		chainID, err := eth_client.ChainID(context.TODO())
		if err != nil {
			util.Logger.Panic(err)
		}

		util.Logger.Info("Getting gas price")
		gasPrice, err := eth_client.SuggestGasPrice(context.Background())
		if err != nil {
			util.Logger.Panic(err)
		}

		auth, err := bind.NewKeyedTransactorWithChainID(privKeyECDSA, chainID)
		if err != nil {
			util.Logger.Panic(err)
		}

		auth.Nonce = big.NewInt(int64(nonce))
		auth.Value = big.NewInt(0)
		auth.GasLimit = uint64(3000000)
		auth.GasPrice = gasPrice

		auth_instance = auth

		_, _, instance, err := library.DeployLibrary(auth, eth_client)
		if err != nil {
			util.Logger.Panic(err)
		}

		lib_instance = instance
		util.Logger.Info("Deployed library")

		err = ReadPreviousGameEvents()
		if err != nil {
			util.Logger.Errorf("Error reading previous games: %s", err)
		}

		err = watchNewGameEvent()
		if err != nil {
			util.Logger.Errorf("Error watching for new games: %s", err)
		}
	})

	return auth_instance, lib_instance, nil
}

// checks if a game has already been uploaded to the blockchain
func isGameOnBlockchain(rootHash [32]byte) (bool, error) {
	gameSize, err := lib_instance.LibSize(nil)
	if err != nil {
		return false, err
	}

	one := big.NewInt(1)
	for i := big.NewInt(0); i.Cmp(gameSize) < 0; i.Add(i, one) {
		gameHash, err := lib_instance.GameHashes(nil, i)
		if err != nil {
			util.Logger.Warnf("Error getting game hash %s", err)
			continue
		}

		if bytes.Equal(gameHash[:], rootHash[:]) {
			return true, nil
		}
	}

	return false, nil
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
			Title:           game.Title,
			Version:         game.Version,
			ReleaseDate:     game.ReleaseDate,
			Developer:       game.Developer,
			RootHash:        game.RootHash,
			IPFSId:          game.IpfsAddress,
			Uploader:        game.Uploader,
			Price:           game.Price,
			PreviousVersion: game.PreviousVersion,
		})
	}

	util.Logger.Infof("Fetched %d games from ethereum", len(gs))
	return gs, nil
}

// upload a new game to ethereum as well as its data from IPFS
func uploadToEthereum(g *games.Game) error {

	// upload data to IPFS
	util.Logger.Infof("Uploading game data for %s to IPFS", g.Title)
	err := g.UploadDataToIPFS()
	if err != nil {
		return err
	}
	util.Logger.Infof("Successfully uploaded game data for %s to IPFS", g.Title)

	// upload metadata to blockchain
	util.Logger.Infof("Uploading game metadata for %s to Ethereum", g.Title)
	_, err = lib_instance.UploadGame(auth_instance, library.LibraryGameEntry{
		Title:           g.Title,
		Version:         g.Version,
		ReleaseDate:     g.ReleaseDate,
		Developer:       g.Developer,
		RootHash:        g.RootHash,
		IpfsAddress:     g.IPFSId,
		PreviousVersion: g.PreviousVersion,
		Price:           g.Price,
		Uploader:        g.Uploader,
		Purchased:       []common.Address{},
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
		p := net.GetPeerInstance()
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
				p.GetLibrary().SetBlockchainGame(newGame.Hash, gameEntryToGame(&newGame.Game))
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

	lib := net.GetPeerInstance().GetLibrary()
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
		Title:           game.Title,
		Version:         game.Version,
		ReleaseDate:     game.ReleaseDate,
		Developer:       game.Developer,
		RootHash:        game.RootHash,
		IPFSId:          game.IpfsAddress,
		Uploader:        game.Uploader,
		Price:           game.Price,
		PreviousVersion: game.PreviousVersion,
	}

}
