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

// CrossDexImplV2MetaData contains all meta data concerning the CrossDexImplV2 contract.
var CrossDexImplV2MetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ROUTER\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allMarkets\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"markets\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"quotes\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"checkTickSizeRoles\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quote\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeCollector\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"createMarket\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_routerImpl\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_findPrevPriceCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxMatchCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_cancelLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_marketImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pairImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tickSizeSetter\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"}],\"name\":\"isMarket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketImpl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"name\":\"pairCreated\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pairImpl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"name\":\"pairToMarket\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_marketImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pairImpl\",\"type\":\"address\"}],\"name\":\"reinitialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_marketImpl\",\"type\":\"address\"}],\"name\":\"setMarketImpl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pairImpl\",\"type\":\"address\"}],\"name\":\"setPairImpl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"setter\",\"type\":\"address\"}],\"name\":\"setTickSizeSetter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tickSizeSetter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"quote\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fee_collector\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"MarketCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"before\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"MarketImplSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"before\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"PairImplSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"before\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"TickSizeSetterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Create2EmptyBytecode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"CrossDexInitializeData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"CrossDexInvalidMarketAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"input\",\"type\":\"address\"}],\"name\":\"CrossDexInvalidTickSizeSetter\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"CrossDexUnauthorizedChangeTickSizes\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedDeployment\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"}]",
	ID:  "CrossDexImplV2",
	Bin: "0x60a060405230608052348015610013575f5ffd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b608051612ad26100f95f395f818161179c015281816117c501526119e60152612ad25ff3fe608060405260043610610162575f3560e01c8063715018a6116100c6578063a1eea7781161007c578063e492c27011610057578063e492c27014610474578063e9f7206b14610493578063f2fde38b146104b2575f5ffd5b8063a1eea778146103e1578063a9d951a314610400578063ad3cb1cc1461041f575f5ffd5b80637f764a85116100ac5780637f764a851461034d5780638da5cb5b1461037957806398715ee7146103c2575f5ffd5b8063715018a61461031a578063786b0b111461032e575f5ffd5b8063375a7cba1161011b5780634f1ef286116101015780634f1ef286146102b657806352d1902d146102c95780636ec934da146102eb575f5ffd5b8063375a7cba146102755780633cc047aa14610297575f5ffd5b806332fe7b261161014b57806332fe7b26146101f257806334eaeeb91461021d57806335f7d45614610249575f5ffd5b8063082705731461016657806331ed1bcc146101d1575b5f5ffd5b348015610171575f5ffd5b506101a761018036600461219c565b60066020525f908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b3480156101dc575f5ffd5b506101f06101eb3660046121b5565b6104d1565b005b3480156101fd575f5ffd5b505f546101a79073ffffffffffffffffffffffffffffffffffffffff1681565b348015610228575f5ffd5b506001546101a79073ffffffffffffffffffffffffffffffffffffffff1681565b348015610254575f5ffd5b506002546101a79073ffffffffffffffffffffffffffffffffffffffff1681565b348015610280575f5ffd5b50610289610950565b6040516101c8929190612282565b3480156102a2575f5ffd5b506101f06102b136600461219c565b610a5f565b6101f06102c4366004612386565b610b63565b3480156102d4575f5ffd5b506102dd610b82565b6040519081526020016101c8565b3480156102f6575f5ffd5b5061030a61030536600461219c565b610bb0565b60405190151581526020016101c8565b348015610325575f5ffd5b506101f0610bc2565b348015610339575f5ffd5b506101a76103483660046123d1565b610bd5565b348015610358575f5ffd5b506007546101a79073ffffffffffffffffffffffffffffffffffffffff1681565b348015610384575f5ffd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c1993005473ffffffffffffffffffffffffffffffffffffffff166101a7565b3480156103cd575f5ffd5b506101f06103dc36600461219c565b610e00565b3480156103ec575f5ffd5b506101f06103fb36600461219c565b610f2d565b34801561040b575f5ffd5b506101f061041a366004612479565b610fbf565b34801561042a575f5ffd5b506104676040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b6040516101c891906124f6565b34801561047f575f5ffd5b506101f061048e36600461219c565b6114ce565b34801561049e575f5ffd5b506101f06104ad36600461219c565b6115d2565b3480156104bd575f5ffd5b506101f06104cc36600461219c565b611660565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff165f8115801561051b5750825b90505f8267ffffffffffffffff1660011480156105375750303b155b905081158015610545575080155b1561057c576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000016600117855583156105dd5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b6105e68d6116c0565b73ffffffffffffffffffffffffffffffffffffffff8c1661065a576040517fc53a07b90000000000000000000000000000000000000000000000000000000081527f726f75746572496d706c0000000000000000000000000000000000000000000060048201526024015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff88166106c9576040517fc53a07b90000000000000000000000000000000000000000000000000000000081527f6d61726b6574496d706c000000000000000000000000000000000000000000006004820152602401610651565b73ffffffffffffffffffffffffffffffffffffffff8716610738576040517fc53a07b90000000000000000000000000000000000000000000000000000000081527f70616972496d706c0000000000000000000000000000000000000000000000006004820152602401610651565b5f8c60405161074690612167565b73ffffffffffffffffffffffffffffffffffffffff90911681526040602082018190525f90820152606001604051809103905ff08015801561078a573d5f5f3e3d5ffd5b505f80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040517f80d85911000000000000000000000000000000000000000000000000000000008152600481018f9052602481018e9052604481018d9052919250906380d85911906064015f604051808303815f87803b15801561082c575f5ffd5b505af115801561083e573d5f5f3e3d5ffd5b50506001805473ffffffffffffffffffffffffffffffffffffffff808e167fffffffffffffffffffffffff000000000000000000000000000000000000000092831617909255600280548d8416921691909117905589161592506108e091505057600780547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff88161790555b83156109415784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050505050505050565b6060805f61095e60036116d1565b90508067ffffffffffffffff811115610979576109796122a6565b6040519080825280602002602001820160405280156109a2578160200160208202803683370190505b5092508067ffffffffffffffff8111156109be576109be6122a6565b6040519080825280602002602001820160405280156109e7578160200160208202803683370190505b5091505f5b81811015610a59576109ff6003826116db565b858381518110610a1157610a11612508565b60200260200101858481518110610a2a57610a2a612508565b73ffffffffffffffffffffffffffffffffffffffff9384166020918202929092010152911690526001016109ec565b50509091565b610a676116f6565b73ffffffffffffffffffffffffffffffffffffffff8116610ad6576040517fc53a07b90000000000000000000000000000000000000000000000000000000081527f70616972496d706c0000000000000000000000000000000000000000000000006004820152602401610651565b60025460405173ffffffffffffffffffffffffffffffffffffffff8084169216907f685c330e9b04243f99ff39d8de8578c9af2bfaff120e17a1b5f1d61e004935e2905f90a3600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b610b6b611784565b610b7482611888565b610b7e8282611890565b5050565b5f610b8b6119ce565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b5f610bbc600383611a3d565b92915050565b610bca6116f6565b610bd35f611a65565b565b5f610bde6116f6565b5f60405180602001610bef90612167565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082820381018352601f9091011660408190526001545f5460025473ffffffffffffffffffffffffffffffffffffffff92831693610c5d938d938116928d929116908c908c90602401612535565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f01f79fb6000000000000000000000000000000000000000000000000000000001790529051610ce2939291016125df565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815290829052610d1e9291602001612624565b60405160208183030381529060405290505f8684604051602001610d439291906125df565b6040516020818303038152906040528051906020012090505f610d675f8385611afa565b9050610d756003828a611be6565b508873ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff167f3fd2f44bbde90ebd71c132db1c1c094924bbda4b1437f153ac54830124655c848a89604051610dec9291906125df565b60405180910390a498975050505050505050565b610e086116f6565b73ffffffffffffffffffffffffffffffffffffffff81161580610e45575060075473ffffffffffffffffffffffffffffffffffffffff8281169116145b15610ea0576007546040517f18ddd67e00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff91821660048201529082166024820152604401610651565b60075460405173ffffffffffffffffffffffffffffffffffffffff8084169216907f4f1656ebcdda14dfa0f9951b46485570075b9730e1f1d22cb4ddacc35374fde8905f90a3600780547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60075473ffffffffffffffffffffffffffffffffffffffff161580610f6d575060075473ffffffffffffffffffffffffffffffffffffffff828116911614155b15610fbc576040517ff42eaafb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152602401610651565b50565b610fc76116f6565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0080546002919068010000000000000000900460ff16806110165750805467ffffffffffffffff808416911610155b1561104d576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80547fffffffffffffffffffffffffffffffffffffffffffffff0000000000000000001667ffffffffffffffff8316176801000000000000000017815573ffffffffffffffffffffffffffffffffffffffff84166110f9576040517fc53a07b90000000000000000000000000000000000000000000000000000000081527f6d61726b6574496d706c000000000000000000000000000000000000000000006004820152602401610651565b73ffffffffffffffffffffffffffffffffffffffff8316611168576040517fc53a07b90000000000000000000000000000000000000000000000000000000081527f70616972496d706c0000000000000000000000000000000000000000000000006004820152602401610651565b5f61117360036116d1565b90505f8167ffffffffffffffff81111561118f5761118f6122a6565b6040519080825280602002602001820160405280156111b8578160200160208202803683370190505b5090505f8267ffffffffffffffff8111156111d5576111d56122a6565b6040519080825280602002602001820160405280156111fe578160200160208202803683370190505b5090505f5b83811015611270576112166003826116db565b84838151811061122857611228612508565b6020026020010184848151811061124157611241612508565b73ffffffffffffffffffffffffffffffffffffffff938416602091820292909201015291169052600101611203565b50825b80156112b357806001900390506112ad83828151811061129557611295612508565b60200260200101516003611c1390919063ffffffff16565b50611273565b505f5b8381101561130f576113068282815181106112d3576112d3612508565b60200260200101518483815181106112ed576112ed612508565b60200260200101516003611be69092919063ffffffff16565b506001016112b6565b5060015473ffffffffffffffffffffffffffffffffffffffff8881169116146113b95760015460405173ffffffffffffffffffffffffffffffffffffffff808a169216907f081eadab9f9e460b348376155bdd3b4aed9b2a33469b489416d7416b4cffe49c905f90a3600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff89161790555b60025473ffffffffffffffffffffffffffffffffffffffff8781169116146114625760025460405173ffffffffffffffffffffffffffffffffffffffff8089169216907f685c330e9b04243f99ff39d8de8578c9af2bfaff120e17a1b5f1d61e004935e2905f90a3600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff88161790555b505081547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff1682555060405167ffffffffffffffff831681527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a150505050565b6114d66116f6565b73ffffffffffffffffffffffffffffffffffffffff8116611545576040517fc53a07b90000000000000000000000000000000000000000000000000000000081527f6d61726b6574496d706c000000000000000000000000000000000000000000006004820152602401610651565b60015460405173ffffffffffffffffffffffffffffffffffffffff8084169216907f081eadab9f9e460b348376155bdd3b4aed9b2a33469b489416d7416b4cffe49c905f90a3600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b6115db33610bb0565b611613576040517f74be5d0c000000000000000000000000000000000000000000000000000000008152336004820152602401610651565b73ffffffffffffffffffffffffffffffffffffffff165f90815260066020526040902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001633179055565b6116686116f6565b73ffffffffffffffffffffffffffffffffffffffff81166116b7576040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081525f6004820152602401610651565b610fbc81611a65565b6116c8611c34565b610fbc81611c9b565b5f610bbc82611ca3565b5f8080806116e98686611cad565b9097909650945050505050565b336117357f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c1993005473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614610bd3576040517f118cdaa7000000000000000000000000000000000000000000000000000000008152336004820152602401610651565b3073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016148061185157507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166118387f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614155b15610bd3576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610fbc6116f6565b8173ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015611915575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016820190925261191291810190612638565b60015b611963576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610651565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc81146119bf576040517faa1d49a400000000000000000000000000000000000000000000000000000000815260048101829052602401610651565b6119c98383611cd6565b505050565b3073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610bd3576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f611a5e8373ffffffffffffffffffffffffffffffffffffffff8416611d38565b9392505050565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080547fffffffffffffffffffffffff0000000000000000000000000000000000000000811673ffffffffffffffffffffffffffffffffffffffff848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b5f83471015611b3e576040517fcf47918100000000000000000000000000000000000000000000000000000000815247600482015260248101859052604401610651565b81515f03611b78576040517f4ca249dc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8282516020840186f590503d151981151615611b99576040513d5f823e3d81fd5b73ffffffffffffffffffffffffffffffffffffffff8116611a5e576040517fb06ebf3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f611c0b8473ffffffffffffffffffffffffffffffffffffffff808616908516611d43565b949350505050565b5f611a5e8373ffffffffffffffffffffffffffffffffffffffff8416611d5f565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff16610bd3576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611668611c34565b5f610bbc82611d7b565b5f8080611cba8585611d84565b5f81815260029690960160205260409095205494959350505050565b611cdf82611d8f565b60405173ffffffffffffffffffffffffffffffffffffffff8316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115611d30576119c98282611e5d565b610b7e611edc565b5f611a5e8383611f14565b5f8281526002840160205260408120829055611c0b8484611f2b565b5f8181526002830160205260408120819055611a5e8383611f36565b5f610bbc825490565b5f611a5e8383611f41565b8073ffffffffffffffffffffffffffffffffffffffff163b5f03611df7576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152602401610651565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60605f5f8473ffffffffffffffffffffffffffffffffffffffff1684604051611e86919061264f565b5f60405180830381855af49150503d805f8114611ebe576040519150601f19603f3d011682016040523d82523d5f602084013e611ec3565b606091505b5091509150611ed3858383611f67565b95945050505050565b3415610bd3576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8181526001830160205260408120541515611a5e565b5f611a5e8383611ff6565b5f611a5e8383612042565b5f825f018281548110611f5657611f56612508565b905f5260205f200154905092915050565b606082611f7c57611f7782612125565b611a5e565b8151158015611fa0575073ffffffffffffffffffffffffffffffffffffffff84163b155b15611fef576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610651565b5080611a5e565b5f81815260018301602052604081205461203b57508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610bbc565b505f610bbc565b5f818152600183016020526040812054801561211c575f61206460018361265a565b85549091505f906120779060019061265a565b90508082146120d6575f865f01828154811061209557612095612508565b905f5260205f200154905080875f0184815481106120b5576120b5612508565b5f918252602080832090910192909255918252600188019052604090208390555b85548690806120e7576120e7612692565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050610bbc565b5f915050610bbc565b8051156121355780518082602001fd5b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103dd806126c083390190565b803573ffffffffffffffffffffffffffffffffffffffff81168114612197575f5ffd5b919050565b5f602082840312156121ac575f5ffd5b611a5e82612174565b5f5f5f5f5f5f5f5f610100898b0312156121cd575f5ffd5b6121d689612174565b97506121e460208a01612174565b965060408901359550606089013594506080890135935061220760a08a01612174565b925061221560c08a01612174565b915061222360e08a01612174565b90509295985092959890939650565b5f8151808452602084019350602083015f5b8281101561227857815173ffffffffffffffffffffffffffffffffffffffff16865260209586019590910190600101612244565b5093949350505050565b604081525f6122946040830185612232565b8281036020840152611ed38185612232565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f5f67ffffffffffffffff8411156122ed576122ed6122a6565b506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85018116603f0116810181811067ffffffffffffffff8211171561233a5761233a6122a6565b604052838152905080828401851015612351575f5ffd5b838360208301375f60208583010152509392505050565b5f82601f830112612377575f5ffd5b611a5e838335602085016122d3565b5f5f60408385031215612397575f5ffd5b6123a083612174565b9150602083013567ffffffffffffffff8111156123bb575f5ffd5b6123c785828601612368565b9150509250929050565b5f5f5f5f5f60a086880312156123e5575f5ffd5b6123ee86612174565b94506123fc60208701612174565b935061240a60408701612174565b9250606086013567ffffffffffffffff811115612425575f5ffd5b61243188828901612368565b925050608086013567ffffffffffffffff81111561244d575f5ffd5b8601601f8101881361245d575f5ffd5b61246c888235602084016122d3565b9150509295509295909350565b5f5f6040838503121561248a575f5ffd5b61249383612174565b91506124a160208401612174565b90509250929050565b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f611a5e60208301846124aa565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b73ffffffffffffffffffffffffffffffffffffffff8716815273ffffffffffffffffffffffffffffffffffffffff8616602082015273ffffffffffffffffffffffffffffffffffffffff8516604082015273ffffffffffffffffffffffffffffffffffffffff8416606082015273ffffffffffffffffffffffffffffffffffffffff8316608082015260c060a08201525f6125d360c08301846124aa565b98975050505050505050565b73ffffffffffffffffffffffffffffffffffffffff83168152604060208201525f611c0b60408301846124aa565b5f81518060208401855e5f93019283525090919050565b5f611c0b612632838661260d565b8461260d565b5f60208284031215612648575f5ffd5b5051919050565b5f611a5e828461260d565b81810381811115610bbc577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffdfe60806040526040516103dd3803806103dd8339810160408190526100229161023c565b61002c8282610033565b5050610321565b61003c82610091565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a280511561008557610080828261010c565b505050565b61008d61017f565b5050565b806001600160a01b03163b5f036100cb57604051634c9c8ce360e01b81526001600160a01b03821660048201526024015b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b031684604051610128919061030b565b5f60405180830381855af49150503d805f8114610160576040519150601f19603f3d011682016040523d82523d5f602084013e610165565b606091505b5090925090506101768583836101a0565b95945050505050565b341561019e5760405163b398979f60e01b815260040160405180910390fd5b565b6060826101b5576101b0826101ff565b6101f8565b81511580156101cc57506001600160a01b0384163b155b156101f557604051639996b31560e01b81526001600160a01b03851660048201526024016100c2565b50805b9392505050565b80511561020f5780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b634e487b7160e01b5f52604160045260245ffd5b5f5f6040838503121561024d575f5ffd5b82516001600160a01b0381168114610263575f5ffd5b60208401519092506001600160401b0381111561027e575f5ffd5b8301601f8101851361028e575f5ffd5b80516001600160401b038111156102a7576102a7610228565b604051601f8201601f19908116603f011681016001600160401b03811182821017156102d5576102d5610228565b6040528181528282016020018710156102ec575f5ffd5b8160208401602083015e5f602083830101528093505050509250929050565b5f82518060208501845e5f920191825250919050565b60b08061032d5f395ff3fe6080604052600a600c565b005b60186014601a565b605d565b565b5f60587f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5473ffffffffffffffffffffffffffffffffffffffff1690565b905090565b365f5f375f5f365f845af43d5f5f3e8080156076573d5ff35b3d5ffdfea264697066735822122053cea113bf96df2d97b042d29245e0a41b1290b530767f8b25bfeaed774391b764736f6c634300081c0033a26469706673582212204da2ce1642a2eeee13c4aff6e9f92a9cf4204fedafa07f43a028ab51a84f260b64736f6c634300081c0033",
}

