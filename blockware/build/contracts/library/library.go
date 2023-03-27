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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"releaseDate\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"developer\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"previousVersion\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"uploader\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"hashTreeIPFSAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"assetsIPFSAddress\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structLibrary.GameEntry\",\"name\":\"game\",\"type\":\"tuple\"}],\"name\":\"NewGame\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"Purchase\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"gameHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"games\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"releaseDate\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"developer\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"previousVersion\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"uploader\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"hashTreeIPFSAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"assetsIPFSAddress\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_game\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"hasPurchased\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"libSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_game\",\"type\":\"bytes32\"}],\"name\":\"purchaseGame\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"releaseDate\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"developer\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"previousVersion\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"uploader\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"hashTreeIPFSAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"assetsIPFSAddress\",\"type\":\"string\"}],\"internalType\":\"structLibrary.GameEntry\",\"name\":\"_game\",\"type\":\"tuple\"}],\"name\":\"uploadGame\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506121ce806100206000396000f3fe6080604052600436106100555760003560e01c80630d0e12341461005a5780632d139a1b146100835780633e093f79146100c057806350e0c46e146100dc578063dc164c8214610107578063f579f88214610144575b600080fd5b34801561006657600080fd5b50610081600480360381019061007c919061147e565b61018a565b005b34801561008f57600080fd5b506100aa60048036038101906100a591906114c7565b610970565b6040516100b79190611522565b60405180910390f35b6100da60048036038101906100d5919061153d565b6109dd565b005b3480156100e857600080fd5b506100f1610cdd565b6040516100fe9190611579565b60405180910390f35b34801561011357600080fd5b5061012e60048036038101906101299190611594565b610cea565b60405161013b91906115d0565b60405180910390f35b34801561015057600080fd5b5061016b6004803603810190610166919061153d565b610d0e565b6040516101819a99989796959493929190611679565b60405180910390f35b6000816080015150602060ff16116101d7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101ce9061178b565b60405180910390fd5b60008161010001515111610220576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102179061181d565b60405180910390fd5b60008161012001515111610269576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610260906118af565b60405180910390fd5b6000801b8160a00151146107c65760008060008360a001518152602001908152602001600020600001805461029d906118fe565b9050116102df576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102d6906119a1565b60405180910390fd5b60008060008360a00151815260200190815260200160002060405180610140016040529081600082018054610313906118fe565b80601f016020809104026020016040519081016040528092919081815260200182805461033f906118fe565b801561038c5780601f106103615761010080835404028352916020019161038c565b820191906000526020600020905b81548152906001019060200180831161036f57829003601f168201915b505050505081526020016001820180546103a5906118fe565b80601f01602080910402602001604051908101604052809291908181526020018280546103d1906118fe565b801561041e5780601f106103f35761010080835404028352916020019161041e565b820191906000526020600020905b81548152906001019060200180831161040157829003601f168201915b50505050508152602001600282018054610437906118fe565b80601f0160208091040260200160405190810160405280929190818152602001828054610463906118fe565b80156104b05780601f10610485576101008083540402835291602001916104b0565b820191906000526020600020905b81548152906001019060200180831161049357829003601f168201915b505050505081526020016003820180546104c9906118fe565b80601f01602080910402602001604051908101604052809291908181526020018280546104f5906118fe565b80156105425780601f1061051757610100808354040283529160200191610542565b820191906000526020600020905b81548152906001019060200180831161052557829003601f168201915b505050505081526020016004820154815260200160058201548152602001600682015481526020016007820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016008820180546105cf906118fe565b80601f01602080910402602001604051908101604052809291908181526020018280546105fb906118fe565b80156106485780601f1061061d57610100808354040283529160200191610648565b820191906000526020600020905b81548152906001019060200180831161062b57829003601f168201915b50505050508152602001600982018054610661906118fe565b80601f016020809104026020016040519081016040528092919081815260200182805461068d906118fe565b80156106da5780601f106106af576101008083540402835291602001916106da565b820191906000526020600020905b8154815290600101906020018083116106bd57829003601f168201915b50505050508152505090503373ffffffffffffffffffffffffffffffffffffffff168160e0015173ffffffffffffffffffffffffffffffffffffffff1614610757576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161074e90611a33565b60405180910390fd5b60018060008460800151815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908360ff160217905550505b338160e0019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508060008083608001518152602001908152602001600020600082015181600001908161082a9190611bff565b5060208201518160010190816108409190611bff565b5060408201518160020190816108569190611bff565b50606082015181600301908161086c9190611bff565b506080820151816004015560a0820151816005015560c0820151816006015560e08201518160070160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506101008201518160080190816108e89190611bff565b506101208201518160090190816108ff9190611bff565b509050506002816080015190806001815401808255809150506001900390600052602060002001600090919091909150557f0184d8edb4833799540b2e0c1ae01aba9d7355f6267e124ee0ed0f2ca94f084c816080015182604051610965929190611e45565b60405180910390a150565b6000600180600085815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1660ff1614905092915050565b600080600083815260200190815260200160002060000180546109ff906118fe565b905011610a41576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a3890611ec1565b60405180910390fd5b60006001600083815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1660ff1614610ae4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610adb90611f2d565b60405180910390fd5b346000808381526020019081526020016000206006015414610b3b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b3290611fbf565b60405180910390fd5b60008060008084815260200190815260200160002060070160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1634600080868152602001908152602001600020600001604051610bad9190612082565b60006040518083038185875af1925050503d8060008114610bea576040519150601f19603f3d011682016040523d82523d6000602084013e610bef565b606091505b509150915081610c34576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c2b906120e5565b60405180910390fd5b600180600085815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908360ff1602179055507f4cd43aaf0718852d4b731ab37f54b2e990206041b88e7fffc13205577f96550a833383604051610cd09392919061215a565b60405180910390a1505050565b6000600280549050905090565b60028181548110610cfa57600080fd5b906000526020600020016000915090505481565b6000602052806000526040600020600091509050806000018054610d31906118fe565b80601f0160208091040260200160405190810160405280929190818152602001828054610d5d906118fe565b8015610daa5780601f10610d7f57610100808354040283529160200191610daa565b820191906000526020600020905b815481529060010190602001808311610d8d57829003601f168201915b505050505090806001018054610dbf906118fe565b80601f0160208091040260200160405190810160405280929190818152602001828054610deb906118fe565b8015610e385780601f10610e0d57610100808354040283529160200191610e38565b820191906000526020600020905b815481529060010190602001808311610e1b57829003601f168201915b505050505090806002018054610e4d906118fe565b80601f0160208091040260200160405190810160405280929190818152602001828054610e79906118fe565b8015610ec65780601f10610e9b57610100808354040283529160200191610ec6565b820191906000526020600020905b815481529060010190602001808311610ea957829003601f168201915b505050505090806003018054610edb906118fe565b80601f0160208091040260200160405190810160405280929190818152602001828054610f07906118fe565b8015610f545780601f10610f2957610100808354040283529160200191610f54565b820191906000526020600020905b815481529060010190602001808311610f3757829003601f168201915b5050505050908060040154908060050154908060060154908060070160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806008018054610fa1906118fe565b80601f0160208091040260200160405190810160405280929190818152602001828054610fcd906118fe565b801561101a5780601f10610fef5761010080835404028352916020019161101a565b820191906000526020600020905b815481529060010190602001808311610ffd57829003601f168201915b50505050509080600901805461102f906118fe565b80601f016020809104026020016040519081016040528092919081815260200182805461105b906118fe565b80156110a85780601f1061107d576101008083540402835291602001916110a8565b820191906000526020600020905b81548152906001019060200180831161108b57829003601f168201915b505050505090508a565b6000604051905090565b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b611114826110cb565b810181811067ffffffffffffffff82111715611133576111326110dc565b5b80604052505050565b60006111466110b2565b9050611152828261110b565b919050565b600080fd5b600080fd5b600080fd5b600067ffffffffffffffff821115611181576111806110dc565b5b61118a826110cb565b9050602081019050919050565b82818337600083830152505050565b60006111b96111b484611166565b61113c565b9050828152602081018484840111156111d5576111d4611161565b5b6111e0848285611197565b509392505050565b600082601f8301126111fd576111fc61115c565b5b813561120d8482602086016111a6565b91505092915050565b6000819050919050565b61122981611216565b811461123457600080fd5b50565b60008135905061124681611220565b92915050565b6000819050919050565b61125f8161124c565b811461126a57600080fd5b50565b60008135905061127c81611256565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006112ad82611282565b9050919050565b6112bd816112a2565b81146112c857600080fd5b50565b6000813590506112da816112b4565b92915050565b600061014082840312156112f7576112f66110c6565b5b61130261014061113c565b9050600082013567ffffffffffffffff81111561132257611321611157565b5b61132e848285016111e8565b600083015250602082013567ffffffffffffffff81111561135257611351611157565b5b61135e848285016111e8565b602083015250604082013567ffffffffffffffff81111561138257611381611157565b5b61138e848285016111e8565b604083015250606082013567ffffffffffffffff8111156113b2576113b1611157565b5b6113be848285016111e8565b60608301525060806113d284828501611237565b60808301525060a06113e684828501611237565b60a08301525060c06113fa8482850161126d565b60c08301525060e061140e848285016112cb565b60e08301525061010082013567ffffffffffffffff81111561143357611432611157565b5b61143f848285016111e8565b6101008301525061012082013567ffffffffffffffff81111561146557611464611157565b5b611471848285016111e8565b6101208301525092915050565b600060208284031215611494576114936110bc565b5b600082013567ffffffffffffffff8111156114b2576114b16110c1565b5b6114be848285016112e0565b91505092915050565b600080604083850312156114de576114dd6110bc565b5b60006114ec85828601611237565b92505060206114fd858286016112cb565b9150509250929050565b60008115159050919050565b61151c81611507565b82525050565b60006020820190506115376000830184611513565b92915050565b600060208284031215611553576115526110bc565b5b600061156184828501611237565b91505092915050565b6115738161124c565b82525050565b600060208201905061158e600083018461156a565b92915050565b6000602082840312156115aa576115a96110bc565b5b60006115b88482850161126d565b91505092915050565b6115ca81611216565b82525050565b60006020820190506115e560008301846115c1565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561162557808201518184015260208101905061160a565b60008484015250505050565b600061163c826115eb565b61164681856115f6565b9350611656818560208601611607565b61165f816110cb565b840191505092915050565b611673816112a2565b82525050565b6000610140820190508181036000830152611694818d611631565b905081810360208301526116a8818c611631565b905081810360408301526116bc818b611631565b905081810360608301526116d0818a611631565b90506116df60808301896115c1565b6116ec60a08301886115c1565b6116f960c083018761156a565b61170660e083018661166a565b8181036101008301526117198185611631565b905081810361012083015261172e8184611631565b90509b9a5050505050505050505050565b7f6e6f20726f6f74206861736820676976656e0000000000000000000000000000600082015250565b60006117756012836115f6565b91506117808261173f565b602082019050919050565b600060208201905081810360008301526117a481611768565b9050919050565b7f6e6f2049504653206164647265737320676976656e20666f722068617368207460008201527f7265656500000000000000000000000000000000000000000000000000000000602082015250565b60006118076024836115f6565b9150611812826117ab565b604082019050919050565b60006020820190508181036000830152611836816117fa565b9050919050565b7f6e6f2049504653206164647265737320676976656e20666f722074686520617360008201527f7365747320000000000000000000000000000000000000000000000000000000602082015250565b60006118996025836115f6565b91506118a48261183d565b604082019050919050565b600060208201905081810360008301526118c88161188c565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061191657607f821691505b602082108103611929576119286118cf565b5b50919050565b7f70726576696f75732076657273696f6e206f662067616d65206e6f7420666f7560008201527f6e64000000000000000000000000000000000000000000000000000000000000602082015250565b600061198b6022836115f6565b91506119968261192f565b604082019050919050565b600060208201905081810360008301526119ba8161197e565b9050919050565b7f6f6e6c7920746865206f726967696e616c2075706c6f616465722063616e207560008201527f70646174652074686569722067616d6500000000000000000000000000000000602082015250565b6000611a1d6030836115f6565b9150611a28826119c1565b604082019050919050565b60006020820190508181036000830152611a4c81611a10565b9050919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302611ab57fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611a78565b611abf8683611a78565b95508019841693508086168417925050509392505050565b6000819050919050565b6000611afc611af7611af28461124c565b611ad7565b61124c565b9050919050565b6000819050919050565b611b1683611ae1565b611b2a611b2282611b03565b848454611a85565b825550505050565b600090565b611b3f611b32565b611b4a818484611b0d565b505050565b5b81811015611b6e57611b63600082611b37565b600181019050611b50565b5050565b601f821115611bb357611b8481611a53565b611b8d84611a68565b81016020851015611b9c578190505b611bb0611ba885611a68565b830182611b4f565b50505b505050565b600082821c905092915050565b6000611bd660001984600802611bb8565b1980831691505092915050565b6000611bef8383611bc5565b9150826002028217905092915050565b611c08826115eb565b67ffffffffffffffff811115611c2157611c206110dc565b5b611c2b82546118fe565b611c36828285611b72565b600060209050601f831160018114611c695760008415611c57578287015190505b611c618582611be3565b865550611cc9565b601f198416611c7786611a53565b60005b82811015611c9f57848901518255600182019150602085019450602081019050611c7a565b86831015611cbc5784890151611cb8601f891682611bc5565b8355505b6001600288020188555050505b505050505050565b600082825260208201905092915050565b6000611ced826115eb565b611cf78185611cd1565b9350611d07818560208601611607565b611d10816110cb565b840191505092915050565b611d2481611216565b82525050565b611d338161124c565b82525050565b611d42816112a2565b82525050565b6000610140830160008301518482036000860152611d668282611ce2565b91505060208301518482036020860152611d808282611ce2565b91505060408301518482036040860152611d9a8282611ce2565b91505060608301518482036060860152611db48282611ce2565b9150506080830151611dc96080860182611d1b565b5060a0830151611ddc60a0860182611d1b565b5060c0830151611def60c0860182611d2a565b5060e0830151611e0260e0860182611d39565b50610100830151848203610100860152611e1c8282611ce2565b915050610120830151848203610120860152611e388282611ce2565b9150508091505092915050565b6000604082019050611e5a60008301856115c1565b8181036020830152611e6c8184611d48565b90509392505050565b7f67616d65206e6f7420666f756e64000000000000000000000000000000000000600082015250565b6000611eab600e836115f6565b9150611eb682611e75565b602082019050919050565b60006020820190508181036000830152611eda81611e9e565b9050919050565b7f7573657220616c7265616479206f776e732067616d6500000000000000000000600082015250565b6000611f176016836115f6565b9150611f2282611ee1565b602082019050919050565b60006020820190508181036000830152611f4681611f0a565b9050919050565b7f756e6578706563746564207072696365203d3e2076616c75652073686f756c6460008201527f20657175616c207468652067616d652773207072696363650000000000000000602082015250565b6000611fa96038836115f6565b9150611fb482611f4d565b604082019050919050565b60006020820190508181036000830152611fd881611f9c565b9050919050565b600081905092915050565b60008190508160005260206000209050919050565b6000815461200c816118fe565b6120168186611fdf565b94506001821660008114612031576001811461204657612079565b60ff1983168652811515820286019350612079565b61204f85611fea565b60005b8381101561207157815481890152600182019150602081019050612052565b838801955050505b50505092915050565b600061208e8284611fff565b915081905092915050565b7f4661696c656420746f207472616e736665722045746865720000000000000000600082015250565b60006120cf6018836115f6565b91506120da82612099565b602082019050919050565b600060208201905081810360008301526120fe816120c2565b9050919050565b600081519050919050565b600082825260208201905092915050565b600061212c82612105565b6121368185612110565b9350612146818560208601611607565b61214f816110cb565b840191505092915050565b600060608201905061216f60008301866115c1565b61217c602083018561166a565b818103604083015261218e8184612121565b905094935050505056fea2646970667358221220c1f3f77ed3d434fd7bdd61530d2a7efe52e6462e30fa575171a487358d1ad75464736f6c63430008120033",
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

