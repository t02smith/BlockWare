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
	Purchased       []common.Address
	IpfsAddress     string
}

// LibraryMetaData contains all meta data concerning the Library contract.
var LibraryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"releaseDate\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"developer\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"previousVersion\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"uploader\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"purchased\",\"type\":\"address[]\"},{\"internalType\":\"string\",\"name\":\"ipfsAddress\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structLibrary.GameEntry\",\"name\":\"game\",\"type\":\"tuple\"}],\"name\":\"NewGame\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"gameHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"games\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"releaseDate\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"developer\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"previousVersion\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"uploader\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ipfsAddress\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"libSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_game\",\"type\":\"bytes32\"}],\"name\":\"purchaseGame\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_game\",\"type\":\"bytes32\"}],\"name\":\"purchasedSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"releaseDate\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"developer\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"previousVersion\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"uploader\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"purchased\",\"type\":\"address[]\"},{\"internalType\":\"string\",\"name\":\"ipfsAddress\",\"type\":\"string\"}],\"internalType\":\"structLibrary.GameEntry\",\"name\":\"_game\",\"type\":\"tuple\"}],\"name\":\"uploadGame\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061206f806100206000396000f3fe6080604052600436106100555760003560e01c80633e093f791461005a57806350e0c46e14610076578063741df964146100a157806390691d6a146100de578063dc164c8214610107578063f579f88214610144575b600080fd5b610074600480360381019061006f919061100a565b610189565b005b34801561008257600080fd5b5061008b610407565b6040516100989190611050565b60405180910390f35b3480156100ad57600080fd5b506100c860048036038101906100c3919061100a565b610414565b6040516100d59190611050565b60405180910390f35b3480156100ea57600080fd5b50610105600480360381019061010091906114e9565b610499565b005b34801561011357600080fd5b5061012e60048036038101906101299190611532565b610bdf565b60405161013b919061156e565b60405180910390f35b34801561015057600080fd5b5061016b6004803603810190610166919061100a565b610c03565b60405161018099989796959493929190611617565b60405180910390f35b600080600083815260200190815260200160002060000180546101ab906116f6565b9050116101ed576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101e490611773565b60405180910390fd5b60008060008381526020019081526020016000209050806006015434101561024a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610241906117df565b60405180910390fd5b6000805b82600801805490508110156102ec573373ffffffffffffffffffffffffffffffffffffffff1683600801828154811061028a576102896117ff565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036102d957600191506102ec565b80806102e49061185d565b91505061024e565b50801561032e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610325906118f1565b60405180910390fd5b8160070160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc83600601549081150290604051600060405180830381858888f1935050505015801561039c573d6000803e3d6000fd5b5081600801339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050565b6000600180549050905090565b6000806000808481526020019081526020016000206000018054610437906116f6565b905011610479576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161047090611773565b60405180910390fd5b600080838152602001908152602001600020600801805490509050919050565b6000816080015150602060ff16116104e6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104dd9061195d565b60405180910390fd5b6000816101200151511161052f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610526906119ef565b60405180910390fd5b60008160a0015150602060ff1614610a2e5760008060008360a0015181526020019081526020016000206000018054610567906116f6565b9050116105a9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105a090611a81565b60405180910390fd5b60008060008360a001518152602001908152602001600020604051806101400160405290816000820180546105dd906116f6565b80601f0160208091040260200160405190810160405280929190818152602001828054610609906116f6565b80156106565780601f1061062b57610100808354040283529160200191610656565b820191906000526020600020905b81548152906001019060200180831161063957829003601f168201915b5050505050815260200160018201805461066f906116f6565b80601f016020809104026020016040519081016040528092919081815260200182805461069b906116f6565b80156106e85780601f106106bd576101008083540402835291602001916106e8565b820191906000526020600020905b8154815290600101906020018083116106cb57829003601f168201915b50505050508152602001600282018054610701906116f6565b80601f016020809104026020016040519081016040528092919081815260200182805461072d906116f6565b801561077a5780601f1061074f5761010080835404028352916020019161077a565b820191906000526020600020905b81548152906001019060200180831161075d57829003601f168201915b50505050508152602001600382018054610793906116f6565b80601f01602080910402602001604051908101604052809291908181526020018280546107bf906116f6565b801561080c5780601f106107e15761010080835404028352916020019161080c565b820191906000526020600020905b8154815290600101906020018083116107ef57829003601f168201915b505050505081526020016004820154815260200160058201548152602001600682015481526020016007820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016008820180548060200260200160405190810160405280929190818152602001828054801561090e57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116108c4575b50505050508152602001600982018054610927906116f6565b80601f0160208091040260200160405190810160405280929190818152602001828054610953906116f6565b80156109a05780601f10610975576101008083540402835291602001916109a0565b820191906000526020600020905b81548152906001019060200180831161098357829003601f168201915b50505050508152505090503373ffffffffffffffffffffffffffffffffffffffff168160e0015173ffffffffffffffffffffffffffffffffffffffff1614610a1d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a1490611b13565b60405180910390fd5b806101000151826101000181905250505b338160e0019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505080600080836080015181526020019081526020016000206000820151816000019081610a929190611cdf565b506020820151816001019081610aa89190611cdf565b506040820151816002019081610abe9190611cdf565b506060820151816003019081610ad49190611cdf565b506080820151816004015560a0820151816005015560c0820151816006015560e08201518160070160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610100820151816008019080519060200190610b57929190610f19565b50610120820151816009019081610b6e9190611cdf565b509050506001816080015190806001815401808255809150506001900390600052602060002001600090919091909150557f1e55381ba3dfad1800fac2963559018a3bf1b9b90b187ca535709c9f0905cc05816080015182604051610bd4929190611fe3565b60405180910390a150565b60018181548110610bef57600080fd5b906000526020600020016000915090505481565b6000602052806000526040600020600091509050806000018054610c26906116f6565b80601f0160208091040260200160405190810160405280929190818152602001828054610c52906116f6565b8015610c9f5780601f10610c7457610100808354040283529160200191610c9f565b820191906000526020600020905b815481529060010190602001808311610c8257829003601f168201915b505050505090806001018054610cb4906116f6565b80601f0160208091040260200160405190810160405280929190818152602001828054610ce0906116f6565b8015610d2d5780601f10610d0257610100808354040283529160200191610d2d565b820191906000526020600020905b815481529060010190602001808311610d1057829003601f168201915b505050505090806002018054610d42906116f6565b80601f0160208091040260200160405190810160405280929190818152602001828054610d6e906116f6565b8015610dbb5780601f10610d9057610100808354040283529160200191610dbb565b820191906000526020600020905b815481529060010190602001808311610d9e57829003601f168201915b505050505090806003018054610dd0906116f6565b80601f0160208091040260200160405190810160405280929190818152602001828054610dfc906116f6565b8015610e495780601f10610e1e57610100808354040283529160200191610e49565b820191906000526020600020905b815481529060010190602001808311610e2c57829003601f168201915b5050505050908060040154908060050154908060060154908060070160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806009018054610e96906116f6565b80601f0160208091040260200160405190810160405280929190818152602001828054610ec2906116f6565b8015610f0f5780601f10610ee457610100808354040283529160200191610f0f565b820191906000526020600020905b815481529060010190602001808311610ef257829003601f168201915b5050505050905089565b828054828255906000526020600020908101928215610f92579160200282015b82811115610f915782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555091602001919060010190610f39565b5b509050610f9f9190610fa3565b5090565b5b80821115610fbc576000816000905550600101610fa4565b5090565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b610fe781610fd4565b8114610ff257600080fd5b50565b60008135905061100481610fde565b92915050565b6000602082840312156110205761101f610fca565b5b600061102e84828501610ff5565b91505092915050565b6000819050919050565b61104a81611037565b82525050565b60006020820190506110656000830184611041565b92915050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6110b982611070565b810181811067ffffffffffffffff821117156110d8576110d7611081565b5b80604052505050565b60006110eb610fc0565b90506110f782826110b0565b919050565b600080fd5b600080fd5b600080fd5b600067ffffffffffffffff82111561112657611125611081565b5b61112f82611070565b9050602081019050919050565b82818337600083830152505050565b600061115e6111598461110b565b6110e1565b90508281526020810184848401111561117a57611179611106565b5b61118584828561113c565b509392505050565b600082601f8301126111a2576111a1611101565b5b81356111b284826020860161114b565b91505092915050565b6111c481611037565b81146111cf57600080fd5b50565b6000813590506111e1816111bb565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000611212826111e7565b9050919050565b61122281611207565b811461122d57600080fd5b50565b60008135905061123f81611219565b92915050565b600067ffffffffffffffff8211156112605761125f611081565b5b602082029050602081019050919050565b600080fd5b6000611281826111e7565b9050919050565b61129181611276565b811461129c57600080fd5b50565b6000813590506112ae81611288565b92915050565b60006112c76112c284611245565b6110e1565b905080838252602082019050602084028301858111156112ea576112e9611271565b5b835b8181101561131357806112ff888261129f565b8452602084019350506020810190506112ec565b5050509392505050565b600082601f83011261133257611331611101565b5b81356113428482602086016112b4565b91505092915050565b600061014082840312156113625761136161106b565b5b61136d6101406110e1565b9050600082013567ffffffffffffffff81111561138d5761138c6110fc565b5b6113998482850161118d565b600083015250602082013567ffffffffffffffff8111156113bd576113bc6110fc565b5b6113c98482850161118d565b602083015250604082013567ffffffffffffffff8111156113ed576113ec6110fc565b5b6113f98482850161118d565b604083015250606082013567ffffffffffffffff81111561141d5761141c6110fc565b5b6114298482850161118d565b606083015250608061143d84828501610ff5565b60808301525060a061145184828501610ff5565b60a08301525060c0611465848285016111d2565b60c08301525060e061147984828501611230565b60e08301525061010082013567ffffffffffffffff81111561149e5761149d6110fc565b5b6114aa8482850161131d565b6101008301525061012082013567ffffffffffffffff8111156114d0576114cf6110fc565b5b6114dc8482850161118d565b6101208301525092915050565b6000602082840312156114ff576114fe610fca565b5b600082013567ffffffffffffffff81111561151d5761151c610fcf565b5b6115298482850161134b565b91505092915050565b60006020828403121561154857611547610fca565b5b6000611556848285016111d2565b91505092915050565b61156881610fd4565b82525050565b6000602082019050611583600083018461155f565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156115c35780820151818401526020810190506115a8565b60008484015250505050565b60006115da82611589565b6115e48185611594565b93506115f48185602086016115a5565b6115fd81611070565b840191505092915050565b61161181611207565b82525050565b6000610120820190508181036000830152611632818c6115cf565b90508181036020830152611646818b6115cf565b9050818103604083015261165a818a6115cf565b9050818103606083015261166e81896115cf565b905061167d608083018861155f565b61168a60a083018761155f565b61169760c0830186611041565b6116a460e0830185611608565b8181036101008301526116b781846115cf565b90509a9950505050505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061170e57607f821691505b602082108103611721576117206116c7565b5b50919050565b7f67616d65206e6f7420666f756e64000000000000000000000000000000000000600082015250565b600061175d600e83611594565b915061176882611727565b602082019050919050565b6000602082019050818103600083015261178c81611750565b9050919050565b7f757365722063616e6e6f74206166666f72642067616d65000000000000000000600082015250565b60006117c9601783611594565b91506117d482611793565b602082019050919050565b600060208201905081810360008301526117f8816117bc565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061186882611037565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361189a5761189961182e565b5b600182019050919050565b7f7573657220616c7265616479206f776e732067616d6500000000000000000000600082015250565b60006118db601683611594565b91506118e6826118a5565b602082019050919050565b6000602082019050818103600083015261190a816118ce565b9050919050565b7f6e6f20726f6f74206861736820676976656e0000000000000000000000000000600082015250565b6000611947601283611594565b915061195282611911565b602082019050919050565b600060208201905081810360008301526119768161193a565b9050919050565b7f6e6f2049504653206164647265737320676976656e20666f722068617368207460008201527f7265656500000000000000000000000000000000000000000000000000000000602082015250565b60006119d9602483611594565b91506119e48261197d565b604082019050919050565b60006020820190508181036000830152611a08816119cc565b9050919050565b7f70726576696f75732076657273696f6e206f662067616d65206e6f7420666f7560008201527f6e64000000000000000000000000000000000000000000000000000000000000602082015250565b6000611a6b602283611594565b9150611a7682611a0f565b604082019050919050565b60006020820190508181036000830152611a9a81611a5e565b9050919050565b7f6f6e6c7920746865206f726967696e616c2075706c6f616465722063616e207560008201527f70646174652074686569722067616d6500000000000000000000000000000000602082015250565b6000611afd603083611594565b9150611b0882611aa1565b604082019050919050565b60006020820190508181036000830152611b2c81611af0565b9050919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302611b957fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611b58565b611b9f8683611b58565b95508019841693508086168417925050509392505050565b6000819050919050565b6000611bdc611bd7611bd284611037565b611bb7565b611037565b9050919050565b6000819050919050565b611bf683611bc1565b611c0a611c0282611be3565b848454611b65565b825550505050565b600090565b611c1f611c12565b611c2a818484611bed565b505050565b5b81811015611c4e57611c43600082611c17565b600181019050611c30565b5050565b601f821115611c9357611c6481611b33565b611c6d84611b48565b81016020851015611c7c578190505b611c90611c8885611b48565b830182611c2f565b50505b505050565b600082821c905092915050565b6000611cb660001984600802611c98565b1980831691505092915050565b6000611ccf8383611ca5565b9150826002028217905092915050565b611ce882611589565b67ffffffffffffffff811115611d0157611d00611081565b5b611d0b82546116f6565b611d16828285611c52565b600060209050601f831160018114611d495760008415611d37578287015190505b611d418582611cc3565b865550611da9565b601f198416611d5786611b33565b60005b82811015611d7f57848901518255600182019150602085019450602081019050611d5a565b86831015611d9c5784890151611d98601f891682611ca5565b8355505b6001600288020188555050505b505050505050565b600082825260208201905092915050565b6000611dcd82611589565b611dd78185611db1565b9350611de78185602086016115a5565b611df081611070565b840191505092915050565b611e0481610fd4565b82525050565b611e1381611037565b82525050565b611e2281611207565b82525050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b611e5d81611276565b82525050565b6000611e6f8383611e54565b60208301905092915050565b6000602082019050919050565b6000611e9382611e28565b611e9d8185611e33565b9350611ea883611e44565b8060005b83811015611ed9578151611ec08882611e63565b9750611ecb83611e7b565b925050600181019050611eac565b5085935050505092915050565b6000610140830160008301518482036000860152611f048282611dc2565b91505060208301518482036020860152611f1e8282611dc2565b91505060408301518482036040860152611f388282611dc2565b91505060608301518482036060860152611f528282611dc2565b9150506080830151611f676080860182611dfb565b5060a0830151611f7a60a0860182611dfb565b5060c0830151611f8d60c0860182611e0a565b5060e0830151611fa060e0860182611e19565b50610100830151848203610100860152611fba8282611e88565b915050610120830151848203610120860152611fd68282611dc2565b9150508091505092915050565b6000604082019050611ff8600083018561155f565b818103602083015261200a8184611ee6565b9050939250505056fea2646970667358221220c47d9ddc1a445224c89cd5a05c1f7ed6c0b99e840c66da8e35f9b2dcc8ba8d8264736f6c637828302e382e31392d646576656c6f702e323032332e322e31302b636f6d6d69742e35396639616234640059",
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

