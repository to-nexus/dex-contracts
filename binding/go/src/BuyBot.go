// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package src

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

// BuyBotMetaData contains all meta data concerning the BuyBot contract.
var BuyBotMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_minOrderAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_interval\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"buyMarket\",\"inputs\":[{\"name\":\"pair\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxMatchCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"canBuyMarket\",\"inputs\":[{\"name\":\"pair\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"canBuy\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBalance\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"interval\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastBuyTime\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minOrderAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"router\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRouter\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setInterval\",\"inputs\":[{\"name\":\"_interval\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinOrderAmount\",\"inputs\":[{\"name\":\"_minOrderAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawETH\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"IntervalSet\",\"inputs\":[{\"name\":\"before\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"current\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MarketBuyExecuted\",\"inputs\":[{\"name\":\"pair\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"quoteToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"baseToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"quoteAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"executor\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinOrderAmountSet\",\"inputs\":[{\"name\":\"before\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"current\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BuyBotInsufficientBalance\",\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minOrderAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"BuyBotIntervalNotPassed\",\"inputs\":[{\"name\":\"timeSinceLastBuy\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requiredInterval\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"BuyBotInvalidMinOrderAmount\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"BuyBotInvalidPair\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"BuyBotInvalidRouter\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x608060405234801561000f575f5ffd5b5060405161158b38038061158b83398101604081905261002e916101ac565b836001600160a01b03811661005d57604051631e4fbdf760e01b81525f60048201526024015b60405180910390fd5b61006681610142565b506001600160a01b03831661009957604051631561f27b60e21b81526001600160a01b0384166004820152602401610054565b815f036100bc576040516328bfc81960e11b815260048101839052602401610054565b600180546001600160a01b0319166001600160a01b0385161790556002829055600381905560405182905f907fd6d62b78d21fc4f5151029ca37079020e8a8815c5a488b28ee7499d4c2a19e55908290a360405181905f907f3fb677206c0b314c404bae3da94bee6bda11375c9fe986f266fb033d6cddbbd7908290a3505050506101ec565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b80516001600160a01b03811681146101a7575f5ffd5b919050565b5f5f5f5f608085870312156101bf575f5ffd5b6101c885610191565b93506101d660208601610191565b6040860151606090960151949790965092505050565b611392806101f95f395ff3fe6080604052600436106100dc575f3560e01c8063a3b8ef041161007c578063f2fde38b11610057578063f2fde38b1461024b578063f3fef3a31461026a578063f887ea4014610289578063f8b2cb4f146102b5575f5ffd5b8063a3b8ef04146101f8578063f14210a614610217578063f29f4d0b14610236575f5ffd5b80635959cdb7116100b75780635959cdb71461014f578063715018a6146101855780638da5cb5b14610199578063947a36fb146101e3575f5ffd5b806322a90082146100e7578063413ec7781461010857806346b62c4a14610127575f5ffd5b366100e357005b5f5ffd5b3480156100f2575f5ffd5b506101066101013660046111b4565b6102d4565b005b348015610113575f5ffd5b506101066101223660046111ec565b61030f565b348015610132575f5ffd5b5061013c60025481565b6040519081526020015b60405180910390f35b34801561015a575f5ffd5b5061016e610169366004611231565b610845565b604080519215158352602083019190915201610146565b348015610190575f5ffd5b5061010661097a565b3480156101a4575f5ffd5b505f5473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610146565b3480156101ee575f5ffd5b5061013c60035481565b348015610203575f5ffd5b506101066102123660046111b4565b61098d565b348015610222575f5ffd5b506101066102313660046111b4565b610a04565b348015610241575f5ffd5b5061013c60045481565b348015610256575f5ffd5b50610106610265366004611231565b610ba5565b348015610275575f5ffd5b50610106610284366004611253565b610c08565b348015610294575f5ffd5b506001546101be9073ffffffffffffffffffffffffffffffffffffffff1681565b3480156102c0575f5ffd5b5061013c6102cf366004611231565b610de1565b6102dc610e96565b6003546040518291907f3fb677206c0b314c404bae3da94bee6bda11375c9fe986f266fb033d6cddbbd7905f90a3600355565b73ffffffffffffffffffffffffffffffffffffffff8416610379576040517f235aafe400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024015b60405180910390fd5b5f60035411801561038b57505f600454115b156103ed575f6004544261039f919061127d565b90506003548110156103eb576003546040517f1e379906000000000000000000000000000000000000000000000000000000008152610370918391600401918252602082015260400190565b505b5f8473ffffffffffffffffffffffffffffffffffffffff1663c3f909d46040518163ffffffff1660e01b8152600401606060405180830381865afa158015610437573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061045b91906112b5565b805160208201516040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015292935090915f9073ffffffffffffffffffffffffffffffffffffffff8416906370a0823190602401602060405180830381865afa1580156104d1573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104f59190611345565b90505f87156105045787610506565b815b9050600254811015610552576002546040517f013fafe2000000000000000000000000000000000000000000000000000000008152610370918391600401918252602082015260400190565b81811115610596576040517f013fafe20000000000000000000000000000000000000000000000000000000081526004810183905260248101829052604401610370565b6001546040517fdd62ed3e00000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff918216602482018190529183919087169063dd62ed3e90604401602060405180830381865afa15801561060f573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106339190611345565b101561067a5761067a73ffffffffffffffffffffffffffffffffffffffff8616827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff610ee8565b6001546040517f1e92008400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8c8116600483015260248201859052604482018a905290911690631e920084906064015f604051808303815f87803b1580156106f2575f5ffd5b505af1158015610704573d5f5f3e3d5ffd5b50506040805185815233602082015273ffffffffffffffffffffffffffffffffffffffff808916945089811693508e16917fc7b4d815bd0a3b8348577971cb55a9334c80a0b3ccdfe287b55b5a6320d8480a910160405180910390a44260045573ffffffffffffffffffffffffffffffffffffffff881615610839576040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201525f9073ffffffffffffffffffffffffffffffffffffffff8616906370a0823190602401602060405180830381865afa1580156107ea573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061080e9190611345565b905080156108375761083773ffffffffffffffffffffffffffffffffffffffff86168a83611008565b505b50505050505050505050565b5f8073ffffffffffffffffffffffffffffffffffffffff831661086c57505f928392509050565b5f8373ffffffffffffffffffffffffffffffffffffffff1663c3f909d46040518163ffffffff1660e01b8152600401606060405180830381865afa1580156108b6573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108da91906112b5565b80516040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015291925073ffffffffffffffffffffffffffffffffffffffff16906370a0823190602401602060405180830381865afa158015610946573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061096a9190611345565b9150600254821015925050915091565b610982610e96565b61098b5f61104b565b565b610995610e96565b805f036109d1576040517f517f903200000000000000000000000000000000000000000000000000000000815260048101829052602401610370565b6002546040518291907fd6d62b78d21fc4f5151029ca37079020e8a8815c5a488b28ee7499d4c2a19e55905f90a3600255565b610a0c610e96565b475f8215610a1a5782610a1c565b815b905081811115610a88576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f496e73756666696369656e742062616c616e63650000000000000000000000006044820152606401610370565b5f805460405173ffffffffffffffffffffffffffffffffffffffff9091169083908381818185875af1925050503d805f8114610adf576040519150601f19603f3d011682016040523d82523d5f602084013e610ae4565b606091505b5050905080610b4f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f455448207472616e73666572206661696c6564000000000000000000000000006044820152606401610370565b5f805460405184815273ffffffffffffffffffffffffffffffffffffffff90911691907fd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb9060200160405180910390a350505050565b610bad610e96565b73ffffffffffffffffffffffffffffffffffffffff8116610bfc576040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081525f6004820152602401610370565b610c058161104b565b50565b610c10610e96565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015282905f9073ffffffffffffffffffffffffffffffffffffffff8316906370a0823190602401602060405180830381865afa158015610c7c573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ca09190611345565b90505f8315610caf5783610cb1565b815b905081811115610d1d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f496e73756666696369656e742062616c616e63650000000000000000000000006044820152606401610370565b610d5c610d3e5f5473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff85169083611008565b5f5473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167fd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb83604051610dd291815260200190565b60405180910390a35050505050565b5f73ffffffffffffffffffffffffffffffffffffffff8216610e04575047919050565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff8316906370a0823190602401602060405180830381865afa158015610e6c573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e909190611345565b92915050565b5f5473ffffffffffffffffffffffffffffffffffffffff16331461098b576040517f118cdaa7000000000000000000000000000000000000000000000000000000008152336004820152602401610370565b6040805173ffffffffffffffffffffffffffffffffffffffff8416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f095ea7b300000000000000000000000000000000000000000000000000000000179052610f7484826110bf565b6110025760405173ffffffffffffffffffffffffffffffffffffffff84811660248301525f6044830152610ff891869182169063095ea7b3906064015b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611115565b6110028482611115565b50505050565b60405173ffffffffffffffffffffffffffffffffffffffff83811660248301526044820183905261104691859182169063a9059cbb90606401610fb1565b505050565b5f805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b5f5f5f5f60205f8651602088015f8a5af192503d91505f51905082801561110b575081156110f0578060011461110b565b5f8673ffffffffffffffffffffffffffffffffffffffff163b115b9695505050505050565b5f5f60205f8451602086015f885af180611134576040513d5f823e3d81fd5b50505f513d9150811561114b578060011415611165565b73ffffffffffffffffffffffffffffffffffffffff84163b155b15611002576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610370565b5f602082840312156111c4575f5ffd5b5035919050565b73ffffffffffffffffffffffffffffffffffffffff81168114610c05575f5ffd5b5f5f5f5f608085870312156111ff575f5ffd5b843561120a816111cb565b9350602085013592506040850135611221816111cb565b9396929550929360600135925050565b5f60208284031215611241575f5ffd5b813561124c816111cb565b9392505050565b5f5f60408385031215611264575f5ffd5b823561126f816111cb565b946020939093013593505050565b81810381811115610e90577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f60608284031280156112c6575f5ffd5b506040516060810167ffffffffffffffff8111828210171561130f577f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604052825161131d816111cb565b8152602083015161132d816111cb565b60208201526040928301519281019290925250919050565b5f60208284031215611355575f5ffd5b505191905056fea2646970667358221220acff4ca952838a28697624b35f0dfd429ef2b7d38cc71e762099d5a835bbaa3464736f6c634300081c0033",
}

