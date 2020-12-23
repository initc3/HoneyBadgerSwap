// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hbswap

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// HbSwapABI is the input ABI used to generate the binding from.
const HbSwapABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_servers\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"SecretDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"SecretWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"idxA\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"idxB\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maskedA\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maskedB\",\"type\":\"uint256\"}],\"name\":\"Trade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"idxA\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"idxB\",\"type\":\"uint256\"}],\"name\":\"TradePrep\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"constant\":true,\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"consentRecord\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inputmaskCnt\",\"constant\":true,\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"secretWithdrawCnt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"secretWithdrawMap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"serverNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"servers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amt\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amt\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amt\",\"type\":\"uint256\"}],\"name\":\"secretDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amt\",\"type\":\"uint256\"}],\"name\":\"secretWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_seq\",\"type\":\"uint256\"}],\"name\":\"consent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tradePrep\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_idxA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_idxB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maskedA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maskedB\",\"type\":\"uint256\"}],\"name\":\"trade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// HbSwapBin is the compiled bytecode used for deploying new contracts.
var HbSwapBin = "0x60806040523480156200001157600080fd5b5060405162001b7b38038062001b7b833981810160405260408110156200003757600080fd5b81019080805160405193929190846401000000008211156200005857600080fd5b838201915060208201858111156200006f57600080fd5b82518660208202830111640100000000821117156200008d57600080fd5b8083526020830192505050908051906020019060200280838360005b83811015620000c6578082015181840152602081019050620000a9565b505050509050016040526020018051906020019092919050505060008090505b82518110156200016a576001600460008584815181106200010357fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508080600101915050620000e6565b508060008190555050506119f780620001846000396000f3fe6080604052600436106100f35760003560e01c806393910e661161008a578063c23f001f11610059578063c23f001f146104c9578063d826f88f1461054e578063dee4059514610565578063f3fef3a3146105c0576100f3565b806393910e661461036f5780639eef6ce81461039a578063ade28aad146103d5578063b72a213914610430576100f3565b806333d5e5d2116100c657806333d5e5d21461021657806342cde4e81461024157806347e7ef241461026c5780637aa6fd65146102ba576100f3565b8063064d4810146100f857806312ada8de1461010f5780631768af96146101785780633394dc6f146101eb575b600080fd5b34801561010457600080fd5b5061010d61061b565b005b34801561011b57600080fd5b5061015e6004803603602081101561013257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506106a7565b604051808215151515815260200191505060405180910390f35b34801561018457600080fd5b506101d16004803603604081101561019b57600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506106c7565b604051808215151515815260200191505060405180910390f35b3480156101f757600080fd5b506102006106f6565b6040518082815260200191505060405180910390f35b34801561022257600080fd5b5061022b6106fc565b6040518082815260200191505060405180910390f35b34801561024d57600080fd5b50610256610702565b6040518082815260200191505060405180910390f35b6102b86004803603604081101561028257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610708565b005b3480156102c657600080fd5b506102f3600480360360208110156102dd57600080fd5b8101908080359060200190929190505050610894565b604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060405180910390f35b34801561037b57600080fd5b506103846108fe565b6040518082815260200191505060405180910390f35b3480156103a657600080fd5b506103d3600480360360208110156103bd57600080fd5b8101908080359060200190929190505050610904565b005b3480156103e157600080fd5b5061042e600480360360408110156103f857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610d11565b005b34801561043c57600080fd5b506104c7600480360360c081101561045357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001909291908035906020019092919080359060200190929190505050610f36565b005b3480156104d557600080fd5b50610538600480360360408110156104ec57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506110ca565b6040518082815260200191505060405180910390f35b34801561055a57600080fd5b506105636110ef565b005b34801561057157600080fd5b506105be6004803603604081101561058857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506110f9565b005b3480156105cc57600080fd5b50610619600480360360408110156105e357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506112b4565b005b7fb2fd402d6b838b10cf190139b9d4495eefcfea7543bc1056544d13732d82e6ac336001546001805401604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001838152602001828152602001935050505060405180910390a16002600160008282540192505081905550565b60046020528060005260406000206000915054906101000a900460ff1681565b60076020528160005260406000206020528060005260406000206000915091509054906101000a900460ff1681565b60055481565b60035481565b60005481565b6000339050600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614156107d15734600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254019250508190555061088f565b60008390506108038230858473ffffffffffffffffffffffffffffffffffffffff16611485909392919063ffffffff16565b82600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282540192505081905550505b505050565b60066020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020154905083565b60015481565b6000339050600460008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166109c8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f6e6f7420612076616c696420736572766572000000000000000000000000000081525060200191505060405180910390fd5b6007600083815260200190815260200160002060008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615610a99576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600f8152602001807f616c726561647920636f6e73656e74000000000000000000000000000000000081525060200191505060405180910390fd5b60016007600084815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600160086000848152602001908152602001600020600082825401925050819055506000546008600084815260200190815260200160002054118015610b6657506009600083815260200190815260200160002060009054906101000a900460ff16155b15610d0d57610b7361194a565b600660008481526020019081526020016000206040518060600160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820154815250509050806040015160026000836000015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000836020015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254019250508190555060016009600085815260200190815260200160002060006101000a81548160ff021916908315150217905550505b5050565b600033905081600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015610e08576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f6e6f7420656e6f7567682062616c616e6365000000000000000000000000000081525060200191505060405180910390fd5b81600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825403925050819055507f07c06144435b7d2bdccf9ee7e5a7022c63382ac7c3a0e14ed08b5969dedf0ecf838284604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060405180910390a1505050565b8473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff1610610fd7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f696e76616c69642074726164696e67207061697200000000000000000000000081525060200191505060405180910390fd5b7ffcd787469e8a6f203bfd16a0cb03516726a80f30d3adb2d93af6739e007b54b233878787878787604051808873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200185815260200184815260200183815260200182815260200197505050505050505060405180910390a1505050505050565b6002602052816000526040600020602052806000526040600020600091509150505481565b6000600181905550565b6000339050600160056000828254019250508190555060405180606001604052808473ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1681526020018381525060066000600554815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550604082015181600201559050507f4ef3cc4825a92c3b6922acc8a45152cc96ef48463e8ed500dacd5df9e29a67f3600554848385604051808581526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200194505050505060405180910390a1505050565b6000339050600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614156113c45781600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825403925050819055508073ffffffffffffffffffffffffffffffffffffffff166108fc839081150290604051600060405180830381858888f193505050501580156113be573d6000803e3d6000fd5b50611480565b60008390506113f482848373ffffffffffffffffffffffffffffffffffffffff166115729092919063ffffffff16565b82600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282540392505081905550505b505050565b61156c846323b872dd60e01b858585604051602401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505061162a565b50505050565b6116258363a9059cbb60e01b8484604051602401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505061162a565b505050565b606061168c826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166117199092919063ffffffff16565b9050600081511115611714578080602001905160208110156116ad57600080fd5b8101908080519060200190929190505050611713576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a815260200180611998602a913960400191505060405180910390fd5b5b505050565b60606117288484600085611731565b90509392505050565b606061173c85611937565b6117ae576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000081525060200191505060405180910390fd5b600060608673ffffffffffffffffffffffffffffffffffffffff1685876040518082805190602001908083835b602083106117fe57805182526020820191506020810190506020830392506117db565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d8060008114611860576040519150601f19603f3d011682016040523d82523d6000602084013e611865565b606091505b5091509150811561187a57809250505061192f565b60008151111561188d5780518082602001fd5b836040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b838110156118f45780820151818401526020810190506118d9565b50505050905090810190601f1680156119215780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b949350505050565b600080823b905060008111915050919050565b6040518060600160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152509056fe5361666545524332303a204552433230206f7065726174696f6e20646964206e6f742073756363656564a264697066735822122008fd8db82ad014bdfad8d0844c60ae19c4ff68b91097f3cb530c98c85d5db94964736f6c634300060a0033"

