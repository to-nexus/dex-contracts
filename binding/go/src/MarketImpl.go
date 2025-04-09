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

// MarketImplMetaData contains all meta data concerning the MarketImpl contract.
var MarketImplMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CROSS_DEX\",\"outputs\":[{\"internalType\":\"contractICrossDex\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"QUOTE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROUTER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allPairs\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"bases\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"pairs\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"}],\"name\":\"baseToPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"checkTickSizeRoles\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tickSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lotSize\",\"type\":\"uint256\"}],\"name\":\"createPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deployed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeBps\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeCollector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_quote\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pairImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeCollector\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_feeBPS\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pairImpl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_feeBps\",\"type\":\"uint256\"}],\"name\":\"setFeeBps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeCollector\",\"type\":\"address\"}],\"name\":\"setFeeCollector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"before\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"current\",\"type\":\"uint256\"}],\"name\":\"FeeBpsChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"before\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"FeeCollectorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"PairCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Create2EmptyBytecode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"EnumerableMapNonexistentKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedDeployment\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"MarketAlreadyCreatedBaseAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MarketDeployPair\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"MarketInvalidBaseAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"MarketInvalidInitializeData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"b820d8aa": "CROSS_DEX()",
		"9c579839": "QUOTE()",
		"32fe7b26": "ROUTER()",
		"ad3cb1cc": "UPGRADE_INTERFACE_VERSION()",
		"c97682f8": "allPairs()",
		"364e9a19": "baseToPair(address)",
		"a1eea778": "checkTickSizeRoles(address)",
		"284cfbe0": "createPair(address,uint256,uint256)",
		"f905c15a": "deployed()",
		"24a9d853": "feeBps()",
		"c415b95c": "feeCollector()",
		"95b6ef0c": "initialize(address,address,address,address,address,uint256)",
		"8da5cb5b": "owner()",
		"35f7d456": "pairImpl()",
		"52d1902d": "proxiableUUID()",
		"715018a6": "renounceOwnership()",
		"72c27b62": "setFeeBps(uint256)",
		"a42dce80": "setFeeCollector(address)",
		"f2fde38b": "transferOwnership(address)",
		"4f1ef286": "upgradeToAndCall(address,bytes)",
	},
	Bin: "0x60a060405230608052348015610013575f5ffd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6080516124d26100f95f395f81816115030152818161152c015261174d01526124d25ff3fe608060405260043610610157575f3560e01c806395b6ef0c116100bb578063b820d8aa11610071578063c97682f811610057578063c97682f81461044d578063f2fde38b1461046f578063f905c15a1461048e575f5ffd5b8063b820d8aa146103f5578063c415b95c14610421575f5ffd5b8063a1eea778116100a1578063a1eea77814610362578063a42dce8014610381578063ad3cb1cc146103a0575f5ffd5b806395b6ef0c146103175780639c57983914610336575f5ffd5b80634f1ef28611610110578063715018a6116100f6578063715018a61461029b57806372c27b62146102af5780638da5cb5b146102ce575f5ffd5b80634f1ef2861461026457806352d1902d14610279575f5ffd5b806332fe7b261161014057806332fe7b26146101ed57806335f7d45614610219578063364e9a1914610245575f5ffd5b806324a9d8531461015b578063284cfbe0146101a9575b5f5ffd5b348015610166575f5ffd5b5060055461018f9074010000000000000000000000000000000000000000900463ffffffff1681565b60405163ffffffff90911681526020015b60405180910390f35b3480156101b4575f5ffd5b506101c86101c3366004611d2c565b6104a2565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101a0565b3480156101f8575f5ffd5b506003546101c89073ffffffffffffffffffffffffffffffffffffffff1681565b348015610224575f5ffd5b506004546101c89073ffffffffffffffffffffffffffffffffffffffff1681565b348015610250575f5ffd5b506101c861025f366004611d5c565b6109ca565b610277610272366004611da2565b6109dc565b005b348015610284575f5ffd5b5061028d6109fb565b6040519081526020016101a0565b3480156102a6575f5ffd5b50610277610a29565b3480156102ba575f5ffd5b506102776102c9366004611ea3565b610a3c565b3480156102d9575f5ffd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c1993005473ffffffffffffffffffffffffffffffffffffffff166101c8565b348015610322575f5ffd5b50610277610331366004611eba565b610b3f565b348015610341575f5ffd5b506002546101c89073ffffffffffffffffffffffffffffffffffffffff1681565b34801561036d575f5ffd5b5061027761037c366004611d5c565b61100a565b34801561038c575f5ffd5b5061027761039b366004611d5c565b61108c565b3480156103ab575f5ffd5b506103e86040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b6040516101a09190611f71565b348015610400575f5ffd5b506001546101c89073ffffffffffffffffffffffffffffffffffffffff1681565b34801561042c575f5ffd5b506005546101c89073ffffffffffffffffffffffffffffffffffffffff1681565b348015610458575f5ffd5b50610461611190565b6040516101a0929190611fd3565b34801561047a575f5ffd5b50610277610489366004611d5c565b61129f565b348015610499575f5ffd5b5061028d5f5481565b5f6104ab611302565b73ffffffffffffffffffffffffffffffffffffffff841615806104e8575060025473ffffffffffffffffffffffffffffffffffffffff8581169116145b1561053c576040517f20f8254c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024015b60405180910390fd5b5f8473ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa158015610586573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105aa9190611ff7565b60ff169050805f03610600576040517f20f8254c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff86166004820152602401610533565b61060b600686611390565b1561065a576040517fc13fcfdc00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff86166004820152602401610533565b5f6040518060200161066b90611cf7565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082820381018352601f90910116604081905260045460035460025473ffffffffffffffffffffffffffffffffffffffff91821660248501528116604484015289811660648401526084830189905260a48301889052169060c401604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa6b63eb800000000000000000000000000000000000000000000000000000000179052905161076b93929101612017565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290526107a7929160200161205c565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606089901b16602083015291505f906034016040516020818303038152906040528051906020012090505f61082d5f83856113b1565b905073ffffffffffffffffffffffffffffffffffffffff811661087c576040517f3ff6a26100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6108886006898361149d565b6108d6576040517fc13fcfdc00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff89166004820152602401610533565b6001546040517fe9f7206b00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83811660048301529091169063e9f7206b906024015f604051808303815f87803b158015610940575f5ffd5b505af1158015610952573d5f5f3e3d5ffd5b505050508773ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167fc1db9ba7c4b7ce660fe8d17bbcf07167549381df2abd694a970bd1402d86d313426040516109b591815260200190565b60405180910390a393505050505b9392505050565b5f6109d66006836114ca565b92915050565b6109e46114eb565b6109ed826115ef565b6109f782826115f7565b5050565b5f610a04611735565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b610a31611302565b610a3a5f6117a4565b565b610a44611302565b63ffffffff811115610aa4576040517fa25d2db50000000000000000000000000000000000000000000000000000000081527f66656542505300000000000000000000000000000000000000000000000000006004820152602401610533565b600554604051829174010000000000000000000000000000000000000000900463ffffffff16907ff878449785c2378ffd246b9c0c874a3aeec940d231ad03fe4a62a1cce095632d905f90a36005805463ffffffff90921674010000000000000000000000000000000000000000027fffffffffffffffff00000000ffffffffffffffffffffffffffffffffffffffff909216919091179055565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff165f81158015610b895750825b90505f8267ffffffffffffffff166001148015610ba55750303b155b905081158015610bb3575080155b15610bea576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001660011785558315610c4b5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b73ffffffffffffffffffffffffffffffffffffffff8b16610cba576040517fa25d2db50000000000000000000000000000000000000000000000000000000081527f6f776e65720000000000000000000000000000000000000000000000000000006004820152602401610533565b73ffffffffffffffffffffffffffffffffffffffff8a16610d29576040517fa25d2db50000000000000000000000000000000000000000000000000000000081527f726f7574657200000000000000000000000000000000000000000000000000006004820152602401610533565b73ffffffffffffffffffffffffffffffffffffffff8916610d98576040517fa25d2db50000000000000000000000000000000000000000000000000000000081527f71756f74650000000000000000000000000000000000000000000000000000006004820152602401610533565b73ffffffffffffffffffffffffffffffffffffffff8816610e07576040517fa25d2db50000000000000000000000000000000000000000000000000000000081527f70616972496d706c0000000000000000000000000000000000000000000000006004820152602401610533565b73ffffffffffffffffffffffffffffffffffffffff8716610e76576040517fa25d2db50000000000000000000000000000000000000000000000000000000081527f666565436f6c6c6563746f7200000000000000000000000000000000000000006004820152602401610533565b63ffffffff861115610ed6576040517fa25d2db50000000000000000000000000000000000000000000000000000000081527f66656542505300000000000000000000000000000000000000000000000000006004820152602401610533565b435f5533600180547fffffffffffffffffffffffff000000000000000000000000000000000000000090811673ffffffffffffffffffffffffffffffffffffffff938416179091556002805482168c84161790556003805482168d8416179055600480549091168a8316179055600580549189167fffffffffffffffff000000000000000000000000000000000000000000000000909216919091177401000000000000000000000000000000000000000063ffffffff891602179055610f9c8b611839565b8315610ffd5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050505050565b6001546040517fa1eea77800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83811660048301529091169063a1eea778906024015f6040518083038186803b158015611073575f5ffd5b505afa158015611085573d5f5f3e3d5ffd5b5050505050565b611094611302565b73ffffffffffffffffffffffffffffffffffffffff8116611103576040517fa25d2db50000000000000000000000000000000000000000000000000000000081527f666565436f6c6c6563746f7200000000000000000000000000000000000000006004820152602401610533565b60055460405173ffffffffffffffffffffffffffffffffffffffff8084169216907f649c5e3d0ed183894196148e193af316452b0037e77d2ff0fef23b7dc722bed0905f90a3600580547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b6060805f61119e600661184a565b90508067ffffffffffffffff8111156111b9576111b9611d75565b6040519080825280602002602001820160405280156111e2578160200160208202803683370190505b5092508067ffffffffffffffff8111156111fe576111fe611d75565b604051908082528060200260200182016040528015611227578160200160208202803683370190505b5091505f5b818110156112995761123f600682611854565b85838151811061125157611251612070565b6020026020010185848151811061126a5761126a612070565b73ffffffffffffffffffffffffffffffffffffffff93841660209182029290920101529116905260010161122c565b50509091565b6112a7611302565b73ffffffffffffffffffffffffffffffffffffffff81166112f6576040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081525f6004820152602401610533565b6112ff816117a4565b50565b336113417f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c1993005473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614610a3a576040517f118cdaa7000000000000000000000000000000000000000000000000000000008152336004820152602401610533565b5f6109c38373ffffffffffffffffffffffffffffffffffffffff841661186f565b5f834710156113f5576040517fcf47918100000000000000000000000000000000000000000000000000000000815247600482015260248101859052604401610533565b81515f0361142f576040517f4ca249dc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8282516020840186f590503d151981151615611450576040513d5f823e3d81fd5b73ffffffffffffffffffffffffffffffffffffffff81166109c3576040517fb06ebf3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6114c28473ffffffffffffffffffffffffffffffffffffffff80861690851661187a565b949350505050565b5f6109c38373ffffffffffffffffffffffffffffffffffffffff8416611896565b3073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614806115b857507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1661159f7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614155b15610a3a576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6112ff611302565b8173ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa92505050801561167c575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682019092526116799181019061209d565b60015b6116ca576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610533565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8114611726576040517faa1d49a400000000000000000000000000000000000000000000000000000000815260048101829052602401610533565b61173083836118f5565b505050565b3073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610a3a576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080547fffffffffffffffffffffffff0000000000000000000000000000000000000000811673ffffffffffffffffffffffffffffffffffffffff848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b611841611957565b6112ff816119be565b5f6109d6826119c6565b5f80808061186286866119d0565b9097909650945050505050565b5f6109c383836119f9565b5f82815260028401602052604081208290556114c28484611a10565b5f818152600283016020526040812054801580156118bb57506118b9848461186f565b155b156109c3576040517f02b5668600000000000000000000000000000000000000000000000000000000815260048101849052602401610533565b6118fe82611a1b565b60405173ffffffffffffffffffffffffffffffffffffffff8316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a280511561194f576117308282611ae9565b6109f7611b68565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff16610a3a576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6112a7611957565b5f6109d682611ba0565b5f80806119dd8585611ba9565b5f81815260029690960160205260409095205494959350505050565b5f81815260018301602052604081205415156109c3565b5f6109c38383611bb4565b8073ffffffffffffffffffffffffffffffffffffffff163b5f03611a83576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152602401610533565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60605f5f8473ffffffffffffffffffffffffffffffffffffffff1684604051611b1291906120b4565b5f60405180830381855af49150503d805f8114611b4a576040519150601f19603f3d011682016040523d82523d5f602084013e611b4f565b606091505b5091509150611b5f858383611c00565b95945050505050565b3415610a3a576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6109d6825490565b5f6109c38383611c8f565b5f818152600183016020526040812054611bf957508154600181810184555f8481526020808220909301849055845484825282860190935260409020919091556109d6565b505f6109d6565b606082611c1557611c1082611cb5565b6109c3565b8151158015611c39575073ffffffffffffffffffffffffffffffffffffffff84163b155b15611c88576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610533565b50806109c3565b5f825f018281548110611ca457611ca4612070565b905f5260205f200154905092915050565b805115611cc55780518082602001fd5b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103dd806120c083390190565b803573ffffffffffffffffffffffffffffffffffffffff81168114611d27575f5ffd5b919050565b5f5f5f60608486031215611d3e575f5ffd5b611d4784611d04565b95602085013595506040909401359392505050565b5f60208284031215611d6c575f5ffd5b6109c382611d04565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f5f60408385031215611db3575f5ffd5b611dbc83611d04565b9150602083013567ffffffffffffffff811115611dd7575f5ffd5b8301601f81018513611de7575f5ffd5b803567ffffffffffffffff811115611e0157611e01611d75565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff82111715611e6d57611e6d611d75565b604052818152828201602001871015611e84575f5ffd5b816020840160208301375f602083830101528093505050509250929050565b5f60208284031215611eb3575f5ffd5b5035919050565b5f5f5f5f5f5f60c08789031215611ecf575f5ffd5b611ed887611d04565b9550611ee660208801611d04565b9450611ef460408801611d04565b9350611f0260608801611d04565b9250611f1060808801611d04565b9598949750929591949360a090920135925050565b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f6109c36020830184611f25565b5f8151808452602084019350602083015f5b82811015611fc957815173ffffffffffffffffffffffffffffffffffffffff16865260209586019590910190600101611f95565b5093949350505050565b604081525f611fe56040830185611f83565b8281036020840152611b5f8185611f83565b5f60208284031215612007575f5ffd5b815160ff811681146109c3575f5ffd5b73ffffffffffffffffffffffffffffffffffffffff83168152604060208201525f6114c26040830184611f25565b5f81518060208401855e5f93019283525090919050565b5f6114c261206a8386612045565b84612045565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f602082840312156120ad575f5ffd5b5051919050565b5f6109c3828461204556fe60806040526040516103dd3803806103dd8339810160408190526100229161023c565b61002c8282610033565b5050610321565b61003c82610091565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a280511561008557610080828261010c565b505050565b61008d61017f565b5050565b806001600160a01b03163b5f036100cb57604051634c9c8ce360e01b81526001600160a01b03821660048201526024015b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b031684604051610128919061030b565b5f60405180830381855af49150503d805f8114610160576040519150601f19603f3d011682016040523d82523d5f602084013e610165565b606091505b5090925090506101768583836101a0565b95945050505050565b341561019e5760405163b398979f60e01b815260040160405180910390fd5b565b6060826101b5576101b0826101ff565b6101f8565b81511580156101cc57506001600160a01b0384163b155b156101f557604051639996b31560e01b81526001600160a01b03851660048201526024016100c2565b50805b9392505050565b80511561020f5780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b634e487b7160e01b5f52604160045260245ffd5b5f5f6040838503121561024d575f5ffd5b82516001600160a01b0381168114610263575f5ffd5b60208401519092506001600160401b0381111561027e575f5ffd5b8301601f8101851361028e575f5ffd5b80516001600160401b038111156102a7576102a7610228565b604051601f8201601f19908116603f011681016001600160401b03811182821017156102d5576102d5610228565b6040528181528282016020018710156102ec575f5ffd5b8160208401602083015e5f602083830101528093505050509250929050565b5f82518060208501845e5f920191825250919050565b60b08061032d5f395ff3fe6080604052600a600c565b005b60186014601a565b605d565b565b5f60587f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5473ffffffffffffffffffffffffffffffffffffffff1690565b905090565b365f5f375f5f365f845af43d5f5f3e8080156076573d5ff35b3d5ffdfea26469706673582212209fd71e6f878ca8212eaa80602ace1507d0020bad3381eef3782249983b65738964736f6c634300081c0033a2646970667358221220e9b06d3942c509b3f2e6432bd15c900ca22e37825ac71c08ff158c536d229a6264736f6c634300081c0033",
}

