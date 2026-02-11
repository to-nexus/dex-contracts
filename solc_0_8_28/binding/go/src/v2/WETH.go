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

// WETHMetaData contains all meta data concerning the WETH contract.
var WETHMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"ROUTER\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mintTo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"}]",
	ID:  "WETH",
	Bin: "0x60a060405234801561000f575f5ffd5b506040518060400160405280601681526020017f43726f737344455820577261707065642043524f5353000000000000000000008152506040518060400160405280600581526020016443524f535360d81b81525081600390816100739190610134565b5060046100808282610134565b5061008b9150503390565b6001600160a01b03166080526101ee565b634e487b7160e01b5f52604160045260245ffd5b600181811c908216806100c457607f821691505b6020821081036100e257634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561012f57805f5260205f20601f840160051c8101602085101561010d5750805b601f840160051c820191505b8181101561012c575f8155600101610119565b50505b505050565b81516001600160401b0381111561014d5761014d61009c565b6101618161015b84546100b0565b846100e8565b6020601f821160018114610193575f831561017c5750848201515b5f19600385901b1c1916600184901b17845561012c565b5f84815260208120601f198516915b828110156101c257878501518255602094850194600190920191016101a2565b50848210156101df57868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b608051610cdb61020d5f395f818161019201526104750152610cdb5ff3fe6080604052600436106100bb575f3560e01c806332fe7b261161007157806395d89b411161004c57806395d89b411461022d578063a9059cbb14610241578063dd62ed3e14610260575f5ffd5b806332fe7b261461018157806370a08231146101d9578063755edd171461021a575f5ffd5b806318160ddd116100a157806318160ddd1461012957806323b872dd14610147578063313ce56714610166575f5ffd5b806306fdde03146100d0578063095ea7b3146100fa575f5ffd5b366100cc576100ca33346102b1565b005b5f5ffd5b3480156100db575f5ffd5b506100e4610314565b6040516100f19190610acf565b60405180910390f35b348015610105575f5ffd5b50610119610114366004610b4a565b6103a4565b60405190151581526020016100f1565b348015610134575f5ffd5b506002545b6040519081526020016100f1565b348015610152575f5ffd5b50610119610161366004610b72565b6103bd565b348015610171575f5ffd5b50604051601281526020016100f1565b34801561018c575f5ffd5b506101b47f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100f1565b3480156101e4575f5ffd5b506101396101f3366004610bac565b73ffffffffffffffffffffffffffffffffffffffff165f9081526020819052604090205490565b6100ca610228366004610bac565b6103e0565b348015610238575f5ffd5b506100e46103ed565b34801561024c575f5ffd5b5061011961025b366004610b4a565b6103fc565b34801561026b575f5ffd5b5061013961027a366004610bcc565b73ffffffffffffffffffffffffffffffffffffffff9182165f90815260016020908152604080832093909416825291909152205490565b73ffffffffffffffffffffffffffffffffffffffff8216610305576040517fec442f050000000000000000000000000000000000000000000000000000000081525f60048201526024015b60405180910390fd5b6103105f8383610409565b5050565b60606003805461032390610bfd565b80601f016020809104026020016040519081016040528092919081815260200182805461034f90610bfd565b801561039a5780601f106103715761010080835404028352916020019161039a565b820191905f5260205f20905b81548152906001019060200180831161037d57829003601f168201915b5050505050905090565b5f336103b1818585610511565b60019150505b92915050565b5f336103ca85828561051e565b6103d58585856105ec565b506001949350505050565b6103ea81346102b1565b50565b60606004805461032390610bfd565b5f336103b18185856105ec565b610414838383610695565b73ffffffffffffffffffffffffffffffffffffffff82161561050c576040517fe5e31b1300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063e5e31b1390602401602060405180830381865afa1580156104ba573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104de9190610c4e565b61050c576104ec828261083c565b61050c73ffffffffffffffffffffffffffffffffffffffff831682610896565b505050565b61050c8383836001610948565b73ffffffffffffffffffffffffffffffffffffffff8381165f908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8110156105e657818110156105d8576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8416600482015260248101829052604481018390526064016102fc565b6105e684848484035f610948565b50505050565b73ffffffffffffffffffffffffffffffffffffffff831661063b576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f60048201526024016102fc565b73ffffffffffffffffffffffffffffffffffffffff821661068a576040517fec442f050000000000000000000000000000000000000000000000000000000081525f60048201526024016102fc565b61050c838383610409565b73ffffffffffffffffffffffffffffffffffffffff83166106cc578060025f8282546106c19190610c6d565b9091555061077c9050565b73ffffffffffffffffffffffffffffffffffffffff83165f9081526020819052604090205481811015610751576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015260248101829052604481018390526064016102fc565b73ffffffffffffffffffffffffffffffffffffffff84165f9081526020819052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff82166107a5576002805482900390556107d0565b73ffffffffffffffffffffffffffffffffffffffff82165f9081526020819052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161082f91815260200190565b60405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff821661088b576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f60048201526024016102fc565b610310825f83610409565b804710156108d9576040517fcf479181000000000000000000000000000000000000000000000000000000008152476004820152602481018290526044016102fc565b5f5f8373ffffffffffffffffffffffffffffffffffffffff16836040515f6040518083038185875af1925050503d805f8114610930576040519150601f19603f3d011682016040523d82523d5f602084013e610935565b606091505b5091509150816105e6576105e681610a8d565b73ffffffffffffffffffffffffffffffffffffffff8416610997576040517fe602df050000000000000000000000000000000000000000000000000000000081525f60048201526024016102fc565b73ffffffffffffffffffffffffffffffffffffffff83166109e6576040517f94280d620000000000000000000000000000000000000000000000000000000081525f60048201526024016102fc565b73ffffffffffffffffffffffffffffffffffffffff8085165f90815260016020908152604080832093871683529290522082905580156105e6578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051610a7f91815260200190565b60405180910390a350505050565b805115610a9d5780518082602001fd5b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610b45575f5ffd5b919050565b5f5f60408385031215610b5b575f5ffd5b610b6483610b22565b946020939093013593505050565b5f5f5f60608486031215610b84575f5ffd5b610b8d84610b22565b9250610b9b60208501610b22565b929592945050506040919091013590565b5f60208284031215610bbc575f5ffd5b610bc582610b22565b9392505050565b5f5f60408385031215610bdd575f5ffd5b610be683610b22565b9150610bf460208401610b22565b90509250929050565b600181811c90821680610c1157607f821691505b602082108103610c48577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b5f60208284031215610c5e575f5ffd5b81518015158114610bc5575f5ffd5b808201808211156103b7577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffdfea2646970667358221220ff2c94170bfb1a4b9862ab30fb98d16e587528e43a5fcafadc87c86eb6afbc4e64736f6c634300081c0033",
}

