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

// BuyBotMetaData contains all meta data concerning the BuyBot contract.
var BuyBotMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"_initialDelay\",\"type\":\"uint48\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minOrderAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_interval\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_swapRouter\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"BUYER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NATIVE_COIN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"beginDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxMatchCount\",\"type\":\"uint256\"}],\"name\":\"buyMarket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"canBuyMarket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"canBuy\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"}],\"name\":\"changeDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelayIncreaseWait\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastBuyTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minOrderAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"recipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollbackDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"contractIRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_interval\",\"type\":\"uint256\"}],\"name\":\"setInterval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minOrderAmount\",\"type\":\"uint256\"}],\"name\":\"setMinOrderAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"setRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"setSwapPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_swapRouter\",\"type\":\"address\"}],\"name\":\"setSwapRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_swapToken\",\"type\":\"address\"}],\"name\":\"setSwapToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"swapPools\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapRouter\",\"outputs\":[{\"internalType\":\"contractISwapRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"uniswapFee\",\"type\":\"uint24\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"}],\"name\":\"swapToQuote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminDelayChangeCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"effectSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminDelayChangeScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminTransferCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"acceptSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminTransferScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"before\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"current\",\"type\":\"uint256\"}],\"name\":\"IntervalSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"MarketBuyExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"before\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"current\",\"type\":\"uint256\"}],\"name\":\"MinOrderAmountSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"before\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"RecipientSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"SwapExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"SwapPoolSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"before\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"SwapRouterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"before\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"SwapTokenSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"name\":\"AccessControlEnforcedDefaultAdminDelay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccessControlEnforcedDefaultAdminRules\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"defaultAdmin\",\"type\":\"address\"}],\"name\":\"AccessControlInvalidDefaultAdmin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oldValue\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"BuyBotAddressNotChanged\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BuyBotETHTransferFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minOrderAmount\",\"type\":\"uint256\"}],\"name\":\"BuyBotInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"BuyBotInsufficientSwapBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BuyBotInsufficientWithdrawBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timeSinceLastBuy\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredInterval\",\"type\":\"uint256\"}],\"name\":\"BuyBotIntervalNotPassed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"BuyBotInvalidAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"}],\"name\":\"BuyBotInvalidBuyer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"BuyBotInvalidManager\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"BuyBotInvalidMinOrderAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"BuyBotInvalidPair\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"BuyBotInvalidPool\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"}],\"name\":\"BuyBotInvalidPoolTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"BuyBotInvalidRouter\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"BuyBotInvalidSwapRouter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BuyBotInvalidTokenAddresses\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BuyBotNoSwapToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"oldValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"BuyBotNotChanged\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"}],\"name\":\"BuyBotPoolNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"}]",
	ID:  "BuyBot",
	Bin: "0x608060405234801561000f575f5ffd5b50604051613b0f380380613b0f83398101604081905261002e916103fc565b88886001600160a01b03811661005e57604051636116401160e11b81525f60048201526024015b60405180910390fd5b600180546001600160d01b0316600160d01b65ffffffffffff8516021790556100875f826102cb565b50506001600355506001600160a01b0387166100c157604051631561f27b60e21b81526001600160a01b0388166004820152602401610055565b855f036100e4576040516328bfc81960e11b815260048101879052602401610055565b6001600160a01b03831661011657604051632a7409a960e01b81526001600160a01b0384166004820152602401610055565b6001600160a01b0382166101485760405163192178dd60e11b81526001600160a01b0383166004820152602401610055565b600480546001600160a01b03808a166001600160a01b031992831617909255600588905560068790556008805487841690831617905560098054928416929091169190911790556101a65f516020613aef5f395f51905f52896102cb565b506101be5f516020613aef5f395f51905f52846102cb565b506101d65f516020613acf5f395f51905f52896102cb565b506101ee5f516020613acf5f395f51905f52836102cb565b5060405186905f907fd6d62b78d21fc4f5151029ca37079020e8a8815c5a488b28ee7499d4c2a19e55908290a360405185905f907f3fb677206c0b314c404bae3da94bee6bda11375c9fe986f266fb033d6cddbbd7908290a36040516001600160a01b038516905f907fc1416b5cdab50a9fbc872236e1aa54566c6deb40024e63a4b1737ecacf09d6f9908290a3604080515f81526001600160a01b03831660208201527fc7324ad5feb4318ddf48817d97597d40855e50e83dc1b1b796bd5fb48dd9379f910160405180910390a150505050505050505061049d565b5f82610327575f6102e46002546001600160a01b031690565b6001600160a01b03161461030b57604051631fe1e13d60e11b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b0384161790555b610331838361033a565b90505b92915050565b5f828152602081815260408083206001600160a01b038516845290915281205460ff166103da575f838152602081815260408083206001600160a01b03861684529091529020805460ff191660011790556103923390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4506001610334565b505f610334565b80516001600160a01b03811681146103f7575f5ffd5b919050565b5f5f5f5f5f5f5f5f5f6101208a8c031215610415575f5ffd5b895165ffffffffffff8116811461042a575f5ffd5b985061043860208b016103e1565b975061044660408b016103e1565b60608b015160808c01519198509650945061046360a08b016103e1565b935061047160c08b016103e1565b925061047f60e08b016103e1565b915061048e6101008b016103e1565b90509295985092959850929598565b613625806104aa5f395ff3fe6080604052600436106102b9575f3560e01c80638bfe0df111610170578063cefc1429116100d1578063ec87621c11610087578063f3fef3a311610062578063f3fef3a31461086f578063f887ea401461088e578063f8b2cb4f146108ba575f5ffd5b8063ec87621c14610808578063f14210a61461083b578063f29f4d0b1461085a575f5ffd5b8063d547741f116100b7578063d547741f146107a9578063d602b9fd146107c8578063dc73e49c146107dc575f5ffd5b8063cefc14291461073d578063cf6eefb714610751575f5ffd5b8063a217fddf11610126578063b851b7ca1161010c578063b851b7ca146106de578063c31c9c07146106fd578063cc8463c814610729575f5ffd5b8063a217fddf146106ac578063a3b8ef04146106bf575f5ffd5b806391d148541161015657806391d1485414610615578063947a36fb14610664578063a1eda53c14610679575f5ffd5b80638bfe0df1146105b55780638da5cb5b14610601575f5ffd5b80633bbed4a01161021a578063649a5ec7116101d057806367ea88eb116101b657806367ea88eb146105395780637a01a1da1461055857806384ef8ffc1461058b575f5ffd5b8063649a5ec7146104ee57806366d003ac1461050d575f5ffd5b806346b62c4a1161020057806346b62c4a1461049b5780635e1c17ed146104b0578063634e93da146104cf575f5ffd5b80633bbed4a01461045d578063412736571461047c575f5ffd5b806322a900821161026f5780632aaa9628116102555780632aaa9628146103e95780632f2ff15d1461041f57806336568abe1461043e575f5ffd5b806322a900821461038e578063248a9ca3146103ad575f5ffd5b806304a411591161029f57806304a41159146103205780630aa6220b146103595780630f3fffbf1461036f575f5ffd5b806301ffc9a7146102c4578063022d63fb146102f8575f5ffd5b366102c057005b5f5ffd5b3480156102cf575f5ffd5b506102e36102de3660046131f2565b6108d9565b60405190151581526020015b60405180910390f35b348015610303575f5ffd5b50620697805b60405165ffffffffffff90911681526020016102ef565b34801561032b575f5ffd5b50610334600181565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016102ef565b348015610364575f5ffd5b5061036d610934565b005b34801561037a575f5ffd5b5061036d610389366004613252565b610949565b348015610399575f5ffd5b5061036d6103a836600461329a565b610cee565b3480156103b8575f5ffd5b506103db6103c736600461329a565b5f9081526020819052604090206001015490565b6040519081526020016102ef565b3480156103f4575f5ffd5b506104086104033660046132b1565b610d9a565b6040805192151583526020830191909152016102ef565b34801561042a575f5ffd5b5061036d6104393660046132e8565b610f74565b348015610449575f5ffd5b5061036d6104583660046132e8565b610fb9565b348015610468575f5ffd5b5061036d61047736600461330b565b6110be565b348015610487575f5ffd5b5061036d61049636600461330b565b6111b7565b3480156104a6575f5ffd5b506103db60055481565b3480156104bb575f5ffd5b506103db6104ca366004613326565b6112c0565b3480156104da575f5ffd5b5061036d6104e936600461330b565b611823565b3480156104f9575f5ffd5b5061036d61050836600461336b565b611836565b348015610518575f5ffd5b506008546103349073ffffffffffffffffffffffffffffffffffffffff1681565b348015610544575f5ffd5b5061036d610553366004613390565b611849565b348015610563575f5ffd5b506103db7ff8cd32ed93fc2f9fc78152a14807c9609af3d99c5fe4dc6b106a801aaddfe90e81565b348015610596575f5ffd5b5060025473ffffffffffffffffffffffffffffffffffffffff16610334565b3480156105c0575f5ffd5b506103346105cf3660046132b1565b600b60209081525f928352604080842090915290825290205473ffffffffffffffffffffffffffffffffffffffff1681565b34801561060c575f5ffd5b50610334611e51565b348015610620575f5ffd5b506102e361062f3660046132e8565b5f9182526020828152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b34801561066f575f5ffd5b506103db60065481565b348015610684575f5ffd5b5061068d611e76565b6040805165ffffffffffff9384168152929091166020830152016102ef565b3480156106b7575f5ffd5b506103db5f81565b3480156106ca575f5ffd5b5061036d6106d936600461329a565b611ef0565b3480156106e9575f5ffd5b5061036d6106f836600461330b565b611fd8565b348015610708575f5ffd5b506009546103349073ffffffffffffffffffffffffffffffffffffffff1681565b348015610734575f5ffd5b506103096120f1565b348015610748575f5ffd5b5061036d61218e565b34801561075c575f5ffd5b506001546040805173ffffffffffffffffffffffffffffffffffffffff831681527401000000000000000000000000000000000000000090920465ffffffffffff166020830152016102ef565b3480156107b4575f5ffd5b5061036d6107c33660046132e8565b6121ea565b3480156107d3575f5ffd5b5061036d61222b565b3480156107e7575f5ffd5b50600a546103349073ffffffffffffffffffffffffffffffffffffffff1681565b348015610813575f5ffd5b506103db7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b0881565b348015610846575f5ffd5b5061036d61085536600461329a565b61223d565b348015610865575f5ffd5b506103db60075481565b34801561087a575f5ffd5b5061036d6108893660046133c2565b6123af565b348015610899575f5ffd5b506004546103349073ffffffffffffffffffffffffffffffffffffffff1681565b3480156108c5575f5ffd5b506103db6108d436600461330b565b61256b565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f3149878600000000000000000000000000000000000000000000000000000000148061092e575061092e8261263c565b92915050565b5f61093e816126d2565b6109466126dc565b50565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b08610973816126d2565b73ffffffffffffffffffffffffffffffffffffffff841615806109aa575073ffffffffffffffffffffffffffffffffffffffff8316155b156109e1576040517f5159e7e000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216610a4b576040517febb9ebd700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff831660048201526024015b60405180910390fd5b5f8273ffffffffffffffffffffffffffffffffffffffff16630dfe16816040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a95573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ab991906133ec565b90505f8373ffffffffffffffffffffffffffffffffffffffff1663d21220a76040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b05573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b2991906133ec565b90505f8273ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff16148015610b9257508173ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff16145b80610bfe57508173ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff16148015610bfe57508273ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff16145b905080610c5f576040517fd697964600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8087166004830152808916602483015287166044820152606401610a42565b73ffffffffffffffffffffffffffffffffffffffff8781165f818152600b602090815260408083208b8616808552925280832080547fffffffffffffffffffffffff000000000000000000000000000000000000000016958b169586179055519092917fff59a584b1c3a296ea9baaa7fb8580f277f4763522ee6998dc8c8fcfad11934591a450505050505050565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b08610d18816126d2565b8160065403610d61576006546040517f51baf7be000000000000000000000000000000000000000000000000000000008152600481019190915260248101839052604401610a42565b6006805490839055604051839082907f3fb677206c0b314c404bae3da94bee6bda11375c9fe986f266fb033d6cddbbd7905f90a3505050565b5f8073ffffffffffffffffffffffffffffffffffffffff8416610dc157505f905080610f6d565b73ffffffffffffffffffffffffffffffffffffffff83165f9081527ff3fa603c74bfe2a4719960e47343678c3dc690d2b27a2295acc6fc430833aaf9602052604090205460ff16610e1657505f905080610f6d565b5f8473ffffffffffffffffffffffffffffffffffffffff1663c3f909d46040518163ffffffff1660e01b8152600401606060405180830381865afa158015610e60573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e849190613407565b80516040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015291925073ffffffffffffffffffffffffffffffffffffffff16906370a0823190602401602060405180830381865afa158015610ef0573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f149190613497565b9150600554821015610f29575f925050610f6d565b5f600654118015610f3b57505f600754115b15610f67575f60075442610f4f91906134db565b9050600654811015610f65575f93505050610f6d565b505b60019250505b9250929050565b81610fab576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610fb582826126e8565b5050565b81158015610fe1575060025473ffffffffffffffffffffffffffffffffffffffff8281169116145b156110b45760015473ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900465ffffffffffff1681151580611035575065ffffffffffff8116155b8061104857504265ffffffffffff821610155b15611089576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff82166004820152602401610a42565b5050600180547fffffffffffff000000000000ffffffffffffffffffffffffffffffffffffffff1690555b610fb58282612712565b5f6110c8816126d2565b60085473ffffffffffffffffffffffffffffffffffffffff808416911603611140576008546040517fea9ebefb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff91821660048201529083166024820152604401610a42565b6008805473ffffffffffffffffffffffffffffffffffffffff8481167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907fc1416b5cdab50a9fbc872236e1aa54566c6deb40024e63a4b1737ecacf09d6f9905f90a3505050565b5f6111c1816126d2565b60095473ffffffffffffffffffffffffffffffffffffffff808416911603611239576009546040517fea9ebefb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff91821660048201529083166024820152604401610a42565b6009805473ffffffffffffffffffffffffffffffffffffffff8481167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527fc7324ad5feb4318ddf48817d97597d40855e50e83dc1b1b796bd5fb48dd9379f910160405180910390a1505050565b5f6112c961276b565b7ff8cd32ed93fc2f9fc78152a14807c9609af3d99c5fe4dc6b106a801aaddfe90e6112f3816126d2565b73ffffffffffffffffffffffffffffffffffffffff8516611358576040517f235aafe400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff86166004820152602401610a42565b60095473ffffffffffffffffffffffffffffffffffffffff166113a9576040517f5da5a13d0000000000000000000000000000000000000000000000000000000081525f6004820152602401610a42565b600a5473ffffffffffffffffffffffffffffffffffffffff166113f8576040517fffb5baac00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8573ffffffffffffffffffffffffffffffffffffffff1663c3f909d46040518163ffffffff1660e01b8152600401606060405180830381865afa158015611442573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906114669190613407565b8051600a546040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015292935090915f9173ffffffffffffffffffffffffffffffffffffffff16906370a0823190602401602060405180830381865afa1580156114d9573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906114fd9190613497565b9050805f0361155a57600a546040517f1a89d8d400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911660048201525f6024820152604401610a42565b600a5473ffffffffffffffffffffffffffffffffffffffff9081165f908152600b60209081526040808320868516845290915290205416806115ec57600a546040517fa06d7bd800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff91821660048201529084166024820152604401610a42565b600a546009546040517fdd62ed3e00000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff9182166024820181905291909216918490839063dd62ed3e90604401602060405180830381865afa15801561166a573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061168e9190613497565b10156116d5576116d573ffffffffffffffffffffffffffffffffffffffff8316827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6127ae565b6040805161010081018252600a5473ffffffffffffffffffffffffffffffffffffffff9081168252878116602083015262ffffff8d168284015230606083015242608083015260a0820187905260c082018c90525f60e083015260095492517f414bf3890000000000000000000000000000000000000000000000000000000081529192169063414bf3899061176f9084906004016134ee565b6020604051808303815f875af115801561178b573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906117af9190613497565b600a546040805188815260208101849052929b5073ffffffffffffffffffffffffffffffffffffffff898116939216917fdd36740e2a012d93061a0d99eaa9107860955de4e90027d3cf465a055026c407910160405180910390a3505050505050505061181c6001600355565b9392505050565b5f61182d816126d2565b610fb5826128c8565b5f611840816126d2565b610fb582612947565b61185161276b565b7ff8cd32ed93fc2f9fc78152a14807c9609af3d99c5fe4dc6b106a801aaddfe90e61187b816126d2565b73ffffffffffffffffffffffffffffffffffffffff84166118e0576040517f235aafe400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610a42565b5f8473ffffffffffffffffffffffffffffffffffffffff1663c3f909d46040518163ffffffff1660e01b8152600401606060405180830381865afa15801561192a573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061194e9190613407565b8051602082015160055492935090918610156119a4576005546040517f013fafe2000000000000000000000000000000000000000000000000000000008152610a42918891600401918252602082015260400190565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201525f9073ffffffffffffffffffffffffffffffffffffffff8416906370a0823190602401602060405180830381865afa158015611a0e573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611a329190613497565b905080871115611a78576040517f013fafe20000000000000000000000000000000000000000000000000000000081526004810182905260248101889052604401610a42565b5f600654118015611a8a57505f600754115b15611aec575f60075442611a9e91906134db565b9050600654811015611aea576006546040517f1e379906000000000000000000000000000000000000000000000000000000008152610a42918391600401918252602082015260400190565b505b600480546040517fdd62ed3e000000000000000000000000000000000000000000000000000000008152309281019290925273ffffffffffffffffffffffffffffffffffffffff9081166024830181905291899186169063dd62ed3e90604401602060405180830381865afa158015611b67573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611b8b9190613497565b1015611bd257611bd273ffffffffffffffffffffffffffffffffffffffff8516827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6127ae565b600480546040517f1e92008400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8c811693820193909352602481018b9052604481018a9052911690631e920084906064015f604051808303815f87803b158015611c4c575f5ffd5b505af1158015611c5e573d5f5f3e3d5ffd5b5050604080518b815233602082015273ffffffffffffffffffffffffffffffffffffffff808816945088811693508d16917fc7b4d815bd0a3b8348577971cb55a9334c80a0b3ccdfe287b55b5a6320d8480a910160405180910390a44260075560085473ffffffffffffffffffffffffffffffffffffffff1615611e3c57478015611d7c576008546040515f9173ffffffffffffffffffffffffffffffffffffffff169083908381818185875af1925050503d805f8114611d3a576040519150601f19603f3d011682016040523d82523d5f602084013e611d3f565b606091505b5050905080611d7a576040517f87fd589900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201525f9073ffffffffffffffffffffffffffffffffffffffff8616906370a0823190602401602060405180830381865afa158015611de6573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611e0a9190613497565b90508015611e3957600854611e399073ffffffffffffffffffffffffffffffffffffffff8781169116836129b6565b50505b505050505050611e4c6001600355565b505050565b5f611e7160025473ffffffffffffffffffffffffffffffffffffffff1690565b905090565b6002545f907a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015158015611eb857504265ffffffffffff821610155b611ec3575f5f611ee8565b60025474010000000000000000000000000000000000000000900465ffffffffffff16815b915091509091565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b08611f1a816126d2565b815f03611f56576040517f517f903200000000000000000000000000000000000000000000000000000000815260048101839052602401610a42565b8160055403611f9f576005546040517f51baf7be000000000000000000000000000000000000000000000000000000008152600481019190915260248101839052604401610a42565b6005805490839055604051839082907fd6d62b78d21fc4f5151029ca37079020e8a8815c5a488b28ee7499d4c2a19e55905f90a3505050565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b08612002816126d2565b600a5473ffffffffffffffffffffffffffffffffffffffff80841691160361207a57600a546040517fea9ebefb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff91821660048201529083166024820152604401610a42565b600a805473ffffffffffffffffffffffffffffffffffffffff8481167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f36b1a7df795121e2bd67c2b04fc89f611b7b21ee10f345226e13bc1b2ebda30d905f90a3505050565b6002545f907a010000000000000000000000000000000000000000000000000000900465ffffffffffff16801515801561213257504265ffffffffffff8216105b612164576001547a010000000000000000000000000000000000000000000000000000900465ffffffffffff16612188565b60025474010000000000000000000000000000000000000000900465ffffffffffff165b91505090565b60015473ffffffffffffffffffffffffffffffffffffffff163381146121e2576040517fc22c8022000000000000000000000000000000000000000000000000000000008152336004820152602401610a42565b6109466129f4565b81612221576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610fb58282612ae5565b5f612235816126d2565b610946612b09565b61224561276b565b5f61224f816126d2565b475f831561225d578361225f565b815b90508181111561229b576040517f26f4246a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6122bb60025473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff16826040515f6040518083038185875af1925050503d805f811461230f576040519150601f19603f3d011682016040523d82523d5f602084013e612314565b606091505b505090508061234f576040517f87fd589900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60025460405183815273ffffffffffffffffffffffffffffffffffffffff909116906001907fd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb9060200160405180910390a3505050506109466001600355565b6123b761276b565b5f6123c1816126d2565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015283905f9073ffffffffffffffffffffffffffffffffffffffff8316906370a0823190602401602060405180830381865afa15801561242d573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906124519190613497565b90505f84156124605784612462565b815b90508181111561249e576040517f26f4246a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6124de6124c060025473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff851690836129b6565b60025473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff167fd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb8360405161255591815260200190565b60405180910390a350505050610fb56001600355565b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff73ffffffffffffffffffffffffffffffffffffffff8316016125b0575047919050565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff8316906370a0823190602401602060405180830381865afa158015612618573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061092e9190613497565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061092e57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000083161461092e565b6109468133612b13565b6126e65f5f612b98565b565b5f82815260208190526040902060010154612702816126d2565b61270c8383612cf1565b50505050565b73ffffffffffffffffffffffffffffffffffffffff81163314612761576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611e4c8282612daf565b6002600354036127a7576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600355565b6040805173ffffffffffffffffffffffffffffffffffffffff8416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f095ea7b30000000000000000000000000000000000000000000000000000000017905261283a8482612e10565b61270c5760405173ffffffffffffffffffffffffffffffffffffffff84811660248301525f60448301526128be91869182169063095ea7b3906064015b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050612e66565b61270c8482612e66565b5f6128d16120f1565b6128da42612f05565b6128e491906135b3565b90506128f08282612f54565b60405165ffffffffffff8216815273ffffffffffffffffffffffffffffffffffffffff8316907f3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed69060200160405180910390a25050565b5f61295182612fef565b61295a42612f05565b61296491906135b3565b90506129708282612b98565b6040805165ffffffffffff8085168252831660208201527ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b910160405180910390a15050565b60405173ffffffffffffffffffffffffffffffffffffffff838116602483015260448201839052611e4c91859182169063a9059cbb90606401612877565b60015473ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900465ffffffffffff16801580612a4457504265ffffffffffff821610155b15612a85576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff82166004820152602401610a42565b612aad5f612aa860025473ffffffffffffffffffffffffffffffffffffffff1690565b612daf565b50612ab85f83612cf1565b5050600180547fffffffffffff000000000000000000000000000000000000000000000000000016905550565b5f82815260208190526040902060010154612aff816126d2565b61270c8383612daf565b6126e65f5f612f54565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16610fb5576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8216600482015260248101839052604401610a42565b6002547a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015612c6c574265ffffffffffff82161015612c43576002546001805479ffffffffffffffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000090920465ffffffffffff167a01000000000000000000000000000000000000000000000000000002919091179055612c6c565b6040517f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5905f90a15b506002805473ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000065ffffffffffff9485160279ffffffffffffffffffffffffffffffffffffffffffffffffffff16177a0100000000000000000000000000000000000000000000000000009290931691909102919091179055565b5f82612da5575f612d1760025473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614612d64576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84161790555b61181c8383613040565b5f82158015612dd8575060025473ffffffffffffffffffffffffffffffffffffffff8381169116145b15612e0657600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001690555b61181c8383613139565b5f5f5f5f60205f8651602088015f8a5af192503d91505f519050828015612e5c57508115612e415780600114612e5c565b5f8673ffffffffffffffffffffffffffffffffffffffff163b115b9695505050505050565b5f5f60205f8451602086015f885af180612e85576040513d5f823e3d81fd5b50505f513d91508115612e9c578060011415612eb6565b73ffffffffffffffffffffffffffffffffffffffff84163b155b1561270c576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610a42565b5f65ffffffffffff821115612f50576040517f6dfcc6500000000000000000000000000000000000000000000000000000000081526030600482015260248101839052604401610a42565b5090565b600180547401000000000000000000000000000000000000000065ffffffffffff84811682027fffffffffffff0000000000000000000000000000000000000000000000000000841673ffffffffffffffffffffffffffffffffffffffff881617179093559004168015611e4c576040517f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109905f90a1505050565b5f5f612ff96120f1565b90508065ffffffffffff168365ffffffffffff16116130215761301c83826135d1565b61181c565b61181c65ffffffffffff8416620697805f82821882841002821861181c565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff16613132575f8381526020818152604080832073ffffffffffffffffffffffffffffffffffffffff86168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556130d03390565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a450600161092e565b505f61092e565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff1615613132575f8381526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8616808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a450600161092e565b5f60208284031215613202575f5ffd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461181c575f5ffd5b73ffffffffffffffffffffffffffffffffffffffff81168114610946575f5ffd5b5f5f5f60608486031215613264575f5ffd5b833561326f81613231565b9250602084013561327f81613231565b9150604084013561328f81613231565b809150509250925092565b5f602082840312156132aa575f5ffd5b5035919050565b5f5f604083850312156132c2575f5ffd5b82356132cd81613231565b915060208301356132dd81613231565b809150509250929050565b5f5f604083850312156132f9575f5ffd5b8235915060208301356132dd81613231565b5f6020828403121561331b575f5ffd5b813561181c81613231565b5f5f5f60608486031215613338575f5ffd5b833561334381613231565b9250602084013562ffffff8116811461335a575f5ffd5b929592945050506040919091013590565b5f6020828403121561337b575f5ffd5b813565ffffffffffff8116811461181c575f5ffd5b5f5f5f606084860312156133a2575f5ffd5b83356133ad81613231565b95602085013595506040909401359392505050565b5f5f604083850312156133d3575f5ffd5b82356133de81613231565b946020939093013593505050565b5f602082840312156133fc575f5ffd5b815161181c81613231565b5f6060828403128015613418575f5ffd5b506040516060810167ffffffffffffffff81118282101715613461577f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604052825161346f81613231565b8152602083015161347f81613231565b60208201526040928301519281019290925250919050565b5f602082840312156134a7575f5ffd5b5051919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b8181038181111561092e5761092e6134ae565b5f6101008201905073ffffffffffffffffffffffffffffffffffffffff835116825273ffffffffffffffffffffffffffffffffffffffff602084015116602083015262ffffff60408401511660408301526060830151613566606084018273ffffffffffffffffffffffffffffffffffffffff169052565b506080830151608083015260a083015160a083015260c083015160c083015260e08301516135ac60e084018273ffffffffffffffffffffffffffffffffffffffff169052565b5092915050565b65ffffffffffff818116838216019081111561092e5761092e6134ae565b65ffffffffffff828116828216039081111561092e5761092e6134ae56fea2646970667358221220b0b6bef686b31cd1fcaf77067778c5a818d4b79f61220738dbba7d3769a46d1464736f6c634300081c0033241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b08f8cd32ed93fc2f9fc78152a14807c9609af3d99c5fe4dc6b106a801aaddfe90e",
}

