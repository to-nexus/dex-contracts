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

// CrossDexImplMetaData contains all meta data concerning the CrossDexImpl contract.
var CrossDexImplMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ROUTER\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allMarkets\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"quotes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"markets\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"checkTickSizeRoles\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_quote\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeCollector\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_feeBps\",\"type\":\"uint256\"}],\"name\":\"createMarket\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_routerImpl\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_findPrevPriceCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxMatchCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_cancelLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_marketImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pairImpl\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"}],\"name\":\"isMarket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketImpl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"name\":\"pairCreated\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pairImpl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"name\":\"pairToMarket\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"quote\",\"type\":\"address\"}],\"name\":\"quoteToMarket\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"setter\",\"type\":\"address\"}],\"name\":\"setTickSizeSetter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tickSizeSetter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"quote\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fee_collector\",\"type\":\"address\"}],\"name\":\"MarketCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"before\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"TickSizeSetterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"CrossDexAlreadyCreatedMarketQuote\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"CrossDexInitializeData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"CrossDexInvalidMarketAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"input\",\"type\":\"address\"}],\"name\":\"CrossDexInvalidTickSizeSetter\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"CrossDexUnauthorizedChangeTickSizes\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"EnumerableMapNonexistentKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"32fe7b26": "ROUTER()",
		"ad3cb1cc": "UPGRADE_INTERFACE_VERSION()",
		"375a7cba": "allMarkets()",
		"a1eea778": "checkTickSizeRoles(address)",
		"2e92c7be": "createMarket(address,address,address,uint256)",
		"0729da0b": "initialize(address,address,uint256,uint256,uint256,address,address)",
		"6ec934da": "isMarket(address)",
		"34eaeeb9": "marketImpl()",
		"8da5cb5b": "owner()",
		"e9f7206b": "pairCreated(address)",
		"35f7d456": "pairImpl()",
		"08270573": "pairToMarket(address)",
		"52d1902d": "proxiableUUID()",
		"beb380f1": "quoteToMarket(address)",
		"715018a6": "renounceOwnership()",
		"98715ee7": "setTickSizeSetter(address)",
		"7f764a85": "tickSizeSetter()",
		"f2fde38b": "transferOwnership(address)",
		"4f1ef286": "upgradeToAndCall(address,bytes)",
	},
	Bin: "0x60a060405230608052348015610013575f5ffd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6080516120296100f95f395f81816110c3015281816110ec015261130d01526120295ff3fe60806040526004361061013d575f3560e01c80636ec934da116100bb578063a1eea77811610071578063beb380f111610057578063beb380f114610411578063e9f7206b14610430578063f2fde38b1461044f575f5ffd5b8063a1eea7781461039d578063ad3cb1cc146103bc575f5ffd5b80637f764a85116100a15780637f764a85146103095780638da5cb5b1461033557806398715ee71461037e575f5ffd5b80636ec934da146102c6578063715018a6146102f5575f5ffd5b806334eaeeb911610110578063375a7cba116100f6578063375a7cba1461026f5780634f1ef2861461029157806352d1902d146102a4575f5ffd5b806334eaeeb91461021757806335f7d45614610243575f5ffd5b80630729da0b1461014157806308270573146101625780632e92c7be146101cd57806332fe7b26146101ec575b5f5ffd5b34801561014c575f5ffd5b5061016061015b3660046118cb565b61046e565b005b34801561016d575f5ffd5b506101a361017c366004611941565b60066020525f908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b3480156101d8575f5ffd5b506101a36101e736600461195c565b6108a6565b3480156101f7575f5ffd5b505f546101a39073ffffffffffffffffffffffffffffffffffffffff1681565b348015610222575f5ffd5b506001546101a39073ffffffffffffffffffffffffffffffffffffffff1681565b34801561024e575f5ffd5b506002546101a39073ffffffffffffffffffffffffffffffffffffffff1681565b34801561027a575f5ffd5b50610283610a84565b6040516101c49291906119fa565b61016061029f366004611a4b565b610b93565b3480156102af575f5ffd5b506102b8610bb2565b6040519081526020016101c4565b3480156102d1575f5ffd5b506102e56102e0366004611941565b610be0565b60405190151581526020016101c4565b348015610300575f5ffd5b50610160610d09565b348015610314575f5ffd5b506007546101a39073ffffffffffffffffffffffffffffffffffffffff1681565b348015610340575f5ffd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c1993005473ffffffffffffffffffffffffffffffffffffffff166101a3565b348015610389575f5ffd5b50610160610398366004611941565b610d1c565b3480156103a8575f5ffd5b506101606103b7366004611941565b610e49565b3480156103c7575f5ffd5b506104046040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b6040516101c49190611b4e565b34801561041c575f5ffd5b506101a361042b366004611941565b610eb8565b34801561043b575f5ffd5b5061016061044a366004611941565b610eca565b34801561045a575f5ffd5b50610160610469366004611941565b610f58565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff165f811580156104b85750825b90505f8267ffffffffffffffff1660011480156104d45750303b155b9050811580156104e2575080155b15610519576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000166001178555831561057a5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b73ffffffffffffffffffffffffffffffffffffffff8b166105ee576040517fc53a07b90000000000000000000000000000000000000000000000000000000081527f726f75746572496d706c0000000000000000000000000000000000000000000060048201526024015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff871661065d576040517fc53a07b90000000000000000000000000000000000000000000000000000000081527f6d61726b6574496d706c0000000000000000000000000000000000000000000060048201526024016105e5565b73ffffffffffffffffffffffffffffffffffffffff86166106cc576040517fc53a07b90000000000000000000000000000000000000000000000000000000081527f70616972496d706c00000000000000000000000000000000000000000000000060048201526024016105e5565b5f8b6040516106da9061189d565b73ffffffffffffffffffffffffffffffffffffffff90911681526040602082018190525f90820152606001604051809103905ff08015801561071e573d5f5f3e3d5ffd5b505f80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040517f80d85911000000000000000000000000000000000000000000000000000000008152600481018e9052602481018d9052604481018c9052919250906380d85911906064015f604051808303815f87803b1580156107c0575f5ffd5b505af11580156107d2573d5f5f3e3d5ffd5b50506001805473ffffffffffffffffffffffffffffffffffffffff808d167fffffffffffffffffffffffff00000000000000000000000000000000000000009283161790925560028054928c16929091169190911790555061083791508d9050610fb8565b83156108985784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050505050505050565b5f6108af610fc9565b6001546040515f9173ffffffffffffffffffffffffffffffffffffffff16906108d79061189d565b73ffffffffffffffffffffffffffffffffffffffff90911681526040602082018190525f90820152606001604051809103905ff08015801561091b573d5f5f3e3d5ffd5b505f546002546040517f95b6ef0c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8a81166004830152928316602482015288831660448201529082166064820152868216608482015260a481018690529192508216906395b6ef0c9060c4015f604051808303815f87803b1580156109b1575f5ffd5b505af11580156109c3573d5f5f3e3d5ffd5b508392506109d79150600390508783611057565b610a25576040517fed703a9800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff871660048201526024016105e5565b60405173ffffffffffffffffffffffffffffffffffffffff868116825280891691838216918916907f5509f6235f0d08c563aedf1b655226cad4a613794902c476ad7a08accb43ebc59060200160405180910390a49695505050505050565b6060805f610a926003611086565b90508067ffffffffffffffff811115610aad57610aad611a1e565b604051908082528060200260200182016040528015610ad6578160200160208202803683370190505b5092508067ffffffffffffffff811115610af257610af2611a1e565b604051908082528060200260200182016040528015610b1b578160200160208202803683370190505b5091505f5b81811015610b8d57610b33600382611090565b858381518110610b4557610b45611ba1565b60200260200101858481518110610b5e57610b5e611ba1565b73ffffffffffffffffffffffffffffffffffffffff938416602091820292909201015291169052600101610b20565b50509091565b610b9b6110ab565b610ba4826111af565b610bae82826111b7565b5050565b5f610bbb6112f5565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b60408051600481526024810182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f9c5798390000000000000000000000000000000000000000000000000000000017905290515f918291829173ffffffffffffffffffffffffffffffffffffffff861691610c609190611bce565b5f60405180830381855afa9150503d805f8114610c98576040519150601f19603f3d011682016040523d82523d5f602084013e610c9d565b606091505b509150915081610cb057505f9392505050565b5f81806020019051810190610cc59190611be4565b905073ffffffffffffffffffffffffffffffffffffffff8516610ce9600383611364565b73ffffffffffffffffffffffffffffffffffffffff161495945050505050565b610d11610fc9565b610d1a5f611385565b565b610d24610fc9565b73ffffffffffffffffffffffffffffffffffffffff81161580610d61575060075473ffffffffffffffffffffffffffffffffffffffff8281169116145b15610dbc576007546040517f18ddd67e00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff918216600482015290821660248201526044016105e5565b60075460405173ffffffffffffffffffffffffffffffffffffffff8084169216907f4f1656ebcdda14dfa0f9951b46485570075b9730e1f1d22cb4ddacc35374fde8905f90a3600780547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60075473ffffffffffffffffffffffffffffffffffffffff828116911614610eb5576040517ff42eaafb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024016105e5565b50565b5f610ec4600383611364565b92915050565b610ed333610be0565b610f0b576040517f74be5d0c0000000000000000000000000000000000000000000000000000000081523360048201526024016105e5565b73ffffffffffffffffffffffffffffffffffffffff165f90815260066020526040902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001633179055565b610f60610fc9565b73ffffffffffffffffffffffffffffffffffffffff8116610faf576040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081525f60048201526024016105e5565b610eb581611385565b610fc061141a565b610eb581611481565b336110087f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c1993005473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614610d1a576040517f118cdaa70000000000000000000000000000000000000000000000000000000081523360048201526024016105e5565b5f61107c8473ffffffffffffffffffffffffffffffffffffffff808616908516611489565b90505b9392505050565b5f610ec4826114a5565b5f80808061109e86866114af565b9097909650945050505050565b3073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016148061117857507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1661115f7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614155b15610d1a576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610eb5610fc9565b8173ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa92505050801561123c575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016820190925261123991810190611bff565b60015b61128a576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff831660048201526024016105e5565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc81146112e6576040517faa1d49a4000000000000000000000000000000000000000000000000000000008152600481018290526024016105e5565b6112f083836114d8565b505050565b3073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610d1a576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f61107f8373ffffffffffffffffffffffffffffffffffffffff841661153a565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080547fffffffffffffffffffffffff0000000000000000000000000000000000000000811673ffffffffffffffffffffffffffffffffffffffff848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff16610d1a576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610f6061141a565b5f828152600284016020526040812082905561107c8484611599565b5f610ec4826115a4565b5f80806114bc85856115ad565b5f81815260029690960160205260409095205494959350505050565b6114e1826115b8565b60405173ffffffffffffffffffffffffffffffffffffffff8316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115611532576112f08282611686565b610bae611705565b5f8181526002830160205260408120548015801561155f575061155d848461173d565b155b1561107f576040517f02b56686000000000000000000000000000000000000000000000000000000008152600481018490526024016105e5565b5f61107f838361175a565b5f610ec4825490565b5f61107f83836117a6565b8073ffffffffffffffffffffffffffffffffffffffff163b5f03611620576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024016105e5565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60605f5f8473ffffffffffffffffffffffffffffffffffffffff16846040516116af9190611bce565b5f60405180830381855af49150503d805f81146116e7576040519150601f19603f3d011682016040523d82523d5f602084013e6116ec565b606091505b50915091506116fc8583836117cc565b95945050505050565b3415610d1a576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f61107f83835f818152600183016020526040812054151561107f565b5f81815260018301602052604081205461179f57508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610ec4565b505f610ec4565b5f825f0182815481106117bb576117bb611ba1565b905f5260205f200154905092915050565b6060826117e1576117dc8261185b565b61107f565b8151158015611805575073ffffffffffffffffffffffffffffffffffffffff84163b155b15611854576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024016105e5565b508061107f565b80511561186b5780518082602001fd5b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103dd80611c1783390190565b73ffffffffffffffffffffffffffffffffffffffff81168114610eb5575f5ffd5b5f5f5f5f5f5f5f60e0888a0312156118e1575f5ffd5b87356118ec816118aa565b965060208801356118fc816118aa565b955060408801359450606088013593506080880135925060a0880135611921816118aa565b915060c0880135611931816118aa565b8091505092959891949750929550565b5f60208284031215611951575f5ffd5b813561107f816118aa565b5f5f5f5f6080858703121561196f575f5ffd5b843561197a816118aa565b9350602085013561198a816118aa565b9250604085013561199a816118aa565b9396929550929360600135925050565b5f8151808452602084019350602083015f5b828110156119f057815173ffffffffffffffffffffffffffffffffffffffff168652602095860195909101906001016119bc565b5093949350505050565b604081525f611a0c60408301856119aa565b82810360208401526116fc81856119aa565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f5f60408385031215611a5c575f5ffd5b8235611a67816118aa565b9150602083013567ffffffffffffffff811115611a82575f5ffd5b8301601f81018513611a92575f5ffd5b803567ffffffffffffffff811115611aac57611aac611a1e565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff82111715611b1857611b18611a1e565b604052818152828201602001871015611b2f575f5ffd5b816020840160208301375f602083830101528093505050509250929050565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f82518060208501845e5f920191825250919050565b5f60208284031215611bf4575f5ffd5b815161107f816118aa565b5f60208284031215611c0f575f5ffd5b505191905056fe60806040526040516103dd3803806103dd8339810160408190526100229161023c565b61002c8282610033565b5050610321565b61003c82610091565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a280511561008557610080828261010c565b505050565b61008d61017f565b5050565b806001600160a01b03163b5f036100cb57604051634c9c8ce360e01b81526001600160a01b03821660048201526024015b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b031684604051610128919061030b565b5f60405180830381855af49150503d805f8114610160576040519150601f19603f3d011682016040523d82523d5f602084013e610165565b606091505b5090925090506101768583836101a0565b95945050505050565b341561019e5760405163b398979f60e01b815260040160405180910390fd5b565b6060826101b5576101b0826101ff565b6101f8565b81511580156101cc57506001600160a01b0384163b155b156101f557604051639996b31560e01b81526001600160a01b03851660048201526024016100c2565b50805b9392505050565b80511561020f5780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b634e487b7160e01b5f52604160045260245ffd5b5f5f6040838503121561024d575f5ffd5b82516001600160a01b0381168114610263575f5ffd5b60208401519092506001600160401b0381111561027e575f5ffd5b8301601f8101851361028e575f5ffd5b80516001600160401b038111156102a7576102a7610228565b604051601f8201601f19908116603f011681016001600160401b03811182821017156102d5576102d5610228565b6040528181528282016020018710156102ec575f5ffd5b8160208401602083015e5f602083830101528093505050509250929050565b5f82518060208501845e5f920191825250919050565b60b08061032d5f395ff3fe6080604052600a600c565b005b60186014601a565b605d565b565b5f60587f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5473ffffffffffffffffffffffffffffffffffffffff1690565b905090565b365f5f375f5f365f845af43d5f5f3e8080156076573d5ff35b3d5ffdfea26469706673582212209fd71e6f878ca8212eaa80602ace1507d0020bad3381eef3782249983b65738964736f6c634300081c0033a26469706673582212204d70e63831a13e9eccd4a937c0a945622d0d0c32a45a29b6c55bf0a6b336b0ac64736f6c634300081c0033",
}

