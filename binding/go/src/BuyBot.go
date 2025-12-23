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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"_initialDelay\",\"type\":\"uint48\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minOrderAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_interval\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_swapRouter\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"_maxTickSlippage\",\"type\":\"uint24\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"BUYER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NATIVE_COIN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"beginDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxMatchCount\",\"type\":\"uint256\"}],\"name\":\"buyMarket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"canBuyMarket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"canBuy\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"}],\"name\":\"changeDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelayIncreaseWait\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastBuyTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxTickSlippage\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minOrderAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"recipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollbackDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"contractIRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_interval\",\"type\":\"uint256\"}],\"name\":\"setInterval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_maxTickSlippage\",\"type\":\"uint24\"}],\"name\":\"setMaxTickSlippage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minOrderAmount\",\"type\":\"uint256\"}],\"name\":\"setMinOrderAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"setRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"setSwapPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_swapRouter\",\"type\":\"address\"}],\"name\":\"setSwapRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_swapToken\",\"type\":\"address\"}],\"name\":\"setSwapToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"swapPools\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapRouter\",\"outputs\":[{\"internalType\":\"contractISwapRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"uniswapFee\",\"type\":\"uint24\"}],\"name\":\"swapToQuote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminDelayChangeCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"effectSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminDelayChangeScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminTransferCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"acceptSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminTransferScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"before\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"current\",\"type\":\"uint256\"}],\"name\":\"IntervalSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"MarketBuyExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"before\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"current\",\"type\":\"uint24\"}],\"name\":\"MaxTickSlippageSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"before\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"current\",\"type\":\"uint256\"}],\"name\":\"MinOrderAmountSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"before\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"RecipientSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"SwapExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"SwapPoolSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"before\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"SwapRouterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"before\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"SwapTokenSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"name\":\"AccessControlEnforcedDefaultAdminDelay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccessControlEnforcedDefaultAdminRules\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"defaultAdmin\",\"type\":\"address\"}],\"name\":\"AccessControlInvalidDefaultAdmin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BuyBotETHTransferFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minOrderAmount\",\"type\":\"uint256\"}],\"name\":\"BuyBotInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"BuyBotInsufficientSwapBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BuyBotInsufficientWithdrawBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timeSinceLastBuy\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredInterval\",\"type\":\"uint256\"}],\"name\":\"BuyBotIntervalNotPassed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"BuyBotInvalidAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"}],\"name\":\"BuyBotInvalidBuyer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"BuyBotInvalidManager\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"BuyBotInvalidMinOrderAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"BuyBotInvalidPair\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"BuyBotInvalidRouter\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"BuyBotInvalidSwapRouter\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"name\":\"BuyBotInvalidTickSlippage\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BuyBotInvalidTokenAddresses\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BuyBotNoSwapToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"}],\"name\":\"BuyBotPoolNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"}]",
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
		"e879a40c": "maxTickSlippage()",
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
		"e1588ff9": "setMaxTickSlippage(uint24)",
		"a3b8ef04": "setMinOrderAmount(uint256)",
		"3bbed4a0": "setRecipient(address)",
		"0f3fffbf": "setSwapPool(address,address,address)",
		"41273657": "setSwapRouter(address)",
		"b851b7ca": "setSwapToken(address)",
		"01ffc9a7": "supportsInterface(bytes4)",
		"8bfe0df1": "swapPools(address,address)",
		"c31c9c07": "swapRouter()",
		"17f8d037": "swapToQuote(address,uint24)",
		"dc73e49c": "swapToken()",
		"f3fef3a3": "withdraw(address,uint256)",
		"f14210a6": "withdrawETH(uint256)",
	},
	Bin: "0x608060405234801561000f575f5ffd5b5060405161412c38038061412c83398101604081905261002e91610522565b89896001600160a01b03811661005e57604051636116401160e11b81525f60048201526024015b60405180910390fd5b600180546001600160d01b0316600160d01b65ffffffffffff8516021790556100875f8261037b565b50506001600355506001600160a01b0388166100c157604051631561f27b60e21b81526001600160a01b0389166004820152602401610055565b865f036100e4576040516328bfc81960e11b815260048101889052602401610055565b6001600160a01b03841661011657604051632a7409a960e01b81526001600160a01b0385166004820152602401610055565b6001600160a01b0383166101485760405163192178dd60e11b81526001600160a01b0384166004820152602401610055565b8062ffffff165f0361017457604051631421ff5b60e01b815262ffffff82166004820152602401610055565b600480546001600160a01b03199081166001600160a01b038b8116919091179092556005899055600688905560088054909116878316179055600980549184166001600160b81b031990921691909117600160a01b62ffffff8416021790556101ea5f51602061410c5f395f51905f525f6103ea565b6102015f5160206140ec5f395f51905f525f6103ea565b6102185f51602061410c5f395f51905f528a61037b565b506102305f51602061410c5f395f51905f528561037b565b506102485f5160206140ec5f395f51905f528a61037b565b506102605f5160206140ec5f395f51905f528461037b565b5060405187905f907fd6d62b78d21fc4f5151029ca37079020e8a8815c5a488b28ee7499d4c2a19e55908290a360405186905f907f3fb677206c0b314c404bae3da94bee6bda11375c9fe986f266fb033d6cddbbd7908290a36040516001600160a01b038616905f907fc1416b5cdab50a9fbc872236e1aa54566c6deb40024e63a4b1737ecacf09d6f9908290a3604080515f81526001600160a01b03841660208201527fc7324ad5feb4318ddf48817d97597d40855e50e83dc1b1b796bd5fb48dd9379f910160405180910390a1604080515f815262ffffff831660208201527fbf29e23461af22c9ea1b162c3d8fa888d73f300b51b7cfbabe9b5c1c8caa61d1910160405180910390a1505050505050505050506105e0565b5f826103d7575f6103946002546001600160a01b031690565b6001600160a01b0316146103bb57604051631fe1e13d60e11b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b0384161790555b6103e18383610416565b90505b92915050565b8161040857604051631fe1e13d60e11b815260040160405180910390fd5b61041282826104bd565b5050565b5f828152602081815260408083206001600160a01b038516845290915281205460ff166104b6575f838152602081815260408083206001600160a01b03861684529091529020805460ff1916600117905561046e3390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45060016103e4565b505f6103e4565b5f82815260208190526040808220600101805490849055905190918391839186917fbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff9190a4505050565b80516001600160a01b038116811461051d575f5ffd5b919050565b5f5f5f5f5f5f5f5f5f5f6101408b8d03121561053c575f5ffd5b8a5165ffffffffffff81168114610551575f5ffd5b995061055f60208c01610507565b985061056d60408c01610507565b60608c015160808d01519199509750955061058a60a08c01610507565b945061059860c08c01610507565b93506105a660e08c01610507565b92506105b56101008c01610507565b91506101208b015162ffffff811681146105cd575f5ffd5b809150509295989b9194979a5092959850565b613aff806105ed5f395ff3fe6080604052600436106102cf575f3560e01c80638da5cb5b1161017b578063d547741f116100d1578063ec87621c11610087578063f3fef3a311610062578063f3fef3a3146108eb578063f887ea401461090a578063f8b2cb4f14610936575f5ffd5b8063ec87621c14610884578063f14210a6146108b7578063f29f4d0b146108d6575f5ffd5b8063dc73e49c116100b7578063dc73e49c146107f2578063e1588ff91461081e578063e879a40c1461083d575f5ffd5b8063d547741f146107bf578063d602b9fd146107de575f5ffd5b8063a3b8ef0411610131578063cc8463c81161010c578063cc8463c81461073f578063cefc142914610753578063cf6eefb714610767575f5ffd5b8063a3b8ef04146106d5578063b851b7ca146106f4578063c31c9c0714610713575f5ffd5b8063947a36fb11610161578063947a36fb1461067a578063a1eda53c1461068f578063a217fddf146106c2575f5ffd5b80638da5cb5b1461061757806391d148541461062b575f5ffd5b806336568abe11610230578063649a5ec7116101e65780637a01a1da116101c15780637a01a1da1461056e57806384ef8ffc146105a15780638bfe0df1146105cb575f5ffd5b8063649a5ec71461050457806366d003ac1461052357806367ea88eb1461054f575f5ffd5b8063412736571161021657806341273657146104b157806346b62c4a146104d0578063634e93da146104e5575f5ffd5b806336568abe146104735780633bbed4a014610492575f5ffd5b806317f8d03711610285578063248a9ca31161026b578063248a9ca3146103f05780632aaa96281461041e5780632f2ff15d14610454575f5ffd5b806317f8d037146103a457806322a90082146103d1575f5ffd5b806304a41159116102b557806304a41159146103365780630aa6220b1461036f5780630f3fffbf14610385575f5ffd5b806301ffc9a7146102da578063022d63fb1461030e575f5ffd5b366102d657005b5f5ffd5b3480156102e5575f5ffd5b506102f96102f4366004613516565b610955565b60405190151581526020015b60405180910390f35b348015610319575f5ffd5b50620697805b60405165ffffffffffff9091168152602001610305565b348015610341575f5ffd5b5061034a600181565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610305565b34801561037a575f5ffd5b506103836109b0565b005b348015610390575f5ffd5b5061038361039f366004613576565b6109c5565b3480156103af575f5ffd5b506103c36103be3660046135d5565b610ae9565b604051908152602001610305565b3480156103dc575f5ffd5b506103836103eb366004613608565b611076565b3480156103fb575f5ffd5b506103c361040a366004613608565b5f9081526020819052604090206001015490565b348015610429575f5ffd5b5061043d61043836600461361f565b6110d9565b604080519215158352602083019190915201610305565b34801561045f575f5ffd5b5061038361046e366004613656565b6112b3565b34801561047e575f5ffd5b5061038361048d366004613656565b6112f8565b34801561049d575f5ffd5b506103836104ac366004613679565b6113fd565b3480156104bc575f5ffd5b506103836104cb366004613679565b61147e565b3480156104db575f5ffd5b506103c360055481565b3480156104f0575f5ffd5b506103836104ff366004613679565b611510565b34801561050f575f5ffd5b5061038361051e366004613694565b611523565b34801561052e575f5ffd5b5060085461034a9073ffffffffffffffffffffffffffffffffffffffff1681565b34801561055a575f5ffd5b506103836105693660046136b9565b611536565b348015610579575f5ffd5b506103c37ff8cd32ed93fc2f9fc78152a14807c9609af3d99c5fe4dc6b106a801aaddfe90e81565b3480156105ac575f5ffd5b5060025473ffffffffffffffffffffffffffffffffffffffff1661034a565b3480156105d6575f5ffd5b5061034a6105e536600461361f565b600b60209081525f928352604080842090915290825290205473ffffffffffffffffffffffffffffffffffffffff1681565b348015610622575f5ffd5b5061034a611b7a565b348015610636575f5ffd5b506102f9610645366004613656565b5f9182526020828152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b348015610685575f5ffd5b506103c360065481565b34801561069a575f5ffd5b506106a3611b9f565b6040805165ffffffffffff938416815292909116602083015201610305565b3480156106cd575f5ffd5b506103c35f81565b3480156106e0575f5ffd5b506103836106ef366004613608565b611c19565b3480156106ff575f5ffd5b5061038361070e366004613679565b611cb8565b34801561071e575f5ffd5b5060095461034a9073ffffffffffffffffffffffffffffffffffffffff1681565b34801561074a575f5ffd5b5061031f611d59565b34801561075e575f5ffd5b50610383611df6565b348015610772575f5ffd5b506001546040805173ffffffffffffffffffffffffffffffffffffffff831681527401000000000000000000000000000000000000000090920465ffffffffffff16602083015201610305565b3480156107ca575f5ffd5b506103836107d9366004613656565b611e52565b3480156107e9575f5ffd5b50610383611e93565b3480156107fd575f5ffd5b50600a5461034a9073ffffffffffffffffffffffffffffffffffffffff1681565b348015610829575f5ffd5b506103836108383660046136eb565b611ea5565b348015610848575f5ffd5b506009546108709074010000000000000000000000000000000000000000900462ffffff1681565b60405162ffffff9091168152602001610305565b34801561088f575f5ffd5b506103c37f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b0881565b3480156108c2575f5ffd5b506103836108d1366004613608565b611f9e565b3480156108e1575f5ffd5b506103c360075481565b3480156108f6575f5ffd5b50610383610905366004613704565b612110565b348015610915575f5ffd5b5060045461034a9073ffffffffffffffffffffffffffffffffffffffff1681565b348015610941575f5ffd5b506103c3610950366004613679565b6122cc565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f314987860000000000000000000000000000000000000000000000000000000014806109aa57506109aa8261239d565b92915050565b5f6109ba81612433565b6109c261243d565b50565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b086109ef81612433565b73ffffffffffffffffffffffffffffffffffffffff84161580610a26575073ffffffffffffffffffffffffffffffffffffffff8316155b15610a5d576040517f5159e7e000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8481165f818152600b60209081526040808320888616808552925280832080547fffffffffffffffffffffffff0000000000000000000000000000000000000000169588169586179055519092917fff59a584b1c3a296ea9baaa7fb8580f277f4763522ee6998dc8c8fcfad11934591a450505050565b5f610af2612449565b7ff8cd32ed93fc2f9fc78152a14807c9609af3d99c5fe4dc6b106a801aaddfe90e610b1c81612433565b73ffffffffffffffffffffffffffffffffffffffff8416610b86576040517f235aafe400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024015b60405180910390fd5b60095473ffffffffffffffffffffffffffffffffffffffff16610bd7576040517f5da5a13d0000000000000000000000000000000000000000000000000000000081525f6004820152602401610b7d565b600a5473ffffffffffffffffffffffffffffffffffffffff16610c26576040517fffb5baac00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8473ffffffffffffffffffffffffffffffffffffffff1663c3f909d46040518163ffffffff1660e01b8152600401606060405180830381865afa158015610c70573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c94919061372e565b8051600a546040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015292935090915f9173ffffffffffffffffffffffffffffffffffffffff16906370a0823190602401602060405180830381865afa158015610d07573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d2b91906137be565b9050805f03610d8857600a546040517f1a89d8d400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911660048201525f6024820152604401610b7d565b600a5473ffffffffffffffffffffffffffffffffffffffff9081165f908152600b6020908152604080832086851684529091529020541680610e1a57600a546040517fa06d7bd800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff91821660048201529084166024820152604401610b7d565b600a546009546040517fdd62ed3e00000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff9182166024820181905291909216918490839063dd62ed3e90604401602060405180830381865afa158015610e98573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ebc91906137be565b1015610f0357610f0373ffffffffffffffffffffffffffffffffffffffff8316827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff61248c565b600a545f90610f2990859073ffffffffffffffffffffffffffffffffffffffff166125ac565b6040805161010081018252600a5473ffffffffffffffffffffffffffffffffffffffff9081168252898116602083015262ffffff8e168284015230606083015242608083015260a082018990525f60c083015280841660e083015260095492517f414bf389000000000000000000000000000000000000000000000000000000008152939450909291169063414bf38990610fc89084906004016137d5565b6020604051808303815f875af1158015610fe4573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061100891906137be565b600a546040805189815260208101849052929c5073ffffffffffffffffffffffffffffffffffffffff8a8116939216917fdd36740e2a012d93061a0d99eaa9107860955de4e90027d3cf465a055026c407910160405180910390a35050505050505050506109aa6001600355565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b086110a081612433565b6006805490839055604051839082907f3fb677206c0b314c404bae3da94bee6bda11375c9fe986f266fb033d6cddbbd7905f90a3505050565b5f8073ffffffffffffffffffffffffffffffffffffffff841661110057505f9050806112ac565b73ffffffffffffffffffffffffffffffffffffffff83165f9081527ff3fa603c74bfe2a4719960e47343678c3dc690d2b27a2295acc6fc430833aaf9602052604090205460ff1661115557505f9050806112ac565b5f8473ffffffffffffffffffffffffffffffffffffffff1663c3f909d46040518163ffffffff1660e01b8152600401606060405180830381865afa15801561119f573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111c3919061372e565b80516040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015291925073ffffffffffffffffffffffffffffffffffffffff16906370a0823190602401602060405180830381865afa15801561122f573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061125391906137be565b9150600554821015611268575f9250506112ac565b5f60065411801561127a57505f600754115b156112a6575f6007544261128e91906138c7565b90506006548110156112a4575f935050506112ac565b505b60019250505b9250929050565b816112ea576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6112f482826127d8565b5050565b81158015611320575060025473ffffffffffffffffffffffffffffffffffffffff8281169116145b156113f35760015473ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900465ffffffffffff1681151580611374575065ffffffffffff8116155b8061138757504265ffffffffffff821610155b156113c8576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff82166004820152602401610b7d565b5050600180547fffffffffffff000000000000ffffffffffffffffffffffffffffffffffffffff1690555b6112f482826127fc565b5f61140781612433565b6008805473ffffffffffffffffffffffffffffffffffffffff8481167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907fc1416b5cdab50a9fbc872236e1aa54566c6deb40024e63a4b1737ecacf09d6f9905f90a3505050565b5f61148881612433565b6009805473ffffffffffffffffffffffffffffffffffffffff8481167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527fc7324ad5feb4318ddf48817d97597d40855e50e83dc1b1b796bd5fb48dd9379f91015b60405180910390a1505050565b5f61151a81612433565b6112f482612855565b5f61152d81612433565b6112f4826128d4565b61153e612449565b7ff8cd32ed93fc2f9fc78152a14807c9609af3d99c5fe4dc6b106a801aaddfe90e61156881612433565b73ffffffffffffffffffffffffffffffffffffffff84166115cd576040517f235aafe400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610b7d565b825f03611609576040517f9717f35f00000000000000000000000000000000000000000000000000000000815260048101849052602401610b7d565b5f60065411801561161b57505f600754115b1561167d575f6007544261162f91906138c7565b905060065481101561167b576006546040517f1e379906000000000000000000000000000000000000000000000000000000008152610b7d918391600401918252602082015260400190565b505b5f8473ffffffffffffffffffffffffffffffffffffffff1663c3f909d46040518163ffffffff1660e01b8152600401606060405180830381865afa1580156116c7573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906116eb919061372e565b805160208201516040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015292935090915f9073ffffffffffffffffffffffffffffffffffffffff8416906370a0823190602401602060405180830381865afa158015611761573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061178591906137be565b90506005548710156117d1576005546040517f013fafe2000000000000000000000000000000000000000000000000000000008152610b7d918991600401918252602082015260400190565b80871115611815576040517f013fafe20000000000000000000000000000000000000000000000000000000081526004810182905260248101889052604401610b7d565b600480546040517fdd62ed3e000000000000000000000000000000000000000000000000000000008152309281019290925273ffffffffffffffffffffffffffffffffffffffff9081166024830181905291899186169063dd62ed3e90604401602060405180830381865afa158015611890573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906118b491906137be565b10156118fb576118fb73ffffffffffffffffffffffffffffffffffffffff8516827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff61248c565b600480546040517f1e92008400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8c811693820193909352602481018b9052604481018a9052911690631e920084906064015f604051808303815f87803b158015611975575f5ffd5b505af1158015611987573d5f5f3e3d5ffd5b5050604080518b815233602082015273ffffffffffffffffffffffffffffffffffffffff808816945088811693508d16917fc7b4d815bd0a3b8348577971cb55a9334c80a0b3ccdfe287b55b5a6320d8480a910160405180910390a44260075560085473ffffffffffffffffffffffffffffffffffffffff1615611b6557478015611aa5576008546040515f9173ffffffffffffffffffffffffffffffffffffffff169083908381818185875af1925050503d805f8114611a63576040519150601f19603f3d011682016040523d82523d5f602084013e611a68565b606091505b5050905080611aa3576040517f87fd589900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201525f9073ffffffffffffffffffffffffffffffffffffffff8616906370a0823190602401602060405180830381865afa158015611b0f573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611b3391906137be565b90508015611b6257600854611b629073ffffffffffffffffffffffffffffffffffffffff878116911683612943565b50505b505050505050611b756001600355565b505050565b5f611b9a60025473ffffffffffffffffffffffffffffffffffffffff1690565b905090565b6002545f907a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015158015611be157504265ffffffffffff821610155b611bec575f5f611c11565b60025474010000000000000000000000000000000000000000900465ffffffffffff16815b915091509091565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b08611c4381612433565b815f03611c7f576040517f517f903200000000000000000000000000000000000000000000000000000000815260048101839052602401610b7d565b6005805490839055604051839082907fd6d62b78d21fc4f5151029ca37079020e8a8815c5a488b28ee7499d4c2a19e55905f90a3505050565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b08611ce281612433565b600a805473ffffffffffffffffffffffffffffffffffffffff8481167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f36b1a7df795121e2bd67c2b04fc89f611b7b21ee10f345226e13bc1b2ebda30d905f90a3505050565b6002545f907a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015158015611d9a57504265ffffffffffff8216105b611dcc576001547a010000000000000000000000000000000000000000000000000000900465ffffffffffff16611df0565b60025474010000000000000000000000000000000000000000900465ffffffffffff165b91505090565b60015473ffffffffffffffffffffffffffffffffffffffff16338114611e4a576040517fc22c8022000000000000000000000000000000000000000000000000000000008152336004820152602401610b7d565b6109c2612981565b81611e89576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6112f48282612a72565b5f611e9d81612433565b6109c2612a96565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b08611ecf81612433565b8162ffffff165f03611f14576040517f1421ff5b00000000000000000000000000000000000000000000000000000000815262ffffff83166004820152602401610b7d565b6009805462ffffff848116740100000000000000000000000000000000000000008181027fffffffffffffffffff000000ffffffffffffffffffffffffffffffffffffffff85161790945560408051949093049091168084526020840191909152917fbf29e23461af22c9ea1b162c3d8fa888d73f300b51b7cfbabe9b5c1c8caa61d19101611503565b611fa6612449565b5f611fb081612433565b475f8315611fbe5783611fc0565b815b905081811115611ffc576040517f26f4246a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f61201c60025473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff16826040515f6040518083038185875af1925050503d805f8114612070576040519150601f19603f3d011682016040523d82523d5f602084013e612075565b606091505b50509050806120b0576040517f87fd589900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60025460405183815273ffffffffffffffffffffffffffffffffffffffff909116906001907fd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb9060200160405180910390a3505050506109c26001600355565b612118612449565b5f61212281612433565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015283905f9073ffffffffffffffffffffffffffffffffffffffff8316906370a0823190602401602060405180830381865afa15801561218e573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906121b291906137be565b90505f84156121c157846121c3565b815b9050818111156121ff576040517f26f4246a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61223f61222160025473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff85169083612943565b60025473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff167fd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb836040516122b691815260200190565b60405180910390a3505050506112f46001600355565b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff73ffffffffffffffffffffffffffffffffffffffff831601612311575047919050565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff8316906370a0823190602401602060405180830381865afa158015612379573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109aa91906137be565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b0000000000000000000000000000000000000000000000000000000014806109aa57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316146109aa565b6109c28133612aa0565b6124475f5f612b25565b565b600260035403612485576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600355565b6040805173ffffffffffffffffffffffffffffffffffffffff8416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f095ea7b3000000000000000000000000000000000000000000000000000000001790526125188482612c7e565b6125a65760405173ffffffffffffffffffffffffffffffffffffffff84811660248301525f604483015261259c91869182169063095ea7b3906064015b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050612cd4565b6125a68482612cd4565b50505050565b5f5f8373ffffffffffffffffffffffffffffffffffffffff16633850c7bd6040518163ffffffff1660e01b815260040160e060405180830381865afa1580156125f7573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061261b91906138eb565b50505050509150505f8473ffffffffffffffffffffffffffffffffffffffff16630dfe16816040518163ffffffff1660e01b8152600401602060405180830381865afa15801561266d573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906126919190613987565b905073ffffffffffffffffffffffffffffffffffffffff848116908216145f8115612736576009546126e09074010000000000000000000000000000000000000000900462ffffff16856139a2565b90507ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff27618600282900b121561273157507ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff276185b6127c4565b6009546127609074010000000000000000000000000000000000000000900462ffffff16856139e3565b905061278b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff27618613a24565b60020b8160020b13156127c4576127c17ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff27618613a24565b90505b6127cd81612d73565b979650505050505050565b5f828152602081905260409020600101546127f281612433565b6125a68383613103565b73ffffffffffffffffffffffffffffffffffffffff8116331461284b576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611b7582826131c8565b5f61285e611d59565b61286742613229565b6128719190613a60565b905061287d8282613278565b60405165ffffffffffff8216815273ffffffffffffffffffffffffffffffffffffffff8316907f3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed69060200160405180910390a25050565b5f6128de82613313565b6128e742613229565b6128f19190613a60565b90506128fd8282612b25565b6040805165ffffffffffff8085168252831660208201527ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b910160405180910390a15050565b60405173ffffffffffffffffffffffffffffffffffffffff838116602483015260448201839052611b7591859182169063a9059cbb90606401612555565b60015473ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900465ffffffffffff168015806129d157504265ffffffffffff821610155b15612a12576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff82166004820152602401610b7d565b612a3a5f612a3560025473ffffffffffffffffffffffffffffffffffffffff1690565b6131c8565b50612a455f83613103565b5050600180547fffffffffffff000000000000000000000000000000000000000000000000000016905550565b5f82815260208190526040902060010154612a8c81612433565b6125a683836131c8565b6124475f5f613278565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff166112f4576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8216600482015260248101839052604401610b7d565b6002547a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015612bf9574265ffffffffffff82161015612bd0576002546001805479ffffffffffffffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000090920465ffffffffffff167a01000000000000000000000000000000000000000000000000000002919091179055612bf9565b6040517f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5905f90a15b506002805473ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000065ffffffffffff9485160279ffffffffffffffffffffffffffffffffffffffffffffffffffff16177a0100000000000000000000000000000000000000000000000000009290931691909102919091179055565b5f5f5f5f60205f8651602088015f8a5af192503d91505f519050828015612cca57508115612caf5780600114612cca565b5f8673ffffffffffffffffffffffffffffffffffffffff163b115b9695505050505050565b5f5f60205f8451602086015f885af180612cf3576040513d5f823e3d81fd5b50505f513d91508115612d0a578060011415612d24565b73ffffffffffffffffffffffffffffffffffffffff84163b155b156125a6576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610b7d565b5f5f5f8360020b12612d88578260020b612d8f565b8260020b5f035b9050620d89e8811115612dfe576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600160248201527f54000000000000000000000000000000000000000000000000000000000000006044820152606401610b7d565b5f816001165f03612e2057700100000000000000000000000000000000612e32565b6ffffcb933bd6fad37aa2d162d1a5940015b70ffffffffffffffffffffffffffffffffff1690506002821615612e66576ffff97272373d413259a46990580e213a0260801c5b6004821615612e85576ffff2e50f5f656932ef12357cf3c7fdcc0260801c5b6008821615612ea4576fffe5caca7e10e4e61c3624eaa0941cd00260801c5b6010821615612ec3576fffcb9843d60f6159c9db58835c9266440260801c5b6020821615612ee2576fff973b41fa98c081472e6896dfb254c00260801c5b6040821615612f01576fff2ea16466c96a3843ec78b326b528610260801c5b6080821615612f20576ffe5dee046a99a2a811c461f1969c30530260801c5b610100821615612f40576ffcbe86c7900a88aedcffc83b479aa3a40260801c5b610200821615612f60576ff987a7253ac413176f2b074cf7815e540260801c5b610400821615612f80576ff3392b0822b70005940c7a398e4b70f30260801c5b610800821615612fa0576fe7159475a2c29b7443b29c7fa6e889d90260801c5b611000821615612fc0576fd097f3bdfd2022b8845ad8f792aa58250260801c5b612000821615612fe0576fa9f746462d870fdf8a65dc1f90e061e50260801c5b614000821615613000576f70d869a156d2a1b890bb3df62baf32f70260801c5b618000821615613020576f31be135f97d08fd981231505542fcfa60260801c5b62010000821615613041576f09aa508b5b7a84e1c677de54f3e99bc90260801c5b62020000821615613061576e5d6af8dedb81196699c329225ee6040260801c5b62040000821615613080576d2216e584f5fa1ea926041bedfe980260801c5b6208000082161561309d576b048a170391f7dc42444e8fa20260801c5b5f8460020b13156130db57807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff816130d7576130d7613a7e565b0490505b6401000000008106156130ef5760016130f1565b5f5b60ff16602082901c0192505050919050565b5f826131b7575f61312960025473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614613176576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84161790555b6131c18383613364565b9392505050565b5f821580156131f1575060025473ffffffffffffffffffffffffffffffffffffffff8381169116145b1561321f57600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001690555b6131c1838361345d565b5f65ffffffffffff821115613274576040517f6dfcc6500000000000000000000000000000000000000000000000000000000081526030600482015260248101839052604401610b7d565b5090565b600180547401000000000000000000000000000000000000000065ffffffffffff84811682027fffffffffffff0000000000000000000000000000000000000000000000000000841673ffffffffffffffffffffffffffffffffffffffff881617179093559004168015611b75576040517f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109905f90a1505050565b5f5f61331d611d59565b90508065ffffffffffff168365ffffffffffff1611613345576133408382613aab565b6131c1565b6131c165ffffffffffff8416620697805f8282188284100282186131c1565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff16613456575f8381526020818152604080832073ffffffffffffffffffffffffffffffffffffffff86168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556133f43390565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45060016109aa565b505f6109aa565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff1615613456575f8381526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8616808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45060016109aa565b5f60208284031215613526575f5ffd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146131c1575f5ffd5b73ffffffffffffffffffffffffffffffffffffffff811681146109c2575f5ffd5b5f5f5f60608486031215613588575f5ffd5b833561359381613555565b925060208401356135a381613555565b915060408401356135b381613555565b809150509250925092565b803562ffffff811681146135d0575f5ffd5b919050565b5f5f604083850312156135e6575f5ffd5b82356135f181613555565b91506135ff602084016135be565b90509250929050565b5f60208284031215613618575f5ffd5b5035919050565b5f5f60408385031215613630575f5ffd5b823561363b81613555565b9150602083013561364b81613555565b809150509250929050565b5f5f60408385031215613667575f5ffd5b82359150602083013561364b81613555565b5f60208284031215613689575f5ffd5b81356131c181613555565b5f602082840312156136a4575f5ffd5b813565ffffffffffff811681146131c1575f5ffd5b5f5f5f606084860312156136cb575f5ffd5b83356136d681613555565b95602085013595506040909401359392505050565b5f602082840312156136fb575f5ffd5b6131c1826135be565b5f5f60408385031215613715575f5ffd5b823561372081613555565b946020939093013593505050565b5f606082840312801561373f575f5ffd5b506040516060810167ffffffffffffffff81118282101715613788577f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604052825161379681613555565b815260208301516137a681613555565b60208201526040928301519281019290925250919050565b5f602082840312156137ce575f5ffd5b5051919050565b5f6101008201905073ffffffffffffffffffffffffffffffffffffffff835116825273ffffffffffffffffffffffffffffffffffffffff602084015116602083015262ffffff6040840151166040830152606083015161384d606084018273ffffffffffffffffffffffffffffffffffffffff169052565b506080830151608083015260a083015160a083015260c083015160c083015260e083015161389360e084018273ffffffffffffffffffffffffffffffffffffffff169052565b5092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b818103818111156109aa576109aa61389a565b805161ffff811681146135d0575f5ffd5b5f5f5f5f5f5f5f60e0888a031215613901575f5ffd5b875161390c81613555565b8097505060208801518060020b8114613923575f5ffd5b9550613931604089016138da565b945061393f606089016138da565b935061394d608089016138da565b925060a088015160ff81168114613962575f5ffd5b60c08901519092508015158114613977575f5ffd5b8091505092959891949750929550565b5f60208284031215613997575f5ffd5b81516131c181613555565b600282810b9082900b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8000008112627fffff821317156109aa576109aa61389a565b600281810b9083900b01627fffff81137fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff800000821217156109aa576109aa61389a565b5f8160020b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8000008103613a5857613a5861389a565b5f0392915050565b65ffffffffffff81811683821601908111156109aa576109aa61389a565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b65ffffffffffff82811682821603908111156109aa576109aa61389a56fea26469706673582212204d799d2e96bddae2261717527223c0fb78c848c9a91e90b190d27ccac3a1300664736f6c634300081c0033241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b08f8cd32ed93fc2f9fc78152a14807c9609af3d99c5fe4dc6b106a801aaddfe90e",
}

