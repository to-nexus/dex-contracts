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

// IPairConfig is an auto generated low-level Go binding around an user-defined struct.

// IPairOrder is an auto generated low-level Go binding around an user-defined struct.

// PairImplV2MetaData contains all meta data concerning the PairImplV2 contract.
var PairImplV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BASE\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DENOMINATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MARKET\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"QUOTE\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROUTER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"accountReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"base\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quote\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseReserve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"orderIds\",\"type\":\"uint256[]\"}],\"name\":\"cancelOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"orderIds\",\"type\":\"uint256[]\"}],\"name\":\"emergencyCancelOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeConfig\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"sellerMakerFeeBps\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"sellerTakerFeeBps\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"buyerMakerFeeBps\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"buyerTakerFeeBps\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIPair.OrderSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256[2]\",\"name\":\"adjacent\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"findMaxCount\",\"type\":\"uint256\"}],\"name\":\"findPrevPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"QUOTE\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"BASE\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"DENOMINATOR\",\"type\":\"uint256\"}],\"internalType\":\"structIPair.Config\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEffectiveFees\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"sellerMakerFeeBps\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"sellerTakerFeeBps\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"buyerMakerFeeBps\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"buyerTakerFeeBps\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quote\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tickSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lotSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lotSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"matchedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"matchedPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minTradeVolume\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"orderById\",\"outputs\":[{\"components\":[{\"internalType\":\"enumIPair.OrderSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"feeBps\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIPair.Order\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIPair.OrderSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"name\":\"ordersByPrices\",\"outputs\":[{\"internalType\":\"uint256[][]\",\"name\":\"\",\"type\":\"uint256[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quoteReserve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"sellerMakerFeeBps_\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"sellerTakerFeeBps_\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"buyerMakerFeeBps_\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"buyerTakerFeeBps_\",\"type\":\"uint32\"}],\"name\":\"setPairFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"setPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lotSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tickSize\",\"type\":\"uint256\"}],\"name\":\"setTickSize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"erc20\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"skim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumIPair.OrderSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"feeBps\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIPair.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"enumIPair.LimitConstraints\",\"name\":\"constraints\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"prevPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxMatchCount\",\"type\":\"uint256\"}],\"name\":\"submitLimitOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumIPair.OrderSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"feeBps\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIPair.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"spendAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxMatchCount\",\"type\":\"uint256\"}],\"name\":\"submitMarketOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tickSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tickSizes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tick\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lot\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ticks\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"sellPrices\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"buyPrices\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeBps\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"FeeCollect\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"enumIPair.CloseType\",\"name\":\"closeType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"OrderClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"enumIPair.OrderSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"OrderCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"sellId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"buyId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"OrderMatched\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"sellerMakerFee\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"sellerTakerFee\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"buyerMakerFee\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"buyerTakerFee\",\"type\":\"uint32\"}],\"name\":\"PairFeesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"erc20\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Skim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"beforeLotSize\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newLotSize\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"beforeTickSize\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTickSize\",\"type\":\"uint256\"}],\"name\":\"TickSizeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ListFailToFindPrev\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ListInvalidFindMaxCount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ListInvalidIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ListInvalidPrevNode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ListZeroData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"PairFillOrKill\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PairInsufficientTradeVolume\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"PairInvalidAccountReserve\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PairInvalidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PairInvalidFeeBps\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"makerFee\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"takerFee\",\"type\":\"uint32\"}],\"name\":\"PairInvalidFeeStructure\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"PairInvalidInitializeData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PairInvalidOrderId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumIPair.OrderSide\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"PairInvalidOrderSide\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumIPair.OrderSide\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PairInvalidPrevPrice\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PairInvalidPrice\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"PairInvalidReserve\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"PairInvalidRouter\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PairInvalidTickSize\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"PairNotOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumIPair.OrderSide\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PairUnknownPrices\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"ec342ad0": "BASE()",
		"918f8674": "DENOMINATOR()",
		"f46f16c2": "MARKET()",
		"9c579839": "QUOTE()",
		"32fe7b26": "ROUTER()",
		"ad3cb1cc": "UPGRADE_INTERFACE_VERSION()",
		"7aae3523": "accountReserves(address)",
		"dfdf2a72": "baseReserve()",
		"1ec482d7": "cancelOrder(address,uint256[])",
		"e077dd8a": "emergencyCancelOrder(uint256[])",
		"1e5eb1d0": "feeConfig()",
		"6c4cfbc8": "findPrevPrice(uint8,uint256,uint256[2],uint256)",
		"c3f909d4": "getConfig()",
		"7cef243c": "getEffectiveFees()",
		"02a204c6": "initialize(address,address,address,uint256,uint256,bytes)",
		"4942f65f": "lotSize()",
		"4d6754f1": "matchedAt()",
		"1edd8614": "matchedPrice()",
		"7cbf6db2": "minTradeVolume()",
		"8d77eba5": "orderById(uint256)",
		"12a808cc": "ordersByPrices(uint8,uint256[])",
		"8da5cb5b": "owner()",
		"5c975abb": "paused()",
		"52d1902d": "proxiableUUID()",
		"9da771f4": "quoteReserve()",
		"803a0386": "setPairFees(uint32,uint32,uint32,uint32)",
		"bedb86fb": "setPause(bool)",
		"fe566349": "setTickSize(uint256,uint256)",
		"1a0361fb": "skim(address,address,uint256)",
		"2dd6a00a": "submitLimitOrder((uint8,address,uint32,uint256,uint256),uint8,uint256,uint256)",
		"cc50e36a": "submitMarketOrder((uint8,address,uint32,uint256,uint256),uint256,uint256)",
		"f210d087": "tickSize()",
		"d6e9cc2a": "tickSizes()",
		"2cfffaf6": "ticks()",
		"4f1ef286": "upgradeToAndCall(address,bytes)",
	},
	Bin: "0x60a060405230608052348015610013575f5ffd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b608051615e666100f95f395f81816135de01528181613607015261385b0152615e665ff3fe608060405260043610610229575f3560e01c80637cef243c11610131578063c3f909d4116100ac578063e077dd8a1161007c578063f210d08711610062578063f210d08714610765578063f46f16c21461077a578063fe566349146107a5575f5ffd5b8063e077dd8a1461071a578063ec342ad014610739575f5ffd5b8063c3f909d414610651578063cc50e36a146106cf578063d6e9cc2a146106ee578063dfdf2a7214610705575f5ffd5b8063918f8674116101015780639da771f4116100e75780639da771f4146105c8578063ad3cb1cc146105dd578063bedb86fb14610632575f5ffd5b8063918f8674146105875780639c5798391461059c575f5ffd5b80637cef243c14610514578063803a0386146105285780638d77eba5146105475780638da5cb5b14610573575f5ffd5b806332fe7b26116101c157806352d1902d116101915780636c4cfbc8116101775780636c4cfbc8146104825780637aae3523146104a15780637cbf6db2146104ff575f5ffd5b806352d1902d1461042d5780635c975abb14610441575f5ffd5b806332fe7b261461039f5780634942f65f146103f05780634d6754f1146104055780634f1ef2861461041a575f5ffd5b80631ec482d7116101fc5780631ec482d71461031c5780631edd86141461033b5780632cfffaf61461035e5780632dd6a00a14610380575f5ffd5b806302a204c61461022d57806312a808cc1461024e5780631a0361fb146102835780631e5eb1d0146102a2575b5f5ffd5b348015610238575f5ffd5b5061024c6102473660046153ac565b6107c4565b005b348015610259575f5ffd5b5061026d610268366004615441565b610cec565b60405161027a9190615500565b60405180910390f35b34801561028e575f5ffd5b5061024c61029d3660046155aa565b610dd2565b3480156102ad575f5ffd5b506019546102ec9063ffffffff8082169164010000000081048216916801000000000000000082048116916c0100000000000000000000000090041684565b6040805163ffffffff9586168152938516602085015291841691830191909152909116606082015260800161027a565b348015610327575f5ffd5b5061024c610336366004615629565b61113a565b348015610346575f5ffd5b5061035060075481565b60405190815260200161027a565b348015610369575f5ffd5b5061037261134a565b60405161027a9291906156b4565b34801561038b575f5ffd5b5061035061039a36600461576e565b611384565b3480156103aa575f5ffd5b506001546103cb9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161027a565b3480156103fb575f5ffd5b50610350600a5481565b348015610410575f5ffd5b5061035060085481565b61024c6104283660046157b4565b611a1b565b348015610438575f5ffd5b50610350611a3a565b34801561044c575f5ffd5b507fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff16604051901515815260200161027a565b34801561048d575f5ffd5b5061035061049c366004615801565b611a68565b3480156104ac575f5ffd5b506104ea6104bb366004615849565b73ffffffffffffffffffffffffffffffffffffffff165f90815260186020526040902080546001909101549091565b6040805192835260208301919091520161027a565b34801561050a575f5ffd5b50610350600b5481565b34801561051f575f5ffd5b506102ec611b4a565b348015610533575f5ffd5b5061024c610542366004615864565b611b78565b348015610552575f5ffd5b506105666105613660046158bd565b611bca565b60405161027a919061593a565b34801561057e575f5ffd5b506103cb611c99565b348015610592575f5ffd5b5061035060045481565b3480156105a7575f5ffd5b506003546103cb9073ffffffffffffffffffffffffffffffffffffffff1681565b3480156105d3575f5ffd5b5061035060065481565b3480156105e8575f5ffd5b506106256040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b60405161027a9190615996565b34801561063d575f5ffd5b5061024c61064c3660046159e9565b611d2c565b34801561065c575f5ffd5b5060408051606080820183525f8083526020808401829052928401528251808201845260035473ffffffffffffffffffffffffffffffffffffffff90811680835260025482168386019081526004549387019384528651918252519091169381019390935251928201929092520161027a565b3480156106da575f5ffd5b5061024c6106e9366004615a08565b611d85565b3480156106f9575f5ffd5b50600954600a546104ea565b348015610710575f5ffd5b5061035060055481565b348015610725575f5ffd5b5061024c610734366004615a39565b611f5b565b348015610744575f5ffd5b506002546103cb9073ffffffffffffffffffffffffffffffffffffffff1681565b348015610770575f5ffd5b5061035060095481565b348015610785575f5ffd5b505f546103cb9073ffffffffffffffffffffffffffffffffffffffff1681565b3480156107b0575f5ffd5b5061024c6107bf366004615a78565b6120a2565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff165f8115801561080e5750825b90505f8267ffffffffffffffff16600114801561082a5750303b155b905081158015610838575080155b1561086f576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000016600117855583156108d05784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b6108d86122ba565b73ffffffffffffffffffffffffffffffffffffffff8b1661094c576040517fc587916e0000000000000000000000000000000000000000000000000000000081527f726f75746572000000000000000000000000000000000000000000000000000060048201526024015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8a166109bb576040517fc587916e0000000000000000000000000000000000000000000000000000000081527f71756f74650000000000000000000000000000000000000000000000000000006004820152602401610943565b73ffffffffffffffffffffffffffffffffffffffff8916610a2a576040517fc587916e0000000000000000000000000000000000000000000000000000000081527f62617365000000000000000000000000000000000000000000000000000000006004820152602401610943565b875f03610a85576040517fc587916e0000000000000000000000000000000000000000000000000000000081527f7469636b53697a650000000000000000000000000000000000000000000000006004820152602401610943565b865f03610ae0576040517fc587916e0000000000000000000000000000000000000000000000000000000081527f6c6f7453697a65000000000000000000000000000000000000000000000000006004820152602401610943565b335f80547fffffffffffffffffffffffff000000000000000000000000000000000000000090811673ffffffffffffffffffffffffffffffffffffffff938416179091556001805482168e84161790556003805482168d841617905560028054909116918b169182179055604080517f313ce567000000000000000000000000000000000000000000000000000000008152905163313ce567916004808201926020929091908290030181865afa158015610b9d573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610bc19190615a98565b610bcc90600a615bc0565b6004819055610bdb888a615bce565b610be59190615c12565b15610c2e57600480546040517f575681350000000000000000000000000000000000000000000000000000000081529182018a9052602482018990526044820152606401610943565b6009889055600a879055600454610c4890899089906122cc565b600b819055505f5f5f5f89806020019051810190610c669190615c4a565b9350935093509350610c7a84848484612383565b505050508315610cdf5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050505050565b60605f80846001811115610d0257610d026158d4565b14610d0e576016610d11565b60155b83519091505f8167ffffffffffffffff811115610d3057610d306152a6565b604051908082528060200260200182016040528015610d6357816020015b6060815260200190600190039081610d4e5790505b5090505f5b82811015610dc657610da1845f888481518110610d8757610d87615c9b565b602002602001015181526020019081526020015f206126bc565b828281518110610db357610db3615c9b565b6020908102919091010152600101610d68565b50925050505b92915050565b610dda611c99565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610e5957335b6040517f118cdaa700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610943565b80156111355760025473ffffffffffffffffffffffffffffffffffffffff8481169116148015610f24575080600554610e929190615cc8565b6002546040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff909116906370a0823190602401602060405180830381865afa158015610efe573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f229190615cdb565b105b15610f77576002546040517f7bdf264d00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610943565b60035473ffffffffffffffffffffffffffffffffffffffff848116911614801561103c575080600654610faa9190615cc8565b6003546040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff909116906370a0823190602401602060405180830381865afa158015611016573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061103a9190615cdb565b105b1561108f576003546040517f7bdf264d00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610943565b6110b073ffffffffffffffffffffffffffffffffffffffff84168383612762565b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff166110e53390565b73ffffffffffffffffffffffffffffffffffffffff167f7e30e394efdaabad9f9ca6fd8f67f9449ba63d674ee8068283c3ae1f49b5b6608460405161112c91815260200190565b60405180910390a45b505050565b60015473ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146111bc57335b6040517f143db19100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610943565b805f5b81811015611343575f8484838181106111da576111da615c9b565b602090810292909201355f81815260179093526040808420815160a08101909252805492955090925090829060ff16600181111561121a5761121a6158d4565b600181111561122b5761122b6158d4565b81528154610100810473ffffffffffffffffffffffffffffffffffffffff908116602080850191909152750100000000000000000000000000000000000000000090920463ffffffff16604084015260018401546060840152600290930154608090920191909152820151919250166112a55750506111bf565b8673ffffffffffffffffffffffffffffffffffffffff16816020015173ffffffffffffffffffffffffffffffffffffffff161461132d576040517f54de53f70000000000000000000000000000000000000000000000000000000081526004810183905273ffffffffffffffffffffffffffffffffffffffff88166024820152604401610943565b611339828260026127ef565b50506001016111bf565b5050505050565b606080611370600d5f5b60ff166002811061136757611367615c9b565b600402016126bc565b915061137e600d6001611354565b90509091565b5f61138d612978565b60015473ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146113c85733611171565b6113d06129d4565b606085015115806113f1575060095485606001516113ee9190615c12565b15155b156114305784606001516040517f4e1404fe00000000000000000000000000000000000000000000000000000000815260040161094391815260200190565b608085015115806114515750600a54856080015161144e9190615c12565b15155b156114905784608001516040517fa334626e00000000000000000000000000000000000000000000000000000000815260040161094391815260200190565b600c5f815461149e90615cf2565b918290555090505f80865160018111156114ba576114ba6158d4565b1490505f5f826114d5576114d084895f88612ac0565b6114e0565b6114e0848987612e35565b91509150811561156e575f847f78615f01112d8477ae8d6fc9a510bcfdf935bc3ab7b02d0ff2316f5b44bd597a4260405161151d91815260200190565b60405180910390a38280156115355750608088015115155b1561156957602088015160808901516002546115699273ffffffffffffffffffffffffffffffffffffffff90911691612762565b6119fd565b6001876002811115611582576115826158d4565b036115f8576001847f78615f01112d8477ae8d6fc9a510bcfdf935bc3ab7b02d0ff2316f5b44bd597a426040516115bb91815260200190565b60405180910390a3821561156957602088015160808901516002546115699273ffffffffffffffffffffffffffffffffffffffff90911691612762565b600287600281111561160c5761160c6158d4565b036116615760208801516040517f1b3c356200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610943565b8515801590611692575082801561167b5750858860600151105b806116925750821580156116925750858860600151115b156116d457875160608901516040517fc18aa6060000000000000000000000000000000000000000000000000000000081526109439291908990600401615d0a565b611712886060015187600d8b5f015160018111156116f4576116f46158d4565b60ff166002811061170757611707615c9b565b6004020191906130db565b505f848152601760205260409020885181548a92919082907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001838181111561175f5761175f6158d4565b021790555060208201518154604084015163ffffffff167501000000000000000000000000000000000000000000027fffffffffffffff00000000ffffffffffffffffffffffffffffffffffffffffff73ffffffffffffffffffffffffffffffffffffffff90931661010002929092167fffffffffffffff000000000000000000000000000000000000000000000000ff909116171781556060820151600182015560809091015160029091015582156118de5761182588602001518960800151613289565b604080516080810182525f80825260208201819052918101829052606001527f1d2ff3fa6980aeeebca4e94965520da48983006e9b1115c1c853cbb10d943d005c515f85815260176020908152604080832080547fffffffffffffff00000000ffffffffffffffffffffffffffffffffffffffffff16750100000000000000000000000000000000000000000063ffffffff9687160217905560608c01518352601590915290206118d89186906133b316565b506119fd565b5f6118f489606001518a608001516004546122cc565b90505f611942604080516080810182525f808252602082018190528183018190526060909101527f1d2ff3fa6980aeeebca4e94965520da48983006e9b1115c1c853cbb10d943d005c015190565b905063ffffffff811615611974575f6119648363ffffffff84166127106122cc565b90506119708184615cc8565b9250505b5f868152601760209081526040909120805463ffffffff9093167501000000000000000000000000000000000000000000027fffffffffffffff00000000ffffffffffffffffffffffffffffffffffffffffff909316929092179091558901516119de90826133c3565b60608901515f9081526016602052604090206119fa90866133b3565b50505b82611a1057611a108860200151826134e9565b505050949350505050565b611a236135c6565b611a2c826136ca565b611a36828261370a565b5050565b5f611a43613843565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b5f80856001811115611a7c57611a7c6158d4565b03611ae557604080518082018252611ade91869190869060029083908390808284375f92019190915250869150600d9050896001811115611abf57611abf6158d4565b60ff1660028110611ad257611ad2615c9b565b600402019291906138b2565b9050611b42565b604080518082018252611ade91869190869060029083908390808284375f92019190915250869150600d9050896001811115611b2357611b236158d4565b60ff1660028110611b3657611b36615c9b565b60040201929190613a58565b949350505050565b5f5f5f5f5f611b57613bc7565b80516020820151604083015160609093015191989097509195509350915050565b611b80611c99565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611bb85733610e0e565b611bc484848484612383565b50505050565b6040805160a0810182525f808252602082018190529181018290526060810182905260808101919091525f8281526017602052604090819020815160a081019092528054829060ff166001811115611c2457611c246158d4565b6001811115611c3557611c356158d4565b81528154610100810473ffffffffffffffffffffffffffffffffffffffff1660208301527501000000000000000000000000000000000000000000900463ffffffff1660408201526001820154606082015260029091015460809091015292915050565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611d03573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611d279190615d29565b905090565b611d34611c99565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611d6c5733610e0e565b8015611d7d57611d7a613e5e565b50565b611d7a613efe565b611d8d612978565b60015473ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611dc85733611171565b611dd06129d4565b5f600c5f8154611ddf90615cf2565b918290555090505f84516001811115611dfa57611dfa6158d4565b03611ea957821580611e175750600a54611e149084615c12565b15155b15611e51576040517fa334626e00000000000000000000000000000000000000000000000000000000815260048101849052602401610943565b5f606085015260808401839052611e69818584612e35565b5050608084015115611ea45760208401516080850151600254611ea49273ffffffffffffffffffffffffffffffffffffffff90911691612762565b611f1a565b600b54831015611ef357600b546040517f0bc1e7dd000000000000000000000000000000000000000000000000000000008152610943918591600401918252602082015260400190565b5f1960608501525f611f0782868686612ac0565b915050611f188560200151826134e9565b505b5f817f78615f01112d8477ae8d6fc9a510bcfdf935bc3ab7b02d0ff2316f5b44bd597a42604051611f4d91815260200190565b60405180910390a350505050565b611f63613f74565b611f6b611c99565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611fa35733610e0e565b805f5b81811015611bc4575f848483818110611fc157611fc1615c9b565b602090810292909201355f81815260179093526040808420815160a08101909252805492955090925090829060ff166001811115612001576120016158d4565b6001811115612012576120126158d4565b81528154610100810473ffffffffffffffffffffffffffffffffffffffff908116602080850191909152750100000000000000000000000000000000000000000090920463ffffffff166040840152600184015460608401526002909301546080909201919091528201519192501661208c575050611fa6565b612098828260036127ef565b5050600101611fa6565b5f5473ffffffffffffffffffffffffffffffffffffffff1663a1eea778336040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815273ffffffffffffffffffffffffffffffffffffffff90911660048201526024015f6040518083038186803b158015612122575f5ffd5b505afa158015612134573d5f5f3e3d5ffd5b50505050805f03612193576040517fc587916e0000000000000000000000000000000000000000000000000000000081527f7469636b53697a650000000000000000000000000000000000000000000000006004820152602401610943565b815f036121ee576040517fc587916e0000000000000000000000000000000000000000000000000000000081527f6c6f7453697a65000000000000000000000000000000000000000000000000006004820152602401610943565b6004546121fb8383615bce565b6122059190615c12565b1561224e57600480546040517f57568135000000000000000000000000000000000000000000000000000000008152918201839052602482018490526044820152606401610943565b600a546009546040805192835260208301859052820152606081018290527f66748bc112f6692a76bebecc33e69768508e27e414e3080f5c02295b6ea3cf3f9060800160405180910390a1600a82905560098190556004546122b390829084906122cc565b600b555050565b6122c2613fcf565b6122ca614036565b565b5f838302815f1985870982811083820303915050805f03612300578382816122f6576122f6615be5565b049250505061237c565b808411612317576123176003851502601118614087565b5f848688095f868103871696879004966002600389028118808a02820302808a02820302808a02820302808a02820302808a02820302808a02909103029181900381900460010186841190950394909402919094039290920491909117919091029150505b9392505050565b61271063ffffffff8516108015906123a1575063ffffffff84811614155b156123d8576040517fac408a2600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61271063ffffffff8416108015906123f6575063ffffffff83811614155b1561242d576040517fac408a2600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61271063ffffffff83161080159061244b575063ffffffff82811614155b15612482576040517fac408a2600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61271063ffffffff8216108015906124a0575063ffffffff81811614155b156124d7576040517fac408a2600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8363ffffffff168363ffffffff161080156124f8575063ffffffff84811614155b1561253f576040517f3018805400000000000000000000000000000000000000000000000000000000815263ffffffff808616600483015284166024820152604401610943565b8163ffffffff168163ffffffff16108015612560575063ffffffff82811614155b156125a7576040517f3018805400000000000000000000000000000000000000000000000000000000815263ffffffff808416600483015282166024820152604401610943565b604080516080808201835263ffffffff87811680845287821660208086018290528884168688018190529388166060968701819052601980547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001685176401000000008502177fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff166801000000000000000087027fffffffffffffffffffffffffffffffff00000000ffffffffffffffffffffffff16176c01000000000000000000000000830217905587519384529083019190915294810191909152918201929092527f86f539dd22f42128e02dc316d053624f3c64c6a38938170bc4bd73e700f1d436910160405180910390a150505050565b60605f825f015467ffffffffffffffff8111156126db576126db6152a6565b604051908082528060200260200182016040528015612704578160200160208202803683370190505b5083546001850154919250905f5b82811015612758578184828151811061272d5761272d615c9b565b6020908102919091018101919091525f92835260038701905260409091206001908101549101612712565b5091949350505050565b6040805173ffffffffffffffffffffffffffffffffffffffff8416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb00000000000000000000000000000000000000000000000000000000179052611135908490614098565b5f808084516001811115612805576128056158d4565b14905080156128735760608401515f9081526015602052604090206080850151909250801561286d5760208501516002546128599173ffffffffffffffffffffffffffffffffffffffff9091169083612762565b61286b8560200151825f600554614137565b505b50612923565b60165f856060015181526020019081526020015f2091505f6128a0856060015186608001516004546122cc565b9050801561292157604085015163ffffffff16156128e2576128d581866040015163ffffffff1661271063ffffffff166122cc565b6128df9082615cc8565b90505b602085015160035461290d9173ffffffffffffffffffffffffffffffffffffffff9091169083612762565b61291f8560200151825f600654614269565b505b505b61292e85848461439a565b8154611343576129708460600151600d865f01516001811115612953576129536158d4565b60ff166002811061296657612966615c9b565b600402019061446c565b505050505050565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff16156122ca576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8054604080517fc415b95c000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff9092169291839163c415b95c9160048083019260209291908290030181865afa158015612a44573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612a689190615d29565b90505f612a73613bc7565b9050817fd6aa07baf8485abf9d26fecf4c935d75b50a73e678db02b944bd3ac8759823005d807f1d2ff3fa6980aeeebca4e94965520da48983006e9b1115c1c853cbb10d943d005d505050565b5f80600185516001811115612ad757612ad76158d4565b14612b10575f6040517f7cbdd0810000000000000000000000000000000000000000000000000000000081526004016109439190615d44565b5f8415612b1e575083612b36565b612b33866060015187608001516004546122cc565b90505b604080516080810182525f8082526020820181905291810182905260609081018290527f1d2ff3fa6980aeeebca4e94965520da48983006e9b1115c1c853cbb10d943d005c015181905f8163ffffffff165f03612b93575f612ba6565b612ba68563ffffffff84166127106122cc565b6003546040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152919250612c5b9173ffffffffffffffffffffffffffffffffffffffff909116906370a0823190602401602060405180830381865afa158015612c19573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612c3d9190615cdb565b8287600654612c4c9190615cc8565b612c569190615cc8565b61453d565b9450925082612cb2576003546040517f7bdf264d00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610943565b89516001811115612cc557612cc56158d4565b8b8b6020015173ffffffffffffffffffffffffffffffffffffffff167f7cde7559ef00f8c3d9244d3fde19e7500979d64f6f9c086a3836a4926abbe1658d606001518e6080015142604051612d2d939291909283526020830191909152604082015260600190565b60405180910390a45050505f5f5f612d478b8b8b8b614561565b925092509250815f14612e24575f612dd58c8c602001518585612d887fd6aa07baf8485abf9d26fecf4c935d75b50a73e678db02b944bd3ac8759823005c90565b604080516080810182525f808252602082018190529181018290526060908101919091527f1d2ff3fa6980aeeebca4e94965520da48983006e9b1115c1c853cbb10d943d005c015161488b565b90508015612e2257612e227fd6aa07baf8485abf9d26fecf4c935d75b50a73e678db02b944bd3ac8759823005c60035473ffffffffffffffffffffffffffffffffffffffff169083612762565b505b509099919850909650505050505050565b5f808084516001811115612e4b57612e4b6158d4565b14612e855760016040517f7cbdd0810000000000000000000000000000000000000000000000000000000081526004016109439190615d44565b8360800151600554612e979190615cc8565b6002546040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff909116906370a0823190602401602060405180830381865afa158015612f03573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612f279190615cdb565b1015612f7b576002546040517f7bdf264d00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610943565b83516001811115612f8e57612f8e6158d4565b85856020015173ffffffffffffffffffffffffffffffffffffffff167f7cde7559ef00f8c3d9244d3fde19e7500979d64f6f9c086a3836a4926abbe1658760600151886080015142604051612ff6939291909283526020830191909152604082015260600190565b60405180910390a45f5f61300b87878761495a565b91509150805f146130cc575f807fd6aa07baf8485abf9d26fecf4c935d75b50a73e678db02b944bd3ac8759823005c613086604080516080810182525f80825260208083018290529282018190526060909101527f1d2ff3fa6980aeeebca4e94965520da48983006e9b1115c1c853cbb10d943d005c015190565b915091505f61309c8a8a60200151868686614bff565b905080156130c8576003546130c89073ffffffffffffffffffffffffffffffffffffffff168483612762565b5050505b5091505f90505b935093915050565b5f825f03613115576040517f0da93b9700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61311f8484614cf9565b1561312b57505f61237c565b5f828152600385016020908152604080832081518083019092528382529181019290925290835f036131975760018601805490869055801561317a575f81815260038801602052604090208690555b60405180604001604052805f815260200182815250915050613239565b6040805180820190915282548152600183015460208201526131b890614d41565b806131c65750856001015484145b6131fc576040517f182dfca500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5060018101546040805180820190915284815260208101829052908015613230575f81815260038801602052604090208690555b50600182018590555b80602001515f0361324c57600286018590555b5f858152600387016020908152604082208351815590830151600190910155865487919061327990615cf2565b9091555060019695505050505050565b5f5f61329760055484614d57565b91509150816132ee576002546040517f7bdf264d00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610943565b600581905573ffffffffffffffffffffffffffffffffffffffff84165f908152601860205260408120613324915b015484614d57565b909250905081613384576002546040517f340b575200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff80871660048301529091166024820152604401610943565b73ffffffffffffffffffffffffffffffffffffffff84165f90815260186020526040812082915b015550505050565b5f61237c838385600201546130db565b5f5f6133d160065484614d57565b9150915081613428576003546040517f7bdf264d00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610943565b600681905573ffffffffffffffffffffffffffffffffffffffff84165f90815260186020526040902061345c90600161331c565b9092509050816134bc576003546040517f340b575200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff80871660048301529091166024820152604401610943565b73ffffffffffffffffffffffffffffffffffffffff84165f908152601860205260409020819060016133ab565b6003546040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201525f9173ffffffffffffffffffffffffffffffffffffffff16906370a0823190602401602060405180830381865afa158015613555573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906135799190615cdb565b90505f6006548361358a9190615cc8565b905080821115611bc4575f61359f8284615d52565b6003549091506113439073ffffffffffffffffffffffffffffffffffffffff168683612762565b3073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016148061369357507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1661367a7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614155b156122ca576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6136d2611c99565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611d7a5733610e0e565b8173ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa92505050801561378f575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016820190925261378c91810190615cdb565b60015b6137dd576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610943565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8114613839576040517faa1d49a400000000000000000000000000000000000000000000000000000000815260048101829052602401610943565b6111358383614d7c565b3073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016146122ca576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f835f036138ec576040517f0da93b9700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6138f68585614cf9565b1561391157505f838152600385016020526040902054611b42565b846001015484108061392257508454155b1561392e57505f611b42565b846002015484111561394557506002840154611b42565b815f0361397e576040517fab7e639300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6139898685614dde565b9050848110156139e9575b80158015906139a257508215155b156139e4575f908152600386016020526040902060010154848111156139d9575f9081526003860160205260409020549050611b42565b5f1990920191613994565b613a26565b80158015906139f757508215155b15613a26575f90815260038601602052604090205484811015613a1b579050611b42565b5f19909201916139e9565b6040517f80ce60d100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f835f03613a92576040517f0da93b9700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b613a9c8585614cf9565b15613ab757505f838152600385016020526040902054611b42565b8460010154841180613ac857508454155b15613ad457505f611b42565b8460020154841015613aeb57506002840154611b42565b815f03613b24576040517fab7e639300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f613b2f8685614dde565b905084811115613b8a575b8015801590613b4857508215155b156139e4575f90815260038601602052604090206001015480851115613b7f575f9081526003860160205260409020549050611b42565b5f1990920191613b3a565b8015801590613b9857508215155b15613a26575f90815260038601602052604090205480851015613bbc579050611b42565b5f1990920191613b8a565b60408051608080820183525f8083526020830181905282840181905260608301819052805484517f5fbbc0d200000000000000000000000000000000000000000000000000000000815294519394919373ffffffffffffffffffffffffffffffffffffffff90911692635fbbc0d292600480820193918290030181865afa158015613c54573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613c789190615d65565b60195490915063ffffffff90811614613c995760195463ffffffff16613c9c565b80515b63ffffffff90811683526019546401000000009004811614613cce57601954640100000000900463ffffffff16613cd4565b80602001515b63ffffffff9081166020840152601954680100000000000000009004811614613d115760195468010000000000000000900463ffffffff16613d17565b80604001515b63ffffffff90811660408401526019546c010000000000000000000000009004811614613d5c576019546c01000000000000000000000000900463ffffffff16613d62565b80606001515b63ffffffff90811660608401526020830151835190821691161115613dd9576019547fffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000164010000000090910463ffffffff1601613dcb57815163ffffffff166020830152613dd9565b602082015163ffffffff1682525b816060015163ffffffff16826040015163ffffffff161115613e5a576019547fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000016c0100000000000000000000000090910463ffffffff1601613e4957604082015163ffffffff1660608301525090565b606082015163ffffffff1660408301525b5090565b613e66612978565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a150565b613f06613f74565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa33613ed3565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff166122ca576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff166122ca576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61403e613fcf565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b634e487b715f52806020526024601cfd5b5f5f60205f8451602086015f885af1806140b7576040513d5f823e3d81fd5b50505f513d915081156140ce5780600114156140e8565b73ffffffffffffffffffffffffffffffffffffffff84163b155b15611bc4576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610943565b73ffffffffffffffffffffffffffffffffffffffff84165f9081526018602052604081208190819061416d90825b01548761453d565b91509150816141cc576002546040517f340b575200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff808a1660048301529091166024820152604401610943565b73ffffffffffffffffffffffffffffffffffffffff87165f9081526018602052604090208190556141fd848761453d565b9350915081614254576002546040517f7bdf264d00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610943565b8461425f5760058390555b5050949350505050565b73ffffffffffffffffffffffffffffffffffffffff84165f9081526018602052604081208190819061429c906001614165565b91509150816142fb576003546040517f340b575200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff808a1660048301529091166024820152604401610943565b73ffffffffffffffffffffffffffffffffffffffff87165f90815260186020526040902081906001015561432f848761453d565b9350915081614386576003546040517f7bdf264d00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610943565b8461425f5760068390555050949350505050565b6143a4818461446c565b6143dd576040517ffcfdf90200000000000000000000000000000000000000000000000000000000815260048101849052602401610943565b5f83815260176020526040812080547fffffffffffffff00000000000000000000000000000000000000000000000000168155600181018290556002015581600381111561442d5761442d6158d4565b837f78615f01112d8477ae8d6fc9a510bcfdf935bc3ab7b02d0ff2316f5b44bd597a4260405161445f91815260200190565b60405180910390a3505050565b5f815f036144a6576040517f0da93b9700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6144b08383614cf9565b6144bb57505f610dcc565b5f82815260038401602052604080822080548352818320600180830154808652939094208482019390935581548355928601549092919085900361450457600180840154908701555b8486600201540361451757825460028701555b5f8581526003870160205260408120818155600101819055865487919061327990615dea565b5f5f8383111561455157505f90508061455a565b50600190508183035b9250929050565b6040805160c0810182525f8082526020820181905291810182905260608101829052600554608082015260a0810182905281908190600d5b8054156147f8576145aa815f614e28565b8083526060890151106147f85781515f90815260156020526040902082516145d190614e96565b8715614630575f6145f584604001518a6145eb9190615d52565b60045486516122cc565b9050600a54816146059190615c12565b61460f9082615d52565b60808b0181905290505f81900361462e575050600160a08301526147f8565b505b80541561479b575f6146428282614e28565b90505f60175f8381526020019081526020015f2090505f5f5f61466c8f8f87878c5f01518b614ebc565b9250925092505f614683895f0151846004546122cc565b90506146b38685837fd6aa07baf8485abf9d26fecf4c935d75b50a73e678db02b944bd3ac8759823005c86614bff565b896060018181516146c49190615cc8565b9052506020890180518491906146db908390615cc8565b9052506040890180518291906146f2908390615cc8565b9150818152505061470a848460018c60800151614137565b6080808b01919091528f0151158061472b57506147268d615dea565b9c508c155b1561479057865461477d57885161474390899061446c565b61477d5788516040517f9f1cfdfe000000000000000000000000000000000000000000000000000000008152610943915f91600401615dff565b5050600160a08801525061479b92505050565b505050505050614630565b8260a00151156147ab57506147f8565b82516147b890839061446c565b6147f25782516040517f9f1cfdfe000000000000000000000000000000000000000000000000000000008152610943915f91600401615dff565b50614599565b60608201511561484c5761484c7fd6aa07baf8485abf9d26fecf4c935d75b50a73e678db02b944bd3ac8759823005c606084015160035473ffffffffffffffffffffffffffffffffffffffff169190612762565b6005548260800151146148625760808201516005555b8160a00151826020015183604001519450945094505050614881614fe2565b9450945094915050565b6002545f906148b19073ffffffffffffffffffffffffffffffffffffffff168787612762565b8163ffffffff165f036148c557505f614950565b5f6148d98563ffffffff85166127106122cc565b905073ffffffffffffffffffffffffffffffffffffffff808516908816897f9814b2b2f43f9ea51ad7e8760fda094ffa3012bfeb10c76a9656389af18956038887866149258184615d52565b6040805194855263ffffffff909316602085015291830152606082015260800160405180910390a490505b9695505050505050565b6040805160a0810182525f80825260208201819052918101829052600654606082015260808101829052819060115b805415614b7d5761499a815f614e28565b808352606087015111614b7d5781515f90815260166020526040902082516149c190614e96565b805415614b1f575f6149d38282614e28565b5f818152601760205260408120865192935091819081906149fc908e908e90889088908b614ebc565b9250925092505f614a13895f0151846004546122cc565b90505f614a45878686857fd6aa07baf8485abf9d26fecf4c935d75b50a73e678db02b944bd3ac8759823005c8861488b565b9050808a604001818151614a599190615cc8565b90525060208a018051839190614a70908390615cc8565b905250614a8d85614a818385615cc8565b60018d60600151614269565b60608b015260808e01511580614aac5750614aa78d615dea565b9c508c155b15614b13578754614aff578951614ac4908a9061446c565b614aff5789516040517f9f1cfdfe00000000000000000000000000000000000000000000000000000000815261094391600191600401615dff565b50506001608089015250614b1f9350505050565b505050505050506149c1565b826080015115614b2f5750614b7d565b8251614b3c90839061446c565b614b775782516040517f9f1cfdfe00000000000000000000000000000000000000000000000000000000815261094391600191600401615dff565b50614989565b604082015115614bd157614bd17fd6aa07baf8485abf9d26fecf4c935d75b50a73e678db02b944bd3ac8759823005c604084015160035473ffffffffffffffffffffffffffffffffffffffff169190612762565b600654826060015114614be75760608201516006555b816080015182602001519350935050506130d3614fe2565b5f8163ffffffff165f03614c3857600354614c319073ffffffffffffffffffffffffffffffffffffffff168686612762565b505f614cf0565b5f614c4c8563ffffffff85166127106122cc565b90505f614c598287615d52565b6040805188815263ffffffff871660208201529081018490526060810182905290915073ffffffffffffffffffffffffffffffffffffffff80871691908916908a907f9814b2b2f43f9ea51ad7e8760fda094ffa3012bfeb10c76a9656389af18956039060800160405180910390a4600354614cec9073ffffffffffffffffffffffffffffffffffffffff168883612762565b5090505b95945050505050565b5f815f03614d0857505f610dcc565b826001015482148061237c57505f828152600384016020908152604091829020825180840190935280548352600101549082015261237c905b80515f90151580610dcc57505060200151151590565b5f8083830184811015614d70575f5f925092505061455a565b60019590945092505050565b614d8582615029565b60405173ffffffffffffffffffffffffffffffffffffffff8316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115614dd65761113582826150f7565b611a3661516d565b5f614df08383835b6020020151614cf9565b15614e0457815f5b60200201519050610dcc565b614e1083836001614de6565b15614e1d57816001614df8565b506001820154610dcc565b81545f908210614e64576040517f39e60f7300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001830154825b8015614e8e575f91825260038501602052604090912060010154905f1901614e6b565b509392505050565b807ffd0e5d4f9b88892d3b04349a0e2bc0d1359414c21932fcd7d5a523a6c0a5cd005d50565b5f5f5f855f0160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16614ef5896080015188600201546151a5565b875491945092507501000000000000000000000000000000000000000000900463ffffffff1690505f80808a516001811115614f3357614f336158d4565b14614f3f57888b614f42565b8a895b915091508681837f6a6896b1e6131e0b8ebae43fdc84cb0178a6eacdf4ee63f15dabae48729a8a038742604051614f83929190918252602082015260400190565b60405180910390a487600201548403614fa657614fa1895f8861439a565b614fb2565b60028801805485900390555b89608001518403614fc8575f60808b0152614fd4565b60808a01805185900390525b505096509650969350505050565b7ffd0e5d4f9b88892d3b04349a0e2bc0d1359414c21932fcd7d5a523a6c0a5cd005c8015611d7a5760075481146150195760078190555b4260085414611d7a574260085550565b8073ffffffffffffffffffffffffffffffffffffffff163b5f03615091576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152602401610943565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60605f5f8473ffffffffffffffffffffffffffffffffffffffff16846040516151209190615e1a565b5f60405180830381855af49150503d805f8114615158576040519150601f19603f3d011682016040523d82523d5f602084013e61515d565b606091505b5091509150614cf08583836151b4565b34156122ca576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f82821882841002821861237c565b6060826151c9576151c482615243565b61237c565b81511580156151ed575073ffffffffffffffffffffffffffffffffffffffff84163b155b1561523c576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610943565b508061237c565b8051156152535780518082602001fd5b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff81168114611d7a575f5ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561531a5761531a6152a6565b604052919050565b5f82601f830112615331575f5ffd5b813567ffffffffffffffff81111561534b5761534b6152a6565b61537c60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116016152d3565b818152846020838601011115615390575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f5f5f5f5f5f60c087890312156153c1575f5ffd5b86356153cc81615285565b955060208701356153dc81615285565b945060408701356153ec81615285565b9350606087013592506080870135915060a087013567ffffffffffffffff811115615415575f5ffd5b61542189828a01615322565b9150509295509295509295565b80356002811061543c575f5ffd5b919050565b5f5f60408385031215615452575f5ffd5b61545b8361542e565b9150602083013567ffffffffffffffff811115615476575f5ffd5b8301601f81018513615486575f5ffd5b803567ffffffffffffffff8111156154a0576154a06152a6565b8060051b6154b0602082016152d3565b918252602081840181019290810190888411156154cb575f5ffd5b6020850194505b838510156154f1578435808352602095860195909350909101906154d2565b80955050505050509250929050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b8281101561559e578685037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0018452815180518087526020918201918701905f5b81811015615585578351835260209384019390920191600101615567565b5090965050506020938401939190910190600101615526565b50929695505050505050565b5f5f5f606084860312156155bc575f5ffd5b83356155c781615285565b925060208401356155d781615285565b929592945050506040919091013590565b5f5f83601f8401126155f8575f5ffd5b50813567ffffffffffffffff81111561560f575f5ffd5b6020830191508360208260051b850101111561455a575f5ffd5b5f5f5f6040848603121561563b575f5ffd5b833561564681615285565b9250602084013567ffffffffffffffff811115615661575f5ffd5b61566d868287016155e8565b9497909650939450505050565b5f8151808452602084019350602083015f5b828110156156aa57815186526020958601959091019060010161568c565b5093949350505050565b604081525f6156c6604083018561567a565b8281036020840152614cf0818561567a565b63ffffffff81168114611d7a575f5ffd5b5f60a082840312156156f9575f5ffd5b60405160a0810167ffffffffffffffff8111828210171561571c5761571c6152a6565b60405290508061572b8361542e565b8152602083013561573b81615285565b6020820152604083013561574e816156d8565b604082015260608381013590820152608092830135920191909152919050565b5f5f5f5f6101008587031215615782575f5ffd5b61578c86866156e9565b935060a08501356003811061579f575f5ffd5b939693955050505060c08201359160e0013590565b5f5f604083850312156157c5575f5ffd5b82356157d081615285565b9150602083013567ffffffffffffffff8111156157eb575f5ffd5b6157f785828601615322565b9150509250929050565b5f5f5f5f60a08587031215615814575f5ffd5b61581d8561542e565b9350602085013592506080850186811115615836575f5ffd5b9396929550505060409290920191903590565b5f60208284031215615859575f5ffd5b813561237c81615285565b5f5f5f5f60808587031215615877575f5ffd5b8435615882816156d8565b93506020850135615892816156d8565b925060408501356158a2816156d8565b915060608501356158b2816156d8565b939692955090935050565b5f602082840312156158cd575f5ffd5b5035919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b60028110615936577f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b9052565b5f60a08201905061594c828451615901565b73ffffffffffffffffffffffffffffffffffffffff602084015116602083015263ffffffff6040840151166040830152606083015160608301526080830151608083015292915050565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b5f602082840312156159f9575f5ffd5b8135801515811461237c575f5ffd5b5f5f5f60e08486031215615a1a575f5ffd5b615a2485856156e9565b9560a0850135955060c0909401359392505050565b5f5f60208385031215615a4a575f5ffd5b823567ffffffffffffffff811115615a60575f5ffd5b615a6c858286016155e8565b90969095509350505050565b5f5f60408385031215615a89575f5ffd5b50508035926020909101359150565b5f60208284031215615aa8575f5ffd5b815160ff8116811461237c575f5ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b6001815b60018411156130d357808504811115615b0457615b04615ab8565b6001841615615b1257908102905b60019390931c928002615ae9565b5f82615b2e57506001610dcc565b81615b3a57505f610dcc565b8160018114615b505760028114615b5a57615b76565b6001915050610dcc565b60ff841115615b6b57615b6b615ab8565b50506001821b610dcc565b5060208310610133831016604e8410600b8410161715615b99575081810a610dcc565b615ba55f198484615ae5565b805f1904821115615bb857615bb8615ab8565b029392505050565b5f61237c60ff841683615b20565b8082028115828204841417610dcc57610dcc615ab8565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f82615c45577f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b500690565b5f5f5f5f60808587031215615c5d575f5ffd5b8451615c68816156d8565b6020860151909450615c79816156d8565b6040860151909350615c8a816156d8565b60608601519092506158b2816156d8565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b80820180821115610dcc57610dcc615ab8565b5f60208284031215615ceb575f5ffd5b5051919050565b5f5f198203615d0357615d03615ab8565b5060010190565b60608101615d188286615901565b602082019390935260400152919050565b5f60208284031215615d39575f5ffd5b815161237c81615285565b60208101610dcc8284615901565b81810381811115610dcc57610dcc615ab8565b5f6080828403128015615d76575f5ffd5b506040516080810167ffffffffffffffff81118282101715615d9a57615d9a6152a6565b6040528251615da8816156d8565b81526020830151615db8816156d8565b60208201526040830151615dcb816156d8565b60408201526060830151615dde816156d8565b60608201529392505050565b5f81615df857615df8615ab8565b505f190190565b60408101615e0d8285615901565b8260208301529392505050565b5f82518060208501845e5f92019182525091905056fea2646970667358221220c251c64827b3f4a735c1d82b304cf5b69969a54b51002026fa6ad9cbdf19fdf764736f6c634300081c0033",
}

