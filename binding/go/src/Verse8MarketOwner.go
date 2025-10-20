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

// Verse8MarketOwnerMetaData contains all meta data concerning the Verse8MarketOwner contract.
var Verse8MarketOwnerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"creators\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"beginDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"}],\"name\":\"changeDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tickSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lotSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"name\":\"createPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tickSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lotSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structVerse8MarketOwner.CreatePairArgs[]\",\"name\":\"args\",\"type\":\"tuple[]\"}],\"name\":\"createPairs\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelayIncreaseWait\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structVerse8MarketOwner.ExecuteBatchArgs[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"executeBatch\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollbackDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminDelayChangeCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"effectSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminDelayChangeScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminTransferCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"acceptSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminTransferScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"name\":\"AccessControlEnforcedDefaultAdminDelay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccessControlEnforcedDefaultAdminRules\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"defaultAdmin\",\"type\":\"address\"}],\"name\":\"AccessControlInvalidDefaultAdmin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Verse8MarketOwner__CallFailed\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"a217fddf": "DEFAULT_ADMIN_ROLE()",
		"cefc1429": "acceptDefaultAdminTransfer()",
		"634e93da": "beginDefaultAdminTransfer(address)",
		"d602b9fd": "cancelDefaultAdminTransfer()",
		"649a5ec7": "changeDefaultAdminDelay(uint48)",
		"8b481bbd": "createPair(address,address,uint256,uint256,bytes)",
		"80d3ff51": "createPairs((address,address,uint256,uint256,bytes)[])",
		"84ef8ffc": "defaultAdmin()",
		"cc8463c8": "defaultAdminDelay()",
		"022d63fb": "defaultAdminDelayIncreaseWait()",
		"b61d27f6": "execute(address,uint256,bytes)",
		"34fcd5be": "executeBatch((address,uint256,bytes)[])",
		"248a9ca3": "getRoleAdmin(bytes32)",
		"2f2ff15d": "grantRole(bytes32,address)",
		"91d14854": "hasRole(bytes32,address)",
		"8da5cb5b": "owner()",
		"cf6eefb7": "pendingDefaultAdmin()",
		"a1eda53c": "pendingDefaultAdminDelay()",
		"36568abe": "renounceRole(bytes32,address)",
		"d547741f": "revokeRole(bytes32,address)",
		"0aa6220b": "rollbackDefaultAdminDelay()",
		"01ffc9a7": "supportsInterface(bytes4)",
	},
	Bin: "0x608060405234801561000f575f5ffd5b5060405161203c38038061203c83398101604081905261002e9161023e565b5f826001600160a01b03811661005d57604051636116401160e11b81525f600482015260240160405180910390fd5b600180546001600160d01b0316600160d01b65ffffffffffff8516021790556100865f826100f9565b5050506100a65f51602061201c5f395f51905f52836100f960201b60201c565b505f5b81518110156100f1576100e85f51602061201c5f395f51905f528383815181106100d5576100d561031c565b60200260200101516100f960201b60201c565b506001016100a9565b505050610330565b5f82610155575f6101126002546001600160a01b031690565b6001600160a01b03161461013957604051631fe1e13d60e11b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b0384161790555b61015f8383610168565b90505b92915050565b5f828152602081815260408083206001600160a01b038516845290915281205460ff16610208575f838152602081815260408083206001600160a01b03861684529091529020805460ff191660011790556101c03390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4506001610162565b505f610162565b80516001600160a01b0381168114610225575f5ffd5b919050565b634e487b7160e01b5f52604160045260245ffd5b5f5f6040838503121561024f575f5ffd5b6102588361020f565b60208401519092506001600160401b03811115610273575f5ffd5b8301601f81018513610283575f5ffd5b80516001600160401b0381111561029c5761029c61022a565b604051600582901b90603f8201601f191681016001600160401b03811182821017156102ca576102ca61022a565b6040529182526020818401810192908101888411156102e7575f5ffd5b6020850194505b8385101561030d576102ff8561020f565b8152602094850194016102ee565b50809450505050509250929050565b634e487b7160e01b5f52603260045260245ffd5b611cdf8061033d5f395ff3fe608060405234801561000f575f5ffd5b5060043610610179575f3560e01c80638b481bbd116100d2578063b61d27f611610088578063cf6eefb711610063578063cf6eefb714610382578063d547741f146103ce578063d602b9fd146103e1575f5ffd5b8063b61d27f614610352578063cc8463c814610372578063cefc14291461037a575f5ffd5b806391d14854116100b857806391d14854146102e1578063a1eda53c14610324578063a217fddf1461034b575f5ffd5b80638b481bbd146102c65780638da5cb5b146102d9575f5ffd5b806334fcd5be11610132578063649a5ec71161010d578063649a5ec71461025457806380d3ff511461026757806384ef8ffc14610287575f5ffd5b806334fcd5be1461020e57806336568abe1461022e578063634e93da14610241575f5ffd5b80630aa6220b116101625780630aa6220b146101c1578063248a9ca3146101cb5780632f2ff15d146101fb575f5ffd5b806301ffc9a71461017d578063022d63fb146101a5575b5f5ffd5b61019061018b3660046115cc565b6103e9565b60405190151581526020015b60405180910390f35b620697805b60405165ffffffffffff909116815260200161019c565b6101c9610444565b005b6101ed6101d936600461160b565b5f9081526020819052604090206001015490565b60405190815260200161019c565b6101c9610209366004611643565b610459565b61022161021c3660046116b9565b61049e565b60405161019c9190611744565b6101c961023c366004611643565b610660565b6101c961024f3660046117c5565b61076a565b6101c96102623660046117e0565b61077d565b61027a6102753660046116b9565b610790565b60405161019c9190611805565b60025473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161019c565b6102a16102d4366004611933565b610930565b6102a16109fe565b6101906102ef366004611643565b5f9182526020828152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b61032c610a23565b6040805165ffffffffffff93841681529290911660208301520161019c565b6101ed5f81565b6103656103603660046119a4565b610a9d565b60405161019c9190611a29565b6101aa610b4e565b6101c9610beb565b6001546040805173ffffffffffffffffffffffffffffffffffffffff831681527401000000000000000000000000000000000000000090920465ffffffffffff1660208301520161019c565b6101c96103dc366004611643565b610c47565b6101c9610c88565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f3149878600000000000000000000000000000000000000000000000000000000148061043e575061043e82610c9a565b92915050565b5f61044e81610d30565b610456610d3a565b50565b81610490576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61049a8282610d46565b5050565b60605f6104aa81610d30565b8267ffffffffffffffff8111156104c3576104c361185d565b6040519080825280602002602001820160405280156104f657816020015b60608152602001906001900390816104e15790505b5091505f5b83811015610658575f5f86868481811061051757610517611a3b565b90506020028101906105299190611a68565b6105379060208101906117c5565b73ffffffffffffffffffffffffffffffffffffffff1687878581811061055f5761055f611a3b565b90506020028101906105719190611a68565b6020013588888681811061058757610587611a3b565b90506020028101906105999190611a68565b6105a7906040810190611aa4565b6040516105b5929190611b05565b5f6040518083038185875af1925050503d805f81146105ef576040519150601f19603f3d011682016040523d82523d5f602084013e6105f4565b606091505b509150915081610630576040517fbc5c9d6b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8085848151811061064357610643611a3b565b602090810291909101015250506001016104fb565b505092915050565b81158015610688575060025473ffffffffffffffffffffffffffffffffffffffff8281169116145b156107605760015473ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900465ffffffffffff16811515806106dc575065ffffffffffff8116155b806106ef57504265ffffffffffff821610155b15610735576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff821660048201526024015b60405180910390fd5b5050600180547fffffffffffff000000000000ffffffffffffffffffffffffffffffffffffffff1690555b61049a8282610d70565b5f61077481610d30565b61049a82610dce565b5f61078781610d30565b61049a82610e4d565b60607f02d639b3242e624c4062ce3346179a769447ef0a01fb09608afc904d0268f1906107bc81610d30565b5f8367ffffffffffffffff8111156107d6576107d661185d565b6040519080825280602002602001820160405280156107ff578160200160208202803683370190505b5090505f5b84811015610927575f86868381811061081f5761081f611a3b565b90506020028101906108319190611b14565b61083a90611b46565b805160208201516040808401516060850151608086015192517f4732c7e500000000000000000000000000000000000000000000000000000000815295965073ffffffffffffffffffffffffffffffffffffffff90941694634732c7e5946108a794939091600401611be1565b6020604051808303815f875af11580156108c3573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108e79190611c25565b8383815181106108f9576108f9611a3b565b73ffffffffffffffffffffffffffffffffffffffff9092166020928302919091019091015250600101610804565b50949350505050565b5f7f02d639b3242e624c4062ce3346179a769447ef0a01fb09608afc904d0268f19061095b81610d30565b6040517f4732c7e500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff881690634732c7e5906109b3908990899089908990600401611be1565b6020604051808303815f875af11580156109cf573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109f39190611c25565b979650505050505050565b5f610a1e60025473ffffffffffffffffffffffffffffffffffffffff1690565b905090565b6002545f907a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015158015610a6557504265ffffffffffff821610155b610a70575f5f610a95565b60025474010000000000000000000000000000000000000000900465ffffffffffff16815b915091509091565b60605f610aa981610d30565b5f5f8773ffffffffffffffffffffffffffffffffffffffff16878787604051610ad3929190611b05565b5f6040518083038185875af1925050503d805f8114610b0d576040519150601f19603f3d011682016040523d82523d5f602084013e610b12565b606091505b5091509150816109f3576040517fbc5c9d6b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002545f907a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015158015610b8f57504265ffffffffffff8216105b610bc1576001547a010000000000000000000000000000000000000000000000000000900465ffffffffffff16610be5565b60025474010000000000000000000000000000000000000000900465ffffffffffff165b91505090565b60015473ffffffffffffffffffffffffffffffffffffffff16338114610c3f576040517fc22c802200000000000000000000000000000000000000000000000000000000815233600482015260240161072c565b610456610ebc565b81610c7e576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61049a8282610fad565b5f610c9281610d30565b610456610fd1565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061043e57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000083161461043e565b6104568133610fdb565b610d445f5f611060565b565b5f82815260208190526040902060010154610d6081610d30565b610d6a83836111b9565b50505050565b73ffffffffffffffffffffffffffffffffffffffff81163314610dbf576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610dc9828261127e565b505050565b5f610dd7610b4e565b610de0426112df565b610dea9190611c6d565b9050610df6828261132e565b60405165ffffffffffff8216815273ffffffffffffffffffffffffffffffffffffffff8316907f3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed69060200160405180910390a25050565b5f610e57826113c9565b610e60426112df565b610e6a9190611c6d565b9050610e768282611060565b6040805165ffffffffffff8085168252831660208201527ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b910160405180910390a15050565b60015473ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900465ffffffffffff16801580610f0c57504265ffffffffffff821610155b15610f4d576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff8216600482015260240161072c565b610f755f610f7060025473ffffffffffffffffffffffffffffffffffffffff1690565b61127e565b50610f805f836111b9565b5050600180547fffffffffffff000000000000000000000000000000000000000000000000000016905550565b5f82815260208190526040902060010154610fc781610d30565b610d6a838361127e565b610d445f5f61132e565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff1661049a576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024810183905260440161072c565b6002547a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015611134574265ffffffffffff8216101561110b576002546001805479ffffffffffffffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000090920465ffffffffffff167a01000000000000000000000000000000000000000000000000000002919091179055611134565b6040517f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5905f90a15b506002805473ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000065ffffffffffff9485160279ffffffffffffffffffffffffffffffffffffffffffffffffffff16177a0100000000000000000000000000000000000000000000000000009290931691909102919091179055565b5f8261126d575f6111df60025473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff161461122c576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84161790555b611277838361141a565b9392505050565b5f821580156112a7575060025473ffffffffffffffffffffffffffffffffffffffff8381169116145b156112d557600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001690555b6112778383611513565b5f65ffffffffffff82111561132a576040517f6dfcc650000000000000000000000000000000000000000000000000000000008152603060048201526024810183905260440161072c565b5090565b600180547401000000000000000000000000000000000000000065ffffffffffff84811682027fffffffffffff0000000000000000000000000000000000000000000000000000841673ffffffffffffffffffffffffffffffffffffffff881617179093559004168015610dc9576040517f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109905f90a1505050565b5f5f6113d3610b4e565b90508065ffffffffffff168365ffffffffffff16116113fb576113f68382611c8b565b611277565b61127765ffffffffffff8416620697805f828218828410028218611277565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff1661150c575f8381526020818152604080832073ffffffffffffffffffffffffffffffffffffffff86168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556114aa3390565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a450600161043e565b505f61043e565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff161561150c575f8381526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8616808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a450600161043e565b5f602082840312156115dc575f5ffd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114611277575f5ffd5b5f6020828403121561161b575f5ffd5b5035919050565b73ffffffffffffffffffffffffffffffffffffffff81168114610456575f5ffd5b5f5f60408385031215611654575f5ffd5b82359150602083013561166681611622565b809150509250929050565b5f5f83601f840112611681575f5ffd5b50813567ffffffffffffffff811115611698575f5ffd5b6020830191508360208260051b85010111156116b2575f5ffd5b9250929050565b5f5f602083850312156116ca575f5ffd5b823567ffffffffffffffff8111156116e0575f5ffd5b6116ec85828601611671565b90969095509350505050565b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b828110156117b9577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08786030184526117a48583516116f8565b9450602093840193919091019060010161176a565b50929695505050505050565b5f602082840312156117d5575f5ffd5b813561127781611622565b5f602082840312156117f0575f5ffd5b813565ffffffffffff81168114611277575f5ffd5b602080825282518282018190525f918401906040840190835b8181101561185257835173ffffffffffffffffffffffffffffffffffffffff1683526020938401939092019160010161181e565b509095945050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f82601f830112611899575f5ffd5b813567ffffffffffffffff8111156118b3576118b361185d565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810167ffffffffffffffff811182821017156119005761190061185d565b604052818152838201602001851015611917575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f5f5f5f5f60a08688031215611947575f5ffd5b853561195281611622565b9450602086013561196281611622565b93506040860135925060608601359150608086013567ffffffffffffffff81111561198b575f5ffd5b6119978882890161188a565b9150509295509295909350565b5f5f5f5f606085870312156119b7575f5ffd5b84356119c281611622565b935060208501359250604085013567ffffffffffffffff8111156119e4575f5ffd5b8501601f810187136119f4575f5ffd5b803567ffffffffffffffff811115611a0a575f5ffd5b876020828401011115611a1b575f5ffd5b949793965060200194505050565b602081525f61127760208301846116f8565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f82357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa1833603018112611a9a575f5ffd5b9190910192915050565b5f5f83357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112611ad7575f5ffd5b83018035915067ffffffffffffffff821115611af1575f5ffd5b6020019150368190038213156116b2575f5ffd5b818382375f9101908152919050565b5f82357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff61833603018112611a9a575f5ffd5b5f60a08236031215611b56575f5ffd5b60405160a0810167ffffffffffffffff81118282101715611b7957611b7961185d565b6040528235611b8781611622565b81526020830135611b9781611622565b60208201526040838101359082015260608084013590820152608083013567ffffffffffffffff811115611bc9575f5ffd5b611bd53682860161188a565b60808301525092915050565b73ffffffffffffffffffffffffffffffffffffffff85168152836020820152826040820152608060608201525f611c1b60808301846116f8565b9695505050505050565b5f60208284031215611c35575f5ffd5b815161127781611622565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b65ffffffffffff818116838216019081111561043e5761043e611c40565b65ffffffffffff828116828216039081111561043e5761043e611c4056fea264697066735822122019dfb29c2cf8b3e0af0133fb4ad0ae9c32950d18fd9c7721056551933463f9fd64736f6c634300081c003302d639b3242e624c4062ce3346179a769447ef0a01fb09608afc904d0268f190",
}

