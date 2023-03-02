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
	Bin: "0x608060405234801561001057600080fd5b50611d1b806100206000396000f3fe6080604052600436106100555760003560e01c80630d0e12341461005a5780632d139a1b146100835780633e093f79146100c057806350e0c46e146100dc578063dc164c8214610107578063f579f88214610144575b600080fd5b34801561006657600080fd5b50610081600480360381019061007c91906112a8565b61018a565b005b34801561008f57600080fd5b506100aa60048036038101906100a591906112f1565b610927565b6040516100b7919061134c565b60405180910390f35b6100da60048036038101906100d59190611367565b610994565b005b3480156100e857600080fd5b506100f1610b07565b6040516100fe91906113a3565b60405180910390f35b34801561011357600080fd5b5061012e600480360381019061012991906113be565b610b14565b60405161013b91906113fa565b60405180910390f35b34801561015057600080fd5b5061016b60048036038101906101669190611367565b610b38565b6040516101819a999897969594939291906114a3565b60405180910390f35b6000816080015150602060ff16116101d7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101ce906115b5565b60405180910390fd5b60008161010001515111610220576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161021790611647565b60405180910390fd5b6000801b8160a001511461077d5760008060008360a001518152602001908152602001600020600001805461025490611696565b905011610296576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161028d90611739565b60405180910390fd5b60008060008360a001518152602001908152602001600020604051806101400160405290816000820180546102ca90611696565b80601f01602080910402602001604051908101604052809291908181526020018280546102f690611696565b80156103435780601f1061031857610100808354040283529160200191610343565b820191906000526020600020905b81548152906001019060200180831161032657829003601f168201915b5050505050815260200160018201805461035c90611696565b80601f016020809104026020016040519081016040528092919081815260200182805461038890611696565b80156103d55780601f106103aa576101008083540402835291602001916103d5565b820191906000526020600020905b8154815290600101906020018083116103b857829003601f168201915b505050505081526020016002820180546103ee90611696565b80601f016020809104026020016040519081016040528092919081815260200182805461041a90611696565b80156104675780601f1061043c57610100808354040283529160200191610467565b820191906000526020600020905b81548152906001019060200180831161044a57829003601f168201915b5050505050815260200160038201805461048090611696565b80601f01602080910402602001604051908101604052809291908181526020018280546104ac90611696565b80156104f95780601f106104ce576101008083540402835291602001916104f9565b820191906000526020600020905b8154815290600101906020018083116104dc57829003601f168201915b505050505081526020016004820154815260200160058201548152602001600682015481526020016007820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160088201805461058690611696565b80601f01602080910402602001604051908101604052809291908181526020018280546105b290611696565b80156105ff5780601f106105d4576101008083540402835291602001916105ff565b820191906000526020600020905b8154815290600101906020018083116105e257829003601f168201915b5050505050815260200160098201805461061890611696565b80601f016020809104026020016040519081016040528092919081815260200182805461064490611696565b80156106915780601f1061066657610100808354040283529160200191610691565b820191906000526020600020905b81548152906001019060200180831161067457829003601f168201915b50505050508152505090503373ffffffffffffffffffffffffffffffffffffffff168160e0015173ffffffffffffffffffffffffffffffffffffffff161461070e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610705906117cb565b60405180910390fd5b60018060008460800151815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908360ff160217905550505b338160e0019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050806000808360800151815260200190815260200160002060008201518160000190816107e19190611997565b5060208201518160010190816107f79190611997565b50604082015181600201908161080d9190611997565b5060608201518160030190816108239190611997565b506080820151816004015560a0820151816005015560c0820151816006015560e08201518160070160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061010082015181600801908161089f9190611997565b506101208201518160090190816108b69190611997565b509050506002816080015190806001815401808255809150506001900390600052602060002001600090919091909150557f0184d8edb4833799540b2e0c1ae01aba9d7355f6267e124ee0ed0f2ca94f084c81608001518260405161091c929190611bdd565b60405180910390a150565b6000600180600085815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1660ff1614905092915050565b600080600083815260200190815260200160002060000180546109b690611696565b9050116109f8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109ef90611c59565b60405180910390fd5b60006001600083815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1660ff1614610a9b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a9290611cc5565b60405180910390fd5b600180600083815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908360ff16021790555050565b6000600280549050905090565b60028181548110610b2457600080fd5b906000526020600020016000915090505481565b6000602052806000526040600020600091509050806000018054610b5b90611696565b80601f0160208091040260200160405190810160405280929190818152602001828054610b8790611696565b8015610bd45780601f10610ba957610100808354040283529160200191610bd4565b820191906000526020600020905b815481529060010190602001808311610bb757829003601f168201915b505050505090806001018054610be990611696565b80601f0160208091040260200160405190810160405280929190818152602001828054610c1590611696565b8015610c625780601f10610c3757610100808354040283529160200191610c62565b820191906000526020600020905b815481529060010190602001808311610c4557829003601f168201915b505050505090806002018054610c7790611696565b80601f0160208091040260200160405190810160405280929190818152602001828054610ca390611696565b8015610cf05780601f10610cc557610100808354040283529160200191610cf0565b820191906000526020600020905b815481529060010190602001808311610cd357829003601f168201915b505050505090806003018054610d0590611696565b80601f0160208091040260200160405190810160405280929190818152602001828054610d3190611696565b8015610d7e5780601f10610d5357610100808354040283529160200191610d7e565b820191906000526020600020905b815481529060010190602001808311610d6157829003601f168201915b5050505050908060040154908060050154908060060154908060070160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806008018054610dcb90611696565b80601f0160208091040260200160405190810160405280929190818152602001828054610df790611696565b8015610e445780601f10610e1957610100808354040283529160200191610e44565b820191906000526020600020905b815481529060010190602001808311610e2757829003601f168201915b505050505090806009018054610e5990611696565b80601f0160208091040260200160405190810160405280929190818152602001828054610e8590611696565b8015610ed25780601f10610ea757610100808354040283529160200191610ed2565b820191906000526020600020905b815481529060010190602001808311610eb557829003601f168201915b505050505090508a565b6000604051905090565b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610f3e82610ef5565b810181811067ffffffffffffffff82111715610f5d57610f5c610f06565b5b80604052505050565b6000610f70610edc565b9050610f7c8282610f35565b919050565b600080fd5b600080fd5b600080fd5b600067ffffffffffffffff821115610fab57610faa610f06565b5b610fb482610ef5565b9050602081019050919050565b82818337600083830152505050565b6000610fe3610fde84610f90565b610f66565b905082815260208101848484011115610fff57610ffe610f8b565b5b61100a848285610fc1565b509392505050565b600082601f83011261102757611026610f86565b5b8135611037848260208601610fd0565b91505092915050565b6000819050919050565b61105381611040565b811461105e57600080fd5b50565b6000813590506110708161104a565b92915050565b6000819050919050565b61108981611076565b811461109457600080fd5b50565b6000813590506110a681611080565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006110d7826110ac565b9050919050565b6110e7816110cc565b81146110f257600080fd5b50565b600081359050611104816110de565b92915050565b6000610140828403121561112157611120610ef0565b5b61112c610140610f66565b9050600082013567ffffffffffffffff81111561114c5761114b610f81565b5b61115884828501611012565b600083015250602082013567ffffffffffffffff81111561117c5761117b610f81565b5b61118884828501611012565b602083015250604082013567ffffffffffffffff8111156111ac576111ab610f81565b5b6111b884828501611012565b604083015250606082013567ffffffffffffffff8111156111dc576111db610f81565b5b6111e884828501611012565b60608301525060806111fc84828501611061565b60808301525060a061121084828501611061565b60a08301525060c061122484828501611097565b60c08301525060e0611238848285016110f5565b60e08301525061010082013567ffffffffffffffff81111561125d5761125c610f81565b5b61126984828501611012565b6101008301525061012082013567ffffffffffffffff81111561128f5761128e610f81565b5b61129b84828501611012565b6101208301525092915050565b6000602082840312156112be576112bd610ee6565b5b600082013567ffffffffffffffff8111156112dc576112db610eeb565b5b6112e88482850161110a565b91505092915050565b6000806040838503121561130857611307610ee6565b5b600061131685828601611061565b9250506020611327858286016110f5565b9150509250929050565b60008115159050919050565b61134681611331565b82525050565b6000602082019050611361600083018461133d565b92915050565b60006020828403121561137d5761137c610ee6565b5b600061138b84828501611061565b91505092915050565b61139d81611076565b82525050565b60006020820190506113b86000830184611394565b92915050565b6000602082840312156113d4576113d3610ee6565b5b60006113e284828501611097565b91505092915050565b6113f481611040565b82525050565b600060208201905061140f60008301846113eb565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561144f578082015181840152602081019050611434565b60008484015250505050565b600061146682611415565b6114708185611420565b9350611480818560208601611431565b61148981610ef5565b840191505092915050565b61149d816110cc565b82525050565b60006101408201905081810360008301526114be818d61145b565b905081810360208301526114d2818c61145b565b905081810360408301526114e6818b61145b565b905081810360608301526114fa818a61145b565b905061150960808301896113eb565b61151660a08301886113eb565b61152360c0830187611394565b61153060e0830186611494565b818103610100830152611543818561145b565b9050818103610120830152611558818461145b565b90509b9a5050505050505050505050565b7f6e6f20726f6f74206861736820676976656e0000000000000000000000000000600082015250565b600061159f601283611420565b91506115aa82611569565b602082019050919050565b600060208201905081810360008301526115ce81611592565b9050919050565b7f6e6f2049504653206164647265737320676976656e20666f722068617368207460008201527f7265656500000000000000000000000000000000000000000000000000000000602082015250565b6000611631602483611420565b915061163c826115d5565b604082019050919050565b6000602082019050818103600083015261166081611624565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806116ae57607f821691505b6020821081036116c1576116c0611667565b5b50919050565b7f70726576696f75732076657273696f6e206f662067616d65206e6f7420666f7560008201527f6e64000000000000000000000000000000000000000000000000000000000000602082015250565b6000611723602283611420565b915061172e826116c7565b604082019050919050565b6000602082019050818103600083015261175281611716565b9050919050565b7f6f6e6c7920746865206f726967696e616c2075706c6f616465722063616e207560008201527f70646174652074686569722067616d6500000000000000000000000000000000602082015250565b60006117b5603083611420565b91506117c082611759565b604082019050919050565b600060208201905081810360008301526117e4816117a8565b9050919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b60006008830261184d7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611810565b6118578683611810565b95508019841693508086168417925050509392505050565b6000819050919050565b600061189461188f61188a84611076565b61186f565b611076565b9050919050565b6000819050919050565b6118ae83611879565b6118c26118ba8261189b565b84845461181d565b825550505050565b600090565b6118d76118ca565b6118e28184846118a5565b505050565b5b81811015611906576118fb6000826118cf565b6001810190506118e8565b5050565b601f82111561194b5761191c816117eb565b61192584611800565b81016020851015611934578190505b61194861194085611800565b8301826118e7565b50505b505050565b600082821c905092915050565b600061196e60001984600802611950565b1980831691505092915050565b6000611987838361195d565b9150826002028217905092915050565b6119a082611415565b67ffffffffffffffff8111156119b9576119b8610f06565b5b6119c38254611696565b6119ce82828561190a565b600060209050601f831160018114611a0157600084156119ef578287015190505b6119f9858261197b565b865550611a61565b601f198416611a0f866117eb565b60005b82811015611a3757848901518255600182019150602085019450602081019050611a12565b86831015611a545784890151611a50601f89168261195d565b8355505b6001600288020188555050505b505050505050565b600082825260208201905092915050565b6000611a8582611415565b611a8f8185611a69565b9350611a9f818560208601611431565b611aa881610ef5565b840191505092915050565b611abc81611040565b82525050565b611acb81611076565b82525050565b611ada816110cc565b82525050565b6000610140830160008301518482036000860152611afe8282611a7a565b91505060208301518482036020860152611b188282611a7a565b91505060408301518482036040860152611b328282611a7a565b91505060608301518482036060860152611b4c8282611a7a565b9150506080830151611b616080860182611ab3565b5060a0830151611b7460a0860182611ab3565b5060c0830151611b8760c0860182611ac2565b5060e0830151611b9a60e0860182611ad1565b50610100830151848203610100860152611bb48282611a7a565b915050610120830151848203610120860152611bd08282611a7a565b9150508091505092915050565b6000604082019050611bf260008301856113eb565b8181036020830152611c048184611ae0565b90509392505050565b7f67616d65206e6f7420666f756e64000000000000000000000000000000000000600082015250565b6000611c43600e83611420565b9150611c4e82611c0d565b602082019050919050565b60006020820190508181036000830152611c7281611c36565b9050919050565b7f7573657220616c7265616479206f776e732067616d6500000000000000000000600082015250565b6000611caf601683611420565b9150611cba82611c79565b602082019050919050565b60006020820190508181036000830152611cde81611ca2565b905091905056fea2646970667358221220a2dc93969ddeac462f3f474ae85328846571c83da20c3ca45772c112a9b23d4364736f6c63430008120033",
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