// PairImplV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use PairImplV2MetaData.ABI instead.
var PairImplV2ABI = PairImplV2MetaData.ABI

// Deprecated: Use PairImplV2MetaData.Sigs instead.
// PairImplV2FuncSigs maps the 4-byte function signature to its string representation.
var PairImplV2FuncSigs = PairImplV2MetaData.Sigs

// PairImplV2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PairImplV2MetaData.Bin instead.
var PairImplV2Bin = PairImplV2MetaData.Bin

// DeployPairImplV2 deploys a new Ethereum contract, binding an instance of PairImplV2 to it.
func DeployPairImplV2(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PairImplV2, error) {
	parsed, err := PairImplV2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PairImplV2Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PairImplV2{PairImplV2Caller: PairImplV2Caller{contract: contract}, PairImplV2Transactor: PairImplV2Transactor{contract: contract}, PairImplV2Filterer: PairImplV2Filterer{contract: contract}}, nil
}

// PairImplV2 is an auto generated Go binding around an Ethereum contract.
type PairImplV2 struct {
	PairImplV2Caller     // Read-only binding to the contract
	PairImplV2Transactor // Write-only binding to the contract
	PairImplV2Filterer   // Log filterer for contract events
}

// PairImplV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type PairImplV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairImplV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type PairImplV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairImplV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PairImplV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairImplV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PairImplV2Session struct {
	Contract     *PairImplV2       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PairImplV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PairImplV2CallerSession struct {
	Contract *PairImplV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PairImplV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PairImplV2TransactorSession struct {
	Contract     *PairImplV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PairImplV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type PairImplV2Raw struct {
	Contract *PairImplV2 // Generic contract binding to access the raw methods on
}

// PairImplV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PairImplV2CallerRaw struct {
	Contract *PairImplV2Caller // Generic read-only contract binding to access the raw methods on
}

// PairImplV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PairImplV2TransactorRaw struct {
	Contract *PairImplV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewPairImplV2 creates a new instance of PairImplV2, bound to a specific deployed contract.
func NewPairImplV2(address common.Address, backend bind.ContractBackend) (*PairImplV2, error) {
	contract, err := bindPairImplV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PairImplV2{PairImplV2Caller: PairImplV2Caller{contract: contract}, PairImplV2Transactor: PairImplV2Transactor{contract: contract}, PairImplV2Filterer: PairImplV2Filterer{contract: contract}}, nil
}

// NewPairImplV2Caller creates a new read-only instance of PairImplV2, bound to a specific deployed contract.
func NewPairImplV2Caller(address common.Address, caller bind.ContractCaller) (*PairImplV2Caller, error) {
	contract, err := bindPairImplV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PairImplV2Caller{contract: contract}, nil
}

// NewPairImplV2Transactor creates a new write-only instance of PairImplV2, bound to a specific deployed contract.
func NewPairImplV2Transactor(address common.Address, transactor bind.ContractTransactor) (*PairImplV2Transactor, error) {
	contract, err := bindPairImplV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PairImplV2Transactor{contract: contract}, nil
}

// NewPairImplV2Filterer creates a new log filterer instance of PairImplV2, bound to a specific deployed contract.
func NewPairImplV2Filterer(address common.Address, filterer bind.ContractFilterer) (*PairImplV2Filterer, error) {
	contract, err := bindPairImplV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PairImplV2Filterer{contract: contract}, nil
}

// bindPairImplV2 binds a generic wrapper to an already deployed contract.
func bindPairImplV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PairImplV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PairImplV2 *PairImplV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PairImplV2.Contract.PairImplV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PairImplV2 *PairImplV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PairImplV2.Contract.PairImplV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PairImplV2 *PairImplV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PairImplV2.Contract.PairImplV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PairImplV2 *PairImplV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PairImplV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PairImplV2 *PairImplV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PairImplV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PairImplV2 *PairImplV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PairImplV2.Contract.contract.Transact(opts, method, params...)
}

