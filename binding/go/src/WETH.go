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
	Bin: "0x60a060405234801561000f575f5ffd5b506040518060400160405280601681526020017f43726f737344455820577261707065642043524f5353000000000000000000008152506040518060400160405280600581526020017f43524f5353000000000000000000000000000000000000000000000000000000815250816003908161008b9190610328565b50806004908161009b9190610328565b5050506100ac6100e460201b60201c565b73ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff16815250506103f7565b5f33905090565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061016657607f821691505b60208210810361017957610178610122565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026101db7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826101a0565b6101e586836101a0565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61022961022461021f846101fd565b610206565b6101fd565b9050919050565b5f819050919050565b6102428361020f565b61025661024e82610230565b8484546101ac565b825550505050565b5f5f905090565b61026d61025e565b610278818484610239565b505050565b5b8181101561029b576102905f82610265565b60018101905061027e565b5050565b601f8211156102e0576102b18161017f565b6102ba84610191565b810160208510156102c9578190505b6102dd6102d585610191565b83018261027d565b50505b505050565b5f82821c905092915050565b5f6103005f19846008026102e5565b1980831691505092915050565b5f61031883836102f1565b9150826002028217905092915050565b610331826100eb565b67ffffffffffffffff81111561034a576103496100f5565b5b610354825461014f565b61035f82828561029f565b5f60209050601f831160018114610390575f841561037e578287015190505b610388858261030d565b8655506103ef565b601f19841661039e8661017f565b5f5b828110156103c5578489015182556001820191506020850194506020810190506103a0565b868310156103e257848901516103de601f8916826102f1565b8355505b6001600288020188555050505b505050505050565b6080516112fd6104165f395f818161044e015261063601526112fd5ff3fe60806040526004361061009f575f3560e01c806332fe7b261161006357806332fe7b26146101b157806370a08231146101db578063755edd171461021757806395d89b4114610233578063a9059cbb1461025d578063dd62ed3e14610299576100b7565b806306fdde03146100bb578063095ea7b3146100e557806318160ddd1461012157806323b872dd1461014b578063313ce56714610187576100b7565b366100b7576100b56100af6102d5565b346102dc565b005b5f5ffd5b3480156100c6575f5ffd5b506100cf61035b565b6040516100dc9190610e80565b60405180910390f35b3480156100f0575f5ffd5b5061010b60048036038101906101069190610f31565b6103eb565b6040516101189190610f89565b60405180910390f35b34801561012c575f5ffd5b5061013561040d565b6040516101429190610fb1565b60405180910390f35b348015610156575f5ffd5b50610171600480360381019061016c9190610fca565b610416565b60405161017e9190610f89565b60405180910390f35b348015610192575f5ffd5b5061019b610444565b6040516101a89190611035565b60405180910390f35b3480156101bc575f5ffd5b506101c561044c565b6040516101d2919061106e565b60405180910390f35b3480156101e6575f5ffd5b5061020160048036038101906101fc9190611087565b610470565b60405161020e9190610fb1565b60405180910390f35b610231600480360381019061022c9190611087565b6104b5565b005b34801561023e575f5ffd5b506102476104c2565b6040516102549190610e80565b60405180910390f35b348015610268575f5ffd5b50610283600480360381019061027e9190610f31565b610552565b6040516102909190610f89565b60405180910390f35b3480156102a4575f5ffd5b506102bf60048036038101906102ba91906110b2565b610574565b6040516102cc9190610fb1565b60405180910390f35b5f33905090565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361034c575f6040517fec442f0500000000000000000000000000000000000000000000000000000000815260040161034391906110ff565b60405180910390fd5b6103575f83836105f6565b5050565b60606003805461036a90611145565b80601f016020809104026020016040519081016040528092919081815260200182805461039690611145565b80156103e15780601f106103b8576101008083540402835291602001916103e1565b820191905f5260205f20905b8154815290600101906020018083116103c457829003601f168201915b5050505050905090565b5f5f6103f56102d5565b905061040281858561070a565b600191505092915050565b5f600254905090565b5f5f6104206102d5565b905061042d85828561071c565b6104388585856107af565b60019150509392505050565b5f6012905090565b7f000000000000000000000000000000000000000000000000000000000000000081565b5f5f5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b6104bf81346102dc565b50565b6060600480546104d190611145565b80601f01602080910402602001604051908101604052809291908181526020018280546104fd90611145565b80156105485780601f1061051f57610100808354040283529160200191610548565b820191905f5260205f20905b81548152906001019060200180831161052b57829003601f168201915b5050505050905090565b5f5f61055c6102d5565b90506105698185856107af565b600191505092915050565b5f60015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905092915050565b61060183838361089f565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614610705577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663e5e31b13836040518263ffffffff1660e01b815260040161068d91906110ff565b602060405180830381865afa1580156106a8573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106cc919061119f565b610704576106da8282610ab8565b610703818373ffffffffffffffffffffffffffffffffffffffff16610b3790919063ffffffff16565b5b5b505050565b6107178383836001610bfd565b505050565b5f6107278484610574565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8110156107a9578181101561079a578281836040517ffb8f41b2000000000000000000000000000000000000000000000000000000008152600401610791939291906111ca565b60405180910390fd5b6107a884848484035f610bfd565b5b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff160361081f575f6040517f96c6fd1e00000000000000000000000000000000000000000000000000000000815260040161081691906110ff565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361088f575f6040517fec442f0500000000000000000000000000000000000000000000000000000000815260040161088691906110ff565b60405180910390fd5b61089a8383836105f6565b505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036108ef578060025f8282546108e3919061122c565b925050819055506109bd565b5f5f5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905081811015610978578381836040517fe450d38c00000000000000000000000000000000000000000000000000000000815260040161096f939291906111ca565b60405180910390fd5b8181035f5f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550505b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610a04578060025f8282540392505081905550610a4e565b805f5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055505b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610aab9190610fb1565b60405180910390a3505050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610b28575f6040517f96c6fd1e000000000000000000000000000000000000000000000000000000008152600401610b1f91906110ff565b60405180910390fd5b610b33825f836105f6565b5050565b80471015610b7e5747816040517fcf479181000000000000000000000000000000000000000000000000000000008152600401610b7592919061125f565b60405180910390fd5b5f5f8373ffffffffffffffffffffffffffffffffffffffff1683604051610ba4906112b3565b5f6040518083038185875af1925050503d805f8114610bde576040519150601f19603f3d011682016040523d82523d5f602084013e610be3565b606091505b509150915081610bf757610bf681610dcc565b5b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603610c6d575f6040517fe602df05000000000000000000000000000000000000000000000000000000008152600401610c6491906110ff565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610cdd575f6040517f94280d62000000000000000000000000000000000000000000000000000000008152600401610cd491906110ff565b60405180910390fd5b8160015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508015610dc6578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051610dbd9190610fb1565b60405180910390a35b50505050565b5f81511115610dde5780518082602001fd5b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610e5282610e10565b610e5c8185610e1a565b9350610e6c818560208601610e2a565b610e7581610e38565b840191505092915050565b5f6020820190508181035f830152610e988184610e48565b905092915050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610ecd82610ea4565b9050919050565b610edd81610ec3565b8114610ee7575f5ffd5b50565b5f81359050610ef881610ed4565b92915050565b5f819050919050565b610f1081610efe565b8114610f1a575f5ffd5b50565b5f81359050610f2b81610f07565b92915050565b5f5f60408385031215610f4757610f46610ea0565b5b5f610f5485828601610eea565b9250506020610f6585828601610f1d565b9150509250929050565b5f8115159050919050565b610f8381610f6f565b82525050565b5f602082019050610f9c5f830184610f7a565b92915050565b610fab81610efe565b82525050565b5f602082019050610fc45f830184610fa2565b92915050565b5f5f5f60608486031215610fe157610fe0610ea0565b5b5f610fee86828701610eea565b9350506020610fff86828701610eea565b925050604061101086828701610f1d565b9150509250925092565b5f60ff82169050919050565b61102f8161101a565b82525050565b5f6020820190506110485f830184611026565b92915050565b5f61105882610ea4565b9050919050565b6110688161104e565b82525050565b5f6020820190506110815f83018461105f565b92915050565b5f6020828403121561109c5761109b610ea0565b5b5f6110a984828501610eea565b91505092915050565b5f5f604083850312156110c8576110c7610ea0565b5b5f6110d585828601610eea565b92505060206110e685828601610eea565b9150509250929050565b6110f981610ec3565b82525050565b5f6020820190506111125f8301846110f0565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061115c57607f821691505b60208210810361116f5761116e611118565b5b50919050565b61117e81610f6f565b8114611188575f5ffd5b50565b5f8151905061119981611175565b92915050565b5f602082840312156111b4576111b3610ea0565b5b5f6111c18482850161118b565b91505092915050565b5f6060820190506111dd5f8301866110f0565b6111ea6020830185610fa2565b6111f76040830184610fa2565b949350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61123682610efe565b915061124183610efe565b9250828201905080821115611259576112586111ff565b5b92915050565b5f6040820190506112725f830185610fa2565b61127f6020830184610fa2565b9392505050565b5f81905092915050565b50565b5f61129e5f83611286565b91506112a982611290565b5f82019050919050565b5f6112bd82611293565b915081905091905056fea264697066735822122091fc5e36755cb2b707be57421ac9754a5efb1ee255f271a964e65d7efe91082264736f6c634300081c0033",
}

