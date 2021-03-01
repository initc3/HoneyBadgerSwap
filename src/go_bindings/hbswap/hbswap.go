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
const HbSwapABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_servers\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amtA\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amtB\",\"type\":\"uint256\"}],\"name\":\"AddLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amtA\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amtB\",\"type\":\"uint256\"}],\"name\":\"InitPool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"RemoveLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"SecretDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"SecretWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tradeSeq\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"idxA\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"idxB\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maskedA\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maskedB\",\"type\":\"uint256\"}],\"name\":\"Trade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"idxA\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"idxB\",\"type\":\"uint256\"}],\"name\":\"TradePrep\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"price\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"UpdatePrice\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"constant\":true,\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"consentRecord\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"inputMaskOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inputmaskCnt\",\"constant\":true,\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastUpdateSeq\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"prices\",\"constant\":true,\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"secretWithdrawCnt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"secretWithdrawMap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"servers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tradeCnt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amtA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amtB\",\"type\":\"uint256\"}],\"name\":\"initPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amtA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amtB\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amt\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amt\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amt\",\"type\":\"uint256\"}],\"name\":\"secretDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amt\",\"type\":\"uint256\"}],\"name\":\"secretWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_seq\",\"type\":\"uint256\"}],\"name\":\"consent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amt\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tradePrep\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_idxA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_idxB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maskedA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maskedB\",\"type\":\"uint256\"}],\"name\":\"trade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_checkpointSeq\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_price\",\"type\":\"string\"}],\"name\":\"updatePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// HbSwapBin is the compiled bytecode used for deploying new contracts.
var HbSwapBin = "0x60806040523480156200001157600080fd5b5060405162002ed538038062002ed5833981810160405260408110156200003757600080fd5b81019080805160405193929190846401000000008211156200005857600080fd5b838201915060208201858111156200006f57600080fd5b82518660208202830111640100000000821117156200008d57600080fd5b8083526020830192505050908051906020019060200280838360005b83811015620000c6578082015181840152602081019050620000a9565b505050509050016040526020018051906020019092919050505060008090505b8251811015620001695760018060008584815181106200010257fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508080600101915050620000e6565b50806000819055505050612d5280620001836000396000f3fe6080604052600436106101355760003560e01c80639eef6ce8116100ab578063b7deef821161006f578063b7deef82146107cf578063c23f001f1461084a578063cf6c62ea146108cf578063d752fab214610954578063dee40595146109cf578063f3fef3a314610a2a57610135565b80639eef6ce8146105f0578063a4bea6c81461062b578063ade28aad146106b0578063af4170c41461070b578063b72a21391461073657610135565b80633394dc6f116100fd5780633394dc6f1461039c5780633766dba7146103c757806342cde4e81461049757806347e7ef24146104c25780637aa6fd651461051057806393910e66146105c557610135565b8063064d48101461013a57806312ada8de146101515780631768af96146101ba5780631f3124041461022d5780632f8a68c514610317575b600080fd5b34801561014657600080fd5b5061014f610a85565b005b34801561015d57600080fd5b506101a06004803603602081101561017457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610bc2565b604051808215151515815260200191505060405180910390f35b3480156101c657600080fd5b50610213600480360360408110156101dd57600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610be2565b604051808215151515815260200191505060405180910390f35b34801561023957600080fd5b5061029c6004803603604081101561025057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610c11565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156102dc5780820151818401526020810190506102c1565b50505050905090810190601f1680156103095780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561032357600080fd5b5061039a6004803603608081101561033a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919080359060200190929190505050610cce565b005b3480156103a857600080fd5b506103b1610e50565b6040518082815260200191505060405180910390f35b3480156103d357600080fd5b50610495600480360360808110156103ea57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019064010000000081111561045157600080fd5b82018360208201111561046357600080fd5b8035906020019184600183028401116401000000008311171561048557600080fd5b9091929391929390505050610e56565b005b3480156104a357600080fd5b506104ac6114a0565b6040518082815260200191505060405180910390f35b61050e600480360360408110156104d857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506114a6565b005b34801561051c57600080fd5b506105496004803603602081101561053357600080fd5b8101908080359060200190929190505050611632565b604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060405180910390f35b3480156105d157600080fd5b506105da61169c565b6040518082815260200191505060405180910390f35b3480156105fc57600080fd5b506106296004803603602081101561061357600080fd5b81019080803590602001909291905050506116a2565b005b34801561063757600080fd5b5061069a6004803603604081101561064e57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611aaf565b6040518082815260200191505060405180910390f35b3480156106bc57600080fd5b50610709600480360360408110156106d357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611ad4565b005b34801561071757600080fd5b50610720611cf9565b6040518082815260200191505060405180910390f35b34801561074257600080fd5b506107cd600480360360c081101561075957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001909291908035906020019092919080359060200190929190505050611cff565b005b3480156107db57600080fd5b50610808600480360360208110156107f257600080fd5b810190808035906020019092919050505061205c565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561085657600080fd5b506108b96004803603604081101561086d57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061208f565b6040518082815260200191505060405180910390f35b3480156108db57600080fd5b50610952600480360360808110156108f257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001909291905050506120b4565b005b34801561096057600080fd5b506109cd6004803603606081101561097757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050612236565b005b3480156109db57600080fd5b50610a28600480360360408110156109f257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506123af565b005b348015610a3657600080fd5b50610a8360048036036040811015610a4d57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061256a565b005b60003390508060046000600254815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060046000600160025401815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fb2fd402d6b838b10cf190139b9d4495eefcfea7543bc1056544d13732d82e6ac33600254600160025401604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001838152602001828152602001935050505060405180910390a16002806000828254019250508190555050565b60016020528060005260406000206000915054906101000a900460ff1681565b60086020528160005260406000206020528060005260406000206000915091509054906101000a900460ff1681565b600b602052816000526040600020602052806000526040600020600091509150508054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610cc65780601f10610c9b57610100808354040283529160200191610cc6565b820191906000526020600020905b815481529060010190602001808311610ca957829003601f168201915b505050505081565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1610610d6f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f696e76616c69642074726164696e67207061697200000000000000000000000081525060200191505060405180910390fd5b7ffaaebcb30b1b421f4f2ca7f2620e5add6a64532c087ee0646fd665a33d36fdf53385858585604051808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018381526020018281526020019550505050505060405180910390a150505050565b60065481565b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1610610ef7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f696e76616c69642074726164696e67207061697200000000000000000000000081525060200191505060405180910390fd5b6000339050600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610fbb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f6e6f7420612076616c696420736572766572000000000000000000000000000081525060200191505060405180910390fd5b600e60008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600085815260200190815260200160002060008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16611498576001600e60008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600086815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506001600d60008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000868152602001908152602001600020848460405180838380828437808301925050509250505090815260200160405180910390206000828254019250508190555060008414806113845750600054600d60008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600086815260200190815260200160002084846040518083838082843780830192505050925050509081526020016040518091039020541180156113835750600c60008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205484115b5b15611497578282600b60008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209190611414929190612c00565b5083600c60008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b5b505050505050565b60005481565b6000339050600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141561156f5734600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254019250508190555061162d565b60008390506115a18230858473ffffffffffffffffffffffffffffffffffffffff1661273b909392919063ffffffff16565b82600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282540192505081905550505b505050565b60076020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020154905083565b60025481565b6000339050600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16611766576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f6e6f7420612076616c696420736572766572000000000000000000000000000081525060200191505060405180910390fd5b6008600083815260200190815260200160002060008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615611837576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600f8152602001807f616c726561647920636f6e73656e74000000000000000000000000000000000081525060200191505060405180910390fd5b60016008600084815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506001600960008481526020019081526020016000206000828254019250508190555060005460096000848152602001908152602001600020541180156119045750600a600083815260200190815260200160002060009054906101000a900460ff16155b15611aab57611911612c80565b600760008481526020019081526020016000206040518060600160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820154815250509050806040015160056000836000015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000836020015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055506001600a600085815260200190815260200160002060006101000a81548160ff021916908315150217905550505b5050565b600c602052816000526040600020602052806000526040600020600091509150505481565b600033905081600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015611bcb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f6e6f7420656e6f7567682062616c616e6365000000000000000000000000000081525060200191505060405180910390fd5b81600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825403925050819055507f07c06144435b7d2bdccf9ee7e5a7022c63382ac7c3a0e14ed08b5969dedf0ecf838284604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060405180910390a1505050565b60035481565b8473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff1610611da0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f696e76616c69642074726164696e67207061697200000000000000000000000081525060200191505060405180910390fd5b60003390508073ffffffffffffffffffffffffffffffffffffffff166004600087815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611e79576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f756e617574686f72697a656420696e7075746d61736b0000000000000000000081525060200191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166004600086815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611f4d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f756e617574686f72697a656420696e7075746d61736b0000000000000000000081525060200191505060405180910390fd5b60016003600082825401925050819055507f2b4d91cd20cc8800407e3614b8466a6f0729ac3b1fa43d4e2b059ff5593cbae660035433898989898989604051808981526020018873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018581526020018481526020018381526020018281526020019850505050505050505060405180910390a150505050505050565b60046020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6005602052816000526040600020602052806000526040600020600091509150505481565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1610612155576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f696e76616c69642074726164696e67207061697200000000000000000000000081525060200191505060405180910390fd5b7fc33fbc9654f9c0dcfcbd829113bdb10afe95619bc0824bc5959ad82fd6952bd93385858585604051808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018381526020018281526020019550505050505060405180910390a150505050565b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16106122d7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f696e76616c69642074726164696e67207061697200000000000000000000000081525060200191505060405180910390fd5b7f2a1a4e62bda5b4987b9aa2f23ddbb29e434808f7a717452a3226bc15243c927733848484604051808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200194505050505060405180910390a1505050565b6000339050600160066000828254019250508190555060405180606001604052808473ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1681526020018381525060076000600654815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550604082015181600201559050507f4ef3cc4825a92c3b6922acc8a45152cc96ef48463e8ed500dacd5df9e29a67f3600654848385604051808581526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200194505050505060405180910390a1505050565b6000339050600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141561267a5781600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825403925050819055508073ffffffffffffffffffffffffffffffffffffffff166108fc839081150290604051600060405180830381858888f19350505050158015612674573d6000803e3d6000fd5b50612736565b60008390506126aa82848373ffffffffffffffffffffffffffffffffffffffff166128289092919063ffffffff16565b82600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282540392505081905550505b505050565b612822846323b872dd60e01b858585604051602401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506128e0565b50505050565b6128db8363a9059cbb60e01b8484604051602401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506128e0565b505050565b6060612942826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166129cf9092919063ffffffff16565b90506000815111156129ca5780806020019051602081101561296357600080fd5b81019080805190602001909291905050506129c9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a815260200180612cf3602a913960400191505060405180910390fd5b5b505050565b60606129de84846000856129e7565b90509392505050565b60606129f285612bed565b612a64576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000081525060200191505060405180910390fd5b600060608673ffffffffffffffffffffffffffffffffffffffff1685876040518082805190602001908083835b60208310612ab45780518252602082019150602081019050602083039250612a91565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d8060008114612b16576040519150601f19603f3d011682016040523d82523d6000602084013e612b1b565b606091505b50915091508115612b30578092505050612be5565b600081511115612b435780518082602001fd5b836040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015612baa578082015181840152602081019050612b8f565b50505050905090810190601f168015612bd75780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b949350505050565b600080823b905060008111915050919050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10612c4157803560ff1916838001178555612c6f565b82800160010185558215612c6f579182015b82811115612c6e578235825591602001919060010190612c53565b5b509050612c7c9190612ccd565b5090565b6040518060600160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b612cef91905b80821115612ceb576000816000905550600101612cd3565b5090565b9056fe5361666545524332303a204552433230206f7065726174696f6e20646964206e6f742073756363656564a2646970667358221220bebe15760418f65ee376d989f912bce38ca540c8fcb27c637d8a701c5aa3c07064736f6c634300060a0033"

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

