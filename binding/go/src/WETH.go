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

// WETHMetaData contains all meta data concerning the WETH contract.
var WETHMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"ROUTER\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mintTo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"32fe7b26": "ROUTER()",
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"313ce567": "decimals()",
		"755edd17": "mintTo(address)",
		"06fdde03": "name()",
		"95d89b41": "symbol()",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
	},
	Bin: "0x60a060405234801561000f575f5ffd5b506040518060400160405280601681526020017f43726f737344455820577261707065642043524f535300000000000000000000815250604051806040016040528060068152602001655743524f535360d01b81525081600390816100749190610135565b5060046100818282610135565b5061008c9150503390565b6001600160a01b03166080526101ef565b634e487b7160e01b5f52604160045260245ffd5b600181811c908216806100c557607f821691505b6020821081036100e357634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561013057805f5260205f20601f840160051c8101602085101561010e5750805b601f840160051c820191505b8181101561012d575f815560010161011a565b50505b505050565b81516001600160401b0381111561014e5761014e61009d565b6101628161015c84546100b1565b846100e9565b6020601f821160018114610194575f831561017d5750848201515b5f19600385901b1c1916600184901b17845561012d565b5f84815260208120601f198516915b828110156101c357878501518255602094850194600190920191016101a3565b50848210156101e057868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b6080516109dc61020e5f395f818161017401526103d701526109dc5ff3fe60806040526004361061009d575f3560e01c806332fe7b261161006257806332fe7b261461016357806370a08231146101ae578063755edd17146101e257806395d89b41146101f5578063a9059cbb14610209578063dd62ed3e14610228575f5ffd5b806306fdde03146100b2578063095ea7b3146100dc57806318160ddd1461010b57806323b872dd14610129578063313ce56714610148575f5ffd5b366100ae576100ac333461026c565b005b5f5ffd5b3480156100bd575f5ffd5b506100c66102a9565b6040516100d3919061082d565b60405180910390f35b3480156100e7575f5ffd5b506100fb6100f636600461087d565b610339565b60405190151581526020016100d3565b348015610116575f5ffd5b506002545b6040519081526020016100d3565b348015610134575f5ffd5b506100fb6101433660046108a5565b610352565b348015610153575f5ffd5b50604051601281526020016100d3565b34801561016e575f5ffd5b506101967f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016100d3565b3480156101b9575f5ffd5b5061011b6101c83660046108df565b6001600160a01b03165f9081526020819052604090205490565b6100ac6101f03660046108df565b610375565b348015610200575f5ffd5b506100c6610382565b348015610214575f5ffd5b506100fb61022336600461087d565b610391565b348015610233575f5ffd5b5061011b6102423660046108ff565b6001600160a01b039182165f90815260016020908152604080832093909416825291909152205490565b6001600160a01b03821661029a5760405163ec442f0560e01b81525f60048201526024015b60405180910390fd5b6102a55f838361039e565b5050565b6060600380546102b890610930565b80601f01602080910402602001604051908101604052809291908181526020018280546102e490610930565b801561032f5780601f106103065761010080835404028352916020019161032f565b820191905f5260205f20905b81548152906001019060200180831161031257829003601f168201915b5050505050905090565b5f33610346818585610466565b60019150505b92915050565b5f3361035f858285610473565b61036a8585856104ef565b506001949350505050565b61037f813461026c565b50565b6060600480546102b890610930565b5f336103468185856104ef565b6103a983838361054c565b6001600160a01b038216156104615760405163e5e31b1360e01b81526001600160a01b0383811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063e5e31b1390602401602060405180830381865afa15801561041c573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104409190610968565b6104615761044e8282610672565b6104616001600160a01b038316826106a6565b505050565b6104618383836001610732565b6001600160a01b038381165f908152600160209081526040808320938616835292905220545f198110156104e957818110156104db57604051637dc7a0d960e11b81526001600160a01b03841660048201526024810182905260448101839052606401610291565b6104e984848484035f610732565b50505050565b6001600160a01b03831661051857604051634b637e8f60e11b81525f6004820152602401610291565b6001600160a01b0382166105415760405163ec442f0560e01b81525f6004820152602401610291565b61046183838361039e565b6001600160a01b038316610576578060025f82825461056b9190610987565b909155506105e69050565b6001600160a01b0383165f90815260208190526040902054818110156105c85760405163391434e360e21b81526001600160a01b03851660048201526024810182905260448101839052606401610291565b6001600160a01b0384165f9081526020819052604090209082900390555b6001600160a01b03821661060257600280548290039055610620565b6001600160a01b0382165f9081526020819052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161066591815260200190565b60405180910390a3505050565b6001600160a01b03821661069b57604051634b637e8f60e11b81525f6004820152602401610291565b6102a5825f8361039e565b804710156106d05760405163cf47918160e01b815247600482015260248101829052604401610291565b5f5f836001600160a01b0316836040515f6040518083038185875af1925050503d805f811461071a576040519150601f19603f3d011682016040523d82523d5f602084013e61071f565b606091505b5091509150816104e9576104e981610804565b6001600160a01b03841661075b5760405163e602df0560e01b81525f6004820152602401610291565b6001600160a01b03831661078457604051634a1406b160e11b81525f6004820152602401610291565b6001600160a01b038085165f90815260016020908152604080832093871683529290522082905580156104e957826001600160a01b0316846001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516107f691815260200190565b60405180910390a350505050565b8051156108145780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b80356001600160a01b0381168114610878575f5ffd5b919050565b5f5f6040838503121561088e575f5ffd5b61089783610862565b946020939093013593505050565b5f5f5f606084860312156108b7575f5ffd5b6108c084610862565b92506108ce60208501610862565b929592945050506040919091013590565b5f602082840312156108ef575f5ffd5b6108f882610862565b9392505050565b5f5f60408385031215610910575f5ffd5b61091983610862565b915061092760208401610862565b90509250929050565b600181811c9082168061094457607f821691505b60208210810361096257634e487b7160e01b5f52602260045260245ffd5b50919050565b5f60208284031215610978575f5ffd5b815180151581146108f8575f5ffd5b8082018082111561034c57634e487b7160e01b5f52601160045260245ffdfea2646970667358221220289c3b572a8895225126481d26bce4a30a84457d016eac383e2b8bf226ae0b0764736f6c634300081c0033",
}

