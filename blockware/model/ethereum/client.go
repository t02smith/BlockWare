package ethereum

import (
	"sync"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/t02smith/part-iii-project/toolkit/util"
)

var (
	eth_client *ethclient.Client
	once       sync.Once
)

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

// func GetKeyStore()
