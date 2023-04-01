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

	_, _, err = library.DeployLibraryContract("af9668cd6ebc3ba4c0e5036c284e128ed66e18ba9e4ed87b2c0c6d9642f2b879")
	if err != nil {
		return fmt.Errorf("error connecting to lib instance %s", err)
	}

	return nil
}
