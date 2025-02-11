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

// MarketImplMetaData contains all meta data concerning the MarketImpl contract.
var MarketImplMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"CROSS_DEX\",\"outputs\":[{\"internalType\":\"contractICrossDex\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"QUOTE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allPairs\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"bases\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"pairs\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"}],\"name\":\"baseToPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quoteTickSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseTickSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feePermil\",\"type\":\"uint256\"}],\"name\":\"createPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deployed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeCollector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeCollector\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_quote\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pairImpl\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pairImpl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"PairCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"EnumerableMapNonexistentKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"MarketAlreadyCreatedBaseAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MarketDeployPair\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"MarketInvalidBaseAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"MarketInvalidInitializeData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"b820d8aa": "CROSS_DEX()",
		"9c579839": "QUOTE()",
		"ad3cb1cc": "UPGRADE_INTERFACE_VERSION()",
		"c97682f8": "allPairs()",
		"364e9a19": "baseToPair(address)",
		"ecaa3911": "createPair(address,uint256,uint256,uint256)",
		"f905c15a": "deployed()",
		"c415b95c": "feeCollector()",
		"1459457a": "initialize(address,address,address,address,address)",
		"8da5cb5b": "owner()",
		"35f7d456": "pairImpl()",
		"52d1902d": "proxiableUUID()",
		"715018a6": "renounceOwnership()",
		"f887ea40": "router()",
		"f2fde38b": "transferOwnership(address)",
		"4f1ef286": "upgradeToAndCall(address,bytes)",
	},
	Bin: "0x60a0604052306080523480156012575f5ffd5b506080516117576100395f395f8181610a2001528181610a490152610b8d01526117575ff3fe6080604052600436106100ef575f3560e01c8063ad3cb1cc11610087578063ecaa391111610057578063ecaa3911146102b0578063f2fde38b146102cf578063f887ea40146102ee578063f905c15a1461030d575f5ffd5b8063ad3cb1cc14610213578063b820d8aa14610250578063c415b95c1461026f578063c97682f81461028e575f5ffd5b806352d1902d116100c257806352d1902d14610182578063715018a6146101a45780638da5cb5b146101b85780639c579839146101f4575f5ffd5b80631459457a146100f357806335f7d45614610114578063364e9a19146101505780634f1ef2861461016f575b5f5ffd5b3480156100fe575f5ffd5b5061011261010d36600461107e565b610321565b005b34801561011f575f5ffd5b50600554610133906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b34801561015b575f5ffd5b5061013361016a3660046110df565b61058f565b61011261017d36600461110c565b6105a1565b34801561018d575f5ffd5b506101966105c0565b604051908152602001610147565b3480156101af575f5ffd5b506101126105db565b3480156101c3575f5ffd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b0316610133565b3480156101ff575f5ffd5b50600254610133906001600160a01b031681565b34801561021e575f5ffd5b50610243604051806040016040528060058152602001640352e302e360dc1b81525081565b60405161014791906111fe565b34801561025b575f5ffd5b50600154610133906001600160a01b031681565b34801561027a575f5ffd5b50600454610133906001600160a01b031681565b348015610299575f5ffd5b506102a26105ee565b604051610147929190611253565b3480156102bb575f5ffd5b506101336102ca366004611277565b6106f0565b3480156102da575f5ffd5b506101126102e93660046110df565b6109ac565b3480156102f9575f5ffd5b50600354610133906001600160a01b031681565b348015610318575f5ffd5b506101965f5481565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff16159067ffffffffffffffff165f811580156103665750825b90505f8267ffffffffffffffff1660011480156103825750303b155b905081158015610390575080155b156103ae5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff1916600117855583156103d857845460ff60401b1916600160401b1785555b6001600160a01b038a1661040e5760405163a25d2db560e01b81526437bbb732b960d91b60048201526024015b60405180910390fd5b6001600160a01b0389166104405760405163a25d2db560e01b8152653937baba32b960d11b6004820152602401610405565b6001600160a01b0388166104785760405163a25d2db560e01b81526b3332b2a1b7b63632b1ba37b960a11b6004820152602401610405565b6001600160a01b0387166104a95760405163a25d2db560e01b81526471756f746560d81b6004820152602401610405565b6001600160a01b0386166104dd5760405163a25d2db560e01b8152671c185a5c925b5c1b60c21b6004820152602401610405565b435f5533600180546001600160a01b03199081166001600160a01b03938416179091556002805482168a84161790556003805482168c84161790556004805482168b84161790556005805490911691881691909117905561053d8a6109e9565b831561058357845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050505050565b5f61059b6006836109fa565b92915050565b6105a9610a15565b6105b282610ab9565b6105bc8282610ac1565b5050565b5f6105c9610b82565b505f5160206117025f395f51905f5290565b6105e3610bcb565b6105ec5f610c26565b565b6060805f6105fc6006610c96565b90508067ffffffffffffffff811115610617576106176110f8565b604051908082528060200260200182016040528015610640578160200160208202803683370190505b5092508067ffffffffffffffff81111561065c5761065c6110f8565b604051908082528060200260200182016040528015610685578160200160208202803683370190505b5091505f5b818110156106ea5761069d600682610ca0565b8583815181106106af576106af6112ad565b602002602001018584815181106106c8576106c86112ad565b6001600160a01b0393841660209182029290920101529116905260010161068a565b50509091565b5f6106f9610bcb565b6001600160a01b03851661072b5760405163083e095360e21b81526001600160a01b0386166004820152602401610405565b5f856001600160a01b031663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa158015610768573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061078c91906112c1565b60ff169050805f036107bc5760405163083e095360e21b81526001600160a01b0387166004820152602401610405565b6107c7600687610cbb565b156107f05760405163304ff3f760e21b81526001600160a01b0387166004820152602401610405565b600554600354600254600454604080516001600160a01b03948516602482015292841660448401528a84166064840152608483018a905260a4830189905290831660c483015260e48083018890528151808403909101815261010490920181526020820180516001600160e01b031663297aaadd60e11b1790525191909216919061087a90611056565b6108859291906112e1565b604051809103905ff08015801561089e573d5f5f3e3d5ffd5b5091506001600160a01b0382166108c857604051633ff6a26160e01b815260040160405180910390fd5b6108d460068784610ccf565b6108fc5760405163304ff3f760e21b81526001600160a01b0387166004820152602401610405565b60015460405163e9f7206b60e01b81526001600160a01b0384811660048301529091169063e9f7206b906024015f604051808303815f87803b158015610940575f5ffd5b505af1158015610952573d5f5f3e3d5ffd5b50505050856001600160a01b0316826001600160a01b03167fc1db9ba7c4b7ce660fe8d17bbcf07167549381df2abd694a970bd1402d86d3134260405161099b91815260200190565b60405180910390a350949350505050565b6109b4610bcb565b6001600160a01b0381166109dd57604051631e4fbdf760e01b81525f6004820152602401610405565b6109e681610c26565b50565b6109f1610cef565b6109e681610d38565b5f610a0e836001600160a01b038416610d40565b9392505050565b306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480610a9b57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316610a8f5f5160206117025f395f51905f52546001600160a01b031690565b6001600160a01b031614155b156105ec5760405163703e46dd60e11b815260040160405180910390fd5b6109e6610bcb565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015610b1b575060408051601f3d908101601f19168201909252610b1891810190611304565b60015b610b4357604051634c9c8ce360e01b81526001600160a01b0383166004820152602401610405565b5f5160206117025f395f51905f528114610b7357604051632a87526960e21b815260048101829052602401610405565b610b7d8383610d86565b505050565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146105ec5760405163703e46dd60e11b815260040160405180910390fd5b33610bfd7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b0316146105ec5760405163118cdaa760e01b8152336004820152602401610405565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b5f61059b82610ddb565b5f808080610cae8686610de5565b9097909650945050505050565b5f610a0e836001600160a01b038416610e0e565b5f610ce7846001600160a01b03808616908516610e19565b949350505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff166105ec57604051631afcd79f60e31b815260040160405180910390fd5b6109b4610cef565b5f81815260028301602052604081205480158015610d655750610d638484610e0e565b155b15610a0e5760405163015ab34360e11b815260048101849052602401610405565b610d8f82610e35565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115610dd357610b7d8282610e98565b6105bc610f0a565b5f61059b82610f29565b5f8080610df28585610f32565b5f81815260029690960160205260409095205494959350505050565b5f610a0e8383610f3d565b5f8281526002840160205260408120829055610ce78484610f54565b806001600160a01b03163b5f03610e6a57604051634c9c8ce360e01b81526001600160a01b0382166004820152602401610405565b5f5160206117025f395f51905f5280546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b031684604051610eb4919061131b565b5f60405180830381855af49150503d805f8114610eec576040519150601f19603f3d011682016040523d82523d5f602084013e610ef1565b606091505b5091509150610f01858383610f5f565b95945050505050565b34156105ec5760405163b398979f60e01b815260040160405180910390fd5b5f61059b825490565b5f610a0e8383610fbb565b5f8181526001830160205260408120541515610a0e565b5f610a0e8383610fe1565b606082610f7457610f6f8261102d565b610a0e565b8151158015610f8b57506001600160a01b0384163b155b15610fb457604051639996b31560e01b81526001600160a01b0385166004820152602401610405565b5080610a0e565b5f825f018281548110610fd057610fd06112ad565b905f5260205f200154905092915050565b5f81815260018301602052604081205461102657508154600181810184555f84815260208082209093018490558454848252828601909352604090209190915561059b565b505f61059b565b80511561103d5780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b6103d08061133283390190565b80356001600160a01b0381168114611079575f5ffd5b919050565b5f5f5f5f5f60a08688031215611092575f5ffd5b61109b86611063565b94506110a960208701611063565b93506110b760408701611063565b92506110c560608701611063565b91506110d360808701611063565b90509295509295909350565b5f602082840312156110ef575f5ffd5b610a0e82611063565b634e487b7160e01b5f52604160045260245ffd5b5f5f6040838503121561111d575f5ffd5b61112683611063565b9150602083013567ffffffffffffffff811115611141575f5ffd5b8301601f81018513611151575f5ffd5b803567ffffffffffffffff81111561116b5761116b6110f8565b604051601f8201601f19908116603f0116810167ffffffffffffffff8111828210171561119a5761119a6110f8565b6040528181528282016020018710156111b1575f5ffd5b816020840160208301375f602083830101528093505050509250929050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f610a0e60208301846111d0565b5f8151808452602084019350602083015f5b828110156112495781516001600160a01b0316865260209586019590910190600101611222565b5093949350505050565b604081525f6112656040830185611210565b8281036020840152610f018185611210565b5f5f5f5f6080858703121561128a575f5ffd5b61129385611063565b966020860135965060408601359560600135945092505050565b634e487b7160e01b5f52603260045260245ffd5b5f602082840312156112d1575f5ffd5b815160ff81168114610a0e575f5ffd5b6001600160a01b03831681526040602082018190525f90610ce7908301846111d0565b5f60208284031215611314575f5ffd5b5051919050565b5f82518060208501845e5f92019182525091905056fe60806040526040516103d03803806103d08339810160408190526100229161023c565b61002c8282610033565b5050610321565b61003c82610091565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a280511561008557610080828261010c565b505050565b61008d61017f565b5050565b806001600160a01b03163b5f036100cb57604051634c9c8ce360e01b81526001600160a01b03821660048201526024015b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b031684604051610128919061030b565b5f60405180830381855af49150503d805f8114610160576040519150601f19603f3d011682016040523d82523d5f602084013e610165565b606091505b5090925090506101768583836101a0565b95945050505050565b341561019e5760405163b398979f60e01b815260040160405180910390fd5b565b6060826101b5576101b0826101ff565b6101f8565b81511580156101cc57506001600160a01b0384163b155b156101f557604051639996b31560e01b81526001600160a01b03851660048201526024016100c2565b50805b9392505050565b80511561020f5780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b634e487b7160e01b5f52604160045260245ffd5b5f5f6040838503121561024d575f5ffd5b82516001600160a01b0381168114610263575f5ffd5b60208401519092506001600160401b0381111561027e575f5ffd5b8301601f8101851361028e575f5ffd5b80516001600160401b038111156102a7576102a7610228565b604051601f8201601f19908116603f011681016001600160401b03811182821017156102d5576102d5610228565b6040528181528282016020018710156102ec575f5ffd5b8160208401602083015e5f602083830101528093505050509250929050565b5f82518060208501845e5f920191825250919050565b60a38061032d5f395ff3fe6080604052600a600c565b005b60186014601a565b6050565b565b5f604b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b905090565b365f5f375f5f365f845af43d5f5f3e8080156069573d5ff35b3d5ffdfea26469706673582212205990887f0037c04aa22ef84499f094f4341c02e52649e71080dd1a78b7edba3e64736f6c634300081c0033360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbca2646970667358221220035ce9a37d5a06cdb43fe62da8a70bf86493427b226ecd35008304b4cb2c44f564736f6c634300081c0033",
}

