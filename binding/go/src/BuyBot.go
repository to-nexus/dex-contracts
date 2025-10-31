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

// BuyBotMetaData contains all meta data concerning the BuyBot contract.
var BuyBotMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"_initialDelay\",\"type\":\"uint48\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minOrderAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_interval\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"BUYER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NATIVE_COIN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"beginDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxMatchCount\",\"type\":\"uint256\"}],\"name\":\"buyMarket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"canBuyMarket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"canBuy\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"}],\"name\":\"changeDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelayIncreaseWait\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastBuyTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minOrderAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"recipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollbackDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"contractIRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_interval\",\"type\":\"uint256\"}],\"name\":\"setInterval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minOrderAmount\",\"type\":\"uint256\"}],\"name\":\"setMinOrderAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"setRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminDelayChangeCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"effectSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminDelayChangeScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminTransferCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"acceptSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminTransferScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"before\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"current\",\"type\":\"uint256\"}],\"name\":\"IntervalSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"MarketBuyExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"before\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"current\",\"type\":\"uint256\"}],\"name\":\"MinOrderAmountSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"before\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"RecipientSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"name\":\"AccessControlEnforcedDefaultAdminDelay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccessControlEnforcedDefaultAdminRules\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"defaultAdmin\",\"type\":\"address\"}],\"name\":\"AccessControlInvalidDefaultAdmin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minOrderAmount\",\"type\":\"uint256\"}],\"name\":\"BuyBotInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timeSinceLastBuy\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredInterval\",\"type\":\"uint256\"}],\"name\":\"BuyBotIntervalNotPassed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"BuyBotInvalidAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"}],\"name\":\"BuyBotInvalidBuyer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"BuyBotInvalidManager\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"BuyBotInvalidMinOrderAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"BuyBotInvalidPair\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"BuyBotInvalidRouter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"7a01a1da": "BUYER_ROLE()",
		"a217fddf": "DEFAULT_ADMIN_ROLE()",
		"ec87621c": "MANAGER_ROLE()",
		"04a41159": "NATIVE_COIN()",
		"cefc1429": "acceptDefaultAdminTransfer()",
		"634e93da": "beginDefaultAdminTransfer(address)",
		"67ea88eb": "buyMarket(address,uint256,uint256)",
		"2aaa9628": "canBuyMarket(address,address)",
		"d602b9fd": "cancelDefaultAdminTransfer()",
		"649a5ec7": "changeDefaultAdminDelay(uint48)",
		"84ef8ffc": "defaultAdmin()",
		"cc8463c8": "defaultAdminDelay()",
		"022d63fb": "defaultAdminDelayIncreaseWait()",
		"f8b2cb4f": "getBalance(address)",
		"248a9ca3": "getRoleAdmin(bytes32)",
		"2f2ff15d": "grantRole(bytes32,address)",
		"91d14854": "hasRole(bytes32,address)",
		"947a36fb": "interval()",
		"f29f4d0b": "lastBuyTime()",
		"46b62c4a": "minOrderAmount()",
		"8da5cb5b": "owner()",
		"cf6eefb7": "pendingDefaultAdmin()",
		"a1eda53c": "pendingDefaultAdminDelay()",
		"66d003ac": "recipient()",
		"36568abe": "renounceRole(bytes32,address)",
		"d547741f": "revokeRole(bytes32,address)",
		"0aa6220b": "rollbackDefaultAdminDelay()",
		"f887ea40": "router()",
		"22a90082": "setInterval(uint256)",
		"a3b8ef04": "setMinOrderAmount(uint256)",
		"3bbed4a0": "setRecipient(address)",
		"01ffc9a7": "supportsInterface(bytes4)",
		"f3fef3a3": "withdraw(address,uint256)",
		"f14210a6": "withdrawETH(uint256)",
	},
	Bin: "0x608060405234801561000f575f5ffd5b50604051612d01380380612d0183398101604081905261002e91610451565b87876001600160a01b03811661005e57604051636116401160e11b81525f60048201526024015b60405180910390fd5b600180546001600160d01b0316600160d01b65ffffffffffff8516021790556100875f826102aa565b50506001600355506001600160a01b0386166100c157604051631561f27b60e21b81526001600160a01b0387166004820152602401610055565b845f036100e4576040516328bfc81960e11b815260048101869052602401610055565b6001600160a01b03821661011657604051632a7409a960e01b81526001600160a01b0383166004820152602401610055565b6001600160a01b0381166101485760405163192178dd60e11b81526001600160a01b0382166004820152602401610055565b600480546001600160a01b038089166001600160a01b0319928316179092556005879055600686905560088054928616929091169190911790556101995f516020612ce15f395f51905f525f610319565b6101b05f516020612cc15f395f51905f525f610319565b6101c75f516020612ce15f395f51905f52886102aa565b506101df5f516020612ce15f395f51905f52836102aa565b506101f75f516020612cc15f395f51905f52886102aa565b5061020f5f516020612cc15f395f51905f52826102aa565b5060405185905f907fd6d62b78d21fc4f5151029ca37079020e8a8815c5a488b28ee7499d4c2a19e55908290a360405184905f907f3fb677206c0b314c404bae3da94bee6bda11375c9fe986f266fb033d6cddbbd7908290a36040516001600160a01b038416905f907fc1416b5cdab50a9fbc872236e1aa54566c6deb40024e63a4b1737ecacf09d6f9908290a350505050505050506104e2565b5f82610306575f6102c36002546001600160a01b031690565b6001600160a01b0316146102ea57604051631fe1e13d60e11b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b0384161790555b6103108383610345565b90505b92915050565b8161033757604051631fe1e13d60e11b815260040160405180910390fd5b61034182826103ec565b5050565b5f828152602081815260408083206001600160a01b038516845290915281205460ff166103e5575f838152602081815260408083206001600160a01b03861684529091529020805460ff1916600117905561039d3390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4506001610313565b505f610313565b5f82815260208190526040808220600101805490849055905190918391839186917fbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff9190a4505050565b80516001600160a01b038116811461044c575f5ffd5b919050565b5f5f5f5f5f5f5f5f610100898b031215610469575f5ffd5b885165ffffffffffff8116811461047e575f5ffd5b975061048c60208a01610436565b965061049a60408a01610436565b60608a015160808b0151919750955093506104b760a08a01610436565b92506104c560c08a01610436565b91506104d360e08a01610436565b90509295985092959890939650565b6127d2806104ef5f395ff3fe608060405260043610610212575f3560e01c80638da5cb5b11610117578063cf6eefb7116100ac578063f14210a61161007c578063f3fef3a311610062578063f3fef3a3146106a8578063f887ea40146106c7578063f8b2cb4f146106f3575f5ffd5b8063f14210a614610674578063f29f4d0b14610693575f5ffd5b8063cf6eefb7146105b6578063d547741f1461060e578063d602b9fd1461062d578063ec87621c14610641575f5ffd5b8063a217fddf116100e7578063a217fddf1461055c578063a3b8ef041461056f578063cc8463c81461058e578063cefc1429146105a2575f5ffd5b80638da5cb5b146104b157806391d14854146104c5578063947a36fb14610514578063a1eda53c14610529575f5ffd5b806336568abe116101a7578063649a5ec71161017757806367ea88eb1161015d57806367ea88eb146104355780637a01a1da1461045457806384ef8ffc14610487575f5ffd5b8063649a5ec7146103ea57806366d003ac14610409575f5ffd5b806336568abe146103785780633bbed4a01461039757806346b62c4a146103b6578063634e93da146103cb575f5ffd5b806322a90082116101e257806322a90082146102c8578063248a9ca3146102e75780632aaa9628146103235780632f2ff15d14610359575f5ffd5b806301ffc9a71461021d578063022d63fb1461025157806304a41159146102795780630aa6220b146102b2575f5ffd5b3661021957005b5f5ffd5b348015610228575f5ffd5b5061023c61023736600461250c565b610712565b60405190151581526020015b60405180910390f35b34801561025c575f5ffd5b50620697805b60405165ffffffffffff9091168152602001610248565b348015610284575f5ffd5b5061028d600181565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610248565b3480156102bd575f5ffd5b506102c661076d565b005b3480156102d3575f5ffd5b506102c66102e236600461254b565b610782565b3480156102f2575f5ffd5b5061031561030136600461254b565b5f9081526020819052604090206001015490565b604051908152602001610248565b34801561032e575f5ffd5b5061034261033d366004612583565b6107e0565b604080519215158352602083019190915201610248565b348015610364575f5ffd5b506102c66103733660046125ba565b6109ba565b348015610383575f5ffd5b506102c66103923660046125ba565b6109ff565b3480156103a2575f5ffd5b506102c66103b13660046125dd565b610b09565b3480156103c1575f5ffd5b5061031560055481565b3480156103d6575f5ffd5b506102c66103e53660046125dd565b610ba1565b3480156103f5575f5ffd5b506102c66104043660046125f8565b610bb4565b348015610414575f5ffd5b5060085461028d9073ffffffffffffffffffffffffffffffffffffffff1681565b348015610440575f5ffd5b506102c661044f36600461261d565b610bc7565b34801561045f575f5ffd5b506103157ff8cd32ed93fc2f9fc78152a14807c9609af3d99c5fe4dc6b106a801aaddfe90e81565b348015610492575f5ffd5b5060025473ffffffffffffffffffffffffffffffffffffffff1661028d565b3480156104bc575f5ffd5b5061028d61123b565b3480156104d0575f5ffd5b5061023c6104df3660046125ba565b5f9182526020828152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b34801561051f575f5ffd5b5061031560065481565b348015610534575f5ffd5b5061053d611260565b6040805165ffffffffffff938416815292909116602083015201610248565b348015610567575f5ffd5b506103155f81565b34801561057a575f5ffd5b506102c661058936600461254b565b6112da565b348015610599575f5ffd5b50610262611374565b3480156105ad575f5ffd5b506102c6611411565b3480156105c1575f5ffd5b506001546040805173ffffffffffffffffffffffffffffffffffffffff831681527401000000000000000000000000000000000000000090920465ffffffffffff16602083015201610248565b348015610619575f5ffd5b506102c66106283660046125ba565b61146d565b348015610638575f5ffd5b506102c66114ae565b34801561064c575f5ffd5b506103157f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b0881565b34801561067f575f5ffd5b506102c661068e36600461254b565b6114c0565b34801561069e575f5ffd5b5061031560075481565b3480156106b3575f5ffd5b506102c66106c236600461264f565b611692565b3480156106d2575f5ffd5b5060045461028d9073ffffffffffffffffffffffffffffffffffffffff1681565b3480156106fe575f5ffd5b5061031561070d3660046125dd565b61187e565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f3149878600000000000000000000000000000000000000000000000000000000148061076757506107678261194f565b92915050565b5f610777816119e5565b61077f6119ef565b50565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b086107ac816119e5565b6006546040518391907f3fb677206c0b314c404bae3da94bee6bda11375c9fe986f266fb033d6cddbbd7905f90a350600655565b5f8073ffffffffffffffffffffffffffffffffffffffff841661080757505f9050806109b3565b73ffffffffffffffffffffffffffffffffffffffff83165f9081527ff3fa603c74bfe2a4719960e47343678c3dc690d2b27a2295acc6fc430833aaf9602052604090205460ff1661085c57505f9050806109b3565b5f8473ffffffffffffffffffffffffffffffffffffffff1663c3f909d46040518163ffffffff1660e01b8152600401606060405180830381865afa1580156108a6573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108ca9190612679565b80516040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015291925073ffffffffffffffffffffffffffffffffffffffff16906370a0823190602401602060405180830381865afa158015610936573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061095a9190612709565b915060055482101561096f575f9250506109b3565b5f60065411801561098157505f600754115b156109ad575f60075442610995919061274d565b90506006548110156109ab575f935050506109b3565b505b60019250505b9250929050565b816109f1576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6109fb82826119fb565b5050565b81158015610a27575060025473ffffffffffffffffffffffffffffffffffffffff8281169116145b15610aff5760015473ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900465ffffffffffff1681151580610a7b575065ffffffffffff8116155b80610a8e57504265ffffffffffff821610155b15610ad4576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff821660048201526024015b60405180910390fd5b5050600180547fffffffffffff000000000000ffffffffffffffffffffffffffffffffffffffff1690555b6109fb8282611a25565b5f610b13816119e5565b60085460405173ffffffffffffffffffffffffffffffffffffffff8085169216907fc1416b5cdab50a9fbc872236e1aa54566c6deb40024e63a4b1737ecacf09d6f9905f90a350600880547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b5f610bab816119e5565b6109fb82611a7e565b5f610bbe816119e5565b6109fb82611afd565b610bcf611b6c565b7ff8cd32ed93fc2f9fc78152a14807c9609af3d99c5fe4dc6b106a801aaddfe90e610bf9816119e5565b73ffffffffffffffffffffffffffffffffffffffff8416610c5e576040517f235aafe400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610acb565b825f03610c9a576040517f9717f35f00000000000000000000000000000000000000000000000000000000815260048101849052602401610acb565b5f600654118015610cac57505f600754115b15610d0e575f60075442610cc0919061274d565b9050600654811015610d0c576006546040517f1e379906000000000000000000000000000000000000000000000000000000008152610acb918391600401918252602082015260400190565b505b5f8473ffffffffffffffffffffffffffffffffffffffff1663c3f909d46040518163ffffffff1660e01b8152600401606060405180830381865afa158015610d58573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d7c9190612679565b805160208201516040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015292935090915f9073ffffffffffffffffffffffffffffffffffffffff8416906370a0823190602401602060405180830381865afa158015610df2573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e169190612709565b9050600554871015610e62576005546040517f013fafe2000000000000000000000000000000000000000000000000000000008152610acb918991600401918252602082015260400190565b80871115610ea6576040517f013fafe20000000000000000000000000000000000000000000000000000000081526004810182905260248101889052604401610acb565b600480546040517fdd62ed3e000000000000000000000000000000000000000000000000000000008152309281019290925273ffffffffffffffffffffffffffffffffffffffff9081166024830181905291899186169063dd62ed3e90604401602060405180830381865afa158015610f21573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f459190612709565b1015610f8c57610f8c73ffffffffffffffffffffffffffffffffffffffff8516827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff611baf565b600480546040517f1e92008400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8c811693820193909352602481018b9052604481018a9052911690631e920084906064015f604051808303815f87803b158015611006575f5ffd5b505af1158015611018573d5f5f3e3d5ffd5b5050604080518b815233602082015273ffffffffffffffffffffffffffffffffffffffff808816945088811693508d16917fc7b4d815bd0a3b8348577971cb55a9334c80a0b3ccdfe287b55b5a6320d8480a910160405180910390a44260075560085473ffffffffffffffffffffffffffffffffffffffff161561122657478015611166576008546040515f9173ffffffffffffffffffffffffffffffffffffffff169083908381818185875af1925050503d805f81146110f4576040519150601f19603f3d011682016040523d82523d5f602084013e6110f9565b606091505b5050905080611164576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f455448207472616e73666572206661696c6564000000000000000000000000006044820152606401610acb565b505b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201525f9073ffffffffffffffffffffffffffffffffffffffff8616906370a0823190602401602060405180830381865afa1580156111d0573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111f49190612709565b90508015611223576008546112239073ffffffffffffffffffffffffffffffffffffffff878116911683611cc9565b50505b5050505050506112366001600355565b505050565b5f61125b60025473ffffffffffffffffffffffffffffffffffffffff1690565b905090565b6002545f907a010000000000000000000000000000000000000000000000000000900465ffffffffffff1680151580156112a257504265ffffffffffff821610155b6112ad575f5f6112d2565b60025474010000000000000000000000000000000000000000900465ffffffffffff16815b915091509091565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b08611304816119e5565b815f03611340576040517f517f903200000000000000000000000000000000000000000000000000000000815260048101839052602401610acb565b6005546040518391907fd6d62b78d21fc4f5151029ca37079020e8a8815c5a488b28ee7499d4c2a19e55905f90a350600555565b6002545f907a010000000000000000000000000000000000000000000000000000900465ffffffffffff1680151580156113b557504265ffffffffffff8216105b6113e7576001547a010000000000000000000000000000000000000000000000000000900465ffffffffffff1661140b565b60025474010000000000000000000000000000000000000000900465ffffffffffff165b91505090565b60015473ffffffffffffffffffffffffffffffffffffffff16338114611465576040517fc22c8022000000000000000000000000000000000000000000000000000000008152336004820152602401610acb565b61077f611d07565b816114a4576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6109fb8282611df8565b5f6114b8816119e5565b61077f611e1c565b6114c8611b6c565b5f6114d2816119e5565b475f83156114e057836114e2565b815b90508181111561154e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f496e73756666696369656e742062616c616e63650000000000000000000000006044820152606401610acb565b5f61156e60025473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff16826040515f6040518083038185875af1925050503d805f81146115c2576040519150601f19603f3d011682016040523d82523d5f602084013e6115c7565b606091505b5050905080611632576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f455448207472616e73666572206661696c6564000000000000000000000000006044820152606401610acb565b60025460405183815273ffffffffffffffffffffffffffffffffffffffff909116906001907fd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb9060200160405180910390a35050505061077f6001600355565b61169a611b6c565b5f6116a4816119e5565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015283905f9073ffffffffffffffffffffffffffffffffffffffff8316906370a0823190602401602060405180830381865afa158015611710573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906117349190612709565b90505f84156117435784611745565b815b9050818111156117b1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f496e73756666696369656e742062616c616e63650000000000000000000000006044820152606401610acb565b6117f16117d360025473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff85169083611cc9565b60025473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff167fd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb8360405161186891815260200190565b60405180910390a3505050506109fb6001600355565b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff73ffffffffffffffffffffffffffffffffffffffff8316016118c3575047919050565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff8316906370a0823190602401602060405180830381865afa15801561192b573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107679190612709565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061076757507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000831614610767565b61077f8133611e26565b6119f95f5f611eab565b565b5f82815260208190526040902060010154611a15816119e5565b611a1f8383612004565b50505050565b73ffffffffffffffffffffffffffffffffffffffff81163314611a74576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61123682826120c9565b5f611a87611374565b611a904261212a565b611a9a9190612760565b9050611aa68282612179565b60405165ffffffffffff8216815273ffffffffffffffffffffffffffffffffffffffff8316907f3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed69060200160405180910390a25050565b5f611b0782612214565b611b104261212a565b611b1a9190612760565b9050611b268282611eab565b6040805165ffffffffffff8085168252831660208201527ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b910160405180910390a15050565b600260035403611ba8576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600355565b6040805173ffffffffffffffffffffffffffffffffffffffff8416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f095ea7b300000000000000000000000000000000000000000000000000000000179052611c3b8482612265565b611a1f5760405173ffffffffffffffffffffffffffffffffffffffff84811660248301525f6044830152611cbf91869182169063095ea7b3906064015b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506122bb565b611a1f84826122bb565b60405173ffffffffffffffffffffffffffffffffffffffff83811660248301526044820183905261123691859182169063a9059cbb90606401611c78565b60015473ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900465ffffffffffff16801580611d5757504265ffffffffffff821610155b15611d98576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff82166004820152602401610acb565b611dc05f611dbb60025473ffffffffffffffffffffffffffffffffffffffff1690565b6120c9565b50611dcb5f83612004565b5050600180547fffffffffffff000000000000000000000000000000000000000000000000000016905550565b5f82815260208190526040902060010154611e12816119e5565b611a1f83836120c9565b6119f95f5f612179565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff166109fb576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8216600482015260248101839052604401610acb565b6002547a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015611f7f574265ffffffffffff82161015611f56576002546001805479ffffffffffffffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000090920465ffffffffffff167a01000000000000000000000000000000000000000000000000000002919091179055611f7f565b6040517f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5905f90a15b506002805473ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000065ffffffffffff9485160279ffffffffffffffffffffffffffffffffffffffffffffffffffff16177a0100000000000000000000000000000000000000000000000000009290931691909102919091179055565b5f826120b8575f61202a60025473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614612077576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84161790555b6120c2838361235a565b9392505050565b5f821580156120f2575060025473ffffffffffffffffffffffffffffffffffffffff8381169116145b1561212057600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001690555b6120c28383612453565b5f65ffffffffffff821115612175576040517f6dfcc6500000000000000000000000000000000000000000000000000000000081526030600482015260248101839052604401610acb565b5090565b600180547401000000000000000000000000000000000000000065ffffffffffff84811682027fffffffffffff0000000000000000000000000000000000000000000000000000841673ffffffffffffffffffffffffffffffffffffffff881617179093559004168015611236576040517f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109905f90a1505050565b5f5f61221e611374565b90508065ffffffffffff168365ffffffffffff161161224657612241838261277e565b6120c2565b6120c265ffffffffffff8416620697805f8282188284100282186120c2565b5f5f5f5f60205f8651602088015f8a5af192503d91505f5190508280156122b15750811561229657806001146122b1565b5f8673ffffffffffffffffffffffffffffffffffffffff163b115b9695505050505050565b5f5f60205f8451602086015f885af1806122da576040513d5f823e3d81fd5b50505f513d915081156122f157806001141561230b565b73ffffffffffffffffffffffffffffffffffffffff84163b155b15611a1f576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610acb565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff1661244c575f8381526020818152604080832073ffffffffffffffffffffffffffffffffffffffff86168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556123ea3390565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4506001610767565b505f610767565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff161561244c575f8381526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8616808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a4506001610767565b5f6020828403121561251c575f5ffd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146120c2575f5ffd5b5f6020828403121561255b575f5ffd5b5035919050565b73ffffffffffffffffffffffffffffffffffffffff8116811461077f575f5ffd5b5f5f60408385031215612594575f5ffd5b823561259f81612562565b915060208301356125af81612562565b809150509250929050565b5f5f604083850312156125cb575f5ffd5b8235915060208301356125af81612562565b5f602082840312156125ed575f5ffd5b81356120c281612562565b5f60208284031215612608575f5ffd5b813565ffffffffffff811681146120c2575f5ffd5b5f5f5f6060848603121561262f575f5ffd5b833561263a81612562565b95602085013595506040909401359392505050565b5f5f60408385031215612660575f5ffd5b823561266b81612562565b946020939093013593505050565b5f606082840312801561268a575f5ffd5b506040516060810167ffffffffffffffff811182821017156126d3577f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b60405282516126e181612562565b815260208301516126f181612562565b60208201526040928301519281019290925250919050565b5f60208284031215612719575f5ffd5b5051919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b8181038181111561076757610767612720565b65ffffffffffff818116838216019081111561076757610767612720565b65ffffffffffff82811682821603908111156107675761076761272056fea2646970667358221220d08781402aaf80cc8eb5a19b063a4445501657f3b2a80d7dc3bd581b03c20aef64736f6c634300081c0033241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b08f8cd32ed93fc2f9fc78152a14807c9609af3d99c5fe4dc6b106a801aaddfe90e",
}