// BuyBotABI is the input ABI used to generate the binding from.
// Deprecated: Use BuyBotMetaData.ABI instead.
var BuyBotABI = BuyBotMetaData.ABI

// BuyBotBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BuyBotMetaData.Bin instead.
var BuyBotBin = BuyBotMetaData.Bin

// DeployBuyBot deploys a new Ethereum contract, binding an instance of BuyBot to it.
func DeployBuyBot(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address, _router common.Address, _minOrderAmount *big.Int, _interval *big.Int) (common.Address, *types.Transaction, *BuyBot, error) {
	parsed, err := BuyBotMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BuyBotBin), backend, _owner, _router, _minOrderAmount, _interval)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BuyBot{BuyBotCaller: BuyBotCaller{contract: contract}, BuyBotTransactor: BuyBotTransactor{contract: contract}, BuyBotFilterer: BuyBotFilterer{contract: contract}}, nil
}

// BuyBot is an auto generated Go binding around an Ethereum contract.
type BuyBot struct {
	BuyBotCaller     // Read-only binding to the contract
	BuyBotTransactor // Write-only binding to the contract
	BuyBotFilterer   // Log filterer for contract events
}

// BuyBotCaller is an auto generated read-only Go binding around an Ethereum contract.
type BuyBotCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BuyBotTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BuyBotTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BuyBotFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BuyBotFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BuyBotSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BuyBotSession struct {
	Contract     *BuyBot           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BuyBotCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BuyBotCallerSession struct {
	Contract *BuyBotCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BuyBotTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BuyBotTransactorSession struct {
	Contract     *BuyBotTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BuyBotRaw is an auto generated low-level Go binding around an Ethereum contract.
type BuyBotRaw struct {
	Contract *BuyBot // Generic contract binding to access the raw methods on
}

// BuyBotCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BuyBotCallerRaw struct {
	Contract *BuyBotCaller // Generic read-only contract binding to access the raw methods on
}

// BuyBotTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BuyBotTransactorRaw struct {
	Contract *BuyBotTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBuyBot creates a new instance of BuyBot, bound to a specific deployed contract.
func NewBuyBot(address common.Address, backend bind.ContractBackend) (*BuyBot, error) {
	contract, err := bindBuyBot(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BuyBot{BuyBotCaller: BuyBotCaller{contract: contract}, BuyBotTransactor: BuyBotTransactor{contract: contract}, BuyBotFilterer: BuyBotFilterer{contract: contract}}, nil
}

// NewBuyBotCaller creates a new read-only instance of BuyBot, bound to a specific deployed contract.
func NewBuyBotCaller(address common.Address, caller bind.ContractCaller) (*BuyBotCaller, error) {
	contract, err := bindBuyBot(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BuyBotCaller{contract: contract}, nil
}

// NewBuyBotTransactor creates a new write-only instance of BuyBot, bound to a specific deployed contract.
func NewBuyBotTransactor(address common.Address, transactor bind.ContractTransactor) (*BuyBotTransactor, error) {
	contract, err := bindBuyBot(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BuyBotTransactor{contract: contract}, nil
}

// NewBuyBotFilterer creates a new log filterer instance of BuyBot, bound to a specific deployed contract.
func NewBuyBotFilterer(address common.Address, filterer bind.ContractFilterer) (*BuyBotFilterer, error) {
	contract, err := bindBuyBot(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BuyBotFilterer{contract: contract}, nil
}

// bindBuyBot binds a generic wrapper to an already deployed contract.
func bindBuyBot(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BuyBotMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BuyBot *BuyBotRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BuyBot.Contract.BuyBotCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BuyBot *BuyBotRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BuyBot.Contract.BuyBotTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BuyBot *BuyBotRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BuyBot.Contract.BuyBotTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BuyBot *BuyBotCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BuyBot.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BuyBot *BuyBotTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BuyBot.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BuyBot *BuyBotTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BuyBot.Contract.contract.Transact(opts, method, params...)
}

// CanBuyMarket is a free data retrieval call binding the contract method 0x5959cdb7.
//
// Solidity: function canBuyMarket(address pair) view returns(bool canBuy, uint256 balance)
func (_BuyBot *BuyBotCaller) CanBuyMarket(opts *bind.CallOpts, pair common.Address) (struct {
	CanBuy  bool
	Balance *big.Int
}, error) {
	var out []interface{}
	err := _BuyBot.contract.Call(opts, &out, "canBuyMarket", pair)

	outstruct := new(struct {
		CanBuy  bool
		Balance *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CanBuy = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Balance = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CanBuyMarket is a free data retrieval call binding the contract method 0x5959cdb7.
//
// Solidity: function canBuyMarket(address pair) view returns(bool canBuy, uint256 balance)
func (_BuyBot *BuyBotSession) CanBuyMarket(pair common.Address) (struct {
	CanBuy  bool
	Balance *big.Int
}, error) {
	return _BuyBot.Contract.CanBuyMarket(&_BuyBot.CallOpts, pair)
}

// CanBuyMarket is a free data retrieval call binding the contract method 0x5959cdb7.
//
// Solidity: function canBuyMarket(address pair) view returns(bool canBuy, uint256 balance)
func (_BuyBot *BuyBotCallerSession) CanBuyMarket(pair common.Address) (struct {
	CanBuy  bool
	Balance *big.Int
}, error) {
	return _BuyBot.Contract.CanBuyMarket(&_BuyBot.CallOpts, pair)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address token) view returns(uint256 balance)
func (_BuyBot *BuyBotCaller) GetBalance(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BuyBot.contract.Call(opts, &out, "getBalance", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address token) view returns(uint256 balance)
func (_BuyBot *BuyBotSession) GetBalance(token common.Address) (*big.Int, error) {
	return _BuyBot.Contract.GetBalance(&_BuyBot.CallOpts, token)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address token) view returns(uint256 balance)
func (_BuyBot *BuyBotCallerSession) GetBalance(token common.Address) (*big.Int, error) {
	return _BuyBot.Contract.GetBalance(&_BuyBot.CallOpts, token)
}

// Interval is a free data retrieval call binding the contract method 0x947a36fb.
//
// Solidity: function interval() view returns(uint256)
func (_BuyBot *BuyBotCaller) Interval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BuyBot.contract.Call(opts, &out, "interval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Interval is a free data retrieval call binding the contract method 0x947a36fb.
//
// Solidity: function interval() view returns(uint256)
func (_BuyBot *BuyBotSession) Interval() (*big.Int, error) {
	return _BuyBot.Contract.Interval(&_BuyBot.CallOpts)
}

// Interval is a free data retrieval call binding the contract method 0x947a36fb.
//
// Solidity: function interval() view returns(uint256)
func (_BuyBot *BuyBotCallerSession) Interval() (*big.Int, error) {
	return _BuyBot.Contract.Interval(&_BuyBot.CallOpts)
}

// LastBuyTime is a free data retrieval call binding the contract method 0xf29f4d0b.
//
// Solidity: function lastBuyTime() view returns(uint256)
func (_BuyBot *BuyBotCaller) LastBuyTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BuyBot.contract.Call(opts, &out, "lastBuyTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastBuyTime is a free data retrieval call binding the contract method 0xf29f4d0b.
//
// Solidity: function lastBuyTime() view returns(uint256)
func (_BuyBot *BuyBotSession) LastBuyTime() (*big.Int, error) {
	return _BuyBot.Contract.LastBuyTime(&_BuyBot.CallOpts)
}

// LastBuyTime is a free data retrieval call binding the contract method 0xf29f4d0b.
//
// Solidity: function lastBuyTime() view returns(uint256)
func (_BuyBot *BuyBotCallerSession) LastBuyTime() (*big.Int, error) {
	return _BuyBot.Contract.LastBuyTime(&_BuyBot.CallOpts)
}

// MinOrderAmount is a free data retrieval call binding the contract method 0x46b62c4a.
//
// Solidity: function minOrderAmount() view returns(uint256)
func (_BuyBot *BuyBotCaller) MinOrderAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BuyBot.contract.Call(opts, &out, "minOrderAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinOrderAmount is a free data retrieval call binding the contract method 0x46b62c4a.
//
// Solidity: function minOrderAmount() view returns(uint256)
func (_BuyBot *BuyBotSession) MinOrderAmount() (*big.Int, error) {
	return _BuyBot.Contract.MinOrderAmount(&_BuyBot.CallOpts)
}

// MinOrderAmount is a free data retrieval call binding the contract method 0x46b62c4a.
//
// Solidity: function minOrderAmount() view returns(uint256)
func (_BuyBot *BuyBotCallerSession) MinOrderAmount() (*big.Int, error) {
	return _BuyBot.Contract.MinOrderAmount(&_BuyBot.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BuyBot *BuyBotCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BuyBot.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BuyBot *BuyBotSession) Owner() (common.Address, error) {
	return _BuyBot.Contract.Owner(&_BuyBot.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BuyBot *BuyBotCallerSession) Owner() (common.Address, error) {
	return _BuyBot.Contract.Owner(&_BuyBot.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_BuyBot *BuyBotCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BuyBot.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_BuyBot *BuyBotSession) Router() (common.Address, error) {
	return _BuyBot.Contract.Router(&_BuyBot.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_BuyBot *BuyBotCallerSession) Router() (common.Address, error) {
	return _BuyBot.Contract.Router(&_BuyBot.CallOpts)
}

// BuyMarket is a paid mutator transaction binding the contract method 0x413ec778.
//
// Solidity: function buyMarket(address pair, uint256 amount, address recipient, uint256 maxMatchCount) returns()
func (_BuyBot *BuyBotTransactor) BuyMarket(opts *bind.TransactOpts, pair common.Address, amount *big.Int, recipient common.Address, maxMatchCount *big.Int) (*types.Transaction, error) {
	return _BuyBot.contract.Transact(opts, "buyMarket", pair, amount, recipient, maxMatchCount)
}

// BuyMarket is a paid mutator transaction binding the contract method 0x413ec778.
//
// Solidity: function buyMarket(address pair, uint256 amount, address recipient, uint256 maxMatchCount) returns()
func (_BuyBot *BuyBotSession) BuyMarket(pair common.Address, amount *big.Int, recipient common.Address, maxMatchCount *big.Int) (*types.Transaction, error) {
	return _BuyBot.Contract.BuyMarket(&_BuyBot.TransactOpts, pair, amount, recipient, maxMatchCount)
}

// BuyMarket is a paid mutator transaction binding the contract method 0x413ec778.
//
// Solidity: function buyMarket(address pair, uint256 amount, address recipient, uint256 maxMatchCount) returns()
func (_BuyBot *BuyBotTransactorSession) BuyMarket(pair common.Address, amount *big.Int, recipient common.Address, maxMatchCount *big.Int) (*types.Transaction, error) {
	return _BuyBot.Contract.BuyMarket(&_BuyBot.TransactOpts, pair, amount, recipient, maxMatchCount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BuyBot *BuyBotTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BuyBot.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BuyBot *BuyBotSession) RenounceOwnership() (*types.Transaction, error) {
	return _BuyBot.Contract.RenounceOwnership(&_BuyBot.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BuyBot *BuyBotTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BuyBot.Contract.RenounceOwnership(&_BuyBot.TransactOpts)
}

// SetInterval is a paid mutator transaction binding the contract method 0x22a90082.
//
// Solidity: function setInterval(uint256 _interval) returns()
func (_BuyBot *BuyBotTransactor) SetInterval(opts *bind.TransactOpts, _interval *big.Int) (*types.Transaction, error) {
	return _BuyBot.contract.Transact(opts, "setInterval", _interval)
}

// SetInterval is a paid mutator transaction binding the contract method 0x22a90082.
//
// Solidity: function setInterval(uint256 _interval) returns()
func (_BuyBot *BuyBotSession) SetInterval(_interval *big.Int) (*types.Transaction, error) {
	return _BuyBot.Contract.SetInterval(&_BuyBot.TransactOpts, _interval)
}

// SetInterval is a paid mutator transaction binding the contract method 0x22a90082.
//
// Solidity: function setInterval(uint256 _interval) returns()
func (_BuyBot *BuyBotTransactorSession) SetInterval(_interval *big.Int) (*types.Transaction, error) {
	return _BuyBot.Contract.SetInterval(&_BuyBot.TransactOpts, _interval)
}

// SetMinOrderAmount is a paid mutator transaction binding the contract method 0xa3b8ef04.
//
// Solidity: function setMinOrderAmount(uint256 _minOrderAmount) returns()
func (_BuyBot *BuyBotTransactor) SetMinOrderAmount(opts *bind.TransactOpts, _minOrderAmount *big.Int) (*types.Transaction, error) {
	return _BuyBot.contract.Transact(opts, "setMinOrderAmount", _minOrderAmount)
}

// SetMinOrderAmount is a paid mutator transaction binding the contract method 0xa3b8ef04.
//
// Solidity: function setMinOrderAmount(uint256 _minOrderAmount) returns()
func (_BuyBot *BuyBotSession) SetMinOrderAmount(_minOrderAmount *big.Int) (*types.Transaction, error) {
	return _BuyBot.Contract.SetMinOrderAmount(&_BuyBot.TransactOpts, _minOrderAmount)
}

// SetMinOrderAmount is a paid mutator transaction binding the contract method 0xa3b8ef04.
//
// Solidity: function setMinOrderAmount(uint256 _minOrderAmount) returns()
func (_BuyBot *BuyBotTransactorSession) SetMinOrderAmount(_minOrderAmount *big.Int) (*types.Transaction, error) {
	return _BuyBot.Contract.SetMinOrderAmount(&_BuyBot.TransactOpts, _minOrderAmount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BuyBot *BuyBotTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BuyBot.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BuyBot *BuyBotSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BuyBot.Contract.TransferOwnership(&_BuyBot.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BuyBot *BuyBotTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BuyBot.Contract.TransferOwnership(&_BuyBot.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address token, uint256 amount) returns()
func (_BuyBot *BuyBotTransactor) Withdraw(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BuyBot.contract.Transact(opts, "withdraw", token, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address token, uint256 amount) returns()
func (_BuyBot *BuyBotSession) Withdraw(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BuyBot.Contract.Withdraw(&_BuyBot.TransactOpts, token, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address token, uint256 amount) returns()
func (_BuyBot *BuyBotTransactorSession) Withdraw(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BuyBot.Contract.Withdraw(&_BuyBot.TransactOpts, token, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xf14210a6.
//
// Solidity: function withdrawETH(uint256 amount) returns()
func (_BuyBot *BuyBotTransactor) WithdrawETH(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _BuyBot.contract.Transact(opts, "withdrawETH", amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xf14210a6.
//
// Solidity: function withdrawETH(uint256 amount) returns()
func (_BuyBot *BuyBotSession) WithdrawETH(amount *big.Int) (*types.Transaction, error) {
	return _BuyBot.Contract.WithdrawETH(&_BuyBot.TransactOpts, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xf14210a6.
//
// Solidity: function withdrawETH(uint256 amount) returns()
func (_BuyBot *BuyBotTransactorSession) WithdrawETH(amount *big.Int) (*types.Transaction, error) {
	return _BuyBot.Contract.WithdrawETH(&_BuyBot.TransactOpts, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BuyBot *BuyBotTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BuyBot.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BuyBot *BuyBotSession) Receive() (*types.Transaction, error) {
	return _BuyBot.Contract.Receive(&_BuyBot.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BuyBot *BuyBotTransactorSession) Receive() (*types.Transaction, error) {
	return _BuyBot.Contract.Receive(&_BuyBot.TransactOpts)
}

// BuyBotIntervalSetIterator is returned from FilterIntervalSet and is used to iterate over the raw logs and unpacked data for IntervalSet events raised by the BuyBot contract.
type BuyBotIntervalSetIterator struct {
	Event *BuyBotIntervalSet // Event containing the contract specifics and raw log

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
func (it *BuyBotIntervalSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuyBotIntervalSet)
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
		it.Event = new(BuyBotIntervalSet)
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
func (it *BuyBotIntervalSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuyBotIntervalSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuyBotIntervalSet represents a IntervalSet event raised by the BuyBot contract.
type BuyBotIntervalSet struct {
	Before  *big.Int
	Current *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterIntervalSet is a free log retrieval operation binding the contract event 0x3fb677206c0b314c404bae3da94bee6bda11375c9fe986f266fb033d6cddbbd7.
//
// Solidity: event IntervalSet(uint256 indexed before, uint256 indexed current)
func (_BuyBot *BuyBotFilterer) FilterIntervalSet(opts *bind.FilterOpts, before []*big.Int, current []*big.Int) (*BuyBotIntervalSetIterator, error) {

	var beforeRule []interface{}
	for _, beforeItem := range before {
		beforeRule = append(beforeRule, beforeItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _BuyBot.contract.FilterLogs(opts, "IntervalSet", beforeRule, currentRule)
	if err != nil {
		return nil, err
	}
	return &BuyBotIntervalSetIterator{contract: _BuyBot.contract, event: "IntervalSet", logs: logs, sub: sub}, nil
}

// WatchIntervalSet is a free log subscription operation binding the contract event 0x3fb677206c0b314c404bae3da94bee6bda11375c9fe986f266fb033d6cddbbd7.
//
// Solidity: event IntervalSet(uint256 indexed before, uint256 indexed current)
func (_BuyBot *BuyBotFilterer) WatchIntervalSet(opts *bind.WatchOpts, sink chan<- *BuyBotIntervalSet, before []*big.Int, current []*big.Int) (event.Subscription, error) {

	var beforeRule []interface{}
	for _, beforeItem := range before {
		beforeRule = append(beforeRule, beforeItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _BuyBot.contract.WatchLogs(opts, "IntervalSet", beforeRule, currentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuyBotIntervalSet)
				if err := _BuyBot.contract.UnpackLog(event, "IntervalSet", log); err != nil {
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

// ParseIntervalSet is a log parse operation binding the contract event 0x3fb677206c0b314c404bae3da94bee6bda11375c9fe986f266fb033d6cddbbd7.
//
// Solidity: event IntervalSet(uint256 indexed before, uint256 indexed current)
func (_BuyBot *BuyBotFilterer) ParseIntervalSet(log types.Log) (*BuyBotIntervalSet, error) {
	event := new(BuyBotIntervalSet)
	if err := _BuyBot.contract.UnpackLog(event, "IntervalSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BuyBotMarketBuyExecutedIterator is returned from FilterMarketBuyExecuted and is used to iterate over the raw logs and unpacked data for MarketBuyExecuted events raised by the BuyBot contract.
type BuyBotMarketBuyExecutedIterator struct {
	Event *BuyBotMarketBuyExecuted // Event containing the contract specifics and raw log

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
func (it *BuyBotMarketBuyExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuyBotMarketBuyExecuted)
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
		it.Event = new(BuyBotMarketBuyExecuted)
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
func (it *BuyBotMarketBuyExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuyBotMarketBuyExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuyBotMarketBuyExecuted represents a MarketBuyExecuted event raised by the BuyBot contract.
type BuyBotMarketBuyExecuted struct {
	Pair        common.Address
	QuoteToken  common.Address
	BaseToken   common.Address
	QuoteAmount *big.Int
	Executor    common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMarketBuyExecuted is a free log retrieval operation binding the contract event 0xc7b4d815bd0a3b8348577971cb55a9334c80a0b3ccdfe287b55b5a6320d8480a.
//
// Solidity: event MarketBuyExecuted(address indexed pair, address indexed quoteToken, address indexed baseToken, uint256 quoteAmount, address executor)
func (_BuyBot *BuyBotFilterer) FilterMarketBuyExecuted(opts *bind.FilterOpts, pair []common.Address, quoteToken []common.Address, baseToken []common.Address) (*BuyBotMarketBuyExecutedIterator, error) {

	var pairRule []interface{}
	for _, pairItem := range pair {
		pairRule = append(pairRule, pairItem)
	}
	var quoteTokenRule []interface{}
	for _, quoteTokenItem := range quoteToken {
		quoteTokenRule = append(quoteTokenRule, quoteTokenItem)
	}
	var baseTokenRule []interface{}
	for _, baseTokenItem := range baseToken {
		baseTokenRule = append(baseTokenRule, baseTokenItem)
	}

	logs, sub, err := _BuyBot.contract.FilterLogs(opts, "MarketBuyExecuted", pairRule, quoteTokenRule, baseTokenRule)
	if err != nil {
		return nil, err
	}
	return &BuyBotMarketBuyExecutedIterator{contract: _BuyBot.contract, event: "MarketBuyExecuted", logs: logs, sub: sub}, nil
}

// WatchMarketBuyExecuted is a free log subscription operation binding the contract event 0xc7b4d815bd0a3b8348577971cb55a9334c80a0b3ccdfe287b55b5a6320d8480a.
//
// Solidity: event MarketBuyExecuted(address indexed pair, address indexed quoteToken, address indexed baseToken, uint256 quoteAmount, address executor)
func (_BuyBot *BuyBotFilterer) WatchMarketBuyExecuted(opts *bind.WatchOpts, sink chan<- *BuyBotMarketBuyExecuted, pair []common.Address, quoteToken []common.Address, baseToken []common.Address) (event.Subscription, error) {

	var pairRule []interface{}
	for _, pairItem := range pair {
		pairRule = append(pairRule, pairItem)
	}
	var quoteTokenRule []interface{}
	for _, quoteTokenItem := range quoteToken {
		quoteTokenRule = append(quoteTokenRule, quoteTokenItem)
	}
	var baseTokenRule []interface{}
	for _, baseTokenItem := range baseToken {
		baseTokenRule = append(baseTokenRule, baseTokenItem)
	}

	logs, sub, err := _BuyBot.contract.WatchLogs(opts, "MarketBuyExecuted", pairRule, quoteTokenRule, baseTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuyBotMarketBuyExecuted)
				if err := _BuyBot.contract.UnpackLog(event, "MarketBuyExecuted", log); err != nil {
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

// ParseMarketBuyExecuted is a log parse operation binding the contract event 0xc7b4d815bd0a3b8348577971cb55a9334c80a0b3ccdfe287b55b5a6320d8480a.
//
// Solidity: event MarketBuyExecuted(address indexed pair, address indexed quoteToken, address indexed baseToken, uint256 quoteAmount, address executor)
func (_BuyBot *BuyBotFilterer) ParseMarketBuyExecuted(log types.Log) (*BuyBotMarketBuyExecuted, error) {
	event := new(BuyBotMarketBuyExecuted)
	if err := _BuyBot.contract.UnpackLog(event, "MarketBuyExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BuyBotMinOrderAmountSetIterator is returned from FilterMinOrderAmountSet and is used to iterate over the raw logs and unpacked data for MinOrderAmountSet events raised by the BuyBot contract.
type BuyBotMinOrderAmountSetIterator struct {
	Event *BuyBotMinOrderAmountSet // Event containing the contract specifics and raw log

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
func (it *BuyBotMinOrderAmountSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuyBotMinOrderAmountSet)
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
		it.Event = new(BuyBotMinOrderAmountSet)
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
func (it *BuyBotMinOrderAmountSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuyBotMinOrderAmountSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuyBotMinOrderAmountSet represents a MinOrderAmountSet event raised by the BuyBot contract.
type BuyBotMinOrderAmountSet struct {
	Before  *big.Int
	Current *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinOrderAmountSet is a free log retrieval operation binding the contract event 0xd6d62b78d21fc4f5151029ca37079020e8a8815c5a488b28ee7499d4c2a19e55.
//
// Solidity: event MinOrderAmountSet(uint256 indexed before, uint256 indexed current)
func (_BuyBot *BuyBotFilterer) FilterMinOrderAmountSet(opts *bind.FilterOpts, before []*big.Int, current []*big.Int) (*BuyBotMinOrderAmountSetIterator, error) {

	var beforeRule []interface{}
	for _, beforeItem := range before {
		beforeRule = append(beforeRule, beforeItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _BuyBot.contract.FilterLogs(opts, "MinOrderAmountSet", beforeRule, currentRule)
	if err != nil {
		return nil, err
	}
	return &BuyBotMinOrderAmountSetIterator{contract: _BuyBot.contract, event: "MinOrderAmountSet", logs: logs, sub: sub}, nil
}

// WatchMinOrderAmountSet is a free log subscription operation binding the contract event 0xd6d62b78d21fc4f5151029ca37079020e8a8815c5a488b28ee7499d4c2a19e55.
//
// Solidity: event MinOrderAmountSet(uint256 indexed before, uint256 indexed current)
func (_BuyBot *BuyBotFilterer) WatchMinOrderAmountSet(opts *bind.WatchOpts, sink chan<- *BuyBotMinOrderAmountSet, before []*big.Int, current []*big.Int) (event.Subscription, error) {

	var beforeRule []interface{}
	for _, beforeItem := range before {
		beforeRule = append(beforeRule, beforeItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _BuyBot.contract.WatchLogs(opts, "MinOrderAmountSet", beforeRule, currentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuyBotMinOrderAmountSet)
				if err := _BuyBot.contract.UnpackLog(event, "MinOrderAmountSet", log); err != nil {
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

// ParseMinOrderAmountSet is a log parse operation binding the contract event 0xd6d62b78d21fc4f5151029ca37079020e8a8815c5a488b28ee7499d4c2a19e55.
//
// Solidity: event MinOrderAmountSet(uint256 indexed before, uint256 indexed current)
func (_BuyBot *BuyBotFilterer) ParseMinOrderAmountSet(log types.Log) (*BuyBotMinOrderAmountSet, error) {
	event := new(BuyBotMinOrderAmountSet)
	if err := _BuyBot.contract.UnpackLog(event, "MinOrderAmountSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BuyBotOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BuyBot contract.
type BuyBotOwnershipTransferredIterator struct {
	Event *BuyBotOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BuyBotOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuyBotOwnershipTransferred)
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
		it.Event = new(BuyBotOwnershipTransferred)
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
func (it *BuyBotOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuyBotOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuyBotOwnershipTransferred represents a OwnershipTransferred event raised by the BuyBot contract.
type BuyBotOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BuyBot *BuyBotFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BuyBotOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BuyBot.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BuyBotOwnershipTransferredIterator{contract: _BuyBot.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BuyBot *BuyBotFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BuyBotOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BuyBot.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuyBotOwnershipTransferred)
				if err := _BuyBot.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BuyBot *BuyBotFilterer) ParseOwnershipTransferred(log types.Log) (*BuyBotOwnershipTransferred, error) {
	event := new(BuyBotOwnershipTransferred)
	if err := _BuyBot.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BuyBotWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the BuyBot contract.
type BuyBotWithdrawnIterator struct {
	Event *BuyBotWithdrawn // Event containing the contract specifics and raw log

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
func (it *BuyBotWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuyBotWithdrawn)
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
		it.Event = new(BuyBotWithdrawn)
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
func (it *BuyBotWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuyBotWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuyBotWithdrawn represents a Withdrawn event raised by the BuyBot contract.
type BuyBotWithdrawn struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed token, address indexed to, uint256 amount)
func (_BuyBot *BuyBotFilterer) FilterWithdrawn(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*BuyBotWithdrawnIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BuyBot.contract.FilterLogs(opts, "Withdrawn", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BuyBotWithdrawnIterator{contract: _BuyBot.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed token, address indexed to, uint256 amount)
func (_BuyBot *BuyBotFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *BuyBotWithdrawn, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BuyBot.contract.WatchLogs(opts, "Withdrawn", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuyBotWithdrawn)
				if err := _BuyBot.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed token, address indexed to, uint256 amount)
func (_BuyBot *BuyBotFilterer) ParseWithdrawn(log types.Log) (*BuyBotWithdrawn, error) {
	event := new(BuyBotWithdrawn)
	if err := _BuyBot.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
