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
	Title       string
	Version     string
	ReleaseDate string
	Developer   string
	Price       *big.Int
	Uploader    common.Address
	RootHash    [32]byte
	IpfsAddress string
	Purchased   []common.Address
}

// LibraryMetaData contains all meta data concerning the Library contract.
var LibraryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"releaseDate\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"developer\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"uploader\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"ipfsAddress\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"purchased\",\"type\":\"address[]\"}],\"indexed\":false,\"internalType\":\"structLibrary.GameEntry\",\"name\":\"game\",\"type\":\"tuple\"}],\"name\":\"NewGame\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"gameHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"games\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"releaseDate\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"developer\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"uploader\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"ipfsAddress\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"libSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_game\",\"type\":\"bytes32\"}],\"name\":\"purchaseGame\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_game\",\"type\":\"bytes32\"}],\"name\":\"purchasedSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"releaseDate\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"developer\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"uploader\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"ipfsAddress\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"purchased\",\"type\":\"address[]\"}],\"internalType\":\"structLibrary.GameEntry\",\"name\":\"_game\",\"type\":\"tuple\"}],\"name\":\"uploadGame\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061186b806100206000396000f3fe6080604052600436106100555760003560e01c80633e093f791461005a57806350e0c46e14610076578063741df964146100a1578063dc164c82146100de578063de611f2e1461011b578063f579f88214610144575b600080fd5b610074600480360381019061006f9190610a63565b610188565b005b34801561008257600080fd5b5061008b610406565b6040516100989190610aa9565b60405180910390f35b3480156100ad57600080fd5b506100c860048036038101906100c39190610a63565b610413565b6040516100d59190610aa9565b60405180910390f35b3480156100ea57600080fd5b5061010560048036038101906101009190610af0565b610498565b6040516101129190610b2c565b60405180910390f35b34801561012757600080fd5b50610142600480360381019061013d9190610f83565b6104bc565b005b34801561015057600080fd5b5061016b60048036038101906101669190610a63565b610662565b60405161017f98979695949392919061105a565b60405180910390f35b600080600083815260200190815260200160002060000180546101aa9061112a565b9050116101ec576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101e3906111a7565b60405180910390fd5b600080600083815260200190815260200160002090508060040154341015610249576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161024090611213565b60405180910390fd5b6000805b82600801805490508110156102eb573373ffffffffffffffffffffffffffffffffffffffff1683600801828154811061028957610288611233565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036102d857600191506102eb565b80806102e390611291565b91505061024d565b50801561032d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161032490611325565b60405180910390fd5b8160050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc83600401549081150290604051600060405180830381858888f1935050505015801561039b573d6000803e3d6000fd5b5081600801339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050565b6000600180549050905090565b60008060008084815260200190815260200160002060000180546104369061112a565b905011610478576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161046f906111a7565b60405180910390fd5b600080838152602001908152602001600020600801805490509050919050565b600181815481106104a857600080fd5b906000526020600020016000915090505481565b338160a0019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050806000808360c001518152602001908152602001600020600082015181600001908161052091906114f1565b50602082015181600101908161053691906114f1565b50604082015181600201908161054c91906114f1565b50606082015181600301908161056291906114f1565b506080820151816004015560a08201518160050160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060c0820151816006015560e08201518160070190816105d391906114f1565b506101008201518160080190805190602001906105f1929190610972565b5090505060018160c0015190806001815401808255809150506001900390600052602060002001600090919091909150557f0135990fdd35acb0b238be5e37d9497950348518ffb6ac6eb549cc3e05d0c42a8160c00151826040516106579291906117e0565b60405180910390a150565b60006020528060005260406000206000915090508060000180546106859061112a565b80601f01602080910402602001604051908101604052809291908181526020018280546106b19061112a565b80156106fe5780601f106106d3576101008083540402835291602001916106fe565b820191906000526020600020905b8154815290600101906020018083116106e157829003601f168201915b5050505050908060010180546107139061112a565b80601f016020809104026020016040519081016040528092919081815260200182805461073f9061112a565b801561078c5780601f106107615761010080835404028352916020019161078c565b820191906000526020600020905b81548152906001019060200180831161076f57829003601f168201915b5050505050908060020180546107a19061112a565b80601f01602080910402602001604051908101604052809291908181526020018280546107cd9061112a565b801561081a5780601f106107ef5761010080835404028352916020019161081a565b820191906000526020600020905b8154815290600101906020018083116107fd57829003601f168201915b50505050509080600301805461082f9061112a565b80601f016020809104026020016040519081016040528092919081815260200182805461085b9061112a565b80156108a85780601f1061087d576101008083540402835291602001916108a8565b820191906000526020600020905b81548152906001019060200180831161088b57829003601f168201915b5050505050908060040154908060050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060060154908060070180546108ef9061112a565b80601f016020809104026020016040519081016040528092919081815260200182805461091b9061112a565b80156109685780601f1061093d57610100808354040283529160200191610968565b820191906000526020600020905b81548152906001019060200180831161094b57829003601f168201915b5050505050905088565b8280548282559060005260206000209081019282156109eb579160200282015b828111156109ea5782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555091602001919060010190610992565b5b5090506109f891906109fc565b5090565b5b80821115610a155760008160009055506001016109fd565b5090565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b610a4081610a2d565b8114610a4b57600080fd5b50565b600081359050610a5d81610a37565b92915050565b600060208284031215610a7957610a78610a23565b5b6000610a8784828501610a4e565b91505092915050565b6000819050919050565b610aa381610a90565b82525050565b6000602082019050610abe6000830184610a9a565b92915050565b610acd81610a90565b8114610ad857600080fd5b50565b600081359050610aea81610ac4565b92915050565b600060208284031215610b0657610b05610a23565b5b6000610b1484828501610adb565b91505092915050565b610b2681610a2d565b82525050565b6000602082019050610b416000830184610b1d565b92915050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610b9582610b4c565b810181811067ffffffffffffffff82111715610bb457610bb3610b5d565b5b80604052505050565b6000610bc7610a19565b9050610bd38282610b8c565b919050565b600080fd5b600080fd5b600080fd5b600067ffffffffffffffff821115610c0257610c01610b5d565b5b610c0b82610b4c565b9050602081019050919050565b82818337600083830152505050565b6000610c3a610c3584610be7565b610bbd565b905082815260208101848484011115610c5657610c55610be2565b5b610c61848285610c18565b509392505050565b600082601f830112610c7e57610c7d610bdd565b5b8135610c8e848260208601610c27565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610cc282610c97565b9050919050565b610cd281610cb7565b8114610cdd57600080fd5b50565b600081359050610cef81610cc9565b92915050565b600067ffffffffffffffff821115610d1057610d0f610b5d565b5b602082029050602081019050919050565b600080fd5b6000610d3182610c97565b9050919050565b610d4181610d26565b8114610d4c57600080fd5b50565b600081359050610d5e81610d38565b92915050565b6000610d77610d7284610cf5565b610bbd565b90508083825260208201905060208402830185811115610d9a57610d99610d21565b5b835b81811015610dc35780610daf8882610d4f565b845260208401935050602081019050610d9c565b5050509392505050565b600082601f830112610de257610de1610bdd565b5b8135610df2848260208601610d64565b91505092915050565b60006101208284031215610e1257610e11610b47565b5b610e1d610120610bbd565b9050600082013567ffffffffffffffff811115610e3d57610e3c610bd8565b5b610e4984828501610c69565b600083015250602082013567ffffffffffffffff811115610e6d57610e6c610bd8565b5b610e7984828501610c69565b602083015250604082013567ffffffffffffffff811115610e9d57610e9c610bd8565b5b610ea984828501610c69565b604083015250606082013567ffffffffffffffff811115610ecd57610ecc610bd8565b5b610ed984828501610c69565b6060830152506080610eed84828501610adb565b60808301525060a0610f0184828501610ce0565b60a08301525060c0610f1584828501610a4e565b60c08301525060e082013567ffffffffffffffff811115610f3957610f38610bd8565b5b610f4584828501610c69565b60e08301525061010082013567ffffffffffffffff811115610f6a57610f69610bd8565b5b610f7684828501610dcd565b6101008301525092915050565b600060208284031215610f9957610f98610a23565b5b600082013567ffffffffffffffff811115610fb757610fb6610a28565b5b610fc384828501610dfb565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015611006578082015181840152602081019050610feb565b60008484015250505050565b600061101d82610fcc565b6110278185610fd7565b9350611037818560208601610fe8565b61104081610b4c565b840191505092915050565b61105481610cb7565b82525050565b6000610100820190508181036000830152611075818b611012565b90508181036020830152611089818a611012565b9050818103604083015261109d8189611012565b905081810360608301526110b18188611012565b90506110c06080830187610a9a565b6110cd60a083018661104b565b6110da60c0830185610b1d565b81810360e08301526110ec8184611012565b90509998505050505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061114257607f821691505b602082108103611155576111546110fb565b5b50919050565b7f67616d65206e6f7420666f756e64000000000000000000000000000000000000600082015250565b6000611191600e83610fd7565b915061119c8261115b565b602082019050919050565b600060208201905081810360008301526111c081611184565b9050919050565b7f757365722063616e6e6f74206166666f72642067616d65000000000000000000600082015250565b60006111fd601783610fd7565b9150611208826111c7565b602082019050919050565b6000602082019050818103600083015261122c816111f0565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061129c82610a90565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036112ce576112cd611262565b5b600182019050919050565b7f7573657220616c7265616479206f776e732067616d6500000000000000000000600082015250565b600061130f601683610fd7565b915061131a826112d9565b602082019050919050565b6000602082019050818103600083015261133e81611302565b9050919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026113a77fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261136a565b6113b1868361136a565b95508019841693508086168417925050509392505050565b6000819050919050565b60006113ee6113e96113e484610a90565b6113c9565b610a90565b9050919050565b6000819050919050565b611408836113d3565b61141c611414826113f5565b848454611377565b825550505050565b600090565b611431611424565b61143c8184846113ff565b505050565b5b8181101561146057611455600082611429565b600181019050611442565b5050565b601f8211156114a55761147681611345565b61147f8461135a565b8101602085101561148e578190505b6114a261149a8561135a565b830182611441565b50505b505050565b600082821c905092915050565b60006114c8600019846008026114aa565b1980831691505092915050565b60006114e183836114b7565b9150826002028217905092915050565b6114fa82610fcc565b67ffffffffffffffff81111561151357611512610b5d565b5b61151d825461112a565b611528828285611464565b600060209050601f83116001811461155b5760008415611549578287015190505b61155385826114d5565b8655506115bb565b601f19841661156986611345565b60005b828110156115915784890151825560018201915060208501945060208101905061156c565b868310156115ae57848901516115aa601f8916826114b7565b8355505b6001600288020188555050505b505050505050565b600082825260208201905092915050565b60006115df82610fcc565b6115e981856115c3565b93506115f9818560208601610fe8565b61160281610b4c565b840191505092915050565b61161681610a90565b82525050565b61162581610cb7565b82525050565b61163481610a2d565b82525050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61166f81610d26565b82525050565b60006116818383611666565b60208301905092915050565b6000602082019050919050565b60006116a58261163a565b6116af8185611645565b93506116ba83611656565b8060005b838110156116eb5781516116d28882611675565b97506116dd8361168d565b9250506001810190506116be565b5085935050505092915050565b600061012083016000830151848203600086015261171682826115d4565b9150506020830151848203602086015261173082826115d4565b9150506040830151848203604086015261174a82826115d4565b9150506060830151848203606086015261176482826115d4565b9150506080830151611779608086018261160d565b5060a083015161178c60a086018261161c565b5060c083015161179f60c086018261162b565b5060e083015184820360e08601526117b782826115d4565b9150506101008301518482036101008601526117d3828261169a565b9150508091505092915050565b60006040820190506117f56000830185610b1d565b818103602083015261180781846116f8565b9050939250505056fea2646970667358221220fd24073a43b336255888efe2082ad5c807fe4e60faba5a5f121b22bff0d4290064736f6c637827302e382e31392d646576656c6f702e323032332e322e382b636f6d6d69742e36363562663239610058",
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
// Solidity: function games(bytes32 ) view returns(string title, string version, string releaseDate, string developer, uint256 price, address uploader, bytes32 rootHash, string ipfsAddress)
func (_Library *LibraryCaller) Games(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Title       string
	Version     string
	ReleaseDate string
	Developer   string
	Price       *big.Int
	Uploader    common.Address
	RootHash    [32]byte
	IpfsAddress string
}, error) {
	var out []interface{}
	err := _Library.contract.Call(opts, &out, "games", arg0)

	outstruct := new(struct {
		Title       string
		Version     string
		ReleaseDate string
		Developer   string
		Price       *big.Int
		Uploader    common.Address
		RootHash    [32]byte
		IpfsAddress string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Title = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.ReleaseDate = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Developer = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Price = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Uploader = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.RootHash = *abi.ConvertType(out[6], new([32]byte)).(*[32]byte)
	outstruct.IpfsAddress = *abi.ConvertType(out[7], new(string)).(*string)

	return *outstruct, err

}

// Games is a free data retrieval call binding the contract method 0xf579f882.
//
// Solidity: function games(bytes32 ) view returns(string title, string version, string releaseDate, string developer, uint256 price, address uploader, bytes32 rootHash, string ipfsAddress)
func (_Library *LibrarySession) Games(arg0 [32]byte) (struct {
	Title       string
	Version     string
	ReleaseDate string
	Developer   string
	Price       *big.Int
	Uploader    common.Address
	RootHash    [32]byte
	IpfsAddress string
}, error) {
	return _Library.Contract.Games(&_Library.CallOpts, arg0)
}

// Games is a free data retrieval call binding the contract method 0xf579f882.
//
// Solidity: function games(bytes32 ) view returns(string title, string version, string releaseDate, string developer, uint256 price, address uploader, bytes32 rootHash, string ipfsAddress)
func (_Library *LibraryCallerSession) Games(arg0 [32]byte) (struct {
	Title       string
	Version     string
	ReleaseDate string
	Developer   string
	Price       *big.Int
	Uploader    common.Address
	RootHash    [32]byte
	IpfsAddress string
}, error) {
	return _Library.Contract.Games(&_Library.CallOpts, arg0)
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

// PurchasedSize is a free data retrieval call binding the contract method 0x741df964.
//
// Solidity: function purchasedSize(bytes32 _game) view returns(uint256)
func (_Library *LibraryCaller) PurchasedSize(opts *bind.CallOpts, _game [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Library.contract.Call(opts, &out, "purchasedSize", _game)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PurchasedSize is a free data retrieval call binding the contract method 0x741df964.
//
// Solidity: function purchasedSize(bytes32 _game) view returns(uint256)
func (_Library *LibrarySession) PurchasedSize(_game [32]byte) (*big.Int, error) {
	return _Library.Contract.PurchasedSize(&_Library.CallOpts, _game)
}

// PurchasedSize is a free data retrieval call binding the contract method 0x741df964.
//
// Solidity: function purchasedSize(bytes32 _game) view returns(uint256)
func (_Library *LibraryCallerSession) PurchasedSize(_game [32]byte) (*big.Int, error) {
	return _Library.Contract.PurchasedSize(&_Library.CallOpts, _game)
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

// UploadGame is a paid mutator transaction binding the contract method 0xde611f2e.
//
// Solidity: function uploadGame((string,string,string,string,uint256,address,bytes32,string,address[]) _game) returns()
func (_Library *LibraryTransactor) UploadGame(opts *bind.TransactOpts, _game LibraryGameEntry) (*types.Transaction, error) {
	return _Library.contract.Transact(opts, "uploadGame", _game)
}

// UploadGame is a paid mutator transaction binding the contract method 0xde611f2e.
//
// Solidity: function uploadGame((string,string,string,string,uint256,address,bytes32,string,address[]) _game) returns()
func (_Library *LibrarySession) UploadGame(_game LibraryGameEntry) (*types.Transaction, error) {
	return _Library.Contract.UploadGame(&_Library.TransactOpts, _game)
}

// UploadGame is a paid mutator transaction binding the contract method 0xde611f2e.
//
// Solidity: function uploadGame((string,string,string,string,uint256,address,bytes32,string,address[]) _game) returns()
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

// FilterNewGame is a free log retrieval operation binding the contract event 0x0135990fdd35acb0b238be5e37d9497950348518ffb6ac6eb549cc3e05d0c42a.
//
// Solidity: event NewGame(bytes32 hash, (string,string,string,string,uint256,address,bytes32,string,address[]) game)
func (_Library *LibraryFilterer) FilterNewGame(opts *bind.FilterOpts) (*LibraryNewGameIterator, error) {

	logs, sub, err := _Library.contract.FilterLogs(opts, "NewGame")
	if err != nil {
		return nil, err
	}
	return &LibraryNewGameIterator{contract: _Library.contract, event: "NewGame", logs: logs, sub: sub}, nil
}

// WatchNewGame is a free log subscription operation binding the contract event 0x0135990fdd35acb0b238be5e37d9497950348518ffb6ac6eb549cc3e05d0c42a.
//
// Solidity: event NewGame(bytes32 hash, (string,string,string,string,uint256,address,bytes32,string,address[]) game)
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

// ParseNewGame is a log parse operation binding the contract event 0x0135990fdd35acb0b238be5e37d9497950348518ffb6ac6eb549cc3e05d0c42a.
//
// Solidity: event NewGame(bytes32 hash, (string,string,string,string,uint256,address,bytes32,string,address[]) game)
func (_Library *LibraryFilterer) ParseNewGame(log types.Log) (*LibraryNewGame, error) {
	event := new(LibraryNewGame)
	if err := _Library.contract.UnpackLog(event, "NewGame", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
