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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"releaseDate\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"developer\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"previousVersion\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"uploader\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ipfsAddress\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structLibrary.GameEntry\",\"name\":\"game\",\"type\":\"tuple\"}],\"name\":\"NewGame\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"gameHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"games\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"releaseDate\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"developer\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"previousVersion\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"uploader\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ipfsAddress\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_game\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"hasPurchased\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"libSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_game\",\"type\":\"bytes32\"}],\"name\":\"purchaseGame\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"releaseDate\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"developer\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"previousVersion\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"uploader\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ipfsAddress\",\"type\":\"string\"}],\"internalType\":\"structLibrary.GameEntry\",\"name\":\"_game\",\"type\":\"tuple\"}],\"name\":\"uploadGame\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611b7f806100206000396000f3fe6080604052600436106100555760003560e01c80632d139a1b1461005a5780633e093f791461009757806350e0c46e146100b357806388729388146100de578063dc164c8214610107578063f579f88214610144575b600080fd5b34801561006657600080fd5b50610081600480360381019061007c9190610e4c565b610189565b60405161008e9190610ea7565b60405180910390f35b6100b160048036038101906100ac9190610ec2565b6101f6565b005b3480156100bf57600080fd5b506100c8610369565b6040516100d59190610f08565b60405180910390f35b3480156100ea57600080fd5b506101056004803603810190610100919061120b565b610376565b005b34801561011357600080fd5b5061012e60048036038101906101299190611254565b610a6a565b60405161013b9190611290565b60405180910390f35b34801561015057600080fd5b5061016b60048036038101906101669190610ec2565b610a8e565b60405161018099989796959493929190611339565b60405180910390f35b6000600180600085815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1660ff1614905092915050565b6000806000838152602001908152602001600020600001805461021890611418565b90501161025a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161025190611495565b60405180910390fd5b60006001600083815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1660ff16146102fd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102f490611501565b60405180910390fd5b600180600083815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908360ff16021790555050565b6000600280549050905090565b6000816080015150602060ff16116103c3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103ba9061156d565b60405180910390fd5b6000816101000151511161040c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610403906115ff565b60405180910390fd5b6000801b8160a00151146108d75760008060008360a001518152602001908152602001600020600001805461044090611418565b905011610482576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161047990611691565b60405180910390fd5b60008060008360a001518152602001908152602001600020604051806101200160405290816000820180546104b690611418565b80601f01602080910402602001604051908101604052809291908181526020018280546104e290611418565b801561052f5780601f106105045761010080835404028352916020019161052f565b820191906000526020600020905b81548152906001019060200180831161051257829003601f168201915b5050505050815260200160018201805461054890611418565b80601f016020809104026020016040519081016040528092919081815260200182805461057490611418565b80156105c15780601f10610596576101008083540402835291602001916105c1565b820191906000526020600020905b8154815290600101906020018083116105a457829003601f168201915b505050505081526020016002820180546105da90611418565b80601f016020809104026020016040519081016040528092919081815260200182805461060690611418565b80156106535780601f1061062857610100808354040283529160200191610653565b820191906000526020600020905b81548152906001019060200180831161063657829003601f168201915b5050505050815260200160038201805461066c90611418565b80601f016020809104026020016040519081016040528092919081815260200182805461069890611418565b80156106e55780601f106106ba576101008083540402835291602001916106e5565b820191906000526020600020905b8154815290600101906020018083116106c857829003601f168201915b505050505081526020016004820154815260200160058201548152602001600682015481526020016007820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160088201805461077290611418565b80601f016020809104026020016040519081016040528092919081815260200182805461079e90611418565b80156107eb5780601f106107c0576101008083540402835291602001916107eb565b820191906000526020600020905b8154815290600101906020018083116107ce57829003601f168201915b50505050508152505090503373ffffffffffffffffffffffffffffffffffffffff168160e0015173ffffffffffffffffffffffffffffffffffffffff1614610868576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161085f90611723565b60405180910390fd5b60018060008460800151815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908360ff160217905550505b338160e0019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508060008083608001518152602001908152602001600020600082015181600001908161093b91906118ef565b50602082015181600101908161095191906118ef565b50604082015181600201908161096791906118ef565b50606082015181600301908161097d91906118ef565b506080820151816004015560a0820151816005015560c0820151816006015560e08201518160070160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506101008201518160080190816109f991906118ef565b509050506002816080015190806001815401808255809150506001900390600052602060002001600090919091909150557f49513426f0a3983e105aed76f3fdd2e5d9c257899e57a502a07512d9218d6422816080015182604051610a5f929190611b19565b60405180910390a150565b60028181548110610a7a57600080fd5b906000526020600020016000915090505481565b6000602052806000526040600020600091509050806000018054610ab190611418565b80601f0160208091040260200160405190810160405280929190818152602001828054610add90611418565b8015610b2a5780601f10610aff57610100808354040283529160200191610b2a565b820191906000526020600020905b815481529060010190602001808311610b0d57829003601f168201915b505050505090806001018054610b3f90611418565b80601f0160208091040260200160405190810160405280929190818152602001828054610b6b90611418565b8015610bb85780601f10610b8d57610100808354040283529160200191610bb8565b820191906000526020600020905b815481529060010190602001808311610b9b57829003601f168201915b505050505090806002018054610bcd90611418565b80601f0160208091040260200160405190810160405280929190818152602001828054610bf990611418565b8015610c465780601f10610c1b57610100808354040283529160200191610c46565b820191906000526020600020905b815481529060010190602001808311610c2957829003601f168201915b505050505090806003018054610c5b90611418565b80601f0160208091040260200160405190810160405280929190818152602001828054610c8790611418565b8015610cd45780601f10610ca957610100808354040283529160200191610cd4565b820191906000526020600020905b815481529060010190602001808311610cb757829003601f168201915b5050505050908060040154908060050154908060060154908060070160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806008018054610d2190611418565b80601f0160208091040260200160405190810160405280929190818152602001828054610d4d90611418565b8015610d9a5780601f10610d6f57610100808354040283529160200191610d9a565b820191906000526020600020905b815481529060010190602001808311610d7d57829003601f168201915b5050505050905089565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b610dcb81610db8565b8114610dd657600080fd5b50565b600081359050610de881610dc2565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610e1982610dee565b9050919050565b610e2981610e0e565b8114610e3457600080fd5b50565b600081359050610e4681610e20565b92915050565b60008060408385031215610e6357610e62610dae565b5b6000610e7185828601610dd9565b9250506020610e8285828601610e37565b9150509250929050565b60008115159050919050565b610ea181610e8c565b82525050565b6000602082019050610ebc6000830184610e98565b92915050565b600060208284031215610ed857610ed7610dae565b5b6000610ee684828501610dd9565b91505092915050565b6000819050919050565b610f0281610eef565b82525050565b6000602082019050610f1d6000830184610ef9565b92915050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610f7182610f28565b810181811067ffffffffffffffff82111715610f9057610f8f610f39565b5b80604052505050565b6000610fa3610da4565b9050610faf8282610f68565b919050565b600080fd5b600080fd5b600080fd5b600067ffffffffffffffff821115610fde57610fdd610f39565b5b610fe782610f28565b9050602081019050919050565b82818337600083830152505050565b600061101661101184610fc3565b610f99565b90508281526020810184848401111561103257611031610fbe565b5b61103d848285610ff4565b509392505050565b600082601f83011261105a57611059610fb9565b5b813561106a848260208601611003565b91505092915050565b61107c81610eef565b811461108757600080fd5b50565b60008135905061109981611073565b92915050565b600061012082840312156110b6576110b5610f23565b5b6110c1610120610f99565b9050600082013567ffffffffffffffff8111156110e1576110e0610fb4565b5b6110ed84828501611045565b600083015250602082013567ffffffffffffffff81111561111157611110610fb4565b5b61111d84828501611045565b602083015250604082013567ffffffffffffffff81111561114157611140610fb4565b5b61114d84828501611045565b604083015250606082013567ffffffffffffffff81111561117157611170610fb4565b5b61117d84828501611045565b606083015250608061119184828501610dd9565b60808301525060a06111a584828501610dd9565b60a08301525060c06111b98482850161108a565b60c08301525060e06111cd84828501610e37565b60e08301525061010082013567ffffffffffffffff8111156111f2576111f1610fb4565b5b6111fe84828501611045565b6101008301525092915050565b60006020828403121561122157611220610dae565b5b600082013567ffffffffffffffff81111561123f5761123e610db3565b5b61124b8482850161109f565b91505092915050565b60006020828403121561126a57611269610dae565b5b60006112788482850161108a565b91505092915050565b61128a81610db8565b82525050565b60006020820190506112a56000830184611281565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156112e55780820151818401526020810190506112ca565b60008484015250505050565b60006112fc826112ab565b61130681856112b6565b93506113168185602086016112c7565b61131f81610f28565b840191505092915050565b61133381610e0e565b82525050565b6000610120820190508181036000830152611354818c6112f1565b90508181036020830152611368818b6112f1565b9050818103604083015261137c818a6112f1565b9050818103606083015261139081896112f1565b905061139f6080830188611281565b6113ac60a0830187611281565b6113b960c0830186610ef9565b6113c660e083018561132a565b8181036101008301526113d981846112f1565b90509a9950505050505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061143057607f821691505b602082108103611443576114426113e9565b5b50919050565b7f67616d65206e6f7420666f756e64000000000000000000000000000000000000600082015250565b600061147f600e836112b6565b915061148a82611449565b602082019050919050565b600060208201905081810360008301526114ae81611472565b9050919050565b7f7573657220616c7265616479206f776e732067616d6500000000000000000000600082015250565b60006114eb6016836112b6565b91506114f6826114b5565b602082019050919050565b6000602082019050818103600083015261151a816114de565b9050919050565b7f6e6f20726f6f74206861736820676976656e0000000000000000000000000000600082015250565b60006115576012836112b6565b915061156282611521565b602082019050919050565b600060208201905081810360008301526115868161154a565b9050919050565b7f6e6f2049504653206164647265737320676976656e20666f722068617368207460008201527f7265656500000000000000000000000000000000000000000000000000000000602082015250565b60006115e96024836112b6565b91506115f48261158d565b604082019050919050565b60006020820190508181036000830152611618816115dc565b9050919050565b7f70726576696f75732076657273696f6e206f662067616d65206e6f7420666f7560008201527f6e64000000000000000000000000000000000000000000000000000000000000602082015250565b600061167b6022836112b6565b91506116868261161f565b604082019050919050565b600060208201905081810360008301526116aa8161166e565b9050919050565b7f6f6e6c7920746865206f726967696e616c2075706c6f616465722063616e207560008201527f70646174652074686569722067616d6500000000000000000000000000000000602082015250565b600061170d6030836112b6565b9150611718826116b1565b604082019050919050565b6000602082019050818103600083015261173c81611700565b9050919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026117a57fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611768565b6117af8683611768565b95508019841693508086168417925050509392505050565b6000819050919050565b60006117ec6117e76117e284610eef565b6117c7565b610eef565b9050919050565b6000819050919050565b611806836117d1565b61181a611812826117f3565b848454611775565b825550505050565b600090565b61182f611822565b61183a8184846117fd565b505050565b5b8181101561185e57611853600082611827565b600181019050611840565b5050565b601f8211156118a35761187481611743565b61187d84611758565b8101602085101561188c578190505b6118a061189885611758565b83018261183f565b50505b505050565b600082821c905092915050565b60006118c6600019846008026118a8565b1980831691505092915050565b60006118df83836118b5565b9150826002028217905092915050565b6118f8826112ab565b67ffffffffffffffff81111561191157611910610f39565b5b61191b8254611418565b611926828285611862565b600060209050601f8311600181146119595760008415611947578287015190505b61195185826118d3565b8655506119b9565b601f19841661196786611743565b60005b8281101561198f5784890151825560018201915060208501945060208101905061196a565b868310156119ac57848901516119a8601f8916826118b5565b8355505b6001600288020188555050505b505050505050565b600082825260208201905092915050565b60006119dd826112ab565b6119e781856119c1565b93506119f78185602086016112c7565b611a0081610f28565b840191505092915050565b611a1481610db8565b82525050565b611a2381610eef565b82525050565b611a3281610e0e565b82525050565b6000610120830160008301518482036000860152611a5682826119d2565b91505060208301518482036020860152611a7082826119d2565b91505060408301518482036040860152611a8a82826119d2565b91505060608301518482036060860152611aa482826119d2565b9150506080830151611ab96080860182611a0b565b5060a0830151611acc60a0860182611a0b565b5060c0830151611adf60c0860182611a1a565b5060e0830151611af260e0860182611a29565b50610100830151848203610100860152611b0c82826119d2565b9150508091505092915050565b6000604082019050611b2e6000830185611281565b8181036020830152611b408184611a38565b9050939250505056fea2646970667358221220e633c049e20aa5eef61e6c5c3eb133e2bade635709a77022c695547ae336dbf564736f6c63430008120033",
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
