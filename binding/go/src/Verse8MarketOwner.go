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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"creators\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAIR_CREATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"beginDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"}],\"name\":\"changeDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tickSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lotSize\",\"type\":\"uint256\"}],\"name\":\"createPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tickSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lotSize\",\"type\":\"uint256\"}],\"internalType\":\"structVerse8MarketOwner.CreatePairArgs[]\",\"name\":\"args\",\"type\":\"tuple[]\"}],\"name\":\"createPairs\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelayIncreaseWait\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structVerse8MarketOwner.ExecuteBatchArgs[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"executeBatch\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollbackDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminDelayChangeCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"effectSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminDelayChangeScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminTransferCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"acceptSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminTransferScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"name\":\"AccessControlEnforcedDefaultAdminDelay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccessControlEnforcedDefaultAdminRules\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"defaultAdmin\",\"type\":\"address\"}],\"name\":\"AccessControlInvalidDefaultAdmin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Verse8MarketOwner__CallFailed\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"a217fddf": "DEFAULT_ADMIN_ROLE()",
		"f6ab9bb0": "PAIR_CREATOR_ROLE()",
		"cefc1429": "acceptDefaultAdminTransfer()",
		"634e93da": "beginDefaultAdminTransfer(address)",
		"d602b9fd": "cancelDefaultAdminTransfer()",
		"649a5ec7": "changeDefaultAdminDelay(uint48)",
		"035f962d": "createPair(address,address,uint256,uint256)",
		"0cd7cf83": "createPairs((address,address,uint256,uint256)[])",
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
	Bin: "0x608060405234801561000f575f5ffd5b50604051611f51380380611f5183398101604081905261002e91610241565b62015180826001600160a01b03811661006057604051636116401160e11b81525f600482015260240160405180910390fd5b600180546001600160d01b0316600160d01b65ffffffffffff8516021790556100895f826100fc565b5050506100a95f516020611f315f395f51905f52836100fc60201b60201c565b505f5b81518110156100f4576100eb5f516020611f315f395f51905f528383815181106100d8576100d861031f565b60200260200101516100fc60201b60201c565b506001016100ac565b505050610333565b5f82610158575f6101156002546001600160a01b031690565b6001600160a01b03161461013c57604051631fe1e13d60e11b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b0384161790555b610162838361016b565b90505b92915050565b5f828152602081815260408083206001600160a01b038516845290915281205460ff1661020b575f838152602081815260408083206001600160a01b03861684529091529020805460ff191660011790556101c33390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4506001610165565b505f610165565b80516001600160a01b0381168114610228575f5ffd5b919050565b634e487b7160e01b5f52604160045260245ffd5b5f5f60408385031215610252575f5ffd5b61025b83610212565b60208401519092506001600160401b03811115610276575f5ffd5b8301601f81018513610286575f5ffd5b80516001600160401b0381111561029f5761029f61022d565b604051600582901b90603f8201601f191681016001600160401b03811182821017156102cd576102cd61022d565b6040529182526020818401810192908101888411156102ea575f5ffd5b6020850194505b838510156103105761030285610212565b8152602094850194016102f1565b50809450505050509250929050565b634e487b7160e01b5f52603260045260245ffd5b611bf1806103405f395ff3fe608060405234801561000f575f5ffd5b5060043610610184575f3560e01c806384ef8ffc116100dd578063cc8463c811610088578063d547741f11610063578063d547741f146103dd578063d602b9fd146103f0578063f6ab9bb0146103f8575f5ffd5b8063cc8463c814610381578063cefc142914610389578063cf6eefb714610391575f5ffd5b8063a1eda53c116100b8578063a1eda53c14610333578063a217fddf1461035a578063b61d27f614610361575f5ffd5b806384ef8ffc146102ca5780638da5cb5b146102e857806391d14854146102f0575f5ffd5b8063248a9ca31161013d57806336568abe1161011857806336568abe14610291578063634e93da146102a4578063649a5ec7146102b7575f5ffd5b8063248a9ca31461022e5780632f2ff15d1461025e57806334fcd5be14610271575f5ffd5b8063035f962d1161016d578063035f962d146101cc5780630aa6220b146102045780630cd7cf831461020e575f5ffd5b806301ffc9a714610188578063022d63fb146101b0575b5f5ffd5b61019b610196366004611674565b61041f565b60405190151581526020015b60405180910390f35b620697805b60405165ffffffffffff90911681526020016101a7565b6101df6101da3660046116d4565b61047a565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101a7565b61020c61054c565b005b61022161021c366004611717565b610561565b6040516101a79190611788565b61025061023c3660046117e0565b5f9081526020819052604090206001015490565b6040519081526020016101a7565b61020c61026c3660046117f7565b610764565b61028461027f366004611825565b6107a9565b6040516101a791906118d2565b61020c61029f3660046117f7565b61096b565b61020c6102b2366004611953565b610a75565b61020c6102c536600461196e565b610a88565b60025473ffffffffffffffffffffffffffffffffffffffff166101df565b6101df610a9b565b61019b6102fe3660046117f7565b5f9182526020828152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b61033b610ac0565b6040805165ffffffffffff9384168152929091166020830152016101a7565b6102505f81565b61037461036f366004611993565b610b3a565b6040516101a79190611a18565b6101b5610bf6565b61020c610c93565b6001546040805173ffffffffffffffffffffffffffffffffffffffff831681527401000000000000000000000000000000000000000090920465ffffffffffff166020830152016101a7565b61020c6103eb3660046117f7565b610cef565b61020c610d30565b6102507f02d639b3242e624c4062ce3346179a769447ef0a01fb09608afc904d0268f19081565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f31498786000000000000000000000000000000000000000000000000000000001480610474575061047482610d42565b92915050565b5f7f02d639b3242e624c4062ce3346179a769447ef0a01fb09608afc904d0268f1906104a581610dd8565b6040517f284cfbe000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8681166004830152602482018690526044820185905287169063284cfbe0906064016020604051808303815f875af115801561051e573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105429190611a2a565b9695505050505050565b5f61055681610dd8565b61055e610de2565b50565b60607f02d639b3242e624c4062ce3346179a769447ef0a01fb09608afc904d0268f19061058d81610dd8565b5f8367ffffffffffffffff8111156105a7576105a7611a45565b6040519080825280602002602001820160405280156105d0578160200160208202803683370190505b5090505f5b8481101561075b578585828181106105ef576105ef611a72565b6106059260206080909202019081019150611953565b73ffffffffffffffffffffffffffffffffffffffff1663284cfbe087878481811061063257610632611a72565b905060800201602001602081019061064a9190611953565b88888581811061065c5761065c611a72565b9050608002016040013589898681811061067857610678611a72565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e088901b16815273ffffffffffffffffffffffffffffffffffffffff90951660048601526024850193909352506060608090920201013560448201526064016020604051808303815f875af11580156106f8573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061071c9190611a2a565b82828151811061072e5761072e611a72565b73ffffffffffffffffffffffffffffffffffffffff909216602092830291909101909101526001016105d5565b50949350505050565b8161079b576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6107a58282610dee565b5050565b60605f6107b581610dd8565b8267ffffffffffffffff8111156107ce576107ce611a45565b60405190808252806020026020018201604052801561080157816020015b60608152602001906001900390816107ec5790505b5091505f5b83811015610963575f5f86868481811061082257610822611a72565b90506020028101906108349190611a9f565b610842906020810190611953565b73ffffffffffffffffffffffffffffffffffffffff1687878581811061086a5761086a611a72565b905060200281019061087c9190611a9f565b6020013588888681811061089257610892611a72565b90506020028101906108a49190611a9f565b6108b2906040810190611adb565b6040516108c0929190611b43565b5f6040518083038185875af1925050503d805f81146108fa576040519150601f19603f3d011682016040523d82523d5f602084013e6108ff565b606091505b50915091508161093b576040517fbc5c9d6b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8085848151811061094e5761094e611a72565b60209081029190910101525050600101610806565b505092915050565b81158015610993575060025473ffffffffffffffffffffffffffffffffffffffff8281169116145b15610a6b5760015473ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900465ffffffffffff16811515806109e7575065ffffffffffff8116155b806109fa57504265ffffffffffff821610155b15610a40576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff821660048201526024015b60405180910390fd5b5050600180547fffffffffffff000000000000ffffffffffffffffffffffffffffffffffffffff1690555b6107a58282610e18565b5f610a7f81610dd8565b6107a582610e76565b5f610a9281610dd8565b6107a582610ef5565b5f610abb60025473ffffffffffffffffffffffffffffffffffffffff1690565b905090565b6002545f907a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015158015610b0257504265ffffffffffff821610155b610b0d575f5f610b32565b60025474010000000000000000000000000000000000000000900465ffffffffffff16815b915091509091565b60605f610b4681610dd8565b5f5f8773ffffffffffffffffffffffffffffffffffffffff16878787604051610b70929190611b43565b5f6040518083038185875af1925050503d805f8114610baa576040519150601f19603f3d011682016040523d82523d5f602084013e610baf565b606091505b509150915081610beb576040517fbc5c9d6b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b979650505050505050565b6002545f907a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015158015610c3757504265ffffffffffff8216105b610c69576001547a010000000000000000000000000000000000000000000000000000900465ffffffffffff16610c8d565b60025474010000000000000000000000000000000000000000900465ffffffffffff165b91505090565b60015473ffffffffffffffffffffffffffffffffffffffff16338114610ce7576040517fc22c8022000000000000000000000000000000000000000000000000000000008152336004820152602401610a37565b61055e610f64565b81610d26576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6107a58282611055565b5f610d3a81610dd8565b61055e611079565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061047457507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000831614610474565b61055e8133611083565b610dec5f5f611108565b565b5f82815260208190526040902060010154610e0881610dd8565b610e128383611261565b50505050565b73ffffffffffffffffffffffffffffffffffffffff81163314610e67576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610e718282611326565b505050565b5f610e7f610bf6565b610e8842611387565b610e929190611b7f565b9050610e9e82826113d6565b60405165ffffffffffff8216815273ffffffffffffffffffffffffffffffffffffffff8316907f3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed69060200160405180910390a25050565b5f610eff82611471565b610f0842611387565b610f129190611b7f565b9050610f1e8282611108565b6040805165ffffffffffff8085168252831660208201527ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b910160405180910390a15050565b60015473ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900465ffffffffffff16801580610fb457504265ffffffffffff821610155b15610ff5576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff82166004820152602401610a37565b61101d5f61101860025473ffffffffffffffffffffffffffffffffffffffff1690565b611326565b506110285f83611261565b5050600180547fffffffffffff000000000000000000000000000000000000000000000000000016905550565b5f8281526020819052604090206001015461106f81610dd8565b610e128383611326565b610dec5f5f6113d6565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff166107a5576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8216600482015260248101839052604401610a37565b6002547a010000000000000000000000000000000000000000000000000000900465ffffffffffff1680156111dc574265ffffffffffff821610156111b3576002546001805479ffffffffffffffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000090920465ffffffffffff167a010000000000000000000000000000000000000000000000000000029190911790556111dc565b6040517f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5905f90a15b506002805473ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000065ffffffffffff9485160279ffffffffffffffffffffffffffffffffffffffffffffffffffff16177a0100000000000000000000000000000000000000000000000000009290931691909102919091179055565b5f82611315575f61128760025473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff16146112d4576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84161790555b61131f83836114c2565b9392505050565b5f8215801561134f575060025473ffffffffffffffffffffffffffffffffffffffff8381169116145b1561137d57600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001690555b61131f83836115bb565b5f65ffffffffffff8211156113d2576040517f6dfcc6500000000000000000000000000000000000000000000000000000000081526030600482015260248101839052604401610a37565b5090565b600180547401000000000000000000000000000000000000000065ffffffffffff84811682027fffffffffffff0000000000000000000000000000000000000000000000000000841673ffffffffffffffffffffffffffffffffffffffff881617179093559004168015610e71576040517f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109905f90a1505050565b5f5f61147b610bf6565b90508065ffffffffffff168365ffffffffffff16116114a35761149e8382611b9d565b61131f565b61131f65ffffffffffff8416620697805f82821882841002821861131f565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff166115b4575f8381526020818152604080832073ffffffffffffffffffffffffffffffffffffffff86168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556115523390565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4506001610474565b505f610474565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff16156115b4575f8381526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8616808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a4506001610474565b5f60208284031215611684575f5ffd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461131f575f5ffd5b73ffffffffffffffffffffffffffffffffffffffff8116811461055e575f5ffd5b5f5f5f5f608085870312156116e7575f5ffd5b84356116f2816116b3565b93506020850135611702816116b3565b93969395505050506040820135916060013590565b5f5f60208385031215611728575f5ffd5b823567ffffffffffffffff81111561173e575f5ffd5b8301601f8101851361174e575f5ffd5b803567ffffffffffffffff811115611764575f5ffd5b8560208260071b8401011115611778575f5ffd5b6020919091019590945092505050565b602080825282518282018190525f918401906040840190835b818110156117d557835173ffffffffffffffffffffffffffffffffffffffff168352602093840193909201916001016117a1565b509095945050505050565b5f602082840312156117f0575f5ffd5b5035919050565b5f5f60408385031215611808575f5ffd5b82359150602083013561181a816116b3565b809150509250929050565b5f5f60208385031215611836575f5ffd5b823567ffffffffffffffff81111561184c575f5ffd5b8301601f8101851361185c575f5ffd5b803567ffffffffffffffff811115611872575f5ffd5b8560208260051b8401011115611778575f5ffd5b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b82811015611947577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452611932858351611886565b945060209384019391909101906001016118f8565b50929695505050505050565b5f60208284031215611963575f5ffd5b813561131f816116b3565b5f6020828403121561197e575f5ffd5b813565ffffffffffff8116811461131f575f5ffd5b5f5f5f5f606085870312156119a6575f5ffd5b84356119b1816116b3565b935060208501359250604085013567ffffffffffffffff8111156119d3575f5ffd5b8501601f810187136119e3575f5ffd5b803567ffffffffffffffff8111156119f9575f5ffd5b876020828401011115611a0a575f5ffd5b949793965060200194505050565b602081525f61131f6020830184611886565b5f60208284031215611a3a575f5ffd5b815161131f816116b3565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f82357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa1833603018112611ad1575f5ffd5b9190910192915050565b5f5f83357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112611b0e575f5ffd5b83018035915067ffffffffffffffff821115611b28575f5ffd5b602001915036819003821315611b3c575f5ffd5b9250929050565b818382375f9101908152919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b65ffffffffffff818116838216019081111561047457610474611b52565b65ffffffffffff828116828216039081111561047457610474611b5256fea26469706673582212207dd26696706a2691cd984fcf5b8e5f1ac05abf20dd9dba0b99d935c84bf41d9c64736f6c634300081c003302d639b3242e624c4062ce3346179a769447ef0a01fb09608afc904d0268f190",
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