// MarketImplABI is the input ABI used to generate the binding from.
// Deprecated: Use MarketImplMetaData.ABI instead.
var MarketImplABI = MarketImplMetaData.ABI

// Deprecated: Use MarketImplMetaData.Sigs instead.
// MarketImplFuncSigs maps the 4-byte function signature to its string representation.
var MarketImplFuncSigs = MarketImplMetaData.Sigs

// MarketImplBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MarketImplMetaData.Bin instead.
var MarketImplBin = MarketImplMetaData.Bin

// DeployMarketImpl deploys a new Ethereum contract, binding an instance of MarketImpl to it.
func DeployMarketImpl(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MarketImpl, error) {
	parsed, err := MarketImplMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MarketImplBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MarketImpl{MarketImplCaller: MarketImplCaller{contract: contract}, MarketImplTransactor: MarketImplTransactor{contract: contract}, MarketImplFilterer: MarketImplFilterer{contract: contract}}, nil
}

// MarketImpl is an auto generated Go binding around an Ethereum contract.
type MarketImpl struct {
	MarketImplCaller     // Read-only binding to the contract
	MarketImplTransactor // Write-only binding to the contract
	MarketImplFilterer   // Log filterer for contract events
}

// MarketImplCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketImplCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketImplTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketImplTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketImplFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketImplFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketImplSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketImplSession struct {
	Contract     *MarketImpl       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketImplCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketImplCallerSession struct {
	Contract *MarketImplCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MarketImplTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketImplTransactorSession struct {
	Contract     *MarketImplTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MarketImplRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketImplRaw struct {
	Contract *MarketImpl // Generic contract binding to access the raw methods on
}

// MarketImplCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketImplCallerRaw struct {
	Contract *MarketImplCaller // Generic read-only contract binding to access the raw methods on
}

// MarketImplTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketImplTransactorRaw struct {
	Contract *MarketImplTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarketImpl creates a new instance of MarketImpl, bound to a specific deployed contract.
func NewMarketImpl(address common.Address, backend bind.ContractBackend) (*MarketImpl, error) {
	contract, err := bindMarketImpl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MarketImpl{MarketImplCaller: MarketImplCaller{contract: contract}, MarketImplTransactor: MarketImplTransactor{contract: contract}, MarketImplFilterer: MarketImplFilterer{contract: contract}}, nil
}

// NewMarketImplCaller creates a new read-only instance of MarketImpl, bound to a specific deployed contract.
func NewMarketImplCaller(address common.Address, caller bind.ContractCaller) (*MarketImplCaller, error) {
	contract, err := bindMarketImpl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketImplCaller{contract: contract}, nil
}

// NewMarketImplTransactor creates a new write-only instance of MarketImpl, bound to a specific deployed contract.
func NewMarketImplTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketImplTransactor, error) {
	contract, err := bindMarketImpl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketImplTransactor{contract: contract}, nil
}

