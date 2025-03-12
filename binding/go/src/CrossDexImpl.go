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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ROUTER\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allMarkets\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"quotes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"markets\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"checkTickSizeRoles\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_fee_collector\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_quote\",\"type\":\"address\"}],\"name\":\"createMarket\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_routerImpl\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxMatchCount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_marketImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pairImpl\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"}],\"name\":\"isMarket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"setter\",\"type\":\"address\"}],\"name\":\"isTickSizeSetter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketImpl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"name\":\"pairCreated\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pairImpl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"name\":\"pairToMarket\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"quote\",\"type\":\"address\"}],\"name\":\"quoteToMarket\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"setter\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"setTickSizeSetter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"quote\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fee_collector\",\"type\":\"address\"}],\"name\":\"MarketCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"setter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"TickSizeSetterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"CrossDexAlreadyCreatedMarketQuote\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"CrossDexAlreadyTickSizeSetter\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"CrossDexInvalidMarketAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"CrossDexUnauthorizedChangeTickSizes\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"EnumerableMapNonexistentKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"32fe7b26": "ROUTER()",
		"ad3cb1cc": "UPGRADE_INTERFACE_VERSION()",
		"375a7cba": "allMarkets()",
		"a1eea778": "checkTickSizeRoles(address)",
		"7416ce60": "createMarket(address,address,address)",
		"33e1a223": "initialize(address,address,uint256,address,address)",
		"6ec934da": "isMarket(address)",
		"de751667": "isTickSizeSetter(address)",
		"34eaeeb9": "marketImpl()",
		"8da5cb5b": "owner()",
		"e9f7206b": "pairCreated(address)",
		"35f7d456": "pairImpl()",
		"08270573": "pairToMarket(address)",
		"52d1902d": "proxiableUUID()",
		"beb380f1": "quoteToMarket(address)",
		"715018a6": "renounceOwnership()",
		"42b63f0f": "setTickSizeSetter(address,bool)",
		"f2fde38b": "transferOwnership(address)",
		"4f1ef286": "upgradeToAndCall(address,bytes)",
	},
	Bin: "0x60a060405230608052348015610013575f5ffd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b60805161185d6100f95f395f8181610b8c01528181610bb50152610cf9015261185d5ff3fe608060405260043610610110575f3560e01c80636ec934da1161009d578063ad3cb1cc11610062578063ad3cb1cc14610315578063beb380f114610352578063de75166714610371578063e9f7206b1461039f578063f2fde38b146103be575f5ffd5b80636ec934da14610258578063715018a6146102875780637416ce601461029b5780638da5cb5b146102ba578063a1eea778146102f6575f5ffd5b806335f7d456116100e357806335f7d456146101c3578063375a7cba146101e257806342b63f0f146102045780634f1ef2861461022357806352d1902d14610236575f5ffd5b8063082705731461011457806332fe7b261461016557806333e1a2231461018357806334eaeeb9146101a4575b5f5ffd5b34801561011f575f5ffd5b5061014861012e366004611163565b60066020525f90815260409020546001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b348015610170575f5ffd5b505f54610148906001600160a01b031681565b34801561018e575f5ffd5b506101a261019d36600461117e565b6103dd565b005b3480156101af575f5ffd5b50600154610148906001600160a01b031681565b3480156101ce575f5ffd5b50600254610148906001600160a01b031681565b3480156101ed575f5ffd5b506101f66105d3565b60405161015c929190611225565b34801561020f575f5ffd5b506101a261021e366004611249565b6106d5565b6101a2610231366004611298565b610788565b348015610241575f5ffd5b5061024a6107a7565b60405190815260200161015c565b348015610263575f5ffd5b50610277610272366004611163565b6107c2565b604051901515815260200161015c565b348015610292575f5ffd5b506101a2610896565b3480156102a6575f5ffd5b506101486102b536600461135e565b6108a9565b3480156102c5575f5ffd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b0316610148565b348015610301575f5ffd5b506101a2610310366004611163565b610a0e565b348015610320575f5ffd5b50610345604051806040016040528060058152602001640352e302e360dc1b81525081565b60405161015c91906113a6565b34801561035d575f5ffd5b5061014861036c366004611163565b610a54565b34801561037c575f5ffd5b5061027761038b366004611163565b60076020525f908152604090205460ff1681565b3480156103aa575f5ffd5b506101a26103b9366004611163565b610a66565b3480156103c9575f5ffd5b506101a26103d8366004611163565b610ab6565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff16159067ffffffffffffffff165f811580156104225750825b90505f8267ffffffffffffffff16600114801561043e5750303b155b90508115801561044c575080155b1561046a5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561049457845460ff60401b1916600160401b1785555b5f896040516104a290611142565b6001600160a01b0390911681526040602082018190525f90820152606001604051809103905ff0801580156104d9573d5f5f3e3d5ffd5b505f80546001600160a01b0319166001600160a01b03831690811790915560405163fe4b84df60e01b8152600481018c90529192509063fe4b84df906024015f604051808303815f87803b15801561052f575f5ffd5b505af1158015610541573d5f5f3e3d5ffd5b5050600180546001600160a01b03808d166001600160a01b03199283161790925560028054928c16929091169190911790555061058191508b9050610af0565b83156105c757845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050505050565b6060805f6105e16003610b01565b90508067ffffffffffffffff8111156105fc576105fc611284565b604051908082528060200260200182016040528015610625578160200160208202803683370190505b5092508067ffffffffffffffff81111561064157610641611284565b60405190808252806020026020018201604052801561066a578160200160208202803683370190505b5091505f5b818110156106cf57610682600382610b0b565b858381518110610694576106946113db565b602002602001018584815181106106ad576106ad6113db565b6001600160a01b0393841660209182029290920101529116905260010161066f565b50509091565b6106dd610b26565b6001600160a01b0382165f9081526007602052604090205481151560ff90911615150361073557604051631e1a433d60e01b81526001600160a01b038316600482015281151560248201526044015b60405180910390fd5b6001600160a01b0382165f81815260076020526040808220805460ff191685151590811790915590519092917fc7bf922bb3bcef719e56c41addde8f3b7ca8d4d37ef2f1c212b2d71a522b6fb091a35050565b610790610b81565b61079982610c25565b6107a38282610c2d565b5050565b5f6107b0610cee565b505f5160206118085f395f51905f5290565b60408051600481526024810182526020810180516001600160e01b0316639c57983960e01b17905290515f91829182916001600160a01b0386169161080791906113ef565b5f60405180830381855afa9150503d805f811461083f576040519150601f19603f3d011682016040523d82523d5f602084013e610844565b606091505b50915091508161085757505f9392505050565b5f8180602001905181019061086c9190611405565b90506001600160a01b038516610883600383610d37565b6001600160a01b03161495945050505050565b61089e610b26565b6108a75f610d4b565b565b5f6108b2610b26565b6001546040515f916001600160a01b0316906108cd90611142565b6001600160a01b0390911681526040602082018190525f90820152606001604051809103905ff080158015610904573d5f5f3e3d5ffd5b505f54600254604051630a2ca2bd60e11b81526001600160a01b0389811660048301529283166024820152878316604482015286831660648201529082166084820152919250821690631459457a9060a4015f604051808303815f87803b15801561096d575f5ffd5b505af115801561097f573d5f5f3e3d5ffd5b508392506109939150600390508583610dbb565b6109bb57604051631dae075360e31b81526001600160a01b038516600482015260240161072c565b6040516001600160a01b03868116825280881691838216918716907f5509f6235f0d08c563aedf1b655226cad4a613794902c476ad7a08accb43ebc59060200160405180910390a49150505b9392505050565b6001600160a01b0381165f9081526007602052604090205460ff16610a515760405163f42eaafb60e01b81526001600160a01b038216600482015260240161072c565b50565b5f610a60600383610d37565b92915050565b610a6f336107c2565b610a8e57604051631d2f974360e21b815233600482015260240161072c565b6001600160a01b03165f90815260066020526040902080546001600160a01b03191633179055565b610abe610b26565b6001600160a01b038116610ae757604051631e4fbdf760e01b81525f600482015260240161072c565b610a5181610d4b565b610af8610ddb565b610a5181610e24565b5f610a6082610e2c565b5f808080610b198686610e36565b9097909650945050505050565b33610b587f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b0316146108a75760405163118cdaa760e01b815233600482015260240161072c565b306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480610c0757507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316610bfb5f5160206118085f395f51905f52546001600160a01b031690565b6001600160a01b031614155b156108a75760405163703e46dd60e11b815260040160405180910390fd5b610a51610b26565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015610c87575060408051601f3d908101601f19168201909252610c8491810190611420565b60015b610caf57604051634c9c8ce360e01b81526001600160a01b038316600482015260240161072c565b5f5160206118085f395f51905f528114610cdf57604051632a87526960e21b81526004810182905260240161072c565b610ce98383610e5f565b505050565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146108a75760405163703e46dd60e11b815260040160405180910390fd5b5f610a07836001600160a01b038416610eb4565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b5f610dd3846001600160a01b03808616908516610efa565b949350505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff166108a757604051631afcd79f60e31b815260040160405180910390fd5b610abe610ddb565b5f610a6082610f16565b5f8080610e438585610f1f565b5f81815260029690960160205260409095205494959350505050565b610e6882610f2a565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115610eac57610ce98282610f8d565b6107a3610fff565b5f81815260028301602052604081205480158015610ed95750610ed7848461101e565b155b15610a075760405163015ab34360e11b81526004810184905260240161072c565b5f8281526002840160205260408120829055610dd38484611029565b5f610a60825490565b5f610a078383611034565b806001600160a01b03163b5f03610f5f57604051634c9c8ce360e01b81526001600160a01b038216600482015260240161072c565b5f5160206118085f395f51905f5280546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b031684604051610fa991906113ef565b5f60405180830381855af49150503d805f8114610fe1576040519150601f19603f3d011682016040523d82523d5f602084013e610fe6565b606091505b5091509150610ff685838361105a565b95945050505050565b34156108a75760405163b398979f60e01b815260040160405180910390fd5b5f610a0783836110b6565b5f610a0783836110cd565b5f825f018281548110611049576110496113db565b905f5260205f200154905092915050565b60608261106f5761106a82611119565b610a07565b815115801561108657506001600160a01b0384163b155b156110af57604051639996b31560e01b81526001600160a01b038516600482015260240161072c565b5080610a07565b5f8181526001830160205260408120541515610a07565b5f81815260018301602052604081205461111257508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610a60565b505f610a60565b8051156111295780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b6103d08061143883390190565b6001600160a01b0381168114610a51575f5ffd5b5f60208284031215611173575f5ffd5b8135610a078161114f565b5f5f5f5f5f60a08688031215611192575f5ffd5b853561119d8161114f565b945060208601356111ad8161114f565b93506040860135925060608601356111c48161114f565b915060808601356111d48161114f565b809150509295509295909350565b5f8151808452602084019350602083015f5b8281101561121b5781516001600160a01b03168652602095860195909101906001016111f4565b5093949350505050565b604081525f61123760408301856111e2565b8281036020840152610ff681856111e2565b5f5f6040838503121561125a575f5ffd5b82356112658161114f565b915060208301358015158114611279575f5ffd5b809150509250929050565b634e487b7160e01b5f52604160045260245ffd5b5f5f604083850312156112a9575f5ffd5b82356112b48161114f565b9150602083013567ffffffffffffffff8111156112cf575f5ffd5b8301601f810185136112df575f5ffd5b803567ffffffffffffffff8111156112f9576112f9611284565b604051601f8201601f19908116603f0116810167ffffffffffffffff8111828210171561132857611328611284565b60405281815282820160200187101561133f575f5ffd5b816020840160208301375f602083830101528093505050509250929050565b5f5f5f60608486031215611370575f5ffd5b833561137b8161114f565b9250602084013561138b8161114f565b9150604084013561139b8161114f565b809150509250925092565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b634e487b7160e01b5f52603260045260245ffd5b5f82518060208501845e5f920191825250919050565b5f60208284031215611415575f5ffd5b8151610a078161114f565b5f60208284031215611430575f5ffd5b505191905056fe60806040526040516103d03803806103d08339810160408190526100229161023c565b61002c8282610033565b5050610321565b61003c82610091565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a280511561008557610080828261010c565b505050565b61008d61017f565b5050565b806001600160a01b03163b5f036100cb57604051634c9c8ce360e01b81526001600160a01b03821660048201526024015b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b031684604051610128919061030b565b5f60405180830381855af49150503d805f8114610160576040519150601f19603f3d011682016040523d82523d5f602084013e610165565b606091505b5090925090506101768583836101a0565b95945050505050565b341561019e5760405163b398979f60e01b815260040160405180910390fd5b565b6060826101b5576101b0826101ff565b6101f8565b81511580156101cc57506001600160a01b0384163b155b156101f557604051639996b31560e01b81526001600160a01b03851660048201526024016100c2565b50805b9392505050565b80511561020f5780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b634e487b7160e01b5f52604160045260245ffd5b5f5f6040838503121561024d575f5ffd5b82516001600160a01b0381168114610263575f5ffd5b60208401519092506001600160401b0381111561027e575f5ffd5b8301601f8101851361028e575f5ffd5b80516001600160401b038111156102a7576102a7610228565b604051601f8201601f19908116603f011681016001600160401b03811182821017156102d5576102d5610228565b6040528181528282016020018710156102ec575f5ffd5b8160208401602083015e5f602083830101528093505050509250929050565b5f82518060208501845e5f920191825250919050565b60a38061032d5f395ff3fe6080604052600a600c565b005b60186014601a565b6050565b565b5f604b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b905090565b365f5f375f5f365f845af43d5f5f3e8080156069573d5ff35b3d5ffdfea26469706673582212205990887f0037c04aa22ef84499f094f4341c02e52649e71080dd1a78b7edba3e64736f6c634300081c0033360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbca2646970667358221220dd087d422343f4faed0823bd7fac3a75bde479aa236c318bfc3f922364ddef3a64736f6c634300081c0033",
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

