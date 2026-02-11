// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package binding

import (
	"bytes"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = bytes.Equal
	_ = errors.New
	_ = big.NewInt
	_ = common.Big1
	_ = types.BloomLookup
	_ = abi.ConvertType
)

// Verse8MarketOwnerCreatePairArgs is an auto generated low-level Go binding around an user-defined struct.
type Verse8MarketOwnerCreatePairArgs struct {
	Market   common.Address
	Base     common.Address
	TickSize *big.Int
	LotSize  *big.Int
	FeeData  []byte
}

// Verse8MarketOwnerExecuteBatchArgs is an auto generated low-level Go binding around an user-defined struct.
type Verse8MarketOwnerExecuteBatchArgs struct {
	To    common.Address
	Value *big.Int
	Data  []byte
}

// Verse8MarketOwnerMetaData contains all meta data concerning the Verse8MarketOwner contract.
var Verse8MarketOwnerMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"creators\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"beginDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"}],\"name\":\"changeDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tickSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lotSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"name\":\"createPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tickSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lotSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structVerse8MarketOwner.CreatePairArgs[]\",\"name\":\"args\",\"type\":\"tuple[]\"}],\"name\":\"createPairs\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelayIncreaseWait\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structVerse8MarketOwner.ExecuteBatchArgs[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"executeBatch\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollbackDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminDelayChangeCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"effectSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminDelayChangeScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminTransferCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"acceptSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminTransferScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"name\":\"AccessControlEnforcedDefaultAdminDelay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccessControlEnforcedDefaultAdminRules\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"defaultAdmin\",\"type\":\"address\"}],\"name\":\"AccessControlInvalidDefaultAdmin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Verse8MarketOwner__CallFailed\",\"type\":\"error\"}]",
	ID:  "Verse8MarketOwner",
	Bin: "0x608060405234801561000f575f5ffd5b5060405161203c38038061203c83398101604081905261002e9161023e565b5f826001600160a01b03811661005d57604051636116401160e11b81525f600482015260240160405180910390fd5b600180546001600160d01b0316600160d01b65ffffffffffff8516021790556100865f826100f9565b5050506100a65f51602061201c5f395f51905f52836100f960201b60201c565b505f5b81518110156100f1576100e85f51602061201c5f395f51905f528383815181106100d5576100d561031c565b60200260200101516100f960201b60201c565b506001016100a9565b505050610330565b5f82610155575f6101126002546001600160a01b031690565b6001600160a01b03161461013957604051631fe1e13d60e11b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b0384161790555b61015f8383610168565b90505b92915050565b5f828152602081815260408083206001600160a01b038516845290915281205460ff16610208575f838152602081815260408083206001600160a01b03861684529091529020805460ff191660011790556101c03390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4506001610162565b505f610162565b80516001600160a01b0381168114610225575f5ffd5b919050565b634e487b7160e01b5f52604160045260245ffd5b5f5f6040838503121561024f575f5ffd5b6102588361020f565b60208401519092506001600160401b03811115610273575f5ffd5b8301601f81018513610283575f5ffd5b80516001600160401b0381111561029c5761029c61022a565b604051600582901b90603f8201601f191681016001600160401b03811182821017156102ca576102ca61022a565b6040529182526020818401810192908101888411156102e7575f5ffd5b6020850194505b8385101561030d576102ff8561020f565b8152602094850194016102ee565b50809450505050509250929050565b634e487b7160e01b5f52603260045260245ffd5b611cdf8061033d5f395ff3fe608060405234801561000f575f5ffd5b5060043610610179575f3560e01c80638b481bbd116100d2578063b61d27f611610088578063cf6eefb711610063578063cf6eefb714610382578063d547741f146103ce578063d602b9fd146103e1575f5ffd5b8063b61d27f614610352578063cc8463c814610372578063cefc14291461037a575f5ffd5b806391d14854116100b857806391d14854146102e1578063a1eda53c14610324578063a217fddf1461034b575f5ffd5b80638b481bbd146102c65780638da5cb5b146102d9575f5ffd5b806334fcd5be11610132578063649a5ec71161010d578063649a5ec71461025457806380d3ff511461026757806384ef8ffc14610287575f5ffd5b806334fcd5be1461020e57806336568abe1461022e578063634e93da14610241575f5ffd5b80630aa6220b116101625780630aa6220b146101c1578063248a9ca3146101cb5780632f2ff15d146101fb575f5ffd5b806301ffc9a71461017d578063022d63fb146101a5575b5f5ffd5b61019061018b3660046115cc565b6103e9565b60405190151581526020015b60405180910390f35b620697805b60405165ffffffffffff909116815260200161019c565b6101c9610444565b005b6101ed6101d936600461160b565b5f9081526020819052604090206001015490565b60405190815260200161019c565b6101c9610209366004611643565b610459565b61022161021c3660046116b9565b61049e565b60405161019c9190611744565b6101c961023c366004611643565b610660565b6101c961024f3660046117c5565b61076a565b6101c96102623660046117e0565b61077d565b61027a6102753660046116b9565b610790565b60405161019c9190611805565b60025473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161019c565b6102a16102d4366004611933565b610930565b6102a16109fe565b6101906102ef366004611643565b5f9182526020828152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b61032c610a23565b6040805165ffffffffffff93841681529290911660208301520161019c565b6101ed5f81565b6103656103603660046119a4565b610a9d565b60405161019c9190611a29565b6101aa610b4e565b6101c9610beb565b6001546040805173ffffffffffffffffffffffffffffffffffffffff831681527401000000000000000000000000000000000000000090920465ffffffffffff1660208301520161019c565b6101c96103dc366004611643565b610c47565b6101c9610c88565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f3149878600000000000000000000000000000000000000000000000000000000148061043e575061043e82610c9a565b92915050565b5f61044e81610d30565b610456610d3a565b50565b81610490576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61049a8282610d46565b5050565b60605f6104aa81610d30565b8267ffffffffffffffff8111156104c3576104c361185d565b6040519080825280602002602001820160405280156104f657816020015b60608152602001906001900390816104e15790505b5091505f5b83811015610658575f5f86868481811061051757610517611a3b565b90506020028101906105299190611a68565b6105379060208101906117c5565b73ffffffffffffffffffffffffffffffffffffffff1687878581811061055f5761055f611a3b565b90506020028101906105719190611a68565b6020013588888681811061058757610587611a3b565b90506020028101906105999190611a68565b6105a7906040810190611aa4565b6040516105b5929190611b05565b5f6040518083038185875af1925050503d805f81146105ef576040519150601f19603f3d011682016040523d82523d5f602084013e6105f4565b606091505b509150915081610630576040517fbc5c9d6b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8085848151811061064357610643611a3b565b602090810291909101015250506001016104fb565b505092915050565b81158015610688575060025473ffffffffffffffffffffffffffffffffffffffff8281169116145b156107605760015473ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900465ffffffffffff16811515806106dc575065ffffffffffff8116155b806106ef57504265ffffffffffff821610155b15610735576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff821660048201526024015b60405180910390fd5b5050600180547fffffffffffff000000000000ffffffffffffffffffffffffffffffffffffffff1690555b61049a8282610d70565b5f61077481610d30565b61049a82610dce565b5f61078781610d30565b61049a82610e4d565b60607f02d639b3242e624c4062ce3346179a769447ef0a01fb09608afc904d0268f1906107bc81610d30565b5f8367ffffffffffffffff8111156107d6576107d661185d565b6040519080825280602002602001820160405280156107ff578160200160208202803683370190505b5090505f5b84811015610927575f86868381811061081f5761081f611a3b565b90506020028101906108319190611b14565b61083a90611b46565b805160208201516040808401516060850151608086015192517f4732c7e500000000000000000000000000000000000000000000000000000000815295965073ffffffffffffffffffffffffffffffffffffffff90941694634732c7e5946108a794939091600401611be1565b6020604051808303815f875af11580156108c3573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108e79190611c25565b8383815181106108f9576108f9611a3b565b73ffffffffffffffffffffffffffffffffffffffff9092166020928302919091019091015250600101610804565b50949350505050565b5f7f02d639b3242e624c4062ce3346179a769447ef0a01fb09608afc904d0268f19061095b81610d30565b6040517f4732c7e500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff881690634732c7e5906109b3908990899089908990600401611be1565b6020604051808303815f875af11580156109cf573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109f39190611c25565b979650505050505050565b5f610a1e60025473ffffffffffffffffffffffffffffffffffffffff1690565b905090565b6002545f907a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015158015610a6557504265ffffffffffff821610155b610a70575f5f610a95565b60025474010000000000000000000000000000000000000000900465ffffffffffff16815b915091509091565b60605f610aa981610d30565b5f5f8773ffffffffffffffffffffffffffffffffffffffff16878787604051610ad3929190611b05565b5f6040518083038185875af1925050503d805f8114610b0d576040519150601f19603f3d011682016040523d82523d5f602084013e610b12565b606091505b5091509150816109f3576040517fbc5c9d6b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002545f907a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015158015610b8f57504265ffffffffffff8216105b610bc1576001547a010000000000000000000000000000000000000000000000000000900465ffffffffffff16610be5565b60025474010000000000000000000000000000000000000000900465ffffffffffff165b91505090565b60015473ffffffffffffffffffffffffffffffffffffffff16338114610c3f576040517fc22c802200000000000000000000000000000000000000000000000000000000815233600482015260240161072c565b610456610ebc565b81610c7e576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61049a8282610fad565b5f610c9281610d30565b610456610fd1565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061043e57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000083161461043e565b6104568133610fdb565b610d445f5f611060565b565b5f82815260208190526040902060010154610d6081610d30565b610d6a83836111b9565b50505050565b73ffffffffffffffffffffffffffffffffffffffff81163314610dbf576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610dc9828261127e565b505050565b5f610dd7610b4e565b610de0426112df565b610dea9190611c6d565b9050610df6828261132e565b60405165ffffffffffff8216815273ffffffffffffffffffffffffffffffffffffffff8316907f3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed69060200160405180910390a25050565b5f610e57826113c9565b610e60426112df565b610e6a9190611c6d565b9050610e768282611060565b6040805165ffffffffffff8085168252831660208201527ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b910160405180910390a15050565b60015473ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900465ffffffffffff16801580610f0c57504265ffffffffffff821610155b15610f4d576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff8216600482015260240161072c565b610f755f610f7060025473ffffffffffffffffffffffffffffffffffffffff1690565b61127e565b50610f805f836111b9565b5050600180547fffffffffffff000000000000000000000000000000000000000000000000000016905550565b5f82815260208190526040902060010154610fc781610d30565b610d6a838361127e565b610d445f5f61132e565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff1661049a576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024810183905260440161072c565b6002547a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015611134574265ffffffffffff8216101561110b576002546001805479ffffffffffffffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000090920465ffffffffffff167a01000000000000000000000000000000000000000000000000000002919091179055611134565b6040517f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5905f90a15b506002805473ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000065ffffffffffff9485160279ffffffffffffffffffffffffffffffffffffffffffffffffffff16177a0100000000000000000000000000000000000000000000000000009290931691909102919091179055565b5f8261126d575f6111df60025473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff161461122c576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84161790555b611277838361141a565b9392505050565b5f821580156112a7575060025473ffffffffffffffffffffffffffffffffffffffff8381169116145b156112d557600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001690555b6112778383611513565b5f65ffffffffffff82111561132a576040517f6dfcc650000000000000000000000000000000000000000000000000000000008152603060048201526024810183905260440161072c565b5090565b600180547401000000000000000000000000000000000000000065ffffffffffff84811682027fffffffffffff0000000000000000000000000000000000000000000000000000841673ffffffffffffffffffffffffffffffffffffffff881617179093559004168015610dc9576040517f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109905f90a1505050565b5f5f6113d3610b4e565b90508065ffffffffffff168365ffffffffffff16116113fb576113f68382611c8b565b611277565b61127765ffffffffffff8416620697805f828218828410028218611277565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff1661150c575f8381526020818152604080832073ffffffffffffffffffffffffffffffffffffffff86168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556114aa3390565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a450600161043e565b505f61043e565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff161561150c575f8381526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8616808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a450600161043e565b5f602082840312156115dc575f5ffd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114611277575f5ffd5b5f6020828403121561161b575f5ffd5b5035919050565b73ffffffffffffffffffffffffffffffffffffffff81168114610456575f5ffd5b5f5f60408385031215611654575f5ffd5b82359150602083013561166681611622565b809150509250929050565b5f5f83601f840112611681575f5ffd5b50813567ffffffffffffffff811115611698575f5ffd5b6020830191508360208260051b85010111156116b2575f5ffd5b9250929050565b5f5f602083850312156116ca575f5ffd5b823567ffffffffffffffff8111156116e0575f5ffd5b6116ec85828601611671565b90969095509350505050565b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b828110156117b9577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08786030184526117a48583516116f8565b9450602093840193919091019060010161176a565b50929695505050505050565b5f602082840312156117d5575f5ffd5b813561127781611622565b5f602082840312156117f0575f5ffd5b813565ffffffffffff81168114611277575f5ffd5b602080825282518282018190525f918401906040840190835b8181101561185257835173ffffffffffffffffffffffffffffffffffffffff1683526020938401939092019160010161181e565b509095945050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f82601f830112611899575f5ffd5b813567ffffffffffffffff8111156118b3576118b361185d565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810167ffffffffffffffff811182821017156119005761190061185d565b604052818152838201602001851015611917575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f5f5f5f5f60a08688031215611947575f5ffd5b853561195281611622565b9450602086013561196281611622565b93506040860135925060608601359150608086013567ffffffffffffffff81111561198b575f5ffd5b6119978882890161188a565b9150509295509295909350565b5f5f5f5f606085870312156119b7575f5ffd5b84356119c281611622565b935060208501359250604085013567ffffffffffffffff8111156119e4575f5ffd5b8501601f810187136119f4575f5ffd5b803567ffffffffffffffff811115611a0a575f5ffd5b876020828401011115611a1b575f5ffd5b949793965060200194505050565b602081525f61127760208301846116f8565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f82357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa1833603018112611a9a575f5ffd5b9190910192915050565b5f5f83357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112611ad7575f5ffd5b83018035915067ffffffffffffffff821115611af1575f5ffd5b6020019150368190038213156116b2575f5ffd5b818382375f9101908152919050565b5f82357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff61833603018112611a9a575f5ffd5b5f60a08236031215611b56575f5ffd5b60405160a0810167ffffffffffffffff81118282101715611b7957611b7961185d565b6040528235611b8781611622565b81526020830135611b9781611622565b60208201526040838101359082015260608084013590820152608083013567ffffffffffffffff811115611bc9575f5ffd5b611bd53682860161188a565b60808301525092915050565b73ffffffffffffffffffffffffffffffffffffffff85168152836020820152826040820152608060608201525f611c1b60808301846116f8565b9695505050505050565b5f60208284031215611c35575f5ffd5b815161127781611622565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b65ffffffffffff818116838216019081111561043e5761043e611c40565b65ffffffffffff828116828216039081111561043e5761043e611c4056fea2646970667358221220bb54034b318c80fe8c088492f91e2d7b2a3fa73439f382bafd2fe0634be187cd64736f6c634300081c003302d639b3242e624c4062ce3346179a769447ef0a01fb09608afc904d0268f190",
}

