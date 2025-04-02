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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BASE\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DENOMINATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MARKET\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"QUOTE\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROUTER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseReserve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"orderIds\",\"type\":\"uint256[]\"}],\"name\":\"cancelOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"orderIds\",\"type\":\"uint256[]\"}],\"name\":\"emergencyCancelOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIPair.OrderSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256[2]\",\"name\":\"adjacent\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"findMaxCount\",\"type\":\"uint256\"}],\"name\":\"findPrevPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"QUOTE\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"BASE\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"DENOMINATOR\",\"type\":\"uint256\"}],\"internalType\":\"structIPair.Config\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quote\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tickSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lotSize\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lotSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"matchedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"matchedPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minTradeVolume\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"orderById\",\"outputs\":[{\"components\":[{\"internalType\":\"enumIPair.OrderSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"feeBps\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIPair.Order\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIPair.OrderSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"name\":\"ordersByPrices\",\"outputs\":[{\"internalType\":\"uint256[][]\",\"name\":\"\",\"type\":\"uint256[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quoteReserve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"setPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lotSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tickSize\",\"type\":\"uint256\"}],\"name\":\"setTickSize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"erc20\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"skim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumIPair.OrderSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"feeBps\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIPair.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"enumIPair.LimitConstraints\",\"name\":\"constraints\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"prevPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxMatchCount\",\"type\":\"uint256\"}],\"name\":\"submitLimitOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumIPair.OrderSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"feeBps\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIPair.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"spendAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxMatchCount\",\"type\":\"uint256\"}],\"name\":\"submitMarketOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tickSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tickSizes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tick\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lot\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ticks\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"sellPrices\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"buyPrices\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeBps\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"FeeCollect\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"enumIPair.CloseType\",\"name\":\"closeType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"OrderClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"enumIPair.OrderSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"OrderCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"sellId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"buyId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"OrderMatched\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"erc20\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Skim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"beforeLotSize\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newLotSize\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"beforeTickSize\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTickSize\",\"type\":\"uint256\"}],\"name\":\"TickSizeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ListFailToFindPrev\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ListInvalidFindMaxCount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ListInvalidIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ListInvalidPrevNode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ListZeroData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"PairFillOrKill\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PairInsufficientTradeVolume\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PairInvalidAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"PairInvalidInitializeData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PairInvalidOrderId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumIPair.OrderSide\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"PairInvalidOrderSide\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumIPair.OrderSide\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PairInvalidPrevPrice\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PairInvalidPrice\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"PairInvalidReserve\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"PairInvalidRouter\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PairInvalidTickSize\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"PairNotOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumIPair.OrderSide\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PairUnknownPrices\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"ec342ad0": "BASE()",
		"918f8674": "DENOMINATOR()",
		"f46f16c2": "MARKET()",
		"9c579839": "QUOTE()",
		"32fe7b26": "ROUTER()",
		"ad3cb1cc": "UPGRADE_INTERFACE_VERSION()",
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
	Bin: "0x60a060405230608052348015610013575f5ffd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b608051614ba36100f95f395f8181612cc701528181612cf00152612f440152614ba35ff3fe6080604052600436106101d0575f3560e01c80638da5cb5b116100fd578063cc50e36a11610092578063ec342ad011610062578063ec342ad0146105e5578063f210d08714610611578063f46f16c214610626578063fe56634914610651575f5ffd5b8063cc50e36a1461056b578063d6e9cc2a1461058a578063dfdf2a72146105b1578063e077dd8a146105c6575f5ffd5b8063a6b63eb8116100cd578063a6b63eb81461045a578063ad3cb1cc14610479578063bedb86fb146104ce578063c3f909d4146104ed575f5ffd5b80638da5cb5b146103f0578063918f8674146104045780639c579839146104195780639da771f414610445575f5ffd5b80634942f65f116101735780635c975abb116101435780635c975abb1461034f5780636c4cfbc8146103905780637cbf6db2146103af5780638d77eba5146103c4575f5ffd5b80634942f65f146102fe5780634d6754f1146103135780634f1ef2861461032857806352d1902d1461033b575f5ffd5b80631edd8614116101ae5780631edd8614146102495780632cfffaf61461026c5780632dd6a00a1461028e57806332fe7b26146102ad575f5ffd5b806312a808cc146101d45780631a0361fb146102095780631ec482d71461022a575b5f5ffd5b3480156101df575f5ffd5b506101f36101ee3660046142ac565b610670565b60405161020091906142f7565b60405180910390f35b348015610214575f5ffd5b506102286102233660046143c2565b610756565b005b348015610235575f5ffd5b50610228610244366004614400565b610ac1565b348015610254575f5ffd5b5061025e60075481565b604051908152602001610200565b348015610277575f5ffd5b50610280610cd1565b604051610200929190614456565b348015610299575f5ffd5b5061025e6102a8366004614510565b610d0b565b3480156102b8575f5ffd5b506001546102d99073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610200565b348015610309575f5ffd5b5061025e600a5481565b34801561031e575f5ffd5b5061025e60085481565b610228610336366004614556565b611365565b348015610346575f5ffd5b5061025e611384565b34801561035a575f5ffd5b507fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff166040519015158152602001610200565b34801561039b575f5ffd5b5061025e6103aa36600461461b565b6113b2565b3480156103ba575f5ffd5b5061025e600b5481565b3480156103cf575f5ffd5b506103e36103de366004614663565b611494565b60405161020091906146e0565b3480156103fb575f5ffd5b506102d9611563565b34801561040f575f5ffd5b5061025e60045481565b348015610424575f5ffd5b506003546102d99073ffffffffffffffffffffffffffffffffffffffff1681565b348015610450575f5ffd5b5061025e60065481565b348015610465575f5ffd5b5061022861047436600461473c565b6115f6565b348015610484575f5ffd5b506104c16040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b6040516102009190614793565b3480156104d9575f5ffd5b506102286104e83660046147e6565b611ae5565b3480156104f8575f5ffd5b5060408051606080820183525f8083526020808401829052928401528251808201845260035473ffffffffffffffffffffffffffffffffffffffff908116808352600254821683860190815260045493870193845286519182525190911693810193909352519282019290925201610200565b348015610576575f5ffd5b50610228610585366004614805565b611b3e565b348015610595575f5ffd5b50600954600a5460408051928352602083019190915201610200565b3480156105bc575f5ffd5b5061025e60055481565b3480156105d1575f5ffd5b506102286105e0366004614836565b611d12565b3480156105f0575f5ffd5b506002546102d99073ffffffffffffffffffffffffffffffffffffffff1681565b34801561061c575f5ffd5b5061025e60095481565b348015610631575f5ffd5b505f546102d99073ffffffffffffffffffffffffffffffffffffffff1681565b34801561065c575f5ffd5b5061022861066b366004614868565b611e5a565b60605f808460018111156106865761068661467a565b14610692576016610695565b60155b83519091505f8167ffffffffffffffff8111156106b4576106b46141b0565b6040519080825280602002602001820160405280156106e757816020015b60608152602001906001900390816106d25790505b5090505f5b8281101561074a57610725845f88848151811061070b5761070b614888565b602002602001015181526020019081526020015f20612072565b82828151811061073757610737614888565b60209081029190910101526001016106ec565b50925050505b92915050565b61075e611563565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146107e257335b6040517f118cdaa700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911660048201526024015b60405180910390fd5b8015610abc5760025473ffffffffffffffffffffffffffffffffffffffff84811691161480156108ac57506005546002546040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152839173ffffffffffffffffffffffffffffffffffffffff16906370a0823190602401602060405180830381865afa15801561087c573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108a091906148b5565b6108aa91906148f9565b105b156108ff576002546040517f7bdf264d00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911660048201526024016107d9565b60035473ffffffffffffffffffffffffffffffffffffffff84811691161480156109c357506006546003546040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152839173ffffffffffffffffffffffffffffffffffffffff16906370a0823190602401602060405180830381865afa158015610993573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109b791906148b5565b6109c191906148f9565b105b15610a16576003546040517f7bdf264d00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911660048201526024016107d9565b610a3773ffffffffffffffffffffffffffffffffffffffff84168383612118565b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16610a6c3390565b73ffffffffffffffffffffffffffffffffffffffff167f7e30e394efdaabad9f9ca6fd8f67f9449ba63d674ee8068283c3ae1f49b5b66084604051610ab391815260200190565b60405180910390a45b505050565b60015473ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610b4357335b6040517f143db19100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911660048201526024016107d9565b80515f5b81811015610ccb575f838281518110610b6257610b62614888565b6020908102919091018101515f81815260179092526040808320815160a0810190925280549294509091829060ff166001811115610ba257610ba261467a565b6001811115610bb357610bb361467a565b81528154610100810473ffffffffffffffffffffffffffffffffffffffff908116602080850191909152750100000000000000000000000000000000000000000090920463ffffffff1660408401526001840154606084015260029093015460809092019190915282015191925016610c2d575050610b47565b8573ffffffffffffffffffffffffffffffffffffffff16816020015173ffffffffffffffffffffffffffffffffffffffff1614610cb5576040517f54de53f70000000000000000000000000000000000000000000000000000000081526004810183905273ffffffffffffffffffffffffffffffffffffffff871660248201526044016107d9565b610cc1828260026121a5565b5050600101610b47565b50505050565b606080610cf7600d5f5b60ff1660028110610cee57610cee614888565b60040201612072565b9150610d05600d6001610cdb565b90509091565b5f610d14612301565b60015473ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610d4f5733610af8565b610d5761235f565b60608501511580610d7857506009548560600151610d759190614939565b15155b15610db75784606001516040517f4e1404fe0000000000000000000000000000000000000000000000000000000081526004016107d991815260200190565b60808501511580610dd85750600a548560800151610dd59190614939565b15155b15610e175784608001516040517fa334626e0000000000000000000000000000000000000000000000000000000081526004016107d991815260200190565b600c5f8154610e2590614971565b918290555090505f8086516001811115610e4157610e4161467a565b1490505f81610e5b57610e5683885f876124ae565b610e66565b610e66838886612709565b905086608001515f03610eb3575f837f78615f01112d8477ae8d6fc9a510bcfdf935bc3ab7b02d0ff2316f5b44bd597a42604051610ea691815260200190565b60405180910390a3611348565b6001866002811115610ec757610ec761467a565b03610f42576001837f78615f01112d8477ae8d6fc9a510bcfdf935bc3ab7b02d0ff2316f5b44bd597a42604051610f0091815260200190565b60405180910390a38115610f3d5760208701516080880151600254610f3d9273ffffffffffffffffffffffffffffffffffffffff90911691612118565b611348565b6002866002811115610f5657610f5661467a565b03610fab5760208701516040517f1b3c356200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911660048201526024016107d9565b8115611193578415801590610fc35750848760600151105b156110035760608701516040517fc18aa6060000000000000000000000000000000000000000000000000000000081526107d9915f918890600401614989565b866080015160055f82825461101891906149a8565b90915550507f1d2ff3fa6980aeeebca4e94965520da48983006e9b1115c1c853cbb10d943d005c63ffffffff166040808901919091525f848152601760205220875181548992919082907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600183818111156110975761109761467a565b021790555060208201518154604084015163ffffffff167501000000000000000000000000000000000000000000027fffffffffffffff00000000ffffffffffffffffffffffffffffffffffffffffff73ffffffffffffffffffffffffffffffffffffffff90931661010002929092167fffffffffffffff000000000000000000000000000000000000000000000000ff9091161717815560608083015160018301556080909201516002909101558701516111709086600d5f5b60ff166002811061116557611165614888565b600402019190612966565b5060608701515f90815260156020526040902061118d9084612b14565b50611348565b84158015906111a55750848760600151115b156111e65760608701516040517fc18aa6060000000000000000000000000000000000000000000000000000000081526107d9916001918890600401614989565b6111fb87606001518860800151600454612b24565b60065f82825461120b91906149a8565b90915550505f604080890182905284825260176020529020875181548992919082907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600183818111156112625761126261467a565b021790555060208201518154604084015163ffffffff167501000000000000000000000000000000000000000000027fffffffffffffff00000000ffffffffffffffffffffffffffffffffffffffffff73ffffffffffffffffffffffffffffffffffffffff90931661010002929092167fffffffffffffff000000000000000000000000000000000000000000000000ff90911617178155606080830151600180840191909155608090930151600290920191909155880151611329918790600d90611152565b5060608701515f9081526016602052604090206113469084612b14565b505b8161135b5761135b876020015182612bda565b5050949350505050565b61136d612caf565b61137682612db3565b6113808282612df3565b5050565b5f61138d612f2c565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b5f808560018111156113c6576113c661467a565b0361142f5760408051808201825261142891869190869060029083908390808284375f92019190915250869150600d90508960018111156114095761140961467a565b60ff166002811061141c5761141c614888565b60040201929190612f9b565b905061148c565b60408051808201825261142891869190869060029083908390808284375f92019190915250869150600d905089600181111561146d5761146d61467a565b60ff166002811061148057611480614888565b60040201929190613141565b949350505050565b6040805160a0810182525f808252602082018190529181018290526060810182905260808101919091525f8281526017602052604090819020815160a081019092528054829060ff1660018111156114ee576114ee61467a565b60018111156114ff576114ff61467a565b81528154610100810473ffffffffffffffffffffffffffffffffffffffff1660208301527501000000000000000000000000000000000000000000900463ffffffff1660408201526001820154606082015260029091015460809091015292915050565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156115cd573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906115f191906149bb565b905090565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff165f811580156116405750825b90505f8267ffffffffffffffff16600114801561165c5750303b155b90508115801561166a575080155b156116a1576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000016600117855583156117025784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b73ffffffffffffffffffffffffffffffffffffffff8a16611771576040517fc587916e0000000000000000000000000000000000000000000000000000000081527f726f75746572000000000000000000000000000000000000000000000000000060048201526024016107d9565b73ffffffffffffffffffffffffffffffffffffffff89166117e0576040517fc587916e0000000000000000000000000000000000000000000000000000000081527f71756f746500000000000000000000000000000000000000000000000000000060048201526024016107d9565b73ffffffffffffffffffffffffffffffffffffffff881661184f576040517fc587916e0000000000000000000000000000000000000000000000000000000081527f626173650000000000000000000000000000000000000000000000000000000060048201526024016107d9565b865f036118aa576040517fc587916e0000000000000000000000000000000000000000000000000000000081527f7469636b53697a6500000000000000000000000000000000000000000000000060048201526024016107d9565b855f03611905576040517fc587916e0000000000000000000000000000000000000000000000000000000081527f6c6f7453697a650000000000000000000000000000000000000000000000000060048201526024016107d9565b335f80547fffffffffffffffffffffffff000000000000000000000000000000000000000090811673ffffffffffffffffffffffffffffffffffffffff938416179091556001805482168d84161790556003805482168c841617905560028054909116918a169182179055604080517f313ce567000000000000000000000000000000000000000000000000000000008152905163313ce567916004808201926020929091908290030181865afa1580156119c2573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906119e691906149d6565b6119f190600a614ad9565b6004819055611a008789614ae7565b611a0a9190614939565b15611a5357600480546040517f575681350000000000000000000000000000000000000000000000000000000081529182018990526024820188905260448201526064016107d9565b6009879055600a869055600454611a6d9088908890612b24565b600b55611a786132b0565b8315611ad95784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050505050565b611aed611563565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611b255733610792565b8015611b3657611b336132c0565b50565b611b33613360565b611b46612301565b60015473ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611b815733610af8565b611b8961235f565b5f600c5f8154611b9890614971565b918290555090505f84516001811115611bb357611bb361467a565b03611c6157821580611bd05750600a54611bcd9084614939565b15155b15611c0a576040517fa334626e000000000000000000000000000000000000000000000000000000008152600481018490526024016107d9565b5f606085015260808401839052611c22818584612709565b50608084015115611c5c5760208401516080850151600254611c5c9273ffffffffffffffffffffffffffffffffffffffff90911691612118565b611cd1565b600b54831015611cab57600b546040517f0bc1e7dd0000000000000000000000000000000000000000000000000000000081526107d9918591600401918252602082015260400190565b5f1960608501525f611cbf828686866124ae565b9050611ccf856020015182612bda565b505b5f817f78615f01112d8477ae8d6fc9a510bcfdf935bc3ab7b02d0ff2316f5b44bd597a42604051611d0491815260200190565b60405180910390a350505050565b611d1a6133d6565b611d22611563565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611d5a5733610792565b80515f5b81811015610abc575f838281518110611d7957611d79614888565b6020908102919091018101515f81815260179092526040808320815160a0810190925280549294509091829060ff166001811115611db957611db961467a565b6001811115611dca57611dca61467a565b81528154610100810473ffffffffffffffffffffffffffffffffffffffff908116602080850191909152750100000000000000000000000000000000000000000090920463ffffffff1660408401526001840154606084015260029093015460809092019190915282015191925016611e44575050611d5e565b611e50828260036121a5565b5050600101611d5e565b5f5473ffffffffffffffffffffffffffffffffffffffff1663a1eea778336040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815273ffffffffffffffffffffffffffffffffffffffff90911660048201526024015f6040518083038186803b158015611eda575f5ffd5b505afa158015611eec573d5f5f3e3d5ffd5b50505050805f03611f4b576040517fc587916e0000000000000000000000000000000000000000000000000000000081527f7469636b53697a6500000000000000000000000000000000000000000000000060048201526024016107d9565b815f03611fa6576040517fc587916e0000000000000000000000000000000000000000000000000000000081527f6c6f7453697a650000000000000000000000000000000000000000000000000060048201526024016107d9565b600454611fb38383614ae7565b611fbd9190614939565b1561200657600480546040517f575681350000000000000000000000000000000000000000000000000000000081529182018390526024820184905260448201526064016107d9565b600a546009546040805192835260208301859052820152606081018290527f66748bc112f6692a76bebecc33e69768508e27e414e3080f5c02295b6ea3cf3f9060800160405180910390a1600a829055600981905560045461206b9082908490612b24565b600b555050565b60605f825f015467ffffffffffffffff811115612091576120916141b0565b6040519080825280602002602001820160405280156120ba578160200160208202803683370190505b5083546001850154919250905f5b8281101561210e57818482815181106120e3576120e3614888565b6020908102919091018101919091525f928352600387019052604090912060019081015491016120c8565b5091949350505050565b6040805173ffffffffffffffffffffffffffffffffffffffff8416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb00000000000000000000000000000000000000000000000000000000179052610abc908490613431565b5f8080845160018111156121bb576121bb61467a565b14905080156122325760608401515f90815260156020526040902060808501519092501561222d57602084015160808501516002546122129273ffffffffffffffffffffffffffffffffffffffff90911691612118565b836080015160055f82825461222791906148f9565b90915550505b6122ab565b60165f856060015181526020019081526020015f2091505f61225f85606001518660800151600454612b24565b905080156122a95760208501516003546122929173ffffffffffffffffffffffffffffffffffffffff9091169083612118565b8060065f8282546122a391906148f9565b90915550505b505b6122b68584846134d0565b81546122fa576122f88460600151600d865f015160018111156122db576122db61467a565b60ff16600281106122ee576122ee614888565b60040201906135a2565b505b5050505050565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff161561235d576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505f5f8273ffffffffffffffffffffffffffffffffffffffff1663c415b95c6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156123ce573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906123f291906149bb565b8373ffffffffffffffffffffffffffffffffffffffff166324a9d8536040518163ffffffff1660e01b8152600401602060405180830381865afa15801561243b573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061245f9190614afe565b91509150817fd6aa07baf8485abf9d26fecf4c935d75b50a73e678db02b944bd3ac8759823005d807f1d2ff3fa6980aeeebca4e94965520da48983006e9b1115c1c853cbb10d943d005d505050565b5f6001845160018111156124c4576124c461467a565b146124fe5783516040517f7cbdd0810000000000000000000000000000000000000000000000000000000081526107d99190600401614b19565b5f831561250c575082612524565b61252185606001518660800151600454612b24565b90505b6003546040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201525f9182916125cd9173ffffffffffffffffffffffffffffffffffffffff16906370a0823190602401602060405180830381865afa158015612596573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906125ba91906148b5565b846006546125c891906149a8565b613673565b9150915081612624576003546040517f7bdf264d00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911660048201526024016107d9565b865160018111156126375761263761467a565b88886020015173ffffffffffffffffffffffffffffffffffffffff167f7cde7559ef00f8c3d9244d3fde19e7500979d64f6f9c086a3836a4926abbe1658a606001518b608001514260405161269f939291909283526020830191909152604082015260600190565b60405180910390a45f6126b489898989613697565b905080156126fd578060055f8282546126cd91906148f9565b909155505060208801516002546126fd9173ffffffffffffffffffffffffffffffffffffffff9091169083612118565b50979650505050505050565b5f808351600181111561271e5761271e61467a565b146127585782516040517f7cbdd0810000000000000000000000000000000000000000000000000000000081526107d99190600401614b19565b826080015160055461276a91906149a8565b6002546040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff909116906370a0823190602401602060405180830381865afa1580156127d6573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906127fa91906148b5565b101561284e576002546040517f7bdf264d00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911660048201526024016107d9565b825160018111156128615761286161467a565b84846020015173ffffffffffffffffffffffffffffffffffffffff167f7cde7559ef00f8c3d9244d3fde19e7500979d64f6f9c086a3836a4926abbe16586606001518760800151426040516128c9939291909283526020830191909152604082015260600190565b60405180910390a45f6128dd8585856138a8565b9050801561295a57600680547fd6aa07baf8485abf9d26fecf4c935d75b50a73e678db02b944bd3ac8759823005c917f1d2ff3fa6980aeeebca4e94965520da48983006e9b1115c1c853cbb10d943d005c918491905f9061293f9084906148f9565b92505081905550612957878760200151858585613a63565b50505b5f9150505b9392505050565b5f825f036129a0576040517f0da93b9700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6129aa8484613b7a565b156129b657505f61295f565b5f828152600385016020908152604080832081518083019092528382529181019290925290835f03612a2257600186018054908690558015612a05575f81815260038801602052604090208690555b60405180604001604052805f815260200182815250915050612ac4565b604080518082019091528254815260018301546020820152612a4390613bc2565b80612a515750856001015484145b612a87576040517f182dfca500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5060018101546040805180820190915284815260208101829052908015612abb575f81815260038801602052604090208690555b50600182018590555b80602001515f03612ad757600286018590555b5f8581526003870160209081526040822083518155908301516001909101558654879190612b0490614971565b9091555060019695505050505050565b5f61295f83838560020154612966565b5f838302815f1985870982811083820303915050805f03612b5857838281612b4e57612b4e61490c565b049250505061295f565b808411612b6f57612b6f6003851502601118613bd8565b5f848688095f868103871696879004966002600389028118808a02820302808a02820302808a02820302808a02820302808a02820302808a02909103029181900381900460010186841190950394909402919094039290920491909117919091029150509392505050565b5f60065482612be991906149a8565b6003546040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff909116906370a0823190602401602060405180830381865afa158015612c55573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612c7991906148b5565b612c8391906148f9565b90508015610abc57600354610abc9073ffffffffffffffffffffffffffffffffffffffff168483612118565b3073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161480612d7c57507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16612d637f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614155b1561235d576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612dbb611563565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611b335733610792565b8173ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015612e78575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201909252612e75918101906148b5565b60015b612ec6576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff831660048201526024016107d9565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8114612f22576040517faa1d49a4000000000000000000000000000000000000000000000000000000008152600481018290526024016107d9565b610abc8383613be9565b3073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161461235d576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f835f03612fd5576040517f0da93b9700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612fdf8585613b7a565b15612ffa57505f83815260038501602052604090205461148c565b846001015484108061300b57508454155b1561301757505f61148c565b846002015484111561302e5750600284015461148c565b815f03613067576040517fab7e639300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6130728685613c4b565b9050848110156130d2575b801580159061309157505f19909201918215155b156130cd575f908152600386016020526040902060010154848111156130c8575f908152600386016020526040902054905061148c565b61307d565b61310f565b80158015906130e657505f19909201918215155b1561310f575f9081526003860160205260409020548481101561310a57905061148c565b6130d2565b6040517f80ce60d100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f835f0361317b576040517f0da93b9700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6131858585613b7a565b156131a057505f83815260038501602052604090205461148c565b84600101548411806131b157508454155b156131bd57505f61148c565b84600201548410156131d45750600284015461148c565b815f0361320d576040517fab7e639300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6132188685613c4b565b905084811115613273575b801580159061323757505f19909201918215155b156130cd575f9081526003860160205260409020600101548085111561326e575f908152600386016020526040902054905061148c565b613223565b801580159061328757505f19909201918215155b1561310f575f908152600386016020526040902054808510156132ab57905061148c565b613273565b6132b8613c95565b61235d613cfc565b6132c8612301565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a150565b6133686133d6565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa33613335565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff1661235d576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5f60205f8451602086015f885af180613450576040513d5f823e3d81fd5b50505f513d91508115613467578060011415613481565b73ffffffffffffffffffffffffffffffffffffffff84163b155b15610ccb576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024016107d9565b6134da81846135a2565b613513576040517ffcfdf902000000000000000000000000000000000000000000000000000000008152600481018490526024016107d9565b5f83815260176020526040812080547fffffffffffffff0000000000000000000000000000000000000000000000000016815560018101829055600201558160038111156135635761356361467a565b837f78615f01112d8477ae8d6fc9a510bcfdf935bc3ab7b02d0ff2316f5b44bd597a4260405161359591815260200190565b60405180910390a3505050565b5f815f036135dc576040517f0da93b9700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6135e68383613b7a565b6135f157505f610750565b5f82815260038401602052604080822080548352818320600180830154808652939094208482019390935581548355928601549092919085900361363a57600180840154908701555b8486600201540361364d57825460028701555b5f85815260038701602052604081208181556001018190558654879190612b0490614b27565b5f5f8383111561368757505f905080613690565b50600190508183035b9250929050565b5f80600d5b80541561389d575f6136ae8282613d4d565b905086606001518111156136c2575061389d565b5f8181526015602052604090206136d882613dbb565b861561372c575f6136f56136ec868a6148f9565b60045485612b24565b9050600a54816137059190614939565b61370f90826148f9565b60808a0181905290505f81900361372a5750505050506138a0565b505b5f613735825490565b90505b8015613851575f6137498382613d4d565b5f81815260176020526040812091925080806137698f8f87878c8c613de1565b9250925092505f61377d8984600454612b24565b90506137ad8685837fd6aa07baf8485abf9d26fecf4c935d75b50a73e678db02b944bd3ac8759823005c86613a63565b6137b7838d6149a8565b9b506137c3818c6149a8565b9a508e608001515f14806137e057506137db8d614b27565b9c508c155b1561383f57875461382f576137f58a8a6135a2565b61382f575f896040517f9f1cfdfe0000000000000000000000000000000000000000000000000000000081526004016107d9929190614b3c565b50505050505050505050506138a0565b86600190039650505050505050613738565b61385b84846135a2565b613895575f836040517f9f1cfdfe0000000000000000000000000000000000000000000000000000000081526004016107d9929190614b3c565b50505061369c565b50505b61148c613f07565b6002546004545f9173ffffffffffffffffffffffffffffffffffffffff169060115b805415613a57575f6138dc8282613d4d565b905086606001518110156138f05750613a57565b5f81815260166020526040902061390682613dbb565b5f61390f825490565b90505b8015613a0a575f6139238382613d4d565b5f818152601760205260408120919250806139428e8e86868b8b613de1565b50909250905061396973ffffffffffffffffffffffffffffffffffffffff8b168383612118565b61397487828b612b24565b61397e908c6149a8565b9a508c608001515f148061399b57506139968c614b27565b9b508b155b156139fa5785546139eb576139b088886135a2565b6139eb576001876040517f9f1cfdfe0000000000000000000000000000000000000000000000000000000081526004016107d9929190614b3c565b50505050505050505050613a5b565b8460019003945050505050613912565b613a1484846135a2565b613a4f576001836040517f9f1cfdfe0000000000000000000000000000000000000000000000000000000081526004016107d9929190614b3c565b5050506138ca565b5050505b61295f613f07565b8063ffffffff165f03613a9957600354613a949073ffffffffffffffffffffffffffffffffffffffff168585612118565b6122fa565b5f613aad848363ffffffff16612710612b24565b90505f613aba82866148f9565b6040805187815263ffffffff861660208201529081018490526060810182905290915073ffffffffffffffffffffffffffffffffffffffff808616919088169089907f9814b2b2f43f9ea51ad7e8760fda094ffa3012bfeb10c76a9656389af18956039060800160405180910390a4600354613b4d9073ffffffffffffffffffffffffffffffffffffffff168584612118565b600354613b719073ffffffffffffffffffffffffffffffffffffffff168783612118565b50505050505050565b5f815f03613b8957505f610750565b826001015482148061295f57505f828152600384016020908152604091829020825180840190935280548352600101549082015261295f905b80515f9015158061075057505060200151151590565b634e487b715f52806020526024601cfd5b613bf282613f38565b60405173ffffffffffffffffffffffffffffffffffffffff8316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115613c4357610abc8282614006565b611380614085565b5f613c5d8383835b6020020151613b7a565b15613c7157815f5b60200201519050610750565b613c7d83836001613c53565b15613c8a57816001613c65565b506001820154610750565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff1661235d576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b613d04613c95565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b81545f908210613d89576040517f39e60f7300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001830154825b8015613db3575f91825260038501602052604090912060010154905f1901613d90565b509392505050565b807ffd0e5d4f9b88892d3b04349a0e2bc0d1359414c21932fcd7d5a523a6c0a5cd005d50565b5f5f5f855f0160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16613e1a896080015188600201546140bd565b875491945092507501000000000000000000000000000000000000000000900463ffffffff1690505f80808a516001811115613e5857613e5861467a565b14613e6457888b613e67565b8a895b915091508681837f6a6896b1e6131e0b8ebae43fdc84cb0178a6eacdf4ee63f15dabae48729a8a038742604051613ea8929190918252602082015260400190565b60405180910390a487600201548403613ecb57613ec6895f886134d0565b613ed7565b60028801805485900390555b89608001518403613eed575f60808b0152613ef9565b60808a01805185900390525b505096509650969350505050565b7ffd0e5d4f9b88892d3b04349a0e2bc0d1359414c21932fcd7d5a523a6c0a5cd005c8015611b335760075542600855565b8073ffffffffffffffffffffffffffffffffffffffff163b5f03613fa0576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024016107d9565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60605f5f8473ffffffffffffffffffffffffffffffffffffffff168460405161402f9190614b57565b5f60405180830381855af49150503d805f8114614067576040519150601f19603f3d011682016040523d82523d5f602084013e61406c565b606091505b509150915061407c8583836140cc565b95945050505050565b341561235d576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f82821882841002821861295f565b6060826140e1576140dc8261415b565b61295f565b8151158015614105575073ffffffffffffffffffffffffffffffffffffffff84163b155b15614154576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024016107d9565b508061295f565b80511561416b5780518082602001fd5b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8035600281106141ab575f5ffd5b919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715614224576142246141b0565b604052919050565b5f82601f83011261423b575f5ffd5b813567ffffffffffffffff811115614255576142556141b0565b8060051b614265602082016141dd565b91825260208185018101929081019086841115614280575f5ffd5b6020860192505b838310156142a2578235825260209283019290910190614287565b9695505050505050565b5f5f604083850312156142bd575f5ffd5b6142c68361419d565b9150602083013567ffffffffffffffff8111156142e1575f5ffd5b6142ed8582860161422c565b9150509250929050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b82811015614395578685037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0018452815180518087526020918201918701905f5b8181101561437c57835183526020938401939092019160010161435e565b509096505050602093840193919091019060010161431d565b50929695505050505050565b73ffffffffffffffffffffffffffffffffffffffff81168114611b33575f5ffd5b5f5f5f606084860312156143d4575f5ffd5b83356143df816143a1565b925060208401356143ef816143a1565b929592945050506040919091013590565b5f5f60408385031215614411575f5ffd5b82356142c6816143a1565b5f8151808452602084019350602083015f5b8281101561444c57815186526020958601959091019060010161442e565b5093949350505050565b604081525f614468604083018561441c565b828103602084015261407c818561441c565b63ffffffff81168114611b33575f5ffd5b5f60a0828403121561449b575f5ffd5b60405160a0810167ffffffffffffffff811182821017156144be576144be6141b0565b6040529050806144cd8361419d565b815260208301356144dd816143a1565b602082015260408301356144f08161447a565b604082015260608381013590820152608092830135920191909152919050565b5f5f5f5f6101008587031215614524575f5ffd5b61452e868661448b565b935060a085013560038110614541575f5ffd5b939693955050505060c08201359160e0013590565b5f5f60408385031215614567575f5ffd5b8235614572816143a1565b9150602083013567ffffffffffffffff81111561458d575f5ffd5b8301601f8101851361459d575f5ffd5b803567ffffffffffffffff8111156145b7576145b76141b0565b6145e860207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116016141dd565b8181528660208385010111156145fc575f5ffd5b816020840160208301375f602083830101528093505050509250929050565b5f5f5f5f60a0858703121561462e575f5ffd5b6146378561419d565b9350602085013592506080850186811115614650575f5ffd5b9396929550505060409290920191903590565b5f60208284031215614673575f5ffd5b5035919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b600281106146dc577f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b9052565b5f60a0820190506146f28284516146a7565b73ffffffffffffffffffffffffffffffffffffffff602084015116602083015263ffffffff6040840151166040830152606083015160608301526080830151608083015292915050565b5f5f5f5f5f60a08688031215614750575f5ffd5b853561475b816143a1565b9450602086013561476b816143a1565b9350604086013561477b816143a1565b94979396509394606081013594506080013592915050565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b5f602082840312156147f6575f5ffd5b8135801515811461295f575f5ffd5b5f5f5f60e08486031215614817575f5ffd5b614821858561448b565b9560a0850135955060c0909401359392505050565b5f60208284031215614846575f5ffd5b813567ffffffffffffffff81111561485c575f5ffd5b61148c8482850161422c565b5f5f60408385031215614879575f5ffd5b50508035926020909101359150565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f602082840312156148c5575f5ffd5b5051919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b81810381811115610750576107506148cc565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f8261496c577f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b500690565b5f5f198203614982576149826148cc565b5060010190565b6060810161499782866146a7565b602082019390935260400152919050565b80820180821115610750576107506148cc565b5f602082840312156149cb575f5ffd5b815161295f816143a1565b5f602082840312156149e6575f5ffd5b815160ff8116811461295f575f5ffd5b6001815b6001841115614a3157808504811115614a1557614a156148cc565b6001841615614a2357908102905b60019390931c9280026149fa565b935093915050565b5f82614a4757506001610750565b81614a5357505f610750565b8160018114614a695760028114614a7357614a8f565b6001915050610750565b60ff841115614a8457614a846148cc565b50506001821b610750565b5060208310610133831016604e8410600b8410161715614ab2575081810a610750565b614abe5f1984846149f6565b805f1904821115614ad157614ad16148cc565b029392505050565b5f61295f60ff841683614a39565b8082028115828204841417610750576107506148cc565b5f60208284031215614b0e575f5ffd5b815161295f8161447a565b6020810161075082846146a7565b5f81614b3557614b356148cc565b505f190190565b60408101614b4a82856146a7565b8260208301529392505050565b5f82518060208501845e5f92019182525091905056fea264697066735822122037d6f8b4fb3aa5e19c34cb759a4f19140f3bc95caaf780a8ff85c8291048755564736f6c634300081c0033",
}