// Verse8MarketOwnerABI is the input ABI used to generate the binding from.
// Deprecated: Use Verse8MarketOwnerMetaData.ABI instead.
var Verse8MarketOwnerABI = Verse8MarketOwnerMetaData.ABI

// Deprecated: Use Verse8MarketOwnerMetaData.Sigs instead.
// Verse8MarketOwnerFuncSigs maps the 4-byte function signature to its string representation.
var Verse8MarketOwnerFuncSigs = Verse8MarketOwnerMetaData.Sigs

// Verse8MarketOwnerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Verse8MarketOwnerMetaData.Bin instead.
var Verse8MarketOwnerBin = Verse8MarketOwnerMetaData.Bin

// DeployVerse8MarketOwner deploys a new Ethereum contract, binding an instance of Verse8MarketOwner to it.
func DeployVerse8MarketOwner(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address, creators []common.Address) (common.Address, *types.Transaction, *Verse8MarketOwner, error) {
	parsed, err := Verse8MarketOwnerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Verse8MarketOwnerBin), backend, _owner, creators)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Verse8MarketOwner{Verse8MarketOwnerCaller: Verse8MarketOwnerCaller{contract: contract}, Verse8MarketOwnerTransactor: Verse8MarketOwnerTransactor{contract: contract}, Verse8MarketOwnerFilterer: Verse8MarketOwnerFilterer{contract: contract}}, nil
}