// NewMarketImplFilterer creates a new log filterer instance of MarketImpl, bound to a specific deployed contract.
func NewMarketImplFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketImplFilterer, error) {
	contract, err := bindMarketImpl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketImplFilterer{contract: contract}, nil
}

// bindMarketImpl binds a generic wrapper to an already deployed contract.
func bindMarketImpl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MarketImplMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MarketImpl *MarketImplRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MarketImpl.Contract.MarketImplCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MarketImpl *MarketImplRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketImpl.Contract.MarketImplTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MarketImpl *MarketImplRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MarketImpl.Contract.MarketImplTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MarketImpl *MarketImplCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MarketImpl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MarketImpl *MarketImplTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketImpl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MarketImpl *MarketImplTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MarketImpl.Contract.contract.Transact(opts, method, params...)
}

// CROSSDEX is a free data retrieval call binding the contract method 0xb820d8aa.
//
// Solidity: function CROSS_DEX() view returns(address)
func (_MarketImpl *MarketImplCaller) CROSSDEX(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarketImpl.contract.Call(opts, &out, "CROSS_DEX")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CROSSDEX is a free data retrieval call binding the contract method 0xb820d8aa.
//
// Solidity: function CROSS_DEX() view returns(address)
func (_MarketImpl *MarketImplSession) CROSSDEX() (common.Address, error) {
	return _MarketImpl.Contract.CROSSDEX(&_MarketImpl.CallOpts)
}

// CROSSDEX is a free data retrieval call binding the contract method 0xb820d8aa.
//
// Solidity: function CROSS_DEX() view returns(address)
func (_MarketImpl *MarketImplCallerSession) CROSSDEX() (common.Address, error) {
	return _MarketImpl.Contract.CROSSDEX(&_MarketImpl.CallOpts)
}

// QUOTE is a free data retrieval call binding the contract method 0x9c579839.
//
// Solidity: function QUOTE() view returns(address)
func (_MarketImpl *MarketImplCaller) QUOTE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarketImpl.contract.Call(opts, &out, "QUOTE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// QUOTE is a free data retrieval call binding the contract method 0x9c579839.
//
// Solidity: function QUOTE() view returns(address)
func (_MarketImpl *MarketImplSession) QUOTE() (common.Address, error) {
	return _MarketImpl.Contract.QUOTE(&_MarketImpl.CallOpts)
}

// QUOTE is a free data retrieval call binding the contract method 0x9c579839.
//
// Solidity: function QUOTE() view returns(address)
func (_MarketImpl *MarketImplCallerSession) QUOTE() (common.Address, error) {
	return _MarketImpl.Contract.QUOTE(&_MarketImpl.CallOpts)
}

// ROUTER is a free data retrieval call binding the contract method 0x32fe7b26.
//
// Solidity: function ROUTER() view returns(address)
func (_MarketImpl *MarketImplCaller) ROUTER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarketImpl.contract.Call(opts, &out, "ROUTER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ROUTER is a free data retrieval call binding the contract method 0x32fe7b26.
//
// Solidity: function ROUTER() view returns(address)
func (_MarketImpl *MarketImplSession) ROUTER() (common.Address, error) {
	return _MarketImpl.Contract.ROUTER(&_MarketImpl.CallOpts)
}

// ROUTER is a free data retrieval call binding the contract method 0x32fe7b26.
//
// Solidity: function ROUTER() view returns(address)
func (_MarketImpl *MarketImplCallerSession) ROUTER() (common.Address, error) {
	return _MarketImpl.Contract.ROUTER(&_MarketImpl.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_MarketImpl *MarketImplCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MarketImpl.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_MarketImpl *MarketImplSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _MarketImpl.Contract.UPGRADEINTERFACEVERSION(&_MarketImpl.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_MarketImpl *MarketImplCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _MarketImpl.Contract.UPGRADEINTERFACEVERSION(&_MarketImpl.CallOpts)
}

// AllPairs is a free data retrieval call binding the contract method 0xc97682f8.
//
// Solidity: function allPairs() view returns(address[] bases, address[] pairs)
func (_MarketImpl *MarketImplCaller) AllPairs(opts *bind.CallOpts) (struct {
	Bases []common.Address
	Pairs []common.Address
}, error) {
	var out []interface{}
	err := _MarketImpl.contract.Call(opts, &out, "allPairs")

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
func (_MarketImpl *MarketImplSession) AllPairs() (struct {
	Bases []common.Address
	Pairs []common.Address
}, error) {
	return _MarketImpl.Contract.AllPairs(&_MarketImpl.CallOpts)
}

// AllPairs is a free data retrieval call binding the contract method 0xc97682f8.
//
// Solidity: function allPairs() view returns(address[] bases, address[] pairs)
func (_MarketImpl *MarketImplCallerSession) AllPairs() (struct {
	Bases []common.Address
	Pairs []common.Address
}, error) {
	return _MarketImpl.Contract.AllPairs(&_MarketImpl.CallOpts)
}

// BaseToPair is a free data retrieval call binding the contract method 0x364e9a19.
//
// Solidity: function baseToPair(address base) view returns(address)
func (_MarketImpl *MarketImplCaller) BaseToPair(opts *bind.CallOpts, base common.Address) (common.Address, error) {
	var out []interface{}
	err := _MarketImpl.contract.Call(opts, &out, "baseToPair", base)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BaseToPair is a free data retrieval call binding the contract method 0x364e9a19.
//
// Solidity: function baseToPair(address base) view returns(address)
func (_MarketImpl *MarketImplSession) BaseToPair(base common.Address) (common.Address, error) {
	return _MarketImpl.Contract.BaseToPair(&_MarketImpl.CallOpts, base)
}

// BaseToPair is a free data retrieval call binding the contract method 0x364e9a19.
//
// Solidity: function baseToPair(address base) view returns(address)
func (_MarketImpl *MarketImplCallerSession) BaseToPair(base common.Address) (common.Address, error) {
	return _MarketImpl.Contract.BaseToPair(&_MarketImpl.CallOpts, base)
}

// CheckTickSizeRoles is a free data retrieval call binding the contract method 0xa1eea778.
//
// Solidity: function checkTickSizeRoles(address account) view returns()
func (_MarketImpl *MarketImplCaller) CheckTickSizeRoles(opts *bind.CallOpts, account common.Address) error {
	var out []interface{}
	err := _MarketImpl.contract.Call(opts, &out, "checkTickSizeRoles", account)

	if err != nil {
		return err
	}

	return err

}

// CheckTickSizeRoles is a free data retrieval call binding the contract method 0xa1eea778.
//
// Solidity: function checkTickSizeRoles(address account) view returns()
func (_MarketImpl *MarketImplSession) CheckTickSizeRoles(account common.Address) error {
	return _MarketImpl.Contract.CheckTickSizeRoles(&_MarketImpl.CallOpts, account)
}

// CheckTickSizeRoles is a free data retrieval call binding the contract method 0xa1eea778.
//
// Solidity: function checkTickSizeRoles(address account) view returns()
func (_MarketImpl *MarketImplCallerSession) CheckTickSizeRoles(account common.Address) error {
	return _MarketImpl.Contract.CheckTickSizeRoles(&_MarketImpl.CallOpts, account)
}

// Deployed is a free data retrieval call binding the contract method 0xf905c15a.
//
// Solidity: function deployed() view returns(uint256)
func (_MarketImpl *MarketImplCaller) Deployed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MarketImpl.contract.Call(opts, &out, "deployed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Deployed is a free data retrieval call binding the contract method 0xf905c15a.
//
// Solidity: function deployed() view returns(uint256)
func (_MarketImpl *MarketImplSession) Deployed() (*big.Int, error) {
	return _MarketImpl.Contract.Deployed(&_MarketImpl.CallOpts)
}

// Deployed is a free data retrieval call binding the contract method 0xf905c15a.
//
// Solidity: function deployed() view returns(uint256)
func (_MarketImpl *MarketImplCallerSession) Deployed() (*big.Int, error) {
	return _MarketImpl.Contract.Deployed(&_MarketImpl.CallOpts)
}

// FeeBps is a free data retrieval call binding the contract method 0x24a9d853.
//
// Solidity: function feeBps() view returns(uint32)
func (_MarketImpl *MarketImplCaller) FeeBps(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _MarketImpl.contract.Call(opts, &out, "feeBps")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// FeeBps is a free data retrieval call binding the contract method 0x24a9d853.
//
// Solidity: function feeBps() view returns(uint32)
func (_MarketImpl *MarketImplSession) FeeBps() (uint32, error) {
	return _MarketImpl.Contract.FeeBps(&_MarketImpl.CallOpts)
}

// FeeBps is a free data retrieval call binding the contract method 0x24a9d853.
//
// Solidity: function feeBps() view returns(uint32)
func (_MarketImpl *MarketImplCallerSession) FeeBps() (uint32, error) {
	return _MarketImpl.Contract.FeeBps(&_MarketImpl.CallOpts)
}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_MarketImpl *MarketImplCaller) FeeCollector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarketImpl.contract.Call(opts, &out, "feeCollector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_MarketImpl *MarketImplSession) FeeCollector() (common.Address, error) {
	return _MarketImpl.Contract.FeeCollector(&_MarketImpl.CallOpts)
}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_MarketImpl *MarketImplCallerSession) FeeCollector() (common.Address, error) {
	return _MarketImpl.Contract.FeeCollector(&_MarketImpl.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MarketImpl *MarketImplCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarketImpl.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MarketImpl *MarketImplSession) Owner() (common.Address, error) {
	return _MarketImpl.Contract.Owner(&_MarketImpl.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MarketImpl *MarketImplCallerSession) Owner() (common.Address, error) {
	return _MarketImpl.Contract.Owner(&_MarketImpl.CallOpts)
}

// PairImpl is a free data retrieval call binding the contract method 0x35f7d456.
//
// Solidity: function pairImpl() view returns(address)
func (_MarketImpl *MarketImplCaller) PairImpl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarketImpl.contract.Call(opts, &out, "pairImpl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PairImpl is a free data retrieval call binding the contract method 0x35f7d456.
//
// Solidity: function pairImpl() view returns(address)
func (_MarketImpl *MarketImplSession) PairImpl() (common.Address, error) {
	return _MarketImpl.Contract.PairImpl(&_MarketImpl.CallOpts)
}

// PairImpl is a free data retrieval call binding the contract method 0x35f7d456.
//
// Solidity: function pairImpl() view returns(address)
func (_MarketImpl *MarketImplCallerSession) PairImpl() (common.Address, error) {
	return _MarketImpl.Contract.PairImpl(&_MarketImpl.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_MarketImpl *MarketImplCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MarketImpl.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_MarketImpl *MarketImplSession) ProxiableUUID() ([32]byte, error) {
	return _MarketImpl.Contract.ProxiableUUID(&_MarketImpl.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_MarketImpl *MarketImplCallerSession) ProxiableUUID() ([32]byte, error) {
	return _MarketImpl.Contract.ProxiableUUID(&_MarketImpl.CallOpts)
}

// CreatePair is a paid mutator transaction binding the contract method 0x284cfbe0.
//
// Solidity: function createPair(address base, uint256 tickSize, uint256 lotSize) returns(address)
func (_MarketImpl *MarketImplTransactor) CreatePair(opts *bind.TransactOpts, base common.Address, tickSize *big.Int, lotSize *big.Int) (*types.Transaction, error) {
	return _MarketImpl.contract.Transact(opts, "createPair", base, tickSize, lotSize)
}

// CreatePair is a paid mutator transaction binding the contract method 0x284cfbe0.
//
// Solidity: function createPair(address base, uint256 tickSize, uint256 lotSize) returns(address)
func (_MarketImpl *MarketImplSession) CreatePair(base common.Address, tickSize *big.Int, lotSize *big.Int) (*types.Transaction, error) {
	return _MarketImpl.Contract.CreatePair(&_MarketImpl.TransactOpts, base, tickSize, lotSize)
}

// CreatePair is a paid mutator transaction binding the contract method 0x284cfbe0.
//
// Solidity: function createPair(address base, uint256 tickSize, uint256 lotSize) returns(address)
func (_MarketImpl *MarketImplTransactorSession) CreatePair(base common.Address, tickSize *big.Int, lotSize *big.Int) (*types.Transaction, error) {
	return _MarketImpl.Contract.CreatePair(&_MarketImpl.TransactOpts, base, tickSize, lotSize)
}

// Initialize is a paid mutator transaction binding the contract method 0x95b6ef0c.
//
// Solidity: function initialize(address _owner, address _router, address _quote, address _pairImpl, address _feeCollector, uint256 _feeBPS) returns()
func (_MarketImpl *MarketImplTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _router common.Address, _quote common.Address, _pairImpl common.Address, _feeCollector common.Address, _feeBPS *big.Int) (*types.Transaction, error) {
	return _MarketImpl.contract.Transact(opts, "initialize", _owner, _router, _quote, _pairImpl, _feeCollector, _feeBPS)
}

// Initialize is a paid mutator transaction binding the contract method 0x95b6ef0c.
//
// Solidity: function initialize(address _owner, address _router, address _quote, address _pairImpl, address _feeCollector, uint256 _feeBPS) returns()
func (_MarketImpl *MarketImplSession) Initialize(_owner common.Address, _router common.Address, _quote common.Address, _pairImpl common.Address, _feeCollector common.Address, _feeBPS *big.Int) (*types.Transaction, error) {
	return _MarketImpl.Contract.Initialize(&_MarketImpl.TransactOpts, _owner, _router, _quote, _pairImpl, _feeCollector, _feeBPS)
}

// Initialize is a paid mutator transaction binding the contract method 0x95b6ef0c.
//
// Solidity: function initialize(address _owner, address _router, address _quote, address _pairImpl, address _feeCollector, uint256 _feeBPS) returns()
func (_MarketImpl *MarketImplTransactorSession) Initialize(_owner common.Address, _router common.Address, _quote common.Address, _pairImpl common.Address, _feeCollector common.Address, _feeBPS *big.Int) (*types.Transaction, error) {
	return _MarketImpl.Contract.Initialize(&_MarketImpl.TransactOpts, _owner, _router, _quote, _pairImpl, _feeCollector, _feeBPS)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MarketImpl *MarketImplTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketImpl.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MarketImpl *MarketImplSession) RenounceOwnership() (*types.Transaction, error) {
	return _MarketImpl.Contract.RenounceOwnership(&_MarketImpl.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MarketImpl *MarketImplTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MarketImpl.Contract.RenounceOwnership(&_MarketImpl.TransactOpts)
}

// SetFeeBps is a paid mutator transaction binding the contract method 0x72c27b62.
//
// Solidity: function setFeeBps(uint256 _feeBps) returns()
func (_MarketImpl *MarketImplTransactor) SetFeeBps(opts *bind.TransactOpts, _feeBps *big.Int) (*types.Transaction, error) {
	return _MarketImpl.contract.Transact(opts, "setFeeBps", _feeBps)
}

// SetFeeBps is a paid mutator transaction binding the contract method 0x72c27b62.
//
// Solidity: function setFeeBps(uint256 _feeBps) returns()
func (_MarketImpl *MarketImplSession) SetFeeBps(_feeBps *big.Int) (*types.Transaction, error) {
	return _MarketImpl.Contract.SetFeeBps(&_MarketImpl.TransactOpts, _feeBps)
}

// SetFeeBps is a paid mutator transaction binding the contract method 0x72c27b62.
//
// Solidity: function setFeeBps(uint256 _feeBps) returns()
func (_MarketImpl *MarketImplTransactorSession) SetFeeBps(_feeBps *big.Int) (*types.Transaction, error) {
	return _MarketImpl.Contract.SetFeeBps(&_MarketImpl.TransactOpts, _feeBps)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address _feeCollector) returns()
func (_MarketImpl *MarketImplTransactor) SetFeeCollector(opts *bind.TransactOpts, _feeCollector common.Address) (*types.Transaction, error) {
	return _MarketImpl.contract.Transact(opts, "setFeeCollector", _feeCollector)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address _feeCollector) returns()
func (_MarketImpl *MarketImplSession) SetFeeCollector(_feeCollector common.Address) (*types.Transaction, error) {
	return _MarketImpl.Contract.SetFeeCollector(&_MarketImpl.TransactOpts, _feeCollector)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address _feeCollector) returns()
func (_MarketImpl *MarketImplTransactorSession) SetFeeCollector(_feeCollector common.Address) (*types.Transaction, error) {
	return _MarketImpl.Contract.SetFeeCollector(&_MarketImpl.TransactOpts, _feeCollector)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MarketImpl *MarketImplTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MarketImpl.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MarketImpl *MarketImplSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MarketImpl.Contract.TransferOwnership(&_MarketImpl.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MarketImpl *MarketImplTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MarketImpl.Contract.TransferOwnership(&_MarketImpl.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_MarketImpl *MarketImplTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _MarketImpl.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_MarketImpl *MarketImplSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _MarketImpl.Contract.UpgradeToAndCall(&_MarketImpl.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_MarketImpl *MarketImplTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _MarketImpl.Contract.UpgradeToAndCall(&_MarketImpl.TransactOpts, newImplementation, data)
}

// MarketImplFeeBpsChangedIterator is returned from FilterFeeBpsChanged and is used to iterate over the raw logs and unpacked data for FeeBpsChanged events raised by the MarketImpl contract.
type MarketImplFeeBpsChangedIterator struct {
	Event *MarketImplFeeBpsChanged // Event containing the contract specifics and raw log

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
func (it *MarketImplFeeBpsChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketImplFeeBpsChanged)
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
		it.Event = new(MarketImplFeeBpsChanged)
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
func (it *MarketImplFeeBpsChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketImplFeeBpsChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketImplFeeBpsChanged represents a FeeBpsChanged event raised by the MarketImpl contract.
type MarketImplFeeBpsChanged struct {
	Before  *big.Int
	Current *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeeBpsChanged is a free log retrieval operation binding the contract event 0xf878449785c2378ffd246b9c0c874a3aeec940d231ad03fe4a62a1cce095632d.
//
// Solidity: event FeeBpsChanged(uint256 indexed before, uint256 indexed current)
func (_MarketImpl *MarketImplFilterer) FilterFeeBpsChanged(opts *bind.FilterOpts, before []*big.Int, current []*big.Int) (*MarketImplFeeBpsChangedIterator, error) {

	var beforeRule []interface{}
	for _, beforeItem := range before {
		beforeRule = append(beforeRule, beforeItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _MarketImpl.contract.FilterLogs(opts, "FeeBpsChanged", beforeRule, currentRule)
	if err != nil {
		return nil, err
	}
	return &MarketImplFeeBpsChangedIterator{contract: _MarketImpl.contract, event: "FeeBpsChanged", logs: logs, sub: sub}, nil
}

// WatchFeeBpsChanged is a free log subscription operation binding the contract event 0xf878449785c2378ffd246b9c0c874a3aeec940d231ad03fe4a62a1cce095632d.
//
// Solidity: event FeeBpsChanged(uint256 indexed before, uint256 indexed current)
func (_MarketImpl *MarketImplFilterer) WatchFeeBpsChanged(opts *bind.WatchOpts, sink chan<- *MarketImplFeeBpsChanged, before []*big.Int, current []*big.Int) (event.Subscription, error) {

	var beforeRule []interface{}
	for _, beforeItem := range before {
		beforeRule = append(beforeRule, beforeItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _MarketImpl.contract.WatchLogs(opts, "FeeBpsChanged", beforeRule, currentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketImplFeeBpsChanged)
				if err := _MarketImpl.contract.UnpackLog(event, "FeeBpsChanged", log); err != nil {
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

// ParseFeeBpsChanged is a log parse operation binding the contract event 0xf878449785c2378ffd246b9c0c874a3aeec940d231ad03fe4a62a1cce095632d.
//
// Solidity: event FeeBpsChanged(uint256 indexed before, uint256 indexed current)
func (_MarketImpl *MarketImplFilterer) ParseFeeBpsChanged(log types.Log) (*MarketImplFeeBpsChanged, error) {
	event := new(MarketImplFeeBpsChanged)
	if err := _MarketImpl.contract.UnpackLog(event, "FeeBpsChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketImplFeeCollectorChangedIterator is returned from FilterFeeCollectorChanged and is used to iterate over the raw logs and unpacked data for FeeCollectorChanged events raised by the MarketImpl contract.
type MarketImplFeeCollectorChangedIterator struct {
	Event *MarketImplFeeCollectorChanged // Event containing the contract specifics and raw log

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
func (it *MarketImplFeeCollectorChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketImplFeeCollectorChanged)
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
		it.Event = new(MarketImplFeeCollectorChanged)
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
func (it *MarketImplFeeCollectorChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketImplFeeCollectorChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketImplFeeCollectorChanged represents a FeeCollectorChanged event raised by the MarketImpl contract.
type MarketImplFeeCollectorChanged struct {
	Before  common.Address
	Current common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeeCollectorChanged is a free log retrieval operation binding the contract event 0x649c5e3d0ed183894196148e193af316452b0037e77d2ff0fef23b7dc722bed0.
//
// Solidity: event FeeCollectorChanged(address indexed before, address indexed current)
func (_MarketImpl *MarketImplFilterer) FilterFeeCollectorChanged(opts *bind.FilterOpts, before []common.Address, current []common.Address) (*MarketImplFeeCollectorChangedIterator, error) {

	var beforeRule []interface{}
	for _, beforeItem := range before {
		beforeRule = append(beforeRule, beforeItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _MarketImpl.contract.FilterLogs(opts, "FeeCollectorChanged", beforeRule, currentRule)
	if err != nil {
		return nil, err
	}
	return &MarketImplFeeCollectorChangedIterator{contract: _MarketImpl.contract, event: "FeeCollectorChanged", logs: logs, sub: sub}, nil
}

// WatchFeeCollectorChanged is a free log subscription operation binding the contract event 0x649c5e3d0ed183894196148e193af316452b0037e77d2ff0fef23b7dc722bed0.
//
// Solidity: event FeeCollectorChanged(address indexed before, address indexed current)
func (_MarketImpl *MarketImplFilterer) WatchFeeCollectorChanged(opts *bind.WatchOpts, sink chan<- *MarketImplFeeCollectorChanged, before []common.Address, current []common.Address) (event.Subscription, error) {

	var beforeRule []interface{}
	for _, beforeItem := range before {
		beforeRule = append(beforeRule, beforeItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _MarketImpl.contract.WatchLogs(opts, "FeeCollectorChanged", beforeRule, currentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketImplFeeCollectorChanged)
				if err := _MarketImpl.contract.UnpackLog(event, "FeeCollectorChanged", log); err != nil {
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
func (_MarketImpl *MarketImplFilterer) ParseFeeCollectorChanged(log types.Log) (*MarketImplFeeCollectorChanged, error) {
	event := new(MarketImplFeeCollectorChanged)
	if err := _MarketImpl.contract.UnpackLog(event, "FeeCollectorChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketImplInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the MarketImpl contract.
type MarketImplInitializedIterator struct {
	Event *MarketImplInitialized // Event containing the contract specifics and raw log

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
func (it *MarketImplInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketImplInitialized)
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
		it.Event = new(MarketImplInitialized)
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
func (it *MarketImplInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketImplInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketImplInitialized represents a Initialized event raised by the MarketImpl contract.
type MarketImplInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MarketImpl *MarketImplFilterer) FilterInitialized(opts *bind.FilterOpts) (*MarketImplInitializedIterator, error) {

	logs, sub, err := _MarketImpl.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MarketImplInitializedIterator{contract: _MarketImpl.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MarketImpl *MarketImplFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MarketImplInitialized) (event.Subscription, error) {

	logs, sub, err := _MarketImpl.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketImplInitialized)
				if err := _MarketImpl.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_MarketImpl *MarketImplFilterer) ParseInitialized(log types.Log) (*MarketImplInitialized, error) {
	event := new(MarketImplInitialized)
	if err := _MarketImpl.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketImplOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MarketImpl contract.
type MarketImplOwnershipTransferredIterator struct {
	Event *MarketImplOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MarketImplOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketImplOwnershipTransferred)
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
		it.Event = new(MarketImplOwnershipTransferred)
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
func (it *MarketImplOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketImplOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketImplOwnershipTransferred represents a OwnershipTransferred event raised by the MarketImpl contract.
type MarketImplOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MarketImpl *MarketImplFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MarketImplOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MarketImpl.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MarketImplOwnershipTransferredIterator{contract: _MarketImpl.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MarketImpl *MarketImplFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MarketImplOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MarketImpl.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketImplOwnershipTransferred)
				if err := _MarketImpl.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MarketImpl *MarketImplFilterer) ParseOwnershipTransferred(log types.Log) (*MarketImplOwnershipTransferred, error) {
	event := new(MarketImplOwnershipTransferred)
	if err := _MarketImpl.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketImplPairCreatedIterator is returned from FilterPairCreated and is used to iterate over the raw logs and unpacked data for PairCreated events raised by the MarketImpl contract.
type MarketImplPairCreatedIterator struct {
	Event *MarketImplPairCreated // Event containing the contract specifics and raw log

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
func (it *MarketImplPairCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketImplPairCreated)
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
		it.Event = new(MarketImplPairCreated)
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
func (it *MarketImplPairCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketImplPairCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketImplPairCreated represents a PairCreated event raised by the MarketImpl contract.
type MarketImplPairCreated struct {
	Pair      common.Address
	Base      common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPairCreated is a free log retrieval operation binding the contract event 0xc1db9ba7c4b7ce660fe8d17bbcf07167549381df2abd694a970bd1402d86d313.
//
// Solidity: event PairCreated(address indexed pair, address indexed base, uint256 timestamp)
func (_MarketImpl *MarketImplFilterer) FilterPairCreated(opts *bind.FilterOpts, pair []common.Address, base []common.Address) (*MarketImplPairCreatedIterator, error) {

	var pairRule []interface{}
	for _, pairItem := range pair {
		pairRule = append(pairRule, pairItem)
	}
	var baseRule []interface{}
	for _, baseItem := range base {
		baseRule = append(baseRule, baseItem)
	}

	logs, sub, err := _MarketImpl.contract.FilterLogs(opts, "PairCreated", pairRule, baseRule)
	if err != nil {
		return nil, err
	}
	return &MarketImplPairCreatedIterator{contract: _MarketImpl.contract, event: "PairCreated", logs: logs, sub: sub}, nil
}

// WatchPairCreated is a free log subscription operation binding the contract event 0xc1db9ba7c4b7ce660fe8d17bbcf07167549381df2abd694a970bd1402d86d313.
//
// Solidity: event PairCreated(address indexed pair, address indexed base, uint256 timestamp)
func (_MarketImpl *MarketImplFilterer) WatchPairCreated(opts *bind.WatchOpts, sink chan<- *MarketImplPairCreated, pair []common.Address, base []common.Address) (event.Subscription, error) {

	var pairRule []interface{}
	for _, pairItem := range pair {
		pairRule = append(pairRule, pairItem)
	}
	var baseRule []interface{}
	for _, baseItem := range base {
		baseRule = append(baseRule, baseItem)
	}

	logs, sub, err := _MarketImpl.contract.WatchLogs(opts, "PairCreated", pairRule, baseRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketImplPairCreated)
				if err := _MarketImpl.contract.UnpackLog(event, "PairCreated", log); err != nil {
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
func (_MarketImpl *MarketImplFilterer) ParsePairCreated(log types.Log) (*MarketImplPairCreated, error) {
	event := new(MarketImplPairCreated)
	if err := _MarketImpl.contract.UnpackLog(event, "PairCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketImplUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the MarketImpl contract.
type MarketImplUpgradedIterator struct {
	Event *MarketImplUpgraded // Event containing the contract specifics and raw log

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
func (it *MarketImplUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketImplUpgraded)
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
		it.Event = new(MarketImplUpgraded)
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
func (it *MarketImplUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketImplUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketImplUpgraded represents a Upgraded event raised by the MarketImpl contract.
type MarketImplUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_MarketImpl *MarketImplFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*MarketImplUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _MarketImpl.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &MarketImplUpgradedIterator{contract: _MarketImpl.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_MarketImpl *MarketImplFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *MarketImplUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _MarketImpl.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketImplUpgraded)
				if err := _MarketImpl.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_MarketImpl *MarketImplFilterer) ParseUpgraded(log types.Log) (*MarketImplUpgraded, error) {
	event := new(MarketImplUpgraded)
	if err := _MarketImpl.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