// PairImplABI is the input ABI used to generate the binding from.
// Deprecated: Use PairImplMetaData.ABI instead.
var PairImplABI = PairImplMetaData.ABI

// PairImplBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const PairImplBinRuntime = "6080604052600436106101d0575f3560e01c80638da5cb5b116100fd578063cc50e36a11610092578063ec342ad011610062578063ec342ad0146105e5578063f210d08714610611578063f46f16c214610626578063fe56634914610651575f5ffd5b8063cc50e36a1461056b578063d6e9cc2a1461058a578063dfdf2a72146105b1578063e077dd8a146105c6575f5ffd5b8063a6b63eb8116100cd578063a6b63eb81461045a578063ad3cb1cc14610479578063bedb86fb146104ce578063c3f909d4146104ed575f5ffd5b80638da5cb5b146103f0578063918f8674146104045780639c579839146104195780639da771f414610445575f5ffd5b80634942f65f116101735780635c975abb116101435780635c975abb1461034f5780636c4cfbc8146103905780637cbf6db2146103af5780638d77eba5146103c4575f5ffd5b80634942f65f146102fe5780634d6754f1146103135780634f1ef2861461032857806352d1902d1461033b575f5ffd5b80631edd8614116101ae5780631edd8614146102495780632cfffaf61461026c5780632dd6a00a1461028e57806332fe7b26146102ad575f5ffd5b806312a808cc146101d45780631a0361fb146102095780631ec482d71461022a575b5f5ffd5b3480156101df575f5ffd5b506101f36101ee3660046142ac565b610670565b60405161020091906142f7565b60405180910390f35b348015610214575f5ffd5b506102286102233660046143c2565b610756565b005b348015610235575f5ffd5b50610228610244366004614400565b610ac1565b348015610254575f5ffd5b5061025e60075481565b604051908152602001610200565b348015610277575f5ffd5b50610280610cd1565b604051610200929190614456565b348015610299575f5ffd5b5061025e6102a8366004614510565b610d0b565b3480156102b8575f5ffd5b506001546102d99073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610200565b348015610309575f5ffd5b5061025e600a5481565b34801561031e575f5ffd5b5061025e60085481565b610228610336366004614556565b611365565b348015610346575f5ffd5b5061025e611384565b34801561035a575f5ffd5b507fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff166040519015158152602001610200565b34801561039b575f5ffd5b5061025e6103aa36600461461b565b6113b2565b3480156103ba575f5ffd5b5061025e600b5481565b3480156103cf575f5ffd5b506103e36103de366004614663565b611494565b60405161020091906146e0565b3480156103fb575f5ffd5b506102d9611563565b34801561040f575f5ffd5b5061025e60045481565b348015610424575f5ffd5b506003546102d99073ffffffffffffffffffffffffffffffffffffffff1681565b348015610450575f5ffd5b5061025e60065481565b348015610465575f5ffd5b5061022861047436600461473c565b6115f6565b348015610484575f5ffd5b506104c16040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b6040516102009190614793565b3480156104d9575f5ffd5b506102286104e83660046147e6565b611ae5565b3480156104f8575f5ffd5b5060408051606080820183525f8083526020808401829052928401528251808201845260035473ffffffffffffffffffffffffffffffffffffffff908116808352600254821683860190815260045493870193845286519182525190911693810193909352519282019290925201610200565b348015610576575f5ffd5b50610228610585366004614805565b611b3e565b348015610595575f5ffd5b50600954600a5460408051928352602083019190915201610200565b3480156105bc575f5ffd5b5061025e60055481565b3480156105d1575f5ffd5b506102286105e0366004614836565b611d12565b3480156105f0575f5ffd5b506002546102d99073ffffffffffffffffffffffffffffffffffffffff1681565b34801561061c575f5ffd5b5061025e60095481565b348015610631575f5ffd5b505f546102d99073ffffffffffffffffffffffffffffffffffffffff1681565b34801561065c575f5ffd5b5061022861066b366004614868565b611e5a565b60605f808460018111156106865761068661467a565b14610692576016610695565b60155b83519091505f8167ffffffffffffffff8111156106b4576106b46141b0565b6040519080825280602002602001820160405280156106e757816020015b60608152602001906001900390816106d25790505b5090505f5b8281101561074a57610725845f88848151811061070b5761070b614888565b602002602001015181526020019081526020015f20612072565b82828151811061073757610737614888565b60209081029190910101526001016106ec565b50925050505b92915050565b61075e611563565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146107e257335b6040517f118cdaa700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911660048201526024015b60405180910390fd5b8015610abc5760025473ffffffffffffffffffffffffffffffffffffffff84811691161480156108ac57506005546002546040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152839173ffffffffffffffffffffffffffffffffffffffff16906370a0823190602401602060405180830381865afa15801561087c573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108a091906148b5565b6108aa91906148f9565b105b156108ff576002546040517f7bdf264d00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911660048201526024016107d9565b60035473ffffffffffffffffffffffffffffffffffffffff84811691161480156109c357506006546003546040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152839173ffffffffffffffffffffffffffffffffffffffff16906370a0823190602401602060405180830381865afa158015610993573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109b791906148b5565b6109c191906148f9565b105b15610a16576003546040517f7bdf264d00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911660048201526024016107d9565b610a3773ffffffffffffffffffffffffffffffffffffffff84168383612118565b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16610a6c3390565b73ffffffffffffffffffffffffffffffffffffffff167f7e30e394efdaabad9f9ca6fd8f67f9449ba63d674ee8068283c3ae1f49b5b66084604051610ab391815260200190565b60405180910390a45b505050565b60015473ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610b4357335b6040517f143db19100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911660048201526024016107d9565b80515f5b81811015610ccb575f838281518110610b6257610b62614888565b6020908102919091018101515f81815260179092526040808320815160a0810190925280549294509091829060ff166001811115610ba257610ba261467a565b6001811115610bb357610bb361467a565b81528154610100810473ffffffffffffffffffffffffffffffffffffffff908116602080850191909152750100000000000000000000000000000000000000000090920463ffffffff1660408401526001840154606084015260029093015460809092019190915282015191925016610c2d575050610b47565b8573ffffffffffffffffffffffffffffffffffffffff16816020015173ffffffffffffffffffffffffffffffffffffffff1614610cb5576040517f54de53f70000000000000000000000000000000000000000000000000000000081526004810183905273ffffffffffffffffffffffffffffffffffffffff871660248201526044016107d9565b610cc1828260026121a5565b5050600101610b47565b50505050565b606080610cf7600d5f5b60ff1660028110610cee57610cee614888565b60040201612072565b9150610d05600d6001610cdb565b90509091565b5f610d14612301565b60015473ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610d4f5733610af8565b610d5761235f565b60608501511580610d7857506009548560600151610d759190614939565b15155b15610db75784606001516040517f4e1404fe0000000000000000000000000000000000000000000000000000000081526004016107d991815260200190565b60808501511580610dd85750600a548560800151610dd59190614939565b15155b15610e175784608001516040517fa334626e0000000000000000000000000000000000000000000000000000000081526004016107d991815260200190565b600c5f8154610e2590614971565b918290555090505f8086516001811115610e4157610e4161467a565b1490505f81610e5b57610e5683885f876124ae565b610e66565b610e66838886612709565b905086608001515f03610eb3575f837f78615f01112d8477ae8d6fc9a510bcfdf935bc3ab7b02d0ff2316f5b44bd597a42604051610ea691815260200190565b60405180910390a3611348565b6001866002811115610ec757610ec761467a565b03610f42576001837f78615f01112d8477ae8d6fc9a510bcfdf935bc3ab7b02d0ff2316f5b44bd597a42604051610f0091815260200190565b60405180910390a38115610f3d5760208701516080880151600254610f3d9273ffffffffffffffffffffffffffffffffffffffff90911691612118565b611348565b6002866002811115610f5657610f5661467a565b03610fab5760208701516040517f1b3c356200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911660048201526024016107d9565b8115611193578415801590610fc35750848760600151105b156110035760608701516040517fc18aa6060000000000000000000000000000000000000000000000000000000081526107d9915f918890600401614989565b866080015160055f82825461101891906149a8565b90915550507f1d2ff3fa6980aeeebca4e94965520da48983006e9b1115c1c853cbb10d943d005c63ffffffff166040808901919091525f848152601760205220875181548992919082907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600183818111156110975761109761467a565b021790555060208201518154604084015163ffffffff167501000000000000000000000000000000000000000000027fffffffffffffff00000000ffffffffffffffffffffffffffffffffffffffffff73ffffffffffffffffffffffffffffffffffffffff90931661010002929092167fffffffffffffff000000000000000000000000000000000000000000000000ff9091161717815560608083015160018301556080909201516002909101558701516111709086600d5f5b60ff166002811061116557611165614888565b600402019190612966565b5060608701515f90815260156020526040902061118d9084612b14565b50611348565b84158015906111a55750848760600151115b156111e65760608701516040517fc18aa6060000000000000000000000000000000000000000000000000000000081526107d9916001918890600401614989565b6111fb87606001518860800151600454612b24565b60065f82825461120b91906149a8565b90915550505f604080890182905284825260176020529020875181548992919082907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600183818111156112625761126261467a565b021790555060208201518154604084015163ffffffff167501000000000000000000000000000000000000000000027fffffffffffffff00000000ffffffffffffffffffffffffffffffffffffffffff73ffffffffffffffffffffffffffffffffffffffff90931661010002929092167fffffffffffffff000000000000000000000000000000000000000000000000ff90911617178155606080830151600180840191909155608090930151600290920191909155880151611329918790600d90611152565b5060608701515f9081526016602052604090206113469084612b14565b505b8161135b5761135b876020015182612bda565b5050949350505050565b61136d612caf565b61137682612db3565b6113808282612df3565b5050565b5f61138d612f2c565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b5f808560018111156113c6576113c661467a565b0361142f5760408051808201825261142891869190869060029083908390808284375f92019190915250869150600d90508960018111156114095761140961467a565b60ff166002811061141c5761141c614888565b60040201929190612f9b565b905061148c565b60408051808201825261142891869190869060029083908390808284375f92019190915250869150600d905089600181111561146d5761146d61467a565b60ff166002811061148057611480614888565b60040201929190613141565b949350505050565b6040805160a0810182525f808252602082018190529181018290526060810182905260808101919091525f8281526017602052604090819020815160a081019092528054829060ff1660018111156114ee576114ee61467a565b60018111156114ff576114ff61467a565b81528154610100810473ffffffffffffffffffffffffffffffffffffffff1660208301527501000000000000000000000000000000000000000000900463ffffffff1660408201526001820154606082015260029091015460809091015292915050565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156115cd573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906115f191906149bb565b905090565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff165f811580156116405750825b90505f8267ffffffffffffffff16600114801561165c5750303b155b90508115801561166a575080155b156116a1576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000016600117855583156117025784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b73ffffffffffffffffffffffffffffffffffffffff8a16611771576040517fc587916e0000000000000000000000000000000000000000000000000000000081527f726f75746572000000000000000000000000000000000000000000000000000060048201526024016107d9565b73ffffffffffffffffffffffffffffffffffffffff89166117e0576040517fc587916e0000000000000000000000000000000000000000000000000000000081527f71756f746500000000000000000000000000000000000000000000000000000060048201526024016107d9565b73ffffffffffffffffffffffffffffffffffffffff881661184f576040517fc587916e0000000000000000000000000000000000000000000000000000000081527f626173650000000000000000000000000000000000000000000000000000000060048201526024016107d9565b865f036118aa576040517fc587916e0000000000000000000000000000000000000000000000000000000081527f7469636b53697a6500000000000000000000000000000000000000000000000060048201526024016107d9565b855f03611905576040517fc587916e0000000000000000000000000000000000000000000000000000000081527f6c6f7453697a650000000000000000000000000000000000000000000000000060048201526024016107d9565b335f80547fffffffffffffffffffffffff000000000000000000000000000000000000000090811673ffffffffffffffffffffffffffffffffffffffff938416179091556001805482168d84161790556003805482168c841617905560028054909116918a169182179055604080517f313ce567000000000000000000000000000000000000000000000000000000008152905163313ce567916004808201926020929091908290030181865afa1580156119c2573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906119e691906149d6565b6119f190600a614ad9565b6004819055611a008789614ae7565b611a0a9190614939565b15611a5357600480546040517f575681350000000000000000000000000000000000000000000000000000000081529182018990526024820188905260448201526064016107d9565b6009879055600a869055600454611a6d9088908890612b24565b600b55611a786132b0565b8315611ad95784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050505050565b611aed611563565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611b255733610792565b8015611b3657611b336132c0565b50565b611b33613360565b611b46612301565b60015473ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611b815733610af8565b611b8961235f565b5f600c5f8154611b9890614971565b918290555090505f84516001811115611bb357611bb361467a565b03611c6157821580611bd05750600a54611bcd9084614939565b15155b15611c0a576040517fa334626e000000000000000000000000000000000000000000000000000000008152600481018490526024016107d9565b5f606085015260808401839052611c22818584612709565b50608084015115611c5c5760208401516080850151600254611c5c9273ffffffffffffffffffffffffffffffffffffffff90911691612118565b611cd1565b600b54831015611cab57600b546040517f0bc1e7dd0000000000000000000000000000000000000000000000000000000081526107d9918591600401918252602082015260400190565b5f1960608501525f611cbf828686866124ae565b9050611ccf856020015182612bda565b505b5f817f78615f01112d8477ae8d6fc9a510bcfdf935bc3ab7b02d0ff2316f5b44bd597a42604051611d0491815260200190565b60405180910390a350505050565b611d1a6133d6565b611d22611563565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611d5a5733610792565b80515f5b81811015610abc575f838281518110611d7957611d79614888565b6020908102919091018101515f81815260179092526040808320815160a0810190925280549294509091829060ff166001811115611db957611db961467a565b6001811115611dca57611dca61467a565b81528154610100810473ffffffffffffffffffffffffffffffffffffffff908116602080850191909152750100000000000000000000000000000000000000000090920463ffffffff1660408401526001840154606084015260029093015460809092019190915282015191925016611e44575050611d5e565b611e50828260036121a5565b5050600101611d5e565b5f5473ffffffffffffffffffffffffffffffffffffffff1663a1eea778336040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815273ffffffffffffffffffffffffffffffffffffffff90911660048201526024015f6040518083038186803b158015611eda575f5ffd5b505afa158015611eec573d5f5f3e3d5ffd5b50505050805f03611f4b576040517fc587916e0000000000000000000000000000000000000000000000000000000081527f7469636b53697a6500000000000000000000000000000000000000000000000060048201526024016107d9565b815f03611fa6576040517fc587916e0000000000000000000000000000000000000000000000000000000081527f6c6f7453697a650000000000000000000000000000000000000000000000000060048201526024016107d9565b600454611fb38383614ae7565b611fbd9190614939565b1561200657600480546040517f575681350000000000000000000000000000000000000000000000000000000081529182018390526024820184905260448201526064016107d9565b600a546009546040805192835260208301859052820152606081018290527f66748bc112f6692a76bebecc33e69768508e27e414e3080f5c02295b6ea3cf3f9060800160405180910390a1600a829055600981905560045461206b9082908490612b24565b600b555050565b60605f825f015467ffffffffffffffff811115612091576120916141b0565b6040519080825280602002602001820160405280156120ba578160200160208202803683370190505b5083546001850154919250905f5b8281101561210e57818482815181106120e3576120e3614888565b6020908102919091018101919091525f928352600387019052604090912060019081015491016120c8565b5091949350505050565b6040805173ffffffffffffffffffffffffffffffffffffffff8416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb00000000000000000000000000000000000000000000000000000000179052610abc908490613431565b5f8080845160018111156121bb576121bb61467a565b14905080156122325760608401515f90815260156020526040902060808501519092501561222d57602084015160808501516002546122129273ffffffffffffffffffffffffffffffffffffffff90911691612118565b836080015160055f82825461222791906148f9565b90915550505b6122ab565b60165f856060015181526020019081526020015f2091505f61225f85606001518660800151600454612b24565b905080156122a95760208501516003546122929173ffffffffffffffffffffffffffffffffffffffff9091169083612118565b8060065f8282546122a391906148f9565b90915550505b505b6122b68584846134d0565b81546122fa576122f88460600151600d865f015160018111156122db576122db61467a565b60ff16600281106122ee576122ee614888565b60040201906135a2565b505b5050505050565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff161561235d576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505f5f8273ffffffffffffffffffffffffffffffffffffffff1663c415b95c6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156123ce573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906123f291906149bb565b8373ffffffffffffffffffffffffffffffffffffffff166324a9d8536040518163ffffffff1660e01b8152600401602060405180830381865afa15801561243b573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061245f9190614afe565b91509150817fd6aa07baf8485abf9d26fecf4c935d75b50a73e678db02b944bd3ac8759823005d807f1d2ff3fa6980aeeebca4e94965520da48983006e9b1115c1c853cbb10d943d005d505050565b5f6001845160018111156124c4576124c461467a565b146124fe5783516040517f7cbdd0810000000000000000000000000000000000000000000000000000000081526107d99190600401614b19565b5f831561250c575082612524565b61252185606001518660800151600454612b24565b90505b6003546040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201525f9182916125cd9173ffffffffffffffffffffffffffffffffffffffff16906370a0823190602401602060405180830381865afa158015612596573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906125ba91906148b5565b846006546125c891906149a8565b613673565b9150915081612624576003546040517f7bdf264d00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911660048201526024016107d9565b865160018111156126375761263761467a565b88886020015173ffffffffffffffffffffffffffffffffffffffff167f7cde7559ef00f8c3d9244d3fde19e7500979d64f6f9c086a3836a4926abbe1658a606001518b608001514260405161269f939291909283526020830191909152604082015260600190565b60405180910390a45f6126b489898989613697565b905080156126fd578060055f8282546126cd91906148f9565b909155505060208801516002546126fd9173ffffffffffffffffffffffffffffffffffffffff9091169083612118565b50979650505050505050565b5f808351600181111561271e5761271e61467a565b146127585782516040517f7cbdd0810000000000000000000000000000000000000000000000000000000081526107d99190600401614b19565b826080015160055461276a91906149a8565b6002546040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff909116906370a0823190602401602060405180830381865afa1580156127d6573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906127fa91906148b5565b101561284e576002546040517f7bdf264d00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911660048201526024016107d9565b825160018111156128615761286161467a565b84846020015173ffffffffffffffffffffffffffffffffffffffff167f7cde7559ef00f8c3d9244d3fde19e7500979d64f6f9c086a3836a4926abbe16586606001518760800151426040516128c9939291909283526020830191909152604082015260600190565b60405180910390a45f6128dd8585856138a8565b9050801561295a57600680547fd6aa07baf8485abf9d26fecf4c935d75b50a73e678db02b944bd3ac8759823005c917f1d2ff3fa6980aeeebca4e94965520da48983006e9b1115c1c853cbb10d943d005c918491905f9061293f9084906148f9565b92505081905550612957878760200151858585613a63565b50505b5f9150505b9392505050565b5f825f036129a0576040517f0da93b9700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6129aa8484613b7a565b156129b657505f61295f565b5f828152600385016020908152604080832081518083019092528382529181019290925290835f03612a2257600186018054908690558015612a05575f81815260038801602052604090208690555b60405180604001604052805f815260200182815250915050612ac4565b604080518082019091528254815260018301546020820152612a4390613bc2565b80612a515750856001015484145b612a87576040517f182dfca500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5060018101546040805180820190915284815260208101829052908015612abb575f81815260038801602052604090208690555b50600182018590555b80602001515f03612ad757600286018590555b5f8581526003870160209081526040822083518155908301516001909101558654879190612b0490614971565b9091555060019695505050505050565b5f61295f83838560020154612966565b5f838302815f1985870982811083820303915050805f03612b5857838281612b4e57612b4e61490c565b049250505061295f565b808411612b6f57612b6f6003851502601118613bd8565b5f848688095f868103871696879004966002600389028118808a02820302808a02820302808a02820302808a02820302808a02820302808a02909103029181900381900460010186841190950394909402919094039290920491909117919091029150509392505050565b5f60065482612be991906149a8565b6003546040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff909116906370a0823190602401602060405180830381865afa158015612c55573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612c7991906148b5565b612c8391906148f9565b90508015610abc57600354610abc9073ffffffffffffffffffffffffffffffffffffffff168483612118565b3073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161480612d7c57507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16612d637f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614155b1561235d576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612dbb611563565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611b335733610792565b8173ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015612e78575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201909252612e75918101906148b5565b60015b612ec6576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff831660048201526024016107d9565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8114612f22576040517faa1d49a4000000000000000000000000000000000000000000000000000000008152600481018290526024016107d9565b610abc8383613be9565b3073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161461235d576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f835f03612fd5576040517f0da93b9700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612fdf8585613b7a565b15612ffa57505f83815260038501602052604090205461148c565b846001015484108061300b57508454155b1561301757505f61148c565b846002015484111561302e5750600284015461148c565b815f03613067576040517fab7e639300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6130728685613c4b565b9050848110156130d2575b801580159061309157505f19909201918215155b156130cd575f908152600386016020526040902060010154848111156130c8575f908152600386016020526040902054905061148c565b61307d565b61310f565b80158015906130e657505f19909201918215155b1561310f575f9081526003860160205260409020548481101561310a57905061148c565b6130d2565b6040517f80ce60d100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f835f0361317b576040517f0da93b9700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6131858585613b7a565b156131a057505f83815260038501602052604090205461148c565b84600101548411806131b157508454155b156131bd57505f61148c565b84600201548410156131d45750600284015461148c565b815f0361320d576040517fab7e639300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6132188685613c4b565b905084811115613273575b801580159061323757505f19909201918215155b156130cd575f9081526003860160205260409020600101548085111561326e575f908152600386016020526040902054905061148c565b613223565b801580159061328757505f19909201918215155b1561310f575f908152600386016020526040902054808510156132ab57905061148c565b613273565b6132b8613c95565b61235d613cfc565b6132c8612301565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a150565b6133686133d6565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa33613335565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff1661235d576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5f60205f8451602086015f885af180613450576040513d5f823e3d81fd5b50505f513d91508115613467578060011415613481565b73ffffffffffffffffffffffffffffffffffffffff84163b155b15610ccb576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024016107d9565b6134da81846135a2565b613513576040517ffcfdf902000000000000000000000000000000000000000000000000000000008152600481018490526024016107d9565b5f83815260176020526040812080547fffffffffffffff0000000000000000000000000000000000000000000000000016815560018101829055600201558160038111156135635761356361467a565b837f78615f01112d8477ae8d6fc9a510bcfdf935bc3ab7b02d0ff2316f5b44bd597a4260405161359591815260200190565b60405180910390a3505050565b5f815f036135dc576040517f0da93b9700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6135e68383613b7a565b6135f157505f610750565b5f82815260038401602052604080822080548352818320600180830154808652939094208482019390935581548355928601549092919085900361363a57600180840154908701555b8486600201540361364d57825460028701555b5f85815260038701602052604081208181556001018190558654879190612b0490614b27565b5f5f8383111561368757505f905080613690565b50600190508183035b9250929050565b5f80600d5b80541561389d575f6136ae8282613d4d565b905086606001518111156136c2575061389d565b5f8181526015602052604090206136d882613dbb565b861561372c575f6136f56136ec868a6148f9565b60045485612b24565b9050600a54816137059190614939565b61370f90826148f9565b60808a0181905290505f81900361372a5750505050506138a0565b505b5f613735825490565b90505b8015613851575f6137498382613d4d565b5f81815260176020526040812091925080806137698f8f87878c8c613de1565b9250925092505f61377d8984600454612b24565b90506137ad8685837fd6aa07baf8485abf9d26fecf4c935d75b50a73e678db02b944bd3ac8759823005c86613a63565b6137b7838d6149a8565b9b506137c3818c6149a8565b9a508e608001515f14806137e057506137db8d614b27565b9c508c155b1561383f57875461382f576137f58a8a6135a2565b61382f575f896040517f9f1cfdfe0000000000000000000000000000000000000000000000000000000081526004016107d9929190614b3c565b50505050505050505050506138a0565b86600190039650505050505050613738565b61385b84846135a2565b613895575f836040517f9f1cfdfe0000000000000000000000000000000000000000000000000000000081526004016107d9929190614b3c565b50505061369c565b50505b61148c613f07565b6002546004545f9173ffffffffffffffffffffffffffffffffffffffff169060115b805415613a57575f6138dc8282613d4d565b905086606001518110156138f05750613a57565b5f81815260166020526040902061390682613dbb565b5f61390f825490565b90505b8015613a0a575f6139238382613d4d565b5f818152601760205260408120919250806139428e8e86868b8b613de1565b50909250905061396973ffffffffffffffffffffffffffffffffffffffff8b168383612118565b61397487828b612b24565b61397e908c6149a8565b9a508c608001515f148061399b57506139968c614b27565b9b508b155b156139fa5785546139eb576139b088886135a2565b6139eb576001876040517f9f1cfdfe0000000000000000000000000000000000000000000000000000000081526004016107d9929190614b3c565b50505050505050505050613a5b565b8460019003945050505050613912565b613a1484846135a2565b613a4f576001836040517f9f1cfdfe0000000000000000000000000000000000000000000000000000000081526004016107d9929190614b3c565b5050506138ca565b5050505b61295f613f07565b8063ffffffff165f03613a9957600354613a949073ffffffffffffffffffffffffffffffffffffffff168585612118565b6122fa565b5f613aad848363ffffffff16612710612b24565b90505f613aba82866148f9565b6040805187815263ffffffff861660208201529081018490526060810182905290915073ffffffffffffffffffffffffffffffffffffffff808616919088169089907f9814b2b2f43f9ea51ad7e8760fda094ffa3012bfeb10c76a9656389af18956039060800160405180910390a4600354613b4d9073ffffffffffffffffffffffffffffffffffffffff168584612118565b600354613b719073ffffffffffffffffffffffffffffffffffffffff168783612118565b50505050505050565b5f815f03613b8957505f610750565b826001015482148061295f57505f828152600384016020908152604091829020825180840190935280548352600101549082015261295f905b80515f9015158061075057505060200151151590565b634e487b715f52806020526024601cfd5b613bf282613f38565b60405173ffffffffffffffffffffffffffffffffffffffff8316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115613c4357610abc8282614006565b611380614085565b5f613c5d8383835b6020020151613b7a565b15613c7157815f5b60200201519050610750565b613c7d83836001613c53565b15613c8a57816001613c65565b506001820154610750565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff1661235d576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b613d04613c95565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b81545f908210613d89576040517f39e60f7300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001830154825b8015613db3575f91825260038501602052604090912060010154905f1901613d90565b509392505050565b807ffd0e5d4f9b88892d3b04349a0e2bc0d1359414c21932fcd7d5a523a6c0a5cd005d50565b5f5f5f855f0160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16613e1a896080015188600201546140bd565b875491945092507501000000000000000000000000000000000000000000900463ffffffff1690505f80808a516001811115613e5857613e5861467a565b14613e6457888b613e67565b8a895b915091508681837f6a6896b1e6131e0b8ebae43fdc84cb0178a6eacdf4ee63f15dabae48729a8a038742604051613ea8929190918252602082015260400190565b60405180910390a487600201548403613ecb57613ec6895f886134d0565b613ed7565b60028801805485900390555b89608001518403613eed575f60808b0152613ef9565b60808a01805185900390525b505096509650969350505050565b7ffd0e5d4f9b88892d3b04349a0e2bc0d1359414c21932fcd7d5a523a6c0a5cd005c8015611b335760075542600855565b8073ffffffffffffffffffffffffffffffffffffffff163b5f03613fa0576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024016107d9565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60605f5f8473ffffffffffffffffffffffffffffffffffffffff168460405161402f9190614b57565b5f60405180830381855af49150503d805f8114614067576040519150601f19603f3d011682016040523d82523d5f602084013e61406c565b606091505b509150915061407c8583836140cc565b95945050505050565b341561235d576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f82821882841002821861295f565b6060826140e1576140dc8261415b565b61295f565b8151158015614105575073ffffffffffffffffffffffffffffffffffffffff84163b155b15614154576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024016107d9565b508061295f565b80511561416b5780518082602001fd5b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8035600281106141ab575f5ffd5b919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715614224576142246141b0565b604052919050565b5f82601f83011261423b575f5ffd5b813567ffffffffffffffff811115614255576142556141b0565b8060051b614265602082016141dd565b91825260208185018101929081019086841115614280575f5ffd5b6020860192505b838310156142a2578235825260209283019290910190614287565b9695505050505050565b5f5f604083850312156142bd575f5ffd5b6142c68361419d565b9150602083013567ffffffffffffffff8111156142e1575f5ffd5b6142ed8582860161422c565b9150509250929050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b82811015614395578685037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0018452815180518087526020918201918701905f5b8181101561437c57835183526020938401939092019160010161435e565b509096505050602093840193919091019060010161431d565b50929695505050505050565b73ffffffffffffffffffffffffffffffffffffffff81168114611b33575f5ffd5b5f5f5f606084860312156143d4575f5ffd5b83356143df816143a1565b925060208401356143ef816143a1565b929592945050506040919091013590565b5f5f60408385031215614411575f5ffd5b82356142c6816143a1565b5f8151808452602084019350602083015f5b8281101561444c57815186526020958601959091019060010161442e565b5093949350505050565b604081525f614468604083018561441c565b828103602084015261407c818561441c565b63ffffffff81168114611b33575f5ffd5b5f60a0828403121561449b575f5ffd5b60405160a0810167ffffffffffffffff811182821017156144be576144be6141b0565b6040529050806144cd8361419d565b815260208301356144dd816143a1565b602082015260408301356144f08161447a565b604082015260608381013590820152608092830135920191909152919050565b5f5f5f5f6101008587031215614524575f5ffd5b61452e868661448b565b935060a085013560038110614541575f5ffd5b939693955050505060c08201359160e0013590565b5f5f60408385031215614567575f5ffd5b8235614572816143a1565b9150602083013567ffffffffffffffff81111561458d575f5ffd5b8301601f8101851361459d575f5ffd5b803567ffffffffffffffff8111156145b7576145b76141b0565b6145e860207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116016141dd565b8181528660208385010111156145fc575f5ffd5b816020840160208301375f602083830101528093505050509250929050565b5f5f5f5f60a0858703121561462e575f5ffd5b6146378561419d565b9350602085013592506080850186811115614650575f5ffd5b9396929550505060409290920191903590565b5f60208284031215614673575f5ffd5b5035919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b600281106146dc577f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b9052565b5f60a0820190506146f28284516146a7565b73ffffffffffffffffffffffffffffffffffffffff602084015116602083015263ffffffff6040840151166040830152606083015160608301526080830151608083015292915050565b5f5f5f5f5f60a08688031215614750575f5ffd5b853561475b816143a1565b9450602086013561476b816143a1565b9350604086013561477b816143a1565b94979396509394606081013594506080013592915050565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b5f602082840312156147f6575f5ffd5b8135801515811461295f575f5ffd5b5f5f5f60e08486031215614817575f5ffd5b614821858561448b565b9560a0850135955060c0909401359392505050565b5f60208284031215614846575f5ffd5b813567ffffffffffffffff81111561485c575f5ffd5b61148c8482850161422c565b5f5f60408385031215614879575f5ffd5b50508035926020909101359150565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f602082840312156148c5575f5ffd5b5051919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b81810381811115610750576107506148cc565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f8261496c577f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b500690565b5f5f198203614982576149826148cc565b5060010190565b6060810161499782866146a7565b602082019390935260400152919050565b80820180821115610750576107506148cc565b5f602082840312156149cb575f5ffd5b815161295f816143a1565b5f602082840312156149e6575f5ffd5b815160ff8116811461295f575f5ffd5b6001815b6001841115614a3157808504811115614a1557614a156148cc565b6001841615614a2357908102905b60019390931c9280026149fa565b935093915050565b5f82614a4757506001610750565b81614a5357505f610750565b8160018114614a695760028114614a7357614a8f565b6001915050610750565b60ff841115614a8457614a846148cc565b50506001821b610750565b5060208310610133831016604e8410600b8410161715614ab2575081810a610750565b614abe5f1984846149f6565b805f1904821115614ad157614ad16148cc565b029392505050565b5f61295f60ff841683614a39565b8082028115828204841417610750576107506148cc565b5f60208284031215614b0e575f5ffd5b815161295f8161447a565b6020810161075082846146a7565b5f81614b3557614b356148cc565b505f190190565b60408101614b4a82856146a7565b8260208301529392505050565b5f82518060208501845e5f92019182525091905056fea264697066735822122037d6f8b4fb3aa5e19c34cb759a4f19140f3bc95caaf780a8ff85c8291048755564736f6c634300081c0033"

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