// WETH is an auto generated Go binding around an Ethereum contract.
type WETH struct {
	abi abi.ABI
}

// NewWETH creates a new instance of WETH.
func NewWETH() *WETH {
	parsed, err := WETHMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &WETH{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *WETH) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackROUTER is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x32fe7b26.
//
// Solidity: function ROUTER() view returns(address)
func (wETH *WETH) PackROUTER() []byte {
	enc, err := wETH.abi.Pack("ROUTER")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackROUTER is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x32fe7b26.
//
// Solidity: function ROUTER() view returns(address)
func (wETH *WETH) UnpackROUTER(data []byte) (common.Address, error) {
	out, err := wETH.abi.Unpack("ROUTER", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackAllowance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (wETH *WETH) PackAllowance(owner common.Address, spender common.Address) []byte {
	enc, err := wETH.abi.Pack("allowance", owner, spender)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackAllowance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (wETH *WETH) UnpackAllowance(data []byte) (*big.Int, error) {
	out, err := wETH.abi.Unpack("allowance", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackApprove is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (wETH *WETH) PackApprove(spender common.Address, value *big.Int) []byte {
	enc, err := wETH.abi.Pack("approve", spender, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackApprove is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (wETH *WETH) UnpackApprove(data []byte) (bool, error) {
	out, err := wETH.abi.Unpack("approve", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackBalanceOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (wETH *WETH) PackBalanceOf(account common.Address) []byte {
	enc, err := wETH.abi.Pack("balanceOf", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackBalanceOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (wETH *WETH) UnpackBalanceOf(data []byte) (*big.Int, error) {
	out, err := wETH.abi.Unpack("balanceOf", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackDecimals is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (wETH *WETH) PackDecimals() []byte {
	enc, err := wETH.abi.Pack("decimals")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDecimals is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (wETH *WETH) UnpackDecimals(data []byte) (uint8, error) {
	out, err := wETH.abi.Unpack("decimals", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, err
}

// PackMintTo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x755edd17.
//
// Solidity: function mintTo(address to) payable returns()
func (wETH *WETH) PackMintTo(to common.Address) []byte {
	enc, err := wETH.abi.Pack("mintTo", to)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (wETH *WETH) PackName() []byte {
	enc, err := wETH.abi.Pack("name")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackName is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (wETH *WETH) UnpackName(data []byte) (string, error) {
	out, err := wETH.abi.Unpack("name", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackSymbol is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (wETH *WETH) PackSymbol() []byte {
	enc, err := wETH.abi.Pack("symbol")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSymbol is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (wETH *WETH) UnpackSymbol(data []byte) (string, error) {
	out, err := wETH.abi.Unpack("symbol", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackTotalSupply is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (wETH *WETH) PackTotalSupply() []byte {
	enc, err := wETH.abi.Pack("totalSupply")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTotalSupply is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (wETH *WETH) UnpackTotalSupply(data []byte) (*big.Int, error) {
	out, err := wETH.abi.Unpack("totalSupply", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (wETH *WETH) PackTransfer(to common.Address, value *big.Int) []byte {
	enc, err := wETH.abi.Pack("transfer", to, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTransfer is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (wETH *WETH) UnpackTransfer(data []byte) (bool, error) {
	out, err := wETH.abi.Unpack("transfer", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (wETH *WETH) PackTransferFrom(from common.Address, to common.Address, value *big.Int) []byte {
	enc, err := wETH.abi.Pack("transferFrom", from, to, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTransferFrom is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (wETH *WETH) UnpackTransferFrom(data []byte) (bool, error) {
	out, err := wETH.abi.Unpack("transferFrom", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// WETHApproval represents a Approval event raised by the WETH contract.
type WETHApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const WETHApprovalEventName = "Approval"

// ContractEventName returns the user-defined event name.
func (WETHApproval) ContractEventName() string {
	return WETHApprovalEventName
}

// UnpackApprovalEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (wETH *WETH) UnpackApprovalEvent(log *types.Log) (*WETHApproval, error) {
	event := "Approval"
	if log.Topics[0] != wETH.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WETHApproval)
	if len(log.Data) > 0 {
		if err := wETH.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range wETH.abi.Events[event].Inputs {
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

// WETHTransfer represents a Transfer event raised by the WETH contract.
type WETHTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   *types.Log // Blockchain specific contextual infos
}

const WETHTransferEventName = "Transfer"

// ContractEventName returns the user-defined event name.
func (WETHTransfer) ContractEventName() string {
	return WETHTransferEventName
}

// UnpackTransferEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (wETH *WETH) UnpackTransferEvent(log *types.Log) (*WETHTransfer, error) {
	event := "Transfer"
	if log.Topics[0] != wETH.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WETHTransfer)
	if len(log.Data) > 0 {
		if err := wETH.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range wETH.abi.Events[event].Inputs {
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
func (wETH *WETH) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], wETH.abi.Errors["ERC20InsufficientAllowance"].ID.Bytes()[:4]) {
		return wETH.UnpackERC20InsufficientAllowanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], wETH.abi.Errors["ERC20InsufficientBalance"].ID.Bytes()[:4]) {
		return wETH.UnpackERC20InsufficientBalanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], wETH.abi.Errors["ERC20InvalidApprover"].ID.Bytes()[:4]) {
		return wETH.UnpackERC20InvalidApproverError(raw[4:])
	}
	if bytes.Equal(raw[:4], wETH.abi.Errors["ERC20InvalidReceiver"].ID.Bytes()[:4]) {
		return wETH.UnpackERC20InvalidReceiverError(raw[4:])
	}
	if bytes.Equal(raw[:4], wETH.abi.Errors["ERC20InvalidSender"].ID.Bytes()[:4]) {
		return wETH.UnpackERC20InvalidSenderError(raw[4:])
	}
	if bytes.Equal(raw[:4], wETH.abi.Errors["ERC20InvalidSpender"].ID.Bytes()[:4]) {
		return wETH.UnpackERC20InvalidSpenderError(raw[4:])
	}
	if bytes.Equal(raw[:4], wETH.abi.Errors["FailedCall"].ID.Bytes()[:4]) {
		return wETH.UnpackFailedCallError(raw[4:])
	}
	if bytes.Equal(raw[:4], wETH.abi.Errors["InsufficientBalance"].ID.Bytes()[:4]) {
		return wETH.UnpackInsufficientBalanceError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// WETHERC20InsufficientAllowance represents a ERC20InsufficientAllowance error raised by the WETH contract.
type WETHERC20InsufficientAllowance struct {
	Spender   common.Address
	Allowance *big.Int
	Needed    *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InsufficientAllowance(address spender, uint256 allowance, uint256 needed)
func WETHERC20InsufficientAllowanceErrorID() common.Hash {
	return common.HexToHash("0xfb8f41b23e99d2101d86da76cdfa87dd51c82ed07d3cb62cbc473e469dbc75c3")
}

// UnpackERC20InsufficientAllowanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InsufficientAllowance(address spender, uint256 allowance, uint256 needed)
func (wETH *WETH) UnpackERC20InsufficientAllowanceError(raw []byte) (*WETHERC20InsufficientAllowance, error) {
	out := new(WETHERC20InsufficientAllowance)
	if err := wETH.abi.UnpackIntoInterface(out, "ERC20InsufficientAllowance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WETHERC20InsufficientBalance represents a ERC20InsufficientBalance error raised by the WETH contract.
type WETHERC20InsufficientBalance struct {
	Sender  common.Address
	Balance *big.Int
	Needed  *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InsufficientBalance(address sender, uint256 balance, uint256 needed)
func WETHERC20InsufficientBalanceErrorID() common.Hash {
	return common.HexToHash("0xe450d38cd8d9f7d95077d567d60ed49c7254716e6ad08fc9872816c97e0ffec6")
}

// UnpackERC20InsufficientBalanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InsufficientBalance(address sender, uint256 balance, uint256 needed)
func (wETH *WETH) UnpackERC20InsufficientBalanceError(raw []byte) (*WETHERC20InsufficientBalance, error) {
	out := new(WETHERC20InsufficientBalance)
	if err := wETH.abi.UnpackIntoInterface(out, "ERC20InsufficientBalance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WETHERC20InvalidApprover represents a ERC20InvalidApprover error raised by the WETH contract.
type WETHERC20InvalidApprover struct {
	Approver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidApprover(address approver)
func WETHERC20InvalidApproverErrorID() common.Hash {
	return common.HexToHash("0xe602df05cc75712490294c6c104ab7c17f4030363910a7a2626411c6d3118847")
}

// UnpackERC20InvalidApproverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidApprover(address approver)
func (wETH *WETH) UnpackERC20InvalidApproverError(raw []byte) (*WETHERC20InvalidApprover, error) {
	out := new(WETHERC20InvalidApprover)
	if err := wETH.abi.UnpackIntoInterface(out, "ERC20InvalidApprover", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WETHERC20InvalidReceiver represents a ERC20InvalidReceiver error raised by the WETH contract.
type WETHERC20InvalidReceiver struct {
	Receiver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidReceiver(address receiver)
func WETHERC20InvalidReceiverErrorID() common.Hash {
	return common.HexToHash("0xec442f055133b72f3b2f9f0bb351c406b178527de2040a7d1feb4e058771f613")
}

// UnpackERC20InvalidReceiverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidReceiver(address receiver)
func (wETH *WETH) UnpackERC20InvalidReceiverError(raw []byte) (*WETHERC20InvalidReceiver, error) {
	out := new(WETHERC20InvalidReceiver)
	if err := wETH.abi.UnpackIntoInterface(out, "ERC20InvalidReceiver", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WETHERC20InvalidSender represents a ERC20InvalidSender error raised by the WETH contract.
type WETHERC20InvalidSender struct {
	Sender common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidSender(address sender)
func WETHERC20InvalidSenderErrorID() common.Hash {
	return common.HexToHash("0x96c6fd1edd0cd6ef7ff0ecc0facdf53148dc0048b57fe58af65755250a7a96bd")
}

// UnpackERC20InvalidSenderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidSender(address sender)
func (wETH *WETH) UnpackERC20InvalidSenderError(raw []byte) (*WETHERC20InvalidSender, error) {
	out := new(WETHERC20InvalidSender)
	if err := wETH.abi.UnpackIntoInterface(out, "ERC20InvalidSender", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WETHERC20InvalidSpender represents a ERC20InvalidSpender error raised by the WETH contract.
type WETHERC20InvalidSpender struct {
	Spender common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidSpender(address spender)
func WETHERC20InvalidSpenderErrorID() common.Hash {
	return common.HexToHash("0x94280d62c347d8d9f4d59a76ea321452406db88df38e0c9da304f58b57b373a2")
}

// UnpackERC20InvalidSpenderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidSpender(address spender)
func (wETH *WETH) UnpackERC20InvalidSpenderError(raw []byte) (*WETHERC20InvalidSpender, error) {
	out := new(WETHERC20InvalidSpender)
	if err := wETH.abi.UnpackIntoInterface(out, "ERC20InvalidSpender", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WETHFailedCall represents a FailedCall error raised by the WETH contract.
type WETHFailedCall struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedCall()
func WETHFailedCallErrorID() common.Hash {
	return common.HexToHash("0xd6bda27508c0fb6d8a39b4b122878dab26f731a7d4e4abe711dd3731899052a4")
}

// UnpackFailedCallError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedCall()
func (wETH *WETH) UnpackFailedCallError(raw []byte) (*WETHFailedCall, error) {
	out := new(WETHFailedCall)
	if err := wETH.abi.UnpackIntoInterface(out, "FailedCall", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WETHInsufficientBalance represents a InsufficientBalance error raised by the WETH contract.
type WETHInsufficientBalance struct {
	Balance *big.Int
	Needed  *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InsufficientBalance(uint256 balance, uint256 needed)
func WETHInsufficientBalanceErrorID() common.Hash {
	return common.HexToHash("0xcf4791818fba6e019216eb4864093b4947f674afada5d305e57d598b641dad1d")
}

// UnpackInsufficientBalanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InsufficientBalance(uint256 balance, uint256 needed)
func (wETH *WETH) UnpackInsufficientBalanceError(raw []byte) (*WETHInsufficientBalance, error) {
	out := new(WETHInsufficientBalance)
	if err := wETH.abi.UnpackIntoInterface(out, "InsufficientBalance", raw); err != nil {
		return nil, err
	}
	return out, nil
}