// WETHABI is the input ABI used to generate the binding from.
// Deprecated: Use WETHMetaData.ABI instead.
var WETHABI = WETHMetaData.ABI

// WETHBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const WETHBinRuntime = "60806040526004361061009d575f3560e01c806332fe7b261161006257806332fe7b261461016357806370a08231146101ae578063755edd17146101e257806395d89b41146101f5578063a9059cbb14610209578063dd62ed3e14610228575f5ffd5b806306fdde03146100b2578063095ea7b3146100dc57806318160ddd1461010b57806323b872dd14610129578063313ce56714610148575f5ffd5b366100ae576100ac333461026c565b005b5f5ffd5b3480156100bd575f5ffd5b506100c66102a9565b6040516100d3919061082d565b60405180910390f35b3480156100e7575f5ffd5b506100fb6100f636600461087d565b610339565b60405190151581526020016100d3565b348015610116575f5ffd5b506002545b6040519081526020016100d3565b348015610134575f5ffd5b506100fb6101433660046108a5565b610352565b348015610153575f5ffd5b50604051601281526020016100d3565b34801561016e575f5ffd5b506101967f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016100d3565b3480156101b9575f5ffd5b5061011b6101c83660046108df565b6001600160a01b03165f9081526020819052604090205490565b6100ac6101f03660046108df565b610375565b348015610200575f5ffd5b506100c6610382565b348015610214575f5ffd5b506100fb61022336600461087d565b610391565b348015610233575f5ffd5b5061011b6102423660046108ff565b6001600160a01b039182165f90815260016020908152604080832093909416825291909152205490565b6001600160a01b03821661029a5760405163ec442f0560e01b81525f60048201526024015b60405180910390fd5b6102a55f838361039e565b5050565b6060600380546102b890610930565b80601f01602080910402602001604051908101604052809291908181526020018280546102e490610930565b801561032f5780601f106103065761010080835404028352916020019161032f565b820191905f5260205f20905b81548152906001019060200180831161031257829003601f168201915b5050505050905090565b5f33610346818585610466565b60019150505b92915050565b5f3361035f858285610473565b61036a8585856104ef565b506001949350505050565b61037f813461026c565b50565b6060600480546102b890610930565b5f336103468185856104ef565b6103a983838361054c565b6001600160a01b038216156104615760405163e5e31b1360e01b81526001600160a01b0383811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063e5e31b1390602401602060405180830381865afa15801561041c573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104409190610968565b6104615761044e8282610672565b6104616001600160a01b038316826106a6565b505050565b6104618383836001610732565b6001600160a01b038381165f908152600160209081526040808320938616835292905220545f198110156104e957818110156104db57604051637dc7a0d960e11b81526001600160a01b03841660048201526024810182905260448101839052606401610291565b6104e984848484035f610732565b50505050565b6001600160a01b03831661051857604051634b637e8f60e11b81525f6004820152602401610291565b6001600160a01b0382166105415760405163ec442f0560e01b81525f6004820152602401610291565b61046183838361039e565b6001600160a01b038316610576578060025f82825461056b9190610987565b909155506105e69050565b6001600160a01b0383165f90815260208190526040902054818110156105c85760405163391434e360e21b81526001600160a01b03851660048201526024810182905260448101839052606401610291565b6001600160a01b0384165f9081526020819052604090209082900390555b6001600160a01b03821661060257600280548290039055610620565b6001600160a01b0382165f9081526020819052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161066591815260200190565b60405180910390a3505050565b6001600160a01b03821661069b57604051634b637e8f60e11b81525f6004820152602401610291565b6102a5825f8361039e565b804710156106d05760405163cf47918160e01b815247600482015260248101829052604401610291565b5f5f836001600160a01b0316836040515f6040518083038185875af1925050503d805f811461071a576040519150601f19603f3d011682016040523d82523d5f602084013e61071f565b606091505b5091509150816104e9576104e981610804565b6001600160a01b03841661075b5760405163e602df0560e01b81525f6004820152602401610291565b6001600160a01b03831661078457604051634a1406b160e11b81525f6004820152602401610291565b6001600160a01b038085165f90815260016020908152604080832093871683529290522082905580156104e957826001600160a01b0316846001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516107f691815260200190565b60405180910390a350505050565b8051156108145780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b80356001600160a01b0381168114610878575f5ffd5b919050565b5f5f6040838503121561088e575f5ffd5b61089783610862565b946020939093013593505050565b5f5f5f606084860312156108b7575f5ffd5b6108c084610862565b92506108ce60208501610862565b929592945050506040919091013590565b5f602082840312156108ef575f5ffd5b6108f882610862565b9392505050565b5f5f60408385031215610910575f5ffd5b61091983610862565b915061092760208401610862565b90509250929050565b600181811c9082168061094457607f821691505b60208210810361096257634e487b7160e01b5f52602260045260245ffd5b50919050565b5f60208284031215610978575f5ffd5b815180151581146108f8575f5ffd5b8082018082111561034c57634e487b7160e01b5f52601160045260245ffdfea2646970667358221220289c3b572a8895225126481d26bce4a30a84457d016eac383e2b8bf226ae0b0764736f6c634300081c0033"