// BuyBotABI is the input ABI used to generate the binding from.
// Deprecated: Use BuyBotMetaData.ABI instead.
var BuyBotABI = BuyBotMetaData.ABI

// Deprecated: Use BuyBotMetaData.Sigs instead.
// BuyBotFuncSigs maps the 4-byte function signature to its string representation.
var BuyBotFuncSigs = BuyBotMetaData.Sigs

// BuyBotBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BuyBotMetaData.Bin instead.
var BuyBotBin = BuyBotMetaData.Bin

// DeployBuyBot deploys a new Ethereum contract, binding an instance of BuyBot to it.
func DeployBuyBot(auth *bind.TransactOpts, backend bind.ContractBackend, _initialDelay *big.Int, _owner common.Address, _router common.Address, _minOrderAmount *big.Int, _interval *big.Int, _recipient common.Address, _buyer common.Address, _manager common.Address) (common.Address, *types.Transaction, *BuyBot, error) {
	parsed, err := BuyBotMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BuyBotBin), backend, _initialDelay, _owner, _router, _minOrderAmount, _interval, _recipient, _buyer, _manager)
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

// BUYERROLE is a free data retrieval call binding the contract method 0x7a01a1da.
//
// Solidity: function BUYER_ROLE() view returns(bytes32)
func (_BuyBot *BuyBotCaller) BUYERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BuyBot.contract.Call(opts, &out, "BUYER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BUYERROLE is a free data retrieval call binding the contract method 0x7a01a1da.
//
// Solidity: function BUYER_ROLE() view returns(bytes32)
func (_BuyBot *BuyBotSession) BUYERROLE() ([32]byte, error) {
	return _BuyBot.Contract.BUYERROLE(&_BuyBot.CallOpts)
}

// BUYERROLE is a free data retrieval call binding the contract method 0x7a01a1da.
//
// Solidity: function BUYER_ROLE() view returns(bytes32)
func (_BuyBot *BuyBotCallerSession) BUYERROLE() ([32]byte, error) {
	return _BuyBot.Contract.BUYERROLE(&_BuyBot.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BuyBot *BuyBotCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BuyBot.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BuyBot *BuyBotSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BuyBot.Contract.DEFAULTADMINROLE(&_BuyBot.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BuyBot *BuyBotCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BuyBot.Contract.DEFAULTADMINROLE(&_BuyBot.CallOpts)
}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_BuyBot *BuyBotCaller) MANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BuyBot.contract.Call(opts, &out, "MANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_BuyBot *BuyBotSession) MANAGERROLE() ([32]byte, error) {
	return _BuyBot.Contract.MANAGERROLE(&_BuyBot.CallOpts)
}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_BuyBot *BuyBotCallerSession) MANAGERROLE() ([32]byte, error) {
	return _BuyBot.Contract.MANAGERROLE(&_BuyBot.CallOpts)
}

// NATIVECOIN is a free data retrieval call binding the contract method 0x04a41159.
//
// Solidity: function NATIVE_COIN() view returns(address)
func (_BuyBot *BuyBotCaller) NATIVECOIN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BuyBot.contract.Call(opts, &out, "NATIVE_COIN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NATIVECOIN is a free data retrieval call binding the contract method 0x04a41159.
//
// Solidity: function NATIVE_COIN() view returns(address)
func (_BuyBot *BuyBotSession) NATIVECOIN() (common.Address, error) {
	return _BuyBot.Contract.NATIVECOIN(&_BuyBot.CallOpts)
}

// NATIVECOIN is a free data retrieval call binding the contract method 0x04a41159.
//
// Solidity: function NATIVE_COIN() view returns(address)
func (_BuyBot *BuyBotCallerSession) NATIVECOIN() (common.Address, error) {
	return _BuyBot.Contract.NATIVECOIN(&_BuyBot.CallOpts)
}

// CanBuyMarket is a free data retrieval call binding the contract method 0x2aaa9628.
//
// Solidity: function canBuyMarket(address pair, address caller) view returns(bool canBuy, uint256 balance)
func (_BuyBot *BuyBotCaller) CanBuyMarket(opts *bind.CallOpts, pair common.Address, caller common.Address) (struct {
	CanBuy  bool
	Balance *big.Int
}, error) {
	var out []interface{}
	err := _BuyBot.contract.Call(opts, &out, "canBuyMarket", pair, caller)

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

// CanBuyMarket is a free data retrieval call binding the contract method 0x2aaa9628.
//
// Solidity: function canBuyMarket(address pair, address caller) view returns(bool canBuy, uint256 balance)
func (_BuyBot *BuyBotSession) CanBuyMarket(pair common.Address, caller common.Address) (struct {
	CanBuy  bool
	Balance *big.Int
}, error) {
	return _BuyBot.Contract.CanBuyMarket(&_BuyBot.CallOpts, pair, caller)
}

// CanBuyMarket is a free data retrieval call binding the contract method 0x2aaa9628.
//
// Solidity: function canBuyMarket(address pair, address caller) view returns(bool canBuy, uint256 balance)
func (_BuyBot *BuyBotCallerSession) CanBuyMarket(pair common.Address, caller common.Address) (struct {
	CanBuy  bool
	Balance *big.Int
}, error) {
	return _BuyBot.Contract.CanBuyMarket(&_BuyBot.CallOpts, pair, caller)
}

// DefaultAdmin is a free data retrieval call binding the contract method 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (_BuyBot *BuyBotCaller) DefaultAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BuyBot.contract.Call(opts, &out, "defaultAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultAdmin is a free data retrieval call binding the contract method 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (_BuyBot *BuyBotSession) DefaultAdmin() (common.Address, error) {
	return _BuyBot.Contract.DefaultAdmin(&_BuyBot.CallOpts)
}

// DefaultAdmin is a free data retrieval call binding the contract method 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (_BuyBot *BuyBotCallerSession) DefaultAdmin() (common.Address, error) {
	return _BuyBot.Contract.DefaultAdmin(&_BuyBot.CallOpts)
}

// DefaultAdminDelay is a free data retrieval call binding the contract method 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (_BuyBot *BuyBotCaller) DefaultAdminDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BuyBot.contract.Call(opts, &out, "defaultAdminDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultAdminDelay is a free data retrieval call binding the contract method 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (_BuyBot *BuyBotSession) DefaultAdminDelay() (*big.Int, error) {
	return _BuyBot.Contract.DefaultAdminDelay(&_BuyBot.CallOpts)
}

