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

// CrossDexImplV2MetaData contains all meta data concerning the CrossDexImplV2 contract.
var CrossDexImplV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ROUTER\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allMarkets\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"quotes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"markets\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"checkTickSizeRoles\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quote\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeCollector\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"createMarket\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_routerImpl\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_findPrevPriceCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxMatchCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_cancelLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_marketImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pairImpl\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"}],\"name\":\"isMarket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketImpl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"name\":\"pairCreated\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pairImpl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"name\":\"pairToMarket\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"quote\",\"type\":\"address\"}],\"name\":\"quoteToMarket\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_marketImpl\",\"type\":\"address\"}],\"name\":\"setMarketImpl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pairImpl\",\"type\":\"address\"}],\"name\":\"setPairImpl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"setter\",\"type\":\"address\"}],\"name\":\"setTickSizeSetter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tickSizeSetter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"quote\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fee_collector\",\"type\":\"address\"}],\"name\":\"MarketCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"before\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"MarketImplSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"before\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"PairImplSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"before\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"TickSizeSetterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Create2EmptyBytecode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"CrossDexAlreadyCreatedMarketQuote\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"CrossDexInitializeData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"CrossDexInvalidMarketAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"input\",\"type\":\"address\"}],\"name\":\"CrossDexInvalidTickSizeSetter\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"CrossDexUnauthorizedChangeTickSizes\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"EnumerableMapNonexistentKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedDeployment\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"32fe7b26": "ROUTER()",
		"ad3cb1cc": "UPGRADE_INTERFACE_VERSION()",
		"375a7cba": "allMarkets()",
		"a1eea778": "checkTickSizeRoles(address)",
		"9967b84e": "createMarket(address,address,address,bytes)",
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
		"e492c270": "setMarketImpl(address)",
		"3cc047aa": "setPairImpl(address)",
		"98715ee7": "setTickSizeSetter(address)",
		"7f764a85": "tickSizeSetter()",
		"f2fde38b": "transferOwnership(address)",
		"4f1ef286": "upgradeToAndCall(address,bytes)",
	},
	Bin: "0x60a060405230608052348015610013575f5ffd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6080516125916100f95f395f81816113ce015281816113f7015261161801526125915ff3fe608060405260043610610162575f3560e01c8063715018a6116100c6578063a1eea7781161007c578063e492c27011610057578063e492c27014610474578063e9f7206b14610493578063f2fde38b146104b2575f5ffd5b8063a1eea778146103e1578063ad3cb1cc14610400578063beb380f114610455575f5ffd5b80638da5cb5b116100ac5780638da5cb5b1461035a57806398715ee7146103a35780639967b84e146103c2575f5ffd5b8063715018a61461031a5780637f764a851461032e575f5ffd5b8063375a7cba1161011b5780634f1ef286116101015780634f1ef286146102b657806352d1902d146102c95780636ec934da146102eb575f5ffd5b8063375a7cba146102755780633cc047aa14610297575f5ffd5b806332fe7b261161014b57806332fe7b26146101f257806334eaeeb91461021d57806335f7d45614610249575f5ffd5b80630729da0b146101665780630827057314610187575b5f5ffd5b348015610171575f5ffd5b50610185610180366004611cfb565b6104d1565b005b348015610192575f5ffd5b506101c86101a1366004611d71565b60066020525f908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b3480156101fd575f5ffd5b505f546101c89073ffffffffffffffffffffffffffffffffffffffff1681565b348015610228575f5ffd5b506001546101c89073ffffffffffffffffffffffffffffffffffffffff1681565b348015610254575f5ffd5b506002546101c89073ffffffffffffffffffffffffffffffffffffffff1681565b348015610280575f5ffd5b50610289610908565b6040516101e9929190611ddc565b3480156102a2575f5ffd5b506101856102b1366004611d71565b610a17565b6101856102c4366004611ef5565b610b1b565b3480156102d4575f5ffd5b506102dd610b3a565b6040519081526020016101e9565b3480156102f6575f5ffd5b5061030a610305366004611d71565b610b68565b60405190151581526020016101e9565b348015610325575f5ffd5b50610185610c91565b348015610339575f5ffd5b506007546101c89073ffffffffffffffffffffffffffffffffffffffff1681565b348015610365575f5ffd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c1993005473ffffffffffffffffffffffffffffffffffffffff166101c8565b3480156103ae575f5ffd5b506101856103bd366004611d71565b610ca4565b3480156103cd575f5ffd5b506101c86103dc366004611f42565b610dd1565b3480156103ec575f5ffd5b506101856103fb366004611d71565b61105c565b34801561040b575f5ffd5b506104486040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b6040516101e99190611fff565b348015610460575f5ffd5b506101c861046f366004611d71565b6110ee565b34801561047f575f5ffd5b5061018561048e366004611d71565b611100565b34801561049e575f5ffd5b506101856104ad366004611d71565b611204565b3480156104bd575f5ffd5b506101856104cc366004611d71565b611292565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff165f8115801561051b5750825b90505f8267ffffffffffffffff1660011480156105375750303b155b905081158015610545575080155b1561057c576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000016600117855583156105dd5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b6105e68c6112f2565b73ffffffffffffffffffffffffffffffffffffffff8b1661065a576040517fc53a07b90000000000000000000000000000000000000000000000000000000081527f726f75746572496d706c0000000000000000000000000000000000000000000060048201526024015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff87166106c9576040517fc53a07b90000000000000000000000000000000000000000000000000000000081527f6d61726b6574496d706c000000000000000000000000000000000000000000006004820152602401610651565b73ffffffffffffffffffffffffffffffffffffffff8616610738576040517fc53a07b90000000000000000000000000000000000000000000000000000000081527f70616972496d706c0000000000000000000000000000000000000000000000006004820152602401610651565b5f8b60405161074690611ccd565b73ffffffffffffffffffffffffffffffffffffffff90911681526040602082018190525f90820152606001604051809103905ff08015801561078a573d5f5f3e3d5ffd5b505f80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040517f80d85911000000000000000000000000000000000000000000000000000000008152600481018e9052602481018d9052604481018c9052919250906380d85911906064015f604051808303815f87803b15801561082c575f5ffd5b505af115801561083e573d5f5f3e3d5ffd5b50506001805473ffffffffffffffffffffffffffffffffffffffff808d167fffffffffffffffffffffffff00000000000000000000000000000000000000009283161790925560028054928c16929091169190911790555050841590506108fa5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050505050505050565b6060805f6109166003611303565b90508067ffffffffffffffff81111561093157610931611e00565b60405190808252806020026020018201604052801561095a578160200160208202803683370190505b5092508067ffffffffffffffff81111561097657610976611e00565b60405190808252806020026020018201604052801561099f578160200160208202803683370190505b5091505f5b81811015610a11576109b760038261130d565b8583815181106109c9576109c9612011565b602002602001018584815181106109e2576109e2612011565b73ffffffffffffffffffffffffffffffffffffffff9384166020918202929092010152911690526001016109a4565b50509091565b610a1f611328565b73ffffffffffffffffffffffffffffffffffffffff8116610a8e576040517fc53a07b90000000000000000000000000000000000000000000000000000000081527f70616972496d706c0000000000000000000000000000000000000000000000006004820152602401610651565b60025460405173ffffffffffffffffffffffffffffffffffffffff8084169216907f685c330e9b04243f99ff39d8de8578c9af2bfaff120e17a1b5f1d61e004935e2905f90a3600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b610b236113b6565b610b2c826114ba565b610b3682826114c2565b5050565b5f610b43611600565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b60408051600481526024810182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f9c5798390000000000000000000000000000000000000000000000000000000017905290515f918291829173ffffffffffffffffffffffffffffffffffffffff861691610be89190612055565b5f60405180830381855afa9150503d805f8114610c20576040519150601f19603f3d011682016040523d82523d5f602084013e610c25565b606091505b509150915081610c3857505f9392505050565b5f81806020019051810190610c4d9190612060565b905073ffffffffffffffffffffffffffffffffffffffff8516610c7160038361166f565b73ffffffffffffffffffffffffffffffffffffffff161495945050505050565b610c99611328565b610ca25f611697565b565b610cac611328565b73ffffffffffffffffffffffffffffffffffffffff81161580610ce9575060075473ffffffffffffffffffffffffffffffffffffffff8281169116145b15610d44576007546040517f18ddd67e00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff91821660048201529082166024820152604401610651565b60075460405173ffffffffffffffffffffffffffffffffffffffff8084169216907f4f1656ebcdda14dfa0f9951b46485570075b9730e1f1d22cb4ddacc35374fde8905f90a3600780547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b5f610dda611328565b5f60405180602001610deb90611ccd565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082820381018352601f9091011660408190526001545f5460025473ffffffffffffffffffffffffffffffffffffffff92831693610e59938c938116928c929116908b908b9060240161207b565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f01f79fb6000000000000000000000000000000000000000000000000000000001790529051610ede93929101612125565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815290829052610f1a9291602001612153565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606088901b16602083015291505f906034016040516020818303038152906040528051906020012090505f610fa05f838561172c565b9050610fae60038883611818565b610ffc576040517fed703a9800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff88166004820152602401610651565b60405173ffffffffffffffffffffffffffffffffffffffff8781168252808a1691838216918a16907f5509f6235f0d08c563aedf1b655226cad4a613794902c476ad7a08accb43ebc59060200160405180910390a4979650505050505050565b60075473ffffffffffffffffffffffffffffffffffffffff16158061109c575060075473ffffffffffffffffffffffffffffffffffffffff828116911614155b156110eb576040517ff42eaafb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152602401610651565b50565b5f6110fa60038361166f565b92915050565b611108611328565b73ffffffffffffffffffffffffffffffffffffffff8116611177576040517fc53a07b90000000000000000000000000000000000000000000000000000000081527f6d61726b6574496d706c000000000000000000000000000000000000000000006004820152602401610651565b60015460405173ffffffffffffffffffffffffffffffffffffffff8084169216907f081eadab9f9e460b348376155bdd3b4aed9b2a33469b489416d7416b4cffe49c905f90a3600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b61120d33610b68565b611245576040517f74be5d0c000000000000000000000000000000000000000000000000000000008152336004820152602401610651565b73ffffffffffffffffffffffffffffffffffffffff165f90815260066020526040902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001633179055565b61129a611328565b73ffffffffffffffffffffffffffffffffffffffff81166112e9576040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081525f6004820152602401610651565b6110eb81611697565b6112fa611845565b6110eb816118ac565b5f6110fa826118b4565b5f80808061131b86866118be565b9097909650945050505050565b336113677f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c1993005473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614610ca2576040517f118cdaa7000000000000000000000000000000000000000000000000000000008152336004820152602401610651565b3073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016148061148357507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1661146a7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614155b15610ca2576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6110eb611328565b8173ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015611547575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016820190925261154491810190612167565b60015b611595576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610651565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc81146115f1576040517faa1d49a400000000000000000000000000000000000000000000000000000000815260048101829052602401610651565b6115fb83836118e7565b505050565b3073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610ca2576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6116908373ffffffffffffffffffffffffffffffffffffffff8416611949565b9392505050565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080547fffffffffffffffffffffffff0000000000000000000000000000000000000000811673ffffffffffffffffffffffffffffffffffffffff848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b5f83471015611770576040517fcf47918100000000000000000000000000000000000000000000000000000000815247600482015260248101859052604401610651565b81515f036117aa576040517f4ca249dc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8282516020840186f590503d1519811516156117cb576040513d5f823e3d81fd5b73ffffffffffffffffffffffffffffffffffffffff8116611690576040517fb06ebf3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f61183d8473ffffffffffffffffffffffffffffffffffffffff8086169085166119a8565b949350505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff16610ca2576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61129a611845565b5f6110fa826119c4565b5f80806118cb85856119cd565b5f81815260029690960160205260409095205494959350505050565b6118f0826119d8565b60405173ffffffffffffffffffffffffffffffffffffffff8316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115611941576115fb8282611aa6565b610b36611b25565b5f8181526002830160205260408120548015801561196e575061196c8484611b5d565b155b15611690576040517f02b5668600000000000000000000000000000000000000000000000000000000815260048101849052602401610651565b5f828152600284016020526040812082905561183d8484611b68565b5f6110fa825490565b5f6116908383611b73565b8073ffffffffffffffffffffffffffffffffffffffff163b5f03611a40576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152602401610651565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60605f5f8473ffffffffffffffffffffffffffffffffffffffff1684604051611acf9190612055565b5f60405180830381855af49150503d805f8114611b07576040519150601f19603f3d011682016040523d82523d5f602084013e611b0c565b606091505b5091509150611b1c858383611b99565b95945050505050565b3415610ca2576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6116908383611c28565b5f6116908383611c3f565b5f825f018281548110611b8857611b88612011565b905f5260205f200154905092915050565b606082611bae57611ba982611c8b565b611690565b8151158015611bd2575073ffffffffffffffffffffffffffffffffffffffff84163b155b15611c21576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610651565b5080611690565b5f8181526001830160205260408120541515611690565b5f818152600183016020526040812054611c8457508154600181810184555f8481526020808220909301849055845484825282860190935260409020919091556110fa565b505f6110fa565b805115611c9b5780518082602001fd5b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103dd8061217f83390190565b73ffffffffffffffffffffffffffffffffffffffff811681146110eb575f5ffd5b5f5f5f5f5f5f5f60e0888a031215611d11575f5ffd5b8735611d1c81611cda565b96506020880135611d2c81611cda565b955060408801359450606088013593506080880135925060a0880135611d5181611cda565b915060c0880135611d6181611cda565b8091505092959891949750929550565b5f60208284031215611d81575f5ffd5b813561169081611cda565b5f8151808452602084019350602083015f5b82811015611dd257815173ffffffffffffffffffffffffffffffffffffffff16865260209586019590910190600101611d9e565b5093949350505050565b604081525f611dee6040830185611d8c565b8281036020840152611b1c8185611d8c565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f82601f830112611e3c575f5ffd5b813567ffffffffffffffff811115611e5657611e56611e00565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff82111715611ec257611ec2611e00565b604052818152838201602001851015611ed9575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f5f60408385031215611f06575f5ffd5b8235611f1181611cda565b9150602083013567ffffffffffffffff811115611f2c575f5ffd5b611f3885828601611e2d565b9150509250929050565b5f5f5f5f60808587031215611f55575f5ffd5b8435611f6081611cda565b93506020850135611f7081611cda565b92506040850135611f8081611cda565b9150606085013567ffffffffffffffff811115611f9b575f5ffd5b611fa787828801611e2d565b91505092959194509250565b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f6116906020830184611fb3565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f81518060208401855e5f93019283525090919050565b5f611690828461203e565b5f60208284031215612070575f5ffd5b815161169081611cda565b73ffffffffffffffffffffffffffffffffffffffff8716815273ffffffffffffffffffffffffffffffffffffffff8616602082015273ffffffffffffffffffffffffffffffffffffffff8516604082015273ffffffffffffffffffffffffffffffffffffffff8416606082015273ffffffffffffffffffffffffffffffffffffffff8316608082015260c060a08201525f61211960c0830184611fb3565b98975050505050505050565b73ffffffffffffffffffffffffffffffffffffffff83168152604060208201525f61183d6040830184611fb3565b5f61183d612161838661203e565b8461203e565b5f60208284031215612177575f5ffd5b505191905056fe60806040526040516103dd3803806103dd8339810160408190526100229161023c565b61002c8282610033565b5050610321565b61003c82610091565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a280511561008557610080828261010c565b505050565b61008d61017f565b5050565b806001600160a01b03163b5f036100cb57604051634c9c8ce360e01b81526001600160a01b03821660048201526024015b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b031684604051610128919061030b565b5f60405180830381855af49150503d805f8114610160576040519150601f19603f3d011682016040523d82523d5f602084013e610165565b606091505b5090925090506101768583836101a0565b95945050505050565b341561019e5760405163b398979f60e01b815260040160405180910390fd5b565b6060826101b5576101b0826101ff565b6101f8565b81511580156101cc57506001600160a01b0384163b155b156101f557604051639996b31560e01b81526001600160a01b03851660048201526024016100c2565b50805b9392505050565b80511561020f5780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b634e487b7160e01b5f52604160045260245ffd5b5f5f6040838503121561024d575f5ffd5b82516001600160a01b0381168114610263575f5ffd5b60208401519092506001600160401b0381111561027e575f5ffd5b8301601f8101851361028e575f5ffd5b80516001600160401b038111156102a7576102a7610228565b604051601f8201601f19908116603f011681016001600160401b03811182821017156102d5576102d5610228565b6040528181528282016020018710156102ec575f5ffd5b8160208401602083015e5f602083830101528093505050509250929050565b5f82518060208501845e5f920191825250919050565b60b08061032d5f395ff3fe6080604052600a600c565b005b60186014601a565b605d565b565b5f60587f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5473ffffffffffffffffffffffffffffffffffffffff1690565b905090565b365f5f375f5f365f845af43d5f5f3e8080156076573d5ff35b3d5ffdfea2646970667358221220f53cdd673e6d446a26f162ed3adcc3948e6ec30e5c1918fa9cfbf44a3f745dd964736f6c634300081c0033a2646970667358221220f42b73ae51dbeba5454b10a6c721ac4e26bb0daf9d9f1d2423e2d478c383f65964736f6c634300081c0033",
}

