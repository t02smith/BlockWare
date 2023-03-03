package ethereum

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
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
	eth_client *ethclient.Client
	once       sync.Once
)

// get the current ETH client
func GetClient() *ethclient.Client {
	return eth_client
}

// functions

// start the ethereum client
func StartClient(addr string) (*ethclient.Client, *accounts.Account, error) {
	var acc *accounts.Account
	once.Do(func() {
		client, err := ethclient.Dial(addr)
		if err != nil {
			util.Logger.Panic(err)
		}

		util.Logger.Infof("Connection to ETH network at %s made", addr)
		eth_client = client
	})

	return eth_client, acc, nil
}

func CloseEthClient() {
	eth_client.Close()
}

//

// creates a new auth instance given a private key for performing non gettet actions
func generateAuthInstance(privateKey string) (*bind.TransactOpts, error) {
	util.Logger.Info("Generating auth instance")
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

	util.Logger.Info("Auth instance generated")
	return auth, nil
}

// key store

// create a new keystore
func CreateKeyStore(keyStorePath string, password string) (*accounts.Account, error) {
	ks := keystore.NewKeyStore(keyStorePath, keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.NewAccount(password)
	if err != nil {
		return nil, err
	}

	return &account, nil
}
