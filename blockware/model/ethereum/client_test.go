package ethereum

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/t02smith/part-iii-project/toolkit/model"
)

func TestStartClient(t *testing.T) {
	t.Skip()
	t.Run("success", func(t *testing.T) {
		client, _, err := StartClient("ws://localhost:8545")
		if err != nil {
			t.Fatalf("Failed to connect to network. Is ganache running?")
		}

		client.Close()
	})

	t.Run("unknown network", func(t *testing.T) {
		t.Skip() // TODO
		_, _, err := StartClient("ws://fake.t02smith.com")
		if err == nil {
			t.Fatalf("Ethereum network not expected to be at domain")
		}
	})

}

func TestCreateKeyStore(t *testing.T) {

	keyStorePath := "../../test/data/tmp/wallets"
	err := model.CreateDirectoryIfNotExist(keyStorePath)
	if err != nil {
		t.Fatalf("Error creating tmp folder")
	}

	t.Run("success", func(t *testing.T) {
		t.Skip() // TODO BROKEN

		password := "secret"
		acc, err := CreateKeyStore(keyStorePath, password)
		if err != nil {
			t.Fatal(err)
		}

		f, err := os.Open(acc.URL.Path)
		if err != nil {
			t.Fatal(err)
		}

		data, err := io.ReadAll(f)
		if err != nil {
			t.Fatal(err)
		}

		ks := keystore.NewKeyStore(keyStorePath, keystore.StandardScryptN, keystore.StandardScryptP)

		accFromFile, err := ks.Import(data, password, password)
		if err != nil {
			t.Fatal(err)
		}

		if !bytes.Equal(acc.Address.Hash().Bytes(), accFromFile.Address.Hash().Bytes()) {
			t.Errorf("accounts do not match")
		}

	})

}
