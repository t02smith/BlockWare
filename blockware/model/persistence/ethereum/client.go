package ethereum

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/t02smith/part-iii-project/toolkit/util"
)

/*

Open and maintain a connection with an ETH blockchain.
This will be used for interacting with our smart contract.

Our ETH client will be a singleton and expected to be persistent
for the entire runtime of the application

*/

var (
	_ethClient  *ethclient.Client
	_privateKey *ecdsa.PrivateKey
	once        sync.Once
)

// Client get the current ETH client
func Client() *ethclient.Client {
	return _ethClient
}

// functions

// StartClient start the ethereum client
func StartClient(addr string) (*ethclient.Client, *accounts.Account, error) {
	var acc *accounts.Account
	once.Do(func() {
		client, err := ethclient.Dial(addr)
		if err != nil {
			util.Logger.Fatal(err)
		}

		util.Logger.Infof("Connection to ETH network at %s made", addr)
		_ethClient = client
	})

	return _ethClient, acc, nil
}

func CloseEthClient() {
	_ethClient.Close()
}

//

// GenerateAuthInstance creates a new auth instance given a private key for performing non gettet actions
func GenerateAuthInstance(privateKey string) (*bind.TransactOpts, error) {
	util.Logger.Info("Generating auth instance")
	privKeyECDSA, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, err
	}
	_privateKey = privKeyECDSA

	pubKeyECDSA, ok := privKeyECDSA.Public().(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("invalid priv key")
	}

	fromAddress := crypto.PubkeyToAddress(*pubKeyECDSA)
	nonce, err := _ethClient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		util.Logger.Info(err)
	}

	util.Logger.Info("Getting chain ID")
	chainID, err := _ethClient.ChainID(context.Background())
	if err != nil {
		util.Logger.Info(err)
	}

	util.Logger.Info("Getting gas price")
	gasPrice, err := _ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		util.Logger.Info(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privKeyECDSA, chainID)
	if err != nil {
		util.Logger.Info(err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = gasPrice

	util.Logger.Info("Auth instance generated")
	return auth, nil
}

func Address() common.Address {
	if _privateKey == nil {
		return common.Address{}
	}

	return crypto.PubkeyToAddress(_privateKey.PublicKey)
}
