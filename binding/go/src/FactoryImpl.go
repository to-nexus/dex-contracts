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

// FactoryImplMetaData contains all meta data concerning the FactoryImpl contract.
var FactoryImplMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"QUOTE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allPairs\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"bases\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"pairs\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quoteTickSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseTickSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFeePermil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFeePermil\",\"type\":\"uint256\"}],\"name\":\"createPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeCollector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeCollector\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_quote\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pairImpl\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"}],\"name\":\"pairByBase\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pairImpl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"PairCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"FactoryAlreadyCreatedBaseAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FactoryDeployPair\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"FactoryInvalidBaseAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"FactoryInvalidInitializeData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"9c579839": "QUOTE()",
		"ad3cb1cc": "UPGRADE_INTERFACE_VERSION()",
		"c97682f8": "allPairs()",
		"9337bd68": "createPair(address,uint256,uint256,uint256,uint256)",
		"c415b95c": "feeCollector()",
		"f8c8765e": "initialize(address,address,address,address)",
		"8da5cb5b": "owner()",
		"43adeda9": "pairByBase(address)",
		"35f7d456": "pairImpl()",
		"52d1902d": "proxiableUUID()",
		"715018a6": "renounceOwnership()",
		"f887ea40": "router()",
		"f2fde38b": "transferOwnership(address)",
		"4f1ef286": "upgradeToAndCall(address,bytes)",
	},
	Bin: "0x60a0604052306080523480156012575f5ffd5b506080516115646100395f395f81816109110152818161093a0152610a7e01526115645ff3fe6080604052600436106100d9575f3560e01c80639c5798391161007c578063c97682f811610057578063c97682f814610248578063f2fde38b1461026a578063f887ea4014610289578063f8c8765e146102a8575f5ffd5b80639c579839146101ce578063ad3cb1cc146101ec578063c415b95c14610229575f5ffd5b806352d1902d116100b757806352d1902d14610165578063715018a6146101875780638da5cb5b1461019b5780639337bd68146101af575f5ffd5b806335f7d456146100dd57806343adeda9146101195780634f1ef28614610150575b5f5ffd5b3480156100e8575f5ffd5b506003546100fc906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b348015610124575f5ffd5b506100fc610133366004610e8b565b6001600160a01b039081165f908152600660205260409020541690565b61016361015e366004610eb8565b6102c7565b005b348015610170575f5ffd5b506101796102e6565b604051908152602001610110565b348015610192575f5ffd5b50610163610301565b3480156101a6575f5ffd5b506100fc610314565b3480156101ba575f5ffd5b506100fc6101c9366004610f7c565b610342565b3480156101d9575f5ffd5b505f546100fc906001600160a01b031681565b3480156101f7575f5ffd5b5061021c604051806040016040528060058152602001640352e302e360dc1b81525081565b6040516101109190610fe8565b348015610234575f5ffd5b506002546100fc906001600160a01b031681565b348015610253575f5ffd5b5061025c6105b8565b60405161011092919061103d565b348015610275575f5ffd5b50610163610284366004610e8b565b6106a0565b348015610294575f5ffd5b506001546100fc906001600160a01b031681565b3480156102b3575f5ffd5b506101636102c2366004611061565b6106dd565b6102cf610906565b6102d8826109aa565b6102e282826109b2565b5050565b5f6102ef610a73565b505f51602061150f5f395f51905f5290565b610309610abc565b6103125f610aee565b565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b5f61034b610abc565b6001600160a01b0386166103825760405163206f7d3560e01b81526001600160a01b03871660048201526024015b60405180910390fd5b61038d600487610b5e565b6103b5576040516339e6195360e01b81526001600160a01b0387166004820152602401610379565b5f866001600160a01b031663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa1580156103f2573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061041691906110b2565b60ff169050805f036104465760405163206f7d3560e01b81526001600160a01b0388166004820152602401610379565b6003546001600160a01b0316637116b01760e01b610462610314565b6001545f546002546040516001600160a01b039485166024820152928416604484015290831660648301528b8316608483015260a482018b905260c482018a90529190911660e4820152610104810187905261012481018690526101440160408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199094169390931790925290516104fd90610e63565b6105089291906110d2565b604051809103905ff080158015610521573d5f5f3e3d5ffd5b5091506001600160a01b03821661054b57604051639cecacaf60e01b815260040160405180910390fd5b6001600160a01b038781165f8181526006602090815260409182902080546001600160a01b031916948716948517905590514281529192917fc1db9ba7c4b7ce660fe8d17bbcf07167549381df2abd694a970bd1402d86d313910160405180910390a35095945050505050565b6060806105c56004610b7b565b80519092508067ffffffffffffffff8111156105e3576105e3610ea4565b60405190808252806020026020018201604052801561060c578160200160208202803683370190505b5091505f5b8181101561069a5760065f85838151811061062e5761062e6110fd565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020015f205f9054906101000a90046001600160a01b031683828151811061067a5761067a6110fd565b6001600160a01b0390921660209283029190910190910152600101610611565b50509091565b6106a8610abc565b6001600160a01b0381166106d157604051631e4fbdf760e01b81525f6004820152602401610379565b6106da81610aee565b50565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff16159067ffffffffffffffff165f811580156107225750825b90505f8267ffffffffffffffff16600114801561073e5750303b155b90508115801561074c575080155b1561076a5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561079457845460ff60401b1916600160401b1785555b6001600160a01b0389166107c65760405163711d871160e11b8152653937baba32b960d11b6004820152602401610379565b6001600160a01b0388166107fe5760405163711d871160e11b81526b3332b2a1b7b63632b1ba37b960a11b6004820152602401610379565b6001600160a01b03871661082f5760405163711d871160e11b81526471756f746560d81b6004820152602401610379565b6001600160a01b0386166108635760405163711d871160e11b8152671c185a5c925b5c1b60c21b6004820152602401610379565b600380546001600160a01b03199081166001600160a01b03898116919091179092556001805482168c84161790556002805482168b84161790555f80549091169189169190911790556108b533610b8e565b83156108fb57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050505050565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148061098c57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166109805f51602061150f5f395f51905f52546001600160a01b031690565b6001600160a01b031614155b156103125760405163703e46dd60e11b815260040160405180910390fd5b6106da610abc565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015610a0c575060408051601f3d908101601f19168201909252610a0991810190611111565b60015b610a3457604051634c9c8ce360e01b81526001600160a01b0383166004820152602401610379565b5f51602061150f5f395f51905f528114610a6457604051632a87526960e21b815260048101829052602401610379565b610a6e8383610b9f565b505050565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146103125760405163703e46dd60e11b815260040160405180910390fd5b33610ac5610314565b6001600160a01b0316146103125760405163118cdaa760e01b8152336004820152602401610379565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b5f610b72836001600160a01b038416610bf4565b90505b92915050565b60605f610b8783610c40565b9392505050565b610b96610c99565b6106da81610ce2565b610ba882610cea565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115610bec57610a6e8282610d4d565b6102e2610dbf565b5f818152600183016020526040812054610c3957508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610b75565b505f610b75565b6060815f01805480602002602001604051908101604052809291908181526020018280548015610c8d57602002820191905f5260205f20905b815481526020019060010190808311610c79575b50505050509050919050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff1661031257604051631afcd79f60e31b815260040160405180910390fd5b6106a8610c99565b806001600160a01b03163b5f03610d1f57604051634c9c8ce360e01b81526001600160a01b0382166004820152602401610379565b5f51602061150f5f395f51905f5280546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b031684604051610d699190611128565b5f60405180830381855af49150503d805f8114610da1576040519150601f19603f3d011682016040523d82523d5f602084013e610da6565b606091505b5091509150610db6858383610dde565b95945050505050565b34156103125760405163b398979f60e01b815260040160405180910390fd5b606082610df357610dee82610e3a565b610b87565b8151158015610e0a57506001600160a01b0384163b155b15610e3357604051639996b31560e01b81526001600160a01b0385166004820152602401610379565b5092915050565b805115610e4a5780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b6103d08061113f83390190565b80356001600160a01b0381168114610e86575f5ffd5b919050565b5f60208284031215610e9b575f5ffd5b610b7282610e70565b634e487b7160e01b5f52604160045260245ffd5b5f5f60408385031215610ec9575f5ffd5b610ed283610e70565b9150602083013567ffffffffffffffff811115610eed575f5ffd5b8301601f81018513610efd575f5ffd5b803567ffffffffffffffff811115610f1757610f17610ea4565b604051601f8201601f19908116603f0116810167ffffffffffffffff81118282101715610f4657610f46610ea4565b604052818152828201602001871015610f5d575f5ffd5b816020840160208301375f602083830101528093505050509250929050565b5f5f5f5f5f60a08688031215610f90575f5ffd5b610f9986610e70565b97602087013597506040870135966060810135965060800135945092505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f610b726020830184610fba565b5f8151808452602084019350602083015f5b828110156110335781516001600160a01b031686526020958601959091019060010161100c565b5093949350505050565b604081525f61104f6040830185610ffa565b8281036020840152610db68185610ffa565b5f5f5f5f60808587031215611074575f5ffd5b61107d85610e70565b935061108b60208601610e70565b925061109960408601610e70565b91506110a760608601610e70565b905092959194509250565b5f602082840312156110c2575f5ffd5b815160ff81168114610b87575f5ffd5b6001600160a01b03831681526040602082018190525f906110f590830184610fba565b949350505050565b634e487b7160e01b5f52603260045260245ffd5b5f60208284031215611121575f5ffd5b5051919050565b5f82518060208501845e5f92019182525091905056fe60806040526040516103d03803806103d08339810160408190526100229161023c565b61002c8282610033565b5050610321565b61003c82610091565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a280511561008557610080828261010c565b505050565b61008d61017f565b5050565b806001600160a01b03163b5f036100cb57604051634c9c8ce360e01b81526001600160a01b03821660048201526024015b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b031684604051610128919061030b565b5f60405180830381855af49150503d805f8114610160576040519150601f19603f3d011682016040523d82523d5f602084013e610165565b606091505b5090925090506101768583836101a0565b95945050505050565b341561019e5760405163b398979f60e01b815260040160405180910390fd5b565b6060826101b5576101b0826101ff565b6101f8565b81511580156101cc57506001600160a01b0384163b155b156101f557604051639996b31560e01b81526001600160a01b03851660048201526024016100c2565b50805b9392505050565b80511561020f5780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b634e487b7160e01b5f52604160045260245ffd5b5f5f6040838503121561024d575f5ffd5b82516001600160a01b0381168114610263575f5ffd5b60208401519092506001600160401b0381111561027e575f5ffd5b8301601f8101851361028e575f5ffd5b80516001600160401b038111156102a7576102a7610228565b604051601f8201601f19908116603f011681016001600160401b03811182821017156102d5576102d5610228565b6040528181528282016020018710156102ec575f5ffd5b8160208401602083015e5f602083830101528093505050509250929050565b5f82518060208501845e5f920191825250919050565b60a38061032d5f395ff3fe6080604052600a600c565b005b60186014601a565b6050565b565b5f604b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b905090565b365f5f375f5f365f845af43d5f5f3e8080156069573d5ff35b3d5ffdfea26469706673582212205990887f0037c04aa22ef84499f094f4341c02e52649e71080dd1a78b7edba3e64736f6c634300081c0033360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbca26469706673582212208f200073b841592b70ffea04b8b65871f5b1e780dd981fc89ac61f22c0afc20264736f6c634300081c0033",
}

