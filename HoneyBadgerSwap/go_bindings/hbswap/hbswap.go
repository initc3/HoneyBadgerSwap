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
const HbSwapABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_servers\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amtA\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amtB\",\"type\":\"uint256\"}],\"name\":\"InitPool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"SecretDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"SecretWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tradeSeq\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"idxA\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"idxB\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maskedA\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maskedB\",\"type\":\"uint256\"}],\"name\":\"Trade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"idxA\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"idxB\",\"type\":\"uint256\"}],\"name\":\"TradePrep\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"constant\":true,\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"consentRecord\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inputmaskCnt\",\"constant\":true,\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"liquidityToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"prices\",\"constant\":true,\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"secretWithdrawCnt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"secretWithdrawMap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"serverNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"servers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tradeCnt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tradingPairs\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"updateTimes\",\"constant\":true,\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amtA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amtB\",\"type\":\"uint256\"}],\"name\":\"initPool\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amt\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amt\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amt\",\"type\":\"uint256\"}],\"name\":\"secretDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amt\",\"type\":\"uint256\"}],\"name\":\"secretWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_seq\",\"type\":\"uint256\"}],\"name\":\"consent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tradePrep\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_idxA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_idxB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maskedA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maskedB\",\"type\":\"uint256\"}],\"name\":\"trade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenB\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_price\",\"type\":\"string\"}],\"name\":\"updatePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// HbSwapBin is the compiled bytecode used for deploying new contracts.
var HbSwapBin = "0x60806040523480156200001157600080fd5b5060405162002a6038038062002a60833981810160405260408110156200003757600080fd5b81019080805160405193929190846401000000008211156200005857600080fd5b838201915060208201858111156200006f57600080fd5b82518660208202830111640100000000821117156200008d57600080fd5b8083526020830192505050908051906020019060200280838360005b83811015620000c6578082015181840152602081019050620000a9565b505050509050016040526020018051906020019092919050505060008090505b82518110156200016a576001600860008584815181106200010357fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508080600101915050620000e6565b508060048190555050506128dc80620001846000396000f3fe6080604052600436106101405760003560e01c806393910e66116100b6578063c23f001f1161006f578063c23f001f146107b1578063c719f4d514610836578063d3fd8987146108db578063d826f88f146109a1578063dee40595146109b8578063f3fef3a314610a1357610140565b806393910e66146105a75780639eef6ce8146105d2578063ade28aad1461060d578063af4170c414610668578063b72a213914610693578063bca8a7c11461072c57610140565b80632f8a68c5116101085780632f8a68c5146103ab5780633394dc6f1461042357806333d5e5d21461044e57806342cde4e81461047957806347e7ef24146104a45780637aa6fd65146104f257610140565b8063064d48101461014557806312ada8de1461015c5780631768af96146101c55780631dc1744d146102385780631f312404146102c1575b600080fd5b34801561015157600080fd5b5061015a610a6e565b005b34801561016857600080fd5b506101ab6004803603602081101561017f57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610afb565b604051808215151515815260200191505060405180910390f35b3480156101d157600080fd5b5061021e600480360360408110156101e857600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610b1b565b604051808215151515815260200191505060405180910390f35b34801561024457600080fd5b506102a76004803603604081101561025b57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610b4a565b604051808215151515815260200191505060405180910390f35b3480156102cd57600080fd5b50610330600480360360408110156102e457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610b79565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610370578082015181840152602081019050610355565b50505050905090810190601f16801561039d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610421600480360360808110156103c157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919080359060200190929190505050610c36565b005b34801561042f57600080fd5b506104386110b2565b6040518082815260200191505060405180910390f35b34801561045a57600080fd5b506104636110b8565b6040518082815260200191505060405180910390f35b34801561048557600080fd5b5061048e6110be565b6040518082815260200191505060405180910390f35b6104f0600480360360408110156104ba57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506110c4565b005b3480156104fe57600080fd5b5061052b6004803603602081101561051557600080fd5b8101908080359060200190929190505050611250565b604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060405180910390f35b3480156105b357600080fd5b506105bc6112ba565b6040518082815260200191505060405180910390f35b3480156105de57600080fd5b5061060b600480360360208110156105f557600080fd5b81019080803590602001909291905050506112c0565b005b34801561061957600080fd5b506106666004803603604081101561063057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506116cd565b005b34801561067457600080fd5b5061067d6118f2565b6040518082815260200191505060405180910390f35b34801561069f57600080fd5b5061072a600480360360c08110156106b657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019092919080359060200190929190803590602001909291905050506118f8565b005b34801561073857600080fd5b5061079b6004803603604081101561074f57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611ba2565b6040518082815260200191505060405180910390f35b3480156107bd57600080fd5b50610820600480360360408110156107d457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611bc7565b6040518082815260200191505060405180910390f35b34801561084257600080fd5b506108c56004803603606081101561085957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611bec565b6040518082815260200191505060405180910390f35b3480156108e757600080fd5b5061099f600480360360608110156108fe57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019064010000000081111561095b57600080fd5b82018360208201111561096d57600080fd5b8035906020019184600183028401116401000000008311171561098f57600080fd5b9091929391929390505050611c1e565b005b3480156109ad57600080fd5b506109b6611ecd565b005b3480156109c457600080fd5b50610a11600480360360408110156109db57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611ed7565b005b348015610a1f57600080fd5b50610a6c60048036036040811015610a3657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050612092565b005b7fb2fd402d6b838b10cf190139b9d4495eefcfea7543bc1056544d13732d82e6ac33600554600160055401604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001838152602001828152602001935050505060405180910390a16002600560008282540192505081905550565b60086020528060005260406000206000915054906101000a900460ff1681565b600c6020528160005260406000206020528060005260406000206000915091509054906101000a900460ff1681565b60006020528160005260406000206020528060005260406000206000915091509054906101000a900460ff1681565b6002602052816000526040600020602052806000526040600020600091509150508054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610c2e5780601f10610c0357610100808354040283529160200191610c2e565b820191906000526020600020905b815481529060010190602001808311610c1157829003601f168201915b505050505081565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1610610cd7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f696e76616c69642074726164696e67207061697200000000000000000000000081525060200191505060405180910390fd5b6000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615610dd3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f706f6f6c20616c726561647920696e697469617465640000000000000000000081525060200191505060405180910390fd5b6000339050600073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1614610e40576000859050610e3e8230868473ffffffffffffffffffffffffffffffffffffffff16612263909392919063ffffffff16565b505b600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1614610ea8576000849050610ea68230858473ffffffffffffffffffffffffffffffffffffffff16612263909392919063ffffffff16565b505b60016000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550610f47828402612350565b600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055507f52ce39f56a81bcdfe306f3ce8d3d56cebf8b1472b62e97786c89c841203edcf885858585604051808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200183815260200182815260200194505050505060405180910390a15050505050565b600a5481565b60075481565b60045481565b6000339050600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141561118d5734600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254019250508190555061124b565b60008390506111bf8230858473ffffffffffffffffffffffffffffffffffffffff16612263909392919063ffffffff16565b82600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282540192505081905550505b505050565b600b6020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020154905083565b60055481565b6000339050600860008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16611384576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f6e6f7420612076616c696420736572766572000000000000000000000000000081525060200191505060405180910390fd5b600c600083815260200190815260200160002060008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615611455576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600f8152602001807f616c726561647920636f6e73656e74000000000000000000000000000000000081525060200191505060405180910390fd5b6001600c600084815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506001600d600084815260200190815260200160002060008282540192505081905550600454600d6000848152602001908152602001600020541180156115225750600e600083815260200190815260200160002060009054906101000a900460ff16155b156116c95761152f61278a565b600b60008481526020019081526020016000206040518060600160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820154815250509050806040015160066000836000015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000836020015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055506001600e600085815260200190815260200160002060006101000a81548160ff021916908315150217905550505b5050565b600033905081600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410156117c4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f6e6f7420656e6f7567682062616c616e6365000000000000000000000000000081525060200191505060405180910390fd5b81600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825403925050819055507f07c06144435b7d2bdccf9ee7e5a7022c63382ac7c3a0e14ed08b5969dedf0ecf838284604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060405180910390a1505050565b60095481565b8473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff1610611999576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f696e76616c69642074726164696e67207061697200000000000000000000000081525060200191505060405180910390fd5b6000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16611a94576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f706f6f6c206e6f7420657869737400000000000000000000000000000000000081525060200191505060405180910390fd5b60016009600082825401925050819055507f2b4d91cd20cc8800407e3614b8466a6f0729ac3b1fa43d4e2b059ff5593cbae660095433888888888888604051808981526020018873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018581526020018481526020018381526020018281526020019850505050505050505060405180910390a1505050505050565b6003602052816000526040600020602052806000526040600020600091509150505481565b6006602052816000526040600020602052806000526040600020600091509150505481565b600160205282600052604060002060205281600052604060002060205280600052604060002060009250925050505481565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1610611cbf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f696e76616c69642074726164696e67207061697200000000000000000000000081525060200191505060405180910390fd5b6000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16611dba576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f706f6f6c206e6f7420657869737400000000000000000000000000000000000081525060200191505060405180910390fd5b8181600260008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209190611e459291906127d7565b5043600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050505050565b6000600581905550565b60003390506001600a6000828254019250508190555060405180606001604052808473ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff16815260200183815250600b6000600a54815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550604082015181600201559050507f4ef3cc4825a92c3b6922acc8a45152cc96ef48463e8ed500dacd5df9e29a67f3600a54848385604051808581526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200194505050505060405180910390a1505050565b6000339050600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614156121a25781600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825403925050819055508073ffffffffffffffffffffffffffffffffffffffff166108fc839081150290604051600060405180830381858888f1935050505015801561219c573d6000803e3d6000fd5b5061225e565b60008390506121d282848373ffffffffffffffffffffffffffffffffffffffff166123b29092919063ffffffff16565b82600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282540392505081905550505b505050565b61234a846323b872dd60e01b858585604051602401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505061246a565b50505050565b6000600382111561239f57819050600060016002848161236c57fe5b040190505b818110156123995780915060028182858161238857fe5b04018161239157fe5b049050612371565b506123ad565b600082146123ac57600190505b5b919050565b6124658363a9059cbb60e01b8484604051602401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505061246a565b505050565b60606124cc826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166125599092919063ffffffff16565b9050600081511115612554578080602001905160208110156124ed57600080fd5b8101908080519060200190929190505050612553576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a81526020018061287d602a913960400191505060405180910390fd5b5b505050565b60606125688484600085612571565b90509392505050565b606061257c85612777565b6125ee576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000081525060200191505060405180910390fd5b600060608673ffffffffffffffffffffffffffffffffffffffff1685876040518082805190602001908083835b6020831061263e578051825260208201915060208101905060208303925061261b565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d80600081146126a0576040519150601f19603f3d011682016040523d82523d6000602084013e6126a5565b606091505b509150915081156126ba57809250505061276f565b6000815111156126cd5780518082602001fd5b836040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015612734578082015181840152602081019050612719565b50505050905090810190601f1680156127615780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b949350505050565b600080823b905060008111915050919050565b6040518060600160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061281857803560ff1916838001178555612846565b82800160010185558215612846579182015b8281111561284557823582559160200191906001019061282a565b5b5090506128539190612857565b5090565b61287991905b8082111561287557600081600090555060010161285d565b5090565b9056fe5361666545524332303a204552433230206f7065726174696f6e20646964206e6f742073756363656564a26469706673582212209d524877ffc9dcb253aadd248955be0ac059e556d58a67239aba9a4ba56298ba64736f6c634300060a0033"

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

