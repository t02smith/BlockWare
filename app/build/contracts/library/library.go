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
	RootHash    [32]byte
	IpfsAddress string
}

// LibraryMetaData contains all meta data concerning the Library contract.
var LibraryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"releaseDate\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"developer\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"ipfsAddress\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structLibrary.GameEntry\",\"name\":\"game\",\"type\":\"tuple\"}],\"name\":\"NewGame\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"gameHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"games\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"releaseDate\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"developer\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"ipfsAddress\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"libSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"releaseDate\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"developer\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"ipfsAddress\",\"type\":\"string\"}],\"internalType\":\"structLibrary.GameEntry\",\"name\":\"_game\",\"type\":\"tuple\"}],\"name\":\"uploadGame\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610e3c806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806350e0c46e14610051578063792876d41461006f578063dc164c821461008b578063f579f882146100bb575b600080fd5b6100596100f0565b604051610066919061051d565b60405180910390f35b610089600480360381019061008491906107fe565b6100fd565b005b6100a560048036038101906100a09190610873565b6101fc565b6040516100b291906108af565b60405180910390f35b6100d560048036038101906100d091906108ca565b610220565b6040516100e796959493929190610976565b60405180910390f35b6000600180549050905090565b806000808360800151815260200190815260200160002060008201518160000190816101299190610c06565b50602082015181600101908161013f9190610c06565b5060408201518160020190816101559190610c06565b50606082015181600301908161016b9190610c06565b506080820151816004015560a082015181600501908161018b9190610c06565b509050506001816080015190806001815401808255809150506001900390600052602060002001600090919091909150557f25a0ddd57e2f4e54cbe40efbe92dcb461730058cfaa6abc4b02904ae94c6107c8160800151826040516101f1929190610dd6565b60405180910390a150565b6001818154811061020c57600080fd5b906000526020600020016000915090505481565b600060205280600052604060002060009150905080600001805461024390610a29565b80601f016020809104026020016040519081016040528092919081815260200182805461026f90610a29565b80156102bc5780601f10610291576101008083540402835291602001916102bc565b820191906000526020600020905b81548152906001019060200180831161029f57829003601f168201915b5050505050908060010180546102d190610a29565b80601f01602080910402602001604051908101604052809291908181526020018280546102fd90610a29565b801561034a5780601f1061031f5761010080835404028352916020019161034a565b820191906000526020600020905b81548152906001019060200180831161032d57829003601f168201915b50505050509080600201805461035f90610a29565b80601f016020809104026020016040519081016040528092919081815260200182805461038b90610a29565b80156103d85780601f106103ad576101008083540402835291602001916103d8565b820191906000526020600020905b8154815290600101906020018083116103bb57829003601f168201915b5050505050908060030180546103ed90610a29565b80601f016020809104026020016040519081016040528092919081815260200182805461041990610a29565b80156104665780601f1061043b57610100808354040283529160200191610466565b820191906000526020600020905b81548152906001019060200180831161044957829003601f168201915b50505050509080600401549080600501805461048190610a29565b80601f01602080910402602001604051908101604052809291908181526020018280546104ad90610a29565b80156104fa5780601f106104cf576101008083540402835291602001916104fa565b820191906000526020600020905b8154815290600101906020018083116104dd57829003601f168201915b5050505050905086565b6000819050919050565b61051781610504565b82525050565b6000602082019050610532600083018461050e565b92915050565b6000604051905090565b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61059a82610551565b810181811067ffffffffffffffff821117156105b9576105b8610562565b5b80604052505050565b60006105cc610538565b90506105d88282610591565b919050565b600080fd5b600080fd5b600080fd5b600067ffffffffffffffff82111561060757610606610562565b5b61061082610551565b9050602081019050919050565b82818337600083830152505050565b600061063f61063a846105ec565b6105c2565b90508281526020810184848401111561065b5761065a6105e7565b5b61066684828561061d565b509392505050565b600082601f830112610683576106826105e2565b5b813561069384826020860161062c565b91505092915050565b6000819050919050565b6106af8161069c565b81146106ba57600080fd5b50565b6000813590506106cc816106a6565b92915050565b600060c082840312156106e8576106e761054c565b5b6106f260c06105c2565b9050600082013567ffffffffffffffff811115610712576107116105dd565b5b61071e8482850161066e565b600083015250602082013567ffffffffffffffff811115610742576107416105dd565b5b61074e8482850161066e565b602083015250604082013567ffffffffffffffff811115610772576107716105dd565b5b61077e8482850161066e565b604083015250606082013567ffffffffffffffff8111156107a2576107a16105dd565b5b6107ae8482850161066e565b60608301525060806107c2848285016106bd565b60808301525060a082013567ffffffffffffffff8111156107e6576107e56105dd565b5b6107f28482850161066e565b60a08301525092915050565b60006020828403121561081457610813610542565b5b600082013567ffffffffffffffff81111561083257610831610547565b5b61083e848285016106d2565b91505092915050565b61085081610504565b811461085b57600080fd5b50565b60008135905061086d81610847565b92915050565b60006020828403121561088957610888610542565b5b60006108978482850161085e565b91505092915050565b6108a98161069c565b82525050565b60006020820190506108c460008301846108a0565b92915050565b6000602082840312156108e0576108df610542565b5b60006108ee848285016106bd565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610931578082015181840152602081019050610916565b60008484015250505050565b6000610948826108f7565b6109528185610902565b9350610962818560208601610913565b61096b81610551565b840191505092915050565b600060c0820190508181036000830152610990818961093d565b905081810360208301526109a4818861093d565b905081810360408301526109b8818761093d565b905081810360608301526109cc818661093d565b90506109db60808301856108a0565b81810360a08301526109ed818461093d565b9050979650505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610a4157607f821691505b602082108103610a5457610a536109fa565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302610abc7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610a7f565b610ac68683610a7f565b95508019841693508086168417925050509392505050565b6000819050919050565b6000610b03610afe610af984610504565b610ade565b610504565b9050919050565b6000819050919050565b610b1d83610ae8565b610b31610b2982610b0a565b848454610a8c565b825550505050565b600090565b610b46610b39565b610b51818484610b14565b505050565b5b81811015610b7557610b6a600082610b3e565b600181019050610b57565b5050565b601f821115610bba57610b8b81610a5a565b610b9484610a6f565b81016020851015610ba3578190505b610bb7610baf85610a6f565b830182610b56565b50505b505050565b600082821c905092915050565b6000610bdd60001984600802610bbf565b1980831691505092915050565b6000610bf68383610bcc565b9150826002028217905092915050565b610c0f826108f7565b67ffffffffffffffff811115610c2857610c27610562565b5b610c328254610a29565b610c3d828285610b79565b600060209050601f831160018114610c705760008415610c5e578287015190505b610c688582610bea565b865550610cd0565b601f198416610c7e86610a5a565b60005b82811015610ca657848901518255600182019150602085019450602081019050610c81565b86831015610cc35784890151610cbf601f891682610bcc565b8355505b6001600288020188555050505b505050505050565b600082825260208201905092915050565b6000610cf4826108f7565b610cfe8185610cd8565b9350610d0e818560208601610913565b610d1781610551565b840191505092915050565b610d2b8161069c565b82525050565b600060c0830160008301518482036000860152610d4e8282610ce9565b91505060208301518482036020860152610d688282610ce9565b91505060408301518482036040860152610d828282610ce9565b91505060608301518482036060860152610d9c8282610ce9565b9150506080830151610db16080860182610d22565b5060a083015184820360a0860152610dc98282610ce9565b9150508091505092915050565b6000604082019050610deb60008301856108a0565b8181036020830152610dfd8184610d31565b9050939250505056fea264697066735822122074dce8048d56046f2fa9bbd30333724d21a17fd19b19bda948b4674ac798d5f764736f6c63430008120033",
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
// Solidity: function games(bytes32 ) view returns(string title, string version, string releaseDate, string developer, bytes32 rootHash, string ipfsAddress)
func (_Library *LibraryCaller) Games(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Title       string
	Version     string
	ReleaseDate string
	Developer   string
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
	outstruct.RootHash = *abi.ConvertType(out[4], new([32]byte)).(*[32]byte)
	outstruct.IpfsAddress = *abi.ConvertType(out[5], new(string)).(*string)

	return *outstruct, err

}

