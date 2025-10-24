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

// BuyBotMetaData contains all meta data concerning the BuyBot contract.
var BuyBotMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_minOrderAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_interval\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_buyer\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"NATIVE_COIN\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"buyMarket\",\"inputs\":[{\"name\":\"pair\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxMatchCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"buyer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"canBuyMarket\",\"inputs\":[{\"name\":\"pair\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"canBuy\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBalance\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"interval\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastBuyTime\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minOrderAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recipient\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"router\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRouter\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setBuyer\",\"inputs\":[{\"name\":\"_buyer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setInterval\",\"inputs\":[{\"name\":\"_interval\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinOrderAmount\",\"inputs\":[{\"name\":\"_minOrderAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRecipient\",\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawETH\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"BuyerSet\",\"inputs\":[{\"name\":\"before\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"current\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IntervalSet\",\"inputs\":[{\"name\":\"before\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"current\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MarketBuyExecuted\",\"inputs\":[{\"name\":\"pair\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"quoteToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"baseToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"quoteAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"executor\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinOrderAmountSet\",\"inputs\":[{\"name\":\"before\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"current\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RecipientSet\",\"inputs\":[{\"name\":\"before\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"current\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BuyBotInsufficientBalance\",\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minOrderAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"BuyBotIntervalNotPassed\",\"inputs\":[{\"name\":\"timeSinceLastBuy\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requiredInterval\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"BuyBotInvalidMinOrderAmount\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"BuyBotInvalidPair\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"BuyBotInvalidRouter\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"BuyBotUnauthorizedCaller\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x608060405234801561000f575f5ffd5b50604051611ab3380380611ab383398101604081905261002e9161023a565b856001600160a01b03811661005d57604051631e4fbdf760e01b81525f60048201526024015b60405180910390fd5b610066816101d0565b506001600160a01b03851661009957604051631561f27b60e21b81526001600160a01b0386166004820152602401610054565b835f036100bc576040516328bfc81960e11b815260048101859052602401610054565b600180546001600160a01b038088166001600160a01b0319928316179092556002869055600385905560058054858416908316179055600680549284169290911691909117905560405184905f907fd6d62b78d21fc4f5151029ca37079020e8a8815c5a488b28ee7499d4c2a19e55908290a360405183905f907f3fb677206c0b314c404bae3da94bee6bda11375c9fe986f266fb033d6cddbbd7908290a36040516001600160a01b038316905f907fc1416b5cdab50a9fbc872236e1aa54566c6deb40024e63a4b1737ecacf09d6f9908290a36040516001600160a01b038216905f907fe4eacf1aafd9e6b52cd356977b5c4ac221b6a852023bd3691b98600a9fb093d0908290a350505050505061029d565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b80516001600160a01b0381168114610235575f5ffd5b919050565b5f5f5f5f5f5f60c0878903121561024f575f5ffd5b6102588761021f565b95506102666020880161021f565b60408801516060890151919650945092506102836080880161021f565b915061029160a0880161021f565b90509295509295509295565b611809806102aa5f395ff3fe608060405260043610610140575f3560e01c80638da5cb5b116100bb578063f29f4d0b11610071578063f3fef3a311610057578063f3fef3a31461037c578063f887ea401461039b578063f8b2cb4f146103c7575f5ffd5b8063f29f4d0b14610348578063f2fde38b1461035d575f5ffd5b8063a3b8ef04116100a1578063a3b8ef04146102eb578063a3f09ad61461030a578063f14210a614610329575f5ffd5b80638da5cb5b146102ad578063947a36fb146102d6575f5ffd5b806346b62c4a1161011057806367ea88eb116100f657806367ea88eb1461024e578063715018a61461026d5780637150d8ae14610281575f5ffd5b806346b62c4a146101ff57806366d003ac14610222575f5ffd5b806304a411591461014b57806322a90082146101895780632aaa9628146101aa5780633bbed4a0146101e0575f5ffd5b3661014757005b5f5ffd5b348015610156575f5ffd5b5061015f600181565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b348015610194575f5ffd5b506101a86101a3366004611607565b6103e6565b005b3480156101b5575f5ffd5b506101c96101c436600461163f565b610421565b604080519215158352602083019190915201610180565b3480156101eb575f5ffd5b506101a86101fa366004611676565b6105fa565b34801561020a575f5ffd5b5061021460025481565b604051908152602001610180565b34801561022d575f5ffd5b5060055461015f9073ffffffffffffffffffffffffffffffffffffffff1681565b348015610259575f5ffd5b506101a8610268366004611698565b61068f565b348015610278575f5ffd5b506101a8610d15565b34801561028c575f5ffd5b5060065461015f9073ffffffffffffffffffffffffffffffffffffffff1681565b3480156102b8575f5ffd5b505f5473ffffffffffffffffffffffffffffffffffffffff1661015f565b3480156102e1575f5ffd5b5061021460035481565b3480156102f6575f5ffd5b506101a8610305366004611607565b610d28565b348015610315575f5ffd5b506101a8610324366004611676565b610d9f565b348015610334575f5ffd5b506101a8610343366004611607565b610e34565b348015610353575f5ffd5b5061021460045481565b348015610368575f5ffd5b506101a8610377366004611676565b610fd6565b348015610387575f5ffd5b506101a86103963660046116ca565b611039565b3480156103a6575f5ffd5b5060015461015f9073ffffffffffffffffffffffffffffffffffffffff1681565b3480156103d2575f5ffd5b506102146103e1366004611676565b611212565b6103ee6112e9565b6003546040518291907f3fb677206c0b314c404bae3da94bee6bda11375c9fe986f266fb033d6cddbbd7905f90a3600355565b5f8073ffffffffffffffffffffffffffffffffffffffff841661044857505f9050806105f3565b5f5473ffffffffffffffffffffffffffffffffffffffff84811691161480159061048d575060065473ffffffffffffffffffffffffffffffffffffffff848116911614155b1561049c57505f9050806105f3565b5f8473ffffffffffffffffffffffffffffffffffffffff1663c3f909d46040518163ffffffff1660e01b8152600401606060405180830381865afa1580156104e6573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061050a91906116f4565b80516040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015291925073ffffffffffffffffffffffffffffffffffffffff16906370a0823190602401602060405180830381865afa158015610576573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061059a9190611784565b91506002548210156105af575f9250506105f3565b5f6003541180156105c157505f600454115b156105ed575f600454426105d5919061179b565b90506003548110156105eb575f935050506105f3565b505b60019250505b9250929050565b6106026112e9565b60055460405173ffffffffffffffffffffffffffffffffffffffff8084169216907fc1416b5cdab50a9fbc872236e1aa54566c6deb40024e63a4b1737ecacf09d6f9905f90a3600580547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b5f5473ffffffffffffffffffffffffffffffffffffffff1633148015906106ce575060065473ffffffffffffffffffffffffffffffffffffffff163314155b1561070c576040517fd9ad42d70000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8316610771576040517f235aafe400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84166004820152602401610703565b5f60035411801561078357505f600454115b156107e5575f60045442610797919061179b565b90506003548110156107e3576003546040517f1e379906000000000000000000000000000000000000000000000000000000008152610703918391600401918252602082015260400190565b505b5f8373ffffffffffffffffffffffffffffffffffffffff1663c3f909d46040518163ffffffff1660e01b8152600401606060405180830381865afa15801561082f573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061085391906116f4565b805160208201516040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015292935090915f9073ffffffffffffffffffffffffffffffffffffffff8416906370a0823190602401602060405180830381865afa1580156108c9573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108ed9190611784565b90505f86156108fc57866108fe565b815b905060025481101561094a576002546040517f013fafe2000000000000000000000000000000000000000000000000000000008152610703918391600401918252602082015260400190565b8181111561098e576040517f013fafe20000000000000000000000000000000000000000000000000000000081526004810183905260248101829052604401610703565b6001546040517fdd62ed3e00000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff918216602482018190529183919087169063dd62ed3e90604401602060405180830381865afa158015610a07573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a2b9190611784565b1015610a7257610a7273ffffffffffffffffffffffffffffffffffffffff8616827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff61133b565b6001546040517f1e92008400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8b8116600483015260248201859052604482018a905290911690631e920084906064015f604051808303815f87803b158015610aea575f5ffd5b505af1158015610afc573d5f5f3e3d5ffd5b50506040805185815233602082015273ffffffffffffffffffffffffffffffffffffffff808916945089811693508d16917fc7b4d815bd0a3b8348577971cb55a9334c80a0b3ccdfe287b55b5a6320d8480a910160405180910390a44260045560055473ffffffffffffffffffffffffffffffffffffffff1615610d0a57478015610c4a576005546040515f9173ffffffffffffffffffffffffffffffffffffffff169083908381818185875af1925050503d805f8114610bd8576040519150601f19603f3d011682016040523d82523d5f602084013e610bdd565b606091505b5050905080610c48576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f455448207472616e73666572206661696c6564000000000000000000000000006044820152606401610703565b505b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201525f9073ffffffffffffffffffffffffffffffffffffffff8716906370a0823190602401602060405180830381865afa158015610cb4573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610cd89190611784565b90508015610d0757600554610d079073ffffffffffffffffffffffffffffffffffffffff88811691168361145b565b50505b505050505050505050565b610d1d6112e9565b610d265f61149e565b565b610d306112e9565b805f03610d6c576040517f517f903200000000000000000000000000000000000000000000000000000000815260048101829052602401610703565b6002546040518291907fd6d62b78d21fc4f5151029ca37079020e8a8815c5a488b28ee7499d4c2a19e55905f90a3600255565b610da76112e9565b60065460405173ffffffffffffffffffffffffffffffffffffffff8084169216907fe4eacf1aafd9e6b52cd356977b5c4ac221b6a852023bd3691b98600a9fb093d0905f90a3600680547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b610e3c6112e9565b475f8215610e4a5782610e4c565b815b905081811115610eb8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f496e73756666696369656e742062616c616e63650000000000000000000000006044820152606401610703565b5f805460405173ffffffffffffffffffffffffffffffffffffffff9091169083908381818185875af1925050503d805f8114610f0f576040519150601f19603f3d011682016040523d82523d5f602084013e610f14565b606091505b5050905080610f7f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f455448207472616e73666572206661696c6564000000000000000000000000006044820152606401610703565b5f5460405183815273ffffffffffffffffffffffffffffffffffffffff909116906001907fd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb9060200160405180910390a350505050565b610fde6112e9565b73ffffffffffffffffffffffffffffffffffffffff811661102d576040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081525f6004820152602401610703565b6110368161149e565b50565b6110416112e9565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015282905f9073ffffffffffffffffffffffffffffffffffffffff8316906370a0823190602401602060405180830381865afa1580156110ad573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110d19190611784565b90505f83156110e057836110e2565b815b90508181111561114e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f496e73756666696369656e742062616c616e63650000000000000000000000006044820152606401610703565b61118d61116f5f5473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff8516908361145b565b5f5473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167fd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb8360405161120391815260200190565b60405180910390a35050505050565b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff73ffffffffffffffffffffffffffffffffffffffff831601611257575047919050565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff8316906370a0823190602401602060405180830381865afa1580156112bf573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906112e39190611784565b92915050565b5f5473ffffffffffffffffffffffffffffffffffffffff163314610d26576040517f118cdaa7000000000000000000000000000000000000000000000000000000008152336004820152602401610703565b6040805173ffffffffffffffffffffffffffffffffffffffff8416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f095ea7b3000000000000000000000000000000000000000000000000000000001790526113c78482611512565b6114555760405173ffffffffffffffffffffffffffffffffffffffff84811660248301525f604483015261144b91869182169063095ea7b3906064015b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611568565b6114558482611568565b50505050565b60405173ffffffffffffffffffffffffffffffffffffffff83811660248301526044820183905261149991859182169063a9059cbb90606401611404565b505050565b5f805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b5f5f5f5f60205f8651602088015f8a5af192503d91505f51905082801561155e57508115611543578060011461155e565b5f8673ffffffffffffffffffffffffffffffffffffffff163b115b9695505050505050565b5f5f60205f8451602086015f885af180611587576040513d5f823e3d81fd5b50505f513d9150811561159e5780600114156115b8565b73ffffffffffffffffffffffffffffffffffffffff84163b155b15611455576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610703565b5f60208284031215611617575f5ffd5b5035919050565b73ffffffffffffffffffffffffffffffffffffffff81168114611036575f5ffd5b5f5f60408385031215611650575f5ffd5b823561165b8161161e565b9150602083013561166b8161161e565b809150509250929050565b5f60208284031215611686575f5ffd5b81356116918161161e565b9392505050565b5f5f5f606084860312156116aa575f5ffd5b83356116b58161161e565b95602085013595506040909401359392505050565b5f5f604083850312156116db575f5ffd5b82356116e68161161e565b946020939093013593505050565b5f6060828403128015611705575f5ffd5b506040516060810167ffffffffffffffff8111828210171561174e577f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604052825161175c8161161e565b8152602083015161176c8161161e565b60208201526040928301519281019290925250919050565b5f60208284031215611794575f5ffd5b5051919050565b818103818111156112e3577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffdfea264697066735822122030ee24c3294a175a78feb83719ea780a5974c6d8833acbd6d07eea6406a09fc164736f6c634300081c0033",
}