// CrossDexImplV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use CrossDexImplV2MetaData.ABI instead.
var CrossDexImplV2ABI = CrossDexImplV2MetaData.ABI

// Deprecated: Use CrossDexImplV2MetaData.Sigs instead.
// CrossDexImplV2FuncSigs maps the 4-byte function signature to its string representation.
var CrossDexImplV2FuncSigs = CrossDexImplV2MetaData.Sigs

// CrossDexImplV2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CrossDexImplV2MetaData.Bin instead.
var CrossDexImplV2Bin = CrossDexImplV2MetaData.Bin

// DeployCrossDexImplV2 deploys a new Ethereum contract, binding an instance of CrossDexImplV2 to it.
func DeployCrossDexImplV2(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CrossDexImplV2, error) {
	parsed, err := CrossDexImplV2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CrossDexImplV2Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CrossDexImplV2{CrossDexImplV2Caller: CrossDexImplV2Caller{contract: contract}, CrossDexImplV2Transactor: CrossDexImplV2Transactor{contract: contract}, CrossDexImplV2Filterer: CrossDexImplV2Filterer{contract: contract}}, nil
}

// CrossDexImplV2 is an auto generated Go binding around an Ethereum contract.
type CrossDexImplV2 struct {
	CrossDexImplV2Caller     // Read-only binding to the contract
	CrossDexImplV2Transactor // Write-only binding to the contract
	CrossDexImplV2Filterer   // Log filterer for contract events
}

