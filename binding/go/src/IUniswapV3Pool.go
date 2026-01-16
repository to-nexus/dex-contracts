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

// IUniswapV3PoolMetaData contains all meta data concerning the IUniswapV3Pool contract.
var IUniswapV3PoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"slot0\",\"outputs\":[{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"},{\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"},{\"internalType\":\"uint16\",\"name\":\"observationIndex\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"observationCardinality\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"observationCardinalityNext\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"feeProtocol\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"unlocked\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"3850c7bd": "slot0()",
		"0dfe1681": "token0()",
		"d21220a7": "token1()",
	},
}

// IUniswapV3PoolABI is the input ABI used to generate the binding from.
// Deprecated: Use IUniswapV3PoolMetaData.ABI instead.
var IUniswapV3PoolABI = IUniswapV3PoolMetaData.ABI

// Deprecated: Use IUniswapV3PoolMetaData.Sigs instead.
// IUniswapV3PoolFuncSigs maps the 4-byte function signature to its string representation.
var IUniswapV3PoolFuncSigs = IUniswapV3PoolMetaData.Sigs

// IUniswapV3Pool is an auto generated Go binding around an Ethereum contract.
type IUniswapV3Pool struct {
	IUniswapV3PoolCaller     // Read-only binding to the contract
	IUniswapV3PoolTransactor // Write-only binding to the contract
	IUniswapV3PoolFilterer   // Log filterer for contract events
}

// IUniswapV3PoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type IUniswapV3PoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV3PoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IUniswapV3PoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV3PoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IUniswapV3PoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV3PoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IUniswapV3PoolSession struct {
	Contract     *IUniswapV3Pool   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IUniswapV3PoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IUniswapV3PoolCallerSession struct {
	Contract *IUniswapV3PoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IUniswapV3PoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IUniswapV3PoolTransactorSession struct {
	Contract     *IUniswapV3PoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IUniswapV3PoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type IUniswapV3PoolRaw struct {
	Contract *IUniswapV3Pool // Generic contract binding to access the raw methods on
}

// IUniswapV3PoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IUniswapV3PoolCallerRaw struct {
	Contract *IUniswapV3PoolCaller // Generic read-only contract binding to access the raw methods on
}

// IUniswapV3PoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IUniswapV3PoolTransactorRaw struct {
	Contract *IUniswapV3PoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIUniswapV3Pool creates a new instance of IUniswapV3Pool, bound to a specific deployed contract.
func NewIUniswapV3Pool(address common.Address, backend bind.ContractBackend) (*IUniswapV3Pool, error) {
	contract, err := bindIUniswapV3Pool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3Pool{IUniswapV3PoolCaller: IUniswapV3PoolCaller{contract: contract}, IUniswapV3PoolTransactor: IUniswapV3PoolTransactor{contract: contract}, IUniswapV3PoolFilterer: IUniswapV3PoolFilterer{contract: contract}}, nil
}

// NewIUniswapV3PoolCaller creates a new read-only instance of IUniswapV3Pool, bound to a specific deployed contract.
func NewIUniswapV3PoolCaller(address common.Address, caller bind.ContractCaller) (*IUniswapV3PoolCaller, error) {
	contract, err := bindIUniswapV3Pool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolCaller{contract: contract}, nil
}

// NewIUniswapV3PoolTransactor creates a new write-only instance of IUniswapV3Pool, bound to a specific deployed contract.
func NewIUniswapV3PoolTransactor(address common.Address, transactor bind.ContractTransactor) (*IUniswapV3PoolTransactor, error) {
	contract, err := bindIUniswapV3Pool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolTransactor{contract: contract}, nil
}

// NewIUniswapV3PoolFilterer creates a new log filterer instance of IUniswapV3Pool, bound to a specific deployed contract.
func NewIUniswapV3PoolFilterer(address common.Address, filterer bind.ContractFilterer) (*IUniswapV3PoolFilterer, error) {
	contract, err := bindIUniswapV3Pool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolFilterer{contract: contract}, nil
}

// bindIUniswapV3Pool binds a generic wrapper to an already deployed contract.
func bindIUniswapV3Pool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IUniswapV3PoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUniswapV3Pool *IUniswapV3PoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUniswapV3Pool.Contract.IUniswapV3PoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUniswapV3Pool *IUniswapV3PoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniswapV3Pool.Contract.IUniswapV3PoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUniswapV3Pool *IUniswapV3PoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUniswapV3Pool.Contract.IUniswapV3PoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUniswapV3Pool *IUniswapV3PoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUniswapV3Pool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUniswapV3Pool *IUniswapV3PoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniswapV3Pool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUniswapV3Pool *IUniswapV3PoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUniswapV3Pool.Contract.contract.Transact(opts, method, params...)
}

