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

// RouterImplMetaData contains all meta data concerning the RouterImpl contract.
var RouterImplMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"CROSS_DEX\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WXCROSS\",\"outputs\":[{\"internalType\":\"contractIWETH\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"orderIds\",\"type\":\"uint256[]\"}],\"name\":\"cancel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxMatchCount\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"name\":\"isPair\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"enumIPair.LimitConstraints\",\"name\":\"constraints\",\"type\":\"uint8\"},{\"internalType\":\"uint256[2]\",\"name\":\"adjacent\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"_maxMatchCount\",\"type\":\"uint256\"}],\"name\":\"limitBuy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"enumIPair.LimitConstraints\",\"name\":\"constraints\",\"type\":\"uint8\"},{\"internalType\":\"uint256[2]\",\"name\":\"adjacent\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"_maxMatchCount\",\"type\":\"uint256\"}],\"name\":\"limitSell\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxMatchCount\",\"type\":\"uint256\"}],\"name\":\"marketBuy\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxMatchCount\",\"type\":\"uint256\"}],\"name\":\"marketSell\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxMatchCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"quote\",\"type\":\"address\"}],\"name\":\"PairAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"name\":\"PairRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"RouterInitializeData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"RouterInvalidPairAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterInvalidValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"b820d8aa": "CROSS_DEX()",
		"ad3cb1cc": "UPGRADE_INTERFACE_VERSION()",
		"3499ee7f": "WXCROSS()",
		"f9f629c5": "cancel(address,uint256[])",
		"fe4b84df": "initialize(uint256)",
		"e5e31b13": "isPair(address)",
		"637d2635": "limitBuy(address,uint256,uint256,uint8,uint256[2],uint256)",
		"473bde60": "limitSell(address,uint256,uint256,uint8,uint256[2],uint256)",
		"bc58e938": "marketBuy(address,uint256,uint256)",
		"26580d11": "marketSell(address,uint256,uint256)",
		"2c76535e": "maxMatchCount()",
		"52d1902d": "proxiableUUID()",
		"4f1ef286": "upgradeToAndCall(address,bytes)",
	},
	Bin: "0x60a060405230608052348015610013575f5ffd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6080516124d36100f95f395f8181610de401528181610e0d0152610fef01526124d35ff3fe6080604052600436106100c2575f3560e01c8063637d26351161007c578063bc58e93811610057578063bc58e93814610206578063e5e31b1314610219578063f9f629c514610248578063fe4b84df14610267575f5ffd5b8063637d263514610198578063ad3cb1cc146101ab578063b820d8aa146101e8575f5ffd5b806326580d11146100ec5780632c76535e146100ff5780633499ee7f14610127578063473bde601461015e5780634f1ef2861461017157806352d1902d14610184575f5ffd5b366100e85747156100e65760405163f85c3ef560e01b815260040160405180910390fd5b005b5f5ffd5b6100e66100fa3660046113b9565b610286565b34801561010a575f5ffd5b5061011460025481565b6040519081526020015b60405180910390f35b348015610132575f5ffd5b50600154610146906001600160a01b031681565b6040516001600160a01b03909116815260200161011e565b61011461016c366004611459565b6104a1565b6100e661017f3660046114f6565b6106c0565b34801561018f575f5ffd5b506101146106df565b6101146101a6366004611459565b6106fa565b3480156101b6575f5ffd5b506101db604051806040016040528060058152602001640352e302e360dc1b81525081565b60405161011e919061159d565b3480156101f3575f5ffd5b505f54610146906001600160a01b031681565b6100e66102143660046113b9565b610910565b348015610224575f5ffd5b506102386102333660046115d2565b610a57565b604051901515815260200161011e565b348015610253575f5ffd5b506100e66102623660046115ed565b610ad4565b348015610272575f5ffd5b506100e6610281366004611670565b610b69565b61028e610d0c565b8261029881610a57565b6102c557604051633b1c760160e21b81526001600160a01b03821660048201526024015b60405180910390fd5b5f3390505f856001600160a01b031663c3f909d46040518163ffffffff1660e01b8152600401606060405180830381865afa158015610306573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061032a9190611687565b6020810151600154919250906001600160a01b03908116908216036103ab5760015460405163755edd1760e01b81526001600160a01b0389811660048301529091169063755edd179088906024015f604051808303818588803b15801561038f575f5ffd5b505af11580156103a1573d5f5f3e3d5ffd5b50505050506103c0565b6103c06001600160a01b038216848989610d43565b6040805160a081019091525f9080825b8152602001856001600160a01b031681526020015f63ffffffff1681526020015f81526020015f8152509050876001600160a01b031663948faf9082896104168a610da3565b6040518463ffffffff1660e01b815260040161043493929190611752565b5f604051808303815f87803b15801561044b575f5ffd5b505af115801561045d573d5f5f3e3d5ffd5b5050505050505050475f146104855760405163f85c3ef560e01b815260040160405180910390fd5b5061049c60015f51602061247e5f395f51905f5255565b505050565b5f6104aa610d0c565b866104b481610a57565b6104dc57604051633b1c760160e21b81526001600160a01b03821660048201526024016102bc565b5f3390505f896001600160a01b031663c3f909d46040518163ffffffff1660e01b8152600401606060405180830381865afa15801561051d573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105419190611687565b6020810151600154919250906001600160a01b03908116908216036105c25760015460405163755edd1760e01b81526001600160a01b038d811660048301529091169063755edd17908b906024015f604051808303818588803b1580156105a6575f5ffd5b505af11580156105b8573d5f5f3e3d5ffd5b50505050506105d7565b6105d76001600160a01b038216848d8c610d43565b6040805160a0810182525f8082526001600160a01b03808716602084015292820152606081018c9052608081018b9052908c16633e7024de828b8b61061b8c610da3565b6040518563ffffffff1660e01b815260040161063a9493929190611771565b6020604051808303815f875af1158015610656573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061067a91906117cf565b955050505050471561069f5760405163f85c3ef560e01b815260040160405180910390fd5b506106b660015f51602061247e5f395f51905f5255565b9695505050505050565b6106c8610dd9565b6106d182610e7f565b6106db8282610f28565b5050565b5f6106e8610fe4565b505f51602061245e5f395f51905f5290565b5f610703610d0c565b8661070d81610a57565b61073557604051633b1c760160e21b81526001600160a01b03821660048201526024016102bc565b5f3390505f896001600160a01b031663c3f909d46040518163ffffffff1660e01b8152600401606060405180830381865afa158015610776573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061079a9190611687565b90505f815f015190505f6107b38b8b856040015161102d565b6001549091506001600160a01b039081169083160361082e5760015460405163755edd1760e01b81526001600160a01b038e811660048301529091169063755edd179083906024015f604051808303818588803b158015610812575f5ffd5b505af1158015610824573d5f5f3e3d5ffd5b5050505050610843565b6108436001600160a01b038316858e84610d43565b6040805160a081018252600181526001600160a01b0380871660208301525f92820192909252606081018d9052608081018c9052908d16633e7024de828c8c61088b8d610da3565b6040518563ffffffff1660e01b81526004016108aa9493929190611771565b6020604051808303815f875af11580156108c6573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108ea91906117cf565b96505050505050471561069f5760405163f85c3ef560e01b815260040160405180910390fd5b610918610d0c565b8261092281610a57565b61094a57604051633b1c760160e21b81526001600160a01b03821660048201526024016102bc565b5f3390505f856001600160a01b031663c3f909d46040518163ffffffff1660e01b8152600401606060405180830381865afa15801561098b573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109af9190611687565b8051600154919250906001600160a01b0390811690821603610a2d5760015460405163755edd1760e01b81526001600160a01b0389811660048301529091169063755edd179088906024015f604051808303818588803b158015610a11575f5ffd5b505af1158015610a23573d5f5f3e3d5ffd5b5050505050610a42565b610a426001600160a01b038216848989610d43565b6040805160a081019091525f908060016103d0565b5f8054604051630827057360e01b81526001600160a01b03848116600483015283921690630827057390602401602060405180830381865afa158015610a9f573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ac391906117e6565b6001600160a01b0316141592915050565b82610ade81610a57565b610b0657604051633b1c760160e21b81526001600160a01b03821660048201526024016102bc565b60405163f9f629c560e01b81526001600160a01b0385169063f9f629c590610b3690339087908790600401611801565b5f604051808303815f87803b158015610b4d575f5ffd5b505af1158015610b5f573d5f5f3e3d5ffd5b5050505050505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff16159067ffffffffffffffff165f81158015610bae5750825b90505f8267ffffffffffffffff166001148015610bca5750303b155b905081158015610bd8575080155b15610bf65760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff191660011785558315610c2057845460ff60401b1916600160401b1785555b855f03610c525760405163b927e50b60e01b81526c1b585e13585d18da10dbdd5b9d609a1b60048201526024016102bc565b5f80546001600160a01b03191633179055604051610c6f90611398565b604051809103905ff080158015610c88573d5f5f3e3d5ffd5b50600180546001600160a01b0319166001600160a01b03929092169190911790556002869055610cb66110e4565b610cbe6110ec565b8315610d0457845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050565b5f51602061247e5f395f51905f52805460011901610d3d57604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b604080516001600160a01b0385811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180516001600160e01b03166323b872dd60e01b179052610d9d9085906110fc565b50505050565b5f811580610db2575060025482115b610dbc5781610dc0565b6002545b92915050565b60015f51602061247e5f395f51905f5255565b306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480610e5f57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316610e535f51602061245e5f395f51905f52546001600160a01b031690565b6001600160a01b031614155b15610e7d5760405163703e46dd60e11b815260040160405180910390fd5b565b5f5f9054906101000a90046001600160a01b03166001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610ece573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ef291906117e6565b6001600160a01b0316336001600160a01b031614610f255760405163118cdaa760e01b81523360048201526024016102bc565b50565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015610f82575060408051601f3d908101601f19168201909252610f7f918101906117cf565b60015b610faa57604051634c9c8ce360e01b81526001600160a01b03831660048201526024016102bc565b5f51602061245e5f395f51905f528114610fda57604051632a87526960e21b8152600481018290526024016102bc565b61049c8383611168565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610e7d5760405163703e46dd60e11b815260040160405180910390fd5b5f838302815f1985870982811083820303915050805f036110615783828161105757611057611849565b04925050506110dd565b8084116110785761107860038515026011186111bd565b5f848688095f868103871696879004966002600389028118808a02820302808a02820302808a02820302808a02820302808a02820302808a02909103029181900381900460010186841190950394909402919094039290920491909117919091029150505b9392505050565b610e7d6111ce565b6110f46111ce565b610e7d611217565b5f5f60205f8451602086015f885af18061111b576040513d5f823e3d81fd5b50505f513d9150811561113257806001141561113f565b6001600160a01b0384163b155b15610d9d57604051635274afe760e01b81526001600160a01b03851660048201526024016102bc565b6111718261121f565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a28051156111b55761049c8282611282565b6106db6112f4565b634e487b715f52806020526024601cfd5b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff16610e7d57604051631afcd79f60e31b815260040160405180910390fd5b610dc66111ce565b806001600160a01b03163b5f0361125457604051634c9c8ce360e01b81526001600160a01b03821660048201526024016102bc565b5f51602061245e5f395f51905f5280546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b03168460405161129e919061185d565b5f60405180830381855af49150503d805f81146112d6576040519150601f19603f3d011682016040523d82523d5f602084013e6112db565b606091505b50915091506112eb858383611313565b95945050505050565b3415610e7d5760405163b398979f60e01b815260040160405180910390fd5b606082611328576113238261136f565b6110dd565b815115801561133f57506001600160a01b0384163b155b1561136857604051639996b31560e01b81526001600160a01b03851660048201526024016102bc565b50806110dd565b80511561137f5780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b610bea8061187483390190565b6001600160a01b0381168114610f25575f5ffd5b5f5f5f606084860312156113cb575f5ffd5b83356113d6816113a5565b95602085013595506040909401359392505050565b634e487b7160e01b5f52604160045260245ffd5b6040805190810167ffffffffffffffff81118282101715611422576114226113eb565b60405290565b604051601f8201601f1916810167ffffffffffffffff81118282101715611451576114516113eb565b604052919050565b5f5f5f5f5f5f60e0878903121561146e575f5ffd5b8635611479816113a5565b9550602087013594506040870135935060608701356003811061149a575f5ffd5b9250609f870188136114aa575f5ffd5b6114b26113ff565b8060c089018a8111156114c3575f5ffd5b60808a015b818110156114e05780358452602093840193016114c8565b50979a9699509497509295939493359392505050565b5f5f60408385031215611507575f5ffd5b8235611512816113a5565b9150602083013567ffffffffffffffff81111561152d575f5ffd5b8301601f8101851361153d575f5ffd5b803567ffffffffffffffff811115611557576115576113eb565b61156a601f8201601f1916602001611428565b81815286602083850101111561157e575f5ffd5b816020840160208301375f602083830101528093505050509250929050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b5f602082840312156115e2575f5ffd5b81356110dd816113a5565b5f5f5f604084860312156115ff575f5ffd5b833561160a816113a5565b9250602084013567ffffffffffffffff811115611625575f5ffd5b8401601f81018613611635575f5ffd5b803567ffffffffffffffff81111561164b575f5ffd5b8660208260051b840101111561165f575f5ffd5b939660209190910195509293505050565b5f60208284031215611680575f5ffd5b5035919050565b5f6060828403128015611698575f5ffd5b506040516060810167ffffffffffffffff811182821017156116bc576116bc6113eb565b60405282516116ca816113a5565b815260208301516116da816113a5565b60208201526040928301519281019290925250919050565b634e487b7160e01b5f52602160045260245ffd5b805160028110611718576117186116f2565b82526020818101516001600160a01b03169083015260408082015163ffffffff169083015260608082015190830152608090810151910152565b60e081016117608286611706565b60a082019390935260c00152919050565b61012081016117808287611706565b60038510611790576117906116f2565b8460a083015260c08201845f5b60028110156117bc57815183526020928301929091019060010161179d565b5050508261010083015295945050505050565b5f602082840312156117df575f5ffd5b5051919050565b5f602082840312156117f6575f5ffd5b81516110dd816113a5565b6001600160a01b038416815260406020820181905281018290525f6001600160fb1b0383111561182f575f5ffd5b8260051b8085606085013791909101606001949350505050565b634e487b7160e01b5f52601260045260245ffd5b5f82518060208501845e5f92019182525091905056fe60a060405234801561000f575f5ffd5b506040518060400160405280601781526020017f43726f73734445582057726170706564205843726f7373000000000000000000815250604051806040016040528060068152602001655843524f535360d01b81525081600390816100749190610135565b5060046100818282610135565b5061008c9150503390565b6001600160a01b03166080526101ef565b634e487b7160e01b5f52604160045260245ffd5b600181811c908216806100c557607f821691505b6020821081036100e357634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561013057805f5260205f20601f840160051c8101602085101561010e5750805b601f840160051c820191505b8181101561012d575f815560010161011a565b50505b505050565b81516001600160401b0381111561014e5761014e61009d565b6101628161015c84546100b1565b846100e9565b6020601f821160018114610194575f831561017d5750848201515b5f19600385901b1c1916600184901b17845561012d565b5f84815260208120601f198516915b828110156101c357878501518255602094850194600190920191016101a3565b50848210156101e057868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b6080516109dc61020e5f395f818161017401526103d701526109dc5ff3fe60806040526004361061009d575f3560e01c806332fe7b261161006257806332fe7b261461016357806370a08231146101ae578063755edd17146101e257806395d89b41146101f5578063a9059cbb14610209578063dd62ed3e14610228575f5ffd5b806306fdde03146100b2578063095ea7b3146100dc57806318160ddd1461010b57806323b872dd14610129578063313ce56714610148575f5ffd5b366100ae576100ac333461026c565b005b5f5ffd5b3480156100bd575f5ffd5b506100c66102a9565b6040516100d3919061082d565b60405180910390f35b3480156100e7575f5ffd5b506100fb6100f636600461087d565b610339565b60405190151581526020016100d3565b348015610116575f5ffd5b506002545b6040519081526020016100d3565b348015610134575f5ffd5b506100fb6101433660046108a5565b610352565b348015610153575f5ffd5b50604051601281526020016100d3565b34801561016e575f5ffd5b506101967f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016100d3565b3480156101b9575f5ffd5b5061011b6101c83660046108df565b6001600160a01b03165f9081526020819052604090205490565b6100ac6101f03660046108df565b610375565b348015610200575f5ffd5b506100c6610382565b348015610214575f5ffd5b506100fb61022336600461087d565b610391565b348015610233575f5ffd5b5061011b6102423660046108ff565b6001600160a01b039182165f90815260016020908152604080832093909416825291909152205490565b6001600160a01b03821661029a5760405163ec442f0560e01b81525f60048201526024015b60405180910390fd5b6102a55f838361039e565b5050565b6060600380546102b890610930565b80601f01602080910402602001604051908101604052809291908181526020018280546102e490610930565b801561032f5780601f106103065761010080835404028352916020019161032f565b820191905f5260205f20905b81548152906001019060200180831161031257829003601f168201915b5050505050905090565b5f33610346818585610466565b60019150505b92915050565b5f3361035f858285610473565b61036a8585856104ef565b506001949350505050565b61037f813461026c565b50565b6060600480546102b890610930565b5f336103468185856104ef565b6103a983838361054c565b6001600160a01b038216156104615760405163e5e31b1360e01b81526001600160a01b0383811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063e5e31b1390602401602060405180830381865afa15801561041c573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104409190610968565b6104615761044e8282610672565b6104616001600160a01b038316826106a6565b505050565b6104618383836001610732565b6001600160a01b038381165f908152600160209081526040808320938616835292905220545f198110156104e957818110156104db57604051637dc7a0d960e11b81526001600160a01b03841660048201526024810182905260448101839052606401610291565b6104e984848484035f610732565b50505050565b6001600160a01b03831661051857604051634b637e8f60e11b81525f6004820152602401610291565b6001600160a01b0382166105415760405163ec442f0560e01b81525f6004820152602401610291565b61046183838361039e565b6001600160a01b038316610576578060025f82825461056b9190610987565b909155506105e69050565b6001600160a01b0383165f90815260208190526040902054818110156105c85760405163391434e360e21b81526001600160a01b03851660048201526024810182905260448101839052606401610291565b6001600160a01b0384165f9081526020819052604090209082900390555b6001600160a01b03821661060257600280548290039055610620565b6001600160a01b0382165f9081526020819052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161066591815260200190565b60405180910390a3505050565b6001600160a01b03821661069b57604051634b637e8f60e11b81525f6004820152602401610291565b6102a5825f8361039e565b804710156106d05760405163cf47918160e01b815247600482015260248101829052604401610291565b5f5f836001600160a01b0316836040515f6040518083038185875af1925050503d805f811461071a576040519150601f19603f3d011682016040523d82523d5f602084013e61071f565b606091505b5091509150816104e9576104e981610804565b6001600160a01b03841661075b5760405163e602df0560e01b81525f6004820152602401610291565b6001600160a01b03831661078457604051634a1406b160e11b81525f6004820152602401610291565b6001600160a01b038085165f90815260016020908152604080832093871683529290522082905580156104e957826001600160a01b0316846001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516107f691815260200190565b60405180910390a350505050565b8051156108145780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b80356001600160a01b0381168114610878575f5ffd5b919050565b5f5f6040838503121561088e575f5ffd5b61089783610862565b946020939093013593505050565b5f5f5f606084860312156108b7575f5ffd5b6108c084610862565b92506108ce60208501610862565b929592945050506040919091013590565b5f602082840312156108ef575f5ffd5b6108f882610862565b9392505050565b5f5f60408385031215610910575f5ffd5b61091983610862565b915061092760208401610862565b90509250929050565b600181811c9082168061094457607f821691505b60208210810361096257634e487b7160e01b5f52602260045260245ffd5b50919050565b5f60208284031215610978575f5ffd5b815180151581146108f8575f5ffd5b8082018082111561034c57634e487b7160e01b5f52601160045260245ffdfea26469706673582212206a099c5f304857a02a0e1a1c0a84b343f0cbf1bb722357365781e1da01adcab164736f6c634300081c0033360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a26469706673582212205fe00fb3a4ba729ed609a9e5cc4d82c4eed9425cc9c68c20fafc9a61336544ff64736f6c634300081c0033",
}

