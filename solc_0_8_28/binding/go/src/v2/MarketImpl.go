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

// MarketImplMetaData contains all meta data concerning the MarketImpl contract.
var MarketImplMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CROSS_DEX\",\"outputs\":[{\"internalType\":\"contractICrossDex\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"QUOTE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROUTER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allPairs\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"bases\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"pairs\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"}],\"name\":\"baseToPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"checkTickSizeRoles\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tickSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lotSize\",\"type\":\"uint256\"}],\"name\":\"createPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deployed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeBps\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeCollector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_quote\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pairImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeCollector\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_feeBPS\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pairImpl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_feeBps\",\"type\":\"uint256\"}],\"name\":\"setFeeBps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeCollector\",\"type\":\"address\"}],\"name\":\"setFeeCollector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"before\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"current\",\"type\":\"uint256\"}],\"name\":\"FeeBpsChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"before\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"FeeCollectorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"PairCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Create2EmptyBytecode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"EnumerableMapNonexistentKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedDeployment\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"MarketAlreadyCreatedBaseAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MarketDeployPair\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"MarketInvalidBaseAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"MarketInvalidInitializeData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"}]",
	ID:  "MarketImpl",
	Bin: "0x60a060405230608052348015610013575f5ffd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6080516124d26100f95f395f81816115030152818161152c015261174d01526124d25ff3fe608060405260043610610157575f3560e01c806395b6ef0c116100bb578063b820d8aa11610071578063c97682f811610057578063c97682f81461044d578063f2fde38b1461046f578063f905c15a1461048e575f5ffd5b8063b820d8aa146103f5578063c415b95c14610421575f5ffd5b8063a1eea778116100a1578063a1eea77814610362578063a42dce8014610381578063ad3cb1cc146103a0575f5ffd5b806395b6ef0c146103175780639c57983914610336575f5ffd5b80634f1ef28611610110578063715018a6116100f6578063715018a61461029b57806372c27b62146102af5780638da5cb5b146102ce575f5ffd5b80634f1ef2861461026457806352d1902d14610279575f5ffd5b806332fe7b261161014057806332fe7b26146101ed57806335f7d45614610219578063364e9a1914610245575f5ffd5b806324a9d8531461015b578063284cfbe0146101a9575b5f5ffd5b348015610166575f5ffd5b5060055461018f9074010000000000000000000000000000000000000000900463ffffffff1681565b60405163ffffffff90911681526020015b60405180910390f35b3480156101b4575f5ffd5b506101c86101c3366004611d2c565b6104a2565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101a0565b3480156101f8575f5ffd5b506003546101c89073ffffffffffffffffffffffffffffffffffffffff1681565b348015610224575f5ffd5b506004546101c89073ffffffffffffffffffffffffffffffffffffffff1681565b348015610250575f5ffd5b506101c861025f366004611d5c565b6109ca565b610277610272366004611da2565b6109dc565b005b348015610284575f5ffd5b5061028d6109fb565b6040519081526020016101a0565b3480156102a6575f5ffd5b50610277610a29565b3480156102ba575f5ffd5b506102776102c9366004611ea3565b610a3c565b3480156102d9575f5ffd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c1993005473ffffffffffffffffffffffffffffffffffffffff166101c8565b348015610322575f5ffd5b50610277610331366004611eba565b610b3f565b348015610341575f5ffd5b506002546101c89073ffffffffffffffffffffffffffffffffffffffff1681565b34801561036d575f5ffd5b5061027761037c366004611d5c565b61100a565b34801561038c575f5ffd5b5061027761039b366004611d5c565b61108c565b3480156103ab575f5ffd5b506103e86040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b6040516101a09190611f71565b348015610400575f5ffd5b506001546101c89073ffffffffffffffffffffffffffffffffffffffff1681565b34801561042c575f5ffd5b506005546101c89073ffffffffffffffffffffffffffffffffffffffff1681565b348015610458575f5ffd5b50610461611190565b6040516101a0929190611fd3565b34801561047a575f5ffd5b50610277610489366004611d5c565b61129f565b348015610499575f5ffd5b5061028d5f5481565b5f6104ab611302565b73ffffffffffffffffffffffffffffffffffffffff841615806104e8575060025473ffffffffffffffffffffffffffffffffffffffff8581169116145b1561053c576040517f20f8254c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024015b60405180910390fd5b5f8473ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa158015610586573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105aa9190611ff7565b60ff169050805f03610600576040517f20f8254c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff86166004820152602401610533565b61060b600686611390565b1561065a576040517fc13fcfdc00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff86166004820152602401610533565b5f6040518060200161066b90611cf7565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082820381018352601f90910116604081905260045460035460025473ffffffffffffffffffffffffffffffffffffffff91821660248501528116604484015289811660648401526084830189905260a48301889052169060c401604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa6b63eb800000000000000000000000000000000000000000000000000000000179052905161076b93929101612017565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290526107a7929160200161205c565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606089901b16602083015291505f906034016040516020818303038152906040528051906020012090505f61082d5f83856113b1565b905073ffffffffffffffffffffffffffffffffffffffff811661087c576040517f3ff6a26100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6108886006898361149d565b6108d6576040517fc13fcfdc00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff89166004820152602401610533565b6001546040517fe9f7206b00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83811660048301529091169063e9f7206b906024015f604051808303815f87803b158015610940575f5ffd5b505af1158015610952573d5f5f3e3d5ffd5b505050508773ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167fc1db9ba7c4b7ce660fe8d17bbcf07167549381df2abd694a970bd1402d86d313426040516109b591815260200190565b60405180910390a393505050505b9392505050565b5f6109d66006836114ca565b92915050565b6109e46114eb565b6109ed826115ef565b6109f782826115f7565b5050565b5f610a04611735565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b610a31611302565b610a3a5f6117a4565b565b610a44611302565b63ffffffff811115610aa4576040517fa25d2db50000000000000000000000000000000000000000000000000000000081527f66656542505300000000000000000000000000000000000000000000000000006004820152602401610533565b600554604051829174010000000000000000000000000000000000000000900463ffffffff16907ff878449785c2378ffd246b9c0c874a3aeec940d231ad03fe4a62a1cce095632d905f90a36005805463ffffffff90921674010000000000000000000000000000000000000000027fffffffffffffffff00000000ffffffffffffffffffffffffffffffffffffffff909216919091179055565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff165f81158015610b895750825b90505f8267ffffffffffffffff166001148015610ba55750303b155b905081158015610bb3575080155b15610bea576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001660011785558315610c4b5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b610c548b611839565b73ffffffffffffffffffffffffffffffffffffffff8b16610cc3576040517fa25d2db50000000000000000000000000000000000000000000000000000000081527f6f776e65720000000000000000000000000000000000000000000000000000006004820152602401610533565b73ffffffffffffffffffffffffffffffffffffffff8a16610d32576040517fa25d2db50000000000000000000000000000000000000000000000000000000081527f726f7574657200000000000000000000000000000000000000000000000000006004820152602401610533565b73ffffffffffffffffffffffffffffffffffffffff8916610da1576040517fa25d2db50000000000000000000000000000000000000000000000000000000081527f71756f74650000000000000000000000000000000000000000000000000000006004820152602401610533565b73ffffffffffffffffffffffffffffffffffffffff8816610e10576040517fa25d2db50000000000000000000000000000000000000000000000000000000081527f70616972496d706c0000000000000000000000000000000000000000000000006004820152602401610533565b73ffffffffffffffffffffffffffffffffffffffff8716610e7f576040517fa25d2db50000000000000000000000000000000000000000000000000000000081527f666565436f6c6c6563746f7200000000000000000000000000000000000000006004820152602401610533565b63ffffffff861115610edf576040517fa25d2db50000000000000000000000000000000000000000000000000000000081527f66656542505300000000000000000000000000000000000000000000000000006004820152602401610533565b435f5533600180547fffffffffffffffffffffffff000000000000000000000000000000000000000090811673ffffffffffffffffffffffffffffffffffffffff938416179091556002805482168c84161790556003805482168d8416179055600480549091168a8316179055600580549189167fffffffffffffffff000000000000000000000000000000000000000000000000909216919091177401000000000000000000000000000000000000000063ffffffff8916021790558315610ffd5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050505050565b6001546040517fa1eea77800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83811660048301529091169063a1eea778906024015f6040518083038186803b158015611073575f5ffd5b505afa158015611085573d5f5f3e3d5ffd5b5050505050565b611094611302565b73ffffffffffffffffffffffffffffffffffffffff8116611103576040517fa25d2db50000000000000000000000000000000000000000000000000000000081527f666565436f6c6c6563746f7200000000000000000000000000000000000000006004820152602401610533565b60055460405173ffffffffffffffffffffffffffffffffffffffff8084169216907f649c5e3d0ed183894196148e193af316452b0037e77d2ff0fef23b7dc722bed0905f90a3600580547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b6060805f61119e600661184a565b90508067ffffffffffffffff8111156111b9576111b9611d75565b6040519080825280602002602001820160405280156111e2578160200160208202803683370190505b5092508067ffffffffffffffff8111156111fe576111fe611d75565b604051908082528060200260200182016040528015611227578160200160208202803683370190505b5091505f5b818110156112995761123f600682611854565b85838151811061125157611251612070565b6020026020010185848151811061126a5761126a612070565b73ffffffffffffffffffffffffffffffffffffffff93841660209182029290920101529116905260010161122c565b50509091565b6112a7611302565b73ffffffffffffffffffffffffffffffffffffffff81166112f6576040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081525f6004820152602401610533565b6112ff816117a4565b50565b336113417f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c1993005473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614610a3a576040517f118cdaa7000000000000000000000000000000000000000000000000000000008152336004820152602401610533565b5f6109c38373ffffffffffffffffffffffffffffffffffffffff841661186f565b5f834710156113f5576040517fcf47918100000000000000000000000000000000000000000000000000000000815247600482015260248101859052604401610533565b81515f0361142f576040517f4ca249dc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8282516020840186f590503d151981151615611450576040513d5f823e3d81fd5b73ffffffffffffffffffffffffffffffffffffffff81166109c3576040517fb06ebf3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6114c28473ffffffffffffffffffffffffffffffffffffffff80861690851661187a565b949350505050565b5f6109c38373ffffffffffffffffffffffffffffffffffffffff8416611896565b3073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614806115b857507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1661159f7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614155b15610a3a576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6112ff611302565b8173ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa92505050801561167c575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682019092526116799181019061209d565b60015b6116ca576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610533565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8114611726576040517faa1d49a400000000000000000000000000000000000000000000000000000000815260048101829052602401610533565b61173083836118f5565b505050565b3073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610a3a576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080547fffffffffffffffffffffffff0000000000000000000000000000000000000000811673ffffffffffffffffffffffffffffffffffffffff848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b611841611957565b6112ff816119be565b5f6109d6826119c6565b5f80808061186286866119d0565b9097909650945050505050565b5f6109c383836119f9565b5f82815260028401602052604081208290556114c28484611a10565b5f818152600283016020526040812054801580156118bb57506118b9848461186f565b155b156109c3576040517f02b5668600000000000000000000000000000000000000000000000000000000815260048101849052602401610533565b6118fe82611a1b565b60405173ffffffffffffffffffffffffffffffffffffffff8316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a280511561194f576117308282611ae9565b6109f7611b68565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff16610a3a576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6112a7611957565b5f6109d682611ba0565b5f80806119dd8585611ba9565b5f81815260029690960160205260409095205494959350505050565b5f81815260018301602052604081205415156109c3565b5f6109c38383611bb4565b8073ffffffffffffffffffffffffffffffffffffffff163b5f03611a83576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152602401610533565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60605f5f8473ffffffffffffffffffffffffffffffffffffffff1684604051611b1291906120b4565b5f60405180830381855af49150503d805f8114611b4a576040519150601f19603f3d011682016040523d82523d5f602084013e611b4f565b606091505b5091509150611b5f858383611c00565b95945050505050565b3415610a3a576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6109d6825490565b5f6109c38383611c8f565b5f818152600183016020526040812054611bf957508154600181810184555f8481526020808220909301849055845484825282860190935260409020919091556109d6565b505f6109d6565b606082611c1557611c1082611cb5565b6109c3565b8151158015611c39575073ffffffffffffffffffffffffffffffffffffffff84163b155b15611c88576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610533565b50806109c3565b5f825f018281548110611ca457611ca4612070565b905f5260205f200154905092915050565b805115611cc55780518082602001fd5b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103dd806120c083390190565b803573ffffffffffffffffffffffffffffffffffffffff81168114611d27575f5ffd5b919050565b5f5f5f60608486031215611d3e575f5ffd5b611d4784611d04565b95602085013595506040909401359392505050565b5f60208284031215611d6c575f5ffd5b6109c382611d04565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f5f60408385031215611db3575f5ffd5b611dbc83611d04565b9150602083013567ffffffffffffffff811115611dd7575f5ffd5b8301601f81018513611de7575f5ffd5b803567ffffffffffffffff811115611e0157611e01611d75565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff82111715611e6d57611e6d611d75565b604052818152828201602001871015611e84575f5ffd5b816020840160208301375f602083830101528093505050509250929050565b5f60208284031215611eb3575f5ffd5b5035919050565b5f5f5f5f5f5f60c08789031215611ecf575f5ffd5b611ed887611d04565b9550611ee660208801611d04565b9450611ef460408801611d04565b9350611f0260608801611d04565b9250611f1060808801611d04565b9598949750929591949360a090920135925050565b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f6109c36020830184611f25565b5f8151808452602084019350602083015f5b82811015611fc957815173ffffffffffffffffffffffffffffffffffffffff16865260209586019590910190600101611f95565b5093949350505050565b604081525f611fe56040830185611f83565b8281036020840152611b5f8185611f83565b5f60208284031215612007575f5ffd5b815160ff811681146109c3575f5ffd5b73ffffffffffffffffffffffffffffffffffffffff83168152604060208201525f6114c26040830184611f25565b5f81518060208401855e5f93019283525090919050565b5f6114c261206a8386612045565b84612045565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f602082840312156120ad575f5ffd5b5051919050565b5f6109c3828461204556fe60806040526040516103dd3803806103dd8339810160408190526100229161023c565b61002c8282610033565b5050610321565b61003c82610091565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a280511561008557610080828261010c565b505050565b61008d61017f565b5050565b806001600160a01b03163b5f036100cb57604051634c9c8ce360e01b81526001600160a01b03821660048201526024015b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b031684604051610128919061030b565b5f60405180830381855af49150503d805f8114610160576040519150601f19603f3d011682016040523d82523d5f602084013e610165565b606091505b5090925090506101768583836101a0565b95945050505050565b341561019e5760405163b398979f60e01b815260040160405180910390fd5b565b6060826101b5576101b0826101ff565b6101f8565b81511580156101cc57506001600160a01b0384163b155b156101f557604051639996b31560e01b81526001600160a01b03851660048201526024016100c2565b50805b9392505050565b80511561020f5780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b634e487b7160e01b5f52604160045260245ffd5b5f5f6040838503121561024d575f5ffd5b82516001600160a01b0381168114610263575f5ffd5b60208401519092506001600160401b0381111561027e575f5ffd5b8301601f8101851361028e575f5ffd5b80516001600160401b038111156102a7576102a7610228565b604051601f8201601f19908116603f011681016001600160401b03811182821017156102d5576102d5610228565b6040528181528282016020018710156102ec575f5ffd5b8160208401602083015e5f602083830101528093505050509250929050565b5f82518060208501845e5f920191825250919050565b60b08061032d5f395ff3fe6080604052600a600c565b005b60186014601a565b605d565b565b5f60587f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5473ffffffffffffffffffffffffffffffffffffffff1690565b905090565b365f5f375f5f365f845af43d5f5f3e8080156076573d5ff35b3d5ffdfea264697066735822122053cea113bf96df2d97b042d29245e0a41b1290b530767f8b25bfeaed774391b764736f6c634300081c0033a264697066735822122032a1c4d5d6b56583e3b13942b91d9214a6524c38118245b2c86e033bcfb00baf64736f6c634300081c0033",
}