// CrossDexImplV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type CrossDexImplV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossDexImplV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type CrossDexImplV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossDexImplV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrossDexImplV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossDexImplV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrossDexImplV2Session struct {
	Contract     *CrossDexImplV2   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrossDexImplV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrossDexImplV2CallerSession struct {
	Contract *CrossDexImplV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// CrossDexImplV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrossDexImplV2TransactorSession struct {
	Contract     *CrossDexImplV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// CrossDexImplV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type CrossDexImplV2Raw struct {
	Contract *CrossDexImplV2 // Generic contract binding to access the raw methods on
}

// CrossDexImplV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrossDexImplV2CallerRaw struct {
	Contract *CrossDexImplV2Caller // Generic read-only contract binding to access the raw methods on
}

// CrossDexImplV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrossDexImplV2TransactorRaw struct {
	Contract *CrossDexImplV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewCrossDexImplV2 creates a new instance of CrossDexImplV2, bound to a specific deployed contract.
func NewCrossDexImplV2(address common.Address, backend bind.ContractBackend) (*CrossDexImplV2, error) {
	contract, err := bindCrossDexImplV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrossDexImplV2{CrossDexImplV2Caller: CrossDexImplV2Caller{contract: contract}, CrossDexImplV2Transactor: CrossDexImplV2Transactor{contract: contract}, CrossDexImplV2Filterer: CrossDexImplV2Filterer{contract: contract}}, nil
}

// NewCrossDexImplV2Caller creates a new read-only instance of CrossDexImplV2, bound to a specific deployed contract.
func NewCrossDexImplV2Caller(address common.Address, caller bind.ContractCaller) (*CrossDexImplV2Caller, error) {
	contract, err := bindCrossDexImplV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrossDexImplV2Caller{contract: contract}, nil
}

// NewCrossDexImplV2Transactor creates a new write-only instance of CrossDexImplV2, bound to a specific deployed contract.
func NewCrossDexImplV2Transactor(address common.Address, transactor bind.ContractTransactor) (*CrossDexImplV2Transactor, error) {
	contract, err := bindCrossDexImplV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrossDexImplV2Transactor{contract: contract}, nil
}

// NewCrossDexImplV2Filterer creates a new log filterer instance of CrossDexImplV2, bound to a specific deployed contract.
func NewCrossDexImplV2Filterer(address common.Address, filterer bind.ContractFilterer) (*CrossDexImplV2Filterer, error) {
	contract, err := bindCrossDexImplV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrossDexImplV2Filterer{contract: contract}, nil
}

// bindCrossDexImplV2 binds a generic wrapper to an already deployed contract.
func bindCrossDexImplV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CrossDexImplV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossDexImplV2 *CrossDexImplV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossDexImplV2.Contract.CrossDexImplV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossDexImplV2 *CrossDexImplV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossDexImplV2.Contract.CrossDexImplV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossDexImplV2 *CrossDexImplV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossDexImplV2.Contract.CrossDexImplV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossDexImplV2 *CrossDexImplV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossDexImplV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossDexImplV2 *CrossDexImplV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossDexImplV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossDexImplV2 *CrossDexImplV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossDexImplV2.Contract.contract.Transact(opts, method, params...)
}

