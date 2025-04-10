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
type IPairConfig struct {
	QUOTE       common.Address
	BASE        common.Address
	DENOMINATOR *big.Int
}

// IPairOrder is an auto generated low-level Go binding around an user-defined struct.
type IPairOrder struct {
	Side   uint8
	Owner  common.Address
	FeeBps uint32
	Price  *big.Int
	Amount *big.Int
}

// PairImplMetaData contains all meta data concerning the PairImpl contract.
var PairImplMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BASE\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DENOMINATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MARKET\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"QUOTE\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROUTER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"accountReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"base\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quote\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseReserve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"orderIds\",\"type\":\"uint256[]\"}],\"name\":\"cancelOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"orderIds\",\"type\":\"uint256[]\"}],\"name\":\"emergencyCancelOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIPair.OrderSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256[2]\",\"name\":\"adjacent\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"findMaxCount\",\"type\":\"uint256\"}],\"name\":\"findPrevPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"QUOTE\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"BASE\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"DENOMINATOR\",\"type\":\"uint256\"}],\"internalType\":\"structIPair.Config\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quote\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tickSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lotSize\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lotSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"matchedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"matchedPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minTradeVolume\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"orderById\",\"outputs\":[{\"components\":[{\"internalType\":\"enumIPair.OrderSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"feeBps\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIPair.Order\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIPair.OrderSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"name\":\"ordersByPrices\",\"outputs\":[{\"internalType\":\"uint256[][]\",\"name\":\"\",\"type\":\"uint256[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quoteReserve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"setPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lotSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tickSize\",\"type\":\"uint256\"}],\"name\":\"setTickSize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"erc20\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"skim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumIPair.OrderSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"feeBps\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIPair.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"enumIPair.LimitConstraints\",\"name\":\"constraints\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"prevPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxMatchCount\",\"type\":\"uint256\"}],\"name\":\"submitLimitOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumIPair.OrderSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"feeBps\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIPair.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"spendAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxMatchCount\",\"type\":\"uint256\"}],\"name\":\"submitMarketOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tickSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tickSizes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tick\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lot\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ticks\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"sellPrices\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"buyPrices\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeBps\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"FeeCollect\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"enumIPair.CloseType\",\"name\":\"closeType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"OrderClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"enumIPair.OrderSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"OrderCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"sellId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"buyId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"OrderMatched\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"erc20\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Skim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"beforeLotSize\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newLotSize\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"beforeTickSize\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTickSize\",\"type\":\"uint256\"}],\"name\":\"TickSizeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ListFailToFindPrev\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ListInvalidFindMaxCount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ListInvalidIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ListInvalidPrevNode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ListZeroData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"PairFillOrKill\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PairInsufficientTradeVolume\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"PairInvalidAccountReserve\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PairInvalidAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"PairInvalidInitializeData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PairInvalidOrderId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumIPair.OrderSide\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"PairInvalidOrderSide\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumIPair.OrderSide\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PairInvalidPrevPrice\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PairInvalidPrice\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"PairInvalidReserve\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"PairInvalidRouter\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PairInvalidTickSize\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"PairNotOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumIPair.OrderSide\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PairUnknownPrices\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"}]",
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
		"6c4cfbc8": "findPrevPrice(uint8,uint256,uint256[2],uint256)",
		"c3f909d4": "getConfig()",
		"a6b63eb8": "initialize(address,address,address,uint256,uint256)",
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
	Bin: "0x60a060405230608052348015610013575f5ffd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b608051614f576100f95f395f8181612e5d01528181612e8601526130da0152614f575ff3fe6080604052600436106101db575f3560e01c80638da5cb5b116100fd578063cc50e36a11610092578063ec342ad011610062578063ec342ad01461063e578063f210d0871461066a578063f46f16c21461067f578063fe566349146106aa575f5ffd5b8063cc50e36a146105d4578063d6e9cc2a146105f3578063dfdf2a721461060a578063e077dd8a1461061f575f5ffd5b8063a6b63eb8116100cd578063a6b63eb8146104c3578063ad3cb1cc146104e2578063bedb86fb14610537578063c3f909d414610556575f5ffd5b80638da5cb5b14610459578063918f86741461046d5780639c579839146104825780639da771f4146104ae575f5ffd5b80634d6754f1116101735780636c4cfbc8116101435780636c4cfbc81461039b5780637aae3523146103ba5780637cbf6db2146104185780638d77eba51461042d575f5ffd5b80634d6754f11461031e5780634f1ef2861461033357806352d1902d146103465780635c975abb1461035a575f5ffd5b80632cfffaf6116101ae5780632cfffaf6146102775780632dd6a00a1461029957806332fe7b26146102b85780634942f65f14610309575f5ffd5b806312a808cc146101df5780631a0361fb146102145780631ec482d7146102355780631edd861414610254575b5f5ffd5b3480156101ea575f5ffd5b506101fe6101f93660046144da565b6106c9565b60405161020b9190614599565b60405180910390f35b34801561021f575f5ffd5b5061023361022e366004614664565b6107af565b005b348015610240575f5ffd5b5061023361024f3660046146e3565b610b1a565b34801561025f575f5ffd5b5061026960075481565b60405190815260200161020b565b348015610282575f5ffd5b5061028b610d2a565b60405161020b92919061476e565b3480156102a4575f5ffd5b506102696102b3366004614828565b610d64565b3480156102c3575f5ffd5b506001546102e49073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161020b565b348015610314575f5ffd5b50610269600a5481565b348015610329575f5ffd5b5061026960085481565b61023361034136600461486e565b6112ff565b348015610351575f5ffd5b5061026961131e565b348015610365575f5ffd5b507fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff16604051901515815260200161020b565b3480156103a6575f5ffd5b506102696103b5366004614933565b61134c565b3480156103c5575f5ffd5b506104036103d436600461497b565b73ffffffffffffffffffffffffffffffffffffffff165f90815260186020526040902080546001909101549091565b6040805192835260208301919091520161020b565b348015610423575f5ffd5b50610269600b5481565b348015610438575f5ffd5b5061044c610447366004614996565b61142e565b60405161020b9190614a13565b348015610464575f5ffd5b506102e46114fd565b348015610478575f5ffd5b5061026960045481565b34801561048d575f5ffd5b506003546102e49073ffffffffffffffffffffffffffffffffffffffff1681565b3480156104b9575f5ffd5b5061026960065481565b3480156104ce575f5ffd5b506102336104dd366004614a6f565b611590565b3480156104ed575f5ffd5b5061052a6040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b60405161020b9190614ac6565b348015610542575f5ffd5b50610233610551366004614b19565b611a7f565b348015610561575f5ffd5b5060408051606080820183525f8083526020808401829052928401528251808201845260035473ffffffffffffffffffffffffffffffffffffffff90811680835260025482168386019081526004549387019384528651918252519091169381019390935251928201929092520161020b565b3480156105df575f5ffd5b506102336105ee366004614b38565b611ad8565b3480156105fe575f5ffd5b50600954600a54610403565b348015610615575f5ffd5b5061026960055481565b34801561062a575f5ffd5b50610233610639366004614b69565b611ccd565b348015610649575f5ffd5b506002546102e49073ffffffffffffffffffffffffffffffffffffffff1681565b348015610675575f5ffd5b5061026960095481565b34801561068a575f5ffd5b505f546102e49073ffffffffffffffffffffffffffffffffffffffff1681565b3480156106b5575f5ffd5b506102336106c4366004614ba8565b611e1a565b60605f808460018111156106df576106df6149ad565b146106eb5760166106ee565b60155b83519091505f8167ffffffffffffffff81111561070d5761070d61445e565b60405190808252806020026020018201604052801561074057816020015b606081526020019060019003908161072b5790505b5090505f5b828110156107a35761077e845f88848151811061076457610764614bc8565b602002602001015181526020019081526020015f20612032565b82828151811061079057610790614bc8565b6020908102919091010152600101610745565b50925050505b92915050565b6107b76114fd565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461083b57335b6040517f118cdaa700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911660048201526024015b60405180910390fd5b8015610b155760025473ffffffffffffffffffffffffffffffffffffffff848116911614801561090557506005546002546040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152839173ffffffffffffffffffffffffffffffffffffffff16906370a0823190602401602060405180830381865afa1580156108d5573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108f99190614bf5565b6109039190614c39565b105b15610958576002546040517f7bdf264d00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610832565b60035473ffffffffffffffffffffffffffffffffffffffff8481169116148015610a1c57506006546003546040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152839173ffffffffffffffffffffffffffffffffffffffff16906370a0823190602401602060405180830381865afa1580156109ec573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a109190614bf5565b610a1a9190614c39565b105b15610a6f576003546040517f7bdf264d00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610832565b610a9073ffffffffffffffffffffffffffffffffffffffff841683836120d8565b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16610ac53390565b73ffffffffffffffffffffffffffffffffffffffff167f7e30e394efdaabad9f9ca6fd8f67f9449ba63d674ee8068283c3ae1f49b5b66084604051610b0c91815260200190565b60405180910390a45b505050565b60015473ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610b9c57335b6040517f143db19100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610832565b805f5b81811015610d23575f848483818110610bba57610bba614bc8565b602090810292909201355f81815260179093526040808420815160a08101909252805492955090925090829060ff166001811115610bfa57610bfa6149ad565b6001811115610c0b57610c0b6149ad565b81528154610100810473ffffffffffffffffffffffffffffffffffffffff908116602080850191909152750100000000000000000000000000000000000000000090920463ffffffff1660408401526001840154606084015260029093015460809092019190915282015191925016610c85575050610b9f565b8673ffffffffffffffffffffffffffffffffffffffff16816020015173ffffffffffffffffffffffffffffffffffffffff1614610d0d576040517f54de53f70000000000000000000000000000000000000000000000000000000081526004810183905273ffffffffffffffffffffffffffffffffffffffff88166024820152604401610832565b610d1982826002612165565b5050600101610b9f565b5050505050565b606080610d50600d5f5b60ff1660028110610d4757610d47614bc8565b60040201612032565b9150610d5e600d6001610d34565b90509091565b5f610d6d6122b1565b60015473ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610da85733610b51565b610db061230f565b60608501511580610dd157506009548560600151610dce9190614c79565b15155b15610e105784606001516040517f4e1404fe00000000000000000000000000000000000000000000000000000000815260040161083291815260200190565b60808501511580610e315750600a548560800151610e2e9190614c79565b15155b15610e705784608001516040517fa334626e00000000000000000000000000000000000000000000000000000000815260040161083291815260200190565b600c5f8154610e7e90614cb1565b918290555090505f8086516001811115610e9a57610e9a6149ad565b1490505f5f82610eb557610eb084895f8861245e565b610ec0565b610ec08489876126c2565b915091508115610f4e575f847f78615f01112d8477ae8d6fc9a510bcfdf935bc3ab7b02d0ff2316f5b44bd597a42604051610efd91815260200190565b60405180910390a3828015610f155750608088015115155b15610f495760208801516080890151600254610f499273ffffffffffffffffffffffffffffffffffffffff909116916120d8565b6112e1565b6001876002811115610f6257610f626149ad565b03610fd8576001847f78615f01112d8477ae8d6fc9a510bcfdf935bc3ab7b02d0ff2316f5b44bd597a42604051610f9b91815260200190565b60405180910390a38215610f495760208801516080890151600254610f499273ffffffffffffffffffffffffffffffffffffffff909116916120d8565b6002876002811115610fec57610fec6149ad565b036110415760208801516040517f1b3c356200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610832565b8515801590611072575082801561105b5750858860600151105b806110725750821580156110725750858860600151115b156110b457875160608901516040517fc18aa6060000000000000000000000000000000000000000000000000000000081526108329291908990600401614ce8565b6110eb83158960200151856110dd576110d88b606001518c60800151600454612927565b6110e3565b8a608001515b6129fd612a24565b611129886060015187600d8b5f0151600181111561110b5761110b6149ad565b60ff166002811061111e5761111e614bc8565b600402019190612bb2565b505f848152601760205260409020885181548a92919082907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660018381811115611176576111766149ad565b021790555060208201518154604084015163ffffffff167501000000000000000000000000000000000000000000027fffffffffffffff00000000ffffffffffffffffffffffffffffffffffffffffff73ffffffffffffffffffffffffffffffffffffffff90931661010002929092167fffffffffffffff000000000000000000000000000000000000000000000000ff909116171781556060820151600182015560809091015160029091015582156112c3577f1d2ff3fa6980aeeebca4e94965520da48983006e9b1115c1c853cbb10d943d005c5f85815260176020908152604080832080547fffffffffffffff00000000ffffffffffffffffffffffffffffffffffffffffff16750100000000000000000000000000000000000000000063ffffffff9687160217905560608c01518352601590915290206112bd918690612d6016565b506112e1565b60608801515f9081526016602052604090206112df9085612d60565b505b826112f4576112f4886020015182612d70565b505050949350505050565b611307612e45565b61131082612f49565b61131a8282612f89565b5050565b5f6113276130c2565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b5f80856001811115611360576113606149ad565b036113c9576040805180820182526113c291869190869060029083908390808284375f92019190915250869150600d90508960018111156113a3576113a36149ad565b60ff16600281106113b6576113b6614bc8565b60040201929190613131565b9050611426565b6040805180820182526113c291869190869060029083908390808284375f92019190915250869150600d9050896001811115611407576114076149ad565b60ff166002811061141a5761141a614bc8565b60040201929190613315565b949350505050565b6040805160a0810182525f808252602082018190529181018290526060810182905260808101919091525f8281526017602052604090819020815160a081019092528054829060ff166001811115611488576114886149ad565b6001811115611499576114996149ad565b81528154610100810473ffffffffffffffffffffffffffffffffffffffff1660208301527501000000000000000000000000000000000000000000900463ffffffff1660408201526001820154606082015260029091015460809091015292915050565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611567573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061158b9190614d07565b905090565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff165f811580156115da5750825b90505f8267ffffffffffffffff1660011480156115f65750303b155b905081158015611604575080155b1561163b576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000166001178555831561169c5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b6116a46134c2565b73ffffffffffffffffffffffffffffffffffffffff8a16611713576040517fc587916e0000000000000000000000000000000000000000000000000000000081527f726f7574657200000000000000000000000000000000000000000000000000006004820152602401610832565b73ffffffffffffffffffffffffffffffffffffffff8916611782576040517fc587916e0000000000000000000000000000000000000000000000000000000081527f71756f74650000000000000000000000000000000000000000000000000000006004820152602401610832565b73ffffffffffffffffffffffffffffffffffffffff88166117f1576040517fc587916e0000000000000000000000000000000000000000000000000000000081527f62617365000000000000000000000000000000000000000000000000000000006004820152602401610832565b865f0361184c576040517fc587916e0000000000000000000000000000000000000000000000000000000081527f7469636b53697a650000000000000000000000000000000000000000000000006004820152602401610832565b855f036118a7576040517fc587916e0000000000000000000000000000000000000000000000000000000081527f6c6f7453697a65000000000000000000000000000000000000000000000000006004820152602401610832565b335f80547fffffffffffffffffffffffff000000000000000000000000000000000000000090811673ffffffffffffffffffffffffffffffffffffffff938416179091556001805482168d84161790556003805482168c841617905560028054909116918a169182179055604080517f313ce567000000000000000000000000000000000000000000000000000000008152905163313ce567916004808201926020929091908290030181865afa158015611964573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906119889190614d22565b61199390600a614e5b565b60048190556119a28789614e69565b6119ac9190614c79565b156119f557600480546040517f57568135000000000000000000000000000000000000000000000000000000008152918201899052602482018890526044820152606401610832565b6009879055600a869055600454611a0f9088908890612927565b600b558315611a735784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050505050565b611a876114fd565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611abf57336107eb565b8015611ad057611acd6134d2565b50565b611acd613572565b611ae06122b1565b60015473ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611b1b5733610b51565b611b2361230f565b5f600c5f8154611b3290614cb1565b918290555090505f84516001811115611b4d57611b4d6149ad565b03611bfc57821580611b6a5750600a54611b679084614c79565b15155b15611ba4576040517fa334626e00000000000000000000000000000000000000000000000000000000815260048101849052602401610832565b5f606085015260808401839052611bbc8185846126c2565b5050608084015115611bf75760208401516080850151600254611bf79273ffffffffffffffffffffffffffffffffffffffff909116916120d8565b611c8c565b600b54831015611c4657600b546040517f0bc1e7dd000000000000000000000000000000000000000000000000000000008152610832918591600401918252602082015260400190565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60608501525f611c798286868661245e565b915050611c8a856020015182612d70565b505b5f817f78615f01112d8477ae8d6fc9a510bcfdf935bc3ab7b02d0ff2316f5b44bd597a42604051611cbf91815260200190565b60405180910390a350505050565b611cd56135e8565b611cdd6114fd565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611d1557336107eb565b805f5b81811015611e14575f848483818110611d3357611d33614bc8565b602090810292909201355f81815260179093526040808420815160a08101909252805492955090925090829060ff166001811115611d7357611d736149ad565b6001811115611d8457611d846149ad565b81528154610100810473ffffffffffffffffffffffffffffffffffffffff908116602080850191909152750100000000000000000000000000000000000000000090920463ffffffff1660408401526001840154606084015260029093015460809092019190915282015191925016611dfe575050611d18565b611e0a82826003612165565b5050600101611d18565b50505050565b5f5473ffffffffffffffffffffffffffffffffffffffff1663a1eea778336040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815273ffffffffffffffffffffffffffffffffffffffff90911660048201526024015f6040518083038186803b158015611e9a575f5ffd5b505afa158015611eac573d5f5f3e3d5ffd5b50505050805f03611f0b576040517fc587916e0000000000000000000000000000000000000000000000000000000081527f7469636b53697a650000000000000000000000000000000000000000000000006004820152602401610832565b815f03611f66576040517fc587916e0000000000000000000000000000000000000000000000000000000081527f6c6f7453697a65000000000000000000000000000000000000000000000000006004820152602401610832565b600454611f738383614e69565b611f7d9190614c79565b15611fc657600480546040517f57568135000000000000000000000000000000000000000000000000000000008152918201839052602482018490526044820152606401610832565b600a546009546040805192835260208301859052820152606081018290527f66748bc112f6692a76bebecc33e69768508e27e414e3080f5c02295b6ea3cf3f9060800160405180910390a1600a829055600981905560045461202b9082908490612927565b600b555050565b60605f825f015467ffffffffffffffff8111156120515761205161445e565b60405190808252806020026020018201604052801561207a578160200160208202803683370190505b5083546001850154919250905f5b828110156120ce57818482815181106120a3576120a3614bc8565b6020908102919091018101919091525f92835260038701905260409091206001908101549101612088565b5091949350505050565b6040805173ffffffffffffffffffffffffffffffffffffffff8416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb00000000000000000000000000000000000000000000000000000000179052610b15908490613643565b5f80808451600181111561217b5761217b6149ad565b14905080156121e75760608401515f908152601560205260409020608085015190925080156121e15760208501516002546121cf9173ffffffffffffffffffffffffffffffffffffffff90911690836120d8565b6121e15f8660200151836136e2612a24565b5061225c565b60165f856060015181526020019081526020015f2091505f61221485606001518660800151600454612927565b9050801561225a5760208501516003546122479173ffffffffffffffffffffffffffffffffffffffff90911690836120d8565b61225a60018660200151836136e2612a24565b505b612267858484613702565b8154610d23576122a98460600151600d865f0151600181111561228c5761228c6149ad565b60ff166002811061229f5761229f614bc8565b60040201906137d4565b505050505050565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff161561230d576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505f5f8273ffffffffffffffffffffffffffffffffffffffff1663c415b95c6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561237e573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906123a29190614d07565b8373ffffffffffffffffffffffffffffffffffffffff166324a9d8536040518163ffffffff1660e01b8152600401602060405180830381865afa1580156123eb573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061240f9190614e80565b91509150817fd6aa07baf8485abf9d26fecf4c935d75b50a73e678db02b944bd3ac8759823005d807f1d2ff3fa6980aeeebca4e94965520da48983006e9b1115c1c853cbb10d943d005d505050565b5f80600185516001811115612475576124756149ad565b146124ae575f6040517f7cbdd0810000000000000000000000000000000000000000000000000000000081526004016108329190614e9b565b5f84156124bc5750836124d4565b6124d186606001518760800151600454612927565b90505b6003546040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201525f91829161257d9173ffffffffffffffffffffffffffffffffffffffff16906370a0823190602401602060405180830381865afa158015612546573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061256a9190614bf5565b846006546125789190614ea9565b6136e2565b91509150816125d4576003546040517f7bdf264d00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610832565b875160018111156125e7576125e76149ad565b89896020015173ffffffffffffffffffffffffffffffffffffffff167f7cde7559ef00f8c3d9244d3fde19e7500979d64f6f9c086a3836a4926abbe1658b606001518c608001514260405161264f939291909283526020830191909152604082015260600190565b60405180910390a45f5f6126658b8b8b8b6138a5565b91509150805f146126b1578060055f8282546126819190614c39565b909155505060208a01516002546126b19173ffffffffffffffffffffffffffffffffffffffff90911690836120d8565b509450925050505b94509492505050565b5f8080845160018111156126d8576126d86149ad565b146127125760016040517f7cbdd0810000000000000000000000000000000000000000000000000000000081526004016108329190614e9b565b83608001516005546127249190614ea9565b6002546040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff909116906370a0823190602401602060405180830381865afa158015612790573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906127b49190614bf5565b1015612808576002546040517f7bdf264d00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610832565b8351600181111561281b5761281b6149ad565b85856020015173ffffffffffffffffffffffffffffffffffffffff167f7cde7559ef00f8c3d9244d3fde19e7500979d64f6f9c086a3836a4926abbe1658760600151886080015142604051612883939291909283526020830191909152604082015260600190565b60405180910390a45f5f612898878787613ae8565b91509150805f1461291857600680547fd6aa07baf8485abf9d26fecf4c935d75b50a73e678db02b944bd3ac8759823005c917f1d2ff3fa6980aeeebca4e94965520da48983006e9b1115c1c853cbb10d943d005c918491905f906128fd908490614c39565b92505081905550612915898960200151858585613cd5565b50505b5091505f90505b935093915050565b5f838302817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff85870982811083820303915050805f0361297a5783828161297057612970614c4c565b04925050506129f6565b808411612991576129916003851502601118613dec565b5f848688095f868103871696879004966002600389028118808a02820302808a02820302808a02820302808a02820302808a02820302808a02909103029181900381900460010186841190950394909402919094039290920491909117919091029150505b9392505050565b5f8083830184811015612a16575f5f9250925050612a1d565b6001925090505b9250929050565b5f8415612aad57612a3b600654848463ffffffff16565b60065590508015612aa85773ffffffffffffffffffffffffffffffffffffffff84165f908152601860205260409020612a7d9060015b0154848463ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff86165f9081526018602052604090206001015590505b612b1d565b612abd600554848463ffffffff16565b60055590508015612b1d5773ffffffffffffffffffffffffffffffffffffffff84165f908152601860205260408120612af591612a71565b73ffffffffffffffffffffffffffffffffffffffff86165f9081526018602052604090205590505b80610d23578385612b465760025473ffffffffffffffffffffffffffffffffffffffff16612b60565b60035473ffffffffffffffffffffffffffffffffffffffff165b6040517f340b575200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff928316600482015291166024820152604401610832565b5f825f03612bec576040517f0da93b9700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612bf68484613dfd565b15612c0257505f6129f6565b5f828152600385016020908152604080832081518083019092528382529181019290925290835f03612c6e57600186018054908690558015612c51575f81815260038801602052604090208690555b60405180604001604052805f815260200182815250915050612d10565b604080518082019091528254815260018301546020820152612c8f90613e45565b80612c9d5750856001015484145b612cd3576040517f182dfca500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5060018101546040805180820190915284815260208101829052908015612d07575f81815260038801602052604090208690555b50600182018590555b80602001515f03612d2357600286018590555b5f8581526003870160209081526040822083518155908301516001909101558654879190612d5090614cb1565b9091555060019695505050505050565b5f6129f683838560020154612bb2565b5f60065482612d7f9190614ea9565b6003546040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff909116906370a0823190602401602060405180830381865afa158015612deb573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612e0f9190614bf5565b612e199190614c39565b90508015610b1557600354610b159073ffffffffffffffffffffffffffffffffffffffff1684836120d8565b3073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161480612f1257507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16612ef97f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614155b1561230d576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612f516114fd565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611acd57336107eb565b8173ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa92505050801561300e575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016820190925261300b91810190614bf5565b60015b61305c576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610832565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc81146130b8576040517faa1d49a400000000000000000000000000000000000000000000000000000000815260048101829052602401610832565b610b158383613e5b565b3073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161461230d576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f835f0361316b576040517f0da93b9700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6131758585613dfd565b1561319057505f838152600385016020526040902054611426565b84600101548410806131a157508454155b156131ad57505f611426565b84600201548411156131c457506002840154611426565b815f036131fd576040517fab7e639300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6132088685613ebd565b905084811015613287575b801580159061324657507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff909201918215155b15613282575f9081526003860160205260409020600101548481111561327d575f9081526003860160205260409020549050611426565b613213565b6132e3565b80158015906132ba57507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff909201918215155b156132e3575f908152600386016020526040902054848110156132de579050611426565b613287565b6040517f80ce60d100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f835f0361334f576040517f0da93b9700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6133598585613dfd565b1561337457505f838152600385016020526040902054611426565b846001015484118061338557508454155b1561339157505f611426565b84600201548410156133a857506002840154611426565b815f036133e1576040517fab7e639300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6133ec8685613ebd565b905084811115613466575b801580159061342a57507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff909201918215155b15613282575f90815260038601602052604090206001015480851115613461575f9081526003860160205260409020549050611426565b6133f7565b801580159061349957507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff909201918215155b156132e3575f908152600386016020526040902054808510156134bd579050611426565b613466565b6134ca613f07565b61230d613f6e565b6134da6122b1565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a150565b61357a6135e8565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa33613547565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff1661230d576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5f60205f8451602086015f885af180613662576040513d5f823e3d81fd5b50505f513d91508115613679578060011415613693565b73ffffffffffffffffffffffffffffffffffffffff84163b155b15611e14576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610832565b5f5f838311156136f657505f905080612a1d565b50600193919092039150565b61370c81846137d4565b613745576040517ffcfdf90200000000000000000000000000000000000000000000000000000000815260048101849052602401610832565b5f83815260176020526040812080547fffffffffffffff000000000000000000000000000000000000000000000000001681556001810182905560020155816003811115613795576137956149ad565b837f78615f01112d8477ae8d6fc9a510bcfdf935bc3ab7b02d0ff2316f5b44bd597a426040516137c791815260200190565b60405180910390a3505050565b5f815f0361380e576040517f0da93b9700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6138188383613dfd565b61382357505f6107a9565b5f82815260038401602052604080822080548352818320600180830154808652939094208482019390935581548355928601549092919085900361386c57600180840154908701555b8486600201540361387f57825460028701555b5f85815260038701602052604081208181556001018190558654879190612d5090614ebc565b5f8080600d5b805415613ada575f6138bd8282613fbf565b905087606001518111156138d15750613ada565b5f8181526015602052604090206138e78261404c565b871561393f575f6139046138fb868b614c39565b60045485612927565b9050600a54816139149190614c79565b61391e9082614c39565b60808b0181905290505f81900361393d57600196505050505050613ae0565b505b805415613a8f575f6139518282613fbf565b5f81815260176020526040812091925080806139718f8f87878b8b614072565b9250925092505f6139858884600454612927565b90506139b58685837fd6aa07baf8485abf9d26fecf4c935d75b50a73e678db02b944bd3ac8759823005c86613cd5565b6139bf838c614ea9565b9a506139cb818b614ea9565b73ffffffffffffffffffffffffffffffffffffffff85165f908152601860205260408120919b508491015f828254613a039190614c39565b909155505060808f01511580613a225750613a1d8d614ebc565b9c508c155b15613a84578654613a7157613a3789896137d4565b613a71575f886040517f9f1cfdfe000000000000000000000000000000000000000000000000000000008152600401610832929190614ef0565b60019b5050505050505050505050613ae0565b50505050505061393f565b613a9983836137d4565b613ad3575f826040517f9f1cfdfe000000000000000000000000000000000000000000000000000000008152600401610832929190614ef0565b50506138ab565b5f935050505b6126b9614198565b6002545f90819073ffffffffffffffffffffffffffffffffffffffff1660115b805415613cc7575f613b1a8282613fbf565b90508660600151811015613b2e5750613cc7565b5f818152601660205260409020613b448261404c565b805415613c7b575f613b568282613fbf565b5f81815260176020526040812091925080613b758d8d86868a8a614072565b509092509050613b9c73ffffffffffffffffffffffffffffffffffffffff891683836120d8565b5f613baa8783600454612927565b9050613bb6818b614ea9565b73ffffffffffffffffffffffffffffffffffffffff84165f908152601860205260409020909a5081906001015f828254613bf09190614c39565b909155505060808d01511580613c0f5750613c0a8c614ebc565b9b508b155b15613c71578554613c5f57613c2488886137d4565b613c5f576001876040517f9f1cfdfe000000000000000000000000000000000000000000000000000000008152600401610832929190614ef0565b60019a50505050505050505050613ccd565b5050505050613b44565b613c8583836137d4565b613cc0576001826040517f9f1cfdfe000000000000000000000000000000000000000000000000000000008152600401610832929190614ef0565b5050613b08565b5f935050505b61291f614198565b8063ffffffff165f03613d0b57600354613d069073ffffffffffffffffffffffffffffffffffffffff1685856120d8565b610d23565b5f613d1f848363ffffffff16612710612927565b90505f613d2c8286614c39565b6040805187815263ffffffff861660208201529081018490526060810182905290915073ffffffffffffffffffffffffffffffffffffffff808616919088169089907f9814b2b2f43f9ea51ad7e8760fda094ffa3012bfeb10c76a9656389af18956039060800160405180910390a4600354613dbf9073ffffffffffffffffffffffffffffffffffffffff1685846120d8565b600354613de39073ffffffffffffffffffffffffffffffffffffffff1687836120d8565b50505050505050565b634e487b715f52806020526024601cfd5b5f815f03613e0c57505f6107a9565b82600101548214806129f657505f82815260038401602090815260409182902082518084019093528054835260010154908201526129f6905b80515f901515806107a957505060200151151590565b613e64826141e6565b60405173ffffffffffffffffffffffffffffffffffffffff8316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115613eb557610b1582826142b4565b61131a614333565b5f613ecf8383835b6020020151613dfd565b15613ee357815f5b602002015190506107a9565b613eef83836001613ec5565b15613efc57816001613ed7565b5060018201546107a9565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff1661230d576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b613f76613f07565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b81545f908210613ffb576040517f39e60f7300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001830154825b8015614044575f91825260038501602052604090912060010154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01614002565b509392505050565b807ffd0e5d4f9b88892d3b04349a0e2bc0d1359414c21932fcd7d5a523a6c0a5cd005d50565b5f5f5f855f0160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff166140ab8960800151886002015461436b565b875491945092507501000000000000000000000000000000000000000000900463ffffffff1690505f80808a5160018111156140e9576140e96149ad565b146140f557888b6140f8565b8a895b915091508681837f6a6896b1e6131e0b8ebae43fdc84cb0178a6eacdf4ee63f15dabae48729a8a038742604051614139929190918252602082015260400190565b60405180910390a48760020154840361415c57614157895f88613702565b614168565b60028801805485900390555b8960800151840361417e575f60808b015261418a565b60808a01805185900390525b505096509650969350505050565b7ffd0e5d4f9b88892d3b04349a0e2bc0d1359414c21932fcd7d5a523a6c0a5cd005c80158015906141cb57506007548114155b156141d65760078190555b4260085414611acd574260085550565b8073ffffffffffffffffffffffffffffffffffffffff163b5f0361424e576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152602401610832565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60605f5f8473ffffffffffffffffffffffffffffffffffffffff16846040516142dd9190614f0b565b5f60405180830381855af49150503d805f8114614315576040519150601f19603f3d011682016040523d82523d5f602084013e61431a565b606091505b509150915061432a85838361437a565b95945050505050565b341561230d576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8282188284100282186129f6565b60608261438f5761438a82614409565b6129f6565b81511580156143b3575073ffffffffffffffffffffffffffffffffffffffff84163b155b15614402576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610832565b50806129f6565b8051156144195780518082602001fd5b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b803560028110614459575f5ffd5b919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156144d2576144d261445e565b604052919050565b5f5f604083850312156144eb575f5ffd5b6144f48361444b565b9150602083013567ffffffffffffffff81111561450f575f5ffd5b8301601f8101851361451f575f5ffd5b803567ffffffffffffffff8111156145395761453961445e565b8060051b6145496020820161448b565b91825260208184018101929081019088841115614564575f5ffd5b6020850194505b8385101561458a5784358083526020958601959093509091019061456b565b80955050505050509250929050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b82811015614637578685037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0018452815180518087526020918201918701905f5b8181101561461e578351835260209384019390920191600101614600565b50909650505060209384019391909101906001016145bf565b50929695505050505050565b73ffffffffffffffffffffffffffffffffffffffff81168114611acd575f5ffd5b5f5f5f60608486031215614676575f5ffd5b833561468181614643565b9250602084013561469181614643565b929592945050506040919091013590565b5f5f83601f8401126146b2575f5ffd5b50813567ffffffffffffffff8111156146c9575f5ffd5b6020830191508360208260051b8501011115612a1d575f5ffd5b5f5f5f604084860312156146f5575f5ffd5b833561470081614643565b9250602084013567ffffffffffffffff81111561471b575f5ffd5b614727868287016146a2565b9497909650939450505050565b5f8151808452602084019350602083015f5b82811015614764578151865260209586019590910190600101614746565b5093949350505050565b604081525f6147806040830185614734565b828103602084015261432a8185614734565b63ffffffff81168114611acd575f5ffd5b5f60a082840312156147b3575f5ffd5b60405160a0810167ffffffffffffffff811182821017156147d6576147d661445e565b6040529050806147e58361444b565b815260208301356147f581614643565b6020820152604083013561480881614792565b604082015260608381013590820152608092830135920191909152919050565b5f5f5f5f610100858703121561483c575f5ffd5b61484686866147a3565b935060a085013560038110614859575f5ffd5b939693955050505060c08201359160e0013590565b5f5f6040838503121561487f575f5ffd5b823561488a81614643565b9150602083013567ffffffffffffffff8111156148a5575f5ffd5b8301601f810185136148b5575f5ffd5b803567ffffffffffffffff8111156148cf576148cf61445e565b61490060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8401160161448b565b818152866020838501011115614914575f5ffd5b816020840160208301375f602083830101528093505050509250929050565b5f5f5f5f60a08587031215614946575f5ffd5b61494f8561444b565b9350602085013592506080850186811115614968575f5ffd5b9396929550505060409290920191903590565b5f6020828403121561498b575f5ffd5b81356129f681614643565b5f602082840312156149a6575f5ffd5b5035919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b60028110614a0f577f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b9052565b5f60a082019050614a258284516149da565b73ffffffffffffffffffffffffffffffffffffffff602084015116602083015263ffffffff6040840151166040830152606083015160608301526080830151608083015292915050565b5f5f5f5f5f60a08688031215614a83575f5ffd5b8535614a8e81614643565b94506020860135614a9e81614643565b93506040860135614aae81614643565b94979396509394606081013594506080013592915050565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b5f60208284031215614b29575f5ffd5b813580151581146129f6575f5ffd5b5f5f5f60e08486031215614b4a575f5ffd5b614b5485856147a3565b9560a0850135955060c0909401359392505050565b5f5f60208385031215614b7a575f5ffd5b823567ffffffffffffffff811115614b90575f5ffd5b614b9c858286016146a2565b90969095509350505050565b5f5f60408385031215614bb9575f5ffd5b50508035926020909101359150565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f60208284031215614c05575f5ffd5b5051919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b818103818111156107a9576107a9614c0c565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f82614cac577f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b500690565b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203614ce157614ce1614c0c565b5060010190565b60608101614cf682866149da565b602082019390935260400152919050565b5f60208284031215614d17575f5ffd5b81516129f681614643565b5f60208284031215614d32575f5ffd5b815160ff811681146129f6575f5ffd5b6001815b600184111561291f57808504811115614d6157614d61614c0c565b6001841615614d6f57908102905b60019390931c928002614d46565b5f82614d8b575060016107a9565b81614d9757505f6107a9565b8160018114614dad5760028114614db757614dd3565b60019150506107a9565b60ff841115614dc857614dc8614c0c565b50506001821b6107a9565b5060208310610133831016604e8410600b8410161715614df6575081810a6107a9565b614e217fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8484614d42565b807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115614e5357614e53614c0c565b029392505050565b5f6129f660ff841683614d7d565b80820281158282048414176107a9576107a9614c0c565b5f60208284031215614e90575f5ffd5b81516129f681614792565b602081016107a982846149da565b808201808211156107a9576107a9614c0c565b5f81614eca57614eca614c0c565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b60408101614efe82856149da565b8260208301529392505050565b5f82518060208501845e5f92019182525091905056fea264697066735822122046112a9be4b11f99334ffa50806d0cef26027927c7532488608ed315b81636a264736f6c634300081c0033",
}