// DefaultAdminDelay is a free data retrieval call binding the contract method 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (_BuyBot *BuyBotCallerSession) DefaultAdminDelay() (*big.Int, error) {
	return _BuyBot.Contract.DefaultAdminDelay(&_BuyBot.CallOpts)
}

// DefaultAdminDelayIncreaseWait is a free data retrieval call binding the contract method 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (_BuyBot *BuyBotCaller) DefaultAdminDelayIncreaseWait(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BuyBot.contract.Call(opts, &out, "defaultAdminDelayIncreaseWait")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultAdminDelayIncreaseWait is a free data retrieval call binding the contract method 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (_BuyBot *BuyBotSession) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _BuyBot.Contract.DefaultAdminDelayIncreaseWait(&_BuyBot.CallOpts)
}

// DefaultAdminDelayIncreaseWait is a free data retrieval call binding the contract method 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (_BuyBot *BuyBotCallerSession) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _BuyBot.Contract.DefaultAdminDelayIncreaseWait(&_BuyBot.CallOpts)
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

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BuyBot *BuyBotCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _BuyBot.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BuyBot *BuyBotSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BuyBot.Contract.GetRoleAdmin(&_BuyBot.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BuyBot *BuyBotCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BuyBot.Contract.GetRoleAdmin(&_BuyBot.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BuyBot *BuyBotCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _BuyBot.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BuyBot *BuyBotSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BuyBot.Contract.HasRole(&_BuyBot.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BuyBot *BuyBotCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BuyBot.Contract.HasRole(&_BuyBot.CallOpts, role, account)
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

// PendingDefaultAdmin is a free data retrieval call binding the contract method 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (_BuyBot *BuyBotCaller) PendingDefaultAdmin(opts *bind.CallOpts) (struct {
	NewAdmin common.Address
	Schedule *big.Int
}, error) {
	var out []interface{}
	err := _BuyBot.contract.Call(opts, &out, "pendingDefaultAdmin")

	outstruct := new(struct {
		NewAdmin common.Address
		Schedule *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NewAdmin = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Schedule = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PendingDefaultAdmin is a free data retrieval call binding the contract method 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (_BuyBot *BuyBotSession) PendingDefaultAdmin() (struct {
	NewAdmin common.Address
	Schedule *big.Int
}, error) {
	return _BuyBot.Contract.PendingDefaultAdmin(&_BuyBot.CallOpts)
}

// PendingDefaultAdmin is a free data retrieval call binding the contract method 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (_BuyBot *BuyBotCallerSession) PendingDefaultAdmin() (struct {
	NewAdmin common.Address
	Schedule *big.Int
}, error) {
	return _BuyBot.Contract.PendingDefaultAdmin(&_BuyBot.CallOpts)
}

// PendingDefaultAdminDelay is a free data retrieval call binding the contract method 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (_BuyBot *BuyBotCaller) PendingDefaultAdminDelay(opts *bind.CallOpts) (struct {
	NewDelay *big.Int
	Schedule *big.Int
}, error) {
	var out []interface{}
	err := _BuyBot.contract.Call(opts, &out, "pendingDefaultAdminDelay")

	outstruct := new(struct {
		NewDelay *big.Int
		Schedule *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NewDelay = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Schedule = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PendingDefaultAdminDelay is a free data retrieval call binding the contract method 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (_BuyBot *BuyBotSession) PendingDefaultAdminDelay() (struct {
	NewDelay *big.Int
	Schedule *big.Int
}, error) {
	return _BuyBot.Contract.PendingDefaultAdminDelay(&_BuyBot.CallOpts)
}

// PendingDefaultAdminDelay is a free data retrieval call binding the contract method 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (_BuyBot *BuyBotCallerSession) PendingDefaultAdminDelay() (struct {
	NewDelay *big.Int
	Schedule *big.Int
}, error) {
	return _BuyBot.Contract.PendingDefaultAdminDelay(&_BuyBot.CallOpts)
}

// Recipient is a free data retrieval call binding the contract method 0x66d003ac.
//
// Solidity: function recipient() view returns(address)
func (_BuyBot *BuyBotCaller) Recipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BuyBot.contract.Call(opts, &out, "recipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Recipient is a free data retrieval call binding the contract method 0x66d003ac.
//
// Solidity: function recipient() view returns(address)
func (_BuyBot *BuyBotSession) Recipient() (common.Address, error) {
	return _BuyBot.Contract.Recipient(&_BuyBot.CallOpts)
}

// Recipient is a free data retrieval call binding the contract method 0x66d003ac.
//
// Solidity: function recipient() view returns(address)
func (_BuyBot *BuyBotCallerSession) Recipient() (common.Address, error) {
	return _BuyBot.Contract.Recipient(&_BuyBot.CallOpts)
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

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BuyBot *BuyBotCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BuyBot.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BuyBot *BuyBotSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BuyBot.Contract.SupportsInterface(&_BuyBot.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BuyBot *BuyBotCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BuyBot.Contract.SupportsInterface(&_BuyBot.CallOpts, interfaceId)
}

// AcceptDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xcefc1429.
//
// Solidity: function acceptDefaultAdminTransfer() returns()
func (_BuyBot *BuyBotTransactor) AcceptDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BuyBot.contract.Transact(opts, "acceptDefaultAdminTransfer")
}

// AcceptDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xcefc1429.
//
// Solidity: function acceptDefaultAdminTransfer() returns()
func (_BuyBot *BuyBotSession) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _BuyBot.Contract.AcceptDefaultAdminTransfer(&_BuyBot.TransactOpts)
}

// AcceptDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xcefc1429.
//
// Solidity: function acceptDefaultAdminTransfer() returns()
func (_BuyBot *BuyBotTransactorSession) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _BuyBot.Contract.AcceptDefaultAdminTransfer(&_BuyBot.TransactOpts)
}