// ROUTER is a free data retrieval call binding the contract method 0x32fe7b26.
//
// Solidity: function ROUTER() view returns(address)
func (_CrossDexImplV2 *CrossDexImplV2Caller) ROUTER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossDexImplV2.contract.Call(opts, &out, "ROUTER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ROUTER is a free data retrieval call binding the contract method 0x32fe7b26.
//
// Solidity: function ROUTER() view returns(address)
func (_CrossDexImplV2 *CrossDexImplV2Session) ROUTER() (common.Address, error) {
	return _CrossDexImplV2.Contract.ROUTER(&_CrossDexImplV2.CallOpts)
}

// ROUTER is a free data retrieval call binding the contract method 0x32fe7b26.
//
// Solidity: function ROUTER() view returns(address)
func (_CrossDexImplV2 *CrossDexImplV2CallerSession) ROUTER() (common.Address, error) {
	return _CrossDexImplV2.Contract.ROUTER(&_CrossDexImplV2.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_CrossDexImplV2 *CrossDexImplV2Caller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CrossDexImplV2.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_CrossDexImplV2 *CrossDexImplV2Session) UPGRADEINTERFACEVERSION() (string, error) {
	return _CrossDexImplV2.Contract.UPGRADEINTERFACEVERSION(&_CrossDexImplV2.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_CrossDexImplV2 *CrossDexImplV2CallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _CrossDexImplV2.Contract.UPGRADEINTERFACEVERSION(&_CrossDexImplV2.CallOpts)
}

// AllMarkets is a free data retrieval call binding the contract method 0x375a7cba.
//
// Solidity: function allMarkets() view returns(address[] quotes, address[] markets)
func (_CrossDexImplV2 *CrossDexImplV2Caller) AllMarkets(opts *bind.CallOpts) (struct {
	Quotes  []common.Address
	Markets []common.Address
}, error) {
	var out []interface{}
	err := _CrossDexImplV2.contract.Call(opts, &out, "allMarkets")

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
func (_CrossDexImplV2 *CrossDexImplV2Session) AllMarkets() (struct {
	Quotes  []common.Address
	Markets []common.Address
}, error) {
	return _CrossDexImplV2.Contract.AllMarkets(&_CrossDexImplV2.CallOpts)
}

// AllMarkets is a free data retrieval call binding the contract method 0x375a7cba.
//
// Solidity: function allMarkets() view returns(address[] quotes, address[] markets)
func (_CrossDexImplV2 *CrossDexImplV2CallerSession) AllMarkets() (struct {
	Quotes  []common.Address
	Markets []common.Address
}, error) {
	return _CrossDexImplV2.Contract.AllMarkets(&_CrossDexImplV2.CallOpts)
}

// CheckTickSizeRoles is a free data retrieval call binding the contract method 0xa1eea778.
//
// Solidity: function checkTickSizeRoles(address account) view returns()
func (_CrossDexImplV2 *CrossDexImplV2Caller) CheckTickSizeRoles(opts *bind.CallOpts, account common.Address) error {
	var out []interface{}
	err := _CrossDexImplV2.contract.Call(opts, &out, "checkTickSizeRoles", account)

	if err != nil {
		return err
	}

	return err

}

// CheckTickSizeRoles is a free data retrieval call binding the contract method 0xa1eea778.
//
// Solidity: function checkTickSizeRoles(address account) view returns()
func (_CrossDexImplV2 *CrossDexImplV2Session) CheckTickSizeRoles(account common.Address) error {
	return _CrossDexImplV2.Contract.CheckTickSizeRoles(&_CrossDexImplV2.CallOpts, account)
}

// CheckTickSizeRoles is a free data retrieval call binding the contract method 0xa1eea778.
//
// Solidity: function checkTickSizeRoles(address account) view returns()
func (_CrossDexImplV2 *CrossDexImplV2CallerSession) CheckTickSizeRoles(account common.Address) error {
	return _CrossDexImplV2.Contract.CheckTickSizeRoles(&_CrossDexImplV2.CallOpts, account)
}

// IsMarket is a free data retrieval call binding the contract method 0x6ec934da.
//
// Solidity: function isMarket(address market) view returns(bool)
func (_CrossDexImplV2 *CrossDexImplV2Caller) IsMarket(opts *bind.CallOpts, market common.Address) (bool, error) {
	var out []interface{}
	err := _CrossDexImplV2.contract.Call(opts, &out, "isMarket", market)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMarket is a free data retrieval call binding the contract method 0x6ec934da.
//
// Solidity: function isMarket(address market) view returns(bool)
func (_CrossDexImplV2 *CrossDexImplV2Session) IsMarket(market common.Address) (bool, error) {
	return _CrossDexImplV2.Contract.IsMarket(&_CrossDexImplV2.CallOpts, market)
}

// IsMarket is a free data retrieval call binding the contract method 0x6ec934da.
//
// Solidity: function isMarket(address market) view returns(bool)
func (_CrossDexImplV2 *CrossDexImplV2CallerSession) IsMarket(market common.Address) (bool, error) {
	return _CrossDexImplV2.Contract.IsMarket(&_CrossDexImplV2.CallOpts, market)
}

// MarketImpl is a free data retrieval call binding the contract method 0x34eaeeb9.
//
// Solidity: function marketImpl() view returns(address)
func (_CrossDexImplV2 *CrossDexImplV2Caller) MarketImpl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossDexImplV2.contract.Call(opts, &out, "marketImpl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MarketImpl is a free data retrieval call binding the contract method 0x34eaeeb9.
//
// Solidity: function marketImpl() view returns(address)
func (_CrossDexImplV2 *CrossDexImplV2Session) MarketImpl() (common.Address, error) {
	return _CrossDexImplV2.Contract.MarketImpl(&_CrossDexImplV2.CallOpts)
}

// MarketImpl is a free data retrieval call binding the contract method 0x34eaeeb9.
//
// Solidity: function marketImpl() view returns(address)
func (_CrossDexImplV2 *CrossDexImplV2CallerSession) MarketImpl() (common.Address, error) {
	return _CrossDexImplV2.Contract.MarketImpl(&_CrossDexImplV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CrossDexImplV2 *CrossDexImplV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossDexImplV2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CrossDexImplV2 *CrossDexImplV2Session) Owner() (common.Address, error) {
	return _CrossDexImplV2.Contract.Owner(&_CrossDexImplV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CrossDexImplV2 *CrossDexImplV2CallerSession) Owner() (common.Address, error) {
	return _CrossDexImplV2.Contract.Owner(&_CrossDexImplV2.CallOpts)
}

// PairImpl is a free data retrieval call binding the contract method 0x35f7d456.
//
// Solidity: function pairImpl() view returns(address)
func (_CrossDexImplV2 *CrossDexImplV2Caller) PairImpl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossDexImplV2.contract.Call(opts, &out, "pairImpl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PairImpl is a free data retrieval call binding the contract method 0x35f7d456.
//
// Solidity: function pairImpl() view returns(address)
func (_CrossDexImplV2 *CrossDexImplV2Session) PairImpl() (common.Address, error) {
	return _CrossDexImplV2.Contract.PairImpl(&_CrossDexImplV2.CallOpts)
}

// PairImpl is a free data retrieval call binding the contract method 0x35f7d456.
//
// Solidity: function pairImpl() view returns(address)
func (_CrossDexImplV2 *CrossDexImplV2CallerSession) PairImpl() (common.Address, error) {
	return _CrossDexImplV2.Contract.PairImpl(&_CrossDexImplV2.CallOpts)
}

// PairToMarket is a free data retrieval call binding the contract method 0x08270573.
//
// Solidity: function pairToMarket(address pair) view returns(address)
func (_CrossDexImplV2 *CrossDexImplV2Caller) PairToMarket(opts *bind.CallOpts, pair common.Address) (common.Address, error) {
	var out []interface{}
	err := _CrossDexImplV2.contract.Call(opts, &out, "pairToMarket", pair)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PairToMarket is a free data retrieval call binding the contract method 0x08270573.
//
// Solidity: function pairToMarket(address pair) view returns(address)
func (_CrossDexImplV2 *CrossDexImplV2Session) PairToMarket(pair common.Address) (common.Address, error) {
	return _CrossDexImplV2.Contract.PairToMarket(&_CrossDexImplV2.CallOpts, pair)
}

// PairToMarket is a free data retrieval call binding the contract method 0x08270573.
//
// Solidity: function pairToMarket(address pair) view returns(address)
func (_CrossDexImplV2 *CrossDexImplV2CallerSession) PairToMarket(pair common.Address) (common.Address, error) {
	return _CrossDexImplV2.Contract.PairToMarket(&_CrossDexImplV2.CallOpts, pair)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CrossDexImplV2 *CrossDexImplV2Caller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CrossDexImplV2.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CrossDexImplV2 *CrossDexImplV2Session) ProxiableUUID() ([32]byte, error) {
	return _CrossDexImplV2.Contract.ProxiableUUID(&_CrossDexImplV2.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CrossDexImplV2 *CrossDexImplV2CallerSession) ProxiableUUID() ([32]byte, error) {
	return _CrossDexImplV2.Contract.ProxiableUUID(&_CrossDexImplV2.CallOpts)
}

// QuoteToMarket is a free data retrieval call binding the contract method 0xbeb380f1.
//
// Solidity: function quoteToMarket(address quote) view returns(address)
func (_CrossDexImplV2 *CrossDexImplV2Caller) QuoteToMarket(opts *bind.CallOpts, quote common.Address) (common.Address, error) {
	var out []interface{}
	err := _CrossDexImplV2.contract.Call(opts, &out, "quoteToMarket", quote)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// QuoteToMarket is a free data retrieval call binding the contract method 0xbeb380f1.
//
// Solidity: function quoteToMarket(address quote) view returns(address)
func (_CrossDexImplV2 *CrossDexImplV2Session) QuoteToMarket(quote common.Address) (common.Address, error) {
	return _CrossDexImplV2.Contract.QuoteToMarket(&_CrossDexImplV2.CallOpts, quote)
}

// QuoteToMarket is a free data retrieval call binding the contract method 0xbeb380f1.
//
// Solidity: function quoteToMarket(address quote) view returns(address)
func (_CrossDexImplV2 *CrossDexImplV2CallerSession) QuoteToMarket(quote common.Address) (common.Address, error) {
	return _CrossDexImplV2.Contract.QuoteToMarket(&_CrossDexImplV2.CallOpts, quote)
}

// TickSizeSetter is a free data retrieval call binding the contract method 0x7f764a85.
//
// Solidity: function tickSizeSetter() view returns(address)
func (_CrossDexImplV2 *CrossDexImplV2Caller) TickSizeSetter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossDexImplV2.contract.Call(opts, &out, "tickSizeSetter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TickSizeSetter is a free data retrieval call binding the contract method 0x7f764a85.
//
// Solidity: function tickSizeSetter() view returns(address)
func (_CrossDexImplV2 *CrossDexImplV2Session) TickSizeSetter() (common.Address, error) {
	return _CrossDexImplV2.Contract.TickSizeSetter(&_CrossDexImplV2.CallOpts)
}

// TickSizeSetter is a free data retrieval call binding the contract method 0x7f764a85.
//
// Solidity: function tickSizeSetter() view returns(address)
func (_CrossDexImplV2 *CrossDexImplV2CallerSession) TickSizeSetter() (common.Address, error) {
	return _CrossDexImplV2.Contract.TickSizeSetter(&_CrossDexImplV2.CallOpts)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x9967b84e.
//
// Solidity: function createMarket(address _owner, address quote, address feeCollector, bytes data) returns(address)
func (_CrossDexImplV2 *CrossDexImplV2Transactor) CreateMarket(opts *bind.TransactOpts, _owner common.Address, quote common.Address, feeCollector common.Address, data []byte) (*types.Transaction, error) {
	return _CrossDexImplV2.contract.Transact(opts, "createMarket", _owner, quote, feeCollector, data)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x9967b84e.
//
// Solidity: function createMarket(address _owner, address quote, address feeCollector, bytes data) returns(address)
func (_CrossDexImplV2 *CrossDexImplV2Session) CreateMarket(_owner common.Address, quote common.Address, feeCollector common.Address, data []byte) (*types.Transaction, error) {
	return _CrossDexImplV2.Contract.CreateMarket(&_CrossDexImplV2.TransactOpts, _owner, quote, feeCollector, data)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x9967b84e.
//
// Solidity: function createMarket(address _owner, address quote, address feeCollector, bytes data) returns(address)
func (_CrossDexImplV2 *CrossDexImplV2TransactorSession) CreateMarket(_owner common.Address, quote common.Address, feeCollector common.Address, data []byte) (*types.Transaction, error) {
	return _CrossDexImplV2.Contract.CreateMarket(&_CrossDexImplV2.TransactOpts, _owner, quote, feeCollector, data)
}

// Initialize is a paid mutator transaction binding the contract method 0x0729da0b.
//
// Solidity: function initialize(address _owner, address _routerImpl, uint256 _findPrevPriceCount, uint256 _maxMatchCount, uint256 _cancelLimit, address _marketImpl, address _pairImpl) returns()
func (_CrossDexImplV2 *CrossDexImplV2Transactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _routerImpl common.Address, _findPrevPriceCount *big.Int, _maxMatchCount *big.Int, _cancelLimit *big.Int, _marketImpl common.Address, _pairImpl common.Address) (*types.Transaction, error) {
	return _CrossDexImplV2.contract.Transact(opts, "initialize", _owner, _routerImpl, _findPrevPriceCount, _maxMatchCount, _cancelLimit, _marketImpl, _pairImpl)
}

// Initialize is a paid mutator transaction binding the contract method 0x0729da0b.
//
// Solidity: function initialize(address _owner, address _routerImpl, uint256 _findPrevPriceCount, uint256 _maxMatchCount, uint256 _cancelLimit, address _marketImpl, address _pairImpl) returns()
func (_CrossDexImplV2 *CrossDexImplV2Session) Initialize(_owner common.Address, _routerImpl common.Address, _findPrevPriceCount *big.Int, _maxMatchCount *big.Int, _cancelLimit *big.Int, _marketImpl common.Address, _pairImpl common.Address) (*types.Transaction, error) {
	return _CrossDexImplV2.Contract.Initialize(&_CrossDexImplV2.TransactOpts, _owner, _routerImpl, _findPrevPriceCount, _maxMatchCount, _cancelLimit, _marketImpl, _pairImpl)
}

// Initialize is a paid mutator transaction binding the contract method 0x0729da0b.
//
// Solidity: function initialize(address _owner, address _routerImpl, uint256 _findPrevPriceCount, uint256 _maxMatchCount, uint256 _cancelLimit, address _marketImpl, address _pairImpl) returns()
func (_CrossDexImplV2 *CrossDexImplV2TransactorSession) Initialize(_owner common.Address, _routerImpl common.Address, _findPrevPriceCount *big.Int, _maxMatchCount *big.Int, _cancelLimit *big.Int, _marketImpl common.Address, _pairImpl common.Address) (*types.Transaction, error) {
	return _CrossDexImplV2.Contract.Initialize(&_CrossDexImplV2.TransactOpts, _owner, _routerImpl, _findPrevPriceCount, _maxMatchCount, _cancelLimit, _marketImpl, _pairImpl)
}

// PairCreated is a paid mutator transaction binding the contract method 0xe9f7206b.
//
// Solidity: function pairCreated(address pair) returns()
func (_CrossDexImplV2 *CrossDexImplV2Transactor) PairCreated(opts *bind.TransactOpts, pair common.Address) (*types.Transaction, error) {
	return _CrossDexImplV2.contract.Transact(opts, "pairCreated", pair)
}

// PairCreated is a paid mutator transaction binding the contract method 0xe9f7206b.
//
// Solidity: function pairCreated(address pair) returns()
func (_CrossDexImplV2 *CrossDexImplV2Session) PairCreated(pair common.Address) (*types.Transaction, error) {
	return _CrossDexImplV2.Contract.PairCreated(&_CrossDexImplV2.TransactOpts, pair)
}

// PairCreated is a paid mutator transaction binding the contract method 0xe9f7206b.
//
// Solidity: function pairCreated(address pair) returns()
func (_CrossDexImplV2 *CrossDexImplV2TransactorSession) PairCreated(pair common.Address) (*types.Transaction, error) {
	return _CrossDexImplV2.Contract.PairCreated(&_CrossDexImplV2.TransactOpts, pair)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CrossDexImplV2 *CrossDexImplV2Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossDexImplV2.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CrossDexImplV2 *CrossDexImplV2Session) RenounceOwnership() (*types.Transaction, error) {
	return _CrossDexImplV2.Contract.RenounceOwnership(&_CrossDexImplV2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CrossDexImplV2 *CrossDexImplV2TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CrossDexImplV2.Contract.RenounceOwnership(&_CrossDexImplV2.TransactOpts)
}

// SetMarketImpl is a paid mutator transaction binding the contract method 0xe492c270.
//
// Solidity: function setMarketImpl(address _marketImpl) returns()
func (_CrossDexImplV2 *CrossDexImplV2Transactor) SetMarketImpl(opts *bind.TransactOpts, _marketImpl common.Address) (*types.Transaction, error) {
	return _CrossDexImplV2.contract.Transact(opts, "setMarketImpl", _marketImpl)
}

// SetMarketImpl is a paid mutator transaction binding the contract method 0xe492c270.
//
// Solidity: function setMarketImpl(address _marketImpl) returns()
func (_CrossDexImplV2 *CrossDexImplV2Session) SetMarketImpl(_marketImpl common.Address) (*types.Transaction, error) {
	return _CrossDexImplV2.Contract.SetMarketImpl(&_CrossDexImplV2.TransactOpts, _marketImpl)
}

// SetMarketImpl is a paid mutator transaction binding the contract method 0xe492c270.
//
// Solidity: function setMarketImpl(address _marketImpl) returns()
func (_CrossDexImplV2 *CrossDexImplV2TransactorSession) SetMarketImpl(_marketImpl common.Address) (*types.Transaction, error) {
	return _CrossDexImplV2.Contract.SetMarketImpl(&_CrossDexImplV2.TransactOpts, _marketImpl)
}

// SetPairImpl is a paid mutator transaction binding the contract method 0x3cc047aa.
//
// Solidity: function setPairImpl(address _pairImpl) returns()
func (_CrossDexImplV2 *CrossDexImplV2Transactor) SetPairImpl(opts *bind.TransactOpts, _pairImpl common.Address) (*types.Transaction, error) {
	return _CrossDexImplV2.contract.Transact(opts, "setPairImpl", _pairImpl)
}

// SetPairImpl is a paid mutator transaction binding the contract method 0x3cc047aa.
//
// Solidity: function setPairImpl(address _pairImpl) returns()
func (_CrossDexImplV2 *CrossDexImplV2Session) SetPairImpl(_pairImpl common.Address) (*types.Transaction, error) {
	return _CrossDexImplV2.Contract.SetPairImpl(&_CrossDexImplV2.TransactOpts, _pairImpl)
}

// SetPairImpl is a paid mutator transaction binding the contract method 0x3cc047aa.
//
// Solidity: function setPairImpl(address _pairImpl) returns()
func (_CrossDexImplV2 *CrossDexImplV2TransactorSession) SetPairImpl(_pairImpl common.Address) (*types.Transaction, error) {
	return _CrossDexImplV2.Contract.SetPairImpl(&_CrossDexImplV2.TransactOpts, _pairImpl)
}

// SetTickSizeSetter is a paid mutator transaction binding the contract method 0x98715ee7.
//
// Solidity: function setTickSizeSetter(address setter) returns()
func (_CrossDexImplV2 *CrossDexImplV2Transactor) SetTickSizeSetter(opts *bind.TransactOpts, setter common.Address) (*types.Transaction, error) {
	return _CrossDexImplV2.contract.Transact(opts, "setTickSizeSetter", setter)
}

// SetTickSizeSetter is a paid mutator transaction binding the contract method 0x98715ee7.
//
// Solidity: function setTickSizeSetter(address setter) returns()
func (_CrossDexImplV2 *CrossDexImplV2Session) SetTickSizeSetter(setter common.Address) (*types.Transaction, error) {
	return _CrossDexImplV2.Contract.SetTickSizeSetter(&_CrossDexImplV2.TransactOpts, setter)
}

// SetTickSizeSetter is a paid mutator transaction binding the contract method 0x98715ee7.
//
// Solidity: function setTickSizeSetter(address setter) returns()
func (_CrossDexImplV2 *CrossDexImplV2TransactorSession) SetTickSizeSetter(setter common.Address) (*types.Transaction, error) {
	return _CrossDexImplV2.Contract.SetTickSizeSetter(&_CrossDexImplV2.TransactOpts, setter)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CrossDexImplV2 *CrossDexImplV2Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CrossDexImplV2.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CrossDexImplV2 *CrossDexImplV2Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CrossDexImplV2.Contract.TransferOwnership(&_CrossDexImplV2.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CrossDexImplV2 *CrossDexImplV2TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CrossDexImplV2.Contract.TransferOwnership(&_CrossDexImplV2.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CrossDexImplV2 *CrossDexImplV2Transactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CrossDexImplV2.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CrossDexImplV2 *CrossDexImplV2Session) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CrossDexImplV2.Contract.UpgradeToAndCall(&_CrossDexImplV2.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CrossDexImplV2 *CrossDexImplV2TransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CrossDexImplV2.Contract.UpgradeToAndCall(&_CrossDexImplV2.TransactOpts, newImplementation, data)
}

// CrossDexImplV2InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the CrossDexImplV2 contract.
type CrossDexImplV2InitializedIterator struct {
	Event *CrossDexImplV2Initialized // Event containing the contract specifics and raw log

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
func (it *CrossDexImplV2InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossDexImplV2Initialized)
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
		it.Event = new(CrossDexImplV2Initialized)
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
func (it *CrossDexImplV2InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossDexImplV2InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossDexImplV2Initialized represents a Initialized event raised by the CrossDexImplV2 contract.
type CrossDexImplV2Initialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CrossDexImplV2 *CrossDexImplV2Filterer) FilterInitialized(opts *bind.FilterOpts) (*CrossDexImplV2InitializedIterator, error) {

	logs, sub, err := _CrossDexImplV2.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CrossDexImplV2InitializedIterator{contract: _CrossDexImplV2.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CrossDexImplV2 *CrossDexImplV2Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CrossDexImplV2Initialized) (event.Subscription, error) {

	logs, sub, err := _CrossDexImplV2.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossDexImplV2Initialized)
				if err := _CrossDexImplV2.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_CrossDexImplV2 *CrossDexImplV2Filterer) ParseInitialized(log types.Log) (*CrossDexImplV2Initialized, error) {
	event := new(CrossDexImplV2Initialized)
	if err := _CrossDexImplV2.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossDexImplV2MarketCreatedIterator is returned from FilterMarketCreated and is used to iterate over the raw logs and unpacked data for MarketCreated events raised by the CrossDexImplV2 contract.
type CrossDexImplV2MarketCreatedIterator struct {
	Event *CrossDexImplV2MarketCreated // Event containing the contract specifics and raw log

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
func (it *CrossDexImplV2MarketCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossDexImplV2MarketCreated)
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
		it.Event = new(CrossDexImplV2MarketCreated)
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
func (it *CrossDexImplV2MarketCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossDexImplV2MarketCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossDexImplV2MarketCreated represents a MarketCreated event raised by the CrossDexImplV2 contract.
type CrossDexImplV2MarketCreated struct {
	Quote        common.Address
	Market       common.Address
	Owner        common.Address
	FeeCollector common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMarketCreated is a free log retrieval operation binding the contract event 0x5509f6235f0d08c563aedf1b655226cad4a613794902c476ad7a08accb43ebc5.
//
// Solidity: event MarketCreated(address indexed quote, address indexed market, address indexed owner, address fee_collector)
func (_CrossDexImplV2 *CrossDexImplV2Filterer) FilterMarketCreated(opts *bind.FilterOpts, quote []common.Address, market []common.Address, owner []common.Address) (*CrossDexImplV2MarketCreatedIterator, error) {

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

	logs, sub, err := _CrossDexImplV2.contract.FilterLogs(opts, "MarketCreated", quoteRule, marketRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &CrossDexImplV2MarketCreatedIterator{contract: _CrossDexImplV2.contract, event: "MarketCreated", logs: logs, sub: sub}, nil
}

// WatchMarketCreated is a free log subscription operation binding the contract event 0x5509f6235f0d08c563aedf1b655226cad4a613794902c476ad7a08accb43ebc5.
//
// Solidity: event MarketCreated(address indexed quote, address indexed market, address indexed owner, address fee_collector)
func (_CrossDexImplV2 *CrossDexImplV2Filterer) WatchMarketCreated(opts *bind.WatchOpts, sink chan<- *CrossDexImplV2MarketCreated, quote []common.Address, market []common.Address, owner []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CrossDexImplV2.contract.WatchLogs(opts, "MarketCreated", quoteRule, marketRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossDexImplV2MarketCreated)
				if err := _CrossDexImplV2.contract.UnpackLog(event, "MarketCreated", log); err != nil {
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
func (_CrossDexImplV2 *CrossDexImplV2Filterer) ParseMarketCreated(log types.Log) (*CrossDexImplV2MarketCreated, error) {
	event := new(CrossDexImplV2MarketCreated)
	if err := _CrossDexImplV2.contract.UnpackLog(event, "MarketCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossDexImplV2MarketImplSetIterator is returned from FilterMarketImplSet and is used to iterate over the raw logs and unpacked data for MarketImplSet events raised by the CrossDexImplV2 contract.
type CrossDexImplV2MarketImplSetIterator struct {
	Event *CrossDexImplV2MarketImplSet // Event containing the contract specifics and raw log

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
func (it *CrossDexImplV2MarketImplSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossDexImplV2MarketImplSet)
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
		it.Event = new(CrossDexImplV2MarketImplSet)
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
func (it *CrossDexImplV2MarketImplSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossDexImplV2MarketImplSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossDexImplV2MarketImplSet represents a MarketImplSet event raised by the CrossDexImplV2 contract.
type CrossDexImplV2MarketImplSet struct {
	Before  common.Address
	Current common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMarketImplSet is a free log retrieval operation binding the contract event 0x081eadab9f9e460b348376155bdd3b4aed9b2a33469b489416d7416b4cffe49c.
//
// Solidity: event MarketImplSet(address indexed before, address indexed current)
func (_CrossDexImplV2 *CrossDexImplV2Filterer) FilterMarketImplSet(opts *bind.FilterOpts, before []common.Address, current []common.Address) (*CrossDexImplV2MarketImplSetIterator, error) {

	var beforeRule []interface{}
	for _, beforeItem := range before {
		beforeRule = append(beforeRule, beforeItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _CrossDexImplV2.contract.FilterLogs(opts, "MarketImplSet", beforeRule, currentRule)
	if err != nil {
		return nil, err
	}
	return &CrossDexImplV2MarketImplSetIterator{contract: _CrossDexImplV2.contract, event: "MarketImplSet", logs: logs, sub: sub}, nil
}

// WatchMarketImplSet is a free log subscription operation binding the contract event 0x081eadab9f9e460b348376155bdd3b4aed9b2a33469b489416d7416b4cffe49c.
//
// Solidity: event MarketImplSet(address indexed before, address indexed current)
func (_CrossDexImplV2 *CrossDexImplV2Filterer) WatchMarketImplSet(opts *bind.WatchOpts, sink chan<- *CrossDexImplV2MarketImplSet, before []common.Address, current []common.Address) (event.Subscription, error) {

	var beforeRule []interface{}
	for _, beforeItem := range before {
		beforeRule = append(beforeRule, beforeItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _CrossDexImplV2.contract.WatchLogs(opts, "MarketImplSet", beforeRule, currentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossDexImplV2MarketImplSet)
				if err := _CrossDexImplV2.contract.UnpackLog(event, "MarketImplSet", log); err != nil {
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

// ParseMarketImplSet is a log parse operation binding the contract event 0x081eadab9f9e460b348376155bdd3b4aed9b2a33469b489416d7416b4cffe49c.
//
// Solidity: event MarketImplSet(address indexed before, address indexed current)
func (_CrossDexImplV2 *CrossDexImplV2Filterer) ParseMarketImplSet(log types.Log) (*CrossDexImplV2MarketImplSet, error) {
	event := new(CrossDexImplV2MarketImplSet)
	if err := _CrossDexImplV2.contract.UnpackLog(event, "MarketImplSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossDexImplV2OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CrossDexImplV2 contract.
type CrossDexImplV2OwnershipTransferredIterator struct {
	Event *CrossDexImplV2OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CrossDexImplV2OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossDexImplV2OwnershipTransferred)
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
		it.Event = new(CrossDexImplV2OwnershipTransferred)
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
func (it *CrossDexImplV2OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossDexImplV2OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossDexImplV2OwnershipTransferred represents a OwnershipTransferred event raised by the CrossDexImplV2 contract.
type CrossDexImplV2OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CrossDexImplV2 *CrossDexImplV2Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CrossDexImplV2OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CrossDexImplV2.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CrossDexImplV2OwnershipTransferredIterator{contract: _CrossDexImplV2.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CrossDexImplV2 *CrossDexImplV2Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CrossDexImplV2OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CrossDexImplV2.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossDexImplV2OwnershipTransferred)
				if err := _CrossDexImplV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_CrossDexImplV2 *CrossDexImplV2Filterer) ParseOwnershipTransferred(log types.Log) (*CrossDexImplV2OwnershipTransferred, error) {
	event := new(CrossDexImplV2OwnershipTransferred)
	if err := _CrossDexImplV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossDexImplV2PairImplSetIterator is returned from FilterPairImplSet and is used to iterate over the raw logs and unpacked data for PairImplSet events raised by the CrossDexImplV2 contract.
type CrossDexImplV2PairImplSetIterator struct {
	Event *CrossDexImplV2PairImplSet // Event containing the contract specifics and raw log

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
func (it *CrossDexImplV2PairImplSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossDexImplV2PairImplSet)
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
		it.Event = new(CrossDexImplV2PairImplSet)
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
func (it *CrossDexImplV2PairImplSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossDexImplV2PairImplSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossDexImplV2PairImplSet represents a PairImplSet event raised by the CrossDexImplV2 contract.
type CrossDexImplV2PairImplSet struct {
	Before  common.Address
	Current common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPairImplSet is a free log retrieval operation binding the contract event 0x685c330e9b04243f99ff39d8de8578c9af2bfaff120e17a1b5f1d61e004935e2.
//
// Solidity: event PairImplSet(address indexed before, address indexed current)
func (_CrossDexImplV2 *CrossDexImplV2Filterer) FilterPairImplSet(opts *bind.FilterOpts, before []common.Address, current []common.Address) (*CrossDexImplV2PairImplSetIterator, error) {

	var beforeRule []interface{}
	for _, beforeItem := range before {
		beforeRule = append(beforeRule, beforeItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _CrossDexImplV2.contract.FilterLogs(opts, "PairImplSet", beforeRule, currentRule)
	if err != nil {
		return nil, err
	}
	return &CrossDexImplV2PairImplSetIterator{contract: _CrossDexImplV2.contract, event: "PairImplSet", logs: logs, sub: sub}, nil
}

// WatchPairImplSet is a free log subscription operation binding the contract event 0x685c330e9b04243f99ff39d8de8578c9af2bfaff120e17a1b5f1d61e004935e2.
//
// Solidity: event PairImplSet(address indexed before, address indexed current)
func (_CrossDexImplV2 *CrossDexImplV2Filterer) WatchPairImplSet(opts *bind.WatchOpts, sink chan<- *CrossDexImplV2PairImplSet, before []common.Address, current []common.Address) (event.Subscription, error) {

	var beforeRule []interface{}
	for _, beforeItem := range before {
		beforeRule = append(beforeRule, beforeItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _CrossDexImplV2.contract.WatchLogs(opts, "PairImplSet", beforeRule, currentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossDexImplV2PairImplSet)
				if err := _CrossDexImplV2.contract.UnpackLog(event, "PairImplSet", log); err != nil {
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
func (_CrossDexImplV2 *CrossDexImplV2Filterer) ParsePairImplSet(log types.Log) (*CrossDexImplV2PairImplSet, error) {
	event := new(CrossDexImplV2PairImplSet)
	if err := _CrossDexImplV2.contract.UnpackLog(event, "PairImplSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossDexImplV2TickSizeSetterSetIterator is returned from FilterTickSizeSetterSet and is used to iterate over the raw logs and unpacked data for TickSizeSetterSet events raised by the CrossDexImplV2 contract.
type CrossDexImplV2TickSizeSetterSetIterator struct {
	Event *CrossDexImplV2TickSizeSetterSet // Event containing the contract specifics and raw log

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
func (it *CrossDexImplV2TickSizeSetterSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossDexImplV2TickSizeSetterSet)
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
		it.Event = new(CrossDexImplV2TickSizeSetterSet)
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
func (it *CrossDexImplV2TickSizeSetterSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossDexImplV2TickSizeSetterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossDexImplV2TickSizeSetterSet represents a TickSizeSetterSet event raised by the CrossDexImplV2 contract.
type CrossDexImplV2TickSizeSetterSet struct {
	Before  common.Address
	Current common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTickSizeSetterSet is a free log retrieval operation binding the contract event 0x4f1656ebcdda14dfa0f9951b46485570075b9730e1f1d22cb4ddacc35374fde8.
//
// Solidity: event TickSizeSetterSet(address indexed before, address indexed current)
func (_CrossDexImplV2 *CrossDexImplV2Filterer) FilterTickSizeSetterSet(opts *bind.FilterOpts, before []common.Address, current []common.Address) (*CrossDexImplV2TickSizeSetterSetIterator, error) {

	var beforeRule []interface{}
	for _, beforeItem := range before {
		beforeRule = append(beforeRule, beforeItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _CrossDexImplV2.contract.FilterLogs(opts, "TickSizeSetterSet", beforeRule, currentRule)
	if err != nil {
		return nil, err
	}
	return &CrossDexImplV2TickSizeSetterSetIterator{contract: _CrossDexImplV2.contract, event: "TickSizeSetterSet", logs: logs, sub: sub}, nil
}

// WatchTickSizeSetterSet is a free log subscription operation binding the contract event 0x4f1656ebcdda14dfa0f9951b46485570075b9730e1f1d22cb4ddacc35374fde8.
//
// Solidity: event TickSizeSetterSet(address indexed before, address indexed current)
func (_CrossDexImplV2 *CrossDexImplV2Filterer) WatchTickSizeSetterSet(opts *bind.WatchOpts, sink chan<- *CrossDexImplV2TickSizeSetterSet, before []common.Address, current []common.Address) (event.Subscription, error) {

	var beforeRule []interface{}
	for _, beforeItem := range before {
		beforeRule = append(beforeRule, beforeItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _CrossDexImplV2.contract.WatchLogs(opts, "TickSizeSetterSet", beforeRule, currentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossDexImplV2TickSizeSetterSet)
				if err := _CrossDexImplV2.contract.UnpackLog(event, "TickSizeSetterSet", log); err != nil {
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
func (_CrossDexImplV2 *CrossDexImplV2Filterer) ParseTickSizeSetterSet(log types.Log) (*CrossDexImplV2TickSizeSetterSet, error) {
	event := new(CrossDexImplV2TickSizeSetterSet)
	if err := _CrossDexImplV2.contract.UnpackLog(event, "TickSizeSetterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossDexImplV2UpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the CrossDexImplV2 contract.
type CrossDexImplV2UpgradedIterator struct {
	Event *CrossDexImplV2Upgraded // Event containing the contract specifics and raw log

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
func (it *CrossDexImplV2UpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossDexImplV2Upgraded)
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
		it.Event = new(CrossDexImplV2Upgraded)
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
func (it *CrossDexImplV2UpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossDexImplV2UpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossDexImplV2Upgraded represents a Upgraded event raised by the CrossDexImplV2 contract.
type CrossDexImplV2Upgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_CrossDexImplV2 *CrossDexImplV2Filterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*CrossDexImplV2UpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _CrossDexImplV2.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &CrossDexImplV2UpgradedIterator{contract: _CrossDexImplV2.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_CrossDexImplV2 *CrossDexImplV2Filterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *CrossDexImplV2Upgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _CrossDexImplV2.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossDexImplV2Upgraded)
				if err := _CrossDexImplV2.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_CrossDexImplV2 *CrossDexImplV2Filterer) ParseUpgraded(log types.Log) (*CrossDexImplV2Upgraded, error) {
	event := new(CrossDexImplV2Upgraded)
	if err := _CrossDexImplV2.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}