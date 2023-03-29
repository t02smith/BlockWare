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

	_, _, err = library.DeployLibraryContract("a368fddde78a49b3e415b280ebd36003b0aa93ab07bc97b5aabeaf70835fe778")
	if err != nil {
		return fmt.Errorf("error connecting to lib instance %s", err)
	}

	return nil
}