// RouterImplABI is the input ABI used to generate the binding from.
// Deprecated: Use RouterImplMetaData.ABI instead.
var RouterImplABI = RouterImplMetaData.ABI

// Deprecated: Use RouterImplMetaData.Sigs instead.
// RouterImplFuncSigs maps the 4-byte function signature to its string representation.
var RouterImplFuncSigs = RouterImplMetaData.Sigs

// RouterImplBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RouterImplMetaData.Bin instead.
var RouterImplBin = RouterImplMetaData.Bin

// DeployRouterImpl deploys a new Ethereum contract, binding an instance of RouterImpl to it.
func DeployRouterImpl(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RouterImpl, error) {
	parsed, err := RouterImplMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RouterImplBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RouterImpl{RouterImplCaller: RouterImplCaller{contract: contract}, RouterImplTransactor: RouterImplTransactor{contract: contract}, RouterImplFilterer: RouterImplFilterer{contract: contract}}, nil
}

// RouterImpl is an auto generated Go binding around an Ethereum contract.
type RouterImpl struct {
	RouterImplCaller     // Read-only binding to the contract
	RouterImplTransactor // Write-only binding to the contract
	RouterImplFilterer   // Log filterer for contract events
}

// RouterImplCaller is an auto generated read-only Go binding around an Ethereum contract.
type RouterImplCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterImplTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RouterImplTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterImplFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RouterImplFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterImplSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RouterImplSession struct {
	Contract     *RouterImpl       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RouterImplCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RouterImplCallerSession struct {
	Contract *RouterImplCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// RouterImplTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RouterImplTransactorSession struct {
	Contract     *RouterImplTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// RouterImplRaw is an auto generated low-level Go binding around an Ethereum contract.
type RouterImplRaw struct {
	Contract *RouterImpl // Generic contract binding to access the raw methods on
}

// RouterImplCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RouterImplCallerRaw struct {
	Contract *RouterImplCaller // Generic read-only contract binding to access the raw methods on
}

// RouterImplTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RouterImplTransactorRaw struct {
	Contract *RouterImplTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRouterImpl creates a new instance of RouterImpl, bound to a specific deployed contract.
func NewRouterImpl(address common.Address, backend bind.ContractBackend) (*RouterImpl, error) {
	contract, err := bindRouterImpl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RouterImpl{RouterImplCaller: RouterImplCaller{contract: contract}, RouterImplTransactor: RouterImplTransactor{contract: contract}, RouterImplFilterer: RouterImplFilterer{contract: contract}}, nil
}

// NewRouterImplCaller creates a new read-only instance of RouterImpl, bound to a specific deployed contract.
func NewRouterImplCaller(address common.Address, caller bind.ContractCaller) (*RouterImplCaller, error) {
	contract, err := bindRouterImpl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RouterImplCaller{contract: contract}, nil
}

// NewRouterImplTransactor creates a new write-only instance of RouterImpl, bound to a specific deployed contract.
func NewRouterImplTransactor(address common.Address, transactor bind.ContractTransactor) (*RouterImplTransactor, error) {
	contract, err := bindRouterImpl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RouterImplTransactor{contract: contract}, nil
}

// NewRouterImplFilterer creates a new log filterer instance of RouterImpl, bound to a specific deployed contract.
func NewRouterImplFilterer(address common.Address, filterer bind.ContractFilterer) (*RouterImplFilterer, error) {
	contract, err := bindRouterImpl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RouterImplFilterer{contract: contract}, nil
}

// bindRouterImpl binds a generic wrapper to an already deployed contract.
func bindRouterImpl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RouterImplMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RouterImpl *RouterImplRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RouterImpl.Contract.RouterImplCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RouterImpl *RouterImplRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RouterImpl.Contract.RouterImplTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RouterImpl *RouterImplRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RouterImpl.Contract.RouterImplTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RouterImpl *RouterImplCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RouterImpl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RouterImpl *RouterImplTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RouterImpl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RouterImpl *RouterImplTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RouterImpl.Contract.contract.Transact(opts, method, params...)
}