// Games is a free data retrieval call binding the contract method 0xf579f882.
//
// Solidity: function games(bytes32 ) view returns(string title, string version, string releaseDate, string developer, bytes32 rootHash, string ipfsAddress)
func (_Library *LibrarySession) Games(arg0 [32]byte) (struct {
	Title       string
	Version     string
	ReleaseDate string
	Developer   string
	RootHash    [32]byte
	IpfsAddress string
}, error) {
	return _Library.Contract.Games(&_Library.CallOpts, arg0)
}

// Games is a free data retrieval call binding the contract method 0xf579f882.
//
// Solidity: function games(bytes32 ) view returns(string title, string version, string releaseDate, string developer, bytes32 rootHash, string ipfsAddress)
func (_Library *LibraryCallerSession) Games(arg0 [32]byte) (struct {
	Title       string
	Version     string
	ReleaseDate string
	Developer   string
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

// UploadGame is a paid mutator transaction binding the contract method 0x792876d4.
//
// Solidity: function uploadGame((string,string,string,string,bytes32,string) _game) returns()
func (_Library *LibraryTransactor) UploadGame(opts *bind.TransactOpts, _game LibraryGameEntry) (*types.Transaction, error) {
	return _Library.contract.Transact(opts, "uploadGame", _game)
}

// UploadGame is a paid mutator transaction binding the contract method 0x792876d4.
//
// Solidity: function uploadGame((string,string,string,string,bytes32,string) _game) returns()
func (_Library *LibrarySession) UploadGame(_game LibraryGameEntry) (*types.Transaction, error) {
	return _Library.Contract.UploadGame(&_Library.TransactOpts, _game)
}

// UploadGame is a paid mutator transaction binding the contract method 0x792876d4.
//
// Solidity: function uploadGame((string,string,string,string,bytes32,string) _game) returns()
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

// FilterNewGame is a free log retrieval operation binding the contract event 0x25a0ddd57e2f4e54cbe40efbe92dcb461730058cfaa6abc4b02904ae94c6107c.
//
// Solidity: event NewGame(bytes32 hash, (string,string,string,string,bytes32,string) game)
func (_Library *LibraryFilterer) FilterNewGame(opts *bind.FilterOpts) (*LibraryNewGameIterator, error) {

	logs, sub, err := _Library.contract.FilterLogs(opts, "NewGame")
	if err != nil {
		return nil, err
	}
	return &LibraryNewGameIterator{contract: _Library.contract, event: "NewGame", logs: logs, sub: sub}, nil
}

// WatchNewGame is a free log subscription operation binding the contract event 0x25a0ddd57e2f4e54cbe40efbe92dcb461730058cfaa6abc4b02904ae94c6107c.
//
// Solidity: event NewGame(bytes32 hash, (string,string,string,string,bytes32,string) game)
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

// ParseNewGame is a log parse operation binding the contract event 0x25a0ddd57e2f4e54cbe40efbe92dcb461730058cfaa6abc4b02904ae94c6107c.
//
// Solidity: event NewGame(bytes32 hash, (string,string,string,string,bytes32,string) game)
func (_Library *LibraryFilterer) ParseNewGame(log types.Log) (*LibraryNewGame, error) {
	event := new(LibraryNewGame)
	if err := _Library.contract.UnpackLog(event, "NewGame", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
