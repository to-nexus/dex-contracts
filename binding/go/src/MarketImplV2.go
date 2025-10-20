// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package binding

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
	_ = abi.ConvertType
)

// MarketImplV2MetaData contains all meta data concerning the MarketImplV2 contract.
var MarketImplV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CROSS_DEX\",\"outputs\":[{\"internalType\":\"contractICrossDex\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"QUOTE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROUTER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allPairs\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"bases\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"pairs\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"}],\"name\":\"baseToPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"checkTickSizeRoles\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tickSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lotSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"name\":\"createPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deployed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeCollector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"sellerMakerFeeBps\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"sellerTakerFeeBps\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"buyerMakerFeeBps\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"buyerTakerFeeBps\",\"type\":\"uint32\"}],\"internalType\":\"structIMarketV2.FeeConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_quote\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pairImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeCollector\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pairImpl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeCollector\",\"type\":\"address\"}],\"name\":\"setFeeCollector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_sellerMakerFeeBps\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_sellerTakerFeeBps\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_buyerMakerFeeBps\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_buyerTakerFeeBps\",\"type\":\"uint32\"}],\"name\":\"setMarketFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pairImpl\",\"type\":\"address\"}],\"name\":\"setPairImpl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"before\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"FeeCollectorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"sellerMakerFee\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"sellerTakerFee\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"buyerMakerFee\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"buyerTakerFee\",\"type\":\"uint32\"}],\"name\":\"MarketFeesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"PairCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"before\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"PairImplSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Create2EmptyBytecode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"EnumerableMapNonexistentKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedDeployment\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"MarketAlreadyCreatedBaseAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MarketDeployPair\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"MarketInvalidBaseAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"makerFee\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"takerFee\",\"type\":\"uint32\"}],\"name\":\"MarketInvalidFeeStructure\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"MarketInvalidInitializeData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"b820d8aa": "CROSS_DEX()",
		"9c579839": "QUOTE()",
		"32fe7b26": "ROUTER()",
		"ad3cb1cc": "UPGRADE_INTERFACE_VERSION()",
		"c97682f8": "allPairs()",
		"364e9a19": "baseToPair(address)",
		"a1eea778": "checkTickSizeRoles(address)",
		"4732c7e5": "createPair(address,uint256,uint256,bytes)",
		"f905c15a": "deployed()",
		"c415b95c": "feeCollector()",
		"5fbbc0d2": "getFeeConfig()",
		"01f79fb6": "initialize(address,address,address,address,address,bytes)",
		"8da5cb5b": "owner()",
		"35f7d456": "pairImpl()",
		"52d1902d": "proxiableUUID()",
		"715018a6": "renounceOwnership()",
		"a42dce80": "setFeeCollector(address)",
		"584f6af3": "setMarketFees(uint32,uint32,uint32,uint32)",
		"3cc047aa": "setPairImpl(address)",
		"f2fde38b": "transferOwnership(address)",
		"4f1ef286": "upgradeToAndCall(address,bytes)",
	},
	Bin: "0x60a060405230608052348015610013575f5ffd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b608051612d7a6100f95f395f8181611c3c01528181611c650152611e860152612d7a5ff3fe608060405260043610610162575f3560e01c8063715018a6116100c6578063ad3cb1cc1161007c578063c97682f811610057578063c97682f814610506578063f2fde38b14610528578063f905c15a14610547575f5ffd5b8063ad3cb1cc14610459578063b820d8aa146104ae578063c415b95c146104da575f5ffd5b80639c579839116100ac5780639c579839146103ef578063a1eea7781461041b578063a42dce801461043a575f5ffd5b8063715018a6146103925780638da5cb5b146103a6575f5ffd5b80634732c7e51161011b57806352d1902d1161010157806352d1902d14610279578063584f6af31461029b5780635fbbc0d2146102ba575f5ffd5b80634732c7e5146102475780634f1ef28614610266575f5ffd5b806335f7d4561161014b57806335f7d456146101dd578063364e9a19146102095780633cc047aa14610228575f5ffd5b806301f79fb61461016657806332fe7b2614610187575b5f5ffd5b348015610171575f5ffd5b50610185610180366004612549565b61055b565b005b348015610192575f5ffd5b506003546101b39073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b3480156101e8575f5ffd5b506004546101b39073ffffffffffffffffffffffffffffffffffffffff1681565b348015610214575f5ffd5b506101b36102233660046125d3565b610ce2565b348015610233575f5ffd5b506101856102423660046125d3565b610cf4565b348015610252575f5ffd5b506101b36102613660046125ec565b610df8565b610185610274366004612649565b611311565b348015610284575f5ffd5b5061028d611330565b6040519081526020016101d4565b3480156102a6575f5ffd5b506101856102b53660046126a5565b61135e565b3480156102c5575f5ffd5b50610343604080516080810182525f808252602082018190529181018290526060810191909152506040805160808101825260095463ffffffff80821683526401000000008204811660208401526801000000000000000082048116938301939093526c010000000000000000000000009004909116606082015290565b6040516101d491905f60808201905063ffffffff835116825263ffffffff602084015116602083015263ffffffff604084015116604083015263ffffffff606084015116606083015292915050565b34801561039d575f5ffd5b50610185611718565b3480156103b1575f5ffd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c1993005473ffffffffffffffffffffffffffffffffffffffff166101b3565b3480156103fa575f5ffd5b506002546101b39073ffffffffffffffffffffffffffffffffffffffff1681565b348015610426575f5ffd5b506101856104353660046125d3565b61172b565b348015610445575f5ffd5b506101856104543660046125d3565b6117ad565b348015610464575f5ffd5b506104a16040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b6040516101d4919061274a565b3480156104b9575f5ffd5b506001546101b39073ffffffffffffffffffffffffffffffffffffffff1681565b3480156104e5575f5ffd5b506005546101b39073ffffffffffffffffffffffffffffffffffffffff1681565b348015610511575f5ffd5b5061051a6118b1565b6040516101d49291906127ac565b348015610533575f5ffd5b506101856105423660046125d3565b6119c0565b348015610552575f5ffd5b5061028d5f5481565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff165f811580156105a55750825b90505f8267ffffffffffffffff1660011480156105c15750303b155b9050811580156105cf575080155b15610606576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000016600117855583156106675784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b6106708b611a23565b73ffffffffffffffffffffffffffffffffffffffff8b166106e4576040517fa25d2db50000000000000000000000000000000000000000000000000000000081527f6f776e657200000000000000000000000000000000000000000000000000000060048201526024015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8a16610753576040517fa25d2db50000000000000000000000000000000000000000000000000000000081527f726f75746572000000000000000000000000000000000000000000000000000060048201526024016106db565b73ffffffffffffffffffffffffffffffffffffffff89166107c2576040517fa25d2db50000000000000000000000000000000000000000000000000000000081527f71756f746500000000000000000000000000000000000000000000000000000060048201526024016106db565b73ffffffffffffffffffffffffffffffffffffffff8816610831576040517fa25d2db50000000000000000000000000000000000000000000000000000000081527f70616972496d706c00000000000000000000000000000000000000000000000060048201526024016106db565b73ffffffffffffffffffffffffffffffffffffffff87166108a0576040517fa25d2db50000000000000000000000000000000000000000000000000000000081527f666565436f6c6c6563746f72000000000000000000000000000000000000000060048201526024016106db565b5f5f5f5f898060200190518101906108b891906127d0565b935093509350935063ffffffff80168463ffffffff16141580156108e4575061271063ffffffff851610155b1561093d576040517fa25d2db50000000000000000000000000000000000000000000000000000000081527f73656c6c65724d616b657246656542707300000000000000000000000000000060048201526024016106db565b63ffffffff8381161480159061095b575061271063ffffffff841610155b156109b4576040517fa25d2db50000000000000000000000000000000000000000000000000000000081527f73656c6c657254616b657246656542707300000000000000000000000000000060048201526024016106db565b63ffffffff828116148015906109d2575061271063ffffffff831610155b15610a2b576040517fa25d2db50000000000000000000000000000000000000000000000000000000081527f62757965724d616b65724665654270730000000000000000000000000000000060048201526024016106db565b63ffffffff81811614801590610a49575061271063ffffffff821610155b15610aa2576040517fa25d2db50000000000000000000000000000000000000000000000000000000081527f627579657254616b65724665654270730000000000000000000000000000000060048201526024016106db565b435f553360015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508c60025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508d60035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508b60045f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508a60055f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508360095f015f6101000a81548163ffffffff021916908363ffffffff1602179055508260095f0160046101000a81548163ffffffff021916908363ffffffff1602179055508160095f0160086101000a81548163ffffffff021916908363ffffffff1602179055508060095f01600c6101000a81548163ffffffff021916908363ffffffff160217905550505050508315610cd55784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050505050565b5f610cee600683611a34565b92915050565b610cfc611a5c565b73ffffffffffffffffffffffffffffffffffffffff8116610d6b576040517fa25d2db50000000000000000000000000000000000000000000000000000000081527f70616972496d706c00000000000000000000000000000000000000000000000060048201526024016106db565b60045460405173ffffffffffffffffffffffffffffffffffffffff8084169216907f685c330e9b04243f99ff39d8de8578c9af2bfaff120e17a1b5f1d61e004935e2905f90a3600480547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b5f610e01611a5c565b73ffffffffffffffffffffffffffffffffffffffff85161580610e3e575060025473ffffffffffffffffffffffffffffffffffffffff8681169116145b15610e8d576040517f20f8254c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff861660048201526024016106db565b5f8573ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa158015610ed7573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610efb9190612821565b60ff169050805f03610f51576040517f20f8254c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff871660048201526024016106db565b610f5c600687611aea565b15610fab576040517fc13fcfdc00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff871660048201526024016106db565b5f60405180602001610fbc9061241f565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082820381018352601f90910116604081905260045460035460025473ffffffffffffffffffffffffffffffffffffffff9283169361102e939283169291909116908c908c908c908c90602401612841565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f02a204c60000000000000000000000000000000000000000000000000000000017905290516110b3939291016128bf565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290526110ef9291602001612904565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290527fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060608a901b16602083015291505f906034016040516020818303038152906040528051906020012090505f6111755f8385611b0b565b905073ffffffffffffffffffffffffffffffffffffffff81166111c4576040517f3ff6a26100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6111d060068a83611bf7565b61121e576040517fc13fcfdc00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8a1660048201526024016106db565b6001546040517fe9f7206b00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83811660048301529091169063e9f7206b906024015f604051808303815f87803b158015611288575f5ffd5b505af115801561129a573d5f5f3e3d5ffd5b505050508873ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167fc1db9ba7c4b7ce660fe8d17bbcf07167549381df2abd694a970bd1402d86d313426040516112fd91815260200190565b60405180910390a398975050505050505050565b611319611c24565b61132282611d28565b61132c8282611d30565b5050565b5f611339611e6e565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b611366611a5c565b63ffffffff84811614801590611384575061271063ffffffff851610155b156113dd576040517fa25d2db50000000000000000000000000000000000000000000000000000000081527f73656c6c65724d616b657246656542707300000000000000000000000000000060048201526024016106db565b63ffffffff838116148015906113fb575061271063ffffffff841610155b15611454576040517fa25d2db50000000000000000000000000000000000000000000000000000000081527f73656c6c657254616b657246656542707300000000000000000000000000000060048201526024016106db565b63ffffffff82811614801590611472575061271063ffffffff831610155b156114cb576040517fa25d2db50000000000000000000000000000000000000000000000000000000081527f62757965724d616b65724665654270730000000000000000000000000000000060048201526024016106db565b63ffffffff818116148015906114e9575061271063ffffffff821610155b15611542576040517fa25d2db50000000000000000000000000000000000000000000000000000000081527f627579657254616b65724665654270730000000000000000000000000000000060048201526024016106db565b8363ffffffff168363ffffffff16108015611563575063ffffffff84811614155b156115aa576040517f38642d5f00000000000000000000000000000000000000000000000000000000815263ffffffff8086166004830152841660248201526044016106db565b8163ffffffff168163ffffffff161080156115cb575063ffffffff82811614155b15611612576040517f38642d5f00000000000000000000000000000000000000000000000000000000815263ffffffff8084166004830152821660248201526044016106db565b6040805163ffffffff86811682528581166020830152848116828401528316606082015290517f68657e5573176ce9149cd289efa48a58a9db7703c3b1bef2e30e41647f5ab2089181900360800190a16009805463ffffffff9586167fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000009091161764010000000094861694909402939093177fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff1668010000000000000000928516929092027fffffffffffffffffffffffffffffffff00000000ffffffffffffffffffffffff16919091176c010000000000000000000000009190931602919091179055565b611720611a5c565b6117295f611edd565b565b6001546040517fa1eea77800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83811660048301529091169063a1eea778906024015f6040518083038186803b158015611794575f5ffd5b505afa1580156117a6573d5f5f3e3d5ffd5b5050505050565b6117b5611a5c565b73ffffffffffffffffffffffffffffffffffffffff8116611824576040517fa25d2db50000000000000000000000000000000000000000000000000000000081527f666565436f6c6c6563746f72000000000000000000000000000000000000000060048201526024016106db565b60055460405173ffffffffffffffffffffffffffffffffffffffff8084169216907f649c5e3d0ed183894196148e193af316452b0037e77d2ff0fef23b7dc722bed0905f90a3600580547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b6060805f6118bf6006611f72565b90508067ffffffffffffffff8111156118da576118da612454565b604051908082528060200260200182016040528015611903578160200160208202803683370190505b5092508067ffffffffffffffff81111561191f5761191f612454565b604051908082528060200260200182016040528015611948578160200160208202803683370190505b5091505f5b818110156119ba57611960600682611f7c565b85838151811061197257611972612918565b6020026020010185848151811061198b5761198b612918565b73ffffffffffffffffffffffffffffffffffffffff93841660209182029290920101529116905260010161194d565b50509091565b6119c8611a5c565b73ffffffffffffffffffffffffffffffffffffffff8116611a17576040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081525f60048201526024016106db565b611a2081611edd565b50565b611a2b611f97565b611a2081611ffe565b5f611a558373ffffffffffffffffffffffffffffffffffffffff8416612006565b9392505050565b33611a9b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c1993005473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614611729576040517f118cdaa70000000000000000000000000000000000000000000000000000000081523360048201526024016106db565b5f611a558373ffffffffffffffffffffffffffffffffffffffff8416612065565b5f83471015611b4f576040517fcf479181000000000000000000000000000000000000000000000000000000008152476004820152602481018590526044016106db565b81515f03611b89576040517f4ca249dc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8282516020840186f590503d151981151615611baa576040513d5f823e3d81fd5b73ffffffffffffffffffffffffffffffffffffffff8116611a55576040517fb06ebf3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f611c1c8473ffffffffffffffffffffffffffffffffffffffff808616908516612070565b949350505050565b3073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161480611cf157507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16611cd87f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614155b15611729576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611a20611a5c565b8173ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015611db5575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201909252611db291810190612945565b60015b611e03576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff831660048201526024016106db565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8114611e5f576040517faa1d49a4000000000000000000000000000000000000000000000000000000008152600481018290526024016106db565b611e69838361208c565b505050565b3073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614611729576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080547fffffffffffffffffffffffff0000000000000000000000000000000000000000811673ffffffffffffffffffffffffffffffffffffffff848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b5f610cee826120ee565b5f808080611f8a86866120f8565b9097909650945050505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff16611729576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6119c8611f97565b5f8181526002830160205260408120548015801561202b57506120298484612065565b155b15611a55576040517f02b56686000000000000000000000000000000000000000000000000000000008152600481018490526024016106db565b5f611a558383612121565b5f8281526002840160205260408120829055611c1c8484612138565b61209582612143565b60405173ffffffffffffffffffffffffffffffffffffffff8316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a28051156120e657611e698282612211565b61132c612290565b5f610cee826122c8565b5f808061210585856122d1565b5f81815260029690960160205260409095205494959350505050565b5f8181526001830160205260408120541515611a55565b5f611a5583836122dc565b8073ffffffffffffffffffffffffffffffffffffffff163b5f036121ab576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024016106db565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60605f5f8473ffffffffffffffffffffffffffffffffffffffff168460405161223a919061295c565b5f60405180830381855af49150503d805f8114612272576040519150601f19603f3d011682016040523d82523d5f602084013e612277565b606091505b5091509150612287858383612328565b95945050505050565b3415611729576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f610cee825490565b5f611a5583836123b7565b5f81815260018301602052604081205461232157508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610cee565b505f610cee565b60608261233d57612338826123dd565b611a55565b8151158015612361575073ffffffffffffffffffffffffffffffffffffffff84163b155b156123b0576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024016106db565b5080611a55565b5f825f0182815481106123cc576123cc612918565b905f5260205f200154905092915050565b8051156123ed5780518082602001fd5b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103dd8061296883390190565b803573ffffffffffffffffffffffffffffffffffffffff8116811461244f575f5ffd5b919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f82601f830112612490575f5ffd5b813567ffffffffffffffff8111156124aa576124aa612454565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff8211171561251657612516612454565b60405281815283820160200185101561252d575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f5f5f5f5f5f60c0878903121561255e575f5ffd5b6125678761242c565b95506125756020880161242c565b94506125836040880161242c565b93506125916060880161242c565b925061259f6080880161242c565b915060a087013567ffffffffffffffff8111156125ba575f5ffd5b6125c689828a01612481565b9150509295509295509295565b5f602082840312156125e3575f5ffd5b611a558261242c565b5f5f5f5f608085870312156125ff575f5ffd5b6126088561242c565b93506020850135925060408501359150606085013567ffffffffffffffff811115612631575f5ffd5b61263d87828801612481565b91505092959194509250565b5f5f6040838503121561265a575f5ffd5b6126638361242c565b9150602083013567ffffffffffffffff81111561267e575f5ffd5b61268a85828601612481565b9150509250929050565b63ffffffff81168114611a20575f5ffd5b5f5f5f5f608085870312156126b8575f5ffd5b84356126c381612694565b935060208501356126d381612694565b925060408501356126e381612694565b915060608501356126f381612694565b939692955090935050565b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f611a5560208301846126fe565b5f8151808452602084019350602083015f5b828110156127a257815173ffffffffffffffffffffffffffffffffffffffff1686526020958601959091019060010161276e565b5093949350505050565b604081525f6127be604083018561275c565b8281036020840152612287818561275c565b5f5f5f5f608085870312156127e3575f5ffd5b84516127ee81612694565b60208601519094506127ff81612694565b604086015190935061281081612694565b60608601519092506126f381612694565b5f60208284031215612831575f5ffd5b815160ff81168114611a55575f5ffd5b73ffffffffffffffffffffffffffffffffffffffff8716815273ffffffffffffffffffffffffffffffffffffffff8616602082015273ffffffffffffffffffffffffffffffffffffffff8516604082015283606082015282608082015260c060a08201525f6128b360c08301846126fe565b98975050505050505050565b73ffffffffffffffffffffffffffffffffffffffff83168152604060208201525f611c1c60408301846126fe565b5f81518060208401855e5f93019283525090919050565b5f611c1c61291283866128ed565b846128ed565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f60208284031215612955575f5ffd5b5051919050565b5f611a5582846128ed56fe60806040526040516103dd3803806103dd8339810160408190526100229161023c565b61002c8282610033565b5050610321565b61003c82610091565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a280511561008557610080828261010c565b505050565b61008d61017f565b5050565b806001600160a01b03163b5f036100cb57604051634c9c8ce360e01b81526001600160a01b03821660048201526024015b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b031684604051610128919061030b565b5f60405180830381855af49150503d805f8114610160576040519150601f19603f3d011682016040523d82523d5f602084013e610165565b606091505b5090925090506101768583836101a0565b95945050505050565b341561019e5760405163b398979f60e01b815260040160405180910390fd5b565b6060826101b5576101b0826101ff565b6101f8565b81511580156101cc57506001600160a01b0384163b155b156101f557604051639996b31560e01b81526001600160a01b03851660048201526024016100c2565b50805b9392505050565b80511561020f5780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b634e487b7160e01b5f52604160045260245ffd5b5f5f6040838503121561024d575f5ffd5b82516001600160a01b0381168114610263575f5ffd5b60208401519092506001600160401b0381111561027e575f5ffd5b8301601f8101851361028e575f5ffd5b80516001600160401b038111156102a7576102a7610228565b604051601f8201601f19908116603f011681016001600160401b03811182821017156102d5576102d5610228565b6040528181528282016020018710156102ec575f5ffd5b8160208401602083015e5f602083830101528093505050509250929050565b5f82518060208501845e5f920191825250919050565b60b08061032d5f395ff3fe6080604052600a600c565b005b60186014601a565b605d565b565b5f60587f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5473ffffffffffffffffffffffffffffffffffffffff1690565b905090565b365f5f375f5f365f845af43d5f5f3e8080156076573d5ff35b3d5ffdfea2646970667358221220f53cdd673e6d446a26f162ed3adcc3948e6ec30e5c1918fa9cfbf44a3f745dd964736f6c634300081c0033a26469706673582212209dada83115d8f996eb6a8d6aff844b308f1d2a4954841e791f071ee76f4a680664736f6c634300081c0033",
}