// FactoryImplABI is the input ABI used to generate the binding from.
// Deprecated: Use FactoryImplMetaData.ABI instead.
var FactoryImplABI = FactoryImplMetaData.ABI

// Deprecated: Use FactoryImplMetaData.Sigs instead.
// FactoryImplFuncSigs maps the 4-byte function signature to its string representation.
var FactoryImplFuncSigs = FactoryImplMetaData.Sigs

// FactoryImplBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FactoryImplMetaData.Bin instead.
var FactoryImplBin = FactoryImplMetaData.Bin

// DeployFactoryImpl deploys a new Ethereum contract, binding an instance of FactoryImpl to it.
func DeployFactoryImpl(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FactoryImpl, error) {
	parsed, err := FactoryImplMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FactoryImplBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FactoryImpl{FactoryImplCaller: FactoryImplCaller{contract: contract}, FactoryImplTransactor: FactoryImplTransactor{contract: contract}, FactoryImplFilterer: FactoryImplFilterer{contract: contract}}, nil
}

// FactoryImpl is an auto generated Go binding around an Ethereum contract.
type FactoryImpl struct {
	FactoryImplCaller     // Read-only binding to the contract
	FactoryImplTransactor // Write-only binding to the contract
	FactoryImplFilterer   // Log filterer for contract events
}