// Verse8MarketOwner is an auto generated Go binding around an Ethereum contract.
type Verse8MarketOwner struct {
	Verse8MarketOwnerCaller     // Read-only binding to the contract
	Verse8MarketOwnerTransactor // Write-only binding to the contract
	Verse8MarketOwnerFilterer   // Log filterer for contract events
}

// Verse8MarketOwnerCaller is an auto generated read-only Go binding around an Ethereum contract.
type Verse8MarketOwnerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Verse8MarketOwnerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Verse8MarketOwnerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Verse8MarketOwnerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Verse8MarketOwnerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Verse8MarketOwnerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Verse8MarketOwnerSession struct {
	Contract     *Verse8MarketOwner // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// Verse8MarketOwnerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Verse8MarketOwnerCallerSession struct {
	Contract *Verse8MarketOwnerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// Verse8MarketOwnerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Verse8MarketOwnerTransactorSession struct {
	Contract     *Verse8MarketOwnerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// Verse8MarketOwnerRaw is an auto generated low-level Go binding around an Ethereum contract.
type Verse8MarketOwnerRaw struct {
	Contract *Verse8MarketOwner // Generic contract binding to access the raw methods on
}

// Verse8MarketOwnerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Verse8MarketOwnerCallerRaw struct {
	Contract *Verse8MarketOwnerCaller // Generic read-only contract binding to access the raw methods on
}

// Verse8MarketOwnerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Verse8MarketOwnerTransactorRaw struct {
	Contract *Verse8MarketOwnerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVerse8MarketOwner creates a new instance of Verse8MarketOwner, bound to a specific deployed contract.