// DeployHbSwap deploys a new Ethereum contract, binding an instance of HbSwap to it.
func DeployHbSwap(auth *bind.TransactOpts, backend bind.ContractBackend, _servers []common.Address, _threshold *big.Int) (common.Address, *types.Transaction, *HbSwap, error) {
	parsed, err := abi.JSON(strings.NewReader(HbSwapABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HbSwapBin), backend, _servers, _threshold)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HbSwap{HbSwapCaller: HbSwapCaller{contract: contract}, HbSwapTransactor: HbSwapTransactor{contract: contract}, HbSwapFilterer: HbSwapFilterer{contract: contract}}, nil
}

// HbSwap is an auto generated Go binding around an Ethereum contract.
type HbSwap struct {
	HbSwapCaller     // Read-only binding to the contract
	HbSwapTransactor // Write-only binding to the contract
	HbSwapFilterer   // Log filterer for contract events
}

// HbSwapCaller is an auto generated read-only Go binding around an Ethereum contract.
type HbSwapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HbSwapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HbSwapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HbSwapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HbSwapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HbSwapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HbSwapSession struct {
	Contract     *HbSwap           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HbSwapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HbSwapCallerSession struct {
	Contract *HbSwapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// HbSwapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HbSwapTransactorSession struct {
	Contract     *HbSwapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HbSwapRaw is an auto generated low-level Go binding around an Ethereum contract.
type HbSwapRaw struct {
	Contract *HbSwap // Generic contract binding to access the raw methods on
}

// HbSwapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HbSwapCallerRaw struct {
	Contract *HbSwapCaller // Generic read-only contract binding to access the raw methods on
}

// HbSwapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HbSwapTransactorRaw struct {
	Contract *HbSwapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHbSwap creates a new instance of HbSwap, bound to a specific deployed contract.
func NewHbSwap(address common.Address, backend bind.ContractBackend) (*HbSwap, error) {
	contract, err := bindHbSwap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HbSwap{HbSwapCaller: HbSwapCaller{contract: contract}, HbSwapTransactor: HbSwapTransactor{contract: contract}, HbSwapFilterer: HbSwapFilterer{contract: contract}}, nil
}

// NewHbSwapCaller creates a new read-only instance of HbSwap, bound to a specific deployed contract.
func NewHbSwapCaller(address common.Address, caller bind.ContractCaller) (*HbSwapCaller, error) {
	contract, err := bindHbSwap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HbSwapCaller{contract: contract}, nil
}

// NewHbSwapTransactor creates a new write-only instance of HbSwap, bound to a specific deployed contract.
func NewHbSwapTransactor(address common.Address, transactor bind.ContractTransactor) (*HbSwapTransactor, error) {
	contract, err := bindHbSwap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HbSwapTransactor{contract: contract}, nil
}

// NewHbSwapFilterer creates a new log filterer instance of HbSwap, bound to a specific deployed contract.
func NewHbSwapFilterer(address common.Address, filterer bind.ContractFilterer) (*HbSwapFilterer, error) {
	contract, err := bindHbSwap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HbSwapFilterer{contract: contract}, nil
}

// bindHbSwap binds a generic wrapper to an already deployed contract.
func bindHbSwap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HbSwapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HbSwap *HbSwapRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HbSwap.Contract.HbSwapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HbSwap *HbSwapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HbSwap.Contract.HbSwapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HbSwap *HbSwapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HbSwap.Contract.HbSwapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HbSwap *HbSwapCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HbSwap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HbSwap *HbSwapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HbSwap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HbSwap *HbSwapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HbSwap.Contract.contract.Transact(opts, method, params...)
}