// BuyBotABI is the input ABI used to generate the binding from.
// Deprecated: Use BuyBotMetaData.ABI instead.
var BuyBotABI = BuyBotMetaData.ABI

// BuyBotBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const BuyBotBinRuntime = "6080604052600436106102cf575f3560e01c80638da5cb5b1161017b578063d547741f116100d1578063ec87621c11610087578063f3fef3a311610062578063f3fef3a3146108eb578063f887ea401461090a578063f8b2cb4f14610936575f5ffd5b8063ec87621c14610884578063f14210a6146108b7578063f29f4d0b146108d6575f5ffd5b8063dc73e49c116100b7578063dc73e49c146107f2578063e1588ff91461081e578063e879a40c1461083d575f5ffd5b8063d547741f146107bf578063d602b9fd146107de575f5ffd5b8063a3b8ef0411610131578063cc8463c81161010c578063cc8463c81461073f578063cefc142914610753578063cf6eefb714610767575f5ffd5b8063a3b8ef04146106d5578063b851b7ca146106f4578063c31c9c0714610713575f5ffd5b8063947a36fb11610161578063947a36fb1461067a578063a1eda53c1461068f578063a217fddf146106c2575f5ffd5b80638da5cb5b1461061757806391d148541461062b575f5ffd5b806336568abe11610230578063649a5ec7116101e65780637a01a1da116101c15780637a01a1da1461056e57806384ef8ffc146105a15780638bfe0df1146105cb575f5ffd5b8063649a5ec71461050457806366d003ac1461052357806367ea88eb1461054f575f5ffd5b8063412736571161021657806341273657146104b157806346b62c4a146104d0578063634e93da146104e5575f5ffd5b806336568abe146104735780633bbed4a014610492575f5ffd5b806317f8d03711610285578063248a9ca31161026b578063248a9ca3146103f05780632aaa96281461041e5780632f2ff15d14610454575f5ffd5b806317f8d037146103a457806322a90082146103d1575f5ffd5b806304a41159116102b557806304a41159146103365780630aa6220b1461036f5780630f3fffbf14610385575f5ffd5b806301ffc9a7146102da578063022d63fb1461030e575f5ffd5b366102d657005b5f5ffd5b3480156102e5575f5ffd5b506102f96102f4366004613516565b610955565b60405190151581526020015b60405180910390f35b348015610319575f5ffd5b50620697805b60405165ffffffffffff9091168152602001610305565b348015610341575f5ffd5b5061034a600181565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610305565b34801561037a575f5ffd5b506103836109b0565b005b348015610390575f5ffd5b5061038361039f366004613576565b6109c5565b3480156103af575f5ffd5b506103c36103be3660046135d5565b610ae9565b604051908152602001610305565b3480156103dc575f5ffd5b506103836103eb366004613608565b611076565b3480156103fb575f5ffd5b506103c361040a366004613608565b5f9081526020819052604090206001015490565b348015610429575f5ffd5b5061043d61043836600461361f565b6110d9565b604080519215158352602083019190915201610305565b34801561045f575f5ffd5b5061038361046e366004613656565b6112b3565b34801561047e575f5ffd5b5061038361048d366004613656565b6112f8565b34801561049d575f5ffd5b506103836104ac366004613679565b6113fd565b3480156104bc575f5ffd5b506103836104cb366004613679565b61147e565b3480156104db575f5ffd5b506103c360055481565b3480156104f0575f5ffd5b506103836104ff366004613679565b611510565b34801561050f575f5ffd5b5061038361051e366004613694565b611523565b34801561052e575f5ffd5b5060085461034a9073ffffffffffffffffffffffffffffffffffffffff1681565b34801561055a575f5ffd5b506103836105693660046136b9565b611536565b348015610579575f5ffd5b506103c37ff8cd32ed93fc2f9fc78152a14807c9609af3d99c5fe4dc6b106a801aaddfe90e81565b3480156105ac575f5ffd5b5060025473ffffffffffffffffffffffffffffffffffffffff1661034a565b3480156105d6575f5ffd5b5061034a6105e536600461361f565b600b60209081525f928352604080842090915290825290205473ffffffffffffffffffffffffffffffffffffffff1681565b348015610622575f5ffd5b5061034a611b7a565b348015610636575f5ffd5b506102f9610645366004613656565b5f9182526020828152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b348015610685575f5ffd5b506103c360065481565b34801561069a575f5ffd5b506106a3611b9f565b6040805165ffffffffffff938416815292909116602083015201610305565b3480156106cd575f5ffd5b506103c35f81565b3480156106e0575f5ffd5b506103836106ef366004613608565b611c19565b3480156106ff575f5ffd5b5061038361070e366004613679565b611cb8565b34801561071e575f5ffd5b5060095461034a9073ffffffffffffffffffffffffffffffffffffffff1681565b34801561074a575f5ffd5b5061031f611d59565b34801561075e575f5ffd5b50610383611df6565b348015610772575f5ffd5b506001546040805173ffffffffffffffffffffffffffffffffffffffff831681527401000000000000000000000000000000000000000090920465ffffffffffff16602083015201610305565b3480156107ca575f5ffd5b506103836107d9366004613656565b611e52565b3480156107e9575f5ffd5b50610383611e93565b3480156107fd575f5ffd5b50600a5461034a9073ffffffffffffffffffffffffffffffffffffffff1681565b348015610829575f5ffd5b506103836108383660046136eb565b611ea5565b348015610848575f5ffd5b506009546108709074010000000000000000000000000000000000000000900462ffffff1681565b60405162ffffff9091168152602001610305565b34801561088f575f5ffd5b506103c37f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b0881565b3480156108c2575f5ffd5b506103836108d1366004613608565b611f9e565b3480156108e1575f5ffd5b506103c360075481565b3480156108f6575f5ffd5b50610383610905366004613704565b612110565b348015610915575f5ffd5b5060045461034a9073ffffffffffffffffffffffffffffffffffffffff1681565b348015610941575f5ffd5b506103c3610950366004613679565b6122cc565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f314987860000000000000000000000000000000000000000000000000000000014806109aa57506109aa8261239d565b92915050565b5f6109ba81612433565b6109c261243d565b50565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b086109ef81612433565b73ffffffffffffffffffffffffffffffffffffffff84161580610a26575073ffffffffffffffffffffffffffffffffffffffff8316155b15610a5d576040517f5159e7e000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8481165f818152600b60209081526040808320888616808552925280832080547fffffffffffffffffffffffff0000000000000000000000000000000000000000169588169586179055519092917fff59a584b1c3a296ea9baaa7fb8580f277f4763522ee6998dc8c8fcfad11934591a450505050565b5f610af2612449565b7ff8cd32ed93fc2f9fc78152a14807c9609af3d99c5fe4dc6b106a801aaddfe90e610b1c81612433565b73ffffffffffffffffffffffffffffffffffffffff8416610b86576040517f235aafe400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024015b60405180910390fd5b60095473ffffffffffffffffffffffffffffffffffffffff16610bd7576040517f5da5a13d0000000000000000000000000000000000000000000000000000000081525f6004820152602401610b7d565b600a5473ffffffffffffffffffffffffffffffffffffffff16610c26576040517fffb5baac00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8473ffffffffffffffffffffffffffffffffffffffff1663c3f909d46040518163ffffffff1660e01b8152600401606060405180830381865afa158015610c70573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c94919061372e565b8051600a546040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015292935090915f9173ffffffffffffffffffffffffffffffffffffffff16906370a0823190602401602060405180830381865afa158015610d07573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d2b91906137be565b9050805f03610d8857600a546040517f1a89d8d400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911660048201525f6024820152604401610b7d565b600a5473ffffffffffffffffffffffffffffffffffffffff9081165f908152600b6020908152604080832086851684529091529020541680610e1a57600a546040517fa06d7bd800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff91821660048201529084166024820152604401610b7d565b600a546009546040517fdd62ed3e00000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff9182166024820181905291909216918490839063dd62ed3e90604401602060405180830381865afa158015610e98573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ebc91906137be565b1015610f0357610f0373ffffffffffffffffffffffffffffffffffffffff8316827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff61248c565b600a545f90610f2990859073ffffffffffffffffffffffffffffffffffffffff166125ac565b6040805161010081018252600a5473ffffffffffffffffffffffffffffffffffffffff9081168252898116602083015262ffffff8e168284015230606083015242608083015260a082018990525f60c083015280841660e083015260095492517f414bf389000000000000000000000000000000000000000000000000000000008152939450909291169063414bf38990610fc89084906004016137d5565b6020604051808303815f875af1158015610fe4573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061100891906137be565b600a546040805189815260208101849052929c5073ffffffffffffffffffffffffffffffffffffffff8a8116939216917fdd36740e2a012d93061a0d99eaa9107860955de4e90027d3cf465a055026c407910160405180910390a35050505050505050506109aa6001600355565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b086110a081612433565b6006805490839055604051839082907f3fb677206c0b314c404bae3da94bee6bda11375c9fe986f266fb033d6cddbbd7905f90a3505050565b5f8073ffffffffffffffffffffffffffffffffffffffff841661110057505f9050806112ac565b73ffffffffffffffffffffffffffffffffffffffff83165f9081527ff3fa603c74bfe2a4719960e47343678c3dc690d2b27a2295acc6fc430833aaf9602052604090205460ff1661115557505f9050806112ac565b5f8473ffffffffffffffffffffffffffffffffffffffff1663c3f909d46040518163ffffffff1660e01b8152600401606060405180830381865afa15801561119f573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111c3919061372e565b80516040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015291925073ffffffffffffffffffffffffffffffffffffffff16906370a0823190602401602060405180830381865afa15801561122f573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061125391906137be565b9150600554821015611268575f9250506112ac565b5f60065411801561127a57505f600754115b156112a6575f6007544261128e91906138c7565b90506006548110156112a4575f935050506112ac565b505b60019250505b9250929050565b816112ea576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6112f482826127d8565b5050565b81158015611320575060025473ffffffffffffffffffffffffffffffffffffffff8281169116145b156113f35760015473ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900465ffffffffffff1681151580611374575065ffffffffffff8116155b8061138757504265ffffffffffff821610155b156113c8576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff82166004820152602401610b7d565b5050600180547fffffffffffff000000000000ffffffffffffffffffffffffffffffffffffffff1690555b6112f482826127fc565b5f61140781612433565b6008805473ffffffffffffffffffffffffffffffffffffffff8481167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907fc1416b5cdab50a9fbc872236e1aa54566c6deb40024e63a4b1737ecacf09d6f9905f90a3505050565b5f61148881612433565b6009805473ffffffffffffffffffffffffffffffffffffffff8481167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527fc7324ad5feb4318ddf48817d97597d40855e50e83dc1b1b796bd5fb48dd9379f91015b60405180910390a1505050565b5f61151a81612433565b6112f482612855565b5f61152d81612433565b6112f4826128d4565b61153e612449565b7ff8cd32ed93fc2f9fc78152a14807c9609af3d99c5fe4dc6b106a801aaddfe90e61156881612433565b73ffffffffffffffffffffffffffffffffffffffff84166115cd576040517f235aafe400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610b7d565b825f03611609576040517f9717f35f00000000000000000000000000000000000000000000000000000000815260048101849052602401610b7d565b5f60065411801561161b57505f600754115b1561167d575f6007544261162f91906138c7565b905060065481101561167b576006546040517f1e379906000000000000000000000000000000000000000000000000000000008152610b7d918391600401918252602082015260400190565b505b5f8473ffffffffffffffffffffffffffffffffffffffff1663c3f909d46040518163ffffffff1660e01b8152600401606060405180830381865afa1580156116c7573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906116eb919061372e565b805160208201516040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015292935090915f9073ffffffffffffffffffffffffffffffffffffffff8416906370a0823190602401602060405180830381865afa158015611761573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061178591906137be565b90506005548710156117d1576005546040517f013fafe2000000000000000000000000000000000000000000000000000000008152610b7d918991600401918252602082015260400190565b80871115611815576040517f013fafe20000000000000000000000000000000000000000000000000000000081526004810182905260248101889052604401610b7d565b600480546040517fdd62ed3e000000000000000000000000000000000000000000000000000000008152309281019290925273ffffffffffffffffffffffffffffffffffffffff9081166024830181905291899186169063dd62ed3e90604401602060405180830381865afa158015611890573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906118b491906137be565b10156118fb576118fb73ffffffffffffffffffffffffffffffffffffffff8516827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff61248c565b600480546040517f1e92008400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8c811693820193909352602481018b9052604481018a9052911690631e920084906064015f604051808303815f87803b158015611975575f5ffd5b505af1158015611987573d5f5f3e3d5ffd5b5050604080518b815233602082015273ffffffffffffffffffffffffffffffffffffffff808816945088811693508d16917fc7b4d815bd0a3b8348577971cb55a9334c80a0b3ccdfe287b55b5a6320d8480a910160405180910390a44260075560085473ffffffffffffffffffffffffffffffffffffffff1615611b6557478015611aa5576008546040515f9173ffffffffffffffffffffffffffffffffffffffff169083908381818185875af1925050503d805f8114611a63576040519150601f19603f3d011682016040523d82523d5f602084013e611a68565b606091505b5050905080611aa3576040517f87fd589900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201525f9073ffffffffffffffffffffffffffffffffffffffff8616906370a0823190602401602060405180830381865afa158015611b0f573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611b3391906137be565b90508015611b6257600854611b629073ffffffffffffffffffffffffffffffffffffffff878116911683612943565b50505b505050505050611b756001600355565b505050565b5f611b9a60025473ffffffffffffffffffffffffffffffffffffffff1690565b905090565b6002545f907a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015158015611be157504265ffffffffffff821610155b611bec575f5f611c11565b60025474010000000000000000000000000000000000000000900465ffffffffffff16815b915091509091565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b08611c4381612433565b815f03611c7f576040517f517f903200000000000000000000000000000000000000000000000000000000815260048101839052602401610b7d565b6005805490839055604051839082907fd6d62b78d21fc4f5151029ca37079020e8a8815c5a488b28ee7499d4c2a19e55905f90a3505050565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b08611ce281612433565b600a805473ffffffffffffffffffffffffffffffffffffffff8481167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f36b1a7df795121e2bd67c2b04fc89f611b7b21ee10f345226e13bc1b2ebda30d905f90a3505050565b6002545f907a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015158015611d9a57504265ffffffffffff8216105b611dcc576001547a010000000000000000000000000000000000000000000000000000900465ffffffffffff16611df0565b60025474010000000000000000000000000000000000000000900465ffffffffffff165b91505090565b60015473ffffffffffffffffffffffffffffffffffffffff16338114611e4a576040517fc22c8022000000000000000000000000000000000000000000000000000000008152336004820152602401610b7d565b6109c2612981565b81611e89576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6112f48282612a72565b5f611e9d81612433565b6109c2612a96565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b08611ecf81612433565b8162ffffff165f03611f14576040517f1421ff5b00000000000000000000000000000000000000000000000000000000815262ffffff83166004820152602401610b7d565b6009805462ffffff848116740100000000000000000000000000000000000000008181027fffffffffffffffffff000000ffffffffffffffffffffffffffffffffffffffff85161790945560408051949093049091168084526020840191909152917fbf29e23461af22c9ea1b162c3d8fa888d73f300b51b7cfbabe9b5c1c8caa61d19101611503565b611fa6612449565b5f611fb081612433565b475f8315611fbe5783611fc0565b815b905081811115611ffc576040517f26f4246a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f61201c60025473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff16826040515f6040518083038185875af1925050503d805f8114612070576040519150601f19603f3d011682016040523d82523d5f602084013e612075565b606091505b50509050806120b0576040517f87fd589900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60025460405183815273ffffffffffffffffffffffffffffffffffffffff909116906001907fd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb9060200160405180910390a3505050506109c26001600355565b612118612449565b5f61212281612433565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015283905f9073ffffffffffffffffffffffffffffffffffffffff8316906370a0823190602401602060405180830381865afa15801561218e573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906121b291906137be565b90505f84156121c157846121c3565b815b9050818111156121ff576040517f26f4246a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61223f61222160025473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff85169083612943565b60025473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff167fd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb836040516122b691815260200190565b60405180910390a3505050506112f46001600355565b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff73ffffffffffffffffffffffffffffffffffffffff831601612311575047919050565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff8316906370a0823190602401602060405180830381865afa158015612379573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109aa91906137be565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b0000000000000000000000000000000000000000000000000000000014806109aa57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316146109aa565b6109c28133612aa0565b6124475f5f612b25565b565b600260035403612485576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600355565b6040805173ffffffffffffffffffffffffffffffffffffffff8416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f095ea7b3000000000000000000000000000000000000000000000000000000001790526125188482612c7e565b6125a65760405173ffffffffffffffffffffffffffffffffffffffff84811660248301525f604483015261259c91869182169063095ea7b3906064015b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050612cd4565b6125a68482612cd4565b50505050565b5f5f8373ffffffffffffffffffffffffffffffffffffffff16633850c7bd6040518163ffffffff1660e01b815260040160e060405180830381865afa1580156125f7573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061261b91906138eb565b50505050509150505f8473ffffffffffffffffffffffffffffffffffffffff16630dfe16816040518163ffffffff1660e01b8152600401602060405180830381865afa15801561266d573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906126919190613987565b905073ffffffffffffffffffffffffffffffffffffffff848116908216145f8115612736576009546126e09074010000000000000000000000000000000000000000900462ffffff16856139a2565b90507ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff27618600282900b121561273157507ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff276185b6127c4565b6009546127609074010000000000000000000000000000000000000000900462ffffff16856139e3565b905061278b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff27618613a24565b60020b8160020b13156127c4576127c17ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff27618613a24565b90505b6127cd81612d73565b979650505050505050565b5f828152602081905260409020600101546127f281612433565b6125a68383613103565b73ffffffffffffffffffffffffffffffffffffffff8116331461284b576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611b7582826131c8565b5f61285e611d59565b61286742613229565b6128719190613a60565b905061287d8282613278565b60405165ffffffffffff8216815273ffffffffffffffffffffffffffffffffffffffff8316907f3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed69060200160405180910390a25050565b5f6128de82613313565b6128e742613229565b6128f19190613a60565b90506128fd8282612b25565b6040805165ffffffffffff8085168252831660208201527ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b910160405180910390a15050565b60405173ffffffffffffffffffffffffffffffffffffffff838116602483015260448201839052611b7591859182169063a9059cbb90606401612555565b60015473ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900465ffffffffffff168015806129d157504265ffffffffffff821610155b15612a12576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff82166004820152602401610b7d565b612a3a5f612a3560025473ffffffffffffffffffffffffffffffffffffffff1690565b6131c8565b50612a455f83613103565b5050600180547fffffffffffff000000000000000000000000000000000000000000000000000016905550565b5f82815260208190526040902060010154612a8c81612433565b6125a683836131c8565b6124475f5f613278565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff166112f4576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8216600482015260248101839052604401610b7d565b6002547a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015612bf9574265ffffffffffff82161015612bd0576002546001805479ffffffffffffffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000090920465ffffffffffff167a01000000000000000000000000000000000000000000000000000002919091179055612bf9565b6040517f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5905f90a15b506002805473ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000065ffffffffffff9485160279ffffffffffffffffffffffffffffffffffffffffffffffffffff16177a0100000000000000000000000000000000000000000000000000009290931691909102919091179055565b5f5f5f5f60205f8651602088015f8a5af192503d91505f519050828015612cca57508115612caf5780600114612cca565b5f8673ffffffffffffffffffffffffffffffffffffffff163b115b9695505050505050565b5f5f60205f8451602086015f885af180612cf3576040513d5f823e3d81fd5b50505f513d91508115612d0a578060011415612d24565b73ffffffffffffffffffffffffffffffffffffffff84163b155b156125a6576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610b7d565b5f5f5f8360020b12612d88578260020b612d8f565b8260020b5f035b9050620d89e8811115612dfe576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600160248201527f54000000000000000000000000000000000000000000000000000000000000006044820152606401610b7d565b5f816001165f03612e2057700100000000000000000000000000000000612e32565b6ffffcb933bd6fad37aa2d162d1a5940015b70ffffffffffffffffffffffffffffffffff1690506002821615612e66576ffff97272373d413259a46990580e213a0260801c5b6004821615612e85576ffff2e50f5f656932ef12357cf3c7fdcc0260801c5b6008821615612ea4576fffe5caca7e10e4e61c3624eaa0941cd00260801c5b6010821615612ec3576fffcb9843d60f6159c9db58835c9266440260801c5b6020821615612ee2576fff973b41fa98c081472e6896dfb254c00260801c5b6040821615612f01576fff2ea16466c96a3843ec78b326b528610260801c5b6080821615612f20576ffe5dee046a99a2a811c461f1969c30530260801c5b610100821615612f40576ffcbe86c7900a88aedcffc83b479aa3a40260801c5b610200821615612f60576ff987a7253ac413176f2b074cf7815e540260801c5b610400821615612f80576ff3392b0822b70005940c7a398e4b70f30260801c5b610800821615612fa0576fe7159475a2c29b7443b29c7fa6e889d90260801c5b611000821615612fc0576fd097f3bdfd2022b8845ad8f792aa58250260801c5b612000821615612fe0576fa9f746462d870fdf8a65dc1f90e061e50260801c5b614000821615613000576f70d869a156d2a1b890bb3df62baf32f70260801c5b618000821615613020576f31be135f97d08fd981231505542fcfa60260801c5b62010000821615613041576f09aa508b5b7a84e1c677de54f3e99bc90260801c5b62020000821615613061576e5d6af8dedb81196699c329225ee6040260801c5b62040000821615613080576d2216e584f5fa1ea926041bedfe980260801c5b6208000082161561309d576b048a170391f7dc42444e8fa20260801c5b5f8460020b13156130db57807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff816130d7576130d7613a7e565b0490505b6401000000008106156130ef5760016130f1565b5f5b60ff16602082901c0192505050919050565b5f826131b7575f61312960025473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614613176576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84161790555b6131c18383613364565b9392505050565b5f821580156131f1575060025473ffffffffffffffffffffffffffffffffffffffff8381169116145b1561321f57600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001690555b6131c1838361345d565b5f65ffffffffffff821115613274576040517f6dfcc6500000000000000000000000000000000000000000000000000000000081526030600482015260248101839052604401610b7d565b5090565b600180547401000000000000000000000000000000000000000065ffffffffffff84811682027fffffffffffff0000000000000000000000000000000000000000000000000000841673ffffffffffffffffffffffffffffffffffffffff881617179093559004168015611b75576040517f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109905f90a1505050565b5f5f61331d611d59565b90508065ffffffffffff168365ffffffffffff1611613345576133408382613aab565b6131c1565b6131c165ffffffffffff8416620697805f8282188284100282186131c1565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff16613456575f8381526020818152604080832073ffffffffffffffffffffffffffffffffffffffff86168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556133f43390565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45060016109aa565b505f6109aa565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff1615613456575f8381526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8616808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45060016109aa565b5f60208284031215613526575f5ffd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146131c1575f5ffd5b73ffffffffffffffffffffffffffffffffffffffff811681146109c2575f5ffd5b5f5f5f60608486031215613588575f5ffd5b833561359381613555565b925060208401356135a381613555565b915060408401356135b381613555565b809150509250925092565b803562ffffff811681146135d0575f5ffd5b919050565b5f5f604083850312156135e6575f5ffd5b82356135f181613555565b91506135ff602084016135be565b90509250929050565b5f60208284031215613618575f5ffd5b5035919050565b5f5f60408385031215613630575f5ffd5b823561363b81613555565b9150602083013561364b81613555565b809150509250929050565b5f5f60408385031215613667575f5ffd5b82359150602083013561364b81613555565b5f60208284031215613689575f5ffd5b81356131c181613555565b5f602082840312156136a4575f5ffd5b813565ffffffffffff811681146131c1575f5ffd5b5f5f5f606084860312156136cb575f5ffd5b83356136d681613555565b95602085013595506040909401359392505050565b5f602082840312156136fb575f5ffd5b6131c1826135be565b5f5f60408385031215613715575f5ffd5b823561372081613555565b946020939093013593505050565b5f606082840312801561373f575f5ffd5b506040516060810167ffffffffffffffff81118282101715613788577f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604052825161379681613555565b815260208301516137a681613555565b60208201526040928301519281019290925250919050565b5f602082840312156137ce575f5ffd5b5051919050565b5f6101008201905073ffffffffffffffffffffffffffffffffffffffff835116825273ffffffffffffffffffffffffffffffffffffffff602084015116602083015262ffffff6040840151166040830152606083015161384d606084018273ffffffffffffffffffffffffffffffffffffffff169052565b506080830151608083015260a083015160a083015260c083015160c083015260e083015161389360e084018273ffffffffffffffffffffffffffffffffffffffff169052565b5092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b818103818111156109aa576109aa61389a565b805161ffff811681146135d0575f5ffd5b5f5f5f5f5f5f5f60e0888a031215613901575f5ffd5b875161390c81613555565b8097505060208801518060020b8114613923575f5ffd5b9550613931604089016138da565b945061393f606089016138da565b935061394d608089016138da565b925060a088015160ff81168114613962575f5ffd5b60c08901519092508015158114613977575f5ffd5b8091505092959891949750929550565b5f60208284031215613997575f5ffd5b81516131c181613555565b600282810b9082900b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8000008112627fffff821317156109aa576109aa61389a565b600281810b9083900b01627fffff81137fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff800000821217156109aa576109aa61389a565b5f8160020b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8000008103613a5857613a5861389a565b5f0392915050565b65ffffffffffff81811683821601908111156109aa576109aa61389a565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b65ffffffffffff82811682821603908111156109aa576109aa61389a56fea26469706673582212204d799d2e96bddae2261717527223c0fb78c848c9a91e90b190d27ccac3a1300664736f6c634300081c0033"

