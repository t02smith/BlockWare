// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package library

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// LibraryGameEntry is an auto generated low-level Go binding around an user-defined struct.
type LibraryGameEntry struct {
	Title               string
	Version             string
	ReleaseDate         string
	Developer           string
	RootHash            [32]byte
	PreviousVersion     [32]byte
	Price               *big.Int
	Uploader            common.Address
	HashTreeIPFSAddress string
	AssetsIPFSAddress   string
}

// LibraryMetaData contains all meta data concerning the Library contract.
var LibraryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"releaseDate\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"developer\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"previousVersion\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"uploader\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"hashTreeIPFSAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"assetsIPFSAddress\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structLibrary.GameEntry\",\"name\":\"game\",\"type\":\"tuple\"}],\"name\":\"NewGame\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"gameHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"games\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"releaseDate\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"developer\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"previousVersion\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"uploader\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"hashTreeIPFSAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"assetsIPFSAddress\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_game\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"hasPurchased\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"libSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_game\",\"type\":\"bytes32\"}],\"name\":\"purchaseGame\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"releaseDate\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"developer\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"previousVersion\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"uploader\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"hashTreeIPFSAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"assetsIPFSAddress\",\"type\":\"string\"}],\"internalType\":\"structLibrary.GameEntry\",\"name\":\"_game\",\"type\":\"tuple\"}],\"name\":\"uploadGame\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611df6806100206000396000f3fe6080604052600436106100555760003560e01c80630d0e12341461005a5780632d139a1b146100835780633e093f79146100c057806350e0c46e146100dc578063dc164c8214610107578063f579f88214610144575b600080fd5b34801561006657600080fd5b50610081600480360381019061007c91906112f1565b61018a565b005b34801561008f57600080fd5b506100aa60048036038101906100a5919061133a565b610970565b6040516100b79190611395565b60405180910390f35b6100da60048036038101906100d591906113b0565b6109dd565b005b3480156100e857600080fd5b506100f1610b50565b6040516100fe91906113ec565b60405180910390f35b34801561011357600080fd5b5061012e60048036038101906101299190611407565b610b5d565b60405161013b9190611443565b60405180910390f35b34801561015057600080fd5b5061016b600480360381019061016691906113b0565b610b81565b6040516101819a999897969594939291906114ec565b60405180910390f35b6000816080015150602060ff16116101d7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101ce906115fe565b60405180910390fd5b60008161010001515111610220576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161021790611690565b60405180910390fd5b60008161012001515111610269576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161026090611722565b60405180910390fd5b6000801b8160a00151146107c65760008060008360a001518152602001908152602001600020600001805461029d90611771565b9050116102df576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102d690611814565b60405180910390fd5b60008060008360a0015181526020019081526020016000206040518061014001604052908160008201805461031390611771565b80601f016020809104026020016040519081016040528092919081815260200182805461033f90611771565b801561038c5780601f106103615761010080835404028352916020019161038c565b820191906000526020600020905b81548152906001019060200180831161036f57829003601f168201915b505050505081526020016001820180546103a590611771565b80601f01602080910402602001604051908101604052809291908181526020018280546103d190611771565b801561041e5780601f106103f35761010080835404028352916020019161041e565b820191906000526020600020905b81548152906001019060200180831161040157829003601f168201915b5050505050815260200160028201805461043790611771565b80601f016020809104026020016040519081016040528092919081815260200182805461046390611771565b80156104b05780601f10610485576101008083540402835291602001916104b0565b820191906000526020600020905b81548152906001019060200180831161049357829003601f168201915b505050505081526020016003820180546104c990611771565b80601f01602080910402602001604051908101604052809291908181526020018280546104f590611771565b80156105425780601f1061051757610100808354040283529160200191610542565b820191906000526020600020905b81548152906001019060200180831161052557829003601f168201915b505050505081526020016004820154815260200160058201548152602001600682015481526020016007820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016008820180546105cf90611771565b80601f01602080910402602001604051908101604052809291908181526020018280546105fb90611771565b80156106485780601f1061061d57610100808354040283529160200191610648565b820191906000526020600020905b81548152906001019060200180831161062b57829003601f168201915b5050505050815260200160098201805461066190611771565b80601f016020809104026020016040519081016040528092919081815260200182805461068d90611771565b80156106da5780601f106106af576101008083540402835291602001916106da565b820191906000526020600020905b8154815290600101906020018083116106bd57829003601f168201915b50505050508152505090503373ffffffffffffffffffffffffffffffffffffffff168160e0015173ffffffffffffffffffffffffffffffffffffffff1614610757576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161074e906118a6565b60405180910390fd5b60018060008460800151815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908360ff160217905550505b338160e0019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508060008083608001518152602001908152602001600020600082015181600001908161082a9190611a72565b5060208201518160010190816108409190611a72565b5060408201518160020190816108569190611a72565b50606082015181600301908161086c9190611a72565b506080820151816004015560a0820151816005015560c0820151816006015560e08201518160070160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506101008201518160080190816108e89190611a72565b506101208201518160090190816108ff9190611a72565b509050506002816080015190806001815401808255809150506001900390600052602060002001600090919091909150557f0184d8edb4833799540b2e0c1ae01aba9d7355f6267e124ee0ed0f2ca94f084c816080015182604051610965929190611cb8565b60405180910390a150565b6000600180600085815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1660ff1614905092915050565b600080600083815260200190815260200160002060000180546109ff90611771565b905011610a41576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a3890611d34565b60405180910390fd5b60006001600083815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1660ff1614610ae4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610adb90611da0565b60405180910390fd5b600180600083815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908360ff16021790555050565b6000600280549050905090565b60028181548110610b6d57600080fd5b906000526020600020016000915090505481565b6000602052806000526040600020600091509050806000018054610ba490611771565b80601f0160208091040260200160405190810160405280929190818152602001828054610bd090611771565b8015610c1d5780601f10610bf257610100808354040283529160200191610c1d565b820191906000526020600020905b815481529060010190602001808311610c0057829003601f168201915b505050505090806001018054610c3290611771565b80601f0160208091040260200160405190810160405280929190818152602001828054610c5e90611771565b8015610cab5780601f10610c8057610100808354040283529160200191610cab565b820191906000526020600020905b815481529060010190602001808311610c8e57829003601f168201915b505050505090806002018054610cc090611771565b80601f0160208091040260200160405190810160405280929190818152602001828054610cec90611771565b8015610d395780601f10610d0e57610100808354040283529160200191610d39565b820191906000526020600020905b815481529060010190602001808311610d1c57829003601f168201915b505050505090806003018054610d4e90611771565b80601f0160208091040260200160405190810160405280929190818152602001828054610d7a90611771565b8015610dc75780601f10610d9c57610100808354040283529160200191610dc7565b820191906000526020600020905b815481529060010190602001808311610daa57829003601f168201915b5050505050908060040154908060050154908060060154908060070160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806008018054610e1490611771565b80601f0160208091040260200160405190810160405280929190818152602001828054610e4090611771565b8015610e8d5780601f10610e6257610100808354040283529160200191610e8d565b820191906000526020600020905b815481529060010190602001808311610e7057829003601f168201915b505050505090806009018054610ea290611771565b80601f0160208091040260200160405190810160405280929190818152602001828054610ece90611771565b8015610f1b5780601f10610ef057610100808354040283529160200191610f1b565b820191906000526020600020905b815481529060010190602001808311610efe57829003601f168201915b505050505090508a565b6000604051905090565b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610f8782610f3e565b810181811067ffffffffffffffff82111715610fa657610fa5610f4f565b5b80604052505050565b6000610fb9610f25565b9050610fc58282610f7e565b919050565b600080fd5b600080fd5b600080fd5b600067ffffffffffffffff821115610ff457610ff3610f4f565b5b610ffd82610f3e565b9050602081019050919050565b82818337600083830152505050565b600061102c61102784610fd9565b610faf565b90508281526020810184848401111561104857611047610fd4565b5b61105384828561100a565b509392505050565b600082601f8301126110705761106f610fcf565b5b8135611080848260208601611019565b91505092915050565b6000819050919050565b61109c81611089565b81146110a757600080fd5b50565b6000813590506110b981611093565b92915050565b6000819050919050565b6110d2816110bf565b81146110dd57600080fd5b50565b6000813590506110ef816110c9565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000611120826110f5565b9050919050565b61113081611115565b811461113b57600080fd5b50565b60008135905061114d81611127565b92915050565b6000610140828403121561116a57611169610f39565b5b611175610140610faf565b9050600082013567ffffffffffffffff81111561119557611194610fca565b5b6111a18482850161105b565b600083015250602082013567ffffffffffffffff8111156111c5576111c4610fca565b5b6111d18482850161105b565b602083015250604082013567ffffffffffffffff8111156111f5576111f4610fca565b5b6112018482850161105b565b604083015250606082013567ffffffffffffffff81111561122557611224610fca565b5b6112318482850161105b565b6060830152506080611245848285016110aa565b60808301525060a0611259848285016110aa565b60a08301525060c061126d848285016110e0565b60c08301525060e06112818482850161113e565b60e08301525061010082013567ffffffffffffffff8111156112a6576112a5610fca565b5b6112b28482850161105b565b6101008301525061012082013567ffffffffffffffff8111156112d8576112d7610fca565b5b6112e48482850161105b565b6101208301525092915050565b60006020828403121561130757611306610f2f565b5b600082013567ffffffffffffffff81111561132557611324610f34565b5b61133184828501611153565b91505092915050565b6000806040838503121561135157611350610f2f565b5b600061135f858286016110aa565b92505060206113708582860161113e565b9150509250929050565b60008115159050919050565b61138f8161137a565b82525050565b60006020820190506113aa6000830184611386565b92915050565b6000602082840312156113c6576113c5610f2f565b5b60006113d4848285016110aa565b91505092915050565b6113e6816110bf565b82525050565b600060208201905061140160008301846113dd565b92915050565b60006020828403121561141d5761141c610f2f565b5b600061142b848285016110e0565b91505092915050565b61143d81611089565b82525050565b60006020820190506114586000830184611434565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561149857808201518184015260208101905061147d565b60008484015250505050565b60006114af8261145e565b6114b98185611469565b93506114c981856020860161147a565b6114d281610f3e565b840191505092915050565b6114e681611115565b82525050565b6000610140820190508181036000830152611507818d6114a4565b9050818103602083015261151b818c6114a4565b9050818103604083015261152f818b6114a4565b90508181036060830152611543818a6114a4565b90506115526080830189611434565b61155f60a0830188611434565b61156c60c08301876113dd565b61157960e08301866114dd565b81810361010083015261158c81856114a4565b90508181036101208301526115a181846114a4565b90509b9a5050505050505050505050565b7f6e6f20726f6f74206861736820676976656e0000000000000000000000000000600082015250565b60006115e8601283611469565b91506115f3826115b2565b602082019050919050565b60006020820190508181036000830152611617816115db565b9050919050565b7f6e6f2049504653206164647265737320676976656e20666f722068617368207460008201527f7265656500000000000000000000000000000000000000000000000000000000602082015250565b600061167a602483611469565b91506116858261161e565b604082019050919050565b600060208201905081810360008301526116a98161166d565b9050919050565b7f6e6f2049504653206164647265737320676976656e20666f722074686520617360008201527f7365747320000000000000000000000000000000000000000000000000000000602082015250565b600061170c602583611469565b9150611717826116b0565b604082019050919050565b6000602082019050818103600083015261173b816116ff565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061178957607f821691505b60208210810361179c5761179b611742565b5b50919050565b7f70726576696f75732076657273696f6e206f662067616d65206e6f7420666f7560008201527f6e64000000000000000000000000000000000000000000000000000000000000602082015250565b60006117fe602283611469565b9150611809826117a2565b604082019050919050565b6000602082019050818103600083015261182d816117f1565b9050919050565b7f6f6e6c7920746865206f726967696e616c2075706c6f616465722063616e207560008201527f70646174652074686569722067616d6500000000000000000000000000000000602082015250565b6000611890603083611469565b915061189b82611834565b604082019050919050565b600060208201905081810360008301526118bf81611883565b9050919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026119287fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826118eb565b61193286836118eb565b95508019841693508086168417925050509392505050565b6000819050919050565b600061196f61196a611965846110bf565b61194a565b6110bf565b9050919050565b6000819050919050565b61198983611954565b61199d61199582611976565b8484546118f8565b825550505050565b600090565b6119b26119a5565b6119bd818484611980565b505050565b5b818110156119e1576119d66000826119aa565b6001810190506119c3565b5050565b601f821115611a26576119f7816118c6565b611a00846118db565b81016020851015611a0f578190505b611a23611a1b856118db565b8301826119c2565b50505b505050565b600082821c905092915050565b6000611a4960001984600802611a2b565b1980831691505092915050565b6000611a628383611a38565b9150826002028217905092915050565b611a7b8261145e565b67ffffffffffffffff811115611a9457611a93610f4f565b5b611a9e8254611771565b611aa98282856119e5565b600060209050601f831160018114611adc5760008415611aca578287015190505b611ad48582611a56565b865550611b3c565b601f198416611aea866118c6565b60005b82811015611b1257848901518255600182019150602085019450602081019050611aed565b86831015611b2f5784890151611b2b601f891682611a38565b8355505b6001600288020188555050505b505050505050565b600082825260208201905092915050565b6000611b608261145e565b611b6a8185611b44565b9350611b7a81856020860161147a565b611b8381610f3e565b840191505092915050565b611b9781611089565b82525050565b611ba6816110bf565b82525050565b611bb581611115565b82525050565b6000610140830160008301518482036000860152611bd98282611b55565b91505060208301518482036020860152611bf38282611b55565b91505060408301518482036040860152611c0d8282611b55565b91505060608301518482036060860152611c278282611b55565b9150506080830151611c3c6080860182611b8e565b5060a0830151611c4f60a0860182611b8e565b5060c0830151611c6260c0860182611b9d565b5060e0830151611c7560e0860182611bac565b50610100830151848203610100860152611c8f8282611b55565b915050610120830151848203610120860152611cab8282611b55565b9150508091505092915050565b6000604082019050611ccd6000830185611434565b8181036020830152611cdf8184611bbb565b90509392505050565b7f67616d65206e6f7420666f756e64000000000000000000000000000000000000600082015250565b6000611d1e600e83611469565b9150611d2982611ce8565b602082019050919050565b60006020820190508181036000830152611d4d81611d11565b9050919050565b7f7573657220616c7265616479206f776e732067616d6500000000000000000000600082015250565b6000611d8a601683611469565b9150611d9582611d54565b602082019050919050565b60006020820190508181036000830152611db981611d7d565b905091905056fea26469706673582212201a21015e4210d50f3186ec150975677f785233a7c511b489298cf486d18566ed64736f6c63430008120033",
}