// BASE is a free data retrieval call binding the contract method 0xec342ad0.
//
// Solidity: function BASE() view returns(address)
func (_PairImplV2 *PairImplV2Caller) BASE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PairImplV2.contract.Call(opts, &out, "BASE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BASE is a free data retrieval call binding the contract method 0xec342ad0.
//
// Solidity: function BASE() view returns(address)
func (_PairImplV2 *PairImplV2Session) BASE() (common.Address, error) {
	return _PairImplV2.Contract.BASE(&_PairImplV2.CallOpts)
}

// BASE is a free data retrieval call binding the contract method 0xec342ad0.
//
// Solidity: function BASE() view returns(address)
func (_PairImplV2 *PairImplV2CallerSession) BASE() (common.Address, error) {
	return _PairImplV2.Contract.BASE(&_PairImplV2.CallOpts)
}

// DENOMINATOR is a free data retrieval call binding the contract method 0x918f8674.
//
// Solidity: function DENOMINATOR() view returns(uint256)
func (_PairImplV2 *PairImplV2Caller) DENOMINATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PairImplV2.contract.Call(opts, &out, "DENOMINATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DENOMINATOR is a free data retrieval call binding the contract method 0x918f8674.
//
// Solidity: function DENOMINATOR() view returns(uint256)
func (_PairImplV2 *PairImplV2Session) DENOMINATOR() (*big.Int, error) {
	return _PairImplV2.Contract.DENOMINATOR(&_PairImplV2.CallOpts)
}

// DENOMINATOR is a free data retrieval call binding the contract method 0x918f8674.
//
// Solidity: function DENOMINATOR() view returns(uint256)
func (_PairImplV2 *PairImplV2CallerSession) DENOMINATOR() (*big.Int, error) {
	return _PairImplV2.Contract.DENOMINATOR(&_PairImplV2.CallOpts)
}

// MARKET is a free data retrieval call binding the contract method 0xf46f16c2.
//
// Solidity: function MARKET() view returns(address)
func (_PairImplV2 *PairImplV2Caller) MARKET(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PairImplV2.contract.Call(opts, &out, "MARKET")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MARKET is a free data retrieval call binding the contract method 0xf46f16c2.
//
// Solidity: function MARKET() view returns(address)
func (_PairImplV2 *PairImplV2Session) MARKET() (common.Address, error) {
	return _PairImplV2.Contract.MARKET(&_PairImplV2.CallOpts)
}

// MARKET is a free data retrieval call binding the contract method 0xf46f16c2.
//
// Solidity: function MARKET() view returns(address)
func (_PairImplV2 *PairImplV2CallerSession) MARKET() (common.Address, error) {
	return _PairImplV2.Contract.MARKET(&_PairImplV2.CallOpts)
}

// QUOTE is a free data retrieval call binding the contract method 0x9c579839.
//
// Solidity: function QUOTE() view returns(address)
func (_PairImplV2 *PairImplV2Caller) QUOTE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PairImplV2.contract.Call(opts, &out, "QUOTE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// QUOTE is a free data retrieval call binding the contract method 0x9c579839.
//
// Solidity: function QUOTE() view returns(address)
func (_PairImplV2 *PairImplV2Session) QUOTE() (common.Address, error) {
	return _PairImplV2.Contract.QUOTE(&_PairImplV2.CallOpts)
}

// QUOTE is a free data retrieval call binding the contract method 0x9c579839.
//
// Solidity: function QUOTE() view returns(address)
func (_PairImplV2 *PairImplV2CallerSession) QUOTE() (common.Address, error) {
	return _PairImplV2.Contract.QUOTE(&_PairImplV2.CallOpts)
}

// ROUTER is a free data retrieval call binding the contract method 0x32fe7b26.
//
// Solidity: function ROUTER() view returns(address)
func (_PairImplV2 *PairImplV2Caller) ROUTER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PairImplV2.contract.Call(opts, &out, "ROUTER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ROUTER is a free data retrieval call binding the contract method 0x32fe7b26.
//
// Solidity: function ROUTER() view returns(address)
func (_PairImplV2 *PairImplV2Session) ROUTER() (common.Address, error) {
	return _PairImplV2.Contract.ROUTER(&_PairImplV2.CallOpts)
}

// ROUTER is a free data retrieval call binding the contract method 0x32fe7b26.
//
// Solidity: function ROUTER() view returns(address)
func (_PairImplV2 *PairImplV2CallerSession) ROUTER() (common.Address, error) {
	return _PairImplV2.Contract.ROUTER(&_PairImplV2.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_PairImplV2 *PairImplV2Caller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PairImplV2.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_PairImplV2 *PairImplV2Session) UPGRADEINTERFACEVERSION() (string, error) {
	return _PairImplV2.Contract.UPGRADEINTERFACEVERSION(&_PairImplV2.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_PairImplV2 *PairImplV2CallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _PairImplV2.Contract.UPGRADEINTERFACEVERSION(&_PairImplV2.CallOpts)
}

// AccountReserves is a free data retrieval call binding the contract method 0x7aae3523.
//
// Solidity: function accountReserves(address account) view returns(uint256 base, uint256 quote)
func (_PairImplV2 *PairImplV2Caller) AccountReserves(opts *bind.CallOpts, account common.Address) (struct {
	Base  *big.Int
	Quote *big.Int
}, error) {
	var out []interface{}
	err := _PairImplV2.contract.Call(opts, &out, "accountReserves", account)

	outstruct := new(struct {
		Base  *big.Int
		Quote *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Base = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Quote = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// AccountReserves is a free data retrieval call binding the contract method 0x7aae3523.
//
// Solidity: function accountReserves(address account) view returns(uint256 base, uint256 quote)
func (_PairImplV2 *PairImplV2Session) AccountReserves(account common.Address) (struct {
	Base  *big.Int
	Quote *big.Int
}, error) {
	return _PairImplV2.Contract.AccountReserves(&_PairImplV2.CallOpts, account)
}

// AccountReserves is a free data retrieval call binding the contract method 0x7aae3523.
//
// Solidity: function accountReserves(address account) view returns(uint256 base, uint256 quote)
func (_PairImplV2 *PairImplV2CallerSession) AccountReserves(account common.Address) (struct {
	Base  *big.Int
	Quote *big.Int
}, error) {
	return _PairImplV2.Contract.AccountReserves(&_PairImplV2.CallOpts, account)
}

// BaseReserve is a free data retrieval call binding the contract method 0xdfdf2a72.
//
// Solidity: function baseReserve() view returns(uint256)
func (_PairImplV2 *PairImplV2Caller) BaseReserve(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PairImplV2.contract.Call(opts, &out, "baseReserve")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseReserve is a free data retrieval call binding the contract method 0xdfdf2a72.
//
// Solidity: function baseReserve() view returns(uint256)
func (_PairImplV2 *PairImplV2Session) BaseReserve() (*big.Int, error) {
	return _PairImplV2.Contract.BaseReserve(&_PairImplV2.CallOpts)
}

// BaseReserve is a free data retrieval call binding the contract method 0xdfdf2a72.
//
// Solidity: function baseReserve() view returns(uint256)
func (_PairImplV2 *PairImplV2CallerSession) BaseReserve() (*big.Int, error) {
	return _PairImplV2.Contract.BaseReserve(&_PairImplV2.CallOpts)
}

// FeeConfig is a free data retrieval call binding the contract method 0x1e5eb1d0.
//
// Solidity: function feeConfig() view returns(uint32 sellerMakerFeeBps, uint32 sellerTakerFeeBps, uint32 buyerMakerFeeBps, uint32 buyerTakerFeeBps)
func (_PairImplV2 *PairImplV2Caller) FeeConfig(opts *bind.CallOpts) (struct {
	SellerMakerFeeBps uint32
	SellerTakerFeeBps uint32
	BuyerMakerFeeBps  uint32
	BuyerTakerFeeBps  uint32
}, error) {
	var out []interface{}
	err := _PairImplV2.contract.Call(opts, &out, "feeConfig")

	outstruct := new(struct {
		SellerMakerFeeBps uint32
		SellerTakerFeeBps uint32
		BuyerMakerFeeBps  uint32
		BuyerTakerFeeBps  uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SellerMakerFeeBps = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.SellerTakerFeeBps = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.BuyerMakerFeeBps = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.BuyerTakerFeeBps = *abi.ConvertType(out[3], new(uint32)).(*uint32)

	return *outstruct, err

}

// FeeConfig is a free data retrieval call binding the contract method 0x1e5eb1d0.
//
// Solidity: function feeConfig() view returns(uint32 sellerMakerFeeBps, uint32 sellerTakerFeeBps, uint32 buyerMakerFeeBps, uint32 buyerTakerFeeBps)
func (_PairImplV2 *PairImplV2Session) FeeConfig() (struct {
	SellerMakerFeeBps uint32
	SellerTakerFeeBps uint32
	BuyerMakerFeeBps  uint32
	BuyerTakerFeeBps  uint32
}, error) {
	return _PairImplV2.Contract.FeeConfig(&_PairImplV2.CallOpts)
}

// FeeConfig is a free data retrieval call binding the contract method 0x1e5eb1d0.
//
// Solidity: function feeConfig() view returns(uint32 sellerMakerFeeBps, uint32 sellerTakerFeeBps, uint32 buyerMakerFeeBps, uint32 buyerTakerFeeBps)
func (_PairImplV2 *PairImplV2CallerSession) FeeConfig() (struct {
	SellerMakerFeeBps uint32
	SellerTakerFeeBps uint32
	BuyerMakerFeeBps  uint32
	BuyerTakerFeeBps  uint32
}, error) {
	return _PairImplV2.Contract.FeeConfig(&_PairImplV2.CallOpts)
}

// FindPrevPrice is a free data retrieval call binding the contract method 0x6c4cfbc8.
//
// Solidity: function findPrevPrice(uint8 side, uint256 price, uint256[2] adjacent, uint256 findMaxCount) view returns(uint256)
func (_PairImplV2 *PairImplV2Caller) FindPrevPrice(opts *bind.CallOpts, side uint8, price *big.Int, adjacent [2]*big.Int, findMaxCount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PairImplV2.contract.Call(opts, &out, "findPrevPrice", side, price, adjacent, findMaxCount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FindPrevPrice is a free data retrieval call binding the contract method 0x6c4cfbc8.
//
// Solidity: function findPrevPrice(uint8 side, uint256 price, uint256[2] adjacent, uint256 findMaxCount) view returns(uint256)
func (_PairImplV2 *PairImplV2Session) FindPrevPrice(side uint8, price *big.Int, adjacent [2]*big.Int, findMaxCount *big.Int) (*big.Int, error) {
	return _PairImplV2.Contract.FindPrevPrice(&_PairImplV2.CallOpts, side, price, adjacent, findMaxCount)
}

// FindPrevPrice is a free data retrieval call binding the contract method 0x6c4cfbc8.
//
// Solidity: function findPrevPrice(uint8 side, uint256 price, uint256[2] adjacent, uint256 findMaxCount) view returns(uint256)
func (_PairImplV2 *PairImplV2CallerSession) FindPrevPrice(side uint8, price *big.Int, adjacent [2]*big.Int, findMaxCount *big.Int) (*big.Int, error) {
	return _PairImplV2.Contract.FindPrevPrice(&_PairImplV2.CallOpts, side, price, adjacent, findMaxCount)
}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() view returns((address,address,uint256))
func (_PairImplV2 *PairImplV2Caller) GetConfig(opts *bind.CallOpts) (IPairConfig, error) {
	var out []interface{}
	err := _PairImplV2.contract.Call(opts, &out, "getConfig")

	if err != nil {
		return *new(IPairConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IPairConfig)).(*IPairConfig)

	return out0, err

}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() view returns((address,address,uint256))
func (_PairImplV2 *PairImplV2Session) GetConfig() (IPairConfig, error) {
	return _PairImplV2.Contract.GetConfig(&_PairImplV2.CallOpts)
}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() view returns((address,address,uint256))
func (_PairImplV2 *PairImplV2CallerSession) GetConfig() (IPairConfig, error) {
	return _PairImplV2.Contract.GetConfig(&_PairImplV2.CallOpts)
}

// GetEffectiveFees is a free data retrieval call binding the contract method 0x7cef243c.
//
// Solidity: function getEffectiveFees() view returns(uint32 sellerMakerFeeBps, uint32 sellerTakerFeeBps, uint32 buyerMakerFeeBps, uint32 buyerTakerFeeBps)
func (_PairImplV2 *PairImplV2Caller) GetEffectiveFees(opts *bind.CallOpts) (struct {
	SellerMakerFeeBps uint32
	SellerTakerFeeBps uint32
	BuyerMakerFeeBps  uint32
	BuyerTakerFeeBps  uint32
}, error) {
	var out []interface{}
	err := _PairImplV2.contract.Call(opts, &out, "getEffectiveFees")

	outstruct := new(struct {
		SellerMakerFeeBps uint32
		SellerTakerFeeBps uint32
		BuyerMakerFeeBps  uint32
		BuyerTakerFeeBps  uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SellerMakerFeeBps = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.SellerTakerFeeBps = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.BuyerMakerFeeBps = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.BuyerTakerFeeBps = *abi.ConvertType(out[3], new(uint32)).(*uint32)

	return *outstruct, err

}

// GetEffectiveFees is a free data retrieval call binding the contract method 0x7cef243c.
//
// Solidity: function getEffectiveFees() view returns(uint32 sellerMakerFeeBps, uint32 sellerTakerFeeBps, uint32 buyerMakerFeeBps, uint32 buyerTakerFeeBps)
func (_PairImplV2 *PairImplV2Session) GetEffectiveFees() (struct {
	SellerMakerFeeBps uint32
	SellerTakerFeeBps uint32
	BuyerMakerFeeBps  uint32
	BuyerTakerFeeBps  uint32
}, error) {
	return _PairImplV2.Contract.GetEffectiveFees(&_PairImplV2.CallOpts)
}

// GetEffectiveFees is a free data retrieval call binding the contract method 0x7cef243c.
//
// Solidity: function getEffectiveFees() view returns(uint32 sellerMakerFeeBps, uint32 sellerTakerFeeBps, uint32 buyerMakerFeeBps, uint32 buyerTakerFeeBps)
func (_PairImplV2 *PairImplV2CallerSession) GetEffectiveFees() (struct {
	SellerMakerFeeBps uint32
	SellerTakerFeeBps uint32
	BuyerMakerFeeBps  uint32
	BuyerTakerFeeBps  uint32
}, error) {
	return _PairImplV2.Contract.GetEffectiveFees(&_PairImplV2.CallOpts)
}

// LotSize is a free data retrieval call binding the contract method 0x4942f65f.
//
// Solidity: function lotSize() view returns(uint256)
func (_PairImplV2 *PairImplV2Caller) LotSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PairImplV2.contract.Call(opts, &out, "lotSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LotSize is a free data retrieval call binding the contract method 0x4942f65f.
//
// Solidity: function lotSize() view returns(uint256)
func (_PairImplV2 *PairImplV2Session) LotSize() (*big.Int, error) {
	return _PairImplV2.Contract.LotSize(&_PairImplV2.CallOpts)
}

// LotSize is a free data retrieval call binding the contract method 0x4942f65f.
//
// Solidity: function lotSize() view returns(uint256)
func (_PairImplV2 *PairImplV2CallerSession) LotSize() (*big.Int, error) {
	return _PairImplV2.Contract.LotSize(&_PairImplV2.CallOpts)
}

// MatchedAt is a free data retrieval call binding the contract method 0x4d6754f1.
//
// Solidity: function matchedAt() view returns(uint256)
func (_PairImplV2 *PairImplV2Caller) MatchedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PairImplV2.contract.Call(opts, &out, "matchedAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MatchedAt is a free data retrieval call binding the contract method 0x4d6754f1.
//
// Solidity: function matchedAt() view returns(uint256)
func (_PairImplV2 *PairImplV2Session) MatchedAt() (*big.Int, error) {
	return _PairImplV2.Contract.MatchedAt(&_PairImplV2.CallOpts)
}

// MatchedAt is a free data retrieval call binding the contract method 0x4d6754f1.
//
// Solidity: function matchedAt() view returns(uint256)
func (_PairImplV2 *PairImplV2CallerSession) MatchedAt() (*big.Int, error) {
	return _PairImplV2.Contract.MatchedAt(&_PairImplV2.CallOpts)
}

// MatchedPrice is a free data retrieval call binding the contract method 0x1edd8614.
//
// Solidity: function matchedPrice() view returns(uint256)
func (_PairImplV2 *PairImplV2Caller) MatchedPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PairImplV2.contract.Call(opts, &out, "matchedPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MatchedPrice is a free data retrieval call binding the contract method 0x1edd8614.
//
// Solidity: function matchedPrice() view returns(uint256)
func (_PairImplV2 *PairImplV2Session) MatchedPrice() (*big.Int, error) {
	return _PairImplV2.Contract.MatchedPrice(&_PairImplV2.CallOpts)
}

// MatchedPrice is a free data retrieval call binding the contract method 0x1edd8614.
//
// Solidity: function matchedPrice() view returns(uint256)
func (_PairImplV2 *PairImplV2CallerSession) MatchedPrice() (*big.Int, error) {
	return _PairImplV2.Contract.MatchedPrice(&_PairImplV2.CallOpts)
}

// MinTradeVolume is a free data retrieval call binding the contract method 0x7cbf6db2.
//
// Solidity: function minTradeVolume() view returns(uint256)
func (_PairImplV2 *PairImplV2Caller) MinTradeVolume(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PairImplV2.contract.Call(opts, &out, "minTradeVolume")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinTradeVolume is a free data retrieval call binding the contract method 0x7cbf6db2.
//
// Solidity: function minTradeVolume() view returns(uint256)
func (_PairImplV2 *PairImplV2Session) MinTradeVolume() (*big.Int, error) {
	return _PairImplV2.Contract.MinTradeVolume(&_PairImplV2.CallOpts)
}

// MinTradeVolume is a free data retrieval call binding the contract method 0x7cbf6db2.
//
// Solidity: function minTradeVolume() view returns(uint256)
func (_PairImplV2 *PairImplV2CallerSession) MinTradeVolume() (*big.Int, error) {
	return _PairImplV2.Contract.MinTradeVolume(&_PairImplV2.CallOpts)
}

// OrderById is a free data retrieval call binding the contract method 0x8d77eba5.
//
// Solidity: function orderById(uint256 id) view returns((uint8,address,uint32,uint256,uint256))
func (_PairImplV2 *PairImplV2Caller) OrderById(opts *bind.CallOpts, id *big.Int) (IPairOrder, error) {
	var out []interface{}
	err := _PairImplV2.contract.Call(opts, &out, "orderById", id)

	if err != nil {
		return *new(IPairOrder), err
	}

	out0 := *abi.ConvertType(out[0], new(IPairOrder)).(*IPairOrder)

	return out0, err

}

// OrderById is a free data retrieval call binding the contract method 0x8d77eba5.
//
// Solidity: function orderById(uint256 id) view returns((uint8,address,uint32,uint256,uint256))
func (_PairImplV2 *PairImplV2Session) OrderById(id *big.Int) (IPairOrder, error) {
	return _PairImplV2.Contract.OrderById(&_PairImplV2.CallOpts, id)
}

// OrderById is a free data retrieval call binding the contract method 0x8d77eba5.
//
// Solidity: function orderById(uint256 id) view returns((uint8,address,uint32,uint256,uint256))
func (_PairImplV2 *PairImplV2CallerSession) OrderById(id *big.Int) (IPairOrder, error) {
	return _PairImplV2.Contract.OrderById(&_PairImplV2.CallOpts, id)
}

// OrdersByPrices is a free data retrieval call binding the contract method 0x12a808cc.
//
// Solidity: function ordersByPrices(uint8 side, uint256[] prices) view returns(uint256[][])
func (_PairImplV2 *PairImplV2Caller) OrdersByPrices(opts *bind.CallOpts, side uint8, prices []*big.Int) ([][]*big.Int, error) {
	var out []interface{}
	err := _PairImplV2.contract.Call(opts, &out, "ordersByPrices", side, prices)

	if err != nil {
		return *new([][]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([][]*big.Int)).(*[][]*big.Int)

	return out0, err

}

// OrdersByPrices is a free data retrieval call binding the contract method 0x12a808cc.
//
// Solidity: function ordersByPrices(uint8 side, uint256[] prices) view returns(uint256[][])
func (_PairImplV2 *PairImplV2Session) OrdersByPrices(side uint8, prices []*big.Int) ([][]*big.Int, error) {
	return _PairImplV2.Contract.OrdersByPrices(&_PairImplV2.CallOpts, side, prices)
}

// OrdersByPrices is a free data retrieval call binding the contract method 0x12a808cc.
//
// Solidity: function ordersByPrices(uint8 side, uint256[] prices) view returns(uint256[][])
func (_PairImplV2 *PairImplV2CallerSession) OrdersByPrices(side uint8, prices []*big.Int) ([][]*big.Int, error) {
	return _PairImplV2.Contract.OrdersByPrices(&_PairImplV2.CallOpts, side, prices)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PairImplV2 *PairImplV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PairImplV2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PairImplV2 *PairImplV2Session) Owner() (common.Address, error) {
	return _PairImplV2.Contract.Owner(&_PairImplV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PairImplV2 *PairImplV2CallerSession) Owner() (common.Address, error) {
	return _PairImplV2.Contract.Owner(&_PairImplV2.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PairImplV2 *PairImplV2Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PairImplV2.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PairImplV2 *PairImplV2Session) Paused() (bool, error) {
	return _PairImplV2.Contract.Paused(&_PairImplV2.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PairImplV2 *PairImplV2CallerSession) Paused() (bool, error) {
	return _PairImplV2.Contract.Paused(&_PairImplV2.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_PairImplV2 *PairImplV2Caller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PairImplV2.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_PairImplV2 *PairImplV2Session) ProxiableUUID() ([32]byte, error) {
	return _PairImplV2.Contract.ProxiableUUID(&_PairImplV2.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_PairImplV2 *PairImplV2CallerSession) ProxiableUUID() ([32]byte, error) {
	return _PairImplV2.Contract.ProxiableUUID(&_PairImplV2.CallOpts)
}

// QuoteReserve is a free data retrieval call binding the contract method 0x9da771f4.
//
// Solidity: function quoteReserve() view returns(uint256)
func (_PairImplV2 *PairImplV2Caller) QuoteReserve(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PairImplV2.contract.Call(opts, &out, "quoteReserve")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuoteReserve is a free data retrieval call binding the contract method 0x9da771f4.
//
// Solidity: function quoteReserve() view returns(uint256)
func (_PairImplV2 *PairImplV2Session) QuoteReserve() (*big.Int, error) {
	return _PairImplV2.Contract.QuoteReserve(&_PairImplV2.CallOpts)
}

// QuoteReserve is a free data retrieval call binding the contract method 0x9da771f4.
//
// Solidity: function quoteReserve() view returns(uint256)
func (_PairImplV2 *PairImplV2CallerSession) QuoteReserve() (*big.Int, error) {
	return _PairImplV2.Contract.QuoteReserve(&_PairImplV2.CallOpts)
}

// TickSize is a free data retrieval call binding the contract method 0xf210d087.
//
// Solidity: function tickSize() view returns(uint256)
func (_PairImplV2 *PairImplV2Caller) TickSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PairImplV2.contract.Call(opts, &out, "tickSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TickSize is a free data retrieval call binding the contract method 0xf210d087.
//
// Solidity: function tickSize() view returns(uint256)
func (_PairImplV2 *PairImplV2Session) TickSize() (*big.Int, error) {
	return _PairImplV2.Contract.TickSize(&_PairImplV2.CallOpts)
}

// TickSize is a free data retrieval call binding the contract method 0xf210d087.
//
// Solidity: function tickSize() view returns(uint256)
func (_PairImplV2 *PairImplV2CallerSession) TickSize() (*big.Int, error) {
	return _PairImplV2.Contract.TickSize(&_PairImplV2.CallOpts)
}

// TickSizes is a free data retrieval call binding the contract method 0xd6e9cc2a.
//
// Solidity: function tickSizes() view returns(uint256 tick, uint256 lot)
func (_PairImplV2 *PairImplV2Caller) TickSizes(opts *bind.CallOpts) (struct {
	Tick *big.Int
	Lot  *big.Int
}, error) {
	var out []interface{}
	err := _PairImplV2.contract.Call(opts, &out, "tickSizes")

	outstruct := new(struct {
		Tick *big.Int
		Lot  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Tick = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Lot = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TickSizes is a free data retrieval call binding the contract method 0xd6e9cc2a.
//
// Solidity: function tickSizes() view returns(uint256 tick, uint256 lot)
func (_PairImplV2 *PairImplV2Session) TickSizes() (struct {
	Tick *big.Int
	Lot  *big.Int
}, error) {
	return _PairImplV2.Contract.TickSizes(&_PairImplV2.CallOpts)
}

// TickSizes is a free data retrieval call binding the contract method 0xd6e9cc2a.
//
// Solidity: function tickSizes() view returns(uint256 tick, uint256 lot)
func (_PairImplV2 *PairImplV2CallerSession) TickSizes() (struct {
	Tick *big.Int
	Lot  *big.Int
}, error) {
	return _PairImplV2.Contract.TickSizes(&_PairImplV2.CallOpts)
}

// Ticks is a free data retrieval call binding the contract method 0x2cfffaf6.
//
// Solidity: function ticks() view returns(uint256[] sellPrices, uint256[] buyPrices)
func (_PairImplV2 *PairImplV2Caller) Ticks(opts *bind.CallOpts) (struct {
	SellPrices []*big.Int
	BuyPrices  []*big.Int
}, error) {
	var out []interface{}
	err := _PairImplV2.contract.Call(opts, &out, "ticks")

	outstruct := new(struct {
		SellPrices []*big.Int
		BuyPrices  []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SellPrices = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.BuyPrices = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Ticks is a free data retrieval call binding the contract method 0x2cfffaf6.
//
// Solidity: function ticks() view returns(uint256[] sellPrices, uint256[] buyPrices)
func (_PairImplV2 *PairImplV2Session) Ticks() (struct {
	SellPrices []*big.Int
	BuyPrices  []*big.Int
}, error) {
	return _PairImplV2.Contract.Ticks(&_PairImplV2.CallOpts)
}

// Ticks is a free data retrieval call binding the contract method 0x2cfffaf6.
//
// Solidity: function ticks() view returns(uint256[] sellPrices, uint256[] buyPrices)
func (_PairImplV2 *PairImplV2CallerSession) Ticks() (struct {
	SellPrices []*big.Int
	BuyPrices  []*big.Int
}, error) {
	return _PairImplV2.Contract.Ticks(&_PairImplV2.CallOpts)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x1ec482d7.
//
// Solidity: function cancelOrder(address caller, uint256[] orderIds) returns()
func (_PairImplV2 *PairImplV2Transactor) CancelOrder(opts *bind.TransactOpts, caller common.Address, orderIds []*big.Int) (*types.Transaction, error) {
	return _PairImplV2.contract.Transact(opts, "cancelOrder", caller, orderIds)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x1ec482d7.
//
// Solidity: function cancelOrder(address caller, uint256[] orderIds) returns()
func (_PairImplV2 *PairImplV2Session) CancelOrder(caller common.Address, orderIds []*big.Int) (*types.Transaction, error) {
	return _PairImplV2.Contract.CancelOrder(&_PairImplV2.TransactOpts, caller, orderIds)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x1ec482d7.
//
// Solidity: function cancelOrder(address caller, uint256[] orderIds) returns()
func (_PairImplV2 *PairImplV2TransactorSession) CancelOrder(caller common.Address, orderIds []*big.Int) (*types.Transaction, error) {
	return _PairImplV2.Contract.CancelOrder(&_PairImplV2.TransactOpts, caller, orderIds)
}

// EmergencyCancelOrder is a paid mutator transaction binding the contract method 0xe077dd8a.
//
// Solidity: function emergencyCancelOrder(uint256[] orderIds) returns()
func (_PairImplV2 *PairImplV2Transactor) EmergencyCancelOrder(opts *bind.TransactOpts, orderIds []*big.Int) (*types.Transaction, error) {
	return _PairImplV2.contract.Transact(opts, "emergencyCancelOrder", orderIds)
}

// EmergencyCancelOrder is a paid mutator transaction binding the contract method 0xe077dd8a.
//
// Solidity: function emergencyCancelOrder(uint256[] orderIds) returns()
func (_PairImplV2 *PairImplV2Session) EmergencyCancelOrder(orderIds []*big.Int) (*types.Transaction, error) {
	return _PairImplV2.Contract.EmergencyCancelOrder(&_PairImplV2.TransactOpts, orderIds)
}

// EmergencyCancelOrder is a paid mutator transaction binding the contract method 0xe077dd8a.
//
// Solidity: function emergencyCancelOrder(uint256[] orderIds) returns()
func (_PairImplV2 *PairImplV2TransactorSession) EmergencyCancelOrder(orderIds []*big.Int) (*types.Transaction, error) {
	return _PairImplV2.Contract.EmergencyCancelOrder(&_PairImplV2.TransactOpts, orderIds)
}

// Initialize is a paid mutator transaction binding the contract method 0x02a204c6.
//
// Solidity: function initialize(address router, address quote, address base, uint256 _tickSize, uint256 _lotSize, bytes feeData) returns()
func (_PairImplV2 *PairImplV2Transactor) Initialize(opts *bind.TransactOpts, router common.Address, quote common.Address, base common.Address, _tickSize *big.Int, _lotSize *big.Int, feeData []byte) (*types.Transaction, error) {
	return _PairImplV2.contract.Transact(opts, "initialize", router, quote, base, _tickSize, _lotSize, feeData)
}

// Initialize is a paid mutator transaction binding the contract method 0x02a204c6.
//
// Solidity: function initialize(address router, address quote, address base, uint256 _tickSize, uint256 _lotSize, bytes feeData) returns()
func (_PairImplV2 *PairImplV2Session) Initialize(router common.Address, quote common.Address, base common.Address, _tickSize *big.Int, _lotSize *big.Int, feeData []byte) (*types.Transaction, error) {
	return _PairImplV2.Contract.Initialize(&_PairImplV2.TransactOpts, router, quote, base, _tickSize, _lotSize, feeData)
}

// Initialize is a paid mutator transaction binding the contract method 0x02a204c6.
//
// Solidity: function initialize(address router, address quote, address base, uint256 _tickSize, uint256 _lotSize, bytes feeData) returns()
func (_PairImplV2 *PairImplV2TransactorSession) Initialize(router common.Address, quote common.Address, base common.Address, _tickSize *big.Int, _lotSize *big.Int, feeData []byte) (*types.Transaction, error) {
	return _PairImplV2.Contract.Initialize(&_PairImplV2.TransactOpts, router, quote, base, _tickSize, _lotSize, feeData)
}

// SetPairFees is a paid mutator transaction binding the contract method 0x803a0386.
//
// Solidity: function setPairFees(uint32 sellerMakerFeeBps_, uint32 sellerTakerFeeBps_, uint32 buyerMakerFeeBps_, uint32 buyerTakerFeeBps_) returns()
func (_PairImplV2 *PairImplV2Transactor) SetPairFees(opts *bind.TransactOpts, sellerMakerFeeBps_ uint32, sellerTakerFeeBps_ uint32, buyerMakerFeeBps_ uint32, buyerTakerFeeBps_ uint32) (*types.Transaction, error) {
	return _PairImplV2.contract.Transact(opts, "setPairFees", sellerMakerFeeBps_, sellerTakerFeeBps_, buyerMakerFeeBps_, buyerTakerFeeBps_)
}

// SetPairFees is a paid mutator transaction binding the contract method 0x803a0386.
//
// Solidity: function setPairFees(uint32 sellerMakerFeeBps_, uint32 sellerTakerFeeBps_, uint32 buyerMakerFeeBps_, uint32 buyerTakerFeeBps_) returns()
func (_PairImplV2 *PairImplV2Session) SetPairFees(sellerMakerFeeBps_ uint32, sellerTakerFeeBps_ uint32, buyerMakerFeeBps_ uint32, buyerTakerFeeBps_ uint32) (*types.Transaction, error) {
	return _PairImplV2.Contract.SetPairFees(&_PairImplV2.TransactOpts, sellerMakerFeeBps_, sellerTakerFeeBps_, buyerMakerFeeBps_, buyerTakerFeeBps_)
}

// SetPairFees is a paid mutator transaction binding the contract method 0x803a0386.
//
// Solidity: function setPairFees(uint32 sellerMakerFeeBps_, uint32 sellerTakerFeeBps_, uint32 buyerMakerFeeBps_, uint32 buyerTakerFeeBps_) returns()
func (_PairImplV2 *PairImplV2TransactorSession) SetPairFees(sellerMakerFeeBps_ uint32, sellerTakerFeeBps_ uint32, buyerMakerFeeBps_ uint32, buyerTakerFeeBps_ uint32) (*types.Transaction, error) {
	return _PairImplV2.Contract.SetPairFees(&_PairImplV2.TransactOpts, sellerMakerFeeBps_, sellerTakerFeeBps_, buyerMakerFeeBps_, buyerTakerFeeBps_)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool pause) returns()
func (_PairImplV2 *PairImplV2Transactor) SetPause(opts *bind.TransactOpts, pause bool) (*types.Transaction, error) {
	return _PairImplV2.contract.Transact(opts, "setPause", pause)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool pause) returns()
func (_PairImplV2 *PairImplV2Session) SetPause(pause bool) (*types.Transaction, error) {
	return _PairImplV2.Contract.SetPause(&_PairImplV2.TransactOpts, pause)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool pause) returns()
func (_PairImplV2 *PairImplV2TransactorSession) SetPause(pause bool) (*types.Transaction, error) {
	return _PairImplV2.Contract.SetPause(&_PairImplV2.TransactOpts, pause)
}

// SetTickSize is a paid mutator transaction binding the contract method 0xfe566349.
//
// Solidity: function setTickSize(uint256 _lotSize, uint256 _tickSize) returns()
func (_PairImplV2 *PairImplV2Transactor) SetTickSize(opts *bind.TransactOpts, _lotSize *big.Int, _tickSize *big.Int) (*types.Transaction, error) {
	return _PairImplV2.contract.Transact(opts, "setTickSize", _lotSize, _tickSize)
}

// SetTickSize is a paid mutator transaction binding the contract method 0xfe566349.
//
// Solidity: function setTickSize(uint256 _lotSize, uint256 _tickSize) returns()
func (_PairImplV2 *PairImplV2Session) SetTickSize(_lotSize *big.Int, _tickSize *big.Int) (*types.Transaction, error) {
	return _PairImplV2.Contract.SetTickSize(&_PairImplV2.TransactOpts, _lotSize, _tickSize)
}

// SetTickSize is a paid mutator transaction binding the contract method 0xfe566349.
//
// Solidity: function setTickSize(uint256 _lotSize, uint256 _tickSize) returns()
func (_PairImplV2 *PairImplV2TransactorSession) SetTickSize(_lotSize *big.Int, _tickSize *big.Int) (*types.Transaction, error) {
	return _PairImplV2.Contract.SetTickSize(&_PairImplV2.TransactOpts, _lotSize, _tickSize)
}

// Skim is a paid mutator transaction binding the contract method 0x1a0361fb.
//
// Solidity: function skim(address erc20, address to, uint256 amount) returns()
func (_PairImplV2 *PairImplV2Transactor) Skim(opts *bind.TransactOpts, erc20 common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PairImplV2.contract.Transact(opts, "skim", erc20, to, amount)
}

// Skim is a paid mutator transaction binding the contract method 0x1a0361fb.
//
// Solidity: function skim(address erc20, address to, uint256 amount) returns()
func (_PairImplV2 *PairImplV2Session) Skim(erc20 common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PairImplV2.Contract.Skim(&_PairImplV2.TransactOpts, erc20, to, amount)
}

// Skim is a paid mutator transaction binding the contract method 0x1a0361fb.
//
// Solidity: function skim(address erc20, address to, uint256 amount) returns()
func (_PairImplV2 *PairImplV2TransactorSession) Skim(erc20 common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PairImplV2.Contract.Skim(&_PairImplV2.TransactOpts, erc20, to, amount)
}

// SubmitLimitOrder is a paid mutator transaction binding the contract method 0x2dd6a00a.
//
// Solidity: function submitLimitOrder((uint8,address,uint32,uint256,uint256) order, uint8 constraints, uint256 prevPrice, uint256 maxMatchCount) returns(uint256 orderId)
func (_PairImplV2 *PairImplV2Transactor) SubmitLimitOrder(opts *bind.TransactOpts, order IPairOrder, constraints uint8, prevPrice *big.Int, maxMatchCount *big.Int) (*types.Transaction, error) {
	return _PairImplV2.contract.Transact(opts, "submitLimitOrder", order, constraints, prevPrice, maxMatchCount)
}

// SubmitLimitOrder is a paid mutator transaction binding the contract method 0x2dd6a00a.
//
// Solidity: function submitLimitOrder((uint8,address,uint32,uint256,uint256) order, uint8 constraints, uint256 prevPrice, uint256 maxMatchCount) returns(uint256 orderId)
func (_PairImplV2 *PairImplV2Session) SubmitLimitOrder(order IPairOrder, constraints uint8, prevPrice *big.Int, maxMatchCount *big.Int) (*types.Transaction, error) {
	return _PairImplV2.Contract.SubmitLimitOrder(&_PairImplV2.TransactOpts, order, constraints, prevPrice, maxMatchCount)
}

// SubmitLimitOrder is a paid mutator transaction binding the contract method 0x2dd6a00a.
//
// Solidity: function submitLimitOrder((uint8,address,uint32,uint256,uint256) order, uint8 constraints, uint256 prevPrice, uint256 maxMatchCount) returns(uint256 orderId)
func (_PairImplV2 *PairImplV2TransactorSession) SubmitLimitOrder(order IPairOrder, constraints uint8, prevPrice *big.Int, maxMatchCount *big.Int) (*types.Transaction, error) {
	return _PairImplV2.Contract.SubmitLimitOrder(&_PairImplV2.TransactOpts, order, constraints, prevPrice, maxMatchCount)
}

// SubmitMarketOrder is a paid mutator transaction binding the contract method 0xcc50e36a.
//
// Solidity: function submitMarketOrder((uint8,address,uint32,uint256,uint256) order, uint256 spendAmount, uint256 maxMatchCount) returns()
func (_PairImplV2 *PairImplV2Transactor) SubmitMarketOrder(opts *bind.TransactOpts, order IPairOrder, spendAmount *big.Int, maxMatchCount *big.Int) (*types.Transaction, error) {
	return _PairImplV2.contract.Transact(opts, "submitMarketOrder", order, spendAmount, maxMatchCount)
}

// SubmitMarketOrder is a paid mutator transaction binding the contract method 0xcc50e36a.
//
// Solidity: function submitMarketOrder((uint8,address,uint32,uint256,uint256) order, uint256 spendAmount, uint256 maxMatchCount) returns()
func (_PairImplV2 *PairImplV2Session) SubmitMarketOrder(order IPairOrder, spendAmount *big.Int, maxMatchCount *big.Int) (*types.Transaction, error) {
	return _PairImplV2.Contract.SubmitMarketOrder(&_PairImplV2.TransactOpts, order, spendAmount, maxMatchCount)
}

// SubmitMarketOrder is a paid mutator transaction binding the contract method 0xcc50e36a.
//
// Solidity: function submitMarketOrder((uint8,address,uint32,uint256,uint256) order, uint256 spendAmount, uint256 maxMatchCount) returns()
func (_PairImplV2 *PairImplV2TransactorSession) SubmitMarketOrder(order IPairOrder, spendAmount *big.Int, maxMatchCount *big.Int) (*types.Transaction, error) {
	return _PairImplV2.Contract.SubmitMarketOrder(&_PairImplV2.TransactOpts, order, spendAmount, maxMatchCount)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_PairImplV2 *PairImplV2Transactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _PairImplV2.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_PairImplV2 *PairImplV2Session) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _PairImplV2.Contract.UpgradeToAndCall(&_PairImplV2.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_PairImplV2 *PairImplV2TransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _PairImplV2.Contract.UpgradeToAndCall(&_PairImplV2.TransactOpts, newImplementation, data)
}

// PairImplV2FeeCollectIterator is returned from FilterFeeCollect and is used to iterate over the raw logs and unpacked data for FeeCollect events raised by the PairImplV2 contract.
type PairImplV2FeeCollectIterator struct {
	Event *PairImplV2FeeCollect // Event containing the contract specifics and raw log

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
func (it *PairImplV2FeeCollectIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PairImplV2FeeCollect)
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
		it.Event = new(PairImplV2FeeCollect)
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
func (it *PairImplV2FeeCollectIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PairImplV2FeeCollectIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PairImplV2FeeCollect represents a FeeCollect event raised by the PairImplV2 contract.
type PairImplV2FeeCollect struct {
	OrderId   *big.Int
	Owner     common.Address
	Amount    *big.Int
	Recipient common.Address
	FeeBps    *big.Int
	Fee       *big.Int
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFeeCollect is a free log retrieval operation binding the contract event 0x9814b2b2f43f9ea51ad7e8760fda094ffa3012bfeb10c76a9656389af1895603.
//
// Solidity: event FeeCollect(uint256 indexed orderId, address indexed owner, uint256 amount, address indexed recipient, uint256 feeBps, uint256 fee, uint256 value)
func (_PairImplV2 *PairImplV2Filterer) FilterFeeCollect(opts *bind.FilterOpts, orderId []*big.Int, owner []common.Address, recipient []common.Address) (*PairImplV2FeeCollectIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _PairImplV2.contract.FilterLogs(opts, "FeeCollect", orderIdRule, ownerRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &PairImplV2FeeCollectIterator{contract: _PairImplV2.contract, event: "FeeCollect", logs: logs, sub: sub}, nil
}

// WatchFeeCollect is a free log subscription operation binding the contract event 0x9814b2b2f43f9ea51ad7e8760fda094ffa3012bfeb10c76a9656389af1895603.
//
// Solidity: event FeeCollect(uint256 indexed orderId, address indexed owner, uint256 amount, address indexed recipient, uint256 feeBps, uint256 fee, uint256 value)
func (_PairImplV2 *PairImplV2Filterer) WatchFeeCollect(opts *bind.WatchOpts, sink chan<- *PairImplV2FeeCollect, orderId []*big.Int, owner []common.Address, recipient []common.Address) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _PairImplV2.contract.WatchLogs(opts, "FeeCollect", orderIdRule, ownerRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PairImplV2FeeCollect)
				if err := _PairImplV2.contract.UnpackLog(event, "FeeCollect", log); err != nil {
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

// ParseFeeCollect is a log parse operation binding the contract event 0x9814b2b2f43f9ea51ad7e8760fda094ffa3012bfeb10c76a9656389af1895603.
//
// Solidity: event FeeCollect(uint256 indexed orderId, address indexed owner, uint256 amount, address indexed recipient, uint256 feeBps, uint256 fee, uint256 value)
func (_PairImplV2 *PairImplV2Filterer) ParseFeeCollect(log types.Log) (*PairImplV2FeeCollect, error) {
	event := new(PairImplV2FeeCollect)
	if err := _PairImplV2.contract.UnpackLog(event, "FeeCollect", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PairImplV2InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the PairImplV2 contract.
type PairImplV2InitializedIterator struct {
	Event *PairImplV2Initialized // Event containing the contract specifics and raw log

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
func (it *PairImplV2InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PairImplV2Initialized)
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
		it.Event = new(PairImplV2Initialized)
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
func (it *PairImplV2InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PairImplV2InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PairImplV2Initialized represents a Initialized event raised by the PairImplV2 contract.
type PairImplV2Initialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_PairImplV2 *PairImplV2Filterer) FilterInitialized(opts *bind.FilterOpts) (*PairImplV2InitializedIterator, error) {

	logs, sub, err := _PairImplV2.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PairImplV2InitializedIterator{contract: _PairImplV2.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_PairImplV2 *PairImplV2Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PairImplV2Initialized) (event.Subscription, error) {

	logs, sub, err := _PairImplV2.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PairImplV2Initialized)
				if err := _PairImplV2.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_PairImplV2 *PairImplV2Filterer) ParseInitialized(log types.Log) (*PairImplV2Initialized, error) {
	event := new(PairImplV2Initialized)
	if err := _PairImplV2.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PairImplV2OrderClosedIterator is returned from FilterOrderClosed and is used to iterate over the raw logs and unpacked data for OrderClosed events raised by the PairImplV2 contract.
type PairImplV2OrderClosedIterator struct {
	Event *PairImplV2OrderClosed // Event containing the contract specifics and raw log

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
func (it *PairImplV2OrderClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PairImplV2OrderClosed)
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
		it.Event = new(PairImplV2OrderClosed)
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
func (it *PairImplV2OrderClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PairImplV2OrderClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PairImplV2OrderClosed represents a OrderClosed event raised by the PairImplV2 contract.
type PairImplV2OrderClosed struct {
	OrderId   *big.Int
	CloseType uint8
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOrderClosed is a free log retrieval operation binding the contract event 0x78615f01112d8477ae8d6fc9a510bcfdf935bc3ab7b02d0ff2316f5b44bd597a.
//
// Solidity: event OrderClosed(uint256 indexed orderId, uint8 indexed closeType, uint256 timestamp)
func (_PairImplV2 *PairImplV2Filterer) FilterOrderClosed(opts *bind.FilterOpts, orderId []*big.Int, closeType []uint8) (*PairImplV2OrderClosedIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var closeTypeRule []interface{}
	for _, closeTypeItem := range closeType {
		closeTypeRule = append(closeTypeRule, closeTypeItem)
	}

	logs, sub, err := _PairImplV2.contract.FilterLogs(opts, "OrderClosed", orderIdRule, closeTypeRule)
	if err != nil {
		return nil, err
	}
	return &PairImplV2OrderClosedIterator{contract: _PairImplV2.contract, event: "OrderClosed", logs: logs, sub: sub}, nil
}

// WatchOrderClosed is a free log subscription operation binding the contract event 0x78615f01112d8477ae8d6fc9a510bcfdf935bc3ab7b02d0ff2316f5b44bd597a.
//
// Solidity: event OrderClosed(uint256 indexed orderId, uint8 indexed closeType, uint256 timestamp)
func (_PairImplV2 *PairImplV2Filterer) WatchOrderClosed(opts *bind.WatchOpts, sink chan<- *PairImplV2OrderClosed, orderId []*big.Int, closeType []uint8) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var closeTypeRule []interface{}
	for _, closeTypeItem := range closeType {
		closeTypeRule = append(closeTypeRule, closeTypeItem)
	}

	logs, sub, err := _PairImplV2.contract.WatchLogs(opts, "OrderClosed", orderIdRule, closeTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PairImplV2OrderClosed)
				if err := _PairImplV2.contract.UnpackLog(event, "OrderClosed", log); err != nil {
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

// ParseOrderClosed is a log parse operation binding the contract event 0x78615f01112d8477ae8d6fc9a510bcfdf935bc3ab7b02d0ff2316f5b44bd597a.
//
// Solidity: event OrderClosed(uint256 indexed orderId, uint8 indexed closeType, uint256 timestamp)
func (_PairImplV2 *PairImplV2Filterer) ParseOrderClosed(log types.Log) (*PairImplV2OrderClosed, error) {
	event := new(PairImplV2OrderClosed)
	if err := _PairImplV2.contract.UnpackLog(event, "OrderClosed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PairImplV2OrderCreatedIterator is returned from FilterOrderCreated and is used to iterate over the raw logs and unpacked data for OrderCreated events raised by the PairImplV2 contract.
type PairImplV2OrderCreatedIterator struct {
	Event *PairImplV2OrderCreated // Event containing the contract specifics and raw log

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
func (it *PairImplV2OrderCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PairImplV2OrderCreated)
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
		it.Event = new(PairImplV2OrderCreated)
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
func (it *PairImplV2OrderCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PairImplV2OrderCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PairImplV2OrderCreated represents a OrderCreated event raised by the PairImplV2 contract.
type PairImplV2OrderCreated struct {
	Owner     common.Address
	OrderId   *big.Int
	Side      uint8
	Price     *big.Int
	Amount    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOrderCreated is a free log retrieval operation binding the contract event 0x7cde7559ef00f8c3d9244d3fde19e7500979d64f6f9c086a3836a4926abbe165.
//
// Solidity: event OrderCreated(address indexed owner, uint256 indexed orderId, uint8 indexed side, uint256 price, uint256 amount, uint256 timestamp)
func (_PairImplV2 *PairImplV2Filterer) FilterOrderCreated(opts *bind.FilterOpts, owner []common.Address, orderId []*big.Int, side []uint8) (*PairImplV2OrderCreatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var sideRule []interface{}
	for _, sideItem := range side {
		sideRule = append(sideRule, sideItem)
	}

	logs, sub, err := _PairImplV2.contract.FilterLogs(opts, "OrderCreated", ownerRule, orderIdRule, sideRule)
	if err != nil {
		return nil, err
	}
	return &PairImplV2OrderCreatedIterator{contract: _PairImplV2.contract, event: "OrderCreated", logs: logs, sub: sub}, nil
}

// WatchOrderCreated is a free log subscription operation binding the contract event 0x7cde7559ef00f8c3d9244d3fde19e7500979d64f6f9c086a3836a4926abbe165.
//
// Solidity: event OrderCreated(address indexed owner, uint256 indexed orderId, uint8 indexed side, uint256 price, uint256 amount, uint256 timestamp)
func (_PairImplV2 *PairImplV2Filterer) WatchOrderCreated(opts *bind.WatchOpts, sink chan<- *PairImplV2OrderCreated, owner []common.Address, orderId []*big.Int, side []uint8) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var sideRule []interface{}
	for _, sideItem := range side {
		sideRule = append(sideRule, sideItem)
	}

	logs, sub, err := _PairImplV2.contract.WatchLogs(opts, "OrderCreated", ownerRule, orderIdRule, sideRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PairImplV2OrderCreated)
				if err := _PairImplV2.contract.UnpackLog(event, "OrderCreated", log); err != nil {
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

// ParseOrderCreated is a log parse operation binding the contract event 0x7cde7559ef00f8c3d9244d3fde19e7500979d64f6f9c086a3836a4926abbe165.
//
// Solidity: event OrderCreated(address indexed owner, uint256 indexed orderId, uint8 indexed side, uint256 price, uint256 amount, uint256 timestamp)
func (_PairImplV2 *PairImplV2Filterer) ParseOrderCreated(log types.Log) (*PairImplV2OrderCreated, error) {
	event := new(PairImplV2OrderCreated)
	if err := _PairImplV2.contract.UnpackLog(event, "OrderCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PairImplV2OrderMatchedIterator is returned from FilterOrderMatched and is used to iterate over the raw logs and unpacked data for OrderMatched events raised by the PairImplV2 contract.
type PairImplV2OrderMatchedIterator struct {
	Event *PairImplV2OrderMatched // Event containing the contract specifics and raw log

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
func (it *PairImplV2OrderMatchedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PairImplV2OrderMatched)
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
		it.Event = new(PairImplV2OrderMatched)
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
func (it *PairImplV2OrderMatchedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PairImplV2OrderMatchedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PairImplV2OrderMatched represents a OrderMatched event raised by the PairImplV2 contract.
type PairImplV2OrderMatched struct {
	SellId    *big.Int
	BuyId     *big.Int
	Price     *big.Int
	Amount    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOrderMatched is a free log retrieval operation binding the contract event 0x6a6896b1e6131e0b8ebae43fdc84cb0178a6eacdf4ee63f15dabae48729a8a03.
//
// Solidity: event OrderMatched(uint256 indexed sellId, uint256 indexed buyId, uint256 indexed price, uint256 amount, uint256 timestamp)
func (_PairImplV2 *PairImplV2Filterer) FilterOrderMatched(opts *bind.FilterOpts, sellId []*big.Int, buyId []*big.Int, price []*big.Int) (*PairImplV2OrderMatchedIterator, error) {

	var sellIdRule []interface{}
	for _, sellIdItem := range sellId {
		sellIdRule = append(sellIdRule, sellIdItem)
	}
	var buyIdRule []interface{}
	for _, buyIdItem := range buyId {
		buyIdRule = append(buyIdRule, buyIdItem)
	}
	var priceRule []interface{}
	for _, priceItem := range price {
		priceRule = append(priceRule, priceItem)
	}

	logs, sub, err := _PairImplV2.contract.FilterLogs(opts, "OrderMatched", sellIdRule, buyIdRule, priceRule)
	if err != nil {
		return nil, err
	}
	return &PairImplV2OrderMatchedIterator{contract: _PairImplV2.contract, event: "OrderMatched", logs: logs, sub: sub}, nil
}

// WatchOrderMatched is a free log subscription operation binding the contract event 0x6a6896b1e6131e0b8ebae43fdc84cb0178a6eacdf4ee63f15dabae48729a8a03.
//
// Solidity: event OrderMatched(uint256 indexed sellId, uint256 indexed buyId, uint256 indexed price, uint256 amount, uint256 timestamp)
func (_PairImplV2 *PairImplV2Filterer) WatchOrderMatched(opts *bind.WatchOpts, sink chan<- *PairImplV2OrderMatched, sellId []*big.Int, buyId []*big.Int, price []*big.Int) (event.Subscription, error) {

	var sellIdRule []interface{}
	for _, sellIdItem := range sellId {
		sellIdRule = append(sellIdRule, sellIdItem)
	}
	var buyIdRule []interface{}
	for _, buyIdItem := range buyId {
		buyIdRule = append(buyIdRule, buyIdItem)
	}
	var priceRule []interface{}
	for _, priceItem := range price {
		priceRule = append(priceRule, priceItem)
	}

	logs, sub, err := _PairImplV2.contract.WatchLogs(opts, "OrderMatched", sellIdRule, buyIdRule, priceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PairImplV2OrderMatched)
				if err := _PairImplV2.contract.UnpackLog(event, "OrderMatched", log); err != nil {
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

// ParseOrderMatched is a log parse operation binding the contract event 0x6a6896b1e6131e0b8ebae43fdc84cb0178a6eacdf4ee63f15dabae48729a8a03.
//
// Solidity: event OrderMatched(uint256 indexed sellId, uint256 indexed buyId, uint256 indexed price, uint256 amount, uint256 timestamp)
func (_PairImplV2 *PairImplV2Filterer) ParseOrderMatched(log types.Log) (*PairImplV2OrderMatched, error) {
	event := new(PairImplV2OrderMatched)
	if err := _PairImplV2.contract.UnpackLog(event, "OrderMatched", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PairImplV2PairFeesUpdatedIterator is returned from FilterPairFeesUpdated and is used to iterate over the raw logs and unpacked data for PairFeesUpdated events raised by the PairImplV2 contract.
type PairImplV2PairFeesUpdatedIterator struct {
	Event *PairImplV2PairFeesUpdated // Event containing the contract specifics and raw log

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
func (it *PairImplV2PairFeesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PairImplV2PairFeesUpdated)
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
		it.Event = new(PairImplV2PairFeesUpdated)
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
func (it *PairImplV2PairFeesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PairImplV2PairFeesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PairImplV2PairFeesUpdated represents a PairFeesUpdated event raised by the PairImplV2 contract.
type PairImplV2PairFeesUpdated struct {
	SellerMakerFee uint32
	SellerTakerFee uint32
	BuyerMakerFee  uint32
	BuyerTakerFee  uint32
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPairFeesUpdated is a free log retrieval operation binding the contract event 0x86f539dd22f42128e02dc316d053624f3c64c6a38938170bc4bd73e700f1d436.
//
// Solidity: event PairFeesUpdated(uint32 sellerMakerFee, uint32 sellerTakerFee, uint32 buyerMakerFee, uint32 buyerTakerFee)
func (_PairImplV2 *PairImplV2Filterer) FilterPairFeesUpdated(opts *bind.FilterOpts) (*PairImplV2PairFeesUpdatedIterator, error) {

	logs, sub, err := _PairImplV2.contract.FilterLogs(opts, "PairFeesUpdated")
	if err != nil {
		return nil, err
	}
	return &PairImplV2PairFeesUpdatedIterator{contract: _PairImplV2.contract, event: "PairFeesUpdated", logs: logs, sub: sub}, nil
}

// WatchPairFeesUpdated is a free log subscription operation binding the contract event 0x86f539dd22f42128e02dc316d053624f3c64c6a38938170bc4bd73e700f1d436.
//
// Solidity: event PairFeesUpdated(uint32 sellerMakerFee, uint32 sellerTakerFee, uint32 buyerMakerFee, uint32 buyerTakerFee)
func (_PairImplV2 *PairImplV2Filterer) WatchPairFeesUpdated(opts *bind.WatchOpts, sink chan<- *PairImplV2PairFeesUpdated) (event.Subscription, error) {

	logs, sub, err := _PairImplV2.contract.WatchLogs(opts, "PairFeesUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PairImplV2PairFeesUpdated)
				if err := _PairImplV2.contract.UnpackLog(event, "PairFeesUpdated", log); err != nil {
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

// ParsePairFeesUpdated is a log parse operation binding the contract event 0x86f539dd22f42128e02dc316d053624f3c64c6a38938170bc4bd73e700f1d436.
//
// Solidity: event PairFeesUpdated(uint32 sellerMakerFee, uint32 sellerTakerFee, uint32 buyerMakerFee, uint32 buyerTakerFee)
func (_PairImplV2 *PairImplV2Filterer) ParsePairFeesUpdated(log types.Log) (*PairImplV2PairFeesUpdated, error) {
	event := new(PairImplV2PairFeesUpdated)
	if err := _PairImplV2.contract.UnpackLog(event, "PairFeesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PairImplV2PausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the PairImplV2 contract.
type PairImplV2PausedIterator struct {
	Event *PairImplV2Paused // Event containing the contract specifics and raw log

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
func (it *PairImplV2PausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PairImplV2Paused)
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
		it.Event = new(PairImplV2Paused)
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
func (it *PairImplV2PausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PairImplV2PausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PairImplV2Paused represents a Paused event raised by the PairImplV2 contract.
type PairImplV2Paused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PairImplV2 *PairImplV2Filterer) FilterPaused(opts *bind.FilterOpts) (*PairImplV2PausedIterator, error) {

	logs, sub, err := _PairImplV2.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PairImplV2PausedIterator{contract: _PairImplV2.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PairImplV2 *PairImplV2Filterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PairImplV2Paused) (event.Subscription, error) {

	logs, sub, err := _PairImplV2.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PairImplV2Paused)
				if err := _PairImplV2.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PairImplV2 *PairImplV2Filterer) ParsePaused(log types.Log) (*PairImplV2Paused, error) {
	event := new(PairImplV2Paused)
	if err := _PairImplV2.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PairImplV2SkimIterator is returned from FilterSkim and is used to iterate over the raw logs and unpacked data for Skim events raised by the PairImplV2 contract.
type PairImplV2SkimIterator struct {
	Event *PairImplV2Skim // Event containing the contract specifics and raw log

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
func (it *PairImplV2SkimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PairImplV2Skim)
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
		it.Event = new(PairImplV2Skim)
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
func (it *PairImplV2SkimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PairImplV2SkimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PairImplV2Skim represents a Skim event raised by the PairImplV2 contract.
type PairImplV2Skim struct {
	Caller common.Address
	Erc20  common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSkim is a free log retrieval operation binding the contract event 0x7e30e394efdaabad9f9ca6fd8f67f9449ba63d674ee8068283c3ae1f49b5b660.
//
// Solidity: event Skim(address indexed caller, address indexed erc20, address indexed to, uint256 amount)
func (_PairImplV2 *PairImplV2Filterer) FilterSkim(opts *bind.FilterOpts, caller []common.Address, erc20 []common.Address, to []common.Address) (*PairImplV2SkimIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var erc20Rule []interface{}
	for _, erc20Item := range erc20 {
		erc20Rule = append(erc20Rule, erc20Item)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PairImplV2.contract.FilterLogs(opts, "Skim", callerRule, erc20Rule, toRule)
	if err != nil {
		return nil, err
	}
	return &PairImplV2SkimIterator{contract: _PairImplV2.contract, event: "Skim", logs: logs, sub: sub}, nil
}

// WatchSkim is a free log subscription operation binding the contract event 0x7e30e394efdaabad9f9ca6fd8f67f9449ba63d674ee8068283c3ae1f49b5b660.
//
// Solidity: event Skim(address indexed caller, address indexed erc20, address indexed to, uint256 amount)
func (_PairImplV2 *PairImplV2Filterer) WatchSkim(opts *bind.WatchOpts, sink chan<- *PairImplV2Skim, caller []common.Address, erc20 []common.Address, to []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var erc20Rule []interface{}
	for _, erc20Item := range erc20 {
		erc20Rule = append(erc20Rule, erc20Item)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PairImplV2.contract.WatchLogs(opts, "Skim", callerRule, erc20Rule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PairImplV2Skim)
				if err := _PairImplV2.contract.UnpackLog(event, "Skim", log); err != nil {
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

// ParseSkim is a log parse operation binding the contract event 0x7e30e394efdaabad9f9ca6fd8f67f9449ba63d674ee8068283c3ae1f49b5b660.
//
// Solidity: event Skim(address indexed caller, address indexed erc20, address indexed to, uint256 amount)
func (_PairImplV2 *PairImplV2Filterer) ParseSkim(log types.Log) (*PairImplV2Skim, error) {
	event := new(PairImplV2Skim)
	if err := _PairImplV2.contract.UnpackLog(event, "Skim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PairImplV2TickSizeUpdatedIterator is returned from FilterTickSizeUpdated and is used to iterate over the raw logs and unpacked data for TickSizeUpdated events raised by the PairImplV2 contract.
type PairImplV2TickSizeUpdatedIterator struct {
	Event *PairImplV2TickSizeUpdated // Event containing the contract specifics and raw log

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
func (it *PairImplV2TickSizeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PairImplV2TickSizeUpdated)
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
		it.Event = new(PairImplV2TickSizeUpdated)
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
func (it *PairImplV2TickSizeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PairImplV2TickSizeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PairImplV2TickSizeUpdated represents a TickSizeUpdated event raised by the PairImplV2 contract.
type PairImplV2TickSizeUpdated struct {
	BeforeLotSize  *big.Int
	NewLotSize     *big.Int
	BeforeTickSize *big.Int
	NewTickSize    *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTickSizeUpdated is a free log retrieval operation binding the contract event 0x66748bc112f6692a76bebecc33e69768508e27e414e3080f5c02295b6ea3cf3f.
//
// Solidity: event TickSizeUpdated(uint256 beforeLotSize, uint256 newLotSize, uint256 beforeTickSize, uint256 newTickSize)
func (_PairImplV2 *PairImplV2Filterer) FilterTickSizeUpdated(opts *bind.FilterOpts) (*PairImplV2TickSizeUpdatedIterator, error) {

	logs, sub, err := _PairImplV2.contract.FilterLogs(opts, "TickSizeUpdated")
	if err != nil {
		return nil, err
	}
	return &PairImplV2TickSizeUpdatedIterator{contract: _PairImplV2.contract, event: "TickSizeUpdated", logs: logs, sub: sub}, nil
}

// WatchTickSizeUpdated is a free log subscription operation binding the contract event 0x66748bc112f6692a76bebecc33e69768508e27e414e3080f5c02295b6ea3cf3f.
//
// Solidity: event TickSizeUpdated(uint256 beforeLotSize, uint256 newLotSize, uint256 beforeTickSize, uint256 newTickSize)
func (_PairImplV2 *PairImplV2Filterer) WatchTickSizeUpdated(opts *bind.WatchOpts, sink chan<- *PairImplV2TickSizeUpdated) (event.Subscription, error) {

	logs, sub, err := _PairImplV2.contract.WatchLogs(opts, "TickSizeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PairImplV2TickSizeUpdated)
				if err := _PairImplV2.contract.UnpackLog(event, "TickSizeUpdated", log); err != nil {
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

// ParseTickSizeUpdated is a log parse operation binding the contract event 0x66748bc112f6692a76bebecc33e69768508e27e414e3080f5c02295b6ea3cf3f.
//
// Solidity: event TickSizeUpdated(uint256 beforeLotSize, uint256 newLotSize, uint256 beforeTickSize, uint256 newTickSize)
func (_PairImplV2 *PairImplV2Filterer) ParseTickSizeUpdated(log types.Log) (*PairImplV2TickSizeUpdated, error) {
	event := new(PairImplV2TickSizeUpdated)
	if err := _PairImplV2.contract.UnpackLog(event, "TickSizeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PairImplV2UnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the PairImplV2 contract.
type PairImplV2UnpausedIterator struct {
	Event *PairImplV2Unpaused // Event containing the contract specifics and raw log

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
func (it *PairImplV2UnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PairImplV2Unpaused)
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
		it.Event = new(PairImplV2Unpaused)
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
func (it *PairImplV2UnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PairImplV2UnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PairImplV2Unpaused represents a Unpaused event raised by the PairImplV2 contract.
type PairImplV2Unpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PairImplV2 *PairImplV2Filterer) FilterUnpaused(opts *bind.FilterOpts) (*PairImplV2UnpausedIterator, error) {

	logs, sub, err := _PairImplV2.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PairImplV2UnpausedIterator{contract: _PairImplV2.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PairImplV2 *PairImplV2Filterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PairImplV2Unpaused) (event.Subscription, error) {

	logs, sub, err := _PairImplV2.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PairImplV2Unpaused)
				if err := _PairImplV2.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PairImplV2 *PairImplV2Filterer) ParseUnpaused(log types.Log) (*PairImplV2Unpaused, error) {
	event := new(PairImplV2Unpaused)
	if err := _PairImplV2.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PairImplV2UpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the PairImplV2 contract.
type PairImplV2UpgradedIterator struct {
	Event *PairImplV2Upgraded // Event containing the contract specifics and raw log

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
func (it *PairImplV2UpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PairImplV2Upgraded)
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
		it.Event = new(PairImplV2Upgraded)
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
func (it *PairImplV2UpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PairImplV2UpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PairImplV2Upgraded represents a Upgraded event raised by the PairImplV2 contract.
type PairImplV2Upgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_PairImplV2 *PairImplV2Filterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*PairImplV2UpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _PairImplV2.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &PairImplV2UpgradedIterator{contract: _PairImplV2.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_PairImplV2 *PairImplV2Filterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *PairImplV2Upgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _PairImplV2.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PairImplV2Upgraded)
				if err := _PairImplV2.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_PairImplV2 *PairImplV2Filterer) ParseUpgraded(log types.Log) (*PairImplV2Upgraded, error) {
	event := new(PairImplV2Upgraded)
	if err := _PairImplV2.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
