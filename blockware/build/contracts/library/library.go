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
	Title           string
	Version         string
	ReleaseDate     string
	Developer       string
	RootHash        [32]byte
	PreviousVersion [32]byte
	Price           *big.Int
	Uploader        common.Address
	IpfsAddress     string
}

// LibraryMetaData contains all meta data concerning the Library contract.
var LibraryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"releaseDate\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"developer\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"previousVersion\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"uploader\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ipfsAddress\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structLibrary.GameEntry\",\"name\":\"game\",\"type\":\"tuple\"}],\"name\":\"NewGame\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"gameHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"games\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"releaseDate\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"developer\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"previousVersion\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"uploader\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ipfsAddress\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_game\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"hasPurchased\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"libSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_game\",\"type\":\"bytes32\"}],\"name\":\"purchaseGame\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"releaseDate\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"developer\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"previousVersion\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"uploader\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ipfsAddress\",\"type\":\"string\"}],\"internalType\":\"structLibrary.GameEntry\",\"name\":\"_game\",\"type\":\"tuple\"}],\"name\":\"uploadGame\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611ce9806100206000396000f3fe6080604052600436106100555760003560e01c80632d139a1b1461005a5780633e093f791461009757806350e0c46e146100b357806388729388146100de578063dc164c8214610107578063f579f88214610144575b600080fd5b34801561006657600080fd5b50610081600480360381019061007c9190610f0c565b610189565b60405161008e9190610f67565b60405180910390f35b6100b160048036038101906100ac9190610f82565b6101f1565b005b3480156100bf57600080fd5b506100c861042a565b6040516100d59190610fc8565b60405180910390f35b3480156100ea57600080fd5b5061010560048036038101906101009190611309565b610437565b005b34801561011357600080fd5b5061012e60048036038101906101299190611352565b610b2a565b60405161013b919061138e565b60405180910390f35b34801561015057600080fd5b5061016b60048036038101906101669190610f82565b610b4e565b60405161018099989796959493929190611437565b60405180910390f35b60006001600084815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b6000806000838152602001908152602001600020600001805461021390611516565b905011610255576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161024c90611593565b60405180910390fd5b6000806000838152602001908152602001600020905080600601543410156102b2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102a9906115ff565b60405180910390fd5b6001600083815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1661034f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103469061166b565b60405180910390fd5b8060070160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc82600601549081150290604051600060405180830381858888f193505050501580156103bd573d6000803e3d6000fd5b50600180600084815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b6000600280549050905090565b6000816080015150602060ff1611610484576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161047b906116d7565b60405180910390fd5b600081610100015151116104cd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104c490611769565b60405180910390fd5b6000801b8160a00151146109975760008060008360a001518152602001908152602001600020600001805461050190611516565b905011610543576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161053a906117fb565b60405180910390fd5b60008060008360a0015181526020019081526020016000206040518061012001604052908160008201805461057790611516565b80601f01602080910402602001604051908101604052809291908181526020018280546105a390611516565b80156105f05780601f106105c5576101008083540402835291602001916105f0565b820191906000526020600020905b8154815290600101906020018083116105d357829003601f168201915b5050505050815260200160018201805461060990611516565b80601f016020809104026020016040519081016040528092919081815260200182805461063590611516565b80156106825780601f1061065757610100808354040283529160200191610682565b820191906000526020600020905b81548152906001019060200180831161066557829003601f168201915b5050505050815260200160028201805461069b90611516565b80601f01602080910402602001604051908101604052809291908181526020018280546106c790611516565b80156107145780601f106106e957610100808354040283529160200191610714565b820191906000526020600020905b8154815290600101906020018083116106f757829003601f168201915b5050505050815260200160038201805461072d90611516565b80601f016020809104026020016040519081016040528092919081815260200182805461075990611516565b80156107a65780601f1061077b576101008083540402835291602001916107a6565b820191906000526020600020905b81548152906001019060200180831161078957829003601f168201915b505050505081526020016004820154815260200160058201548152602001600682015481526020016007820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160088201805461083390611516565b80601f016020809104026020016040519081016040528092919081815260200182805461085f90611516565b80156108ac5780601f10610881576101008083540402835291602001916108ac565b820191906000526020600020905b81548152906001019060200180831161088f57829003601f168201915b50505050508152505090503373ffffffffffffffffffffffffffffffffffffffff168160e0015173ffffffffffffffffffffffffffffffffffffffff1614610929576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109209061188d565b60405180910390fd5b60018060008460800151815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550505b338160e0019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050806000808360800151815260200190815260200160002060008201518160000190816109fb9190611a59565b506020820151816001019081610a119190611a59565b506040820151816002019081610a279190611a59565b506060820151816003019081610a3d9190611a59565b506080820151816004015560a0820151816005015560c0820151816006015560e08201518160070160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610100820151816008019081610ab99190611a59565b509050506002816080015190806001815401808255809150506001900390600052602060002001600090919091909150557f49513426f0a3983e105aed76f3fdd2e5d9c257899e57a502a07512d9218d6422816080015182604051610b1f929190611c83565b60405180910390a150565b60028181548110610b3a57600080fd5b906000526020600020016000915090505481565b6000602052806000526040600020600091509050806000018054610b7190611516565b80601f0160208091040260200160405190810160405280929190818152602001828054610b9d90611516565b8015610bea5780601f10610bbf57610100808354040283529160200191610bea565b820191906000526020600020905b815481529060010190602001808311610bcd57829003601f168201915b505050505090806001018054610bff90611516565b80601f0160208091040260200160405190810160405280929190818152602001828054610c2b90611516565b8015610c785780601f10610c4d57610100808354040283529160200191610c78565b820191906000526020600020905b815481529060010190602001808311610c5b57829003601f168201915b505050505090806002018054610c8d90611516565b80601f0160208091040260200160405190810160405280929190818152602001828054610cb990611516565b8015610d065780601f10610cdb57610100808354040283529160200191610d06565b820191906000526020600020905b815481529060010190602001808311610ce957829003601f168201915b505050505090806003018054610d1b90611516565b80601f0160208091040260200160405190810160405280929190818152602001828054610d4790611516565b8015610d945780601f10610d6957610100808354040283529160200191610d94565b820191906000526020600020905b815481529060010190602001808311610d7757829003601f168201915b5050505050908060040154908060050154908060060154908060070160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806008018054610de190611516565b80601f0160208091040260200160405190810160405280929190818152602001828054610e0d90611516565b8015610e5a5780601f10610e2f57610100808354040283529160200191610e5a565b820191906000526020600020905b815481529060010190602001808311610e3d57829003601f168201915b5050505050905089565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b610e8b81610e78565b8114610e9657600080fd5b50565b600081359050610ea881610e82565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610ed982610eae565b9050919050565b610ee981610ece565b8114610ef457600080fd5b50565b600081359050610f0681610ee0565b92915050565b60008060408385031215610f2357610f22610e6e565b5b6000610f3185828601610e99565b9250506020610f4285828601610ef7565b9150509250929050565b60008115159050919050565b610f6181610f4c565b82525050565b6000602082019050610f7c6000830184610f58565b92915050565b600060208284031215610f9857610f97610e6e565b5b6000610fa684828501610e99565b91505092915050565b6000819050919050565b610fc281610faf565b82525050565b6000602082019050610fdd6000830184610fb9565b92915050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61103182610fe8565b810181811067ffffffffffffffff821117156110505761104f610ff9565b5b80604052505050565b6000611063610e64565b905061106f8282611028565b919050565b600080fd5b600080fd5b600080fd5b600067ffffffffffffffff82111561109e5761109d610ff9565b5b6110a782610fe8565b9050602081019050919050565b82818337600083830152505050565b60006110d66110d184611083565b611059565b9050828152602081018484840111156110f2576110f161107e565b5b6110fd8482856110b4565b509392505050565b600082601f83011261111a57611119611079565b5b813561112a8482602086016110c3565b91505092915050565b61113c81610faf565b811461114757600080fd5b50565b60008135905061115981611133565b92915050565b600061116a82610eae565b9050919050565b61117a8161115f565b811461118557600080fd5b50565b60008135905061119781611171565b92915050565b600061012082840312156111b4576111b3610fe3565b5b6111bf610120611059565b9050600082013567ffffffffffffffff8111156111df576111de611074565b5b6111eb84828501611105565b600083015250602082013567ffffffffffffffff81111561120f5761120e611074565b5b61121b84828501611105565b602083015250604082013567ffffffffffffffff81111561123f5761123e611074565b5b61124b84828501611105565b604083015250606082013567ffffffffffffffff81111561126f5761126e611074565b5b61127b84828501611105565b606083015250608061128f84828501610e99565b60808301525060a06112a384828501610e99565b60a08301525060c06112b78482850161114a565b60c08301525060e06112cb84828501611188565b60e08301525061010082013567ffffffffffffffff8111156112f0576112ef611074565b5b6112fc84828501611105565b6101008301525092915050565b60006020828403121561131f5761131e610e6e565b5b600082013567ffffffffffffffff81111561133d5761133c610e73565b5b6113498482850161119d565b91505092915050565b60006020828403121561136857611367610e6e565b5b60006113768482850161114a565b91505092915050565b61138881610e78565b82525050565b60006020820190506113a3600083018461137f565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156113e35780820151818401526020810190506113c8565b60008484015250505050565b60006113fa826113a9565b61140481856113b4565b93506114148185602086016113c5565b61141d81610fe8565b840191505092915050565b6114318161115f565b82525050565b6000610120820190508181036000830152611452818c6113ef565b90508181036020830152611466818b6113ef565b9050818103604083015261147a818a6113ef565b9050818103606083015261148e81896113ef565b905061149d608083018861137f565b6114aa60a083018761137f565b6114b760c0830186610fb9565b6114c460e0830185611428565b8181036101008301526114d781846113ef565b90509a9950505050505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061152e57607f821691505b602082108103611541576115406114e7565b5b50919050565b7f67616d65206e6f7420666f756e64000000000000000000000000000000000000600082015250565b600061157d600e836113b4565b915061158882611547565b602082019050919050565b600060208201905081810360008301526115ac81611570565b9050919050565b7f757365722063616e6e6f74206166666f72642067616d65000000000000000000600082015250565b60006115e96017836113b4565b91506115f4826115b3565b602082019050919050565b60006020820190508181036000830152611618816115dc565b9050919050565b7f7573657220616c7265616479206f776e732067616d6500000000000000000000600082015250565b60006116556016836113b4565b91506116608261161f565b602082019050919050565b6000602082019050818103600083015261168481611648565b9050919050565b7f6e6f20726f6f74206861736820676976656e0000000000000000000000000000600082015250565b60006116c16012836113b4565b91506116cc8261168b565b602082019050919050565b600060208201905081810360008301526116f0816116b4565b9050919050565b7f6e6f2049504653206164647265737320676976656e20666f722068617368207460008201527f7265656500000000000000000000000000000000000000000000000000000000602082015250565b60006117536024836113b4565b915061175e826116f7565b604082019050919050565b6000602082019050818103600083015261178281611746565b9050919050565b7f70726576696f75732076657273696f6e206f662067616d65206e6f7420666f7560008201527f6e64000000000000000000000000000000000000000000000000000000000000602082015250565b60006117e56022836113b4565b91506117f082611789565b604082019050919050565b60006020820190508181036000830152611814816117d8565b9050919050565b7f6f6e6c7920746865206f726967696e616c2075706c6f616465722063616e207560008201527f70646174652074686569722067616d6500000000000000000000000000000000602082015250565b60006118776030836113b4565b91506118828261181b565b604082019050919050565b600060208201905081810360008301526118a68161186a565b9050919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b60006008830261190f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826118d2565b61191986836118d2565b95508019841693508086168417925050509392505050565b6000819050919050565b600061195661195161194c84610faf565b611931565b610faf565b9050919050565b6000819050919050565b6119708361193b565b61198461197c8261195d565b8484546118df565b825550505050565b600090565b61199961198c565b6119a4818484611967565b505050565b5b818110156119c8576119bd600082611991565b6001810190506119aa565b5050565b601f821115611a0d576119de816118ad565b6119e7846118c2565b810160208510156119f6578190505b611a0a611a02856118c2565b8301826119a9565b50505b505050565b600082821c905092915050565b6000611a3060001984600802611a12565b1980831691505092915050565b6000611a498383611a1f565b9150826002028217905092915050565b611a62826113a9565b67ffffffffffffffff811115611a7b57611a7a610ff9565b5b611a858254611516565b611a908282856119cc565b600060209050601f831160018114611ac35760008415611ab1578287015190505b611abb8582611a3d565b865550611b23565b601f198416611ad1866118ad565b60005b82811015611af957848901518255600182019150602085019450602081019050611ad4565b86831015611b165784890151611b12601f891682611a1f565b8355505b6001600288020188555050505b505050505050565b600082825260208201905092915050565b6000611b47826113a9565b611b518185611b2b565b9350611b618185602086016113c5565b611b6a81610fe8565b840191505092915050565b611b7e81610e78565b82525050565b611b8d81610faf565b82525050565b611b9c8161115f565b82525050565b6000610120830160008301518482036000860152611bc08282611b3c565b91505060208301518482036020860152611bda8282611b3c565b91505060408301518482036040860152611bf48282611b3c565b91505060608301518482036060860152611c0e8282611b3c565b9150506080830151611c236080860182611b75565b5060a0830151611c3660a0860182611b75565b5060c0830151611c4960c0860182611b84565b5060e0830151611c5c60e0860182611b93565b50610100830151848203610100860152611c768282611b3c565b9150508091505092915050565b6000604082019050611c98600083018561137f565b8181036020830152611caa8184611ba2565b9050939250505056fea26469706673582212206d74e2a66c02e8bbdc91143437636c4a2a92dfec0bec4fe9af60258505ad21e164736f6c63430008120033",
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
// Solidity: function games(bytes32 ) view returns(string title, string version, string releaseDate, string developer, bytes32 rootHash, bytes32 previousVersion, uint256 price, address uploader, string ipfsAddress)
func (_Library *LibraryCaller) Games(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Title           string
	Version         string
	ReleaseDate     string
	Developer       string
	RootHash        [32]byte
	PreviousVersion [32]byte
	Price           *big.Int
	Uploader        common.Address
	IpfsAddress     string
}, error) {
	var out []interface{}
	err := _Library.contract.Call(opts, &out, "games", arg0)

	outstruct := new(struct {
		Title           string
		Version         string
		ReleaseDate     string
		Developer       string
		RootHash        [32]byte
		PreviousVersion [32]byte
		Price           *big.Int
		Uploader        common.Address
		IpfsAddress     string
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
	outstruct.IpfsAddress = *abi.ConvertType(out[8], new(string)).(*string)

	return *outstruct, err

}

// Games is a free data retrieval call binding the contract method 0xf579f882.
//
// Solidity: function games(bytes32 ) view returns(string title, string version, string releaseDate, string developer, bytes32 rootHash, bytes32 previousVersion, uint256 price, address uploader, string ipfsAddress)
func (_Library *LibrarySession) Games(arg0 [32]byte) (struct {
	Title           string
	Version         string
	ReleaseDate     string
	Developer       string
	RootHash        [32]byte
	PreviousVersion [32]byte
	Price           *big.Int
	Uploader        common.Address
	IpfsAddress     string
}, error) {
	return _Library.Contract.Games(&_Library.CallOpts, arg0)
}

// Games is a free data retrieval call binding the contract method 0xf579f882.
//
// Solidity: function games(bytes32 ) view returns(string title, string version, string releaseDate, string developer, bytes32 rootHash, bytes32 previousVersion, uint256 price, address uploader, string ipfsAddress)
func (_Library *LibraryCallerSession) Games(arg0 [32]byte) (struct {
	Title           string
	Version         string
	ReleaseDate     string
	Developer       string
	RootHash        [32]byte
	PreviousVersion [32]byte
	Price           *big.Int
	Uploader        common.Address
	IpfsAddress     string
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

// UploadGame is a paid mutator transaction binding the contract method 0x88729388.
//
// Solidity: function uploadGame((string,string,string,string,bytes32,bytes32,uint256,address,string) _game) returns()
func (_Library *LibraryTransactor) UploadGame(opts *bind.TransactOpts, _game LibraryGameEntry) (*types.Transaction, error) {
	return _Library.contract.Transact(opts, "uploadGame", _game)
}

// UploadGame is a paid mutator transaction binding the contract method 0x88729388.
//
// Solidity: function uploadGame((string,string,string,string,bytes32,bytes32,uint256,address,string) _game) returns()
func (_Library *LibrarySession) UploadGame(_game LibraryGameEntry) (*types.Transaction, error) {
	return _Library.Contract.UploadGame(&_Library.TransactOpts, _game)
}

// UploadGame is a paid mutator transaction binding the contract method 0x88729388.
//
// Solidity: function uploadGame((string,string,string,string,bytes32,bytes32,uint256,address,string) _game) returns()
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

// FilterNewGame is a free log retrieval operation binding the contract event 0x49513426f0a3983e105aed76f3fdd2e5d9c257899e57a502a07512d9218d6422.
//
// Solidity: event NewGame(bytes32 hash, (string,string,string,string,bytes32,bytes32,uint256,address,string) game)
func (_Library *LibraryFilterer) FilterNewGame(opts *bind.FilterOpts) (*LibraryNewGameIterator, error) {

	logs, sub, err := _Library.contract.FilterLogs(opts, "NewGame")
	if err != nil {
		return nil, err
	}
	return &LibraryNewGameIterator{contract: _Library.contract, event: "NewGame", logs: logs, sub: sub}, nil
}

// WatchNewGame is a free log subscription operation binding the contract event 0x49513426f0a3983e105aed76f3fdd2e5d9c257899e57a502a07512d9218d6422.
//
// Solidity: event NewGame(bytes32 hash, (string,string,string,string,bytes32,bytes32,uint256,address,string) game)
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

// ParseNewGame is a log parse operation binding the contract event 0x49513426f0a3983e105aed76f3fdd2e5d9c257899e57a502a07512d9218d6422.
//
// Solidity: event NewGame(bytes32 hash, (string,string,string,string,bytes32,bytes32,uint256,address,string) game)
func (_Library *LibraryFilterer) ParseNewGame(log types.Log) (*LibraryNewGame, error) {
	event := new(LibraryNewGame)
	if err := _Library.contract.UnpackLog(event, "NewGame", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