// MarketImplV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use MarketImplV2MetaData.ABI instead.
var MarketImplV2ABI = MarketImplV2MetaData.ABI

// Deprecated: Use MarketImplV2MetaData.Sigs instead.
// MarketImplV2FuncSigs maps the 4-byte function signature to its string representation.
var MarketImplV2FuncSigs = MarketImplV2MetaData.Sigs

// MarketImplV2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MarketImplV2MetaData.Bin instead.
var MarketImplV2Bin = MarketImplV2MetaData.Bin

// DeployMarketImplV2 deploys a new Ethereum contract, binding an instance of MarketImplV2 to it.
func DeployMarketImplV2(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MarketImplV2, error) {
	parsed, err := MarketImplV2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MarketImplV2Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MarketImplV2{MarketImplV2Caller: MarketImplV2Caller{contract: contract}, MarketImplV2Transactor: MarketImplV2Transactor{contract: contract}, MarketImplV2Filterer: MarketImplV2Filterer{contract: contract}}, nil
}

// MarketImplV2 is an auto generated Go binding around an Ethereum contract.
type MarketImplV2 struct {
	MarketImplV2Caller     // Read-only binding to the contract
	MarketImplV2Transactor // Write-only binding to the contract
	MarketImplV2Filterer   // Log filterer for contract events
}