func NewVerse8MarketOwner(address common.Address, backend bind.ContractBackend) (*Verse8MarketOwner, error) {
	contract, err := bindVerse8MarketOwner(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Verse8MarketOwner{Verse8MarketOwnerCaller: Verse8MarketOwnerCaller{contract: contract}, Verse8MarketOwnerTransactor: Verse8MarketOwnerTransactor{contract: contract}, Verse8MarketOwnerFilterer: Verse8MarketOwnerFilterer{contract: contract}}, nil
}

// NewVerse8MarketOwnerCaller creates a new read-only instance of Verse8MarketOwner, bound to a specific deployed contract.
func NewVerse8MarketOwnerCaller(address common.Address, caller bind.ContractCaller) (*Verse8MarketOwnerCaller, error) {
	contract, err := bindVerse8MarketOwner(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Verse8MarketOwnerCaller{contract: contract}, nil
}

// NewVerse8MarketOwnerTransactor creates a new write-only instance of Verse8MarketOwner, bound to a specific deployed contract.
func NewVerse8MarketOwnerTransactor(address common.Address, transactor bind.ContractTransactor) (*Verse8MarketOwnerTransactor, error) {
	contract, err := bindVerse8MarketOwner(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Verse8MarketOwnerTransactor{contract: contract}, nil
}

// NewVerse8MarketOwnerFilterer creates a new log filterer instance of Verse8MarketOwner, bound to a specific deployed contract.
func NewVerse8MarketOwnerFilterer(address common.Address, filterer bind.ContractFilterer) (*Verse8MarketOwnerFilterer, error) {
	contract, err := bindVerse8MarketOwner(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Verse8MarketOwnerFilterer{contract: contract}, nil
}

// bindVerse8MarketOwner binds a generic wrapper to an already deployed contract.
func bindVerse8MarketOwner(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Verse8MarketOwnerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verse8MarketOwner *Verse8MarketOwnerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Verse8MarketOwner.Contract.Verse8MarketOwnerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verse8MarketOwner *Verse8MarketOwnerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verse8MarketOwner.Contract.Verse8MarketOwnerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verse8MarketOwner *Verse8MarketOwnerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verse8MarketOwner.Contract.Verse8MarketOwnerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verse8MarketOwner *Verse8MarketOwnerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Verse8MarketOwner.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verse8MarketOwner *Verse8MarketOwnerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verse8MarketOwner.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verse8MarketOwner *Verse8MarketOwnerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verse8MarketOwner.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Verse8MarketOwner *Verse8MarketOwnerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Verse8MarketOwner.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Verse8MarketOwner *Verse8MarketOwnerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Verse8MarketOwner.Contract.DEFAULTADMINROLE(&_Verse8MarketOwner.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Verse8MarketOwner *Verse8MarketOwnerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Verse8MarketOwner.Contract.DEFAULTADMINROLE(&_Verse8MarketOwner.CallOpts)
}

// DefaultAdmin is a free data retrieval call binding the contract method 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (_Verse8MarketOwner *Verse8MarketOwnerCaller) DefaultAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Verse8MarketOwner.contract.Call(opts, &out, "defaultAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultAdmin is a free data retrieval call binding the contract method 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (_Verse8MarketOwner *Verse8MarketOwnerSession) DefaultAdmin() (common.Address, error) {
	return _Verse8MarketOwner.Contract.DefaultAdmin(&_Verse8MarketOwner.CallOpts)
}

// DefaultAdmin is a free data retrieval call binding the contract method 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (_Verse8MarketOwner *Verse8MarketOwnerCallerSession) DefaultAdmin() (common.Address, error) {
	return _Verse8MarketOwner.Contract.DefaultAdmin(&_Verse8MarketOwner.CallOpts)
}

// DefaultAdminDelay is a free data retrieval call binding the contract method 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (_Verse8MarketOwner *Verse8MarketOwnerCaller) DefaultAdminDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Verse8MarketOwner.contract.Call(opts, &out, "defaultAdminDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultAdminDelay is a free data retrieval call binding the contract method 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (_Verse8MarketOwner *Verse8MarketOwnerSession) DefaultAdminDelay() (*big.Int, error) {
	return _Verse8MarketOwner.Contract.DefaultAdminDelay(&_Verse8MarketOwner.CallOpts)
}

// DefaultAdminDelay is a free data retrieval call binding the contract method 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (_Verse8MarketOwner *Verse8MarketOwnerCallerSession) DefaultAdminDelay() (*big.Int, error) {
	return _Verse8MarketOwner.Contract.DefaultAdminDelay(&_Verse8MarketOwner.CallOpts)
}

// DefaultAdminDelayIncreaseWait is a free data retrieval call binding the contract method 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (_Verse8MarketOwner *Verse8MarketOwnerCaller) DefaultAdminDelayIncreaseWait(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Verse8MarketOwner.contract.Call(opts, &out, "defaultAdminDelayIncreaseWait")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultAdminDelayIncreaseWait is a free data retrieval call binding the contract method 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (_Verse8MarketOwner *Verse8MarketOwnerSession) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _Verse8MarketOwner.Contract.DefaultAdminDelayIncreaseWait(&_Verse8MarketOwner.CallOpts)
}

// DefaultAdminDelayIncreaseWait is a free data retrieval call binding the contract method 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (_Verse8MarketOwner *Verse8MarketOwnerCallerSession) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _Verse8MarketOwner.Contract.DefaultAdminDelayIncreaseWait(&_Verse8MarketOwner.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Verse8MarketOwner *Verse8MarketOwnerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Verse8MarketOwner.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Verse8MarketOwner *Verse8MarketOwnerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Verse8MarketOwner.Contract.GetRoleAdmin(&_Verse8MarketOwner.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Verse8MarketOwner *Verse8MarketOwnerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Verse8MarketOwner.Contract.GetRoleAdmin(&_Verse8MarketOwner.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Verse8MarketOwner *Verse8MarketOwnerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Verse8MarketOwner.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Verse8MarketOwner *Verse8MarketOwnerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Verse8MarketOwner.Contract.HasRole(&_Verse8MarketOwner.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Verse8MarketOwner *Verse8MarketOwnerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Verse8MarketOwner.Contract.HasRole(&_Verse8MarketOwner.CallOpts, role, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Verse8MarketOwner *Verse8MarketOwnerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Verse8MarketOwner.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Verse8MarketOwner *Verse8MarketOwnerSession) Owner() (common.Address, error) {
	return _Verse8MarketOwner.Contract.Owner(&_Verse8MarketOwner.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Verse8MarketOwner *Verse8MarketOwnerCallerSession) Owner() (common.Address, error) {
	return _Verse8MarketOwner.Contract.Owner(&_Verse8MarketOwner.CallOpts)
}

// PendingDefaultAdmin is a free data retrieval call binding the contract method 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (_Verse8MarketOwner *Verse8MarketOwnerCaller) PendingDefaultAdmin(opts *bind.CallOpts) (struct {
	NewAdmin common.Address
	Schedule *big.Int
}, error) {
	var out []interface{}
	err := _Verse8MarketOwner.contract.Call(opts, &out, "pendingDefaultAdmin")

	outstruct := new(struct {
		NewAdmin common.Address
		Schedule *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NewAdmin = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Schedule = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PendingDefaultAdmin is a free data retrieval call binding the contract method 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (_Verse8MarketOwner *Verse8MarketOwnerSession) PendingDefaultAdmin() (struct {
	NewAdmin common.Address
	Schedule *big.Int
}, error) {
	return _Verse8MarketOwner.Contract.PendingDefaultAdmin(&_Verse8MarketOwner.CallOpts)
}

// PendingDefaultAdmin is a free data retrieval call binding the contract method 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (_Verse8MarketOwner *Verse8MarketOwnerCallerSession) PendingDefaultAdmin() (struct {
	NewAdmin common.Address
	Schedule *big.Int
}, error) {
	return _Verse8MarketOwner.Contract.PendingDefaultAdmin(&_Verse8MarketOwner.CallOpts)
}

// PendingDefaultAdminDelay is a free data retrieval call binding the contract method 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (_Verse8MarketOwner *Verse8MarketOwnerCaller) PendingDefaultAdminDelay(opts *bind.CallOpts) (struct {
	NewDelay *big.Int
	Schedule *big.Int
}, error) {
	var out []interface{}
	err := _Verse8MarketOwner.contract.Call(opts, &out, "pendingDefaultAdminDelay")

	outstruct := new(struct {
		NewDelay *big.Int
		Schedule *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NewDelay = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Schedule = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PendingDefaultAdminDelay is a free data retrieval call binding the contract method 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (_Verse8MarketOwner *Verse8MarketOwnerSession) PendingDefaultAdminDelay() (struct {
	NewDelay *big.Int
	Schedule *big.Int
}, error) {
	return _Verse8MarketOwner.Contract.PendingDefaultAdminDelay(&_Verse8MarketOwner.CallOpts)
}

// PendingDefaultAdminDelay is a free data retrieval call binding the contract method 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (_Verse8MarketOwner *Verse8MarketOwnerCallerSession) PendingDefaultAdminDelay() (struct {
	NewDelay *big.Int
	Schedule *big.Int
}, error) {
	return _Verse8MarketOwner.Contract.PendingDefaultAdminDelay(&_Verse8MarketOwner.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Verse8MarketOwner *Verse8MarketOwnerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Verse8MarketOwner.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Verse8MarketOwner *Verse8MarketOwnerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Verse8MarketOwner.Contract.SupportsInterface(&_Verse8MarketOwner.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Verse8MarketOwner *Verse8MarketOwnerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Verse8MarketOwner.Contract.SupportsInterface(&_Verse8MarketOwner.CallOpts, interfaceId)
}

// AcceptDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xcefc1429.
//
// Solidity: function acceptDefaultAdminTransfer() returns()
func (_Verse8MarketOwner *Verse8MarketOwnerTransactor) AcceptDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verse8MarketOwner.contract.Transact(opts, "acceptDefaultAdminTransfer")
}

// AcceptDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xcefc1429.
//
// Solidity: function acceptDefaultAdminTransfer() returns()
func (_Verse8MarketOwner *Verse8MarketOwnerSession) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _Verse8MarketOwner.Contract.AcceptDefaultAdminTransfer(&_Verse8MarketOwner.TransactOpts)
}

// AcceptDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xcefc1429.
//
// Solidity: function acceptDefaultAdminTransfer() returns()
func (_Verse8MarketOwner *Verse8MarketOwnerTransactorSession) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _Verse8MarketOwner.Contract.AcceptDefaultAdminTransfer(&_Verse8MarketOwner.TransactOpts)
}