// PairImplABI is the input ABI used to generate the binding from.
// Deprecated: Use PairImplMetaData.ABI instead.
var PairImplABI = PairImplMetaData.ABI

// Deprecated: Use PairImplMetaData.Sigs instead.
// PairImplFuncSigs maps the 4-byte function signature to its string representation.
var PairImplFuncSigs = PairImplMetaData.Sigs

// PairImplBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PairImplMetaData.Bin instead.
var PairImplBin = PairImplMetaData.Bin

// DeployPairImpl deploys a new Ethereum contract, binding an instance of PairImpl to it.
func DeployPairImpl(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PairImpl, error) {
	parsed, err := PairImplMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PairImplBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PairImpl{PairImplCaller: PairImplCaller{contract: contract}, PairImplTransactor: PairImplTransactor{contract: contract}, PairImplFilterer: PairImplFilterer{contract: contract}}, nil
}

// PairImpl is an auto generated Go binding around an Ethereum contract.
type PairImpl struct {
	PairImplCaller     // Read-only binding to the contract
	PairImplTransactor // Write-only binding to the contract
	PairImplFilterer   // Log filterer for contract events
}

// PairImplCaller is an auto generated read-only Go binding around an Ethereum contract.
type PairImplCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairImplTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PairImplTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairImplFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PairImplFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairImplSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PairImplSession struct {
	Contract     *PairImpl         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PairImplCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PairImplCallerSession struct {
	Contract *PairImplCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PairImplTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PairImplTransactorSession struct {
	Contract     *PairImplTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PairImplRaw is an auto generated low-level Go binding around an Ethereum contract.
type PairImplRaw struct {
	Contract *PairImpl // Generic contract binding to access the raw methods on
}

// PairImplCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PairImplCallerRaw struct {
	Contract *PairImplCaller // Generic read-only contract binding to access the raw methods on
}

// PairImplTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PairImplTransactorRaw struct {
	Contract *PairImplTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPairImpl creates a new instance of PairImpl, bound to a specific deployed contract.
func NewPairImpl(address common.Address, backend bind.ContractBackend) (*PairImpl, error) {
	contract, err := bindPairImpl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PairImpl{PairImplCaller: PairImplCaller{contract: contract}, PairImplTransactor: PairImplTransactor{contract: contract}, PairImplFilterer: PairImplFilterer{contract: contract}}, nil
}

// NewPairImplCaller creates a new read-only instance of PairImpl, bound to a specific deployed contract.
func NewPairImplCaller(address common.Address, caller bind.ContractCaller) (*PairImplCaller, error) {
	contract, err := bindPairImpl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PairImplCaller{contract: contract}, nil
}

// NewPairImplTransactor creates a new write-only instance of PairImpl, bound to a specific deployed contract.
func NewPairImplTransactor(address common.Address, transactor bind.ContractTransactor) (*PairImplTransactor, error) {
	contract, err := bindPairImpl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PairImplTransactor{contract: contract}, nil
}

// NewPairImplFilterer creates a new log filterer instance of PairImpl, bound to a specific deployed contract.
func NewPairImplFilterer(address common.Address, filterer bind.ContractFilterer) (*PairImplFilterer, error) {
	contract, err := bindPairImpl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PairImplFilterer{contract: contract}, nil
}

// bindPairImpl binds a generic wrapper to an already deployed contract.
func bindPairImpl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PairImplMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PairImpl *PairImplRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PairImpl.Contract.PairImplCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PairImpl *PairImplRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PairImpl.Contract.PairImplTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PairImpl *PairImplRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PairImpl.Contract.PairImplTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PairImpl *PairImplCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PairImpl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PairImpl *PairImplTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PairImpl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PairImpl *PairImplTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PairImpl.Contract.contract.Transact(opts, method, params...)
}

// BASE is a free data retrieval call binding the contract method 0xec342ad0.
//
// Solidity: function BASE() view returns(address)
func (_PairImpl *PairImplCaller) BASE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PairImpl.contract.Call(opts, &out, "BASE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BASE is a free data retrieval call binding the contract method 0xec342ad0.
//
// Solidity: function BASE() view returns(address)
func (_PairImpl *PairImplSession) BASE() (common.Address, error) {
	return _PairImpl.Contract.BASE(&_PairImpl.CallOpts)
}

// BASE is a free data retrieval call binding the contract method 0xec342ad0.
//
// Solidity: function BASE() view returns(address)
func (_PairImpl *PairImplCallerSession) BASE() (common.Address, error) {
	return _PairImpl.Contract.BASE(&_PairImpl.CallOpts)
}

// DENOMINATOR is a free data retrieval call binding the contract method 0x918f8674.
//
// Solidity: function DENOMINATOR() view returns(uint256)
func (_PairImpl *PairImplCaller) DENOMINATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PairImpl.contract.Call(opts, &out, "DENOMINATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DENOMINATOR is a free data retrieval call binding the contract method 0x918f8674.
//
// Solidity: function DENOMINATOR() view returns(uint256)
func (_PairImpl *PairImplSession) DENOMINATOR() (*big.Int, error) {
	return _PairImpl.Contract.DENOMINATOR(&_PairImpl.CallOpts)
}

// DENOMINATOR is a free data retrieval call binding the contract method 0x918f8674.
//
// Solidity: function DENOMINATOR() view returns(uint256)
func (_PairImpl *PairImplCallerSession) DENOMINATOR() (*big.Int, error) {
	return _PairImpl.Contract.DENOMINATOR(&_PairImpl.CallOpts)
}

// MARKET is a free data retrieval call binding the contract method 0xf46f16c2.
//
// Solidity: function MARKET() view returns(address)
func (_PairImpl *PairImplCaller) MARKET(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PairImpl.contract.Call(opts, &out, "MARKET")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MARKET is a free data retrieval call binding the contract method 0xf46f16c2.
//
// Solidity: function MARKET() view returns(address)
func (_PairImpl *PairImplSession) MARKET() (common.Address, error) {
	return _PairImpl.Contract.MARKET(&_PairImpl.CallOpts)
}

// MARKET is a free data retrieval call binding the contract method 0xf46f16c2.
//
// Solidity: function MARKET() view returns(address)
func (_PairImpl *PairImplCallerSession) MARKET() (common.Address, error) {
	return _PairImpl.Contract.MARKET(&_PairImpl.CallOpts)
}

// QUOTE is a free data retrieval call binding the contract method 0x9c579839.
//
// Solidity: function QUOTE() view returns(address)
func (_PairImpl *PairImplCaller) QUOTE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PairImpl.contract.Call(opts, &out, "QUOTE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// QUOTE is a free data retrieval call binding the contract method 0x9c579839.
//
// Solidity: function QUOTE() view returns(address)
func (_PairImpl *PairImplSession) QUOTE() (common.Address, error) {
	return _PairImpl.Contract.QUOTE(&_PairImpl.CallOpts)
}

// QUOTE is a free data retrieval call binding the contract method 0x9c579839.
//
// Solidity: function QUOTE() view returns(address)
func (_PairImpl *PairImplCallerSession) QUOTE() (common.Address, error) {
	return _PairImpl.Contract.QUOTE(&_PairImpl.CallOpts)
}

// ROUTER is a free data retrieval call binding the contract method 0x32fe7b26.
//
// Solidity: function ROUTER() view returns(address)
func (_PairImpl *PairImplCaller) ROUTER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PairImpl.contract.Call(opts, &out, "ROUTER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ROUTER is a free data retrieval call binding the contract method 0x32fe7b26.
//
// Solidity: function ROUTER() view returns(address)
func (_PairImpl *PairImplSession) ROUTER() (common.Address, error) {
	return _PairImpl.Contract.ROUTER(&_PairImpl.CallOpts)
}

// ROUTER is a free data retrieval call binding the contract method 0x32fe7b26.
//
// Solidity: function ROUTER() view returns(address)
func (_PairImpl *PairImplCallerSession) ROUTER() (common.Address, error) {
	return _PairImpl.Contract.ROUTER(&_PairImpl.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_PairImpl *PairImplCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PairImpl.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_PairImpl *PairImplSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _PairImpl.Contract.UPGRADEINTERFACEVERSION(&_PairImpl.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_PairImpl *PairImplCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _PairImpl.Contract.UPGRADEINTERFACEVERSION(&_PairImpl.CallOpts)
}

// AccountReserves is a free data retrieval call binding the contract method 0x7aae3523.
//
// Solidity: function accountReserves(address account) view returns(uint256 base, uint256 quote)
func (_PairImpl *PairImplCaller) AccountReserves(opts *bind.CallOpts, account common.Address) (struct {
	Base  *big.Int
	Quote *big.Int
}, error) {
	var out []interface{}
	err := _PairImpl.contract.Call(opts, &out, "accountReserves", account)

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
func (_PairImpl *PairImplSession) AccountReserves(account common.Address) (struct {
	Base  *big.Int
	Quote *big.Int
}, error) {
	return _PairImpl.Contract.AccountReserves(&_PairImpl.CallOpts, account)
}

// AccountReserves is a free data retrieval call binding the contract method 0x7aae3523.
//
// Solidity: function accountReserves(address account) view returns(uint256 base, uint256 quote)
func (_PairImpl *PairImplCallerSession) AccountReserves(account common.Address) (struct {
	Base  *big.Int
	Quote *big.Int
}, error) {
	return _PairImpl.Contract.AccountReserves(&_PairImpl.CallOpts, account)
}

// BaseReserve is a free data retrieval call binding the contract method 0xdfdf2a72.
//
// Solidity: function baseReserve() view returns(uint256)
func (_PairImpl *PairImplCaller) BaseReserve(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PairImpl.contract.Call(opts, &out, "baseReserve")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseReserve is a free data retrieval call binding the contract method 0xdfdf2a72.
//
// Solidity: function baseReserve() view returns(uint256)
func (_PairImpl *PairImplSession) BaseReserve() (*big.Int, error) {
	return _PairImpl.Contract.BaseReserve(&_PairImpl.CallOpts)
}

// BaseReserve is a free data retrieval call binding the contract method 0xdfdf2a72.
//
// Solidity: function baseReserve() view returns(uint256)
func (_PairImpl *PairImplCallerSession) BaseReserve() (*big.Int, error) {
	return _PairImpl.Contract.BaseReserve(&_PairImpl.CallOpts)
}

// FindPrevPrice is a free data retrieval call binding the contract method 0x6c4cfbc8.
//
// Solidity: function findPrevPrice(uint8 side, uint256 price, uint256[2] adjacent, uint256 findMaxCount) view returns(uint256)
func (_PairImpl *PairImplCaller) FindPrevPrice(opts *bind.CallOpts, side uint8, price *big.Int, adjacent [2]*big.Int, findMaxCount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PairImpl.contract.Call(opts, &out, "findPrevPrice", side, price, adjacent, findMaxCount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FindPrevPrice is a free data retrieval call binding the contract method 0x6c4cfbc8.
//
// Solidity: function findPrevPrice(uint8 side, uint256 price, uint256[2] adjacent, uint256 findMaxCount) view returns(uint256)
func (_PairImpl *PairImplSession) FindPrevPrice(side uint8, price *big.Int, adjacent [2]*big.Int, findMaxCount *big.Int) (*big.Int, error) {
	return _PairImpl.Contract.FindPrevPrice(&_PairImpl.CallOpts, side, price, adjacent, findMaxCount)
}

// FindPrevPrice is a free data retrieval call binding the contract method 0x6c4cfbc8.
//
// Solidity: function findPrevPrice(uint8 side, uint256 price, uint256[2] adjacent, uint256 findMaxCount) view returns(uint256)
func (_PairImpl *PairImplCallerSession) FindPrevPrice(side uint8, price *big.Int, adjacent [2]*big.Int, findMaxCount *big.Int) (*big.Int, error) {
	return _PairImpl.Contract.FindPrevPrice(&_PairImpl.CallOpts, side, price, adjacent, findMaxCount)
}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() view returns((address,address,uint256))
func (_PairImpl *PairImplCaller) GetConfig(opts *bind.CallOpts) (IPairConfig, error) {
	var out []interface{}
	err := _PairImpl.contract.Call(opts, &out, "getConfig")

	if err != nil {
		return *new(IPairConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IPairConfig)).(*IPairConfig)

	return out0, err

}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() view returns((address,address,uint256))
func (_PairImpl *PairImplSession) GetConfig() (IPairConfig, error) {
	return _PairImpl.Contract.GetConfig(&_PairImpl.CallOpts)
}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() view returns((address,address,uint256))
func (_PairImpl *PairImplCallerSession) GetConfig() (IPairConfig, error) {
	return _PairImpl.Contract.GetConfig(&_PairImpl.CallOpts)
}

// LotSize is a free data retrieval call binding the contract method 0x4942f65f.
//
// Solidity: function lotSize() view returns(uint256)
func (_PairImpl *PairImplCaller) LotSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PairImpl.contract.Call(opts, &out, "lotSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LotSize is a free data retrieval call binding the contract method 0x4942f65f.
//
// Solidity: function lotSize() view returns(uint256)
func (_PairImpl *PairImplSession) LotSize() (*big.Int, error) {
	return _PairImpl.Contract.LotSize(&_PairImpl.CallOpts)
}

// LotSize is a free data retrieval call binding the contract method 0x4942f65f.
//
// Solidity: function lotSize() view returns(uint256)
func (_PairImpl *PairImplCallerSession) LotSize() (*big.Int, error) {
	return _PairImpl.Contract.LotSize(&_PairImpl.CallOpts)
}

// MatchedAt is a free data retrieval call binding the contract method 0x4d6754f1.
//
// Solidity: function matchedAt() view returns(uint256)
func (_PairImpl *PairImplCaller) MatchedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PairImpl.contract.Call(opts, &out, "matchedAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MatchedAt is a free data retrieval call binding the contract method 0x4d6754f1.
//
// Solidity: function matchedAt() view returns(uint256)
func (_PairImpl *PairImplSession) MatchedAt() (*big.Int, error) {
	return _PairImpl.Contract.MatchedAt(&_PairImpl.CallOpts)
}

// MatchedAt is a free data retrieval call binding the contract method 0x4d6754f1.
//
// Solidity: function matchedAt() view returns(uint256)
func (_PairImpl *PairImplCallerSession) MatchedAt() (*big.Int, error) {
	return _PairImpl.Contract.MatchedAt(&_PairImpl.CallOpts)
}

// MatchedPrice is a free data retrieval call binding the contract method 0x1edd8614.
//
// Solidity: function matchedPrice() view returns(uint256)
func (_PairImpl *PairImplCaller) MatchedPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PairImpl.contract.Call(opts, &out, "matchedPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MatchedPrice is a free data retrieval call binding the contract method 0x1edd8614.
//
// Solidity: function matchedPrice() view returns(uint256)
func (_PairImpl *PairImplSession) MatchedPrice() (*big.Int, error) {
	return _PairImpl.Contract.MatchedPrice(&_PairImpl.CallOpts)
}

// MatchedPrice is a free data retrieval call binding the contract method 0x1edd8614.
//
// Solidity: function matchedPrice() view returns(uint256)
func (_PairImpl *PairImplCallerSession) MatchedPrice() (*big.Int, error) {
	return _PairImpl.Contract.MatchedPrice(&_PairImpl.CallOpts)
}

// MinTradeVolume is a free data retrieval call binding the contract method 0x7cbf6db2.
//
// Solidity: function minTradeVolume() view returns(uint256)
func (_PairImpl *PairImplCaller) MinTradeVolume(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PairImpl.contract.Call(opts, &out, "minTradeVolume")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinTradeVolume is a free data retrieval call binding the contract method 0x7cbf6db2.
//
// Solidity: function minTradeVolume() view returns(uint256)
func (_PairImpl *PairImplSession) MinTradeVolume() (*big.Int, error) {
	return _PairImpl.Contract.MinTradeVolume(&_PairImpl.CallOpts)
}

// MinTradeVolume is a free data retrieval call binding the contract method 0x7cbf6db2.
//
// Solidity: function minTradeVolume() view returns(uint256)
func (_PairImpl *PairImplCallerSession) MinTradeVolume() (*big.Int, error) {
	return _PairImpl.Contract.MinTradeVolume(&_PairImpl.CallOpts)
}

// OrderById is a free data retrieval call binding the contract method 0x8d77eba5.
//
// Solidity: function orderById(uint256 id) view returns((uint8,address,uint32,uint256,uint256))
func (_PairImpl *PairImplCaller) OrderById(opts *bind.CallOpts, id *big.Int) (IPairOrder, error) {
	var out []interface{}
	err := _PairImpl.contract.Call(opts, &out, "orderById", id)

	if err != nil {
		return *new(IPairOrder), err
	}

	out0 := *abi.ConvertType(out[0], new(IPairOrder)).(*IPairOrder)

	return out0, err

}

// OrderById is a free data retrieval call binding the contract method 0x8d77eba5.
//
// Solidity: function orderById(uint256 id) view returns((uint8,address,uint32,uint256,uint256))
func (_PairImpl *PairImplSession) OrderById(id *big.Int) (IPairOrder, error) {
	return _PairImpl.Contract.OrderById(&_PairImpl.CallOpts, id)
}

// OrderById is a free data retrieval call binding the contract method 0x8d77eba5.
//
// Solidity: function orderById(uint256 id) view returns((uint8,address,uint32,uint256,uint256))
func (_PairImpl *PairImplCallerSession) OrderById(id *big.Int) (IPairOrder, error) {
	return _PairImpl.Contract.OrderById(&_PairImpl.CallOpts, id)
}

// OrdersByPrices is a free data retrieval call binding the contract method 0x12a808cc.
//
// Solidity: function ordersByPrices(uint8 side, uint256[] prices) view returns(uint256[][])
func (_PairImpl *PairImplCaller) OrdersByPrices(opts *bind.CallOpts, side uint8, prices []*big.Int) ([][]*big.Int, error) {
	var out []interface{}
	err := _PairImpl.contract.Call(opts, &out, "ordersByPrices", side, prices)

	if err != nil {
		return *new([][]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([][]*big.Int)).(*[][]*big.Int)

	return out0, err

}

// OrdersByPrices is a free data retrieval call binding the contract method 0x12a808cc.
//
// Solidity: function ordersByPrices(uint8 side, uint256[] prices) view returns(uint256[][])
func (_PairImpl *PairImplSession) OrdersByPrices(side uint8, prices []*big.Int) ([][]*big.Int, error) {
	return _PairImpl.Contract.OrdersByPrices(&_PairImpl.CallOpts, side, prices)
}

// OrdersByPrices is a free data retrieval call binding the contract method 0x12a808cc.
//
// Solidity: function ordersByPrices(uint8 side, uint256[] prices) view returns(uint256[][])
func (_PairImpl *PairImplCallerSession) OrdersByPrices(side uint8, prices []*big.Int) ([][]*big.Int, error) {
	return _PairImpl.Contract.OrdersByPrices(&_PairImpl.CallOpts, side, prices)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PairImpl *PairImplCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PairImpl.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PairImpl *PairImplSession) Owner() (common.Address, error) {
	return _PairImpl.Contract.Owner(&_PairImpl.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PairImpl *PairImplCallerSession) Owner() (common.Address, error) {
	return _PairImpl.Contract.Owner(&_PairImpl.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PairImpl *PairImplCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PairImpl.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PairImpl *PairImplSession) Paused() (bool, error) {
	return _PairImpl.Contract.Paused(&_PairImpl.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PairImpl *PairImplCallerSession) Paused() (bool, error) {
	return _PairImpl.Contract.Paused(&_PairImpl.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_PairImpl *PairImplCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PairImpl.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_PairImpl *PairImplSession) ProxiableUUID() ([32]byte, error) {
	return _PairImpl.Contract.ProxiableUUID(&_PairImpl.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_PairImpl *PairImplCallerSession) ProxiableUUID() ([32]byte, error) {
	return _PairImpl.Contract.ProxiableUUID(&_PairImpl.CallOpts)
}

// QuoteReserve is a free data retrieval call binding the contract method 0x9da771f4.
//
// Solidity: function quoteReserve() view returns(uint256)
func (_PairImpl *PairImplCaller) QuoteReserve(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PairImpl.contract.Call(opts, &out, "quoteReserve")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuoteReserve is a free data retrieval call binding the contract method 0x9da771f4.
//
// Solidity: function quoteReserve() view returns(uint256)
func (_PairImpl *PairImplSession) QuoteReserve() (*big.Int, error) {
	return _PairImpl.Contract.QuoteReserve(&_PairImpl.CallOpts)
}

// QuoteReserve is a free data retrieval call binding the contract method 0x9da771f4.
//
// Solidity: function quoteReserve() view returns(uint256)
func (_PairImpl *PairImplCallerSession) QuoteReserve() (*big.Int, error) {
	return _PairImpl.Contract.QuoteReserve(&_PairImpl.CallOpts)
}

// TickSize is a free data retrieval call binding the contract method 0xf210d087.
//
// Solidity: function tickSize() view returns(uint256)
func (_PairImpl *PairImplCaller) TickSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PairImpl.contract.Call(opts, &out, "tickSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TickSize is a free data retrieval call binding the contract method 0xf210d087.
//
// Solidity: function tickSize() view returns(uint256)
func (_PairImpl *PairImplSession) TickSize() (*big.Int, error) {
	return _PairImpl.Contract.TickSize(&_PairImpl.CallOpts)
}

// TickSize is a free data retrieval call binding the contract method 0xf210d087.
//
// Solidity: function tickSize() view returns(uint256)
func (_PairImpl *PairImplCallerSession) TickSize() (*big.Int, error) {
	return _PairImpl.Contract.TickSize(&_PairImpl.CallOpts)
}

// TickSizes is a free data retrieval call binding the contract method 0xd6e9cc2a.
//
// Solidity: function tickSizes() view returns(uint256 tick, uint256 lot)
func (_PairImpl *PairImplCaller) TickSizes(opts *bind.CallOpts) (struct {
	Tick *big.Int
	Lot  *big.Int
}, error) {
	var out []interface{}
	err := _PairImpl.contract.Call(opts, &out, "tickSizes")

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
func (_PairImpl *PairImplSession) TickSizes() (struct {
	Tick *big.Int
	Lot  *big.Int
}, error) {
	return _PairImpl.Contract.TickSizes(&_PairImpl.CallOpts)
}

// TickSizes is a free data retrieval call binding the contract method 0xd6e9cc2a.
//
// Solidity: function tickSizes() view returns(uint256 tick, uint256 lot)
func (_PairImpl *PairImplCallerSession) TickSizes() (struct {
	Tick *big.Int
	Lot  *big.Int
}, error) {
	return _PairImpl.Contract.TickSizes(&_PairImpl.CallOpts)
}

// Ticks is a free data retrieval call binding the contract method 0x2cfffaf6.
//
// Solidity: function ticks() view returns(uint256[] sellPrices, uint256[] buyPrices)
func (_PairImpl *PairImplCaller) Ticks(opts *bind.CallOpts) (struct {
	SellPrices []*big.Int
	BuyPrices  []*big.Int
}, error) {
	var out []interface{}
	err := _PairImpl.contract.Call(opts, &out, "ticks")

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
func (_PairImpl *PairImplSession) Ticks() (struct {
	SellPrices []*big.Int
	BuyPrices  []*big.Int
}, error) {
	return _PairImpl.Contract.Ticks(&_PairImpl.CallOpts)
}

// Ticks is a free data retrieval call binding the contract method 0x2cfffaf6.
//
// Solidity: function ticks() view returns(uint256[] sellPrices, uint256[] buyPrices)
func (_PairImpl *PairImplCallerSession) Ticks() (struct {
	SellPrices []*big.Int
	BuyPrices  []*big.Int
}, error) {
	return _PairImpl.Contract.Ticks(&_PairImpl.CallOpts)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x1ec482d7.
//
// Solidity: function cancelOrder(address caller, uint256[] orderIds) returns()
func (_PairImpl *PairImplTransactor) CancelOrder(opts *bind.TransactOpts, caller common.Address, orderIds []*big.Int) (*types.Transaction, error) {
	return _PairImpl.contract.Transact(opts, "cancelOrder", caller, orderIds)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x1ec482d7.
//
// Solidity: function cancelOrder(address caller, uint256[] orderIds) returns()
func (_PairImpl *PairImplSession) CancelOrder(caller common.Address, orderIds []*big.Int) (*types.Transaction, error) {
	return _PairImpl.Contract.CancelOrder(&_PairImpl.TransactOpts, caller, orderIds)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x1ec482d7.
//
// Solidity: function cancelOrder(address caller, uint256[] orderIds) returns()
func (_PairImpl *PairImplTransactorSession) CancelOrder(caller common.Address, orderIds []*big.Int) (*types.Transaction, error) {
	return _PairImpl.Contract.CancelOrder(&_PairImpl.TransactOpts, caller, orderIds)
}

// EmergencyCancelOrder is a paid mutator transaction binding the contract method 0xe077dd8a.
//
// Solidity: function emergencyCancelOrder(uint256[] orderIds) returns()
func (_PairImpl *PairImplTransactor) EmergencyCancelOrder(opts *bind.TransactOpts, orderIds []*big.Int) (*types.Transaction, error) {
	return _PairImpl.contract.Transact(opts, "emergencyCancelOrder", orderIds)
}

// EmergencyCancelOrder is a paid mutator transaction binding the contract method 0xe077dd8a.
//
// Solidity: function emergencyCancelOrder(uint256[] orderIds) returns()
func (_PairImpl *PairImplSession) EmergencyCancelOrder(orderIds []*big.Int) (*types.Transaction, error) {
	return _PairImpl.Contract.EmergencyCancelOrder(&_PairImpl.TransactOpts, orderIds)
}

// EmergencyCancelOrder is a paid mutator transaction binding the contract method 0xe077dd8a.
//
// Solidity: function emergencyCancelOrder(uint256[] orderIds) returns()
func (_PairImpl *PairImplTransactorSession) EmergencyCancelOrder(orderIds []*big.Int) (*types.Transaction, error) {
	return _PairImpl.Contract.EmergencyCancelOrder(&_PairImpl.TransactOpts, orderIds)
}

// Initialize is a paid mutator transaction binding the contract method 0xa6b63eb8.
//
// Solidity: function initialize(address router, address quote, address base, uint256 _tickSize, uint256 _lotSize) returns()
func (_PairImpl *PairImplTransactor) Initialize(opts *bind.TransactOpts, router common.Address, quote common.Address, base common.Address, _tickSize *big.Int, _lotSize *big.Int) (*types.Transaction, error) {
	return _PairImpl.contract.Transact(opts, "initialize", router, quote, base, _tickSize, _lotSize)
}

// Initialize is a paid mutator transaction binding the contract method 0xa6b63eb8.
//
// Solidity: function initialize(address router, address quote, address base, uint256 _tickSize, uint256 _lotSize) returns()
func (_PairImpl *PairImplSession) Initialize(router common.Address, quote common.Address, base common.Address, _tickSize *big.Int, _lotSize *big.Int) (*types.Transaction, error) {
	return _PairImpl.Contract.Initialize(&_PairImpl.TransactOpts, router, quote, base, _tickSize, _lotSize)
}

// Initialize is a paid mutator transaction binding the contract method 0xa6b63eb8.
//
// Solidity: function initialize(address router, address quote, address base, uint256 _tickSize, uint256 _lotSize) returns()
func (_PairImpl *PairImplTransactorSession) Initialize(router common.Address, quote common.Address, base common.Address, _tickSize *big.Int, _lotSize *big.Int) (*types.Transaction, error) {
	return _PairImpl.Contract.Initialize(&_PairImpl.TransactOpts, router, quote, base, _tickSize, _lotSize)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool pause) returns()
func (_PairImpl *PairImplTransactor) SetPause(opts *bind.TransactOpts, pause bool) (*types.Transaction, error) {
	return _PairImpl.contract.Transact(opts, "setPause", pause)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool pause) returns()
func (_PairImpl *PairImplSession) SetPause(pause bool) (*types.Transaction, error) {
	return _PairImpl.Contract.SetPause(&_PairImpl.TransactOpts, pause)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool pause) returns()
func (_PairImpl *PairImplTransactorSession) SetPause(pause bool) (*types.Transaction, error) {
	return _PairImpl.Contract.SetPause(&_PairImpl.TransactOpts, pause)
}

// SetTickSize is a paid mutator transaction binding the contract method 0xfe566349.
//
// Solidity: function setTickSize(uint256 _lotSize, uint256 _tickSize) returns()
func (_PairImpl *PairImplTransactor) SetTickSize(opts *bind.TransactOpts, _lotSize *big.Int, _tickSize *big.Int) (*types.Transaction, error) {
	return _PairImpl.contract.Transact(opts, "setTickSize", _lotSize, _tickSize)
}

// SetTickSize is a paid mutator transaction binding the contract method 0xfe566349.
//
// Solidity: function setTickSize(uint256 _lotSize, uint256 _tickSize) returns()
func (_PairImpl *PairImplSession) SetTickSize(_lotSize *big.Int, _tickSize *big.Int) (*types.Transaction, error) {
	return _PairImpl.Contract.SetTickSize(&_PairImpl.TransactOpts, _lotSize, _tickSize)
}

// SetTickSize is a paid mutator transaction binding the contract method 0xfe566349.
//
// Solidity: function setTickSize(uint256 _lotSize, uint256 _tickSize) returns()
func (_PairImpl *PairImplTransactorSession) SetTickSize(_lotSize *big.Int, _tickSize *big.Int) (*types.Transaction, error) {
	return _PairImpl.Contract.SetTickSize(&_PairImpl.TransactOpts, _lotSize, _tickSize)
}

// Skim is a paid mutator transaction binding the contract method 0x1a0361fb.
//
// Solidity: function skim(address erc20, address to, uint256 amount) returns()
func (_PairImpl *PairImplTransactor) Skim(opts *bind.TransactOpts, erc20 common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PairImpl.contract.Transact(opts, "skim", erc20, to, amount)
}

// Skim is a paid mutator transaction binding the contract method 0x1a0361fb.
//
// Solidity: function skim(address erc20, address to, uint256 amount) returns()
func (_PairImpl *PairImplSession) Skim(erc20 common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PairImpl.Contract.Skim(&_PairImpl.TransactOpts, erc20, to, amount)
}

// Skim is a paid mutator transaction binding the contract method 0x1a0361fb.
//
// Solidity: function skim(address erc20, address to, uint256 amount) returns()
func (_PairImpl *PairImplTransactorSession) Skim(erc20 common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PairImpl.Contract.Skim(&_PairImpl.TransactOpts, erc20, to, amount)
}

// SubmitLimitOrder is a paid mutator transaction binding the contract method 0x2dd6a00a.
//
// Solidity: function submitLimitOrder((uint8,address,uint32,uint256,uint256) order, uint8 constraints, uint256 prevPrice, uint256 maxMatchCount) returns(uint256 orderId)
func (_PairImpl *PairImplTransactor) SubmitLimitOrder(opts *bind.TransactOpts, order IPairOrder, constraints uint8, prevPrice *big.Int, maxMatchCount *big.Int) (*types.Transaction, error) {
	return _PairImpl.contract.Transact(opts, "submitLimitOrder", order, constraints, prevPrice, maxMatchCount)
}

// SubmitLimitOrder is a paid mutator transaction binding the contract method 0x2dd6a00a.
//
// Solidity: function submitLimitOrder((uint8,address,uint32,uint256,uint256) order, uint8 constraints, uint256 prevPrice, uint256 maxMatchCount) returns(uint256 orderId)
func (_PairImpl *PairImplSession) SubmitLimitOrder(order IPairOrder, constraints uint8, prevPrice *big.Int, maxMatchCount *big.Int) (*types.Transaction, error) {
	return _PairImpl.Contract.SubmitLimitOrder(&_PairImpl.TransactOpts, order, constraints, prevPrice, maxMatchCount)
}

// SubmitLimitOrder is a paid mutator transaction binding the contract method 0x2dd6a00a.
//
// Solidity: function submitLimitOrder((uint8,address,uint32,uint256,uint256) order, uint8 constraints, uint256 prevPrice, uint256 maxMatchCount) returns(uint256 orderId)
func (_PairImpl *PairImplTransactorSession) SubmitLimitOrder(order IPairOrder, constraints uint8, prevPrice *big.Int, maxMatchCount *big.Int) (*types.Transaction, error) {
	return _PairImpl.Contract.SubmitLimitOrder(&_PairImpl.TransactOpts, order, constraints, prevPrice, maxMatchCount)
}

// SubmitMarketOrder is a paid mutator transaction binding the contract method 0xcc50e36a.
//
// Solidity: function submitMarketOrder((uint8,address,uint32,uint256,uint256) order, uint256 spendAmount, uint256 maxMatchCount) returns()
func (_PairImpl *PairImplTransactor) SubmitMarketOrder(opts *bind.TransactOpts, order IPairOrder, spendAmount *big.Int, maxMatchCount *big.Int) (*types.Transaction, error) {
	return _PairImpl.contract.Transact(opts, "submitMarketOrder", order, spendAmount, maxMatchCount)
}

// SubmitMarketOrder is a paid mutator transaction binding the contract method 0xcc50e36a.
//
// Solidity: function submitMarketOrder((uint8,address,uint32,uint256,uint256) order, uint256 spendAmount, uint256 maxMatchCount) returns()
func (_PairImpl *PairImplSession) SubmitMarketOrder(order IPairOrder, spendAmount *big.Int, maxMatchCount *big.Int) (*types.Transaction, error) {
	return _PairImpl.Contract.SubmitMarketOrder(&_PairImpl.TransactOpts, order, spendAmount, maxMatchCount)
}

// SubmitMarketOrder is a paid mutator transaction binding the contract method 0xcc50e36a.
//
// Solidity: function submitMarketOrder((uint8,address,uint32,uint256,uint256) order, uint256 spendAmount, uint256 maxMatchCount) returns()
func (_PairImpl *PairImplTransactorSession) SubmitMarketOrder(order IPairOrder, spendAmount *big.Int, maxMatchCount *big.Int) (*types.Transaction, error) {
	return _PairImpl.Contract.SubmitMarketOrder(&_PairImpl.TransactOpts, order, spendAmount, maxMatchCount)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_PairImpl *PairImplTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _PairImpl.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_PairImpl *PairImplSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _PairImpl.Contract.UpgradeToAndCall(&_PairImpl.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_PairImpl *PairImplTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _PairImpl.Contract.UpgradeToAndCall(&_PairImpl.TransactOpts, newImplementation, data)
}

// PairImplFeeCollectIterator is returned from FilterFeeCollect and is used to iterate over the raw logs and unpacked data for FeeCollect events raised by the PairImpl contract.
type PairImplFeeCollectIterator struct {
	Event *PairImplFeeCollect // Event containing the contract specifics and raw log

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
func (it *PairImplFeeCollectIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PairImplFeeCollect)
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
		it.Event = new(PairImplFeeCollect)
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
func (it *PairImplFeeCollectIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PairImplFeeCollectIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PairImplFeeCollect represents a FeeCollect event raised by the PairImpl contract.
type PairImplFeeCollect struct {
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
func (_PairImpl *PairImplFilterer) FilterFeeCollect(opts *bind.FilterOpts, orderId []*big.Int, owner []common.Address, recipient []common.Address) (*PairImplFeeCollectIterator, error) {

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

	logs, sub, err := _PairImpl.contract.FilterLogs(opts, "FeeCollect", orderIdRule, ownerRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &PairImplFeeCollectIterator{contract: _PairImpl.contract, event: "FeeCollect", logs: logs, sub: sub}, nil
}

// WatchFeeCollect is a free log subscription operation binding the contract event 0x9814b2b2f43f9ea51ad7e8760fda094ffa3012bfeb10c76a9656389af1895603.
//
// Solidity: event FeeCollect(uint256 indexed orderId, address indexed owner, uint256 amount, address indexed recipient, uint256 feeBps, uint256 fee, uint256 value)
func (_PairImpl *PairImplFilterer) WatchFeeCollect(opts *bind.WatchOpts, sink chan<- *PairImplFeeCollect, orderId []*big.Int, owner []common.Address, recipient []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PairImpl.contract.WatchLogs(opts, "FeeCollect", orderIdRule, ownerRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PairImplFeeCollect)
				if err := _PairImpl.contract.UnpackLog(event, "FeeCollect", log); err != nil {
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
func (_PairImpl *PairImplFilterer) ParseFeeCollect(log types.Log) (*PairImplFeeCollect, error) {
	event := new(PairImplFeeCollect)
	if err := _PairImpl.contract.UnpackLog(event, "FeeCollect", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PairImplInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the PairImpl contract.
type PairImplInitializedIterator struct {
	Event *PairImplInitialized // Event containing the contract specifics and raw log

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
func (it *PairImplInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PairImplInitialized)
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
		it.Event = new(PairImplInitialized)
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
func (it *PairImplInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PairImplInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PairImplInitialized represents a Initialized event raised by the PairImpl contract.
type PairImplInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_PairImpl *PairImplFilterer) FilterInitialized(opts *bind.FilterOpts) (*PairImplInitializedIterator, error) {

	logs, sub, err := _PairImpl.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PairImplInitializedIterator{contract: _PairImpl.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_PairImpl *PairImplFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PairImplInitialized) (event.Subscription, error) {

	logs, sub, err := _PairImpl.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PairImplInitialized)
				if err := _PairImpl.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_PairImpl *PairImplFilterer) ParseInitialized(log types.Log) (*PairImplInitialized, error) {
	event := new(PairImplInitialized)
	if err := _PairImpl.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PairImplOrderClosedIterator is returned from FilterOrderClosed and is used to iterate over the raw logs and unpacked data for OrderClosed events raised by the PairImpl contract.
type PairImplOrderClosedIterator struct {
	Event *PairImplOrderClosed // Event containing the contract specifics and raw log

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
func (it *PairImplOrderClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PairImplOrderClosed)
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
		it.Event = new(PairImplOrderClosed)
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
func (it *PairImplOrderClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PairImplOrderClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PairImplOrderClosed represents a OrderClosed event raised by the PairImpl contract.
type PairImplOrderClosed struct {
	OrderId   *big.Int
	CloseType uint8
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOrderClosed is a free log retrieval operation binding the contract event 0x78615f01112d8477ae8d6fc9a510bcfdf935bc3ab7b02d0ff2316f5b44bd597a.
//
// Solidity: event OrderClosed(uint256 indexed orderId, uint8 indexed closeType, uint256 timestamp)
func (_PairImpl *PairImplFilterer) FilterOrderClosed(opts *bind.FilterOpts, orderId []*big.Int, closeType []uint8) (*PairImplOrderClosedIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var closeTypeRule []interface{}
	for _, closeTypeItem := range closeType {
		closeTypeRule = append(closeTypeRule, closeTypeItem)
	}

	logs, sub, err := _PairImpl.contract.FilterLogs(opts, "OrderClosed", orderIdRule, closeTypeRule)
	if err != nil {
		return nil, err
	}
	return &PairImplOrderClosedIterator{contract: _PairImpl.contract, event: "OrderClosed", logs: logs, sub: sub}, nil
}

// WatchOrderClosed is a free log subscription operation binding the contract event 0x78615f01112d8477ae8d6fc9a510bcfdf935bc3ab7b02d0ff2316f5b44bd597a.
//
// Solidity: event OrderClosed(uint256 indexed orderId, uint8 indexed closeType, uint256 timestamp)
func (_PairImpl *PairImplFilterer) WatchOrderClosed(opts *bind.WatchOpts, sink chan<- *PairImplOrderClosed, orderId []*big.Int, closeType []uint8) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var closeTypeRule []interface{}
	for _, closeTypeItem := range closeType {
		closeTypeRule = append(closeTypeRule, closeTypeItem)
	}

	logs, sub, err := _PairImpl.contract.WatchLogs(opts, "OrderClosed", orderIdRule, closeTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PairImplOrderClosed)
				if err := _PairImpl.contract.UnpackLog(event, "OrderClosed", log); err != nil {
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
func (_PairImpl *PairImplFilterer) ParseOrderClosed(log types.Log) (*PairImplOrderClosed, error) {
	event := new(PairImplOrderClosed)
	if err := _PairImpl.contract.UnpackLog(event, "OrderClosed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PairImplOrderCreatedIterator is returned from FilterOrderCreated and is used to iterate over the raw logs and unpacked data for OrderCreated events raised by the PairImpl contract.
type PairImplOrderCreatedIterator struct {
	Event *PairImplOrderCreated // Event containing the contract specifics and raw log

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
func (it *PairImplOrderCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PairImplOrderCreated)
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
		it.Event = new(PairImplOrderCreated)
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
func (it *PairImplOrderCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PairImplOrderCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PairImplOrderCreated represents a OrderCreated event raised by the PairImpl contract.
type PairImplOrderCreated struct {
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
func (_PairImpl *PairImplFilterer) FilterOrderCreated(opts *bind.FilterOpts, owner []common.Address, orderId []*big.Int, side []uint8) (*PairImplOrderCreatedIterator, error) {

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

	logs, sub, err := _PairImpl.contract.FilterLogs(opts, "OrderCreated", ownerRule, orderIdRule, sideRule)
	if err != nil {
		return nil, err
	}
	return &PairImplOrderCreatedIterator{contract: _PairImpl.contract, event: "OrderCreated", logs: logs, sub: sub}, nil
}

// WatchOrderCreated is a free log subscription operation binding the contract event 0x7cde7559ef00f8c3d9244d3fde19e7500979d64f6f9c086a3836a4926abbe165.
//
// Solidity: event OrderCreated(address indexed owner, uint256 indexed orderId, uint8 indexed side, uint256 price, uint256 amount, uint256 timestamp)
func (_PairImpl *PairImplFilterer) WatchOrderCreated(opts *bind.WatchOpts, sink chan<- *PairImplOrderCreated, owner []common.Address, orderId []*big.Int, side []uint8) (event.Subscription, error) {

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

	logs, sub, err := _PairImpl.contract.WatchLogs(opts, "OrderCreated", ownerRule, orderIdRule, sideRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PairImplOrderCreated)
				if err := _PairImpl.contract.UnpackLog(event, "OrderCreated", log); err != nil {
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
func (_PairImpl *PairImplFilterer) ParseOrderCreated(log types.Log) (*PairImplOrderCreated, error) {
	event := new(PairImplOrderCreated)
	if err := _PairImpl.contract.UnpackLog(event, "OrderCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PairImplOrderMatchedIterator is returned from FilterOrderMatched and is used to iterate over the raw logs and unpacked data for OrderMatched events raised by the PairImpl contract.
type PairImplOrderMatchedIterator struct {
	Event *PairImplOrderMatched // Event containing the contract specifics and raw log

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
func (it *PairImplOrderMatchedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PairImplOrderMatched)
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
		it.Event = new(PairImplOrderMatched)
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
func (it *PairImplOrderMatchedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PairImplOrderMatchedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PairImplOrderMatched represents a OrderMatched event raised by the PairImpl contract.
type PairImplOrderMatched struct {
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
func (_PairImpl *PairImplFilterer) FilterOrderMatched(opts *bind.FilterOpts, sellId []*big.Int, buyId []*big.Int, price []*big.Int) (*PairImplOrderMatchedIterator, error) {

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

	logs, sub, err := _PairImpl.contract.FilterLogs(opts, "OrderMatched", sellIdRule, buyIdRule, priceRule)
	if err != nil {
		return nil, err
	}
	return &PairImplOrderMatchedIterator{contract: _PairImpl.contract, event: "OrderMatched", logs: logs, sub: sub}, nil
}

// WatchOrderMatched is a free log subscription operation binding the contract event 0x6a6896b1e6131e0b8ebae43fdc84cb0178a6eacdf4ee63f15dabae48729a8a03.
//
// Solidity: event OrderMatched(uint256 indexed sellId, uint256 indexed buyId, uint256 indexed price, uint256 amount, uint256 timestamp)
func (_PairImpl *PairImplFilterer) WatchOrderMatched(opts *bind.WatchOpts, sink chan<- *PairImplOrderMatched, sellId []*big.Int, buyId []*big.Int, price []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _PairImpl.contract.WatchLogs(opts, "OrderMatched", sellIdRule, buyIdRule, priceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PairImplOrderMatched)
				if err := _PairImpl.contract.UnpackLog(event, "OrderMatched", log); err != nil {
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
func (_PairImpl *PairImplFilterer) ParseOrderMatched(log types.Log) (*PairImplOrderMatched, error) {
	event := new(PairImplOrderMatched)
	if err := _PairImpl.contract.UnpackLog(event, "OrderMatched", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PairImplPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the PairImpl contract.
type PairImplPausedIterator struct {
	Event *PairImplPaused // Event containing the contract specifics and raw log

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
func (it *PairImplPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PairImplPaused)
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
		it.Event = new(PairImplPaused)
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
func (it *PairImplPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PairImplPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PairImplPaused represents a Paused event raised by the PairImpl contract.
type PairImplPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PairImpl *PairImplFilterer) FilterPaused(opts *bind.FilterOpts) (*PairImplPausedIterator, error) {

	logs, sub, err := _PairImpl.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PairImplPausedIterator{contract: _PairImpl.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PairImpl *PairImplFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PairImplPaused) (event.Subscription, error) {

	logs, sub, err := _PairImpl.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PairImplPaused)
				if err := _PairImpl.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_PairImpl *PairImplFilterer) ParsePaused(log types.Log) (*PairImplPaused, error) {
	event := new(PairImplPaused)
	if err := _PairImpl.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PairImplSkimIterator is returned from FilterSkim and is used to iterate over the raw logs and unpacked data for Skim events raised by the PairImpl contract.
type PairImplSkimIterator struct {
	Event *PairImplSkim // Event containing the contract specifics and raw log

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
func (it *PairImplSkimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PairImplSkim)
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
		it.Event = new(PairImplSkim)
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
func (it *PairImplSkimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PairImplSkimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PairImplSkim represents a Skim event raised by the PairImpl contract.
type PairImplSkim struct {
	Caller common.Address
	Erc20  common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSkim is a free log retrieval operation binding the contract event 0x7e30e394efdaabad9f9ca6fd8f67f9449ba63d674ee8068283c3ae1f49b5b660.
//
// Solidity: event Skim(address indexed caller, address indexed erc20, address indexed to, uint256 amount)
func (_PairImpl *PairImplFilterer) FilterSkim(opts *bind.FilterOpts, caller []common.Address, erc20 []common.Address, to []common.Address) (*PairImplSkimIterator, error) {

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

	logs, sub, err := _PairImpl.contract.FilterLogs(opts, "Skim", callerRule, erc20Rule, toRule)
	if err != nil {
		return nil, err
	}
	return &PairImplSkimIterator{contract: _PairImpl.contract, event: "Skim", logs: logs, sub: sub}, nil
}

// WatchSkim is a free log subscription operation binding the contract event 0x7e30e394efdaabad9f9ca6fd8f67f9449ba63d674ee8068283c3ae1f49b5b660.
//
// Solidity: event Skim(address indexed caller, address indexed erc20, address indexed to, uint256 amount)
func (_PairImpl *PairImplFilterer) WatchSkim(opts *bind.WatchOpts, sink chan<- *PairImplSkim, caller []common.Address, erc20 []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PairImpl.contract.WatchLogs(opts, "Skim", callerRule, erc20Rule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PairImplSkim)
				if err := _PairImpl.contract.UnpackLog(event, "Skim", log); err != nil {
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
func (_PairImpl *PairImplFilterer) ParseSkim(log types.Log) (*PairImplSkim, error) {
	event := new(PairImplSkim)
	if err := _PairImpl.contract.UnpackLog(event, "Skim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PairImplTickSizeUpdatedIterator is returned from FilterTickSizeUpdated and is used to iterate over the raw logs and unpacked data for TickSizeUpdated events raised by the PairImpl contract.
type PairImplTickSizeUpdatedIterator struct {
	Event *PairImplTickSizeUpdated // Event containing the contract specifics and raw log

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
func (it *PairImplTickSizeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PairImplTickSizeUpdated)
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
		it.Event = new(PairImplTickSizeUpdated)
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
func (it *PairImplTickSizeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PairImplTickSizeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PairImplTickSizeUpdated represents a TickSizeUpdated event raised by the PairImpl contract.
type PairImplTickSizeUpdated struct {
	BeforeLotSize  *big.Int
	NewLotSize     *big.Int
	BeforeTickSize *big.Int
	NewTickSize    *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTickSizeUpdated is a free log retrieval operation binding the contract event 0x66748bc112f6692a76bebecc33e69768508e27e414e3080f5c02295b6ea3cf3f.
//
// Solidity: event TickSizeUpdated(uint256 beforeLotSize, uint256 newLotSize, uint256 beforeTickSize, uint256 newTickSize)
func (_PairImpl *PairImplFilterer) FilterTickSizeUpdated(opts *bind.FilterOpts) (*PairImplTickSizeUpdatedIterator, error) {

	logs, sub, err := _PairImpl.contract.FilterLogs(opts, "TickSizeUpdated")
	if err != nil {
		return nil, err
	}
	return &PairImplTickSizeUpdatedIterator{contract: _PairImpl.contract, event: "TickSizeUpdated", logs: logs, sub: sub}, nil
}

// WatchTickSizeUpdated is a free log subscription operation binding the contract event 0x66748bc112f6692a76bebecc33e69768508e27e414e3080f5c02295b6ea3cf3f.
//
// Solidity: event TickSizeUpdated(uint256 beforeLotSize, uint256 newLotSize, uint256 beforeTickSize, uint256 newTickSize)
func (_PairImpl *PairImplFilterer) WatchTickSizeUpdated(opts *bind.WatchOpts, sink chan<- *PairImplTickSizeUpdated) (event.Subscription, error) {

	logs, sub, err := _PairImpl.contract.WatchLogs(opts, "TickSizeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PairImplTickSizeUpdated)
				if err := _PairImpl.contract.UnpackLog(event, "TickSizeUpdated", log); err != nil {
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
func (_PairImpl *PairImplFilterer) ParseTickSizeUpdated(log types.Log) (*PairImplTickSizeUpdated, error) {
	event := new(PairImplTickSizeUpdated)
	if err := _PairImpl.contract.UnpackLog(event, "TickSizeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PairImplUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the PairImpl contract.
type PairImplUnpausedIterator struct {
	Event *PairImplUnpaused // Event containing the contract specifics and raw log

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
func (it *PairImplUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PairImplUnpaused)
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
		it.Event = new(PairImplUnpaused)
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
func (it *PairImplUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PairImplUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PairImplUnpaused represents a Unpaused event raised by the PairImpl contract.
type PairImplUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PairImpl *PairImplFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PairImplUnpausedIterator, error) {

	logs, sub, err := _PairImpl.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PairImplUnpausedIterator{contract: _PairImpl.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PairImpl *PairImplFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PairImplUnpaused) (event.Subscription, error) {

	logs, sub, err := _PairImpl.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PairImplUnpaused)
				if err := _PairImpl.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_PairImpl *PairImplFilterer) ParseUnpaused(log types.Log) (*PairImplUnpaused, error) {
	event := new(PairImplUnpaused)
	if err := _PairImpl.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PairImplUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the PairImpl contract.
type PairImplUpgradedIterator struct {
	Event *PairImplUpgraded // Event containing the contract specifics and raw log

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
func (it *PairImplUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PairImplUpgraded)
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
		it.Event = new(PairImplUpgraded)
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
func (it *PairImplUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PairImplUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PairImplUpgraded represents a Upgraded event raised by the PairImpl contract.
type PairImplUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_PairImpl *PairImplFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*PairImplUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _PairImpl.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &PairImplUpgradedIterator{contract: _PairImpl.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_PairImpl *PairImplFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *PairImplUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _PairImpl.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PairImplUpgraded)
				if err := _PairImpl.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_PairImpl *PairImplFilterer) ParseUpgraded(log types.Log) (*PairImplUpgraded, error) {
	event := new(PairImplUpgraded)
	if err := _PairImpl.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