// MarketImplV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type MarketImplV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketImplV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketImplV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketImplV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketImplV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketImplV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketImplV2Session struct {
	Contract     *MarketImplV2     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketImplV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketImplV2CallerSession struct {
	Contract *MarketImplV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// MarketImplV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketImplV2TransactorSession struct {
	Contract     *MarketImplV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MarketImplV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type MarketImplV2Raw struct {
	Contract *MarketImplV2 // Generic contract binding to access the raw methods on
}

// MarketImplV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketImplV2CallerRaw struct {
	Contract *MarketImplV2Caller // Generic read-only contract binding to access the raw methods on
}

// MarketImplV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketImplV2TransactorRaw struct {
	Contract *MarketImplV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMarketImplV2 creates a new instance of MarketImplV2, bound to a specific deployed contract.
func NewMarketImplV2(address common.Address, backend bind.ContractBackend) (*MarketImplV2, error) {
	contract, err := bindMarketImplV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MarketImplV2{MarketImplV2Caller: MarketImplV2Caller{contract: contract}, MarketImplV2Transactor: MarketImplV2Transactor{contract: contract}, MarketImplV2Filterer: MarketImplV2Filterer{contract: contract}}, nil
}

// NewMarketImplV2Caller creates a new read-only instance of MarketImplV2, bound to a specific deployed contract.
func NewMarketImplV2Caller(address common.Address, caller bind.ContractCaller) (*MarketImplV2Caller, error) {
	contract, err := bindMarketImplV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketImplV2Caller{contract: contract}, nil
}

