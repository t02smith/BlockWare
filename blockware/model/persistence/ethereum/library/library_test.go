package library

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/t02smith/part-iii-project/toolkit/model/persistence/ethereum"
	"github.com/t02smith/part-iii-project/toolkit/test/testutil"
)

/*

function: DeployLibraryContract

*/

func TestDeployLibraryContract(t *testing.T) {

	t.Run("no ethereum client", func(t *testing.T) {
		_, _, err := DeployLibraryContract("")
		assert.NotNil(t, err)
		assert.Equal(t, "you need to instantiate the Eth Client => run StartClient", err.Error())
	})

	_, _, err := ethereum.StartClient("http://localhost:8545")
	if err != nil {
		t.Fatal(err)
	}

	t.Run("failure", func(t *testing.T) {
		t.Run("invalid private key", func(t *testing.T) {
			_, _, err := DeployLibraryContract("")
			assert.NotNil(t, err)
			assert.Equal(t, "invalid length, need 256 bits", err.Error())
		})
	})

	t.Run("successful deployment", func(t *testing.T) {
		auth, lib, err := DeployLibraryContract(testutil.Accounts[0][1])
		assert.Nil(t, err)
		assert.NotNil(t, auth)
		assert.Equal(t, auth, authInstance)

		assert.NotNil(t, lib)
		assert.Equal(t, lib, libInstance)

		assert.NotEqual(t, common.Address{}, contractAddress)
	})
}

/*

function: ConnectToLibraryInstance

*/

func TestConnectToLibraryInstance(t *testing.T) {
	_, _, err := ethereum.StartClient("http://localhost:8545")
	if err != nil {
		t.Fatal(err)
	}

	deployContract(t)

	t.Run("errors", func(t *testing.T) {
		t.Run("invalid private key", func(t *testing.T) {
			err := ConnectToLibraryInstance(common.Address{}, "")
			assert.NotNil(t, err)
			assert.Equal(t, "invalid length, need 256 bits", err.Error())
		})
	})
}

/*


 */
