package library

import (
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/t02smith/part-iii-project/toolkit/build/contracts/library"
	"github.com/t02smith/part-iii-project/toolkit/test/testutil"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

func deployContract(t *testing.T) (*bind.TransactOpts, *library.Library) {
	ops, lib, err := DeployLibraryContract(testutil.Accounts[0][1])
	if err != nil {
		t.Fatalf("Err deploying contract. Have you run 'make ganache': %s", err)
	}

	return ops, lib
}

//

func beforeAll() {
	util.InitLogger(true)
	testutil.SetupTestConfig()

}

func TestMain(m *testing.M) {
	beforeAll()
	code := m.Run()

	testutil.ClearTmp("../../../")
	os.Exit(code)
}
