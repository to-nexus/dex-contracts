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
	Bin: "0x60a06040523073ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff16815250348015610042575f5ffd5b5061005161005660201b60201c565b6101b6565b5f61006561015460201b60201c565b9050805f0160089054906101000a900460ff16156100af576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8016815f015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff16146101515767ffffffffffffffff815f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d267ffffffffffffffff604051610148919061019d565b60405180910390a15b50565b5f7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00905090565b5f67ffffffffffffffff82169050919050565b6101978161017b565b82525050565b5f6020820190506101b05f83018461018e565b92915050565b6080516165956101dc5f395f8181613596015281816135eb015261382201526165955ff3fe6080604052600436106101d7575f3560e01c80638da5cb5b11610101578063cc50e36a11610094578063ec342ad011610063578063ec342ad014610679578063f210d087146106a3578063f46f16c2146106cd578063fe566349146106f7576101d7565b8063cc50e36a146105d4578063d6e9cc2a146105fc578063dfdf2a7214610627578063e077dd8a14610651576101d7565b8063a6b63eb8116100d0578063a6b63eb814610530578063ad3cb1cc14610558578063bedb86fb14610582578063c3f909d4146105aa576101d7565b80638da5cb5b14610488578063918f8674146104b25780639c579839146104dc5780639da771f414610506576101d7565b80634942f65f116101795780635c975abb116101485780635c975abb146103bc5780636c4cfbc8146103e65780637cbf6db2146104225780638d77eba51461044c576101d7565b80634942f65f146103225780634d6754f11461034c5780634f1ef2861461037657806352d1902d14610392576101d7565b80631edd8614116101b55780631edd8614146102675780632cfffaf6146102915780632dd6a00a146102bc57806332fe7b26146102f8576101d7565b806312a808cc146101db5780631a0361fb146102175780631ec482d71461023f575b5f5ffd5b3480156101e6575f5ffd5b5061020160048036038101906101fc919061513e565b61071f565b60405161020e919061530a565b60405180910390f35b348015610222575f5ffd5b5061023d600480360381019061023891906153bf565b610824565b005b34801561024a575f5ffd5b506102656004803603810190610260919061540f565b610c2b565b005b348015610272575f5ffd5b5061027b610eb4565b6040516102889190615478565b60405180910390f35b34801561029c575f5ffd5b506102a5610eba565b6040516102b39291906154fd565b60405180910390f35b3480156102c7575f5ffd5b506102e260048036038101906102dd919061561b565b610f2f565b6040516102ef9190615478565b60405180910390f35b348015610303575f5ffd5b5061030c61167e565b604051610319919061568f565b60405180910390f35b34801561032d575f5ffd5b506103366116a3565b6040516103439190615478565b60405180910390f35b348015610357575f5ffd5b506103606116a9565b60405161036d9190615478565b60405180910390f35b610390600480360381019061038b9190615758565b6116af565b005b34801561039d575f5ffd5b506103a66116ce565b6040516103b391906157ca565b60405180910390f35b3480156103c7575f5ffd5b506103d06116ff565b6040516103dd91906157fd565b60405180910390f35b3480156103f1575f5ffd5b5061040c60048036038101906104079190615837565b611721565b6040516104199190615478565b60405180910390f35b34801561042d575f5ffd5b50610436611851565b6040516104439190615478565b60405180910390f35b348015610457575f5ffd5b50610472600480360381019061046d919061589b565b611857565b60405161047f91906159bd565b60405180910390f35b348015610493575f5ffd5b5061049c611949565b6040516104a9919061568f565b60405180910390f35b3480156104bd575f5ffd5b506104c66119dc565b6040516104d39190615478565b60405180910390f35b3480156104e7575f5ffd5b506104f06119e2565b6040516104fd9190615a31565b60405180910390f35b348015610511575f5ffd5b5061051a611a07565b6040516105279190615478565b60405180910390f35b34801561053b575f5ffd5b5061055660048036038101906105519190615a4a565b611a0d565b005b348015610563575f5ffd5b5061056c611f67565b6040516105799190615b21565b60405180910390f35b34801561058d575f5ffd5b506105a860048036038101906105a39190615b6b565b611fa0565b005b3480156105b5575f5ffd5b506105be612044565b6040516105cb9190615be5565b60405180910390f35b3480156105df575f5ffd5b506105fa60048036038101906105f59190615bfe565b6120dc565b005b348015610607575f5ffd5b50610610612397565b60405161061e929190615c4e565b60405180910390f35b348015610632575f5ffd5b5061063b6123a7565b6040516106489190615478565b60405180910390f35b34801561065c575f5ffd5b5061067760048036038101906106729190615c75565b6123ad565b005b348015610684575f5ffd5b5061068d6125ad565b60405161069a9190615a31565b60405180910390f35b3480156106ae575f5ffd5b506106b76125d2565b6040516106c49190615478565b60405180910390f35b3480156106d8575f5ffd5b506106e16125d8565b6040516106ee919061568f565b60405180910390f35b348015610702575f5ffd5b5061071d60048036038101906107189190615cbc565b6125fc565b005b60605f5f6001811115610735576107346158c6565b5b846001811115610748576107476158c6565b5b14610754576016610757565b60155b90505f835190505f8167ffffffffffffffff81111561077957610778614fcf565b5b6040519080825280602002602001820160405280156107ac57816020015b60608152602001906001900390816107975790505b5090505f5f90505b82811015610817576107ee845f8884815181106107d4576107d3615cfa565b5b602002602001015181526020019081526020015f206127d4565b82828151811061080157610800615cfa565b5b60200260200101819052508060010190506107b4565b5080935050505092915050565b61082c611949565b73ffffffffffffffffffffffffffffffffffffffff1661084a612891565b73ffffffffffffffffffffffffffffffffffffffff16146108a95761086d612891565b6040517f118cdaa70000000000000000000000000000000000000000000000000000000081526004016108a0919061568f565b60405180910390fd5b5f810315610c265760025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161480156109b157506005548160025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401610966919061568f565b602060405180830381865afa158015610981573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109a59190615d3b565b6109af9190615d93565b105b15610a145760025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040517f7bdf264d000000000000000000000000000000000000000000000000000000008152600401610a0b919061568f565b60405180910390fd5b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16148015610b1457506006548160035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401610ac9919061568f565b602060405180830381865afa158015610ae4573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b089190615d3b565b610b129190615d93565b105b15610b775760035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040517f7bdf264d000000000000000000000000000000000000000000000000000000008152600401610b6e919061568f565b60405180910390fd5b610ba282828573ffffffffffffffffffffffffffffffffffffffff166128989092919063ffffffff16565b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16610bd8612891565b73ffffffffffffffffffffffffffffffffffffffff167f7e30e394efdaabad9f9ca6fd8f67f9449ba63d674ee8068283c3ae1f49b5b66084604051610c1d9190615478565b60405180910390a45b505050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16610c6b612891565b73ffffffffffffffffffffffffffffffffffffffff1614610cca57610c8e612891565b6040517f143db191000000000000000000000000000000000000000000000000000000008152600401610cc1919061568f565b60405180910390fd5b5f815190505f5f90505b81811015610eae575f838281518110610cf057610cef615cfa565b5b602002602001015190505f60175f8381526020019081526020015f206040518060a00160405290815f82015f9054906101000a900460ff166001811115610d3a57610d396158c6565b5b6001811115610d4c57610d4b6158c6565b5b81526020015f820160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020015f820160159054906101000a900463ffffffff1663ffffffff1663ffffffff1681526020016001820154815260200160028201548152505090505f73ffffffffffffffffffffffffffffffffffffffff16816020015173ffffffffffffffffffffffffffffffffffffffff1603610e1e575050610ea9565b8573ffffffffffffffffffffffffffffffffffffffff16816020015173ffffffffffffffffffffffffffffffffffffffff1614610e945781866040517f54de53f7000000000000000000000000000000000000000000000000000000008152600401610e8b929190615dc6565b60405180910390fd5b610ea082826002612917565b82600101925050505b610cd4565b50505050565b60075481565b606080610ef2600d5f6001811115610ed557610ed46158c6565b5b60ff1660028110610ee957610ee8615cfa565b5b600402016127d4565b9150610f29600d600180811115610f0c57610f0b6158c6565b5b60ff1660028110610f2057610f1f615cfa565b5b600402016127d4565b90509091565b5f610f38612aec565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16610f78612891565b73ffffffffffffffffffffffffffffffffffffffff1614610fd757610f9b612891565b6040517f143db191000000000000000000000000000000000000000000000000000000008152600401610fce919061568f565b60405180910390fd5b610fdf612b2d565b5f8560600151148061100257505f6009548660600151610fff9190615e1a565b14155b156110485784606001516040517f4e1404fe00000000000000000000000000000000000000000000000000000000815260040161103f9190615478565b60405180910390fd5b5f8560800151148061106b57505f600a5486608001516110689190615e1a565b14155b156110b15784608001516040517fa334626e0000000000000000000000000000000000000000000000000000000081526004016110a89190615478565b60405180910390fd5b600c5f81546110bf90615e4a565b91905081905590505f5f60018111156110db576110da6158c6565b5b865f015160018111156110f1576110f06158c6565b5b1490505f8161110b5761110683885f87612c7c565b611117565b611116838886612f2e565b5b90505f876080015103611174575f6003811115611137576111366158c6565b5b837f78615f01112d8477ae8d6fc9a510bcfdf935bc3ab7b02d0ff2316f5b44bd597a426040516111679190615478565b60405180910390a3611660565b60016002811115611188576111876158c6565b5b86600281111561119b5761119a6158c6565b5b0361124c57600160038111156111b4576111b36158c6565b5b837f78615f01112d8477ae8d6fc9a510bcfdf935bc3ab7b02d0ff2316f5b44bd597a426040516111e49190615478565b60405180910390a38115611247576112468760200151886080015160025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166128989092919063ffffffff16565b5b61165f565b60028081111561125f5761125e6158c6565b5b866002811115611272576112716158c6565b5b036112b85786602001516040517f1b3c35620000000000000000000000000000000000000000000000000000000081526004016112af919061568f565b60405180910390fd5b811561148b575f85141580156112d15750848760600151105b1561131b575f8760600151866040517fc18aa60600000000000000000000000000000000000000000000000000000000815260040161131293929190615ea0565b60405180910390fd5b866080015160055f8282546113309190615ed5565b9250508190555061133f613183565b876040019063ffffffff16908163ffffffff16815250508660175f8581526020019081526020015f205f820151815f015f6101000a81548160ff021916908360018111156113905761138f6158c6565b5b02179055506020820151815f0160016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040820151815f0160156101000a81548163ffffffff021916908363ffffffff160217905550606082015181600101556080820151816002015590505061145d876060015186600d5f6001811115611436576114356158c6565b5b60ff166002811061144a57611449615cfa565b5b600402016131ab9092919063ffffffff16565b506114858360155f8a6060015181526020019081526020015f2061338a90919063ffffffff16565b5061165e565b5f851415801561149e5750848760600151115b156114e95760018760600151866040517fc18aa6060000000000000000000000000000000000000000000000000000000081526004016114e093929190615ea0565b60405180910390fd5b6114fe876060015188608001516004546133a2565b60065f82825461150e9190615ed5565b925050819055505f876040019063ffffffff16908163ffffffff16815250508660175f8581526020019081526020015f205f820151815f015f6101000a81548160ff02191690836001811115611567576115666158c6565b5b02179055506020820151815f0160016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040820151815f0160156101000a81548163ffffffff021916908363ffffffff1602179055506060820151816001015560808201518160020155905050611634876060015186600d60018081111561160d5761160c6158c6565b5b60ff166002811061162157611620615cfa565b5b600402016131ab9092919063ffffffff16565b5061165c8360165f8a6060015181526020019081526020015f2061338a90919063ffffffff16565b505b5b5b8161167457611673876020015182613487565b5b5050949350505050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600a5481565b60085481565b6116b7613594565b6116c08261367a565b6116ca8282613702565b5050565b5f6116d7613820565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b905090565b5f5f6117096138a7565b9050805f015f9054906101000a900460ff1691505090565b5f5f6001811115611735576117346158c6565b5b856001811115611748576117476158c6565b5b036117cd576117c6848460028060200260405190810160405280929190826002602002808284375f81840152601f19601f82011690508083019250505050505084600d89600181111561179e5761179d6158c6565b5b60ff16600281106117b2576117b1615cfa565b5b600402016138ce909392919063ffffffff16565b9050611849565b611846848460028060200260405190810160405280929190826002602002808284375f81840152601f19601f82011690508083019250505050505084600d89600181111561181e5761181d6158c6565b5b60ff166002811061183257611831615cfa565b5b60040201613aaa909392919063ffffffff16565b90505b949350505050565b600b5481565b61185f614ecd565b60175f8381526020019081526020015f206040518060a00160405290815f82015f9054906101000a900460ff16600181111561189e5761189d6158c6565b5b60018111156118b0576118af6158c6565b5b81526020015f820160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020015f820160159054906101000a900463ffffffff1663ffffffff1663ffffffff168152602001600182015481526020016002820154815250509050919050565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156119b3573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906119d79190615f1c565b905090565b60045481565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60065481565b5f611a16613c86565b90505f815f0160089054906101000a900460ff161590505f825f015f9054906101000a900467ffffffffffffffff1690505f5f8267ffffffffffffffff16148015611a5e5750825b90505f60018367ffffffffffffffff16148015611a9157505f3073ffffffffffffffffffffffffffffffffffffffff163b145b905081158015611a9f575080155b15611ad6576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001855f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055508315611b23576001855f0160086101000a81548160ff0219169083151502179055505b5f73ffffffffffffffffffffffffffffffffffffffff168a73ffffffffffffffffffffffffffffffffffffffff1603611b91576040517fc587916e000000000000000000000000000000000000000000000000000000008152600401611b8890615f6d565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff1603611bff576040517fc587916e000000000000000000000000000000000000000000000000000000008152600401611bf690615faa565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff1603611c6d576040517fc587916e000000000000000000000000000000000000000000000000000000008152600401611c6490615fe7565b60405180910390fd5b5f8703611caf576040517fc587916e000000000000000000000000000000000000000000000000000000008152600401611ca690616024565b60405180910390fd5b5f8603611cf1576040517fc587916e000000000000000000000000000000000000000000000000000000008152600401611ce890616061565b60405180910390fd5b611cf9612891565b5f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508960015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508860035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508760025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508773ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa158015611e40573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611e6491906160ae565b600a611e709190616208565b6004819055505f6004548789611e869190616252565b611e909190615e1a565b14611ed85786866004546040517f57568135000000000000000000000000000000000000000000000000000000008152600401611ecf93929190616293565b60405180910390fd5b8660098190555085600a81905550611ef387876004546133a2565b600b81905550611f01613cad565b8315611f5b575f855f0160086101000a81548160ff0219169083151502179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d26001604051611f529190616314565b60405180910390a15b50505050505050505050565b6040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b611fa8611949565b73ffffffffffffffffffffffffffffffffffffffff16611fc6612891565b73ffffffffffffffffffffffffffffffffffffffff161461202557611fe9612891565b6040517f118cdaa700000000000000000000000000000000000000000000000000000000815260040161201c919061568f565b60405180910390fd5b801561203857612033613cbf565b612041565b612040613d2e565b5b50565b61204c614f25565b604051806060016040528060035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600454815250905090565b6120e4612aec565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16612124612891565b73ffffffffffffffffffffffffffffffffffffffff161461218357612147612891565b6040517f143db19100000000000000000000000000000000000000000000000000000000815260040161217a919061568f565b60405180910390fd5b61218b612b2d565b5f600c5f815461219a90615e4a565b91905081905590505f60018111156121b5576121b46158c6565b5b845f015160018111156121cb576121ca6158c6565b5b036122b2575f8314806121eb57505f600a54846121e89190615e1a565b14155b1561222d57826040517fa334626e0000000000000000000000000000000000000000000000000000000081526004016122249190615478565b60405180910390fd5b5f8460600181815250508284608001818152505061224c818584612f2e565b505f8460800151146122ad576122ac8460200151856080015160025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166128989092919063ffffffff16565b5b612346565b600b548310156122fd5782600b546040517f0bc1e7dd0000000000000000000000000000000000000000000000000000000081526004016122f4929190615c4e565b60405180910390fd5b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8460600181815250505f61233482868686612c7c565b9050612344856020015182613487565b505b5f6003811115612359576123586158c6565b5b817f78615f01112d8477ae8d6fc9a510bcfdf935bc3ab7b02d0ff2316f5b44bd597a426040516123899190615478565b60405180910390a350505050565b5f5f6009549150600a5490509091565b60055481565b6123b5613d9c565b6123bd611949565b73ffffffffffffffffffffffffffffffffffffffff166123db612891565b73ffffffffffffffffffffffffffffffffffffffff161461243a576123fe612891565b6040517f118cdaa7000000000000000000000000000000000000000000000000000000008152600401612431919061568f565b60405180910390fd5b5f815190505f5f90505b818110156125a8575f8382815181106124605761245f615cfa565b5b602002602001015190505f60175f8381526020019081526020015f206040518060a00160405290815f82015f9054906101000a900460ff1660018111156124aa576124a96158c6565b5b60018111156124bc576124bb6158c6565b5b81526020015f820160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020015f820160159054906101000a900463ffffffff1663ffffffff1663ffffffff1681526020016001820154815260200160028201548152505090505f73ffffffffffffffffffffffffffffffffffffffff16816020015173ffffffffffffffffffffffffffffffffffffffff160361258e5750506125a3565b61259a82826003612917565b82600101925050505b612444565b505050565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60095481565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a1eea778612640612891565b6040518263ffffffff1660e01b815260040161265c919061568f565b5f6040518083038186803b158015612672575f5ffd5b505afa158015612684573d5f5f3e3d5ffd5b505050505f81036126ca576040517fc587916e0000000000000000000000000000000000000000000000000000000081526004016126c190616024565b60405180910390fd5b5f820361270c576040517fc587916e00000000000000000000000000000000000000000000000000000000815260040161270390616061565b60405180910390fd5b5f600454838361271c9190616252565b6127269190615e1a565b1461276e5780826004546040517f5756813500000000000000000000000000000000000000000000000000000000815260040161276593929190616293565b60405180910390fd5b7f66748bc112f6692a76bebecc33e69768508e27e414e3080f5c02295b6ea3cf3f600a5483600954846040516127a7949392919061632d565b60405180910390a181600a81905550806009819055506127ca81836004546133a2565b600b819055505050565b60605f825f015467ffffffffffffffff8111156127f4576127f3614fcf565b5b6040519080825280602002602001820160405280156128225781602001602082028036833780820191505090505b5090505f835f015490505f846001015490505f5f90505b82811015612885578184828151811061285557612854615cfa565b5b602002602001018181525050856003015f8381526020019081526020015f20600101549150806001019050612839565b50829350505050919050565b5f33905090565b612912838473ffffffffffffffffffffffffffffffffffffffff1663a9059cbb85856040516024016128cb929190616370565b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050613ddc565b505050565b5f5f5f600181111561292c5761292b6158c6565b5b845f01516001811115612942576129416158c6565b5b14905080156129e35760155f856060015181526020019081526020015f2091505f8460800151146129de576129c18460200151856080015160025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166128989092919063ffffffff16565b836080015160055f8282546129d69190615d93565b925050819055505b612a84565b60165f856060015181526020019081526020015f2091505f612a10856060015186608001516004546133a2565b90505f8114612a8257612a6985602001518260035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166128989092919063ffffffff16565b8060065f828254612a7a9190615d93565b925050819055505b505b612a8f858484613e77565b612a9882613f8b565b15612ae557612ae38460600151600d865f01516001811115612abd57612abc6158c6565b5b60ff1660028110612ad157612ad0615cfa565b5b60040201613f9990919063ffffffff16565b505b5050505050565b612af46116ff565b15612b2b576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505f5f8273ffffffffffffffffffffffffffffffffffffffff1663c415b95c6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612b9c573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612bc09190615f1c565b8373ffffffffffffffffffffffffffffffffffffffff166324a9d8536040518163ffffffff1660e01b8152600401602060405180830381865afa158015612c09573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612c2d91906163ab565b91509150817fd6aa07baf8485abf9d26fecf4c935d75b50a73e678db02b944bd3ac8759823005d807f1d2ff3fa6980aeeebca4e94965520da48983006e9b1115c1c853cbb10d943d005d505050565b5f600180811115612c9057612c8f6158c6565b5b845f01516001811115612ca657612ca56158c6565b5b14612ceb57835f01516040517f7cbdd081000000000000000000000000000000000000000000000000000000008152600401612ce291906163d6565b60405180910390fd5b5f5f8414612cfb57839050612d13565b612d10856060015186608001516004546133a2565b90505b5f5f612dc460035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401612d72919061568f565b602060405180830381865afa158015612d8d573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612db19190615d3b565b84600654612dbf9190615ed5565b6140c2565b9150915081612e2b5760035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040517f7bdf264d000000000000000000000000000000000000000000000000000000008152600401612e22919061568f565b60405180910390fd5b865f01516001811115612e4157612e406158c6565b5b88886020015173ffffffffffffffffffffffffffffffffffffffff167f7cde7559ef00f8c3d9244d3fde19e7500979d64f6f9c086a3836a4926abbe1658a606001518b6080015142604051612e9893929190616293565b60405180910390a45f612ead898989896140e8565b90505f8114612f1f578060055f828254612ec79190615d93565b92505081905550612f1e88602001518260025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166128989092919063ffffffff16565b5b81945050505050949350505050565b5f5f6001811115612f4257612f416158c6565b5b835f01516001811115612f5857612f576158c6565b5b14612f9d57825f01516040517f7cbdd081000000000000000000000000000000000000000000000000000000008152600401612f9491906163d6565b60405180910390fd5b8260800151600554612faf9190615ed5565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401613009919061568f565b602060405180830381865afa158015613024573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906130489190615d3b565b10156130ac5760025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040517f7bdf264d0000000000000000000000000000000000000000000000000000000081526004016130a3919061568f565b60405180910390fd5b825f015160018111156130c2576130c16158c6565b5b84846020015173ffffffffffffffffffffffffffffffffffffffff167f7cde7559ef00f8c3d9244d3fde19e7500979d64f6f9c086a3836a4926abbe165866060015187608001514260405161311993929190616293565b60405180910390a45f61312d858585614376565b90505f8114613178575f5f6131406145d4565b613148613183565b915091508260065f82825461315d9190615d93565b925050819055506131758787602001518585856145fc565b50505b5f9150509392505050565b5f7f1d2ff3fa6980aeeebca4e94965520da48983006e9b1115c1c853cbb10d943d005c905090565b5f5f83036131e5576040517f0da93b9700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6131ef848461478d565b156131fc575f9050613383565b5f846003015f8481526020019081526020015f20905061321a614f6f565b5f8403613271575f866001015490508587600101819055505f81146132545785876003015f8381526020019081526020015f205f01819055505b60405180604001604052805f815260200182815250915050613328565b613298826040518060400160405290815f82015481526020016001820154815250506147ed565b806132a65750856001015484145b6132dc576040517f182dfca500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8260010154905060405180604001604052808681526020018281525091505f811461331d5785876003015f8381526020019081526020015f205f01819055505b858360010181905550505b5f81602001510361333d578486600201819055505b80866003015f8781526020019081526020015f205f820151815f015560208201518160010155905050855f015f815461337590615e4a565b919050819055506001925050505b9392505050565b5f61339a838385600201546131ab565b905092915050565b5f5f83850290505f5f19858709828110838203039150505f81036133da578382816133d0576133cf615ded565b5b0492505050613480565b8084116133f9576133f86133f35f86146012601161480b565b614824565b5b5f8486880990508281118203915080830392505f855f038616905080860495508084049350600181825f0304019050808302841793505f600287600302189050808702600203810290508087026002038102905080870260020381029050808702600203810290508087026002038102905080870260020381029050808502955050505050505b9392505050565b5f600654826134969190615ed5565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016134f0919061568f565b602060405180830381865afa15801561350b573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061352f9190615d3b565b6135399190615d93565b90505f811461358f5761358e838260035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166128989092919063ffffffff16565b5b505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff16148061364157507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16613628614835565b73ffffffffffffffffffffffffffffffffffffffff1614155b15613678576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b613682611949565b73ffffffffffffffffffffffffffffffffffffffff166136a0612891565b73ffffffffffffffffffffffffffffffffffffffff16146136ff576136c3612891565b6040517f118cdaa70000000000000000000000000000000000000000000000000000000081526004016136f6919061568f565b60405180910390fd5b50565b8173ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa92505050801561376a57506040513d601f19601f820116820180604052508101906137679190616419565b60015b6137ab57816040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526004016137a2919061568f565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b811461381157806040517faa1d49a400000000000000000000000000000000000000000000000000000000815260040161380891906157ca565b60405180910390fd5b61381b8383614888565b505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff16146138a5576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300905090565b5f5f8403613908576040517f0da93b9700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b613912858561478d565b1561393457846003015f8581526020019081526020015f205f01549050613aa2565b846001015484108061394b575061394a85613f8b565b5b15613958575f9050613aa2565b84600201548411156139705784600201549050613aa2565b5f82036139a9576040517fab7e639300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6139b486856148fa565b905084811015613a24575b5f81141580156139d657505f836001900393508314155b15613a1f57856003015f8281526020019081526020015f2060010154905084811115613a1a57856003015f8281526020019081526020015f205f0154915050613aa2565b6139bf565b613a70565b5b5f8114158015613a3c57505f836001900393508314155b15613a6f57856003015f8281526020019081526020015f205f0154905084811015613a6a5780915050613aa2565b613a25565b5b6040517f80ce60d100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b949350505050565b5f5f8403613ae4576040517f0da93b9700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b613aee858561478d565b15613b1057846003015f8581526020019081526020015f205f01549050613c7e565b8460010154841180613b275750613b2685613f8b565b5b15613b34575f9050613c7e565b8460020154841015613b4c5784600201549050613c7e565b5f8203613b85576040517fab7e639300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f613b9086856148fa565b905084811115613c00575b5f8114158015613bb257505f836001900393508314155b15613bfb57856003015f8281526020019081526020015f2060010154905080851115613bf657856003015f8281526020019081526020015f205f0154915050613c7e565b613b9b565b613c4c565b5b5f8114158015613c1857505f836001900393508314155b15613c4b57856003015f8281526020019081526020015f205f0154905080851015613c465780915050613c7e565b613c01565b5b6040517f80ce60d100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b949350505050565b5f7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00905090565b613cb5614995565b613cbd6149d5565b565b613cc7612aec565b5f613cd06138a7565b90506001815f015f6101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258613d16612891565b604051613d23919061568f565b60405180910390a150565b613d36613d9c565b5f613d3f6138a7565b90505f815f015f6101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa613d84612891565b604051613d91919061568f565b60405180910390a150565b613da46116ff565b613dda576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f5f60205f8451602086015f885af180613dfb576040513d5f823e3d81fd5b3d92505f519150505f8214613e14576001811415613e2f565b5f8473ffffffffffffffffffffffffffffffffffffffff163b145b15613e7157836040517f5274afe7000000000000000000000000000000000000000000000000000000008152600401613e68919061568f565b60405180910390fd5b50505050565b613e8a8382613f9990919063ffffffff16565b613ecb57826040517ffcfdf902000000000000000000000000000000000000000000000000000000008152600401613ec29190615478565b60405180910390fd5b60175f8481526020019081526020015f205f5f82015f6101000a81549060ff02191690555f820160016101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690555f820160156101000a81549063ffffffff0219169055600182015f9055600282015f90555050816003811115613f4e57613f4d6158c6565b5b837f78615f01112d8477ae8d6fc9a510bcfdf935bc3ab7b02d0ff2316f5b44bd597a42604051613f7e9190615478565b60405180910390a3505050565b5f5f825f0154149050919050565b5f5f8203613fd3576040517f0da93b9700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b613fdd838361478d565b613fe9575f90506140bc565b5f836003015f8481526020019081526020015f2090505f846003015f835f015481526020019081526020015f2090505f856003015f846001015481526020019081526020015f20905082600101548260010181905550825f0154815f01819055508486600101540361406357826001015486600101819055505b8486600201540361407b57825f015486600201819055505b856003015f8681526020019081526020015f205f5f82015f9055600182015f90555050855f015f81546140ad90616444565b91905081905550600193505050505b92915050565b5f5f838311156140d7575f5f915091506140e1565b6001838503915091505b9250929050565b5f5f5f600d5f6001811115614100576140ff6158c6565b5b60ff166002811061411457614113615cfa565b5b6004020190505b61412481613f8b565b614363575f61413c5f83614a0590919063ffffffff16565b905086606001518111156141505750614363565b5f60155f8381526020019081526020015f20905061416d82614a86565b5f87146141c7575f61418d85896141849190615d93565b600454856133a2565b9050600a548161419d9190615e1a565b816141a89190615d93565b9050808960800181815250505f81036141c5575050505050614366565b505b5f6141d182614aac565b90505b5f8114614305575f6141ef5f84614a0590919063ffffffff16565b90505f60175f8381526020019081526020015f2090505f5f5f6142168f8f87878c8c614ab8565b9250925092505f61422a89846004546133a2565b905061424086858361423a6145d4565b866145fc565b828c61424c9190615ed5565b9b50808b61425a9190615ed5565b9a505f8f60800151148061427957505f8d61427490616444565b9d508d145b156142f35761428788613f8b565b156142e35761429f898b613f9990919063ffffffff16565b6142e2575f896040517f9f1cfdfe0000000000000000000000000000000000000000000000000000000081526004016142d992919061646b565b60405180910390fd5b5b5050505050505050505050614366565b866001900396505050505050506141d4565b6143188385613f9990919063ffffffff16565b61435b575f836040517f9f1cfdfe00000000000000000000000000000000000000000000000000000000815260040161435292919061646b565b60405180910390fd5b50505061411b565b50505b61436e614bf1565b949350505050565b5f5f5f60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600454915091505f600d6001808111156143b8576143b76158c6565b5b60ff16600281106143cc576143cb615cfa565b5b6004020190505b6143dc81613f8b565b6145c1575f6143f45f83614a0590919063ffffffff16565b9050866060015181101561440857506145c1565b5f60165f8381526020019081526020015f20905061442582614a86565b5f61442f82614aac565b90505b5f8114614562575f61444d5f84614a0590919063ffffffff16565b90505f60175f8381526020019081526020015f2090505f5f6144738e8e86868b8b614ab8565b50915091506144a382828c73ffffffffffffffffffffffffffffffffffffffff166128989092919063ffffffff16565b6144ae87828b6133a2565b8b6144b99190615ed5565b9a505f8d6080015114806144d857505f8c6144d390616444565b9c508c145b15614552576144e686613f8b565b15614543576144fe8789613f9990919063ffffffff16565b614542576001876040517f9f1cfdfe00000000000000000000000000000000000000000000000000000000815260040161453992919061646b565b60405180910390fd5b5b505050505050505050506145c5565b8460019003945050505050614432565b6145758385613f9990919063ffffffff16565b6145b9576001836040517f9f1cfdfe0000000000000000000000000000000000000000000000000000000081526004016145b092919061646b565b60405180910390fd5b5050506143d3565b5050505b6145cd614bf1565b9392505050565b5f7fd6aa07baf8485abf9d26fecf4c935d75b50a73e678db02b944bd3ac8759823005c905090565b5f8163ffffffff160361465a57614655848460035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166128989092919063ffffffff16565b614786565b5f61466e848363ffffffff166127106133a2565b90505f818561467d9190615d93565b90508373ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff16887f9814b2b2f43f9ea51ad7e8760fda094ffa3012bfeb10c76a9656389af1895603888787876040516146e394939291906164c2565b60405180910390a4614737848360035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166128989092919063ffffffff16565b614783868260035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166128989092919063ffffffff16565b50505b5050505050565b5f5f820361479d575f90506147e7565b82600101548214806147e457506147e3836003015f8481526020019081526020015f206040518060400160405290815f82015481526020016001820154815250506147ed565b5b90505b92915050565b5f5f825f015114158061480457505f826020015114155b9050919050565b5f61481584614c2f565b82841802821890509392505050565b634e487b715f52806020526024601cfd5b5f6148617f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b614c3a565b5f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b61489182614c43565b8173ffffffffffffffffffffffffffffffffffffffff167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b60405160405180910390a25f815111156148ed576148e78282614d0c565b506148f6565b6148f5614d8c565b5b5050565b5f61491c83835f6002811061491257614911615cfa565b5b602002015161478d565b1561494057815f6002811061493457614933615cfa565b5b6020020151905061498f565b614962838360016002811061495857614957615cfa565b5b602002015161478d565b15614987578160016002811061497b5761497a615cfa565b5b6020020151905061498f565b826001015490505b92915050565b61499d614dc8565b6149d3576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b6149dd614995565b5f6149e66138a7565b90505f815f015f6101000a81548160ff02191690831515021790555050565b5f825f01548210614a42576040517f39e60f7300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f836001015490505f8390505b5f8114614a7b5780600190039050846003015f8381526020019081526020015f20600101549150614a4f565b508091505092915050565b807ffd0e5d4f9b88892d3b04349a0e2bc0d1359414c21932fcd7d5a523a6c0a5cd005d50565b5f815f01549050919050565b5f5f5f855f0160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16614af189608001518860020154614de6565b875f0160159054906101000a900463ffffffff168093508194508295505050505f5f5f6001811115614b2657614b256158c6565b5b8a5f01516001811115614b3c57614b3b6158c6565b5b14614b4857888b614b4b565b8a895b915091508681837f6a6896b1e6131e0b8ebae43fdc84cb0178a6eacdf4ee63f15dabae48729a8a038742604051614b83929190615c4e565b60405180910390a487600201548403614ba657614ba1895f88613e77565b614bb8565b83886002015f82825403925050819055505b89608001518403614bd2575f8a6080018181525050614be3565b838a60800181815103915081815250505b505096509650969350505050565b5f7ffd0e5d4f9b88892d3b04349a0e2bc0d1359414c21932fcd7d5a523a6c0a5cd005c90505f8114614c2c5780600781905550426008819055505b50565b5f8115159050919050565b5f819050919050565b5f8173ffffffffffffffffffffffffffffffffffffffff163b03614c9e57806040517f4c9c8ce3000000000000000000000000000000000000000000000000000000008152600401614c95919061568f565b60405180910390fd5b80614cca7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b614c3a565b5f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60605f5f8473ffffffffffffffffffffffffffffffffffffffff1684604051614d359190616549565b5f60405180830381855af49150503d805f8114614d6d576040519150601f19603f3d011682016040523d82523d5f602084013e614d72565b606091505b5091509150614d82858383614dfc565b9250505092915050565b5f341115614dc6576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f614dd1613c86565b5f0160089054906101000a900460ff16905090565b5f614df4828410848461480b565b905092915050565b606082614e1157614e0c82614e89565b614e81565b5f8251148015614e3757505f8473ffffffffffffffffffffffffffffffffffffffff163b145b15614e7957836040517f9996b315000000000000000000000000000000000000000000000000000000008152600401614e70919061568f565b60405180910390fd5b819050614e82565b5b9392505050565b5f81511115614e9b5780518082602001fd5b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040518060a001604052805f6001811115614eeb57614eea6158c6565b5b81526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020015f63ffffffff1681526020015f81526020015f81525090565b60405180606001604052805f73ffffffffffffffffffffffffffffffffffffffff1681526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020015f81525090565b60405180604001604052805f81526020015f81525090565b5f604051905090565b5f5ffd5b5f5ffd5b60028110614fa4575f5ffd5b50565b5f81359050614fb581614f98565b92915050565b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61500582614fbf565b810181811067ffffffffffffffff8211171561502457615023614fcf565b5b80604052505050565b5f615036614f87565b90506150428282614ffc565b919050565b5f67ffffffffffffffff82111561506157615060614fcf565b5b602082029050602081019050919050565b5f5ffd5b5f819050919050565b61508881615076565b8114615092575f5ffd5b50565b5f813590506150a38161507f565b92915050565b5f6150bb6150b684615047565b61502d565b905080838252602082019050602084028301858111156150de576150dd615072565b5b835b8181101561510757806150f38882615095565b8452602084019350506020810190506150e0565b5050509392505050565b5f82601f83011261512557615124614fbb565b5b81356151358482602086016150a9565b91505092915050565b5f5f6040838503121561515457615153614f90565b5b5f61516185828601614fa7565b925050602083013567ffffffffffffffff81111561518257615181614f94565b5b61518e85828601615111565b9150509250929050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b6151f381615076565b82525050565b5f61520483836151ea565b60208301905092915050565b5f602082019050919050565b5f615226826151c1565b61523081856151cb565b935061523b836151db565b805f5b8381101561526b57815161525288826151f9565b975061525d83615210565b92505060018101905061523e565b5085935050505092915050565b5f615283838361521c565b905092915050565b5f602082019050919050565b5f6152a182615198565b6152ab81856151a2565b9350836020820285016152bd856151b2565b805f5b858110156152f857848403895281516152d98582615278565b94506152e48361528b565b925060208a019950506001810190506152c0565b50829750879550505050505092915050565b5f6020820190508181035f8301526153228184615297565b905092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6153538261532a565b9050919050565b5f61536482615349565b9050919050565b6153748161535a565b811461537e575f5ffd5b50565b5f8135905061538f8161536b565b92915050565b61539e81615349565b81146153a8575f5ffd5b50565b5f813590506153b981615395565b92915050565b5f5f5f606084860312156153d6576153d5614f90565b5b5f6153e386828701615381565b93505060206153f4868287016153ab565b925050604061540586828701615095565b9150509250925092565b5f5f6040838503121561542557615424614f90565b5b5f615432858286016153ab565b925050602083013567ffffffffffffffff81111561545357615452614f94565b5b61545f85828601615111565b9150509250929050565b61547281615076565b82525050565b5f60208201905061548b5f830184615469565b92915050565b5f82825260208201905092915050565b5f6154ab826151c1565b6154b58185615491565b93506154c0836151db565b805f5b838110156154f05781516154d788826151f9565b97506154e283615210565b9250506001810190506154c3565b5085935050505092915050565b5f6040820190508181035f83015261551581856154a1565b9050818103602083015261552981846154a1565b90509392505050565b5f5ffd5b5f63ffffffff82169050919050565b61554e81615536565b8114615558575f5ffd5b50565b5f8135905061556981615545565b92915050565b5f60a0828403121561558457615583615532565b5b61558e60a061502d565b90505f61559d84828501614fa7565b5f8301525060206155b0848285016153ab565b60208301525060406155c48482850161555b565b60408301525060606155d884828501615095565b60608301525060806155ec84828501615095565b60808301525092915050565b60038110615604575f5ffd5b50565b5f81359050615615816155f8565b92915050565b5f5f5f5f610100858703121561563457615633614f90565b5b5f6156418782880161556f565b94505060a061565287828801615607565b93505060c061566387828801615095565b92505060e061567487828801615095565b91505092959194509250565b61568981615349565b82525050565b5f6020820190506156a25f830184615680565b92915050565b5f5ffd5b5f67ffffffffffffffff8211156156c6576156c5614fcf565b5b6156cf82614fbf565b9050602081019050919050565b828183375f83830152505050565b5f6156fc6156f7846156ac565b61502d565b905082815260208101848484011115615718576157176156a8565b5b6157238482856156dc565b509392505050565b5f82601f83011261573f5761573e614fbb565b5b813561574f8482602086016156ea565b91505092915050565b5f5f6040838503121561576e5761576d614f90565b5b5f61577b858286016153ab565b925050602083013567ffffffffffffffff81111561579c5761579b614f94565b5b6157a88582860161572b565b9150509250929050565b5f819050919050565b6157c4816157b2565b82525050565b5f6020820190506157dd5f8301846157bb565b92915050565b5f8115159050919050565b6157f7816157e3565b82525050565b5f6020820190506158105f8301846157ee565b92915050565b5f8190508260206002028201111561583157615830615072565b5b92915050565b5f5f5f5f60a0858703121561584f5761584e614f90565b5b5f61585c87828801614fa7565b945050602061586d87828801615095565b935050604061587e87828801615816565b925050608061588f87828801615095565b91505092959194509250565b5f602082840312156158b0576158af614f90565b5b5f6158bd84828501615095565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b60028110615904576159036158c6565b5b50565b5f819050615914826158f3565b919050565b5f61592382615907565b9050919050565b61593381615919565b82525050565b61594281615349565b82525050565b61595181615536565b82525050565b60a082015f82015161596b5f85018261592a565b50602082015161597e6020850182615939565b5060408201516159916040850182615948565b5060608201516159a460608501826151ea565b5060808201516159b760808501826151ea565b50505050565b5f60a0820190506159d05f830184615957565b92915050565b5f819050919050565b5f6159f96159f46159ef8461532a565b6159d6565b61532a565b9050919050565b5f615a0a826159df565b9050919050565b5f615a1b82615a00565b9050919050565b615a2b81615a11565b82525050565b5f602082019050615a445f830184615a22565b92915050565b5f5f5f5f5f60a08688031215615a6357615a62614f90565b5b5f615a70888289016153ab565b9550506020615a81888289016153ab565b9450506040615a92888289016153ab565b9350506060615aa388828901615095565b9250506080615ab488828901615095565b9150509295509295909350565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f615af382615ac1565b615afd8185615acb565b9350615b0d818560208601615adb565b615b1681614fbf565b840191505092915050565b5f6020820190508181035f830152615b398184615ae9565b905092915050565b615b4a816157e3565b8114615b54575f5ffd5b50565b5f81359050615b6581615b41565b92915050565b5f60208284031215615b8057615b7f614f90565b5b5f615b8d84828501615b57565b91505092915050565b615b9f81615a11565b82525050565b606082015f820151615bb95f850182615b96565b506020820151615bcc6020850182615b96565b506040820151615bdf60408501826151ea565b50505050565b5f606082019050615bf85f830184615ba5565b92915050565b5f5f5f60e08486031215615c1557615c14614f90565b5b5f615c228682870161556f565b93505060a0615c3386828701615095565b92505060c0615c4486828701615095565b9150509250925092565b5f604082019050615c615f830185615469565b615c6e6020830184615469565b9392505050565b5f60208284031215615c8a57615c89614f90565b5b5f82013567ffffffffffffffff811115615ca757615ca6614f94565b5b615cb384828501615111565b91505092915050565b5f5f60408385031215615cd257615cd1614f90565b5b5f615cdf85828601615095565b9250506020615cf085828601615095565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f81519050615d358161507f565b92915050565b5f60208284031215615d5057615d4f614f90565b5b5f615d5d84828501615d27565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f615d9d82615076565b9150615da883615076565b9250828203905081811115615dc057615dbf615d66565b5b92915050565b5f604082019050615dd95f830185615469565b615de66020830184615680565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f615e2482615076565b9150615e2f83615076565b925082615e3f57615e3e615ded565b5b828206905092915050565b5f615e5482615076565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203615e8657615e85615d66565b5b600182019050919050565b615e9a81615919565b82525050565b5f606082019050615eb35f830186615e91565b615ec06020830185615469565b615ecd6040830184615469565b949350505050565b5f615edf82615076565b9150615eea83615076565b9250828201905080821115615f0257615f01615d66565b5b92915050565b5f81519050615f1681615395565b92915050565b5f60208284031215615f3157615f30614f90565b5b5f615f3e84828501615f08565b91505092915050565b7f726f757465720000000000000000000000000000000000000000000000000000815250565b5f602082019050615f7f5f8301615f47565b919050565b7f71756f7465000000000000000000000000000000000000000000000000000000815250565b5f602082019050615fbc5f8301615f84565b919050565b7f6261736500000000000000000000000000000000000000000000000000000000815250565b5f602082019050615ff95f8301615fc1565b919050565b7f7469636b53697a65000000000000000000000000000000000000000000000000815250565b5f6020820190506160365f8301615ffe565b919050565b7f6c6f7453697a6500000000000000000000000000000000000000000000000000815250565b5f6020820190506160735f830161603b565b919050565b5f60ff82169050919050565b61608d81616078565b8114616097575f5ffd5b50565b5f815190506160a881616084565b92915050565b5f602082840312156160c3576160c2614f90565b5b5f6160d08482850161609a565b91505092915050565b5f8160011c9050919050565b5f5f8291508390505b600185111561612e5780860481111561610a57616109615d66565b5b60018516156161195780820291505b8081029050616127856160d9565b94506160ee565b94509492505050565b5f826161465760019050616201565b81616153575f9050616201565b81600181146161695760028114616173576161a2565b6001915050616201565b60ff84111561618557616184615d66565b5b8360020a91508482111561619c5761619b615d66565b5b50616201565b5060208310610133831016604e8410600b84101617156161d75782820a9050838111156161d2576161d1615d66565b5b616201565b6161e484848460016160e5565b925090508184048111156161fb576161fa615d66565b5b81810290505b9392505050565b5f61621282615076565b915061621d83616078565b925061624a7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8484616137565b905092915050565b5f61625c82615076565b915061626783615076565b925082820261627581615076565b9150828204841483151761628c5761628b615d66565b5b5092915050565b5f6060820190506162a65f830186615469565b6162b36020830185615469565b6162c06040830184615469565b949350505050565b5f819050919050565b5f67ffffffffffffffff82169050919050565b5f6162fe6162f96162f4846162c8565b6159d6565b6162d1565b9050919050565b61630e816162e4565b82525050565b5f6020820190506163275f830184616305565b92915050565b5f6080820190506163405f830187615469565b61634d6020830186615469565b61635a6040830185615469565b6163676060830184615469565b95945050505050565b5f6040820190506163835f830185615680565b6163906020830184615469565b9392505050565b5f815190506163a581615545565b92915050565b5f602082840312156163c0576163bf614f90565b5b5f6163cd84828501616397565b91505092915050565b5f6020820190506163e95f830184615e91565b92915050565b6163f8816157b2565b8114616402575f5ffd5b50565b5f81519050616413816163ef565b92915050565b5f6020828403121561642e5761642d614f90565b5b5f61643b84828501616405565b91505092915050565b5f61644e82615076565b91505f82036164605761645f615d66565b5b600182039050919050565b5f60408201905061647e5f830185615e91565b61648b6020830184615469565b9392505050565b5f6164ac6164a76164a284615536565b6159d6565b615076565b9050919050565b6164bc81616492565b82525050565b5f6080820190506164d55f830187615469565b6164e260208301866164b3565b6164ef6040830185615469565b6164fc6060830184615469565b95945050505050565b5f81519050919050565b5f81905092915050565b5f61652382616505565b61652d818561650f565b935061653d818560208601615adb565b80840191505092915050565b5f6165548284616519565b91508190509291505056fea2646970667358221220995d36b9373a93547b3d433298fea200b523885a5836367e9caf9e1a758e4a3064736f6c634300081c0033",
}

