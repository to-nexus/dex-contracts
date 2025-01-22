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
	ABI: "[{\"inputs\":[],\"name\":\"ROUTER\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allMarkets\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allQuotes\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_fee_collector\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_quote\",\"type\":\"address\"}],\"name\":\"createMarket\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_routerImpl\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxMatchCount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_marketImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pairImpl\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketImpl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"name\":\"pairCreated\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pairImpl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"name\":\"pairToMarket\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"quote\",\"type\":\"address\"}],\"name\":\"quoteToMarket\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"quote\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fee_collector\",\"type\":\"address\"}],\"name\":\"CreatedMarket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"CrossDexAlreadyCreatedMarketQuote\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"CrossDexInvalidMarketAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"32fe7b26": "ROUTER()",
		"ad3cb1cc": "UPGRADE_INTERFACE_VERSION()",
		"375a7cba": "allMarkets()",
		"57df6828": "allQuotes()",
		"7416ce60": "createMarket(address,address,address)",
		"33e1a223": "initialize(address,address,uint256,address,address)",
		"34eaeeb9": "marketImpl()",
		"8da5cb5b": "owner()",
		"e9f7206b": "pairCreated(address)",
		"35f7d456": "pairImpl()",
		"08270573": "pairToMarket(address)",
		"52d1902d": "proxiableUUID()",
		"beb380f1": "quoteToMarket(address)",
		"715018a6": "renounceOwnership()",
		"f2fde38b": "transferOwnership(address)",
		"4f1ef286": "upgradeToAndCall(address,bytes)",
	},
	Bin: "0x60a0604052306080523480156012575f5ffd5b506080516113f56100395f395f81816107f701528181610820015261096401526113f55ff3fe6080604052600436106100ef575f3560e01c806357df682811610087578063ad3cb1cc11610057578063ad3cb1cc1461029a578063beb380f1146102d7578063e9f7206b1461030b578063f2fde38b1461032a575f5ffd5b806357df682814610217578063715018a61461022b5780637416ce601461023f5780638da5cb5b1461025e575f5ffd5b806335f7d456116100c257806335f7d456146101a2578063375a7cba146101c15780634f1ef286146101e257806352d1902d146101f5575f5ffd5b806308270573146100f357806332fe7b261461014457806333e1a2231461016257806334eaeeb914610183575b5f5ffd5b3480156100fe575f5ffd5b5061012761010d366004610d97565b60086020525f90815260409020546001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b34801561014f575f5ffd5b505f54610127906001600160a01b031681565b34801561016d575f5ffd5b5061018161017c366004610db0565b610349565b005b34801561018e575f5ffd5b50600154610127906001600160a01b031681565b3480156101ad575f5ffd5b50600254610127906001600160a01b031681565b3480156101cc575f5ffd5b506101d561053f565b60405161013b9190610e0a565b6101816101f0366004610e69565b610550565b348015610200575f5ffd5b5061020961056f565b60405190815260200161013b565b348015610222575f5ffd5b506101d561058a565b348015610236575f5ffd5b50610181610596565b34801561024a575f5ffd5b50610127610259366004610f2d565b6105a9565b348015610269575f5ffd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b0316610127565b3480156102a5575f5ffd5b506102ca604051806040016040528060058152602001640352e302e360dc1b81525081565b60405161013b9190610f6d565b3480156102e2575f5ffd5b506101276102f1366004610d97565b60076020525f90815260409020546001600160a01b031681565b348015610316575f5ffd5b50610181610325366004610d97565b610740565b348015610335575f5ffd5b50610181610344366004610d97565b610792565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff16159067ffffffffffffffff165f8115801561038e5750825b90505f8267ffffffffffffffff1660011480156103aa5750303b155b9050811580156103b8575080155b156103d65760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561040057845460ff60401b1916600160401b1785555b5f8960405161040e90610d6f565b6001600160a01b0390911681526040602082018190525f90820152606001604051809103905ff080158015610445573d5f5f3e3d5ffd5b505f80546001600160a01b0319166001600160a01b03831690811790915560405163fe4b84df60e01b8152600481018c90529192509063fe4b84df906024015f604051808303815f87803b15801561049b575f5ffd5b505af11580156104ad573d5f5f3e3d5ffd5b5050600180546001600160a01b03808d166001600160a01b03199283161790925560028054928c1692909116919091179055506104ed91508b90506107cf565b831561053357845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050505050565b606061054b60056107e0565b905090565b6105586107ec565b61056182610890565b61056b8282610898565b5050565b5f610578610959565b505f5160206113a05f395f51905f5290565b606061054b60036107e0565b61059e6109a2565b6105a75f6109fd565b565b5f6105b26109a2565b6105bd600383610a6d565b6105ea57604051631dae075360e31b81526001600160a01b03831660048201526024015b60405180910390fd5b6001546040515f916001600160a01b03169061060590610d6f565b6001600160a01b0390911681526040602082018190525f90820152606001604051809103905ff08015801561063c573d5f5f3e3d5ffd5b505f54600254604051630a2ca2bd60e11b81526001600160a01b0389811660048301529283166024820152878316604482015286831660648201529082166084820152919250821690631459457a9060a4015f604051808303815f87803b1580156106a5575f5ffd5b505af11580156106b7573d5f5f3e3d5ffd5b508392506106ca91506005905082610a6d565b506001600160a01b038481165f8181526007602090815260409182902080546001600160a01b03191686861690811790915591518985168152938a16939192917f04f42e2f802b5fe8e682dc6e42106de7853d18b5e951696e76b9a1e33e163ceb910160405180910390a49150505b9392505050565b61074b600533610a8a565b61076a57604051631d2f974360e21b81523360048201526024016105e1565b6001600160a01b03165f90815260086020526040902080546001600160a01b03191633179055565b61079a6109a2565b6001600160a01b0381166107c357604051631e4fbdf760e01b81525f60048201526024016105e1565b6107cc816109fd565b50565b6107d7610aab565b6107cc81610af4565b60605f61073983610afc565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148061087257507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166108665f5160206113a05f395f51905f52546001600160a01b031690565b6001600160a01b031614155b156105a75760405163703e46dd60e11b815260040160405180910390fd5b6107cc6109a2565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156108f2575060408051601f3d908101601f191682019092526108ef91810190610fa2565b60015b61091a57604051634c9c8ce360e01b81526001600160a01b03831660048201526024016105e1565b5f5160206113a05f395f51905f52811461094a57604051632a87526960e21b8152600481018290526024016105e1565b6109548383610b55565b505050565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146105a75760405163703e46dd60e11b815260040160405180910390fd5b336109d47f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b0316146105a75760405163118cdaa760e01b81523360048201526024016105e1565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b5f610a81836001600160a01b038416610baa565b90505b92915050565b6001600160a01b0381165f9081526001830160205260408120541515610a81565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff166105a757604051631afcd79f60e31b815260040160405180910390fd5b61079a610aab565b6060815f01805480602002602001604051908101604052809291908181526020018280548015610b4957602002820191905f5260205f20905b815481526020019060010190808311610b35575b50505050509050919050565b610b5e82610bf6565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115610ba2576109548282610c59565b61056b610ccb565b5f818152600183016020526040812054610bef57508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610a84565b505f610a84565b806001600160a01b03163b5f03610c2b57604051634c9c8ce360e01b81526001600160a01b03821660048201526024016105e1565b5f5160206113a05f395f51905f5280546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b031684604051610c759190610fb9565b5f60405180830381855af49150503d805f8114610cad576040519150601f19603f3d011682016040523d82523d5f602084013e610cb2565b606091505b5091509150610cc2858383610cea565b95945050505050565b34156105a75760405163b398979f60e01b815260040160405180910390fd5b606082610cff57610cfa82610d46565b610739565b8151158015610d1657506001600160a01b0384163b155b15610d3f57604051639996b31560e01b81526001600160a01b03851660048201526024016105e1565b5080610739565b805115610d565780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b6103d080610fd083390190565b80356001600160a01b0381168114610d92575f5ffd5b919050565b5f60208284031215610da7575f5ffd5b610a8182610d7c565b5f5f5f5f5f60a08688031215610dc4575f5ffd5b610dcd86610d7c565b9450610ddb60208701610d7c565b935060408601359250610df060608701610d7c565b9150610dfe60808701610d7c565b90509295509295909350565b602080825282518282018190525f918401906040840190835b81811015610e4a5783516001600160a01b0316835260209384019390920191600101610e23565b509095945050505050565b634e487b7160e01b5f52604160045260245ffd5b5f5f60408385031215610e7a575f5ffd5b610e8383610d7c565b9150602083013567ffffffffffffffff811115610e9e575f5ffd5b8301601f81018513610eae575f5ffd5b803567ffffffffffffffff811115610ec857610ec8610e55565b604051601f8201601f19908116603f0116810167ffffffffffffffff81118282101715610ef757610ef7610e55565b604052818152828201602001871015610f0e575f5ffd5b816020840160208301375f602083830101528093505050509250929050565b5f5f5f60608486031215610f3f575f5ffd5b610f4884610d7c565b9250610f5660208501610d7c565b9150610f6460408501610d7c565b90509250925092565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b5f60208284031215610fb2575f5ffd5b5051919050565b5f82518060208501845e5f92019182525091905056fe60806040526040516103d03803806103d08339810160408190526100229161023c565b61002c8282610033565b5050610321565b61003c82610091565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a280511561008557610080828261010c565b505050565b61008d61017f565b5050565b806001600160a01b03163b5f036100cb57604051634c9c8ce360e01b81526001600160a01b03821660048201526024015b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b031684604051610128919061030b565b5f60405180830381855af49150503d805f8114610160576040519150601f19603f3d011682016040523d82523d5f602084013e610165565b606091505b5090925090506101768583836101a0565b95945050505050565b341561019e5760405163b398979f60e01b815260040160405180910390fd5b565b6060826101b5576101b0826101ff565b6101f8565b81511580156101cc57506001600160a01b0384163b155b156101f557604051639996b31560e01b81526001600160a01b03851660048201526024016100c2565b50805b9392505050565b80511561020f5780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b634e487b7160e01b5f52604160045260245ffd5b5f5f6040838503121561024d575f5ffd5b82516001600160a01b0381168114610263575f5ffd5b60208401519092506001600160401b0381111561027e575f5ffd5b8301601f8101851361028e575f5ffd5b80516001600160401b038111156102a7576102a7610228565b604051601f8201601f19908116603f011681016001600160401b03811182821017156102d5576102d5610228565b6040528181528282016020018710156102ec575f5ffd5b8160208401602083015e5f602083830101528093505050509250929050565b5f82518060208501845e5f920191825250919050565b60a38061032d5f395ff3fe6080604052600a600c565b005b60186014601a565b6050565b565b5f604b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b905090565b365f5f375f5f365f845af43d5f5f3e8080156069573d5ff35b3d5ffdfea26469706673582212205990887f0037c04aa22ef84499f094f4341c02e52649e71080dd1a78b7edba3e64736f6c634300081c0033360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbca2646970667358221220060995555ebc67f08623fefd11bc88d1332e39e030230ee0e63a822e7fb36aab64736f6c634300081c0033",
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
// Solidity: function allMarkets() view returns(address[])
func (_CrossDexImpl *CrossDexImplCaller) AllMarkets(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _CrossDexImpl.contract.Call(opts, &out, "allMarkets")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllMarkets is a free data retrieval call binding the contract method 0x375a7cba.
//
// Solidity: function allMarkets() view returns(address[])
func (_CrossDexImpl *CrossDexImplSession) AllMarkets() ([]common.Address, error) {
	return _CrossDexImpl.Contract.AllMarkets(&_CrossDexImpl.CallOpts)
}

// AllMarkets is a free data retrieval call binding the contract method 0x375a7cba.
//
// Solidity: function allMarkets() view returns(address[])
func (_CrossDexImpl *CrossDexImplCallerSession) AllMarkets() ([]common.Address, error) {
	return _CrossDexImpl.Contract.AllMarkets(&_CrossDexImpl.CallOpts)
}

// AllQuotes is a free data retrieval call binding the contract method 0x57df6828.
//
// Solidity: function allQuotes() view returns(address[])
func (_CrossDexImpl *CrossDexImplCaller) AllQuotes(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _CrossDexImpl.contract.Call(opts, &out, "allQuotes")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllQuotes is a free data retrieval call binding the contract method 0x57df6828.
//
// Solidity: function allQuotes() view returns(address[])
func (_CrossDexImpl *CrossDexImplSession) AllQuotes() ([]common.Address, error) {
	return _CrossDexImpl.Contract.AllQuotes(&_CrossDexImpl.CallOpts)
}

// AllQuotes is a free data retrieval call binding the contract method 0x57df6828.
//
// Solidity: function allQuotes() view returns(address[])
func (_CrossDexImpl *CrossDexImplCallerSession) AllQuotes() ([]common.Address, error) {
	return _CrossDexImpl.Contract.AllQuotes(&_CrossDexImpl.CallOpts)
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

// CrossDexImplCreatedMarketIterator is returned from FilterCreatedMarket and is used to iterate over the raw logs and unpacked data for CreatedMarket events raised by the CrossDexImpl contract.
type CrossDexImplCreatedMarketIterator struct {
	Event *CrossDexImplCreatedMarket // Event containing the contract specifics and raw log

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
func (it *CrossDexImplCreatedMarketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossDexImplCreatedMarket)
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
		it.Event = new(CrossDexImplCreatedMarket)
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
func (it *CrossDexImplCreatedMarketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossDexImplCreatedMarketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossDexImplCreatedMarket represents a CreatedMarket event raised by the CrossDexImpl contract.
type CrossDexImplCreatedMarket struct {
	Quote        common.Address
	Market       common.Address
	Owner        common.Address
	FeeCollector common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCreatedMarket is a free log retrieval operation binding the contract event 0x04f42e2f802b5fe8e682dc6e42106de7853d18b5e951696e76b9a1e33e163ceb.
//
// Solidity: event CreatedMarket(address indexed quote, address indexed market, address indexed owner, address fee_collector)
func (_CrossDexImpl *CrossDexImplFilterer) FilterCreatedMarket(opts *bind.FilterOpts, quote []common.Address, market []common.Address, owner []common.Address) (*CrossDexImplCreatedMarketIterator, error) {

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

	logs, sub, err := _CrossDexImpl.contract.FilterLogs(opts, "CreatedMarket", quoteRule, marketRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &CrossDexImplCreatedMarketIterator{contract: _CrossDexImpl.contract, event: "CreatedMarket", logs: logs, sub: sub}, nil
}

// WatchCreatedMarket is a free log subscription operation binding the contract event 0x04f42e2f802b5fe8e682dc6e42106de7853d18b5e951696e76b9a1e33e163ceb.
//
// Solidity: event CreatedMarket(address indexed quote, address indexed market, address indexed owner, address fee_collector)
func (_CrossDexImpl *CrossDexImplFilterer) WatchCreatedMarket(opts *bind.WatchOpts, sink chan<- *CrossDexImplCreatedMarket, quote []common.Address, market []common.Address, owner []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CrossDexImpl.contract.WatchLogs(opts, "CreatedMarket", quoteRule, marketRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossDexImplCreatedMarket)
				if err := _CrossDexImpl.contract.UnpackLog(event, "CreatedMarket", log); err != nil {
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

// ParseCreatedMarket is a log parse operation binding the contract event 0x04f42e2f802b5fe8e682dc6e42106de7853d18b5e951696e76b9a1e33e163ceb.
//
// Solidity: event CreatedMarket(address indexed quote, address indexed market, address indexed owner, address fee_collector)
func (_CrossDexImpl *CrossDexImplFilterer) ParseCreatedMarket(log types.Log) (*CrossDexImplCreatedMarket, error) {
	event := new(CrossDexImplCreatedMarket)
	if err := _CrossDexImpl.contract.UnpackLog(event, "CreatedMarket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