// CrossDexImplV2 is an auto generated Go binding around an Ethereum contract.
type CrossDexImplV2 struct {
	abi abi.ABI
}

// NewCrossDexImplV2 creates a new instance of CrossDexImplV2.
func NewCrossDexImplV2() *CrossDexImplV2 {
	parsed, err := CrossDexImplV2MetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &CrossDexImplV2{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *CrossDexImplV2) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackROUTER is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x32fe7b26.
//
// Solidity: function ROUTER() view returns(address)
func (crossDexImplV2 *CrossDexImplV2) PackROUTER() []byte {
	enc, err := crossDexImplV2.abi.Pack("ROUTER")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackROUTER is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x32fe7b26.
//
// Solidity: function ROUTER() view returns(address)
func (crossDexImplV2 *CrossDexImplV2) UnpackROUTER(data []byte) (common.Address, error) {
	out, err := crossDexImplV2.abi.Unpack("ROUTER", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackUPGRADEINTERFACEVERSION is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (crossDexImplV2 *CrossDexImplV2) PackUPGRADEINTERFACEVERSION() []byte {
	enc, err := crossDexImplV2.abi.Pack("UPGRADE_INTERFACE_VERSION")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackUPGRADEINTERFACEVERSION is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (crossDexImplV2 *CrossDexImplV2) UnpackUPGRADEINTERFACEVERSION(data []byte) (string, error) {
	out, err := crossDexImplV2.abi.Unpack("UPGRADE_INTERFACE_VERSION", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackAllMarkets is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x375a7cba.
//
// Solidity: function allMarkets() view returns(address[] markets, address[] quotes)
func (crossDexImplV2 *CrossDexImplV2) PackAllMarkets() []byte {
	enc, err := crossDexImplV2.abi.Pack("allMarkets")
	if err != nil {
		panic(err)
	}
	return enc
}

// AllMarketsOutput serves as a container for the return parameters of contract
// method AllMarkets.

// UnpackAllMarkets is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x375a7cba.
//
// Solidity: function allMarkets() view returns(address[] markets, address[] quotes)
func (crossDexImplV2 *CrossDexImplV2) UnpackAllMarkets(data []byte) (AllMarketsOutput, error) {
	out, err := crossDexImplV2.abi.Unpack("allMarkets", data)
	outstruct := new(AllMarketsOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Markets = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Quotes = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	return *outstruct, err

}

// PackCheckTickSizeRoles is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa1eea778.
//
// Solidity: function checkTickSizeRoles(address account) view returns()
func (crossDexImplV2 *CrossDexImplV2) PackCheckTickSizeRoles(account common.Address) []byte {
	enc, err := crossDexImplV2.abi.Pack("checkTickSizeRoles", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCreateMarket is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x786b0b11.
//
// Solidity: function createMarket(address _owner, address quote, address feeCollector, bytes data, string message) returns(address)
func (crossDexImplV2 *CrossDexImplV2) PackCreateMarket(owner common.Address, quote common.Address, feeCollector common.Address, data []byte, message string) []byte {
	enc, err := crossDexImplV2.abi.Pack("createMarket", owner, quote, feeCollector, data, message)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackCreateMarket is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x786b0b11.
//
// Solidity: function createMarket(address _owner, address quote, address feeCollector, bytes data, string message) returns(address)
func (crossDexImplV2 *CrossDexImplV2) UnpackCreateMarket(data []byte) (common.Address, error) {
	out, err := crossDexImplV2.abi.Unpack("createMarket", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x31ed1bcc.
//
// Solidity: function initialize(address _owner, address _routerImpl, uint256 _findPrevPriceCount, uint256 _maxMatchCount, uint256 _cancelLimit, address _marketImpl, address _pairImpl, address _tickSizeSetter) returns()
func (crossDexImplV2 *CrossDexImplV2) PackInitialize(owner common.Address, routerImpl common.Address, findPrevPriceCount *big.Int, maxMatchCount *big.Int, cancelLimit *big.Int, marketImpl common.Address, pairImpl common.Address, tickSizeSetter common.Address) []byte {
	enc, err := crossDexImplV2.abi.Pack("initialize", owner, routerImpl, findPrevPriceCount, maxMatchCount, cancelLimit, marketImpl, pairImpl, tickSizeSetter)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackIsMarket is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6ec934da.
//
// Solidity: function isMarket(address market) view returns(bool)
func (crossDexImplV2 *CrossDexImplV2) PackIsMarket(market common.Address) []byte {
	enc, err := crossDexImplV2.abi.Pack("isMarket", market)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsMarket is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6ec934da.
//
// Solidity: function isMarket(address market) view returns(bool)
func (crossDexImplV2 *CrossDexImplV2) UnpackIsMarket(data []byte) (bool, error) {
	out, err := crossDexImplV2.abi.Unpack("isMarket", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackMarketImpl is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x34eaeeb9.
//
// Solidity: function marketImpl() view returns(address)
func (crossDexImplV2 *CrossDexImplV2) PackMarketImpl() []byte {
	enc, err := crossDexImplV2.abi.Pack("marketImpl")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackMarketImpl is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x34eaeeb9.
//
// Solidity: function marketImpl() view returns(address)
func (crossDexImplV2 *CrossDexImplV2) UnpackMarketImpl(data []byte) (common.Address, error) {
	out, err := crossDexImplV2.abi.Unpack("marketImpl", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (crossDexImplV2 *CrossDexImplV2) PackOwner() []byte {
	enc, err := crossDexImplV2.abi.Pack("owner")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (crossDexImplV2 *CrossDexImplV2) UnpackOwner(data []byte) (common.Address, error) {
	out, err := crossDexImplV2.abi.Unpack("owner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackPairCreated is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe9f7206b.
//
// Solidity: function pairCreated(address pair) returns()
func (crossDexImplV2 *CrossDexImplV2) PackPairCreated(pair common.Address) []byte {
	enc, err := crossDexImplV2.abi.Pack("pairCreated", pair)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackPairImpl is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x35f7d456.
//
// Solidity: function pairImpl() view returns(address)
func (crossDexImplV2 *CrossDexImplV2) PackPairImpl() []byte {
	enc, err := crossDexImplV2.abi.Pack("pairImpl")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPairImpl is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x35f7d456.
//
// Solidity: function pairImpl() view returns(address)
func (crossDexImplV2 *CrossDexImplV2) UnpackPairImpl(data []byte) (common.Address, error) {
	out, err := crossDexImplV2.abi.Unpack("pairImpl", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackPairToMarket is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x08270573.
//
// Solidity: function pairToMarket(address pair) view returns(address)
func (crossDexImplV2 *CrossDexImplV2) PackPairToMarket(pair common.Address) []byte {
	enc, err := crossDexImplV2.abi.Pack("pairToMarket", pair)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPairToMarket is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x08270573.
//
// Solidity: function pairToMarket(address pair) view returns(address)
func (crossDexImplV2 *CrossDexImplV2) UnpackPairToMarket(data []byte) (common.Address, error) {
	out, err := crossDexImplV2.abi.Unpack("pairToMarket", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackProxiableUUID is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (crossDexImplV2 *CrossDexImplV2) PackProxiableUUID() []byte {
	enc, err := crossDexImplV2.abi.Pack("proxiableUUID")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackProxiableUUID is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (crossDexImplV2 *CrossDexImplV2) UnpackProxiableUUID(data []byte) ([32]byte, error) {
	out, err := crossDexImplV2.abi.Unpack("proxiableUUID", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackReinitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa9d951a3.
//
// Solidity: function reinitialize(address _marketImpl, address _pairImpl) returns()
func (crossDexImplV2 *CrossDexImplV2) PackReinitialize(marketImpl common.Address, pairImpl common.Address) []byte {
	enc, err := crossDexImplV2.abi.Pack("reinitialize", marketImpl, pairImpl)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRenounceOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (crossDexImplV2 *CrossDexImplV2) PackRenounceOwnership() []byte {
	enc, err := crossDexImplV2.abi.Pack("renounceOwnership")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetMarketImpl is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe492c270.
//
// Solidity: function setMarketImpl(address _marketImpl) returns()
func (crossDexImplV2 *CrossDexImplV2) PackSetMarketImpl(marketImpl common.Address) []byte {
	enc, err := crossDexImplV2.abi.Pack("setMarketImpl", marketImpl)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetPairImpl is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3cc047aa.
//
// Solidity: function setPairImpl(address _pairImpl) returns()
func (crossDexImplV2 *CrossDexImplV2) PackSetPairImpl(pairImpl common.Address) []byte {
	enc, err := crossDexImplV2.abi.Pack("setPairImpl", pairImpl)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetTickSizeSetter is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x98715ee7.
//
// Solidity: function setTickSizeSetter(address setter) returns()
func (crossDexImplV2 *CrossDexImplV2) PackSetTickSizeSetter(setter common.Address) []byte {
	enc, err := crossDexImplV2.abi.Pack("setTickSizeSetter", setter)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackTickSizeSetter is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7f764a85.
//
// Solidity: function tickSizeSetter() view returns(address)
func (crossDexImplV2 *CrossDexImplV2) PackTickSizeSetter() []byte {
	enc, err := crossDexImplV2.abi.Pack("tickSizeSetter")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTickSizeSetter is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7f764a85.
//
// Solidity: function tickSizeSetter() view returns(address)
func (crossDexImplV2 *CrossDexImplV2) UnpackTickSizeSetter(data []byte) (common.Address, error) {
	out, err := crossDexImplV2.abi.Unpack("tickSizeSetter", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackTransferOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (crossDexImplV2 *CrossDexImplV2) PackTransferOwnership(newOwner common.Address) []byte {
	enc, err := crossDexImplV2.abi.Pack("transferOwnership", newOwner)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackUpgradeToAndCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (crossDexImplV2 *CrossDexImplV2) PackUpgradeToAndCall(newImplementation common.Address, data []byte) []byte {
	enc, err := crossDexImplV2.abi.Pack("upgradeToAndCall", newImplementation, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// CrossDexImplV2Initialized represents a Initialized event raised by the CrossDexImplV2 contract.
type CrossDexImplV2Initialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const CrossDexImplV2InitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (CrossDexImplV2Initialized) ContractEventName() string {
	return CrossDexImplV2InitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (crossDexImplV2 *CrossDexImplV2) UnpackInitializedEvent(log *types.Log) (*CrossDexImplV2Initialized, error) {
	event := "Initialized"
	if log.Topics[0] != crossDexImplV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CrossDexImplV2Initialized)
	if len(log.Data) > 0 {
		if err := crossDexImplV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range crossDexImplV2.abi.Events[event].Inputs {
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

// CrossDexImplV2MarketCreated represents a MarketCreated event raised by the CrossDexImplV2 contract.
type CrossDexImplV2MarketCreated struct {
	Quote        common.Address
	Market       common.Address
	Owner        common.Address
	FeeCollector common.Address
	Message      string
	Raw          *types.Log // Blockchain specific contextual infos
}

const CrossDexImplV2MarketCreatedEventName = "MarketCreated"

// ContractEventName returns the user-defined event name.
func (CrossDexImplV2MarketCreated) ContractEventName() string {
	return CrossDexImplV2MarketCreatedEventName
}

// UnpackMarketCreatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MarketCreated(address indexed quote, address indexed market, address indexed owner, address fee_collector, string message)
func (crossDexImplV2 *CrossDexImplV2) UnpackMarketCreatedEvent(log *types.Log) (*CrossDexImplV2MarketCreated, error) {
	event := "MarketCreated"
	if log.Topics[0] != crossDexImplV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CrossDexImplV2MarketCreated)
	if len(log.Data) > 0 {
		if err := crossDexImplV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range crossDexImplV2.abi.Events[event].Inputs {
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

// CrossDexImplV2MarketImplSet represents a MarketImplSet event raised by the CrossDexImplV2 contract.
type CrossDexImplV2MarketImplSet struct {
	Before  common.Address
	Current common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const CrossDexImplV2MarketImplSetEventName = "MarketImplSet"

// ContractEventName returns the user-defined event name.
func (CrossDexImplV2MarketImplSet) ContractEventName() string {
	return CrossDexImplV2MarketImplSetEventName
}

// UnpackMarketImplSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MarketImplSet(address indexed before, address indexed current)
func (crossDexImplV2 *CrossDexImplV2) UnpackMarketImplSetEvent(log *types.Log) (*CrossDexImplV2MarketImplSet, error) {
	event := "MarketImplSet"
	if log.Topics[0] != crossDexImplV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CrossDexImplV2MarketImplSet)
	if len(log.Data) > 0 {
		if err := crossDexImplV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range crossDexImplV2.abi.Events[event].Inputs {
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

// CrossDexImplV2OwnershipTransferred represents a OwnershipTransferred event raised by the CrossDexImplV2 contract.
type CrossDexImplV2OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const CrossDexImplV2OwnershipTransferredEventName = "OwnershipTransferred"

// ContractEventName returns the user-defined event name.
func (CrossDexImplV2OwnershipTransferred) ContractEventName() string {
	return CrossDexImplV2OwnershipTransferredEventName
}

// UnpackOwnershipTransferredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (crossDexImplV2 *CrossDexImplV2) UnpackOwnershipTransferredEvent(log *types.Log) (*CrossDexImplV2OwnershipTransferred, error) {
	event := "OwnershipTransferred"
	if log.Topics[0] != crossDexImplV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CrossDexImplV2OwnershipTransferred)
	if len(log.Data) > 0 {
		if err := crossDexImplV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range crossDexImplV2.abi.Events[event].Inputs {
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

// CrossDexImplV2PairImplSet represents a PairImplSet event raised by the CrossDexImplV2 contract.
type CrossDexImplV2PairImplSet struct {
	Before  common.Address
	Current common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const CrossDexImplV2PairImplSetEventName = "PairImplSet"

// ContractEventName returns the user-defined event name.
func (CrossDexImplV2PairImplSet) ContractEventName() string {
	return CrossDexImplV2PairImplSetEventName
}

// UnpackPairImplSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PairImplSet(address indexed before, address indexed current)
func (crossDexImplV2 *CrossDexImplV2) UnpackPairImplSetEvent(log *types.Log) (*CrossDexImplV2PairImplSet, error) {
	event := "PairImplSet"
	if log.Topics[0] != crossDexImplV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CrossDexImplV2PairImplSet)
	if len(log.Data) > 0 {
		if err := crossDexImplV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range crossDexImplV2.abi.Events[event].Inputs {
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

// CrossDexImplV2TickSizeSetterSet represents a TickSizeSetterSet event raised by the CrossDexImplV2 contract.
type CrossDexImplV2TickSizeSetterSet struct {
	Before  common.Address
	Current common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const CrossDexImplV2TickSizeSetterSetEventName = "TickSizeSetterSet"

// ContractEventName returns the user-defined event name.
func (CrossDexImplV2TickSizeSetterSet) ContractEventName() string {
	return CrossDexImplV2TickSizeSetterSetEventName
}

// UnpackTickSizeSetterSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TickSizeSetterSet(address indexed before, address indexed current)
func (crossDexImplV2 *CrossDexImplV2) UnpackTickSizeSetterSetEvent(log *types.Log) (*CrossDexImplV2TickSizeSetterSet, error) {
	event := "TickSizeSetterSet"
	if log.Topics[0] != crossDexImplV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CrossDexImplV2TickSizeSetterSet)
	if len(log.Data) > 0 {
		if err := crossDexImplV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range crossDexImplV2.abi.Events[event].Inputs {
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

// CrossDexImplV2Upgraded represents a Upgraded event raised by the CrossDexImplV2 contract.
type CrossDexImplV2Upgraded struct {
	Implementation common.Address
	Raw            *types.Log // Blockchain specific contextual infos
}

const CrossDexImplV2UpgradedEventName = "Upgraded"

// ContractEventName returns the user-defined event name.
func (CrossDexImplV2Upgraded) ContractEventName() string {
	return CrossDexImplV2UpgradedEventName
}

// UnpackUpgradedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Upgraded(address indexed implementation)
func (crossDexImplV2 *CrossDexImplV2) UnpackUpgradedEvent(log *types.Log) (*CrossDexImplV2Upgraded, error) {
	event := "Upgraded"
	if log.Topics[0] != crossDexImplV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CrossDexImplV2Upgraded)
	if len(log.Data) > 0 {
		if err := crossDexImplV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range crossDexImplV2.abi.Events[event].Inputs {
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
func (crossDexImplV2 *CrossDexImplV2) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], crossDexImplV2.abi.Errors["AddressEmptyCode"].ID.Bytes()[:4]) {
		return crossDexImplV2.UnpackAddressEmptyCodeError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossDexImplV2.abi.Errors["Create2EmptyBytecode"].ID.Bytes()[:4]) {
		return crossDexImplV2.UnpackCreate2EmptyBytecodeError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossDexImplV2.abi.Errors["CrossDexInitializeData"].ID.Bytes()[:4]) {
		return crossDexImplV2.UnpackCrossDexInitializeDataError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossDexImplV2.abi.Errors["CrossDexInvalidMarketAddress"].ID.Bytes()[:4]) {
		return crossDexImplV2.UnpackCrossDexInvalidMarketAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossDexImplV2.abi.Errors["CrossDexInvalidTickSizeSetter"].ID.Bytes()[:4]) {
		return crossDexImplV2.UnpackCrossDexInvalidTickSizeSetterError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossDexImplV2.abi.Errors["CrossDexUnauthorizedChangeTickSizes"].ID.Bytes()[:4]) {
		return crossDexImplV2.UnpackCrossDexUnauthorizedChangeTickSizesError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossDexImplV2.abi.Errors["ERC1967InvalidImplementation"].ID.Bytes()[:4]) {
		return crossDexImplV2.UnpackERC1967InvalidImplementationError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossDexImplV2.abi.Errors["ERC1967NonPayable"].ID.Bytes()[:4]) {
		return crossDexImplV2.UnpackERC1967NonPayableError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossDexImplV2.abi.Errors["FailedCall"].ID.Bytes()[:4]) {
		return crossDexImplV2.UnpackFailedCallError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossDexImplV2.abi.Errors["FailedDeployment"].ID.Bytes()[:4]) {
		return crossDexImplV2.UnpackFailedDeploymentError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossDexImplV2.abi.Errors["InsufficientBalance"].ID.Bytes()[:4]) {
		return crossDexImplV2.UnpackInsufficientBalanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossDexImplV2.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return crossDexImplV2.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossDexImplV2.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return crossDexImplV2.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossDexImplV2.abi.Errors["OwnableInvalidOwner"].ID.Bytes()[:4]) {
		return crossDexImplV2.UnpackOwnableInvalidOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossDexImplV2.abi.Errors["OwnableUnauthorizedAccount"].ID.Bytes()[:4]) {
		return crossDexImplV2.UnpackOwnableUnauthorizedAccountError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossDexImplV2.abi.Errors["UUPSUnauthorizedCallContext"].ID.Bytes()[:4]) {
		return crossDexImplV2.UnpackUUPSUnauthorizedCallContextError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossDexImplV2.abi.Errors["UUPSUnsupportedProxiableUUID"].ID.Bytes()[:4]) {
		return crossDexImplV2.UnpackUUPSUnsupportedProxiableUUIDError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// CrossDexImplV2AddressEmptyCode represents a AddressEmptyCode error raised by the CrossDexImplV2 contract.
type CrossDexImplV2AddressEmptyCode struct {
	Target common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressEmptyCode(address target)
func CrossDexImplV2AddressEmptyCodeErrorID() common.Hash {
	return common.HexToHash("0x9996b315c842ff135b8fc4a08ad5df1c344efbc03d2687aecc0678050d2aac89")
}

// UnpackAddressEmptyCodeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressEmptyCode(address target)
func (crossDexImplV2 *CrossDexImplV2) UnpackAddressEmptyCodeError(raw []byte) (*CrossDexImplV2AddressEmptyCode, error) {
	out := new(CrossDexImplV2AddressEmptyCode)
	if err := crossDexImplV2.abi.UnpackIntoInterface(out, "AddressEmptyCode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossDexImplV2Create2EmptyBytecode represents a Create2EmptyBytecode error raised by the CrossDexImplV2 contract.
type CrossDexImplV2Create2EmptyBytecode struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error Create2EmptyBytecode()
func CrossDexImplV2Create2EmptyBytecodeErrorID() common.Hash {
	return common.HexToHash("0x4ca249dcffe41558ef8b961d71c905e4fa4317a1663f377b9610642e4e0abdb6")
}

// UnpackCreate2EmptyBytecodeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error Create2EmptyBytecode()
func (crossDexImplV2 *CrossDexImplV2) UnpackCreate2EmptyBytecodeError(raw []byte) (*CrossDexImplV2Create2EmptyBytecode, error) {
	out := new(CrossDexImplV2Create2EmptyBytecode)
	if err := crossDexImplV2.abi.UnpackIntoInterface(out, "Create2EmptyBytecode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossDexImplV2CrossDexInitializeData represents a CrossDexInitializeData error raised by the CrossDexImplV2 contract.
type CrossDexImplV2CrossDexInitializeData struct {
	Arg0 [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error CrossDexInitializeData(bytes32 arg0)
func CrossDexImplV2CrossDexInitializeDataErrorID() common.Hash {
	return common.HexToHash("0xc53a07b94abfc701bef7b9c8d77ba8466e4369b8995a2d5796d69e7c3c2fcfc2")
}

// UnpackCrossDexInitializeDataError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error CrossDexInitializeData(bytes32 arg0)
func (crossDexImplV2 *CrossDexImplV2) UnpackCrossDexInitializeDataError(raw []byte) (*CrossDexImplV2CrossDexInitializeData, error) {
	out := new(CrossDexImplV2CrossDexInitializeData)
	if err := crossDexImplV2.abi.UnpackIntoInterface(out, "CrossDexInitializeData", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossDexImplV2CrossDexInvalidMarketAddress represents a CrossDexInvalidMarketAddress error raised by the CrossDexImplV2 contract.
type CrossDexImplV2CrossDexInvalidMarketAddress struct {
	Arg0 common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error CrossDexInvalidMarketAddress(address arg0)
func CrossDexImplV2CrossDexInvalidMarketAddressErrorID() common.Hash {
	return common.HexToHash("0x74be5d0cefe8d064e961c4e0e557948c66ec877bbec4e6cb767d44a87b72a332")
}

// UnpackCrossDexInvalidMarketAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error CrossDexInvalidMarketAddress(address arg0)
func (crossDexImplV2 *CrossDexImplV2) UnpackCrossDexInvalidMarketAddressError(raw []byte) (*CrossDexImplV2CrossDexInvalidMarketAddress, error) {
	out := new(CrossDexImplV2CrossDexInvalidMarketAddress)
	if err := crossDexImplV2.abi.UnpackIntoInterface(out, "CrossDexInvalidMarketAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossDexImplV2CrossDexInvalidTickSizeSetter represents a CrossDexInvalidTickSizeSetter error raised by the CrossDexImplV2 contract.
type CrossDexImplV2CrossDexInvalidTickSizeSetter struct {
	Current common.Address
	Input   common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error CrossDexInvalidTickSizeSetter(address current, address input)
func CrossDexImplV2CrossDexInvalidTickSizeSetterErrorID() common.Hash {
	return common.HexToHash("0x18ddd67ebded54806618f9d60db7e0282fa9ae4a9c1139ff7ffde791b19783d1")
}

// UnpackCrossDexInvalidTickSizeSetterError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error CrossDexInvalidTickSizeSetter(address current, address input)
func (crossDexImplV2 *CrossDexImplV2) UnpackCrossDexInvalidTickSizeSetterError(raw []byte) (*CrossDexImplV2CrossDexInvalidTickSizeSetter, error) {
	out := new(CrossDexImplV2CrossDexInvalidTickSizeSetter)
	if err := crossDexImplV2.abi.UnpackIntoInterface(out, "CrossDexInvalidTickSizeSetter", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossDexImplV2CrossDexUnauthorizedChangeTickSizes represents a CrossDexUnauthorizedChangeTickSizes error raised by the CrossDexImplV2 contract.
type CrossDexImplV2CrossDexUnauthorizedChangeTickSizes struct {
	Arg0 common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error CrossDexUnauthorizedChangeTickSizes(address arg0)
func CrossDexImplV2CrossDexUnauthorizedChangeTickSizesErrorID() common.Hash {
	return common.HexToHash("0xf42eaafb7cf2208fd186d90528cb751ff79bfcbd79260c8b1448289b8146376a")
}

// UnpackCrossDexUnauthorizedChangeTickSizesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error CrossDexUnauthorizedChangeTickSizes(address arg0)
func (crossDexImplV2 *CrossDexImplV2) UnpackCrossDexUnauthorizedChangeTickSizesError(raw []byte) (*CrossDexImplV2CrossDexUnauthorizedChangeTickSizes, error) {
	out := new(CrossDexImplV2CrossDexUnauthorizedChangeTickSizes)
	if err := crossDexImplV2.abi.UnpackIntoInterface(out, "CrossDexUnauthorizedChangeTickSizes", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossDexImplV2ERC1967InvalidImplementation represents a ERC1967InvalidImplementation error raised by the CrossDexImplV2 contract.
type CrossDexImplV2ERC1967InvalidImplementation struct {
	Implementation common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func CrossDexImplV2ERC1967InvalidImplementationErrorID() common.Hash {
	return common.HexToHash("0x4c9c8ce3ceb3130f17f7cdba48d89b5b0129f266a8bac114e6e315a41879b617")
}

// UnpackERC1967InvalidImplementationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func (crossDexImplV2 *CrossDexImplV2) UnpackERC1967InvalidImplementationError(raw []byte) (*CrossDexImplV2ERC1967InvalidImplementation, error) {
	out := new(CrossDexImplV2ERC1967InvalidImplementation)
	if err := crossDexImplV2.abi.UnpackIntoInterface(out, "ERC1967InvalidImplementation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossDexImplV2ERC1967NonPayable represents a ERC1967NonPayable error raised by the CrossDexImplV2 contract.
type CrossDexImplV2ERC1967NonPayable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967NonPayable()
func CrossDexImplV2ERC1967NonPayableErrorID() common.Hash {
	return common.HexToHash("0xb398979fa84f543c8e222f17890372c487baf85e062276c127fef521eea7224b")
}

// UnpackERC1967NonPayableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967NonPayable()
func (crossDexImplV2 *CrossDexImplV2) UnpackERC1967NonPayableError(raw []byte) (*CrossDexImplV2ERC1967NonPayable, error) {
	out := new(CrossDexImplV2ERC1967NonPayable)
	if err := crossDexImplV2.abi.UnpackIntoInterface(out, "ERC1967NonPayable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossDexImplV2FailedCall represents a FailedCall error raised by the CrossDexImplV2 contract.
type CrossDexImplV2FailedCall struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedCall()
func CrossDexImplV2FailedCallErrorID() common.Hash {
	return common.HexToHash("0xd6bda27508c0fb6d8a39b4b122878dab26f731a7d4e4abe711dd3731899052a4")
}

// UnpackFailedCallError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedCall()
func (crossDexImplV2 *CrossDexImplV2) UnpackFailedCallError(raw []byte) (*CrossDexImplV2FailedCall, error) {
	out := new(CrossDexImplV2FailedCall)
	if err := crossDexImplV2.abi.UnpackIntoInterface(out, "FailedCall", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossDexImplV2FailedDeployment represents a FailedDeployment error raised by the CrossDexImplV2 contract.
type CrossDexImplV2FailedDeployment struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedDeployment()
func CrossDexImplV2FailedDeploymentErrorID() common.Hash {
	return common.HexToHash("0xb06ebf3d5067824a3fe5d5ba19471e035a7de6c88dac362c77b162830a5b9093")
}

// UnpackFailedDeploymentError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedDeployment()
func (crossDexImplV2 *CrossDexImplV2) UnpackFailedDeploymentError(raw []byte) (*CrossDexImplV2FailedDeployment, error) {
	out := new(CrossDexImplV2FailedDeployment)
	if err := crossDexImplV2.abi.UnpackIntoInterface(out, "FailedDeployment", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossDexImplV2InsufficientBalance represents a InsufficientBalance error raised by the CrossDexImplV2 contract.
type CrossDexImplV2InsufficientBalance struct {
	Balance *big.Int
	Needed  *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InsufficientBalance(uint256 balance, uint256 needed)
func CrossDexImplV2InsufficientBalanceErrorID() common.Hash {
	return common.HexToHash("0xcf4791818fba6e019216eb4864093b4947f674afada5d305e57d598b641dad1d")
}

// UnpackInsufficientBalanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InsufficientBalance(uint256 balance, uint256 needed)
func (crossDexImplV2 *CrossDexImplV2) UnpackInsufficientBalanceError(raw []byte) (*CrossDexImplV2InsufficientBalance, error) {
	out := new(CrossDexImplV2InsufficientBalance)
	if err := crossDexImplV2.abi.UnpackIntoInterface(out, "InsufficientBalance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossDexImplV2InvalidInitialization represents a InvalidInitialization error raised by the CrossDexImplV2 contract.
type CrossDexImplV2InvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func CrossDexImplV2InvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (crossDexImplV2 *CrossDexImplV2) UnpackInvalidInitializationError(raw []byte) (*CrossDexImplV2InvalidInitialization, error) {
	out := new(CrossDexImplV2InvalidInitialization)
	if err := crossDexImplV2.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossDexImplV2NotInitializing represents a NotInitializing error raised by the CrossDexImplV2 contract.
type CrossDexImplV2NotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func CrossDexImplV2NotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (crossDexImplV2 *CrossDexImplV2) UnpackNotInitializingError(raw []byte) (*CrossDexImplV2NotInitializing, error) {
	out := new(CrossDexImplV2NotInitializing)
	if err := crossDexImplV2.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossDexImplV2OwnableInvalidOwner represents a OwnableInvalidOwner error raised by the CrossDexImplV2 contract.
type CrossDexImplV2OwnableInvalidOwner struct {
	Owner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OwnableInvalidOwner(address owner)
func CrossDexImplV2OwnableInvalidOwnerErrorID() common.Hash {
	return common.HexToHash("0x1e4fbdf7f3ef8bcaa855599e3abf48b232380f183f08f6f813d9ffa5bd585188")
}

// UnpackOwnableInvalidOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OwnableInvalidOwner(address owner)
func (crossDexImplV2 *CrossDexImplV2) UnpackOwnableInvalidOwnerError(raw []byte) (*CrossDexImplV2OwnableInvalidOwner, error) {
	out := new(CrossDexImplV2OwnableInvalidOwner)
	if err := crossDexImplV2.abi.UnpackIntoInterface(out, "OwnableInvalidOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossDexImplV2OwnableUnauthorizedAccount represents a OwnableUnauthorizedAccount error raised by the CrossDexImplV2 contract.
type CrossDexImplV2OwnableUnauthorizedAccount struct {
	Account common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OwnableUnauthorizedAccount(address account)
func CrossDexImplV2OwnableUnauthorizedAccountErrorID() common.Hash {
	return common.HexToHash("0x118cdaa7a341953d1887a2245fd6665d741c67c8c50581daa59e1d03373fa188")
}

// UnpackOwnableUnauthorizedAccountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OwnableUnauthorizedAccount(address account)
func (crossDexImplV2 *CrossDexImplV2) UnpackOwnableUnauthorizedAccountError(raw []byte) (*CrossDexImplV2OwnableUnauthorizedAccount, error) {
	out := new(CrossDexImplV2OwnableUnauthorizedAccount)
	if err := crossDexImplV2.abi.UnpackIntoInterface(out, "OwnableUnauthorizedAccount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossDexImplV2UUPSUnauthorizedCallContext represents a UUPSUnauthorizedCallContext error raised by the CrossDexImplV2 contract.
type CrossDexImplV2UUPSUnauthorizedCallContext struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnauthorizedCallContext()
func CrossDexImplV2UUPSUnauthorizedCallContextErrorID() common.Hash {
	return common.HexToHash("0xe07c8dba242a06571ac65fe4bbe20522c9fb111cb33599b799ff8039c1ed18f4")
}

// UnpackUUPSUnauthorizedCallContextError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnauthorizedCallContext()
func (crossDexImplV2 *CrossDexImplV2) UnpackUUPSUnauthorizedCallContextError(raw []byte) (*CrossDexImplV2UUPSUnauthorizedCallContext, error) {
	out := new(CrossDexImplV2UUPSUnauthorizedCallContext)
	if err := crossDexImplV2.abi.UnpackIntoInterface(out, "UUPSUnauthorizedCallContext", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossDexImplV2UUPSUnsupportedProxiableUUID represents a UUPSUnsupportedProxiableUUID error raised by the CrossDexImplV2 contract.
type CrossDexImplV2UUPSUnsupportedProxiableUUID struct {
	Slot [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func CrossDexImplV2UUPSUnsupportedProxiableUUIDErrorID() common.Hash {
	return common.HexToHash("0xaa1d49a4c084bfa9aeeee2a0be65267a7f19ba7e1476b114dac513d2c14cb563")
}

// UnpackUUPSUnsupportedProxiableUUIDError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func (crossDexImplV2 *CrossDexImplV2) UnpackUUPSUnsupportedProxiableUUIDError(raw []byte) (*CrossDexImplV2UUPSUnsupportedProxiableUUID, error) {
	out := new(CrossDexImplV2UUPSUnsupportedProxiableUUID)
	if err := crossDexImplV2.abi.UnpackIntoInterface(out, "UUPSUnsupportedProxiableUUID", raw); err != nil {
		return nil, err
	}
	return out, nil
}