// BuyBotABI is the input ABI used to generate the binding from.
// Deprecated: Use BuyBotMetaData.ABI instead.
var BuyBotABI = BuyBotMetaData.ABI

// BuyBotBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BuyBotMetaData.Bin instead.
var BuyBotBin = BuyBotMetaData.Bin

// DeployBuyBot deploys a new Ethereum contract, binding an instance of BuyBot to it.
func DeployBuyBot(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address, _router common.Address, _minOrderAmount *big.Int, _interval *big.Int, _recipient common.Address, _buyer common.Address) (common.Address, *types.Transaction, *BuyBot, error) {
	parsed, err := BuyBotMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BuyBotBin), backend, _owner, _router, _minOrderAmount, _interval, _recipient, _buyer)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BuyBot{BuyBotCaller: BuyBotCaller{contract: contract}, BuyBotTransactor: BuyBotTransactor{contract: contract}, BuyBotFilterer: BuyBotFilterer{contract: contract}}, nil
}

// BuyBot is an auto generated Go binding around an Ethereum contract.
type BuyBot struct {
	BuyBotCaller     // Read-only binding to the contract
	BuyBotTransactor // Write-only binding to the contract
	BuyBotFilterer   // Log filterer for contract events
}

// BuyBotCaller is an auto generated read-only Go binding around an Ethereum contract.
type BuyBotCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BuyBotTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BuyBotTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BuyBotFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BuyBotFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BuyBotSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BuyBotSession struct {
	Contract     *BuyBot           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BuyBotCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BuyBotCallerSession struct {
	Contract *BuyBotCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BuyBotTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BuyBotTransactorSession struct {
	Contract     *BuyBotTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BuyBotRaw is an auto generated low-level Go binding around an Ethereum contract.
type BuyBotRaw struct {
	Contract *BuyBot // Generic contract binding to access the raw methods on
}

// BuyBotCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BuyBotCallerRaw struct {
	Contract *BuyBotCaller // Generic read-only contract binding to access the raw methods on
}

// BuyBotTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BuyBotTransactorRaw struct {
	Contract *BuyBotTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBuyBot creates a new instance of BuyBot, bound to a specific deployed contract.
func NewBuyBot(address common.Address, backend bind.ContractBackend) (*BuyBot, error) {
	contract, err := bindBuyBot(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BuyBot{BuyBotCaller: BuyBotCaller{contract: contract}, BuyBotTransactor: BuyBotTransactor{contract: contract}, BuyBotFilterer: BuyBotFilterer{contract: contract}}, nil
}

// NewBuyBotCaller creates a new read-only instance of BuyBot, bound to a specific deployed contract.
func NewBuyBotCaller(address common.Address, caller bind.ContractCaller) (*BuyBotCaller, error) {
	contract, err := bindBuyBot(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BuyBotCaller{contract: contract}, nil
}

// NewBuyBotTransactor creates a new write-only instance of BuyBot, bound to a specific deployed contract.
func NewBuyBotTransactor(address common.Address, transactor bind.ContractTransactor) (*BuyBotTransactor, error) {
	contract, err := bindBuyBot(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BuyBotTransactor{contract: contract}, nil
}

// NewBuyBotFilterer creates a new log filterer instance of BuyBot, bound to a specific deployed contract.
func NewBuyBotFilterer(address common.Address, filterer bind.ContractFilterer) (*BuyBotFilterer, error) {
	contract, err := bindBuyBot(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BuyBotFilterer{contract: contract}, nil
}

// bindBuyBot binds a generic wrapper to an already deployed contract.
func bindBuyBot(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BuyBotMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BuyBot *BuyBotRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BuyBot.Contract.BuyBotCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BuyBot *BuyBotRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BuyBot.Contract.BuyBotTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BuyBot *BuyBotRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BuyBot.Contract.BuyBotTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BuyBot *BuyBotCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BuyBot.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BuyBot *BuyBotTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BuyBot.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BuyBot *BuyBotTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BuyBot.Contract.contract.Transact(opts, method, params...)
}

// NATIVECOIN is a free data retrieval call binding the contract method 0x04a41159.
//
// Solidity: function NATIVE_COIN() view returns(address)
func (_BuyBot *BuyBotCaller) NATIVECOIN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BuyBot.contract.Call(opts, &out, "NATIVE_COIN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NATIVECOIN is a free data retrieval call binding the contract method 0x04a41159.
//
// Solidity: function NATIVE_COIN() view returns(address)
func (_BuyBot *BuyBotSession) NATIVECOIN() (common.Address, error) {
	return _BuyBot.Contract.NATIVECOIN(&_BuyBot.CallOpts)
}

// NATIVECOIN is a free data retrieval call binding the contract method 0x04a41159.
//
// Solidity: function NATIVE_COIN() view returns(address)
func (_BuyBot *BuyBotCallerSession) NATIVECOIN() (common.Address, error) {
	return _BuyBot.Contract.NATIVECOIN(&_BuyBot.CallOpts)
}

// Buyer is a free data retrieval call binding the contract method 0x7150d8ae.
//
// Solidity: function buyer() view returns(address)
func (_BuyBot *BuyBotCaller) Buyer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BuyBot.contract.Call(opts, &out, "buyer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Buyer is a free data retrieval call binding the contract method 0x7150d8ae.
//
// Solidity: function buyer() view returns(address)
func (_BuyBot *BuyBotSession) Buyer() (common.Address, error) {
	return _BuyBot.Contract.Buyer(&_BuyBot.CallOpts)
}

// Buyer is a free data retrieval call binding the contract method 0x7150d8ae.
//
// Solidity: function buyer() view returns(address)
func (_BuyBot *BuyBotCallerSession) Buyer() (common.Address, error) {
	return _BuyBot.Contract.Buyer(&_BuyBot.CallOpts)
}

// CanBuyMarket is a free data retrieval call binding the contract method 0x2aaa9628.
//
// Solidity: function canBuyMarket(address pair, address caller) view returns(bool canBuy, uint256 balance)
func (_BuyBot *BuyBotCaller) CanBuyMarket(opts *bind.CallOpts, pair common.Address, caller common.Address) (struct {
	CanBuy  bool
	Balance *big.Int
}, error) {
	var out []interface{}
	err := _BuyBot.contract.Call(opts, &out, "canBuyMarket", pair, caller)

	outstruct := new(struct {
		CanBuy  bool
		Balance *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CanBuy = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Balance = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CanBuyMarket is a free data retrieval call binding the contract method 0x2aaa9628.
//
// Solidity: function canBuyMarket(address pair, address caller) view returns(bool canBuy, uint256 balance)
func (_BuyBot *BuyBotSession) CanBuyMarket(pair common.Address, caller common.Address) (struct {
	CanBuy  bool
	Balance *big.Int
}, error) {
	return _BuyBot.Contract.CanBuyMarket(&_BuyBot.CallOpts, pair, caller)
}

// CanBuyMarket is a free data retrieval call binding the contract method 0x2aaa9628.
//
// Solidity: function canBuyMarket(address pair, address caller) view returns(bool canBuy, uint256 balance)
func (_BuyBot *BuyBotCallerSession) CanBuyMarket(pair common.Address, caller common.Address) (struct {
	CanBuy  bool
	Balance *big.Int
}, error) {
	return _BuyBot.Contract.CanBuyMarket(&_BuyBot.CallOpts, pair, caller)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address token) view returns(uint256 balance)
func (_BuyBot *BuyBotCaller) GetBalance(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BuyBot.contract.Call(opts, &out, "getBalance", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address token) view returns(uint256 balance)
func (_BuyBot *BuyBotSession) GetBalance(token common.Address) (*big.Int, error) {
	return _BuyBot.Contract.GetBalance(&_BuyBot.CallOpts, token)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address token) view returns(uint256 balance)
func (_BuyBot *BuyBotCallerSession) GetBalance(token common.Address) (*big.Int, error) {
	return _BuyBot.Contract.GetBalance(&_BuyBot.CallOpts, token)
}

// Interval is a free data retrieval call binding the contract method 0x947a36fb.
//
// Solidity: function interval() view returns(uint256)
func (_BuyBot *BuyBotCaller) Interval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BuyBot.contract.Call(opts, &out, "interval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Interval is a free data retrieval call binding the contract method 0x947a36fb.
//
// Solidity: function interval() view returns(uint256)
func (_BuyBot *BuyBotSession) Interval() (*big.Int, error) {
	return _BuyBot.Contract.Interval(&_BuyBot.CallOpts)
}

// Interval is a free data retrieval call binding the contract method 0x947a36fb.
//
// Solidity: function interval() view returns(uint256)
func (_BuyBot *BuyBotCallerSession) Interval() (*big.Int, error) {
	return _BuyBot.Contract.Interval(&_BuyBot.CallOpts)
}

// LastBuyTime is a free data retrieval call binding the contract method 0xf29f4d0b.
//
// Solidity: function lastBuyTime() view returns(uint256)
func (_BuyBot *BuyBotCaller) LastBuyTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BuyBot.contract.Call(opts, &out, "lastBuyTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastBuyTime is a free data retrieval call binding the contract method 0xf29f4d0b.
//
// Solidity: function lastBuyTime() view returns(uint256)
func (_BuyBot *BuyBotSession) LastBuyTime() (*big.Int, error) {
	return _BuyBot.Contract.LastBuyTime(&_BuyBot.CallOpts)
}

// LastBuyTime is a free data retrieval call binding the contract method 0xf29f4d0b.
//
// Solidity: function lastBuyTime() view returns(uint256)
func (_BuyBot *BuyBotCallerSession) LastBuyTime() (*big.Int, error) {
	return _BuyBot.Contract.LastBuyTime(&_BuyBot.CallOpts)
}

// MinOrderAmount is a free data retrieval call binding the contract method 0x46b62c4a.
//
// Solidity: function minOrderAmount() view returns(uint256)
func (_BuyBot *BuyBotCaller) MinOrderAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BuyBot.contract.Call(opts, &out, "minOrderAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinOrderAmount is a free data retrieval call binding the contract method 0x46b62c4a.
//
// Solidity: function minOrderAmount() view returns(uint256)
func (_BuyBot *BuyBotSession) MinOrderAmount() (*big.Int, error) {
	return _BuyBot.Contract.MinOrderAmount(&_BuyBot.CallOpts)
}

// MinOrderAmount is a free data retrieval call binding the contract method 0x46b62c4a.
//
// Solidity: function minOrderAmount() view returns(uint256)
func (_BuyBot *BuyBotCallerSession) MinOrderAmount() (*big.Int, error) {
	return _BuyBot.Contract.MinOrderAmount(&_BuyBot.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BuyBot *BuyBotCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BuyBot.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BuyBot *BuyBotSession) Owner() (common.Address, error) {
	return _BuyBot.Contract.Owner(&_BuyBot.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BuyBot *BuyBotCallerSession) Owner() (common.Address, error) {
	return _BuyBot.Contract.Owner(&_BuyBot.CallOpts)
}

// Recipient is a free data retrieval call binding the contract method 0x66d003ac.
//
// Solidity: function recipient() view returns(address)
func (_BuyBot *BuyBotCaller) Recipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BuyBot.contract.Call(opts, &out, "recipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Recipient is a free data retrieval call binding the contract method 0x66d003ac.
//
// Solidity: function recipient() view returns(address)
func (_BuyBot *BuyBotSession) Recipient() (common.Address, error) {
	return _BuyBot.Contract.Recipient(&_BuyBot.CallOpts)
}

// Recipient is a free data retrieval call binding the contract method 0x66d003ac.
//
// Solidity: function recipient() view returns(address)
func (_BuyBot *BuyBotCallerSession) Recipient() (common.Address, error) {
	return _BuyBot.Contract.Recipient(&_BuyBot.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_BuyBot *BuyBotCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BuyBot.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_BuyBot *BuyBotSession) Router() (common.Address, error) {
	return _BuyBot.Contract.Router(&_BuyBot.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_BuyBot *BuyBotCallerSession) Router() (common.Address, error) {
	return _BuyBot.Contract.Router(&_BuyBot.CallOpts)
}

// BuyMarket is a paid mutator transaction binding the contract method 0x67ea88eb.
//
// Solidity: function buyMarket(address pair, uint256 amount, uint256 maxMatchCount) returns()
func (_BuyBot *BuyBotTransactor) BuyMarket(opts *bind.TransactOpts, pair common.Address, amount *big.Int, maxMatchCount *big.Int) (*types.Transaction, error) {
	return _BuyBot.contract.Transact(opts, "buyMarket", pair, amount, maxMatchCount)
}

// BuyMarket is a paid mutator transaction binding the contract method 0x67ea88eb.
//
// Solidity: function buyMarket(address pair, uint256 amount, uint256 maxMatchCount) returns()
func (_BuyBot *BuyBotSession) BuyMarket(pair common.Address, amount *big.Int, maxMatchCount *big.Int) (*types.Transaction, error) {
	return _BuyBot.Contract.BuyMarket(&_BuyBot.TransactOpts, pair, amount, maxMatchCount)
}

// BuyMarket is a paid mutator transaction binding the contract method 0x67ea88eb.
//
// Solidity: function buyMarket(address pair, uint256 amount, uint256 maxMatchCount) returns()
func (_BuyBot *BuyBotTransactorSession) BuyMarket(pair common.Address, amount *big.Int, maxMatchCount *big.Int) (*types.Transaction, error) {
	return _BuyBot.Contract.BuyMarket(&_BuyBot.TransactOpts, pair, amount, maxMatchCount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BuyBot *BuyBotTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BuyBot.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BuyBot *BuyBotSession) RenounceOwnership() (*types.Transaction, error) {
	return _BuyBot.Contract.RenounceOwnership(&_BuyBot.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BuyBot *BuyBotTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BuyBot.Contract.RenounceOwnership(&_BuyBot.TransactOpts)
}

// SetBuyer is a paid mutator transaction binding the contract method 0xa3f09ad6.
//
// Solidity: function setBuyer(address _buyer) returns()
func (_BuyBot *BuyBotTransactor) SetBuyer(opts *bind.TransactOpts, _buyer common.Address) (*types.Transaction, error) {
	return _BuyBot.contract.Transact(opts, "setBuyer", _buyer)
}

// SetBuyer is a paid mutator transaction binding the contract method 0xa3f09ad6.
//
// Solidity: function setBuyer(address _buyer) returns()
func (_BuyBot *BuyBotSession) SetBuyer(_buyer common.Address) (*types.Transaction, error) {
	return _BuyBot.Contract.SetBuyer(&_BuyBot.TransactOpts, _buyer)
}

// SetBuyer is a paid mutator transaction binding the contract method 0xa3f09ad6.
//
// Solidity: function setBuyer(address _buyer) returns()
func (_BuyBot *BuyBotTransactorSession) SetBuyer(_buyer common.Address) (*types.Transaction, error) {
	return _BuyBot.Contract.SetBuyer(&_BuyBot.TransactOpts, _buyer)
}

// SetInterval is a paid mutator transaction binding the contract method 0x22a90082.
//
// Solidity: function setInterval(uint256 _interval) returns()
func (_BuyBot *BuyBotTransactor) SetInterval(opts *bind.TransactOpts, _interval *big.Int) (*types.Transaction, error) {
	return _BuyBot.contract.Transact(opts, "setInterval", _interval)
}

// SetInterval is a paid mutator transaction binding the contract method 0x22a90082.
//
// Solidity: function setInterval(uint256 _interval) returns()
func (_BuyBot *BuyBotSession) SetInterval(_interval *big.Int) (*types.Transaction, error) {
	return _BuyBot.Contract.SetInterval(&_BuyBot.TransactOpts, _interval)
}

// SetInterval is a paid mutator transaction binding the contract method 0x22a90082.
//
// Solidity: function setInterval(uint256 _interval) returns()
func (_BuyBot *BuyBotTransactorSession) SetInterval(_interval *big.Int) (*types.Transaction, error) {
	return _BuyBot.Contract.SetInterval(&_BuyBot.TransactOpts, _interval)
}

// SetMinOrderAmount is a paid mutator transaction binding the contract method 0xa3b8ef04.
//
// Solidity: function setMinOrderAmount(uint256 _minOrderAmount) returns()
func (_BuyBot *BuyBotTransactor) SetMinOrderAmount(opts *bind.TransactOpts, _minOrderAmount *big.Int) (*types.Transaction, error) {
	return _BuyBot.contract.Transact(opts, "setMinOrderAmount", _minOrderAmount)
}

// SetMinOrderAmount is a paid mutator transaction binding the contract method 0xa3b8ef04.
//
// Solidity: function setMinOrderAmount(uint256 _minOrderAmount) returns()
func (_BuyBot *BuyBotSession) SetMinOrderAmount(_minOrderAmount *big.Int) (*types.Transaction, error) {
	return _BuyBot.Contract.SetMinOrderAmount(&_BuyBot.TransactOpts, _minOrderAmount)
}

// SetMinOrderAmount is a paid mutator transaction binding the contract method 0xa3b8ef04.
//
// Solidity: function setMinOrderAmount(uint256 _minOrderAmount) returns()
func (_BuyBot *BuyBotTransactorSession) SetMinOrderAmount(_minOrderAmount *big.Int) (*types.Transaction, error) {
	return _BuyBot.Contract.SetMinOrderAmount(&_BuyBot.TransactOpts, _minOrderAmount)
}

// SetRecipient is a paid mutator transaction binding the contract method 0x3bbed4a0.
//
// Solidity: function setRecipient(address _recipient) returns()
func (_BuyBot *BuyBotTransactor) SetRecipient(opts *bind.TransactOpts, _recipient common.Address) (*types.Transaction, error) {
	return _BuyBot.contract.Transact(opts, "setRecipient", _recipient)
}

// SetRecipient is a paid mutator transaction binding the contract method 0x3bbed4a0.
//
// Solidity: function setRecipient(address _recipient) returns()
func (_BuyBot *BuyBotSession) SetRecipient(_recipient common.Address) (*types.Transaction, error) {
	return _BuyBot.Contract.SetRecipient(&_BuyBot.TransactOpts, _recipient)
}

// SetRecipient is a paid mutator transaction binding the contract method 0x3bbed4a0.
//
// Solidity: function setRecipient(address _recipient) returns()
func (_BuyBot *BuyBotTransactorSession) SetRecipient(_recipient common.Address) (*types.Transaction, error) {
	return _BuyBot.Contract.SetRecipient(&_BuyBot.TransactOpts, _recipient)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BuyBot *BuyBotTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BuyBot.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BuyBot *BuyBotSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BuyBot.Contract.TransferOwnership(&_BuyBot.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BuyBot *BuyBotTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BuyBot.Contract.TransferOwnership(&_BuyBot.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address token, uint256 amount) returns()
func (_BuyBot *BuyBotTransactor) Withdraw(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BuyBot.contract.Transact(opts, "withdraw", token, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address token, uint256 amount) returns()
func (_BuyBot *BuyBotSession) Withdraw(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BuyBot.Contract.Withdraw(&_BuyBot.TransactOpts, token, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address token, uint256 amount) returns()
func (_BuyBot *BuyBotTransactorSession) Withdraw(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BuyBot.Contract.Withdraw(&_BuyBot.TransactOpts, token, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xf14210a6.
//
// Solidity: function withdrawETH(uint256 amount) returns()
func (_BuyBot *BuyBotTransactor) WithdrawETH(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _BuyBot.contract.Transact(opts, "withdrawETH", amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xf14210a6.
//
// Solidity: function withdrawETH(uint256 amount) returns()
func (_BuyBot *BuyBotSession) WithdrawETH(amount *big.Int) (*types.Transaction, error) {
	return _BuyBot.Contract.WithdrawETH(&_BuyBot.TransactOpts, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xf14210a6.
//
// Solidity: function withdrawETH(uint256 amount) returns()
func (_BuyBot *BuyBotTransactorSession) WithdrawETH(amount *big.Int) (*types.Transaction, error) {
	return _BuyBot.Contract.WithdrawETH(&_BuyBot.TransactOpts, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BuyBot *BuyBotTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BuyBot.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BuyBot *BuyBotSession) Receive() (*types.Transaction, error) {
	return _BuyBot.Contract.Receive(&_BuyBot.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BuyBot *BuyBotTransactorSession) Receive() (*types.Transaction, error) {
	return _BuyBot.Contract.Receive(&_BuyBot.TransactOpts)
}

// BuyBotBuyerSetIterator is returned from FilterBuyerSet and is used to iterate over the raw logs and unpacked data for BuyerSet events raised by the BuyBot contract.
type BuyBotBuyerSetIterator struct {
	Event *BuyBotBuyerSet // Event containing the contract specifics and raw log

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
func (it *BuyBotBuyerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuyBotBuyerSet)
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
		it.Event = new(BuyBotBuyerSet)
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
func (it *BuyBotBuyerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuyBotBuyerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuyBotBuyerSet represents a BuyerSet event raised by the BuyBot contract.
type BuyBotBuyerSet struct {
	Before  common.Address
	Current common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBuyerSet is a free log retrieval operation binding the contract event 0xe4eacf1aafd9e6b52cd356977b5c4ac221b6a852023bd3691b98600a9fb093d0.
//
// Solidity: event BuyerSet(address indexed before, address indexed current)
func (_BuyBot *BuyBotFilterer) FilterBuyerSet(opts *bind.FilterOpts, before []common.Address, current []common.Address) (*BuyBotBuyerSetIterator, error) {

	var beforeRule []interface{}
	for _, beforeItem := range before {
		beforeRule = append(beforeRule, beforeItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _BuyBot.contract.FilterLogs(opts, "BuyerSet", beforeRule, currentRule)
	if err != nil {
		return nil, err
	}
	return &BuyBotBuyerSetIterator{contract: _BuyBot.contract, event: "BuyerSet", logs: logs, sub: sub}, nil
}

// WatchBuyerSet is a free log subscription operation binding the contract event 0xe4eacf1aafd9e6b52cd356977b5c4ac221b6a852023bd3691b98600a9fb093d0.
//
// Solidity: event BuyerSet(address indexed before, address indexed current)
func (_BuyBot *BuyBotFilterer) WatchBuyerSet(opts *bind.WatchOpts, sink chan<- *BuyBotBuyerSet, before []common.Address, current []common.Address) (event.Subscription, error) {

	var beforeRule []interface{}
	for _, beforeItem := range before {
		beforeRule = append(beforeRule, beforeItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _BuyBot.contract.WatchLogs(opts, "BuyerSet", beforeRule, currentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuyBotBuyerSet)
				if err := _BuyBot.contract.UnpackLog(event, "BuyerSet", log); err != nil {
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

// ParseBuyerSet is a log parse operation binding the contract event 0xe4eacf1aafd9e6b52cd356977b5c4ac221b6a852023bd3691b98600a9fb093d0.
//
// Solidity: event BuyerSet(address indexed before, address indexed current)
func (_BuyBot *BuyBotFilterer) ParseBuyerSet(log types.Log) (*BuyBotBuyerSet, error) {
	event := new(BuyBotBuyerSet)
	if err := _BuyBot.contract.UnpackLog(event, "BuyerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BuyBotIntervalSetIterator is returned from FilterIntervalSet and is used to iterate over the raw logs and unpacked data for IntervalSet events raised by the BuyBot contract.
type BuyBotIntervalSetIterator struct {
	Event *BuyBotIntervalSet // Event containing the contract specifics and raw log

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
func (it *BuyBotIntervalSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuyBotIntervalSet)
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
		it.Event = new(BuyBotIntervalSet)
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
func (it *BuyBotIntervalSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuyBotIntervalSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuyBotIntervalSet represents a IntervalSet event raised by the BuyBot contract.
type BuyBotIntervalSet struct {
	Before  *big.Int
	Current *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterIntervalSet is a free log retrieval operation binding the contract event 0x3fb677206c0b314c404bae3da94bee6bda11375c9fe986f266fb033d6cddbbd7.
//
// Solidity: event IntervalSet(uint256 indexed before, uint256 indexed current)
func (_BuyBot *BuyBotFilterer) FilterIntervalSet(opts *bind.FilterOpts, before []*big.Int, current []*big.Int) (*BuyBotIntervalSetIterator, error) {

	var beforeRule []interface{}
	for _, beforeItem := range before {
		beforeRule = append(beforeRule, beforeItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _BuyBot.contract.FilterLogs(opts, "IntervalSet", beforeRule, currentRule)
	if err != nil {
		return nil, err
	}
	return &BuyBotIntervalSetIterator{contract: _BuyBot.contract, event: "IntervalSet", logs: logs, sub: sub}, nil
}

// WatchIntervalSet is a free log subscription operation binding the contract event 0x3fb677206c0b314c404bae3da94bee6bda11375c9fe986f266fb033d6cddbbd7.
//
// Solidity: event IntervalSet(uint256 indexed before, uint256 indexed current)
func (_BuyBot *BuyBotFilterer) WatchIntervalSet(opts *bind.WatchOpts, sink chan<- *BuyBotIntervalSet, before []*big.Int, current []*big.Int) (event.Subscription, error) {

	var beforeRule []interface{}
	for _, beforeItem := range before {
		beforeRule = append(beforeRule, beforeItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _BuyBot.contract.WatchLogs(opts, "IntervalSet", beforeRule, currentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuyBotIntervalSet)
				if err := _BuyBot.contract.UnpackLog(event, "IntervalSet", log); err != nil {
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

// ParseIntervalSet is a log parse operation binding the contract event 0x3fb677206c0b314c404bae3da94bee6bda11375c9fe986f266fb033d6cddbbd7.
//
// Solidity: event IntervalSet(uint256 indexed before, uint256 indexed current)
func (_BuyBot *BuyBotFilterer) ParseIntervalSet(log types.Log) (*BuyBotIntervalSet, error) {
	event := new(BuyBotIntervalSet)
	if err := _BuyBot.contract.UnpackLog(event, "IntervalSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BuyBotMarketBuyExecutedIterator is returned from FilterMarketBuyExecuted and is used to iterate over the raw logs and unpacked data for MarketBuyExecuted events raised by the BuyBot contract.
type BuyBotMarketBuyExecutedIterator struct {
	Event *BuyBotMarketBuyExecuted // Event containing the contract specifics and raw log

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
func (it *BuyBotMarketBuyExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuyBotMarketBuyExecuted)
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
		it.Event = new(BuyBotMarketBuyExecuted)
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
func (it *BuyBotMarketBuyExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuyBotMarketBuyExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuyBotMarketBuyExecuted represents a MarketBuyExecuted event raised by the BuyBot contract.
type BuyBotMarketBuyExecuted struct {
	Pair        common.Address
	QuoteToken  common.Address
	BaseToken   common.Address
	QuoteAmount *big.Int
	Executor    common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMarketBuyExecuted is a free log retrieval operation binding the contract event 0xc7b4d815bd0a3b8348577971cb55a9334c80a0b3ccdfe287b55b5a6320d8480a.
//
// Solidity: event MarketBuyExecuted(address indexed pair, address indexed quoteToken, address indexed baseToken, uint256 quoteAmount, address executor)
func (_BuyBot *BuyBotFilterer) FilterMarketBuyExecuted(opts *bind.FilterOpts, pair []common.Address, quoteToken []common.Address, baseToken []common.Address) (*BuyBotMarketBuyExecutedIterator, error) {

	var pairRule []interface{}
	for _, pairItem := range pair {
		pairRule = append(pairRule, pairItem)
	}
	var quoteTokenRule []interface{}
	for _, quoteTokenItem := range quoteToken {
		quoteTokenRule = append(quoteTokenRule, quoteTokenItem)
	}
	var baseTokenRule []interface{}
	for _, baseTokenItem := range baseToken {
		baseTokenRule = append(baseTokenRule, baseTokenItem)
	}

	logs, sub, err := _BuyBot.contract.FilterLogs(opts, "MarketBuyExecuted", pairRule, quoteTokenRule, baseTokenRule)
	if err != nil {
		return nil, err
	}
	return &BuyBotMarketBuyExecutedIterator{contract: _BuyBot.contract, event: "MarketBuyExecuted", logs: logs, sub: sub}, nil
}

// WatchMarketBuyExecuted is a free log subscription operation binding the contract event 0xc7b4d815bd0a3b8348577971cb55a9334c80a0b3ccdfe287b55b5a6320d8480a.
//
// Solidity: event MarketBuyExecuted(address indexed pair, address indexed quoteToken, address indexed baseToken, uint256 quoteAmount, address executor)
func (_BuyBot *BuyBotFilterer) WatchMarketBuyExecuted(opts *bind.WatchOpts, sink chan<- *BuyBotMarketBuyExecuted, pair []common.Address, quoteToken []common.Address, baseToken []common.Address) (event.Subscription, error) {

	var pairRule []interface{}
	for _, pairItem := range pair {
		pairRule = append(pairRule, pairItem)
	}
	var quoteTokenRule []interface{}
	for _, quoteTokenItem := range quoteToken {
		quoteTokenRule = append(quoteTokenRule, quoteTokenItem)
	}
	var baseTokenRule []interface{}
	for _, baseTokenItem := range baseToken {
		baseTokenRule = append(baseTokenRule, baseTokenItem)
	}

	logs, sub, err := _BuyBot.contract.WatchLogs(opts, "MarketBuyExecuted", pairRule, quoteTokenRule, baseTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuyBotMarketBuyExecuted)
				if err := _BuyBot.contract.UnpackLog(event, "MarketBuyExecuted", log); err != nil {
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

// ParseMarketBuyExecuted is a log parse operation binding the contract event 0xc7b4d815bd0a3b8348577971cb55a9334c80a0b3ccdfe287b55b5a6320d8480a.
//
// Solidity: event MarketBuyExecuted(address indexed pair, address indexed quoteToken, address indexed baseToken, uint256 quoteAmount, address executor)
func (_BuyBot *BuyBotFilterer) ParseMarketBuyExecuted(log types.Log) (*BuyBotMarketBuyExecuted, error) {
	event := new(BuyBotMarketBuyExecuted)
	if err := _BuyBot.contract.UnpackLog(event, "MarketBuyExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BuyBotMinOrderAmountSetIterator is returned from FilterMinOrderAmountSet and is used to iterate over the raw logs and unpacked data for MinOrderAmountSet events raised by the BuyBot contract.
type BuyBotMinOrderAmountSetIterator struct {
	Event *BuyBotMinOrderAmountSet // Event containing the contract specifics and raw log

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
func (it *BuyBotMinOrderAmountSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuyBotMinOrderAmountSet)
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
		it.Event = new(BuyBotMinOrderAmountSet)
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
func (it *BuyBotMinOrderAmountSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuyBotMinOrderAmountSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuyBotMinOrderAmountSet represents a MinOrderAmountSet event raised by the BuyBot contract.
type BuyBotMinOrderAmountSet struct {
	Before  *big.Int
	Current *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinOrderAmountSet is a free log retrieval operation binding the contract event 0xd6d62b78d21fc4f5151029ca37079020e8a8815c5a488b28ee7499d4c2a19e55.
//
// Solidity: event MinOrderAmountSet(uint256 indexed before, uint256 indexed current)
func (_BuyBot *BuyBotFilterer) FilterMinOrderAmountSet(opts *bind.FilterOpts, before []*big.Int, current []*big.Int) (*BuyBotMinOrderAmountSetIterator, error) {

	var beforeRule []interface{}
	for _, beforeItem := range before {
		beforeRule = append(beforeRule, beforeItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _BuyBot.contract.FilterLogs(opts, "MinOrderAmountSet", beforeRule, currentRule)
	if err != nil {
		return nil, err
	}
	return &BuyBotMinOrderAmountSetIterator{contract: _BuyBot.contract, event: "MinOrderAmountSet", logs: logs, sub: sub}, nil
}

// WatchMinOrderAmountSet is a free log subscription operation binding the contract event 0xd6d62b78d21fc4f5151029ca37079020e8a8815c5a488b28ee7499d4c2a19e55.
//
// Solidity: event MinOrderAmountSet(uint256 indexed before, uint256 indexed current)
func (_BuyBot *BuyBotFilterer) WatchMinOrderAmountSet(opts *bind.WatchOpts, sink chan<- *BuyBotMinOrderAmountSet, before []*big.Int, current []*big.Int) (event.Subscription, error) {

	var beforeRule []interface{}
	for _, beforeItem := range before {
		beforeRule = append(beforeRule, beforeItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _BuyBot.contract.WatchLogs(opts, "MinOrderAmountSet", beforeRule, currentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuyBotMinOrderAmountSet)
				if err := _BuyBot.contract.UnpackLog(event, "MinOrderAmountSet", log); err != nil {
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

// ParseMinOrderAmountSet is a log parse operation binding the contract event 0xd6d62b78d21fc4f5151029ca37079020e8a8815c5a488b28ee7499d4c2a19e55.
//
// Solidity: event MinOrderAmountSet(uint256 indexed before, uint256 indexed current)
func (_BuyBot *BuyBotFilterer) ParseMinOrderAmountSet(log types.Log) (*BuyBotMinOrderAmountSet, error) {
	event := new(BuyBotMinOrderAmountSet)
	if err := _BuyBot.contract.UnpackLog(event, "MinOrderAmountSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BuyBotOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BuyBot contract.
type BuyBotOwnershipTransferredIterator struct {
	Event *BuyBotOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BuyBotOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuyBotOwnershipTransferred)
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
		it.Event = new(BuyBotOwnershipTransferred)
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
func (it *BuyBotOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuyBotOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuyBotOwnershipTransferred represents a OwnershipTransferred event raised by the BuyBot contract.
type BuyBotOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BuyBot *BuyBotFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BuyBotOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BuyBot.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BuyBotOwnershipTransferredIterator{contract: _BuyBot.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BuyBot *BuyBotFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BuyBotOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BuyBot.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuyBotOwnershipTransferred)
				if err := _BuyBot.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BuyBot *BuyBotFilterer) ParseOwnershipTransferred(log types.Log) (*BuyBotOwnershipTransferred, error) {
	event := new(BuyBotOwnershipTransferred)
	if err := _BuyBot.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BuyBotRecipientSetIterator is returned from FilterRecipientSet and is used to iterate over the raw logs and unpacked data for RecipientSet events raised by the BuyBot contract.
type BuyBotRecipientSetIterator struct {
	Event *BuyBotRecipientSet // Event containing the contract specifics and raw log

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
func (it *BuyBotRecipientSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuyBotRecipientSet)
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
		it.Event = new(BuyBotRecipientSet)
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
func (it *BuyBotRecipientSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuyBotRecipientSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuyBotRecipientSet represents a RecipientSet event raised by the BuyBot contract.
type BuyBotRecipientSet struct {
	Before  common.Address
	Current common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRecipientSet is a free log retrieval operation binding the contract event 0xc1416b5cdab50a9fbc872236e1aa54566c6deb40024e63a4b1737ecacf09d6f9.
//
// Solidity: event RecipientSet(address indexed before, address indexed current)
func (_BuyBot *BuyBotFilterer) FilterRecipientSet(opts *bind.FilterOpts, before []common.Address, current []common.Address) (*BuyBotRecipientSetIterator, error) {

	var beforeRule []interface{}
	for _, beforeItem := range before {
		beforeRule = append(beforeRule, beforeItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _BuyBot.contract.FilterLogs(opts, "RecipientSet", beforeRule, currentRule)
	if err != nil {
		return nil, err
	}
	return &BuyBotRecipientSetIterator{contract: _BuyBot.contract, event: "RecipientSet", logs: logs, sub: sub}, nil
}

// WatchRecipientSet is a free log subscription operation binding the contract event 0xc1416b5cdab50a9fbc872236e1aa54566c6deb40024e63a4b1737ecacf09d6f9.
//
// Solidity: event RecipientSet(address indexed before, address indexed current)
func (_BuyBot *BuyBotFilterer) WatchRecipientSet(opts *bind.WatchOpts, sink chan<- *BuyBotRecipientSet, before []common.Address, current []common.Address) (event.Subscription, error) {

	var beforeRule []interface{}
	for _, beforeItem := range before {
		beforeRule = append(beforeRule, beforeItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _BuyBot.contract.WatchLogs(opts, "RecipientSet", beforeRule, currentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuyBotRecipientSet)
				if err := _BuyBot.contract.UnpackLog(event, "RecipientSet", log); err != nil {
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

// ParseRecipientSet is a log parse operation binding the contract event 0xc1416b5cdab50a9fbc872236e1aa54566c6deb40024e63a4b1737ecacf09d6f9.
//
// Solidity: event RecipientSet(address indexed before, address indexed current)
func (_BuyBot *BuyBotFilterer) ParseRecipientSet(log types.Log) (*BuyBotRecipientSet, error) {
	event := new(BuyBotRecipientSet)
	if err := _BuyBot.contract.UnpackLog(event, "RecipientSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BuyBotWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the BuyBot contract.
type BuyBotWithdrawnIterator struct {
	Event *BuyBotWithdrawn // Event containing the contract specifics and raw log

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
func (it *BuyBotWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuyBotWithdrawn)
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
		it.Event = new(BuyBotWithdrawn)
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
func (it *BuyBotWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuyBotWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuyBotWithdrawn represents a Withdrawn event raised by the BuyBot contract.
type BuyBotWithdrawn struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed token, address indexed to, uint256 amount)
func (_BuyBot *BuyBotFilterer) FilterWithdrawn(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*BuyBotWithdrawnIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BuyBot.contract.FilterLogs(opts, "Withdrawn", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BuyBotWithdrawnIterator{contract: _BuyBot.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed token, address indexed to, uint256 amount)
func (_BuyBot *BuyBotFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *BuyBotWithdrawn, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BuyBot.contract.WatchLogs(opts, "Withdrawn", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuyBotWithdrawn)
				if err := _BuyBot.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed token, address indexed to, uint256 amount)
func (_BuyBot *BuyBotFilterer) ParseWithdrawn(log types.Log) (*BuyBotWithdrawn, error) {
	event := new(BuyBotWithdrawn)
	if err := _BuyBot.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
