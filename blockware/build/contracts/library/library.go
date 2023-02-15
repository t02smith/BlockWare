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
	Bin: "0x608060405234801561001057600080fd5b50612045806100206000396000f3fe6080604052600436106100555760003560e01c80633e093f791461005a57806350e0c46e14610076578063741df964146100a157806390691d6a146100de578063dc164c8214610107578063f579f88214610144575b600080fd5b610074600480360381019061006f9190611006565b610189565b005b34801561008257600080fd5b5061008b610407565b604051610098919061104c565b60405180910390f35b3480156100ad57600080fd5b506100c860048036038101906100c39190611006565b610414565b6040516100d5919061104c565b60405180910390f35b3480156100ea57600080fd5b50610105600480360381019061010091906114e5565b610499565b005b34801561011357600080fd5b5061012e6004803603810190610129919061152e565b610bdb565b60405161013b919061156a565b60405180910390f35b34801561015057600080fd5b5061016b60048036038101906101669190611006565b610bff565b60405161018099989796959493929190611613565b60405180910390f35b600080600083815260200190815260200160002060000180546101ab906116f2565b9050116101ed576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101e49061176f565b60405180910390fd5b60008060008381526020019081526020016000209050806006015434101561024a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610241906117db565b60405180910390fd5b6000805b82600801805490508110156102ec573373ffffffffffffffffffffffffffffffffffffffff1683600801828154811061028a576102896117fb565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036102d957600191506102ec565b80806102e490611859565b91505061024e565b50801561032e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610325906118ed565b60405180910390fd5b8160070160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc83600601549081150290604051600060405180830381858888f1935050505015801561039c573d6000803e3d6000fd5b5081600801339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050565b6000600180549050905090565b6000806000808481526020019081526020016000206000018054610437906116f2565b905011610479576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104709061176f565b60405180910390fd5b600080838152602001908152602001600020600801805490509050919050565b6000816080015150602060ff16116104e6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104dd90611959565b60405180910390fd5b6000816101200151511161052f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610526906119eb565b60405180910390fd5b6000801b8160a0015114610a2a5760008060008360a0015181526020019081526020016000206000018054610563906116f2565b9050116105a5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161059c90611a7d565b60405180910390fd5b60008060008360a001518152602001908152602001600020604051806101400160405290816000820180546105d9906116f2565b80601f0160208091040260200160405190810160405280929190818152602001828054610605906116f2565b80156106525780601f1061062757610100808354040283529160200191610652565b820191906000526020600020905b81548152906001019060200180831161063557829003601f168201915b5050505050815260200160018201805461066b906116f2565b80601f0160208091040260200160405190810160405280929190818152602001828054610697906116f2565b80156106e45780601f106106b9576101008083540402835291602001916106e4565b820191906000526020600020905b8154815290600101906020018083116106c757829003601f168201915b505050505081526020016002820180546106fd906116f2565b80601f0160208091040260200160405190810160405280929190818152602001828054610729906116f2565b80156107765780601f1061074b57610100808354040283529160200191610776565b820191906000526020600020905b81548152906001019060200180831161075957829003601f168201915b5050505050815260200160038201805461078f906116f2565b80601f01602080910402602001604051908101604052809291908181526020018280546107bb906116f2565b80156108085780601f106107dd57610100808354040283529160200191610808565b820191906000526020600020905b8154815290600101906020018083116107eb57829003601f168201915b505050505081526020016004820154815260200160058201548152602001600682015481526020016007820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016008820180548060200260200160405190810160405280929190818152602001828054801561090a57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116108c0575b50505050508152602001600982018054610923906116f2565b80601f016020809104026020016040519081016040528092919081815260200182805461094f906116f2565b801561099c5780601f106109715761010080835404028352916020019161099c565b820191906000526020600020905b81548152906001019060200180831161097f57829003601f168201915b50505050508152505090503373ffffffffffffffffffffffffffffffffffffffff168160e0015173ffffffffffffffffffffffffffffffffffffffff1614610a19576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a1090611b0f565b60405180910390fd5b806101000151826101000181905250505b338160e0019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505080600080836080015181526020019081526020016000206000820151816000019081610a8e9190611cdb565b506020820151816001019081610aa49190611cdb565b506040820151816002019081610aba9190611cdb565b506060820151816003019081610ad09190611cdb565b506080820151816004015560a0820151816005015560c0820151816006015560e08201518160070160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610100820151816008019080519060200190610b53929190610f15565b50610120820151816009019081610b6a9190611cdb565b509050506001816080015190806001815401808255809150506001900390600052602060002001600090919091909150557f1e55381ba3dfad1800fac2963559018a3bf1b9b90b187ca535709c9f0905cc05816080015182604051610bd0929190611fdf565b60405180910390a150565b60018181548110610beb57600080fd5b906000526020600020016000915090505481565b6000602052806000526040600020600091509050806000018054610c22906116f2565b80601f0160208091040260200160405190810160405280929190818152602001828054610c4e906116f2565b8015610c9b5780601f10610c7057610100808354040283529160200191610c9b565b820191906000526020600020905b815481529060010190602001808311610c7e57829003601f168201915b505050505090806001018054610cb0906116f2565b80601f0160208091040260200160405190810160405280929190818152602001828054610cdc906116f2565b8015610d295780601f10610cfe57610100808354040283529160200191610d29565b820191906000526020600020905b815481529060010190602001808311610d0c57829003601f168201915b505050505090806002018054610d3e906116f2565b80601f0160208091040260200160405190810160405280929190818152602001828054610d6a906116f2565b8015610db75780601f10610d8c57610100808354040283529160200191610db7565b820191906000526020600020905b815481529060010190602001808311610d9a57829003601f168201915b505050505090806003018054610dcc906116f2565b80601f0160208091040260200160405190810160405280929190818152602001828054610df8906116f2565b8015610e455780601f10610e1a57610100808354040283529160200191610e45565b820191906000526020600020905b815481529060010190602001808311610e2857829003601f168201915b5050505050908060040154908060050154908060060154908060070160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806009018054610e92906116f2565b80601f0160208091040260200160405190810160405280929190818152602001828054610ebe906116f2565b8015610f0b5780601f10610ee057610100808354040283529160200191610f0b565b820191906000526020600020905b815481529060010190602001808311610eee57829003601f168201915b5050505050905089565b828054828255906000526020600020908101928215610f8e579160200282015b82811115610f8d5782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555091602001919060010190610f35565b5b509050610f9b9190610f9f565b5090565b5b80821115610fb8576000816000905550600101610fa0565b5090565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b610fe381610fd0565b8114610fee57600080fd5b50565b60008135905061100081610fda565b92915050565b60006020828403121561101c5761101b610fc6565b5b600061102a84828501610ff1565b91505092915050565b6000819050919050565b61104681611033565b82525050565b6000602082019050611061600083018461103d565b92915050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6110b58261106c565b810181811067ffffffffffffffff821117156110d4576110d361107d565b5b80604052505050565b60006110e7610fbc565b90506110f382826110ac565b919050565b600080fd5b600080fd5b600080fd5b600067ffffffffffffffff8211156111225761112161107d565b5b61112b8261106c565b9050602081019050919050565b82818337600083830152505050565b600061115a61115584611107565b6110dd565b90508281526020810184848401111561117657611175611102565b5b611181848285611138565b509392505050565b600082601f83011261119e5761119d6110fd565b5b81356111ae848260208601611147565b91505092915050565b6111c081611033565b81146111cb57600080fd5b50565b6000813590506111dd816111b7565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061120e826111e3565b9050919050565b61121e81611203565b811461122957600080fd5b50565b60008135905061123b81611215565b92915050565b600067ffffffffffffffff82111561125c5761125b61107d565b5b602082029050602081019050919050565b600080fd5b600061127d826111e3565b9050919050565b61128d81611272565b811461129857600080fd5b50565b6000813590506112aa81611284565b92915050565b60006112c36112be84611241565b6110dd565b905080838252602082019050602084028301858111156112e6576112e561126d565b5b835b8181101561130f57806112fb888261129b565b8452602084019350506020810190506112e8565b5050509392505050565b600082601f83011261132e5761132d6110fd565b5b813561133e8482602086016112b0565b91505092915050565b6000610140828403121561135e5761135d611067565b5b6113696101406110dd565b9050600082013567ffffffffffffffff811115611389576113886110f8565b5b61139584828501611189565b600083015250602082013567ffffffffffffffff8111156113b9576113b86110f8565b5b6113c584828501611189565b602083015250604082013567ffffffffffffffff8111156113e9576113e86110f8565b5b6113f584828501611189565b604083015250606082013567ffffffffffffffff811115611419576114186110f8565b5b61142584828501611189565b606083015250608061143984828501610ff1565b60808301525060a061144d84828501610ff1565b60a08301525060c0611461848285016111ce565b60c08301525060e06114758482850161122c565b60e08301525061010082013567ffffffffffffffff81111561149a576114996110f8565b5b6114a684828501611319565b6101008301525061012082013567ffffffffffffffff8111156114cc576114cb6110f8565b5b6114d884828501611189565b6101208301525092915050565b6000602082840312156114fb576114fa610fc6565b5b600082013567ffffffffffffffff81111561151957611518610fcb565b5b61152584828501611347565b91505092915050565b60006020828403121561154457611543610fc6565b5b6000611552848285016111ce565b91505092915050565b61156481610fd0565b82525050565b600060208201905061157f600083018461155b565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156115bf5780820151818401526020810190506115a4565b60008484015250505050565b60006115d682611585565b6115e08185611590565b93506115f08185602086016115a1565b6115f98161106c565b840191505092915050565b61160d81611203565b82525050565b600061012082019050818103600083015261162e818c6115cb565b90508181036020830152611642818b6115cb565b90508181036040830152611656818a6115cb565b9050818103606083015261166a81896115cb565b9050611679608083018861155b565b61168660a083018761155b565b61169360c083018661103d565b6116a060e0830185611604565b8181036101008301526116b381846115cb565b90509a9950505050505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061170a57607f821691505b60208210810361171d5761171c6116c3565b5b50919050565b7f67616d65206e6f7420666f756e64000000000000000000000000000000000000600082015250565b6000611759600e83611590565b915061176482611723565b602082019050919050565b600060208201905081810360008301526117888161174c565b9050919050565b7f757365722063616e6e6f74206166666f72642067616d65000000000000000000600082015250565b60006117c5601783611590565b91506117d08261178f565b602082019050919050565b600060208201905081810360008301526117f4816117b8565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061186482611033565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036118965761189561182a565b5b600182019050919050565b7f7573657220616c7265616479206f776e732067616d6500000000000000000000600082015250565b60006118d7601683611590565b91506118e2826118a1565b602082019050919050565b60006020820190508181036000830152611906816118ca565b9050919050565b7f6e6f20726f6f74206861736820676976656e0000000000000000000000000000600082015250565b6000611943601283611590565b915061194e8261190d565b602082019050919050565b6000602082019050818103600083015261197281611936565b9050919050565b7f6e6f2049504653206164647265737320676976656e20666f722068617368207460008201527f7265656500000000000000000000000000000000000000000000000000000000602082015250565b60006119d5602483611590565b91506119e082611979565b604082019050919050565b60006020820190508181036000830152611a04816119c8565b9050919050565b7f70726576696f75732076657273696f6e206f662067616d65206e6f7420666f7560008201527f6e64000000000000000000000000000000000000000000000000000000000000602082015250565b6000611a67602283611590565b9150611a7282611a0b565b604082019050919050565b60006020820190508181036000830152611a9681611a5a565b9050919050565b7f6f6e6c7920746865206f726967696e616c2075706c6f616465722063616e207560008201527f70646174652074686569722067616d6500000000000000000000000000000000602082015250565b6000611af9603083611590565b9150611b0482611a9d565b604082019050919050565b60006020820190508181036000830152611b2881611aec565b9050919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302611b917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611b54565b611b9b8683611b54565b95508019841693508086168417925050509392505050565b6000819050919050565b6000611bd8611bd3611bce84611033565b611bb3565b611033565b9050919050565b6000819050919050565b611bf283611bbd565b611c06611bfe82611bdf565b848454611b61565b825550505050565b600090565b611c1b611c0e565b611c26818484611be9565b505050565b5b81811015611c4a57611c3f600082611c13565b600181019050611c2c565b5050565b601f821115611c8f57611c6081611b2f565b611c6984611b44565b81016020851015611c78578190505b611c8c611c8485611b44565b830182611c2b565b50505b505050565b600082821c905092915050565b6000611cb260001984600802611c94565b1980831691505092915050565b6000611ccb8383611ca1565b9150826002028217905092915050565b611ce482611585565b67ffffffffffffffff811115611cfd57611cfc61107d565b5b611d0782546116f2565b611d12828285611c4e565b600060209050601f831160018114611d455760008415611d33578287015190505b611d3d8582611cbf565b865550611da5565b601f198416611d5386611b2f565b60005b82811015611d7b57848901518255600182019150602085019450602081019050611d56565b86831015611d985784890151611d94601f891682611ca1565b8355505b6001600288020188555050505b505050505050565b600082825260208201905092915050565b6000611dc982611585565b611dd38185611dad565b9350611de38185602086016115a1565b611dec8161106c565b840191505092915050565b611e0081610fd0565b82525050565b611e0f81611033565b82525050565b611e1e81611203565b82525050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b611e5981611272565b82525050565b6000611e6b8383611e50565b60208301905092915050565b6000602082019050919050565b6000611e8f82611e24565b611e998185611e2f565b9350611ea483611e40565b8060005b83811015611ed5578151611ebc8882611e5f565b9750611ec783611e77565b925050600181019050611ea8565b5085935050505092915050565b6000610140830160008301518482036000860152611f008282611dbe565b91505060208301518482036020860152611f1a8282611dbe565b91505060408301518482036040860152611f348282611dbe565b91505060608301518482036060860152611f4e8282611dbe565b9150506080830151611f636080860182611df7565b5060a0830151611f7660a0860182611df7565b5060c0830151611f8960c0860182611e06565b5060e0830151611f9c60e0860182611e15565b50610100830151848203610100860152611fb68282611e84565b915050610120830151848203610120860152611fd28282611dbe565b9150508091505092915050565b6000604082019050611ff4600083018561155b565b81810360208301526120068184611ee2565b9050939250505056fea2646970667358221220ee57bc576109d4e0842067a9027bac0478dfcea4125b48b1f18feda673c0402464736f6c63430008120033",
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
