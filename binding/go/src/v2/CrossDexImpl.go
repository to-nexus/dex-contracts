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

// CrossDexImplMetaData contains all meta data concerning the CrossDexImpl contract.
var CrossDexImplMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ROUTER\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allMarkets\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"quotes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"markets\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"checkTickSizeRoles\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quote\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeCollector\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBps\",\"type\":\"uint256\"}],\"name\":\"createMarket\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_routerImpl\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_findPrevPriceCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxMatchCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_cancelLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_marketImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pairImpl\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"}],\"name\":\"isMarket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketImpl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"name\":\"pairCreated\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pairImpl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"name\":\"pairToMarket\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"quote\",\"type\":\"address\"}],\"name\":\"quoteToMarket\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"setter\",\"type\":\"address\"}],\"name\":\"setTickSizeSetter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tickSizeSetter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"quote\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fee_collector\",\"type\":\"address\"}],\"name\":\"MarketCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"before\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"TickSizeSetterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Create2EmptyBytecode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"CrossDexAlreadyCreatedMarketQuote\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"CrossDexInitializeData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"CrossDexInvalidMarketAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"input\",\"type\":\"address\"}],\"name\":\"CrossDexInvalidTickSizeSetter\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"CrossDexUnauthorizedChangeTickSizes\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"EnumerableMapNonexistentKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedDeployment\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"}]",
	ID:  "CrossDexImpl",
	Bin: "0x60a060405230608052348015610013575f5ffd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6080516122576100f95f395f8181611298015281816112c101526114e201526122575ff3fe60806040526004361061013d575f3560e01c80636ec934da116100bb578063a1eea77811610071578063beb380f111610057578063beb380f114610411578063e9f7206b14610430578063f2fde38b1461044f575f5ffd5b8063a1eea7781461039d578063ad3cb1cc146103bc575f5ffd5b80637f764a85116100a15780637f764a85146103095780638da5cb5b1461033557806398715ee71461037e575f5ffd5b80636ec934da146102c6578063715018a6146102f5575f5ffd5b806334eaeeb911610110578063375a7cba116100f6578063375a7cba1461026f5780634f1ef2861461029157806352d1902d146102a4575f5ffd5b806334eaeeb91461021757806335f7d45614610243575f5ffd5b80630729da0b1461014157806308270573146101625780632e92c7be146101cd57806332fe7b26146101ec575b5f5ffd5b34801561014c575f5ffd5b5061016061015b366004611aa0565b61046e565b005b34801561016d575f5ffd5b506101a361017c366004611b16565b60066020525f908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b3480156101d8575f5ffd5b506101a36101e7366004611b31565b6108a5565b3480156101f7575f5ffd5b505f546101a39073ffffffffffffffffffffffffffffffffffffffff1681565b348015610222575f5ffd5b506001546101a39073ffffffffffffffffffffffffffffffffffffffff1681565b34801561024e575f5ffd5b506002546101a39073ffffffffffffffffffffffffffffffffffffffff1681565b34801561027a575f5ffd5b50610283610b45565b6040516101c4929190611bcf565b61016061029f366004611c20565b610c54565b3480156102af575f5ffd5b506102b8610c73565b6040519081526020016101c4565b3480156102d1575f5ffd5b506102e56102e0366004611b16565b610ca1565b60405190151581526020016101c4565b348015610300575f5ffd5b50610160610dca565b348015610314575f5ffd5b506007546101a39073ffffffffffffffffffffffffffffffffffffffff1681565b348015610340575f5ffd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c1993005473ffffffffffffffffffffffffffffffffffffffff166101a3565b348015610389575f5ffd5b50610160610398366004611b16565b610ddd565b3480156103a8575f5ffd5b506101606103b7366004611b16565b610f0a565b3480156103c7575f5ffd5b506104046040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b6040516101c49190611d6f565b34801561041c575f5ffd5b506101a361042b366004611b16565b610f9c565b34801561043b575f5ffd5b5061016061044a366004611b16565b610fae565b34801561045a575f5ffd5b50610160610469366004611b16565b61103c565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff165f811580156104b85750825b90505f8267ffffffffffffffff1660011480156104d45750303b155b9050811580156104e2575080155b15610519576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000166001178555831561057a5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b6105838c61109c565b73ffffffffffffffffffffffffffffffffffffffff8b166105f7576040517fc53a07b90000000000000000000000000000000000000000000000000000000081527f726f75746572496d706c0000000000000000000000000000000000000000000060048201526024015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8716610666576040517fc53a07b90000000000000000000000000000000000000000000000000000000081527f6d61726b6574496d706c0000000000000000000000000000000000000000000060048201526024016105ee565b73ffffffffffffffffffffffffffffffffffffffff86166106d5576040517fc53a07b90000000000000000000000000000000000000000000000000000000081527f70616972496d706c00000000000000000000000000000000000000000000000060048201526024016105ee565b5f8b6040516106e390611a72565b73ffffffffffffffffffffffffffffffffffffffff90911681526040602082018190525f90820152606001604051809103905ff080158015610727573d5f5f3e3d5ffd5b505f80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040517f80d85911000000000000000000000000000000000000000000000000000000008152600481018e9052602481018d9052604481018c9052919250906380d85911906064015f604051808303815f87803b1580156107c9575f5ffd5b505af11580156107db573d5f5f3e3d5ffd5b50506001805473ffffffffffffffffffffffffffffffffffffffff808d167fffffffffffffffffffffffff00000000000000000000000000000000000000009283161790925560028054928c16929091169190911790555050841590506108975784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050505050505050565b5f6108ae6110ad565b5f604051806020016108bf90611a72565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082820381018352601f9091011660408190526001545f5460025473ffffffffffffffffffffffffffffffffffffffff8b81166024860152918216604485015289821660648501528116608484015287811660a484015260c48301879052169060e401604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f95b6ef0c0000000000000000000000000000000000000000000000000000000017905290516109c793929101611d81565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815290829052610a039291602001611dc6565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606088901b16602083015291505f906034016040516020818303038152906040528051906020012090505f610a895f838561113b565b9050610a976003888361122e565b610ae5576040517fed703a9800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff881660048201526024016105ee565b60405173ffffffffffffffffffffffffffffffffffffffff8781168252808a1691838216918a16907f5509f6235f0d08c563aedf1b655226cad4a613794902c476ad7a08accb43ebc59060200160405180910390a4979650505050505050565b6060805f610b53600361125b565b90508067ffffffffffffffff811115610b6e57610b6e611bf3565b604051908082528060200260200182016040528015610b97578160200160208202803683370190505b5092508067ffffffffffffffff811115610bb357610bb3611bf3565b604051908082528060200260200182016040528015610bdc578160200160208202803683370190505b5091505f5b81811015610c4e57610bf4600382611265565b858381518110610c0657610c06611dda565b60200260200101858481518110610c1f57610c1f611dda565b73ffffffffffffffffffffffffffffffffffffffff938416602091820292909201015291169052600101610be1565b50509091565b610c5c611280565b610c6582611384565b610c6f828261138c565b5050565b5f610c7c6114ca565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b60408051600481526024810182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f9c5798390000000000000000000000000000000000000000000000000000000017905290515f918291829173ffffffffffffffffffffffffffffffffffffffff861691610d219190611e07565b5f60405180830381855afa9150503d805f8114610d59576040519150601f19603f3d011682016040523d82523d5f602084013e610d5e565b606091505b509150915081610d7157505f9392505050565b5f81806020019051810190610d869190611e12565b905073ffffffffffffffffffffffffffffffffffffffff8516610daa600383611539565b73ffffffffffffffffffffffffffffffffffffffff161495945050505050565b610dd26110ad565b610ddb5f61155a565b565b610de56110ad565b73ffffffffffffffffffffffffffffffffffffffff81161580610e22575060075473ffffffffffffffffffffffffffffffffffffffff8281169116145b15610e7d576007546040517f18ddd67e00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff918216600482015290821660248201526044016105ee565b60075460405173ffffffffffffffffffffffffffffffffffffffff8084169216907f4f1656ebcdda14dfa0f9951b46485570075b9730e1f1d22cb4ddacc35374fde8905f90a3600780547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60075473ffffffffffffffffffffffffffffffffffffffff161580610f4a575060075473ffffffffffffffffffffffffffffffffffffffff828116911614155b15610f99576040517ff42eaafb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024016105ee565b50565b5f610fa8600383611539565b92915050565b610fb733610ca1565b610fef576040517f74be5d0c0000000000000000000000000000000000000000000000000000000081523360048201526024016105ee565b73ffffffffffffffffffffffffffffffffffffffff165f90815260066020526040902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001633179055565b6110446110ad565b73ffffffffffffffffffffffffffffffffffffffff8116611093576040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081525f60048201526024016105ee565b610f998161155a565b6110a46115ef565b610f9981611656565b336110ec7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c1993005473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614610ddb576040517f118cdaa70000000000000000000000000000000000000000000000000000000081523360048201526024016105ee565b5f8347101561117f576040517fcf479181000000000000000000000000000000000000000000000000000000008152476004820152602481018590526044016105ee565b81515f036111b9576040517f4ca249dc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8282516020840186f590503d1519811516156111da576040513d5f823e3d81fd5b73ffffffffffffffffffffffffffffffffffffffff8116611227576040517fb06ebf3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b9392505050565b5f6112538473ffffffffffffffffffffffffffffffffffffffff80861690851661165e565b949350505050565b5f610fa88261167a565b5f8080806112738686611684565b9097909650945050505050565b3073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016148061134d57507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166113347f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614155b15610ddb576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610f996110ad565b8173ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015611411575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016820190925261140e91810190611e2d565b60015b61145f576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff831660048201526024016105ee565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc81146114bb576040517faa1d49a4000000000000000000000000000000000000000000000000000000008152600481018290526024016105ee565b6114c583836116ad565b505050565b3073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610ddb576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6112278373ffffffffffffffffffffffffffffffffffffffff841661170f565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080547fffffffffffffffffffffffff0000000000000000000000000000000000000000811673ffffffffffffffffffffffffffffffffffffffff848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff16610ddb576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6110446115ef565b5f8281526002840160205260408120829055611253848461176e565b5f610fa882611779565b5f80806116918585611782565b5f81815260029690960160205260409095205494959350505050565b6116b68261178d565b60405173ffffffffffffffffffffffffffffffffffffffff8316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115611707576114c5828261185b565b610c6f6118da565b5f8181526002830160205260408120548015801561173457506117328484611912565b155b15611227576040517f02b56686000000000000000000000000000000000000000000000000000000008152600481018490526024016105ee565b5f611227838361192f565b5f610fa8825490565b5f611227838361197b565b8073ffffffffffffffffffffffffffffffffffffffff163b5f036117f5576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024016105ee565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60605f5f8473ffffffffffffffffffffffffffffffffffffffff16846040516118849190611e07565b5f60405180830381855af49150503d805f81146118bc576040519150601f19603f3d011682016040523d82523d5f602084013e6118c1565b606091505b50915091506118d18583836119a1565b95945050505050565b3415610ddb576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f61122783835f8181526001830160205260408120541515611227565b5f81815260018301602052604081205461197457508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610fa8565b505f610fa8565b5f825f01828154811061199057611990611dda565b905f5260205f200154905092915050565b6060826119b6576119b182611a30565b611227565b81511580156119da575073ffffffffffffffffffffffffffffffffffffffff84163b155b15611a29576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024016105ee565b5080611227565b805115611a405780518082602001fd5b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103dd80611e4583390190565b73ffffffffffffffffffffffffffffffffffffffff81168114610f99575f5ffd5b5f5f5f5f5f5f5f60e0888a031215611ab6575f5ffd5b8735611ac181611a7f565b96506020880135611ad181611a7f565b955060408801359450606088013593506080880135925060a0880135611af681611a7f565b915060c0880135611b0681611a7f565b8091505092959891949750929550565b5f60208284031215611b26575f5ffd5b813561122781611a7f565b5f5f5f5f60808587031215611b44575f5ffd5b8435611b4f81611a7f565b93506020850135611b5f81611a7f565b92506040850135611b6f81611a7f565b9396929550929360600135925050565b5f8151808452602084019350602083015f5b82811015611bc557815173ffffffffffffffffffffffffffffffffffffffff16865260209586019590910190600101611b91565b5093949350505050565b604081525f611be16040830185611b7f565b82810360208401526118d18185611b7f565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f5f60408385031215611c31575f5ffd5b8235611c3c81611a7f565b9150602083013567ffffffffffffffff811115611c57575f5ffd5b8301601f81018513611c67575f5ffd5b803567ffffffffffffffff811115611c8157611c81611bf3565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff82111715611ced57611ced611bf3565b604052818152828201602001871015611d04575f5ffd5b816020840160208301375f602083830101528093505050509250929050565b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f6112276020830184611d23565b73ffffffffffffffffffffffffffffffffffffffff83168152604060208201525f6112536040830184611d23565b5f81518060208401855e5f93019283525090919050565b5f611253611dd48386611daf565b84611daf565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f6112278284611daf565b5f60208284031215611e22575f5ffd5b815161122781611a7f565b5f60208284031215611e3d575f5ffd5b505191905056fe60806040526040516103dd3803806103dd8339810160408190526100229161023c565b61002c8282610033565b5050610321565b61003c82610091565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a280511561008557610080828261010c565b505050565b61008d61017f565b5050565b806001600160a01b03163b5f036100cb57604051634c9c8ce360e01b81526001600160a01b03821660048201526024015b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b031684604051610128919061030b565b5f60405180830381855af49150503d805f8114610160576040519150601f19603f3d011682016040523d82523d5f602084013e610165565b606091505b5090925090506101768583836101a0565b95945050505050565b341561019e5760405163b398979f60e01b815260040160405180910390fd5b565b6060826101b5576101b0826101ff565b6101f8565b81511580156101cc57506001600160a01b0384163b155b156101f557604051639996b31560e01b81526001600160a01b03851660048201526024016100c2565b50805b9392505050565b80511561020f5780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b634e487b7160e01b5f52604160045260245ffd5b5f5f6040838503121561024d575f5ffd5b82516001600160a01b0381168114610263575f5ffd5b60208401519092506001600160401b0381111561027e575f5ffd5b8301601f8101851361028e575f5ffd5b80516001600160401b038111156102a7576102a7610228565b604051601f8201601f19908116603f011681016001600160401b03811182821017156102d5576102d5610228565b6040528181528282016020018710156102ec575f5ffd5b8160208401602083015e5f602083830101528093505050509250929050565b5f82518060208501845e5f920191825250919050565b60b08061032d5f395ff3fe6080604052600a600c565b005b60186014601a565b605d565b565b5f60587f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5473ffffffffffffffffffffffffffffffffffffffff1690565b905090565b365f5f375f5f365f845af43d5f5f3e8080156076573d5ff35b3d5ffdfea264697066735822122053cea113bf96df2d97b042d29245e0a41b1290b530767f8b25bfeaed774391b764736f6c634300081c0033a2646970667358221220967f768dd8cbab77474517ca17dfd2568afbae713116e3eb87953d672c921bac64736f6c634300081c0033",
}

