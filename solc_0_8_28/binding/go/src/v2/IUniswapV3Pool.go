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

// IUniswapV3PoolMetaData contains all meta data concerning the IUniswapV3Pool contract.
var IUniswapV3PoolMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"slot0\",\"outputs\":[{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"},{\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"},{\"internalType\":\"uint16\",\"name\":\"observationIndex\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"observationCardinality\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"observationCardinalityNext\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"feeProtocol\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"unlocked\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	ID:  "IUniswapV3Pool",
}

// IUniswapV3Pool is an auto generated Go binding around an Ethereum contract.
type IUniswapV3Pool struct {
	abi abi.ABI
}

// NewIUniswapV3Pool creates a new instance of IUniswapV3Pool.
func NewIUniswapV3Pool() *IUniswapV3Pool {
	parsed, err := IUniswapV3PoolMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &IUniswapV3Pool{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *IUniswapV3Pool) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackSlot0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3850c7bd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function slot0() view returns(uint160 sqrtPriceX96, int24 tick, uint16 observationIndex, uint16 observationCardinality, uint16 observationCardinalityNext, uint8 feeProtocol, bool unlocked)
func (iUniswapV3Pool *IUniswapV3Pool) PackSlot0() []byte {
	enc, err := iUniswapV3Pool.abi.Pack("slot0")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSlot0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3850c7bd.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function slot0() view returns(uint160 sqrtPriceX96, int24 tick, uint16 observationIndex, uint16 observationCardinality, uint16 observationCardinalityNext, uint8 feeProtocol, bool unlocked)
func (iUniswapV3Pool *IUniswapV3Pool) TryPackSlot0() ([]byte, error) {
	return iUniswapV3Pool.abi.Pack("slot0")
}

// Slot0Output serves as a container for the return parameters of contract
// method Slot0.
type Slot0Output struct {
	SqrtPriceX96               *big.Int
	Tick                       *big.Int
	ObservationIndex           uint16
	ObservationCardinality     uint16
	ObservationCardinalityNext uint16
	FeeProtocol                uint8
	Unlocked                   bool
}

// UnpackSlot0 is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3850c7bd.
//
// Solidity: function slot0() view returns(uint160 sqrtPriceX96, int24 tick, uint16 observationIndex, uint16 observationCardinality, uint16 observationCardinalityNext, uint8 feeProtocol, bool unlocked)
func (iUniswapV3Pool *IUniswapV3Pool) UnpackSlot0(data []byte) (Slot0Output, error) {
	out, err := iUniswapV3Pool.abi.Unpack("slot0", data)
	outstruct := new(Slot0Output)
	if err != nil {
		return *outstruct, err
	}
	outstruct.SqrtPriceX96 = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.Tick = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	outstruct.ObservationIndex = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.ObservationCardinality = *abi.ConvertType(out[3], new(uint16)).(*uint16)
	outstruct.ObservationCardinalityNext = *abi.ConvertType(out[4], new(uint16)).(*uint16)
	outstruct.FeeProtocol = *abi.ConvertType(out[5], new(uint8)).(*uint8)
	outstruct.Unlocked = *abi.ConvertType(out[6], new(bool)).(*bool)
	return *outstruct, nil
}

// PackToken0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0dfe1681.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function token0() view returns(address)
func (iUniswapV3Pool *IUniswapV3Pool) PackToken0() []byte {
	enc, err := iUniswapV3Pool.abi.Pack("token0")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackToken0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0dfe1681.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function token0() view returns(address)
func (iUniswapV3Pool *IUniswapV3Pool) TryPackToken0() ([]byte, error) {
	return iUniswapV3Pool.abi.Pack("token0")
}

// UnpackToken0 is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (iUniswapV3Pool *IUniswapV3Pool) UnpackToken0(data []byte) (common.Address, error) {
	out, err := iUniswapV3Pool.abi.Unpack("token0", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackToken1 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd21220a7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function token1() view returns(address)
func (iUniswapV3Pool *IUniswapV3Pool) PackToken1() []byte {
	enc, err := iUniswapV3Pool.abi.Pack("token1")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackToken1 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd21220a7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function token1() view returns(address)
func (iUniswapV3Pool *IUniswapV3Pool) TryPackToken1() ([]byte, error) {
	return iUniswapV3Pool.abi.Pack("token1")
}

// UnpackToken1 is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (iUniswapV3Pool *IUniswapV3Pool) UnpackToken1(data []byte) (common.Address, error) {
	out, err := iUniswapV3Pool.abi.Unpack("token1", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}
