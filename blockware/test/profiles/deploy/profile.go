package profileDeploy

import (
	"fmt"

	"github.com/t02smith/part-iii-project/toolkit/model/persistence/ethereum"
	"github.com/t02smith/part-iii-project/toolkit/model/persistence/ethereum/library"
)

func Run(privateKey string) error {
	_, _, err := ethereum.StartClient("http://localhost:8545")
	if err != nil {
		return fmt.Errorf("error starting eth client %s", err)
	}

	_, _, err = library.DeployLibraryContract("b3868cc6652c9279c088be8dbfe6f4ef2ab39ecc80b3c69602217fac64ed6ad4")
	if err != nil {
		return fmt.Errorf("error connecting to lib instance %s", err)
	}

	return nil
}
