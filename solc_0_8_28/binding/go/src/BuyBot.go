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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"_initialDelay\",\"type\":\"uint48\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minOrderAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_interval\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_swapRouter\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"BUYER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NATIVE_COIN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"beginDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxMatchCount\",\"type\":\"uint256\"}],\"name\":\"buyMarket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"canBuyMarket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"canBuy\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"}],\"name\":\"changeDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelayIncreaseWait\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastBuyTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minOrderAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"recipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollbackDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"contractIRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_interval\",\"type\":\"uint256\"}],\"name\":\"setInterval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minOrderAmount\",\"type\":\"uint256\"}],\"name\":\"setMinOrderAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"setRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"setSwapPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_swapRouter\",\"type\":\"address\"}],\"name\":\"setSwapRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_swapToken\",\"type\":\"address\"}],\"name\":\"setSwapToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"swapPools\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapRouter\",\"outputs\":[{\"internalType\":\"contractISwapRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"uniswapFee\",\"type\":\"uint24\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"}],\"name\":\"swapToQuote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminDelayChangeCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"effectSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminDelayChangeScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminTransferCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"acceptSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminTransferScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"before\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"current\",\"type\":\"uint256\"}],\"name\":\"IntervalSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"MarketBuyExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"before\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"current\",\"type\":\"uint256\"}],\"name\":\"MinOrderAmountSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"before\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"RecipientSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"SwapExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"SwapPoolSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"before\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"SwapRouterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"before\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"SwapTokenSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"name\":\"AccessControlEnforcedDefaultAdminDelay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccessControlEnforcedDefaultAdminRules\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"defaultAdmin\",\"type\":\"address\"}],\"name\":\"AccessControlInvalidDefaultAdmin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oldValue\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"BuyBotAddressNotChanged\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BuyBotETHTransferFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minOrderAmount\",\"type\":\"uint256\"}],\"name\":\"BuyBotInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"BuyBotInsufficientSwapBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BuyBotInsufficientWithdrawBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timeSinceLastBuy\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredInterval\",\"type\":\"uint256\"}],\"name\":\"BuyBotIntervalNotPassed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"BuyBotInvalidAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"}],\"name\":\"BuyBotInvalidBuyer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"BuyBotInvalidManager\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"BuyBotInvalidMinOrderAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"BuyBotInvalidPair\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"BuyBotInvalidPool\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"}],\"name\":\"BuyBotInvalidPoolTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"BuyBotInvalidRouter\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"BuyBotInvalidSwapRouter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BuyBotInvalidTokenAddresses\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BuyBotNoSwapToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"oldValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"BuyBotNotChanged\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"}],\"name\":\"BuyBotPoolNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"}]",
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
		"0f3fffbf": "setSwapPool(address,address,address)",
		"41273657": "setSwapRouter(address)",
		"b851b7ca": "setSwapToken(address)",
		"01ffc9a7": "supportsInterface(bytes4)",
		"8bfe0df1": "swapPools(address,address)",
		"c31c9c07": "swapRouter()",
		"5e1c17ed": "swapToQuote(address,uint24,uint256)",
		"dc73e49c": "swapToken()",
		"f3fef3a3": "withdraw(address,uint256)",
		"f14210a6": "withdrawETH(uint256)",
	},
	Bin: "0x608060405234801561000f575f5ffd5b50604051613b0f380380613b0f83398101604081905261002e916103fc565b88886001600160a01b03811661005e57604051636116401160e11b81525f60048201526024015b60405180910390fd5b600180546001600160d01b0316600160d01b65ffffffffffff8516021790556100875f826102cb565b50506001600355506001600160a01b0387166100c157604051631561f27b60e21b81526001600160a01b0388166004820152602401610055565b855f036100e4576040516328bfc81960e11b815260048101879052602401610055565b6001600160a01b03831661011657604051632a7409a960e01b81526001600160a01b0384166004820152602401610055565b6001600160a01b0382166101485760405163192178dd60e11b81526001600160a01b0383166004820152602401610055565b600480546001600160a01b03808a166001600160a01b031992831617909255600588905560068790556008805487841690831617905560098054928416929091169190911790556101a65f516020613aef5f395f51905f52896102cb565b506101be5f516020613aef5f395f51905f52846102cb565b506101d65f516020613acf5f395f51905f52896102cb565b506101ee5f516020613acf5f395f51905f52836102cb565b5060405186905f907fd6d62b78d21fc4f5151029ca37079020e8a8815c5a488b28ee7499d4c2a19e55908290a360405185905f907f3fb677206c0b314c404bae3da94bee6bda11375c9fe986f266fb033d6cddbbd7908290a36040516001600160a01b038516905f907fc1416b5cdab50a9fbc872236e1aa54566c6deb40024e63a4b1737ecacf09d6f9908290a3604080515f81526001600160a01b03831660208201527fc7324ad5feb4318ddf48817d97597d40855e50e83dc1b1b796bd5fb48dd9379f910160405180910390a150505050505050505061049d565b5f82610327575f6102e46002546001600160a01b031690565b6001600160a01b03161461030b57604051631fe1e13d60e11b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b0384161790555b610331838361033a565b90505b92915050565b5f828152602081815260408083206001600160a01b038516845290915281205460ff166103da575f838152602081815260408083206001600160a01b03861684529091529020805460ff191660011790556103923390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4506001610334565b505f610334565b80516001600160a01b03811681146103f7575f5ffd5b919050565b5f5f5f5f5f5f5f5f5f6101208a8c031215610415575f5ffd5b895165ffffffffffff8116811461042a575f5ffd5b985061043860208b016103e1565b975061044660408b016103e1565b60608b015160808c01519198509650945061046360a08b016103e1565b935061047160c08b016103e1565b925061047f60e08b016103e1565b915061048e6101008b016103e1565b90509295985092959850929598565b613625806104aa5f395ff3fe6080604052600436106102b9575f3560e01c80638bfe0df111610170578063cefc1429116100d1578063ec87621c11610087578063f3fef3a311610062578063f3fef3a31461086f578063f887ea401461088e578063f8b2cb4f146108ba575f5ffd5b8063ec87621c14610808578063f14210a61461083b578063f29f4d0b1461085a575f5ffd5b8063d547741f116100b7578063d547741f146107a9578063d602b9fd146107c8578063dc73e49c146107dc575f5ffd5b8063cefc14291461073d578063cf6eefb714610751575f5ffd5b8063a217fddf11610126578063b851b7ca1161010c578063b851b7ca146106de578063c31c9c07146106fd578063cc8463c814610729575f5ffd5b8063a217fddf146106ac578063a3b8ef04146106bf575f5ffd5b806391d148541161015657806391d1485414610615578063947a36fb14610664578063a1eda53c14610679575f5ffd5b80638bfe0df1146105b55780638da5cb5b14610601575f5ffd5b80633bbed4a01161021a578063649a5ec7116101d057806367ea88eb116101b657806367ea88eb146105395780637a01a1da1461055857806384ef8ffc1461058b575f5ffd5b8063649a5ec7146104ee57806366d003ac1461050d575f5ffd5b806346b62c4a1161020057806346b62c4a1461049b5780635e1c17ed146104b0578063634e93da146104cf575f5ffd5b80633bbed4a01461045d578063412736571461047c575f5ffd5b806322a900821161026f5780632aaa9628116102555780632aaa9628146103e95780632f2ff15d1461041f57806336568abe1461043e575f5ffd5b806322a900821461038e578063248a9ca3146103ad575f5ffd5b806304a411591161029f57806304a41159146103205780630aa6220b146103595780630f3fffbf1461036f575f5ffd5b806301ffc9a7146102c4578063022d63fb146102f8575f5ffd5b366102c057005b5f5ffd5b3480156102cf575f5ffd5b506102e36102de3660046131f2565b6108d9565b60405190151581526020015b60405180910390f35b348015610303575f5ffd5b50620697805b60405165ffffffffffff90911681526020016102ef565b34801561032b575f5ffd5b50610334600181565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016102ef565b348015610364575f5ffd5b5061036d610934565b005b34801561037a575f5ffd5b5061036d610389366004613252565b610949565b348015610399575f5ffd5b5061036d6103a836600461329a565b610cee565b3480156103b8575f5ffd5b506103db6103c736600461329a565b5f9081526020819052604090206001015490565b6040519081526020016102ef565b3480156103f4575f5ffd5b506104086104033660046132b1565b610d9a565b6040805192151583526020830191909152016102ef565b34801561042a575f5ffd5b5061036d6104393660046132e8565b610f74565b348015610449575f5ffd5b5061036d6104583660046132e8565b610fb9565b348015610468575f5ffd5b5061036d61047736600461330b565b6110be565b348015610487575f5ffd5b5061036d61049636600461330b565b6111b7565b3480156104a6575f5ffd5b506103db60055481565b3480156104bb575f5ffd5b506103db6104ca366004613326565b6112c0565b3480156104da575f5ffd5b5061036d6104e936600461330b565b611823565b3480156104f9575f5ffd5b5061036d61050836600461336b565b611836565b348015610518575f5ffd5b506008546103349073ffffffffffffffffffffffffffffffffffffffff1681565b348015610544575f5ffd5b5061036d610553366004613390565b611849565b348015610563575f5ffd5b506103db7ff8cd32ed93fc2f9fc78152a14807c9609af3d99c5fe4dc6b106a801aaddfe90e81565b348015610596575f5ffd5b5060025473ffffffffffffffffffffffffffffffffffffffff16610334565b3480156105c0575f5ffd5b506103346105cf3660046132b1565b600b60209081525f928352604080842090915290825290205473ffffffffffffffffffffffffffffffffffffffff1681565b34801561060c575f5ffd5b50610334611e51565b348015610620575f5ffd5b506102e361062f3660046132e8565b5f9182526020828152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b34801561066f575f5ffd5b506103db60065481565b348015610684575f5ffd5b5061068d611e76565b6040805165ffffffffffff9384168152929091166020830152016102ef565b3480156106b7575f5ffd5b506103db5f81565b3480156106ca575f5ffd5b5061036d6106d936600461329a565b611ef0565b3480156106e9575f5ffd5b5061036d6106f836600461330b565b611fd8565b348015610708575f5ffd5b506009546103349073ffffffffffffffffffffffffffffffffffffffff1681565b348015610734575f5ffd5b506103096120f1565b348015610748575f5ffd5b5061036d61218e565b34801561075c575f5ffd5b506001546040805173ffffffffffffffffffffffffffffffffffffffff831681527401000000000000000000000000000000000000000090920465ffffffffffff166020830152016102ef565b3480156107b4575f5ffd5b5061036d6107c33660046132e8565b6121ea565b3480156107d3575f5ffd5b5061036d61222b565b3480156107e7575f5ffd5b50600a546103349073ffffffffffffffffffffffffffffffffffffffff1681565b348015610813575f5ffd5b506103db7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b0881565b348015610846575f5ffd5b5061036d61085536600461329a565b61223d565b348015610865575f5ffd5b506103db60075481565b34801561087a575f5ffd5b5061036d6108893660046133c2565b6123af565b348015610899575f5ffd5b506004546103349073ffffffffffffffffffffffffffffffffffffffff1681565b3480156108c5575f5ffd5b506103db6108d436600461330b565b61256b565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f3149878600000000000000000000000000000000000000000000000000000000148061092e575061092e8261263c565b92915050565b5f61093e816126d2565b6109466126dc565b50565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b08610973816126d2565b73ffffffffffffffffffffffffffffffffffffffff841615806109aa575073ffffffffffffffffffffffffffffffffffffffff8316155b156109e1576040517f5159e7e000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216610a4b576040517febb9ebd700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff831660048201526024015b60405180910390fd5b5f8273ffffffffffffffffffffffffffffffffffffffff16630dfe16816040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a95573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ab991906133ec565b90505f8373ffffffffffffffffffffffffffffffffffffffff1663d21220a76040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b05573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b2991906133ec565b90505f8273ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff16148015610b9257508173ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff16145b80610bfe57508173ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff16148015610bfe57508273ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff16145b905080610c5f576040517fd697964600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8087166004830152808916602483015287166044820152606401610a42565b73ffffffffffffffffffffffffffffffffffffffff8781165f818152600b602090815260408083208b8616808552925280832080547fffffffffffffffffffffffff000000000000000000000000000000000000000016958b169586179055519092917fff59a584b1c3a296ea9baaa7fb8580f277f4763522ee6998dc8c8fcfad11934591a450505050505050565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b08610d18816126d2565b8160065403610d61576006546040517f51baf7be000000000000000000000000000000000000000000000000000000008152600481019190915260248101839052604401610a42565b6006805490839055604051839082907f3fb677206c0b314c404bae3da94bee6bda11375c9fe986f266fb033d6cddbbd7905f90a3505050565b5f8073ffffffffffffffffffffffffffffffffffffffff8416610dc157505f905080610f6d565b73ffffffffffffffffffffffffffffffffffffffff83165f9081527ff3fa603c74bfe2a4719960e47343678c3dc690d2b27a2295acc6fc430833aaf9602052604090205460ff16610e1657505f905080610f6d565b5f8473ffffffffffffffffffffffffffffffffffffffff1663c3f909d46040518163ffffffff1660e01b8152600401606060405180830381865afa158015610e60573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e849190613407565b80516040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015291925073ffffffffffffffffffffffffffffffffffffffff16906370a0823190602401602060405180830381865afa158015610ef0573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f149190613497565b9150600554821015610f29575f925050610f6d565b5f600654118015610f3b57505f600754115b15610f67575f60075442610f4f91906134db565b9050600654811015610f65575f93505050610f6d565b505b60019250505b9250929050565b81610fab576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610fb582826126e8565b5050565b81158015610fe1575060025473ffffffffffffffffffffffffffffffffffffffff8281169116145b156110b45760015473ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900465ffffffffffff1681151580611035575065ffffffffffff8116155b8061104857504265ffffffffffff821610155b15611089576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff82166004820152602401610a42565b5050600180547fffffffffffff000000000000ffffffffffffffffffffffffffffffffffffffff1690555b610fb58282612712565b5f6110c8816126d2565b60085473ffffffffffffffffffffffffffffffffffffffff808416911603611140576008546040517fea9ebefb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff91821660048201529083166024820152604401610a42565b6008805473ffffffffffffffffffffffffffffffffffffffff8481167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907fc1416b5cdab50a9fbc872236e1aa54566c6deb40024e63a4b1737ecacf09d6f9905f90a3505050565b5f6111c1816126d2565b60095473ffffffffffffffffffffffffffffffffffffffff808416911603611239576009546040517fea9ebefb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff91821660048201529083166024820152604401610a42565b6009805473ffffffffffffffffffffffffffffffffffffffff8481167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527fc7324ad5feb4318ddf48817d97597d40855e50e83dc1b1b796bd5fb48dd9379f910160405180910390a1505050565b5f6112c961276b565b7ff8cd32ed93fc2f9fc78152a14807c9609af3d99c5fe4dc6b106a801aaddfe90e6112f3816126d2565b73ffffffffffffffffffffffffffffffffffffffff8516611358576040517f235aafe400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff86166004820152602401610a42565b60095473ffffffffffffffffffffffffffffffffffffffff166113a9576040517f5da5a13d0000000000000000000000000000000000000000000000000000000081525f6004820152602401610a42565b600a5473ffffffffffffffffffffffffffffffffffffffff166113f8576040517fffb5baac00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8573ffffffffffffffffffffffffffffffffffffffff1663c3f909d46040518163ffffffff1660e01b8152600401606060405180830381865afa158015611442573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906114669190613407565b8051600a546040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015292935090915f9173ffffffffffffffffffffffffffffffffffffffff16906370a0823190602401602060405180830381865afa1580156114d9573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906114fd9190613497565b9050805f0361155a57600a546040517f1a89d8d400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911660048201525f6024820152604401610a42565b600a5473ffffffffffffffffffffffffffffffffffffffff9081165f908152600b60209081526040808320868516845290915290205416806115ec57600a546040517fa06d7bd800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff91821660048201529084166024820152604401610a42565b600a546009546040517fdd62ed3e00000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff9182166024820181905291909216918490839063dd62ed3e90604401602060405180830381865afa15801561166a573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061168e9190613497565b10156116d5576116d573ffffffffffffffffffffffffffffffffffffffff8316827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6127ae565b6040805161010081018252600a5473ffffffffffffffffffffffffffffffffffffffff9081168252878116602083015262ffffff8d168284015230606083015242608083015260a0820187905260c082018c90525f60e083015260095492517f414bf3890000000000000000000000000000000000000000000000000000000081529192169063414bf3899061176f9084906004016134ee565b6020604051808303815f875af115801561178b573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906117af9190613497565b600a546040805188815260208101849052929b5073ffffffffffffffffffffffffffffffffffffffff898116939216917fdd36740e2a012d93061a0d99eaa9107860955de4e90027d3cf465a055026c407910160405180910390a3505050505050505061181c6001600355565b9392505050565b5f61182d816126d2565b610fb5826128c8565b5f611840816126d2565b610fb582612947565b61185161276b565b7ff8cd32ed93fc2f9fc78152a14807c9609af3d99c5fe4dc6b106a801aaddfe90e61187b816126d2565b73ffffffffffffffffffffffffffffffffffffffff84166118e0576040517f235aafe400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610a42565b5f8473ffffffffffffffffffffffffffffffffffffffff1663c3f909d46040518163ffffffff1660e01b8152600401606060405180830381865afa15801561192a573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061194e9190613407565b8051602082015160055492935090918610156119a4576005546040517f013fafe2000000000000000000000000000000000000000000000000000000008152610a42918891600401918252602082015260400190565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201525f9073ffffffffffffffffffffffffffffffffffffffff8416906370a0823190602401602060405180830381865afa158015611a0e573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611a329190613497565b905080871115611a78576040517f013fafe20000000000000000000000000000000000000000000000000000000081526004810182905260248101889052604401610a42565b5f600654118015611a8a57505f600754115b15611aec575f60075442611a9e91906134db565b9050600654811015611aea576006546040517f1e379906000000000000000000000000000000000000000000000000000000008152610a42918391600401918252602082015260400190565b505b600480546040517fdd62ed3e000000000000000000000000000000000000000000000000000000008152309281019290925273ffffffffffffffffffffffffffffffffffffffff9081166024830181905291899186169063dd62ed3e90604401602060405180830381865afa158015611b67573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611b8b9190613497565b1015611bd257611bd273ffffffffffffffffffffffffffffffffffffffff8516827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6127ae565b600480546040517f1e92008400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8c811693820193909352602481018b9052604481018a9052911690631e920084906064015f604051808303815f87803b158015611c4c575f5ffd5b505af1158015611c5e573d5f5f3e3d5ffd5b5050604080518b815233602082015273ffffffffffffffffffffffffffffffffffffffff808816945088811693508d16917fc7b4d815bd0a3b8348577971cb55a9334c80a0b3ccdfe287b55b5a6320d8480a910160405180910390a44260075560085473ffffffffffffffffffffffffffffffffffffffff1615611e3c57478015611d7c576008546040515f9173ffffffffffffffffffffffffffffffffffffffff169083908381818185875af1925050503d805f8114611d3a576040519150601f19603f3d011682016040523d82523d5f602084013e611d3f565b606091505b5050905080611d7a576040517f87fd589900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201525f9073ffffffffffffffffffffffffffffffffffffffff8616906370a0823190602401602060405180830381865afa158015611de6573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611e0a9190613497565b90508015611e3957600854611e399073ffffffffffffffffffffffffffffffffffffffff8781169116836129b6565b50505b505050505050611e4c6001600355565b505050565b5f611e7160025473ffffffffffffffffffffffffffffffffffffffff1690565b905090565b6002545f907a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015158015611eb857504265ffffffffffff821610155b611ec3575f5f611ee8565b60025474010000000000000000000000000000000000000000900465ffffffffffff16815b915091509091565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b08611f1a816126d2565b815f03611f56576040517f517f903200000000000000000000000000000000000000000000000000000000815260048101839052602401610a42565b8160055403611f9f576005546040517f51baf7be000000000000000000000000000000000000000000000000000000008152600481019190915260248101839052604401610a42565b6005805490839055604051839082907fd6d62b78d21fc4f5151029ca37079020e8a8815c5a488b28ee7499d4c2a19e55905f90a3505050565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b08612002816126d2565b600a5473ffffffffffffffffffffffffffffffffffffffff80841691160361207a57600a546040517fea9ebefb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff91821660048201529083166024820152604401610a42565b600a805473ffffffffffffffffffffffffffffffffffffffff8481167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f36b1a7df795121e2bd67c2b04fc89f611b7b21ee10f345226e13bc1b2ebda30d905f90a3505050565b6002545f907a010000000000000000000000000000000000000000000000000000900465ffffffffffff16801515801561213257504265ffffffffffff8216105b612164576001547a010000000000000000000000000000000000000000000000000000900465ffffffffffff16612188565b60025474010000000000000000000000000000000000000000900465ffffffffffff165b91505090565b60015473ffffffffffffffffffffffffffffffffffffffff163381146121e2576040517fc22c8022000000000000000000000000000000000000000000000000000000008152336004820152602401610a42565b6109466129f4565b81612221576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610fb58282612ae5565b5f612235816126d2565b610946612b09565b61224561276b565b5f61224f816126d2565b475f831561225d578361225f565b815b90508181111561229b576040517f26f4246a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6122bb60025473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff16826040515f6040518083038185875af1925050503d805f811461230f576040519150601f19603f3d011682016040523d82523d5f602084013e612314565b606091505b505090508061234f576040517f87fd589900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60025460405183815273ffffffffffffffffffffffffffffffffffffffff909116906001907fd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb9060200160405180910390a3505050506109466001600355565b6123b761276b565b5f6123c1816126d2565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015283905f9073ffffffffffffffffffffffffffffffffffffffff8316906370a0823190602401602060405180830381865afa15801561242d573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906124519190613497565b90505f84156124605784612462565b815b90508181111561249e576040517f26f4246a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6124de6124c060025473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff851690836129b6565b60025473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff167fd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb8360405161255591815260200190565b60405180910390a350505050610fb56001600355565b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff73ffffffffffffffffffffffffffffffffffffffff8316016125b0575047919050565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff8316906370a0823190602401602060405180830381865afa158015612618573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061092e9190613497565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061092e57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000083161461092e565b6109468133612b13565b6126e65f5f612b98565b565b5f82815260208190526040902060010154612702816126d2565b61270c8383612cf1565b50505050565b73ffffffffffffffffffffffffffffffffffffffff81163314612761576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611e4c8282612daf565b6002600354036127a7576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600355565b6040805173ffffffffffffffffffffffffffffffffffffffff8416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f095ea7b30000000000000000000000000000000000000000000000000000000017905261283a8482612e10565b61270c5760405173ffffffffffffffffffffffffffffffffffffffff84811660248301525f60448301526128be91869182169063095ea7b3906064015b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050612e66565b61270c8482612e66565b5f6128d16120f1565b6128da42612f05565b6128e491906135b3565b90506128f08282612f54565b60405165ffffffffffff8216815273ffffffffffffffffffffffffffffffffffffffff8316907f3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed69060200160405180910390a25050565b5f61295182612fef565b61295a42612f05565b61296491906135b3565b90506129708282612b98565b6040805165ffffffffffff8085168252831660208201527ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b910160405180910390a15050565b60405173ffffffffffffffffffffffffffffffffffffffff838116602483015260448201839052611e4c91859182169063a9059cbb90606401612877565b60015473ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900465ffffffffffff16801580612a4457504265ffffffffffff821610155b15612a85576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff82166004820152602401610a42565b612aad5f612aa860025473ffffffffffffffffffffffffffffffffffffffff1690565b612daf565b50612ab85f83612cf1565b5050600180547fffffffffffff000000000000000000000000000000000000000000000000000016905550565b5f82815260208190526040902060010154612aff816126d2565b61270c8383612daf565b6126e65f5f612f54565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16610fb5576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8216600482015260248101839052604401610a42565b6002547a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015612c6c574265ffffffffffff82161015612c43576002546001805479ffffffffffffffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000090920465ffffffffffff167a01000000000000000000000000000000000000000000000000000002919091179055612c6c565b6040517f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5905f90a15b506002805473ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000065ffffffffffff9485160279ffffffffffffffffffffffffffffffffffffffffffffffffffff16177a0100000000000000000000000000000000000000000000000000009290931691909102919091179055565b5f82612da5575f612d1760025473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614612d64576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84161790555b61181c8383613040565b5f82158015612dd8575060025473ffffffffffffffffffffffffffffffffffffffff8381169116145b15612e0657600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001690555b61181c8383613139565b5f5f5f5f60205f8651602088015f8a5af192503d91505f519050828015612e5c57508115612e415780600114612e5c565b5f8673ffffffffffffffffffffffffffffffffffffffff163b115b9695505050505050565b5f5f60205f8451602086015f885af180612e85576040513d5f823e3d81fd5b50505f513d91508115612e9c578060011415612eb6565b73ffffffffffffffffffffffffffffffffffffffff84163b155b1561270c576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610a42565b5f65ffffffffffff821115612f50576040517f6dfcc6500000000000000000000000000000000000000000000000000000000081526030600482015260248101839052604401610a42565b5090565b600180547401000000000000000000000000000000000000000065ffffffffffff84811682027fffffffffffff0000000000000000000000000000000000000000000000000000841673ffffffffffffffffffffffffffffffffffffffff881617179093559004168015611e4c576040517f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109905f90a1505050565b5f5f612ff96120f1565b90508065ffffffffffff168365ffffffffffff16116130215761301c83826135d1565b61181c565b61181c65ffffffffffff8416620697805f82821882841002821861181c565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff16613132575f8381526020818152604080832073ffffffffffffffffffffffffffffffffffffffff86168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556130d03390565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a450600161092e565b505f61092e565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff1615613132575f8381526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8616808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a450600161092e565b5f60208284031215613202575f5ffd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461181c575f5ffd5b73ffffffffffffffffffffffffffffffffffffffff81168114610946575f5ffd5b5f5f5f60608486031215613264575f5ffd5b833561326f81613231565b9250602084013561327f81613231565b9150604084013561328f81613231565b809150509250925092565b5f602082840312156132aa575f5ffd5b5035919050565b5f5f604083850312156132c2575f5ffd5b82356132cd81613231565b915060208301356132dd81613231565b809150509250929050565b5f5f604083850312156132f9575f5ffd5b8235915060208301356132dd81613231565b5f6020828403121561331b575f5ffd5b813561181c81613231565b5f5f5f60608486031215613338575f5ffd5b833561334381613231565b9250602084013562ffffff8116811461335a575f5ffd5b929592945050506040919091013590565b5f6020828403121561337b575f5ffd5b813565ffffffffffff8116811461181c575f5ffd5b5f5f5f606084860312156133a2575f5ffd5b83356133ad81613231565b95602085013595506040909401359392505050565b5f5f604083850312156133d3575f5ffd5b82356133de81613231565b946020939093013593505050565b5f602082840312156133fc575f5ffd5b815161181c81613231565b5f6060828403128015613418575f5ffd5b506040516060810167ffffffffffffffff81118282101715613461577f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604052825161346f81613231565b8152602083015161347f81613231565b60208201526040928301519281019290925250919050565b5f602082840312156134a7575f5ffd5b5051919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b8181038181111561092e5761092e6134ae565b5f6101008201905073ffffffffffffffffffffffffffffffffffffffff835116825273ffffffffffffffffffffffffffffffffffffffff602084015116602083015262ffffff60408401511660408301526060830151613566606084018273ffffffffffffffffffffffffffffffffffffffff169052565b506080830151608083015260a083015160a083015260c083015160c083015260e08301516135ac60e084018273ffffffffffffffffffffffffffffffffffffffff169052565b5092915050565b65ffffffffffff818116838216019081111561092e5761092e6134ae565b65ffffffffffff828116828216039081111561092e5761092e6134ae56fea2646970667358221220b0b6bef686b31cd1fcaf77067778c5a818d4b79f61220738dbba7d3769a46d1464736f6c634300081c0033241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b08f8cd32ed93fc2f9fc78152a14807c9609af3d99c5fe4dc6b106a801aaddfe90e",
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
func DeployBuyBot(auth *bind.TransactOpts, backend bind.ContractBackend, _initialDelay *big.Int, _owner common.Address, _router common.Address, _minOrderAmount *big.Int, _interval *big.Int, _recipient common.Address, _buyer common.Address, _manager common.Address, _swapRouter common.Address) (common.Address, *types.Transaction, *BuyBot, error) {
	parsed, err := BuyBotMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BuyBotBin), backend, _initialDelay, _owner, _router, _minOrderAmount, _interval, _recipient, _buyer, _manager, _swapRouter)
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