// LibraryABI is the input ABI used to generate the binding from.
// Deprecated: Use LibraryMetaData.ABI instead.
var LibraryABI = LibraryMetaData.ABI

// LibraryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LibraryMetaData.Bin instead.
var LibraryBin = LibraryMetaData.Bin

// DeployLibrary deploys a new Ethereum contract, binding an instance of Library to it.
func DeployLibrary(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Library, error) {
	parsed, err := LibraryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LibraryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Library{LibraryCaller: LibraryCaller{contract: contract}, LibraryTransactor: LibraryTransactor{contract: contract}, LibraryFilterer: LibraryFilterer{contract: contract}}, nil
}

// Library is an auto generated Go binding around an Ethereum contract.
type Library struct {
	LibraryCaller     // Read-only binding to the contract
	LibraryTransactor // Write-only binding to the contract
	LibraryFilterer   // Log filterer for contract events
}

// LibraryCaller is an auto generated read-only Go binding around an Ethereum contract.
type LibraryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibraryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LibraryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibraryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LibraryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibrarySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LibrarySession struct {
	Contract     *Library          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LibraryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LibraryCallerSession struct {
	Contract *LibraryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// LibraryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LibraryTransactorSession struct {
	Contract     *LibraryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// LibraryRaw is an auto generated low-level Go binding around an Ethereum contract.
type LibraryRaw struct {
	Contract *Library // Generic contract binding to access the raw methods on
}

// LibraryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LibraryCallerRaw struct {
	Contract *LibraryCaller // Generic read-only contract binding to access the raw methods on
}

// LibraryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LibraryTransactorRaw struct {
	Contract *LibraryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLibrary creates a new instance of Library, bound to a specific deployed contract.
func NewLibrary(address common.Address, backend bind.ContractBackend) (*Library, error) {
	contract, err := bindLibrary(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Library{LibraryCaller: LibraryCaller{contract: contract}, LibraryTransactor: LibraryTransactor{contract: contract}, LibraryFilterer: LibraryFilterer{contract: contract}}, nil
}

// NewLibraryCaller creates a new read-only instance of Library, bound to a specific deployed contract.
func NewLibraryCaller(address common.Address, caller bind.ContractCaller) (*LibraryCaller, error) {
	contract, err := bindLibrary(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LibraryCaller{contract: contract}, nil
}

// NewLibraryTransactor creates a new write-only instance of Library, bound to a specific deployed contract.
func NewLibraryTransactor(address common.Address, transactor bind.ContractTransactor) (*LibraryTransactor, error) {
	contract, err := bindLibrary(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LibraryTransactor{contract: contract}, nil
}

// NewLibraryFilterer creates a new log filterer instance of Library, bound to a specific deployed contract.
func NewLibraryFilterer(address common.Address, filterer bind.ContractFilterer) (*LibraryFilterer, error) {
	contract, err := bindLibrary(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LibraryFilterer{contract: contract}, nil
}

// bindLibrary binds a generic wrapper to an already deployed contract.
func bindLibrary(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LibraryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Library *LibraryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Library.Contract.LibraryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Library *LibraryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Library.Contract.LibraryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Library *LibraryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Library.Contract.LibraryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Library *LibraryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Library.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Library *LibraryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Library.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Library *LibraryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Library.Contract.contract.Transact(opts, method, params...)
}

// GameHashes is a free data retrieval call binding the contract method 0xdc164c82.
//
// Solidity: function gameHashes(uint256 ) view returns(bytes32)
func (_Library *LibraryCaller) GameHashes(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Library.contract.Call(opts, &out, "gameHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GameHashes is a free data retrieval call binding the contract method 0xdc164c82.
//
// Solidity: function gameHashes(uint256 ) view returns(bytes32)
func (_Library *LibrarySession) GameHashes(arg0 *big.Int) ([32]byte, error) {
	return _Library.Contract.GameHashes(&_Library.CallOpts, arg0)
}

// GameHashes is a free data retrieval call binding the contract method 0xdc164c82.
//
// Solidity: function gameHashes(uint256 ) view returns(bytes32)
func (_Library *LibraryCallerSession) GameHashes(arg0 *big.Int) ([32]byte, error) {
	return _Library.Contract.GameHashes(&_Library.CallOpts, arg0)
}

// Games is a free data retrieval call binding the contract method 0xf579f882.
//
// Solidity: function games(bytes32 ) view returns(string title, string version, string releaseDate, string developer, bytes32 rootHash, bytes32 previousVersion, uint256 price, address uploader, string hashTreeIPFSAddress, string assetsIPFSAddress)
func (_Library *LibraryCaller) Games(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Title               string
	Version             string
	ReleaseDate         string
	Developer           string
	RootHash            [32]byte
	PreviousVersion     [32]byte
	Price               *big.Int
	Uploader            common.Address
	HashTreeIPFSAddress string
	AssetsIPFSAddress   string
}, error) {
	var out []interface{}
	err := _Library.contract.Call(opts, &out, "games", arg0)

	outstruct := new(struct {
		Title               string
		Version             string
		ReleaseDate         string
		Developer           string
		RootHash            [32]byte
		PreviousVersion     [32]byte
		Price               *big.Int
		Uploader            common.Address
		HashTreeIPFSAddress string
		AssetsIPFSAddress   string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Title = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.ReleaseDate = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Developer = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.RootHash = *abi.ConvertType(out[4], new([32]byte)).(*[32]byte)
	outstruct.PreviousVersion = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Price = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Uploader = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)
	outstruct.HashTreeIPFSAddress = *abi.ConvertType(out[8], new(string)).(*string)
	outstruct.AssetsIPFSAddress = *abi.ConvertType(out[9], new(string)).(*string)

	return *outstruct, err

}

// Games is a free data retrieval call binding the contract method 0xf579f882.
//
// Solidity: function games(bytes32 ) view returns(string title, string version, string releaseDate, string developer, bytes32 rootHash, bytes32 previousVersion, uint256 price, address uploader, string hashTreeIPFSAddress, string assetsIPFSAddress)
func (_Library *LibrarySession) Games(arg0 [32]byte) (struct {
	Title               string
	Version             string
	ReleaseDate         string
	Developer           string
	RootHash            [32]byte
	PreviousVersion     [32]byte
	Price               *big.Int
	Uploader            common.Address
	HashTreeIPFSAddress string
	AssetsIPFSAddress   string
}, error) {
	return _Library.Contract.Games(&_Library.CallOpts, arg0)
}

// Games is a free data retrieval call binding the contract method 0xf579f882.
//
// Solidity: function games(bytes32 ) view returns(string title, string version, string releaseDate, string developer, bytes32 rootHash, bytes32 previousVersion, uint256 price, address uploader, string hashTreeIPFSAddress, string assetsIPFSAddress)
func (_Library *LibraryCallerSession) Games(arg0 [32]byte) (struct {
	Title               string
	Version             string
	ReleaseDate         string
	Developer           string
	RootHash            [32]byte
	PreviousVersion     [32]byte
	Price               *big.Int
	Uploader            common.Address
	HashTreeIPFSAddress string
	AssetsIPFSAddress   string
}, error) {
	return _Library.Contract.Games(&_Library.CallOpts, arg0)
}

// HasPurchased is a free data retrieval call binding the contract method 0x2d139a1b.
//
// Solidity: function hasPurchased(bytes32 _game, address _addr) view returns(bool)
func (_Library *LibraryCaller) HasPurchased(opts *bind.CallOpts, _game [32]byte, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _Library.contract.Call(opts, &out, "hasPurchased", _game, _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasPurchased is a free data retrieval call binding the contract method 0x2d139a1b.
//
// Solidity: function hasPurchased(bytes32 _game, address _addr) view returns(bool)
func (_Library *LibrarySession) HasPurchased(_game [32]byte, _addr common.Address) (bool, error) {
	return _Library.Contract.HasPurchased(&_Library.CallOpts, _game, _addr)
}

// HasPurchased is a free data retrieval call binding the contract method 0x2d139a1b.
//
// Solidity: function hasPurchased(bytes32 _game, address _addr) view returns(bool)
func (_Library *LibraryCallerSession) HasPurchased(_game [32]byte, _addr common.Address) (bool, error) {
	return _Library.Contract.HasPurchased(&_Library.CallOpts, _game, _addr)
}

// LibSize is a free data retrieval call binding the contract method 0x50e0c46e.
//
// Solidity: function libSize() view returns(uint256)
func (_Library *LibraryCaller) LibSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Library.contract.Call(opts, &out, "libSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LibSize is a free data retrieval call binding the contract method 0x50e0c46e.
//
// Solidity: function libSize() view returns(uint256)
func (_Library *LibrarySession) LibSize() (*big.Int, error) {
	return _Library.Contract.LibSize(&_Library.CallOpts)
}

// LibSize is a free data retrieval call binding the contract method 0x50e0c46e.
//
// Solidity: function libSize() view returns(uint256)
func (_Library *LibraryCallerSession) LibSize() (*big.Int, error) {
	return _Library.Contract.LibSize(&_Library.CallOpts)
}

// PurchaseGame is a paid mutator transaction binding the contract method 0x3e093f79.
//
// Solidity: function purchaseGame(bytes32 _game) payable returns()
func (_Library *LibraryTransactor) PurchaseGame(opts *bind.TransactOpts, _game [32]byte) (*types.Transaction, error) {
	return _Library.contract.Transact(opts, "purchaseGame", _game)
}

// PurchaseGame is a paid mutator transaction binding the contract method 0x3e093f79.
//
// Solidity: function purchaseGame(bytes32 _game) payable returns()
func (_Library *LibrarySession) PurchaseGame(_game [32]byte) (*types.Transaction, error) {
	return _Library.Contract.PurchaseGame(&_Library.TransactOpts, _game)
}

// PurchaseGame is a paid mutator transaction binding the contract method 0x3e093f79.
//
// Solidity: function purchaseGame(bytes32 _game) payable returns()
func (_Library *LibraryTransactorSession) PurchaseGame(_game [32]byte) (*types.Transaction, error) {
	return _Library.Contract.PurchaseGame(&_Library.TransactOpts, _game)
}

// UploadGame is a paid mutator transaction binding the contract method 0x0d0e1234.
//
// Solidity: function uploadGame((string,string,string,string,bytes32,bytes32,uint256,address,string,string) _game) returns()
func (_Library *LibraryTransactor) UploadGame(opts *bind.TransactOpts, _game LibraryGameEntry) (*types.Transaction, error) {
	return _Library.contract.Transact(opts, "uploadGame", _game)
}

// UploadGame is a paid mutator transaction binding the contract method 0x0d0e1234.
//
// Solidity: function uploadGame((string,string,string,string,bytes32,bytes32,uint256,address,string,string) _game) returns()
func (_Library *LibrarySession) UploadGame(_game LibraryGameEntry) (*types.Transaction, error) {
	return _Library.Contract.UploadGame(&_Library.TransactOpts, _game)
}

// UploadGame is a paid mutator transaction binding the contract method 0x0d0e1234.
//
// Solidity: function uploadGame((string,string,string,string,bytes32,bytes32,uint256,address,string,string) _game) returns()
func (_Library *LibraryTransactorSession) UploadGame(_game LibraryGameEntry) (*types.Transaction, error) {
	return _Library.Contract.UploadGame(&_Library.TransactOpts, _game)
}

// LibraryNewGameIterator is returned from FilterNewGame and is used to iterate over the raw logs and unpacked data for NewGame events raised by the Library contract.
type LibraryNewGameIterator struct {
	Event *LibraryNewGame // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LibraryNewGameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LibraryNewGame)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LibraryNewGame)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LibraryNewGameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LibraryNewGameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LibraryNewGame represents a NewGame event raised by the Library contract.
type LibraryNewGame struct {
	Hash [32]byte
	Game LibraryGameEntry
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNewGame is a free log retrieval operation binding the contract event 0x0184d8edb4833799540b2e0c1ae01aba9d7355f6267e124ee0ed0f2ca94f084c.
//
// Solidity: event NewGame(bytes32 hash, (string,string,string,string,bytes32,bytes32,uint256,address,string,string) game)
func (_Library *LibraryFilterer) FilterNewGame(opts *bind.FilterOpts) (*LibraryNewGameIterator, error) {

	logs, sub, err := _Library.contract.FilterLogs(opts, "NewGame")
	if err != nil {
		return nil, err
	}
	return &LibraryNewGameIterator{contract: _Library.contract, event: "NewGame", logs: logs, sub: sub}, nil
}

// WatchNewGame is a free log subscription operation binding the contract event 0x0184d8edb4833799540b2e0c1ae01aba9d7355f6267e124ee0ed0f2ca94f084c.
//
// Solidity: event NewGame(bytes32 hash, (string,string,string,string,bytes32,bytes32,uint256,address,string,string) game)
func (_Library *LibraryFilterer) WatchNewGame(opts *bind.WatchOpts, sink chan<- *LibraryNewGame) (event.Subscription, error) {

	logs, sub, err := _Library.contract.WatchLogs(opts, "NewGame")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LibraryNewGame)
				if err := _Library.contract.UnpackLog(event, "NewGame", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewGame is a log parse operation binding the contract event 0x0184d8edb4833799540b2e0c1ae01aba9d7355f6267e124ee0ed0f2ca94f084c.
//
// Solidity: event NewGame(bytes32 hash, (string,string,string,string,bytes32,bytes32,uint256,address,string,string) game)
func (_Library *LibraryFilterer) ParseNewGame(log types.Log) (*LibraryNewGame, error) {
	event := new(LibraryNewGame)
	if err := _Library.contract.UnpackLog(event, "NewGame", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