// FactoryImplCaller is an auto generated read-only Go binding around an Ethereum contract.
type FactoryImplCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactoryImplTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FactoryImplTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactoryImplFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FactoryImplFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactoryImplSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FactoryImplSession struct {
	Contract     *FactoryImpl      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FactoryImplCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FactoryImplCallerSession struct {
	Contract *FactoryImplCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// FactoryImplTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FactoryImplTransactorSession struct {
	Contract     *FactoryImplTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// FactoryImplRaw is an auto generated low-level Go binding around an Ethereum contract.
type FactoryImplRaw struct {
	Contract *FactoryImpl // Generic contract binding to access the raw methods on
}

// FactoryImplCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FactoryImplCallerRaw struct {
	Contract *FactoryImplCaller // Generic read-only contract binding to access the raw methods on
}

// FactoryImplTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FactoryImplTransactorRaw struct {
	Contract *FactoryImplTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFactoryImpl creates a new instance of FactoryImpl, bound to a specific deployed contract.
func NewFactoryImpl(address common.Address, backend bind.ContractBackend) (*FactoryImpl, error) {
	contract, err := bindFactoryImpl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FactoryImpl{FactoryImplCaller: FactoryImplCaller{contract: contract}, FactoryImplTransactor: FactoryImplTransactor{contract: contract}, FactoryImplFilterer: FactoryImplFilterer{contract: contract}}, nil
}

// NewFactoryImplCaller creates a new read-only instance of FactoryImpl, bound to a specific deployed contract.
func NewFactoryImplCaller(address common.Address, caller bind.ContractCaller) (*FactoryImplCaller, error) {
	contract, err := bindFactoryImpl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FactoryImplCaller{contract: contract}, nil
}

// NewFactoryImplTransactor creates a new write-only instance of FactoryImpl, bound to a specific deployed contract.
func NewFactoryImplTransactor(address common.Address, transactor bind.ContractTransactor) (*FactoryImplTransactor, error) {
	contract, err := bindFactoryImpl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FactoryImplTransactor{contract: contract}, nil
}

// NewFactoryImplFilterer creates a new log filterer instance of FactoryImpl, bound to a specific deployed contract.
func NewFactoryImplFilterer(address common.Address, filterer bind.ContractFilterer) (*FactoryImplFilterer, error) {
	contract, err := bindFactoryImpl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FactoryImplFilterer{contract: contract}, nil
}

// bindFactoryImpl binds a generic wrapper to an already deployed contract.
func bindFactoryImpl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FactoryImplMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FactoryImpl *FactoryImplRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FactoryImpl.Contract.FactoryImplCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FactoryImpl *FactoryImplRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FactoryImpl.Contract.FactoryImplTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FactoryImpl *FactoryImplRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FactoryImpl.Contract.FactoryImplTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FactoryImpl *FactoryImplCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FactoryImpl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FactoryImpl *FactoryImplTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FactoryImpl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FactoryImpl *FactoryImplTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FactoryImpl.Contract.contract.Transact(opts, method, params...)
}