// BeginDefaultAdminTransfer is a paid mutator transaction binding the contract method 0x634e93da.
//
// Solidity: function beginDefaultAdminTransfer(address newAdmin) returns()
func (_Verse8MarketOwner *Verse8MarketOwnerTransactor) BeginDefaultAdminTransfer(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _Verse8MarketOwner.contract.Transact(opts, "beginDefaultAdminTransfer", newAdmin)
}

// BeginDefaultAdminTransfer is a paid mutator transaction binding the contract method 0x634e93da.
//
// Solidity: function beginDefaultAdminTransfer(address newAdmin) returns()
func (_Verse8MarketOwner *Verse8MarketOwnerSession) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _Verse8MarketOwner.Contract.BeginDefaultAdminTransfer(&_Verse8MarketOwner.TransactOpts, newAdmin)
}

// BeginDefaultAdminTransfer is a paid mutator transaction binding the contract method 0x634e93da.
//
// Solidity: function beginDefaultAdminTransfer(address newAdmin) returns()
func (_Verse8MarketOwner *Verse8MarketOwnerTransactorSession) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _Verse8MarketOwner.Contract.BeginDefaultAdminTransfer(&_Verse8MarketOwner.TransactOpts, newAdmin)
}

// CancelDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (_Verse8MarketOwner *Verse8MarketOwnerTransactor) CancelDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verse8MarketOwner.contract.Transact(opts, "cancelDefaultAdminTransfer")
}

// CancelDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (_Verse8MarketOwner *Verse8MarketOwnerSession) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _Verse8MarketOwner.Contract.CancelDefaultAdminTransfer(&_Verse8MarketOwner.TransactOpts)
}

// CancelDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (_Verse8MarketOwner *Verse8MarketOwnerTransactorSession) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _Verse8MarketOwner.Contract.CancelDefaultAdminTransfer(&_Verse8MarketOwner.TransactOpts)
}

// ChangeDefaultAdminDelay is a paid mutator transaction binding the contract method 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (_Verse8MarketOwner *Verse8MarketOwnerTransactor) ChangeDefaultAdminDelay(opts *bind.TransactOpts, newDelay *big.Int) (*types.Transaction, error) {
	return _Verse8MarketOwner.contract.Transact(opts, "changeDefaultAdminDelay", newDelay)
}

// ChangeDefaultAdminDelay is a paid mutator transaction binding the contract method 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (_Verse8MarketOwner *Verse8MarketOwnerSession) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _Verse8MarketOwner.Contract.ChangeDefaultAdminDelay(&_Verse8MarketOwner.TransactOpts, newDelay)
}

// ChangeDefaultAdminDelay is a paid mutator transaction binding the contract method 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (_Verse8MarketOwner *Verse8MarketOwnerTransactorSession) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _Verse8MarketOwner.Contract.ChangeDefaultAdminDelay(&_Verse8MarketOwner.TransactOpts, newDelay)
}

// CreatePair is a paid mutator transaction binding the contract method 0x8b481bbd.
//
// Solidity: function createPair(address market, address base, uint256 tickSize, uint256 lotSize, bytes feeData) returns(address)
func (_Verse8MarketOwner *Verse8MarketOwnerTransactor) CreatePair(opts *bind.TransactOpts, market common.Address, base common.Address, tickSize *big.Int, lotSize *big.Int, feeData []byte) (*types.Transaction, error) {
	return _Verse8MarketOwner.contract.Transact(opts, "createPair", market, base, tickSize, lotSize, feeData)
}

// CreatePair is a paid mutator transaction binding the contract method 0x8b481bbd.
//
// Solidity: function createPair(address market, address base, uint256 tickSize, uint256 lotSize, bytes feeData) returns(address)
func (_Verse8MarketOwner *Verse8MarketOwnerSession) CreatePair(market common.Address, base common.Address, tickSize *big.Int, lotSize *big.Int, feeData []byte) (*types.Transaction, error) {
	return _Verse8MarketOwner.Contract.CreatePair(&_Verse8MarketOwner.TransactOpts, market, base, tickSize, lotSize, feeData)
}

// CreatePair is a paid mutator transaction binding the contract method 0x8b481bbd.
//
// Solidity: function createPair(address market, address base, uint256 tickSize, uint256 lotSize, bytes feeData) returns(address)
func (_Verse8MarketOwner *Verse8MarketOwnerTransactorSession) CreatePair(market common.Address, base common.Address, tickSize *big.Int, lotSize *big.Int, feeData []byte) (*types.Transaction, error) {
	return _Verse8MarketOwner.Contract.CreatePair(&_Verse8MarketOwner.TransactOpts, market, base, tickSize, lotSize, feeData)
}

// CreatePairs is a paid mutator transaction binding the contract method 0x80d3ff51.
//
// Solidity: function createPairs((address,address,uint256,uint256,bytes)[] args) returns(address[])
func (_Verse8MarketOwner *Verse8MarketOwnerTransactor) CreatePairs(opts *bind.TransactOpts, args []Verse8MarketOwnerCreatePairArgs) (*types.Transaction, error) {
	return _Verse8MarketOwner.contract.Transact(opts, "createPairs", args)
}

