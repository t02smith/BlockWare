package profileDeploy

import (
	"fmt"

	"github.com/t02smith/part-iii-project/toolkit/model/ethereum"
)

func Run(privateKey string) error {
	_, _, err := ethereum.StartClient("ws://localhost:8545")
	if err != nil {
		return fmt.Errorf("error starting eth client %s", err)
	}

	_, _, err = ethereum.DeployLibraryContract(privateKey)
	if err != nil {
		return fmt.Errorf("error connecting to lib instance %s", err)
	}

	return nil
}
