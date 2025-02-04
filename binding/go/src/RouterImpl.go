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
	ABI: "[{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"CROSS_DEX\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WCross\",\"outputs\":[{\"internalType\":\"contractIWCROSS\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"orderIds\",\"type\":\"uint256[]\"}],\"name\":\"cancel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxMatchCount\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"name\":\"isPair\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"searchPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxMatchCount\",\"type\":\"uint256\"}],\"name\":\"limitBuy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"searchPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxMatchCount\",\"type\":\"uint256\"}],\"name\":\"limitSell\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxMatchCount\",\"type\":\"uint256\"}],\"name\":\"marketBuy\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxMatchCount\",\"type\":\"uint256\"}],\"name\":\"marketSell\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxMatchCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"quote\",\"type\":\"address\"}],\"name\":\"PairAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"name\":\"PairRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"RouterInitializeData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"RouterInvalidPairAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterInvalidValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"b820d8aa": "CROSS_DEX()",
		"ad3cb1cc": "UPGRADE_INTERFACE_VERSION()",
		"7f0d2bdc": "WCross()",
		"f9f629c5": "cancel(address,uint256[])",
		"fe4b84df": "initialize(uint256)",
		"e5e31b13": "isPair(address)",
		"5d2b4c09": "limitBuy(address,uint256,uint256,uint256,uint256)",
		"f3ea4d4d": "limitSell(address,uint256,uint256,uint256,uint256)",
		"bc58e938": "marketBuy(address,uint256,uint256)",
		"26580d11": "marketSell(address,uint256,uint256)",
		"2c76535e": "maxMatchCount()",
		"52d1902d": "proxiableUUID()",
		"4f1ef286": "upgradeToAndCall(address,bytes)",
	},
	Bin: "0x60a0604052306080523480156012575f5ffd5b506080516124156100395f395f8181610ddf01528181610e080152610fea01526124155ff3fe6080604052600436106100c2575f3560e01c8063ad3cb1cc1161007c578063e5e31b1311610057578063e5e31b1314610206578063f3ea4d4d14610235578063f9f629c514610248578063fe4b84df14610267575f5ffd5b8063ad3cb1cc14610198578063b820d8aa146101d5578063bc58e938146101f3575f5ffd5b806326580d11146100ec5780632c76535e146100ff5780634f1ef2861461012757806352d1902d1461013a5780635d2b4c091461014e5780637f0d2bdc14610161575f5ffd5b366100e85747156100e65760405163f85c3ef560e01b815260040160405180910390fd5b005b5f5ffd5b6100e66100fa3660046113ab565b610286565b34801561010a575f5ffd5b5061011460025481565b6040519081526020015b60405180910390f35b6100e6610135366004611422565b6104a1565b348015610145575f5ffd5b506101146104c0565b61011461015c3660046114c9565b6104db565b34801561016c575f5ffd5b50600154610180906001600160a01b031681565b6040516001600160a01b03909116815260200161011e565b3480156101a3575f5ffd5b506101c8604051806040016040528060058152602001640352e302e360dc1b81525081565b60405161011e9190611509565b3480156101e0575f5ffd5b505f54610180906001600160a01b031681565b6100e66102013660046113ab565b61070f565b348015610211575f5ffd5b5061022561022036600461153e565b610856565b604051901515815260200161011e565b6101146102433660046114c9565b6108d3565b348015610253575f5ffd5b506100e6610262366004611559565b610acf565b348015610272575f5ffd5b506100e66102813660046115dc565b610b64565b61028e610d07565b8261029881610856565b6102c557604051633b1c760160e21b81526001600160a01b03821660048201526024015b60405180910390fd5b5f3390505f856001600160a01b031663a1441f4f6040518163ffffffff1660e01b8152600401606060405180830381865afa158015610306573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061032a91906115f3565b6020810151600154919250906001600160a01b03908116908216036103ab5760015460405163755edd1760e01b81526001600160a01b0389811660048301529091169063755edd179088906024015f604051808303818588803b15801561038f575f5ffd5b505af11580156103a1573d5f5f3e3d5ffd5b50505050506103c0565b6103c06001600160a01b038216848989610d3e565b6040805160a081019091525f9080825b8152602001856001600160a01b031681526020015f63ffffffff1681526020015f81526020015f8152509050876001600160a01b031663948faf9082896104168a610d9e565b6040518463ffffffff1660e01b81526004016104349392919061165e565b5f604051808303815f87803b15801561044b575f5ffd5b505af115801561045d573d5f5f3e3d5ffd5b5050505050505050475f146104855760405163f85c3ef560e01b815260040160405180910390fd5b5061049c60015f5160206123c05f395f51905f5255565b505050565b6104a9610dd4565b6104b282610e7a565b6104bc8282610f23565b5050565b5f6104c9610fdf565b505f5160206123a05f395f51905f5290565b5f6104e4610d07565b856104ee81610856565b61051657604051633b1c760160e21b81526001600160a01b03821660048201526024016102bc565b5f3390505f886001600160a01b031663a1441f4f6040518163ffffffff1660e01b8152600401606060405180830381865afa158015610557573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061057b91906115f3565b90505f815f015190505f6105948a8a8560400151611028565b6001549091506001600160a01b039081169083160361060f5760015460405163755edd1760e01b81526001600160a01b038d811660048301529091169063755edd179083906024015f604051808303818588803b1580156105f3575f5ffd5b505af1158015610605573d5f5f3e3d5ffd5b5050505050610624565b6106246001600160a01b038316858d84610d3e565b6040805160a081018252600181526001600160a01b0380871660208301525f92820192909252606081018c9052608081018b9052908c1663f938c5b1828b61066b8c610d9e565b6040518463ffffffff1660e01b81526004016106899392919061165e565b6020604051808303815f875af11580156106a5573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106c991906116ce565b9650505050505047156106ef5760405163f85c3ef560e01b815260040160405180910390fd5b5061070660015f5160206123c05f395f51905f5255565b95945050505050565b610717610d07565b8261072181610856565b61074957604051633b1c760160e21b81526001600160a01b03821660048201526024016102bc565b5f3390505f856001600160a01b031663a1441f4f6040518163ffffffff1660e01b8152600401606060405180830381865afa15801561078a573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107ae91906115f3565b8051600154919250906001600160a01b039081169082160361082c5760015460405163755edd1760e01b81526001600160a01b0389811660048301529091169063755edd179088906024015f604051808303818588803b158015610810575f5ffd5b505af1158015610822573d5f5f3e3d5ffd5b5050505050610841565b6108416001600160a01b038216848989610d3e565b6040805160a081019091525f908060016103d0565b5f8054604051630827057360e01b81526001600160a01b03848116600483015283921690630827057390602401602060405180830381865afa15801561089e573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108c291906116e5565b6001600160a01b0316141592915050565b5f6108dc610d07565b856108e681610856565b61090e57604051633b1c760160e21b81526001600160a01b03821660048201526024016102bc565b5f3390505f886001600160a01b031663a1441f4f6040518163ffffffff1660e01b8152600401606060405180830381865afa15801561094f573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061097391906115f3565b6020810151600154919250906001600160a01b03908116908216036109f45760015460405163755edd1760e01b81526001600160a01b038c811660048301529091169063755edd17908a906024015f604051808303818588803b1580156109d8575f5ffd5b505af11580156109ea573d5f5f3e3d5ffd5b5050505050610a09565b610a096001600160a01b038216848c8b610d3e565b6040805160a0810182525f8082526001600160a01b03808716602084015292820152606081018b9052608081018a9052908b1663f938c5b1828a610a4c8b610d9e565b6040518463ffffffff1660e01b8152600401610a6a9392919061165e565b6020604051808303815f875af1158015610a86573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610aaa91906116ce565b95505050505047156106ef5760405163f85c3ef560e01b815260040160405180910390fd5b82610ad981610856565b610b0157604051633b1c760160e21b81526001600160a01b03821660048201526024016102bc565b60405163f9f629c560e01b81526001600160a01b0385169063f9f629c590610b3190339087908790600401611700565b5f604051808303815f87803b158015610b48575f5ffd5b505af1158015610b5a573d5f5f3e3d5ffd5b5050505050505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff16159067ffffffffffffffff165f81158015610ba95750825b90505f8267ffffffffffffffff166001148015610bc55750303b155b905081158015610bd3575080155b15610bf15760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff191660011785558315610c1b57845460ff60401b1916600160401b1785555b855f03610c4d5760405163b927e50b60e01b81526c1b585e13585d18da10dbdd5b9d609a1b60048201526024016102bc565b5f80546001600160a01b03191633179055604051610c6a9061138a565b604051809103905ff080158015610c83573d5f5f3e3d5ffd5b50600180546001600160a01b0319166001600160a01b03929092169190911790556002869055610cb16110df565b610cb96110e7565b8315610cff57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050565b5f5160206123c05f395f51905f52805460011901610d3857604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b604080516001600160a01b0385811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180516001600160e01b03166323b872dd60e01b179052610d989085906110f7565b50505050565b5f811580610dad575060025482115b610db75781610dbb565b6002545b92915050565b60015f5160206123c05f395f51905f5255565b306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480610e5a57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316610e4e5f5160206123a05f395f51905f52546001600160a01b031690565b6001600160a01b031614155b15610e785760405163703e46dd60e11b815260040160405180910390fd5b565b5f5f9054906101000a90046001600160a01b03166001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610ec9573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610eed91906116e5565b6001600160a01b0316336001600160a01b031614610f205760405163118cdaa760e01b81523360048201526024016102bc565b50565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015610f7d575060408051601f3d908101601f19168201909252610f7a918101906116ce565b60015b610fa557604051634c9c8ce360e01b81526001600160a01b03831660048201526024016102bc565b5f5160206123a05f395f51905f528114610fd557604051632a87526960e21b8152600481018290526024016102bc565b61049c8383611163565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610e785760405163703e46dd60e11b815260040160405180910390fd5b5f838302815f1985870982811083820303915050805f0361105c5783828161105257611052611748565b04925050506110d8565b8084116110735761107360038515026011186111b8565b5f848688095f868103871696879004966002600389028118808a02820302808a02820302808a02820302808a02820302808a02820302808a02909103029181900381900460010186841190950394909402919094039290920491909117919091029150505b9392505050565b610e786111c9565b6110ef6111c9565b610e78611212565b5f5f60205f8451602086015f885af180611116576040513d5f823e3d81fd5b50505f513d9150811561112d57806001141561113a565b6001600160a01b0384163b155b15610d9857604051635274afe760e01b81526001600160a01b03851660048201526024016102bc565b61116c8261121a565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a28051156111b05761049c828261127d565b6104bc6112e6565b634e487b715f52806020526024601cfd5b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff16610e7857604051631afcd79f60e31b815260040160405180910390fd5b610dc16111c9565b806001600160a01b03163b5f0361124f57604051634c9c8ce360e01b81526001600160a01b03821660048201526024016102bc565b5f5160206123a05f395f51905f5280546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b031684604051611299919061175c565b5f60405180830381855af49150503d805f81146112d1576040519150601f19603f3d011682016040523d82523d5f602084013e6112d6565b606091505b5091509150610706858383611305565b3415610e785760405163b398979f60e01b815260040160405180910390fd5b60608261131a5761131582611361565b6110d8565b815115801561133157506001600160a01b0384163b155b1561135a57604051639996b31560e01b81526001600160a01b03851660048201526024016102bc565b50806110d8565b8051156113715780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b610c2d8061177383390190565b6001600160a01b0381168114610f20575f5ffd5b5f5f5f606084860312156113bd575f5ffd5b83356113c881611397565b95602085013595506040909401359392505050565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff8111828210171561141a5761141a6113dd565b604052919050565b5f5f60408385031215611433575f5ffd5b823561143e81611397565b9150602083013567ffffffffffffffff811115611459575f5ffd5b8301601f81018513611469575f5ffd5b803567ffffffffffffffff811115611483576114836113dd565b611496601f8201601f19166020016113f1565b8181528660208385010111156114aa575f5ffd5b816020840160208301375f602083830101528093505050509250929050565b5f5f5f5f5f60a086880312156114dd575f5ffd5b85356114e881611397565b97602087013597506040870135966060810135965060800135945092505050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b5f6020828403121561154e575f5ffd5b81356110d881611397565b5f5f5f6040848603121561156b575f5ffd5b833561157681611397565b9250602084013567ffffffffffffffff811115611591575f5ffd5b8401601f810186136115a1575f5ffd5b803567ffffffffffffffff8111156115b7575f5ffd5b8660208260051b84010111156115cb575f5ffd5b939660209190910195509293505050565b5f602082840312156115ec575f5ffd5b5035919050565b5f6060828403128015611604575f5ffd5b506040516060810167ffffffffffffffff81118282101715611628576116286113dd565b604052825161163681611397565b8152602083015161164681611397565b60208201526040928301519281019290925250919050565b835160e08201906002811061168157634e487b7160e01b5f52602160045260245ffd5b82526020858101516001600160a01b03169083015260408086015163ffffffff1690830152606080860151908301526080948501519482019490945260a081019290925260c09091015290565b5f602082840312156116de575f5ffd5b5051919050565b5f602082840312156116f5575f5ffd5b81516110d881611397565b6001600160a01b038416815260406020820181905281018290525f6001600160fb1b0383111561172e575f5ffd5b8260051b8085606085013791909101606001949350505050565b634e487b7160e01b5f52601260045260245ffd5b5f82518060208501845e5f92019182525091905056fe60a060405234801561000f575f5ffd5b506040518060400160405280601681526020017f43726f737344455820577261707065642043726f73730000000000000000000081525060405180604001604052806007815260200166636443524f535360c81b81525081600390816100759190610136565b5060046100828282610136565b5061008d9150503390565b6001600160a01b03166080526101f0565b634e487b7160e01b5f52604160045260245ffd5b600181811c908216806100c657607f821691505b6020821081036100e457634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561013157805f5260205f20601f840160051c8101602085101561010f5750805b601f840160051c820191505b8181101561012e575f815560010161011b565b50505b505050565b81516001600160401b0381111561014f5761014f61009e565b6101638161015d84546100b2565b846100ea565b6020601f821160018114610195575f831561017e5750848201515b5f19600385901b1c1916600184901b17845561012e565b5f84815260208120601f198516915b828110156101c457878501518255602094850194600190920191016101a4565b50848210156101e157868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b608051610a176102165f395f8181610174015281816103ba01526104120152610a175ff3fe60806040526004361061009d575f3560e01c806332fe7b261161006257806332fe7b261461016357806370a08231146101ae578063755edd17146101e257806395d89b41146101f5578063a9059cbb14610209578063dd62ed3e14610228575f5ffd5b806306fdde03146100b2578063095ea7b3146100dc57806318160ddd1461010b57806323b872dd14610129578063313ce56714610148575f5ffd5b366100ae576100ac333461026c565b005b5f5ffd5b3480156100bd575f5ffd5b506100c66102a9565b6040516100d39190610868565b60405180910390f35b3480156100e7575f5ffd5b506100fb6100f63660046108b8565b610339565b60405190151581526020016100d3565b348015610116575f5ffd5b506002545b6040519081526020016100d3565b348015610134575f5ffd5b506100fb6101433660046108e0565b610352565b348015610153575f5ffd5b50604051601281526020016100d3565b34801561016e575f5ffd5b506101967f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016100d3565b3480156101b9575f5ffd5b5061011b6101c836600461091a565b6001600160a01b03165f9081526020819052604090205490565b6100ac6101f036600461091a565b610375565b348015610200575f5ffd5b506100c6610382565b348015610214575f5ffd5b506100fb6102233660046108b8565b610391565b348015610233575f5ffd5b5061011b61024236600461093a565b6001600160a01b039182165f90815260016020908152604080832093909416825291909152205490565b6001600160a01b03821661029a5760405163ec442f0560e01b81525f60048201526024015b60405180910390fd5b6102a55f838361039e565b5050565b6060600380546102b89061096b565b80601f01602080910402602001604051908101604052809291908181526020018280546102e49061096b565b801561032f5780601f106103065761010080835404028352916020019161032f565b820191905f5260205f20905b81548152906001019060200180831161031257829003601f168201915b5050505050905090565b5f336103468185856104a1565b60019150505b92915050565b5f3361035f8582856104ae565b61036a85858561052a565b506001949350505050565b61037f813461026c565b50565b6060600480546102b89061096b565b5f3361034681858561052a565b6103a9838383610587565b6001600160a01b0382161561049c577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316826001600160a01b0316148061047b575060405163e5e31b1360e01b81526001600160a01b0383811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063e5e31b1390602401602060405180830381865afa158015610457573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061047b91906109a3565b61049c5761048982826106ad565b61049c6001600160a01b038316826106e1565b505050565b61049c838383600161076d565b6001600160a01b038381165f908152600160209081526040808320938616835292905220545f19811015610524578181101561051657604051637dc7a0d960e11b81526001600160a01b03841660048201526024810182905260448101839052606401610291565b61052484848484035f61076d565b50505050565b6001600160a01b03831661055357604051634b637e8f60e11b81525f6004820152602401610291565b6001600160a01b03821661057c5760405163ec442f0560e01b81525f6004820152602401610291565b61049c83838361039e565b6001600160a01b0383166105b1578060025f8282546105a691906109c2565b909155506106219050565b6001600160a01b0383165f90815260208190526040902054818110156106035760405163391434e360e21b81526001600160a01b03851660048201526024810182905260448101839052606401610291565b6001600160a01b0384165f9081526020819052604090209082900390555b6001600160a01b03821661063d5760028054829003905561065b565b6001600160a01b0382165f9081526020819052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516106a091815260200190565b60405180910390a3505050565b6001600160a01b0382166106d657604051634b637e8f60e11b81525f6004820152602401610291565b6102a5825f8361039e565b8047101561070b5760405163cf47918160e01b815247600482015260248101829052604401610291565b5f5f836001600160a01b0316836040515f6040518083038185875af1925050503d805f8114610755576040519150601f19603f3d011682016040523d82523d5f602084013e61075a565b606091505b509150915081610524576105248161083f565b6001600160a01b0384166107965760405163e602df0560e01b81525f6004820152602401610291565b6001600160a01b0383166107bf57604051634a1406b160e11b81525f6004820152602401610291565b6001600160a01b038085165f908152600160209081526040808320938716835292905220829055801561052457826001600160a01b0316846001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161083191815260200190565b60405180910390a350505050565b80511561084f5780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b80356001600160a01b03811681146108b3575f5ffd5b919050565b5f5f604083850312156108c9575f5ffd5b6108d28361089d565b946020939093013593505050565b5f5f5f606084860312156108f2575f5ffd5b6108fb8461089d565b92506109096020850161089d565b929592945050506040919091013590565b5f6020828403121561092a575f5ffd5b6109338261089d565b9392505050565b5f5f6040838503121561094b575f5ffd5b6109548361089d565b91506109626020840161089d565b90509250929050565b600181811c9082168061097f57607f821691505b60208210810361099d57634e487b7160e01b5f52602260045260245ffd5b50919050565b5f602082840312156109b3575f5ffd5b81518015158114610933575f5ffd5b8082018082111561034c57634e487b7160e01b5f52601160045260245ffdfea264697066735822122073bcef6bf2768756a5c355f6a7f748703b5ee064a0e9414e5762222cd42e0e3f64736f6c634300081c0033360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a264697066735822122004e7c4f5fc32938acddfcbf91f362de239e445b7c06dcf870c4cd4da195ab39e64736f6c634300081c0033",
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