// CreatePairs is a paid mutator transaction binding the contract method 0x80d3ff51.
//
// Solidity: function createPairs((address,address,uint256,uint256,bytes)[] args) returns(address[])
func (_Verse8MarketOwner *Verse8MarketOwnerSession) CreatePairs(args []Verse8MarketOwnerCreatePairArgs) (*types.Transaction, error) {
	return _Verse8MarketOwner.Contract.CreatePairs(&_Verse8MarketOwner.TransactOpts, args)
}

// CreatePairs is a paid mutator transaction binding the contract method 0x80d3ff51.
//
// Solidity: function createPairs((address,address,uint256,uint256,bytes)[] args) returns(address[])
func (_Verse8MarketOwner *Verse8MarketOwnerTransactorSession) CreatePairs(args []Verse8MarketOwnerCreatePairArgs) (*types.Transaction, error) {
	return _Verse8MarketOwner.Contract.CreatePairs(&_Verse8MarketOwner.TransactOpts, args)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address to, uint256 value, bytes data) returns(bytes)
func (_Verse8MarketOwner *Verse8MarketOwnerTransactor) Execute(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Verse8MarketOwner.contract.Transact(opts, "execute", to, value, data)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address to, uint256 value, bytes data) returns(bytes)
func (_Verse8MarketOwner *Verse8MarketOwnerSession) Execute(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Verse8MarketOwner.Contract.Execute(&_Verse8MarketOwner.TransactOpts, to, value, data)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address to, uint256 value, bytes data) returns(bytes)
func (_Verse8MarketOwner *Verse8MarketOwnerTransactorSession) Execute(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Verse8MarketOwner.Contract.Execute(&_Verse8MarketOwner.TransactOpts, to, value, data)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x34fcd5be.
//
// Solidity: function executeBatch((address,uint256,bytes)[] calls) returns(bytes[] results)
func (_Verse8MarketOwner *Verse8MarketOwnerTransactor) ExecuteBatch(opts *bind.TransactOpts, calls []Verse8MarketOwnerExecuteBatchArgs) (*types.Transaction, error) {
	return _Verse8MarketOwner.contract.Transact(opts, "executeBatch", calls)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x34fcd5be.
//
// Solidity: function executeBatch((address,uint256,bytes)[] calls) returns(bytes[] results)
func (_Verse8MarketOwner *Verse8MarketOwnerSession) ExecuteBatch(calls []Verse8MarketOwnerExecuteBatchArgs) (*types.Transaction, error) {
	return _Verse8MarketOwner.Contract.ExecuteBatch(&_Verse8MarketOwner.TransactOpts, calls)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x34fcd5be.
//
// Solidity: function executeBatch((address,uint256,bytes)[] calls) returns(bytes[] results)
func (_Verse8MarketOwner *Verse8MarketOwnerTransactorSession) ExecuteBatch(calls []Verse8MarketOwnerExecuteBatchArgs) (*types.Transaction, error) {
	return _Verse8MarketOwner.Contract.ExecuteBatch(&_Verse8MarketOwner.TransactOpts, calls)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Verse8MarketOwner *Verse8MarketOwnerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Verse8MarketOwner.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Verse8MarketOwner *Verse8MarketOwnerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Verse8MarketOwner.Contract.GrantRole(&_Verse8MarketOwner.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Verse8MarketOwner *Verse8MarketOwnerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Verse8MarketOwner.Contract.GrantRole(&_Verse8MarketOwner.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Verse8MarketOwner *Verse8MarketOwnerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Verse8MarketOwner.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Verse8MarketOwner *Verse8MarketOwnerSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Verse8MarketOwner.Contract.RenounceRole(&_Verse8MarketOwner.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Verse8MarketOwner *Verse8MarketOwnerTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Verse8MarketOwner.Contract.RenounceRole(&_Verse8MarketOwner.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Verse8MarketOwner *Verse8MarketOwnerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Verse8MarketOwner.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Verse8MarketOwner *Verse8MarketOwnerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Verse8MarketOwner.Contract.RevokeRole(&_Verse8MarketOwner.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Verse8MarketOwner *Verse8MarketOwnerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Verse8MarketOwner.Contract.RevokeRole(&_Verse8MarketOwner.TransactOpts, role, account)
}

// RollbackDefaultAdminDelay is a paid mutator transaction binding the contract method 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (_Verse8MarketOwner *Verse8MarketOwnerTransactor) RollbackDefaultAdminDelay(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verse8MarketOwner.contract.Transact(opts, "rollbackDefaultAdminDelay")
}

// RollbackDefaultAdminDelay is a paid mutator transaction binding the contract method 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (_Verse8MarketOwner *Verse8MarketOwnerSession) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _Verse8MarketOwner.Contract.RollbackDefaultAdminDelay(&_Verse8MarketOwner.TransactOpts)
}

// RollbackDefaultAdminDelay is a paid mutator transaction binding the contract method 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (_Verse8MarketOwner *Verse8MarketOwnerTransactorSession) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _Verse8MarketOwner.Contract.RollbackDefaultAdminDelay(&_Verse8MarketOwner.TransactOpts)
}

// Verse8MarketOwnerDefaultAdminDelayChangeCanceledIterator is returned from FilterDefaultAdminDelayChangeCanceled and is used to iterate over the raw logs and unpacked data for DefaultAdminDelayChangeCanceled events raised by the Verse8MarketOwner contract.
type Verse8MarketOwnerDefaultAdminDelayChangeCanceledIterator struct {
	Event *Verse8MarketOwnerDefaultAdminDelayChangeCanceled // Event containing the contract specifics and raw log

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
func (it *Verse8MarketOwnerDefaultAdminDelayChangeCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Verse8MarketOwnerDefaultAdminDelayChangeCanceled)
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
		it.Event = new(Verse8MarketOwnerDefaultAdminDelayChangeCanceled)
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
func (it *Verse8MarketOwnerDefaultAdminDelayChangeCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Verse8MarketOwnerDefaultAdminDelayChangeCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Verse8MarketOwnerDefaultAdminDelayChangeCanceled represents a DefaultAdminDelayChangeCanceled event raised by the Verse8MarketOwner contract.
type Verse8MarketOwnerDefaultAdminDelayChangeCanceled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminDelayChangeCanceled is a free log retrieval operation binding the contract event 0x2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5.
//
// Solidity: event DefaultAdminDelayChangeCanceled()
func (_Verse8MarketOwner *Verse8MarketOwnerFilterer) FilterDefaultAdminDelayChangeCanceled(opts *bind.FilterOpts) (*Verse8MarketOwnerDefaultAdminDelayChangeCanceledIterator, error) {

	logs, sub, err := _Verse8MarketOwner.contract.FilterLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return &Verse8MarketOwnerDefaultAdminDelayChangeCanceledIterator{contract: _Verse8MarketOwner.contract, event: "DefaultAdminDelayChangeCanceled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminDelayChangeCanceled is a free log subscription operation binding the contract event 0x2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5.
//
// Solidity: event DefaultAdminDelayChangeCanceled()
func (_Verse8MarketOwner *Verse8MarketOwnerFilterer) WatchDefaultAdminDelayChangeCanceled(opts *bind.WatchOpts, sink chan<- *Verse8MarketOwnerDefaultAdminDelayChangeCanceled) (event.Subscription, error) {

	logs, sub, err := _Verse8MarketOwner.contract.WatchLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Verse8MarketOwnerDefaultAdminDelayChangeCanceled)
				if err := _Verse8MarketOwner.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
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

// ParseDefaultAdminDelayChangeCanceled is a log parse operation binding the contract event 0x2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5.
//
// Solidity: event DefaultAdminDelayChangeCanceled()
func (_Verse8MarketOwner *Verse8MarketOwnerFilterer) ParseDefaultAdminDelayChangeCanceled(log types.Log) (*Verse8MarketOwnerDefaultAdminDelayChangeCanceled, error) {
	event := new(Verse8MarketOwnerDefaultAdminDelayChangeCanceled)
	if err := _Verse8MarketOwner.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Verse8MarketOwnerDefaultAdminDelayChangeScheduledIterator is returned from FilterDefaultAdminDelayChangeScheduled and is used to iterate over the raw logs and unpacked data for DefaultAdminDelayChangeScheduled events raised by the Verse8MarketOwner contract.
type Verse8MarketOwnerDefaultAdminDelayChangeScheduledIterator struct {
	Event *Verse8MarketOwnerDefaultAdminDelayChangeScheduled // Event containing the contract specifics and raw log

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
func (it *Verse8MarketOwnerDefaultAdminDelayChangeScheduledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Verse8MarketOwnerDefaultAdminDelayChangeScheduled)
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
		it.Event = new(Verse8MarketOwnerDefaultAdminDelayChangeScheduled)
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
func (it *Verse8MarketOwnerDefaultAdminDelayChangeScheduledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Verse8MarketOwnerDefaultAdminDelayChangeScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Verse8MarketOwnerDefaultAdminDelayChangeScheduled represents a DefaultAdminDelayChangeScheduled event raised by the Verse8MarketOwner contract.
type Verse8MarketOwnerDefaultAdminDelayChangeScheduled struct {
	NewDelay       *big.Int
	EffectSchedule *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminDelayChangeScheduled is a free log retrieval operation binding the contract event 0xf1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b.
//
// Solidity: event DefaultAdminDelayChangeScheduled(uint48 newDelay, uint48 effectSchedule)
func (_Verse8MarketOwner *Verse8MarketOwnerFilterer) FilterDefaultAdminDelayChangeScheduled(opts *bind.FilterOpts) (*Verse8MarketOwnerDefaultAdminDelayChangeScheduledIterator, error) {

	logs, sub, err := _Verse8MarketOwner.contract.FilterLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return &Verse8MarketOwnerDefaultAdminDelayChangeScheduledIterator{contract: _Verse8MarketOwner.contract, event: "DefaultAdminDelayChangeScheduled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminDelayChangeScheduled is a free log subscription operation binding the contract event 0xf1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b.
//
// Solidity: event DefaultAdminDelayChangeScheduled(uint48 newDelay, uint48 effectSchedule)
func (_Verse8MarketOwner *Verse8MarketOwnerFilterer) WatchDefaultAdminDelayChangeScheduled(opts *bind.WatchOpts, sink chan<- *Verse8MarketOwnerDefaultAdminDelayChangeScheduled) (event.Subscription, error) {

	logs, sub, err := _Verse8MarketOwner.contract.WatchLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Verse8MarketOwnerDefaultAdminDelayChangeScheduled)
				if err := _Verse8MarketOwner.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
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

// ParseDefaultAdminDelayChangeScheduled is a log parse operation binding the contract event 0xf1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b.
//
// Solidity: event DefaultAdminDelayChangeScheduled(uint48 newDelay, uint48 effectSchedule)
func (_Verse8MarketOwner *Verse8MarketOwnerFilterer) ParseDefaultAdminDelayChangeScheduled(log types.Log) (*Verse8MarketOwnerDefaultAdminDelayChangeScheduled, error) {
	event := new(Verse8MarketOwnerDefaultAdminDelayChangeScheduled)
	if err := _Verse8MarketOwner.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Verse8MarketOwnerDefaultAdminTransferCanceledIterator is returned from FilterDefaultAdminTransferCanceled and is used to iterate over the raw logs and unpacked data for DefaultAdminTransferCanceled events raised by the Verse8MarketOwner contract.
type Verse8MarketOwnerDefaultAdminTransferCanceledIterator struct {
	Event *Verse8MarketOwnerDefaultAdminTransferCanceled // Event containing the contract specifics and raw log

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
func (it *Verse8MarketOwnerDefaultAdminTransferCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Verse8MarketOwnerDefaultAdminTransferCanceled)
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
		it.Event = new(Verse8MarketOwnerDefaultAdminTransferCanceled)
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
func (it *Verse8MarketOwnerDefaultAdminTransferCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Verse8MarketOwnerDefaultAdminTransferCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Verse8MarketOwnerDefaultAdminTransferCanceled represents a DefaultAdminTransferCanceled event raised by the Verse8MarketOwner contract.
type Verse8MarketOwnerDefaultAdminTransferCanceled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminTransferCanceled is a free log retrieval operation binding the contract event 0x8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109.
//
// Solidity: event DefaultAdminTransferCanceled()
func (_Verse8MarketOwner *Verse8MarketOwnerFilterer) FilterDefaultAdminTransferCanceled(opts *bind.FilterOpts) (*Verse8MarketOwnerDefaultAdminTransferCanceledIterator, error) {

	logs, sub, err := _Verse8MarketOwner.contract.FilterLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return &Verse8MarketOwnerDefaultAdminTransferCanceledIterator{contract: _Verse8MarketOwner.contract, event: "DefaultAdminTransferCanceled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminTransferCanceled is a free log subscription operation binding the contract event 0x8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109.
//
// Solidity: event DefaultAdminTransferCanceled()
func (_Verse8MarketOwner *Verse8MarketOwnerFilterer) WatchDefaultAdminTransferCanceled(opts *bind.WatchOpts, sink chan<- *Verse8MarketOwnerDefaultAdminTransferCanceled) (event.Subscription, error) {

	logs, sub, err := _Verse8MarketOwner.contract.WatchLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Verse8MarketOwnerDefaultAdminTransferCanceled)
				if err := _Verse8MarketOwner.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
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

// ParseDefaultAdminTransferCanceled is a log parse operation binding the contract event 0x8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109.
//
// Solidity: event DefaultAdminTransferCanceled()
func (_Verse8MarketOwner *Verse8MarketOwnerFilterer) ParseDefaultAdminTransferCanceled(log types.Log) (*Verse8MarketOwnerDefaultAdminTransferCanceled, error) {
	event := new(Verse8MarketOwnerDefaultAdminTransferCanceled)
	if err := _Verse8MarketOwner.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Verse8MarketOwnerDefaultAdminTransferScheduledIterator is returned from FilterDefaultAdminTransferScheduled and is used to iterate over the raw logs and unpacked data for DefaultAdminTransferScheduled events raised by the Verse8MarketOwner contract.
type Verse8MarketOwnerDefaultAdminTransferScheduledIterator struct {
	Event *Verse8MarketOwnerDefaultAdminTransferScheduled // Event containing the contract specifics and raw log

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
func (it *Verse8MarketOwnerDefaultAdminTransferScheduledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Verse8MarketOwnerDefaultAdminTransferScheduled)
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
		it.Event = new(Verse8MarketOwnerDefaultAdminTransferScheduled)
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
func (it *Verse8MarketOwnerDefaultAdminTransferScheduledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Verse8MarketOwnerDefaultAdminTransferScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Verse8MarketOwnerDefaultAdminTransferScheduled represents a DefaultAdminTransferScheduled event raised by the Verse8MarketOwner contract.
type Verse8MarketOwnerDefaultAdminTransferScheduled struct {
	NewAdmin       common.Address
	AcceptSchedule *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminTransferScheduled is a free log retrieval operation binding the contract event 0x3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6.
//
// Solidity: event DefaultAdminTransferScheduled(address indexed newAdmin, uint48 acceptSchedule)
func (_Verse8MarketOwner *Verse8MarketOwnerFilterer) FilterDefaultAdminTransferScheduled(opts *bind.FilterOpts, newAdmin []common.Address) (*Verse8MarketOwnerDefaultAdminTransferScheduledIterator, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _Verse8MarketOwner.contract.FilterLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return &Verse8MarketOwnerDefaultAdminTransferScheduledIterator{contract: _Verse8MarketOwner.contract, event: "DefaultAdminTransferScheduled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminTransferScheduled is a free log subscription operation binding the contract event 0x3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6.
//
// Solidity: event DefaultAdminTransferScheduled(address indexed newAdmin, uint48 acceptSchedule)
func (_Verse8MarketOwner *Verse8MarketOwnerFilterer) WatchDefaultAdminTransferScheduled(opts *bind.WatchOpts, sink chan<- *Verse8MarketOwnerDefaultAdminTransferScheduled, newAdmin []common.Address) (event.Subscription, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _Verse8MarketOwner.contract.WatchLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Verse8MarketOwnerDefaultAdminTransferScheduled)
				if err := _Verse8MarketOwner.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
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

// ParseDefaultAdminTransferScheduled is a log parse operation binding the contract event 0x3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6.
//
// Solidity: event DefaultAdminTransferScheduled(address indexed newAdmin, uint48 acceptSchedule)
func (_Verse8MarketOwner *Verse8MarketOwnerFilterer) ParseDefaultAdminTransferScheduled(log types.Log) (*Verse8MarketOwnerDefaultAdminTransferScheduled, error) {
	event := new(Verse8MarketOwnerDefaultAdminTransferScheduled)
	if err := _Verse8MarketOwner.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Verse8MarketOwnerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Verse8MarketOwner contract.
type Verse8MarketOwnerRoleAdminChangedIterator struct {
	Event *Verse8MarketOwnerRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *Verse8MarketOwnerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Verse8MarketOwnerRoleAdminChanged)
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
		it.Event = new(Verse8MarketOwnerRoleAdminChanged)
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
func (it *Verse8MarketOwnerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Verse8MarketOwnerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Verse8MarketOwnerRoleAdminChanged represents a RoleAdminChanged event raised by the Verse8MarketOwner contract.
type Verse8MarketOwnerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Verse8MarketOwner *Verse8MarketOwnerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*Verse8MarketOwnerRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Verse8MarketOwner.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &Verse8MarketOwnerRoleAdminChangedIterator{contract: _Verse8MarketOwner.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Verse8MarketOwner *Verse8MarketOwnerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *Verse8MarketOwnerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Verse8MarketOwner.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Verse8MarketOwnerRoleAdminChanged)
				if err := _Verse8MarketOwner.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Verse8MarketOwner *Verse8MarketOwnerFilterer) ParseRoleAdminChanged(log types.Log) (*Verse8MarketOwnerRoleAdminChanged, error) {
	event := new(Verse8MarketOwnerRoleAdminChanged)
	if err := _Verse8MarketOwner.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Verse8MarketOwnerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Verse8MarketOwner contract.
type Verse8MarketOwnerRoleGrantedIterator struct {
	Event *Verse8MarketOwnerRoleGranted // Event containing the contract specifics and raw log

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
func (it *Verse8MarketOwnerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Verse8MarketOwnerRoleGranted)
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
		it.Event = new(Verse8MarketOwnerRoleGranted)
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
func (it *Verse8MarketOwnerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Verse8MarketOwnerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Verse8MarketOwnerRoleGranted represents a RoleGranted event raised by the Verse8MarketOwner contract.
type Verse8MarketOwnerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Verse8MarketOwner *Verse8MarketOwnerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*Verse8MarketOwnerRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Verse8MarketOwner.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &Verse8MarketOwnerRoleGrantedIterator{contract: _Verse8MarketOwner.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Verse8MarketOwner *Verse8MarketOwnerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *Verse8MarketOwnerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Verse8MarketOwner.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Verse8MarketOwnerRoleGranted)
				if err := _Verse8MarketOwner.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Verse8MarketOwner *Verse8MarketOwnerFilterer) ParseRoleGranted(log types.Log) (*Verse8MarketOwnerRoleGranted, error) {
	event := new(Verse8MarketOwnerRoleGranted)
	if err := _Verse8MarketOwner.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Verse8MarketOwnerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Verse8MarketOwner contract.
type Verse8MarketOwnerRoleRevokedIterator struct {
	Event *Verse8MarketOwnerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *Verse8MarketOwnerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Verse8MarketOwnerRoleRevoked)
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
		it.Event = new(Verse8MarketOwnerRoleRevoked)
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
func (it *Verse8MarketOwnerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Verse8MarketOwnerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Verse8MarketOwnerRoleRevoked represents a RoleRevoked event raised by the Verse8MarketOwner contract.
type Verse8MarketOwnerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Verse8MarketOwner *Verse8MarketOwnerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*Verse8MarketOwnerRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Verse8MarketOwner.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &Verse8MarketOwnerRoleRevokedIterator{contract: _Verse8MarketOwner.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Verse8MarketOwner *Verse8MarketOwnerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *Verse8MarketOwnerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Verse8MarketOwner.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Verse8MarketOwnerRoleRevoked)
				if err := _Verse8MarketOwner.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Verse8MarketOwner *Verse8MarketOwnerFilterer) ParseRoleRevoked(log types.Log) (*Verse8MarketOwnerRoleRevoked, error) {
	event := new(Verse8MarketOwnerRoleRevoked)
	if err := _Verse8MarketOwner.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}