// BeginDefaultAdminTransfer is a paid mutator transaction binding the contract method 0x634e93da.
//
// Solidity: function beginDefaultAdminTransfer(address newAdmin) returns()
func (_BuyBot *BuyBotTransactor) BeginDefaultAdminTransfer(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _BuyBot.contract.Transact(opts, "beginDefaultAdminTransfer", newAdmin)
}

// BeginDefaultAdminTransfer is a paid mutator transaction binding the contract method 0x634e93da.
//
// Solidity: function beginDefaultAdminTransfer(address newAdmin) returns()
func (_BuyBot *BuyBotSession) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _BuyBot.Contract.BeginDefaultAdminTransfer(&_BuyBot.TransactOpts, newAdmin)
}

// BeginDefaultAdminTransfer is a paid mutator transaction binding the contract method 0x634e93da.
//
// Solidity: function beginDefaultAdminTransfer(address newAdmin) returns()
func (_BuyBot *BuyBotTransactorSession) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _BuyBot.Contract.BeginDefaultAdminTransfer(&_BuyBot.TransactOpts, newAdmin)
}

// BuyMarket is a paid mutator transaction binding the contract method 0x67ea88eb.
//
// Solidity: function buyMarket(address pair, uint256 amount, uint256 maxMatchCount) returns()
func (_BuyBot *BuyBotTransactor) BuyMarket(opts *bind.TransactOpts, pair common.Address, amount *big.Int, maxMatchCount *big.Int) (*types.Transaction, error) {
	return _BuyBot.contract.Transact(opts, "buyMarket", pair, amount, maxMatchCount)
}