// CrossDexImplABI is the input ABI used to generate the binding from.
// Deprecated: Use CrossDexImplMetaData.ABI instead.
var CrossDexImplABI = CrossDexImplMetaData.ABI

// Deprecated: Use CrossDexImplMetaData.Sigs instead.
// CrossDexImplFuncSigs maps the 4-byte function signature to its string representation.
var CrossDexImplFuncSigs = CrossDexImplMetaData.Sigs

// CrossDexImplBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CrossDexImplMetaData.Bin instead.
var CrossDexImplBin = CrossDexImplMetaData.Bin

// DeployCrossDexImpl deploys a new Ethereum contract, binding an instance of CrossDexImpl to it.
func DeployCrossDexImpl(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CrossDexImpl, error) {
	parsed, err := CrossDexImplMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CrossDexImplBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CrossDexImpl{CrossDexImplCaller: CrossDexImplCaller{contract: contract}, CrossDexImplTransactor: CrossDexImplTransactor{contract: contract}, CrossDexImplFilterer: CrossDexImplFilterer{contract: contract}}, nil
}

// CrossDexImpl is an auto generated Go binding around an Ethereum contract.
type CrossDexImpl struct {
	CrossDexImplCaller     // Read-only binding to the contract
	CrossDexImplTransactor // Write-only binding to the contract
	CrossDexImplFilterer   // Log filterer for contract events
}

// CrossDexImplCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrossDexImplCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossDexImplTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrossDexImplTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossDexImplFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrossDexImplFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossDexImplSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrossDexImplSession struct {
	Contract     *CrossDexImpl     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrossDexImplCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrossDexImplCallerSession struct {
	Contract *CrossDexImplCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// CrossDexImplTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrossDexImplTransactorSession struct {
	Contract     *CrossDexImplTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CrossDexImplRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrossDexImplRaw struct {
	Contract *CrossDexImpl // Generic contract binding to access the raw methods on
}

// CrossDexImplCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrossDexImplCallerRaw struct {
	Contract *CrossDexImplCaller // Generic read-only contract binding to access the raw methods on
}

// CrossDexImplTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrossDexImplTransactorRaw struct {
	Contract *CrossDexImplTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrossDexImpl creates a new instance of CrossDexImpl, bound to a specific deployed contract.
func NewCrossDexImpl(address common.Address, backend bind.ContractBackend) (*CrossDexImpl, error) {
	contract, err := bindCrossDexImpl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrossDexImpl{CrossDexImplCaller: CrossDexImplCaller{contract: contract}, CrossDexImplTransactor: CrossDexImplTransactor{contract: contract}, CrossDexImplFilterer: CrossDexImplFilterer{contract: contract}}, nil
}