// LibraryPurchaseIterator is returned from FilterPurchase and is used to iterate over the raw logs and unpacked data for Purchase events raised by the Library contract.
type LibraryPurchaseIterator struct {
	Event *LibraryPurchase // Event containing the contract specifics and raw log

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
func (it *LibraryPurchaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LibraryPurchase)
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
		it.Event = new(LibraryPurchase)
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
func (it *LibraryPurchaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LibraryPurchaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LibraryPurchase represents a Purchase event raised by the Library contract.
type LibraryPurchase struct {
	Arg0 [32]byte
	Arg1 common.Address
	Arg2 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPurchase is a free log retrieval operation binding the contract event 0x4cd43aaf0718852d4b731ab37f54b2e990206041b88e7fffc13205577f96550a.
//
// Solidity: event Purchase(bytes32 arg0, address arg1, bytes arg2)
func (_Library *LibraryFilterer) FilterPurchase(opts *bind.FilterOpts) (*LibraryPurchaseIterator, error) {

	logs, sub, err := _Library.contract.FilterLogs(opts, "Purchase")
	if err != nil {
		return nil, err
	}
	return &LibraryPurchaseIterator{contract: _Library.contract, event: "Purchase", logs: logs, sub: sub}, nil
}

// WatchPurchase is a free log subscription operation binding the contract event 0x4cd43aaf0718852d4b731ab37f54b2e990206041b88e7fffc13205577f96550a.
//
// Solidity: event Purchase(bytes32 arg0, address arg1, bytes arg2)
func (_Library *LibraryFilterer) WatchPurchase(opts *bind.WatchOpts, sink chan<- *LibraryPurchase) (event.Subscription, error) {

	logs, sub, err := _Library.contract.WatchLogs(opts, "Purchase")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LibraryPurchase)
				if err := _Library.contract.UnpackLog(event, "Purchase", log); err != nil {
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

// ParsePurchase is a log parse operation binding the contract event 0x4cd43aaf0718852d4b731ab37f54b2e990206041b88e7fffc13205577f96550a.
//
// Solidity: event Purchase(bytes32 arg0, address arg1, bytes arg2)
func (_Library *LibraryFilterer) ParsePurchase(log types.Log) (*LibraryPurchase, error) {
	event := new(LibraryPurchase)
	if err := _Library.contract.UnpackLog(event, "Purchase", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