// Deprecated: Use WETHMetaData.Sigs instead.
// WETHFuncSigs maps the 4-byte function signature to its string representation.
var WETHFuncSigs = WETHMetaData.Sigs

// WETHBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WETHMetaData.Bin instead.
var WETHBin = WETHMetaData.Bin

// DeployWETH deploys a new Ethereum contract, binding an instance of WETH to it.
func DeployWETH(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WETH, error) {
	parsed, err := WETHMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WETHBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WETH{WETHCaller: WETHCaller{contract: contract}, WETHTransactor: WETHTransactor{contract: contract}, WETHFilterer: WETHFilterer{contract: contract}}, nil
}

// WETH is an auto generated Go binding around an Ethereum contract.
type WETH struct {
	WETHCaller     // Read-only binding to the contract
	WETHTransactor // Write-only binding to the contract
	WETHFilterer   // Log filterer for contract events
}

// WETHCaller is an auto generated read-only Go binding around an Ethereum contract.
type WETHCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WETHTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WETHTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WETHFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WETHFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WETHSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WETHSession struct {
	Contract     *WETH             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WETHCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WETHCallerSession struct {
	Contract *WETHCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// WETHTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WETHTransactorSession struct {
	Contract     *WETHTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WETHRaw is an auto generated low-level Go binding around an Ethereum contract.
type WETHRaw struct {
	Contract *WETH // Generic contract binding to access the raw methods on
}

// WETHCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WETHCallerRaw struct {
	Contract *WETHCaller // Generic read-only contract binding to access the raw methods on
}

// WETHTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WETHTransactorRaw struct {
	Contract *WETHTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWETH creates a new instance of WETH, bound to a specific deployed contract.
func NewWETH(address common.Address, backend bind.ContractBackend) (*WETH, error) {
	contract, err := bindWETH(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WETH{WETHCaller: WETHCaller{contract: contract}, WETHTransactor: WETHTransactor{contract: contract}, WETHFilterer: WETHFilterer{contract: contract}}, nil
}

// NewWETHCaller creates a new read-only instance of WETH, bound to a specific deployed contract.
func NewWETHCaller(address common.Address, caller bind.ContractCaller) (*WETHCaller, error) {
	contract, err := bindWETH(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WETHCaller{contract: contract}, nil
}

// NewWETHTransactor creates a new write-only instance of WETH, bound to a specific deployed contract.
func NewWETHTransactor(address common.Address, transactor bind.ContractTransactor) (*WETHTransactor, error) {
	contract, err := bindWETH(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WETHTransactor{contract: contract}, nil
}

// NewWETHFilterer creates a new log filterer instance of WETH, bound to a specific deployed contract.
func NewWETHFilterer(address common.Address, filterer bind.ContractFilterer) (*WETHFilterer, error) {
	contract, err := bindWETH(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WETHFilterer{contract: contract}, nil
}

// bindWETH binds a generic wrapper to an already deployed contract.
func bindWETH(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WETHMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WETH *WETHRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WETH.Contract.WETHCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WETH *WETHRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WETH.Contract.WETHTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WETH *WETHRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WETH.Contract.WETHTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WETH *WETHCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WETH.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WETH *WETHTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WETH.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WETH *WETHTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WETH.Contract.contract.Transact(opts, method, params...)
}

// ROUTER is a free data retrieval call binding the contract method 0x32fe7b26.
//
// Solidity: function ROUTER() view returns(address)
func (_WETH *WETHCaller) ROUTER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WETH.contract.Call(opts, &out, "ROUTER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ROUTER is a free data retrieval call binding the contract method 0x32fe7b26.
//
// Solidity: function ROUTER() view returns(address)
func (_WETH *WETHSession) ROUTER() (common.Address, error) {
	return _WETH.Contract.ROUTER(&_WETH.CallOpts)
}

// ROUTER is a free data retrieval call binding the contract method 0x32fe7b26.
//
// Solidity: function ROUTER() view returns(address)
func (_WETH *WETHCallerSession) ROUTER() (common.Address, error) {
	return _WETH.Contract.ROUTER(&_WETH.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WETH *WETHCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WETH.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WETH *WETHSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _WETH.Contract.Allowance(&_WETH.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WETH *WETHCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _WETH.Contract.Allowance(&_WETH.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WETH *WETHCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WETH.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WETH *WETHSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _WETH.Contract.BalanceOf(&_WETH.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WETH *WETHCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _WETH.Contract.BalanceOf(&_WETH.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WETH *WETHCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _WETH.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WETH *WETHSession) Decimals() (uint8, error) {
	return _WETH.Contract.Decimals(&_WETH.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WETH *WETHCallerSession) Decimals() (uint8, error) {
	return _WETH.Contract.Decimals(&_WETH.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WETH *WETHCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WETH.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WETH *WETHSession) Name() (string, error) {
	return _WETH.Contract.Name(&_WETH.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WETH *WETHCallerSession) Name() (string, error) {
	return _WETH.Contract.Name(&_WETH.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WETH *WETHCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WETH.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WETH *WETHSession) Symbol() (string, error) {
	return _WETH.Contract.Symbol(&_WETH.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WETH *WETHCallerSession) Symbol() (string, error) {
	return _WETH.Contract.Symbol(&_WETH.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WETH *WETHCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WETH.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WETH *WETHSession) TotalSupply() (*big.Int, error) {
	return _WETH.Contract.TotalSupply(&_WETH.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WETH *WETHCallerSession) TotalSupply() (*big.Int, error) {
	return _WETH.Contract.TotalSupply(&_WETH.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_WETH *WETHTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _WETH.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_WETH *WETHSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _WETH.Contract.Approve(&_WETH.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_WETH *WETHTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _WETH.Contract.Approve(&_WETH.TransactOpts, spender, value)
}

// MintTo is a paid mutator transaction binding the contract method 0x755edd17.
//
// Solidity: function mintTo(address to) payable returns()
func (_WETH *WETHTransactor) MintTo(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _WETH.contract.Transact(opts, "mintTo", to)
}

// MintTo is a paid mutator transaction binding the contract method 0x755edd17.
//
// Solidity: function mintTo(address to) payable returns()
func (_WETH *WETHSession) MintTo(to common.Address) (*types.Transaction, error) {
	return _WETH.Contract.MintTo(&_WETH.TransactOpts, to)
}

// MintTo is a paid mutator transaction binding the contract method 0x755edd17.
//
// Solidity: function mintTo(address to) payable returns()
func (_WETH *WETHTransactorSession) MintTo(to common.Address) (*types.Transaction, error) {
	return _WETH.Contract.MintTo(&_WETH.TransactOpts, to)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_WETH *WETHTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WETH.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_WETH *WETHSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WETH.Contract.Transfer(&_WETH.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_WETH *WETHTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WETH.Contract.Transfer(&_WETH.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_WETH *WETHTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WETH.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_WETH *WETHSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WETH.Contract.TransferFrom(&_WETH.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_WETH *WETHTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WETH.Contract.TransferFrom(&_WETH.TransactOpts, from, to, value)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WETH *WETHTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WETH.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WETH *WETHSession) Receive() (*types.Transaction, error) {
	return _WETH.Contract.Receive(&_WETH.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WETH *WETHTransactorSession) Receive() (*types.Transaction, error) {
	return _WETH.Contract.Receive(&_WETH.TransactOpts)
}

// WETHApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the WETH contract.
type WETHApprovalIterator struct {
	Event *WETHApproval // Event containing the contract specifics and raw log

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
func (it *WETHApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WETHApproval)
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
		it.Event = new(WETHApproval)
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
func (it *WETHApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WETHApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WETHApproval represents a Approval event raised by the WETH contract.
type WETHApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WETH *WETHFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*WETHApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WETH.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &WETHApprovalIterator{contract: _WETH.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WETH *WETHFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *WETHApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WETH.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WETHApproval)
				if err := _WETH.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WETH *WETHFilterer) ParseApproval(log types.Log) (*WETHApproval, error) {
	event := new(WETHApproval)
	if err := _WETH.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WETHTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the WETH contract.
type WETHTransferIterator struct {
	Event *WETHTransfer // Event containing the contract specifics and raw log

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
func (it *WETHTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WETHTransfer)
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
		it.Event = new(WETHTransfer)
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
func (it *WETHTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WETHTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WETHTransfer represents a Transfer event raised by the WETH contract.
type WETHTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WETH *WETHFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WETHTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WETH.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WETHTransferIterator{contract: _WETH.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WETH *WETHFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *WETHTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WETH.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WETHTransfer)
				if err := _WETH.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WETH *WETHFilterer) ParseTransfer(log types.Log) (*WETHTransfer, error) {
	event := new(WETHTransfer)
	if err := _WETH.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