// QUOTE is a free data retrieval call binding the contract method 0x9c579839.
//
// Solidity: function QUOTE() view returns(address)
func (_FactoryImpl *FactoryImplCaller) QUOTE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FactoryImpl.contract.Call(opts, &out, "QUOTE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// QUOTE is a free data retrieval call binding the contract method 0x9c579839.
//
// Solidity: function QUOTE() view returns(address)
func (_FactoryImpl *FactoryImplSession) QUOTE() (common.Address, error) {
	return _FactoryImpl.Contract.QUOTE(&_FactoryImpl.CallOpts)
}

// QUOTE is a free data retrieval call binding the contract method 0x9c579839.
//
// Solidity: function QUOTE() view returns(address)
func (_FactoryImpl *FactoryImplCallerSession) QUOTE() (common.Address, error) {
	return _FactoryImpl.Contract.QUOTE(&_FactoryImpl.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_FactoryImpl *FactoryImplCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FactoryImpl.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_FactoryImpl *FactoryImplSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _FactoryImpl.Contract.UPGRADEINTERFACEVERSION(&_FactoryImpl.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_FactoryImpl *FactoryImplCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _FactoryImpl.Contract.UPGRADEINTERFACEVERSION(&_FactoryImpl.CallOpts)
}

// AllPairs is a free data retrieval call binding the contract method 0xc97682f8.
//
// Solidity: function allPairs() view returns(address[] bases, address[] pairs)
func (_FactoryImpl *FactoryImplCaller) AllPairs(opts *bind.CallOpts) (struct {
	Bases []common.Address
	Pairs []common.Address
}, error) {
	var out []interface{}
	err := _FactoryImpl.contract.Call(opts, &out, "allPairs")

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
func (_FactoryImpl *FactoryImplSession) AllPairs() (struct {
	Bases []common.Address
	Pairs []common.Address
}, error) {
	return _FactoryImpl.Contract.AllPairs(&_FactoryImpl.CallOpts)
}

// AllPairs is a free data retrieval call binding the contract method 0xc97682f8.
//
// Solidity: function allPairs() view returns(address[] bases, address[] pairs)
func (_FactoryImpl *FactoryImplCallerSession) AllPairs() (struct {
	Bases []common.Address
	Pairs []common.Address
}, error) {
	return _FactoryImpl.Contract.AllPairs(&_FactoryImpl.CallOpts)
}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_FactoryImpl *FactoryImplCaller) FeeCollector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FactoryImpl.contract.Call(opts, &out, "feeCollector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_FactoryImpl *FactoryImplSession) FeeCollector() (common.Address, error) {
	return _FactoryImpl.Contract.FeeCollector(&_FactoryImpl.CallOpts)
}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_FactoryImpl *FactoryImplCallerSession) FeeCollector() (common.Address, error) {
	return _FactoryImpl.Contract.FeeCollector(&_FactoryImpl.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FactoryImpl *FactoryImplCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FactoryImpl.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FactoryImpl *FactoryImplSession) Owner() (common.Address, error) {
	return _FactoryImpl.Contract.Owner(&_FactoryImpl.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FactoryImpl *FactoryImplCallerSession) Owner() (common.Address, error) {
	return _FactoryImpl.Contract.Owner(&_FactoryImpl.CallOpts)
}

// PairByBase is a free data retrieval call binding the contract method 0x43adeda9.
//
// Solidity: function pairByBase(address base) view returns(address)
func (_FactoryImpl *FactoryImplCaller) PairByBase(opts *bind.CallOpts, base common.Address) (common.Address, error) {
	var out []interface{}
	err := _FactoryImpl.contract.Call(opts, &out, "pairByBase", base)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PairByBase is a free data retrieval call binding the contract method 0x43adeda9.
//
// Solidity: function pairByBase(address base) view returns(address)
func (_FactoryImpl *FactoryImplSession) PairByBase(base common.Address) (common.Address, error) {
	return _FactoryImpl.Contract.PairByBase(&_FactoryImpl.CallOpts, base)
}

// PairByBase is a free data retrieval call binding the contract method 0x43adeda9.
//
// Solidity: function pairByBase(address base) view returns(address)
func (_FactoryImpl *FactoryImplCallerSession) PairByBase(base common.Address) (common.Address, error) {
	return _FactoryImpl.Contract.PairByBase(&_FactoryImpl.CallOpts, base)
}

// PairImpl is a free data retrieval call binding the contract method 0x35f7d456.
//
// Solidity: function pairImpl() view returns(address)
func (_FactoryImpl *FactoryImplCaller) PairImpl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FactoryImpl.contract.Call(opts, &out, "pairImpl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PairImpl is a free data retrieval call binding the contract method 0x35f7d456.
//
// Solidity: function pairImpl() view returns(address)
func (_FactoryImpl *FactoryImplSession) PairImpl() (common.Address, error) {
	return _FactoryImpl.Contract.PairImpl(&_FactoryImpl.CallOpts)
}

// PairImpl is a free data retrieval call binding the contract method 0x35f7d456.
//
// Solidity: function pairImpl() view returns(address)
func (_FactoryImpl *FactoryImplCallerSession) PairImpl() (common.Address, error) {
	return _FactoryImpl.Contract.PairImpl(&_FactoryImpl.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_FactoryImpl *FactoryImplCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FactoryImpl.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_FactoryImpl *FactoryImplSession) ProxiableUUID() ([32]byte, error) {
	return _FactoryImpl.Contract.ProxiableUUID(&_FactoryImpl.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_FactoryImpl *FactoryImplCallerSession) ProxiableUUID() ([32]byte, error) {
	return _FactoryImpl.Contract.ProxiableUUID(&_FactoryImpl.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_FactoryImpl *FactoryImplCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FactoryImpl.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_FactoryImpl *FactoryImplSession) Router() (common.Address, error) {
	return _FactoryImpl.Contract.Router(&_FactoryImpl.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_FactoryImpl *FactoryImplCallerSession) Router() (common.Address, error) {
	return _FactoryImpl.Contract.Router(&_FactoryImpl.CallOpts)
}

// CreatePair is a paid mutator transaction binding the contract method 0x9337bd68.
//
// Solidity: function createPair(address base, uint256 quoteTickSize, uint256 baseTickSize, uint256 makerFeePermil, uint256 takerFeePermil) returns(address pair)
func (_FactoryImpl *FactoryImplTransactor) CreatePair(opts *bind.TransactOpts, base common.Address, quoteTickSize *big.Int, baseTickSize *big.Int, makerFeePermil *big.Int, takerFeePermil *big.Int) (*types.Transaction, error) {
	return _FactoryImpl.contract.Transact(opts, "createPair", base, quoteTickSize, baseTickSize, makerFeePermil, takerFeePermil)
}

// CreatePair is a paid mutator transaction binding the contract method 0x9337bd68.
//
// Solidity: function createPair(address base, uint256 quoteTickSize, uint256 baseTickSize, uint256 makerFeePermil, uint256 takerFeePermil) returns(address pair)
func (_FactoryImpl *FactoryImplSession) CreatePair(base common.Address, quoteTickSize *big.Int, baseTickSize *big.Int, makerFeePermil *big.Int, takerFeePermil *big.Int) (*types.Transaction, error) {
	return _FactoryImpl.Contract.CreatePair(&_FactoryImpl.TransactOpts, base, quoteTickSize, baseTickSize, makerFeePermil, takerFeePermil)
}

// CreatePair is a paid mutator transaction binding the contract method 0x9337bd68.
//
// Solidity: function createPair(address base, uint256 quoteTickSize, uint256 baseTickSize, uint256 makerFeePermil, uint256 takerFeePermil) returns(address pair)
func (_FactoryImpl *FactoryImplTransactorSession) CreatePair(base common.Address, quoteTickSize *big.Int, baseTickSize *big.Int, makerFeePermil *big.Int, takerFeePermil *big.Int) (*types.Transaction, error) {
	return _FactoryImpl.Contract.CreatePair(&_FactoryImpl.TransactOpts, base, quoteTickSize, baseTickSize, makerFeePermil, takerFeePermil)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _router, address _feeCollector, address _quote, address _pairImpl) returns()
func (_FactoryImpl *FactoryImplTransactor) Initialize(opts *bind.TransactOpts, _router common.Address, _feeCollector common.Address, _quote common.Address, _pairImpl common.Address) (*types.Transaction, error) {
	return _FactoryImpl.contract.Transact(opts, "initialize", _router, _feeCollector, _quote, _pairImpl)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _router, address _feeCollector, address _quote, address _pairImpl) returns()
func (_FactoryImpl *FactoryImplSession) Initialize(_router common.Address, _feeCollector common.Address, _quote common.Address, _pairImpl common.Address) (*types.Transaction, error) {
	return _FactoryImpl.Contract.Initialize(&_FactoryImpl.TransactOpts, _router, _feeCollector, _quote, _pairImpl)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _router, address _feeCollector, address _quote, address _pairImpl) returns()
func (_FactoryImpl *FactoryImplTransactorSession) Initialize(_router common.Address, _feeCollector common.Address, _quote common.Address, _pairImpl common.Address) (*types.Transaction, error) {
	return _FactoryImpl.Contract.Initialize(&_FactoryImpl.TransactOpts, _router, _feeCollector, _quote, _pairImpl)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FactoryImpl *FactoryImplTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FactoryImpl.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FactoryImpl *FactoryImplSession) RenounceOwnership() (*types.Transaction, error) {
	return _FactoryImpl.Contract.RenounceOwnership(&_FactoryImpl.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FactoryImpl *FactoryImplTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FactoryImpl.Contract.RenounceOwnership(&_FactoryImpl.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FactoryImpl *FactoryImplTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FactoryImpl.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FactoryImpl *FactoryImplSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FactoryImpl.Contract.TransferOwnership(&_FactoryImpl.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FactoryImpl *FactoryImplTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FactoryImpl.Contract.TransferOwnership(&_FactoryImpl.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_FactoryImpl *FactoryImplTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _FactoryImpl.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_FactoryImpl *FactoryImplSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _FactoryImpl.Contract.UpgradeToAndCall(&_FactoryImpl.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_FactoryImpl *FactoryImplTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _FactoryImpl.Contract.UpgradeToAndCall(&_FactoryImpl.TransactOpts, newImplementation, data)
}

// FactoryImplInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the FactoryImpl contract.
type FactoryImplInitializedIterator struct {
	Event *FactoryImplInitialized // Event containing the contract specifics and raw log

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
func (it *FactoryImplInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactoryImplInitialized)
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
		it.Event = new(FactoryImplInitialized)
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
func (it *FactoryImplInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FactoryImplInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FactoryImplInitialized represents a Initialized event raised by the FactoryImpl contract.
type FactoryImplInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_FactoryImpl *FactoryImplFilterer) FilterInitialized(opts *bind.FilterOpts) (*FactoryImplInitializedIterator, error) {

	logs, sub, err := _FactoryImpl.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &FactoryImplInitializedIterator{contract: _FactoryImpl.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_FactoryImpl *FactoryImplFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *FactoryImplInitialized) (event.Subscription, error) {

	logs, sub, err := _FactoryImpl.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FactoryImplInitialized)
				if err := _FactoryImpl.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_FactoryImpl *FactoryImplFilterer) ParseInitialized(log types.Log) (*FactoryImplInitialized, error) {
	event := new(FactoryImplInitialized)
	if err := _FactoryImpl.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FactoryImplOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FactoryImpl contract.
type FactoryImplOwnershipTransferredIterator struct {
	Event *FactoryImplOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FactoryImplOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactoryImplOwnershipTransferred)
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
		it.Event = new(FactoryImplOwnershipTransferred)
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
func (it *FactoryImplOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FactoryImplOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FactoryImplOwnershipTransferred represents a OwnershipTransferred event raised by the FactoryImpl contract.
type FactoryImplOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FactoryImpl *FactoryImplFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FactoryImplOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FactoryImpl.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FactoryImplOwnershipTransferredIterator{contract: _FactoryImpl.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FactoryImpl *FactoryImplFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FactoryImplOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FactoryImpl.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FactoryImplOwnershipTransferred)
				if err := _FactoryImpl.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_FactoryImpl *FactoryImplFilterer) ParseOwnershipTransferred(log types.Log) (*FactoryImplOwnershipTransferred, error) {
	event := new(FactoryImplOwnershipTransferred)
	if err := _FactoryImpl.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FactoryImplPairCreatedIterator is returned from FilterPairCreated and is used to iterate over the raw logs and unpacked data for PairCreated events raised by the FactoryImpl contract.
type FactoryImplPairCreatedIterator struct {
	Event *FactoryImplPairCreated // Event containing the contract specifics and raw log

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
func (it *FactoryImplPairCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactoryImplPairCreated)
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
		it.Event = new(FactoryImplPairCreated)
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
func (it *FactoryImplPairCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FactoryImplPairCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FactoryImplPairCreated represents a PairCreated event raised by the FactoryImpl contract.
type FactoryImplPairCreated struct {
	Pair      common.Address
	Base      common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPairCreated is a free log retrieval operation binding the contract event 0xc1db9ba7c4b7ce660fe8d17bbcf07167549381df2abd694a970bd1402d86d313.
//
// Solidity: event PairCreated(address indexed pair, address indexed base, uint256 timestamp)
func (_FactoryImpl *FactoryImplFilterer) FilterPairCreated(opts *bind.FilterOpts, pair []common.Address, base []common.Address) (*FactoryImplPairCreatedIterator, error) {

	var pairRule []interface{}
	for _, pairItem := range pair {
		pairRule = append(pairRule, pairItem)
	}
	var baseRule []interface{}
	for _, baseItem := range base {
		baseRule = append(baseRule, baseItem)
	}

	logs, sub, err := _FactoryImpl.contract.FilterLogs(opts, "PairCreated", pairRule, baseRule)
	if err != nil {
		return nil, err
	}
	return &FactoryImplPairCreatedIterator{contract: _FactoryImpl.contract, event: "PairCreated", logs: logs, sub: sub}, nil
}

// WatchPairCreated is a free log subscription operation binding the contract event 0xc1db9ba7c4b7ce660fe8d17bbcf07167549381df2abd694a970bd1402d86d313.
//
// Solidity: event PairCreated(address indexed pair, address indexed base, uint256 timestamp)
func (_FactoryImpl *FactoryImplFilterer) WatchPairCreated(opts *bind.WatchOpts, sink chan<- *FactoryImplPairCreated, pair []common.Address, base []common.Address) (event.Subscription, error) {

	var pairRule []interface{}
	for _, pairItem := range pair {
		pairRule = append(pairRule, pairItem)
	}
	var baseRule []interface{}
	for _, baseItem := range base {
		baseRule = append(baseRule, baseItem)
	}

	logs, sub, err := _FactoryImpl.contract.WatchLogs(opts, "PairCreated", pairRule, baseRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FactoryImplPairCreated)
				if err := _FactoryImpl.contract.UnpackLog(event, "PairCreated", log); err != nil {
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
func (_FactoryImpl *FactoryImplFilterer) ParsePairCreated(log types.Log) (*FactoryImplPairCreated, error) {
	event := new(FactoryImplPairCreated)
	if err := _FactoryImpl.contract.UnpackLog(event, "PairCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FactoryImplUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the FactoryImpl contract.
type FactoryImplUpgradedIterator struct {
	Event *FactoryImplUpgraded // Event containing the contract specifics and raw log

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
func (it *FactoryImplUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactoryImplUpgraded)
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
		it.Event = new(FactoryImplUpgraded)
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
func (it *FactoryImplUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FactoryImplUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FactoryImplUpgraded represents a Upgraded event raised by the FactoryImpl contract.
type FactoryImplUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_FactoryImpl *FactoryImplFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*FactoryImplUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _FactoryImpl.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &FactoryImplUpgradedIterator{contract: _FactoryImpl.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_FactoryImpl *FactoryImplFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *FactoryImplUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _FactoryImpl.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FactoryImplUpgraded)
				if err := _FactoryImpl.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_FactoryImpl *FactoryImplFilterer) ParseUpgraded(log types.Log) (*FactoryImplUpgraded, error) {
	event := new(FactoryImplUpgraded)
	if err := _FactoryImpl.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