// NewCrossDexImplCaller creates a new read-only instance of CrossDexImpl, bound to a specific deployed contract.
func NewCrossDexImplCaller(address common.Address, caller bind.ContractCaller) (*CrossDexImplCaller, error) {
	contract, err := bindCrossDexImpl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrossDexImplCaller{contract: contract}, nil
}

// NewCrossDexImplTransactor creates a new write-only instance of CrossDexImpl, bound to a specific deployed contract.
func NewCrossDexImplTransactor(address common.Address, transactor bind.ContractTransactor) (*CrossDexImplTransactor, error) {
	contract, err := bindCrossDexImpl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrossDexImplTransactor{contract: contract}, nil
}

// NewCrossDexImplFilterer creates a new log filterer instance of CrossDexImpl, bound to a specific deployed contract.
func NewCrossDexImplFilterer(address common.Address, filterer bind.ContractFilterer) (*CrossDexImplFilterer, error) {
	contract, err := bindCrossDexImpl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrossDexImplFilterer{contract: contract}, nil
}

// bindCrossDexImpl binds a generic wrapper to an already deployed contract.
func bindCrossDexImpl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CrossDexImplMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossDexImpl *CrossDexImplRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossDexImpl.Contract.CrossDexImplCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossDexImpl *CrossDexImplRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossDexImpl.Contract.CrossDexImplTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossDexImpl *CrossDexImplRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossDexImpl.Contract.CrossDexImplTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossDexImpl *CrossDexImplCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossDexImpl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossDexImpl *CrossDexImplTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossDexImpl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossDexImpl *CrossDexImplTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossDexImpl.Contract.contract.Transact(opts, method, params...)
}