// UpdateTimes is a free data retrieval call binding the contract method 0xbca8a7c1.
//
// Solidity: function updateTimes(address , address ) constant returns(uint256)
func (_HbSwap *HbSwapCaller) UpdateTimes(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HbSwap.contract.Call(opts, out, "updateTimes", arg0, arg1)
	return *ret0, err
}

// UpdateTimes is a free data retrieval call binding the contract method 0xbca8a7c1.
//
// Solidity: function updateTimes(address , address ) constant returns(uint256)
func (_HbSwap *HbSwapSession) UpdateTimes(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _HbSwap.Contract.UpdateTimes(&_HbSwap.CallOpts, arg0, arg1)
}

// UpdateTimes is a free data retrieval call binding the contract method 0xbca8a7c1.
//
// Solidity: function updateTimes(address , address ) constant returns(uint256)
func (_HbSwap *HbSwapCallerSession) UpdateTimes(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _HbSwap.Contract.UpdateTimes(&_HbSwap.CallOpts, arg0, arg1)
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

// LiquidityToken is a paid mutator transaction binding the contract method 0xc719f4d5.
//
// Solidity: function liquidityToken(address , address , address ) returns(uint256)
func (_HbSwap *HbSwapTransactor) LiquidityToken(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 common.Address) (*types.Transaction, error) {
	return _HbSwap.contract.Transact(opts, "liquidityToken", arg0, arg1, arg2)
}

// LiquidityToken is a paid mutator transaction binding the contract method 0xc719f4d5.
//
// Solidity: function liquidityToken(address , address , address ) returns(uint256)
func (_HbSwap *HbSwapSession) LiquidityToken(arg0 common.Address, arg1 common.Address, arg2 common.Address) (*types.Transaction, error) {
	return _HbSwap.Contract.LiquidityToken(&_HbSwap.TransactOpts, arg0, arg1, arg2)
}

// LiquidityToken is a paid mutator transaction binding the contract method 0xc719f4d5.
//
// Solidity: function liquidityToken(address , address , address ) returns(uint256)
func (_HbSwap *HbSwapTransactorSession) LiquidityToken(arg0 common.Address, arg1 common.Address, arg2 common.Address) (*types.Transaction, error) {
	return _HbSwap.Contract.LiquidityToken(&_HbSwap.TransactOpts, arg0, arg1, arg2)
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

// TradingPairs is a paid mutator transaction binding the contract method 0x1dc1744d.
//
// Solidity: function tradingPairs(address , address ) returns(bool)
func (_HbSwap *HbSwapTransactor) TradingPairs(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address) (*types.Transaction, error) {
	return _HbSwap.contract.Transact(opts, "tradingPairs", arg0, arg1)
}

// TradingPairs is a paid mutator transaction binding the contract method 0x1dc1744d.
//
// Solidity: function tradingPairs(address , address ) returns(bool)
func (_HbSwap *HbSwapSession) TradingPairs(arg0 common.Address, arg1 common.Address) (*types.Transaction, error) {
	return _HbSwap.Contract.TradingPairs(&_HbSwap.TransactOpts, arg0, arg1)
}

// TradingPairs is a paid mutator transaction binding the contract method 0x1dc1744d.
//
// Solidity: function tradingPairs(address , address ) returns(bool)
func (_HbSwap *HbSwapTransactorSession) TradingPairs(arg0 common.Address, arg1 common.Address) (*types.Transaction, error) {
	return _HbSwap.Contract.TradingPairs(&_HbSwap.TransactOpts, arg0, arg1)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0xd3fd8987.
//
// Solidity: function updatePrice(address _tokenA, address _tokenB, string _price) returns()
func (_HbSwap *HbSwapTransactor) UpdatePrice(opts *bind.TransactOpts, _tokenA common.Address, _tokenB common.Address, _price string) (*types.Transaction, error) {
	return _HbSwap.contract.Transact(opts, "updatePrice", _tokenA, _tokenB, _price)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0xd3fd8987.
//
// Solidity: function updatePrice(address _tokenA, address _tokenB, string _price) returns()
func (_HbSwap *HbSwapSession) UpdatePrice(_tokenA common.Address, _tokenB common.Address, _price string) (*types.Transaction, error) {
	return _HbSwap.Contract.UpdatePrice(&_HbSwap.TransactOpts, _tokenA, _tokenB, _price)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0xd3fd8987.
//
// Solidity: function updatePrice(address _tokenA, address _tokenB, string _price) returns()
func (_HbSwap *HbSwapTransactorSession) UpdatePrice(_tokenA common.Address, _tokenB common.Address, _price string) (*types.Transaction, error) {
	return _HbSwap.Contract.UpdatePrice(&_HbSwap.TransactOpts, _tokenA, _tokenB, _price)
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
	TokenA common.Address
	TokenB common.Address
	AmtA   *big.Int
	AmtB   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterInitPool is a free log retrieval operation binding the contract event 0x52ce39f56a81bcdfe306f3ce8d3d56cebf8b1472b62e97786c89c841203edcf8.
//
// Solidity: event InitPool(address tokenA, address tokenB, uint256 amtA, uint256 amtB)
func (_HbSwap *HbSwapFilterer) FilterInitPool(opts *bind.FilterOpts) (*HbSwapInitPoolIterator, error) {

	logs, sub, err := _HbSwap.contract.FilterLogs(opts, "InitPool")
	if err != nil {
		return nil, err
	}
	return &HbSwapInitPoolIterator{contract: _HbSwap.contract, event: "InitPool", logs: logs, sub: sub}, nil
}

// WatchInitPool is a free log subscription operation binding the contract event 0x52ce39f56a81bcdfe306f3ce8d3d56cebf8b1472b62e97786c89c841203edcf8.
//
// Solidity: event InitPool(address tokenA, address tokenB, uint256 amtA, uint256 amtB)
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

// ParseInitPool is a log parse operation binding the contract event 0x52ce39f56a81bcdfe306f3ce8d3d56cebf8b1472b62e97786c89c841203edcf8.
//
// Solidity: event InitPool(address tokenA, address tokenB, uint256 amtA, uint256 amtB)
func (_HbSwap *HbSwapFilterer) ParseInitPool(log types.Log) (*HbSwapInitPool, error) {
	event := new(HbSwapInitPool)
	if err := _HbSwap.contract.UnpackLog(event, "InitPool", log); err != nil {
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