// Slot0 is a free data retrieval call binding the contract method 0x3850c7bd.
//
// Solidity: function slot0() view returns(uint160 sqrtPriceX96, int24 tick, uint16 observationIndex, uint16 observationCardinality, uint16 observationCardinalityNext, uint8 feeProtocol, bool unlocked)
func (_IUniswapV3Pool *IUniswapV3PoolCaller) Slot0(opts *bind.CallOpts) (struct {
	SqrtPriceX96               *big.Int
	Tick                       *big.Int
	ObservationIndex           uint16
	ObservationCardinality     uint16
	ObservationCardinalityNext uint16
	FeeProtocol                uint8
	Unlocked                   bool
}, error) {
	var out []interface{}
	err := _IUniswapV3Pool.contract.Call(opts, &out, "slot0")

	outstruct := new(struct {
		SqrtPriceX96               *big.Int
		Tick                       *big.Int
		ObservationIndex           uint16
		ObservationCardinality     uint16
		ObservationCardinalityNext uint16
		FeeProtocol                uint8
		Unlocked                   bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SqrtPriceX96 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Tick = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ObservationIndex = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.ObservationCardinality = *abi.ConvertType(out[3], new(uint16)).(*uint16)
	outstruct.ObservationCardinalityNext = *abi.ConvertType(out[4], new(uint16)).(*uint16)
	outstruct.FeeProtocol = *abi.ConvertType(out[5], new(uint8)).(*uint8)
	outstruct.Unlocked = *abi.ConvertType(out[6], new(bool)).(*bool)

	return *outstruct, err

}

// Slot0 is a free data retrieval call binding the contract method 0x3850c7bd.
//
// Solidity: function slot0() view returns(uint160 sqrtPriceX96, int24 tick, uint16 observationIndex, uint16 observationCardinality, uint16 observationCardinalityNext, uint8 feeProtocol, bool unlocked)
func (_IUniswapV3Pool *IUniswapV3PoolSession) Slot0() (struct {
	SqrtPriceX96               *big.Int
	Tick                       *big.Int
	ObservationIndex           uint16
	ObservationCardinality     uint16
	ObservationCardinalityNext uint16
	FeeProtocol                uint8
	Unlocked                   bool
}, error) {
	return _IUniswapV3Pool.Contract.Slot0(&_IUniswapV3Pool.CallOpts)
}

// Slot0 is a free data retrieval call binding the contract method 0x3850c7bd.
//
// Solidity: function slot0() view returns(uint160 sqrtPriceX96, int24 tick, uint16 observationIndex, uint16 observationCardinality, uint16 observationCardinalityNext, uint8 feeProtocol, bool unlocked)
func (_IUniswapV3Pool *IUniswapV3PoolCallerSession) Slot0() (struct {
	SqrtPriceX96               *big.Int
	Tick                       *big.Int
	ObservationIndex           uint16
	ObservationCardinality     uint16
	ObservationCardinalityNext uint16
	FeeProtocol                uint8
	Unlocked                   bool
}, error) {
	return _IUniswapV3Pool.Contract.Slot0(&_IUniswapV3Pool.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_IUniswapV3Pool *IUniswapV3PoolCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IUniswapV3Pool.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_IUniswapV3Pool *IUniswapV3PoolSession) Token0() (common.Address, error) {
	return _IUniswapV3Pool.Contract.Token0(&_IUniswapV3Pool.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_IUniswapV3Pool *IUniswapV3PoolCallerSession) Token0() (common.Address, error) {
	return _IUniswapV3Pool.Contract.Token0(&_IUniswapV3Pool.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_IUniswapV3Pool *IUniswapV3PoolCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IUniswapV3Pool.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_IUniswapV3Pool *IUniswapV3PoolSession) Token1() (common.Address, error) {
	return _IUniswapV3Pool.Contract.Token1(&_IUniswapV3Pool.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_IUniswapV3Pool *IUniswapV3PoolCallerSession) Token1() (common.Address, error) {
	return _IUniswapV3Pool.Contract.Token1(&_IUniswapV3Pool.CallOpts)
}