// Prices is a free data retrieval call binding the contract method 0x1f312404.
//
// Solidity: function prices(address , address ) constant returns(string)
func (_HbSwap *HbSwapCaller) Prices(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _HbSwap.contract.Call(opts, out, "prices", arg0, arg1)
	return *ret0, err
}

// Prices is a free data retrieval call binding the contract method 0x1f312404.
//
// Solidity: function prices(address , address ) constant returns(string)
func (_HbSwap *HbSwapSession) Prices(arg0 common.Address, arg1 common.Address) (string, error) {
	return _HbSwap.Contract.Prices(&_HbSwap.CallOpts, arg0, arg1)
}

// Prices is a free data retrieval call binding the contract method 0x1f312404.
//
// Solidity: function prices(address , address ) constant returns(string)
func (_HbSwap *HbSwapCallerSession) Prices(arg0 common.Address, arg1 common.Address) (string, error) {
	return _HbSwap.Contract.Prices(&_HbSwap.CallOpts, arg0, arg1)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xcf6c62ea.
//
// Solidity: function addLiquidity(address _tokenA, address _tokenB, uint256 _amtA, uint256 _amtB) returns()
func (_HbSwap *HbSwapTransactor) AddLiquidity(opts *bind.TransactOpts, _tokenA common.Address, _tokenB common.Address, _amtA *big.Int, _amtB *big.Int) (*types.Transaction, error) {
	return _HbSwap.contract.Transact(opts, "addLiquidity", _tokenA, _tokenB, _amtA, _amtB)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xcf6c62ea.
//
// Solidity: function addLiquidity(address _tokenA, address _tokenB, uint256 _amtA, uint256 _amtB) returns()
func (_HbSwap *HbSwapSession) AddLiquidity(_tokenA common.Address, _tokenB common.Address, _amtA *big.Int, _amtB *big.Int) (*types.Transaction, error) {
	return _HbSwap.Contract.AddLiquidity(&_HbSwap.TransactOpts, _tokenA, _tokenB, _amtA, _amtB)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xcf6c62ea.
//
// Solidity: function addLiquidity(address _tokenA, address _tokenB, uint256 _amtA, uint256 _amtB) returns()
func (_HbSwap *HbSwapTransactorSession) AddLiquidity(_tokenA common.Address, _tokenB common.Address, _amtA *big.Int, _amtB *big.Int) (*types.Transaction, error) {
	return _HbSwap.Contract.AddLiquidity(&_HbSwap.TransactOpts, _tokenA, _tokenB, _amtA, _amtB)
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

// InitPool is a paid mutator transaction binding the contract method 0x2f8a68c5.
//
// Solidity: function initPool(address _tokenA, address _tokenB, uint256 _amtA, uint256 _amtB) returns()
func (_HbSwap *HbSwapTransactor) InitPool(opts *bind.TransactOpts, _tokenA common.Address, _tokenB common.Address, _amtA *big.Int, _amtB *big.Int) (*types.Transaction, error) {
	return _HbSwap.contract.Transact(opts, "initPool", _tokenA, _tokenB, _amtA, _amtB)
}

// InitPool is a paid mutator transaction binding the contract method 0x2f8a68c5.
//
// Solidity: function initPool(address _tokenA, address _tokenB, uint256 _amtA, uint256 _amtB) returns()
func (_HbSwap *HbSwapSession) InitPool(_tokenA common.Address, _tokenB common.Address, _amtA *big.Int, _amtB *big.Int) (*types.Transaction, error) {
	return _HbSwap.Contract.InitPool(&_HbSwap.TransactOpts, _tokenA, _tokenB, _amtA, _amtB)
}

// InitPool is a paid mutator transaction binding the contract method 0x2f8a68c5.
//
// Solidity: function initPool(address _tokenA, address _tokenB, uint256 _amtA, uint256 _amtB) returns()
func (_HbSwap *HbSwapTransactorSession) InitPool(_tokenA common.Address, _tokenB common.Address, _amtA *big.Int, _amtB *big.Int) (*types.Transaction, error) {
	return _HbSwap.Contract.InitPool(&_HbSwap.TransactOpts, _tokenA, _tokenB, _amtA, _amtB)
}

// InputMaskOwner is a paid mutator transaction binding the contract method 0xb7deef82.
//
// Solidity: function inputMaskOwner(uint256 ) returns(address)
func (_HbSwap *HbSwapTransactor) InputMaskOwner(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _HbSwap.contract.Transact(opts, "inputMaskOwner", arg0)
}

// InputMaskOwner is a paid mutator transaction binding the contract method 0xb7deef82.
//
// Solidity: function inputMaskOwner(uint256 ) returns(address)
func (_HbSwap *HbSwapSession) InputMaskOwner(arg0 *big.Int) (*types.Transaction, error) {
	return _HbSwap.Contract.InputMaskOwner(&_HbSwap.TransactOpts, arg0)
}

// InputMaskOwner is a paid mutator transaction binding the contract method 0xb7deef82.
//
// Solidity: function inputMaskOwner(uint256 ) returns(address)
func (_HbSwap *HbSwapTransactorSession) InputMaskOwner(arg0 *big.Int) (*types.Transaction, error) {
	return _HbSwap.Contract.InputMaskOwner(&_HbSwap.TransactOpts, arg0)
}

// LastUpdateSeq is a paid mutator transaction binding the contract method 0xa4bea6c8.
//
// Solidity: function lastUpdateSeq(address , address ) returns(uint256)
func (_HbSwap *HbSwapTransactor) LastUpdateSeq(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address) (*types.Transaction, error) {
	return _HbSwap.contract.Transact(opts, "lastUpdateSeq", arg0, arg1)
}

// LastUpdateSeq is a paid mutator transaction binding the contract method 0xa4bea6c8.
//
// Solidity: function lastUpdateSeq(address , address ) returns(uint256)
func (_HbSwap *HbSwapSession) LastUpdateSeq(arg0 common.Address, arg1 common.Address) (*types.Transaction, error) {
	return _HbSwap.Contract.LastUpdateSeq(&_HbSwap.TransactOpts, arg0, arg1)
}

// LastUpdateSeq is a paid mutator transaction binding the contract method 0xa4bea6c8.
//
// Solidity: function lastUpdateSeq(address , address ) returns(uint256)
func (_HbSwap *HbSwapTransactorSession) LastUpdateSeq(arg0 common.Address, arg1 common.Address) (*types.Transaction, error) {
	return _HbSwap.Contract.LastUpdateSeq(&_HbSwap.TransactOpts, arg0, arg1)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xd752fab2.
//
// Solidity: function removeLiquidity(address _tokenA, address _tokenB, uint256 _amt) returns()
func (_HbSwap *HbSwapTransactor) RemoveLiquidity(opts *bind.TransactOpts, _tokenA common.Address, _tokenB common.Address, _amt *big.Int) (*types.Transaction, error) {
	return _HbSwap.contract.Transact(opts, "removeLiquidity", _tokenA, _tokenB, _amt)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xd752fab2.
//
// Solidity: function removeLiquidity(address _tokenA, address _tokenB, uint256 _amt) returns()
func (_HbSwap *HbSwapSession) RemoveLiquidity(_tokenA common.Address, _tokenB common.Address, _amt *big.Int) (*types.Transaction, error) {
	return _HbSwap.Contract.RemoveLiquidity(&_HbSwap.TransactOpts, _tokenA, _tokenB, _amt)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xd752fab2.
//
// Solidity: function removeLiquidity(address _tokenA, address _tokenB, uint256 _amt) returns()
func (_HbSwap *HbSwapTransactorSession) RemoveLiquidity(_tokenA common.Address, _tokenB common.Address, _amt *big.Int) (*types.Transaction, error) {
	return _HbSwap.Contract.RemoveLiquidity(&_HbSwap.TransactOpts, _tokenA, _tokenB, _amt)
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

// TradeCnt is a paid mutator transaction binding the contract method 0xaf4170c4.
//
// Solidity: function tradeCnt() returns(uint256)
func (_HbSwap *HbSwapTransactor) TradeCnt(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HbSwap.contract.Transact(opts, "tradeCnt")
}

// TradeCnt is a paid mutator transaction binding the contract method 0xaf4170c4.
//
// Solidity: function tradeCnt() returns(uint256)
func (_HbSwap *HbSwapSession) TradeCnt() (*types.Transaction, error) {
	return _HbSwap.Contract.TradeCnt(&_HbSwap.TransactOpts)
}

// TradeCnt is a paid mutator transaction binding the contract method 0xaf4170c4.
//
// Solidity: function tradeCnt() returns(uint256)
func (_HbSwap *HbSwapTransactorSession) TradeCnt() (*types.Transaction, error) {
	return _HbSwap.Contract.TradeCnt(&_HbSwap.TransactOpts)
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

// UpdatePrice is a paid mutator transaction binding the contract method 0x3766dba7.
//
// Solidity: function updatePrice(address _tokenA, address _tokenB, uint256 _checkpointSeq, string _price) returns()
func (_HbSwap *HbSwapTransactor) UpdatePrice(opts *bind.TransactOpts, _tokenA common.Address, _tokenB common.Address, _checkpointSeq *big.Int, _price string) (*types.Transaction, error) {
	return _HbSwap.contract.Transact(opts, "updatePrice", _tokenA, _tokenB, _checkpointSeq, _price)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x3766dba7.
//
// Solidity: function updatePrice(address _tokenA, address _tokenB, uint256 _checkpointSeq, string _price) returns()
func (_HbSwap *HbSwapSession) UpdatePrice(_tokenA common.Address, _tokenB common.Address, _checkpointSeq *big.Int, _price string) (*types.Transaction, error) {
	return _HbSwap.Contract.UpdatePrice(&_HbSwap.TransactOpts, _tokenA, _tokenB, _checkpointSeq, _price)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x3766dba7.
//
// Solidity: function updatePrice(address _tokenA, address _tokenB, uint256 _checkpointSeq, string _price) returns()
func (_HbSwap *HbSwapTransactorSession) UpdatePrice(_tokenA common.Address, _tokenB common.Address, _checkpointSeq *big.Int, _price string) (*types.Transaction, error) {
	return _HbSwap.Contract.UpdatePrice(&_HbSwap.TransactOpts, _tokenA, _tokenB, _checkpointSeq, _price)
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

// HbSwapAddLiquidityIterator is returned from FilterAddLiquidity and is used to iterate over the raw logs and unpacked data for AddLiquidity events raised by the HbSwap contract.
type HbSwapAddLiquidityIterator struct {
	Event *HbSwapAddLiquidity // Event containing the contract specifics and raw log

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
func (it *HbSwapAddLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HbSwapAddLiquidity)
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
		it.Event = new(HbSwapAddLiquidity)
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
func (it *HbSwapAddLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HbSwapAddLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HbSwapAddLiquidity represents a AddLiquidity event raised by the HbSwap contract.
type HbSwapAddLiquidity struct {
	User   common.Address
	TokenA common.Address
	TokenB common.Address
	AmtA   *big.Int
	AmtB   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAddLiquidity is a free log retrieval operation binding the contract event 0xc33fbc9654f9c0dcfcbd829113bdb10afe95619bc0824bc5959ad82fd6952bd9.
//
// Solidity: event AddLiquidity(address user, address tokenA, address tokenB, uint256 amtA, uint256 amtB)
func (_HbSwap *HbSwapFilterer) FilterAddLiquidity(opts *bind.FilterOpts) (*HbSwapAddLiquidityIterator, error) {

	logs, sub, err := _HbSwap.contract.FilterLogs(opts, "AddLiquidity")
	if err != nil {
		return nil, err
	}
	return &HbSwapAddLiquidityIterator{contract: _HbSwap.contract, event: "AddLiquidity", logs: logs, sub: sub}, nil
}

// WatchAddLiquidity is a free log subscription operation binding the contract event 0xc33fbc9654f9c0dcfcbd829113bdb10afe95619bc0824bc5959ad82fd6952bd9.
//
// Solidity: event AddLiquidity(address user, address tokenA, address tokenB, uint256 amtA, uint256 amtB)
func (_HbSwap *HbSwapFilterer) WatchAddLiquidity(opts *bind.WatchOpts, sink chan<- *HbSwapAddLiquidity) (event.Subscription, error) {

	logs, sub, err := _HbSwap.contract.WatchLogs(opts, "AddLiquidity")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HbSwapAddLiquidity)
				if err := _HbSwap.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
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

// ParseAddLiquidity is a log parse operation binding the contract event 0xc33fbc9654f9c0dcfcbd829113bdb10afe95619bc0824bc5959ad82fd6952bd9.
//
// Solidity: event AddLiquidity(address user, address tokenA, address tokenB, uint256 amtA, uint256 amtB)
func (_HbSwap *HbSwapFilterer) ParseAddLiquidity(log types.Log) (*HbSwapAddLiquidity, error) {
	event := new(HbSwapAddLiquidity)
	if err := _HbSwap.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
		return nil, err
	}
	return event, nil
}

// HbSwapInitPoolIterator is returned from FilterInitPool and is used to iterate over the raw logs and unpacked data for InitPool events raised by the HbSwap contract.
type HbSwapInitPoolIterator struct {
	Event *HbSwapInitPool // Event containing the contract specifics and raw log

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
func (it *HbSwapInitPoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HbSwapInitPool)
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
		it.Event = new(HbSwapInitPool)
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
func (it *HbSwapInitPoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HbSwapInitPoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HbSwapInitPool represents a InitPool event raised by the HbSwap contract.
type HbSwapInitPool struct {
	User   common.Address
	TokenA common.Address
	TokenB common.Address
	AmtA   *big.Int
	AmtB   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterInitPool is a free log retrieval operation binding the contract event 0xfaaebcb30b1b421f4f2ca7f2620e5add6a64532c087ee0646fd665a33d36fdf5.
//
// Solidity: event InitPool(address user, address tokenA, address tokenB, uint256 amtA, uint256 amtB)
func (_HbSwap *HbSwapFilterer) FilterInitPool(opts *bind.FilterOpts) (*HbSwapInitPoolIterator, error) {

	logs, sub, err := _HbSwap.contract.FilterLogs(opts, "InitPool")
	if err != nil {
		return nil, err
	}
	return &HbSwapInitPoolIterator{contract: _HbSwap.contract, event: "InitPool", logs: logs, sub: sub}, nil
}

// WatchInitPool is a free log subscription operation binding the contract event 0xfaaebcb30b1b421f4f2ca7f2620e5add6a64532c087ee0646fd665a33d36fdf5.
//
// Solidity: event InitPool(address user, address tokenA, address tokenB, uint256 amtA, uint256 amtB)
func (_HbSwap *HbSwapFilterer) WatchInitPool(opts *bind.WatchOpts, sink chan<- *HbSwapInitPool) (event.Subscription, error) {

	logs, sub, err := _HbSwap.contract.WatchLogs(opts, "InitPool")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HbSwapInitPool)
				if err := _HbSwap.contract.UnpackLog(event, "InitPool", log); err != nil {
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

// ParseInitPool is a log parse operation binding the contract event 0xfaaebcb30b1b421f4f2ca7f2620e5add6a64532c087ee0646fd665a33d36fdf5.
//
// Solidity: event InitPool(address user, address tokenA, address tokenB, uint256 amtA, uint256 amtB)
func (_HbSwap *HbSwapFilterer) ParseInitPool(log types.Log) (*HbSwapInitPool, error) {
	event := new(HbSwapInitPool)
	if err := _HbSwap.contract.UnpackLog(event, "InitPool", log); err != nil {
		return nil, err
	}
	return event, nil
}

// HbSwapRemoveLiquidityIterator is returned from FilterRemoveLiquidity and is used to iterate over the raw logs and unpacked data for RemoveLiquidity events raised by the HbSwap contract.
type HbSwapRemoveLiquidityIterator struct {
	Event *HbSwapRemoveLiquidity // Event containing the contract specifics and raw log

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
func (it *HbSwapRemoveLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HbSwapRemoveLiquidity)
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
		it.Event = new(HbSwapRemoveLiquidity)
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
func (it *HbSwapRemoveLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HbSwapRemoveLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HbSwapRemoveLiquidity represents a RemoveLiquidity event raised by the HbSwap contract.
type HbSwapRemoveLiquidity struct {
	User   common.Address
	TokenA common.Address
	TokenB common.Address
	Amt    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidity is a free log retrieval operation binding the contract event 0x2a1a4e62bda5b4987b9aa2f23ddbb29e434808f7a717452a3226bc15243c9277.
//
// Solidity: event RemoveLiquidity(address user, address tokenA, address tokenB, uint256 amt)
func (_HbSwap *HbSwapFilterer) FilterRemoveLiquidity(opts *bind.FilterOpts) (*HbSwapRemoveLiquidityIterator, error) {

	logs, sub, err := _HbSwap.contract.FilterLogs(opts, "RemoveLiquidity")
	if err != nil {
		return nil, err
	}
	return &HbSwapRemoveLiquidityIterator{contract: _HbSwap.contract, event: "RemoveLiquidity", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidity is a free log subscription operation binding the contract event 0x2a1a4e62bda5b4987b9aa2f23ddbb29e434808f7a717452a3226bc15243c9277.
//
// Solidity: event RemoveLiquidity(address user, address tokenA, address tokenB, uint256 amt)
func (_HbSwap *HbSwapFilterer) WatchRemoveLiquidity(opts *bind.WatchOpts, sink chan<- *HbSwapRemoveLiquidity) (event.Subscription, error) {

	logs, sub, err := _HbSwap.contract.WatchLogs(opts, "RemoveLiquidity")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HbSwapRemoveLiquidity)
				if err := _HbSwap.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
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

// ParseRemoveLiquidity is a log parse operation binding the contract event 0x2a1a4e62bda5b4987b9aa2f23ddbb29e434808f7a717452a3226bc15243c9277.
//
// Solidity: event RemoveLiquidity(address user, address tokenA, address tokenB, uint256 amt)
func (_HbSwap *HbSwapFilterer) ParseRemoveLiquidity(log types.Log) (*HbSwapRemoveLiquidity, error) {
	event := new(HbSwapRemoveLiquidity)
	if err := _HbSwap.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
		return nil, err
	}
	return event, nil
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
	TradeSeq *big.Int
	User     common.Address
	TokenA   common.Address
	TokenB   common.Address
	IdxA     *big.Int
	IdxB     *big.Int
	MaskedA  *big.Int
	MaskedB  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTrade is a free log retrieval operation binding the contract event 0x2b4d91cd20cc8800407e3614b8466a6f0729ac3b1fa43d4e2b059ff5593cbae6.
//
// Solidity: event Trade(uint256 tradeSeq, address user, address tokenA, address tokenB, uint256 idxA, uint256 idxB, uint256 maskedA, uint256 maskedB)
func (_HbSwap *HbSwapFilterer) FilterTrade(opts *bind.FilterOpts) (*HbSwapTradeIterator, error) {

	logs, sub, err := _HbSwap.contract.FilterLogs(opts, "Trade")
	if err != nil {
		return nil, err
	}
	return &HbSwapTradeIterator{contract: _HbSwap.contract, event: "Trade", logs: logs, sub: sub}, nil
}

// WatchTrade is a free log subscription operation binding the contract event 0x2b4d91cd20cc8800407e3614b8466a6f0729ac3b1fa43d4e2b059ff5593cbae6.
//
// Solidity: event Trade(uint256 tradeSeq, address user, address tokenA, address tokenB, uint256 idxA, uint256 idxB, uint256 maskedA, uint256 maskedB)
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

// ParseTrade is a log parse operation binding the contract event 0x2b4d91cd20cc8800407e3614b8466a6f0729ac3b1fa43d4e2b059ff5593cbae6.
//
// Solidity: event Trade(uint256 tradeSeq, address user, address tokenA, address tokenB, uint256 idxA, uint256 idxB, uint256 maskedA, uint256 maskedB)
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

// HbSwapUpdatePriceIterator is returned from FilterUpdatePrice and is used to iterate over the raw logs and unpacked data for UpdatePrice events raised by the HbSwap contract.
type HbSwapUpdatePriceIterator struct {
	Event *HbSwapUpdatePrice // Event containing the contract specifics and raw log

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
func (it *HbSwapUpdatePriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HbSwapUpdatePrice)
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
		it.Event = new(HbSwapUpdatePrice)
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
func (it *HbSwapUpdatePriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HbSwapUpdatePriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HbSwapUpdatePrice represents a UpdatePrice event raised by the HbSwap contract.
type HbSwapUpdatePrice struct {
	TokenA common.Address
	TokenB common.Address
	Price  string
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUpdatePrice is a free log retrieval operation binding the contract event 0x768a9235e77345cf2e5d38042be1a15e329aa5d7b3a228aa9a881a0b919500b3.
//
// Solidity: event UpdatePrice(address tokenA, address tokenB, string price, uint256 time)
func (_HbSwap *HbSwapFilterer) FilterUpdatePrice(opts *bind.FilterOpts) (*HbSwapUpdatePriceIterator, error) {

	logs, sub, err := _HbSwap.contract.FilterLogs(opts, "UpdatePrice")
	if err != nil {
		return nil, err
	}
	return &HbSwapUpdatePriceIterator{contract: _HbSwap.contract, event: "UpdatePrice", logs: logs, sub: sub}, nil
}

// WatchUpdatePrice is a free log subscription operation binding the contract event 0x768a9235e77345cf2e5d38042be1a15e329aa5d7b3a228aa9a881a0b919500b3.
//
// Solidity: event UpdatePrice(address tokenA, address tokenB, string price, uint256 time)
func (_HbSwap *HbSwapFilterer) WatchUpdatePrice(opts *bind.WatchOpts, sink chan<- *HbSwapUpdatePrice) (event.Subscription, error) {

	logs, sub, err := _HbSwap.contract.WatchLogs(opts, "UpdatePrice")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HbSwapUpdatePrice)
				if err := _HbSwap.contract.UnpackLog(event, "UpdatePrice", log); err != nil {
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

// ParseUpdatePrice is a log parse operation binding the contract event 0x768a9235e77345cf2e5d38042be1a15e329aa5d7b3a228aa9a881a0b919500b3.
//
// Solidity: event UpdatePrice(address tokenA, address tokenB, string price, uint256 time)
func (_HbSwap *HbSwapFilterer) ParseUpdatePrice(log types.Log) (*HbSwapUpdatePrice, error) {
	event := new(HbSwapUpdatePrice)
	if err := _HbSwap.contract.UnpackLog(event, "UpdatePrice", log); err != nil {
		return nil, err
	}
	return event, nil
}