// UploadGame is a paid mutator transaction binding the contract method 0x90691d6a.
//
// Solidity: function uploadGame((string,string,string,string,bytes32,bytes32,uint256,address,address[],string) _game) returns()
func (_Library *LibraryTransactor) UploadGame(opts *bind.TransactOpts, _game LibraryGameEntry) (*types.Transaction, error) {
	return _Library.contract.Transact(opts, "uploadGame", _game)
}

// UploadGame is a paid mutator transaction binding the contract method 0x90691d6a.
//
// Solidity: function uploadGame((string,string,string,string,bytes32,bytes32,uint256,address,address[],string) _game) returns()
func (_Library *LibrarySession) UploadGame(_game LibraryGameEntry) (*types.Transaction, error) {
	return _Library.Contract.UploadGame(&_Library.TransactOpts, _game)
}

// UploadGame is a paid mutator transaction binding the contract method 0x90691d6a.
//
// Solidity: function uploadGame((string,string,string,string,bytes32,bytes32,uint256,address,address[],string) _game) returns()
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

// FilterNewGame is a free log retrieval operation binding the contract event 0x1e55381ba3dfad1800fac2963559018a3bf1b9b90b187ca535709c9f0905cc05.
//
// Solidity: event NewGame(bytes32 hash, (string,string,string,string,bytes32,bytes32,uint256,address,address[],string) game)
func (_Library *LibraryFilterer) FilterNewGame(opts *bind.FilterOpts) (*LibraryNewGameIterator, error) {

	logs, sub, err := _Library.contract.FilterLogs(opts, "NewGame")
	if err != nil {
		return nil, err
	}
	return &LibraryNewGameIterator{contract: _Library.contract, event: "NewGame", logs: logs, sub: sub}, nil
}

// WatchNewGame is a free log subscription operation binding the contract event 0x1e55381ba3dfad1800fac2963559018a3bf1b9b90b187ca535709c9f0905cc05.
//
// Solidity: event NewGame(bytes32 hash, (string,string,string,string,bytes32,bytes32,uint256,address,address[],string) game)
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

// ParseNewGame is a log parse operation binding the contract event 0x1e55381ba3dfad1800fac2963559018a3bf1b9b90b187ca535709c9f0905cc05.
//
// Solidity: event NewGame(bytes32 hash, (string,string,string,string,bytes32,bytes32,uint256,address,address[],string) game)
func (_Library *LibraryFilterer) ParseNewGame(log types.Log) (*LibraryNewGame, error) {
	event := new(LibraryNewGame)
	if err := _Library.contract.UnpackLog(event, "NewGame", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
