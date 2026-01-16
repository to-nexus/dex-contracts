// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package binding

import (
	"bytes"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = bytes.Equal
	_ = errors.New
	_ = big.NewInt
	_ = common.Big1
	_ = types.BloomLookup
	_ = abi.ConvertType
)

// IMarketV2FeeConfig is an auto generated low-level Go binding around an user-defined struct.
type IMarketV2FeeConfig struct {
	SellerMakerFeeBps uint32
	SellerTakerFeeBps uint32
	BuyerMakerFeeBps  uint32
	BuyerTakerFeeBps  uint32
}

// MarketImplV2MetaData contains all meta data concerning the MarketImplV2 contract.
var MarketImplV2MetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CROSS_DEX\",\"outputs\":[{\"internalType\":\"contractICrossDex\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"QUOTE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROUTER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allPairs\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"bases\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"pairs\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"}],\"name\":\"baseToPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"checkTickSizeRoles\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tickSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lotSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"name\":\"createPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deployed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeCollector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"sellerMakerFeeBps\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"sellerTakerFeeBps\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"buyerMakerFeeBps\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"buyerTakerFeeBps\",\"type\":\"uint32\"}],\"internalType\":\"structIMarketV2.FeeConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_quote\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pairImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeCollector\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pairImpl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeCollector\",\"type\":\"address\"}],\"name\":\"setFeeCollector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_sellerMakerFeeBps\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_sellerTakerFeeBps\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_buyerMakerFeeBps\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_buyerTakerFeeBps\",\"type\":\"uint32\"}],\"name\":\"setMarketFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pairImpl\",\"type\":\"address\"}],\"name\":\"setPairImpl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"before\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"FeeCollectorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"sellerMakerFee\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"sellerTakerFee\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"buyerMakerFee\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"buyerTakerFee\",\"type\":\"uint32\"}],\"name\":\"MarketFeesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"PairCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"before\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"PairImplSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Create2EmptyBytecode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"EnumerableMapNonexistentKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedDeployment\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"MarketAlreadyCreatedBaseAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MarketDeployPair\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"MarketInvalidBaseAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"makerFee\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"takerFee\",\"type\":\"uint32\"}],\"name\":\"MarketInvalidFeeStructure\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"MarketInvalidInitializeData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"}]",
	ID:  "MarketImplV2",
	Bin: "0x60a060405230608052348015610013575f5ffd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b608051612ac86100f95f395f818161198a015281816119b30152611bd40152612ac85ff3fe608060405260043610610162575f3560e01c8063715018a6116100c6578063ad3cb1cc1161007c578063c97682f811610057578063c97682f814610506578063f2fde38b14610528578063f905c15a14610547575f5ffd5b8063ad3cb1cc14610459578063b820d8aa146104ae578063c415b95c146104da575f5ffd5b80639c579839116100ac5780639c579839146103ef578063a1eea7781461041b578063a42dce801461043a575f5ffd5b8063715018a6146103925780638da5cb5b146103a6575f5ffd5b80634732c7e51161011b57806352d1902d1161010157806352d1902d14610279578063584f6af31461029b5780635fbbc0d2146102ba575f5ffd5b80634732c7e5146102475780634f1ef28614610266575f5ffd5b806335f7d4561161014b57806335f7d456146101dd578063364e9a19146102095780633cc047aa14610228575f5ffd5b806301f79fb61461016657806332fe7b2614610187575b5f5ffd5b348015610171575f5ffd5b50610185610180366004612297565b61055b565b005b348015610192575f5ffd5b506003546101b39073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b3480156101e8575f5ffd5b506004546101b39073ffffffffffffffffffffffffffffffffffffffff1681565b348015610214575f5ffd5b506101b3610223366004612321565b610a83565b348015610233575f5ffd5b50610185610242366004612321565b610a95565b348015610252575f5ffd5b506101b361026136600461233a565b610b99565b610185610274366004612397565b6110b2565b348015610284575f5ffd5b5061028d6110d1565b6040519081526020016101d4565b3480156102a6575f5ffd5b506101856102b53660046123f3565b6110ff565b3480156102c5575f5ffd5b50610343604080516080810182525f808252602082018190529181018290526060810191909152506040805160808101825260095463ffffffff80821683526401000000008204811660208401526801000000000000000082048116938301939093526c010000000000000000000000009004909116606082015290565b6040516101d491905f60808201905063ffffffff835116825263ffffffff602084015116602083015263ffffffff604084015116604083015263ffffffff606084015116606083015292915050565b34801561039d575f5ffd5b50610185611119565b3480156103b1575f5ffd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c1993005473ffffffffffffffffffffffffffffffffffffffff166101b3565b3480156103fa575f5ffd5b506002546101b39073ffffffffffffffffffffffffffffffffffffffff1681565b348015610426575f5ffd5b50610185610435366004612321565b61112c565b348015610445575f5ffd5b50610185610454366004612321565b6111ae565b348015610464575f5ffd5b506104a16040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b6040516101d49190612498565b3480156104b9575f5ffd5b506001546101b39073ffffffffffffffffffffffffffffffffffffffff1681565b3480156104e5575f5ffd5b506005546101b39073ffffffffffffffffffffffffffffffffffffffff1681565b348015610511575f5ffd5b5061051a6112b2565b6040516101d49291906124fa565b348015610533575f5ffd5b50610185610542366004612321565b6113c1565b348015610552575f5ffd5b5061028d5f5481565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff165f811580156105a55750825b90505f8267ffffffffffffffff1660011480156105c15750303b155b9050811580156105cf575080155b15610606576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000016600117855583156106675784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b6106708b611424565b73ffffffffffffffffffffffffffffffffffffffff8b166106e4576040517fa25d2db50000000000000000000000000000000000000000000000000000000081527f6f776e657200000000000000000000000000000000000000000000000000000060048201526024015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8a16610753576040517fa25d2db50000000000000000000000000000000000000000000000000000000081527f726f75746572000000000000000000000000000000000000000000000000000060048201526024016106db565b73ffffffffffffffffffffffffffffffffffffffff89166107c2576040517fa25d2db50000000000000000000000000000000000000000000000000000000081527f71756f746500000000000000000000000000000000000000000000000000000060048201526024016106db565b73ffffffffffffffffffffffffffffffffffffffff8816610831576040517fa25d2db50000000000000000000000000000000000000000000000000000000081527f70616972496d706c00000000000000000000000000000000000000000000000060048201526024016106db565b73ffffffffffffffffffffffffffffffffffffffff87166108a0576040517fa25d2db50000000000000000000000000000000000000000000000000000000081527f666565436f6c6c6563746f72000000000000000000000000000000000000000060048201526024016106db565b5f5f5f5f898060200190518101906108b8919061251e565b435f55929650909450925090503360015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508c60025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508d60035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508b60045f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508a60055f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610a1184848484611435565b505050508315610a765784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050505050565b5f610a8f600683611782565b92915050565b610a9d6117aa565b73ffffffffffffffffffffffffffffffffffffffff8116610b0c576040517fa25d2db50000000000000000000000000000000000000000000000000000000081527f70616972496d706c00000000000000000000000000000000000000000000000060048201526024016106db565b60045460405173ffffffffffffffffffffffffffffffffffffffff8084169216907f685c330e9b04243f99ff39d8de8578c9af2bfaff120e17a1b5f1d61e004935e2905f90a3600480547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b5f610ba26117aa565b73ffffffffffffffffffffffffffffffffffffffff85161580610bdf575060025473ffffffffffffffffffffffffffffffffffffffff8681169116145b15610c2e576040517f20f8254c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff861660048201526024016106db565b5f8573ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa158015610c78573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c9c919061256f565b60ff169050805f03610cf2576040517f20f8254c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff871660048201526024016106db565b610cfd600687611838565b15610d4c576040517fc13fcfdc00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff871660048201526024016106db565b5f60405180602001610d5d9061216d565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082820381018352601f90910116604081905260045460035460025473ffffffffffffffffffffffffffffffffffffffff92831693610dcf939283169291909116908c908c908c908c9060240161258f565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f02a204c6000000000000000000000000000000000000000000000000000000001790529051610e549392910161260d565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815290829052610e909291602001612652565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290527fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060608a901b16602083015291505f906034016040516020818303038152906040528051906020012090505f610f165f8385611859565b905073ffffffffffffffffffffffffffffffffffffffff8116610f65576040517f3ff6a26100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610f7160068a83611945565b610fbf576040517fc13fcfdc00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8a1660048201526024016106db565b6001546040517fe9f7206b00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83811660048301529091169063e9f7206b906024015f604051808303815f87803b158015611029575f5ffd5b505af115801561103b573d5f5f3e3d5ffd5b505050508873ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167fc1db9ba7c4b7ce660fe8d17bbcf07167549381df2abd694a970bd1402d86d3134260405161109e91815260200190565b60405180910390a398975050505050505050565b6110ba611972565b6110c382611a76565b6110cd8282611a7e565b5050565b5f6110da611bbc565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b6111076117aa565b61111384848484611435565b50505050565b6111216117aa565b61112a5f611c2b565b565b6001546040517fa1eea77800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83811660048301529091169063a1eea778906024015f6040518083038186803b158015611195575f5ffd5b505afa1580156111a7573d5f5f3e3d5ffd5b5050505050565b6111b66117aa565b73ffffffffffffffffffffffffffffffffffffffff8116611225576040517fa25d2db50000000000000000000000000000000000000000000000000000000081527f666565436f6c6c6563746f72000000000000000000000000000000000000000060048201526024016106db565b60055460405173ffffffffffffffffffffffffffffffffffffffff8084169216907f649c5e3d0ed183894196148e193af316452b0037e77d2ff0fef23b7dc722bed0905f90a3600580547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b6060805f6112c06006611cc0565b90508067ffffffffffffffff8111156112db576112db6121a2565b604051908082528060200260200182016040528015611304578160200160208202803683370190505b5092508067ffffffffffffffff811115611320576113206121a2565b604051908082528060200260200182016040528015611349578160200160208202803683370190505b5091505f5b818110156113bb57611361600682611cca565b85838151811061137357611373612666565b6020026020010185848151811061138c5761138c612666565b73ffffffffffffffffffffffffffffffffffffffff93841660209182029290920101529116905260010161134e565b50509091565b6113c96117aa565b73ffffffffffffffffffffffffffffffffffffffff8116611418576040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081525f60048201526024016106db565b61142181611c2b565b50565b61142c611ce5565b61142181611d4c565b61271063ffffffff851610611498576040517fa25d2db50000000000000000000000000000000000000000000000000000000081527f73656c6c65724d616b657246656542707300000000000000000000000000000060048201526024016106db565b61271063ffffffff8416106114fb576040517fa25d2db50000000000000000000000000000000000000000000000000000000081527f73656c6c657254616b657246656542707300000000000000000000000000000060048201526024016106db565b61271063ffffffff83161061155e576040517fa25d2db50000000000000000000000000000000000000000000000000000000081527f62757965724d616b65724665654270730000000000000000000000000000000060048201526024016106db565b61271063ffffffff8216106115c1576040517fa25d2db50000000000000000000000000000000000000000000000000000000081527f627579657254616b65724665654270730000000000000000000000000000000060048201526024016106db565b8363ffffffff168363ffffffff161015611617576040517f38642d5f00000000000000000000000000000000000000000000000000000000815263ffffffff8086166004830152841660248201526044016106db565b8163ffffffff168163ffffffff16101561166d576040517f38642d5f00000000000000000000000000000000000000000000000000000000815263ffffffff8084166004830152821660248201526044016106db565b604080516080808201835263ffffffff87811680845287821660208086018290528884168688018190529388166060968701819052600980547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001685176401000000008502177fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff166801000000000000000087027fffffffffffffffffffffffffffffffff00000000ffffffffffffffffffffffff16176c01000000000000000000000000830217905587519384529083019190915294810191909152918201929092527f68657e5573176ce9149cd289efa48a58a9db7703c3b1bef2e30e41647f5ab208910160405180910390a150505050565b5f6117a38373ffffffffffffffffffffffffffffffffffffffff8416611d54565b9392505050565b336117e97f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c1993005473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff161461112a576040517f118cdaa70000000000000000000000000000000000000000000000000000000081523360048201526024016106db565b5f6117a38373ffffffffffffffffffffffffffffffffffffffff8416611db3565b5f8347101561189d576040517fcf479181000000000000000000000000000000000000000000000000000000008152476004820152602481018590526044016106db565b81515f036118d7576040517f4ca249dc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8282516020840186f590503d1519811516156118f8576040513d5f823e3d81fd5b73ffffffffffffffffffffffffffffffffffffffff81166117a3576040517fb06ebf3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f61196a8473ffffffffffffffffffffffffffffffffffffffff808616908516611dbe565b949350505050565b3073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161480611a3f57507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16611a267f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614155b1561112a576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6114216117aa565b8173ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015611b03575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201909252611b0091810190612693565b60015b611b51576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff831660048201526024016106db565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8114611bad576040517faa1d49a4000000000000000000000000000000000000000000000000000000008152600481018290526024016106db565b611bb78383611dda565b505050565b3073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161461112a576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080547fffffffffffffffffffffffff0000000000000000000000000000000000000000811673ffffffffffffffffffffffffffffffffffffffff848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b5f610a8f82611e3c565b5f808080611cd88686611e46565b9097909650945050505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff1661112a576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6113c9611ce5565b5f81815260028301602052604081205480158015611d795750611d778484611db3565b155b156117a3576040517f02b56686000000000000000000000000000000000000000000000000000000008152600481018490526024016106db565b5f6117a38383611e6f565b5f828152600284016020526040812082905561196a8484611e86565b611de382611e91565b60405173ffffffffffffffffffffffffffffffffffffffff8316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115611e3457611bb78282611f5f565b6110cd611fde565b5f610a8f82612016565b5f8080611e53858561201f565b5f81815260029690960160205260409095205494959350505050565b5f81815260018301602052604081205415156117a3565b5f6117a3838361202a565b8073ffffffffffffffffffffffffffffffffffffffff163b5f03611ef9576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024016106db565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60605f5f8473ffffffffffffffffffffffffffffffffffffffff1684604051611f8891906126aa565b5f60405180830381855af49150503d805f8114611fc0576040519150601f19603f3d011682016040523d82523d5f602084013e611fc5565b606091505b5091509150611fd5858383612076565b95945050505050565b341561112a576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f610a8f825490565b5f6117a38383612105565b5f81815260018301602052604081205461206f57508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610a8f565b505f610a8f565b60608261208b576120868261212b565b6117a3565b81511580156120af575073ffffffffffffffffffffffffffffffffffffffff84163b155b156120fe576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024016106db565b50806117a3565b5f825f01828154811061211a5761211a612666565b905f5260205f200154905092915050565b80511561213b5780518082602001fd5b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103dd806126b683390190565b803573ffffffffffffffffffffffffffffffffffffffff8116811461219d575f5ffd5b919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f82601f8301126121de575f5ffd5b813567ffffffffffffffff8111156121f8576121f86121a2565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff82111715612264576122646121a2565b60405281815283820160200185101561227b575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f5f5f5f5f5f60c087890312156122ac575f5ffd5b6122b58761217a565b95506122c36020880161217a565b94506122d16040880161217a565b93506122df6060880161217a565b92506122ed6080880161217a565b915060a087013567ffffffffffffffff811115612308575f5ffd5b61231489828a016121cf565b9150509295509295509295565b5f60208284031215612331575f5ffd5b6117a38261217a565b5f5f5f5f6080858703121561234d575f5ffd5b6123568561217a565b93506020850135925060408501359150606085013567ffffffffffffffff81111561237f575f5ffd5b61238b878288016121cf565b91505092959194509250565b5f5f604083850312156123a8575f5ffd5b6123b18361217a565b9150602083013567ffffffffffffffff8111156123cc575f5ffd5b6123d8858286016121cf565b9150509250929050565b63ffffffff81168114611421575f5ffd5b5f5f5f5f60808587031215612406575f5ffd5b8435612411816123e2565b93506020850135612421816123e2565b92506040850135612431816123e2565b91506060850135612441816123e2565b939692955090935050565b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f6117a3602083018461244c565b5f8151808452602084019350602083015f5b828110156124f057815173ffffffffffffffffffffffffffffffffffffffff168652602095860195909101906001016124bc565b5093949350505050565b604081525f61250c60408301856124aa565b8281036020840152611fd581856124aa565b5f5f5f5f60808587031215612531575f5ffd5b845161253c816123e2565b602086015190945061254d816123e2565b604086015190935061255e816123e2565b6060860151909250612441816123e2565b5f6020828403121561257f575f5ffd5b815160ff811681146117a3575f5ffd5b73ffffffffffffffffffffffffffffffffffffffff8716815273ffffffffffffffffffffffffffffffffffffffff8616602082015273ffffffffffffffffffffffffffffffffffffffff8516604082015283606082015282608082015260c060a08201525f61260160c083018461244c565b98975050505050505050565b73ffffffffffffffffffffffffffffffffffffffff83168152604060208201525f61196a604083018461244c565b5f81518060208401855e5f93019283525090919050565b5f61196a612660838661263b565b8461263b565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f602082840312156126a3575f5ffd5b5051919050565b5f6117a3828461263b56fe60806040526040516103dd3803806103dd8339810160408190526100229161023c565b61002c8282610033565b5050610321565b61003c82610091565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a280511561008557610080828261010c565b505050565b61008d61017f565b5050565b806001600160a01b03163b5f036100cb57604051634c9c8ce360e01b81526001600160a01b03821660048201526024015b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b031684604051610128919061030b565b5f60405180830381855af49150503d805f8114610160576040519150601f19603f3d011682016040523d82523d5f602084013e610165565b606091505b5090925090506101768583836101a0565b95945050505050565b341561019e5760405163b398979f60e01b815260040160405180910390fd5b565b6060826101b5576101b0826101ff565b6101f8565b81511580156101cc57506001600160a01b0384163b155b156101f557604051639996b31560e01b81526001600160a01b03851660048201526024016100c2565b50805b9392505050565b80511561020f5780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b634e487b7160e01b5f52604160045260245ffd5b5f5f6040838503121561024d575f5ffd5b82516001600160a01b0381168114610263575f5ffd5b60208401519092506001600160401b0381111561027e575f5ffd5b8301601f8101851361028e575f5ffd5b80516001600160401b038111156102a7576102a7610228565b604051601f8201601f19908116603f011681016001600160401b03811182821017156102d5576102d5610228565b6040528181528282016020018710156102ec575f5ffd5b8160208401602083015e5f602083830101528093505050509250929050565b5f82518060208501845e5f920191825250919050565b60b08061032d5f395ff3fe6080604052600a600c565b005b60186014601a565b605d565b565b5f60587f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5473ffffffffffffffffffffffffffffffffffffffff1690565b905090565b365f5f375f5f365f845af43d5f5f3e8080156076573d5ff35b3d5ffdfea264697066735822122053cea113bf96df2d97b042d29245e0a41b1290b530767f8b25bfeaed774391b764736f6c634300081c0033a26469706673582212208b32427808fbaad5b47c71f5e845f692bf40e18f969efa7dacbbc05d748a58cb64736f6c634300081c0033",
}

