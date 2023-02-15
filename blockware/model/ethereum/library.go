package ethereum

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/t02smith/part-iii-project/toolkit/build/contracts/library"
	"github.com/t02smith/part-iii-project/toolkit/model/games"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

var (
	lib_instance  *library.Library
	auth_instance *bind.TransactOpts
	lib_auth_once sync.Once
)

func DeployLibraryContract(privateKey string) (*bind.TransactOpts, *library.Library, error) {
	if eth_client == nil {
		return nil, nil, errors.New("you need to instantiate the Eth Client => run StartClient")
	}

	lib_auth_once.Do(func() {
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

		chainID, err := eth_client.ChainID(context.Background())
		if err != nil {
			util.Logger.Panic(err)
		}

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
		Title:       g.Title,
		Version:     g.Version,
		ReleaseDate: g.ReleaseDate,
		Developer:   g.Developer,
		RootHash:    g.RootHash,
		IpfsAddress: g.IPFSId,
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

	// ? is game already on eth
	onChain, err := isGameOnBlockchain(g.RootHash)
	if err != nil {
		return err
	}
	if onChain {
		return fmt.Errorf("game %s with hash %x already present on blockchain", g.Title, g.RootHash)
	}

	// * upload
	return uploadToEthereum(g)
}