// NewMarketImplV2Transactor creates a new write-only instance of MarketImplV2, bound to a specific deployed contract.
func NewMarketImplV2Transactor(address common.Address, transactor bind.ContractTransactor) (*MarketImplV2Transactor, error) {
	contract, err := bindMarketImplV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketImplV2Transactor{contract: contract}, nil
}

// NewMarketImplV2Filterer creates a new log filterer instance of MarketImplV2, bound to a specific deployed contract.
func NewMarketImplV2Filterer(address common.Address, filterer bind.ContractFilterer) (*MarketImplV2Filterer, error) {
	contract, err := bindMarketImplV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketImplV2Filterer{contract: contract}, nil
}

// bindMarketImplV2 binds a generic wrapper to an already deployed contract.
func bindMarketImplV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MarketImplV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MarketImplV2 *MarketImplV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MarketImplV2.Contract.MarketImplV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MarketImplV2 *MarketImplV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketImplV2.Contract.MarketImplV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MarketImplV2 *MarketImplV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MarketImplV2.Contract.MarketImplV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MarketImplV2 *MarketImplV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MarketImplV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MarketImplV2 *MarketImplV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketImplV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MarketImplV2 *MarketImplV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MarketImplV2.Contract.contract.Transact(opts, method, params...)
}