// WCross is a free data retrieval call binding the contract method 0x7f0d2bdc.
//
// Solidity: function WCross() view returns(address)
func (_RouterImpl *RouterImplCaller) WCross(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RouterImpl.contract.Call(opts, &out, "WCross")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WCross is a free data retrieval call binding the contract method 0x7f0d2bdc.
//
// Solidity: function WCross() view returns(address)
func (_RouterImpl *RouterImplSession) WCross() (common.Address, error) {
	return _RouterImpl.Contract.WCross(&_RouterImpl.CallOpts)
}

// WCross is a free data retrieval call binding the contract method 0x7f0d2bdc.
//
// Solidity: function WCross() view returns(address)
func (_RouterImpl *RouterImplCallerSession) WCross() (common.Address, error) {
	return _RouterImpl.Contract.WCross(&_RouterImpl.CallOpts)
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

// LimitBuy is a paid mutator transaction binding the contract method 0x5d2b4c09.
//
// Solidity: function limitBuy(address pair, uint256 price, uint256 amount, uint256 searchPrice, uint256 _maxMatchCount) payable returns(uint256)
func (_RouterImpl *RouterImplTransactor) LimitBuy(opts *bind.TransactOpts, pair common.Address, price *big.Int, amount *big.Int, searchPrice *big.Int, _maxMatchCount *big.Int) (*types.Transaction, error) {
	return _RouterImpl.contract.Transact(opts, "limitBuy", pair, price, amount, searchPrice, _maxMatchCount)
}

// LimitBuy is a paid mutator transaction binding the contract method 0x5d2b4c09.
//
// Solidity: function limitBuy(address pair, uint256 price, uint256 amount, uint256 searchPrice, uint256 _maxMatchCount) payable returns(uint256)
func (_RouterImpl *RouterImplSession) LimitBuy(pair common.Address, price *big.Int, amount *big.Int, searchPrice *big.Int, _maxMatchCount *big.Int) (*types.Transaction, error) {
	return _RouterImpl.Contract.LimitBuy(&_RouterImpl.TransactOpts, pair, price, amount, searchPrice, _maxMatchCount)
}

// LimitBuy is a paid mutator transaction binding the contract method 0x5d2b4c09.
//
// Solidity: function limitBuy(address pair, uint256 price, uint256 amount, uint256 searchPrice, uint256 _maxMatchCount) payable returns(uint256)
func (_RouterImpl *RouterImplTransactorSession) LimitBuy(pair common.Address, price *big.Int, amount *big.Int, searchPrice *big.Int, _maxMatchCount *big.Int) (*types.Transaction, error) {
	return _RouterImpl.Contract.LimitBuy(&_RouterImpl.TransactOpts, pair, price, amount, searchPrice, _maxMatchCount)
}

// LimitSell is a paid mutator transaction binding the contract method 0xf3ea4d4d.
//
// Solidity: function limitSell(address pair, uint256 price, uint256 amount, uint256 searchPrice, uint256 _maxMatchCount) payable returns(uint256)
func (_RouterImpl *RouterImplTransactor) LimitSell(opts *bind.TransactOpts, pair common.Address, price *big.Int, amount *big.Int, searchPrice *big.Int, _maxMatchCount *big.Int) (*types.Transaction, error) {
	return _RouterImpl.contract.Transact(opts, "limitSell", pair, price, amount, searchPrice, _maxMatchCount)
}

// LimitSell is a paid mutator transaction binding the contract method 0xf3ea4d4d.
//
// Solidity: function limitSell(address pair, uint256 price, uint256 amount, uint256 searchPrice, uint256 _maxMatchCount) payable returns(uint256)
func (_RouterImpl *RouterImplSession) LimitSell(pair common.Address, price *big.Int, amount *big.Int, searchPrice *big.Int, _maxMatchCount *big.Int) (*types.Transaction, error) {
	return _RouterImpl.Contract.LimitSell(&_RouterImpl.TransactOpts, pair, price, amount, searchPrice, _maxMatchCount)
}

// LimitSell is a paid mutator transaction binding the contract method 0xf3ea4d4d.
//
// Solidity: function limitSell(address pair, uint256 price, uint256 amount, uint256 searchPrice, uint256 _maxMatchCount) payable returns(uint256)
func (_RouterImpl *RouterImplTransactorSession) LimitSell(pair common.Address, price *big.Int, amount *big.Int, searchPrice *big.Int, _maxMatchCount *big.Int) (*types.Transaction, error) {
	return _RouterImpl.Contract.LimitSell(&_RouterImpl.TransactOpts, pair, price, amount, searchPrice, _maxMatchCount)
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