// Balances is a free data retrieval call binding the contract method 0xc23f001f.
//
// Solidity: function balances(address , address ) constant returns(uint256)
func (_HbSwap *HbSwapCaller) Balances(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HbSwap.contract.Call(opts, out, "balances", arg0, arg1)
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0xc23f001f.
//
// Solidity: function balances(address , address ) constant returns(uint256)
func (_HbSwap *HbSwapSession) Balances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _HbSwap.Contract.Balances(&_HbSwap.CallOpts, arg0, arg1)
}

// Balances is a free data retrieval call binding the contract method 0xc23f001f.
//
// Solidity: function balances(address , address ) constant returns(uint256)
func (_HbSwap *HbSwapCallerSession) Balances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _HbSwap.Contract.Balances(&_HbSwap.CallOpts, arg0, arg1)
}

// InputmaskCnt is a free data retrieval call binding the contract method 0x93910e66.
//
// Solidity: function inputmaskCnt() constant returns(uint256)
func (_HbSwap *HbSwapCaller) InputmaskCnt(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HbSwap.contract.Call(opts, out, "inputmaskCnt")
	return *ret0, err
}

// InputmaskCnt is a free data retrieval call binding the contract method 0x93910e66.
//
// Solidity: function inputmaskCnt() constant returns(uint256)
func (_HbSwap *HbSwapSession) InputmaskCnt() (*big.Int, error) {
	return _HbSwap.Contract.InputmaskCnt(&_HbSwap.CallOpts)
}

// InputmaskCnt is a free data retrieval call binding the contract method 0x93910e66.
//
// Solidity: function inputmaskCnt() constant returns(uint256)
func (_HbSwap *HbSwapCallerSession) InputmaskCnt() (*big.Int, error) {
	return _HbSwap.Contract.InputmaskCnt(&_HbSwap.CallOpts)
}

// Consent is a paid mutator transaction binding the contract method 0x9eef6ce8.
//
// Solidity: function consent(uint256 _seq) returns()
func (_HbSwap *HbSwapTransactor) Consent(opts *bind.TransactOpts, _seq *big.Int) (*types.Transaction, error) {
	return _HbSwap.contract.Transact(opts, "consent", _seq)
}

// Consent is a paid mutator transaction binding the contract method 0x9eef6ce8.
//
// Solidity: function consent(uint256 _seq) returns()
func (_HbSwap *HbSwapSession) Consent(_seq *big.Int) (*types.Transaction, error) {
	return _HbSwap.Contract.Consent(&_HbSwap.TransactOpts, _seq)
}

// Consent is a paid mutator transaction binding the contract method 0x9eef6ce8.
//
// Solidity: function consent(uint256 _seq) returns()
func (_HbSwap *HbSwapTransactorSession) Consent(_seq *big.Int) (*types.Transaction, error) {
	return _HbSwap.Contract.Consent(&_HbSwap.TransactOpts, _seq)
}

// ConsentRecord is a paid mutator transaction binding the contract method 0x1768af96.
//
// Solidity: function consentRecord(uint256 , address ) returns(bool)
func (_HbSwap *HbSwapTransactor) ConsentRecord(opts *bind.TransactOpts, arg0 *big.Int, arg1 common.Address) (*types.Transaction, error) {
	return _HbSwap.contract.Transact(opts, "consentRecord", arg0, arg1)
}