// BuyMarket is a paid mutator transaction binding the contract method 0x67ea88eb.
//
// Solidity: function buyMarket(address pair, uint256 amount, uint256 maxMatchCount) returns()
func (_BuyBot *BuyBotSession) BuyMarket(pair common.Address, amount *big.Int, maxMatchCount *big.Int) (*types.Transaction, error) {
	return _BuyBot.Contract.BuyMarket(&_BuyBot.TransactOpts, pair, amount, maxMatchCount)
}

// BuyMarket is a paid mutator transaction binding the contract method 0x67ea88eb.
//
// Solidity: function buyMarket(address pair, uint256 amount, uint256 maxMatchCount) returns()
func (_BuyBot *BuyBotTransactorSession) BuyMarket(pair common.Address, amount *big.Int, maxMatchCount *big.Int) (*types.Transaction, error) {
	return _BuyBot.Contract.BuyMarket(&_BuyBot.TransactOpts, pair, amount, maxMatchCount)
}

// CancelDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (_BuyBot *BuyBotTransactor) CancelDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BuyBot.contract.Transact(opts, "cancelDefaultAdminTransfer")
}

// CancelDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (_BuyBot *BuyBotSession) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _BuyBot.Contract.CancelDefaultAdminTransfer(&_BuyBot.TransactOpts)
}

// CancelDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (_BuyBot *BuyBotTransactorSession) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _BuyBot.Contract.CancelDefaultAdminTransfer(&_BuyBot.TransactOpts)
}

// ChangeDefaultAdminDelay is a paid mutator transaction binding the contract method 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (_BuyBot *BuyBotTransactor) ChangeDefaultAdminDelay(opts *bind.TransactOpts, newDelay *big.Int) (*types.Transaction, error) {
	return _BuyBot.contract.Transact(opts, "changeDefaultAdminDelay", newDelay)
}

// ChangeDefaultAdminDelay is a paid mutator transaction binding the contract method 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (_BuyBot *BuyBotSession) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _BuyBot.Contract.ChangeDefaultAdminDelay(&_BuyBot.TransactOpts, newDelay)
}

// ChangeDefaultAdminDelay is a paid mutator transaction binding the contract method 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (_BuyBot *BuyBotTransactorSession) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _BuyBot.Contract.ChangeDefaultAdminDelay(&_BuyBot.TransactOpts, newDelay)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BuyBot *BuyBotTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BuyBot.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BuyBot *BuyBotSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BuyBot.Contract.GrantRole(&_BuyBot.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BuyBot *BuyBotTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BuyBot.Contract.GrantRole(&_BuyBot.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_BuyBot *BuyBotTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BuyBot.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_BuyBot *BuyBotSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BuyBot.Contract.RenounceRole(&_BuyBot.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_BuyBot *BuyBotTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BuyBot.Contract.RenounceRole(&_BuyBot.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BuyBot *BuyBotTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BuyBot.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BuyBot *BuyBotSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BuyBot.Contract.RevokeRole(&_BuyBot.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BuyBot *BuyBotTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BuyBot.Contract.RevokeRole(&_BuyBot.TransactOpts, role, account)
}

// RollbackDefaultAdminDelay is a paid mutator transaction binding the contract method 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (_BuyBot *BuyBotTransactor) RollbackDefaultAdminDelay(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BuyBot.contract.Transact(opts, "rollbackDefaultAdminDelay")
}

// RollbackDefaultAdminDelay is a paid mutator transaction binding the contract method 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (_BuyBot *BuyBotSession) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _BuyBot.Contract.RollbackDefaultAdminDelay(&_BuyBot.TransactOpts)
}

// RollbackDefaultAdminDelay is a paid mutator transaction binding the contract method 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (_BuyBot *BuyBotTransactorSession) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _BuyBot.Contract.RollbackDefaultAdminDelay(&_BuyBot.TransactOpts)
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

// SetRecipient is a paid mutator transaction binding the contract method 0x3bbed4a0.
//
// Solidity: function setRecipient(address _recipient) returns()
func (_BuyBot *BuyBotTransactor) SetRecipient(opts *bind.TransactOpts, _recipient common.Address) (*types.Transaction, error) {
	return _BuyBot.contract.Transact(opts, "setRecipient", _recipient)
}

// SetRecipient is a paid mutator transaction binding the contract method 0x3bbed4a0.
//
// Solidity: function setRecipient(address _recipient) returns()
func (_BuyBot *BuyBotSession) SetRecipient(_recipient common.Address) (*types.Transaction, error) {
	return _BuyBot.Contract.SetRecipient(&_BuyBot.TransactOpts, _recipient)
}

