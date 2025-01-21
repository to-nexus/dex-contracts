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
	ABI: "[{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"name\":\"addPair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"orderIds\",\"type\":\"uint256[]\"}],\"name\":\"cancel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"weth\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxMatchCount\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"name\":\"isPair\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"searchPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxMatchCount\",\"type\":\"uint256\"}],\"name\":\"limitBuy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"searchPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxMatchCount\",\"type\":\"uint256\"}],\"name\":\"limitSell\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxMatchCount\",\"type\":\"uint256\"}],\"name\":\"marketBuy\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxMatchCount\",\"type\":\"uint256\"}],\"name\":\"marketSell\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxMatchCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"name\":\"pairInfo\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"BASE\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"QUOTE\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"DENOMINATOR\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"name\":\"removePair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"quote\",\"type\":\"address\"}],\"name\":\"PairAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"name\":\"PairRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"RouterAlreadyAddedPair\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"RouterInitializeData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"RouterInvalidPairAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterInvalidValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"ad3cb1cc": "UPGRADE_INTERFACE_VERSION()",
		"ad5c4648": "WETH()",
		"c2b7bbb6": "addPair(address)",
		"f9f629c5": "cancel(address,uint256[])",
		"cd6dc687": "initialize(address,uint256)",
		"e5e31b13": "isPair(address)",
		"5d2b4c09": "limitBuy(address,uint256,uint256,uint256,uint256)",
		"f3ea4d4d": "limitSell(address,uint256,uint256,uint256,uint256)",
		"bc58e938": "marketBuy(address,uint256,uint256)",
		"26580d11": "marketSell(address,uint256,uint256)",
		"2c76535e": "maxMatchCount()",
		"8da5cb5b": "owner()",
		"dc0a4229": "pairInfo(address)",
		"52d1902d": "proxiableUUID()",
		"af6c9c1d": "removePair(address)",
		"715018a6": "renounceOwnership()",
		"f2fde38b": "transferOwnership(address)",
		"4f1ef286": "upgradeToAndCall(address,bytes)",
	},
	Bin: "0x60a0604052306080523480156012575f5ffd5b50608051611d3f6100395f395f81816110ff0152818161112801526112670152611d3f5ff3fe608060405260043610610108575f3560e01c8063af6c9c1d11610092578063dc0a422911610062578063dc0a4229146102f5578063e5e31b1314610362578063f2fde38b14610391578063f3ea4d4d146103b0578063f9f629c5146103c3575f5ffd5b8063af6c9c1d14610285578063bc58e938146102a4578063c2b7bbb6146102b7578063cd6dc687146102d6575f5ffd5b80635d2b4c09116100d85780635d2b4c09146101b3578063715018a6146101c65780638da5cb5b146101da578063ad3cb1cc1461022a578063ad5c464814610267575f5ffd5b806326580d11146101515780632c76535e146101645780634f1ef2861461018c57806352d1902d1461019f575f5ffd5b3661014d57341561012c5760405163f85c3ef560e01b815260040160405180910390fd5b471561014b5760405163f85c3ef560e01b815260040160405180910390fd5b005b5f5ffd5b61014b61015f366004611925565b6103e2565b34801561016f575f5ffd5b5061017960015481565b6040519081526020015b60405180910390f35b61014b61019a36600461196b565b61055a565b3480156101aa575f5ffd5b50610179610579565b6101796101c1366004611a31565b610594565b3480156101d1575f5ffd5b5061014b610726565b3480156101e5575f5ffd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b03165b6040516001600160a01b039091168152602001610183565b348015610235575f5ffd5b5061025a604051806040016040528060058152602001640352e302e360dc1b81525081565b6040516101839190611a71565b348015610272575f5ffd5b505f54610212906001600160a01b031681565b348015610290575f5ffd5b5061014b61029f366004611aa6565b610739565b61014b6102b2366004611925565b610812565b3480156102c2575f5ffd5b5061014b6102d1366004611aa6565b6108bb565b3480156102e1575f5ffd5b5061014b6102f0366004611ac1565b610b38565b348015610300575f5ffd5b5061033c61030f366004611aa6565b60046020525f90815260409020805460018201546002909201546001600160a01b03918216929091169083565b604080516001600160a01b03948516815293909216602084015290820152606001610183565b34801561036d575f5ffd5b5061038161037c366004611aa6565b610cd0565b6040519015158152602001610183565b34801561039c575f5ffd5b5061014b6103ab366004611aa6565b610ce2565b6101796103be366004611a31565b610d1f565b3480156103ce575f5ffd5b5061014b6103dd366004611aeb565b610e7b565b6103ea610f12565b826103f6600282610f49565b61042357604051633b1c760160e21b81526001600160a01b03821660048201526024015b60405180910390fd5b335f61042e86610f6d565b80515f54919250906001600160a01b0390811690821603610463575f5461045e906001600160a01b031687610fc7565b610478565b6104786001600160a01b038216843089611059565b6040805160a081019091525f908060015b8152602001856001600160a01b031681526020015f63ffffffff1681526020015f81526020015f8152509050876001600160a01b031663948faf9082896104cf8a6110c0565b6040518463ffffffff1660e01b81526004016104ed93929190611b6e565b5f604051808303815f87803b158015610504575f5ffd5b505af1158015610516573d5f5f3e3d5ffd5b5050505050505050475f1461053e5760405163f85c3ef560e01b815260040160405180910390fd5b5061055560015f516020611cea5f395f51905f5255565b505050565b6105626110f4565b61056b82611198565b61057582826111a0565b5050565b5f61058261125c565b505f516020611cca5f395f51905f5290565b5f61059d610f12565b856105a9600282610f49565b6105d157604051633b1c760160e21b81526001600160a01b038216600482015260240161041a565b335f6105dc89610f6d565b90505f6105ee898984604001516112a5565b60208301515f54919250906001600160a01b0390811690821603610626575f54610621906001600160a01b031683610fc7565b61063b565b61063b6001600160a01b038216853085611059565b6040805160a081018252600281526001600160a01b0380871660208301525f92820192909252606081018c9052608081018b9052908c1663f938c5b1828b6106828c6110c0565b6040518463ffffffff1660e01b81526004016106a093929190611b6e565b6020604051808303815f875af11580156106bc573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106e09190611bde565b9650505050505047156107065760405163f85c3ef560e01b815260040160405180910390fd5b5061071d60015f516020611cea5f395f51905f5255565b95945050505050565b61072e61135b565b6107375f6113b6565b565b61074161135b565b61074c600282611426565b61077457604051633b1c760160e21b81526001600160a01b038216600482015260240161041a565b5f61077e82610f6d565b8051909150610797906001600160a01b0316835f61143a565b60208101516107b0906001600160a01b0316835f61143a565b6001600160a01b0382165f8181526004602052604080822080546001600160a01b031990811682556001820180549091169055600201829055517f9493af33d363a70a1cd9cc0ceb347e3c0e16b5d9371695618e1ac5fb1c255b7c9190a25050565b61081a610f12565b82610826600282610f49565b61084e57604051633b1c760160e21b81526001600160a01b038216600482015260240161041a565b335f61085986610f6d565b60208101515f54919250906001600160a01b0390811690821603610891575f5461088c906001600160a01b031687610fc7565b6108a6565b6108a66001600160a01b038216843089611059565b6040805160a081019091525f90806002610489565b6108c361135b565b6108ce6002826114c9565b6108f657604051636befb49760e01b81526001600160a01b038216600482015260240161041a565b5f8190505f816001600160a01b031663918f86746040518163ffffffff1660e01b8152600401602060405180830381865afa158015610937573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061095b9190611bde565b90505f826001600160a01b031663ec342ad06040518163ffffffff1660e01b8152600401602060405180830381865afa15801561099a573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109be9190611bf5565b90505f836001600160a01b0316639c5798396040518163ffffffff1660e01b8152600401602060405180830381865afa1580156109fd573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a219190611bf5565b9050821580610a3757506001600160a01b038116155b80610a4957506001600160a01b038216155b15610a7257604051633b1c760160e21b81526001600160a01b038616600482015260240161041a565b610a876001600160a01b038316865f1961143a565b610a9c6001600160a01b038216865f1961143a565b604080516060810182526001600160a01b0380851680835284821660208085018281528587018a81528c86165f8181526004909452888420975188549088166001600160a01b03199182161789559251600189018054919098169316929092179095559351600290950194909455935190927f7f46075c67ca5bbd3aaf82d8e324282141f27b7cba3376e5498a6b70c9931c2a91a45050505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff16159067ffffffffffffffff165f81158015610b7d5750825b90505f8267ffffffffffffffff166001148015610b995750303b155b905081158015610ba7575080155b15610bc55760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff191660011785558315610bef57845460ff60401b1916600160401b1785555b6001600160a01b038716610c1f5760405163b927e50b60e01b8152630eecae8d60e31b600482015260240161041a565b855f03610c515760405163b927e50b60e01b81526c1b585e13585d18da10dbdd5b9d609a1b600482015260240161041a565b5f80546001600160a01b0319166001600160a01b0389161790556001869055610c79336114dd565b610c816114ee565b8315610cc757845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050565b5f610cdc600283610f49565b92915050565b610cea61135b565b6001600160a01b038116610d1357604051631e4fbdf760e01b81525f600482015260240161041a565b610d1c816113b6565b50565b5f610d28610f12565b85610d34600282610f49565b610d5c57604051633b1c760160e21b81526001600160a01b038216600482015260240161041a565b335f610d6789610f6d565b80515f54919250906001600160a01b0390811690821603610d9c575f54610d97906001600160a01b031689610fc7565b610db1565b610db16001600160a01b03821684308b611059565b6040805160a081018252600181526001600160a01b0380861660208301525f92820192909252606081018b9052608081018a9052908b1663f938c5b1828a610df88b6110c0565b6040518463ffffffff1660e01b8152600401610e1693929190611b6e565b6020604051808303815f875af1158015610e32573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e569190611bde565b95505050505047156107065760405163f85c3ef560e01b815260040160405180910390fd5b82610e87600282610f49565b610eaf57604051633b1c760160e21b81526001600160a01b038216600482015260240161041a565b60405163f9f629c560e01b81526001600160a01b0385169063f9f629c590610edf90339087908790600401611c10565b5f604051808303815f87803b158015610ef6575f5ffd5b505af1158015610f08573d5f5f3e3d5ffd5b5050505050505050565b5f516020611cea5f395f51905f52805460011901610f4357604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b6001600160a01b0381165f90815260018301602052604081205415155b9392505050565b60408051606080820183525f80835260208084018290529284018190526001600160a01b03948516815260048352839020835191820184528054851682526001810154909416918101919091526002909201549082015290565b80471015610ff15760405163cf47918160e01b81524760048201526024810182905260440161041a565b5f5f836001600160a01b0316836040515f6040518083038185875af1925050503d805f811461103b576040519150601f19603f3d011682016040523d82523d5f602084013e611040565b606091505b50915091508161105357611053816114fe565b50505050565b6040516001600160a01b0384811660248301528381166044830152606482018390526110539186918216906323b872dd906084015b604051602081830303815290604052915060e01b6020820180516001600160e01b038381831617835250505050611527565b5f8115806110cf575060015482115b6110d95781610cdc565b505060015490565b60015f516020611cea5f395f51905f5255565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148061117a57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031661116e5f516020611cca5f395f51905f52546001600160a01b031690565b6001600160a01b031614155b156107375760405163703e46dd60e11b815260040160405180910390fd5b610d1c61135b565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156111fa575060408051601f3d908101601f191682019092526111f791810190611bde565b60015b61122257604051634c9c8ce360e01b81526001600160a01b038316600482015260240161041a565b5f516020611cca5f395f51905f52811461125257604051632a87526960e21b81526004810182905260240161041a565b6105558383611593565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146107375760405163703e46dd60e11b815260040160405180910390fd5b5f838302815f1985870982811083820303915050805f036112d9578382816112cf576112cf611c58565b0492505050610f66565b8084116112f0576112f060038515026011186115e8565b5f848688095f868103871696879004966002600389028118808a02820302808a02820302808a02820302808a02820302808a02820302808a02909103029181900381900460010186841190950394909402919094039290920491909117919091029150509392505050565b3361138d7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b0316146107375760405163118cdaa760e01b815233600482015260240161041a565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b5f610f66836001600160a01b0384166115f9565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663095ea7b360e01b17905261148b84826116dc565b611053576040516001600160a01b0384811660248301525f60448301526114bf91869182169063095ea7b39060640161108e565b6110538482611527565b5f610f66836001600160a01b038416611725565b6114e5611771565b610d1c816117ba565b6114f6611771565b6107376117c2565b80511561150e5780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b5f5f60205f8451602086015f885af180611546576040513d5f823e3d81fd5b50505f513d9150811561155d57806001141561156a565b6001600160a01b0384163b155b1561105357604051635274afe760e01b81526001600160a01b038516600482015260240161041a565b61159c826117ca565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a28051156115e057610555828261182d565b610575611896565b634e487b715f52806020526024601cfd5b5f81815260018301602052604081205480156116d3575f61161b600183611c6c565b85549091505f9061162e90600190611c6c565b905080821461168d575f865f01828154811061164c5761164c611c8b565b905f5260205f200154905080875f01848154811061166c5761166c611c8b565b5f918252602080832090910192909255918252600188019052604090208390555b855486908061169e5761169e611c9f565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050610cdc565b5f915050610cdc565b5f5f5f5f60205f8651602088015f8a5af192503d91505f51905082801561171b5750811561170d578060011461171b565b5f866001600160a01b03163b115b9695505050505050565b5f81815260018301602052604081205461176a57508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610cdc565b505f610cdc565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff1661073757604051631afcd79f60e31b815260040160405180910390fd5b610cea611771565b6110e1611771565b806001600160a01b03163b5f036117ff57604051634c9c8ce360e01b81526001600160a01b038216600482015260240161041a565b5f516020611cca5f395f51905f5280546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b0316846040516118499190611cb3565b5f60405180830381855af49150503d805f8114611881576040519150601f19603f3d011682016040523d82523d5f602084013e611886565b606091505b509150915061071d8583836118b5565b34156107375760405163b398979f60e01b815260040160405180910390fd5b6060826118ca576118c5826114fe565b610f66565b81511580156118e157506001600160a01b0384163b155b1561190a57604051639996b31560e01b81526001600160a01b038516600482015260240161041a565b5080610f66565b6001600160a01b0381168114610d1c575f5ffd5b5f5f5f60608486031215611937575f5ffd5b833561194281611911565b95602085013595506040909401359392505050565b634e487b7160e01b5f52604160045260245ffd5b5f5f6040838503121561197c575f5ffd5b823561198781611911565b9150602083013567ffffffffffffffff8111156119a2575f5ffd5b8301601f810185136119b2575f5ffd5b803567ffffffffffffffff8111156119cc576119cc611957565b604051601f8201601f19908116603f0116810167ffffffffffffffff811182821017156119fb576119fb611957565b604052818152828201602001871015611a12575f5ffd5b816020840160208301375f602083830101528093505050509250929050565b5f5f5f5f5f60a08688031215611a45575f5ffd5b8535611a5081611911565b97602087013597506040870135966060810135965060800135945092505050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b5f60208284031215611ab6575f5ffd5b8135610f6681611911565b5f5f60408385031215611ad2575f5ffd5b8235611add81611911565b946020939093013593505050565b5f5f5f60408486031215611afd575f5ffd5b8335611b0881611911565b9250602084013567ffffffffffffffff811115611b23575f5ffd5b8401601f81018613611b33575f5ffd5b803567ffffffffffffffff811115611b49575f5ffd5b8660208260051b8401011115611b5d575f5ffd5b939660209190910195509293505050565b835160e082019060038110611b9157634e487b7160e01b5f52602160045260245ffd5b82526020858101516001600160a01b03169083015260408086015163ffffffff1690830152606080860151908301526080948501519482019490945260a081019290925260c09091015290565b5f60208284031215611bee575f5ffd5b5051919050565b5f60208284031215611c05575f5ffd5b8151610f6681611911565b6001600160a01b038416815260406020820181905281018290525f6001600160fb1b03831115611c3e575f5ffd5b8260051b8085606085013791909101606001949350505050565b634e487b7160e01b5f52601260045260245ffd5b81810381811115610cdc57634e487b7160e01b5f52601160045260245ffd5b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52603160045260245ffd5b5f82518060208501845e5f92019182525091905056fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a2646970667358221220db9d682c467d48ef2f622492bdbee7a9966f022858b065e73cedbeabb7af5fa364736f6c634300081c0033",
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

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_RouterImpl *RouterImplCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RouterImpl.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_RouterImpl *RouterImplSession) WETH() (common.Address, error) {
	return _RouterImpl.Contract.WETH(&_RouterImpl.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_RouterImpl *RouterImplCallerSession) WETH() (common.Address, error) {
	return _RouterImpl.Contract.WETH(&_RouterImpl.CallOpts)
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

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RouterImpl *RouterImplCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RouterImpl.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RouterImpl *RouterImplSession) Owner() (common.Address, error) {
	return _RouterImpl.Contract.Owner(&_RouterImpl.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RouterImpl *RouterImplCallerSession) Owner() (common.Address, error) {
	return _RouterImpl.Contract.Owner(&_RouterImpl.CallOpts)
}

// PairInfo is a free data retrieval call binding the contract method 0xdc0a4229.
//
// Solidity: function pairInfo(address pair) view returns(address BASE, address QUOTE, uint256 DENOMINATOR)
func (_RouterImpl *RouterImplCaller) PairInfo(opts *bind.CallOpts, pair common.Address) (struct {
	BASE        common.Address
	QUOTE       common.Address
	DENOMINATOR *big.Int
}, error) {
	var out []interface{}
	err := _RouterImpl.contract.Call(opts, &out, "pairInfo", pair)

	outstruct := new(struct {
		BASE        common.Address
		QUOTE       common.Address
		DENOMINATOR *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BASE = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.QUOTE = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.DENOMINATOR = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PairInfo is a free data retrieval call binding the contract method 0xdc0a4229.
//
// Solidity: function pairInfo(address pair) view returns(address BASE, address QUOTE, uint256 DENOMINATOR)
func (_RouterImpl *RouterImplSession) PairInfo(pair common.Address) (struct {
	BASE        common.Address
	QUOTE       common.Address
	DENOMINATOR *big.Int
}, error) {
	return _RouterImpl.Contract.PairInfo(&_RouterImpl.CallOpts, pair)
}

// PairInfo is a free data retrieval call binding the contract method 0xdc0a4229.
//
// Solidity: function pairInfo(address pair) view returns(address BASE, address QUOTE, uint256 DENOMINATOR)
func (_RouterImpl *RouterImplCallerSession) PairInfo(pair common.Address) (struct {
	BASE        common.Address
	QUOTE       common.Address
	DENOMINATOR *big.Int
}, error) {
	return _RouterImpl.Contract.PairInfo(&_RouterImpl.CallOpts, pair)
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

// AddPair is a paid mutator transaction binding the contract method 0xc2b7bbb6.
//
// Solidity: function addPair(address pair) returns()
func (_RouterImpl *RouterImplTransactor) AddPair(opts *bind.TransactOpts, pair common.Address) (*types.Transaction, error) {
	return _RouterImpl.contract.Transact(opts, "addPair", pair)
}

// AddPair is a paid mutator transaction binding the contract method 0xc2b7bbb6.
//
// Solidity: function addPair(address pair) returns()
func (_RouterImpl *RouterImplSession) AddPair(pair common.Address) (*types.Transaction, error) {
	return _RouterImpl.Contract.AddPair(&_RouterImpl.TransactOpts, pair)
}

// AddPair is a paid mutator transaction binding the contract method 0xc2b7bbb6.
//
// Solidity: function addPair(address pair) returns()
func (_RouterImpl *RouterImplTransactorSession) AddPair(pair common.Address) (*types.Transaction, error) {
	return _RouterImpl.Contract.AddPair(&_RouterImpl.TransactOpts, pair)
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

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address weth, uint256 _maxMatchCount) returns()
func (_RouterImpl *RouterImplTransactor) Initialize(opts *bind.TransactOpts, weth common.Address, _maxMatchCount *big.Int) (*types.Transaction, error) {
	return _RouterImpl.contract.Transact(opts, "initialize", weth, _maxMatchCount)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address weth, uint256 _maxMatchCount) returns()
func (_RouterImpl *RouterImplSession) Initialize(weth common.Address, _maxMatchCount *big.Int) (*types.Transaction, error) {
	return _RouterImpl.Contract.Initialize(&_RouterImpl.TransactOpts, weth, _maxMatchCount)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address weth, uint256 _maxMatchCount) returns()
func (_RouterImpl *RouterImplTransactorSession) Initialize(weth common.Address, _maxMatchCount *big.Int) (*types.Transaction, error) {
	return _RouterImpl.Contract.Initialize(&_RouterImpl.TransactOpts, weth, _maxMatchCount)
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

// RemovePair is a paid mutator transaction binding the contract method 0xaf6c9c1d.
//
// Solidity: function removePair(address pair) returns()
func (_RouterImpl *RouterImplTransactor) RemovePair(opts *bind.TransactOpts, pair common.Address) (*types.Transaction, error) {
	return _RouterImpl.contract.Transact(opts, "removePair", pair)
}

// RemovePair is a paid mutator transaction binding the contract method 0xaf6c9c1d.
//
// Solidity: function removePair(address pair) returns()
func (_RouterImpl *RouterImplSession) RemovePair(pair common.Address) (*types.Transaction, error) {
	return _RouterImpl.Contract.RemovePair(&_RouterImpl.TransactOpts, pair)
}

// RemovePair is a paid mutator transaction binding the contract method 0xaf6c9c1d.
//
// Solidity: function removePair(address pair) returns()
func (_RouterImpl *RouterImplTransactorSession) RemovePair(pair common.Address) (*types.Transaction, error) {
	return _RouterImpl.Contract.RemovePair(&_RouterImpl.TransactOpts, pair)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RouterImpl *RouterImplTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RouterImpl.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RouterImpl *RouterImplSession) RenounceOwnership() (*types.Transaction, error) {
	return _RouterImpl.Contract.RenounceOwnership(&_RouterImpl.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RouterImpl *RouterImplTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RouterImpl.Contract.RenounceOwnership(&_RouterImpl.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RouterImpl *RouterImplTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RouterImpl.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RouterImpl *RouterImplSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RouterImpl.Contract.TransferOwnership(&_RouterImpl.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RouterImpl *RouterImplTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RouterImpl.Contract.TransferOwnership(&_RouterImpl.TransactOpts, newOwner)
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

// RouterImplOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RouterImpl contract.
type RouterImplOwnershipTransferredIterator struct {
	Event *RouterImplOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RouterImplOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterImplOwnershipTransferred)
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
		it.Event = new(RouterImplOwnershipTransferred)
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
func (it *RouterImplOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterImplOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterImplOwnershipTransferred represents a OwnershipTransferred event raised by the RouterImpl contract.
type RouterImplOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RouterImpl *RouterImplFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RouterImplOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RouterImpl.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RouterImplOwnershipTransferredIterator{contract: _RouterImpl.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RouterImpl *RouterImplFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RouterImplOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RouterImpl.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterImplOwnershipTransferred)
				if err := _RouterImpl.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RouterImpl *RouterImplFilterer) ParseOwnershipTransferred(log types.Log) (*RouterImplOwnershipTransferred, error) {
	event := new(RouterImplOwnershipTransferred)
	if err := _RouterImpl.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