// PAIRCREATORROLE is a free data retrieval call binding the contract method 0xf6ab9bb0.
//
// Solidity: function PAIR_CREATOR_ROLE() view returns(bytes32)
func (_Verse8MarketOwner *Verse8MarketOwnerCaller) PAIRCREATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Verse8MarketOwner.contract.Call(opts, &out, "PAIR_CREATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAIRCREATORROLE is a free data retrieval call binding the contract method 0xf6ab9bb0.
//
// Solidity: function PAIR_CREATOR_ROLE() view returns(bytes32)
func (_Verse8MarketOwner *Verse8MarketOwnerSession) PAIRCREATORROLE() ([32]byte, error) {
	return _Verse8MarketOwner.Contract.PAIRCREATORROLE(&_Verse8MarketOwner.CallOpts)
}

// PAIRCREATORROLE is a free data retrieval call binding the contract method 0xf6ab9bb0.
//
// Solidity: function PAIR_CREATOR_ROLE() view returns(bytes32)
func (_Verse8MarketOwner *Verse8MarketOwnerCallerSession) PAIRCREATORROLE() ([32]byte, error) {
	return _Verse8MarketOwner.Contract.PAIRCREATORROLE(&_Verse8MarketOwner.CallOpts)
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

// CreatePair is a paid mutator transaction binding the contract method 0x035f962d.
//
// Solidity: function createPair(address market, address base, uint256 tickSize, uint256 lotSize) returns(address)
func (_Verse8MarketOwner *Verse8MarketOwnerTransactor) CreatePair(opts *bind.TransactOpts, market common.Address, base common.Address, tickSize *big.Int, lotSize *big.Int) (*types.Transaction, error) {
	return _Verse8MarketOwner.contract.Transact(opts, "createPair", market, base, tickSize, lotSize)
}

// CreatePair is a paid mutator transaction binding the contract method 0x035f962d.
//
// Solidity: function createPair(address market, address base, uint256 tickSize, uint256 lotSize) returns(address)
func (_Verse8MarketOwner *Verse8MarketOwnerSession) CreatePair(market common.Address, base common.Address, tickSize *big.Int, lotSize *big.Int) (*types.Transaction, error) {
	return _Verse8MarketOwner.Contract.CreatePair(&_Verse8MarketOwner.TransactOpts, market, base, tickSize, lotSize)
}

// CreatePair is a paid mutator transaction binding the contract method 0x035f962d.
//
// Solidity: function createPair(address market, address base, uint256 tickSize, uint256 lotSize) returns(address)
func (_Verse8MarketOwner *Verse8MarketOwnerTransactorSession) CreatePair(market common.Address, base common.Address, tickSize *big.Int, lotSize *big.Int) (*types.Transaction, error) {
	return _Verse8MarketOwner.Contract.CreatePair(&_Verse8MarketOwner.TransactOpts, market, base, tickSize, lotSize)
}

// CreatePairs is a paid mutator transaction binding the contract method 0x0cd7cf83.
//
// Solidity: function createPairs((address,address,uint256,uint256)[] args) returns(address[])
func (_Verse8MarketOwner *Verse8MarketOwnerTransactor) CreatePairs(opts *bind.TransactOpts, args []Verse8MarketOwnerCreatePairArgs) (*types.Transaction, error) {
	return _Verse8MarketOwner.contract.Transact(opts, "createPairs", args)
}

// CreatePairs is a paid mutator transaction binding the contract method 0x0cd7cf83.
//
// Solidity: function createPairs((address,address,uint256,uint256)[] args) returns(address[])
func (_Verse8MarketOwner *Verse8MarketOwnerSession) CreatePairs(args []Verse8MarketOwnerCreatePairArgs) (*types.Transaction, error) {
	return _Verse8MarketOwner.Contract.CreatePairs(&_Verse8MarketOwner.TransactOpts, args)
}

// CreatePairs is a paid mutator transaction binding the contract method 0x0cd7cf83.
//
// Solidity: function createPairs((address,address,uint256,uint256)[] args) returns(address[])
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