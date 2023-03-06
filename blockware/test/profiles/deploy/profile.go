package profileDeploy

import (
	"fmt"

	"github.com/t02smith/part-iii-project/toolkit/model/persistence/ethereum"
	"github.com/t02smith/part-iii-project/toolkit/model/persistence/ethereum/library"
)

func Run(privateKey string) error {
	_, _, err := ethereum.StartClient("ws://localhost:8545")
	if err != nil {
		return fmt.Errorf("error starting eth client %s", err)
	}

	_, _, err = library.DeployLibraryContract(privateKey)
	if err != nil {
		return fmt.Errorf("error connecting to lib instance %s", err)
	}

	return nil
}