// IsTickSizeSetter is a free data retrieval call binding the contract method 0xde751667.
//
// Solidity: function isTickSizeSetter(address setter) view returns(bool)
func (_CrossDexImpl *CrossDexImplCaller) IsTickSizeSetter(opts *bind.CallOpts, setter common.Address) (bool, error) {
	var out []interface{}
	err := _CrossDexImpl.contract.Call(opts, &out, "isTickSizeSetter", setter)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTickSizeSetter is a free data retrieval call binding the contract method 0xde751667.
//
// Solidity: function isTickSizeSetter(address setter) view returns(bool)
func (_CrossDexImpl *CrossDexImplSession) IsTickSizeSetter(setter common.Address) (bool, error) {
	return _CrossDexImpl.Contract.IsTickSizeSetter(&_CrossDexImpl.CallOpts, setter)
}

// IsTickSizeSetter is a free data retrieval call binding the contract method 0xde751667.
//
// Solidity: function isTickSizeSetter(address setter) view returns(bool)
func (_CrossDexImpl *CrossDexImplCallerSession) IsTickSizeSetter(setter common.Address) (bool, error) {
	return _CrossDexImpl.Contract.IsTickSizeSetter(&_CrossDexImpl.CallOpts, setter)
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

// CreateMarket is a paid mutator transaction binding the contract method 0x7416ce60.
//
// Solidity: function createMarket(address _owner, address _fee_collector, address _quote) returns(address)
func (_CrossDexImpl *CrossDexImplTransactor) CreateMarket(opts *bind.TransactOpts, _owner common.Address, _fee_collector common.Address, _quote common.Address) (*types.Transaction, error) {
	return _CrossDexImpl.contract.Transact(opts, "createMarket", _owner, _fee_collector, _quote)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x7416ce60.
//
// Solidity: function createMarket(address _owner, address _fee_collector, address _quote) returns(address)
func (_CrossDexImpl *CrossDexImplSession) CreateMarket(_owner common.Address, _fee_collector common.Address, _quote common.Address) (*types.Transaction, error) {
	return _CrossDexImpl.Contract.CreateMarket(&_CrossDexImpl.TransactOpts, _owner, _fee_collector, _quote)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x7416ce60.
//
// Solidity: function createMarket(address _owner, address _fee_collector, address _quote) returns(address)
func (_CrossDexImpl *CrossDexImplTransactorSession) CreateMarket(_owner common.Address, _fee_collector common.Address, _quote common.Address) (*types.Transaction, error) {
	return _CrossDexImpl.Contract.CreateMarket(&_CrossDexImpl.TransactOpts, _owner, _fee_collector, _quote)
}

// Initialize is a paid mutator transaction binding the contract method 0x33e1a223.
//
// Solidity: function initialize(address _owner, address _routerImpl, uint256 _maxMatchCount, address _marketImpl, address _pairImpl) returns()
func (_CrossDexImpl *CrossDexImplTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _routerImpl common.Address, _maxMatchCount *big.Int, _marketImpl common.Address, _pairImpl common.Address) (*types.Transaction, error) {
	return _CrossDexImpl.contract.Transact(opts, "initialize", _owner, _routerImpl, _maxMatchCount, _marketImpl, _pairImpl)
}

// Initialize is a paid mutator transaction binding the contract method 0x33e1a223.
//
// Solidity: function initialize(address _owner, address _routerImpl, uint256 _maxMatchCount, address _marketImpl, address _pairImpl) returns()
func (_CrossDexImpl *CrossDexImplSession) Initialize(_owner common.Address, _routerImpl common.Address, _maxMatchCount *big.Int, _marketImpl common.Address, _pairImpl common.Address) (*types.Transaction, error) {
	return _CrossDexImpl.Contract.Initialize(&_CrossDexImpl.TransactOpts, _owner, _routerImpl, _maxMatchCount, _marketImpl, _pairImpl)
}

// Initialize is a paid mutator transaction binding the contract method 0x33e1a223.
//
// Solidity: function initialize(address _owner, address _routerImpl, uint256 _maxMatchCount, address _marketImpl, address _pairImpl) returns()
func (_CrossDexImpl *CrossDexImplTransactorSession) Initialize(_owner common.Address, _routerImpl common.Address, _maxMatchCount *big.Int, _marketImpl common.Address, _pairImpl common.Address) (*types.Transaction, error) {
	return _CrossDexImpl.Contract.Initialize(&_CrossDexImpl.TransactOpts, _owner, _routerImpl, _maxMatchCount, _marketImpl, _pairImpl)
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

// SetTickSizeSetter is a paid mutator transaction binding the contract method 0x42b63f0f.
//
// Solidity: function setTickSizeSetter(address setter, bool allowed) returns()
func (_CrossDexImpl *CrossDexImplTransactor) SetTickSizeSetter(opts *bind.TransactOpts, setter common.Address, allowed bool) (*types.Transaction, error) {
	return _CrossDexImpl.contract.Transact(opts, "setTickSizeSetter", setter, allowed)
}

// SetTickSizeSetter is a paid mutator transaction binding the contract method 0x42b63f0f.
//
// Solidity: function setTickSizeSetter(address setter, bool allowed) returns()
func (_CrossDexImpl *CrossDexImplSession) SetTickSizeSetter(setter common.Address, allowed bool) (*types.Transaction, error) {
	return _CrossDexImpl.Contract.SetTickSizeSetter(&_CrossDexImpl.TransactOpts, setter, allowed)
}

// SetTickSizeSetter is a paid mutator transaction binding the contract method 0x42b63f0f.
//
// Solidity: function setTickSizeSetter(address setter, bool allowed) returns()
func (_CrossDexImpl *CrossDexImplTransactorSession) SetTickSizeSetter(setter common.Address, allowed bool) (*types.Transaction, error) {
	return _CrossDexImpl.Contract.SetTickSizeSetter(&_CrossDexImpl.TransactOpts, setter, allowed)
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
	Setter  common.Address
	Allowed bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTickSizeSetterSet is a free log retrieval operation binding the contract event 0xc7bf922bb3bcef719e56c41addde8f3b7ca8d4d37ef2f1c212b2d71a522b6fb0.
//
// Solidity: event TickSizeSetterSet(address indexed setter, bool indexed allowed)
func (_CrossDexImpl *CrossDexImplFilterer) FilterTickSizeSetterSet(opts *bind.FilterOpts, setter []common.Address, allowed []bool) (*CrossDexImplTickSizeSetterSetIterator, error) {

	var setterRule []interface{}
	for _, setterItem := range setter {
		setterRule = append(setterRule, setterItem)
	}
	var allowedRule []interface{}
	for _, allowedItem := range allowed {
		allowedRule = append(allowedRule, allowedItem)
	}

	logs, sub, err := _CrossDexImpl.contract.FilterLogs(opts, "TickSizeSetterSet", setterRule, allowedRule)
	if err != nil {
		return nil, err
	}
	return &CrossDexImplTickSizeSetterSetIterator{contract: _CrossDexImpl.contract, event: "TickSizeSetterSet", logs: logs, sub: sub}, nil
}

// WatchTickSizeSetterSet is a free log subscription operation binding the contract event 0xc7bf922bb3bcef719e56c41addde8f3b7ca8d4d37ef2f1c212b2d71a522b6fb0.
//
// Solidity: event TickSizeSetterSet(address indexed setter, bool indexed allowed)
func (_CrossDexImpl *CrossDexImplFilterer) WatchTickSizeSetterSet(opts *bind.WatchOpts, sink chan<- *CrossDexImplTickSizeSetterSet, setter []common.Address, allowed []bool) (event.Subscription, error) {

	var setterRule []interface{}
	for _, setterItem := range setter {
		setterRule = append(setterRule, setterItem)
	}
	var allowedRule []interface{}
	for _, allowedItem := range allowed {
		allowedRule = append(allowedRule, allowedItem)
	}

	logs, sub, err := _CrossDexImpl.contract.WatchLogs(opts, "TickSizeSetterSet", setterRule, allowedRule)
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

// ParseTickSizeSetterSet is a log parse operation binding the contract event 0xc7bf922bb3bcef719e56c41addde8f3b7ca8d4d37ef2f1c212b2d71a522b6fb0.
//
// Solidity: event TickSizeSetterSet(address indexed setter, bool indexed allowed)
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