// MarketImpl is an auto generated Go binding around an Ethereum contract.
type MarketImpl struct {
	abi abi.ABI
}

// NewMarketImpl creates a new instance of MarketImpl.
func NewMarketImpl() *MarketImpl {
	parsed, err := MarketImplMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &MarketImpl{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *MarketImpl) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackCROSSDEX is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb820d8aa.
//
// Solidity: function CROSS_DEX() view returns(address)
func (marketImpl *MarketImpl) PackCROSSDEX() []byte {
	enc, err := marketImpl.abi.Pack("CROSS_DEX")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackCROSSDEX is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb820d8aa.
//
// Solidity: function CROSS_DEX() view returns(address)
func (marketImpl *MarketImpl) UnpackCROSSDEX(data []byte) (common.Address, error) {
	out, err := marketImpl.abi.Unpack("CROSS_DEX", data)
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
func (marketImpl *MarketImpl) PackQUOTE() []byte {
	enc, err := marketImpl.abi.Pack("QUOTE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackQUOTE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9c579839.
//
// Solidity: function QUOTE() view returns(address)
func (marketImpl *MarketImpl) UnpackQUOTE(data []byte) (common.Address, error) {
	out, err := marketImpl.abi.Unpack("QUOTE", data)
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
func (marketImpl *MarketImpl) PackROUTER() []byte {
	enc, err := marketImpl.abi.Pack("ROUTER")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackROUTER is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x32fe7b26.
//
// Solidity: function ROUTER() view returns(address)
func (marketImpl *MarketImpl) UnpackROUTER(data []byte) (common.Address, error) {
	out, err := marketImpl.abi.Unpack("ROUTER", data)
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
func (marketImpl *MarketImpl) PackUPGRADEINTERFACEVERSION() []byte {
	enc, err := marketImpl.abi.Pack("UPGRADE_INTERFACE_VERSION")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackUPGRADEINTERFACEVERSION is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (marketImpl *MarketImpl) UnpackUPGRADEINTERFACEVERSION(data []byte) (string, error) {
	out, err := marketImpl.abi.Unpack("UPGRADE_INTERFACE_VERSION", data)
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
func (marketImpl *MarketImpl) PackAllPairs() []byte {
	enc, err := marketImpl.abi.Pack("allPairs")
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
func (marketImpl *MarketImpl) UnpackAllPairs(data []byte) (AllPairsOutput, error) {
	out, err := marketImpl.abi.Unpack("allPairs", data)
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
func (marketImpl *MarketImpl) PackBaseToPair(base common.Address) []byte {
	enc, err := marketImpl.abi.Pack("baseToPair", base)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackBaseToPair is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x364e9a19.
//
// Solidity: function baseToPair(address base) view returns(address)
func (marketImpl *MarketImpl) UnpackBaseToPair(data []byte) (common.Address, error) {
	out, err := marketImpl.abi.Unpack("baseToPair", data)
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
func (marketImpl *MarketImpl) PackCheckTickSizeRoles(account common.Address) []byte {
	enc, err := marketImpl.abi.Pack("checkTickSizeRoles", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCreatePair is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x284cfbe0.
//
// Solidity: function createPair(address base, uint256 tickSize, uint256 lotSize) returns(address)
func (marketImpl *MarketImpl) PackCreatePair(base common.Address, tickSize *big.Int, lotSize *big.Int) []byte {
	enc, err := marketImpl.abi.Pack("createPair", base, tickSize, lotSize)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackCreatePair is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x284cfbe0.
//
// Solidity: function createPair(address base, uint256 tickSize, uint256 lotSize) returns(address)
func (marketImpl *MarketImpl) UnpackCreatePair(data []byte) (common.Address, error) {
	out, err := marketImpl.abi.Unpack("createPair", data)
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
func (marketImpl *MarketImpl) PackDeployed() []byte {
	enc, err := marketImpl.abi.Pack("deployed")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDeployed is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf905c15a.
//
// Solidity: function deployed() view returns(uint256)
func (marketImpl *MarketImpl) UnpackDeployed(data []byte) (*big.Int, error) {
	out, err := marketImpl.abi.Unpack("deployed", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackFeeBps is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x24a9d853.
//
// Solidity: function feeBps() view returns(uint32)
func (marketImpl *MarketImpl) PackFeeBps() []byte {
	enc, err := marketImpl.abi.Pack("feeBps")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackFeeBps is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x24a9d853.
//
// Solidity: function feeBps() view returns(uint32)
func (marketImpl *MarketImpl) UnpackFeeBps(data []byte) (uint32, error) {
	out, err := marketImpl.abi.Unpack("feeBps", data)
	if err != nil {
		return *new(uint32), err
	}
	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
	return out0, err
}

// PackFeeCollector is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (marketImpl *MarketImpl) PackFeeCollector() []byte {
	enc, err := marketImpl.abi.Pack("feeCollector")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackFeeCollector is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (marketImpl *MarketImpl) UnpackFeeCollector(data []byte) (common.Address, error) {
	out, err := marketImpl.abi.Unpack("feeCollector", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x95b6ef0c.
//
// Solidity: function initialize(address _owner, address _router, address _quote, address _pairImpl, address _feeCollector, uint256 _feeBPS) returns()
func (marketImpl *MarketImpl) PackInitialize(owner common.Address, router common.Address, quote common.Address, pairImpl common.Address, feeCollector common.Address, feeBPS *big.Int) []byte {
	enc, err := marketImpl.abi.Pack("initialize", owner, router, quote, pairImpl, feeCollector, feeBPS)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (marketImpl *MarketImpl) PackOwner() []byte {
	enc, err := marketImpl.abi.Pack("owner")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (marketImpl *MarketImpl) UnpackOwner(data []byte) (common.Address, error) {
	out, err := marketImpl.abi.Unpack("owner", data)
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
func (marketImpl *MarketImpl) PackPairImpl() []byte {
	enc, err := marketImpl.abi.Pack("pairImpl")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPairImpl is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x35f7d456.
//
// Solidity: function pairImpl() view returns(address)
func (marketImpl *MarketImpl) UnpackPairImpl(data []byte) (common.Address, error) {
	out, err := marketImpl.abi.Unpack("pairImpl", data)
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
func (marketImpl *MarketImpl) PackProxiableUUID() []byte {
	enc, err := marketImpl.abi.Pack("proxiableUUID")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackProxiableUUID is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (marketImpl *MarketImpl) UnpackProxiableUUID(data []byte) ([32]byte, error) {
	out, err := marketImpl.abi.Unpack("proxiableUUID", data)
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
func (marketImpl *MarketImpl) PackRenounceOwnership() []byte {
	enc, err := marketImpl.abi.Pack("renounceOwnership")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetFeeBps is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x72c27b62.
//
// Solidity: function setFeeBps(uint256 _feeBps) returns()
func (marketImpl *MarketImpl) PackSetFeeBps(feeBps *big.Int) []byte {
	enc, err := marketImpl.abi.Pack("setFeeBps", feeBps)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetFeeCollector is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa42dce80.
//
// Solidity: function setFeeCollector(address _feeCollector) returns()
func (marketImpl *MarketImpl) PackSetFeeCollector(feeCollector common.Address) []byte {
	enc, err := marketImpl.abi.Pack("setFeeCollector", feeCollector)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackTransferOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (marketImpl *MarketImpl) PackTransferOwnership(newOwner common.Address) []byte {
	enc, err := marketImpl.abi.Pack("transferOwnership", newOwner)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackUpgradeToAndCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (marketImpl *MarketImpl) PackUpgradeToAndCall(newImplementation common.Address, data []byte) []byte {
	enc, err := marketImpl.abi.Pack("upgradeToAndCall", newImplementation, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// MarketImplFeeBpsChanged represents a FeeBpsChanged event raised by the MarketImpl contract.
type MarketImplFeeBpsChanged struct {
	Before  *big.Int
	Current *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const MarketImplFeeBpsChangedEventName = "FeeBpsChanged"

// ContractEventName returns the user-defined event name.
func (MarketImplFeeBpsChanged) ContractEventName() string {
	return MarketImplFeeBpsChangedEventName
}

// UnpackFeeBpsChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event FeeBpsChanged(uint256 indexed before, uint256 indexed current)
func (marketImpl *MarketImpl) UnpackFeeBpsChangedEvent(log *types.Log) (*MarketImplFeeBpsChanged, error) {
	event := "FeeBpsChanged"
	if log.Topics[0] != marketImpl.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MarketImplFeeBpsChanged)
	if len(log.Data) > 0 {
		if err := marketImpl.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range marketImpl.abi.Events[event].Inputs {
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

// MarketImplFeeCollectorChanged represents a FeeCollectorChanged event raised by the MarketImpl contract.
type MarketImplFeeCollectorChanged struct {
	Before  common.Address
	Current common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const MarketImplFeeCollectorChangedEventName = "FeeCollectorChanged"

// ContractEventName returns the user-defined event name.
func (MarketImplFeeCollectorChanged) ContractEventName() string {
	return MarketImplFeeCollectorChangedEventName
}

// UnpackFeeCollectorChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event FeeCollectorChanged(address indexed before, address indexed current)
func (marketImpl *MarketImpl) UnpackFeeCollectorChangedEvent(log *types.Log) (*MarketImplFeeCollectorChanged, error) {
	event := "FeeCollectorChanged"
	if log.Topics[0] != marketImpl.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MarketImplFeeCollectorChanged)
	if len(log.Data) > 0 {
		if err := marketImpl.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range marketImpl.abi.Events[event].Inputs {
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

// MarketImplInitialized represents a Initialized event raised by the MarketImpl contract.
type MarketImplInitialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const MarketImplInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (MarketImplInitialized) ContractEventName() string {
	return MarketImplInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (marketImpl *MarketImpl) UnpackInitializedEvent(log *types.Log) (*MarketImplInitialized, error) {
	event := "Initialized"
	if log.Topics[0] != marketImpl.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MarketImplInitialized)
	if len(log.Data) > 0 {
		if err := marketImpl.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range marketImpl.abi.Events[event].Inputs {
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

// MarketImplOwnershipTransferred represents a OwnershipTransferred event raised by the MarketImpl contract.
type MarketImplOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const MarketImplOwnershipTransferredEventName = "OwnershipTransferred"

// ContractEventName returns the user-defined event name.
func (MarketImplOwnershipTransferred) ContractEventName() string {
	return MarketImplOwnershipTransferredEventName
}

// UnpackOwnershipTransferredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (marketImpl *MarketImpl) UnpackOwnershipTransferredEvent(log *types.Log) (*MarketImplOwnershipTransferred, error) {
	event := "OwnershipTransferred"
	if log.Topics[0] != marketImpl.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MarketImplOwnershipTransferred)
	if len(log.Data) > 0 {
		if err := marketImpl.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range marketImpl.abi.Events[event].Inputs {
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

// MarketImplPairCreated represents a PairCreated event raised by the MarketImpl contract.
type MarketImplPairCreated struct {
	Pair      common.Address
	Base      common.Address
	Timestamp *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const MarketImplPairCreatedEventName = "PairCreated"

// ContractEventName returns the user-defined event name.
func (MarketImplPairCreated) ContractEventName() string {
	return MarketImplPairCreatedEventName
}

// UnpackPairCreatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PairCreated(address indexed pair, address indexed base, uint256 timestamp)
func (marketImpl *MarketImpl) UnpackPairCreatedEvent(log *types.Log) (*MarketImplPairCreated, error) {
	event := "PairCreated"
	if log.Topics[0] != marketImpl.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MarketImplPairCreated)
	if len(log.Data) > 0 {
		if err := marketImpl.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range marketImpl.abi.Events[event].Inputs {
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

// MarketImplUpgraded represents a Upgraded event raised by the MarketImpl contract.
type MarketImplUpgraded struct {
	Implementation common.Address
	Raw            *types.Log // Blockchain specific contextual infos
}

const MarketImplUpgradedEventName = "Upgraded"

// ContractEventName returns the user-defined event name.
func (MarketImplUpgraded) ContractEventName() string {
	return MarketImplUpgradedEventName
}

// UnpackUpgradedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Upgraded(address indexed implementation)
func (marketImpl *MarketImpl) UnpackUpgradedEvent(log *types.Log) (*MarketImplUpgraded, error) {
	event := "Upgraded"
	if log.Topics[0] != marketImpl.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MarketImplUpgraded)
	if len(log.Data) > 0 {
		if err := marketImpl.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range marketImpl.abi.Events[event].Inputs {
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
func (marketImpl *MarketImpl) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], marketImpl.abi.Errors["AddressEmptyCode"].ID.Bytes()[:4]) {
		return marketImpl.UnpackAddressEmptyCodeError(raw[4:])
	}
	if bytes.Equal(raw[:4], marketImpl.abi.Errors["Create2EmptyBytecode"].ID.Bytes()[:4]) {
		return marketImpl.UnpackCreate2EmptyBytecodeError(raw[4:])
	}
	if bytes.Equal(raw[:4], marketImpl.abi.Errors["ERC1967InvalidImplementation"].ID.Bytes()[:4]) {
		return marketImpl.UnpackERC1967InvalidImplementationError(raw[4:])
	}
	if bytes.Equal(raw[:4], marketImpl.abi.Errors["ERC1967NonPayable"].ID.Bytes()[:4]) {
		return marketImpl.UnpackERC1967NonPayableError(raw[4:])
	}
	if bytes.Equal(raw[:4], marketImpl.abi.Errors["EnumerableMapNonexistentKey"].ID.Bytes()[:4]) {
		return marketImpl.UnpackEnumerableMapNonexistentKeyError(raw[4:])
	}
	if bytes.Equal(raw[:4], marketImpl.abi.Errors["FailedCall"].ID.Bytes()[:4]) {
		return marketImpl.UnpackFailedCallError(raw[4:])
	}
	if bytes.Equal(raw[:4], marketImpl.abi.Errors["FailedDeployment"].ID.Bytes()[:4]) {
		return marketImpl.UnpackFailedDeploymentError(raw[4:])
	}
	if bytes.Equal(raw[:4], marketImpl.abi.Errors["InsufficientBalance"].ID.Bytes()[:4]) {
		return marketImpl.UnpackInsufficientBalanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], marketImpl.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return marketImpl.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], marketImpl.abi.Errors["MarketAlreadyCreatedBaseAddress"].ID.Bytes()[:4]) {
		return marketImpl.UnpackMarketAlreadyCreatedBaseAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], marketImpl.abi.Errors["MarketDeployPair"].ID.Bytes()[:4]) {
		return marketImpl.UnpackMarketDeployPairError(raw[4:])
	}
	if bytes.Equal(raw[:4], marketImpl.abi.Errors["MarketInvalidBaseAddress"].ID.Bytes()[:4]) {
		return marketImpl.UnpackMarketInvalidBaseAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], marketImpl.abi.Errors["MarketInvalidInitializeData"].ID.Bytes()[:4]) {
		return marketImpl.UnpackMarketInvalidInitializeDataError(raw[4:])
	}
	if bytes.Equal(raw[:4], marketImpl.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return marketImpl.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], marketImpl.abi.Errors["OwnableInvalidOwner"].ID.Bytes()[:4]) {
		return marketImpl.UnpackOwnableInvalidOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], marketImpl.abi.Errors["OwnableUnauthorizedAccount"].ID.Bytes()[:4]) {
		return marketImpl.UnpackOwnableUnauthorizedAccountError(raw[4:])
	}
	if bytes.Equal(raw[:4], marketImpl.abi.Errors["UUPSUnauthorizedCallContext"].ID.Bytes()[:4]) {
		return marketImpl.UnpackUUPSUnauthorizedCallContextError(raw[4:])
	}
	if bytes.Equal(raw[:4], marketImpl.abi.Errors["UUPSUnsupportedProxiableUUID"].ID.Bytes()[:4]) {
		return marketImpl.UnpackUUPSUnsupportedProxiableUUIDError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// MarketImplAddressEmptyCode represents a AddressEmptyCode error raised by the MarketImpl contract.
type MarketImplAddressEmptyCode struct {
	Target common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressEmptyCode(address target)
func MarketImplAddressEmptyCodeErrorID() common.Hash {
	return common.HexToHash("0x9996b315c842ff135b8fc4a08ad5df1c344efbc03d2687aecc0678050d2aac89")
}

// UnpackAddressEmptyCodeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressEmptyCode(address target)
func (marketImpl *MarketImpl) UnpackAddressEmptyCodeError(raw []byte) (*MarketImplAddressEmptyCode, error) {
	out := new(MarketImplAddressEmptyCode)
	if err := marketImpl.abi.UnpackIntoInterface(out, "AddressEmptyCode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MarketImplCreate2EmptyBytecode represents a Create2EmptyBytecode error raised by the MarketImpl contract.
type MarketImplCreate2EmptyBytecode struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error Create2EmptyBytecode()
func MarketImplCreate2EmptyBytecodeErrorID() common.Hash {
	return common.HexToHash("0x4ca249dcffe41558ef8b961d71c905e4fa4317a1663f377b9610642e4e0abdb6")
}

// UnpackCreate2EmptyBytecodeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error Create2EmptyBytecode()
func (marketImpl *MarketImpl) UnpackCreate2EmptyBytecodeError(raw []byte) (*MarketImplCreate2EmptyBytecode, error) {
	out := new(MarketImplCreate2EmptyBytecode)
	if err := marketImpl.abi.UnpackIntoInterface(out, "Create2EmptyBytecode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MarketImplERC1967InvalidImplementation represents a ERC1967InvalidImplementation error raised by the MarketImpl contract.
type MarketImplERC1967InvalidImplementation struct {
	Implementation common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func MarketImplERC1967InvalidImplementationErrorID() common.Hash {
	return common.HexToHash("0x4c9c8ce3ceb3130f17f7cdba48d89b5b0129f266a8bac114e6e315a41879b617")
}

// UnpackERC1967InvalidImplementationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func (marketImpl *MarketImpl) UnpackERC1967InvalidImplementationError(raw []byte) (*MarketImplERC1967InvalidImplementation, error) {
	out := new(MarketImplERC1967InvalidImplementation)
	if err := marketImpl.abi.UnpackIntoInterface(out, "ERC1967InvalidImplementation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MarketImplERC1967NonPayable represents a ERC1967NonPayable error raised by the MarketImpl contract.
type MarketImplERC1967NonPayable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967NonPayable()
func MarketImplERC1967NonPayableErrorID() common.Hash {
	return common.HexToHash("0xb398979fa84f543c8e222f17890372c487baf85e062276c127fef521eea7224b")
}

// UnpackERC1967NonPayableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967NonPayable()
func (marketImpl *MarketImpl) UnpackERC1967NonPayableError(raw []byte) (*MarketImplERC1967NonPayable, error) {
	out := new(MarketImplERC1967NonPayable)
	if err := marketImpl.abi.UnpackIntoInterface(out, "ERC1967NonPayable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MarketImplEnumerableMapNonexistentKey represents a EnumerableMapNonexistentKey error raised by the MarketImpl contract.
type MarketImplEnumerableMapNonexistentKey struct {
	Key [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EnumerableMapNonexistentKey(bytes32 key)
func MarketImplEnumerableMapNonexistentKeyErrorID() common.Hash {
	return common.HexToHash("0x02b566865b1da2a2deb61443712fe7f812ffd7a1fce56446ff0fe3db9f3484ef")
}

// UnpackEnumerableMapNonexistentKeyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EnumerableMapNonexistentKey(bytes32 key)
func (marketImpl *MarketImpl) UnpackEnumerableMapNonexistentKeyError(raw []byte) (*MarketImplEnumerableMapNonexistentKey, error) {
	out := new(MarketImplEnumerableMapNonexistentKey)
	if err := marketImpl.abi.UnpackIntoInterface(out, "EnumerableMapNonexistentKey", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MarketImplFailedCall represents a FailedCall error raised by the MarketImpl contract.
type MarketImplFailedCall struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedCall()
func MarketImplFailedCallErrorID() common.Hash {
	return common.HexToHash("0xd6bda27508c0fb6d8a39b4b122878dab26f731a7d4e4abe711dd3731899052a4")
}

// UnpackFailedCallError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedCall()
func (marketImpl *MarketImpl) UnpackFailedCallError(raw []byte) (*MarketImplFailedCall, error) {
	out := new(MarketImplFailedCall)
	if err := marketImpl.abi.UnpackIntoInterface(out, "FailedCall", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MarketImplFailedDeployment represents a FailedDeployment error raised by the MarketImpl contract.
type MarketImplFailedDeployment struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedDeployment()
func MarketImplFailedDeploymentErrorID() common.Hash {
	return common.HexToHash("0xb06ebf3d5067824a3fe5d5ba19471e035a7de6c88dac362c77b162830a5b9093")
}

// UnpackFailedDeploymentError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedDeployment()
func (marketImpl *MarketImpl) UnpackFailedDeploymentError(raw []byte) (*MarketImplFailedDeployment, error) {
	out := new(MarketImplFailedDeployment)
	if err := marketImpl.abi.UnpackIntoInterface(out, "FailedDeployment", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MarketImplInsufficientBalance represents a InsufficientBalance error raised by the MarketImpl contract.
type MarketImplInsufficientBalance struct {
	Balance *big.Int
	Needed  *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InsufficientBalance(uint256 balance, uint256 needed)
func MarketImplInsufficientBalanceErrorID() common.Hash {
	return common.HexToHash("0xcf4791818fba6e019216eb4864093b4947f674afada5d305e57d598b641dad1d")
}

// UnpackInsufficientBalanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InsufficientBalance(uint256 balance, uint256 needed)
func (marketImpl *MarketImpl) UnpackInsufficientBalanceError(raw []byte) (*MarketImplInsufficientBalance, error) {
	out := new(MarketImplInsufficientBalance)
	if err := marketImpl.abi.UnpackIntoInterface(out, "InsufficientBalance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MarketImplInvalidInitialization represents a InvalidInitialization error raised by the MarketImpl contract.
type MarketImplInvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func MarketImplInvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (marketImpl *MarketImpl) UnpackInvalidInitializationError(raw []byte) (*MarketImplInvalidInitialization, error) {
	out := new(MarketImplInvalidInitialization)
	if err := marketImpl.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MarketImplMarketAlreadyCreatedBaseAddress represents a MarketAlreadyCreatedBaseAddress error raised by the MarketImpl contract.
type MarketImplMarketAlreadyCreatedBaseAddress struct {
	Arg0 common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error MarketAlreadyCreatedBaseAddress(address arg0)
func MarketImplMarketAlreadyCreatedBaseAddressErrorID() common.Hash {
	return common.HexToHash("0xc13fcfdc9c07ead596b870feafa1c7da90f68b06105d707eafb3632707684ca2")
}

// UnpackMarketAlreadyCreatedBaseAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error MarketAlreadyCreatedBaseAddress(address arg0)
func (marketImpl *MarketImpl) UnpackMarketAlreadyCreatedBaseAddressError(raw []byte) (*MarketImplMarketAlreadyCreatedBaseAddress, error) {
	out := new(MarketImplMarketAlreadyCreatedBaseAddress)
	if err := marketImpl.abi.UnpackIntoInterface(out, "MarketAlreadyCreatedBaseAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MarketImplMarketDeployPair represents a MarketDeployPair error raised by the MarketImpl contract.
type MarketImplMarketDeployPair struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error MarketDeployPair()
func MarketImplMarketDeployPairErrorID() common.Hash {
	return common.HexToHash("0x3ff6a261f592a9c9e4b2201a0b6349edfb07458d95b8d396f4b58be24bcdc6e2")
}

// UnpackMarketDeployPairError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error MarketDeployPair()
func (marketImpl *MarketImpl) UnpackMarketDeployPairError(raw []byte) (*MarketImplMarketDeployPair, error) {
	out := new(MarketImplMarketDeployPair)
	if err := marketImpl.abi.UnpackIntoInterface(out, "MarketDeployPair", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MarketImplMarketInvalidBaseAddress represents a MarketInvalidBaseAddress error raised by the MarketImpl contract.
type MarketImplMarketInvalidBaseAddress struct {
	Arg0 common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error MarketInvalidBaseAddress(address arg0)
func MarketImplMarketInvalidBaseAddressErrorID() common.Hash {
	return common.HexToHash("0x20f8254c2e2f68f995cc9b2605eef3477e0f8b491020c7f5bac08276c16a6c3c")
}

// UnpackMarketInvalidBaseAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error MarketInvalidBaseAddress(address arg0)
func (marketImpl *MarketImpl) UnpackMarketInvalidBaseAddressError(raw []byte) (*MarketImplMarketInvalidBaseAddress, error) {
	out := new(MarketImplMarketInvalidBaseAddress)
	if err := marketImpl.abi.UnpackIntoInterface(out, "MarketInvalidBaseAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MarketImplMarketInvalidInitializeData represents a MarketInvalidInitializeData error raised by the MarketImpl contract.
type MarketImplMarketInvalidInitializeData struct {
	Arg0 [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error MarketInvalidInitializeData(bytes32 arg0)
func MarketImplMarketInvalidInitializeDataErrorID() common.Hash {
	return common.HexToHash("0xa25d2db5f604033642dd6fbeddebf2802c569afa70e0e2031384d873de82c9c5")
}

// UnpackMarketInvalidInitializeDataError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error MarketInvalidInitializeData(bytes32 arg0)
func (marketImpl *MarketImpl) UnpackMarketInvalidInitializeDataError(raw []byte) (*MarketImplMarketInvalidInitializeData, error) {
	out := new(MarketImplMarketInvalidInitializeData)
	if err := marketImpl.abi.UnpackIntoInterface(out, "MarketInvalidInitializeData", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MarketImplNotInitializing represents a NotInitializing error raised by the MarketImpl contract.
type MarketImplNotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func MarketImplNotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (marketImpl *MarketImpl) UnpackNotInitializingError(raw []byte) (*MarketImplNotInitializing, error) {
	out := new(MarketImplNotInitializing)
	if err := marketImpl.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MarketImplOwnableInvalidOwner represents a OwnableInvalidOwner error raised by the MarketImpl contract.
type MarketImplOwnableInvalidOwner struct {
	Owner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OwnableInvalidOwner(address owner)
func MarketImplOwnableInvalidOwnerErrorID() common.Hash {
	return common.HexToHash("0x1e4fbdf7f3ef8bcaa855599e3abf48b232380f183f08f6f813d9ffa5bd585188")
}

// UnpackOwnableInvalidOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OwnableInvalidOwner(address owner)
func (marketImpl *MarketImpl) UnpackOwnableInvalidOwnerError(raw []byte) (*MarketImplOwnableInvalidOwner, error) {
	out := new(MarketImplOwnableInvalidOwner)
	if err := marketImpl.abi.UnpackIntoInterface(out, "OwnableInvalidOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MarketImplOwnableUnauthorizedAccount represents a OwnableUnauthorizedAccount error raised by the MarketImpl contract.
type MarketImplOwnableUnauthorizedAccount struct {
	Account common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OwnableUnauthorizedAccount(address account)
func MarketImplOwnableUnauthorizedAccountErrorID() common.Hash {
	return common.HexToHash("0x118cdaa7a341953d1887a2245fd6665d741c67c8c50581daa59e1d03373fa188")
}

// UnpackOwnableUnauthorizedAccountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OwnableUnauthorizedAccount(address account)
func (marketImpl *MarketImpl) UnpackOwnableUnauthorizedAccountError(raw []byte) (*MarketImplOwnableUnauthorizedAccount, error) {
	out := new(MarketImplOwnableUnauthorizedAccount)
	if err := marketImpl.abi.UnpackIntoInterface(out, "OwnableUnauthorizedAccount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MarketImplUUPSUnauthorizedCallContext represents a UUPSUnauthorizedCallContext error raised by the MarketImpl contract.
type MarketImplUUPSUnauthorizedCallContext struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnauthorizedCallContext()
func MarketImplUUPSUnauthorizedCallContextErrorID() common.Hash {
	return common.HexToHash("0xe07c8dba242a06571ac65fe4bbe20522c9fb111cb33599b799ff8039c1ed18f4")
}

// UnpackUUPSUnauthorizedCallContextError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnauthorizedCallContext()
func (marketImpl *MarketImpl) UnpackUUPSUnauthorizedCallContextError(raw []byte) (*MarketImplUUPSUnauthorizedCallContext, error) {
	out := new(MarketImplUUPSUnauthorizedCallContext)
	if err := marketImpl.abi.UnpackIntoInterface(out, "UUPSUnauthorizedCallContext", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MarketImplUUPSUnsupportedProxiableUUID represents a UUPSUnsupportedProxiableUUID error raised by the MarketImpl contract.
type MarketImplUUPSUnsupportedProxiableUUID struct {
	Slot [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func MarketImplUUPSUnsupportedProxiableUUIDErrorID() common.Hash {
	return common.HexToHash("0xaa1d49a4c084bfa9aeeee2a0be65267a7f19ba7e1476b114dac513d2c14cb563")
}

// UnpackUUPSUnsupportedProxiableUUIDError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func (marketImpl *MarketImpl) UnpackUUPSUnsupportedProxiableUUIDError(raw []byte) (*MarketImplUUPSUnsupportedProxiableUUID, error) {
	out := new(MarketImplUUPSUnsupportedProxiableUUID)
	if err := marketImpl.abi.UnpackIntoInterface(out, "UUPSUnsupportedProxiableUUID", raw); err != nil {
		return nil, err
	}
	return out, nil
}