// CROSSDEX is a free data retrieval call binding the contract method 0xb820d8aa.
//
// Solidity: function CROSS_DEX() view returns(address)
func (_RouterImpl *RouterImplCaller) CROSSDEX(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RouterImpl.contract.Call(opts, &out, "CROSS_DEX")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CROSSDEX is a free data retrieval call binding the contract method 0xb820d8aa.
//
// Solidity: function CROSS_DEX() view returns(address)
func (_RouterImpl *RouterImplSession) CROSSDEX() (common.Address, error) {
	return _RouterImpl.Contract.CROSSDEX(&_RouterImpl.CallOpts)
}

// CROSSDEX is a free data retrieval call binding the contract method 0xb820d8aa.
//
// Solidity: function CROSS_DEX() view returns(address)
func (_RouterImpl *RouterImplCallerSession) CROSSDEX() (common.Address, error) {
	return _RouterImpl.Contract.CROSSDEX(&_RouterImpl.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_RouterImpl *RouterImplCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RouterImpl.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_RouterImpl *RouterImplSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _RouterImpl.Contract.UPGRADEINTERFACEVERSION(&_RouterImpl.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_RouterImpl *RouterImplCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _RouterImpl.Contract.UPGRADEINTERFACEVERSION(&_RouterImpl.CallOpts)
}

// WXCROSS is a free data retrieval call binding the contract method 0x3499ee7f.
//
// Solidity: function WXCROSS() view returns(address)
func (_RouterImpl *RouterImplCaller) WXCROSS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RouterImpl.contract.Call(opts, &out, "WXCROSS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WXCROSS is a free data retrieval call binding the contract method 0x3499ee7f.
//
// Solidity: function WXCROSS() view returns(address)
func (_RouterImpl *RouterImplSession) WXCROSS() (common.Address, error) {
	return _RouterImpl.Contract.WXCROSS(&_RouterImpl.CallOpts)
}

// WXCROSS is a free data retrieval call binding the contract method 0x3499ee7f.
//
// Solidity: function WXCROSS() view returns(address)
func (_RouterImpl *RouterImplCallerSession) WXCROSS() (common.Address, error) {
	return _RouterImpl.Contract.WXCROSS(&_RouterImpl.CallOpts)
}

// IsPair is a free data retrieval call binding the contract method 0xe5e31b13.
//
// Solidity: function isPair(address pair) view returns(bool)
func (_RouterImpl *RouterImplCaller) IsPair(opts *bind.CallOpts, pair common.Address) (bool, error) {
	var out []interface{}
	err := _RouterImpl.contract.Call(opts, &out, "isPair", pair)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPair is a free data retrieval call binding the contract method 0xe5e31b13.
//
// Solidity: function isPair(address pair) view returns(bool)
func (_RouterImpl *RouterImplSession) IsPair(pair common.Address) (bool, error) {
	return _RouterImpl.Contract.IsPair(&_RouterImpl.CallOpts, pair)
}

// IsPair is a free data retrieval call binding the contract method 0xe5e31b13.
//
// Solidity: function isPair(address pair) view returns(bool)
func (_RouterImpl *RouterImplCallerSession) IsPair(pair common.Address) (bool, error) {
	return _RouterImpl.Contract.IsPair(&_RouterImpl.CallOpts, pair)
}

// MaxMatchCount is a free data retrieval call binding the contract method 0x2c76535e.
//
// Solidity: function maxMatchCount() view returns(uint256)
func (_RouterImpl *RouterImplCaller) MaxMatchCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RouterImpl.contract.Call(opts, &out, "maxMatchCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMatchCount is a free data retrieval call binding the contract method 0x2c76535e.
//
// Solidity: function maxMatchCount() view returns(uint256)
func (_RouterImpl *RouterImplSession) MaxMatchCount() (*big.Int, error) {
	return _RouterImpl.Contract.MaxMatchCount(&_RouterImpl.CallOpts)
}

// MaxMatchCount is a free data retrieval call binding the contract method 0x2c76535e.
//
// Solidity: function maxMatchCount() view returns(uint256)
func (_RouterImpl *RouterImplCallerSession) MaxMatchCount() (*big.Int, error) {
	return _RouterImpl.Contract.MaxMatchCount(&_RouterImpl.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_RouterImpl *RouterImplCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RouterImpl.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_RouterImpl *RouterImplSession) ProxiableUUID() ([32]byte, error) {
	return _RouterImpl.Contract.ProxiableUUID(&_RouterImpl.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_RouterImpl *RouterImplCallerSession) ProxiableUUID() ([32]byte, error) {
	return _RouterImpl.Contract.ProxiableUUID(&_RouterImpl.CallOpts)
}

// Cancel is a paid mutator transaction binding the contract method 0xf9f629c5.
//
// Solidity: function cancel(address pair, uint256[] orderIds) returns()
func (_RouterImpl *RouterImplTransactor) Cancel(opts *bind.TransactOpts, pair common.Address, orderIds []*big.Int) (*types.Transaction, error) {
	return _RouterImpl.contract.Transact(opts, "cancel", pair, orderIds)
}

// Cancel is a paid mutator transaction binding the contract method 0xf9f629c5.
//
// Solidity: function cancel(address pair, uint256[] orderIds) returns()
func (_RouterImpl *RouterImplSession) Cancel(pair common.Address, orderIds []*big.Int) (*types.Transaction, error) {
	return _RouterImpl.Contract.Cancel(&_RouterImpl.TransactOpts, pair, orderIds)
}

// Cancel is a paid mutator transaction binding the contract method 0xf9f629c5.
//
// Solidity: function cancel(address pair, uint256[] orderIds) returns()
func (_RouterImpl *RouterImplTransactorSession) Cancel(pair common.Address, orderIds []*big.Int) (*types.Transaction, error) {
	return _RouterImpl.Contract.Cancel(&_RouterImpl.TransactOpts, pair, orderIds)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 _maxMatchCount) returns()
func (_RouterImpl *RouterImplTransactor) Initialize(opts *bind.TransactOpts, _maxMatchCount *big.Int) (*types.Transaction, error) {
	return _RouterImpl.contract.Transact(opts, "initialize", _maxMatchCount)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 _maxMatchCount) returns()
func (_RouterImpl *RouterImplSession) Initialize(_maxMatchCount *big.Int) (*types.Transaction, error) {
	return _RouterImpl.Contract.Initialize(&_RouterImpl.TransactOpts, _maxMatchCount)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 _maxMatchCount) returns()
func (_RouterImpl *RouterImplTransactorSession) Initialize(_maxMatchCount *big.Int) (*types.Transaction, error) {
	return _RouterImpl.Contract.Initialize(&_RouterImpl.TransactOpts, _maxMatchCount)
}

// LimitBuy is a paid mutator transaction binding the contract method 0x637d2635.
//
// Solidity: function limitBuy(address pair, uint256 price, uint256 amount, uint8 constraints, uint256[2] adjacent, uint256 _maxMatchCount) payable returns(uint256)
func (_RouterImpl *RouterImplTransactor) LimitBuy(opts *bind.TransactOpts, pair common.Address, price *big.Int, amount *big.Int, constraints uint8, adjacent [2]*big.Int, _maxMatchCount *big.Int) (*types.Transaction, error) {
	return _RouterImpl.contract.Transact(opts, "limitBuy", pair, price, amount, constraints, adjacent, _maxMatchCount)
}

// LimitBuy is a paid mutator transaction binding the contract method 0x637d2635.
//
// Solidity: function limitBuy(address pair, uint256 price, uint256 amount, uint8 constraints, uint256[2] adjacent, uint256 _maxMatchCount) payable returns(uint256)
func (_RouterImpl *RouterImplSession) LimitBuy(pair common.Address, price *big.Int, amount *big.Int, constraints uint8, adjacent [2]*big.Int, _maxMatchCount *big.Int) (*types.Transaction, error) {
	return _RouterImpl.Contract.LimitBuy(&_RouterImpl.TransactOpts, pair, price, amount, constraints, adjacent, _maxMatchCount)
}

// LimitBuy is a paid mutator transaction binding the contract method 0x637d2635.
//
// Solidity: function limitBuy(address pair, uint256 price, uint256 amount, uint8 constraints, uint256[2] adjacent, uint256 _maxMatchCount) payable returns(uint256)
func (_RouterImpl *RouterImplTransactorSession) LimitBuy(pair common.Address, price *big.Int, amount *big.Int, constraints uint8, adjacent [2]*big.Int, _maxMatchCount *big.Int) (*types.Transaction, error) {
	return _RouterImpl.Contract.LimitBuy(&_RouterImpl.TransactOpts, pair, price, amount, constraints, adjacent, _maxMatchCount)
}

// LimitSell is a paid mutator transaction binding the contract method 0x473bde60.
//
// Solidity: function limitSell(address pair, uint256 price, uint256 amount, uint8 constraints, uint256[2] adjacent, uint256 _maxMatchCount) payable returns(uint256)
func (_RouterImpl *RouterImplTransactor) LimitSell(opts *bind.TransactOpts, pair common.Address, price *big.Int, amount *big.Int, constraints uint8, adjacent [2]*big.Int, _maxMatchCount *big.Int) (*types.Transaction, error) {
	return _RouterImpl.contract.Transact(opts, "limitSell", pair, price, amount, constraints, adjacent, _maxMatchCount)
}

// LimitSell is a paid mutator transaction binding the contract method 0x473bde60.
//
// Solidity: function limitSell(address pair, uint256 price, uint256 amount, uint8 constraints, uint256[2] adjacent, uint256 _maxMatchCount) payable returns(uint256)
func (_RouterImpl *RouterImplSession) LimitSell(pair common.Address, price *big.Int, amount *big.Int, constraints uint8, adjacent [2]*big.Int, _maxMatchCount *big.Int) (*types.Transaction, error) {
	return _RouterImpl.Contract.LimitSell(&_RouterImpl.TransactOpts, pair, price, amount, constraints, adjacent, _maxMatchCount)
}

// LimitSell is a paid mutator transaction binding the contract method 0x473bde60.
//
// Solidity: function limitSell(address pair, uint256 price, uint256 amount, uint8 constraints, uint256[2] adjacent, uint256 _maxMatchCount) payable returns(uint256)
func (_RouterImpl *RouterImplTransactorSession) LimitSell(pair common.Address, price *big.Int, amount *big.Int, constraints uint8, adjacent [2]*big.Int, _maxMatchCount *big.Int) (*types.Transaction, error) {
	return _RouterImpl.Contract.LimitSell(&_RouterImpl.TransactOpts, pair, price, amount, constraints, adjacent, _maxMatchCount)
}

// MarketBuy is a paid mutator transaction binding the contract method 0xbc58e938.
//
// Solidity: function marketBuy(address pair, uint256 amount, uint256 _maxMatchCount) payable returns()
func (_RouterImpl *RouterImplTransactor) MarketBuy(opts *bind.TransactOpts, pair common.Address, amount *big.Int, _maxMatchCount *big.Int) (*types.Transaction, error) {
	return _RouterImpl.contract.Transact(opts, "marketBuy", pair, amount, _maxMatchCount)
}

// MarketBuy is a paid mutator transaction binding the contract method 0xbc58e938.
//
// Solidity: function marketBuy(address pair, uint256 amount, uint256 _maxMatchCount) payable returns()
func (_RouterImpl *RouterImplSession) MarketBuy(pair common.Address, amount *big.Int, _maxMatchCount *big.Int) (*types.Transaction, error) {
	return _RouterImpl.Contract.MarketBuy(&_RouterImpl.TransactOpts, pair, amount, _maxMatchCount)
}

// MarketBuy is a paid mutator transaction binding the contract method 0xbc58e938.
//
// Solidity: function marketBuy(address pair, uint256 amount, uint256 _maxMatchCount) payable returns()
func (_RouterImpl *RouterImplTransactorSession) MarketBuy(pair common.Address, amount *big.Int, _maxMatchCount *big.Int) (*types.Transaction, error) {
	return _RouterImpl.Contract.MarketBuy(&_RouterImpl.TransactOpts, pair, amount, _maxMatchCount)
}

// MarketSell is a paid mutator transaction binding the contract method 0x26580d11.
//
// Solidity: function marketSell(address pair, uint256 amount, uint256 _maxMatchCount) payable returns()
func (_RouterImpl *RouterImplTransactor) MarketSell(opts *bind.TransactOpts, pair common.Address, amount *big.Int, _maxMatchCount *big.Int) (*types.Transaction, error) {
	return _RouterImpl.contract.Transact(opts, "marketSell", pair, amount, _maxMatchCount)
}

// MarketSell is a paid mutator transaction binding the contract method 0x26580d11.
//
// Solidity: function marketSell(address pair, uint256 amount, uint256 _maxMatchCount) payable returns()
func (_RouterImpl *RouterImplSession) MarketSell(pair common.Address, amount *big.Int, _maxMatchCount *big.Int) (*types.Transaction, error) {
	return _RouterImpl.Contract.MarketSell(&_RouterImpl.TransactOpts, pair, amount, _maxMatchCount)
}

// MarketSell is a paid mutator transaction binding the contract method 0x26580d11.
//
// Solidity: function marketSell(address pair, uint256 amount, uint256 _maxMatchCount) payable returns()
func (_RouterImpl *RouterImplTransactorSession) MarketSell(pair common.Address, amount *big.Int, _maxMatchCount *big.Int) (*types.Transaction, error) {
	return _RouterImpl.Contract.MarketSell(&_RouterImpl.TransactOpts, pair, amount, _maxMatchCount)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_RouterImpl *RouterImplTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _RouterImpl.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_RouterImpl *RouterImplSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _RouterImpl.Contract.UpgradeToAndCall(&_RouterImpl.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_RouterImpl *RouterImplTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _RouterImpl.Contract.UpgradeToAndCall(&_RouterImpl.TransactOpts, newImplementation, data)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RouterImpl *RouterImplTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RouterImpl.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RouterImpl *RouterImplSession) Receive() (*types.Transaction, error) {
	return _RouterImpl.Contract.Receive(&_RouterImpl.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RouterImpl *RouterImplTransactorSession) Receive() (*types.Transaction, error) {
	return _RouterImpl.Contract.Receive(&_RouterImpl.TransactOpts)
}

// RouterImplInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the RouterImpl contract.
type RouterImplInitializedIterator struct {
	Event *RouterImplInitialized // Event containing the contract specifics and raw log

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
func (it *RouterImplInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterImplInitialized)
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
		it.Event = new(RouterImplInitialized)
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
func (it *RouterImplInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterImplInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterImplInitialized represents a Initialized event raised by the RouterImpl contract.
type RouterImplInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_RouterImpl *RouterImplFilterer) FilterInitialized(opts *bind.FilterOpts) (*RouterImplInitializedIterator, error) {

	logs, sub, err := _RouterImpl.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RouterImplInitializedIterator{contract: _RouterImpl.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_RouterImpl *RouterImplFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RouterImplInitialized) (event.Subscription, error) {

	logs, sub, err := _RouterImpl.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterImplInitialized)
				if err := _RouterImpl.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_RouterImpl *RouterImplFilterer) ParseInitialized(log types.Log) (*RouterImplInitialized, error) {
	event := new(RouterImplInitialized)
	if err := _RouterImpl.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RouterImplPairAddedIterator is returned from FilterPairAdded and is used to iterate over the raw logs and unpacked data for PairAdded events raised by the RouterImpl contract.
type RouterImplPairAddedIterator struct {
	Event *RouterImplPairAdded // Event containing the contract specifics and raw log

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
func (it *RouterImplPairAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterImplPairAdded)
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
		it.Event = new(RouterImplPairAdded)
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
func (it *RouterImplPairAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterImplPairAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterImplPairAdded represents a PairAdded event raised by the RouterImpl contract.
type RouterImplPairAdded struct {
	Pair  common.Address
	Base  common.Address
	Quote common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPairAdded is a free log retrieval operation binding the contract event 0x7f46075c67ca5bbd3aaf82d8e324282141f27b7cba3376e5498a6b70c9931c2a.
//
// Solidity: event PairAdded(address indexed pair, address indexed base, address indexed quote)
func (_RouterImpl *RouterImplFilterer) FilterPairAdded(opts *bind.FilterOpts, pair []common.Address, base []common.Address, quote []common.Address) (*RouterImplPairAddedIterator, error) {

	var pairRule []interface{}
	for _, pairItem := range pair {
		pairRule = append(pairRule, pairItem)
	}
	var baseRule []interface{}
	for _, baseItem := range base {
		baseRule = append(baseRule, baseItem)
	}
	var quoteRule []interface{}
	for _, quoteItem := range quote {
		quoteRule = append(quoteRule, quoteItem)
	}

	logs, sub, err := _RouterImpl.contract.FilterLogs(opts, "PairAdded", pairRule, baseRule, quoteRule)
	if err != nil {
		return nil, err
	}
	return &RouterImplPairAddedIterator{contract: _RouterImpl.contract, event: "PairAdded", logs: logs, sub: sub}, nil
}

// WatchPairAdded is a free log subscription operation binding the contract event 0x7f46075c67ca5bbd3aaf82d8e324282141f27b7cba3376e5498a6b70c9931c2a.
//
// Solidity: event PairAdded(address indexed pair, address indexed base, address indexed quote)
func (_RouterImpl *RouterImplFilterer) WatchPairAdded(opts *bind.WatchOpts, sink chan<- *RouterImplPairAdded, pair []common.Address, base []common.Address, quote []common.Address) (event.Subscription, error) {

	var pairRule []interface{}
	for _, pairItem := range pair {
		pairRule = append(pairRule, pairItem)
	}
	var baseRule []interface{}
	for _, baseItem := range base {
		baseRule = append(baseRule, baseItem)
	}
	var quoteRule []interface{}
	for _, quoteItem := range quote {
		quoteRule = append(quoteRule, quoteItem)
	}

	logs, sub, err := _RouterImpl.contract.WatchLogs(opts, "PairAdded", pairRule, baseRule, quoteRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterImplPairAdded)
				if err := _RouterImpl.contract.UnpackLog(event, "PairAdded", log); err != nil {
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

// ParsePairAdded is a log parse operation binding the contract event 0x7f46075c67ca5bbd3aaf82d8e324282141f27b7cba3376e5498a6b70c9931c2a.
//
// Solidity: event PairAdded(address indexed pair, address indexed base, address indexed quote)
func (_RouterImpl *RouterImplFilterer) ParsePairAdded(log types.Log) (*RouterImplPairAdded, error) {
	event := new(RouterImplPairAdded)
	if err := _RouterImpl.contract.UnpackLog(event, "PairAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RouterImplPairRemovedIterator is returned from FilterPairRemoved and is used to iterate over the raw logs and unpacked data for PairRemoved events raised by the RouterImpl contract.
type RouterImplPairRemovedIterator struct {
	Event *RouterImplPairRemoved // Event containing the contract specifics and raw log

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
func (it *RouterImplPairRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterImplPairRemoved)
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
		it.Event = new(RouterImplPairRemoved)
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
func (it *RouterImplPairRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterImplPairRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterImplPairRemoved represents a PairRemoved event raised by the RouterImpl contract.
type RouterImplPairRemoved struct {
	Pair common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPairRemoved is a free log retrieval operation binding the contract event 0x9493af33d363a70a1cd9cc0ceb347e3c0e16b5d9371695618e1ac5fb1c255b7c.
//
// Solidity: event PairRemoved(address indexed pair)
func (_RouterImpl *RouterImplFilterer) FilterPairRemoved(opts *bind.FilterOpts, pair []common.Address) (*RouterImplPairRemovedIterator, error) {

	var pairRule []interface{}
	for _, pairItem := range pair {
		pairRule = append(pairRule, pairItem)
	}

	logs, sub, err := _RouterImpl.contract.FilterLogs(opts, "PairRemoved", pairRule)
	if err != nil {
		return nil, err
	}
	return &RouterImplPairRemovedIterator{contract: _RouterImpl.contract, event: "PairRemoved", logs: logs, sub: sub}, nil
}

// WatchPairRemoved is a free log subscription operation binding the contract event 0x9493af33d363a70a1cd9cc0ceb347e3c0e16b5d9371695618e1ac5fb1c255b7c.
//
// Solidity: event PairRemoved(address indexed pair)
func (_RouterImpl *RouterImplFilterer) WatchPairRemoved(opts *bind.WatchOpts, sink chan<- *RouterImplPairRemoved, pair []common.Address) (event.Subscription, error) {

	var pairRule []interface{}
	for _, pairItem := range pair {
		pairRule = append(pairRule, pairItem)
	}

	logs, sub, err := _RouterImpl.contract.WatchLogs(opts, "PairRemoved", pairRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterImplPairRemoved)
				if err := _RouterImpl.contract.UnpackLog(event, "PairRemoved", log); err != nil {
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

// ParsePairRemoved is a log parse operation binding the contract event 0x9493af33d363a70a1cd9cc0ceb347e3c0e16b5d9371695618e1ac5fb1c255b7c.
//
// Solidity: event PairRemoved(address indexed pair)
func (_RouterImpl *RouterImplFilterer) ParsePairRemoved(log types.Log) (*RouterImplPairRemoved, error) {
	event := new(RouterImplPairRemoved)
	if err := _RouterImpl.contract.UnpackLog(event, "PairRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RouterImplUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the RouterImpl contract.
type RouterImplUpgradedIterator struct {
	Event *RouterImplUpgraded // Event containing the contract specifics and raw log

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
func (it *RouterImplUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterImplUpgraded)
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
		it.Event = new(RouterImplUpgraded)
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
func (it *RouterImplUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterImplUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterImplUpgraded represents a Upgraded event raised by the RouterImpl contract.
type RouterImplUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_RouterImpl *RouterImplFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*RouterImplUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _RouterImpl.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &RouterImplUpgradedIterator{contract: _RouterImpl.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_RouterImpl *RouterImplFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *RouterImplUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _RouterImpl.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterImplUpgraded)
				if err := _RouterImpl.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_RouterImpl *RouterImplFilterer) ParseUpgraded(log types.Log) (*RouterImplUpgraded, error) {
	event := new(RouterImplUpgraded)
	if err := _RouterImpl.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
