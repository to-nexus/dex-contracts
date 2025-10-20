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
	Bin: "0x60a060405230608052348015610013575f5ffd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b608051615df86100f95f395f81816136460152818161366f01526138c30152615df85ff3fe608060405260043610610229575f3560e01c80637cef243c11610131578063c3f909d4116100ac578063e077dd8a1161007c578063f210d08711610062578063f210d08714610765578063f46f16c21461077a578063fe566349146107a5575f5ffd5b8063e077dd8a1461071a578063ec342ad014610739575f5ffd5b8063c3f909d414610651578063cc50e36a146106cf578063d6e9cc2a146106ee578063dfdf2a7214610705575f5ffd5b8063918f8674116101015780639da771f4116100e75780639da771f4146105c8578063ad3cb1cc146105dd578063bedb86fb14610632575f5ffd5b8063918f8674146105875780639c5798391461059c575f5ffd5b80637cef243c14610514578063803a0386146105285780638d77eba5146105475780638da5cb5b14610573575f5ffd5b806332fe7b26116101c157806352d1902d116101915780636c4cfbc8116101775780636c4cfbc8146104825780637aae3523146104a15780637cbf6db2146104ff575f5ffd5b806352d1902d1461042d5780635c975abb14610441575f5ffd5b806332fe7b261461039f5780634942f65f146103f05780634d6754f1146104055780634f1ef2861461041a575f5ffd5b80631ec482d7116101fc5780631ec482d71461031c5780631edd86141461033b5780632cfffaf61461035e5780632dd6a00a14610380575f5ffd5b806302a204c61461022d57806312a808cc1461024e5780631a0361fb146102835780631e5eb1d0146102a2575b5f5ffd5b348015610238575f5ffd5b5061024c61024736600461533e565b6107c4565b005b348015610259575f5ffd5b5061026d6102683660046153d3565b610cec565b60405161027a9190615492565b60405180910390f35b34801561028e575f5ffd5b5061024c61029d36600461553c565b610dd2565b3480156102ad575f5ffd5b506019546102ec9063ffffffff8082169164010000000081048216916801000000000000000082048116916c0100000000000000000000000090041684565b6040805163ffffffff9586168152938516602085015291841691830191909152909116606082015260800161027a565b348015610327575f5ffd5b5061024c6103363660046155bb565b61113a565b348015610346575f5ffd5b5061035060075481565b60405190815260200161027a565b348015610369575f5ffd5b5061037261134a565b60405161027a929190615646565b34801561038b575f5ffd5b5061035061039a366004615700565b611384565b3480156103aa575f5ffd5b506001546103cb9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161027a565b3480156103fb575f5ffd5b50610350600a5481565b348015610410575f5ffd5b5061035060085481565b61024c610428366004615746565b611a73565b348015610438575f5ffd5b50610350611a92565b34801561044c575f5ffd5b507fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff16604051901515815260200161027a565b34801561048d575f5ffd5b5061035061049c366004615793565b611ac0565b3480156104ac575f5ffd5b506104ea6104bb3660046157db565b73ffffffffffffffffffffffffffffffffffffffff165f90815260186020526040902080546001909101549091565b6040805192835260208301919091520161027a565b34801561050a575f5ffd5b50610350600b5481565b34801561051f575f5ffd5b506102ec611ba2565b348015610533575f5ffd5b5061024c6105423660046157f6565b611bd0565b348015610552575f5ffd5b5061056661056136600461584f565b611c22565b60405161027a91906158cc565b34801561057e575f5ffd5b506103cb611cf1565b348015610592575f5ffd5b5061035060045481565b3480156105a7575f5ffd5b506003546103cb9073ffffffffffffffffffffffffffffffffffffffff1681565b3480156105d3575f5ffd5b5061035060065481565b3480156105e8575f5ffd5b506106256040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b60405161027a9190615928565b34801561063d575f5ffd5b5061024c61064c36600461597b565b611d84565b34801561065c575f5ffd5b5060408051606080820183525f8083526020808401829052928401528251808201845260035473ffffffffffffffffffffffffffffffffffffffff90811680835260025482168386019081526004549387019384528651918252519091169381019390935251928201929092520161027a565b3480156106da575f5ffd5b5061024c6106e936600461599a565b611ddd565b3480156106f9575f5ffd5b50600954600a546104ea565b348015610710575f5ffd5b5061035060055481565b348015610725575f5ffd5b5061024c6107343660046159cb565b611fb3565b348015610744575f5ffd5b506002546103cb9073ffffffffffffffffffffffffffffffffffffffff1681565b348015610770575f5ffd5b5061035060095481565b348015610785575f5ffd5b505f546103cb9073ffffffffffffffffffffffffffffffffffffffff1681565b3480156107b0575f5ffd5b5061024c6107bf366004615a0a565b6120fa565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff165f8115801561080e5750825b90505f8267ffffffffffffffff16600114801561082a5750303b155b905081158015610838575080155b1561086f576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000016600117855583156108d05784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b6108d8612312565b73ffffffffffffffffffffffffffffffffffffffff8b1661094c576040517fc587916e0000000000000000000000000000000000000000000000000000000081527f726f75746572000000000000000000000000000000000000000000000000000060048201526024015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8a166109bb576040517fc587916e0000000000000000000000000000000000000000000000000000000081527f71756f74650000000000000000000000000000000000000000000000000000006004820152602401610943565b73ffffffffffffffffffffffffffffffffffffffff8916610a2a576040517fc587916e0000000000000000000000000000000000000000000000000000000081527f62617365000000000000000000000000000000000000000000000000000000006004820152602401610943565b875f03610a85576040517fc587916e0000000000000000000000000000000000000000000000000000000081527f7469636b53697a650000000000000000000000000000000000000000000000006004820152602401610943565b865f03610ae0576040517fc587916e0000000000000000000000000000000000000000000000000000000081527f6c6f7453697a65000000000000000000000000000000000000000000000000006004820152602401610943565b335f80547fffffffffffffffffffffffff000000000000000000000000000000000000000090811673ffffffffffffffffffffffffffffffffffffffff938416179091556001805482168e84161790556003805482168d841617905560028054909116918b169182179055604080517f313ce567000000000000000000000000000000000000000000000000000000008152905163313ce567916004808201926020929091908290030181865afa158015610b9d573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610bc19190615a2a565b610bcc90600a615b52565b6004819055610bdb888a615b60565b610be59190615ba4565b15610c2e57600480546040517f575681350000000000000000000000000000000000000000000000000000000081529182018a9052602482018990526044820152606401610943565b6009889055600a879055600454610c489089908990612324565b600b819055505f5f5f5f89806020019051810190610c669190615bdc565b9350935093509350610c7a848484846123db565b505050508315610cdf5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050505050565b60605f80846001811115610d0257610d02615866565b14610d0e576016610d11565b60155b83519091505f8167ffffffffffffffff811115610d3057610d30615238565b604051908082528060200260200182016040528015610d6357816020015b6060815260200190600190039081610d4e5790505b5090505f5b82811015610dc657610da1845f888481518110610d8757610d87615c2d565b602002602001015181526020019081526020015f20612714565b828281518110610db357610db3615c2d565b6020908102919091010152600101610d68565b50925050505b92915050565b610dda611cf1565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610e5957335b6040517f118cdaa700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610943565b80156111355760025473ffffffffffffffffffffffffffffffffffffffff8481169116148015610f24575080600554610e929190615c5a565b6002546040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff909116906370a0823190602401602060405180830381865afa158015610efe573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f229190615c6d565b105b15610f77576002546040517f7bdf264d00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610943565b60035473ffffffffffffffffffffffffffffffffffffffff848116911614801561103c575080600654610faa9190615c5a565b6003546040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff909116906370a0823190602401602060405180830381865afa158015611016573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061103a9190615c6d565b105b1561108f576003546040517f7bdf264d00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610943565b6110b073ffffffffffffffffffffffffffffffffffffffff841683836127ba565b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff166110e53390565b73ffffffffffffffffffffffffffffffffffffffff167f7e30e394efdaabad9f9ca6fd8f67f9449ba63d674ee8068283c3ae1f49b5b6608460405161112c91815260200190565b60405180910390a45b505050565b60015473ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146111bc57335b6040517f143db19100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610943565b805f5b81811015611343575f8484838181106111da576111da615c2d565b602090810292909201355f81815260179093526040808420815160a08101909252805492955090925090829060ff16600181111561121a5761121a615866565b600181111561122b5761122b615866565b81528154610100810473ffffffffffffffffffffffffffffffffffffffff908116602080850191909152750100000000000000000000000000000000000000000090920463ffffffff16604084015260018401546060840152600290930154608090920191909152820151919250166112a55750506111bf565b8673ffffffffffffffffffffffffffffffffffffffff16816020015173ffffffffffffffffffffffffffffffffffffffff161461132d576040517f54de53f70000000000000000000000000000000000000000000000000000000081526004810183905273ffffffffffffffffffffffffffffffffffffffff88166024820152604401610943565b61133982826002612847565b50506001016111bf565b5050505050565b606080611370600d5f5b60ff166002811061136757611367615c2d565b60040201612714565b915061137e600d6001611354565b90509091565b5f61138d6129d0565b60015473ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146113c85733611171565b6113d0612a2c565b606085015115806113f1575060095485606001516113ee9190615ba4565b15155b156114305784606001516040517f4e1404fe00000000000000000000000000000000000000000000000000000000815260040161094391815260200190565b608085015115806114515750600a54856080015161144e9190615ba4565b15155b156114905784608001516040517fa334626e00000000000000000000000000000000000000000000000000000000815260040161094391815260200190565b600c5f815461149e90615c84565b918290555090505f80865160018111156114ba576114ba615866565b1490505f5f826114d5576114d084895f88612b18565b6114e0565b6114e0848987612e8d565b91509150811561156e575f847f78615f01112d8477ae8d6fc9a510bcfdf935bc3ab7b02d0ff2316f5b44bd597a4260405161151d91815260200190565b60405180910390a38280156115355750608088015115155b1561156957602088015160808901516002546115699273ffffffffffffffffffffffffffffffffffffffff909116916127ba565b611a55565b600187600281111561158257611582615866565b036115f8576001847f78615f01112d8477ae8d6fc9a510bcfdf935bc3ab7b02d0ff2316f5b44bd597a426040516115bb91815260200190565b60405180910390a3821561156957602088015160808901516002546115699273ffffffffffffffffffffffffffffffffffffffff909116916127ba565b600287600281111561160c5761160c615866565b036116615760208801516040517f1b3c356200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610943565b8515801590611692575082801561167b5750858860600151105b806116925750821580156116925750858860600151115b156116d457875160608901516040517fc18aa6060000000000000000000000000000000000000000000000000000000081526109439291908990600401615c9c565b611712886060015187600d8b5f015160018111156116f4576116f4615866565b60ff166002811061170757611707615c2d565b600402019190613133565b505f848152601760205260409020885181548a92919082907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001838181111561175f5761175f615866565b021790555060208201518154604084015163ffffffff167501000000000000000000000000000000000000000000027fffffffffffffff00000000ffffffffffffffffffffffffffffffffffffffffff73ffffffffffffffffffffffffffffffffffffffff90931661010002929092167fffffffffffffff000000000000000000000000000000000000000000000000ff909116171781556060820151600182015560809091015160029091015582156118de57611825886020015189608001516132e1565b604080516080810182525f80825260208201819052918101829052606001527f1d2ff3fa6980aeeebca4e94965520da48983006e9b1115c1c853cbb10d943d005c515f85815260176020908152604080832080547fffffffffffffff00000000ffffffffffffffffffffffffffffffffffffffffff16750100000000000000000000000000000000000000000063ffffffff9687160217905560608c01518352601590915290206118d891869061341316565b50611a55565b5f61192a604080516080810182525f808252602082018190528183018190526060909101527f1d2ff3fa6980aeeebca4e94965520da48983006e9b1115c1c853cbb10d943d005c015190565b6003546040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201529192505f91849173ffffffffffffffffffffffffffffffffffffffff16906370a0823190602401602060405180830381865afa15801561199b573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906119bf9190615c6d565b6119c99190615cbb565b90506119d98a6020015182613423565b5f868152601760209081526040808320805463ffffffff8088167501000000000000000000000000000000000000000000027fffffffffffffff00000000ffffffffffffffffffffffffffffffffffffffffff9092169190911790915560608e015184526016909252909120611a5191889061341316565b5050505b82611a6857611a68886020015182613551565b505050949350505050565b611a7b61362e565b611a8482613732565b611a8e8282613772565b5050565b5f611a9b6138ab565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b5f80856001811115611ad457611ad4615866565b03611b3d57604080518082018252611b3691869190869060029083908390808284375f92019190915250869150600d9050896001811115611b1757611b17615866565b60ff1660028110611b2a57611b2a615c2d565b6004020192919061391a565b9050611b9a565b604080518082018252611b3691869190869060029083908390808284375f92019190915250869150600d9050896001811115611b7b57611b7b615866565b60ff1660028110611b8e57611b8e615c2d565b60040201929190613ac0565b949350505050565b5f5f5f5f5f611baf613c2f565b80516020820151604083015160609093015191989097509195509350915050565b611bd8611cf1565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611c105733610e0e565b611c1c848484846123db565b50505050565b6040805160a0810182525f808252602082018190529181018290526060810182905260808101919091525f8281526017602052604090819020815160a081019092528054829060ff166001811115611c7c57611c7c615866565b6001811115611c8d57611c8d615866565b81528154610100810473ffffffffffffffffffffffffffffffffffffffff1660208301527501000000000000000000000000000000000000000000900463ffffffff1660408201526001820154606082015260029091015460809091015292915050565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611d5b573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611d7f9190615cce565b905090565b611d8c611cf1565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611dc45733610e0e565b8015611dd557611dd2613dd9565b50565b611dd2613e79565b611de56129d0565b60015473ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611e205733611171565b611e28612a2c565b5f600c5f8154611e3790615c84565b918290555090505f84516001811115611e5257611e52615866565b03611f0157821580611e6f5750600a54611e6c9084615ba4565b15155b15611ea9576040517fa334626e00000000000000000000000000000000000000000000000000000000815260048101849052602401610943565b5f606085015260808401839052611ec1818584612e8d565b5050608084015115611efc5760208401516080850151600254611efc9273ffffffffffffffffffffffffffffffffffffffff909116916127ba565b611f72565b600b54831015611f4b57600b546040517f0bc1e7dd000000000000000000000000000000000000000000000000000000008152610943918591600401918252602082015260400190565b5f1960608501525f611f5f82868686612b18565b915050611f70856020015182613551565b505b5f817f78615f01112d8477ae8d6fc9a510bcfdf935bc3ab7b02d0ff2316f5b44bd597a42604051611fa591815260200190565b60405180910390a350505050565b611fbb613eef565b611fc3611cf1565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611ffb5733610e0e565b805f5b81811015611c1c575f84848381811061201957612019615c2d565b602090810292909201355f81815260179093526040808420815160a08101909252805492955090925090829060ff16600181111561205957612059615866565b600181111561206a5761206a615866565b81528154610100810473ffffffffffffffffffffffffffffffffffffffff908116602080850191909152750100000000000000000000000000000000000000000090920463ffffffff16604084015260018401546060840152600290930154608090920191909152820151919250166120e4575050611ffe565b6120f082826003612847565b5050600101611ffe565b5f5473ffffffffffffffffffffffffffffffffffffffff1663a1eea778336040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815273ffffffffffffffffffffffffffffffffffffffff90911660048201526024015f6040518083038186803b15801561217a575f5ffd5b505afa15801561218c573d5f5f3e3d5ffd5b50505050805f036121eb576040517fc587916e0000000000000000000000000000000000000000000000000000000081527f7469636b53697a650000000000000000000000000000000000000000000000006004820152602401610943565b815f03612246576040517fc587916e0000000000000000000000000000000000000000000000000000000081527f6c6f7453697a65000000000000000000000000000000000000000000000000006004820152602401610943565b6004546122538383615b60565b61225d9190615ba4565b156122a657600480546040517f57568135000000000000000000000000000000000000000000000000000000008152918201839052602482018490526044820152606401610943565b600a546009546040805192835260208301859052820152606081018290527f66748bc112f6692a76bebecc33e69768508e27e414e3080f5c02295b6ea3cf3f9060800160405180910390a1600a829055600981905560045461230b9082908490612324565b600b555050565b61231a613f4a565b612322613fb1565b565b5f838302815f1985870982811083820303915050805f036123585783828161234e5761234e615b77565b04925050506123d4565b80841161236f5761236f6003851502601118614002565b5f848688095f868103871696879004966002600389028118808a02820302808a02820302808a02820302808a02820302808a02820302808a02909103029181900381900460010186841190950394909402919094039290920491909117919091029150505b9392505050565b61271063ffffffff8516108015906123f9575063ffffffff84811614155b15612430576040517fac408a2600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61271063ffffffff84161080159061244e575063ffffffff83811614155b15612485576040517fac408a2600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61271063ffffffff8316108015906124a3575063ffffffff82811614155b156124da576040517fac408a2600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61271063ffffffff8216108015906124f8575063ffffffff81811614155b1561252f576040517fac408a2600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8363ffffffff168363ffffffff16108015612550575063ffffffff84811614155b15612597576040517f3018805400000000000000000000000000000000000000000000000000000000815263ffffffff808616600483015284166024820152604401610943565b8163ffffffff168163ffffffff161080156125b8575063ffffffff82811614155b156125ff576040517f3018805400000000000000000000000000000000000000000000000000000000815263ffffffff808416600483015282166024820152604401610943565b604080516080808201835263ffffffff87811680845287821660208086018290528884168688018190529388166060968701819052601980547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001685176401000000008502177fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff166801000000000000000087027fffffffffffffffffffffffffffffffff00000000ffffffffffffffffffffffff16176c01000000000000000000000000830217905587519384529083019190915294810191909152918201929092527f86f539dd22f42128e02dc316d053624f3c64c6a38938170bc4bd73e700f1d436910160405180910390a150505050565b60605f825f015467ffffffffffffffff81111561273357612733615238565b60405190808252806020026020018201604052801561275c578160200160208202803683370190505b5083546001850154919250905f5b828110156127b0578184828151811061278557612785615c2d565b6020908102919091018101919091525f9283526003870190526040909120600190810154910161276a565b5091949350505050565b6040805173ffffffffffffffffffffffffffffffffffffffff8416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb00000000000000000000000000000000000000000000000000000000179052611135908490614013565b5f80808451600181111561285d5761285d615866565b14905080156128cb5760608401515f908152601560205260409020608085015190925080156128c55760208501516002546128b19173ffffffffffffffffffffffffffffffffffffffff90911690836127ba565b6128c38560200151825f6005546140b2565b505b5061297b565b60165f856060015181526020019081526020015f2091505f6128f885606001518660800151600454612324565b9050801561297957604085015163ffffffff161561293a5761292d81866040015163ffffffff1661271063ffffffff16612324565b6129379082615c5a565b90505b60208501516003546129659173ffffffffffffffffffffffffffffffffffffffff90911690836127ba565b6129778560200151825f6006546141ec565b505b505b612986858484614325565b8154611343576129c88460600151600d865f015160018111156129ab576129ab615866565b60ff16600281106129be576129be615c2d565b60040201906143f7565b505050505050565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff1615612322576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8054604080517fc415b95c000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff9092169291839163c415b95c9160048083019260209291908290030181865afa158015612a9c573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612ac09190615cce565b90505f612acb613c2f565b9050817fd6aa07baf8485abf9d26fecf4c935d75b50a73e678db02b944bd3ac8759823005d807f1d2ff3fa6980aeeebca4e94965520da48983006e9b1115c1c853cbb10d943d005d505050565b5f80600185516001811115612b2f57612b2f615866565b14612b68575f6040517f7cbdd0810000000000000000000000000000000000000000000000000000000081526004016109439190615ce9565b5f8415612b76575083612b8e565b612b8b86606001518760800151600454612324565b90505b604080516080810182525f8082526020820181905291810182905260609081018290527f1d2ff3fa6980aeeebca4e94965520da48983006e9b1115c1c853cbb10d943d005c015181905f8163ffffffff165f03612beb575f612bfe565b612bfe8563ffffffff8416612710612324565b6003546040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152919250612cb39173ffffffffffffffffffffffffffffffffffffffff909116906370a0823190602401602060405180830381865afa158015612c71573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612c959190615c6d565b8287600654612ca49190615c5a565b612cae9190615c5a565b6144c8565b9450925082612d0a576003546040517f7bdf264d00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610943565b89516001811115612d1d57612d1d615866565b8b8b6020015173ffffffffffffffffffffffffffffffffffffffff167f7cde7559ef00f8c3d9244d3fde19e7500979d64f6f9c086a3836a4926abbe1658d606001518e6080015142604051612d85939291909283526020830191909152604082015260600190565b60405180910390a45050505f5f5f612d9f8b8b8b8b6144ec565b925092509250815f14612e7c575f612e2d8c8c602001518585612de07fd6aa07baf8485abf9d26fecf4c935d75b50a73e678db02b944bd3ac8759823005c90565b604080516080810182525f808252602082018190529181018290526060908101919091527f1d2ff3fa6980aeeebca4e94965520da48983006e9b1115c1c853cbb10d943d005c0151614816565b90508015612e7a57612e7a7fd6aa07baf8485abf9d26fecf4c935d75b50a73e678db02b944bd3ac8759823005c60035473ffffffffffffffffffffffffffffffffffffffff1690836127ba565b505b509099919850909650505050505050565b5f808084516001811115612ea357612ea3615866565b14612edd5760016040517f7cbdd0810000000000000000000000000000000000000000000000000000000081526004016109439190615ce9565b8360800151600554612eef9190615c5a565b6002546040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff909116906370a0823190602401602060405180830381865afa158015612f5b573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612f7f9190615c6d565b1015612fd3576002546040517f7bdf264d00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610943565b83516001811115612fe657612fe6615866565b85856020015173ffffffffffffffffffffffffffffffffffffffff167f7cde7559ef00f8c3d9244d3fde19e7500979d64f6f9c086a3836a4926abbe165876060015188608001514260405161304e939291909283526020830191909152604082015260600190565b60405180910390a45f5f6130638787876148e5565b91509150805f14613124575f807fd6aa07baf8485abf9d26fecf4c935d75b50a73e678db02b944bd3ac8759823005c6130de604080516080810182525f80825260208083018290529282018190526060909101527f1d2ff3fa6980aeeebca4e94965520da48983006e9b1115c1c853cbb10d943d005c015190565b915091505f6130f48a8a60200151868686614b8a565b90508015613120576003546131209073ffffffffffffffffffffffffffffffffffffffff1684836127ba565b5050505b5091505f90505b935093915050565b5f825f0361316d576040517f0da93b9700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6131778484614c84565b1561318357505f6123d4565b5f828152600385016020908152604080832081518083019092528382529181019290925290835f036131ef576001860180549086905580156131d2575f81815260038801602052604090208690555b60405180604001604052805f815260200182815250915050613291565b60408051808201909152825481526001830154602082015261321090614ccc565b8061321e5750856001015484145b613254576040517f182dfca500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5060018101546040805180820190915284815260208101829052908015613288575f81815260038801602052604090208690555b50600182018590555b80602001515f036132a457600286018590555b5f85815260038701602090815260408220835181559083015160019091015586548791906132d190615c84565b9091555060019695505050505050565b5f5f6132ef60055484614ce2565b915091508161334e576002546040517f340b575200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff80871660048301529091166024820152604401610943565b600581905573ffffffffffffffffffffffffffffffffffffffff84165f908152601860205260408120613384915b015484614ce2565b9092509050816133e4576002546040517f340b575200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff80871660048301529091166024820152604401610943565b73ffffffffffffffffffffffffffffffffffffffff84165f90815260186020526040812082915b015550505050565b5f6123d483838560020154613133565b5f5f61343160065484614ce2565b9150915081613490576003546040517f340b575200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff80871660048301529091166024820152604401610943565b600681905573ffffffffffffffffffffffffffffffffffffffff84165f9081526018602052604090206134c490600161337c565b909250905081613524576003546040517f340b575200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff80871660048301529091166024820152604401610943565b73ffffffffffffffffffffffffffffffffffffffff84165f9081526018602052604090208190600161340b565b6003546040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201525f9173ffffffffffffffffffffffffffffffffffffffff16906370a0823190602401602060405180830381865afa1580156135bd573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906135e19190615c6d565b90505f600654836135f29190615c5a565b905080821115611c1c575f6136078284615cbb565b6003549091506113439073ffffffffffffffffffffffffffffffffffffffff1686836127ba565b3073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614806136fb57507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166136e27f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614155b15612322576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61373a611cf1565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611dd25733610e0e565b8173ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156137f7575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682019092526137f491810190615c6d565b60015b613845576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610943565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc81146138a1576040517faa1d49a400000000000000000000000000000000000000000000000000000000815260048101829052602401610943565b6111358383614d07565b3073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614612322576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f835f03613954576040517f0da93b9700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61395e8585614c84565b1561397957505f838152600385016020526040902054611b9a565b846001015484108061398a57508454155b1561399657505f611b9a565b84600201548411156139ad57506002840154611b9a565b815f036139e6576040517fab7e639300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6139f18685614d69565b905084811015613a51575b8015801590613a0a57508215155b15613a4c575f90815260038601602052604090206001015484811115613a41575f9081526003860160205260409020549050611b9a565b5f19909201916139fc565b613a8e565b8015801590613a5f57508215155b15613a8e575f90815260038601602052604090205484811015613a83579050611b9a565b5f1990920191613a51565b6040517f80ce60d100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f835f03613afa576040517f0da93b9700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b613b048585614c84565b15613b1f57505f838152600385016020526040902054611b9a565b8460010154841180613b3057508454155b15613b3c57505f611b9a565b8460020154841015613b5357506002840154611b9a565b815f03613b8c576040517fab7e639300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f613b978685614d69565b905084811115613bf2575b8015801590613bb057508215155b15613a4c575f90815260038601602052604090206001015480851115613be7575f9081526003860160205260409020549050611b9a565b5f1990920191613ba2565b8015801590613c0057508215155b15613a8e575f90815260038601602052604090205480851015613c24579050611b9a565b5f1990920191613bf2565b60408051608080820183525f8083526020830181905282840181905260608301819052805484517f5fbbc0d200000000000000000000000000000000000000000000000000000000815294519394919373ffffffffffffffffffffffffffffffffffffffff90911692635fbbc0d292600480820193918290030181865afa158015613cbc573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613ce09190615cf7565b60195490915063ffffffff90811614613d015760195463ffffffff16613d04565b80515b63ffffffff90811683526019546401000000009004811614613d3657601954640100000000900463ffffffff16613d3c565b80602001515b63ffffffff9081166020840152601954680100000000000000009004811614613d795760195468010000000000000000900463ffffffff16613d7f565b80604001515b63ffffffff90811660408401526019546c010000000000000000000000009004811614613dc4576019546c01000000000000000000000000900463ffffffff16613dca565b80606001515b63ffffffff1660608301525090565b613de16129d0565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a150565b613e81613eef565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa33613e4e565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff16612322576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff16612322576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b613fb9613f4a565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b634e487b715f52806020526024601cfd5b5f5f60205f8451602086015f885af180614032576040513d5f823e3d81fd5b50505f513d91508115614049578060011415614063565b73ffffffffffffffffffffffffffffffffffffffff84163b155b15611c1c576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610943565b73ffffffffffffffffffffffffffffffffffffffff84165f908152601860205260408120819081906140e890825b0154876144c8565b9150915081614147576002546040517f340b575200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff808a1660048301529091166024820152604401610943565b73ffffffffffffffffffffffffffffffffffffffff87165f90815260186020526040902081905561417884876144c8565b93509150816141d7576002546040517f340b575200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff808a1660048301529091166024820152604401610943565b846141e25760058390555b5050949350505050565b73ffffffffffffffffffffffffffffffffffffffff84165f9081526018602052604081208190819061421f9060016140e0565b915091508161427e576003546040517f340b575200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff808a1660048301529091166024820152604401610943565b73ffffffffffffffffffffffffffffffffffffffff87165f9081526018602052604090208190600101556142b284876144c8565b9350915081614311576003546040517f340b575200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff808a1660048301529091166024820152604401610943565b846141e25760068390555050949350505050565b61432f81846143f7565b614368576040517ffcfdf90200000000000000000000000000000000000000000000000000000000815260048101849052602401610943565b5f83815260176020526040812080547fffffffffffffff0000000000000000000000000000000000000000000000000016815560018101829055600201558160038111156143b8576143b8615866565b837f78615f01112d8477ae8d6fc9a510bcfdf935bc3ab7b02d0ff2316f5b44bd597a426040516143ea91815260200190565b60405180910390a3505050565b5f815f03614431576040517f0da93b9700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61443b8383614c84565b61444657505f610dcc565b5f82815260038401602052604080822080548352818320600180830154808652939094208482019390935581548355928601549092919085900361448f57600180840154908701555b848660020154036144a257825460028701555b5f858152600387016020526040812081815560010181905586548791906132d190615d7c565b5f5f838311156144dc57505f9050806144e5565b50600190508183035b9250929050565b6040805160c0810182525f8082526020820181905291810182905260608101829052600554608082015260a0810182905281908190600d5b80541561478357614535815f614db3565b8083526060890151106147835781515f908152601560205260409020825161455c90614e21565b87156145bb575f61458084604001518a6145769190615cbb565b6004548651612324565b9050600a54816145909190615ba4565b61459a9082615cbb565b60808b0181905290505f8190036145b9575050600160a0830152614783565b505b805415614726575f6145cd8282614db3565b90505f60175f8381526020019081526020015f2090505f5f5f6145f78f8f87878c5f01518b614e47565b9250925092505f61460e895f015184600454612324565b905061463e8685837fd6aa07baf8485abf9d26fecf4c935d75b50a73e678db02b944bd3ac8759823005c86614b8a565b8960600181815161464f9190615c5a565b905250602089018051849190614666908390615c5a565b90525060408901805182919061467d908390615c5a565b91508181525050614695848460018c608001516140b2565b6080808b01919091528f015115806146b657506146b18d615d7c565b9c508c155b1561471b5786546147085788516146ce9089906143f7565b6147085788516040517f9f1cfdfe000000000000000000000000000000000000000000000000000000008152610943915f91600401615d91565b5050600160a08801525061472692505050565b5050505050506145bb565b8260a00151156147365750614783565b82516147439083906143f7565b61477d5782516040517f9f1cfdfe000000000000000000000000000000000000000000000000000000008152610943915f91600401615d91565b50614524565b6060820151156147d7576147d77fd6aa07baf8485abf9d26fecf4c935d75b50a73e678db02b944bd3ac8759823005c606084015160035473ffffffffffffffffffffffffffffffffffffffff1691906127ba565b6005548260800151146147ed5760808201516005555b8160a0015182602001518360400151945094509450505061480c614f6d565b9450945094915050565b6002545f9061483c9073ffffffffffffffffffffffffffffffffffffffff1687876127ba565b8163ffffffff165f0361485057505f6148db565b5f6148648563ffffffff8516612710612324565b905073ffffffffffffffffffffffffffffffffffffffff808516908816897f9814b2b2f43f9ea51ad7e8760fda094ffa3012bfeb10c76a9656389af18956038887866148b08184615cbb565b6040805194855263ffffffff909316602085015291830152606082015260800160405180910390a490505b9695505050505050565b6040805160a0810182525f80825260208201819052918101829052600654606082015260808101829052819060115b805415614b0857614925815f614db3565b808352606087015111614b085781515f908152601660205260409020825161494c90614e21565b805415614aaa575f61495e8282614db3565b5f81815260176020526040812086519293509181908190614987908e908e90889088908b614e47565b9250925092505f61499e895f015184600454612324565b90505f6149d0878686857fd6aa07baf8485abf9d26fecf4c935d75b50a73e678db02b944bd3ac8759823005c88614816565b9050808a6040018181516149e49190615c5a565b90525060208a0180518391906149fb908390615c5a565b905250614a1885614a0c8385615c5a565b60018d606001516141ec565b60608b015260808e01511580614a375750614a328d615d7c565b9c508c155b15614a9e578754614a8a578951614a4f908a906143f7565b614a8a5789516040517f9f1cfdfe00000000000000000000000000000000000000000000000000000000815261094391600191600401615d91565b50506001608089015250614aaa9350505050565b5050505050505061494c565b826080015115614aba5750614b08565b8251614ac79083906143f7565b614b025782516040517f9f1cfdfe00000000000000000000000000000000000000000000000000000000815261094391600191600401615d91565b50614914565b604082015115614b5c57614b5c7fd6aa07baf8485abf9d26fecf4c935d75b50a73e678db02b944bd3ac8759823005c604084015160035473ffffffffffffffffffffffffffffffffffffffff1691906127ba565b600654826060015114614b725760608201516006555b8160800151826020015193509350505061312b614f6d565b5f8163ffffffff165f03614bc357600354614bbc9073ffffffffffffffffffffffffffffffffffffffff1686866127ba565b505f614c7b565b5f614bd78563ffffffff8516612710612324565b90505f614be48287615cbb565b6040805188815263ffffffff871660208201529081018490526060810182905290915073ffffffffffffffffffffffffffffffffffffffff80871691908916908a907f9814b2b2f43f9ea51ad7e8760fda094ffa3012bfeb10c76a9656389af18956039060800160405180910390a4600354614c779073ffffffffffffffffffffffffffffffffffffffff1688836127ba565b5090505b95945050505050565b5f815f03614c9357505f610dcc565b82600101548214806123d457505f82815260038401602090815260409182902082518084019093528054835260010154908201526123d4905b80515f90151580610dcc57505060200151151590565b5f8083830184811015614cfb575f5f92509250506144e5565b60019590945092505050565b614d1082614fbb565b60405173ffffffffffffffffffffffffffffffffffffffff8316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115614d61576111358282615089565b611a8e6150ff565b5f614d7b8383835b6020020151614c84565b15614d8f57815f5b60200201519050610dcc565b614d9b83836001614d71565b15614da857816001614d83565b506001820154610dcc565b81545f908210614def576040517f39e60f7300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001830154825b8015614e19575f91825260038501602052604090912060010154905f1901614df6565b509392505050565b807ffd0e5d4f9b88892d3b04349a0e2bc0d1359414c21932fcd7d5a523a6c0a5cd005d50565b5f5f5f855f0160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16614e8089608001518860020154615137565b875491945092507501000000000000000000000000000000000000000000900463ffffffff1690505f80808a516001811115614ebe57614ebe615866565b14614eca57888b614ecd565b8a895b915091508681837f6a6896b1e6131e0b8ebae43fdc84cb0178a6eacdf4ee63f15dabae48729a8a038742604051614f0e929190918252602082015260400190565b60405180910390a487600201548403614f3157614f2c895f88614325565b614f3d565b60028801805485900390555b89608001518403614f53575f60808b0152614f5f565b60808a01805185900390525b505096509650969350505050565b7ffd0e5d4f9b88892d3b04349a0e2bc0d1359414c21932fcd7d5a523a6c0a5cd005c8015801590614fa057506007548114155b15614fab5760078190555b4260085414611dd2574260085550565b8073ffffffffffffffffffffffffffffffffffffffff163b5f03615023576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152602401610943565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60605f5f8473ffffffffffffffffffffffffffffffffffffffff16846040516150b29190615dac565b5f60405180830381855af49150503d805f81146150ea576040519150601f19603f3d011682016040523d82523d5f602084013e6150ef565b606091505b5091509150614c7b858383615146565b3415612322576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8282188284100282186123d4565b60608261515b57615156826151d5565b6123d4565b815115801561517f575073ffffffffffffffffffffffffffffffffffffffff84163b155b156151ce576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610943565b50806123d4565b8051156151e55780518082602001fd5b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff81168114611dd2575f5ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156152ac576152ac615238565b604052919050565b5f82601f8301126152c3575f5ffd5b813567ffffffffffffffff8111156152dd576152dd615238565b61530e60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601615265565b818152846020838601011115615322575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f5f5f5f5f5f60c08789031215615353575f5ffd5b863561535e81615217565b9550602087013561536e81615217565b9450604087013561537e81615217565b9350606087013592506080870135915060a087013567ffffffffffffffff8111156153a7575f5ffd5b6153b389828a016152b4565b9150509295509295509295565b8035600281106153ce575f5ffd5b919050565b5f5f604083850312156153e4575f5ffd5b6153ed836153c0565b9150602083013567ffffffffffffffff811115615408575f5ffd5b8301601f81018513615418575f5ffd5b803567ffffffffffffffff81111561543257615432615238565b8060051b61544260208201615265565b9182526020818401810192908101908884111561545d575f5ffd5b6020850194505b8385101561548357843580835260209586019590935090910190615464565b80955050505050509250929050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b82811015615530578685037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0018452815180518087526020918201918701905f5b818110156155175783518352602093840193909201916001016154f9565b50909650505060209384019391909101906001016154b8565b50929695505050505050565b5f5f5f6060848603121561554e575f5ffd5b833561555981615217565b9250602084013561556981615217565b929592945050506040919091013590565b5f5f83601f84011261558a575f5ffd5b50813567ffffffffffffffff8111156155a1575f5ffd5b6020830191508360208260051b85010111156144e5575f5ffd5b5f5f5f604084860312156155cd575f5ffd5b83356155d881615217565b9250602084013567ffffffffffffffff8111156155f3575f5ffd5b6155ff8682870161557a565b9497909650939450505050565b5f8151808452602084019350602083015f5b8281101561563c57815186526020958601959091019060010161561e565b5093949350505050565b604081525f615658604083018561560c565b8281036020840152614c7b818561560c565b63ffffffff81168114611dd2575f5ffd5b5f60a0828403121561568b575f5ffd5b60405160a0810167ffffffffffffffff811182821017156156ae576156ae615238565b6040529050806156bd836153c0565b815260208301356156cd81615217565b602082015260408301356156e08161566a565b604082015260608381013590820152608092830135920191909152919050565b5f5f5f5f6101008587031215615714575f5ffd5b61571e868661567b565b935060a085013560038110615731575f5ffd5b939693955050505060c08201359160e0013590565b5f5f60408385031215615757575f5ffd5b823561576281615217565b9150602083013567ffffffffffffffff81111561577d575f5ffd5b615789858286016152b4565b9150509250929050565b5f5f5f5f60a085870312156157a6575f5ffd5b6157af856153c0565b93506020850135925060808501868111156157c8575f5ffd5b9396929550505060409290920191903590565b5f602082840312156157eb575f5ffd5b81356123d481615217565b5f5f5f5f60808587031215615809575f5ffd5b84356158148161566a565b935060208501356158248161566a565b925060408501356158348161566a565b915060608501356158448161566a565b939692955090935050565b5f6020828403121561585f575f5ffd5b5035919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b600281106158c8577f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b9052565b5f60a0820190506158de828451615893565b73ffffffffffffffffffffffffffffffffffffffff602084015116602083015263ffffffff6040840151166040830152606083015160608301526080830151608083015292915050565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b5f6020828403121561598b575f5ffd5b813580151581146123d4575f5ffd5b5f5f5f60e084860312156159ac575f5ffd5b6159b6858561567b565b9560a0850135955060c0909401359392505050565b5f5f602083850312156159dc575f5ffd5b823567ffffffffffffffff8111156159f2575f5ffd5b6159fe8582860161557a565b90969095509350505050565b5f5f60408385031215615a1b575f5ffd5b50508035926020909101359150565b5f60208284031215615a3a575f5ffd5b815160ff811681146123d4575f5ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b6001815b600184111561312b57808504811115615a9657615a96615a4a565b6001841615615aa457908102905b60019390931c928002615a7b565b5f82615ac057506001610dcc565b81615acc57505f610dcc565b8160018114615ae25760028114615aec57615b08565b6001915050610dcc565b60ff841115615afd57615afd615a4a565b50506001821b610dcc565b5060208310610133831016604e8410600b8410161715615b2b575081810a610dcc565b615b375f198484615a77565b805f1904821115615b4a57615b4a615a4a565b029392505050565b5f6123d460ff841683615ab2565b8082028115828204841417610dcc57610dcc615a4a565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f82615bd7577f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b500690565b5f5f5f5f60808587031215615bef575f5ffd5b8451615bfa8161566a565b6020860151909450615c0b8161566a565b6040860151909350615c1c8161566a565b60608601519092506158448161566a565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b80820180821115610dcc57610dcc615a4a565b5f60208284031215615c7d575f5ffd5b5051919050565b5f5f198203615c9557615c95615a4a565b5060010190565b60608101615caa8286615893565b602082019390935260400152919050565b81810381811115610dcc57610dcc615a4a565b5f60208284031215615cde575f5ffd5b81516123d481615217565b60208101610dcc8284615893565b5f6080828403128015615d08575f5ffd5b506040516080810167ffffffffffffffff81118282101715615d2c57615d2c615238565b6040528251615d3a8161566a565b81526020830151615d4a8161566a565b60208201526040830151615d5d8161566a565b60408201526060830151615d708161566a565b60608201529392505050565b5f81615d8a57615d8a615a4a565b505f190190565b60408101615d9f8285615893565b8260208301529392505050565b5f82518060208501845e5f92019182525091905056fea264697066735822122009976cd5664760255d7e6b19c6d3b9dd6dad6928e808c49409053b7ea95612b064736f6c634300081c0033",
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