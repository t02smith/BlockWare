package ethereum

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/t02smith/part-iii-project/toolkit/build/contracts"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

func StartClient(addr string) (*ethclient.Client, error) {
	client, err := ethclient.Dial(addr)
	if err != nil {
		util.Logger.Errorf("Error connecting to ETH network: %s", err)
		return nil, err
	}

	util.Logger.Infof("Connection to ETH network at %s made", addr)
	return client, nil
}

func DeployContracts(client *ethclient.Client, privateKey string) error {

	privKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return err
	}

	pubKey := privKey.Public()
	pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("public key is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*pubKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	chainID, err := client.ChainID(context.TODO())
	if err != nil {
		return err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privKey, chainID)
	if err != nil {
		return err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(0)
	auth.GasPrice = gasPrice

	input := "1.0"
	address, tx, instance, err := contracts.DeployContracts(auth, client, input)
	if err != nil {
		return err
	}

	fmt.Println(address.Hex())
	fmt.Println(tx.Hash().Hex())
	_ = instance

	return nil
}