// Verse8MarketOwner is an auto generated Go binding around an Ethereum contract.
type Verse8MarketOwner struct {
	abi abi.ABI
}

// NewVerse8MarketOwner creates a new instance of Verse8MarketOwner.
func NewVerse8MarketOwner() *Verse8MarketOwner {
	parsed, err := Verse8MarketOwnerMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Verse8MarketOwner{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Verse8MarketOwner) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address _owner, address[] creators) returns()
func (verse8MarketOwner *Verse8MarketOwner) PackConstructor(_owner common.Address, creators []common.Address) []byte {
	enc, err := verse8MarketOwner.abi.Pack("", _owner, creators)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackDEFAULTADMINROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (verse8MarketOwner *Verse8MarketOwner) PackDEFAULTADMINROLE() []byte {
	enc, err := verse8MarketOwner.abi.Pack("DEFAULT_ADMIN_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDEFAULTADMINROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (verse8MarketOwner *Verse8MarketOwner) UnpackDEFAULTADMINROLE(data []byte) ([32]byte, error) {
	out, err := verse8MarketOwner.abi.Unpack("DEFAULT_ADMIN_ROLE", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackAcceptDefaultAdminTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcefc1429.
//
// Solidity: function acceptDefaultAdminTransfer() returns()
func (verse8MarketOwner *Verse8MarketOwner) PackAcceptDefaultAdminTransfer() []byte {
	enc, err := verse8MarketOwner.abi.Pack("acceptDefaultAdminTransfer")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackBeginDefaultAdminTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x634e93da.
//
// Solidity: function beginDefaultAdminTransfer(address newAdmin) returns()
func (verse8MarketOwner *Verse8MarketOwner) PackBeginDefaultAdminTransfer(newAdmin common.Address) []byte {
	enc, err := verse8MarketOwner.abi.Pack("beginDefaultAdminTransfer", newAdmin)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCancelDefaultAdminTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (verse8MarketOwner *Verse8MarketOwner) PackCancelDefaultAdminTransfer() []byte {
	enc, err := verse8MarketOwner.abi.Pack("cancelDefaultAdminTransfer")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackChangeDefaultAdminDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (verse8MarketOwner *Verse8MarketOwner) PackChangeDefaultAdminDelay(newDelay *big.Int) []byte {
	enc, err := verse8MarketOwner.abi.Pack("changeDefaultAdminDelay", newDelay)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCreatePair is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8b481bbd.
//
// Solidity: function createPair(address market, address base, uint256 tickSize, uint256 lotSize, bytes feeData) returns(address)
func (verse8MarketOwner *Verse8MarketOwner) PackCreatePair(market common.Address, base common.Address, tickSize *big.Int, lotSize *big.Int, feeData []byte) []byte {
	enc, err := verse8MarketOwner.abi.Pack("createPair", market, base, tickSize, lotSize, feeData)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackCreatePair is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8b481bbd.
//
// Solidity: function createPair(address market, address base, uint256 tickSize, uint256 lotSize, bytes feeData) returns(address)
func (verse8MarketOwner *Verse8MarketOwner) UnpackCreatePair(data []byte) (common.Address, error) {
	out, err := verse8MarketOwner.abi.Unpack("createPair", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackCreatePairs is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x80d3ff51.
//
// Solidity: function createPairs((address,address,uint256,uint256,bytes)[] args) returns(address[])
func (verse8MarketOwner *Verse8MarketOwner) PackCreatePairs(args []Verse8MarketOwnerCreatePairArgs) []byte {
	enc, err := verse8MarketOwner.abi.Pack("createPairs", args)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackCreatePairs is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x80d3ff51.
//
// Solidity: function createPairs((address,address,uint256,uint256,bytes)[] args) returns(address[])
func (verse8MarketOwner *Verse8MarketOwner) UnpackCreatePairs(data []byte) ([]common.Address, error) {
	out, err := verse8MarketOwner.abi.Unpack("createPairs", data)
	if err != nil {
		return *new([]common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	return out0, err
}

// PackDefaultAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (verse8MarketOwner *Verse8MarketOwner) PackDefaultAdmin() []byte {
	enc, err := verse8MarketOwner.abi.Pack("defaultAdmin")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDefaultAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (verse8MarketOwner *Verse8MarketOwner) UnpackDefaultAdmin(data []byte) (common.Address, error) {
	out, err := verse8MarketOwner.abi.Unpack("defaultAdmin", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackDefaultAdminDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (verse8MarketOwner *Verse8MarketOwner) PackDefaultAdminDelay() []byte {
	enc, err := verse8MarketOwner.abi.Pack("defaultAdminDelay")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDefaultAdminDelay is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (verse8MarketOwner *Verse8MarketOwner) UnpackDefaultAdminDelay(data []byte) (*big.Int, error) {
	out, err := verse8MarketOwner.abi.Unpack("defaultAdminDelay", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackDefaultAdminDelayIncreaseWait is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (verse8MarketOwner *Verse8MarketOwner) PackDefaultAdminDelayIncreaseWait() []byte {
	enc, err := verse8MarketOwner.abi.Pack("defaultAdminDelayIncreaseWait")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDefaultAdminDelayIncreaseWait is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (verse8MarketOwner *Verse8MarketOwner) UnpackDefaultAdminDelayIncreaseWait(data []byte) (*big.Int, error) {
	out, err := verse8MarketOwner.abi.Unpack("defaultAdminDelayIncreaseWait", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackExecute is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb61d27f6.
//
// Solidity: function execute(address to, uint256 value, bytes data) returns(bytes)
func (verse8MarketOwner *Verse8MarketOwner) PackExecute(to common.Address, value *big.Int, data []byte) []byte {
	enc, err := verse8MarketOwner.abi.Pack("execute", to, value, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackExecute is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb61d27f6.
//
// Solidity: function execute(address to, uint256 value, bytes data) returns(bytes)
func (verse8MarketOwner *Verse8MarketOwner) UnpackExecute(data []byte) ([]byte, error) {
	out, err := verse8MarketOwner.abi.Unpack("execute", data)
	if err != nil {
		return *new([]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	return out0, err
}

// PackExecuteBatch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x34fcd5be.
//
// Solidity: function executeBatch((address,uint256,bytes)[] calls) returns(bytes[] results)
func (verse8MarketOwner *Verse8MarketOwner) PackExecuteBatch(calls []Verse8MarketOwnerExecuteBatchArgs) []byte {
	enc, err := verse8MarketOwner.abi.Pack("executeBatch", calls)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackExecuteBatch is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x34fcd5be.
//
// Solidity: function executeBatch((address,uint256,bytes)[] calls) returns(bytes[] results)
func (verse8MarketOwner *Verse8MarketOwner) UnpackExecuteBatch(data []byte) ([][]byte, error) {
	out, err := verse8MarketOwner.abi.Unpack("executeBatch", data)
	if err != nil {
		return *new([][]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)
	return out0, err
}

// PackGetRoleAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (verse8MarketOwner *Verse8MarketOwner) PackGetRoleAdmin(role [32]byte) []byte {
	enc, err := verse8MarketOwner.abi.Pack("getRoleAdmin", role)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (verse8MarketOwner *Verse8MarketOwner) UnpackGetRoleAdmin(data []byte) ([32]byte, error) {
	out, err := verse8MarketOwner.abi.Unpack("getRoleAdmin", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackGrantRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (verse8MarketOwner *Verse8MarketOwner) PackGrantRole(role [32]byte, account common.Address) []byte {
	enc, err := verse8MarketOwner.abi.Pack("grantRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackHasRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (verse8MarketOwner *Verse8MarketOwner) PackHasRole(role [32]byte, account common.Address) []byte {
	enc, err := verse8MarketOwner.abi.Pack("hasRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackHasRole is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (verse8MarketOwner *Verse8MarketOwner) UnpackHasRole(data []byte) (bool, error) {
	out, err := verse8MarketOwner.abi.Unpack("hasRole", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (verse8MarketOwner *Verse8MarketOwner) PackOwner() []byte {
	enc, err := verse8MarketOwner.abi.Pack("owner")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (verse8MarketOwner *Verse8MarketOwner) UnpackOwner(data []byte) (common.Address, error) {
	out, err := verse8MarketOwner.abi.Unpack("owner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackPendingDefaultAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (verse8MarketOwner *Verse8MarketOwner) PackPendingDefaultAdmin() []byte {
	enc, err := verse8MarketOwner.abi.Pack("pendingDefaultAdmin")
	if err != nil {
		panic(err)
	}
	return enc
}

// PendingDefaultAdminOutput serves as a container for the return parameters of contract
// method PendingDefaultAdmin.

// UnpackPendingDefaultAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (verse8MarketOwner *Verse8MarketOwner) UnpackPendingDefaultAdmin(data []byte) (PendingDefaultAdminOutput, error) {
	out, err := verse8MarketOwner.abi.Unpack("pendingDefaultAdmin", data)
	outstruct := new(PendingDefaultAdminOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.NewAdmin = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Schedule = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	return *outstruct, err

}

// PackPendingDefaultAdminDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (verse8MarketOwner *Verse8MarketOwner) PackPendingDefaultAdminDelay() []byte {
	enc, err := verse8MarketOwner.abi.Pack("pendingDefaultAdminDelay")
	if err != nil {
		panic(err)
	}
	return enc
}

// PendingDefaultAdminDelayOutput serves as a container for the return parameters of contract
// method PendingDefaultAdminDelay.

// UnpackPendingDefaultAdminDelay is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (verse8MarketOwner *Verse8MarketOwner) UnpackPendingDefaultAdminDelay(data []byte) (PendingDefaultAdminDelayOutput, error) {
	out, err := verse8MarketOwner.abi.Unpack("pendingDefaultAdminDelay", data)
	outstruct := new(PendingDefaultAdminDelayOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.NewDelay = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.Schedule = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	return *outstruct, err

}

// PackRenounceRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (verse8MarketOwner *Verse8MarketOwner) PackRenounceRole(role [32]byte, account common.Address) []byte {
	enc, err := verse8MarketOwner.abi.Pack("renounceRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRevokeRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (verse8MarketOwner *Verse8MarketOwner) PackRevokeRole(role [32]byte, account common.Address) []byte {
	enc, err := verse8MarketOwner.abi.Pack("revokeRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRollbackDefaultAdminDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (verse8MarketOwner *Verse8MarketOwner) PackRollbackDefaultAdminDelay() []byte {
	enc, err := verse8MarketOwner.abi.Pack("rollbackDefaultAdminDelay")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (verse8MarketOwner *Verse8MarketOwner) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := verse8MarketOwner.abi.Pack("supportsInterface", interfaceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (verse8MarketOwner *Verse8MarketOwner) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := verse8MarketOwner.abi.Unpack("supportsInterface", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// Verse8MarketOwnerDefaultAdminDelayChangeCanceled represents a DefaultAdminDelayChangeCanceled event raised by the Verse8MarketOwner contract.
type Verse8MarketOwnerDefaultAdminDelayChangeCanceled struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const Verse8MarketOwnerDefaultAdminDelayChangeCanceledEventName = "DefaultAdminDelayChangeCanceled"

// ContractEventName returns the user-defined event name.
func (Verse8MarketOwnerDefaultAdminDelayChangeCanceled) ContractEventName() string {
	return Verse8MarketOwnerDefaultAdminDelayChangeCanceledEventName
}

// UnpackDefaultAdminDelayChangeCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DefaultAdminDelayChangeCanceled()
func (verse8MarketOwner *Verse8MarketOwner) UnpackDefaultAdminDelayChangeCanceledEvent(log *types.Log) (*Verse8MarketOwnerDefaultAdminDelayChangeCanceled, error) {
	event := "DefaultAdminDelayChangeCanceled"
	if log.Topics[0] != verse8MarketOwner.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Verse8MarketOwnerDefaultAdminDelayChangeCanceled)
	if len(log.Data) > 0 {
		if err := verse8MarketOwner.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range verse8MarketOwner.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// Verse8MarketOwnerDefaultAdminDelayChangeScheduled represents a DefaultAdminDelayChangeScheduled event raised by the Verse8MarketOwner contract.
type Verse8MarketOwnerDefaultAdminDelayChangeScheduled struct {
	NewDelay       *big.Int
	EffectSchedule *big.Int
	Raw            *types.Log // Blockchain specific contextual infos
}

const Verse8MarketOwnerDefaultAdminDelayChangeScheduledEventName = "DefaultAdminDelayChangeScheduled"

// ContractEventName returns the user-defined event name.
func (Verse8MarketOwnerDefaultAdminDelayChangeScheduled) ContractEventName() string {
	return Verse8MarketOwnerDefaultAdminDelayChangeScheduledEventName
}

// UnpackDefaultAdminDelayChangeScheduledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DefaultAdminDelayChangeScheduled(uint48 newDelay, uint48 effectSchedule)
func (verse8MarketOwner *Verse8MarketOwner) UnpackDefaultAdminDelayChangeScheduledEvent(log *types.Log) (*Verse8MarketOwnerDefaultAdminDelayChangeScheduled, error) {
	event := "DefaultAdminDelayChangeScheduled"
	if log.Topics[0] != verse8MarketOwner.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Verse8MarketOwnerDefaultAdminDelayChangeScheduled)
	if len(log.Data) > 0 {
		if err := verse8MarketOwner.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range verse8MarketOwner.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// Verse8MarketOwnerDefaultAdminTransferCanceled represents a DefaultAdminTransferCanceled event raised by the Verse8MarketOwner contract.
type Verse8MarketOwnerDefaultAdminTransferCanceled struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const Verse8MarketOwnerDefaultAdminTransferCanceledEventName = "DefaultAdminTransferCanceled"

// ContractEventName returns the user-defined event name.
func (Verse8MarketOwnerDefaultAdminTransferCanceled) ContractEventName() string {
	return Verse8MarketOwnerDefaultAdminTransferCanceledEventName
}

// UnpackDefaultAdminTransferCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DefaultAdminTransferCanceled()
func (verse8MarketOwner *Verse8MarketOwner) UnpackDefaultAdminTransferCanceledEvent(log *types.Log) (*Verse8MarketOwnerDefaultAdminTransferCanceled, error) {
	event := "DefaultAdminTransferCanceled"
	if log.Topics[0] != verse8MarketOwner.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Verse8MarketOwnerDefaultAdminTransferCanceled)
	if len(log.Data) > 0 {
		if err := verse8MarketOwner.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range verse8MarketOwner.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// Verse8MarketOwnerDefaultAdminTransferScheduled represents a DefaultAdminTransferScheduled event raised by the Verse8MarketOwner contract.
type Verse8MarketOwnerDefaultAdminTransferScheduled struct {
	NewAdmin       common.Address
	AcceptSchedule *big.Int
	Raw            *types.Log // Blockchain specific contextual infos
}

const Verse8MarketOwnerDefaultAdminTransferScheduledEventName = "DefaultAdminTransferScheduled"

// ContractEventName returns the user-defined event name.
func (Verse8MarketOwnerDefaultAdminTransferScheduled) ContractEventName() string {
	return Verse8MarketOwnerDefaultAdminTransferScheduledEventName
}

// UnpackDefaultAdminTransferScheduledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DefaultAdminTransferScheduled(address indexed newAdmin, uint48 acceptSchedule)
func (verse8MarketOwner *Verse8MarketOwner) UnpackDefaultAdminTransferScheduledEvent(log *types.Log) (*Verse8MarketOwnerDefaultAdminTransferScheduled, error) {
	event := "DefaultAdminTransferScheduled"
	if log.Topics[0] != verse8MarketOwner.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Verse8MarketOwnerDefaultAdminTransferScheduled)
	if len(log.Data) > 0 {
		if err := verse8MarketOwner.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range verse8MarketOwner.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// Verse8MarketOwnerRoleAdminChanged represents a RoleAdminChanged event raised by the Verse8MarketOwner contract.
type Verse8MarketOwnerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               *types.Log // Blockchain specific contextual infos
}

const Verse8MarketOwnerRoleAdminChangedEventName = "RoleAdminChanged"

// ContractEventName returns the user-defined event name.
func (Verse8MarketOwnerRoleAdminChanged) ContractEventName() string {
	return Verse8MarketOwnerRoleAdminChangedEventName
}

// UnpackRoleAdminChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (verse8MarketOwner *Verse8MarketOwner) UnpackRoleAdminChangedEvent(log *types.Log) (*Verse8MarketOwnerRoleAdminChanged, error) {
	event := "RoleAdminChanged"
	if log.Topics[0] != verse8MarketOwner.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Verse8MarketOwnerRoleAdminChanged)
	if len(log.Data) > 0 {
		if err := verse8MarketOwner.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range verse8MarketOwner.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// Verse8MarketOwnerRoleGranted represents a RoleGranted event raised by the Verse8MarketOwner contract.
type Verse8MarketOwnerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const Verse8MarketOwnerRoleGrantedEventName = "RoleGranted"

// ContractEventName returns the user-defined event name.
func (Verse8MarketOwnerRoleGranted) ContractEventName() string {
	return Verse8MarketOwnerRoleGrantedEventName
}

// UnpackRoleGrantedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (verse8MarketOwner *Verse8MarketOwner) UnpackRoleGrantedEvent(log *types.Log) (*Verse8MarketOwnerRoleGranted, error) {
	event := "RoleGranted"
	if log.Topics[0] != verse8MarketOwner.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Verse8MarketOwnerRoleGranted)
	if len(log.Data) > 0 {
		if err := verse8MarketOwner.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range verse8MarketOwner.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// Verse8MarketOwnerRoleRevoked represents a RoleRevoked event raised by the Verse8MarketOwner contract.
type Verse8MarketOwnerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const Verse8MarketOwnerRoleRevokedEventName = "RoleRevoked"

// ContractEventName returns the user-defined event name.
func (Verse8MarketOwnerRoleRevoked) ContractEventName() string {
	return Verse8MarketOwnerRoleRevokedEventName
}

// UnpackRoleRevokedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (verse8MarketOwner *Verse8MarketOwner) UnpackRoleRevokedEvent(log *types.Log) (*Verse8MarketOwnerRoleRevoked, error) {
	event := "RoleRevoked"
	if log.Topics[0] != verse8MarketOwner.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Verse8MarketOwnerRoleRevoked)
	if len(log.Data) > 0 {
		if err := verse8MarketOwner.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range verse8MarketOwner.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// UnpackError attempts to decode the provided error data using user-defined
// error definitions.
func (verse8MarketOwner *Verse8MarketOwner) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], verse8MarketOwner.abi.Errors["AccessControlBadConfirmation"].ID.Bytes()[:4]) {
		return verse8MarketOwner.UnpackAccessControlBadConfirmationError(raw[4:])
	}
	if bytes.Equal(raw[:4], verse8MarketOwner.abi.Errors["AccessControlEnforcedDefaultAdminDelay"].ID.Bytes()[:4]) {
		return verse8MarketOwner.UnpackAccessControlEnforcedDefaultAdminDelayError(raw[4:])
	}
	if bytes.Equal(raw[:4], verse8MarketOwner.abi.Errors["AccessControlEnforcedDefaultAdminRules"].ID.Bytes()[:4]) {
		return verse8MarketOwner.UnpackAccessControlEnforcedDefaultAdminRulesError(raw[4:])
	}
	if bytes.Equal(raw[:4], verse8MarketOwner.abi.Errors["AccessControlInvalidDefaultAdmin"].ID.Bytes()[:4]) {
		return verse8MarketOwner.UnpackAccessControlInvalidDefaultAdminError(raw[4:])
	}
	if bytes.Equal(raw[:4], verse8MarketOwner.abi.Errors["AccessControlUnauthorizedAccount"].ID.Bytes()[:4]) {
		return verse8MarketOwner.UnpackAccessControlUnauthorizedAccountError(raw[4:])
	}
	if bytes.Equal(raw[:4], verse8MarketOwner.abi.Errors["SafeCastOverflowedUintDowncast"].ID.Bytes()[:4]) {
		return verse8MarketOwner.UnpackSafeCastOverflowedUintDowncastError(raw[4:])
	}
	if bytes.Equal(raw[:4], verse8MarketOwner.abi.Errors["Verse8MarketOwnerCallFailed"].ID.Bytes()[:4]) {
		return verse8MarketOwner.UnpackVerse8MarketOwnerCallFailedError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// Verse8MarketOwnerAccessControlBadConfirmation represents a AccessControlBadConfirmation error raised by the Verse8MarketOwner contract.
type Verse8MarketOwnerAccessControlBadConfirmation struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlBadConfirmation()
func Verse8MarketOwnerAccessControlBadConfirmationErrorID() common.Hash {
	return common.HexToHash("0x6697b23232a647058342c0724fe7c415cab25915b54e5dbc03f233173d37b41c")
}

// UnpackAccessControlBadConfirmationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlBadConfirmation()
func (verse8MarketOwner *Verse8MarketOwner) UnpackAccessControlBadConfirmationError(raw []byte) (*Verse8MarketOwnerAccessControlBadConfirmation, error) {
	out := new(Verse8MarketOwnerAccessControlBadConfirmation)
	if err := verse8MarketOwner.abi.UnpackIntoInterface(out, "AccessControlBadConfirmation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Verse8MarketOwnerAccessControlEnforcedDefaultAdminDelay represents a AccessControlEnforcedDefaultAdminDelay error raised by the Verse8MarketOwner contract.
type Verse8MarketOwnerAccessControlEnforcedDefaultAdminDelay struct {
	Schedule *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlEnforcedDefaultAdminDelay(uint48 schedule)
func Verse8MarketOwnerAccessControlEnforcedDefaultAdminDelayErrorID() common.Hash {
	return common.HexToHash("0x19ca5ebb8fb33f00e502c9392eddab1501674629178bf69b853cf037aaf4bb5d")
}

// UnpackAccessControlEnforcedDefaultAdminDelayError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlEnforcedDefaultAdminDelay(uint48 schedule)
func (verse8MarketOwner *Verse8MarketOwner) UnpackAccessControlEnforcedDefaultAdminDelayError(raw []byte) (*Verse8MarketOwnerAccessControlEnforcedDefaultAdminDelay, error) {
	out := new(Verse8MarketOwnerAccessControlEnforcedDefaultAdminDelay)
	if err := verse8MarketOwner.abi.UnpackIntoInterface(out, "AccessControlEnforcedDefaultAdminDelay", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Verse8MarketOwnerAccessControlEnforcedDefaultAdminRules represents a AccessControlEnforcedDefaultAdminRules error raised by the Verse8MarketOwner contract.
type Verse8MarketOwnerAccessControlEnforcedDefaultAdminRules struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlEnforcedDefaultAdminRules()
func Verse8MarketOwnerAccessControlEnforcedDefaultAdminRulesErrorID() common.Hash {
	return common.HexToHash("0x3fc3c27ae3db78c81b8f6e685172134623efa268ee8cd8d54be38ad2a74fc13b")
}

// UnpackAccessControlEnforcedDefaultAdminRulesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlEnforcedDefaultAdminRules()
func (verse8MarketOwner *Verse8MarketOwner) UnpackAccessControlEnforcedDefaultAdminRulesError(raw []byte) (*Verse8MarketOwnerAccessControlEnforcedDefaultAdminRules, error) {
	out := new(Verse8MarketOwnerAccessControlEnforcedDefaultAdminRules)
	if err := verse8MarketOwner.abi.UnpackIntoInterface(out, "AccessControlEnforcedDefaultAdminRules", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Verse8MarketOwnerAccessControlInvalidDefaultAdmin represents a AccessControlInvalidDefaultAdmin error raised by the Verse8MarketOwner contract.
type Verse8MarketOwnerAccessControlInvalidDefaultAdmin struct {
	DefaultAdmin common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlInvalidDefaultAdmin(address defaultAdmin)
func Verse8MarketOwnerAccessControlInvalidDefaultAdminErrorID() common.Hash {
	return common.HexToHash("0xc22c8022f2a840d6b6a9f113407715f5bbd4e88c1b0dd9434dc00700ba609ed4")
}

// UnpackAccessControlInvalidDefaultAdminError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlInvalidDefaultAdmin(address defaultAdmin)
func (verse8MarketOwner *Verse8MarketOwner) UnpackAccessControlInvalidDefaultAdminError(raw []byte) (*Verse8MarketOwnerAccessControlInvalidDefaultAdmin, error) {
	out := new(Verse8MarketOwnerAccessControlInvalidDefaultAdmin)
	if err := verse8MarketOwner.abi.UnpackIntoInterface(out, "AccessControlInvalidDefaultAdmin", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Verse8MarketOwnerAccessControlUnauthorizedAccount represents a AccessControlUnauthorizedAccount error raised by the Verse8MarketOwner contract.
type Verse8MarketOwnerAccessControlUnauthorizedAccount struct {
	Account    common.Address
	NeededRole [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlUnauthorizedAccount(address account, bytes32 neededRole)
func Verse8MarketOwnerAccessControlUnauthorizedAccountErrorID() common.Hash {
	return common.HexToHash("0xe2517d3fbfae6f8515ef5ff1ccedc3933ab0cbbda0b492c06eb54ad10ef03b3e")
}

// UnpackAccessControlUnauthorizedAccountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlUnauthorizedAccount(address account, bytes32 neededRole)
func (verse8MarketOwner *Verse8MarketOwner) UnpackAccessControlUnauthorizedAccountError(raw []byte) (*Verse8MarketOwnerAccessControlUnauthorizedAccount, error) {
	out := new(Verse8MarketOwnerAccessControlUnauthorizedAccount)
	if err := verse8MarketOwner.abi.UnpackIntoInterface(out, "AccessControlUnauthorizedAccount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Verse8MarketOwnerSafeCastOverflowedUintDowncast represents a SafeCastOverflowedUintDowncast error raised by the Verse8MarketOwner contract.
type Verse8MarketOwnerSafeCastOverflowedUintDowncast struct {
	Bits  uint8
	Value *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SafeCastOverflowedUintDowncast(uint8 bits, uint256 value)
func Verse8MarketOwnerSafeCastOverflowedUintDowncastErrorID() common.Hash {
	return common.HexToHash("0x6dfcc6503a32754ce7a89698e18201fc5294fd4aad43edefee786f88423b1a12")
}

// UnpackSafeCastOverflowedUintDowncastError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SafeCastOverflowedUintDowncast(uint8 bits, uint256 value)
func (verse8MarketOwner *Verse8MarketOwner) UnpackSafeCastOverflowedUintDowncastError(raw []byte) (*Verse8MarketOwnerSafeCastOverflowedUintDowncast, error) {
	out := new(Verse8MarketOwnerSafeCastOverflowedUintDowncast)
	if err := verse8MarketOwner.abi.UnpackIntoInterface(out, "SafeCastOverflowedUintDowncast", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Verse8MarketOwnerVerse8MarketOwnerCallFailed represents a Verse8MarketOwner__CallFailed error raised by the Verse8MarketOwner contract.
type Verse8MarketOwnerVerse8MarketOwnerCallFailed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error Verse8MarketOwner__CallFailed()
func Verse8MarketOwnerVerse8MarketOwnerCallFailedErrorID() common.Hash {
	return common.HexToHash("0xbc5c9d6b3f8cc8b060bd0e3b5e62873549b63809aae0511d1a90f58ecafc4493")
}

// UnpackVerse8MarketOwnerCallFailedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error Verse8MarketOwner__CallFailed()
func (verse8MarketOwner *Verse8MarketOwner) UnpackVerse8MarketOwnerCallFailedError(raw []byte) (*Verse8MarketOwnerVerse8MarketOwnerCallFailed, error) {
	out := new(Verse8MarketOwnerVerse8MarketOwnerCallFailed)
	if err := verse8MarketOwner.abi.UnpackIntoInterface(out, "Verse8MarketOwnerCallFailed", raw); err != nil {
		return nil, err
	}
	return out, nil
}