// CrossDexImpl is an auto generated Go binding around an Ethereum contract.
type CrossDexImpl struct {
	abi abi.ABI
}

// NewCrossDexImpl creates a new instance of CrossDexImpl.
func NewCrossDexImpl() *CrossDexImpl {
	parsed, err := CrossDexImplMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &CrossDexImpl{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *CrossDexImpl) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackROUTER is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x32fe7b26.
//
// Solidity: function ROUTER() view returns(address)
func (crossDexImpl *CrossDexImpl) PackROUTER() []byte {
	enc, err := crossDexImpl.abi.Pack("ROUTER")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackROUTER is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x32fe7b26.
//
// Solidity: function ROUTER() view returns(address)
func (crossDexImpl *CrossDexImpl) UnpackROUTER(data []byte) (common.Address, error) {
	out, err := crossDexImpl.abi.Unpack("ROUTER", data)
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
func (crossDexImpl *CrossDexImpl) PackUPGRADEINTERFACEVERSION() []byte {
	enc, err := crossDexImpl.abi.Pack("UPGRADE_INTERFACE_VERSION")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackUPGRADEINTERFACEVERSION is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (crossDexImpl *CrossDexImpl) UnpackUPGRADEINTERFACEVERSION(data []byte) (string, error) {
	out, err := crossDexImpl.abi.Unpack("UPGRADE_INTERFACE_VERSION", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackAllMarkets is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x375a7cba.
//
// Solidity: function allMarkets() view returns(address[] quotes, address[] markets)
func (crossDexImpl *CrossDexImpl) PackAllMarkets() []byte {
	enc, err := crossDexImpl.abi.Pack("allMarkets")
	if err != nil {
		panic(err)
	}
	return enc
}

// AllMarketsOutput serves as a container for the return parameters of contract
// method AllMarkets.

// UnpackAllMarkets is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x375a7cba.
//
// Solidity: function allMarkets() view returns(address[] quotes, address[] markets)
func (crossDexImpl *CrossDexImpl) UnpackAllMarkets(data []byte) (AllMarketsOutput, error) {
	out, err := crossDexImpl.abi.Unpack("allMarkets", data)
	outstruct := new(AllMarketsOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Quotes = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Markets = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	return *outstruct, err

}

// PackCheckTickSizeRoles is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa1eea778.
//
// Solidity: function checkTickSizeRoles(address account) view returns()
func (crossDexImpl *CrossDexImpl) PackCheckTickSizeRoles(account common.Address) []byte {
	enc, err := crossDexImpl.abi.Pack("checkTickSizeRoles", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCreateMarket is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2e92c7be.
//
// Solidity: function createMarket(address _owner, address quote, address feeCollector, uint256 feeBps) returns(address)
func (crossDexImpl *CrossDexImpl) PackCreateMarket(owner common.Address, quote common.Address, feeCollector common.Address, feeBps *big.Int) []byte {
	enc, err := crossDexImpl.abi.Pack("createMarket", owner, quote, feeCollector, feeBps)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackCreateMarket is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2e92c7be.
//
// Solidity: function createMarket(address _owner, address quote, address feeCollector, uint256 feeBps) returns(address)
func (crossDexImpl *CrossDexImpl) UnpackCreateMarket(data []byte) (common.Address, error) {
	out, err := crossDexImpl.abi.Unpack("createMarket", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0729da0b.
//
// Solidity: function initialize(address _owner, address _routerImpl, uint256 _findPrevPriceCount, uint256 _maxMatchCount, uint256 _cancelLimit, address _marketImpl, address _pairImpl) returns()
func (crossDexImpl *CrossDexImpl) PackInitialize(owner common.Address, routerImpl common.Address, findPrevPriceCount *big.Int, maxMatchCount *big.Int, cancelLimit *big.Int, marketImpl common.Address, pairImpl common.Address) []byte {
	enc, err := crossDexImpl.abi.Pack("initialize", owner, routerImpl, findPrevPriceCount, maxMatchCount, cancelLimit, marketImpl, pairImpl)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackIsMarket is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6ec934da.
//
// Solidity: function isMarket(address market) view returns(bool)
func (crossDexImpl *CrossDexImpl) PackIsMarket(market common.Address) []byte {
	enc, err := crossDexImpl.abi.Pack("isMarket", market)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsMarket is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6ec934da.
//
// Solidity: function isMarket(address market) view returns(bool)
func (crossDexImpl *CrossDexImpl) UnpackIsMarket(data []byte) (bool, error) {
	out, err := crossDexImpl.abi.Unpack("isMarket", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackMarketImpl is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x34eaeeb9.
//
// Solidity: function marketImpl() view returns(address)
func (crossDexImpl *CrossDexImpl) PackMarketImpl() []byte {
	enc, err := crossDexImpl.abi.Pack("marketImpl")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackMarketImpl is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x34eaeeb9.
//
// Solidity: function marketImpl() view returns(address)
func (crossDexImpl *CrossDexImpl) UnpackMarketImpl(data []byte) (common.Address, error) {
	out, err := crossDexImpl.abi.Unpack("marketImpl", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (crossDexImpl *CrossDexImpl) PackOwner() []byte {
	enc, err := crossDexImpl.abi.Pack("owner")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (crossDexImpl *CrossDexImpl) UnpackOwner(data []byte) (common.Address, error) {
	out, err := crossDexImpl.abi.Unpack("owner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackPairCreated is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe9f7206b.
//
// Solidity: function pairCreated(address pair) returns()
func (crossDexImpl *CrossDexImpl) PackPairCreated(pair common.Address) []byte {
	enc, err := crossDexImpl.abi.Pack("pairCreated", pair)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackPairImpl is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x35f7d456.
//
// Solidity: function pairImpl() view returns(address)
func (crossDexImpl *CrossDexImpl) PackPairImpl() []byte {
	enc, err := crossDexImpl.abi.Pack("pairImpl")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPairImpl is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x35f7d456.
//
// Solidity: function pairImpl() view returns(address)
func (crossDexImpl *CrossDexImpl) UnpackPairImpl(data []byte) (common.Address, error) {
	out, err := crossDexImpl.abi.Unpack("pairImpl", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackPairToMarket is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x08270573.
//
// Solidity: function pairToMarket(address pair) view returns(address)
func (crossDexImpl *CrossDexImpl) PackPairToMarket(pair common.Address) []byte {
	enc, err := crossDexImpl.abi.Pack("pairToMarket", pair)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPairToMarket is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x08270573.
//
// Solidity: function pairToMarket(address pair) view returns(address)
func (crossDexImpl *CrossDexImpl) UnpackPairToMarket(data []byte) (common.Address, error) {
	out, err := crossDexImpl.abi.Unpack("pairToMarket", data)
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
func (crossDexImpl *CrossDexImpl) PackProxiableUUID() []byte {
	enc, err := crossDexImpl.abi.Pack("proxiableUUID")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackProxiableUUID is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (crossDexImpl *CrossDexImpl) UnpackProxiableUUID(data []byte) ([32]byte, error) {
	out, err := crossDexImpl.abi.Unpack("proxiableUUID", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackQuoteToMarket is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbeb380f1.
//
// Solidity: function quoteToMarket(address quote) view returns(address)
func (crossDexImpl *CrossDexImpl) PackQuoteToMarket(quote common.Address) []byte {
	enc, err := crossDexImpl.abi.Pack("quoteToMarket", quote)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackQuoteToMarket is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbeb380f1.
//
// Solidity: function quoteToMarket(address quote) view returns(address)
func (crossDexImpl *CrossDexImpl) UnpackQuoteToMarket(data []byte) (common.Address, error) {
	out, err := crossDexImpl.abi.Unpack("quoteToMarket", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackRenounceOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (crossDexImpl *CrossDexImpl) PackRenounceOwnership() []byte {
	enc, err := crossDexImpl.abi.Pack("renounceOwnership")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetTickSizeSetter is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x98715ee7.
//
// Solidity: function setTickSizeSetter(address setter) returns()
func (crossDexImpl *CrossDexImpl) PackSetTickSizeSetter(setter common.Address) []byte {
	enc, err := crossDexImpl.abi.Pack("setTickSizeSetter", setter)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackTickSizeSetter is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7f764a85.
//
// Solidity: function tickSizeSetter() view returns(address)
func (crossDexImpl *CrossDexImpl) PackTickSizeSetter() []byte {
	enc, err := crossDexImpl.abi.Pack("tickSizeSetter")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTickSizeSetter is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7f764a85.
//
// Solidity: function tickSizeSetter() view returns(address)
func (crossDexImpl *CrossDexImpl) UnpackTickSizeSetter(data []byte) (common.Address, error) {
	out, err := crossDexImpl.abi.Unpack("tickSizeSetter", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackTransferOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (crossDexImpl *CrossDexImpl) PackTransferOwnership(newOwner common.Address) []byte {
	enc, err := crossDexImpl.abi.Pack("transferOwnership", newOwner)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackUpgradeToAndCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (crossDexImpl *CrossDexImpl) PackUpgradeToAndCall(newImplementation common.Address, data []byte) []byte {
	enc, err := crossDexImpl.abi.Pack("upgradeToAndCall", newImplementation, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// CrossDexImplInitialized represents a Initialized event raised by the CrossDexImpl contract.
type CrossDexImplInitialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const CrossDexImplInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (CrossDexImplInitialized) ContractEventName() string {
	return CrossDexImplInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (crossDexImpl *CrossDexImpl) UnpackInitializedEvent(log *types.Log) (*CrossDexImplInitialized, error) {
	event := "Initialized"
	if log.Topics[0] != crossDexImpl.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CrossDexImplInitialized)
	if len(log.Data) > 0 {
		if err := crossDexImpl.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range crossDexImpl.abi.Events[event].Inputs {
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

// CrossDexImplMarketCreated represents a MarketCreated event raised by the CrossDexImpl contract.
type CrossDexImplMarketCreated struct {
	Quote        common.Address
	Market       common.Address
	Owner        common.Address
	FeeCollector common.Address
	Raw          *types.Log // Blockchain specific contextual infos
}

const CrossDexImplMarketCreatedEventName = "MarketCreated"

// ContractEventName returns the user-defined event name.
func (CrossDexImplMarketCreated) ContractEventName() string {
	return CrossDexImplMarketCreatedEventName
}

// UnpackMarketCreatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MarketCreated(address indexed quote, address indexed market, address indexed owner, address fee_collector)
func (crossDexImpl *CrossDexImpl) UnpackMarketCreatedEvent(log *types.Log) (*CrossDexImplMarketCreated, error) {
	event := "MarketCreated"
	if log.Topics[0] != crossDexImpl.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CrossDexImplMarketCreated)
	if len(log.Data) > 0 {
		if err := crossDexImpl.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range crossDexImpl.abi.Events[event].Inputs {
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

// CrossDexImplOwnershipTransferred represents a OwnershipTransferred event raised by the CrossDexImpl contract.
type CrossDexImplOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const CrossDexImplOwnershipTransferredEventName = "OwnershipTransferred"

// ContractEventName returns the user-defined event name.
func (CrossDexImplOwnershipTransferred) ContractEventName() string {
	return CrossDexImplOwnershipTransferredEventName
}

// UnpackOwnershipTransferredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (crossDexImpl *CrossDexImpl) UnpackOwnershipTransferredEvent(log *types.Log) (*CrossDexImplOwnershipTransferred, error) {
	event := "OwnershipTransferred"
	if log.Topics[0] != crossDexImpl.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CrossDexImplOwnershipTransferred)
	if len(log.Data) > 0 {
		if err := crossDexImpl.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range crossDexImpl.abi.Events[event].Inputs {
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

// CrossDexImplTickSizeSetterSet represents a TickSizeSetterSet event raised by the CrossDexImpl contract.
type CrossDexImplTickSizeSetterSet struct {
	Before  common.Address
	Current common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const CrossDexImplTickSizeSetterSetEventName = "TickSizeSetterSet"

// ContractEventName returns the user-defined event name.
func (CrossDexImplTickSizeSetterSet) ContractEventName() string {
	return CrossDexImplTickSizeSetterSetEventName
}

// UnpackTickSizeSetterSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TickSizeSetterSet(address indexed before, address indexed current)
func (crossDexImpl *CrossDexImpl) UnpackTickSizeSetterSetEvent(log *types.Log) (*CrossDexImplTickSizeSetterSet, error) {
	event := "TickSizeSetterSet"
	if log.Topics[0] != crossDexImpl.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CrossDexImplTickSizeSetterSet)
	if len(log.Data) > 0 {
		if err := crossDexImpl.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range crossDexImpl.abi.Events[event].Inputs {
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

// CrossDexImplUpgraded represents a Upgraded event raised by the CrossDexImpl contract.
type CrossDexImplUpgraded struct {
	Implementation common.Address
	Raw            *types.Log // Blockchain specific contextual infos
}

const CrossDexImplUpgradedEventName = "Upgraded"

// ContractEventName returns the user-defined event name.
func (CrossDexImplUpgraded) ContractEventName() string {
	return CrossDexImplUpgradedEventName
}

// UnpackUpgradedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Upgraded(address indexed implementation)
func (crossDexImpl *CrossDexImpl) UnpackUpgradedEvent(log *types.Log) (*CrossDexImplUpgraded, error) {
	event := "Upgraded"
	if log.Topics[0] != crossDexImpl.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CrossDexImplUpgraded)
	if len(log.Data) > 0 {
		if err := crossDexImpl.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range crossDexImpl.abi.Events[event].Inputs {
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
func (crossDexImpl *CrossDexImpl) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], crossDexImpl.abi.Errors["AddressEmptyCode"].ID.Bytes()[:4]) {
		return crossDexImpl.UnpackAddressEmptyCodeError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossDexImpl.abi.Errors["Create2EmptyBytecode"].ID.Bytes()[:4]) {
		return crossDexImpl.UnpackCreate2EmptyBytecodeError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossDexImpl.abi.Errors["CrossDexAlreadyCreatedMarketQuote"].ID.Bytes()[:4]) {
		return crossDexImpl.UnpackCrossDexAlreadyCreatedMarketQuoteError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossDexImpl.abi.Errors["CrossDexInitializeData"].ID.Bytes()[:4]) {
		return crossDexImpl.UnpackCrossDexInitializeDataError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossDexImpl.abi.Errors["CrossDexInvalidMarketAddress"].ID.Bytes()[:4]) {
		return crossDexImpl.UnpackCrossDexInvalidMarketAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossDexImpl.abi.Errors["CrossDexInvalidTickSizeSetter"].ID.Bytes()[:4]) {
		return crossDexImpl.UnpackCrossDexInvalidTickSizeSetterError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossDexImpl.abi.Errors["CrossDexUnauthorizedChangeTickSizes"].ID.Bytes()[:4]) {
		return crossDexImpl.UnpackCrossDexUnauthorizedChangeTickSizesError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossDexImpl.abi.Errors["ERC1967InvalidImplementation"].ID.Bytes()[:4]) {
		return crossDexImpl.UnpackERC1967InvalidImplementationError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossDexImpl.abi.Errors["ERC1967NonPayable"].ID.Bytes()[:4]) {
		return crossDexImpl.UnpackERC1967NonPayableError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossDexImpl.abi.Errors["EnumerableMapNonexistentKey"].ID.Bytes()[:4]) {
		return crossDexImpl.UnpackEnumerableMapNonexistentKeyError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossDexImpl.abi.Errors["FailedCall"].ID.Bytes()[:4]) {
		return crossDexImpl.UnpackFailedCallError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossDexImpl.abi.Errors["FailedDeployment"].ID.Bytes()[:4]) {
		return crossDexImpl.UnpackFailedDeploymentError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossDexImpl.abi.Errors["InsufficientBalance"].ID.Bytes()[:4]) {
		return crossDexImpl.UnpackInsufficientBalanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossDexImpl.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return crossDexImpl.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossDexImpl.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return crossDexImpl.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossDexImpl.abi.Errors["OwnableInvalidOwner"].ID.Bytes()[:4]) {
		return crossDexImpl.UnpackOwnableInvalidOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossDexImpl.abi.Errors["OwnableUnauthorizedAccount"].ID.Bytes()[:4]) {
		return crossDexImpl.UnpackOwnableUnauthorizedAccountError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossDexImpl.abi.Errors["UUPSUnauthorizedCallContext"].ID.Bytes()[:4]) {
		return crossDexImpl.UnpackUUPSUnauthorizedCallContextError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossDexImpl.abi.Errors["UUPSUnsupportedProxiableUUID"].ID.Bytes()[:4]) {
		return crossDexImpl.UnpackUUPSUnsupportedProxiableUUIDError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// CrossDexImplAddressEmptyCode represents a AddressEmptyCode error raised by the CrossDexImpl contract.
type CrossDexImplAddressEmptyCode struct {
	Target common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressEmptyCode(address target)
func CrossDexImplAddressEmptyCodeErrorID() common.Hash {
	return common.HexToHash("0x9996b315c842ff135b8fc4a08ad5df1c344efbc03d2687aecc0678050d2aac89")
}

// UnpackAddressEmptyCodeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressEmptyCode(address target)
func (crossDexImpl *CrossDexImpl) UnpackAddressEmptyCodeError(raw []byte) (*CrossDexImplAddressEmptyCode, error) {
	out := new(CrossDexImplAddressEmptyCode)
	if err := crossDexImpl.abi.UnpackIntoInterface(out, "AddressEmptyCode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossDexImplCreate2EmptyBytecode represents a Create2EmptyBytecode error raised by the CrossDexImpl contract.
type CrossDexImplCreate2EmptyBytecode struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error Create2EmptyBytecode()
func CrossDexImplCreate2EmptyBytecodeErrorID() common.Hash {
	return common.HexToHash("0x4ca249dcffe41558ef8b961d71c905e4fa4317a1663f377b9610642e4e0abdb6")
}

// UnpackCreate2EmptyBytecodeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error Create2EmptyBytecode()
func (crossDexImpl *CrossDexImpl) UnpackCreate2EmptyBytecodeError(raw []byte) (*CrossDexImplCreate2EmptyBytecode, error) {
	out := new(CrossDexImplCreate2EmptyBytecode)
	if err := crossDexImpl.abi.UnpackIntoInterface(out, "Create2EmptyBytecode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossDexImplCrossDexAlreadyCreatedMarketQuote represents a CrossDexAlreadyCreatedMarketQuote error raised by the CrossDexImpl contract.
type CrossDexImplCrossDexAlreadyCreatedMarketQuote struct {
	Arg0 common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error CrossDexAlreadyCreatedMarketQuote(address arg0)
func CrossDexImplCrossDexAlreadyCreatedMarketQuoteErrorID() common.Hash {
	return common.HexToHash("0xed703a980240a0d2b459c71f53e7f8481420e17246a1c86f7d2311e4f9f44378")
}

// UnpackCrossDexAlreadyCreatedMarketQuoteError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error CrossDexAlreadyCreatedMarketQuote(address arg0)
func (crossDexImpl *CrossDexImpl) UnpackCrossDexAlreadyCreatedMarketQuoteError(raw []byte) (*CrossDexImplCrossDexAlreadyCreatedMarketQuote, error) {
	out := new(CrossDexImplCrossDexAlreadyCreatedMarketQuote)
	if err := crossDexImpl.abi.UnpackIntoInterface(out, "CrossDexAlreadyCreatedMarketQuote", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossDexImplCrossDexInitializeData represents a CrossDexInitializeData error raised by the CrossDexImpl contract.
type CrossDexImplCrossDexInitializeData struct {
	Arg0 [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error CrossDexInitializeData(bytes32 arg0)
func CrossDexImplCrossDexInitializeDataErrorID() common.Hash {
	return common.HexToHash("0xc53a07b94abfc701bef7b9c8d77ba8466e4369b8995a2d5796d69e7c3c2fcfc2")
}

// UnpackCrossDexInitializeDataError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error CrossDexInitializeData(bytes32 arg0)
func (crossDexImpl *CrossDexImpl) UnpackCrossDexInitializeDataError(raw []byte) (*CrossDexImplCrossDexInitializeData, error) {
	out := new(CrossDexImplCrossDexInitializeData)
	if err := crossDexImpl.abi.UnpackIntoInterface(out, "CrossDexInitializeData", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossDexImplCrossDexInvalidMarketAddress represents a CrossDexInvalidMarketAddress error raised by the CrossDexImpl contract.
type CrossDexImplCrossDexInvalidMarketAddress struct {
	Arg0 common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error CrossDexInvalidMarketAddress(address arg0)
func CrossDexImplCrossDexInvalidMarketAddressErrorID() common.Hash {
	return common.HexToHash("0x74be5d0cefe8d064e961c4e0e557948c66ec877bbec4e6cb767d44a87b72a332")
}

// UnpackCrossDexInvalidMarketAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error CrossDexInvalidMarketAddress(address arg0)
func (crossDexImpl *CrossDexImpl) UnpackCrossDexInvalidMarketAddressError(raw []byte) (*CrossDexImplCrossDexInvalidMarketAddress, error) {
	out := new(CrossDexImplCrossDexInvalidMarketAddress)
	if err := crossDexImpl.abi.UnpackIntoInterface(out, "CrossDexInvalidMarketAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossDexImplCrossDexInvalidTickSizeSetter represents a CrossDexInvalidTickSizeSetter error raised by the CrossDexImpl contract.
type CrossDexImplCrossDexInvalidTickSizeSetter struct {
	Current common.Address
	Input   common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error CrossDexInvalidTickSizeSetter(address current, address input)
func CrossDexImplCrossDexInvalidTickSizeSetterErrorID() common.Hash {
	return common.HexToHash("0x18ddd67ebded54806618f9d60db7e0282fa9ae4a9c1139ff7ffde791b19783d1")
}

// UnpackCrossDexInvalidTickSizeSetterError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error CrossDexInvalidTickSizeSetter(address current, address input)
func (crossDexImpl *CrossDexImpl) UnpackCrossDexInvalidTickSizeSetterError(raw []byte) (*CrossDexImplCrossDexInvalidTickSizeSetter, error) {
	out := new(CrossDexImplCrossDexInvalidTickSizeSetter)
	if err := crossDexImpl.abi.UnpackIntoInterface(out, "CrossDexInvalidTickSizeSetter", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossDexImplCrossDexUnauthorizedChangeTickSizes represents a CrossDexUnauthorizedChangeTickSizes error raised by the CrossDexImpl contract.
type CrossDexImplCrossDexUnauthorizedChangeTickSizes struct {
	Arg0 common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error CrossDexUnauthorizedChangeTickSizes(address arg0)
func CrossDexImplCrossDexUnauthorizedChangeTickSizesErrorID() common.Hash {
	return common.HexToHash("0xf42eaafb7cf2208fd186d90528cb751ff79bfcbd79260c8b1448289b8146376a")
}

// UnpackCrossDexUnauthorizedChangeTickSizesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error CrossDexUnauthorizedChangeTickSizes(address arg0)
func (crossDexImpl *CrossDexImpl) UnpackCrossDexUnauthorizedChangeTickSizesError(raw []byte) (*CrossDexImplCrossDexUnauthorizedChangeTickSizes, error) {
	out := new(CrossDexImplCrossDexUnauthorizedChangeTickSizes)
	if err := crossDexImpl.abi.UnpackIntoInterface(out, "CrossDexUnauthorizedChangeTickSizes", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossDexImplERC1967InvalidImplementation represents a ERC1967InvalidImplementation error raised by the CrossDexImpl contract.
type CrossDexImplERC1967InvalidImplementation struct {
	Implementation common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func CrossDexImplERC1967InvalidImplementationErrorID() common.Hash {
	return common.HexToHash("0x4c9c8ce3ceb3130f17f7cdba48d89b5b0129f266a8bac114e6e315a41879b617")
}

// UnpackERC1967InvalidImplementationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func (crossDexImpl *CrossDexImpl) UnpackERC1967InvalidImplementationError(raw []byte) (*CrossDexImplERC1967InvalidImplementation, error) {
	out := new(CrossDexImplERC1967InvalidImplementation)
	if err := crossDexImpl.abi.UnpackIntoInterface(out, "ERC1967InvalidImplementation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossDexImplERC1967NonPayable represents a ERC1967NonPayable error raised by the CrossDexImpl contract.
type CrossDexImplERC1967NonPayable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967NonPayable()
func CrossDexImplERC1967NonPayableErrorID() common.Hash {
	return common.HexToHash("0xb398979fa84f543c8e222f17890372c487baf85e062276c127fef521eea7224b")
}

// UnpackERC1967NonPayableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967NonPayable()
func (crossDexImpl *CrossDexImpl) UnpackERC1967NonPayableError(raw []byte) (*CrossDexImplERC1967NonPayable, error) {
	out := new(CrossDexImplERC1967NonPayable)
	if err := crossDexImpl.abi.UnpackIntoInterface(out, "ERC1967NonPayable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossDexImplEnumerableMapNonexistentKey represents a EnumerableMapNonexistentKey error raised by the CrossDexImpl contract.
type CrossDexImplEnumerableMapNonexistentKey struct {
	Key [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EnumerableMapNonexistentKey(bytes32 key)
func CrossDexImplEnumerableMapNonexistentKeyErrorID() common.Hash {
	return common.HexToHash("0x02b566865b1da2a2deb61443712fe7f812ffd7a1fce56446ff0fe3db9f3484ef")
}

// UnpackEnumerableMapNonexistentKeyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EnumerableMapNonexistentKey(bytes32 key)
func (crossDexImpl *CrossDexImpl) UnpackEnumerableMapNonexistentKeyError(raw []byte) (*CrossDexImplEnumerableMapNonexistentKey, error) {
	out := new(CrossDexImplEnumerableMapNonexistentKey)
	if err := crossDexImpl.abi.UnpackIntoInterface(out, "EnumerableMapNonexistentKey", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossDexImplFailedCall represents a FailedCall error raised by the CrossDexImpl contract.
type CrossDexImplFailedCall struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedCall()
func CrossDexImplFailedCallErrorID() common.Hash {
	return common.HexToHash("0xd6bda27508c0fb6d8a39b4b122878dab26f731a7d4e4abe711dd3731899052a4")
}

// UnpackFailedCallError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedCall()
func (crossDexImpl *CrossDexImpl) UnpackFailedCallError(raw []byte) (*CrossDexImplFailedCall, error) {
	out := new(CrossDexImplFailedCall)
	if err := crossDexImpl.abi.UnpackIntoInterface(out, "FailedCall", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossDexImplFailedDeployment represents a FailedDeployment error raised by the CrossDexImpl contract.
type CrossDexImplFailedDeployment struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedDeployment()
func CrossDexImplFailedDeploymentErrorID() common.Hash {
	return common.HexToHash("0xb06ebf3d5067824a3fe5d5ba19471e035a7de6c88dac362c77b162830a5b9093")
}

// UnpackFailedDeploymentError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedDeployment()
func (crossDexImpl *CrossDexImpl) UnpackFailedDeploymentError(raw []byte) (*CrossDexImplFailedDeployment, error) {
	out := new(CrossDexImplFailedDeployment)
	if err := crossDexImpl.abi.UnpackIntoInterface(out, "FailedDeployment", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossDexImplInsufficientBalance represents a InsufficientBalance error raised by the CrossDexImpl contract.
type CrossDexImplInsufficientBalance struct {
	Balance *big.Int
	Needed  *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InsufficientBalance(uint256 balance, uint256 needed)
func CrossDexImplInsufficientBalanceErrorID() common.Hash {
	return common.HexToHash("0xcf4791818fba6e019216eb4864093b4947f674afada5d305e57d598b641dad1d")
}

// UnpackInsufficientBalanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InsufficientBalance(uint256 balance, uint256 needed)
func (crossDexImpl *CrossDexImpl) UnpackInsufficientBalanceError(raw []byte) (*CrossDexImplInsufficientBalance, error) {
	out := new(CrossDexImplInsufficientBalance)
	if err := crossDexImpl.abi.UnpackIntoInterface(out, "InsufficientBalance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossDexImplInvalidInitialization represents a InvalidInitialization error raised by the CrossDexImpl contract.
type CrossDexImplInvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func CrossDexImplInvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (crossDexImpl *CrossDexImpl) UnpackInvalidInitializationError(raw []byte) (*CrossDexImplInvalidInitialization, error) {
	out := new(CrossDexImplInvalidInitialization)
	if err := crossDexImpl.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossDexImplNotInitializing represents a NotInitializing error raised by the CrossDexImpl contract.
type CrossDexImplNotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func CrossDexImplNotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (crossDexImpl *CrossDexImpl) UnpackNotInitializingError(raw []byte) (*CrossDexImplNotInitializing, error) {
	out := new(CrossDexImplNotInitializing)
	if err := crossDexImpl.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossDexImplOwnableInvalidOwner represents a OwnableInvalidOwner error raised by the CrossDexImpl contract.
type CrossDexImplOwnableInvalidOwner struct {
	Owner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OwnableInvalidOwner(address owner)
func CrossDexImplOwnableInvalidOwnerErrorID() common.Hash {
	return common.HexToHash("0x1e4fbdf7f3ef8bcaa855599e3abf48b232380f183f08f6f813d9ffa5bd585188")
}

// UnpackOwnableInvalidOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OwnableInvalidOwner(address owner)
func (crossDexImpl *CrossDexImpl) UnpackOwnableInvalidOwnerError(raw []byte) (*CrossDexImplOwnableInvalidOwner, error) {
	out := new(CrossDexImplOwnableInvalidOwner)
	if err := crossDexImpl.abi.UnpackIntoInterface(out, "OwnableInvalidOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossDexImplOwnableUnauthorizedAccount represents a OwnableUnauthorizedAccount error raised by the CrossDexImpl contract.
type CrossDexImplOwnableUnauthorizedAccount struct {
	Account common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OwnableUnauthorizedAccount(address account)
func CrossDexImplOwnableUnauthorizedAccountErrorID() common.Hash {
	return common.HexToHash("0x118cdaa7a341953d1887a2245fd6665d741c67c8c50581daa59e1d03373fa188")
}

// UnpackOwnableUnauthorizedAccountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OwnableUnauthorizedAccount(address account)
func (crossDexImpl *CrossDexImpl) UnpackOwnableUnauthorizedAccountError(raw []byte) (*CrossDexImplOwnableUnauthorizedAccount, error) {
	out := new(CrossDexImplOwnableUnauthorizedAccount)
	if err := crossDexImpl.abi.UnpackIntoInterface(out, "OwnableUnauthorizedAccount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossDexImplUUPSUnauthorizedCallContext represents a UUPSUnauthorizedCallContext error raised by the CrossDexImpl contract.
type CrossDexImplUUPSUnauthorizedCallContext struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnauthorizedCallContext()
func CrossDexImplUUPSUnauthorizedCallContextErrorID() common.Hash {
	return common.HexToHash("0xe07c8dba242a06571ac65fe4bbe20522c9fb111cb33599b799ff8039c1ed18f4")
}

// UnpackUUPSUnauthorizedCallContextError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnauthorizedCallContext()
func (crossDexImpl *CrossDexImpl) UnpackUUPSUnauthorizedCallContextError(raw []byte) (*CrossDexImplUUPSUnauthorizedCallContext, error) {
	out := new(CrossDexImplUUPSUnauthorizedCallContext)
	if err := crossDexImpl.abi.UnpackIntoInterface(out, "UUPSUnauthorizedCallContext", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossDexImplUUPSUnsupportedProxiableUUID represents a UUPSUnsupportedProxiableUUID error raised by the CrossDexImpl contract.
type CrossDexImplUUPSUnsupportedProxiableUUID struct {
	Slot [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func CrossDexImplUUPSUnsupportedProxiableUUIDErrorID() common.Hash {
	return common.HexToHash("0xaa1d49a4c084bfa9aeeee2a0be65267a7f19ba7e1476b114dac513d2c14cb563")
}

// UnpackUUPSUnsupportedProxiableUUIDError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func (crossDexImpl *CrossDexImpl) UnpackUUPSUnsupportedProxiableUUIDError(raw []byte) (*CrossDexImplUUPSUnsupportedProxiableUUID, error) {
	out := new(CrossDexImplUUPSUnsupportedProxiableUUID)
	if err := crossDexImpl.abi.UnpackIntoInterface(out, "UUPSUnsupportedProxiableUUID", raw); err != nil {
		return nil, err
	}
	return out, nil
}