// CROSSDEX is a free data retrieval call binding the contract method 0xb820d8aa.
//
// Solidity: function CROSS_DEX() view returns(address)
func (_MarketImplV2 *MarketImplV2Caller) CROSSDEX(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarketImplV2.contract.Call(opts, &out, "CROSS_DEX")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CROSSDEX is a free data retrieval call binding the contract method 0xb820d8aa.
//
// Solidity: function CROSS_DEX() view returns(address)
func (_MarketImplV2 *MarketImplV2Session) CROSSDEX() (common.Address, error) {
	return _MarketImplV2.Contract.CROSSDEX(&_MarketImplV2.CallOpts)
}

// CROSSDEX is a free data retrieval call binding the contract method 0xb820d8aa.
//
// Solidity: function CROSS_DEX() view returns(address)
func (_MarketImplV2 *MarketImplV2CallerSession) CROSSDEX() (common.Address, error) {
	return _MarketImplV2.Contract.CROSSDEX(&_MarketImplV2.CallOpts)
}

// QUOTE is a free data retrieval call binding the contract method 0x9c579839.
//
// Solidity: function QUOTE() view returns(address)
func (_MarketImplV2 *MarketImplV2Caller) QUOTE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarketImplV2.contract.Call(opts, &out, "QUOTE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// QUOTE is a free data retrieval call binding the contract method 0x9c579839.
//
// Solidity: function QUOTE() view returns(address)
func (_MarketImplV2 *MarketImplV2Session) QUOTE() (common.Address, error) {
	return _MarketImplV2.Contract.QUOTE(&_MarketImplV2.CallOpts)
}

// QUOTE is a free data retrieval call binding the contract method 0x9c579839.
//
// Solidity: function QUOTE() view returns(address)
func (_MarketImplV2 *MarketImplV2CallerSession) QUOTE() (common.Address, error) {
	return _MarketImplV2.Contract.QUOTE(&_MarketImplV2.CallOpts)
}

// ROUTER is a free data retrieval call binding the contract method 0x32fe7b26.
//
// Solidity: function ROUTER() view returns(address)
func (_MarketImplV2 *MarketImplV2Caller) ROUTER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarketImplV2.contract.Call(opts, &out, "ROUTER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ROUTER is a free data retrieval call binding the contract method 0x32fe7b26.
//
// Solidity: function ROUTER() view returns(address)
func (_MarketImplV2 *MarketImplV2Session) ROUTER() (common.Address, error) {
	return _MarketImplV2.Contract.ROUTER(&_MarketImplV2.CallOpts)
}

// ROUTER is a free data retrieval call binding the contract method 0x32fe7b26.
//
// Solidity: function ROUTER() view returns(address)
func (_MarketImplV2 *MarketImplV2CallerSession) ROUTER() (common.Address, error) {
	return _MarketImplV2.Contract.ROUTER(&_MarketImplV2.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_MarketImplV2 *MarketImplV2Caller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MarketImplV2.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_MarketImplV2 *MarketImplV2Session) UPGRADEINTERFACEVERSION() (string, error) {
	return _MarketImplV2.Contract.UPGRADEINTERFACEVERSION(&_MarketImplV2.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_MarketImplV2 *MarketImplV2CallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _MarketImplV2.Contract.UPGRADEINTERFACEVERSION(&_MarketImplV2.CallOpts)
}

// AllPairs is a free data retrieval call binding the contract method 0xc97682f8.
//
// Solidity: function allPairs() view returns(address[] bases, address[] pairs)
func (_MarketImplV2 *MarketImplV2Caller) AllPairs(opts *bind.CallOpts) (struct {
	Bases []common.Address
	Pairs []common.Address
}, error) {
	var out []interface{}
	err := _MarketImplV2.contract.Call(opts, &out, "allPairs")

	outstruct := new(struct {
		Bases []common.Address
		Pairs []common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Bases = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Pairs = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

// AllPairs is a free data retrieval call binding the contract method 0xc97682f8.
//
// Solidity: function allPairs() view returns(address[] bases, address[] pairs)
func (_MarketImplV2 *MarketImplV2Session) AllPairs() (struct {
	Bases []common.Address
	Pairs []common.Address
}, error) {
	return _MarketImplV2.Contract.AllPairs(&_MarketImplV2.CallOpts)
}

// AllPairs is a free data retrieval call binding the contract method 0xc97682f8.
//
// Solidity: function allPairs() view returns(address[] bases, address[] pairs)
func (_MarketImplV2 *MarketImplV2CallerSession) AllPairs() (struct {
	Bases []common.Address
	Pairs []common.Address
}, error) {
	return _MarketImplV2.Contract.AllPairs(&_MarketImplV2.CallOpts)
}

// BaseToPair is a free data retrieval call binding the contract method 0x364e9a19.
//
// Solidity: function baseToPair(address base) view returns(address)
func (_MarketImplV2 *MarketImplV2Caller) BaseToPair(opts *bind.CallOpts, base common.Address) (common.Address, error) {
	var out []interface{}
	err := _MarketImplV2.contract.Call(opts, &out, "baseToPair", base)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BaseToPair is a free data retrieval call binding the contract method 0x364e9a19.
//
// Solidity: function baseToPair(address base) view returns(address)
func (_MarketImplV2 *MarketImplV2Session) BaseToPair(base common.Address) (common.Address, error) {
	return _MarketImplV2.Contract.BaseToPair(&_MarketImplV2.CallOpts, base)
}

// BaseToPair is a free data retrieval call binding the contract method 0x364e9a19.
//
// Solidity: function baseToPair(address base) view returns(address)
func (_MarketImplV2 *MarketImplV2CallerSession) BaseToPair(base common.Address) (common.Address, error) {
	return _MarketImplV2.Contract.BaseToPair(&_MarketImplV2.CallOpts, base)
}

// CheckTickSizeRoles is a free data retrieval call binding the contract method 0xa1eea778.
//
// Solidity: function checkTickSizeRoles(address account) view returns()
func (_MarketImplV2 *MarketImplV2Caller) CheckTickSizeRoles(opts *bind.CallOpts, account common.Address) error {
	var out []interface{}
	err := _MarketImplV2.contract.Call(opts, &out, "checkTickSizeRoles", account)

	if err != nil {
		return err
	}

	return err

}

// CheckTickSizeRoles is a free data retrieval call binding the contract method 0xa1eea778.
//
// Solidity: function checkTickSizeRoles(address account) view returns()
func (_MarketImplV2 *MarketImplV2Session) CheckTickSizeRoles(account common.Address) error {
	return _MarketImplV2.Contract.CheckTickSizeRoles(&_MarketImplV2.CallOpts, account)
}

// CheckTickSizeRoles is a free data retrieval call binding the contract method 0xa1eea778.
//
// Solidity: function checkTickSizeRoles(address account) view returns()
func (_MarketImplV2 *MarketImplV2CallerSession) CheckTickSizeRoles(account common.Address) error {
	return _MarketImplV2.Contract.CheckTickSizeRoles(&_MarketImplV2.CallOpts, account)
}

// Deployed is a free data retrieval call binding the contract method 0xf905c15a.
//
// Solidity: function deployed() view returns(uint256)
func (_MarketImplV2 *MarketImplV2Caller) Deployed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MarketImplV2.contract.Call(opts, &out, "deployed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Deployed is a free data retrieval call binding the contract method 0xf905c15a.
//
// Solidity: function deployed() view returns(uint256)
func (_MarketImplV2 *MarketImplV2Session) Deployed() (*big.Int, error) {
	return _MarketImplV2.Contract.Deployed(&_MarketImplV2.CallOpts)
}

// Deployed is a free data retrieval call binding the contract method 0xf905c15a.
//
// Solidity: function deployed() view returns(uint256)
func (_MarketImplV2 *MarketImplV2CallerSession) Deployed() (*big.Int, error) {
	return _MarketImplV2.Contract.Deployed(&_MarketImplV2.CallOpts)
}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_MarketImplV2 *MarketImplV2Caller) FeeCollector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarketImplV2.contract.Call(opts, &out, "feeCollector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_MarketImplV2 *MarketImplV2Session) FeeCollector() (common.Address, error) {
	return _MarketImplV2.Contract.FeeCollector(&_MarketImplV2.CallOpts)
}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_MarketImplV2 *MarketImplV2CallerSession) FeeCollector() (common.Address, error) {
	return _MarketImplV2.Contract.FeeCollector(&_MarketImplV2.CallOpts)
}

// GetFeeConfig is a free data retrieval call binding the contract method 0x5fbbc0d2.
//
// Solidity: function getFeeConfig() view returns((uint32,uint32,uint32,uint32))
func (_MarketImplV2 *MarketImplV2Caller) GetFeeConfig(opts *bind.CallOpts) (IMarketV2FeeConfig, error) {
	var out []interface{}
	err := _MarketImplV2.contract.Call(opts, &out, "getFeeConfig")

	if err != nil {
		return *new(IMarketV2FeeConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IMarketV2FeeConfig)).(*IMarketV2FeeConfig)

	return out0, err

}

// GetFeeConfig is a free data retrieval call binding the contract method 0x5fbbc0d2.
//
// Solidity: function getFeeConfig() view returns((uint32,uint32,uint32,uint32))
func (_MarketImplV2 *MarketImplV2Session) GetFeeConfig() (IMarketV2FeeConfig, error) {
	return _MarketImplV2.Contract.GetFeeConfig(&_MarketImplV2.CallOpts)
}

// GetFeeConfig is a free data retrieval call binding the contract method 0x5fbbc0d2.
//
// Solidity: function getFeeConfig() view returns((uint32,uint32,uint32,uint32))
func (_MarketImplV2 *MarketImplV2CallerSession) GetFeeConfig() (IMarketV2FeeConfig, error) {
	return _MarketImplV2.Contract.GetFeeConfig(&_MarketImplV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MarketImplV2 *MarketImplV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarketImplV2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MarketImplV2 *MarketImplV2Session) Owner() (common.Address, error) {
	return _MarketImplV2.Contract.Owner(&_MarketImplV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MarketImplV2 *MarketImplV2CallerSession) Owner() (common.Address, error) {
	return _MarketImplV2.Contract.Owner(&_MarketImplV2.CallOpts)
}

// PairImpl is a free data retrieval call binding the contract method 0x35f7d456.
//
// Solidity: function pairImpl() view returns(address)
func (_MarketImplV2 *MarketImplV2Caller) PairImpl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarketImplV2.contract.Call(opts, &out, "pairImpl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PairImpl is a free data retrieval call binding the contract method 0x35f7d456.
//
// Solidity: function pairImpl() view returns(address)
func (_MarketImplV2 *MarketImplV2Session) PairImpl() (common.Address, error) {
	return _MarketImplV2.Contract.PairImpl(&_MarketImplV2.CallOpts)
}

// PairImpl is a free data retrieval call binding the contract method 0x35f7d456.
//
// Solidity: function pairImpl() view returns(address)
func (_MarketImplV2 *MarketImplV2CallerSession) PairImpl() (common.Address, error) {
	return _MarketImplV2.Contract.PairImpl(&_MarketImplV2.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_MarketImplV2 *MarketImplV2Caller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MarketImplV2.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_MarketImplV2 *MarketImplV2Session) ProxiableUUID() ([32]byte, error) {
	return _MarketImplV2.Contract.ProxiableUUID(&_MarketImplV2.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_MarketImplV2 *MarketImplV2CallerSession) ProxiableUUID() ([32]byte, error) {
	return _MarketImplV2.Contract.ProxiableUUID(&_MarketImplV2.CallOpts)
}

// CreatePair is a paid mutator transaction binding the contract method 0x4732c7e5.
//
// Solidity: function createPair(address base, uint256 tickSize, uint256 lotSize, bytes feeData) returns(address)
func (_MarketImplV2 *MarketImplV2Transactor) CreatePair(opts *bind.TransactOpts, base common.Address, tickSize *big.Int, lotSize *big.Int, feeData []byte) (*types.Transaction, error) {
	return _MarketImplV2.contract.Transact(opts, "createPair", base, tickSize, lotSize, feeData)
}

// CreatePair is a paid mutator transaction binding the contract method 0x4732c7e5.
//
// Solidity: function createPair(address base, uint256 tickSize, uint256 lotSize, bytes feeData) returns(address)
func (_MarketImplV2 *MarketImplV2Session) CreatePair(base common.Address, tickSize *big.Int, lotSize *big.Int, feeData []byte) (*types.Transaction, error) {
	return _MarketImplV2.Contract.CreatePair(&_MarketImplV2.TransactOpts, base, tickSize, lotSize, feeData)
}

// CreatePair is a paid mutator transaction binding the contract method 0x4732c7e5.
//
// Solidity: function createPair(address base, uint256 tickSize, uint256 lotSize, bytes feeData) returns(address)
func (_MarketImplV2 *MarketImplV2TransactorSession) CreatePair(base common.Address, tickSize *big.Int, lotSize *big.Int, feeData []byte) (*types.Transaction, error) {
	return _MarketImplV2.Contract.CreatePair(&_MarketImplV2.TransactOpts, base, tickSize, lotSize, feeData)
}

// Initialize is a paid mutator transaction binding the contract method 0x01f79fb6.
//
// Solidity: function initialize(address _owner, address _router, address _quote, address _pairImpl, address _feeCollector, bytes feeData) returns()
func (_MarketImplV2 *MarketImplV2Transactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _router common.Address, _quote common.Address, _pairImpl common.Address, _feeCollector common.Address, feeData []byte) (*types.Transaction, error) {
	return _MarketImplV2.contract.Transact(opts, "initialize", _owner, _router, _quote, _pairImpl, _feeCollector, feeData)
}

// Initialize is a paid mutator transaction binding the contract method 0x01f79fb6.
//
// Solidity: function initialize(address _owner, address _router, address _quote, address _pairImpl, address _feeCollector, bytes feeData) returns()
func (_MarketImplV2 *MarketImplV2Session) Initialize(_owner common.Address, _router common.Address, _quote common.Address, _pairImpl common.Address, _feeCollector common.Address, feeData []byte) (*types.Transaction, error) {
	return _MarketImplV2.Contract.Initialize(&_MarketImplV2.TransactOpts, _owner, _router, _quote, _pairImpl, _feeCollector, feeData)
}

// Initialize is a paid mutator transaction binding the contract method 0x01f79fb6.
//
// Solidity: function initialize(address _owner, address _router, address _quote, address _pairImpl, address _feeCollector, bytes feeData) returns()
func (_MarketImplV2 *MarketImplV2TransactorSession) Initialize(_owner common.Address, _router common.Address, _quote common.Address, _pairImpl common.Address, _feeCollector common.Address, feeData []byte) (*types.Transaction, error) {
	return _MarketImplV2.Contract.Initialize(&_MarketImplV2.TransactOpts, _owner, _router, _quote, _pairImpl, _feeCollector, feeData)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MarketImplV2 *MarketImplV2Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketImplV2.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MarketImplV2 *MarketImplV2Session) RenounceOwnership() (*types.Transaction, error) {
	return _MarketImplV2.Contract.RenounceOwnership(&_MarketImplV2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MarketImplV2 *MarketImplV2TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MarketImplV2.Contract.RenounceOwnership(&_MarketImplV2.TransactOpts)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address _feeCollector) returns()
func (_MarketImplV2 *MarketImplV2Transactor) SetFeeCollector(opts *bind.TransactOpts, _feeCollector common.Address) (*types.Transaction, error) {
	return _MarketImplV2.contract.Transact(opts, "setFeeCollector", _feeCollector)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address _feeCollector) returns()
func (_MarketImplV2 *MarketImplV2Session) SetFeeCollector(_feeCollector common.Address) (*types.Transaction, error) {
	return _MarketImplV2.Contract.SetFeeCollector(&_MarketImplV2.TransactOpts, _feeCollector)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address _feeCollector) returns()
func (_MarketImplV2 *MarketImplV2TransactorSession) SetFeeCollector(_feeCollector common.Address) (*types.Transaction, error) {
	return _MarketImplV2.Contract.SetFeeCollector(&_MarketImplV2.TransactOpts, _feeCollector)
}

// SetMarketFees is a paid mutator transaction binding the contract method 0x584f6af3.
//
// Solidity: function setMarketFees(uint32 _sellerMakerFeeBps, uint32 _sellerTakerFeeBps, uint32 _buyerMakerFeeBps, uint32 _buyerTakerFeeBps) returns()
func (_MarketImplV2 *MarketImplV2Transactor) SetMarketFees(opts *bind.TransactOpts, _sellerMakerFeeBps uint32, _sellerTakerFeeBps uint32, _buyerMakerFeeBps uint32, _buyerTakerFeeBps uint32) (*types.Transaction, error) {
	return _MarketImplV2.contract.Transact(opts, "setMarketFees", _sellerMakerFeeBps, _sellerTakerFeeBps, _buyerMakerFeeBps, _buyerTakerFeeBps)
}

// SetMarketFees is a paid mutator transaction binding the contract method 0x584f6af3.
//
// Solidity: function setMarketFees(uint32 _sellerMakerFeeBps, uint32 _sellerTakerFeeBps, uint32 _buyerMakerFeeBps, uint32 _buyerTakerFeeBps) returns()
func (_MarketImplV2 *MarketImplV2Session) SetMarketFees(_sellerMakerFeeBps uint32, _sellerTakerFeeBps uint32, _buyerMakerFeeBps uint32, _buyerTakerFeeBps uint32) (*types.Transaction, error) {
	return _MarketImplV2.Contract.SetMarketFees(&_MarketImplV2.TransactOpts, _sellerMakerFeeBps, _sellerTakerFeeBps, _buyerMakerFeeBps, _buyerTakerFeeBps)
}

// SetMarketFees is a paid mutator transaction binding the contract method 0x584f6af3.
//
// Solidity: function setMarketFees(uint32 _sellerMakerFeeBps, uint32 _sellerTakerFeeBps, uint32 _buyerMakerFeeBps, uint32 _buyerTakerFeeBps) returns()
func (_MarketImplV2 *MarketImplV2TransactorSession) SetMarketFees(_sellerMakerFeeBps uint32, _sellerTakerFeeBps uint32, _buyerMakerFeeBps uint32, _buyerTakerFeeBps uint32) (*types.Transaction, error) {
	return _MarketImplV2.Contract.SetMarketFees(&_MarketImplV2.TransactOpts, _sellerMakerFeeBps, _sellerTakerFeeBps, _buyerMakerFeeBps, _buyerTakerFeeBps)
}

// SetPairImpl is a paid mutator transaction binding the contract method 0x3cc047aa.
//
// Solidity: function setPairImpl(address _pairImpl) returns()
func (_MarketImplV2 *MarketImplV2Transactor) SetPairImpl(opts *bind.TransactOpts, _pairImpl common.Address) (*types.Transaction, error) {
	return _MarketImplV2.contract.Transact(opts, "setPairImpl", _pairImpl)
}

// SetPairImpl is a paid mutator transaction binding the contract method 0x3cc047aa.
//
// Solidity: function setPairImpl(address _pairImpl) returns()
func (_MarketImplV2 *MarketImplV2Session) SetPairImpl(_pairImpl common.Address) (*types.Transaction, error) {
	return _MarketImplV2.Contract.SetPairImpl(&_MarketImplV2.TransactOpts, _pairImpl)
}

// SetPairImpl is a paid mutator transaction binding the contract method 0x3cc047aa.
//
// Solidity: function setPairImpl(address _pairImpl) returns()
func (_MarketImplV2 *MarketImplV2TransactorSession) SetPairImpl(_pairImpl common.Address) (*types.Transaction, error) {
	return _MarketImplV2.Contract.SetPairImpl(&_MarketImplV2.TransactOpts, _pairImpl)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MarketImplV2 *MarketImplV2Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MarketImplV2.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MarketImplV2 *MarketImplV2Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MarketImplV2.Contract.TransferOwnership(&_MarketImplV2.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MarketImplV2 *MarketImplV2TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MarketImplV2.Contract.TransferOwnership(&_MarketImplV2.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_MarketImplV2 *MarketImplV2Transactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _MarketImplV2.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_MarketImplV2 *MarketImplV2Session) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _MarketImplV2.Contract.UpgradeToAndCall(&_MarketImplV2.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_MarketImplV2 *MarketImplV2TransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _MarketImplV2.Contract.UpgradeToAndCall(&_MarketImplV2.TransactOpts, newImplementation, data)
}

// MarketImplV2FeeCollectorChangedIterator is returned from FilterFeeCollectorChanged and is used to iterate over the raw logs and unpacked data for FeeCollectorChanged events raised by the MarketImplV2 contract.
type MarketImplV2FeeCollectorChangedIterator struct {
	Event *MarketImplV2FeeCollectorChanged // Event containing the contract specifics and raw log

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
func (it *MarketImplV2FeeCollectorChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketImplV2FeeCollectorChanged)
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
		it.Event = new(MarketImplV2FeeCollectorChanged)
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
func (it *MarketImplV2FeeCollectorChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketImplV2FeeCollectorChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketImplV2FeeCollectorChanged represents a FeeCollectorChanged event raised by the MarketImplV2 contract.
type MarketImplV2FeeCollectorChanged struct {
	Before  common.Address
	Current common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeeCollectorChanged is a free log retrieval operation binding the contract event 0x649c5e3d0ed183894196148e193af316452b0037e77d2ff0fef23b7dc722bed0.
//
// Solidity: event FeeCollectorChanged(address indexed before, address indexed current)
func (_MarketImplV2 *MarketImplV2Filterer) FilterFeeCollectorChanged(opts *bind.FilterOpts, before []common.Address, current []common.Address) (*MarketImplV2FeeCollectorChangedIterator, error) {

	var beforeRule []interface{}
	for _, beforeItem := range before {
		beforeRule = append(beforeRule, beforeItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _MarketImplV2.contract.FilterLogs(opts, "FeeCollectorChanged", beforeRule, currentRule)
	if err != nil {
		return nil, err
	}
	return &MarketImplV2FeeCollectorChangedIterator{contract: _MarketImplV2.contract, event: "FeeCollectorChanged", logs: logs, sub: sub}, nil
}

// WatchFeeCollectorChanged is a free log subscription operation binding the contract event 0x649c5e3d0ed183894196148e193af316452b0037e77d2ff0fef23b7dc722bed0.
//
// Solidity: event FeeCollectorChanged(address indexed before, address indexed current)
func (_MarketImplV2 *MarketImplV2Filterer) WatchFeeCollectorChanged(opts *bind.WatchOpts, sink chan<- *MarketImplV2FeeCollectorChanged, before []common.Address, current []common.Address) (event.Subscription, error) {

	var beforeRule []interface{}
	for _, beforeItem := range before {
		beforeRule = append(beforeRule, beforeItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _MarketImplV2.contract.WatchLogs(opts, "FeeCollectorChanged", beforeRule, currentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketImplV2FeeCollectorChanged)
				if err := _MarketImplV2.contract.UnpackLog(event, "FeeCollectorChanged", log); err != nil {
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

// ParseFeeCollectorChanged is a log parse operation binding the contract event 0x649c5e3d0ed183894196148e193af316452b0037e77d2ff0fef23b7dc722bed0.
//
// Solidity: event FeeCollectorChanged(address indexed before, address indexed current)
func (_MarketImplV2 *MarketImplV2Filterer) ParseFeeCollectorChanged(log types.Log) (*MarketImplV2FeeCollectorChanged, error) {
	event := new(MarketImplV2FeeCollectorChanged)
	if err := _MarketImplV2.contract.UnpackLog(event, "FeeCollectorChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketImplV2InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the MarketImplV2 contract.
type MarketImplV2InitializedIterator struct {
	Event *MarketImplV2Initialized // Event containing the contract specifics and raw log

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
func (it *MarketImplV2InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketImplV2Initialized)
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
		it.Event = new(MarketImplV2Initialized)
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
func (it *MarketImplV2InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketImplV2InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketImplV2Initialized represents a Initialized event raised by the MarketImplV2 contract.
type MarketImplV2Initialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MarketImplV2 *MarketImplV2Filterer) FilterInitialized(opts *bind.FilterOpts) (*MarketImplV2InitializedIterator, error) {

	logs, sub, err := _MarketImplV2.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MarketImplV2InitializedIterator{contract: _MarketImplV2.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MarketImplV2 *MarketImplV2Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MarketImplV2Initialized) (event.Subscription, error) {

	logs, sub, err := _MarketImplV2.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketImplV2Initialized)
				if err := _MarketImplV2.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MarketImplV2 *MarketImplV2Filterer) ParseInitialized(log types.Log) (*MarketImplV2Initialized, error) {
	event := new(MarketImplV2Initialized)
	if err := _MarketImplV2.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketImplV2MarketFeesUpdatedIterator is returned from FilterMarketFeesUpdated and is used to iterate over the raw logs and unpacked data for MarketFeesUpdated events raised by the MarketImplV2 contract.
type MarketImplV2MarketFeesUpdatedIterator struct {
	Event *MarketImplV2MarketFeesUpdated // Event containing the contract specifics and raw log

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
func (it *MarketImplV2MarketFeesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketImplV2MarketFeesUpdated)
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
		it.Event = new(MarketImplV2MarketFeesUpdated)
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
func (it *MarketImplV2MarketFeesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketImplV2MarketFeesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketImplV2MarketFeesUpdated represents a MarketFeesUpdated event raised by the MarketImplV2 contract.
type MarketImplV2MarketFeesUpdated struct {
	SellerMakerFee uint32
	SellerTakerFee uint32
	BuyerMakerFee  uint32
	BuyerTakerFee  uint32
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMarketFeesUpdated is a free log retrieval operation binding the contract event 0x68657e5573176ce9149cd289efa48a58a9db7703c3b1bef2e30e41647f5ab208.
//
// Solidity: event MarketFeesUpdated(uint32 sellerMakerFee, uint32 sellerTakerFee, uint32 buyerMakerFee, uint32 buyerTakerFee)
func (_MarketImplV2 *MarketImplV2Filterer) FilterMarketFeesUpdated(opts *bind.FilterOpts) (*MarketImplV2MarketFeesUpdatedIterator, error) {

	logs, sub, err := _MarketImplV2.contract.FilterLogs(opts, "MarketFeesUpdated")
	if err != nil {
		return nil, err
	}
	return &MarketImplV2MarketFeesUpdatedIterator{contract: _MarketImplV2.contract, event: "MarketFeesUpdated", logs: logs, sub: sub}, nil
}

// WatchMarketFeesUpdated is a free log subscription operation binding the contract event 0x68657e5573176ce9149cd289efa48a58a9db7703c3b1bef2e30e41647f5ab208.
//
// Solidity: event MarketFeesUpdated(uint32 sellerMakerFee, uint32 sellerTakerFee, uint32 buyerMakerFee, uint32 buyerTakerFee)
func (_MarketImplV2 *MarketImplV2Filterer) WatchMarketFeesUpdated(opts *bind.WatchOpts, sink chan<- *MarketImplV2MarketFeesUpdated) (event.Subscription, error) {

	logs, sub, err := _MarketImplV2.contract.WatchLogs(opts, "MarketFeesUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketImplV2MarketFeesUpdated)
				if err := _MarketImplV2.contract.UnpackLog(event, "MarketFeesUpdated", log); err != nil {
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

// ParseMarketFeesUpdated is a log parse operation binding the contract event 0x68657e5573176ce9149cd289efa48a58a9db7703c3b1bef2e30e41647f5ab208.
//
// Solidity: event MarketFeesUpdated(uint32 sellerMakerFee, uint32 sellerTakerFee, uint32 buyerMakerFee, uint32 buyerTakerFee)
func (_MarketImplV2 *MarketImplV2Filterer) ParseMarketFeesUpdated(log types.Log) (*MarketImplV2MarketFeesUpdated, error) {
	event := new(MarketImplV2MarketFeesUpdated)
	if err := _MarketImplV2.contract.UnpackLog(event, "MarketFeesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketImplV2OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MarketImplV2 contract.
type MarketImplV2OwnershipTransferredIterator struct {
	Event *MarketImplV2OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MarketImplV2OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketImplV2OwnershipTransferred)
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
		it.Event = new(MarketImplV2OwnershipTransferred)
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
func (it *MarketImplV2OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketImplV2OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketImplV2OwnershipTransferred represents a OwnershipTransferred event raised by the MarketImplV2 contract.
type MarketImplV2OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MarketImplV2 *MarketImplV2Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MarketImplV2OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MarketImplV2.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MarketImplV2OwnershipTransferredIterator{contract: _MarketImplV2.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MarketImplV2 *MarketImplV2Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MarketImplV2OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MarketImplV2.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketImplV2OwnershipTransferred)
				if err := _MarketImplV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MarketImplV2 *MarketImplV2Filterer) ParseOwnershipTransferred(log types.Log) (*MarketImplV2OwnershipTransferred, error) {
	event := new(MarketImplV2OwnershipTransferred)
	if err := _MarketImplV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketImplV2PairCreatedIterator is returned from FilterPairCreated and is used to iterate over the raw logs and unpacked data for PairCreated events raised by the MarketImplV2 contract.
type MarketImplV2PairCreatedIterator struct {
	Event *MarketImplV2PairCreated // Event containing the contract specifics and raw log

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
func (it *MarketImplV2PairCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketImplV2PairCreated)
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
		it.Event = new(MarketImplV2PairCreated)
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
func (it *MarketImplV2PairCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketImplV2PairCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketImplV2PairCreated represents a PairCreated event raised by the MarketImplV2 contract.
type MarketImplV2PairCreated struct {
	Pair      common.Address
	Base      common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPairCreated is a free log retrieval operation binding the contract event 0xc1db9ba7c4b7ce660fe8d17bbcf07167549381df2abd694a970bd1402d86d313.
//
// Solidity: event PairCreated(address indexed pair, address indexed base, uint256 timestamp)
func (_MarketImplV2 *MarketImplV2Filterer) FilterPairCreated(opts *bind.FilterOpts, pair []common.Address, base []common.Address) (*MarketImplV2PairCreatedIterator, error) {

	var pairRule []interface{}
	for _, pairItem := range pair {
		pairRule = append(pairRule, pairItem)
	}
	var baseRule []interface{}
	for _, baseItem := range base {
		baseRule = append(baseRule, baseItem)
	}

	logs, sub, err := _MarketImplV2.contract.FilterLogs(opts, "PairCreated", pairRule, baseRule)
	if err != nil {
		return nil, err
	}
	return &MarketImplV2PairCreatedIterator{contract: _MarketImplV2.contract, event: "PairCreated", logs: logs, sub: sub}, nil
}

// WatchPairCreated is a free log subscription operation binding the contract event 0xc1db9ba7c4b7ce660fe8d17bbcf07167549381df2abd694a970bd1402d86d313.
//
// Solidity: event PairCreated(address indexed pair, address indexed base, uint256 timestamp)
func (_MarketImplV2 *MarketImplV2Filterer) WatchPairCreated(opts *bind.WatchOpts, sink chan<- *MarketImplV2PairCreated, pair []common.Address, base []common.Address) (event.Subscription, error) {

	var pairRule []interface{}
	for _, pairItem := range pair {
		pairRule = append(pairRule, pairItem)
	}
	var baseRule []interface{}
	for _, baseItem := range base {
		baseRule = append(baseRule, baseItem)
	}

	logs, sub, err := _MarketImplV2.contract.WatchLogs(opts, "PairCreated", pairRule, baseRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketImplV2PairCreated)
				if err := _MarketImplV2.contract.UnpackLog(event, "PairCreated", log); err != nil {
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

// ParsePairCreated is a log parse operation binding the contract event 0xc1db9ba7c4b7ce660fe8d17bbcf07167549381df2abd694a970bd1402d86d313.
//
// Solidity: event PairCreated(address indexed pair, address indexed base, uint256 timestamp)
func (_MarketImplV2 *MarketImplV2Filterer) ParsePairCreated(log types.Log) (*MarketImplV2PairCreated, error) {
	event := new(MarketImplV2PairCreated)
	if err := _MarketImplV2.contract.UnpackLog(event, "PairCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketImplV2PairImplSetIterator is returned from FilterPairImplSet and is used to iterate over the raw logs and unpacked data for PairImplSet events raised by the MarketImplV2 contract.
type MarketImplV2PairImplSetIterator struct {
	Event *MarketImplV2PairImplSet // Event containing the contract specifics and raw log

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
func (it *MarketImplV2PairImplSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketImplV2PairImplSet)
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
		it.Event = new(MarketImplV2PairImplSet)
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
func (it *MarketImplV2PairImplSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketImplV2PairImplSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketImplV2PairImplSet represents a PairImplSet event raised by the MarketImplV2 contract.
type MarketImplV2PairImplSet struct {
	Before  common.Address
	Current common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPairImplSet is a free log retrieval operation binding the contract event 0x685c330e9b04243f99ff39d8de8578c9af2bfaff120e17a1b5f1d61e004935e2.
//
// Solidity: event PairImplSet(address indexed before, address indexed current)
func (_MarketImplV2 *MarketImplV2Filterer) FilterPairImplSet(opts *bind.FilterOpts, before []common.Address, current []common.Address) (*MarketImplV2PairImplSetIterator, error) {

	var beforeRule []interface{}
	for _, beforeItem := range before {
		beforeRule = append(beforeRule, beforeItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _MarketImplV2.contract.FilterLogs(opts, "PairImplSet", beforeRule, currentRule)
	if err != nil {
		return nil, err
	}
	return &MarketImplV2PairImplSetIterator{contract: _MarketImplV2.contract, event: "PairImplSet", logs: logs, sub: sub}, nil
}

// WatchPairImplSet is a free log subscription operation binding the contract event 0x685c330e9b04243f99ff39d8de8578c9af2bfaff120e17a1b5f1d61e004935e2.
//
// Solidity: event PairImplSet(address indexed before, address indexed current)
func (_MarketImplV2 *MarketImplV2Filterer) WatchPairImplSet(opts *bind.WatchOpts, sink chan<- *MarketImplV2PairImplSet, before []common.Address, current []common.Address) (event.Subscription, error) {

	var beforeRule []interface{}
	for _, beforeItem := range before {
		beforeRule = append(beforeRule, beforeItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _MarketImplV2.contract.WatchLogs(opts, "PairImplSet", beforeRule, currentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketImplV2PairImplSet)
				if err := _MarketImplV2.contract.UnpackLog(event, "PairImplSet", log); err != nil {
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

// ParsePairImplSet is a log parse operation binding the contract event 0x685c330e9b04243f99ff39d8de8578c9af2bfaff120e17a1b5f1d61e004935e2.
//
// Solidity: event PairImplSet(address indexed before, address indexed current)
func (_MarketImplV2 *MarketImplV2Filterer) ParsePairImplSet(log types.Log) (*MarketImplV2PairImplSet, error) {
	event := new(MarketImplV2PairImplSet)
	if err := _MarketImplV2.contract.UnpackLog(event, "PairImplSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketImplV2UpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the MarketImplV2 contract.
type MarketImplV2UpgradedIterator struct {
	Event *MarketImplV2Upgraded // Event containing the contract specifics and raw log

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
func (it *MarketImplV2UpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketImplV2Upgraded)
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
		it.Event = new(MarketImplV2Upgraded)
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
func (it *MarketImplV2UpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketImplV2UpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketImplV2Upgraded represents a Upgraded event raised by the MarketImplV2 contract.
type MarketImplV2Upgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_MarketImplV2 *MarketImplV2Filterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*MarketImplV2UpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _MarketImplV2.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &MarketImplV2UpgradedIterator{contract: _MarketImplV2.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_MarketImplV2 *MarketImplV2Filterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *MarketImplV2Upgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _MarketImplV2.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketImplV2Upgraded)
				if err := _MarketImplV2.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_MarketImplV2 *MarketImplV2Filterer) ParseUpgraded(log types.Log) (*MarketImplV2Upgraded, error) {
	event := new(MarketImplV2Upgraded)
	if err := _MarketImplV2.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}