// PairImplABI is the input ABI used to generate the binding from.
// Deprecated: Use PairImplMetaData.ABI instead.
var PairImplABI = PairImplMetaData.ABI

// PairImplBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const PairImplBinRuntime = "6080604052600436106101d7575f3560e01c80638da5cb5b11610101578063cc50e36a11610094578063ec342ad011610063578063ec342ad014610679578063f210d087146106a3578063f46f16c2146106cd578063fe566349146106f7576101d7565b8063cc50e36a146105d4578063d6e9cc2a146105fc578063dfdf2a7214610627578063e077dd8a14610651576101d7565b8063a6b63eb8116100d0578063a6b63eb814610530578063ad3cb1cc14610558578063bedb86fb14610582578063c3f909d4146105aa576101d7565b80638da5cb5b14610488578063918f8674146104b25780639c579839146104dc5780639da771f414610506576101d7565b80634942f65f116101795780635c975abb116101485780635c975abb146103bc5780636c4cfbc8146103e65780637cbf6db2146104225780638d77eba51461044c576101d7565b80634942f65f146103225780634d6754f11461034c5780634f1ef2861461037657806352d1902d14610392576101d7565b80631edd8614116101b55780631edd8614146102675780632cfffaf6146102915780632dd6a00a146102bc57806332fe7b26146102f8576101d7565b806312a808cc146101db5780631a0361fb146102175780631ec482d71461023f575b5f5ffd5b3480156101e6575f5ffd5b5061020160048036038101906101fc919061513e565b61071f565b60405161020e919061530a565b60405180910390f35b348015610222575f5ffd5b5061023d600480360381019061023891906153bf565b610824565b005b34801561024a575f5ffd5b506102656004803603810190610260919061540f565b610c2b565b005b348015610272575f5ffd5b5061027b610eb4565b6040516102889190615478565b60405180910390f35b34801561029c575f5ffd5b506102a5610eba565b6040516102b39291906154fd565b60405180910390f35b3480156102c7575f5ffd5b506102e260048036038101906102dd919061561b565b610f2f565b6040516102ef9190615478565b60405180910390f35b348015610303575f5ffd5b5061030c61167e565b604051610319919061568f565b60405180910390f35b34801561032d575f5ffd5b506103366116a3565b6040516103439190615478565b60405180910390f35b348015610357575f5ffd5b506103606116a9565b60405161036d9190615478565b60405180910390f35b610390600480360381019061038b9190615758565b6116af565b005b34801561039d575f5ffd5b506103a66116ce565b6040516103b391906157ca565b60405180910390f35b3480156103c7575f5ffd5b506103d06116ff565b6040516103dd91906157fd565b60405180910390f35b3480156103f1575f5ffd5b5061040c60048036038101906104079190615837565b611721565b6040516104199190615478565b60405180910390f35b34801561042d575f5ffd5b50610436611851565b6040516104439190615478565b60405180910390f35b348015610457575f5ffd5b50610472600480360381019061046d919061589b565b611857565b60405161047f91906159bd565b60405180910390f35b348015610493575f5ffd5b5061049c611949565b6040516104a9919061568f565b60405180910390f35b3480156104bd575f5ffd5b506104c66119dc565b6040516104d39190615478565b60405180910390f35b3480156104e7575f5ffd5b506104f06119e2565b6040516104fd9190615a31565b60405180910390f35b348015610511575f5ffd5b5061051a611a07565b6040516105279190615478565b60405180910390f35b34801561053b575f5ffd5b5061055660048036038101906105519190615a4a565b611a0d565b005b348015610563575f5ffd5b5061056c611f67565b6040516105799190615b21565b60405180910390f35b34801561058d575f5ffd5b506105a860048036038101906105a39190615b6b565b611fa0565b005b3480156105b5575f5ffd5b506105be612044565b6040516105cb9190615be5565b60405180910390f35b3480156105df575f5ffd5b506105fa60048036038101906105f59190615bfe565b6120dc565b005b348015610607575f5ffd5b50610610612397565b60405161061e929190615c4e565b60405180910390f35b348015610632575f5ffd5b5061063b6123a7565b6040516106489190615478565b60405180910390f35b34801561065c575f5ffd5b5061067760048036038101906106729190615c75565b6123ad565b005b348015610684575f5ffd5b5061068d6125ad565b60405161069a9190615a31565b60405180910390f35b3480156106ae575f5ffd5b506106b76125d2565b6040516106c49190615478565b60405180910390f35b3480156106d8575f5ffd5b506106e16125d8565b6040516106ee919061568f565b60405180910390f35b348015610702575f5ffd5b5061071d60048036038101906107189190615cbc565b6125fc565b005b60605f5f6001811115610735576107346158c6565b5b846001811115610748576107476158c6565b5b14610754576016610757565b60155b90505f835190505f8167ffffffffffffffff81111561077957610778614fcf565b5b6040519080825280602002602001820160405280156107ac57816020015b60608152602001906001900390816107975790505b5090505f5f90505b82811015610817576107ee845f8884815181106107d4576107d3615cfa565b5b602002602001015181526020019081526020015f206127d4565b82828151811061080157610800615cfa565b5b60200260200101819052508060010190506107b4565b5080935050505092915050565b61082c611949565b73ffffffffffffffffffffffffffffffffffffffff1661084a612891565b73ffffffffffffffffffffffffffffffffffffffff16146108a95761086d612891565b6040517f118cdaa70000000000000000000000000000000000000000000000000000000081526004016108a0919061568f565b60405180910390fd5b5f810315610c265760025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161480156109b157506005548160025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401610966919061568f565b602060405180830381865afa158015610981573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109a59190615d3b565b6109af9190615d93565b105b15610a145760025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040517f7bdf264d000000000000000000000000000000000000000000000000000000008152600401610a0b919061568f565b60405180910390fd5b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16148015610b1457506006548160035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401610ac9919061568f565b602060405180830381865afa158015610ae4573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b089190615d3b565b610b129190615d93565b105b15610b775760035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040517f7bdf264d000000000000000000000000000000000000000000000000000000008152600401610b6e919061568f565b60405180910390fd5b610ba282828573ffffffffffffffffffffffffffffffffffffffff166128989092919063ffffffff16565b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16610bd8612891565b73ffffffffffffffffffffffffffffffffffffffff167f7e30e394efdaabad9f9ca6fd8f67f9449ba63d674ee8068283c3ae1f49b5b66084604051610c1d9190615478565b60405180910390a45b505050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16610c6b612891565b73ffffffffffffffffffffffffffffffffffffffff1614610cca57610c8e612891565b6040517f143db191000000000000000000000000000000000000000000000000000000008152600401610cc1919061568f565b60405180910390fd5b5f815190505f5f90505b81811015610eae575f838281518110610cf057610cef615cfa565b5b602002602001015190505f60175f8381526020019081526020015f206040518060a00160405290815f82015f9054906101000a900460ff166001811115610d3a57610d396158c6565b5b6001811115610d4c57610d4b6158c6565b5b81526020015f820160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020015f820160159054906101000a900463ffffffff1663ffffffff1663ffffffff1681526020016001820154815260200160028201548152505090505f73ffffffffffffffffffffffffffffffffffffffff16816020015173ffffffffffffffffffffffffffffffffffffffff1603610e1e575050610ea9565b8573ffffffffffffffffffffffffffffffffffffffff16816020015173ffffffffffffffffffffffffffffffffffffffff1614610e945781866040517f54de53f7000000000000000000000000000000000000000000000000000000008152600401610e8b929190615dc6565b60405180910390fd5b610ea082826002612917565b82600101925050505b610cd4565b50505050565b60075481565b606080610ef2600d5f6001811115610ed557610ed46158c6565b5b60ff1660028110610ee957610ee8615cfa565b5b600402016127d4565b9150610f29600d600180811115610f0c57610f0b6158c6565b5b60ff1660028110610f2057610f1f615cfa565b5b600402016127d4565b90509091565b5f610f38612aec565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16610f78612891565b73ffffffffffffffffffffffffffffffffffffffff1614610fd757610f9b612891565b6040517f143db191000000000000000000000000000000000000000000000000000000008152600401610fce919061568f565b60405180910390fd5b610fdf612b2d565b5f8560600151148061100257505f6009548660600151610fff9190615e1a565b14155b156110485784606001516040517f4e1404fe00000000000000000000000000000000000000000000000000000000815260040161103f9190615478565b60405180910390fd5b5f8560800151148061106b57505f600a5486608001516110689190615e1a565b14155b156110b15784608001516040517fa334626e0000000000000000000000000000000000000000000000000000000081526004016110a89190615478565b60405180910390fd5b600c5f81546110bf90615e4a565b91905081905590505f5f60018111156110db576110da6158c6565b5b865f015160018111156110f1576110f06158c6565b5b1490505f8161110b5761110683885f87612c7c565b611117565b611116838886612f2e565b5b90505f876080015103611174575f6003811115611137576111366158c6565b5b837f78615f01112d8477ae8d6fc9a510bcfdf935bc3ab7b02d0ff2316f5b44bd597a426040516111679190615478565b60405180910390a3611660565b60016002811115611188576111876158c6565b5b86600281111561119b5761119a6158c6565b5b0361124c57600160038111156111b4576111b36158c6565b5b837f78615f01112d8477ae8d6fc9a510bcfdf935bc3ab7b02d0ff2316f5b44bd597a426040516111e49190615478565b60405180910390a38115611247576112468760200151886080015160025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166128989092919063ffffffff16565b5b61165f565b60028081111561125f5761125e6158c6565b5b866002811115611272576112716158c6565b5b036112b85786602001516040517f1b3c35620000000000000000000000000000000000000000000000000000000081526004016112af919061568f565b60405180910390fd5b811561148b575f85141580156112d15750848760600151105b1561131b575f8760600151866040517fc18aa60600000000000000000000000000000000000000000000000000000000815260040161131293929190615ea0565b60405180910390fd5b866080015160055f8282546113309190615ed5565b9250508190555061133f613183565b876040019063ffffffff16908163ffffffff16815250508660175f8581526020019081526020015f205f820151815f015f6101000a81548160ff021916908360018111156113905761138f6158c6565b5b02179055506020820151815f0160016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040820151815f0160156101000a81548163ffffffff021916908363ffffffff160217905550606082015181600101556080820151816002015590505061145d876060015186600d5f6001811115611436576114356158c6565b5b60ff166002811061144a57611449615cfa565b5b600402016131ab9092919063ffffffff16565b506114858360155f8a6060015181526020019081526020015f2061338a90919063ffffffff16565b5061165e565b5f851415801561149e5750848760600151115b156114e95760018760600151866040517fc18aa6060000000000000000000000000000000000000000000000000000000081526004016114e093929190615ea0565b60405180910390fd5b6114fe876060015188608001516004546133a2565b60065f82825461150e9190615ed5565b925050819055505f876040019063ffffffff16908163ffffffff16815250508660175f8581526020019081526020015f205f820151815f015f6101000a81548160ff02191690836001811115611567576115666158c6565b5b02179055506020820151815f0160016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040820151815f0160156101000a81548163ffffffff021916908363ffffffff1602179055506060820151816001015560808201518160020155905050611634876060015186600d60018081111561160d5761160c6158c6565b5b60ff166002811061162157611620615cfa565b5b600402016131ab9092919063ffffffff16565b5061165c8360165f8a6060015181526020019081526020015f2061338a90919063ffffffff16565b505b5b5b8161167457611673876020015182613487565b5b5050949350505050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600a5481565b60085481565b6116b7613594565b6116c08261367a565b6116ca8282613702565b5050565b5f6116d7613820565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b905090565b5f5f6117096138a7565b9050805f015f9054906101000a900460ff1691505090565b5f5f6001811115611735576117346158c6565b5b856001811115611748576117476158c6565b5b036117cd576117c6848460028060200260405190810160405280929190826002602002808284375f81840152601f19601f82011690508083019250505050505084600d89600181111561179e5761179d6158c6565b5b60ff16600281106117b2576117b1615cfa565b5b600402016138ce909392919063ffffffff16565b9050611849565b611846848460028060200260405190810160405280929190826002602002808284375f81840152601f19601f82011690508083019250505050505084600d89600181111561181e5761181d6158c6565b5b60ff166002811061183257611831615cfa565b5b60040201613aaa909392919063ffffffff16565b90505b949350505050565b600b5481565b61185f614ecd565b60175f8381526020019081526020015f206040518060a00160405290815f82015f9054906101000a900460ff16600181111561189e5761189d6158c6565b5b60018111156118b0576118af6158c6565b5b81526020015f820160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020015f820160159054906101000a900463ffffffff1663ffffffff1663ffffffff168152602001600182015481526020016002820154815250509050919050565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156119b3573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906119d79190615f1c565b905090565b60045481565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60065481565b5f611a16613c86565b90505f815f0160089054906101000a900460ff161590505f825f015f9054906101000a900467ffffffffffffffff1690505f5f8267ffffffffffffffff16148015611a5e5750825b90505f60018367ffffffffffffffff16148015611a9157505f3073ffffffffffffffffffffffffffffffffffffffff163b145b905081158015611a9f575080155b15611ad6576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001855f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055508315611b23576001855f0160086101000a81548160ff0219169083151502179055505b5f73ffffffffffffffffffffffffffffffffffffffff168a73ffffffffffffffffffffffffffffffffffffffff1603611b91576040517fc587916e000000000000000000000000000000000000000000000000000000008152600401611b8890615f6d565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff1603611bff576040517fc587916e000000000000000000000000000000000000000000000000000000008152600401611bf690615faa565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff1603611c6d576040517fc587916e000000000000000000000000000000000000000000000000000000008152600401611c6490615fe7565b60405180910390fd5b5f8703611caf576040517fc587916e000000000000000000000000000000000000000000000000000000008152600401611ca690616024565b60405180910390fd5b5f8603611cf1576040517fc587916e000000000000000000000000000000000000000000000000000000008152600401611ce890616061565b60405180910390fd5b611cf9612891565b5f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508960015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508860035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508760025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508773ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa158015611e40573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611e6491906160ae565b600a611e709190616208565b6004819055505f6004548789611e869190616252565b611e909190615e1a565b14611ed85786866004546040517f57568135000000000000000000000000000000000000000000000000000000008152600401611ecf93929190616293565b60405180910390fd5b8660098190555085600a81905550611ef387876004546133a2565b600b81905550611f01613cad565b8315611f5b575f855f0160086101000a81548160ff0219169083151502179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d26001604051611f529190616314565b60405180910390a15b50505050505050505050565b6040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b611fa8611949565b73ffffffffffffffffffffffffffffffffffffffff16611fc6612891565b73ffffffffffffffffffffffffffffffffffffffff161461202557611fe9612891565b6040517f118cdaa700000000000000000000000000000000000000000000000000000000815260040161201c919061568f565b60405180910390fd5b801561203857612033613cbf565b612041565b612040613d2e565b5b50565b61204c614f25565b604051806060016040528060035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600454815250905090565b6120e4612aec565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16612124612891565b73ffffffffffffffffffffffffffffffffffffffff161461218357612147612891565b6040517f143db19100000000000000000000000000000000000000000000000000000000815260040161217a919061568f565b60405180910390fd5b61218b612b2d565b5f600c5f815461219a90615e4a565b91905081905590505f60018111156121b5576121b46158c6565b5b845f015160018111156121cb576121ca6158c6565b5b036122b2575f8314806121eb57505f600a54846121e89190615e1a565b14155b1561222d57826040517fa334626e0000000000000000000000000000000000000000000000000000000081526004016122249190615478565b60405180910390fd5b5f8460600181815250508284608001818152505061224c818584612f2e565b505f8460800151146122ad576122ac8460200151856080015160025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166128989092919063ffffffff16565b5b612346565b600b548310156122fd5782600b546040517f0bc1e7dd0000000000000000000000000000000000000000000000000000000081526004016122f4929190615c4e565b60405180910390fd5b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8460600181815250505f61233482868686612c7c565b9050612344856020015182613487565b505b5f6003811115612359576123586158c6565b5b817f78615f01112d8477ae8d6fc9a510bcfdf935bc3ab7b02d0ff2316f5b44bd597a426040516123899190615478565b60405180910390a350505050565b5f5f6009549150600a5490509091565b60055481565b6123b5613d9c565b6123bd611949565b73ffffffffffffffffffffffffffffffffffffffff166123db612891565b73ffffffffffffffffffffffffffffffffffffffff161461243a576123fe612891565b6040517f118cdaa7000000000000000000000000000000000000000000000000000000008152600401612431919061568f565b60405180910390fd5b5f815190505f5f90505b818110156125a8575f8382815181106124605761245f615cfa565b5b602002602001015190505f60175f8381526020019081526020015f206040518060a00160405290815f82015f9054906101000a900460ff1660018111156124aa576124a96158c6565b5b60018111156124bc576124bb6158c6565b5b81526020015f820160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020015f820160159054906101000a900463ffffffff1663ffffffff1663ffffffff1681526020016001820154815260200160028201548152505090505f73ffffffffffffffffffffffffffffffffffffffff16816020015173ffffffffffffffffffffffffffffffffffffffff160361258e5750506125a3565b61259a82826003612917565b82600101925050505b612444565b505050565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60095481565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a1eea778612640612891565b6040518263ffffffff1660e01b815260040161265c919061568f565b5f6040518083038186803b158015612672575f5ffd5b505afa158015612684573d5f5f3e3d5ffd5b505050505f81036126ca576040517fc587916e0000000000000000000000000000000000000000000000000000000081526004016126c190616024565b60405180910390fd5b5f820361270c576040517fc587916e00000000000000000000000000000000000000000000000000000000815260040161270390616061565b60405180910390fd5b5f600454838361271c9190616252565b6127269190615e1a565b1461276e5780826004546040517f5756813500000000000000000000000000000000000000000000000000000000815260040161276593929190616293565b60405180910390fd5b7f66748bc112f6692a76bebecc33e69768508e27e414e3080f5c02295b6ea3cf3f600a5483600954846040516127a7949392919061632d565b60405180910390a181600a81905550806009819055506127ca81836004546133a2565b600b819055505050565b60605f825f015467ffffffffffffffff8111156127f4576127f3614fcf565b5b6040519080825280602002602001820160405280156128225781602001602082028036833780820191505090505b5090505f835f015490505f846001015490505f5f90505b82811015612885578184828151811061285557612854615cfa565b5b602002602001018181525050856003015f8381526020019081526020015f20600101549150806001019050612839565b50829350505050919050565b5f33905090565b612912838473ffffffffffffffffffffffffffffffffffffffff1663a9059cbb85856040516024016128cb929190616370565b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050613ddc565b505050565b5f5f5f600181111561292c5761292b6158c6565b5b845f01516001811115612942576129416158c6565b5b14905080156129e35760155f856060015181526020019081526020015f2091505f8460800151146129de576129c18460200151856080015160025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166128989092919063ffffffff16565b836080015160055f8282546129d69190615d93565b925050819055505b612a84565b60165f856060015181526020019081526020015f2091505f612a10856060015186608001516004546133a2565b90505f8114612a8257612a6985602001518260035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166128989092919063ffffffff16565b8060065f828254612a7a9190615d93565b925050819055505b505b612a8f858484613e77565b612a9882613f8b565b15612ae557612ae38460600151600d865f01516001811115612abd57612abc6158c6565b5b60ff1660028110612ad157612ad0615cfa565b5b60040201613f9990919063ffffffff16565b505b5050505050565b612af46116ff565b15612b2b576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505f5f8273ffffffffffffffffffffffffffffffffffffffff1663c415b95c6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612b9c573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612bc09190615f1c565b8373ffffffffffffffffffffffffffffffffffffffff166324a9d8536040518163ffffffff1660e01b8152600401602060405180830381865afa158015612c09573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612c2d91906163ab565b91509150817fd6aa07baf8485abf9d26fecf4c935d75b50a73e678db02b944bd3ac8759823005d807f1d2ff3fa6980aeeebca4e94965520da48983006e9b1115c1c853cbb10d943d005d505050565b5f600180811115612c9057612c8f6158c6565b5b845f01516001811115612ca657612ca56158c6565b5b14612ceb57835f01516040517f7cbdd081000000000000000000000000000000000000000000000000000000008152600401612ce291906163d6565b60405180910390fd5b5f5f8414612cfb57839050612d13565b612d10856060015186608001516004546133a2565b90505b5f5f612dc460035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401612d72919061568f565b602060405180830381865afa158015612d8d573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612db19190615d3b565b84600654612dbf9190615ed5565b6140c2565b9150915081612e2b5760035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040517f7bdf264d000000000000000000000000000000000000000000000000000000008152600401612e22919061568f565b60405180910390fd5b865f01516001811115612e4157612e406158c6565b5b88886020015173ffffffffffffffffffffffffffffffffffffffff167f7cde7559ef00f8c3d9244d3fde19e7500979d64f6f9c086a3836a4926abbe1658a606001518b6080015142604051612e9893929190616293565b60405180910390a45f612ead898989896140e8565b90505f8114612f1f578060055f828254612ec79190615d93565b92505081905550612f1e88602001518260025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166128989092919063ffffffff16565b5b81945050505050949350505050565b5f5f6001811115612f4257612f416158c6565b5b835f01516001811115612f5857612f576158c6565b5b14612f9d57825f01516040517f7cbdd081000000000000000000000000000000000000000000000000000000008152600401612f9491906163d6565b60405180910390fd5b8260800151600554612faf9190615ed5565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401613009919061568f565b602060405180830381865afa158015613024573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906130489190615d3b565b10156130ac5760025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040517f7bdf264d0000000000000000000000000000000000000000000000000000000081526004016130a3919061568f565b60405180910390fd5b825f015160018111156130c2576130c16158c6565b5b84846020015173ffffffffffffffffffffffffffffffffffffffff167f7cde7559ef00f8c3d9244d3fde19e7500979d64f6f9c086a3836a4926abbe165866060015187608001514260405161311993929190616293565b60405180910390a45f61312d858585614376565b90505f8114613178575f5f6131406145d4565b613148613183565b915091508260065f82825461315d9190615d93565b925050819055506131758787602001518585856145fc565b50505b5f9150509392505050565b5f7f1d2ff3fa6980aeeebca4e94965520da48983006e9b1115c1c853cbb10d943d005c905090565b5f5f83036131e5576040517f0da93b9700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6131ef848461478d565b156131fc575f9050613383565b5f846003015f8481526020019081526020015f20905061321a614f6f565b5f8403613271575f866001015490508587600101819055505f81146132545785876003015f8381526020019081526020015f205f01819055505b60405180604001604052805f815260200182815250915050613328565b613298826040518060400160405290815f82015481526020016001820154815250506147ed565b806132a65750856001015484145b6132dc576040517f182dfca500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8260010154905060405180604001604052808681526020018281525091505f811461331d5785876003015f8381526020019081526020015f205f01819055505b858360010181905550505b5f81602001510361333d578486600201819055505b80866003015f8781526020019081526020015f205f820151815f015560208201518160010155905050855f015f815461337590615e4a565b919050819055506001925050505b9392505050565b5f61339a838385600201546131ab565b905092915050565b5f5f83850290505f5f19858709828110838203039150505f81036133da578382816133d0576133cf615ded565b5b0492505050613480565b8084116133f9576133f86133f35f86146012601161480b565b614824565b5b5f8486880990508281118203915080830392505f855f038616905080860495508084049350600181825f0304019050808302841793505f600287600302189050808702600203810290508087026002038102905080870260020381029050808702600203810290508087026002038102905080870260020381029050808502955050505050505b9392505050565b5f600654826134969190615ed5565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016134f0919061568f565b602060405180830381865afa15801561350b573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061352f9190615d3b565b6135399190615d93565b90505f811461358f5761358e838260035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166128989092919063ffffffff16565b5b505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff16148061364157507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16613628614835565b73ffffffffffffffffffffffffffffffffffffffff1614155b15613678576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b613682611949565b73ffffffffffffffffffffffffffffffffffffffff166136a0612891565b73ffffffffffffffffffffffffffffffffffffffff16146136ff576136c3612891565b6040517f118cdaa70000000000000000000000000000000000000000000000000000000081526004016136f6919061568f565b60405180910390fd5b50565b8173ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa92505050801561376a57506040513d601f19601f820116820180604052508101906137679190616419565b60015b6137ab57816040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526004016137a2919061568f565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b811461381157806040517faa1d49a400000000000000000000000000000000000000000000000000000000815260040161380891906157ca565b60405180910390fd5b61381b8383614888565b505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff16146138a5576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300905090565b5f5f8403613908576040517f0da93b9700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b613912858561478d565b1561393457846003015f8581526020019081526020015f205f01549050613aa2565b846001015484108061394b575061394a85613f8b565b5b15613958575f9050613aa2565b84600201548411156139705784600201549050613aa2565b5f82036139a9576040517fab7e639300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6139b486856148fa565b905084811015613a24575b5f81141580156139d657505f836001900393508314155b15613a1f57856003015f8281526020019081526020015f2060010154905084811115613a1a57856003015f8281526020019081526020015f205f0154915050613aa2565b6139bf565b613a70565b5b5f8114158015613a3c57505f836001900393508314155b15613a6f57856003015f8281526020019081526020015f205f0154905084811015613a6a5780915050613aa2565b613a25565b5b6040517f80ce60d100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b949350505050565b5f5f8403613ae4576040517f0da93b9700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b613aee858561478d565b15613b1057846003015f8581526020019081526020015f205f01549050613c7e565b8460010154841180613b275750613b2685613f8b565b5b15613b34575f9050613c7e565b8460020154841015613b4c5784600201549050613c7e565b5f8203613b85576040517fab7e639300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f613b9086856148fa565b905084811115613c00575b5f8114158015613bb257505f836001900393508314155b15613bfb57856003015f8281526020019081526020015f2060010154905080851115613bf657856003015f8281526020019081526020015f205f0154915050613c7e565b613b9b565b613c4c565b5b5f8114158015613c1857505f836001900393508314155b15613c4b57856003015f8281526020019081526020015f205f0154905080851015613c465780915050613c7e565b613c01565b5b6040517f80ce60d100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b949350505050565b5f7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00905090565b613cb5614995565b613cbd6149d5565b565b613cc7612aec565b5f613cd06138a7565b90506001815f015f6101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258613d16612891565b604051613d23919061568f565b60405180910390a150565b613d36613d9c565b5f613d3f6138a7565b90505f815f015f6101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa613d84612891565b604051613d91919061568f565b60405180910390a150565b613da46116ff565b613dda576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f5f60205f8451602086015f885af180613dfb576040513d5f823e3d81fd5b3d92505f519150505f8214613e14576001811415613e2f565b5f8473ffffffffffffffffffffffffffffffffffffffff163b145b15613e7157836040517f5274afe7000000000000000000000000000000000000000000000000000000008152600401613e68919061568f565b60405180910390fd5b50505050565b613e8a8382613f9990919063ffffffff16565b613ecb57826040517ffcfdf902000000000000000000000000000000000000000000000000000000008152600401613ec29190615478565b60405180910390fd5b60175f8481526020019081526020015f205f5f82015f6101000a81549060ff02191690555f820160016101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690555f820160156101000a81549063ffffffff0219169055600182015f9055600282015f90555050816003811115613f4e57613f4d6158c6565b5b837f78615f01112d8477ae8d6fc9a510bcfdf935bc3ab7b02d0ff2316f5b44bd597a42604051613f7e9190615478565b60405180910390a3505050565b5f5f825f0154149050919050565b5f5f8203613fd3576040517f0da93b9700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b613fdd838361478d565b613fe9575f90506140bc565b5f836003015f8481526020019081526020015f2090505f846003015f835f015481526020019081526020015f2090505f856003015f846001015481526020019081526020015f20905082600101548260010181905550825f0154815f01819055508486600101540361406357826001015486600101819055505b8486600201540361407b57825f015486600201819055505b856003015f8681526020019081526020015f205f5f82015f9055600182015f90555050855f015f81546140ad90616444565b91905081905550600193505050505b92915050565b5f5f838311156140d7575f5f915091506140e1565b6001838503915091505b9250929050565b5f5f5f600d5f6001811115614100576140ff6158c6565b5b60ff166002811061411457614113615cfa565b5b6004020190505b61412481613f8b565b614363575f61413c5f83614a0590919063ffffffff16565b905086606001518111156141505750614363565b5f60155f8381526020019081526020015f20905061416d82614a86565b5f87146141c7575f61418d85896141849190615d93565b600454856133a2565b9050600a548161419d9190615e1a565b816141a89190615d93565b9050808960800181815250505f81036141c5575050505050614366565b505b5f6141d182614aac565b90505b5f8114614305575f6141ef5f84614a0590919063ffffffff16565b90505f60175f8381526020019081526020015f2090505f5f5f6142168f8f87878c8c614ab8565b9250925092505f61422a89846004546133a2565b905061424086858361423a6145d4565b866145fc565b828c61424c9190615ed5565b9b50808b61425a9190615ed5565b9a505f8f60800151148061427957505f8d61427490616444565b9d508d145b156142f35761428788613f8b565b156142e35761429f898b613f9990919063ffffffff16565b6142e2575f896040517f9f1cfdfe0000000000000000000000000000000000000000000000000000000081526004016142d992919061646b565b60405180910390fd5b5b5050505050505050505050614366565b866001900396505050505050506141d4565b6143188385613f9990919063ffffffff16565b61435b575f836040517f9f1cfdfe00000000000000000000000000000000000000000000000000000000815260040161435292919061646b565b60405180910390fd5b50505061411b565b50505b61436e614bf1565b949350505050565b5f5f5f60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600454915091505f600d6001808111156143b8576143b76158c6565b5b60ff16600281106143cc576143cb615cfa565b5b6004020190505b6143dc81613f8b565b6145c1575f6143f45f83614a0590919063ffffffff16565b9050866060015181101561440857506145c1565b5f60165f8381526020019081526020015f20905061442582614a86565b5f61442f82614aac565b90505b5f8114614562575f61444d5f84614a0590919063ffffffff16565b90505f60175f8381526020019081526020015f2090505f5f6144738e8e86868b8b614ab8565b50915091506144a382828c73ffffffffffffffffffffffffffffffffffffffff166128989092919063ffffffff16565b6144ae87828b6133a2565b8b6144b99190615ed5565b9a505f8d6080015114806144d857505f8c6144d390616444565b9c508c145b15614552576144e686613f8b565b15614543576144fe8789613f9990919063ffffffff16565b614542576001876040517f9f1cfdfe00000000000000000000000000000000000000000000000000000000815260040161453992919061646b565b60405180910390fd5b5b505050505050505050506145c5565b8460019003945050505050614432565b6145758385613f9990919063ffffffff16565b6145b9576001836040517f9f1cfdfe0000000000000000000000000000000000000000000000000000000081526004016145b092919061646b565b60405180910390fd5b5050506143d3565b5050505b6145cd614bf1565b9392505050565b5f7fd6aa07baf8485abf9d26fecf4c935d75b50a73e678db02b944bd3ac8759823005c905090565b5f8163ffffffff160361465a57614655848460035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166128989092919063ffffffff16565b614786565b5f61466e848363ffffffff166127106133a2565b90505f818561467d9190615d93565b90508373ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff16887f9814b2b2f43f9ea51ad7e8760fda094ffa3012bfeb10c76a9656389af1895603888787876040516146e394939291906164c2565b60405180910390a4614737848360035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166128989092919063ffffffff16565b614783868260035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166128989092919063ffffffff16565b50505b5050505050565b5f5f820361479d575f90506147e7565b82600101548214806147e457506147e3836003015f8481526020019081526020015f206040518060400160405290815f82015481526020016001820154815250506147ed565b5b90505b92915050565b5f5f825f015114158061480457505f826020015114155b9050919050565b5f61481584614c2f565b82841802821890509392505050565b634e487b715f52806020526024601cfd5b5f6148617f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b614c3a565b5f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b61489182614c43565b8173ffffffffffffffffffffffffffffffffffffffff167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b60405160405180910390a25f815111156148ed576148e78282614d0c565b506148f6565b6148f5614d8c565b5b5050565b5f61491c83835f6002811061491257614911615cfa565b5b602002015161478d565b1561494057815f6002811061493457614933615cfa565b5b6020020151905061498f565b614962838360016002811061495857614957615cfa565b5b602002015161478d565b15614987578160016002811061497b5761497a615cfa565b5b6020020151905061498f565b826001015490505b92915050565b61499d614dc8565b6149d3576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b6149dd614995565b5f6149e66138a7565b90505f815f015f6101000a81548160ff02191690831515021790555050565b5f825f01548210614a42576040517f39e60f7300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f836001015490505f8390505b5f8114614a7b5780600190039050846003015f8381526020019081526020015f20600101549150614a4f565b508091505092915050565b807ffd0e5d4f9b88892d3b04349a0e2bc0d1359414c21932fcd7d5a523a6c0a5cd005d50565b5f815f01549050919050565b5f5f5f855f0160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16614af189608001518860020154614de6565b875f0160159054906101000a900463ffffffff168093508194508295505050505f5f5f6001811115614b2657614b256158c6565b5b8a5f01516001811115614b3c57614b3b6158c6565b5b14614b4857888b614b4b565b8a895b915091508681837f6a6896b1e6131e0b8ebae43fdc84cb0178a6eacdf4ee63f15dabae48729a8a038742604051614b83929190615c4e565b60405180910390a487600201548403614ba657614ba1895f88613e77565b614bb8565b83886002015f82825403925050819055505b89608001518403614bd2575f8a6080018181525050614be3565b838a60800181815103915081815250505b505096509650969350505050565b5f7ffd0e5d4f9b88892d3b04349a0e2bc0d1359414c21932fcd7d5a523a6c0a5cd005c90505f8114614c2c5780600781905550426008819055505b50565b5f8115159050919050565b5f819050919050565b5f8173ffffffffffffffffffffffffffffffffffffffff163b03614c9e57806040517f4c9c8ce3000000000000000000000000000000000000000000000000000000008152600401614c95919061568f565b60405180910390fd5b80614cca7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b614c3a565b5f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60605f5f8473ffffffffffffffffffffffffffffffffffffffff1684604051614d359190616549565b5f60405180830381855af49150503d805f8114614d6d576040519150601f19603f3d011682016040523d82523d5f602084013e614d72565b606091505b5091509150614d82858383614dfc565b9250505092915050565b5f341115614dc6576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f614dd1613c86565b5f0160089054906101000a900460ff16905090565b5f614df4828410848461480b565b905092915050565b606082614e1157614e0c82614e89565b614e81565b5f8251148015614e3757505f8473ffffffffffffffffffffffffffffffffffffffff163b145b15614e7957836040517f9996b315000000000000000000000000000000000000000000000000000000008152600401614e70919061568f565b60405180910390fd5b819050614e82565b5b9392505050565b5f81511115614e9b5780518082602001fd5b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040518060a001604052805f6001811115614eeb57614eea6158c6565b5b81526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020015f63ffffffff1681526020015f81526020015f81525090565b60405180606001604052805f73ffffffffffffffffffffffffffffffffffffffff1681526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020015f81525090565b60405180604001604052805f81526020015f81525090565b5f604051905090565b5f5ffd5b5f5ffd5b60028110614fa4575f5ffd5b50565b5f81359050614fb581614f98565b92915050565b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61500582614fbf565b810181811067ffffffffffffffff8211171561502457615023614fcf565b5b80604052505050565b5f615036614f87565b90506150428282614ffc565b919050565b5f67ffffffffffffffff82111561506157615060614fcf565b5b602082029050602081019050919050565b5f5ffd5b5f819050919050565b61508881615076565b8114615092575f5ffd5b50565b5f813590506150a38161507f565b92915050565b5f6150bb6150b684615047565b61502d565b905080838252602082019050602084028301858111156150de576150dd615072565b5b835b8181101561510757806150f38882615095565b8452602084019350506020810190506150e0565b5050509392505050565b5f82601f83011261512557615124614fbb565b5b81356151358482602086016150a9565b91505092915050565b5f5f6040838503121561515457615153614f90565b5b5f61516185828601614fa7565b925050602083013567ffffffffffffffff81111561518257615181614f94565b5b61518e85828601615111565b9150509250929050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b6151f381615076565b82525050565b5f61520483836151ea565b60208301905092915050565b5f602082019050919050565b5f615226826151c1565b61523081856151cb565b935061523b836151db565b805f5b8381101561526b57815161525288826151f9565b975061525d83615210565b92505060018101905061523e565b5085935050505092915050565b5f615283838361521c565b905092915050565b5f602082019050919050565b5f6152a182615198565b6152ab81856151a2565b9350836020820285016152bd856151b2565b805f5b858110156152f857848403895281516152d98582615278565b94506152e48361528b565b925060208a019950506001810190506152c0565b50829750879550505050505092915050565b5f6020820190508181035f8301526153228184615297565b905092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6153538261532a565b9050919050565b5f61536482615349565b9050919050565b6153748161535a565b811461537e575f5ffd5b50565b5f8135905061538f8161536b565b92915050565b61539e81615349565b81146153a8575f5ffd5b50565b5f813590506153b981615395565b92915050565b5f5f5f606084860312156153d6576153d5614f90565b5b5f6153e386828701615381565b93505060206153f4868287016153ab565b925050604061540586828701615095565b9150509250925092565b5f5f6040838503121561542557615424614f90565b5b5f615432858286016153ab565b925050602083013567ffffffffffffffff81111561545357615452614f94565b5b61545f85828601615111565b9150509250929050565b61547281615076565b82525050565b5f60208201905061548b5f830184615469565b92915050565b5f82825260208201905092915050565b5f6154ab826151c1565b6154b58185615491565b93506154c0836151db565b805f5b838110156154f05781516154d788826151f9565b97506154e283615210565b9250506001810190506154c3565b5085935050505092915050565b5f6040820190508181035f83015261551581856154a1565b9050818103602083015261552981846154a1565b90509392505050565b5f5ffd5b5f63ffffffff82169050919050565b61554e81615536565b8114615558575f5ffd5b50565b5f8135905061556981615545565b92915050565b5f60a0828403121561558457615583615532565b5b61558e60a061502d565b90505f61559d84828501614fa7565b5f8301525060206155b0848285016153ab565b60208301525060406155c48482850161555b565b60408301525060606155d884828501615095565b60608301525060806155ec84828501615095565b60808301525092915050565b60038110615604575f5ffd5b50565b5f81359050615615816155f8565b92915050565b5f5f5f5f610100858703121561563457615633614f90565b5b5f6156418782880161556f565b94505060a061565287828801615607565b93505060c061566387828801615095565b92505060e061567487828801615095565b91505092959194509250565b61568981615349565b82525050565b5f6020820190506156a25f830184615680565b92915050565b5f5ffd5b5f67ffffffffffffffff8211156156c6576156c5614fcf565b5b6156cf82614fbf565b9050602081019050919050565b828183375f83830152505050565b5f6156fc6156f7846156ac565b61502d565b905082815260208101848484011115615718576157176156a8565b5b6157238482856156dc565b509392505050565b5f82601f83011261573f5761573e614fbb565b5b813561574f8482602086016156ea565b91505092915050565b5f5f6040838503121561576e5761576d614f90565b5b5f61577b858286016153ab565b925050602083013567ffffffffffffffff81111561579c5761579b614f94565b5b6157a88582860161572b565b9150509250929050565b5f819050919050565b6157c4816157b2565b82525050565b5f6020820190506157dd5f8301846157bb565b92915050565b5f8115159050919050565b6157f7816157e3565b82525050565b5f6020820190506158105f8301846157ee565b92915050565b5f8190508260206002028201111561583157615830615072565b5b92915050565b5f5f5f5f60a0858703121561584f5761584e614f90565b5b5f61585c87828801614fa7565b945050602061586d87828801615095565b935050604061587e87828801615816565b925050608061588f87828801615095565b91505092959194509250565b5f602082840312156158b0576158af614f90565b5b5f6158bd84828501615095565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b60028110615904576159036158c6565b5b50565b5f819050615914826158f3565b919050565b5f61592382615907565b9050919050565b61593381615919565b82525050565b61594281615349565b82525050565b61595181615536565b82525050565b60a082015f82015161596b5f85018261592a565b50602082015161597e6020850182615939565b5060408201516159916040850182615948565b5060608201516159a460608501826151ea565b5060808201516159b760808501826151ea565b50505050565b5f60a0820190506159d05f830184615957565b92915050565b5f819050919050565b5f6159f96159f46159ef8461532a565b6159d6565b61532a565b9050919050565b5f615a0a826159df565b9050919050565b5f615a1b82615a00565b9050919050565b615a2b81615a11565b82525050565b5f602082019050615a445f830184615a22565b92915050565b5f5f5f5f5f60a08688031215615a6357615a62614f90565b5b5f615a70888289016153ab565b9550506020615a81888289016153ab565b9450506040615a92888289016153ab565b9350506060615aa388828901615095565b9250506080615ab488828901615095565b9150509295509295909350565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f615af382615ac1565b615afd8185615acb565b9350615b0d818560208601615adb565b615b1681614fbf565b840191505092915050565b5f6020820190508181035f830152615b398184615ae9565b905092915050565b615b4a816157e3565b8114615b54575f5ffd5b50565b5f81359050615b6581615b41565b92915050565b5f60208284031215615b8057615b7f614f90565b5b5f615b8d84828501615b57565b91505092915050565b615b9f81615a11565b82525050565b606082015f820151615bb95f850182615b96565b506020820151615bcc6020850182615b96565b506040820151615bdf60408501826151ea565b50505050565b5f606082019050615bf85f830184615ba5565b92915050565b5f5f5f60e08486031215615c1557615c14614f90565b5b5f615c228682870161556f565b93505060a0615c3386828701615095565b92505060c0615c4486828701615095565b9150509250925092565b5f604082019050615c615f830185615469565b615c6e6020830184615469565b9392505050565b5f60208284031215615c8a57615c89614f90565b5b5f82013567ffffffffffffffff811115615ca757615ca6614f94565b5b615cb384828501615111565b91505092915050565b5f5f60408385031215615cd257615cd1614f90565b5b5f615cdf85828601615095565b9250506020615cf085828601615095565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f81519050615d358161507f565b92915050565b5f60208284031215615d5057615d4f614f90565b5b5f615d5d84828501615d27565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f615d9d82615076565b9150615da883615076565b9250828203905081811115615dc057615dbf615d66565b5b92915050565b5f604082019050615dd95f830185615469565b615de66020830184615680565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f615e2482615076565b9150615e2f83615076565b925082615e3f57615e3e615ded565b5b828206905092915050565b5f615e5482615076565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203615e8657615e85615d66565b5b600182019050919050565b615e9a81615919565b82525050565b5f606082019050615eb35f830186615e91565b615ec06020830185615469565b615ecd6040830184615469565b949350505050565b5f615edf82615076565b9150615eea83615076565b9250828201905080821115615f0257615f01615d66565b5b92915050565b5f81519050615f1681615395565b92915050565b5f60208284031215615f3157615f30614f90565b5b5f615f3e84828501615f08565b91505092915050565b7f726f757465720000000000000000000000000000000000000000000000000000815250565b5f602082019050615f7f5f8301615f47565b919050565b7f71756f7465000000000000000000000000000000000000000000000000000000815250565b5f602082019050615fbc5f8301615f84565b919050565b7f6261736500000000000000000000000000000000000000000000000000000000815250565b5f602082019050615ff95f8301615fc1565b919050565b7f7469636b53697a65000000000000000000000000000000000000000000000000815250565b5f6020820190506160365f8301615ffe565b919050565b7f6c6f7453697a6500000000000000000000000000000000000000000000000000815250565b5f6020820190506160735f830161603b565b919050565b5f60ff82169050919050565b61608d81616078565b8114616097575f5ffd5b50565b5f815190506160a881616084565b92915050565b5f602082840312156160c3576160c2614f90565b5b5f6160d08482850161609a565b91505092915050565b5f8160011c9050919050565b5f5f8291508390505b600185111561612e5780860481111561610a57616109615d66565b5b60018516156161195780820291505b8081029050616127856160d9565b94506160ee565b94509492505050565b5f826161465760019050616201565b81616153575f9050616201565b81600181146161695760028114616173576161a2565b6001915050616201565b60ff84111561618557616184615d66565b5b8360020a91508482111561619c5761619b615d66565b5b50616201565b5060208310610133831016604e8410600b84101617156161d75782820a9050838111156161d2576161d1615d66565b5b616201565b6161e484848460016160e5565b925090508184048111156161fb576161fa615d66565b5b81810290505b9392505050565b5f61621282615076565b915061621d83616078565b925061624a7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8484616137565b905092915050565b5f61625c82615076565b915061626783615076565b925082820261627581615076565b9150828204841483151761628c5761628b615d66565b5b5092915050565b5f6060820190506162a65f830186615469565b6162b36020830185615469565b6162c06040830184615469565b949350505050565b5f819050919050565b5f67ffffffffffffffff82169050919050565b5f6162fe6162f96162f4846162c8565b6159d6565b6162d1565b9050919050565b61630e816162e4565b82525050565b5f6020820190506163275f830184616305565b92915050565b5f6080820190506163405f830187615469565b61634d6020830186615469565b61635a6040830185615469565b6163676060830184615469565b95945050505050565b5f6040820190506163835f830185615680565b6163906020830184615469565b9392505050565b5f815190506163a581615545565b92915050565b5f602082840312156163c0576163bf614f90565b5b5f6163cd84828501616397565b91505092915050565b5f6020820190506163e95f830184615e91565b92915050565b6163f8816157b2565b8114616402575f5ffd5b50565b5f81519050616413816163ef565b92915050565b5f6020828403121561642e5761642d614f90565b5b5f61643b84828501616405565b91505092915050565b5f61644e82615076565b91505f82036164605761645f615d66565b5b600182039050919050565b5f60408201905061647e5f830185615e91565b61648b6020830184615469565b9392505050565b5f6164ac6164a76164a284615536565b6159d6565b615076565b9050919050565b6164bc81616492565b82525050565b5f6080820190506164d55f830187615469565b6164e260208301866164b3565b6164ef6040830185615469565b6164fc6060830184615469565b95945050505050565b5f81519050919050565b5f81905092915050565b5f61652382616505565b61652d818561650f565b935061653d818560208601615adb565b80840191505092915050565b5f6165548284616519565b91508190509291505056fea2646970667358221220995d36b9373a93547b3d433298fea200b523885a5836367e9caf9e1a758e4a3064736f6c634300081c0033"

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
