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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"_initialDelay\",\"type\":\"uint48\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minOrderAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_interval\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"BUYER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NATIVE_COIN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"beginDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxMatchCount\",\"type\":\"uint256\"}],\"name\":\"buyMarket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"canBuyMarket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"canBuy\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"}],\"name\":\"changeDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelayIncreaseWait\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastBuyTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minOrderAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"recipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollbackDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"contractIRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_interval\",\"type\":\"uint256\"}],\"name\":\"setInterval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minOrderAmount\",\"type\":\"uint256\"}],\"name\":\"setMinOrderAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"setRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminDelayChangeCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"effectSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminDelayChangeScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminTransferCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"acceptSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminTransferScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"before\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"current\",\"type\":\"uint256\"}],\"name\":\"IntervalSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"MarketBuyExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"before\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"current\",\"type\":\"uint256\"}],\"name\":\"MinOrderAmountSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"before\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"RecipientSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"name\":\"AccessControlEnforcedDefaultAdminDelay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccessControlEnforcedDefaultAdminRules\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"defaultAdmin\",\"type\":\"address\"}],\"name\":\"AccessControlInvalidDefaultAdmin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minOrderAmount\",\"type\":\"uint256\"}],\"name\":\"BuyBotInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timeSinceLastBuy\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredInterval\",\"type\":\"uint256\"}],\"name\":\"BuyBotIntervalNotPassed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"BuyBotInvalidAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"}],\"name\":\"BuyBotInvalidBuyer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"BuyBotInvalidManager\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"BuyBotInvalidMinOrderAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"BuyBotInvalidPair\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"BuyBotInvalidRouter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"}]",
	ID:  "BuyBot",
	Bin: "0x608060405234801561000f575f5ffd5b50604051612d01380380612d0183398101604081905261002e91610451565b87876001600160a01b03811661005e57604051636116401160e11b81525f60048201526024015b60405180910390fd5b600180546001600160d01b0316600160d01b65ffffffffffff8516021790556100875f826102aa565b50506001600355506001600160a01b0386166100c157604051631561f27b60e21b81526001600160a01b0387166004820152602401610055565b845f036100e4576040516328bfc81960e11b815260048101869052602401610055565b6001600160a01b03821661011657604051632a7409a960e01b81526001600160a01b0383166004820152602401610055565b6001600160a01b0381166101485760405163192178dd60e11b81526001600160a01b0382166004820152602401610055565b600480546001600160a01b038089166001600160a01b0319928316179092556005879055600686905560088054928616929091169190911790556101995f516020612ce15f395f51905f525f610319565b6101b05f516020612cc15f395f51905f525f610319565b6101c75f516020612ce15f395f51905f52886102aa565b506101df5f516020612ce15f395f51905f52836102aa565b506101f75f516020612cc15f395f51905f52886102aa565b5061020f5f516020612cc15f395f51905f52826102aa565b5060405185905f907fd6d62b78d21fc4f5151029ca37079020e8a8815c5a488b28ee7499d4c2a19e55908290a360405184905f907f3fb677206c0b314c404bae3da94bee6bda11375c9fe986f266fb033d6cddbbd7908290a36040516001600160a01b038416905f907fc1416b5cdab50a9fbc872236e1aa54566c6deb40024e63a4b1737ecacf09d6f9908290a350505050505050506104e2565b5f82610306575f6102c36002546001600160a01b031690565b6001600160a01b0316146102ea57604051631fe1e13d60e11b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b0384161790555b6103108383610345565b90505b92915050565b8161033757604051631fe1e13d60e11b815260040160405180910390fd5b61034182826103ec565b5050565b5f828152602081815260408083206001600160a01b038516845290915281205460ff166103e5575f838152602081815260408083206001600160a01b03861684529091529020805460ff1916600117905561039d3390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4506001610313565b505f610313565b5f82815260208190526040808220600101805490849055905190918391839186917fbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff9190a4505050565b80516001600160a01b038116811461044c575f5ffd5b919050565b5f5f5f5f5f5f5f5f610100898b031215610469575f5ffd5b885165ffffffffffff8116811461047e575f5ffd5b975061048c60208a01610436565b965061049a60408a01610436565b60608a015160808b0151919750955093506104b760a08a01610436565b92506104c560c08a01610436565b91506104d360e08a01610436565b90509295985092959890939650565b6127d2806104ef5f395ff3fe608060405260043610610212575f3560e01c80638da5cb5b11610117578063cf6eefb7116100ac578063f14210a61161007c578063f3fef3a311610062578063f3fef3a3146106a8578063f887ea40146106c7578063f8b2cb4f146106f3575f5ffd5b8063f14210a614610674578063f29f4d0b14610693575f5ffd5b8063cf6eefb7146105b6578063d547741f1461060e578063d602b9fd1461062d578063ec87621c14610641575f5ffd5b8063a217fddf116100e7578063a217fddf1461055c578063a3b8ef041461056f578063cc8463c81461058e578063cefc1429146105a2575f5ffd5b80638da5cb5b146104b157806391d14854146104c5578063947a36fb14610514578063a1eda53c14610529575f5ffd5b806336568abe116101a7578063649a5ec71161017757806367ea88eb1161015d57806367ea88eb146104355780637a01a1da1461045457806384ef8ffc14610487575f5ffd5b8063649a5ec7146103ea57806366d003ac14610409575f5ffd5b806336568abe146103785780633bbed4a01461039757806346b62c4a146103b6578063634e93da146103cb575f5ffd5b806322a90082116101e257806322a90082146102c8578063248a9ca3146102e75780632aaa9628146103235780632f2ff15d14610359575f5ffd5b806301ffc9a71461021d578063022d63fb1461025157806304a41159146102795780630aa6220b146102b2575f5ffd5b3661021957005b5f5ffd5b348015610228575f5ffd5b5061023c61023736600461250c565b610712565b60405190151581526020015b60405180910390f35b34801561025c575f5ffd5b50620697805b60405165ffffffffffff9091168152602001610248565b348015610284575f5ffd5b5061028d600181565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610248565b3480156102bd575f5ffd5b506102c661076d565b005b3480156102d3575f5ffd5b506102c66102e236600461254b565b610782565b3480156102f2575f5ffd5b5061031561030136600461254b565b5f9081526020819052604090206001015490565b604051908152602001610248565b34801561032e575f5ffd5b5061034261033d366004612583565b6107e0565b604080519215158352602083019190915201610248565b348015610364575f5ffd5b506102c66103733660046125ba565b6109ba565b348015610383575f5ffd5b506102c66103923660046125ba565b6109ff565b3480156103a2575f5ffd5b506102c66103b13660046125dd565b610b09565b3480156103c1575f5ffd5b5061031560055481565b3480156103d6575f5ffd5b506102c66103e53660046125dd565b610ba1565b3480156103f5575f5ffd5b506102c66104043660046125f8565b610bb4565b348015610414575f5ffd5b5060085461028d9073ffffffffffffffffffffffffffffffffffffffff1681565b348015610440575f5ffd5b506102c661044f36600461261d565b610bc7565b34801561045f575f5ffd5b506103157ff8cd32ed93fc2f9fc78152a14807c9609af3d99c5fe4dc6b106a801aaddfe90e81565b348015610492575f5ffd5b5060025473ffffffffffffffffffffffffffffffffffffffff1661028d565b3480156104bc575f5ffd5b5061028d61123b565b3480156104d0575f5ffd5b5061023c6104df3660046125ba565b5f9182526020828152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b34801561051f575f5ffd5b5061031560065481565b348015610534575f5ffd5b5061053d611260565b6040805165ffffffffffff938416815292909116602083015201610248565b348015610567575f5ffd5b506103155f81565b34801561057a575f5ffd5b506102c661058936600461254b565b6112da565b348015610599575f5ffd5b50610262611374565b3480156105ad575f5ffd5b506102c6611411565b3480156105c1575f5ffd5b506001546040805173ffffffffffffffffffffffffffffffffffffffff831681527401000000000000000000000000000000000000000090920465ffffffffffff16602083015201610248565b348015610619575f5ffd5b506102c66106283660046125ba565b61146d565b348015610638575f5ffd5b506102c66114ae565b34801561064c575f5ffd5b506103157f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b0881565b34801561067f575f5ffd5b506102c661068e36600461254b565b6114c0565b34801561069e575f5ffd5b5061031560075481565b3480156106b3575f5ffd5b506102c66106c236600461264f565b611692565b3480156106d2575f5ffd5b5060045461028d9073ffffffffffffffffffffffffffffffffffffffff1681565b3480156106fe575f5ffd5b5061031561070d3660046125dd565b61187e565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f3149878600000000000000000000000000000000000000000000000000000000148061076757506107678261194f565b92915050565b5f610777816119e5565b61077f6119ef565b50565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b086107ac816119e5565b6006546040518391907f3fb677206c0b314c404bae3da94bee6bda11375c9fe986f266fb033d6cddbbd7905f90a350600655565b5f8073ffffffffffffffffffffffffffffffffffffffff841661080757505f9050806109b3565b73ffffffffffffffffffffffffffffffffffffffff83165f9081527ff3fa603c74bfe2a4719960e47343678c3dc690d2b27a2295acc6fc430833aaf9602052604090205460ff1661085c57505f9050806109b3565b5f8473ffffffffffffffffffffffffffffffffffffffff1663c3f909d46040518163ffffffff1660e01b8152600401606060405180830381865afa1580156108a6573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108ca9190612679565b80516040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015291925073ffffffffffffffffffffffffffffffffffffffff16906370a0823190602401602060405180830381865afa158015610936573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061095a9190612709565b915060055482101561096f575f9250506109b3565b5f60065411801561098157505f600754115b156109ad575f60075442610995919061274d565b90506006548110156109ab575f935050506109b3565b505b60019250505b9250929050565b816109f1576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6109fb82826119fb565b5050565b81158015610a27575060025473ffffffffffffffffffffffffffffffffffffffff8281169116145b15610aff5760015473ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900465ffffffffffff1681151580610a7b575065ffffffffffff8116155b80610a8e57504265ffffffffffff821610155b15610ad4576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff821660048201526024015b60405180910390fd5b5050600180547fffffffffffff000000000000ffffffffffffffffffffffffffffffffffffffff1690555b6109fb8282611a25565b5f610b13816119e5565b60085460405173ffffffffffffffffffffffffffffffffffffffff8085169216907fc1416b5cdab50a9fbc872236e1aa54566c6deb40024e63a4b1737ecacf09d6f9905f90a350600880547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b5f610bab816119e5565b6109fb82611a7e565b5f610bbe816119e5565b6109fb82611afd565b610bcf611b6c565b7ff8cd32ed93fc2f9fc78152a14807c9609af3d99c5fe4dc6b106a801aaddfe90e610bf9816119e5565b73ffffffffffffffffffffffffffffffffffffffff8416610c5e576040517f235aafe400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610acb565b825f03610c9a576040517f9717f35f00000000000000000000000000000000000000000000000000000000815260048101849052602401610acb565b5f600654118015610cac57505f600754115b15610d0e575f60075442610cc0919061274d565b9050600654811015610d0c576006546040517f1e379906000000000000000000000000000000000000000000000000000000008152610acb918391600401918252602082015260400190565b505b5f8473ffffffffffffffffffffffffffffffffffffffff1663c3f909d46040518163ffffffff1660e01b8152600401606060405180830381865afa158015610d58573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d7c9190612679565b805160208201516040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015292935090915f9073ffffffffffffffffffffffffffffffffffffffff8416906370a0823190602401602060405180830381865afa158015610df2573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e169190612709565b9050600554871015610e62576005546040517f013fafe2000000000000000000000000000000000000000000000000000000008152610acb918991600401918252602082015260400190565b80871115610ea6576040517f013fafe20000000000000000000000000000000000000000000000000000000081526004810182905260248101889052604401610acb565b600480546040517fdd62ed3e000000000000000000000000000000000000000000000000000000008152309281019290925273ffffffffffffffffffffffffffffffffffffffff9081166024830181905291899186169063dd62ed3e90604401602060405180830381865afa158015610f21573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f459190612709565b1015610f8c57610f8c73ffffffffffffffffffffffffffffffffffffffff8516827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff611baf565b600480546040517f1e92008400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8c811693820193909352602481018b9052604481018a9052911690631e920084906064015f604051808303815f87803b158015611006575f5ffd5b505af1158015611018573d5f5f3e3d5ffd5b5050604080518b815233602082015273ffffffffffffffffffffffffffffffffffffffff808816945088811693508d16917fc7b4d815bd0a3b8348577971cb55a9334c80a0b3ccdfe287b55b5a6320d8480a910160405180910390a44260075560085473ffffffffffffffffffffffffffffffffffffffff161561122657478015611166576008546040515f9173ffffffffffffffffffffffffffffffffffffffff169083908381818185875af1925050503d805f81146110f4576040519150601f19603f3d011682016040523d82523d5f602084013e6110f9565b606091505b5050905080611164576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f455448207472616e73666572206661696c6564000000000000000000000000006044820152606401610acb565b505b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201525f9073ffffffffffffffffffffffffffffffffffffffff8616906370a0823190602401602060405180830381865afa1580156111d0573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111f49190612709565b90508015611223576008546112239073ffffffffffffffffffffffffffffffffffffffff878116911683611cc9565b50505b5050505050506112366001600355565b505050565b5f61125b60025473ffffffffffffffffffffffffffffffffffffffff1690565b905090565b6002545f907a010000000000000000000000000000000000000000000000000000900465ffffffffffff1680151580156112a257504265ffffffffffff821610155b6112ad575f5f6112d2565b60025474010000000000000000000000000000000000000000900465ffffffffffff16815b915091509091565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b08611304816119e5565b815f03611340576040517f517f903200000000000000000000000000000000000000000000000000000000815260048101839052602401610acb565b6005546040518391907fd6d62b78d21fc4f5151029ca37079020e8a8815c5a488b28ee7499d4c2a19e55905f90a350600555565b6002545f907a010000000000000000000000000000000000000000000000000000900465ffffffffffff1680151580156113b557504265ffffffffffff8216105b6113e7576001547a010000000000000000000000000000000000000000000000000000900465ffffffffffff1661140b565b60025474010000000000000000000000000000000000000000900465ffffffffffff165b91505090565b60015473ffffffffffffffffffffffffffffffffffffffff16338114611465576040517fc22c8022000000000000000000000000000000000000000000000000000000008152336004820152602401610acb565b61077f611d07565b816114a4576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6109fb8282611df8565b5f6114b8816119e5565b61077f611e1c565b6114c8611b6c565b5f6114d2816119e5565b475f83156114e057836114e2565b815b90508181111561154e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f496e73756666696369656e742062616c616e63650000000000000000000000006044820152606401610acb565b5f61156e60025473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff16826040515f6040518083038185875af1925050503d805f81146115c2576040519150601f19603f3d011682016040523d82523d5f602084013e6115c7565b606091505b5050905080611632576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f455448207472616e73666572206661696c6564000000000000000000000000006044820152606401610acb565b60025460405183815273ffffffffffffffffffffffffffffffffffffffff909116906001907fd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb9060200160405180910390a35050505061077f6001600355565b61169a611b6c565b5f6116a4816119e5565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015283905f9073ffffffffffffffffffffffffffffffffffffffff8316906370a0823190602401602060405180830381865afa158015611710573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906117349190612709565b90505f84156117435784611745565b815b9050818111156117b1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f496e73756666696369656e742062616c616e63650000000000000000000000006044820152606401610acb565b6117f16117d360025473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff85169083611cc9565b60025473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff167fd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb8360405161186891815260200190565b60405180910390a3505050506109fb6001600355565b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff73ffffffffffffffffffffffffffffffffffffffff8316016118c3575047919050565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff8316906370a0823190602401602060405180830381865afa15801561192b573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107679190612709565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061076757507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000831614610767565b61077f8133611e26565b6119f95f5f611eab565b565b5f82815260208190526040902060010154611a15816119e5565b611a1f8383612004565b50505050565b73ffffffffffffffffffffffffffffffffffffffff81163314611a74576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61123682826120c9565b5f611a87611374565b611a904261212a565b611a9a9190612760565b9050611aa68282612179565b60405165ffffffffffff8216815273ffffffffffffffffffffffffffffffffffffffff8316907f3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed69060200160405180910390a25050565b5f611b0782612214565b611b104261212a565b611b1a9190612760565b9050611b268282611eab565b6040805165ffffffffffff8085168252831660208201527ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b910160405180910390a15050565b600260035403611ba8576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600355565b6040805173ffffffffffffffffffffffffffffffffffffffff8416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f095ea7b300000000000000000000000000000000000000000000000000000000179052611c3b8482612265565b611a1f5760405173ffffffffffffffffffffffffffffffffffffffff84811660248301525f6044830152611cbf91869182169063095ea7b3906064015b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506122bb565b611a1f84826122bb565b60405173ffffffffffffffffffffffffffffffffffffffff83811660248301526044820183905261123691859182169063a9059cbb90606401611c78565b60015473ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900465ffffffffffff16801580611d5757504265ffffffffffff821610155b15611d98576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff82166004820152602401610acb565b611dc05f611dbb60025473ffffffffffffffffffffffffffffffffffffffff1690565b6120c9565b50611dcb5f83612004565b5050600180547fffffffffffff000000000000000000000000000000000000000000000000000016905550565b5f82815260208190526040902060010154611e12816119e5565b611a1f83836120c9565b6119f95f5f612179565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff166109fb576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8216600482015260248101839052604401610acb565b6002547a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015611f7f574265ffffffffffff82161015611f56576002546001805479ffffffffffffffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000090920465ffffffffffff167a01000000000000000000000000000000000000000000000000000002919091179055611f7f565b6040517f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5905f90a15b506002805473ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000065ffffffffffff9485160279ffffffffffffffffffffffffffffffffffffffffffffffffffff16177a0100000000000000000000000000000000000000000000000000009290931691909102919091179055565b5f826120b8575f61202a60025473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614612077576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84161790555b6120c2838361235a565b9392505050565b5f821580156120f2575060025473ffffffffffffffffffffffffffffffffffffffff8381169116145b1561212057600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001690555b6120c28383612453565b5f65ffffffffffff821115612175576040517f6dfcc6500000000000000000000000000000000000000000000000000000000081526030600482015260248101839052604401610acb565b5090565b600180547401000000000000000000000000000000000000000065ffffffffffff84811682027fffffffffffff0000000000000000000000000000000000000000000000000000841673ffffffffffffffffffffffffffffffffffffffff881617179093559004168015611236576040517f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109905f90a1505050565b5f5f61221e611374565b90508065ffffffffffff168365ffffffffffff161161224657612241838261277e565b6120c2565b6120c265ffffffffffff8416620697805f8282188284100282186120c2565b5f5f5f5f60205f8651602088015f8a5af192503d91505f5190508280156122b15750811561229657806001146122b1565b5f8673ffffffffffffffffffffffffffffffffffffffff163b115b9695505050505050565b5f5f60205f8451602086015f885af1806122da576040513d5f823e3d81fd5b50505f513d915081156122f157806001141561230b565b73ffffffffffffffffffffffffffffffffffffffff84163b155b15611a1f576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610acb565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff1661244c575f8381526020818152604080832073ffffffffffffffffffffffffffffffffffffffff86168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556123ea3390565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4506001610767565b505f610767565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff161561244c575f8381526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8616808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a4506001610767565b5f6020828403121561251c575f5ffd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146120c2575f5ffd5b5f6020828403121561255b575f5ffd5b5035919050565b73ffffffffffffffffffffffffffffffffffffffff8116811461077f575f5ffd5b5f5f60408385031215612594575f5ffd5b823561259f81612562565b915060208301356125af81612562565b809150509250929050565b5f5f604083850312156125cb575f5ffd5b8235915060208301356125af81612562565b5f602082840312156125ed575f5ffd5b81356120c281612562565b5f60208284031215612608575f5ffd5b813565ffffffffffff811681146120c2575f5ffd5b5f5f5f6060848603121561262f575f5ffd5b833561263a81612562565b95602085013595506040909401359392505050565b5f5f60408385031215612660575f5ffd5b823561266b81612562565b946020939093013593505050565b5f606082840312801561268a575f5ffd5b506040516060810167ffffffffffffffff811182821017156126d3577f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b60405282516126e181612562565b815260208301516126f181612562565b60208201526040928301519281019290925250919050565b5f60208284031215612719575f5ffd5b5051919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b8181038181111561076757610767612720565b65ffffffffffff818116838216019081111561076757610767612720565b65ffffffffffff82811682821603908111156107675761076761272056fea2646970667358221220d08781402aaf80cc8eb5a19b063a4445501657f3b2a80d7dc3bd581b03c20aef64736f6c634300081c0033241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b08f8cd32ed93fc2f9fc78152a14807c9609af3d99c5fe4dc6b106a801aaddfe90e",
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
// Solidity: constructor(uint48 _initialDelay, address _owner, address _router, uint256 _minOrderAmount, uint256 _interval, address _recipient, address _buyer, address _manager) returns()
func (buyBot *BuyBot) PackConstructor(_initialDelay *big.Int, _owner common.Address, _router common.Address, _minOrderAmount *big.Int, _interval *big.Int, _recipient common.Address, _buyer common.Address, _manager common.Address) []byte {
	enc, err := buyBot.abi.Pack("", _initialDelay, _owner, _router, _minOrderAmount, _interval, _recipient, _buyer, _manager)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackBUYERROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7a01a1da.
//
// Solidity: function BUYER_ROLE() view returns(bytes32)
func (buyBot *BuyBot) PackBUYERROLE() []byte {
	enc, err := buyBot.abi.Pack("BUYER_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
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
	return out0, err
}

// PackDEFAULTADMINROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (buyBot *BuyBot) PackDEFAULTADMINROLE() []byte {
	enc, err := buyBot.abi.Pack("DEFAULT_ADMIN_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
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
	return out0, err
}

// PackMANAGERROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (buyBot *BuyBot) PackMANAGERROLE() []byte {
	enc, err := buyBot.abi.Pack("MANAGER_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
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
	return out0, err
}

// PackNATIVECOIN is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x04a41159.
//
// Solidity: function NATIVE_COIN() view returns(address)
func (buyBot *BuyBot) PackNATIVECOIN() []byte {
	enc, err := buyBot.abi.Pack("NATIVE_COIN")
	if err != nil {
		panic(err)
	}
	return enc
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
	return out0, err
}

// PackAcceptDefaultAdminTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcefc1429.
//
// Solidity: function acceptDefaultAdminTransfer() returns()
func (buyBot *BuyBot) PackAcceptDefaultAdminTransfer() []byte {
	enc, err := buyBot.abi.Pack("acceptDefaultAdminTransfer")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackBeginDefaultAdminTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x634e93da.
//
// Solidity: function beginDefaultAdminTransfer(address newAdmin) returns()
func (buyBot *BuyBot) PackBeginDefaultAdminTransfer(newAdmin common.Address) []byte {
	enc, err := buyBot.abi.Pack("beginDefaultAdminTransfer", newAdmin)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackBuyMarket is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x67ea88eb.
//
// Solidity: function buyMarket(address pair, uint256 amount, uint256 maxMatchCount) returns()
func (buyBot *BuyBot) PackBuyMarket(pair common.Address, amount *big.Int, maxMatchCount *big.Int) []byte {
	enc, err := buyBot.abi.Pack("buyMarket", pair, amount, maxMatchCount)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCanBuyMarket is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2aaa9628.
//
// Solidity: function canBuyMarket(address pair, address caller) view returns(bool canBuy, uint256 balance)
func (buyBot *BuyBot) PackCanBuyMarket(pair common.Address, caller common.Address) []byte {
	enc, err := buyBot.abi.Pack("canBuyMarket", pair, caller)
	if err != nil {
		panic(err)
	}
	return enc
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
	return *outstruct, err

}

// PackCancelDefaultAdminTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (buyBot *BuyBot) PackCancelDefaultAdminTransfer() []byte {
	enc, err := buyBot.abi.Pack("cancelDefaultAdminTransfer")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackChangeDefaultAdminDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (buyBot *BuyBot) PackChangeDefaultAdminDelay(newDelay *big.Int) []byte {
	enc, err := buyBot.abi.Pack("changeDefaultAdminDelay", newDelay)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackDefaultAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (buyBot *BuyBot) PackDefaultAdmin() []byte {
	enc, err := buyBot.abi.Pack("defaultAdmin")
	if err != nil {
		panic(err)
	}
	return enc
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
	return out0, err
}

// PackDefaultAdminDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (buyBot *BuyBot) PackDefaultAdminDelay() []byte {
	enc, err := buyBot.abi.Pack("defaultAdminDelay")
	if err != nil {
		panic(err)
	}
	return enc
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
	return out0, err
}

// PackDefaultAdminDelayIncreaseWait is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (buyBot *BuyBot) PackDefaultAdminDelayIncreaseWait() []byte {
	enc, err := buyBot.abi.Pack("defaultAdminDelayIncreaseWait")
	if err != nil {
		panic(err)
	}
	return enc
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
	return out0, err
}

// PackGetBalance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf8b2cb4f.
//
// Solidity: function getBalance(address token) view returns(uint256 balance)
func (buyBot *BuyBot) PackGetBalance(token common.Address) []byte {
	enc, err := buyBot.abi.Pack("getBalance", token)
	if err != nil {
		panic(err)
	}
	return enc
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
	return out0, err
}

// PackGetRoleAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (buyBot *BuyBot) PackGetRoleAdmin(role [32]byte) []byte {
	enc, err := buyBot.abi.Pack("getRoleAdmin", role)
	if err != nil {
		panic(err)
	}
	return enc
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
	return out0, err
}

// PackGrantRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (buyBot *BuyBot) PackGrantRole(role [32]byte, account common.Address) []byte {
	enc, err := buyBot.abi.Pack("grantRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackHasRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (buyBot *BuyBot) PackHasRole(role [32]byte, account common.Address) []byte {
	enc, err := buyBot.abi.Pack("hasRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
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
	return out0, err
}

// PackInterval is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x947a36fb.
//
// Solidity: function interval() view returns(uint256)
func (buyBot *BuyBot) PackInterval() []byte {
	enc, err := buyBot.abi.Pack("interval")
	if err != nil {
		panic(err)
	}
	return enc
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
	return out0, err
}

// PackLastBuyTime is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf29f4d0b.
//
// Solidity: function lastBuyTime() view returns(uint256)
func (buyBot *BuyBot) PackLastBuyTime() []byte {
	enc, err := buyBot.abi.Pack("lastBuyTime")
	if err != nil {
		panic(err)
	}
	return enc
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
	return out0, err
}

// PackMinOrderAmount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x46b62c4a.
//
// Solidity: function minOrderAmount() view returns(uint256)
func (buyBot *BuyBot) PackMinOrderAmount() []byte {
	enc, err := buyBot.abi.Pack("minOrderAmount")
	if err != nil {
		panic(err)
	}
	return enc
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
	return out0, err
}

// PackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (buyBot *BuyBot) PackOwner() []byte {
	enc, err := buyBot.abi.Pack("owner")
	if err != nil {
		panic(err)
	}
	return enc
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
	return out0, err
}

// PackPendingDefaultAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (buyBot *BuyBot) PackPendingDefaultAdmin() []byte {
	enc, err := buyBot.abi.Pack("pendingDefaultAdmin")
	if err != nil {
		panic(err)
	}
	return enc
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
	return *outstruct, err

}

// PackPendingDefaultAdminDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (buyBot *BuyBot) PackPendingDefaultAdminDelay() []byte {
	enc, err := buyBot.abi.Pack("pendingDefaultAdminDelay")
	if err != nil {
		panic(err)
	}
	return enc
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
	return *outstruct, err

}

// PackRecipient is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x66d003ac.
//
// Solidity: function recipient() view returns(address)
func (buyBot *BuyBot) PackRecipient() []byte {
	enc, err := buyBot.abi.Pack("recipient")
	if err != nil {
		panic(err)
	}
	return enc
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
	return out0, err
}

// PackRenounceRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (buyBot *BuyBot) PackRenounceRole(role [32]byte, account common.Address) []byte {
	enc, err := buyBot.abi.Pack("renounceRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRevokeRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (buyBot *BuyBot) PackRevokeRole(role [32]byte, account common.Address) []byte {
	enc, err := buyBot.abi.Pack("revokeRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRollbackDefaultAdminDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (buyBot *BuyBot) PackRollbackDefaultAdminDelay() []byte {
	enc, err := buyBot.abi.Pack("rollbackDefaultAdminDelay")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRouter is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (buyBot *BuyBot) PackRouter() []byte {
	enc, err := buyBot.abi.Pack("router")
	if err != nil {
		panic(err)
	}
	return enc
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
	return out0, err
}

// PackSetInterval is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x22a90082.
//
// Solidity: function setInterval(uint256 _interval) returns()
func (buyBot *BuyBot) PackSetInterval(interval *big.Int) []byte {
	enc, err := buyBot.abi.Pack("setInterval", interval)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetMinOrderAmount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa3b8ef04.
//
// Solidity: function setMinOrderAmount(uint256 _minOrderAmount) returns()
func (buyBot *BuyBot) PackSetMinOrderAmount(minOrderAmount *big.Int) []byte {
	enc, err := buyBot.abi.Pack("setMinOrderAmount", minOrderAmount)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetRecipient is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3bbed4a0.
//
// Solidity: function setRecipient(address _recipient) returns()
func (buyBot *BuyBot) PackSetRecipient(recipient common.Address) []byte {
	enc, err := buyBot.abi.Pack("setRecipient", recipient)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (buyBot *BuyBot) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := buyBot.abi.Pack("supportsInterface", interfaceId)
	if err != nil {
		panic(err)
	}
	return enc
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
	return out0, err
}

// PackWithdraw is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf3fef3a3.
//
// Solidity: function withdraw(address token, uint256 amount) returns()
func (buyBot *BuyBot) PackWithdraw(token common.Address, amount *big.Int) []byte {
	enc, err := buyBot.abi.Pack("withdraw", token, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackWithdrawETH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf14210a6.
//
// Solidity: function withdrawETH(uint256 amount) returns()
func (buyBot *BuyBot) PackWithdrawETH(amount *big.Int) []byte {
	enc, err := buyBot.abi.Pack("withdrawETH", amount)
	if err != nil {
		panic(err)
	}
	return enc
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
	if log.Topics[0] != buyBot.abi.Events[event].ID {
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
	if log.Topics[0] != buyBot.abi.Events[event].ID {
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
	if log.Topics[0] != buyBot.abi.Events[event].ID {
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
	if log.Topics[0] != buyBot.abi.Events[event].ID {
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
	if log.Topics[0] != buyBot.abi.Events[event].ID {
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
	if log.Topics[0] != buyBot.abi.Events[event].ID {
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
	if log.Topics[0] != buyBot.abi.Events[event].ID {
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
	if log.Topics[0] != buyBot.abi.Events[event].ID {
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
	if log.Topics[0] != buyBot.abi.Events[event].ID {
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
	if log.Topics[0] != buyBot.abi.Events[event].ID {
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
	if log.Topics[0] != buyBot.abi.Events[event].ID {
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
	if log.Topics[0] != buyBot.abi.Events[event].ID {
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
	if bytes.Equal(raw[:4], buyBot.abi.Errors["BuyBotInsufficientBalance"].ID.Bytes()[:4]) {
		return buyBot.UnpackBuyBotInsufficientBalanceError(raw[4:])
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
	if bytes.Equal(raw[:4], buyBot.abi.Errors["BuyBotInvalidRouter"].ID.Bytes()[:4]) {
		return buyBot.UnpackBuyBotInvalidRouterError(raw[4:])
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