// MarketImplABI is the input ABI used to generate the binding from.
// Deprecated: Use MarketImplMetaData.ABI instead.
var MarketImplABI = MarketImplMetaData.ABI

// Deprecated: Use MarketImplMetaData.Sigs instead.
// MarketImplFuncSigs maps the 4-byte function signature to its string representation.
var MarketImplFuncSigs = MarketImplMetaData.Sigs

// MarketImplBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MarketImplMetaData.Bin instead.
var MarketImplBin = MarketImplMetaData.Bin

// DeployMarketImpl deploys a new Ethereum contract, binding an instance of MarketImpl to it.
func DeployMarketImpl(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MarketImpl, error) {
	parsed, err := MarketImplMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MarketImplBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MarketImpl{MarketImplCaller: MarketImplCaller{contract: contract}, MarketImplTransactor: MarketImplTransactor{contract: contract}, MarketImplFilterer: MarketImplFilterer{contract: contract}}, nil
}

// MarketImpl is an auto generated Go binding around an Ethereum contract.
type MarketImpl struct {
	MarketImplCaller     // Read-only binding to the contract
	MarketImplTransactor // Write-only binding to the contract
	MarketImplFilterer   // Log filterer for contract events
}

// MarketImplCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketImplCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketImplTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketImplTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketImplFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketImplFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketImplSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketImplSession struct {
	Contract     *MarketImpl       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketImplCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketImplCallerSession struct {
	Contract *MarketImplCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MarketImplTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketImplTransactorSession struct {
	Contract     *MarketImplTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MarketImplRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketImplRaw struct {
	Contract *MarketImpl // Generic contract binding to access the raw methods on
}

// MarketImplCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketImplCallerRaw struct {
	Contract *MarketImplCaller // Generic read-only contract binding to access the raw methods on
}

// MarketImplTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketImplTransactorRaw struct {
	Contract *MarketImplTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarketImpl creates a new instance of MarketImpl, bound to a specific deployed contract.
func NewMarketImpl(address common.Address, backend bind.ContractBackend) (*MarketImpl, error) {
	contract, err := bindMarketImpl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MarketImpl{MarketImplCaller: MarketImplCaller{contract: contract}, MarketImplTransactor: MarketImplTransactor{contract: contract}, MarketImplFilterer: MarketImplFilterer{contract: contract}}, nil
}

// NewMarketImplCaller creates a new read-only instance of MarketImpl, bound to a specific deployed contract.
func NewMarketImplCaller(address common.Address, caller bind.ContractCaller) (*MarketImplCaller, error) {
	contract, err := bindMarketImpl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketImplCaller{contract: contract}, nil
}

// NewMarketImplTransactor creates a new write-only instance of MarketImpl, bound to a specific deployed contract.
func NewMarketImplTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketImplTransactor, error) {
	contract, err := bindMarketImpl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketImplTransactor{contract: contract}, nil
}