// ConsentRecord is a paid mutator transaction binding the contract method 0x1768af96.
//
// Solidity: function consentRecord(uint256 , address ) returns(bool)
func (_HbSwap *HbSwapSession) ConsentRecord(arg0 *big.Int, arg1 common.Address) (*types.Transaction, error) {
	return _HbSwap.Contract.ConsentRecord(&_HbSwap.TransactOpts, arg0, arg1)
}

// ConsentRecord is a paid mutator transaction binding the contract method 0x1768af96.
//
// Solidity: function consentRecord(uint256 , address ) returns(bool)
func (_HbSwap *HbSwapTransactorSession) ConsentRecord(arg0 *big.Int, arg1 common.Address) (*types.Transaction, error) {
	return _HbSwap.Contract.ConsentRecord(&_HbSwap.TransactOpts, arg0, arg1)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address _token, uint256 _amt) returns()
func (_HbSwap *HbSwapTransactor) Deposit(opts *bind.TransactOpts, _token common.Address, _amt *big.Int) (*types.Transaction, error) {
	return _HbSwap.contract.Transact(opts, "deposit", _token, _amt)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address _token, uint256 _amt) returns()
func (_HbSwap *HbSwapSession) Deposit(_token common.Address, _amt *big.Int) (*types.Transaction, error) {
	return _HbSwap.Contract.Deposit(&_HbSwap.TransactOpts, _token, _amt)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address _token, uint256 _amt) returns()
func (_HbSwap *HbSwapTransactorSession) Deposit(_token common.Address, _amt *big.Int) (*types.Transaction, error) {
	return _HbSwap.Contract.Deposit(&_HbSwap.TransactOpts, _token, _amt)
}

// Reset is a paid mutator transaction binding the contract method 0xd826f88f.
//
// Solidity: function reset() returns()
func (_HbSwap *HbSwapTransactor) Reset(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HbSwap.contract.Transact(opts, "reset")
}

// Reset is a paid mutator transaction binding the contract method 0xd826f88f.
//
// Solidity: function reset() returns()
func (_HbSwap *HbSwapSession) Reset() (*types.Transaction, error) {
	return _HbSwap.Contract.Reset(&_HbSwap.TransactOpts)
}

// Reset is a paid mutator transaction binding the contract method 0xd826f88f.
//
// Solidity: function reset() returns()
func (_HbSwap *HbSwapTransactorSession) Reset() (*types.Transaction, error) {
	return _HbSwap.Contract.Reset(&_HbSwap.TransactOpts)
}

// SecretDeposit is a paid mutator transaction binding the contract method 0xade28aad.
//
// Solidity: function secretDeposit(address _token, uint256 _amt) returns()
func (_HbSwap *HbSwapTransactor) SecretDeposit(opts *bind.TransactOpts, _token common.Address, _amt *big.Int) (*types.Transaction, error) {
	return _HbSwap.contract.Transact(opts, "secretDeposit", _token, _amt)
}

// SecretDeposit is a paid mutator transaction binding the contract method 0xade28aad.
//
// Solidity: function secretDeposit(address _token, uint256 _amt) returns()
func (_HbSwap *HbSwapSession) SecretDeposit(_token common.Address, _amt *big.Int) (*types.Transaction, error) {
	return _HbSwap.Contract.SecretDeposit(&_HbSwap.TransactOpts, _token, _amt)
}

// SecretDeposit is a paid mutator transaction binding the contract method 0xade28aad.
//
// Solidity: function secretDeposit(address _token, uint256 _amt) returns()
func (_HbSwap *HbSwapTransactorSession) SecretDeposit(_token common.Address, _amt *big.Int) (*types.Transaction, error) {
	return _HbSwap.Contract.SecretDeposit(&_HbSwap.TransactOpts, _token, _amt)
}

// SecretWithdraw is a paid mutator transaction binding the contract method 0xdee40595.
//
// Solidity: function secretWithdraw(address _token, uint256 _amt) returns()
func (_HbSwap *HbSwapTransactor) SecretWithdraw(opts *bind.TransactOpts, _token common.Address, _amt *big.Int) (*types.Transaction, error) {
	return _HbSwap.contract.Transact(opts, "secretWithdraw", _token, _amt)
}

// SecretWithdraw is a paid mutator transaction binding the contract method 0xdee40595.
//
// Solidity: function secretWithdraw(address _token, uint256 _amt) returns()
func (_HbSwap *HbSwapSession) SecretWithdraw(_token common.Address, _amt *big.Int) (*types.Transaction, error) {
	return _HbSwap.Contract.SecretWithdraw(&_HbSwap.TransactOpts, _token, _amt)
}