// SwapToQuote is a paid mutator transaction binding the contract method 0x5e1c17ed.
//
// Solidity: function swapToQuote(address pair, uint24 uniswapFee, uint256 minAmountOut) returns(uint256 amountOut)
func (_BuyBot *BuyBotTransactor) SwapToQuote(opts *bind.TransactOpts, pair common.Address, uniswapFee *big.Int, minAmountOut *big.Int) (*types.Transaction, error) {
	return _BuyBot.contract.Transact(opts, "swapToQuote", pair, uniswapFee, minAmountOut)
}

// SwapToQuote is a paid mutator transaction binding the contract method 0x5e1c17ed.
//
// Solidity: function swapToQuote(address pair, uint24 uniswapFee, uint256 minAmountOut) returns(uint256 amountOut)
func (_BuyBot *BuyBotSession) SwapToQuote(pair common.Address, uniswapFee *big.Int, minAmountOut *big.Int) (*types.Transaction, error) {
	return _BuyBot.Contract.SwapToQuote(&_BuyBot.TransactOpts, pair, uniswapFee, minAmountOut)
}

// SwapToQuote is a paid mutator transaction binding the contract method 0x5e1c17ed.
//
// Solidity: function swapToQuote(address pair, uint24 uniswapFee, uint256 minAmountOut) returns(uint256 amountOut)
func (_BuyBot *BuyBotTransactorSession) SwapToQuote(pair common.Address, uniswapFee *big.Int, minAmountOut *big.Int) (*types.Transaction, error) {
	return _BuyBot.Contract.SwapToQuote(&_BuyBot.TransactOpts, pair, uniswapFee, minAmountOut)
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