// Deprecated: Use BuyBotMetaData.Sigs instead.
// BuyBotFuncSigs maps the 4-byte function signature to its string representation.
var BuyBotFuncSigs = BuyBotMetaData.Sigs

// BuyBotBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BuyBotMetaData.Bin instead.
var BuyBotBin = BuyBotMetaData.Bin

// DeployBuyBot deploys a new Ethereum contract, binding an instance of BuyBot to it.
func DeployBuyBot(auth *bind.TransactOpts, backend bind.ContractBackend, _initialDelay *big.Int, _owner common.Address, _router common.Address, _minOrderAmount *big.Int, _interval *big.Int, _recipient common.Address, _buyer common.Address, _manager common.Address, _swapRouter common.Address, _maxTickSlippage *big.Int) (common.Address, *types.Transaction, *BuyBot, error) {
	parsed, err := BuyBotMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BuyBotBin), backend, _initialDelay, _owner, _router, _minOrderAmount, _interval, _recipient, _buyer, _manager, _swapRouter, _maxTickSlippage)
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

// MaxTickSlippage is a free data retrieval call binding the contract method 0xe879a40c.
//
// Solidity: function maxTickSlippage() view returns(uint24)
func (_BuyBot *BuyBotCaller) MaxTickSlippage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BuyBot.contract.Call(opts, &out, "maxTickSlippage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxTickSlippage is a free data retrieval call binding the contract method 0xe879a40c.
//
// Solidity: function maxTickSlippage() view returns(uint24)
func (_BuyBot *BuyBotSession) MaxTickSlippage() (*big.Int, error) {
	return _BuyBot.Contract.MaxTickSlippage(&_BuyBot.CallOpts)
}

// MaxTickSlippage is a free data retrieval call binding the contract method 0xe879a40c.
//
// Solidity: function maxTickSlippage() view returns(uint24)
func (_BuyBot *BuyBotCallerSession) MaxTickSlippage() (*big.Int, error) {
	return _BuyBot.Contract.MaxTickSlippage(&_BuyBot.CallOpts)
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

// SwapPools is a free data retrieval call binding the contract method 0x8bfe0df1.
//
// Solidity: function swapPools(address , address ) view returns(address)
func (_BuyBot *BuyBotCaller) SwapPools(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (common.Address, error) {
	var out []interface{}
	err := _BuyBot.contract.Call(opts, &out, "swapPools", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SwapPools is a free data retrieval call binding the contract method 0x8bfe0df1.
//
// Solidity: function swapPools(address , address ) view returns(address)
func (_BuyBot *BuyBotSession) SwapPools(arg0 common.Address, arg1 common.Address) (common.Address, error) {
	return _BuyBot.Contract.SwapPools(&_BuyBot.CallOpts, arg0, arg1)
}

// SwapPools is a free data retrieval call binding the contract method 0x8bfe0df1.
//
// Solidity: function swapPools(address , address ) view returns(address)
func (_BuyBot *BuyBotCallerSession) SwapPools(arg0 common.Address, arg1 common.Address) (common.Address, error) {
	return _BuyBot.Contract.SwapPools(&_BuyBot.CallOpts, arg0, arg1)
}

// SwapRouter is a free data retrieval call binding the contract method 0xc31c9c07.
//
// Solidity: function swapRouter() view returns(address)
func (_BuyBot *BuyBotCaller) SwapRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BuyBot.contract.Call(opts, &out, "swapRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SwapRouter is a free data retrieval call binding the contract method 0xc31c9c07.
//
// Solidity: function swapRouter() view returns(address)
func (_BuyBot *BuyBotSession) SwapRouter() (common.Address, error) {
	return _BuyBot.Contract.SwapRouter(&_BuyBot.CallOpts)
}

// SwapRouter is a free data retrieval call binding the contract method 0xc31c9c07.
//
// Solidity: function swapRouter() view returns(address)
func (_BuyBot *BuyBotCallerSession) SwapRouter() (common.Address, error) {
	return _BuyBot.Contract.SwapRouter(&_BuyBot.CallOpts)
}

// SwapToken is a free data retrieval call binding the contract method 0xdc73e49c.
//
// Solidity: function swapToken() view returns(address)
func (_BuyBot *BuyBotCaller) SwapToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BuyBot.contract.Call(opts, &out, "swapToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SwapToken is a free data retrieval call binding the contract method 0xdc73e49c.
//
// Solidity: function swapToken() view returns(address)
func (_BuyBot *BuyBotSession) SwapToken() (common.Address, error) {
	return _BuyBot.Contract.SwapToken(&_BuyBot.CallOpts)
}

// SwapToken is a free data retrieval call binding the contract method 0xdc73e49c.
//
// Solidity: function swapToken() view returns(address)
func (_BuyBot *BuyBotCallerSession) SwapToken() (common.Address, error) {
	return _BuyBot.Contract.SwapToken(&_BuyBot.CallOpts)
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

// SetMaxTickSlippage is a paid mutator transaction binding the contract method 0xe1588ff9.
//
// Solidity: function setMaxTickSlippage(uint24 _maxTickSlippage) returns()
func (_BuyBot *BuyBotTransactor) SetMaxTickSlippage(opts *bind.TransactOpts, _maxTickSlippage *big.Int) (*types.Transaction, error) {
	return _BuyBot.contract.Transact(opts, "setMaxTickSlippage", _maxTickSlippage)
}

// SetMaxTickSlippage is a paid mutator transaction binding the contract method 0xe1588ff9.
//
// Solidity: function setMaxTickSlippage(uint24 _maxTickSlippage) returns()
func (_BuyBot *BuyBotSession) SetMaxTickSlippage(_maxTickSlippage *big.Int) (*types.Transaction, error) {
	return _BuyBot.Contract.SetMaxTickSlippage(&_BuyBot.TransactOpts, _maxTickSlippage)
}

// SetMaxTickSlippage is a paid mutator transaction binding the contract method 0xe1588ff9.
//
// Solidity: function setMaxTickSlippage(uint24 _maxTickSlippage) returns()
func (_BuyBot *BuyBotTransactorSession) SetMaxTickSlippage(_maxTickSlippage *big.Int) (*types.Transaction, error) {
	return _BuyBot.Contract.SetMaxTickSlippage(&_BuyBot.TransactOpts, _maxTickSlippage)
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

// SetSwapPool is a paid mutator transaction binding the contract method 0x0f3fffbf.
//
// Solidity: function setSwapPool(address tokenIn, address tokenOut, address pool) returns()
func (_BuyBot *BuyBotTransactor) SetSwapPool(opts *bind.TransactOpts, tokenIn common.Address, tokenOut common.Address, pool common.Address) (*types.Transaction, error) {
	return _BuyBot.contract.Transact(opts, "setSwapPool", tokenIn, tokenOut, pool)
}

// SetSwapPool is a paid mutator transaction binding the contract method 0x0f3fffbf.
//
// Solidity: function setSwapPool(address tokenIn, address tokenOut, address pool) returns()
func (_BuyBot *BuyBotSession) SetSwapPool(tokenIn common.Address, tokenOut common.Address, pool common.Address) (*types.Transaction, error) {
	return _BuyBot.Contract.SetSwapPool(&_BuyBot.TransactOpts, tokenIn, tokenOut, pool)
}

// SetSwapPool is a paid mutator transaction binding the contract method 0x0f3fffbf.
//
// Solidity: function setSwapPool(address tokenIn, address tokenOut, address pool) returns()
func (_BuyBot *BuyBotTransactorSession) SetSwapPool(tokenIn common.Address, tokenOut common.Address, pool common.Address) (*types.Transaction, error) {
	return _BuyBot.Contract.SetSwapPool(&_BuyBot.TransactOpts, tokenIn, tokenOut, pool)
}

// SetSwapRouter is a paid mutator transaction binding the contract method 0x41273657.
//
// Solidity: function setSwapRouter(address _swapRouter) returns()
func (_BuyBot *BuyBotTransactor) SetSwapRouter(opts *bind.TransactOpts, _swapRouter common.Address) (*types.Transaction, error) {
	return _BuyBot.contract.Transact(opts, "setSwapRouter", _swapRouter)
}

// SetSwapRouter is a paid mutator transaction binding the contract method 0x41273657.
//
// Solidity: function setSwapRouter(address _swapRouter) returns()
func (_BuyBot *BuyBotSession) SetSwapRouter(_swapRouter common.Address) (*types.Transaction, error) {
	return _BuyBot.Contract.SetSwapRouter(&_BuyBot.TransactOpts, _swapRouter)
}

// SetSwapRouter is a paid mutator transaction binding the contract method 0x41273657.
//
// Solidity: function setSwapRouter(address _swapRouter) returns()
func (_BuyBot *BuyBotTransactorSession) SetSwapRouter(_swapRouter common.Address) (*types.Transaction, error) {
	return _BuyBot.Contract.SetSwapRouter(&_BuyBot.TransactOpts, _swapRouter)
}

// SetSwapToken is a paid mutator transaction binding the contract method 0xb851b7ca.
//
// Solidity: function setSwapToken(address _swapToken) returns()
func (_BuyBot *BuyBotTransactor) SetSwapToken(opts *bind.TransactOpts, _swapToken common.Address) (*types.Transaction, error) {
	return _BuyBot.contract.Transact(opts, "setSwapToken", _swapToken)
}

// SetSwapToken is a paid mutator transaction binding the contract method 0xb851b7ca.
//
// Solidity: function setSwapToken(address _swapToken) returns()
func (_BuyBot *BuyBotSession) SetSwapToken(_swapToken common.Address) (*types.Transaction, error) {
	return _BuyBot.Contract.SetSwapToken(&_BuyBot.TransactOpts, _swapToken)
}

// SetSwapToken is a paid mutator transaction binding the contract method 0xb851b7ca.
//
// Solidity: function setSwapToken(address _swapToken) returns()
func (_BuyBot *BuyBotTransactorSession) SetSwapToken(_swapToken common.Address) (*types.Transaction, error) {
	return _BuyBot.Contract.SetSwapToken(&_BuyBot.TransactOpts, _swapToken)
}

// SwapToQuote is a paid mutator transaction binding the contract method 0x17f8d037.
//
// Solidity: function swapToQuote(address pair, uint24 uniswapFee) returns(uint256 amountOut)
func (_BuyBot *BuyBotTransactor) SwapToQuote(opts *bind.TransactOpts, pair common.Address, uniswapFee *big.Int) (*types.Transaction, error) {
	return _BuyBot.contract.Transact(opts, "swapToQuote", pair, uniswapFee)
}

// SwapToQuote is a paid mutator transaction binding the contract method 0x17f8d037.
//
// Solidity: function swapToQuote(address pair, uint24 uniswapFee) returns(uint256 amountOut)
func (_BuyBot *BuyBotSession) SwapToQuote(pair common.Address, uniswapFee *big.Int) (*types.Transaction, error) {
	return _BuyBot.Contract.SwapToQuote(&_BuyBot.TransactOpts, pair, uniswapFee)
}

// SwapToQuote is a paid mutator transaction binding the contract method 0x17f8d037.
//
// Solidity: function swapToQuote(address pair, uint24 uniswapFee) returns(uint256 amountOut)
func (_BuyBot *BuyBotTransactorSession) SwapToQuote(pair common.Address, uniswapFee *big.Int) (*types.Transaction, error) {
	return _BuyBot.Contract.SwapToQuote(&_BuyBot.TransactOpts, pair, uniswapFee)
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

// BuyBotMaxTickSlippageSetIterator is returned from FilterMaxTickSlippageSet and is used to iterate over the raw logs and unpacked data for MaxTickSlippageSet events raised by the BuyBot contract.
type BuyBotMaxTickSlippageSetIterator struct {
	Event *BuyBotMaxTickSlippageSet // Event containing the contract specifics and raw log

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
func (it *BuyBotMaxTickSlippageSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuyBotMaxTickSlippageSet)
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
		it.Event = new(BuyBotMaxTickSlippageSet)
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
func (it *BuyBotMaxTickSlippageSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuyBotMaxTickSlippageSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuyBotMaxTickSlippageSet represents a MaxTickSlippageSet event raised by the BuyBot contract.
type BuyBotMaxTickSlippageSet struct {
	Before  *big.Int
	Current *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMaxTickSlippageSet is a free log retrieval operation binding the contract event 0xbf29e23461af22c9ea1b162c3d8fa888d73f300b51b7cfbabe9b5c1c8caa61d1.
//
// Solidity: event MaxTickSlippageSet(uint24 before, uint24 current)
func (_BuyBot *BuyBotFilterer) FilterMaxTickSlippageSet(opts *bind.FilterOpts) (*BuyBotMaxTickSlippageSetIterator, error) {

	logs, sub, err := _BuyBot.contract.FilterLogs(opts, "MaxTickSlippageSet")
	if err != nil {
		return nil, err
	}
	return &BuyBotMaxTickSlippageSetIterator{contract: _BuyBot.contract, event: "MaxTickSlippageSet", logs: logs, sub: sub}, nil
}

// WatchMaxTickSlippageSet is a free log subscription operation binding the contract event 0xbf29e23461af22c9ea1b162c3d8fa888d73f300b51b7cfbabe9b5c1c8caa61d1.
//
// Solidity: event MaxTickSlippageSet(uint24 before, uint24 current)
func (_BuyBot *BuyBotFilterer) WatchMaxTickSlippageSet(opts *bind.WatchOpts, sink chan<- *BuyBotMaxTickSlippageSet) (event.Subscription, error) {

	logs, sub, err := _BuyBot.contract.WatchLogs(opts, "MaxTickSlippageSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuyBotMaxTickSlippageSet)
				if err := _BuyBot.contract.UnpackLog(event, "MaxTickSlippageSet", log); err != nil {
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

// ParseMaxTickSlippageSet is a log parse operation binding the contract event 0xbf29e23461af22c9ea1b162c3d8fa888d73f300b51b7cfbabe9b5c1c8caa61d1.
//
// Solidity: event MaxTickSlippageSet(uint24 before, uint24 current)
func (_BuyBot *BuyBotFilterer) ParseMaxTickSlippageSet(log types.Log) (*BuyBotMaxTickSlippageSet, error) {
	event := new(BuyBotMaxTickSlippageSet)
	if err := _BuyBot.contract.UnpackLog(event, "MaxTickSlippageSet", log); err != nil {
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

// BuyBotSwapExecutedIterator is returned from FilterSwapExecuted and is used to iterate over the raw logs and unpacked data for SwapExecuted events raised by the BuyBot contract.
type BuyBotSwapExecutedIterator struct {
	Event *BuyBotSwapExecuted // Event containing the contract specifics and raw log

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
func (it *BuyBotSwapExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuyBotSwapExecuted)
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
		it.Event = new(BuyBotSwapExecuted)
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
func (it *BuyBotSwapExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuyBotSwapExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuyBotSwapExecuted represents a SwapExecuted event raised by the BuyBot contract.
type BuyBotSwapExecuted struct {
	TokenIn   common.Address
	TokenOut  common.Address
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSwapExecuted is a free log retrieval operation binding the contract event 0xdd36740e2a012d93061a0d99eaa9107860955de4e90027d3cf465a055026c407.
//
// Solidity: event SwapExecuted(address indexed tokenIn, address indexed tokenOut, uint256 amountIn, uint256 amountOut)
func (_BuyBot *BuyBotFilterer) FilterSwapExecuted(opts *bind.FilterOpts, tokenIn []common.Address, tokenOut []common.Address) (*BuyBotSwapExecutedIterator, error) {

	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}
	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}

	logs, sub, err := _BuyBot.contract.FilterLogs(opts, "SwapExecuted", tokenInRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return &BuyBotSwapExecutedIterator{contract: _BuyBot.contract, event: "SwapExecuted", logs: logs, sub: sub}, nil
}

// WatchSwapExecuted is a free log subscription operation binding the contract event 0xdd36740e2a012d93061a0d99eaa9107860955de4e90027d3cf465a055026c407.
//
// Solidity: event SwapExecuted(address indexed tokenIn, address indexed tokenOut, uint256 amountIn, uint256 amountOut)
func (_BuyBot *BuyBotFilterer) WatchSwapExecuted(opts *bind.WatchOpts, sink chan<- *BuyBotSwapExecuted, tokenIn []common.Address, tokenOut []common.Address) (event.Subscription, error) {

	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}
	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}

	logs, sub, err := _BuyBot.contract.WatchLogs(opts, "SwapExecuted", tokenInRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuyBotSwapExecuted)
				if err := _BuyBot.contract.UnpackLog(event, "SwapExecuted", log); err != nil {
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

// ParseSwapExecuted is a log parse operation binding the contract event 0xdd36740e2a012d93061a0d99eaa9107860955de4e90027d3cf465a055026c407.
//
// Solidity: event SwapExecuted(address indexed tokenIn, address indexed tokenOut, uint256 amountIn, uint256 amountOut)
func (_BuyBot *BuyBotFilterer) ParseSwapExecuted(log types.Log) (*BuyBotSwapExecuted, error) {
	event := new(BuyBotSwapExecuted)
	if err := _BuyBot.contract.UnpackLog(event, "SwapExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BuyBotSwapPoolSetIterator is returned from FilterSwapPoolSet and is used to iterate over the raw logs and unpacked data for SwapPoolSet events raised by the BuyBot contract.
type BuyBotSwapPoolSetIterator struct {
	Event *BuyBotSwapPoolSet // Event containing the contract specifics and raw log

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
func (it *BuyBotSwapPoolSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuyBotSwapPoolSet)
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
		it.Event = new(BuyBotSwapPoolSet)
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
func (it *BuyBotSwapPoolSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuyBotSwapPoolSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuyBotSwapPoolSet represents a SwapPoolSet event raised by the BuyBot contract.
type BuyBotSwapPoolSet struct {
	TokenIn  common.Address
	TokenOut common.Address
	Pool     common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSwapPoolSet is a free log retrieval operation binding the contract event 0xff59a584b1c3a296ea9baaa7fb8580f277f4763522ee6998dc8c8fcfad119345.
//
// Solidity: event SwapPoolSet(address indexed tokenIn, address indexed tokenOut, address indexed pool)
func (_BuyBot *BuyBotFilterer) FilterSwapPoolSet(opts *bind.FilterOpts, tokenIn []common.Address, tokenOut []common.Address, pool []common.Address) (*BuyBotSwapPoolSetIterator, error) {

	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}
	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _BuyBot.contract.FilterLogs(opts, "SwapPoolSet", tokenInRule, tokenOutRule, poolRule)
	if err != nil {
		return nil, err
	}
	return &BuyBotSwapPoolSetIterator{contract: _BuyBot.contract, event: "SwapPoolSet", logs: logs, sub: sub}, nil
}

// WatchSwapPoolSet is a free log subscription operation binding the contract event 0xff59a584b1c3a296ea9baaa7fb8580f277f4763522ee6998dc8c8fcfad119345.
//
// Solidity: event SwapPoolSet(address indexed tokenIn, address indexed tokenOut, address indexed pool)
func (_BuyBot *BuyBotFilterer) WatchSwapPoolSet(opts *bind.WatchOpts, sink chan<- *BuyBotSwapPoolSet, tokenIn []common.Address, tokenOut []common.Address, pool []common.Address) (event.Subscription, error) {

	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}
	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _BuyBot.contract.WatchLogs(opts, "SwapPoolSet", tokenInRule, tokenOutRule, poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuyBotSwapPoolSet)
				if err := _BuyBot.contract.UnpackLog(event, "SwapPoolSet", log); err != nil {
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

// ParseSwapPoolSet is a log parse operation binding the contract event 0xff59a584b1c3a296ea9baaa7fb8580f277f4763522ee6998dc8c8fcfad119345.
//
// Solidity: event SwapPoolSet(address indexed tokenIn, address indexed tokenOut, address indexed pool)
func (_BuyBot *BuyBotFilterer) ParseSwapPoolSet(log types.Log) (*BuyBotSwapPoolSet, error) {
	event := new(BuyBotSwapPoolSet)
	if err := _BuyBot.contract.UnpackLog(event, "SwapPoolSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BuyBotSwapRouterSetIterator is returned from FilterSwapRouterSet and is used to iterate over the raw logs and unpacked data for SwapRouterSet events raised by the BuyBot contract.
type BuyBotSwapRouterSetIterator struct {
	Event *BuyBotSwapRouterSet // Event containing the contract specifics and raw log

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
func (it *BuyBotSwapRouterSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuyBotSwapRouterSet)
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
		it.Event = new(BuyBotSwapRouterSet)
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
func (it *BuyBotSwapRouterSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuyBotSwapRouterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuyBotSwapRouterSet represents a SwapRouterSet event raised by the BuyBot contract.
type BuyBotSwapRouterSet struct {
	Before  common.Address
	Current common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSwapRouterSet is a free log retrieval operation binding the contract event 0xc7324ad5feb4318ddf48817d97597d40855e50e83dc1b1b796bd5fb48dd9379f.
//
// Solidity: event SwapRouterSet(address before, address current)
func (_BuyBot *BuyBotFilterer) FilterSwapRouterSet(opts *bind.FilterOpts) (*BuyBotSwapRouterSetIterator, error) {

	logs, sub, err := _BuyBot.contract.FilterLogs(opts, "SwapRouterSet")
	if err != nil {
		return nil, err
	}
	return &BuyBotSwapRouterSetIterator{contract: _BuyBot.contract, event: "SwapRouterSet", logs: logs, sub: sub}, nil
}

// WatchSwapRouterSet is a free log subscription operation binding the contract event 0xc7324ad5feb4318ddf48817d97597d40855e50e83dc1b1b796bd5fb48dd9379f.
//
// Solidity: event SwapRouterSet(address before, address current)
func (_BuyBot *BuyBotFilterer) WatchSwapRouterSet(opts *bind.WatchOpts, sink chan<- *BuyBotSwapRouterSet) (event.Subscription, error) {

	logs, sub, err := _BuyBot.contract.WatchLogs(opts, "SwapRouterSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuyBotSwapRouterSet)
				if err := _BuyBot.contract.UnpackLog(event, "SwapRouterSet", log); err != nil {
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

// ParseSwapRouterSet is a log parse operation binding the contract event 0xc7324ad5feb4318ddf48817d97597d40855e50e83dc1b1b796bd5fb48dd9379f.
//
// Solidity: event SwapRouterSet(address before, address current)
func (_BuyBot *BuyBotFilterer) ParseSwapRouterSet(log types.Log) (*BuyBotSwapRouterSet, error) {
	event := new(BuyBotSwapRouterSet)
	if err := _BuyBot.contract.UnpackLog(event, "SwapRouterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BuyBotSwapTokenSetIterator is returned from FilterSwapTokenSet and is used to iterate over the raw logs and unpacked data for SwapTokenSet events raised by the BuyBot contract.
type BuyBotSwapTokenSetIterator struct {
	Event *BuyBotSwapTokenSet // Event containing the contract specifics and raw log

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
func (it *BuyBotSwapTokenSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuyBotSwapTokenSet)
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
		it.Event = new(BuyBotSwapTokenSet)
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
func (it *BuyBotSwapTokenSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuyBotSwapTokenSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuyBotSwapTokenSet represents a SwapTokenSet event raised by the BuyBot contract.
type BuyBotSwapTokenSet struct {
	Before  common.Address
	Current common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSwapTokenSet is a free log retrieval operation binding the contract event 0x36b1a7df795121e2bd67c2b04fc89f611b7b21ee10f345226e13bc1b2ebda30d.
//
// Solidity: event SwapTokenSet(address indexed before, address indexed current)
func (_BuyBot *BuyBotFilterer) FilterSwapTokenSet(opts *bind.FilterOpts, before []common.Address, current []common.Address) (*BuyBotSwapTokenSetIterator, error) {

	var beforeRule []interface{}
	for _, beforeItem := range before {
		beforeRule = append(beforeRule, beforeItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _BuyBot.contract.FilterLogs(opts, "SwapTokenSet", beforeRule, currentRule)
	if err != nil {
		return nil, err
	}
	return &BuyBotSwapTokenSetIterator{contract: _BuyBot.contract, event: "SwapTokenSet", logs: logs, sub: sub}, nil
}

// WatchSwapTokenSet is a free log subscription operation binding the contract event 0x36b1a7df795121e2bd67c2b04fc89f611b7b21ee10f345226e13bc1b2ebda30d.
//
// Solidity: event SwapTokenSet(address indexed before, address indexed current)
func (_BuyBot *BuyBotFilterer) WatchSwapTokenSet(opts *bind.WatchOpts, sink chan<- *BuyBotSwapTokenSet, before []common.Address, current []common.Address) (event.Subscription, error) {

	var beforeRule []interface{}
	for _, beforeItem := range before {
		beforeRule = append(beforeRule, beforeItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _BuyBot.contract.WatchLogs(opts, "SwapTokenSet", beforeRule, currentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuyBotSwapTokenSet)
				if err := _BuyBot.contract.UnpackLog(event, "SwapTokenSet", log); err != nil {
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

// ParseSwapTokenSet is a log parse operation binding the contract event 0x36b1a7df795121e2bd67c2b04fc89f611b7b21ee10f345226e13bc1b2ebda30d.
//
// Solidity: event SwapTokenSet(address indexed before, address indexed current)
func (_BuyBot *BuyBotFilterer) ParseSwapTokenSet(log types.Log) (*BuyBotSwapTokenSet, error) {
	event := new(BuyBotSwapTokenSet)
	if err := _BuyBot.contract.UnpackLog(event, "SwapTokenSet", log); err != nil {
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