// SecretWithdraw is a paid mutator transaction binding the contract method 0xdee40595.
//
// Solidity: function secretWithdraw(address _token, uint256 _amt) returns()
func (_HbSwap *HbSwapTransactorSession) SecretWithdraw(_token common.Address, _amt *big.Int) (*types.Transaction, error) {
	return _HbSwap.Contract.SecretWithdraw(&_HbSwap.TransactOpts, _token, _amt)
}

// SecretWithdrawCnt is a paid mutator transaction binding the contract method 0x3394dc6f.
//
// Solidity: function secretWithdrawCnt() returns(uint256)
func (_HbSwap *HbSwapTransactor) SecretWithdrawCnt(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HbSwap.contract.Transact(opts, "secretWithdrawCnt")
}

// SecretWithdrawCnt is a paid mutator transaction binding the contract method 0x3394dc6f.
//
// Solidity: function secretWithdrawCnt() returns(uint256)
func (_HbSwap *HbSwapSession) SecretWithdrawCnt() (*types.Transaction, error) {
	return _HbSwap.Contract.SecretWithdrawCnt(&_HbSwap.TransactOpts)
}

// SecretWithdrawCnt is a paid mutator transaction binding the contract method 0x3394dc6f.
//
// Solidity: function secretWithdrawCnt() returns(uint256)
func (_HbSwap *HbSwapTransactorSession) SecretWithdrawCnt() (*types.Transaction, error) {
	return _HbSwap.Contract.SecretWithdrawCnt(&_HbSwap.TransactOpts)
}

// SecretWithdrawMap is a paid mutator transaction binding the contract method 0x7aa6fd65.
//
// Solidity: function secretWithdrawMap(uint256 ) returns(address token, address user, uint256 amt)
func (_HbSwap *HbSwapTransactor) SecretWithdrawMap(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _HbSwap.contract.Transact(opts, "secretWithdrawMap", arg0)
}

// SecretWithdrawMap is a paid mutator transaction binding the contract method 0x7aa6fd65.
//
// Solidity: function secretWithdrawMap(uint256 ) returns(address token, address user, uint256 amt)
func (_HbSwap *HbSwapSession) SecretWithdrawMap(arg0 *big.Int) (*types.Transaction, error) {
	return _HbSwap.Contract.SecretWithdrawMap(&_HbSwap.TransactOpts, arg0)
}

// SecretWithdrawMap is a paid mutator transaction binding the contract method 0x7aa6fd65.
//
// Solidity: function secretWithdrawMap(uint256 ) returns(address token, address user, uint256 amt)
func (_HbSwap *HbSwapTransactorSession) SecretWithdrawMap(arg0 *big.Int) (*types.Transaction, error) {
	return _HbSwap.Contract.SecretWithdrawMap(&_HbSwap.TransactOpts, arg0)
}

// ServerNum is a paid mutator transaction binding the contract method 0x33d5e5d2.
//
// Solidity: function serverNum() returns(uint256)
func (_HbSwap *HbSwapTransactor) ServerNum(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HbSwap.contract.Transact(opts, "serverNum")
}

// ServerNum is a paid mutator transaction binding the contract method 0x33d5e5d2.
//
// Solidity: function serverNum() returns(uint256)
func (_HbSwap *HbSwapSession) ServerNum() (*types.Transaction, error) {
	return _HbSwap.Contract.ServerNum(&_HbSwap.TransactOpts)
}

// ServerNum is a paid mutator transaction binding the contract method 0x33d5e5d2.
//
// Solidity: function serverNum() returns(uint256)
func (_HbSwap *HbSwapTransactorSession) ServerNum() (*types.Transaction, error) {
	return _HbSwap.Contract.ServerNum(&_HbSwap.TransactOpts)
}

// Servers is a paid mutator transaction binding the contract method 0x12ada8de.
//
// Solidity: function servers(address ) returns(bool)
func (_HbSwap *HbSwapTransactor) Servers(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _HbSwap.contract.Transact(opts, "servers", arg0)
}

// Servers is a paid mutator transaction binding the contract method 0x12ada8de.
//
// Solidity: function servers(address ) returns(bool)
func (_HbSwap *HbSwapSession) Servers(arg0 common.Address) (*types.Transaction, error) {
	return _HbSwap.Contract.Servers(&_HbSwap.TransactOpts, arg0)
}

// Servers is a paid mutator transaction binding the contract method 0x12ada8de.
//
// Solidity: function servers(address ) returns(bool)
func (_HbSwap *HbSwapTransactorSession) Servers(arg0 common.Address) (*types.Transaction, error) {
	return _HbSwap.Contract.Servers(&_HbSwap.TransactOpts, arg0)
}

// Threshold is a paid mutator transaction binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() returns(uint256)
func (_HbSwap *HbSwapTransactor) Threshold(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HbSwap.contract.Transact(opts, "threshold")
}

// Threshold is a paid mutator transaction binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() returns(uint256)
func (_HbSwap *HbSwapSession) Threshold() (*types.Transaction, error) {
	return _HbSwap.Contract.Threshold(&_HbSwap.TransactOpts)
}