// SetRecipient is a paid mutator transaction binding the contract method 0x3bbed4a0.
//
// Solidity: function setRecipient(address _recipient) returns()
func (_BuyBot *BuyBotTransactorSession) SetRecipient(_recipient common.Address) (*types.Transaction, error) {
	return _BuyBot.Contract.SetRecipient(&_BuyBot.TransactOpts, _recipient)
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

// BuyBotDefaultAdminDelayChangeCanceledIterator is returned from FilterDefaultAdminDelayChangeCanceled and is used to iterate over the raw logs and unpacked data for DefaultAdminDelayChangeCanceled events raised by the BuyBot contract.
type BuyBotDefaultAdminDelayChangeCanceledIterator struct {
	Event *BuyBotDefaultAdminDelayChangeCanceled // Event containing the contract specifics and raw log

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
func (it *BuyBotDefaultAdminDelayChangeCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuyBotDefaultAdminDelayChangeCanceled)
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
		it.Event = new(BuyBotDefaultAdminDelayChangeCanceled)
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
func (it *BuyBotDefaultAdminDelayChangeCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuyBotDefaultAdminDelayChangeCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuyBotDefaultAdminDelayChangeCanceled represents a DefaultAdminDelayChangeCanceled event raised by the BuyBot contract.
type BuyBotDefaultAdminDelayChangeCanceled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminDelayChangeCanceled is a free log retrieval operation binding the contract event 0x2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5.
//
// Solidity: event DefaultAdminDelayChangeCanceled()
func (_BuyBot *BuyBotFilterer) FilterDefaultAdminDelayChangeCanceled(opts *bind.FilterOpts) (*BuyBotDefaultAdminDelayChangeCanceledIterator, error) {

	logs, sub, err := _BuyBot.contract.FilterLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return &BuyBotDefaultAdminDelayChangeCanceledIterator{contract: _BuyBot.contract, event: "DefaultAdminDelayChangeCanceled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminDelayChangeCanceled is a free log subscription operation binding the contract event 0x2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5.
//
// Solidity: event DefaultAdminDelayChangeCanceled()
func (_BuyBot *BuyBotFilterer) WatchDefaultAdminDelayChangeCanceled(opts *bind.WatchOpts, sink chan<- *BuyBotDefaultAdminDelayChangeCanceled) (event.Subscription, error) {

	logs, sub, err := _BuyBot.contract.WatchLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuyBotDefaultAdminDelayChangeCanceled)
				if err := _BuyBot.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
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

// ParseDefaultAdminDelayChangeCanceled is a log parse operation binding the contract event 0x2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5.
//
// Solidity: event DefaultAdminDelayChangeCanceled()
func (_BuyBot *BuyBotFilterer) ParseDefaultAdminDelayChangeCanceled(log types.Log) (*BuyBotDefaultAdminDelayChangeCanceled, error) {
	event := new(BuyBotDefaultAdminDelayChangeCanceled)
	if err := _BuyBot.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BuyBotDefaultAdminDelayChangeScheduledIterator is returned from FilterDefaultAdminDelayChangeScheduled and is used to iterate over the raw logs and unpacked data for DefaultAdminDelayChangeScheduled events raised by the BuyBot contract.
type BuyBotDefaultAdminDelayChangeScheduledIterator struct {
	Event *BuyBotDefaultAdminDelayChangeScheduled // Event containing the contract specifics and raw log

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
func (it *BuyBotDefaultAdminDelayChangeScheduledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuyBotDefaultAdminDelayChangeScheduled)
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
		it.Event = new(BuyBotDefaultAdminDelayChangeScheduled)
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
func (it *BuyBotDefaultAdminDelayChangeScheduledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuyBotDefaultAdminDelayChangeScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuyBotDefaultAdminDelayChangeScheduled represents a DefaultAdminDelayChangeScheduled event raised by the BuyBot contract.
type BuyBotDefaultAdminDelayChangeScheduled struct {
	NewDelay       *big.Int
	EffectSchedule *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminDelayChangeScheduled is a free log retrieval operation binding the contract event 0xf1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b.
//
// Solidity: event DefaultAdminDelayChangeScheduled(uint48 newDelay, uint48 effectSchedule)
func (_BuyBot *BuyBotFilterer) FilterDefaultAdminDelayChangeScheduled(opts *bind.FilterOpts) (*BuyBotDefaultAdminDelayChangeScheduledIterator, error) {

	logs, sub, err := _BuyBot.contract.FilterLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return &BuyBotDefaultAdminDelayChangeScheduledIterator{contract: _BuyBot.contract, event: "DefaultAdminDelayChangeScheduled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminDelayChangeScheduled is a free log subscription operation binding the contract event 0xf1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b.
//
// Solidity: event DefaultAdminDelayChangeScheduled(uint48 newDelay, uint48 effectSchedule)
func (_BuyBot *BuyBotFilterer) WatchDefaultAdminDelayChangeScheduled(opts *bind.WatchOpts, sink chan<- *BuyBotDefaultAdminDelayChangeScheduled) (event.Subscription, error) {

	logs, sub, err := _BuyBot.contract.WatchLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuyBotDefaultAdminDelayChangeScheduled)
				if err := _BuyBot.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
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

// ParseDefaultAdminDelayChangeScheduled is a log parse operation binding the contract event 0xf1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b.
//
// Solidity: event DefaultAdminDelayChangeScheduled(uint48 newDelay, uint48 effectSchedule)
func (_BuyBot *BuyBotFilterer) ParseDefaultAdminDelayChangeScheduled(log types.Log) (*BuyBotDefaultAdminDelayChangeScheduled, error) {
	event := new(BuyBotDefaultAdminDelayChangeScheduled)
	if err := _BuyBot.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BuyBotDefaultAdminTransferCanceledIterator is returned from FilterDefaultAdminTransferCanceled and is used to iterate over the raw logs and unpacked data for DefaultAdminTransferCanceled events raised by the BuyBot contract.
type BuyBotDefaultAdminTransferCanceledIterator struct {
	Event *BuyBotDefaultAdminTransferCanceled // Event containing the contract specifics and raw log

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
func (it *BuyBotDefaultAdminTransferCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuyBotDefaultAdminTransferCanceled)
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
		it.Event = new(BuyBotDefaultAdminTransferCanceled)
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
func (it *BuyBotDefaultAdminTransferCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuyBotDefaultAdminTransferCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuyBotDefaultAdminTransferCanceled represents a DefaultAdminTransferCanceled event raised by the BuyBot contract.
type BuyBotDefaultAdminTransferCanceled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminTransferCanceled is a free log retrieval operation binding the contract event 0x8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109.
//
// Solidity: event DefaultAdminTransferCanceled()
func (_BuyBot *BuyBotFilterer) FilterDefaultAdminTransferCanceled(opts *bind.FilterOpts) (*BuyBotDefaultAdminTransferCanceledIterator, error) {

	logs, sub, err := _BuyBot.contract.FilterLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return &BuyBotDefaultAdminTransferCanceledIterator{contract: _BuyBot.contract, event: "DefaultAdminTransferCanceled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminTransferCanceled is a free log subscription operation binding the contract event 0x8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109.
//
// Solidity: event DefaultAdminTransferCanceled()
func (_BuyBot *BuyBotFilterer) WatchDefaultAdminTransferCanceled(opts *bind.WatchOpts, sink chan<- *BuyBotDefaultAdminTransferCanceled) (event.Subscription, error) {

	logs, sub, err := _BuyBot.contract.WatchLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuyBotDefaultAdminTransferCanceled)
				if err := _BuyBot.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
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

// ParseDefaultAdminTransferCanceled is a log parse operation binding the contract event 0x8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109.
//
// Solidity: event DefaultAdminTransferCanceled()
func (_BuyBot *BuyBotFilterer) ParseDefaultAdminTransferCanceled(log types.Log) (*BuyBotDefaultAdminTransferCanceled, error) {
	event := new(BuyBotDefaultAdminTransferCanceled)
	if err := _BuyBot.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BuyBotDefaultAdminTransferScheduledIterator is returned from FilterDefaultAdminTransferScheduled and is used to iterate over the raw logs and unpacked data for DefaultAdminTransferScheduled events raised by the BuyBot contract.
type BuyBotDefaultAdminTransferScheduledIterator struct {
	Event *BuyBotDefaultAdminTransferScheduled // Event containing the contract specifics and raw log

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
func (it *BuyBotDefaultAdminTransferScheduledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuyBotDefaultAdminTransferScheduled)
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
		it.Event = new(BuyBotDefaultAdminTransferScheduled)
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
func (it *BuyBotDefaultAdminTransferScheduledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuyBotDefaultAdminTransferScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuyBotDefaultAdminTransferScheduled represents a DefaultAdminTransferScheduled event raised by the BuyBot contract.
type BuyBotDefaultAdminTransferScheduled struct {
	NewAdmin       common.Address
	AcceptSchedule *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminTransferScheduled is a free log retrieval operation binding the contract event 0x3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6.
//
// Solidity: event DefaultAdminTransferScheduled(address indexed newAdmin, uint48 acceptSchedule)
func (_BuyBot *BuyBotFilterer) FilterDefaultAdminTransferScheduled(opts *bind.FilterOpts, newAdmin []common.Address) (*BuyBotDefaultAdminTransferScheduledIterator, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _BuyBot.contract.FilterLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return &BuyBotDefaultAdminTransferScheduledIterator{contract: _BuyBot.contract, event: "DefaultAdminTransferScheduled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminTransferScheduled is a free log subscription operation binding the contract event 0x3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6.
//
// Solidity: event DefaultAdminTransferScheduled(address indexed newAdmin, uint48 acceptSchedule)
func (_BuyBot *BuyBotFilterer) WatchDefaultAdminTransferScheduled(opts *bind.WatchOpts, sink chan<- *BuyBotDefaultAdminTransferScheduled, newAdmin []common.Address) (event.Subscription, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _BuyBot.contract.WatchLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuyBotDefaultAdminTransferScheduled)
				if err := _BuyBot.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
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

// ParseDefaultAdminTransferScheduled is a log parse operation binding the contract event 0x3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6.
//
// Solidity: event DefaultAdminTransferScheduled(address indexed newAdmin, uint48 acceptSchedule)
func (_BuyBot *BuyBotFilterer) ParseDefaultAdminTransferScheduled(log types.Log) (*BuyBotDefaultAdminTransferScheduled, error) {
	event := new(BuyBotDefaultAdminTransferScheduled)
	if err := _BuyBot.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// BuyBotRecipientSetIterator is returned from FilterRecipientSet and is used to iterate over the raw logs and unpacked data for RecipientSet events raised by the BuyBot contract.
type BuyBotRecipientSetIterator struct {
	Event *BuyBotRecipientSet // Event containing the contract specifics and raw log

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
func (it *BuyBotRecipientSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuyBotRecipientSet)
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
		it.Event = new(BuyBotRecipientSet)
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
func (it *BuyBotRecipientSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuyBotRecipientSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuyBotRecipientSet represents a RecipientSet event raised by the BuyBot contract.
type BuyBotRecipientSet struct {
	Before  common.Address
	Current common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRecipientSet is a free log retrieval operation binding the contract event 0xc1416b5cdab50a9fbc872236e1aa54566c6deb40024e63a4b1737ecacf09d6f9.
//
// Solidity: event RecipientSet(address indexed before, address indexed current)
func (_BuyBot *BuyBotFilterer) FilterRecipientSet(opts *bind.FilterOpts, before []common.Address, current []common.Address) (*BuyBotRecipientSetIterator, error) {

	var beforeRule []interface{}
	for _, beforeItem := range before {
		beforeRule = append(beforeRule, beforeItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _BuyBot.contract.FilterLogs(opts, "RecipientSet", beforeRule, currentRule)
	if err != nil {
		return nil, err
	}
	return &BuyBotRecipientSetIterator{contract: _BuyBot.contract, event: "RecipientSet", logs: logs, sub: sub}, nil
}

// WatchRecipientSet is a free log subscription operation binding the contract event 0xc1416b5cdab50a9fbc872236e1aa54566c6deb40024e63a4b1737ecacf09d6f9.
//
// Solidity: event RecipientSet(address indexed before, address indexed current)
func (_BuyBot *BuyBotFilterer) WatchRecipientSet(opts *bind.WatchOpts, sink chan<- *BuyBotRecipientSet, before []common.Address, current []common.Address) (event.Subscription, error) {

	var beforeRule []interface{}
	for _, beforeItem := range before {
		beforeRule = append(beforeRule, beforeItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _BuyBot.contract.WatchLogs(opts, "RecipientSet", beforeRule, currentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuyBotRecipientSet)
				if err := _BuyBot.contract.UnpackLog(event, "RecipientSet", log); err != nil {
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

// ParseRecipientSet is a log parse operation binding the contract event 0xc1416b5cdab50a9fbc872236e1aa54566c6deb40024e63a4b1737ecacf09d6f9.
//
// Solidity: event RecipientSet(address indexed before, address indexed current)
func (_BuyBot *BuyBotFilterer) ParseRecipientSet(log types.Log) (*BuyBotRecipientSet, error) {
	event := new(BuyBotRecipientSet)
	if err := _BuyBot.contract.UnpackLog(event, "RecipientSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BuyBotRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the BuyBot contract.
type BuyBotRoleAdminChangedIterator struct {
	Event *BuyBotRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *BuyBotRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuyBotRoleAdminChanged)
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
		it.Event = new(BuyBotRoleAdminChanged)
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
func (it *BuyBotRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuyBotRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuyBotRoleAdminChanged represents a RoleAdminChanged event raised by the BuyBot contract.
type BuyBotRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BuyBot *BuyBotFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BuyBotRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _BuyBot.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &BuyBotRoleAdminChangedIterator{contract: _BuyBot.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BuyBot *BuyBotFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BuyBotRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _BuyBot.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuyBotRoleAdminChanged)
				if err := _BuyBot.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BuyBot *BuyBotFilterer) ParseRoleAdminChanged(log types.Log) (*BuyBotRoleAdminChanged, error) {
	event := new(BuyBotRoleAdminChanged)
	if err := _BuyBot.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BuyBotRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the BuyBot contract.
type BuyBotRoleGrantedIterator struct {
	Event *BuyBotRoleGranted // Event containing the contract specifics and raw log

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
func (it *BuyBotRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuyBotRoleGranted)
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
		it.Event = new(BuyBotRoleGranted)
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
func (it *BuyBotRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuyBotRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuyBotRoleGranted represents a RoleGranted event raised by the BuyBot contract.
type BuyBotRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BuyBot *BuyBotFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BuyBotRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BuyBot.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BuyBotRoleGrantedIterator{contract: _BuyBot.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BuyBot *BuyBotFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BuyBotRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BuyBot.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuyBotRoleGranted)
				if err := _BuyBot.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BuyBot *BuyBotFilterer) ParseRoleGranted(log types.Log) (*BuyBotRoleGranted, error) {
	event := new(BuyBotRoleGranted)
	if err := _BuyBot.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BuyBotRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the BuyBot contract.
type BuyBotRoleRevokedIterator struct {
	Event *BuyBotRoleRevoked // Event containing the contract specifics and raw log

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
func (it *BuyBotRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuyBotRoleRevoked)
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
		it.Event = new(BuyBotRoleRevoked)
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
func (it *BuyBotRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuyBotRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuyBotRoleRevoked represents a RoleRevoked event raised by the BuyBot contract.
type BuyBotRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BuyBot *BuyBotFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BuyBotRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BuyBot.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BuyBotRoleRevokedIterator{contract: _BuyBot.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BuyBot *BuyBotFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BuyBotRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BuyBot.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuyBotRoleRevoked)
				if err := _BuyBot.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BuyBot *BuyBotFilterer) ParseRoleRevoked(log types.Log) (*BuyBotRoleRevoked, error) {
	event := new(BuyBotRoleRevoked)
	if err := _BuyBot.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