// WETHABI is the input ABI used to generate the binding from.
// Deprecated: Use WETHMetaData.ABI instead.
var WETHABI = WETHMetaData.ABI

// WETHBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const WETHBinRuntime = "60806040526004361061009f575f3560e01c806332fe7b261161006357806332fe7b26146101b157806370a08231146101db578063755edd171461021757806395d89b4114610233578063a9059cbb1461025d578063dd62ed3e14610299576100b7565b806306fdde03146100bb578063095ea7b3146100e557806318160ddd1461012157806323b872dd1461014b578063313ce56714610187576100b7565b366100b7576100b56100af6102d5565b346102dc565b005b5f5ffd5b3480156100c6575f5ffd5b506100cf61035b565b6040516100dc9190610e80565b60405180910390f35b3480156100f0575f5ffd5b5061010b60048036038101906101069190610f31565b6103eb565b6040516101189190610f89565b60405180910390f35b34801561012c575f5ffd5b5061013561040d565b6040516101429190610fb1565b60405180910390f35b348015610156575f5ffd5b50610171600480360381019061016c9190610fca565b610416565b60405161017e9190610f89565b60405180910390f35b348015610192575f5ffd5b5061019b610444565b6040516101a89190611035565b60405180910390f35b3480156101bc575f5ffd5b506101c561044c565b6040516101d2919061106e565b60405180910390f35b3480156101e6575f5ffd5b5061020160048036038101906101fc9190611087565b610470565b60405161020e9190610fb1565b60405180910390f35b610231600480360381019061022c9190611087565b6104b5565b005b34801561023e575f5ffd5b506102476104c2565b6040516102549190610e80565b60405180910390f35b348015610268575f5ffd5b50610283600480360381019061027e9190610f31565b610552565b6040516102909190610f89565b60405180910390f35b3480156102a4575f5ffd5b506102bf60048036038101906102ba91906110b2565b610574565b6040516102cc9190610fb1565b60405180910390f35b5f33905090565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361034c575f6040517fec442f0500000000000000000000000000000000000000000000000000000000815260040161034391906110ff565b60405180910390fd5b6103575f83836105f6565b5050565b60606003805461036a90611145565b80601f016020809104026020016040519081016040528092919081815260200182805461039690611145565b80156103e15780601f106103b8576101008083540402835291602001916103e1565b820191905f5260205f20905b8154815290600101906020018083116103c457829003601f168201915b5050505050905090565b5f5f6103f56102d5565b905061040281858561070a565b600191505092915050565b5f600254905090565b5f5f6104206102d5565b905061042d85828561071c565b6104388585856107af565b60019150509392505050565b5f6012905090565b7f000000000000000000000000000000000000000000000000000000000000000081565b5f5f5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b6104bf81346102dc565b50565b6060600480546104d190611145565b80601f01602080910402602001604051908101604052809291908181526020018280546104fd90611145565b80156105485780601f1061051f57610100808354040283529160200191610548565b820191905f5260205f20905b81548152906001019060200180831161052b57829003601f168201915b5050505050905090565b5f5f61055c6102d5565b90506105698185856107af565b600191505092915050565b5f60015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905092915050565b61060183838361089f565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614610705577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663e5e31b13836040518263ffffffff1660e01b815260040161068d91906110ff565b602060405180830381865afa1580156106a8573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106cc919061119f565b610704576106da8282610ab8565b610703818373ffffffffffffffffffffffffffffffffffffffff16610b3790919063ffffffff16565b5b5b505050565b6107178383836001610bfd565b505050565b5f6107278484610574565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8110156107a9578181101561079a578281836040517ffb8f41b2000000000000000000000000000000000000000000000000000000008152600401610791939291906111ca565b60405180910390fd5b6107a884848484035f610bfd565b5b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff160361081f575f6040517f96c6fd1e00000000000000000000000000000000000000000000000000000000815260040161081691906110ff565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361088f575f6040517fec442f0500000000000000000000000000000000000000000000000000000000815260040161088691906110ff565b60405180910390fd5b61089a8383836105f6565b505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036108ef578060025f8282546108e3919061122c565b925050819055506109bd565b5f5f5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905081811015610978578381836040517fe450d38c00000000000000000000000000000000000000000000000000000000815260040161096f939291906111ca565b60405180910390fd5b8181035f5f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550505b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610a04578060025f8282540392505081905550610a4e565b805f5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055505b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610aab9190610fb1565b60405180910390a3505050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610b28575f6040517f96c6fd1e000000000000000000000000000000000000000000000000000000008152600401610b1f91906110ff565b60405180910390fd5b610b33825f836105f6565b5050565b80471015610b7e5747816040517fcf479181000000000000000000000000000000000000000000000000000000008152600401610b7592919061125f565b60405180910390fd5b5f5f8373ffffffffffffffffffffffffffffffffffffffff1683604051610ba4906112b3565b5f6040518083038185875af1925050503d805f8114610bde576040519150601f19603f3d011682016040523d82523d5f602084013e610be3565b606091505b509150915081610bf757610bf681610dcc565b5b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603610c6d575f6040517fe602df05000000000000000000000000000000000000000000000000000000008152600401610c6491906110ff565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610cdd575f6040517f94280d62000000000000000000000000000000000000000000000000000000008152600401610cd491906110ff565b60405180910390fd5b8160015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508015610dc6578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051610dbd9190610fb1565b60405180910390a35b50505050565b5f81511115610dde5780518082602001fd5b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610e5282610e10565b610e5c8185610e1a565b9350610e6c818560208601610e2a565b610e7581610e38565b840191505092915050565b5f6020820190508181035f830152610e988184610e48565b905092915050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610ecd82610ea4565b9050919050565b610edd81610ec3565b8114610ee7575f5ffd5b50565b5f81359050610ef881610ed4565b92915050565b5f819050919050565b610f1081610efe565b8114610f1a575f5ffd5b50565b5f81359050610f2b81610f07565b92915050565b5f5f60408385031215610f4757610f46610ea0565b5b5f610f5485828601610eea565b9250506020610f6585828601610f1d565b9150509250929050565b5f8115159050919050565b610f8381610f6f565b82525050565b5f602082019050610f9c5f830184610f7a565b92915050565b610fab81610efe565b82525050565b5f602082019050610fc45f830184610fa2565b92915050565b5f5f5f60608486031215610fe157610fe0610ea0565b5b5f610fee86828701610eea565b9350506020610fff86828701610eea565b925050604061101086828701610f1d565b9150509250925092565b5f60ff82169050919050565b61102f8161101a565b82525050565b5f6020820190506110485f830184611026565b92915050565b5f61105882610ea4565b9050919050565b6110688161104e565b82525050565b5f6020820190506110815f83018461105f565b92915050565b5f6020828403121561109c5761109b610ea0565b5b5f6110a984828501610eea565b91505092915050565b5f5f604083850312156110c8576110c7610ea0565b5b5f6110d585828601610eea565b92505060206110e685828601610eea565b9150509250929050565b6110f981610ec3565b82525050565b5f6020820190506111125f8301846110f0565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061115c57607f821691505b60208210810361116f5761116e611118565b5b50919050565b61117e81610f6f565b8114611188575f5ffd5b50565b5f8151905061119981611175565b92915050565b5f602082840312156111b4576111b3610ea0565b5b5f6111c18482850161118b565b91505092915050565b5f6060820190506111dd5f8301866110f0565b6111ea6020830185610fa2565b6111f76040830184610fa2565b949350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61123682610efe565b915061124183610efe565b9250828201905080821115611259576112586111ff565b5b92915050565b5f6040820190506112725f830185610fa2565b61127f6020830184610fa2565b9392505050565b5f81905092915050565b50565b5f61129e5f83611286565b91506112a982611290565b5f82019050919050565b5f6112bd82611293565b915081905091905056fea264697066735822122091fc5e36755cb2b707be57421ac9754a5efb1ee255f271a964e65d7efe91082264736f6c634300081c0033"

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