// BuyBot is an auto generated Go binding around an Ethereum contract.
type BuyBot struct {
	abi abi.ABI
}

// NewBuyBot creates a new instance of BuyBot.
func NewBuyBot() *BuyBot {
	parsed, err := BuyBotMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &BuyBot{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *BuyBot) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(uint48 _initialDelay, address _owner, address _router, uint256 _minOrderAmount, uint256 _interval, address _recipient, address _buyer, address _manager, address _swapRouter) returns()
func (buyBot *BuyBot) PackConstructor(_initialDelay *big.Int, _owner common.Address, _router common.Address, _minOrderAmount *big.Int, _interval *big.Int, _recipient common.Address, _buyer common.Address, _manager common.Address, _swapRouter common.Address) []byte {
	enc, err := buyBot.abi.Pack("", _initialDelay, _owner, _router, _minOrderAmount, _interval, _recipient, _buyer, _manager, _swapRouter)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackBUYERROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7a01a1da.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function BUYER_ROLE() view returns(bytes32)
func (buyBot *BuyBot) PackBUYERROLE() []byte {
	enc, err := buyBot.abi.Pack("BUYER_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBUYERROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7a01a1da.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function BUYER_ROLE() view returns(bytes32)
func (buyBot *BuyBot) TryPackBUYERROLE() ([]byte, error) {
	return buyBot.abi.Pack("BUYER_ROLE")
}

// UnpackBUYERROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7a01a1da.
//
// Solidity: function BUYER_ROLE() view returns(bytes32)
func (buyBot *BuyBot) UnpackBUYERROLE(data []byte) ([32]byte, error) {
	out, err := buyBot.abi.Unpack("BUYER_ROLE", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackDEFAULTADMINROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa217fddf.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (buyBot *BuyBot) PackDEFAULTADMINROLE() []byte {
	enc, err := buyBot.abi.Pack("DEFAULT_ADMIN_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDEFAULTADMINROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa217fddf.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (buyBot *BuyBot) TryPackDEFAULTADMINROLE() ([]byte, error) {
	return buyBot.abi.Pack("DEFAULT_ADMIN_ROLE")
}

// UnpackDEFAULTADMINROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (buyBot *BuyBot) UnpackDEFAULTADMINROLE(data []byte) ([32]byte, error) {
	out, err := buyBot.abi.Unpack("DEFAULT_ADMIN_ROLE", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackMANAGERROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xec87621c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (buyBot *BuyBot) PackMANAGERROLE() []byte {
	enc, err := buyBot.abi.Pack("MANAGER_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMANAGERROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xec87621c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (buyBot *BuyBot) TryPackMANAGERROLE() ([]byte, error) {
	return buyBot.abi.Pack("MANAGER_ROLE")
}

// UnpackMANAGERROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (buyBot *BuyBot) UnpackMANAGERROLE(data []byte) ([32]byte, error) {
	out, err := buyBot.abi.Unpack("MANAGER_ROLE", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackNATIVECOIN is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x04a41159.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function NATIVE_COIN() view returns(address)
func (buyBot *BuyBot) PackNATIVECOIN() []byte {
	enc, err := buyBot.abi.Pack("NATIVE_COIN")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackNATIVECOIN is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x04a41159.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function NATIVE_COIN() view returns(address)
func (buyBot *BuyBot) TryPackNATIVECOIN() ([]byte, error) {
	return buyBot.abi.Pack("NATIVE_COIN")
}

// UnpackNATIVECOIN is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x04a41159.
//
// Solidity: function NATIVE_COIN() view returns(address)
func (buyBot *BuyBot) UnpackNATIVECOIN(data []byte) (common.Address, error) {
	out, err := buyBot.abi.Unpack("NATIVE_COIN", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackAcceptDefaultAdminTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcefc1429.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function acceptDefaultAdminTransfer() returns()
func (buyBot *BuyBot) PackAcceptDefaultAdminTransfer() []byte {
	enc, err := buyBot.abi.Pack("acceptDefaultAdminTransfer")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAcceptDefaultAdminTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcefc1429.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function acceptDefaultAdminTransfer() returns()
func (buyBot *BuyBot) TryPackAcceptDefaultAdminTransfer() ([]byte, error) {
	return buyBot.abi.Pack("acceptDefaultAdminTransfer")
}

// PackBeginDefaultAdminTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x634e93da.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function beginDefaultAdminTransfer(address newAdmin) returns()
func (buyBot *BuyBot) PackBeginDefaultAdminTransfer(newAdmin common.Address) []byte {
	enc, err := buyBot.abi.Pack("beginDefaultAdminTransfer", newAdmin)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBeginDefaultAdminTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x634e93da.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function beginDefaultAdminTransfer(address newAdmin) returns()
func (buyBot *BuyBot) TryPackBeginDefaultAdminTransfer(newAdmin common.Address) ([]byte, error) {
	return buyBot.abi.Pack("beginDefaultAdminTransfer", newAdmin)
}

// PackBuyMarket is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x67ea88eb.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function buyMarket(address pair, uint256 amount, uint256 maxMatchCount) returns()
func (buyBot *BuyBot) PackBuyMarket(pair common.Address, amount *big.Int, maxMatchCount *big.Int) []byte {
	enc, err := buyBot.abi.Pack("buyMarket", pair, amount, maxMatchCount)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBuyMarket is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x67ea88eb.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function buyMarket(address pair, uint256 amount, uint256 maxMatchCount) returns()
func (buyBot *BuyBot) TryPackBuyMarket(pair common.Address, amount *big.Int, maxMatchCount *big.Int) ([]byte, error) {
	return buyBot.abi.Pack("buyMarket", pair, amount, maxMatchCount)
}

// PackCanBuyMarket is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2aaa9628.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function canBuyMarket(address pair, address caller) view returns(bool canBuy, uint256 balance)
func (buyBot *BuyBot) PackCanBuyMarket(pair common.Address, caller common.Address) []byte {
	enc, err := buyBot.abi.Pack("canBuyMarket", pair, caller)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCanBuyMarket is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2aaa9628.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function canBuyMarket(address pair, address caller) view returns(bool canBuy, uint256 balance)
func (buyBot *BuyBot) TryPackCanBuyMarket(pair common.Address, caller common.Address) ([]byte, error) {
	return buyBot.abi.Pack("canBuyMarket", pair, caller)
}

// CanBuyMarketOutput serves as a container for the return parameters of contract
// method CanBuyMarket.
type CanBuyMarketOutput struct {
	CanBuy  bool
	Balance *big.Int
}

// UnpackCanBuyMarket is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2aaa9628.
//
// Solidity: function canBuyMarket(address pair, address caller) view returns(bool canBuy, uint256 balance)
func (buyBot *BuyBot) UnpackCanBuyMarket(data []byte) (CanBuyMarketOutput, error) {
	out, err := buyBot.abi.Unpack("canBuyMarket", data)
	outstruct := new(CanBuyMarketOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.CanBuy = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Balance = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	return *outstruct, nil
}

// PackCancelDefaultAdminTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd602b9fd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (buyBot *BuyBot) PackCancelDefaultAdminTransfer() []byte {
	enc, err := buyBot.abi.Pack("cancelDefaultAdminTransfer")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCancelDefaultAdminTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd602b9fd.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (buyBot *BuyBot) TryPackCancelDefaultAdminTransfer() ([]byte, error) {
	return buyBot.abi.Pack("cancelDefaultAdminTransfer")
}

// PackChangeDefaultAdminDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x649a5ec7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (buyBot *BuyBot) PackChangeDefaultAdminDelay(newDelay *big.Int) []byte {
	enc, err := buyBot.abi.Pack("changeDefaultAdminDelay", newDelay)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackChangeDefaultAdminDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x649a5ec7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (buyBot *BuyBot) TryPackChangeDefaultAdminDelay(newDelay *big.Int) ([]byte, error) {
	return buyBot.abi.Pack("changeDefaultAdminDelay", newDelay)
}

// PackDefaultAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x84ef8ffc.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function defaultAdmin() view returns(address)
func (buyBot *BuyBot) PackDefaultAdmin() []byte {
	enc, err := buyBot.abi.Pack("defaultAdmin")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDefaultAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x84ef8ffc.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function defaultAdmin() view returns(address)
func (buyBot *BuyBot) TryPackDefaultAdmin() ([]byte, error) {
	return buyBot.abi.Pack("defaultAdmin")
}

// UnpackDefaultAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (buyBot *BuyBot) UnpackDefaultAdmin(data []byte) (common.Address, error) {
	out, err := buyBot.abi.Unpack("defaultAdmin", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackDefaultAdminDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcc8463c8.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (buyBot *BuyBot) PackDefaultAdminDelay() []byte {
	enc, err := buyBot.abi.Pack("defaultAdminDelay")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDefaultAdminDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcc8463c8.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (buyBot *BuyBot) TryPackDefaultAdminDelay() ([]byte, error) {
	return buyBot.abi.Pack("defaultAdminDelay")
}

// UnpackDefaultAdminDelay is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (buyBot *BuyBot) UnpackDefaultAdminDelay(data []byte) (*big.Int, error) {
	out, err := buyBot.abi.Unpack("defaultAdminDelay", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackDefaultAdminDelayIncreaseWait is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x022d63fb.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (buyBot *BuyBot) PackDefaultAdminDelayIncreaseWait() []byte {
	enc, err := buyBot.abi.Pack("defaultAdminDelayIncreaseWait")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDefaultAdminDelayIncreaseWait is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x022d63fb.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (buyBot *BuyBot) TryPackDefaultAdminDelayIncreaseWait() ([]byte, error) {
	return buyBot.abi.Pack("defaultAdminDelayIncreaseWait")
}

// UnpackDefaultAdminDelayIncreaseWait is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (buyBot *BuyBot) UnpackDefaultAdminDelayIncreaseWait(data []byte) (*big.Int, error) {
	out, err := buyBot.abi.Unpack("defaultAdminDelayIncreaseWait", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetBalance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf8b2cb4f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getBalance(address token) view returns(uint256 balance)
func (buyBot *BuyBot) PackGetBalance(token common.Address) []byte {
	enc, err := buyBot.abi.Pack("getBalance", token)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetBalance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf8b2cb4f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getBalance(address token) view returns(uint256 balance)
func (buyBot *BuyBot) TryPackGetBalance(token common.Address) ([]byte, error) {
	return buyBot.abi.Pack("getBalance", token)
}

// UnpackGetBalance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf8b2cb4f.
//
// Solidity: function getBalance(address token) view returns(uint256 balance)
func (buyBot *BuyBot) UnpackGetBalance(data []byte) (*big.Int, error) {
	out, err := buyBot.abi.Unpack("getBalance", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetRoleAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x248a9ca3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (buyBot *BuyBot) PackGetRoleAdmin(role [32]byte) []byte {
	enc, err := buyBot.abi.Pack("getRoleAdmin", role)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetRoleAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x248a9ca3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (buyBot *BuyBot) TryPackGetRoleAdmin(role [32]byte) ([]byte, error) {
	return buyBot.abi.Pack("getRoleAdmin", role)
}

// UnpackGetRoleAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (buyBot *BuyBot) UnpackGetRoleAdmin(data []byte) ([32]byte, error) {
	out, err := buyBot.abi.Unpack("getRoleAdmin", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackGrantRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2f2ff15d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (buyBot *BuyBot) PackGrantRole(role [32]byte, account common.Address) []byte {
	enc, err := buyBot.abi.Pack("grantRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGrantRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2f2ff15d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (buyBot *BuyBot) TryPackGrantRole(role [32]byte, account common.Address) ([]byte, error) {
	return buyBot.abi.Pack("grantRole", role, account)
}

// PackHasRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x91d14854.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (buyBot *BuyBot) PackHasRole(role [32]byte, account common.Address) []byte {
	enc, err := buyBot.abi.Pack("hasRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackHasRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x91d14854.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (buyBot *BuyBot) TryPackHasRole(role [32]byte, account common.Address) ([]byte, error) {
	return buyBot.abi.Pack("hasRole", role, account)
}

// UnpackHasRole is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (buyBot *BuyBot) UnpackHasRole(data []byte) (bool, error) {
	out, err := buyBot.abi.Unpack("hasRole", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackInterval is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x947a36fb.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function interval() view returns(uint256)
func (buyBot *BuyBot) PackInterval() []byte {
	enc, err := buyBot.abi.Pack("interval")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackInterval is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x947a36fb.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function interval() view returns(uint256)
func (buyBot *BuyBot) TryPackInterval() ([]byte, error) {
	return buyBot.abi.Pack("interval")
}

// UnpackInterval is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x947a36fb.
//
// Solidity: function interval() view returns(uint256)
func (buyBot *BuyBot) UnpackInterval(data []byte) (*big.Int, error) {
	out, err := buyBot.abi.Unpack("interval", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackLastBuyTime is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf29f4d0b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function lastBuyTime() view returns(uint256)
func (buyBot *BuyBot) PackLastBuyTime() []byte {
	enc, err := buyBot.abi.Pack("lastBuyTime")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackLastBuyTime is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf29f4d0b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function lastBuyTime() view returns(uint256)
func (buyBot *BuyBot) TryPackLastBuyTime() ([]byte, error) {
	return buyBot.abi.Pack("lastBuyTime")
}

// UnpackLastBuyTime is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf29f4d0b.
//
// Solidity: function lastBuyTime() view returns(uint256)
func (buyBot *BuyBot) UnpackLastBuyTime(data []byte) (*big.Int, error) {
	out, err := buyBot.abi.Unpack("lastBuyTime", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackMinOrderAmount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x46b62c4a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function minOrderAmount() view returns(uint256)
func (buyBot *BuyBot) PackMinOrderAmount() []byte {
	enc, err := buyBot.abi.Pack("minOrderAmount")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMinOrderAmount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x46b62c4a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function minOrderAmount() view returns(uint256)
func (buyBot *BuyBot) TryPackMinOrderAmount() ([]byte, error) {
	return buyBot.abi.Pack("minOrderAmount")
}

// UnpackMinOrderAmount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x46b62c4a.
//
// Solidity: function minOrderAmount() view returns(uint256)
func (buyBot *BuyBot) UnpackMinOrderAmount(data []byte) (*big.Int, error) {
	out, err := buyBot.abi.Unpack("minOrderAmount", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function owner() view returns(address)
func (buyBot *BuyBot) PackOwner() []byte {
	enc, err := buyBot.abi.Pack("owner")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function owner() view returns(address)
func (buyBot *BuyBot) TryPackOwner() ([]byte, error) {
	return buyBot.abi.Pack("owner")
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (buyBot *BuyBot) UnpackOwner(data []byte) (common.Address, error) {
	out, err := buyBot.abi.Unpack("owner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackPendingDefaultAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcf6eefb7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (buyBot *BuyBot) PackPendingDefaultAdmin() []byte {
	enc, err := buyBot.abi.Pack("pendingDefaultAdmin")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPendingDefaultAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcf6eefb7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (buyBot *BuyBot) TryPackPendingDefaultAdmin() ([]byte, error) {
	return buyBot.abi.Pack("pendingDefaultAdmin")
}

// PendingDefaultAdminOutput serves as a container for the return parameters of contract
// method PendingDefaultAdmin.

// UnpackPendingDefaultAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (buyBot *BuyBot) UnpackPendingDefaultAdmin(data []byte) (PendingDefaultAdminOutput, error) {
	out, err := buyBot.abi.Unpack("pendingDefaultAdmin", data)
	outstruct := new(PendingDefaultAdminOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.NewAdmin = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Schedule = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	return *outstruct, nil
}

// PackPendingDefaultAdminDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa1eda53c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (buyBot *BuyBot) PackPendingDefaultAdminDelay() []byte {
	enc, err := buyBot.abi.Pack("pendingDefaultAdminDelay")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPendingDefaultAdminDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa1eda53c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (buyBot *BuyBot) TryPackPendingDefaultAdminDelay() ([]byte, error) {
	return buyBot.abi.Pack("pendingDefaultAdminDelay")
}

// PendingDefaultAdminDelayOutput serves as a container for the return parameters of contract
// method PendingDefaultAdminDelay.

// UnpackPendingDefaultAdminDelay is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (buyBot *BuyBot) UnpackPendingDefaultAdminDelay(data []byte) (PendingDefaultAdminDelayOutput, error) {
	out, err := buyBot.abi.Unpack("pendingDefaultAdminDelay", data)
	outstruct := new(PendingDefaultAdminDelayOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.NewDelay = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.Schedule = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	return *outstruct, nil
}

// PackRecipient is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x66d003ac.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function recipient() view returns(address)
func (buyBot *BuyBot) PackRecipient() []byte {
	enc, err := buyBot.abi.Pack("recipient")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRecipient is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x66d003ac.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function recipient() view returns(address)
func (buyBot *BuyBot) TryPackRecipient() ([]byte, error) {
	return buyBot.abi.Pack("recipient")
}

// UnpackRecipient is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x66d003ac.
//
// Solidity: function recipient() view returns(address)
func (buyBot *BuyBot) UnpackRecipient(data []byte) (common.Address, error) {
	out, err := buyBot.abi.Unpack("recipient", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackRenounceRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x36568abe.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (buyBot *BuyBot) PackRenounceRole(role [32]byte, account common.Address) []byte {
	enc, err := buyBot.abi.Pack("renounceRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRenounceRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x36568abe.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (buyBot *BuyBot) TryPackRenounceRole(role [32]byte, account common.Address) ([]byte, error) {
	return buyBot.abi.Pack("renounceRole", role, account)
}

// PackRevokeRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd547741f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (buyBot *BuyBot) PackRevokeRole(role [32]byte, account common.Address) []byte {
	enc, err := buyBot.abi.Pack("revokeRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRevokeRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd547741f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (buyBot *BuyBot) TryPackRevokeRole(role [32]byte, account common.Address) ([]byte, error) {
	return buyBot.abi.Pack("revokeRole", role, account)
}

// PackRollbackDefaultAdminDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0aa6220b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (buyBot *BuyBot) PackRollbackDefaultAdminDelay() []byte {
	enc, err := buyBot.abi.Pack("rollbackDefaultAdminDelay")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRollbackDefaultAdminDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0aa6220b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (buyBot *BuyBot) TryPackRollbackDefaultAdminDelay() ([]byte, error) {
	return buyBot.abi.Pack("rollbackDefaultAdminDelay")
}

// PackRouter is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf887ea40.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function router() view returns(address)
func (buyBot *BuyBot) PackRouter() []byte {
	enc, err := buyBot.abi.Pack("router")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRouter is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf887ea40.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function router() view returns(address)
func (buyBot *BuyBot) TryPackRouter() ([]byte, error) {
	return buyBot.abi.Pack("router")
}

// UnpackRouter is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (buyBot *BuyBot) UnpackRouter(data []byte) (common.Address, error) {
	out, err := buyBot.abi.Unpack("router", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackSetInterval is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x22a90082.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setInterval(uint256 _interval) returns()
func (buyBot *BuyBot) PackSetInterval(interval *big.Int) []byte {
	enc, err := buyBot.abi.Pack("setInterval", interval)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetInterval is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x22a90082.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setInterval(uint256 _interval) returns()
func (buyBot *BuyBot) TryPackSetInterval(interval *big.Int) ([]byte, error) {
	return buyBot.abi.Pack("setInterval", interval)
}

// PackSetMinOrderAmount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa3b8ef04.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setMinOrderAmount(uint256 _minOrderAmount) returns()
func (buyBot *BuyBot) PackSetMinOrderAmount(minOrderAmount *big.Int) []byte {
	enc, err := buyBot.abi.Pack("setMinOrderAmount", minOrderAmount)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetMinOrderAmount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa3b8ef04.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setMinOrderAmount(uint256 _minOrderAmount) returns()
func (buyBot *BuyBot) TryPackSetMinOrderAmount(minOrderAmount *big.Int) ([]byte, error) {
	return buyBot.abi.Pack("setMinOrderAmount", minOrderAmount)
}

// PackSetRecipient is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3bbed4a0.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setRecipient(address _recipient) returns()
func (buyBot *BuyBot) PackSetRecipient(recipient common.Address) []byte {
	enc, err := buyBot.abi.Pack("setRecipient", recipient)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetRecipient is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3bbed4a0.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setRecipient(address _recipient) returns()
func (buyBot *BuyBot) TryPackSetRecipient(recipient common.Address) ([]byte, error) {
	return buyBot.abi.Pack("setRecipient", recipient)
}

// PackSetSwapPool is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0f3fffbf.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setSwapPool(address tokenIn, address tokenOut, address pool) returns()
func (buyBot *BuyBot) PackSetSwapPool(tokenIn common.Address, tokenOut common.Address, pool common.Address) []byte {
	enc, err := buyBot.abi.Pack("setSwapPool", tokenIn, tokenOut, pool)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetSwapPool is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0f3fffbf.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setSwapPool(address tokenIn, address tokenOut, address pool) returns()
func (buyBot *BuyBot) TryPackSetSwapPool(tokenIn common.Address, tokenOut common.Address, pool common.Address) ([]byte, error) {
	return buyBot.abi.Pack("setSwapPool", tokenIn, tokenOut, pool)
}

// PackSetSwapRouter is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x41273657.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setSwapRouter(address _swapRouter) returns()
func (buyBot *BuyBot) PackSetSwapRouter(swapRouter common.Address) []byte {
	enc, err := buyBot.abi.Pack("setSwapRouter", swapRouter)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetSwapRouter is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x41273657.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setSwapRouter(address _swapRouter) returns()
func (buyBot *BuyBot) TryPackSetSwapRouter(swapRouter common.Address) ([]byte, error) {
	return buyBot.abi.Pack("setSwapRouter", swapRouter)
}

// PackSetSwapToken is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb851b7ca.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setSwapToken(address _swapToken) returns()
func (buyBot *BuyBot) PackSetSwapToken(swapToken common.Address) []byte {
	enc, err := buyBot.abi.Pack("setSwapToken", swapToken)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetSwapToken is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb851b7ca.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setSwapToken(address _swapToken) returns()
func (buyBot *BuyBot) TryPackSetSwapToken(swapToken common.Address) ([]byte, error) {
	return buyBot.abi.Pack("setSwapToken", swapToken)
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (buyBot *BuyBot) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := buyBot.abi.Pack("supportsInterface", interfaceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (buyBot *BuyBot) TryPackSupportsInterface(interfaceId [4]byte) ([]byte, error) {
	return buyBot.abi.Pack("supportsInterface", interfaceId)
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (buyBot *BuyBot) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := buyBot.abi.Unpack("supportsInterface", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackSwapPools is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8bfe0df1.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function swapPools(address , address ) view returns(address)
func (buyBot *BuyBot) PackSwapPools(arg0 common.Address, arg1 common.Address) []byte {
	enc, err := buyBot.abi.Pack("swapPools", arg0, arg1)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSwapPools is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8bfe0df1.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function swapPools(address , address ) view returns(address)
func (buyBot *BuyBot) TryPackSwapPools(arg0 common.Address, arg1 common.Address) ([]byte, error) {
	return buyBot.abi.Pack("swapPools", arg0, arg1)
}

// UnpackSwapPools is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8bfe0df1.
//
// Solidity: function swapPools(address , address ) view returns(address)
func (buyBot *BuyBot) UnpackSwapPools(data []byte) (common.Address, error) {
	out, err := buyBot.abi.Unpack("swapPools", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackSwapRouter is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc31c9c07.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function swapRouter() view returns(address)
func (buyBot *BuyBot) PackSwapRouter() []byte {
	enc, err := buyBot.abi.Pack("swapRouter")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSwapRouter is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc31c9c07.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function swapRouter() view returns(address)
func (buyBot *BuyBot) TryPackSwapRouter() ([]byte, error) {
	return buyBot.abi.Pack("swapRouter")
}

// UnpackSwapRouter is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc31c9c07.
//
// Solidity: function swapRouter() view returns(address)
func (buyBot *BuyBot) UnpackSwapRouter(data []byte) (common.Address, error) {
	out, err := buyBot.abi.Unpack("swapRouter", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackSwapToQuote is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5e1c17ed.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function swapToQuote(address pair, uint24 uniswapFee, uint256 minAmountOut) returns(uint256 amountOut)
func (buyBot *BuyBot) PackSwapToQuote(pair common.Address, uniswapFee *big.Int, minAmountOut *big.Int) []byte {
	enc, err := buyBot.abi.Pack("swapToQuote", pair, uniswapFee, minAmountOut)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSwapToQuote is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5e1c17ed.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function swapToQuote(address pair, uint24 uniswapFee, uint256 minAmountOut) returns(uint256 amountOut)
func (buyBot *BuyBot) TryPackSwapToQuote(pair common.Address, uniswapFee *big.Int, minAmountOut *big.Int) ([]byte, error) {
	return buyBot.abi.Pack("swapToQuote", pair, uniswapFee, minAmountOut)
}

// UnpackSwapToQuote is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5e1c17ed.
//
// Solidity: function swapToQuote(address pair, uint24 uniswapFee, uint256 minAmountOut) returns(uint256 amountOut)
func (buyBot *BuyBot) UnpackSwapToQuote(data []byte) (*big.Int, error) {
	out, err := buyBot.abi.Unpack("swapToQuote", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackSwapToken is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdc73e49c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function swapToken() view returns(address)
func (buyBot *BuyBot) PackSwapToken() []byte {
	enc, err := buyBot.abi.Pack("swapToken")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSwapToken is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdc73e49c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function swapToken() view returns(address)
func (buyBot *BuyBot) TryPackSwapToken() ([]byte, error) {
	return buyBot.abi.Pack("swapToken")
}

// UnpackSwapToken is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdc73e49c.
//
// Solidity: function swapToken() view returns(address)
func (buyBot *BuyBot) UnpackSwapToken(data []byte) (common.Address, error) {
	out, err := buyBot.abi.Unpack("swapToken", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackWithdraw is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf3fef3a3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function withdraw(address token, uint256 amount) returns()
func (buyBot *BuyBot) PackWithdraw(token common.Address, amount *big.Int) []byte {
	enc, err := buyBot.abi.Pack("withdraw", token, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackWithdraw is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf3fef3a3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function withdraw(address token, uint256 amount) returns()
func (buyBot *BuyBot) TryPackWithdraw(token common.Address, amount *big.Int) ([]byte, error) {
	return buyBot.abi.Pack("withdraw", token, amount)
}

// PackWithdrawETH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf14210a6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function withdrawETH(uint256 amount) returns()
func (buyBot *BuyBot) PackWithdrawETH(amount *big.Int) []byte {
	enc, err := buyBot.abi.Pack("withdrawETH", amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackWithdrawETH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf14210a6.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function withdrawETH(uint256 amount) returns()
func (buyBot *BuyBot) TryPackWithdrawETH(amount *big.Int) ([]byte, error) {
	return buyBot.abi.Pack("withdrawETH", amount)
}

// BuyBotDefaultAdminDelayChangeCanceled represents a DefaultAdminDelayChangeCanceled event raised by the BuyBot contract.
type BuyBotDefaultAdminDelayChangeCanceled struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const BuyBotDefaultAdminDelayChangeCanceledEventName = "DefaultAdminDelayChangeCanceled"

// ContractEventName returns the user-defined event name.
func (BuyBotDefaultAdminDelayChangeCanceled) ContractEventName() string {
	return BuyBotDefaultAdminDelayChangeCanceledEventName
}

// UnpackDefaultAdminDelayChangeCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DefaultAdminDelayChangeCanceled()
func (buyBot *BuyBot) UnpackDefaultAdminDelayChangeCanceledEvent(log *types.Log) (*BuyBotDefaultAdminDelayChangeCanceled, error) {
	event := "DefaultAdminDelayChangeCanceled"
	if len(log.Topics) == 0 || log.Topics[0] != buyBot.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(BuyBotDefaultAdminDelayChangeCanceled)
	if len(log.Data) > 0 {
		if err := buyBot.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range buyBot.abi.Events[event].Inputs {
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

// BuyBotDefaultAdminDelayChangeScheduled represents a DefaultAdminDelayChangeScheduled event raised by the BuyBot contract.
type BuyBotDefaultAdminDelayChangeScheduled struct {
	NewDelay       *big.Int
	EffectSchedule *big.Int
	Raw            *types.Log // Blockchain specific contextual infos
}

const BuyBotDefaultAdminDelayChangeScheduledEventName = "DefaultAdminDelayChangeScheduled"

// ContractEventName returns the user-defined event name.
func (BuyBotDefaultAdminDelayChangeScheduled) ContractEventName() string {
	return BuyBotDefaultAdminDelayChangeScheduledEventName
}

// UnpackDefaultAdminDelayChangeScheduledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DefaultAdminDelayChangeScheduled(uint48 newDelay, uint48 effectSchedule)
func (buyBot *BuyBot) UnpackDefaultAdminDelayChangeScheduledEvent(log *types.Log) (*BuyBotDefaultAdminDelayChangeScheduled, error) {
	event := "DefaultAdminDelayChangeScheduled"
	if len(log.Topics) == 0 || log.Topics[0] != buyBot.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(BuyBotDefaultAdminDelayChangeScheduled)
	if len(log.Data) > 0 {
		if err := buyBot.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range buyBot.abi.Events[event].Inputs {
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

// BuyBotDefaultAdminTransferCanceled represents a DefaultAdminTransferCanceled event raised by the BuyBot contract.
type BuyBotDefaultAdminTransferCanceled struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const BuyBotDefaultAdminTransferCanceledEventName = "DefaultAdminTransferCanceled"

// ContractEventName returns the user-defined event name.
func (BuyBotDefaultAdminTransferCanceled) ContractEventName() string {
	return BuyBotDefaultAdminTransferCanceledEventName
}

// UnpackDefaultAdminTransferCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DefaultAdminTransferCanceled()
func (buyBot *BuyBot) UnpackDefaultAdminTransferCanceledEvent(log *types.Log) (*BuyBotDefaultAdminTransferCanceled, error) {
	event := "DefaultAdminTransferCanceled"
	if len(log.Topics) == 0 || log.Topics[0] != buyBot.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(BuyBotDefaultAdminTransferCanceled)
	if len(log.Data) > 0 {
		if err := buyBot.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range buyBot.abi.Events[event].Inputs {
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

// BuyBotDefaultAdminTransferScheduled represents a DefaultAdminTransferScheduled event raised by the BuyBot contract.
type BuyBotDefaultAdminTransferScheduled struct {
	NewAdmin       common.Address
	AcceptSchedule *big.Int
	Raw            *types.Log // Blockchain specific contextual infos
}

const BuyBotDefaultAdminTransferScheduledEventName = "DefaultAdminTransferScheduled"

// ContractEventName returns the user-defined event name.
func (BuyBotDefaultAdminTransferScheduled) ContractEventName() string {
	return BuyBotDefaultAdminTransferScheduledEventName
}

// UnpackDefaultAdminTransferScheduledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DefaultAdminTransferScheduled(address indexed newAdmin, uint48 acceptSchedule)
func (buyBot *BuyBot) UnpackDefaultAdminTransferScheduledEvent(log *types.Log) (*BuyBotDefaultAdminTransferScheduled, error) {
	event := "DefaultAdminTransferScheduled"
	if len(log.Topics) == 0 || log.Topics[0] != buyBot.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(BuyBotDefaultAdminTransferScheduled)
	if len(log.Data) > 0 {
		if err := buyBot.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range buyBot.abi.Events[event].Inputs {
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

// BuyBotIntervalSet represents a IntervalSet event raised by the BuyBot contract.
type BuyBotIntervalSet struct {
	Before  *big.Int
	Current *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const BuyBotIntervalSetEventName = "IntervalSet"

// ContractEventName returns the user-defined event name.
func (BuyBotIntervalSet) ContractEventName() string {
	return BuyBotIntervalSetEventName
}

// UnpackIntervalSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event IntervalSet(uint256 indexed before, uint256 indexed current)
func (buyBot *BuyBot) UnpackIntervalSetEvent(log *types.Log) (*BuyBotIntervalSet, error) {
	event := "IntervalSet"
	if len(log.Topics) == 0 || log.Topics[0] != buyBot.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(BuyBotIntervalSet)
	if len(log.Data) > 0 {
		if err := buyBot.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range buyBot.abi.Events[event].Inputs {
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

// BuyBotMarketBuyExecuted represents a MarketBuyExecuted event raised by the BuyBot contract.
type BuyBotMarketBuyExecuted struct {
	Pair        common.Address
	QuoteToken  common.Address
	BaseToken   common.Address
	QuoteAmount *big.Int
	Executor    common.Address
	Raw         *types.Log // Blockchain specific contextual infos
}

const BuyBotMarketBuyExecutedEventName = "MarketBuyExecuted"

// ContractEventName returns the user-defined event name.
func (BuyBotMarketBuyExecuted) ContractEventName() string {
	return BuyBotMarketBuyExecutedEventName
}

// UnpackMarketBuyExecutedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MarketBuyExecuted(address indexed pair, address indexed quoteToken, address indexed baseToken, uint256 quoteAmount, address executor)
func (buyBot *BuyBot) UnpackMarketBuyExecutedEvent(log *types.Log) (*BuyBotMarketBuyExecuted, error) {
	event := "MarketBuyExecuted"
	if len(log.Topics) == 0 || log.Topics[0] != buyBot.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(BuyBotMarketBuyExecuted)
	if len(log.Data) > 0 {
		if err := buyBot.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range buyBot.abi.Events[event].Inputs {
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

// BuyBotMinOrderAmountSet represents a MinOrderAmountSet event raised by the BuyBot contract.
type BuyBotMinOrderAmountSet struct {
	Before  *big.Int
	Current *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const BuyBotMinOrderAmountSetEventName = "MinOrderAmountSet"

// ContractEventName returns the user-defined event name.
func (BuyBotMinOrderAmountSet) ContractEventName() string {
	return BuyBotMinOrderAmountSetEventName
}

// UnpackMinOrderAmountSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MinOrderAmountSet(uint256 indexed before, uint256 indexed current)
func (buyBot *BuyBot) UnpackMinOrderAmountSetEvent(log *types.Log) (*BuyBotMinOrderAmountSet, error) {
	event := "MinOrderAmountSet"
	if len(log.Topics) == 0 || log.Topics[0] != buyBot.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(BuyBotMinOrderAmountSet)
	if len(log.Data) > 0 {
		if err := buyBot.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range buyBot.abi.Events[event].Inputs {
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

// BuyBotRecipientSet represents a RecipientSet event raised by the BuyBot contract.
type BuyBotRecipientSet struct {
	Before  common.Address
	Current common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const BuyBotRecipientSetEventName = "RecipientSet"

// ContractEventName returns the user-defined event name.
func (BuyBotRecipientSet) ContractEventName() string {
	return BuyBotRecipientSetEventName
}

// UnpackRecipientSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RecipientSet(address indexed before, address indexed current)
func (buyBot *BuyBot) UnpackRecipientSetEvent(log *types.Log) (*BuyBotRecipientSet, error) {
	event := "RecipientSet"
	if len(log.Topics) == 0 || log.Topics[0] != buyBot.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(BuyBotRecipientSet)
	if len(log.Data) > 0 {
		if err := buyBot.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range buyBot.abi.Events[event].Inputs {
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

// BuyBotRoleAdminChanged represents a RoleAdminChanged event raised by the BuyBot contract.
type BuyBotRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               *types.Log // Blockchain specific contextual infos
}

const BuyBotRoleAdminChangedEventName = "RoleAdminChanged"

// ContractEventName returns the user-defined event name.
func (BuyBotRoleAdminChanged) ContractEventName() string {
	return BuyBotRoleAdminChangedEventName
}

// UnpackRoleAdminChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (buyBot *BuyBot) UnpackRoleAdminChangedEvent(log *types.Log) (*BuyBotRoleAdminChanged, error) {
	event := "RoleAdminChanged"
	if len(log.Topics) == 0 || log.Topics[0] != buyBot.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(BuyBotRoleAdminChanged)
	if len(log.Data) > 0 {
		if err := buyBot.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range buyBot.abi.Events[event].Inputs {
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

// BuyBotRoleGranted represents a RoleGranted event raised by the BuyBot contract.
type BuyBotRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const BuyBotRoleGrantedEventName = "RoleGranted"

// ContractEventName returns the user-defined event name.
func (BuyBotRoleGranted) ContractEventName() string {
	return BuyBotRoleGrantedEventName
}

// UnpackRoleGrantedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (buyBot *BuyBot) UnpackRoleGrantedEvent(log *types.Log) (*BuyBotRoleGranted, error) {
	event := "RoleGranted"
	if len(log.Topics) == 0 || log.Topics[0] != buyBot.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(BuyBotRoleGranted)
	if len(log.Data) > 0 {
		if err := buyBot.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range buyBot.abi.Events[event].Inputs {
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

// BuyBotRoleRevoked represents a RoleRevoked event raised by the BuyBot contract.
type BuyBotRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const BuyBotRoleRevokedEventName = "RoleRevoked"

// ContractEventName returns the user-defined event name.
func (BuyBotRoleRevoked) ContractEventName() string {
	return BuyBotRoleRevokedEventName
}

// UnpackRoleRevokedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (buyBot *BuyBot) UnpackRoleRevokedEvent(log *types.Log) (*BuyBotRoleRevoked, error) {
	event := "RoleRevoked"
	if len(log.Topics) == 0 || log.Topics[0] != buyBot.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(BuyBotRoleRevoked)
	if len(log.Data) > 0 {
		if err := buyBot.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range buyBot.abi.Events[event].Inputs {
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

// BuyBotSwapExecuted represents a SwapExecuted event raised by the BuyBot contract.
type BuyBotSwapExecuted struct {
	TokenIn   common.Address
	TokenOut  common.Address
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const BuyBotSwapExecutedEventName = "SwapExecuted"

// ContractEventName returns the user-defined event name.
func (BuyBotSwapExecuted) ContractEventName() string {
	return BuyBotSwapExecutedEventName
}

// UnpackSwapExecutedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event SwapExecuted(address indexed tokenIn, address indexed tokenOut, uint256 amountIn, uint256 amountOut)
func (buyBot *BuyBot) UnpackSwapExecutedEvent(log *types.Log) (*BuyBotSwapExecuted, error) {
	event := "SwapExecuted"
	if len(log.Topics) == 0 || log.Topics[0] != buyBot.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(BuyBotSwapExecuted)
	if len(log.Data) > 0 {
		if err := buyBot.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range buyBot.abi.Events[event].Inputs {
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

// BuyBotSwapPoolSet represents a SwapPoolSet event raised by the BuyBot contract.
type BuyBotSwapPoolSet struct {
	TokenIn  common.Address
	TokenOut common.Address
	Pool     common.Address
	Raw      *types.Log // Blockchain specific contextual infos
}

const BuyBotSwapPoolSetEventName = "SwapPoolSet"

// ContractEventName returns the user-defined event name.
func (BuyBotSwapPoolSet) ContractEventName() string {
	return BuyBotSwapPoolSetEventName
}

// UnpackSwapPoolSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event SwapPoolSet(address indexed tokenIn, address indexed tokenOut, address indexed pool)
func (buyBot *BuyBot) UnpackSwapPoolSetEvent(log *types.Log) (*BuyBotSwapPoolSet, error) {
	event := "SwapPoolSet"
	if len(log.Topics) == 0 || log.Topics[0] != buyBot.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(BuyBotSwapPoolSet)
	if len(log.Data) > 0 {
		if err := buyBot.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range buyBot.abi.Events[event].Inputs {
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

// BuyBotSwapRouterSet represents a SwapRouterSet event raised by the BuyBot contract.
type BuyBotSwapRouterSet struct {
	Before  common.Address
	Current common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const BuyBotSwapRouterSetEventName = "SwapRouterSet"

// ContractEventName returns the user-defined event name.
func (BuyBotSwapRouterSet) ContractEventName() string {
	return BuyBotSwapRouterSetEventName
}

// UnpackSwapRouterSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event SwapRouterSet(address before, address current)
func (buyBot *BuyBot) UnpackSwapRouterSetEvent(log *types.Log) (*BuyBotSwapRouterSet, error) {
	event := "SwapRouterSet"
	if len(log.Topics) == 0 || log.Topics[0] != buyBot.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(BuyBotSwapRouterSet)
	if len(log.Data) > 0 {
		if err := buyBot.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range buyBot.abi.Events[event].Inputs {
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

// BuyBotSwapTokenSet represents a SwapTokenSet event raised by the BuyBot contract.
type BuyBotSwapTokenSet struct {
	Before  common.Address
	Current common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const BuyBotSwapTokenSetEventName = "SwapTokenSet"

// ContractEventName returns the user-defined event name.
func (BuyBotSwapTokenSet) ContractEventName() string {
	return BuyBotSwapTokenSetEventName
}

// UnpackSwapTokenSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event SwapTokenSet(address indexed before, address indexed current)
func (buyBot *BuyBot) UnpackSwapTokenSetEvent(log *types.Log) (*BuyBotSwapTokenSet, error) {
	event := "SwapTokenSet"
	if len(log.Topics) == 0 || log.Topics[0] != buyBot.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(BuyBotSwapTokenSet)
	if len(log.Data) > 0 {
		if err := buyBot.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range buyBot.abi.Events[event].Inputs {
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

// BuyBotWithdrawn represents a Withdrawn event raised by the BuyBot contract.
type BuyBotWithdrawn struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Raw    *types.Log // Blockchain specific contextual infos
}

const BuyBotWithdrawnEventName = "Withdrawn"

// ContractEventName returns the user-defined event name.
func (BuyBotWithdrawn) ContractEventName() string {
	return BuyBotWithdrawnEventName
}

// UnpackWithdrawnEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Withdrawn(address indexed token, address indexed to, uint256 amount)
func (buyBot *BuyBot) UnpackWithdrawnEvent(log *types.Log) (*BuyBotWithdrawn, error) {
	event := "Withdrawn"
	if len(log.Topics) == 0 || log.Topics[0] != buyBot.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(BuyBotWithdrawn)
	if len(log.Data) > 0 {
		if err := buyBot.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range buyBot.abi.Events[event].Inputs {
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
func (buyBot *BuyBot) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], buyBot.abi.Errors["AccessControlBadConfirmation"].ID.Bytes()[:4]) {
		return buyBot.UnpackAccessControlBadConfirmationError(raw[4:])
	}
	if bytes.Equal(raw[:4], buyBot.abi.Errors["AccessControlEnforcedDefaultAdminDelay"].ID.Bytes()[:4]) {
		return buyBot.UnpackAccessControlEnforcedDefaultAdminDelayError(raw[4:])
	}
	if bytes.Equal(raw[:4], buyBot.abi.Errors["AccessControlEnforcedDefaultAdminRules"].ID.Bytes()[:4]) {
		return buyBot.UnpackAccessControlEnforcedDefaultAdminRulesError(raw[4:])
	}
	if bytes.Equal(raw[:4], buyBot.abi.Errors["AccessControlInvalidDefaultAdmin"].ID.Bytes()[:4]) {
		return buyBot.UnpackAccessControlInvalidDefaultAdminError(raw[4:])
	}
	if bytes.Equal(raw[:4], buyBot.abi.Errors["AccessControlUnauthorizedAccount"].ID.Bytes()[:4]) {
		return buyBot.UnpackAccessControlUnauthorizedAccountError(raw[4:])
	}
	if bytes.Equal(raw[:4], buyBot.abi.Errors["BuyBotAddressNotChanged"].ID.Bytes()[:4]) {
		return buyBot.UnpackBuyBotAddressNotChangedError(raw[4:])
	}
	if bytes.Equal(raw[:4], buyBot.abi.Errors["BuyBotETHTransferFailed"].ID.Bytes()[:4]) {
		return buyBot.UnpackBuyBotETHTransferFailedError(raw[4:])
	}
	if bytes.Equal(raw[:4], buyBot.abi.Errors["BuyBotInsufficientBalance"].ID.Bytes()[:4]) {
		return buyBot.UnpackBuyBotInsufficientBalanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], buyBot.abi.Errors["BuyBotInsufficientSwapBalance"].ID.Bytes()[:4]) {
		return buyBot.UnpackBuyBotInsufficientSwapBalanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], buyBot.abi.Errors["BuyBotInsufficientWithdrawBalance"].ID.Bytes()[:4]) {
		return buyBot.UnpackBuyBotInsufficientWithdrawBalanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], buyBot.abi.Errors["BuyBotIntervalNotPassed"].ID.Bytes()[:4]) {
		return buyBot.UnpackBuyBotIntervalNotPassedError(raw[4:])
	}
	if bytes.Equal(raw[:4], buyBot.abi.Errors["BuyBotInvalidAmount"].ID.Bytes()[:4]) {
		return buyBot.UnpackBuyBotInvalidAmountError(raw[4:])
	}
	if bytes.Equal(raw[:4], buyBot.abi.Errors["BuyBotInvalidBuyer"].ID.Bytes()[:4]) {
		return buyBot.UnpackBuyBotInvalidBuyerError(raw[4:])
	}
	if bytes.Equal(raw[:4], buyBot.abi.Errors["BuyBotInvalidManager"].ID.Bytes()[:4]) {
		return buyBot.UnpackBuyBotInvalidManagerError(raw[4:])
	}
	if bytes.Equal(raw[:4], buyBot.abi.Errors["BuyBotInvalidMinOrderAmount"].ID.Bytes()[:4]) {
		return buyBot.UnpackBuyBotInvalidMinOrderAmountError(raw[4:])
	}
	if bytes.Equal(raw[:4], buyBot.abi.Errors["BuyBotInvalidPair"].ID.Bytes()[:4]) {
		return buyBot.UnpackBuyBotInvalidPairError(raw[4:])
	}
	if bytes.Equal(raw[:4], buyBot.abi.Errors["BuyBotInvalidPool"].ID.Bytes()[:4]) {
		return buyBot.UnpackBuyBotInvalidPoolError(raw[4:])
	}
	if bytes.Equal(raw[:4], buyBot.abi.Errors["BuyBotInvalidPoolTokens"].ID.Bytes()[:4]) {
		return buyBot.UnpackBuyBotInvalidPoolTokensError(raw[4:])
	}
	if bytes.Equal(raw[:4], buyBot.abi.Errors["BuyBotInvalidRouter"].ID.Bytes()[:4]) {
		return buyBot.UnpackBuyBotInvalidRouterError(raw[4:])
	}
	if bytes.Equal(raw[:4], buyBot.abi.Errors["BuyBotInvalidSwapRouter"].ID.Bytes()[:4]) {
		return buyBot.UnpackBuyBotInvalidSwapRouterError(raw[4:])
	}
	if bytes.Equal(raw[:4], buyBot.abi.Errors["BuyBotInvalidTokenAddresses"].ID.Bytes()[:4]) {
		return buyBot.UnpackBuyBotInvalidTokenAddressesError(raw[4:])
	}
	if bytes.Equal(raw[:4], buyBot.abi.Errors["BuyBotNoSwapToken"].ID.Bytes()[:4]) {
		return buyBot.UnpackBuyBotNoSwapTokenError(raw[4:])
	}
	if bytes.Equal(raw[:4], buyBot.abi.Errors["BuyBotNotChanged"].ID.Bytes()[:4]) {
		return buyBot.UnpackBuyBotNotChangedError(raw[4:])
	}
	if bytes.Equal(raw[:4], buyBot.abi.Errors["BuyBotPoolNotFound"].ID.Bytes()[:4]) {
		return buyBot.UnpackBuyBotPoolNotFoundError(raw[4:])
	}
	if bytes.Equal(raw[:4], buyBot.abi.Errors["ReentrancyGuardReentrantCall"].ID.Bytes()[:4]) {
		return buyBot.UnpackReentrancyGuardReentrantCallError(raw[4:])
	}
	if bytes.Equal(raw[:4], buyBot.abi.Errors["SafeCastOverflowedUintDowncast"].ID.Bytes()[:4]) {
		return buyBot.UnpackSafeCastOverflowedUintDowncastError(raw[4:])
	}
	if bytes.Equal(raw[:4], buyBot.abi.Errors["SafeERC20FailedOperation"].ID.Bytes()[:4]) {
		return buyBot.UnpackSafeERC20FailedOperationError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// BuyBotAccessControlBadConfirmation represents a AccessControlBadConfirmation error raised by the BuyBot contract.
type BuyBotAccessControlBadConfirmation struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlBadConfirmation()
func BuyBotAccessControlBadConfirmationErrorID() common.Hash {
	return common.HexToHash("0x6697b23232a647058342c0724fe7c415cab25915b54e5dbc03f233173d37b41c")
}

// UnpackAccessControlBadConfirmationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlBadConfirmation()
func (buyBot *BuyBot) UnpackAccessControlBadConfirmationError(raw []byte) (*BuyBotAccessControlBadConfirmation, error) {
	out := new(BuyBotAccessControlBadConfirmation)
	if err := buyBot.abi.UnpackIntoInterface(out, "AccessControlBadConfirmation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BuyBotAccessControlEnforcedDefaultAdminDelay represents a AccessControlEnforcedDefaultAdminDelay error raised by the BuyBot contract.
type BuyBotAccessControlEnforcedDefaultAdminDelay struct {
	Schedule *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlEnforcedDefaultAdminDelay(uint48 schedule)
func BuyBotAccessControlEnforcedDefaultAdminDelayErrorID() common.Hash {
	return common.HexToHash("0x19ca5ebb8fb33f00e502c9392eddab1501674629178bf69b853cf037aaf4bb5d")
}

// UnpackAccessControlEnforcedDefaultAdminDelayError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlEnforcedDefaultAdminDelay(uint48 schedule)
func (buyBot *BuyBot) UnpackAccessControlEnforcedDefaultAdminDelayError(raw []byte) (*BuyBotAccessControlEnforcedDefaultAdminDelay, error) {
	out := new(BuyBotAccessControlEnforcedDefaultAdminDelay)
	if err := buyBot.abi.UnpackIntoInterface(out, "AccessControlEnforcedDefaultAdminDelay", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BuyBotAccessControlEnforcedDefaultAdminRules represents a AccessControlEnforcedDefaultAdminRules error raised by the BuyBot contract.
type BuyBotAccessControlEnforcedDefaultAdminRules struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlEnforcedDefaultAdminRules()
func BuyBotAccessControlEnforcedDefaultAdminRulesErrorID() common.Hash {
	return common.HexToHash("0x3fc3c27ae3db78c81b8f6e685172134623efa268ee8cd8d54be38ad2a74fc13b")
}

// UnpackAccessControlEnforcedDefaultAdminRulesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlEnforcedDefaultAdminRules()
func (buyBot *BuyBot) UnpackAccessControlEnforcedDefaultAdminRulesError(raw []byte) (*BuyBotAccessControlEnforcedDefaultAdminRules, error) {
	out := new(BuyBotAccessControlEnforcedDefaultAdminRules)
	if err := buyBot.abi.UnpackIntoInterface(out, "AccessControlEnforcedDefaultAdminRules", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BuyBotAccessControlInvalidDefaultAdmin represents a AccessControlInvalidDefaultAdmin error raised by the BuyBot contract.
type BuyBotAccessControlInvalidDefaultAdmin struct {
	DefaultAdmin common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlInvalidDefaultAdmin(address defaultAdmin)
func BuyBotAccessControlInvalidDefaultAdminErrorID() common.Hash {
	return common.HexToHash("0xc22c8022f2a840d6b6a9f113407715f5bbd4e88c1b0dd9434dc00700ba609ed4")
}

// UnpackAccessControlInvalidDefaultAdminError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlInvalidDefaultAdmin(address defaultAdmin)
func (buyBot *BuyBot) UnpackAccessControlInvalidDefaultAdminError(raw []byte) (*BuyBotAccessControlInvalidDefaultAdmin, error) {
	out := new(BuyBotAccessControlInvalidDefaultAdmin)
	if err := buyBot.abi.UnpackIntoInterface(out, "AccessControlInvalidDefaultAdmin", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BuyBotAccessControlUnauthorizedAccount represents a AccessControlUnauthorizedAccount error raised by the BuyBot contract.
type BuyBotAccessControlUnauthorizedAccount struct {
	Account    common.Address
	NeededRole [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlUnauthorizedAccount(address account, bytes32 neededRole)
func BuyBotAccessControlUnauthorizedAccountErrorID() common.Hash {
	return common.HexToHash("0xe2517d3fbfae6f8515ef5ff1ccedc3933ab0cbbda0b492c06eb54ad10ef03b3e")
}

// UnpackAccessControlUnauthorizedAccountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlUnauthorizedAccount(address account, bytes32 neededRole)
func (buyBot *BuyBot) UnpackAccessControlUnauthorizedAccountError(raw []byte) (*BuyBotAccessControlUnauthorizedAccount, error) {
	out := new(BuyBotAccessControlUnauthorizedAccount)
	if err := buyBot.abi.UnpackIntoInterface(out, "AccessControlUnauthorizedAccount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BuyBotBuyBotAddressNotChanged represents a BuyBotAddressNotChanged error raised by the BuyBot contract.
type BuyBotBuyBotAddressNotChanged struct {
	OldValue common.Address
	NewValue common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BuyBotAddressNotChanged(address oldValue, address newValue)
func BuyBotBuyBotAddressNotChangedErrorID() common.Hash {
	return common.HexToHash("0xea9ebefb86ff3ce98cb86bdbab8f04f11a217531e4890628d953bfdb37864d7d")
}

// UnpackBuyBotAddressNotChangedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BuyBotAddressNotChanged(address oldValue, address newValue)
func (buyBot *BuyBot) UnpackBuyBotAddressNotChangedError(raw []byte) (*BuyBotBuyBotAddressNotChanged, error) {
	out := new(BuyBotBuyBotAddressNotChanged)
	if err := buyBot.abi.UnpackIntoInterface(out, "BuyBotAddressNotChanged", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BuyBotBuyBotETHTransferFailed represents a BuyBotETHTransferFailed error raised by the BuyBot contract.
type BuyBotBuyBotETHTransferFailed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BuyBotETHTransferFailed()
func BuyBotBuyBotETHTransferFailedErrorID() common.Hash {
	return common.HexToHash("0x87fd58992384f4c2ca7889c7730b786555394c1a476475e65144e41cc1ef023d")
}

// UnpackBuyBotETHTransferFailedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BuyBotETHTransferFailed()
func (buyBot *BuyBot) UnpackBuyBotETHTransferFailedError(raw []byte) (*BuyBotBuyBotETHTransferFailed, error) {
	out := new(BuyBotBuyBotETHTransferFailed)
	if err := buyBot.abi.UnpackIntoInterface(out, "BuyBotETHTransferFailed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BuyBotBuyBotInsufficientBalance represents a BuyBotInsufficientBalance error raised by the BuyBot contract.
type BuyBotBuyBotInsufficientBalance struct {
	Balance        *big.Int
	MinOrderAmount *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BuyBotInsufficientBalance(uint256 balance, uint256 minOrderAmount)
func BuyBotBuyBotInsufficientBalanceErrorID() common.Hash {
	return common.HexToHash("0x013fafe2906272579304057703df743faf6cd13e4db0801bed26f03bc903bbbc")
}

// UnpackBuyBotInsufficientBalanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BuyBotInsufficientBalance(uint256 balance, uint256 minOrderAmount)
func (buyBot *BuyBot) UnpackBuyBotInsufficientBalanceError(raw []byte) (*BuyBotBuyBotInsufficientBalance, error) {
	out := new(BuyBotBuyBotInsufficientBalance)
	if err := buyBot.abi.UnpackIntoInterface(out, "BuyBotInsufficientBalance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BuyBotBuyBotInsufficientSwapBalance represents a BuyBotInsufficientSwapBalance error raised by the BuyBot contract.
type BuyBotBuyBotInsufficientSwapBalance struct {
	Token   common.Address
	Balance *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BuyBotInsufficientSwapBalance(address token, uint256 balance)
func BuyBotBuyBotInsufficientSwapBalanceErrorID() common.Hash {
	return common.HexToHash("0x1a89d8d42ed3d354336d3e5b11d3863dcf837464a24538b1200b32cd304beef6")
}

// UnpackBuyBotInsufficientSwapBalanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BuyBotInsufficientSwapBalance(address token, uint256 balance)
func (buyBot *BuyBot) UnpackBuyBotInsufficientSwapBalanceError(raw []byte) (*BuyBotBuyBotInsufficientSwapBalance, error) {
	out := new(BuyBotBuyBotInsufficientSwapBalance)
	if err := buyBot.abi.UnpackIntoInterface(out, "BuyBotInsufficientSwapBalance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BuyBotBuyBotInsufficientWithdrawBalance represents a BuyBotInsufficientWithdrawBalance error raised by the BuyBot contract.
type BuyBotBuyBotInsufficientWithdrawBalance struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BuyBotInsufficientWithdrawBalance()
func BuyBotBuyBotInsufficientWithdrawBalanceErrorID() common.Hash {
	return common.HexToHash("0x26f4246a35274076fc9b7b64b90e811d4bfdc4e15297dd47d7c4265df33d31d9")
}

// UnpackBuyBotInsufficientWithdrawBalanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BuyBotInsufficientWithdrawBalance()
func (buyBot *BuyBot) UnpackBuyBotInsufficientWithdrawBalanceError(raw []byte) (*BuyBotBuyBotInsufficientWithdrawBalance, error) {
	out := new(BuyBotBuyBotInsufficientWithdrawBalance)
	if err := buyBot.abi.UnpackIntoInterface(out, "BuyBotInsufficientWithdrawBalance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BuyBotBuyBotIntervalNotPassed represents a BuyBotIntervalNotPassed error raised by the BuyBot contract.
type BuyBotBuyBotIntervalNotPassed struct {
	TimeSinceLastBuy *big.Int
	RequiredInterval *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BuyBotIntervalNotPassed(uint256 timeSinceLastBuy, uint256 requiredInterval)
func BuyBotBuyBotIntervalNotPassedErrorID() common.Hash {
	return common.HexToHash("0x1e379906adfeae49f7702182ee64752eb7616cf6243ee8aeba2ac6c45b360d7e")
}

// UnpackBuyBotIntervalNotPassedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BuyBotIntervalNotPassed(uint256 timeSinceLastBuy, uint256 requiredInterval)
func (buyBot *BuyBot) UnpackBuyBotIntervalNotPassedError(raw []byte) (*BuyBotBuyBotIntervalNotPassed, error) {
	out := new(BuyBotBuyBotIntervalNotPassed)
	if err := buyBot.abi.UnpackIntoInterface(out, "BuyBotIntervalNotPassed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BuyBotBuyBotInvalidAmount represents a BuyBotInvalidAmount error raised by the BuyBot contract.
type BuyBotBuyBotInvalidAmount struct {
	Arg0 *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BuyBotInvalidAmount(uint256 arg0)
func BuyBotBuyBotInvalidAmountErrorID() common.Hash {
	return common.HexToHash("0x9717f35f2b95b5e4602730ee872fc53276f410b1ac152e9b09d3113b05b850af")
}

// UnpackBuyBotInvalidAmountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BuyBotInvalidAmount(uint256 arg0)
func (buyBot *BuyBot) UnpackBuyBotInvalidAmountError(raw []byte) (*BuyBotBuyBotInvalidAmount, error) {
	out := new(BuyBotBuyBotInvalidAmount)
	if err := buyBot.abi.UnpackIntoInterface(out, "BuyBotInvalidAmount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BuyBotBuyBotInvalidBuyer represents a BuyBotInvalidBuyer error raised by the BuyBot contract.
type BuyBotBuyBotInvalidBuyer struct {
	Buyer common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BuyBotInvalidBuyer(address buyer)
func BuyBotBuyBotInvalidBuyerErrorID() common.Hash {
	return common.HexToHash("0x2a7409a9b657ee764ab483d3ea0a280341dc8ef5f15f1811fc36b4aee00e8425")
}

// UnpackBuyBotInvalidBuyerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BuyBotInvalidBuyer(address buyer)
func (buyBot *BuyBot) UnpackBuyBotInvalidBuyerError(raw []byte) (*BuyBotBuyBotInvalidBuyer, error) {
	out := new(BuyBotBuyBotInvalidBuyer)
	if err := buyBot.abi.UnpackIntoInterface(out, "BuyBotInvalidBuyer", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BuyBotBuyBotInvalidManager represents a BuyBotInvalidManager error raised by the BuyBot contract.
type BuyBotBuyBotInvalidManager struct {
	Manager common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BuyBotInvalidManager(address manager)
func BuyBotBuyBotInvalidManagerErrorID() common.Hash {
	return common.HexToHash("0x3242f1baf594d206e10cab2654c7cfcb75ec7a506233b7a0bc0261baa0563400")
}

// UnpackBuyBotInvalidManagerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BuyBotInvalidManager(address manager)
func (buyBot *BuyBot) UnpackBuyBotInvalidManagerError(raw []byte) (*BuyBotBuyBotInvalidManager, error) {
	out := new(BuyBotBuyBotInvalidManager)
	if err := buyBot.abi.UnpackIntoInterface(out, "BuyBotInvalidManager", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BuyBotBuyBotInvalidMinOrderAmount represents a BuyBotInvalidMinOrderAmount error raised by the BuyBot contract.
type BuyBotBuyBotInvalidMinOrderAmount struct {
	Arg0 *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BuyBotInvalidMinOrderAmount(uint256 arg0)
func BuyBotBuyBotInvalidMinOrderAmountErrorID() common.Hash {
	return common.HexToHash("0x517f903233ae5e1183263befde3db719863b7e33a57bec46ed44d9f24e135b83")
}

// UnpackBuyBotInvalidMinOrderAmountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BuyBotInvalidMinOrderAmount(uint256 arg0)
func (buyBot *BuyBot) UnpackBuyBotInvalidMinOrderAmountError(raw []byte) (*BuyBotBuyBotInvalidMinOrderAmount, error) {
	out := new(BuyBotBuyBotInvalidMinOrderAmount)
	if err := buyBot.abi.UnpackIntoInterface(out, "BuyBotInvalidMinOrderAmount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BuyBotBuyBotInvalidPair represents a BuyBotInvalidPair error raised by the BuyBot contract.
type BuyBotBuyBotInvalidPair struct {
	Arg0 common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BuyBotInvalidPair(address arg0)
func BuyBotBuyBotInvalidPairErrorID() common.Hash {
	return common.HexToHash("0x235aafe41858f65cd666c3cfa5b4746f8d4cd2be31eab5fd0a4eaf6fbae6753d")
}

// UnpackBuyBotInvalidPairError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BuyBotInvalidPair(address arg0)
func (buyBot *BuyBot) UnpackBuyBotInvalidPairError(raw []byte) (*BuyBotBuyBotInvalidPair, error) {
	out := new(BuyBotBuyBotInvalidPair)
	if err := buyBot.abi.UnpackIntoInterface(out, "BuyBotInvalidPair", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BuyBotBuyBotInvalidPool represents a BuyBotInvalidPool error raised by the BuyBot contract.
type BuyBotBuyBotInvalidPool struct {
	Pool common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BuyBotInvalidPool(address pool)
func BuyBotBuyBotInvalidPoolErrorID() common.Hash {
	return common.HexToHash("0xebb9ebd73b50d5a50923459f11786b6d277dc91fc8b332f0b82f90b92da948a7")
}

// UnpackBuyBotInvalidPoolError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BuyBotInvalidPool(address pool)
func (buyBot *BuyBot) UnpackBuyBotInvalidPoolError(raw []byte) (*BuyBotBuyBotInvalidPool, error) {
	out := new(BuyBotBuyBotInvalidPool)
	if err := buyBot.abi.UnpackIntoInterface(out, "BuyBotInvalidPool", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BuyBotBuyBotInvalidPoolTokens represents a BuyBotInvalidPoolTokens error raised by the BuyBot contract.
type BuyBotBuyBotInvalidPoolTokens struct {
	Pool     common.Address
	TokenIn  common.Address
	TokenOut common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BuyBotInvalidPoolTokens(address pool, address tokenIn, address tokenOut)
func BuyBotBuyBotInvalidPoolTokensErrorID() common.Hash {
	return common.HexToHash("0xd6979646ac3cf1b94fe33e87bc754e43b64e0abd155fde018d453f368a7b86ab")
}

// UnpackBuyBotInvalidPoolTokensError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BuyBotInvalidPoolTokens(address pool, address tokenIn, address tokenOut)
func (buyBot *BuyBot) UnpackBuyBotInvalidPoolTokensError(raw []byte) (*BuyBotBuyBotInvalidPoolTokens, error) {
	out := new(BuyBotBuyBotInvalidPoolTokens)
	if err := buyBot.abi.UnpackIntoInterface(out, "BuyBotInvalidPoolTokens", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BuyBotBuyBotInvalidRouter represents a BuyBotInvalidRouter error raised by the BuyBot contract.
type BuyBotBuyBotInvalidRouter struct {
	Arg0 common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BuyBotInvalidRouter(address arg0)
func BuyBotBuyBotInvalidRouterErrorID() common.Hash {
	return common.HexToHash("0x5587c9eca80c76d4a2e21651085f445a5307747630a606edb7c80786cb285487")
}

// UnpackBuyBotInvalidRouterError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BuyBotInvalidRouter(address arg0)
func (buyBot *BuyBot) UnpackBuyBotInvalidRouterError(raw []byte) (*BuyBotBuyBotInvalidRouter, error) {
	out := new(BuyBotBuyBotInvalidRouter)
	if err := buyBot.abi.UnpackIntoInterface(out, "BuyBotInvalidRouter", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BuyBotBuyBotInvalidSwapRouter represents a BuyBotInvalidSwapRouter error raised by the BuyBot contract.
type BuyBotBuyBotInvalidSwapRouter struct {
	Arg0 common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BuyBotInvalidSwapRouter(address arg0)
func BuyBotBuyBotInvalidSwapRouterErrorID() common.Hash {
	return common.HexToHash("0x5da5a13d5431c66b52b967c9b530006964a617964287ce75947e93979f41b865")
}

// UnpackBuyBotInvalidSwapRouterError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BuyBotInvalidSwapRouter(address arg0)
func (buyBot *BuyBot) UnpackBuyBotInvalidSwapRouterError(raw []byte) (*BuyBotBuyBotInvalidSwapRouter, error) {
	out := new(BuyBotBuyBotInvalidSwapRouter)
	if err := buyBot.abi.UnpackIntoInterface(out, "BuyBotInvalidSwapRouter", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BuyBotBuyBotInvalidTokenAddresses represents a BuyBotInvalidTokenAddresses error raised by the BuyBot contract.
type BuyBotBuyBotInvalidTokenAddresses struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BuyBotInvalidTokenAddresses()
func BuyBotBuyBotInvalidTokenAddressesErrorID() common.Hash {
	return common.HexToHash("0x5159e7e08228bc5b41d4ef4fd9c4be4405424973e658d3836780aaf2c4af7cdf")
}

// UnpackBuyBotInvalidTokenAddressesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BuyBotInvalidTokenAddresses()
func (buyBot *BuyBot) UnpackBuyBotInvalidTokenAddressesError(raw []byte) (*BuyBotBuyBotInvalidTokenAddresses, error) {
	out := new(BuyBotBuyBotInvalidTokenAddresses)
	if err := buyBot.abi.UnpackIntoInterface(out, "BuyBotInvalidTokenAddresses", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BuyBotBuyBotNoSwapToken represents a BuyBotNoSwapToken error raised by the BuyBot contract.
type BuyBotBuyBotNoSwapToken struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BuyBotNoSwapToken()
func BuyBotBuyBotNoSwapTokenErrorID() common.Hash {
	return common.HexToHash("0xffb5baac980921dc6932054e4b7e5a4298925464ad6a8f227606912aac905cb4")
}

// UnpackBuyBotNoSwapTokenError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BuyBotNoSwapToken()
func (buyBot *BuyBot) UnpackBuyBotNoSwapTokenError(raw []byte) (*BuyBotBuyBotNoSwapToken, error) {
	out := new(BuyBotBuyBotNoSwapToken)
	if err := buyBot.abi.UnpackIntoInterface(out, "BuyBotNoSwapToken", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BuyBotBuyBotNotChanged represents a BuyBotNotChanged error raised by the BuyBot contract.
type BuyBotBuyBotNotChanged struct {
	OldValue *big.Int
	NewValue *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BuyBotNotChanged(uint256 oldValue, uint256 newValue)
func BuyBotBuyBotNotChangedErrorID() common.Hash {
	return common.HexToHash("0x51baf7be0fcc77d47f227de9243e6ad302805ced01d900bf4b509b14a9113cda")
}

// UnpackBuyBotNotChangedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BuyBotNotChanged(uint256 oldValue, uint256 newValue)
func (buyBot *BuyBot) UnpackBuyBotNotChangedError(raw []byte) (*BuyBotBuyBotNotChanged, error) {
	out := new(BuyBotBuyBotNotChanged)
	if err := buyBot.abi.UnpackIntoInterface(out, "BuyBotNotChanged", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BuyBotBuyBotPoolNotFound represents a BuyBotPoolNotFound error raised by the BuyBot contract.
type BuyBotBuyBotPoolNotFound struct {
	TokenIn  common.Address
	TokenOut common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BuyBotPoolNotFound(address tokenIn, address tokenOut)
func BuyBotBuyBotPoolNotFoundErrorID() common.Hash {
	return common.HexToHash("0xa06d7bd88bd48f8c20f6cfa4b74491f49a9982142cff374c8350e65464a2ff48")
}

// UnpackBuyBotPoolNotFoundError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BuyBotPoolNotFound(address tokenIn, address tokenOut)
func (buyBot *BuyBot) UnpackBuyBotPoolNotFoundError(raw []byte) (*BuyBotBuyBotPoolNotFound, error) {
	out := new(BuyBotBuyBotPoolNotFound)
	if err := buyBot.abi.UnpackIntoInterface(out, "BuyBotPoolNotFound", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BuyBotReentrancyGuardReentrantCall represents a ReentrancyGuardReentrantCall error raised by the BuyBot contract.
type BuyBotReentrancyGuardReentrantCall struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ReentrancyGuardReentrantCall()
func BuyBotReentrancyGuardReentrantCallErrorID() common.Hash {
	return common.HexToHash("0x3ee5aeb571de7fc460830b4d0017439a1ca56fb0bc39062227ade4fe4a24c1ca")
}

// UnpackReentrancyGuardReentrantCallError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ReentrancyGuardReentrantCall()
func (buyBot *BuyBot) UnpackReentrancyGuardReentrantCallError(raw []byte) (*BuyBotReentrancyGuardReentrantCall, error) {
	out := new(BuyBotReentrancyGuardReentrantCall)
	if err := buyBot.abi.UnpackIntoInterface(out, "ReentrancyGuardReentrantCall", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BuyBotSafeCastOverflowedUintDowncast represents a SafeCastOverflowedUintDowncast error raised by the BuyBot contract.
type BuyBotSafeCastOverflowedUintDowncast struct {
	Bits  uint8
	Value *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SafeCastOverflowedUintDowncast(uint8 bits, uint256 value)
func BuyBotSafeCastOverflowedUintDowncastErrorID() common.Hash {
	return common.HexToHash("0x6dfcc6503a32754ce7a89698e18201fc5294fd4aad43edefee786f88423b1a12")
}

// UnpackSafeCastOverflowedUintDowncastError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SafeCastOverflowedUintDowncast(uint8 bits, uint256 value)
func (buyBot *BuyBot) UnpackSafeCastOverflowedUintDowncastError(raw []byte) (*BuyBotSafeCastOverflowedUintDowncast, error) {
	out := new(BuyBotSafeCastOverflowedUintDowncast)
	if err := buyBot.abi.UnpackIntoInterface(out, "SafeCastOverflowedUintDowncast", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BuyBotSafeERC20FailedOperation represents a SafeERC20FailedOperation error raised by the BuyBot contract.
type BuyBotSafeERC20FailedOperation struct {
	Token common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SafeERC20FailedOperation(address token)
func BuyBotSafeERC20FailedOperationErrorID() common.Hash {
	return common.HexToHash("0x5274afe73c98b4749fc91ffae6b7b574e7842cb2144a159e9377a5f20b32edf9")
}

// UnpackSafeERC20FailedOperationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SafeERC20FailedOperation(address token)
func (buyBot *BuyBot) UnpackSafeERC20FailedOperationError(raw []byte) (*BuyBotSafeERC20FailedOperation, error) {
	out := new(BuyBotSafeERC20FailedOperation)
	if err := buyBot.abi.UnpackIntoInterface(out, "SafeERC20FailedOperation", raw); err != nil {
		return nil, err
	}
	return out, nil
}