// NewMarketImplFilterer creates a new log filterer instance of MarketImpl, bound to a specific deployed contract.
func NewMarketImplFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketImplFilterer, error) {
	contract, err := bindMarketImpl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketImplFilterer{contract: contract}, nil
}

// bindMarketImpl binds a generic wrapper to an already deployed contract.
func bindMarketImpl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MarketImplMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MarketImpl *MarketImplRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MarketImpl.Contract.MarketImplCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MarketImpl *MarketImplRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketImpl.Contract.MarketImplTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MarketImpl *MarketImplRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MarketImpl.Contract.MarketImplTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MarketImpl *MarketImplCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MarketImpl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MarketImpl *MarketImplTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketImpl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MarketImpl *MarketImplTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MarketImpl.Contract.contract.Transact(opts, method, params...)
}

// CROSSDEX is a free data retrieval call binding the contract method 0xb820d8aa.
//
// Solidity: function CROSS_DEX() view returns(address)
func (_MarketImpl *MarketImplCaller) CROSSDEX(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarketImpl.contract.Call(opts, &out, "CROSS_DEX")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CROSSDEX is a free data retrieval call binding the contract method 0xb820d8aa.
//
// Solidity: function CROSS_DEX() view returns(address)
func (_MarketImpl *MarketImplSession) CROSSDEX() (common.Address, error) {
	return _MarketImpl.Contract.CROSSDEX(&_MarketImpl.CallOpts)
}

// CROSSDEX is a free data retrieval call binding the contract method 0xb820d8aa.
//
// Solidity: function CROSS_DEX() view returns(address)
func (_MarketImpl *MarketImplCallerSession) CROSSDEX() (common.Address, error) {
	return _MarketImpl.Contract.CROSSDEX(&_MarketImpl.CallOpts)
}

// QUOTE is a free data retrieval call binding the contract method 0x9c579839.
//
// Solidity: function QUOTE() view returns(address)
func (_MarketImpl *MarketImplCaller) QUOTE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarketImpl.contract.Call(opts, &out, "QUOTE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// QUOTE is a free data retrieval call binding the contract method 0x9c579839.
//
// Solidity: function QUOTE() view returns(address)
func (_MarketImpl *MarketImplSession) QUOTE() (common.Address, error) {
	return _MarketImpl.Contract.QUOTE(&_MarketImpl.CallOpts)
}

// QUOTE is a free data retrieval call binding the contract method 0x9c579839.
//
// Solidity: function QUOTE() view returns(address)
func (_MarketImpl *MarketImplCallerSession) QUOTE() (common.Address, error) {
	return _MarketImpl.Contract.QUOTE(&_MarketImpl.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_MarketImpl *MarketImplCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MarketImpl.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_MarketImpl *MarketImplSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _MarketImpl.Contract.UPGRADEINTERFACEVERSION(&_MarketImpl.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_MarketImpl *MarketImplCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _MarketImpl.Contract.UPGRADEINTERFACEVERSION(&_MarketImpl.CallOpts)
}

// AllPairs is a free data retrieval call binding the contract method 0xc97682f8.
//
// Solidity: function allPairs() view returns(address[] bases, address[] pairs)
func (_MarketImpl *MarketImplCaller) AllPairs(opts *bind.CallOpts) (struct {
	Bases []common.Address
	Pairs []common.Address
}, error) {
	var out []interface{}
	err := _MarketImpl.contract.Call(opts, &out, "allPairs")

	outstruct := new(struct {
		Bases []common.Address
		Pairs []common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Bases = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Pairs = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

// AllPairs is a free data retrieval call binding the contract method 0xc97682f8.
//
// Solidity: function allPairs() view returns(address[] bases, address[] pairs)
func (_MarketImpl *MarketImplSession) AllPairs() (struct {
	Bases []common.Address
	Pairs []common.Address
}, error) {
	return _MarketImpl.Contract.AllPairs(&_MarketImpl.CallOpts)
}

// AllPairs is a free data retrieval call binding the contract method 0xc97682f8.
//
// Solidity: function allPairs() view returns(address[] bases, address[] pairs)
func (_MarketImpl *MarketImplCallerSession) AllPairs() (struct {
	Bases []common.Address
	Pairs []common.Address
}, error) {
	return _MarketImpl.Contract.AllPairs(&_MarketImpl.CallOpts)
}

// BaseToPair is a free data retrieval call binding the contract method 0x364e9a19.
//
// Solidity: function baseToPair(address base) view returns(address)
func (_MarketImpl *MarketImplCaller) BaseToPair(opts *bind.CallOpts, base common.Address) (common.Address, error) {
	var out []interface{}
	err := _MarketImpl.contract.Call(opts, &out, "baseToPair", base)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BaseToPair is a free data retrieval call binding the contract method 0x364e9a19.
//
// Solidity: function baseToPair(address base) view returns(address)
func (_MarketImpl *MarketImplSession) BaseToPair(base common.Address) (common.Address, error) {
	return _MarketImpl.Contract.BaseToPair(&_MarketImpl.CallOpts, base)
}

// BaseToPair is a free data retrieval call binding the contract method 0x364e9a19.
//
// Solidity: function baseToPair(address base) view returns(address)
func (_MarketImpl *MarketImplCallerSession) BaseToPair(base common.Address) (common.Address, error) {
	return _MarketImpl.Contract.BaseToPair(&_MarketImpl.CallOpts, base)
}

// Deployed is a free data retrieval call binding the contract method 0xf905c15a.
//
// Solidity: function deployed() view returns(uint256)
func (_MarketImpl *MarketImplCaller) Deployed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MarketImpl.contract.Call(opts, &out, "deployed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Deployed is a free data retrieval call binding the contract method 0xf905c15a.
//
// Solidity: function deployed() view returns(uint256)
func (_MarketImpl *MarketImplSession) Deployed() (*big.Int, error) {
	return _MarketImpl.Contract.Deployed(&_MarketImpl.CallOpts)
}

// Deployed is a free data retrieval call binding the contract method 0xf905c15a.
//
// Solidity: function deployed() view returns(uint256)
func (_MarketImpl *MarketImplCallerSession) Deployed() (*big.Int, error) {
	return _MarketImpl.Contract.Deployed(&_MarketImpl.CallOpts)
}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_MarketImpl *MarketImplCaller) FeeCollector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarketImpl.contract.Call(opts, &out, "feeCollector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_MarketImpl *MarketImplSession) FeeCollector() (common.Address, error) {
	return _MarketImpl.Contract.FeeCollector(&_MarketImpl.CallOpts)
}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_MarketImpl *MarketImplCallerSession) FeeCollector() (common.Address, error) {
	return _MarketImpl.Contract.FeeCollector(&_MarketImpl.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MarketImpl *MarketImplCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarketImpl.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MarketImpl *MarketImplSession) Owner() (common.Address, error) {
	return _MarketImpl.Contract.Owner(&_MarketImpl.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MarketImpl *MarketImplCallerSession) Owner() (common.Address, error) {
	return _MarketImpl.Contract.Owner(&_MarketImpl.CallOpts)
}

// PairImpl is a free data retrieval call binding the contract method 0x35f7d456.
//
// Solidity: function pairImpl() view returns(address)
func (_MarketImpl *MarketImplCaller) PairImpl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarketImpl.contract.Call(opts, &out, "pairImpl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PairImpl is a free data retrieval call binding the contract method 0x35f7d456.
//
// Solidity: function pairImpl() view returns(address)
func (_MarketImpl *MarketImplSession) PairImpl() (common.Address, error) {
	return _MarketImpl.Contract.PairImpl(&_MarketImpl.CallOpts)
}

// PairImpl is a free data retrieval call binding the contract method 0x35f7d456.
//
// Solidity: function pairImpl() view returns(address)
func (_MarketImpl *MarketImplCallerSession) PairImpl() (common.Address, error) {
	return _MarketImpl.Contract.PairImpl(&_MarketImpl.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_MarketImpl *MarketImplCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MarketImpl.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_MarketImpl *MarketImplSession) ProxiableUUID() ([32]byte, error) {
	return _MarketImpl.Contract.ProxiableUUID(&_MarketImpl.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_MarketImpl *MarketImplCallerSession) ProxiableUUID() ([32]byte, error) {
	return _MarketImpl.Contract.ProxiableUUID(&_MarketImpl.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_MarketImpl *MarketImplCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarketImpl.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_MarketImpl *MarketImplSession) Router() (common.Address, error) {
	return _MarketImpl.Contract.Router(&_MarketImpl.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_MarketImpl *MarketImplCallerSession) Router() (common.Address, error) {
	return _MarketImpl.Contract.Router(&_MarketImpl.CallOpts)
}

// CreatePair is a paid mutator transaction binding the contract method 0xecaa3911.
//
// Solidity: function createPair(address base, uint256 quoteTickSize, uint256 baseTickSize, uint256 feePermil) returns(address pair)
func (_MarketImpl *MarketImplTransactor) CreatePair(opts *bind.TransactOpts, base common.Address, quoteTickSize *big.Int, baseTickSize *big.Int, feePermil *big.Int) (*types.Transaction, error) {
	return _MarketImpl.contract.Transact(opts, "createPair", base, quoteTickSize, baseTickSize, feePermil)
}

// CreatePair is a paid mutator transaction binding the contract method 0xecaa3911.
//
// Solidity: function createPair(address base, uint256 quoteTickSize, uint256 baseTickSize, uint256 feePermil) returns(address pair)
func (_MarketImpl *MarketImplSession) CreatePair(base common.Address, quoteTickSize *big.Int, baseTickSize *big.Int, feePermil *big.Int) (*types.Transaction, error) {
	return _MarketImpl.Contract.CreatePair(&_MarketImpl.TransactOpts, base, quoteTickSize, baseTickSize, feePermil)
}

// CreatePair is a paid mutator transaction binding the contract method 0xecaa3911.
//
// Solidity: function createPair(address base, uint256 quoteTickSize, uint256 baseTickSize, uint256 feePermil) returns(address pair)
func (_MarketImpl *MarketImplTransactorSession) CreatePair(base common.Address, quoteTickSize *big.Int, baseTickSize *big.Int, feePermil *big.Int) (*types.Transaction, error) {
	return _MarketImpl.Contract.CreatePair(&_MarketImpl.TransactOpts, base, quoteTickSize, baseTickSize, feePermil)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _owner, address _router, address _feeCollector, address _quote, address _pairImpl) returns()
func (_MarketImpl *MarketImplTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _router common.Address, _feeCollector common.Address, _quote common.Address, _pairImpl common.Address) (*types.Transaction, error) {
	return _MarketImpl.contract.Transact(opts, "initialize", _owner, _router, _feeCollector, _quote, _pairImpl)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _owner, address _router, address _feeCollector, address _quote, address _pairImpl) returns()
func (_MarketImpl *MarketImplSession) Initialize(_owner common.Address, _router common.Address, _feeCollector common.Address, _quote common.Address, _pairImpl common.Address) (*types.Transaction, error) {
	return _MarketImpl.Contract.Initialize(&_MarketImpl.TransactOpts, _owner, _router, _feeCollector, _quote, _pairImpl)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _owner, address _router, address _feeCollector, address _quote, address _pairImpl) returns()
func (_MarketImpl *MarketImplTransactorSession) Initialize(_owner common.Address, _router common.Address, _feeCollector common.Address, _quote common.Address, _pairImpl common.Address) (*types.Transaction, error) {
	return _MarketImpl.Contract.Initialize(&_MarketImpl.TransactOpts, _owner, _router, _feeCollector, _quote, _pairImpl)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MarketImpl *MarketImplTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketImpl.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MarketImpl *MarketImplSession) RenounceOwnership() (*types.Transaction, error) {
	return _MarketImpl.Contract.RenounceOwnership(&_MarketImpl.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MarketImpl *MarketImplTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MarketImpl.Contract.RenounceOwnership(&_MarketImpl.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MarketImpl *MarketImplTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MarketImpl.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MarketImpl *MarketImplSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MarketImpl.Contract.TransferOwnership(&_MarketImpl.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MarketImpl *MarketImplTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MarketImpl.Contract.TransferOwnership(&_MarketImpl.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_MarketImpl *MarketImplTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _MarketImpl.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_MarketImpl *MarketImplSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _MarketImpl.Contract.UpgradeToAndCall(&_MarketImpl.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_MarketImpl *MarketImplTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _MarketImpl.Contract.UpgradeToAndCall(&_MarketImpl.TransactOpts, newImplementation, data)
}

// MarketImplInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the MarketImpl contract.
type MarketImplInitializedIterator struct {
	Event *MarketImplInitialized // Event containing the contract specifics and raw log

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
func (it *MarketImplInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketImplInitialized)
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
		it.Event = new(MarketImplInitialized)
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
func (it *MarketImplInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketImplInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketImplInitialized represents a Initialized event raised by the MarketImpl contract.
type MarketImplInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MarketImpl *MarketImplFilterer) FilterInitialized(opts *bind.FilterOpts) (*MarketImplInitializedIterator, error) {

	logs, sub, err := _MarketImpl.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MarketImplInitializedIterator{contract: _MarketImpl.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MarketImpl *MarketImplFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MarketImplInitialized) (event.Subscription, error) {

	logs, sub, err := _MarketImpl.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketImplInitialized)
				if err := _MarketImpl.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_MarketImpl *MarketImplFilterer) ParseInitialized(log types.Log) (*MarketImplInitialized, error) {
	event := new(MarketImplInitialized)
	if err := _MarketImpl.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketImplOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MarketImpl contract.
type MarketImplOwnershipTransferredIterator struct {
	Event *MarketImplOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MarketImplOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketImplOwnershipTransferred)
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
		it.Event = new(MarketImplOwnershipTransferred)
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
func (it *MarketImplOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketImplOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketImplOwnershipTransferred represents a OwnershipTransferred event raised by the MarketImpl contract.
type MarketImplOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MarketImpl *MarketImplFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MarketImplOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MarketImpl.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MarketImplOwnershipTransferredIterator{contract: _MarketImpl.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MarketImpl *MarketImplFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MarketImplOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MarketImpl.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketImplOwnershipTransferred)
				if err := _MarketImpl.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MarketImpl *MarketImplFilterer) ParseOwnershipTransferred(log types.Log) (*MarketImplOwnershipTransferred, error) {
	event := new(MarketImplOwnershipTransferred)
	if err := _MarketImpl.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketImplPairCreatedIterator is returned from FilterPairCreated and is used to iterate over the raw logs and unpacked data for PairCreated events raised by the MarketImpl contract.
type MarketImplPairCreatedIterator struct {
	Event *MarketImplPairCreated // Event containing the contract specifics and raw log

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
func (it *MarketImplPairCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketImplPairCreated)
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
		it.Event = new(MarketImplPairCreated)
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
func (it *MarketImplPairCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketImplPairCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketImplPairCreated represents a PairCreated event raised by the MarketImpl contract.
type MarketImplPairCreated struct {
	Pair      common.Address
	Base      common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPairCreated is a free log retrieval operation binding the contract event 0xc1db9ba7c4b7ce660fe8d17bbcf07167549381df2abd694a970bd1402d86d313.
//
// Solidity: event PairCreated(address indexed pair, address indexed base, uint256 timestamp)
func (_MarketImpl *MarketImplFilterer) FilterPairCreated(opts *bind.FilterOpts, pair []common.Address, base []common.Address) (*MarketImplPairCreatedIterator, error) {

	var pairRule []interface{}
	for _, pairItem := range pair {
		pairRule = append(pairRule, pairItem)
	}
	var baseRule []interface{}
	for _, baseItem := range base {
		baseRule = append(baseRule, baseItem)
	}

	logs, sub, err := _MarketImpl.contract.FilterLogs(opts, "PairCreated", pairRule, baseRule)
	if err != nil {
		return nil, err
	}
	return &MarketImplPairCreatedIterator{contract: _MarketImpl.contract, event: "PairCreated", logs: logs, sub: sub}, nil
}

// WatchPairCreated is a free log subscription operation binding the contract event 0xc1db9ba7c4b7ce660fe8d17bbcf07167549381df2abd694a970bd1402d86d313.
//
// Solidity: event PairCreated(address indexed pair, address indexed base, uint256 timestamp)
func (_MarketImpl *MarketImplFilterer) WatchPairCreated(opts *bind.WatchOpts, sink chan<- *MarketImplPairCreated, pair []common.Address, base []common.Address) (event.Subscription, error) {

	var pairRule []interface{}
	for _, pairItem := range pair {
		pairRule = append(pairRule, pairItem)
	}
	var baseRule []interface{}
	for _, baseItem := range base {
		baseRule = append(baseRule, baseItem)
	}

	logs, sub, err := _MarketImpl.contract.WatchLogs(opts, "PairCreated", pairRule, baseRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketImplPairCreated)
				if err := _MarketImpl.contract.UnpackLog(event, "PairCreated", log); err != nil {
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

// ParsePairCreated is a log parse operation binding the contract event 0xc1db9ba7c4b7ce660fe8d17bbcf07167549381df2abd694a970bd1402d86d313.
//
// Solidity: event PairCreated(address indexed pair, address indexed base, uint256 timestamp)
func (_MarketImpl *MarketImplFilterer) ParsePairCreated(log types.Log) (*MarketImplPairCreated, error) {
	event := new(MarketImplPairCreated)
	if err := _MarketImpl.contract.UnpackLog(event, "PairCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketImplUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the MarketImpl contract.
type MarketImplUpgradedIterator struct {
	Event *MarketImplUpgraded // Event containing the contract specifics and raw log

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
func (it *MarketImplUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketImplUpgraded)
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
		it.Event = new(MarketImplUpgraded)
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
func (it *MarketImplUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketImplUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketImplUpgraded represents a Upgraded event raised by the MarketImpl contract.
type MarketImplUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_MarketImpl *MarketImplFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*MarketImplUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _MarketImpl.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &MarketImplUpgradedIterator{contract: _MarketImpl.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_MarketImpl *MarketImplFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *MarketImplUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _MarketImpl.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketImplUpgraded)
				if err := _MarketImpl.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_MarketImpl *MarketImplFilterer) ParseUpgraded(log types.Log) (*MarketImplUpgraded, error) {
	event := new(MarketImplUpgraded)
	if err := _MarketImpl.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