// Threshold is a paid mutator transaction binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() returns(uint256)
func (_HbSwap *HbSwapTransactorSession) Threshold() (*types.Transaction, error) {
	return _HbSwap.Contract.Threshold(&_HbSwap.TransactOpts)
}

// Trade is a paid mutator transaction binding the contract method 0xb72a2139.
//
// Solidity: function trade(address _tokenA, address _tokenB, uint256 _idxA, uint256 _idxB, uint256 _maskedA, uint256 _maskedB) returns()
func (_HbSwap *HbSwapTransactor) Trade(opts *bind.TransactOpts, _tokenA common.Address, _tokenB common.Address, _idxA *big.Int, _idxB *big.Int, _maskedA *big.Int, _maskedB *big.Int) (*types.Transaction, error) {
	return _HbSwap.contract.Transact(opts, "trade", _tokenA, _tokenB, _idxA, _idxB, _maskedA, _maskedB)
}

// Trade is a paid mutator transaction binding the contract method 0xb72a2139.
//
// Solidity: function trade(address _tokenA, address _tokenB, uint256 _idxA, uint256 _idxB, uint256 _maskedA, uint256 _maskedB) returns()
func (_HbSwap *HbSwapSession) Trade(_tokenA common.Address, _tokenB common.Address, _idxA *big.Int, _idxB *big.Int, _maskedA *big.Int, _maskedB *big.Int) (*types.Transaction, error) {
	return _HbSwap.Contract.Trade(&_HbSwap.TransactOpts, _tokenA, _tokenB, _idxA, _idxB, _maskedA, _maskedB)
}

// Trade is a paid mutator transaction binding the contract method 0xb72a2139.
//
// Solidity: function trade(address _tokenA, address _tokenB, uint256 _idxA, uint256 _idxB, uint256 _maskedA, uint256 _maskedB) returns()
func (_HbSwap *HbSwapTransactorSession) Trade(_tokenA common.Address, _tokenB common.Address, _idxA *big.Int, _idxB *big.Int, _maskedA *big.Int, _maskedB *big.Int) (*types.Transaction, error) {
	return _HbSwap.Contract.Trade(&_HbSwap.TransactOpts, _tokenA, _tokenB, _idxA, _idxB, _maskedA, _maskedB)
}

// TradePrep is a paid mutator transaction binding the contract method 0x064d4810.
//
// Solidity: function tradePrep() returns()
func (_HbSwap *HbSwapTransactor) TradePrep(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HbSwap.contract.Transact(opts, "tradePrep")
}

// TradePrep is a paid mutator transaction binding the contract method 0x064d4810.
//
// Solidity: function tradePrep() returns()
func (_HbSwap *HbSwapSession) TradePrep() (*types.Transaction, error) {
	return _HbSwap.Contract.TradePrep(&_HbSwap.TransactOpts)
}