// MarketImplV2 is an auto generated Go binding around an Ethereum contract.
type MarketImplV2 struct {
	abi abi.ABI
}

// NewMarketImplV2 creates a new instance of MarketImplV2.
func NewMarketImplV2() *MarketImplV2 {
	parsed, err := MarketImplV2MetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &MarketImplV2{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *MarketImplV2) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackCROSSDEX is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb820d8aa.
//
// Solidity: function CROSS_DEX() view returns(address)
func (marketImplV2 *MarketImplV2) PackCROSSDEX() []byte {
	enc, err := marketImplV2.abi.Pack("CROSS_DEX")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackCROSSDEX is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb820d8aa.
//
// Solidity: function CROSS_DEX() view returns(address)
func (marketImplV2 *MarketImplV2) UnpackCROSSDEX(data []byte) (common.Address, error) {
	out, err := marketImplV2.abi.Unpack("CROSS_DEX", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackQUOTE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9c579839.
//
// Solidity: function QUOTE() view returns(address)
func (marketImplV2 *MarketImplV2) PackQUOTE() []byte {
	enc, err := marketImplV2.abi.Pack("QUOTE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackQUOTE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9c579839.
//
// Solidity: function QUOTE() view returns(address)
func (marketImplV2 *MarketImplV2) UnpackQUOTE(data []byte) (common.Address, error) {
	out, err := marketImplV2.abi.Unpack("QUOTE", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackROUTER is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x32fe7b26.
//
// Solidity: function ROUTER() view returns(address)
func (marketImplV2 *MarketImplV2) PackROUTER() []byte {
	enc, err := marketImplV2.abi.Pack("ROUTER")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackROUTER is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x32fe7b26.
//
// Solidity: function ROUTER() view returns(address)
func (marketImplV2 *MarketImplV2) UnpackROUTER(data []byte) (common.Address, error) {
	out, err := marketImplV2.abi.Unpack("ROUTER", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackUPGRADEINTERFACEVERSION is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (marketImplV2 *MarketImplV2) PackUPGRADEINTERFACEVERSION() []byte {
	enc, err := marketImplV2.abi.Pack("UPGRADE_INTERFACE_VERSION")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackUPGRADEINTERFACEVERSION is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (marketImplV2 *MarketImplV2) UnpackUPGRADEINTERFACEVERSION(data []byte) (string, error) {
	out, err := marketImplV2.abi.Unpack("UPGRADE_INTERFACE_VERSION", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackAllPairs is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc97682f8.
//
// Solidity: function allPairs() view returns(address[] bases, address[] pairs)
func (marketImplV2 *MarketImplV2) PackAllPairs() []byte {
	enc, err := marketImplV2.abi.Pack("allPairs")
	if err != nil {
		panic(err)
	}
	return enc
}

// AllPairsOutput serves as a container for the return parameters of contract
// method AllPairs.

// UnpackAllPairs is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc97682f8.
//
// Solidity: function allPairs() view returns(address[] bases, address[] pairs)
func (marketImplV2 *MarketImplV2) UnpackAllPairs(data []byte) (AllPairsOutput, error) {
	out, err := marketImplV2.abi.Unpack("allPairs", data)
	outstruct := new(AllPairsOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Bases = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Pairs = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	return *outstruct, err

}

// PackBaseToPair is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x364e9a19.
//
// Solidity: function baseToPair(address base) view returns(address)
func (marketImplV2 *MarketImplV2) PackBaseToPair(base common.Address) []byte {
	enc, err := marketImplV2.abi.Pack("baseToPair", base)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackBaseToPair is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x364e9a19.
//
// Solidity: function baseToPair(address base) view returns(address)
func (marketImplV2 *MarketImplV2) UnpackBaseToPair(data []byte) (common.Address, error) {
	out, err := marketImplV2.abi.Unpack("baseToPair", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackCheckTickSizeRoles is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa1eea778.
//
// Solidity: function checkTickSizeRoles(address account) view returns()
func (marketImplV2 *MarketImplV2) PackCheckTickSizeRoles(account common.Address) []byte {
	enc, err := marketImplV2.abi.Pack("checkTickSizeRoles", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCreatePair is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4732c7e5.
//
// Solidity: function createPair(address base, uint256 tickSize, uint256 lotSize, bytes feeData) returns(address)
func (marketImplV2 *MarketImplV2) PackCreatePair(base common.Address, tickSize *big.Int, lotSize *big.Int, feeData []byte) []byte {
	enc, err := marketImplV2.abi.Pack("createPair", base, tickSize, lotSize, feeData)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackCreatePair is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4732c7e5.
//
// Solidity: function createPair(address base, uint256 tickSize, uint256 lotSize, bytes feeData) returns(address)
func (marketImplV2 *MarketImplV2) UnpackCreatePair(data []byte) (common.Address, error) {
	out, err := marketImplV2.abi.Unpack("createPair", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackDeployed is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf905c15a.
//
// Solidity: function deployed() view returns(uint256)
func (marketImplV2 *MarketImplV2) PackDeployed() []byte {
	enc, err := marketImplV2.abi.Pack("deployed")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDeployed is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf905c15a.
//
// Solidity: function deployed() view returns(uint256)
func (marketImplV2 *MarketImplV2) UnpackDeployed(data []byte) (*big.Int, error) {
	out, err := marketImplV2.abi.Unpack("deployed", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackFeeCollector is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (marketImplV2 *MarketImplV2) PackFeeCollector() []byte {
	enc, err := marketImplV2.abi.Pack("feeCollector")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackFeeCollector is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (marketImplV2 *MarketImplV2) UnpackFeeCollector(data []byte) (common.Address, error) {
	out, err := marketImplV2.abi.Unpack("feeCollector", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackGetFeeConfig is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5fbbc0d2.
//
// Solidity: function getFeeConfig() view returns((uint32,uint32,uint32,uint32))
func (marketImplV2 *MarketImplV2) PackGetFeeConfig() []byte {
	enc, err := marketImplV2.abi.Pack("getFeeConfig")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetFeeConfig is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5fbbc0d2.
//
// Solidity: function getFeeConfig() view returns((uint32,uint32,uint32,uint32))
func (marketImplV2 *MarketImplV2) UnpackGetFeeConfig(data []byte) (IMarketV2FeeConfig, error) {
	out, err := marketImplV2.abi.Unpack("getFeeConfig", data)
	if err != nil {
		return *new(IMarketV2FeeConfig), err
	}
	out0 := *abi.ConvertType(out[0], new(IMarketV2FeeConfig)).(*IMarketV2FeeConfig)
	return out0, err
}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01f79fb6.
//
// Solidity: function initialize(address _owner, address _router, address _quote, address _pairImpl, address _feeCollector, bytes feeData) returns()
func (marketImplV2 *MarketImplV2) PackInitialize(owner common.Address, router common.Address, quote common.Address, pairImpl common.Address, feeCollector common.Address, feeData []byte) []byte {
	enc, err := marketImplV2.abi.Pack("initialize", owner, router, quote, pairImpl, feeCollector, feeData)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (marketImplV2 *MarketImplV2) PackOwner() []byte {
	enc, err := marketImplV2.abi.Pack("owner")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (marketImplV2 *MarketImplV2) UnpackOwner(data []byte) (common.Address, error) {
	out, err := marketImplV2.abi.Unpack("owner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackPairImpl is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x35f7d456.
//
// Solidity: function pairImpl() view returns(address)
func (marketImplV2 *MarketImplV2) PackPairImpl() []byte {
	enc, err := marketImplV2.abi.Pack("pairImpl")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPairImpl is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x35f7d456.
//
// Solidity: function pairImpl() view returns(address)
func (marketImplV2 *MarketImplV2) UnpackPairImpl(data []byte) (common.Address, error) {
	out, err := marketImplV2.abi.Unpack("pairImpl", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackProxiableUUID is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (marketImplV2 *MarketImplV2) PackProxiableUUID() []byte {
	enc, err := marketImplV2.abi.Pack("proxiableUUID")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackProxiableUUID is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (marketImplV2 *MarketImplV2) UnpackProxiableUUID(data []byte) ([32]byte, error) {
	out, err := marketImplV2.abi.Unpack("proxiableUUID", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackRenounceOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (marketImplV2 *MarketImplV2) PackRenounceOwnership() []byte {
	enc, err := marketImplV2.abi.Pack("renounceOwnership")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetFeeCollector is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa42dce80.
//
// Solidity: function setFeeCollector(address _feeCollector) returns()
func (marketImplV2 *MarketImplV2) PackSetFeeCollector(feeCollector common.Address) []byte {
	enc, err := marketImplV2.abi.Pack("setFeeCollector", feeCollector)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetMarketFees is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x584f6af3.
//
// Solidity: function setMarketFees(uint32 _sellerMakerFeeBps, uint32 _sellerTakerFeeBps, uint32 _buyerMakerFeeBps, uint32 _buyerTakerFeeBps) returns()
func (marketImplV2 *MarketImplV2) PackSetMarketFees(sellerMakerFeeBps uint32, sellerTakerFeeBps uint32, buyerMakerFeeBps uint32, buyerTakerFeeBps uint32) []byte {
	enc, err := marketImplV2.abi.Pack("setMarketFees", sellerMakerFeeBps, sellerTakerFeeBps, buyerMakerFeeBps, buyerTakerFeeBps)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetPairImpl is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3cc047aa.
//
// Solidity: function setPairImpl(address _pairImpl) returns()
func (marketImplV2 *MarketImplV2) PackSetPairImpl(pairImpl common.Address) []byte {
	enc, err := marketImplV2.abi.Pack("setPairImpl", pairImpl)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackTransferOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (marketImplV2 *MarketImplV2) PackTransferOwnership(newOwner common.Address) []byte {
	enc, err := marketImplV2.abi.Pack("transferOwnership", newOwner)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackUpgradeToAndCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (marketImplV2 *MarketImplV2) PackUpgradeToAndCall(newImplementation common.Address, data []byte) []byte {
	enc, err := marketImplV2.abi.Pack("upgradeToAndCall", newImplementation, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// MarketImplV2FeeCollectorChanged represents a FeeCollectorChanged event raised by the MarketImplV2 contract.
type MarketImplV2FeeCollectorChanged struct {
	Before  common.Address
	Current common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const MarketImplV2FeeCollectorChangedEventName = "FeeCollectorChanged"

// ContractEventName returns the user-defined event name.
func (MarketImplV2FeeCollectorChanged) ContractEventName() string {
	return MarketImplV2FeeCollectorChangedEventName
}

// UnpackFeeCollectorChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event FeeCollectorChanged(address indexed before, address indexed current)
func (marketImplV2 *MarketImplV2) UnpackFeeCollectorChangedEvent(log *types.Log) (*MarketImplV2FeeCollectorChanged, error) {
	event := "FeeCollectorChanged"
	if log.Topics[0] != marketImplV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MarketImplV2FeeCollectorChanged)
	if len(log.Data) > 0 {
		if err := marketImplV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range marketImplV2.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// MarketImplV2Initialized represents a Initialized event raised by the MarketImplV2 contract.
type MarketImplV2Initialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const MarketImplV2InitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (MarketImplV2Initialized) ContractEventName() string {
	return MarketImplV2InitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (marketImplV2 *MarketImplV2) UnpackInitializedEvent(log *types.Log) (*MarketImplV2Initialized, error) {
	event := "Initialized"
	if log.Topics[0] != marketImplV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MarketImplV2Initialized)
	if len(log.Data) > 0 {
		if err := marketImplV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range marketImplV2.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// MarketImplV2MarketFeesUpdated represents a MarketFeesUpdated event raised by the MarketImplV2 contract.
type MarketImplV2MarketFeesUpdated struct {
	SellerMakerFee uint32
	SellerTakerFee uint32
	BuyerMakerFee  uint32
	BuyerTakerFee  uint32
	Raw            *types.Log // Blockchain specific contextual infos
}

const MarketImplV2MarketFeesUpdatedEventName = "MarketFeesUpdated"

// ContractEventName returns the user-defined event name.
func (MarketImplV2MarketFeesUpdated) ContractEventName() string {
	return MarketImplV2MarketFeesUpdatedEventName
}

// UnpackMarketFeesUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MarketFeesUpdated(uint32 sellerMakerFee, uint32 sellerTakerFee, uint32 buyerMakerFee, uint32 buyerTakerFee)
func (marketImplV2 *MarketImplV2) UnpackMarketFeesUpdatedEvent(log *types.Log) (*MarketImplV2MarketFeesUpdated, error) {
	event := "MarketFeesUpdated"
	if log.Topics[0] != marketImplV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MarketImplV2MarketFeesUpdated)
	if len(log.Data) > 0 {
		if err := marketImplV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range marketImplV2.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// MarketImplV2OwnershipTransferred represents a OwnershipTransferred event raised by the MarketImplV2 contract.
type MarketImplV2OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const MarketImplV2OwnershipTransferredEventName = "OwnershipTransferred"

// ContractEventName returns the user-defined event name.
func (MarketImplV2OwnershipTransferred) ContractEventName() string {
	return MarketImplV2OwnershipTransferredEventName
}

// UnpackOwnershipTransferredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (marketImplV2 *MarketImplV2) UnpackOwnershipTransferredEvent(log *types.Log) (*MarketImplV2OwnershipTransferred, error) {
	event := "OwnershipTransferred"
	if log.Topics[0] != marketImplV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MarketImplV2OwnershipTransferred)
	if len(log.Data) > 0 {
		if err := marketImplV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range marketImplV2.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// MarketImplV2PairCreated represents a PairCreated event raised by the MarketImplV2 contract.
type MarketImplV2PairCreated struct {
	Pair      common.Address
	Base      common.Address
	Timestamp *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const MarketImplV2PairCreatedEventName = "PairCreated"

// ContractEventName returns the user-defined event name.
func (MarketImplV2PairCreated) ContractEventName() string {
	return MarketImplV2PairCreatedEventName
}

// UnpackPairCreatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PairCreated(address indexed pair, address indexed base, uint256 timestamp)
func (marketImplV2 *MarketImplV2) UnpackPairCreatedEvent(log *types.Log) (*MarketImplV2PairCreated, error) {
	event := "PairCreated"
	if log.Topics[0] != marketImplV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MarketImplV2PairCreated)
	if len(log.Data) > 0 {
		if err := marketImplV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range marketImplV2.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// MarketImplV2PairImplSet represents a PairImplSet event raised by the MarketImplV2 contract.
type MarketImplV2PairImplSet struct {
	Before  common.Address
	Current common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const MarketImplV2PairImplSetEventName = "PairImplSet"

// ContractEventName returns the user-defined event name.
func (MarketImplV2PairImplSet) ContractEventName() string {
	return MarketImplV2PairImplSetEventName
}

// UnpackPairImplSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PairImplSet(address indexed before, address indexed current)
func (marketImplV2 *MarketImplV2) UnpackPairImplSetEvent(log *types.Log) (*MarketImplV2PairImplSet, error) {
	event := "PairImplSet"
	if log.Topics[0] != marketImplV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MarketImplV2PairImplSet)
	if len(log.Data) > 0 {
		if err := marketImplV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range marketImplV2.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// MarketImplV2Upgraded represents a Upgraded event raised by the MarketImplV2 contract.
type MarketImplV2Upgraded struct {
	Implementation common.Address
	Raw            *types.Log // Blockchain specific contextual infos
}

const MarketImplV2UpgradedEventName = "Upgraded"

// ContractEventName returns the user-defined event name.
func (MarketImplV2Upgraded) ContractEventName() string {
	return MarketImplV2UpgradedEventName
}

// UnpackUpgradedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Upgraded(address indexed implementation)
func (marketImplV2 *MarketImplV2) UnpackUpgradedEvent(log *types.Log) (*MarketImplV2Upgraded, error) {
	event := "Upgraded"
	if log.Topics[0] != marketImplV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MarketImplV2Upgraded)
	if len(log.Data) > 0 {
		if err := marketImplV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range marketImplV2.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// UnpackError attempts to decode the provided error data using user-defined
// error definitions.
func (marketImplV2 *MarketImplV2) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], marketImplV2.abi.Errors["AddressEmptyCode"].ID.Bytes()[:4]) {
		return marketImplV2.UnpackAddressEmptyCodeError(raw[4:])
	}
	if bytes.Equal(raw[:4], marketImplV2.abi.Errors["Create2EmptyBytecode"].ID.Bytes()[:4]) {
		return marketImplV2.UnpackCreate2EmptyBytecodeError(raw[4:])
	}
	if bytes.Equal(raw[:4], marketImplV2.abi.Errors["ERC1967InvalidImplementation"].ID.Bytes()[:4]) {
		return marketImplV2.UnpackERC1967InvalidImplementationError(raw[4:])
	}
	if bytes.Equal(raw[:4], marketImplV2.abi.Errors["ERC1967NonPayable"].ID.Bytes()[:4]) {
		return marketImplV2.UnpackERC1967NonPayableError(raw[4:])
	}
	if bytes.Equal(raw[:4], marketImplV2.abi.Errors["EnumerableMapNonexistentKey"].ID.Bytes()[:4]) {
		return marketImplV2.UnpackEnumerableMapNonexistentKeyError(raw[4:])
	}
	if bytes.Equal(raw[:4], marketImplV2.abi.Errors["FailedCall"].ID.Bytes()[:4]) {
		return marketImplV2.UnpackFailedCallError(raw[4:])
	}
	if bytes.Equal(raw[:4], marketImplV2.abi.Errors["FailedDeployment"].ID.Bytes()[:4]) {
		return marketImplV2.UnpackFailedDeploymentError(raw[4:])
	}
	if bytes.Equal(raw[:4], marketImplV2.abi.Errors["InsufficientBalance"].ID.Bytes()[:4]) {
		return marketImplV2.UnpackInsufficientBalanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], marketImplV2.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return marketImplV2.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], marketImplV2.abi.Errors["MarketAlreadyCreatedBaseAddress"].ID.Bytes()[:4]) {
		return marketImplV2.UnpackMarketAlreadyCreatedBaseAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], marketImplV2.abi.Errors["MarketDeployPair"].ID.Bytes()[:4]) {
		return marketImplV2.UnpackMarketDeployPairError(raw[4:])
	}
	if bytes.Equal(raw[:4], marketImplV2.abi.Errors["MarketInvalidBaseAddress"].ID.Bytes()[:4]) {
		return marketImplV2.UnpackMarketInvalidBaseAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], marketImplV2.abi.Errors["MarketInvalidFeeStructure"].ID.Bytes()[:4]) {
		return marketImplV2.UnpackMarketInvalidFeeStructureError(raw[4:])
	}
	if bytes.Equal(raw[:4], marketImplV2.abi.Errors["MarketInvalidInitializeData"].ID.Bytes()[:4]) {
		return marketImplV2.UnpackMarketInvalidInitializeDataError(raw[4:])
	}
	if bytes.Equal(raw[:4], marketImplV2.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return marketImplV2.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], marketImplV2.abi.Errors["OwnableInvalidOwner"].ID.Bytes()[:4]) {
		return marketImplV2.UnpackOwnableInvalidOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], marketImplV2.abi.Errors["OwnableUnauthorizedAccount"].ID.Bytes()[:4]) {
		return marketImplV2.UnpackOwnableUnauthorizedAccountError(raw[4:])
	}
	if bytes.Equal(raw[:4], marketImplV2.abi.Errors["UUPSUnauthorizedCallContext"].ID.Bytes()[:4]) {
		return marketImplV2.UnpackUUPSUnauthorizedCallContextError(raw[4:])
	}
	if bytes.Equal(raw[:4], marketImplV2.abi.Errors["UUPSUnsupportedProxiableUUID"].ID.Bytes()[:4]) {
		return marketImplV2.UnpackUUPSUnsupportedProxiableUUIDError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// MarketImplV2AddressEmptyCode represents a AddressEmptyCode error raised by the MarketImplV2 contract.
type MarketImplV2AddressEmptyCode struct {
	Target common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressEmptyCode(address target)
func MarketImplV2AddressEmptyCodeErrorID() common.Hash {
	return common.HexToHash("0x9996b315c842ff135b8fc4a08ad5df1c344efbc03d2687aecc0678050d2aac89")
}

// UnpackAddressEmptyCodeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressEmptyCode(address target)
func (marketImplV2 *MarketImplV2) UnpackAddressEmptyCodeError(raw []byte) (*MarketImplV2AddressEmptyCode, error) {
	out := new(MarketImplV2AddressEmptyCode)
	if err := marketImplV2.abi.UnpackIntoInterface(out, "AddressEmptyCode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MarketImplV2Create2EmptyBytecode represents a Create2EmptyBytecode error raised by the MarketImplV2 contract.
type MarketImplV2Create2EmptyBytecode struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error Create2EmptyBytecode()
func MarketImplV2Create2EmptyBytecodeErrorID() common.Hash {
	return common.HexToHash("0x4ca249dcffe41558ef8b961d71c905e4fa4317a1663f377b9610642e4e0abdb6")
}

// UnpackCreate2EmptyBytecodeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error Create2EmptyBytecode()
func (marketImplV2 *MarketImplV2) UnpackCreate2EmptyBytecodeError(raw []byte) (*MarketImplV2Create2EmptyBytecode, error) {
	out := new(MarketImplV2Create2EmptyBytecode)
	if err := marketImplV2.abi.UnpackIntoInterface(out, "Create2EmptyBytecode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MarketImplV2ERC1967InvalidImplementation represents a ERC1967InvalidImplementation error raised by the MarketImplV2 contract.
type MarketImplV2ERC1967InvalidImplementation struct {
	Implementation common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func MarketImplV2ERC1967InvalidImplementationErrorID() common.Hash {
	return common.HexToHash("0x4c9c8ce3ceb3130f17f7cdba48d89b5b0129f266a8bac114e6e315a41879b617")
}

// UnpackERC1967InvalidImplementationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func (marketImplV2 *MarketImplV2) UnpackERC1967InvalidImplementationError(raw []byte) (*MarketImplV2ERC1967InvalidImplementation, error) {
	out := new(MarketImplV2ERC1967InvalidImplementation)
	if err := marketImplV2.abi.UnpackIntoInterface(out, "ERC1967InvalidImplementation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MarketImplV2ERC1967NonPayable represents a ERC1967NonPayable error raised by the MarketImplV2 contract.
type MarketImplV2ERC1967NonPayable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967NonPayable()
func MarketImplV2ERC1967NonPayableErrorID() common.Hash {
	return common.HexToHash("0xb398979fa84f543c8e222f17890372c487baf85e062276c127fef521eea7224b")
}

// UnpackERC1967NonPayableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967NonPayable()
func (marketImplV2 *MarketImplV2) UnpackERC1967NonPayableError(raw []byte) (*MarketImplV2ERC1967NonPayable, error) {
	out := new(MarketImplV2ERC1967NonPayable)
	if err := marketImplV2.abi.UnpackIntoInterface(out, "ERC1967NonPayable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MarketImplV2EnumerableMapNonexistentKey represents a EnumerableMapNonexistentKey error raised by the MarketImplV2 contract.
type MarketImplV2EnumerableMapNonexistentKey struct {
	Key [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EnumerableMapNonexistentKey(bytes32 key)
func MarketImplV2EnumerableMapNonexistentKeyErrorID() common.Hash {
	return common.HexToHash("0x02b566865b1da2a2deb61443712fe7f812ffd7a1fce56446ff0fe3db9f3484ef")
}

// UnpackEnumerableMapNonexistentKeyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EnumerableMapNonexistentKey(bytes32 key)
func (marketImplV2 *MarketImplV2) UnpackEnumerableMapNonexistentKeyError(raw []byte) (*MarketImplV2EnumerableMapNonexistentKey, error) {
	out := new(MarketImplV2EnumerableMapNonexistentKey)
	if err := marketImplV2.abi.UnpackIntoInterface(out, "EnumerableMapNonexistentKey", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MarketImplV2FailedCall represents a FailedCall error raised by the MarketImplV2 contract.
type MarketImplV2FailedCall struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedCall()
func MarketImplV2FailedCallErrorID() common.Hash {
	return common.HexToHash("0xd6bda27508c0fb6d8a39b4b122878dab26f731a7d4e4abe711dd3731899052a4")
}

// UnpackFailedCallError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedCall()
func (marketImplV2 *MarketImplV2) UnpackFailedCallError(raw []byte) (*MarketImplV2FailedCall, error) {
	out := new(MarketImplV2FailedCall)
	if err := marketImplV2.abi.UnpackIntoInterface(out, "FailedCall", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MarketImplV2FailedDeployment represents a FailedDeployment error raised by the MarketImplV2 contract.
type MarketImplV2FailedDeployment struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedDeployment()
func MarketImplV2FailedDeploymentErrorID() common.Hash {
	return common.HexToHash("0xb06ebf3d5067824a3fe5d5ba19471e035a7de6c88dac362c77b162830a5b9093")
}

// UnpackFailedDeploymentError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedDeployment()
func (marketImplV2 *MarketImplV2) UnpackFailedDeploymentError(raw []byte) (*MarketImplV2FailedDeployment, error) {
	out := new(MarketImplV2FailedDeployment)
	if err := marketImplV2.abi.UnpackIntoInterface(out, "FailedDeployment", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MarketImplV2InsufficientBalance represents a InsufficientBalance error raised by the MarketImplV2 contract.
type MarketImplV2InsufficientBalance struct {
	Balance *big.Int
	Needed  *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InsufficientBalance(uint256 balance, uint256 needed)
func MarketImplV2InsufficientBalanceErrorID() common.Hash {
	return common.HexToHash("0xcf4791818fba6e019216eb4864093b4947f674afada5d305e57d598b641dad1d")
}

// UnpackInsufficientBalanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InsufficientBalance(uint256 balance, uint256 needed)
func (marketImplV2 *MarketImplV2) UnpackInsufficientBalanceError(raw []byte) (*MarketImplV2InsufficientBalance, error) {
	out := new(MarketImplV2InsufficientBalance)
	if err := marketImplV2.abi.UnpackIntoInterface(out, "InsufficientBalance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MarketImplV2InvalidInitialization represents a InvalidInitialization error raised by the MarketImplV2 contract.
type MarketImplV2InvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func MarketImplV2InvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (marketImplV2 *MarketImplV2) UnpackInvalidInitializationError(raw []byte) (*MarketImplV2InvalidInitialization, error) {
	out := new(MarketImplV2InvalidInitialization)
	if err := marketImplV2.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MarketImplV2MarketAlreadyCreatedBaseAddress represents a MarketAlreadyCreatedBaseAddress error raised by the MarketImplV2 contract.
type MarketImplV2MarketAlreadyCreatedBaseAddress struct {
	Arg0 common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error MarketAlreadyCreatedBaseAddress(address arg0)
func MarketImplV2MarketAlreadyCreatedBaseAddressErrorID() common.Hash {
	return common.HexToHash("0xc13fcfdc9c07ead596b870feafa1c7da90f68b06105d707eafb3632707684ca2")
}

// UnpackMarketAlreadyCreatedBaseAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error MarketAlreadyCreatedBaseAddress(address arg0)
func (marketImplV2 *MarketImplV2) UnpackMarketAlreadyCreatedBaseAddressError(raw []byte) (*MarketImplV2MarketAlreadyCreatedBaseAddress, error) {
	out := new(MarketImplV2MarketAlreadyCreatedBaseAddress)
	if err := marketImplV2.abi.UnpackIntoInterface(out, "MarketAlreadyCreatedBaseAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MarketImplV2MarketDeployPair represents a MarketDeployPair error raised by the MarketImplV2 contract.
type MarketImplV2MarketDeployPair struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error MarketDeployPair()
func MarketImplV2MarketDeployPairErrorID() common.Hash {
	return common.HexToHash("0x3ff6a261f592a9c9e4b2201a0b6349edfb07458d95b8d396f4b58be24bcdc6e2")
}

// UnpackMarketDeployPairError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error MarketDeployPair()
func (marketImplV2 *MarketImplV2) UnpackMarketDeployPairError(raw []byte) (*MarketImplV2MarketDeployPair, error) {
	out := new(MarketImplV2MarketDeployPair)
	if err := marketImplV2.abi.UnpackIntoInterface(out, "MarketDeployPair", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MarketImplV2MarketInvalidBaseAddress represents a MarketInvalidBaseAddress error raised by the MarketImplV2 contract.
type MarketImplV2MarketInvalidBaseAddress struct {
	Arg0 common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error MarketInvalidBaseAddress(address arg0)
func MarketImplV2MarketInvalidBaseAddressErrorID() common.Hash {
	return common.HexToHash("0x20f8254c2e2f68f995cc9b2605eef3477e0f8b491020c7f5bac08276c16a6c3c")
}

// UnpackMarketInvalidBaseAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error MarketInvalidBaseAddress(address arg0)
func (marketImplV2 *MarketImplV2) UnpackMarketInvalidBaseAddressError(raw []byte) (*MarketImplV2MarketInvalidBaseAddress, error) {
	out := new(MarketImplV2MarketInvalidBaseAddress)
	if err := marketImplV2.abi.UnpackIntoInterface(out, "MarketInvalidBaseAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MarketImplV2MarketInvalidFeeStructure represents a MarketInvalidFeeStructure error raised by the MarketImplV2 contract.
type MarketImplV2MarketInvalidFeeStructure struct {
	MakerFee uint32
	TakerFee uint32
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error MarketInvalidFeeStructure(uint32 makerFee, uint32 takerFee)
func MarketImplV2MarketInvalidFeeStructureErrorID() common.Hash {
	return common.HexToHash("0x38642d5f191199c26b5ac197ab5b7159d2887ffd8bf493c688d4c1598aa93e29")
}

// UnpackMarketInvalidFeeStructureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error MarketInvalidFeeStructure(uint32 makerFee, uint32 takerFee)
func (marketImplV2 *MarketImplV2) UnpackMarketInvalidFeeStructureError(raw []byte) (*MarketImplV2MarketInvalidFeeStructure, error) {
	out := new(MarketImplV2MarketInvalidFeeStructure)
	if err := marketImplV2.abi.UnpackIntoInterface(out, "MarketInvalidFeeStructure", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MarketImplV2MarketInvalidInitializeData represents a MarketInvalidInitializeData error raised by the MarketImplV2 contract.
type MarketImplV2MarketInvalidInitializeData struct {
	Arg0 [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error MarketInvalidInitializeData(bytes32 arg0)
func MarketImplV2MarketInvalidInitializeDataErrorID() common.Hash {
	return common.HexToHash("0xa25d2db5f604033642dd6fbeddebf2802c569afa70e0e2031384d873de82c9c5")
}

// UnpackMarketInvalidInitializeDataError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error MarketInvalidInitializeData(bytes32 arg0)
func (marketImplV2 *MarketImplV2) UnpackMarketInvalidInitializeDataError(raw []byte) (*MarketImplV2MarketInvalidInitializeData, error) {
	out := new(MarketImplV2MarketInvalidInitializeData)
	if err := marketImplV2.abi.UnpackIntoInterface(out, "MarketInvalidInitializeData", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MarketImplV2NotInitializing represents a NotInitializing error raised by the MarketImplV2 contract.
type MarketImplV2NotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func MarketImplV2NotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (marketImplV2 *MarketImplV2) UnpackNotInitializingError(raw []byte) (*MarketImplV2NotInitializing, error) {
	out := new(MarketImplV2NotInitializing)
	if err := marketImplV2.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MarketImplV2OwnableInvalidOwner represents a OwnableInvalidOwner error raised by the MarketImplV2 contract.
type MarketImplV2OwnableInvalidOwner struct {
	Owner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OwnableInvalidOwner(address owner)
func MarketImplV2OwnableInvalidOwnerErrorID() common.Hash {
	return common.HexToHash("0x1e4fbdf7f3ef8bcaa855599e3abf48b232380f183f08f6f813d9ffa5bd585188")
}

// UnpackOwnableInvalidOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OwnableInvalidOwner(address owner)
func (marketImplV2 *MarketImplV2) UnpackOwnableInvalidOwnerError(raw []byte) (*MarketImplV2OwnableInvalidOwner, error) {
	out := new(MarketImplV2OwnableInvalidOwner)
	if err := marketImplV2.abi.UnpackIntoInterface(out, "OwnableInvalidOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MarketImplV2OwnableUnauthorizedAccount represents a OwnableUnauthorizedAccount error raised by the MarketImplV2 contract.
type MarketImplV2OwnableUnauthorizedAccount struct {
	Account common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OwnableUnauthorizedAccount(address account)
func MarketImplV2OwnableUnauthorizedAccountErrorID() common.Hash {
	return common.HexToHash("0x118cdaa7a341953d1887a2245fd6665d741c67c8c50581daa59e1d03373fa188")
}

// UnpackOwnableUnauthorizedAccountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OwnableUnauthorizedAccount(address account)
func (marketImplV2 *MarketImplV2) UnpackOwnableUnauthorizedAccountError(raw []byte) (*MarketImplV2OwnableUnauthorizedAccount, error) {
	out := new(MarketImplV2OwnableUnauthorizedAccount)
	if err := marketImplV2.abi.UnpackIntoInterface(out, "OwnableUnauthorizedAccount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MarketImplV2UUPSUnauthorizedCallContext represents a UUPSUnauthorizedCallContext error raised by the MarketImplV2 contract.
type MarketImplV2UUPSUnauthorizedCallContext struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnauthorizedCallContext()
func MarketImplV2UUPSUnauthorizedCallContextErrorID() common.Hash {
	return common.HexToHash("0xe07c8dba242a06571ac65fe4bbe20522c9fb111cb33599b799ff8039c1ed18f4")
}

// UnpackUUPSUnauthorizedCallContextError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnauthorizedCallContext()
func (marketImplV2 *MarketImplV2) UnpackUUPSUnauthorizedCallContextError(raw []byte) (*MarketImplV2UUPSUnauthorizedCallContext, error) {
	out := new(MarketImplV2UUPSUnauthorizedCallContext)
	if err := marketImplV2.abi.UnpackIntoInterface(out, "UUPSUnauthorizedCallContext", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MarketImplV2UUPSUnsupportedProxiableUUID represents a UUPSUnsupportedProxiableUUID error raised by the MarketImplV2 contract.
type MarketImplV2UUPSUnsupportedProxiableUUID struct {
	Slot [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func MarketImplV2UUPSUnsupportedProxiableUUIDErrorID() common.Hash {
	return common.HexToHash("0xaa1d49a4c084bfa9aeeee2a0be65267a7f19ba7e1476b114dac513d2c14cb563")
}

// UnpackUUPSUnsupportedProxiableUUIDError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func (marketImplV2 *MarketImplV2) UnpackUUPSUnsupportedProxiableUUIDError(raw []byte) (*MarketImplV2UUPSUnsupportedProxiableUUID, error) {
	out := new(MarketImplV2UUPSUnsupportedProxiableUUID)
	if err := marketImplV2.abi.UnpackIntoInterface(out, "UUPSUnsupportedProxiableUUID", raw); err != nil {
		return nil, err
	}
	return out, nil
}