// ROUTER is a free data retrieval call binding the contract method 0x32fe7b26.
//
// Solidity: function ROUTER() view returns(address)
func (_CrossDexImpl *CrossDexImplCaller) ROUTER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossDexImpl.contract.Call(opts, &out, "ROUTER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ROUTER is a free data retrieval call binding the contract method 0x32fe7b26.
//
// Solidity: function ROUTER() view returns(address)
func (_CrossDexImpl *CrossDexImplSession) ROUTER() (common.Address, error) {
	return _CrossDexImpl.Contract.ROUTER(&_CrossDexImpl.CallOpts)
}

// ROUTER is a free data retrieval call binding the contract method 0x32fe7b26.
//
// Solidity: function ROUTER() view returns(address)
func (_CrossDexImpl *CrossDexImplCallerSession) ROUTER() (common.Address, error) {
	return _CrossDexImpl.Contract.ROUTER(&_CrossDexImpl.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_CrossDexImpl *CrossDexImplCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CrossDexImpl.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_CrossDexImpl *CrossDexImplSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _CrossDexImpl.Contract.UPGRADEINTERFACEVERSION(&_CrossDexImpl.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_CrossDexImpl *CrossDexImplCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _CrossDexImpl.Contract.UPGRADEINTERFACEVERSION(&_CrossDexImpl.CallOpts)
}

// AllMarkets is a free data retrieval call binding the contract method 0x375a7cba.
//
// Solidity: function allMarkets() view returns(address[] quotes, address[] markets)
func (_CrossDexImpl *CrossDexImplCaller) AllMarkets(opts *bind.CallOpts) (struct {
	Quotes  []common.Address
	Markets []common.Address
}, error) {
	var out []interface{}
	err := _CrossDexImpl.contract.Call(opts, &out, "allMarkets")

	outstruct := new(struct {
		Quotes  []common.Address
		Markets []common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Quotes = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Markets = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

// AllMarkets is a free data retrieval call binding the contract method 0x375a7cba.
//
// Solidity: function allMarkets() view returns(address[] quotes, address[] markets)
func (_CrossDexImpl *CrossDexImplSession) AllMarkets() (struct {
	Quotes  []common.Address
	Markets []common.Address
}, error) {
	return _CrossDexImpl.Contract.AllMarkets(&_CrossDexImpl.CallOpts)
}

// AllMarkets is a free data retrieval call binding the contract method 0x375a7cba.
//
// Solidity: function allMarkets() view returns(address[] quotes, address[] markets)
func (_CrossDexImpl *CrossDexImplCallerSession) AllMarkets() (struct {
	Quotes  []common.Address
	Markets []common.Address
}, error) {
	return _CrossDexImpl.Contract.AllMarkets(&_CrossDexImpl.CallOpts)
}

// CheckTickSizeRoles is a free data retrieval call binding the contract method 0xa1eea778.
//
// Solidity: function checkTickSizeRoles(address account) view returns()
func (_CrossDexImpl *CrossDexImplCaller) CheckTickSizeRoles(opts *bind.CallOpts, account common.Address) error {
	var out []interface{}
	err := _CrossDexImpl.contract.Call(opts, &out, "checkTickSizeRoles", account)

	if err != nil {
		return err
	}

	return err

}

// CheckTickSizeRoles is a free data retrieval call binding the contract method 0xa1eea778.
//
// Solidity: function checkTickSizeRoles(address account) view returns()
func (_CrossDexImpl *CrossDexImplSession) CheckTickSizeRoles(account common.Address) error {
	return _CrossDexImpl.Contract.CheckTickSizeRoles(&_CrossDexImpl.CallOpts, account)
}

// CheckTickSizeRoles is a free data retrieval call binding the contract method 0xa1eea778.
//
// Solidity: function checkTickSizeRoles(address account) view returns()
func (_CrossDexImpl *CrossDexImplCallerSession) CheckTickSizeRoles(account common.Address) error {
	return _CrossDexImpl.Contract.CheckTickSizeRoles(&_CrossDexImpl.CallOpts, account)
}

// IsMarket is a free data retrieval call binding the contract method 0x6ec934da.
//
// Solidity: function isMarket(address market) view returns(bool)
func (_CrossDexImpl *CrossDexImplCaller) IsMarket(opts *bind.CallOpts, market common.Address) (bool, error) {
	var out []interface{}
	err := _CrossDexImpl.contract.Call(opts, &out, "isMarket", market)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMarket is a free data retrieval call binding the contract method 0x6ec934da.
//
// Solidity: function isMarket(address market) view returns(bool)
func (_CrossDexImpl *CrossDexImplSession) IsMarket(market common.Address) (bool, error) {
	return _CrossDexImpl.Contract.IsMarket(&_CrossDexImpl.CallOpts, market)
}

// IsMarket is a free data retrieval call binding the contract method 0x6ec934da.
//
// Solidity: function isMarket(address market) view returns(bool)
func (_CrossDexImpl *CrossDexImplCallerSession) IsMarket(market common.Address) (bool, error) {
	return _CrossDexImpl.Contract.IsMarket(&_CrossDexImpl.CallOpts, market)
}

// MarketImpl is a free data retrieval call binding the contract method 0x34eaeeb9.
//
// Solidity: function marketImpl() view returns(address)
func (_CrossDexImpl *CrossDexImplCaller) MarketImpl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossDexImpl.contract.Call(opts, &out, "marketImpl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MarketImpl is a free data retrieval call binding the contract method 0x34eaeeb9.
//
// Solidity: function marketImpl() view returns(address)
func (_CrossDexImpl *CrossDexImplSession) MarketImpl() (common.Address, error) {
	return _CrossDexImpl.Contract.MarketImpl(&_CrossDexImpl.CallOpts)
}

// MarketImpl is a free data retrieval call binding the contract method 0x34eaeeb9.
//
// Solidity: function marketImpl() view returns(address)
func (_CrossDexImpl *CrossDexImplCallerSession) MarketImpl() (common.Address, error) {
	return _CrossDexImpl.Contract.MarketImpl(&_CrossDexImpl.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CrossDexImpl *CrossDexImplCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossDexImpl.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CrossDexImpl *CrossDexImplSession) Owner() (common.Address, error) {
	return _CrossDexImpl.Contract.Owner(&_CrossDexImpl.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CrossDexImpl *CrossDexImplCallerSession) Owner() (common.Address, error) {
	return _CrossDexImpl.Contract.Owner(&_CrossDexImpl.CallOpts)
}

// PairImpl is a free data retrieval call binding the contract method 0x35f7d456.
//
// Solidity: function pairImpl() view returns(address)
func (_CrossDexImpl *CrossDexImplCaller) PairImpl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossDexImpl.contract.Call(opts, &out, "pairImpl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PairImpl is a free data retrieval call binding the contract method 0x35f7d456.
//
// Solidity: function pairImpl() view returns(address)
func (_CrossDexImpl *CrossDexImplSession) PairImpl() (common.Address, error) {
	return _CrossDexImpl.Contract.PairImpl(&_CrossDexImpl.CallOpts)
}

// PairImpl is a free data retrieval call binding the contract method 0x35f7d456.
//
// Solidity: function pairImpl() view returns(address)
func (_CrossDexImpl *CrossDexImplCallerSession) PairImpl() (common.Address, error) {
	return _CrossDexImpl.Contract.PairImpl(&_CrossDexImpl.CallOpts)
}

// PairToMarket is a free data retrieval call binding the contract method 0x08270573.
//
// Solidity: function pairToMarket(address pair) view returns(address)
func (_CrossDexImpl *CrossDexImplCaller) PairToMarket(opts *bind.CallOpts, pair common.Address) (common.Address, error) {
	var out []interface{}
	err := _CrossDexImpl.contract.Call(opts, &out, "pairToMarket", pair)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PairToMarket is a free data retrieval call binding the contract method 0x08270573.
//
// Solidity: function pairToMarket(address pair) view returns(address)
func (_CrossDexImpl *CrossDexImplSession) PairToMarket(pair common.Address) (common.Address, error) {
	return _CrossDexImpl.Contract.PairToMarket(&_CrossDexImpl.CallOpts, pair)
}

// PairToMarket is a free data retrieval call binding the contract method 0x08270573.
//
// Solidity: function pairToMarket(address pair) view returns(address)
func (_CrossDexImpl *CrossDexImplCallerSession) PairToMarket(pair common.Address) (common.Address, error) {
	return _CrossDexImpl.Contract.PairToMarket(&_CrossDexImpl.CallOpts, pair)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CrossDexImpl *CrossDexImplCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CrossDexImpl.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CrossDexImpl *CrossDexImplSession) ProxiableUUID() ([32]byte, error) {
	return _CrossDexImpl.Contract.ProxiableUUID(&_CrossDexImpl.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CrossDexImpl *CrossDexImplCallerSession) ProxiableUUID() ([32]byte, error) {
	return _CrossDexImpl.Contract.ProxiableUUID(&_CrossDexImpl.CallOpts)
}

// QuoteToMarket is a free data retrieval call binding the contract method 0xbeb380f1.
//
// Solidity: function quoteToMarket(address quote) view returns(address)
func (_CrossDexImpl *CrossDexImplCaller) QuoteToMarket(opts *bind.CallOpts, quote common.Address) (common.Address, error) {
	var out []interface{}
	err := _CrossDexImpl.contract.Call(opts, &out, "quoteToMarket", quote)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// QuoteToMarket is a free data retrieval call binding the contract method 0xbeb380f1.
//
// Solidity: function quoteToMarket(address quote) view returns(address)
func (_CrossDexImpl *CrossDexImplSession) QuoteToMarket(quote common.Address) (common.Address, error) {
	return _CrossDexImpl.Contract.QuoteToMarket(&_CrossDexImpl.CallOpts, quote)
}

// QuoteToMarket is a free data retrieval call binding the contract method 0xbeb380f1.
//
// Solidity: function quoteToMarket(address quote) view returns(address)
func (_CrossDexImpl *CrossDexImplCallerSession) QuoteToMarket(quote common.Address) (common.Address, error) {
	return _CrossDexImpl.Contract.QuoteToMarket(&_CrossDexImpl.CallOpts, quote)
}

// TickSizeSetter is a free data retrieval call binding the contract method 0x7f764a85.
//
// Solidity: function tickSizeSetter() view returns(address)
func (_CrossDexImpl *CrossDexImplCaller) TickSizeSetter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossDexImpl.contract.Call(opts, &out, "tickSizeSetter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TickSizeSetter is a free data retrieval call binding the contract method 0x7f764a85.
//
// Solidity: function tickSizeSetter() view returns(address)
func (_CrossDexImpl *CrossDexImplSession) TickSizeSetter() (common.Address, error) {
	return _CrossDexImpl.Contract.TickSizeSetter(&_CrossDexImpl.CallOpts)
}

// TickSizeSetter is a free data retrieval call binding the contract method 0x7f764a85.
//
// Solidity: function tickSizeSetter() view returns(address)
func (_CrossDexImpl *CrossDexImplCallerSession) TickSizeSetter() (common.Address, error) {
	return _CrossDexImpl.Contract.TickSizeSetter(&_CrossDexImpl.CallOpts)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x2e92c7be.
//
// Solidity: function createMarket(address _owner, address _quote, address _feeCollector, uint256 _feeBps) returns(address)
func (_CrossDexImpl *CrossDexImplTransactor) CreateMarket(opts *bind.TransactOpts, _owner common.Address, _quote common.Address, _feeCollector common.Address, _feeBps *big.Int) (*types.Transaction, error) {
	return _CrossDexImpl.contract.Transact(opts, "createMarket", _owner, _quote, _feeCollector, _feeBps)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x2e92c7be.
//
// Solidity: function createMarket(address _owner, address _quote, address _feeCollector, uint256 _feeBps) returns(address)
func (_CrossDexImpl *CrossDexImplSession) CreateMarket(_owner common.Address, _quote common.Address, _feeCollector common.Address, _feeBps *big.Int) (*types.Transaction, error) {
	return _CrossDexImpl.Contract.CreateMarket(&_CrossDexImpl.TransactOpts, _owner, _quote, _feeCollector, _feeBps)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x2e92c7be.
//
// Solidity: function createMarket(address _owner, address _quote, address _feeCollector, uint256 _feeBps) returns(address)
func (_CrossDexImpl *CrossDexImplTransactorSession) CreateMarket(_owner common.Address, _quote common.Address, _feeCollector common.Address, _feeBps *big.Int) (*types.Transaction, error) {
	return _CrossDexImpl.Contract.CreateMarket(&_CrossDexImpl.TransactOpts, _owner, _quote, _feeCollector, _feeBps)
}

// Initialize is a paid mutator transaction binding the contract method 0x0729da0b.
//
// Solidity: function initialize(address _owner, address _routerImpl, uint256 _findPrevPriceCount, uint256 _maxMatchCount, uint256 _cancelLimit, address _marketImpl, address _pairImpl) returns()
func (_CrossDexImpl *CrossDexImplTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _routerImpl common.Address, _findPrevPriceCount *big.Int, _maxMatchCount *big.Int, _cancelLimit *big.Int, _marketImpl common.Address, _pairImpl common.Address) (*types.Transaction, error) {
	return _CrossDexImpl.contract.Transact(opts, "initialize", _owner, _routerImpl, _findPrevPriceCount, _maxMatchCount, _cancelLimit, _marketImpl, _pairImpl)
}

// Initialize is a paid mutator transaction binding the contract method 0x0729da0b.
//
// Solidity: function initialize(address _owner, address _routerImpl, uint256 _findPrevPriceCount, uint256 _maxMatchCount, uint256 _cancelLimit, address _marketImpl, address _pairImpl) returns()
func (_CrossDexImpl *CrossDexImplSession) Initialize(_owner common.Address, _routerImpl common.Address, _findPrevPriceCount *big.Int, _maxMatchCount *big.Int, _cancelLimit *big.Int, _marketImpl common.Address, _pairImpl common.Address) (*types.Transaction, error) {
	return _CrossDexImpl.Contract.Initialize(&_CrossDexImpl.TransactOpts, _owner, _routerImpl, _findPrevPriceCount, _maxMatchCount, _cancelLimit, _marketImpl, _pairImpl)
}

// Initialize is a paid mutator transaction binding the contract method 0x0729da0b.
//
// Solidity: function initialize(address _owner, address _routerImpl, uint256 _findPrevPriceCount, uint256 _maxMatchCount, uint256 _cancelLimit, address _marketImpl, address _pairImpl) returns()
func (_CrossDexImpl *CrossDexImplTransactorSession) Initialize(_owner common.Address, _routerImpl common.Address, _findPrevPriceCount *big.Int, _maxMatchCount *big.Int, _cancelLimit *big.Int, _marketImpl common.Address, _pairImpl common.Address) (*types.Transaction, error) {
	return _CrossDexImpl.Contract.Initialize(&_CrossDexImpl.TransactOpts, _owner, _routerImpl, _findPrevPriceCount, _maxMatchCount, _cancelLimit, _marketImpl, _pairImpl)
}

// PairCreated is a paid mutator transaction binding the contract method 0xe9f7206b.
//
// Solidity: function pairCreated(address pair) returns()
func (_CrossDexImpl *CrossDexImplTransactor) PairCreated(opts *bind.TransactOpts, pair common.Address) (*types.Transaction, error) {
	return _CrossDexImpl.contract.Transact(opts, "pairCreated", pair)
}

// PairCreated is a paid mutator transaction binding the contract method 0xe9f7206b.
//
// Solidity: function pairCreated(address pair) returns()
func (_CrossDexImpl *CrossDexImplSession) PairCreated(pair common.Address) (*types.Transaction, error) {
	return _CrossDexImpl.Contract.PairCreated(&_CrossDexImpl.TransactOpts, pair)
}

// PairCreated is a paid mutator transaction binding the contract method 0xe9f7206b.
//
// Solidity: function pairCreated(address pair) returns()
func (_CrossDexImpl *CrossDexImplTransactorSession) PairCreated(pair common.Address) (*types.Transaction, error) {
	return _CrossDexImpl.Contract.PairCreated(&_CrossDexImpl.TransactOpts, pair)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CrossDexImpl *CrossDexImplTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossDexImpl.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CrossDexImpl *CrossDexImplSession) RenounceOwnership() (*types.Transaction, error) {
	return _CrossDexImpl.Contract.RenounceOwnership(&_CrossDexImpl.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CrossDexImpl *CrossDexImplTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CrossDexImpl.Contract.RenounceOwnership(&_CrossDexImpl.TransactOpts)
}

// SetTickSizeSetter is a paid mutator transaction binding the contract method 0x98715ee7.
//
// Solidity: function setTickSizeSetter(address setter) returns()
func (_CrossDexImpl *CrossDexImplTransactor) SetTickSizeSetter(opts *bind.TransactOpts, setter common.Address) (*types.Transaction, error) {
	return _CrossDexImpl.contract.Transact(opts, "setTickSizeSetter", setter)
}

// SetTickSizeSetter is a paid mutator transaction binding the contract method 0x98715ee7.
//
// Solidity: function setTickSizeSetter(address setter) returns()
func (_CrossDexImpl *CrossDexImplSession) SetTickSizeSetter(setter common.Address) (*types.Transaction, error) {
	return _CrossDexImpl.Contract.SetTickSizeSetter(&_CrossDexImpl.TransactOpts, setter)
}

// SetTickSizeSetter is a paid mutator transaction binding the contract method 0x98715ee7.
//
// Solidity: function setTickSizeSetter(address setter) returns()
func (_CrossDexImpl *CrossDexImplTransactorSession) SetTickSizeSetter(setter common.Address) (*types.Transaction, error) {
	return _CrossDexImpl.Contract.SetTickSizeSetter(&_CrossDexImpl.TransactOpts, setter)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CrossDexImpl *CrossDexImplTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CrossDexImpl.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CrossDexImpl *CrossDexImplSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CrossDexImpl.Contract.TransferOwnership(&_CrossDexImpl.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CrossDexImpl *CrossDexImplTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CrossDexImpl.Contract.TransferOwnership(&_CrossDexImpl.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CrossDexImpl *CrossDexImplTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CrossDexImpl.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CrossDexImpl *CrossDexImplSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CrossDexImpl.Contract.UpgradeToAndCall(&_CrossDexImpl.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CrossDexImpl *CrossDexImplTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CrossDexImpl.Contract.UpgradeToAndCall(&_CrossDexImpl.TransactOpts, newImplementation, data)
}

// CrossDexImplInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the CrossDexImpl contract.
type CrossDexImplInitializedIterator struct {
	Event *CrossDexImplInitialized // Event containing the contract specifics and raw log

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
func (it *CrossDexImplInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossDexImplInitialized)
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
		it.Event = new(CrossDexImplInitialized)
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
func (it *CrossDexImplInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossDexImplInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossDexImplInitialized represents a Initialized event raised by the CrossDexImpl contract.
type CrossDexImplInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CrossDexImpl *CrossDexImplFilterer) FilterInitialized(opts *bind.FilterOpts) (*CrossDexImplInitializedIterator, error) {

	logs, sub, err := _CrossDexImpl.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CrossDexImplInitializedIterator{contract: _CrossDexImpl.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CrossDexImpl *CrossDexImplFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CrossDexImplInitialized) (event.Subscription, error) {

	logs, sub, err := _CrossDexImpl.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossDexImplInitialized)
				if err := _CrossDexImpl.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_CrossDexImpl *CrossDexImplFilterer) ParseInitialized(log types.Log) (*CrossDexImplInitialized, error) {
	event := new(CrossDexImplInitialized)
	if err := _CrossDexImpl.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossDexImplMarketCreatedIterator is returned from FilterMarketCreated and is used to iterate over the raw logs and unpacked data for MarketCreated events raised by the CrossDexImpl contract.
type CrossDexImplMarketCreatedIterator struct {
	Event *CrossDexImplMarketCreated // Event containing the contract specifics and raw log

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
func (it *CrossDexImplMarketCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossDexImplMarketCreated)
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
		it.Event = new(CrossDexImplMarketCreated)
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
func (it *CrossDexImplMarketCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossDexImplMarketCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossDexImplMarketCreated represents a MarketCreated event raised by the CrossDexImpl contract.
type CrossDexImplMarketCreated struct {
	Quote        common.Address
	Market       common.Address
	Owner        common.Address
	FeeCollector common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMarketCreated is a free log retrieval operation binding the contract event 0x5509f6235f0d08c563aedf1b655226cad4a613794902c476ad7a08accb43ebc5.
//
// Solidity: event MarketCreated(address indexed quote, address indexed market, address indexed owner, address fee_collector)
func (_CrossDexImpl *CrossDexImplFilterer) FilterMarketCreated(opts *bind.FilterOpts, quote []common.Address, market []common.Address, owner []common.Address) (*CrossDexImplMarketCreatedIterator, error) {

	var quoteRule []interface{}
	for _, quoteItem := range quote {
		quoteRule = append(quoteRule, quoteItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _CrossDexImpl.contract.FilterLogs(opts, "MarketCreated", quoteRule, marketRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &CrossDexImplMarketCreatedIterator{contract: _CrossDexImpl.contract, event: "MarketCreated", logs: logs, sub: sub}, nil
}

// WatchMarketCreated is a free log subscription operation binding the contract event 0x5509f6235f0d08c563aedf1b655226cad4a613794902c476ad7a08accb43ebc5.
//
// Solidity: event MarketCreated(address indexed quote, address indexed market, address indexed owner, address fee_collector)
func (_CrossDexImpl *CrossDexImplFilterer) WatchMarketCreated(opts *bind.WatchOpts, sink chan<- *CrossDexImplMarketCreated, quote []common.Address, market []common.Address, owner []common.Address) (event.Subscription, error) {

	var quoteRule []interface{}
	for _, quoteItem := range quote {
		quoteRule = append(quoteRule, quoteItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _CrossDexImpl.contract.WatchLogs(opts, "MarketCreated", quoteRule, marketRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossDexImplMarketCreated)
				if err := _CrossDexImpl.contract.UnpackLog(event, "MarketCreated", log); err != nil {
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

// ParseMarketCreated is a log parse operation binding the contract event 0x5509f6235f0d08c563aedf1b655226cad4a613794902c476ad7a08accb43ebc5.
//
// Solidity: event MarketCreated(address indexed quote, address indexed market, address indexed owner, address fee_collector)
func (_CrossDexImpl *CrossDexImplFilterer) ParseMarketCreated(log types.Log) (*CrossDexImplMarketCreated, error) {
	event := new(CrossDexImplMarketCreated)
	if err := _CrossDexImpl.contract.UnpackLog(event, "MarketCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossDexImplOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CrossDexImpl contract.
type CrossDexImplOwnershipTransferredIterator struct {
	Event *CrossDexImplOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CrossDexImplOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossDexImplOwnershipTransferred)
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
		it.Event = new(CrossDexImplOwnershipTransferred)
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
func (it *CrossDexImplOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossDexImplOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossDexImplOwnershipTransferred represents a OwnershipTransferred event raised by the CrossDexImpl contract.
type CrossDexImplOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CrossDexImpl *CrossDexImplFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CrossDexImplOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CrossDexImpl.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CrossDexImplOwnershipTransferredIterator{contract: _CrossDexImpl.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CrossDexImpl *CrossDexImplFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CrossDexImplOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CrossDexImpl.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossDexImplOwnershipTransferred)
				if err := _CrossDexImpl.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_CrossDexImpl *CrossDexImplFilterer) ParseOwnershipTransferred(log types.Log) (*CrossDexImplOwnershipTransferred, error) {
	event := new(CrossDexImplOwnershipTransferred)
	if err := _CrossDexImpl.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossDexImplTickSizeSetterSetIterator is returned from FilterTickSizeSetterSet and is used to iterate over the raw logs and unpacked data for TickSizeSetterSet events raised by the CrossDexImpl contract.
type CrossDexImplTickSizeSetterSetIterator struct {
	Event *CrossDexImplTickSizeSetterSet // Event containing the contract specifics and raw log

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
func (it *CrossDexImplTickSizeSetterSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossDexImplTickSizeSetterSet)
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
		it.Event = new(CrossDexImplTickSizeSetterSet)
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
func (it *CrossDexImplTickSizeSetterSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossDexImplTickSizeSetterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossDexImplTickSizeSetterSet represents a TickSizeSetterSet event raised by the CrossDexImpl contract.
type CrossDexImplTickSizeSetterSet struct {
	Before  common.Address
	Current common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTickSizeSetterSet is a free log retrieval operation binding the contract event 0x4f1656ebcdda14dfa0f9951b46485570075b9730e1f1d22cb4ddacc35374fde8.
//
// Solidity: event TickSizeSetterSet(address indexed before, address indexed current)
func (_CrossDexImpl *CrossDexImplFilterer) FilterTickSizeSetterSet(opts *bind.FilterOpts, before []common.Address, current []common.Address) (*CrossDexImplTickSizeSetterSetIterator, error) {

	var beforeRule []interface{}
	for _, beforeItem := range before {
		beforeRule = append(beforeRule, beforeItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _CrossDexImpl.contract.FilterLogs(opts, "TickSizeSetterSet", beforeRule, currentRule)
	if err != nil {
		return nil, err
	}
	return &CrossDexImplTickSizeSetterSetIterator{contract: _CrossDexImpl.contract, event: "TickSizeSetterSet", logs: logs, sub: sub}, nil
}

// WatchTickSizeSetterSet is a free log subscription operation binding the contract event 0x4f1656ebcdda14dfa0f9951b46485570075b9730e1f1d22cb4ddacc35374fde8.
//
// Solidity: event TickSizeSetterSet(address indexed before, address indexed current)
func (_CrossDexImpl *CrossDexImplFilterer) WatchTickSizeSetterSet(opts *bind.WatchOpts, sink chan<- *CrossDexImplTickSizeSetterSet, before []common.Address, current []common.Address) (event.Subscription, error) {

	var beforeRule []interface{}
	for _, beforeItem := range before {
		beforeRule = append(beforeRule, beforeItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _CrossDexImpl.contract.WatchLogs(opts, "TickSizeSetterSet", beforeRule, currentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossDexImplTickSizeSetterSet)
				if err := _CrossDexImpl.contract.UnpackLog(event, "TickSizeSetterSet", log); err != nil {
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

// ParseTickSizeSetterSet is a log parse operation binding the contract event 0x4f1656ebcdda14dfa0f9951b46485570075b9730e1f1d22cb4ddacc35374fde8.
//
// Solidity: event TickSizeSetterSet(address indexed before, address indexed current)
func (_CrossDexImpl *CrossDexImplFilterer) ParseTickSizeSetterSet(log types.Log) (*CrossDexImplTickSizeSetterSet, error) {
	event := new(CrossDexImplTickSizeSetterSet)
	if err := _CrossDexImpl.contract.UnpackLog(event, "TickSizeSetterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossDexImplUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the CrossDexImpl contract.
type CrossDexImplUpgradedIterator struct {
	Event *CrossDexImplUpgraded // Event containing the contract specifics and raw log

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
func (it *CrossDexImplUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossDexImplUpgraded)
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
		it.Event = new(CrossDexImplUpgraded)
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
func (it *CrossDexImplUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossDexImplUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossDexImplUpgraded represents a Upgraded event raised by the CrossDexImpl contract.
type CrossDexImplUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_CrossDexImpl *CrossDexImplFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*CrossDexImplUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _CrossDexImpl.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &CrossDexImplUpgradedIterator{contract: _CrossDexImpl.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_CrossDexImpl *CrossDexImplFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *CrossDexImplUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _CrossDexImpl.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossDexImplUpgraded)
				if err := _CrossDexImpl.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_CrossDexImpl *CrossDexImplFilterer) ParseUpgraded(log types.Log) (*CrossDexImplUpgraded, error) {
	event := new(CrossDexImplUpgraded)
	if err := _CrossDexImpl.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