// TradePrep is a paid mutator transaction binding the contract method 0x064d4810.
//
// Solidity: function tradePrep() returns()
func (_HbSwap *HbSwapTransactorSession) TradePrep() (*types.Transaction, error) {
	return _HbSwap.Contract.TradePrep(&_HbSwap.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _token, uint256 _amt) returns()
func (_HbSwap *HbSwapTransactor) Withdraw(opts *bind.TransactOpts, _token common.Address, _amt *big.Int) (*types.Transaction, error) {
	return _HbSwap.contract.Transact(opts, "withdraw", _token, _amt)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _token, uint256 _amt) returns()
func (_HbSwap *HbSwapSession) Withdraw(_token common.Address, _amt *big.Int) (*types.Transaction, error) {
	return _HbSwap.Contract.Withdraw(&_HbSwap.TransactOpts, _token, _amt)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _token, uint256 _amt) returns()
func (_HbSwap *HbSwapTransactorSession) Withdraw(_token common.Address, _amt *big.Int) (*types.Transaction, error) {
	return _HbSwap.Contract.Withdraw(&_HbSwap.TransactOpts, _token, _amt)
}

// HbSwapSecretDepositIterator is returned from FilterSecretDeposit and is used to iterate over the raw logs and unpacked data for SecretDeposit events raised by the HbSwap contract.
type HbSwapSecretDepositIterator struct {
	Event *HbSwapSecretDeposit // Event containing the contract specifics and raw log

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
func (it *HbSwapSecretDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HbSwapSecretDeposit)
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
		it.Event = new(HbSwapSecretDeposit)
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
func (it *HbSwapSecretDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HbSwapSecretDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HbSwapSecretDeposit represents a SecretDeposit event raised by the HbSwap contract.
type HbSwapSecretDeposit struct {
	Token common.Address
	User  common.Address
	Amt   *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSecretDeposit is a free log retrieval operation binding the contract event 0x07c06144435b7d2bdccf9ee7e5a7022c63382ac7c3a0e14ed08b5969dedf0ecf.
//
// Solidity: event SecretDeposit(address token, address user, uint256 amt)
func (_HbSwap *HbSwapFilterer) FilterSecretDeposit(opts *bind.FilterOpts) (*HbSwapSecretDepositIterator, error) {

	logs, sub, err := _HbSwap.contract.FilterLogs(opts, "SecretDeposit")
	if err != nil {
		return nil, err
	}
	return &HbSwapSecretDepositIterator{contract: _HbSwap.contract, event: "SecretDeposit", logs: logs, sub: sub}, nil
}

// WatchSecretDeposit is a free log subscription operation binding the contract event 0x07c06144435b7d2bdccf9ee7e5a7022c63382ac7c3a0e14ed08b5969dedf0ecf.
//
// Solidity: event SecretDeposit(address token, address user, uint256 amt)
func (_HbSwap *HbSwapFilterer) WatchSecretDeposit(opts *bind.WatchOpts, sink chan<- *HbSwapSecretDeposit) (event.Subscription, error) {

	logs, sub, err := _HbSwap.contract.WatchLogs(opts, "SecretDeposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HbSwapSecretDeposit)
				if err := _HbSwap.contract.UnpackLog(event, "SecretDeposit", log); err != nil {
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

// ParseSecretDeposit is a log parse operation binding the contract event 0x07c06144435b7d2bdccf9ee7e5a7022c63382ac7c3a0e14ed08b5969dedf0ecf.
//
// Solidity: event SecretDeposit(address token, address user, uint256 amt)
func (_HbSwap *HbSwapFilterer) ParseSecretDeposit(log types.Log) (*HbSwapSecretDeposit, error) {
	event := new(HbSwapSecretDeposit)
	if err := _HbSwap.contract.UnpackLog(event, "SecretDeposit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// HbSwapSecretWithdrawIterator is returned from FilterSecretWithdraw and is used to iterate over the raw logs and unpacked data for SecretWithdraw events raised by the HbSwap contract.
type HbSwapSecretWithdrawIterator struct {
	Event *HbSwapSecretWithdraw // Event containing the contract specifics and raw log

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
func (it *HbSwapSecretWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HbSwapSecretWithdraw)
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
		it.Event = new(HbSwapSecretWithdraw)
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
func (it *HbSwapSecretWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HbSwapSecretWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HbSwapSecretWithdraw represents a SecretWithdraw event raised by the HbSwap contract.
type HbSwapSecretWithdraw struct {
	Seq   *big.Int
	Token common.Address
	User  common.Address
	Amt   *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSecretWithdraw is a free log retrieval operation binding the contract event 0x4ef3cc4825a92c3b6922acc8a45152cc96ef48463e8ed500dacd5df9e29a67f3.
//
// Solidity: event SecretWithdraw(uint256 seq, address token, address user, uint256 amt)
func (_HbSwap *HbSwapFilterer) FilterSecretWithdraw(opts *bind.FilterOpts) (*HbSwapSecretWithdrawIterator, error) {

	logs, sub, err := _HbSwap.contract.FilterLogs(opts, "SecretWithdraw")
	if err != nil {
		return nil, err
	}
	return &HbSwapSecretWithdrawIterator{contract: _HbSwap.contract, event: "SecretWithdraw", logs: logs, sub: sub}, nil
}

// WatchSecretWithdraw is a free log subscription operation binding the contract event 0x4ef3cc4825a92c3b6922acc8a45152cc96ef48463e8ed500dacd5df9e29a67f3.
//
// Solidity: event SecretWithdraw(uint256 seq, address token, address user, uint256 amt)
func (_HbSwap *HbSwapFilterer) WatchSecretWithdraw(opts *bind.WatchOpts, sink chan<- *HbSwapSecretWithdraw) (event.Subscription, error) {

	logs, sub, err := _HbSwap.contract.WatchLogs(opts, "SecretWithdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HbSwapSecretWithdraw)
				if err := _HbSwap.contract.UnpackLog(event, "SecretWithdraw", log); err != nil {
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

// ParseSecretWithdraw is a log parse operation binding the contract event 0x4ef3cc4825a92c3b6922acc8a45152cc96ef48463e8ed500dacd5df9e29a67f3.
//
// Solidity: event SecretWithdraw(uint256 seq, address token, address user, uint256 amt)
func (_HbSwap *HbSwapFilterer) ParseSecretWithdraw(log types.Log) (*HbSwapSecretWithdraw, error) {
	event := new(HbSwapSecretWithdraw)
	if err := _HbSwap.contract.UnpackLog(event, "SecretWithdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}

// HbSwapTradeIterator is returned from FilterTrade and is used to iterate over the raw logs and unpacked data for Trade events raised by the HbSwap contract.
type HbSwapTradeIterator struct {
	Event *HbSwapTrade // Event containing the contract specifics and raw log

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
func (it *HbSwapTradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HbSwapTrade)
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
		it.Event = new(HbSwapTrade)
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
func (it *HbSwapTradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HbSwapTradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HbSwapTrade represents a Trade event raised by the HbSwap contract.
type HbSwapTrade struct {
	User    common.Address
	TokenA  common.Address
	TokenB  common.Address
	IdxA    *big.Int
	IdxB    *big.Int
	MaskedA *big.Int
	MaskedB *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTrade is a free log retrieval operation binding the contract event 0xfcd787469e8a6f203bfd16a0cb03516726a80f30d3adb2d93af6739e007b54b2.
//
// Solidity: event Trade(address user, address tokenA, address tokenB, uint256 idxA, uint256 idxB, uint256 maskedA, uint256 maskedB)
func (_HbSwap *HbSwapFilterer) FilterTrade(opts *bind.FilterOpts) (*HbSwapTradeIterator, error) {

	logs, sub, err := _HbSwap.contract.FilterLogs(opts, "Trade")
	if err != nil {
		return nil, err
	}
	return &HbSwapTradeIterator{contract: _HbSwap.contract, event: "Trade", logs: logs, sub: sub}, nil
}

// WatchTrade is a free log subscription operation binding the contract event 0xfcd787469e8a6f203bfd16a0cb03516726a80f30d3adb2d93af6739e007b54b2.
//
// Solidity: event Trade(address user, address tokenA, address tokenB, uint256 idxA, uint256 idxB, uint256 maskedA, uint256 maskedB)
func (_HbSwap *HbSwapFilterer) WatchTrade(opts *bind.WatchOpts, sink chan<- *HbSwapTrade) (event.Subscription, error) {

	logs, sub, err := _HbSwap.contract.WatchLogs(opts, "Trade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HbSwapTrade)
				if err := _HbSwap.contract.UnpackLog(event, "Trade", log); err != nil {
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

// ParseTrade is a log parse operation binding the contract event 0xfcd787469e8a6f203bfd16a0cb03516726a80f30d3adb2d93af6739e007b54b2.
//
// Solidity: event Trade(address user, address tokenA, address tokenB, uint256 idxA, uint256 idxB, uint256 maskedA, uint256 maskedB)
func (_HbSwap *HbSwapFilterer) ParseTrade(log types.Log) (*HbSwapTrade, error) {
	event := new(HbSwapTrade)
	if err := _HbSwap.contract.UnpackLog(event, "Trade", log); err != nil {
		return nil, err
	}
	return event, nil
}

// HbSwapTradePrepIterator is returned from FilterTradePrep and is used to iterate over the raw logs and unpacked data for TradePrep events raised by the HbSwap contract.
type HbSwapTradePrepIterator struct {
	Event *HbSwapTradePrep // Event containing the contract specifics and raw log

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
func (it *HbSwapTradePrepIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HbSwapTradePrep)
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
		it.Event = new(HbSwapTradePrep)
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
func (it *HbSwapTradePrepIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HbSwapTradePrepIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HbSwapTradePrep represents a TradePrep event raised by the HbSwap contract.
type HbSwapTradePrep struct {
	User common.Address
	IdxA *big.Int
	IdxB *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTradePrep is a free log retrieval operation binding the contract event 0xb2fd402d6b838b10cf190139b9d4495eefcfea7543bc1056544d13732d82e6ac.
//
// Solidity: event TradePrep(address user, uint256 idxA, uint256 idxB)
func (_HbSwap *HbSwapFilterer) FilterTradePrep(opts *bind.FilterOpts) (*HbSwapTradePrepIterator, error) {

	logs, sub, err := _HbSwap.contract.FilterLogs(opts, "TradePrep")
	if err != nil {
		return nil, err
	}
	return &HbSwapTradePrepIterator{contract: _HbSwap.contract, event: "TradePrep", logs: logs, sub: sub}, nil
}

// WatchTradePrep is a free log subscription operation binding the contract event 0xb2fd402d6b838b10cf190139b9d4495eefcfea7543bc1056544d13732d82e6ac.
//
// Solidity: event TradePrep(address user, uint256 idxA, uint256 idxB)
func (_HbSwap *HbSwapFilterer) WatchTradePrep(opts *bind.WatchOpts, sink chan<- *HbSwapTradePrep) (event.Subscription, error) {

	logs, sub, err := _HbSwap.contract.WatchLogs(opts, "TradePrep")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HbSwapTradePrep)
				if err := _HbSwap.contract.UnpackLog(event, "TradePrep", log); err != nil {
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

// ParseTradePrep is a log parse operation binding the contract event 0xb2fd402d6b838b10cf190139b9d4495eefcfea7543bc1056544d13732d82e6ac.
//
// Solidity: event TradePrep(address user, uint256 idxA, uint256 idxB)
func (_HbSwap *HbSwapFilterer) ParseTradePrep(log types.Log) (*HbSwapTradePrep, error) {
	event := new(HbSwapTradePrep)
	if err := _HbSwap.contract.UnpackLog(event, "TradePrep", log); err != nil {
		return nil, err
	}
	return event, nil
}
