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
const HbSwapABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_servers\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"idxA\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"idxB\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maskedAmtA\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maskedAmtB\",\"type\":\"uint256\"}],\"name\":\"AddLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amtA\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amtB\",\"type\":\"uint256\"}],\"name\":\"InitPool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"inpusMaskIndexes\",\"type\":\"uint256[]\"}],\"name\":\"InputMask\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maskedAmt\",\"type\":\"uint256\"}],\"name\":\"RemoveLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"SecretDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"SecretWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tradeSeq\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"idxA\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"idxB\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maskedAmtA\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maskedAmtB\",\"type\":\"uint256\"}],\"name\":\"Trade\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"Decimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Fp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"consentRecord\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"inputMaskOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"inputmaskCnt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastUpdateSeq\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"prices\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"publicBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"secretWithdrawCnt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"secretWithdrawMap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"servers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tradeCnt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amt\",\"type\":\"uint256\"}],\"name\":\"publicDeposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amt\",\"type\":\"uint256\"}],\"name\":\"secretDeposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amt\",\"type\":\"uint256\"}],\"name\":\"secretWithdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_seq\",\"type\":\"uint256\"}],\"name\":\"consent\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amt\",\"type\":\"uint256\"}],\"name\":\"publicWithdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_num\",\"type\":\"uint256\"}],\"name\":\"reserveInput\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amtA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amtB\",\"type\":\"uint256\"}],\"name\":\"initPool\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_idxA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_idxB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maskedAmtA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maskedAmtB\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_idx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maskedAmt\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_idxA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_idxB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maskedA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maskedB\",\"type\":\"uint256\"}],\"name\":\"trade\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_checkpointSeq\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_price\",\"type\":\"string\"}],\"name\":\"updatePrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenB\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_servers\",\"type\":\"address[]\"}],\"name\":\"resetPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"resetBalance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// HbSwapBin is the compiled bytecode used for deploying new contracts.
var HbSwapBin = "0x60806040523480156200001157600080fd5b506040516200396a3803806200396a833981810160405260408110156200003757600080fd5b81019080805160405193929190846401000000008211156200005857600080fd5b838201915060208201858111156200006f57600080fd5b82518660208202830111640100000000821117156200008d57600080fd5b8083526020830192505050908051906020019060200280838360005b83811015620000c6578082015181840152602081019050620000a9565b505050509050016040526020018051906020019092919050505060008090505b8251811015620001695760018060008584815181106200010257fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508080600101915050620000e6565b508060008190555050506137e780620001836000396000f3fe6080604052600436106101815760003560e01c80639eef6ce8116100d1578063b7deef821161008a578063ca7b58ad11610064578063ca7b58ad14610b95578063cc5bd7fa14610c06578063dee4059514610c41578063e7e8927614610c9c57610181565b8063b7deef8214610a94578063bc94e7cf14610b0f578063c6ceb50b14610b6a57610181565b80639eef6ce81461088a578063a45b6ce6146108c5578063a4bea6c8146108f0578063ade28aad14610975578063af4170c4146109d0578063b72a2139146109fb57610181565b80633766dba71161013e5780637aa6fd65116101185780637aa6fd651461065757806393910e661461070c57806393f54644146107375780639450fb261461078557610181565b80633766dba7146104955780633ee7f9d3146105a757806342cde4e81461062c57610181565b806312ada8de146101865780631768af96146101ef5780631f312404146102625780632f8a68c51461034c5780633351733f146103d15780633394dc6f1461046a575b600080fd5b34801561019257600080fd5b506101d5600480360360208110156101a957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610d21565b604051808215151515815260200191505060405180910390f35b3480156101fb57600080fd5b506102486004803603604081101561021257600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610d41565b604051808215151515815260200191505060405180910390f35b34801561026e57600080fd5b506102d16004803603604081101561028557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610d70565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156103115780820151818401526020810190506102f6565b50505050905090810190601f16801561033e5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561035857600080fd5b506103cf6004803603608081101561036f57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919080359060200190929190505050610e2d565b005b3480156103dd57600080fd5b50610468600480360360c08110156103f457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001909291908035906020019092919080359060200190929190505050611037565b005b34801561047657600080fd5b5061047f611379565b6040518082815260200191505060405180910390f35b3480156104a157600080fd5b506105a5600480360360808110156104b857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019064010000000081111561051f57600080fd5b82018360208201111561053157600080fd5b8035906020019184600183028401116401000000008311171561055357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929050505061137f565b005b3480156105b357600080fd5b5061062a600480360360808110156105ca57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001909291905050506119b3565b005b34801561063857600080fd5b50610641611c0f565b6040518082815260200191505060405180910390f35b34801561066357600080fd5b506106906004803603602081101561067a57600080fd5b8101908080359060200190929190505050611c15565b604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060405180910390f35b34801561071857600080fd5b50610721611c7f565b6040518082815260200191505060405180910390f35b6107836004803603604081101561074d57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611c85565b005b34801561079157600080fd5b50610888600480360360608110156107a857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019064010000000081111561080557600080fd5b82018360208201111561081757600080fd5b8035906020019184602083028401116401000000008311171561083957600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050509192919290505050611e89565b005b34801561089657600080fd5b506108c3600480360360208110156108ad57600080fd5b8101908080359060200190929190505050612192565b005b3480156108d157600080fd5b506108da61259f565b6040518082815260200191505060405180910390f35b3480156108fc57600080fd5b5061095f6004803603604081101561091357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506125a6565b6040518082815260200191505060405180910390f35b34801561098157600080fd5b506109ce6004803603604081101561099857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506125cb565b005b3480156109dc57600080fd5b506109e56127fc565b6040518082815260200191505060405180910390f35b348015610a0757600080fd5b50610a92600480360360c0811015610a1e57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001909291908035906020019092919080359060200190929190505050612802565b005b348015610aa057600080fd5b50610acd60048036036020811015610ab757600080fd5b8101908080359060200190929190505050612b5f565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015610b1b57600080fd5b50610b6860048036036040811015610b3257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050612b92565b005b348015610b7657600080fd5b50610b7f612df5565b6040518082815260200191505060405180910390f35b348015610ba157600080fd5b50610c0460048036036040811015610bb857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050612dfa565b005b348015610c1257600080fd5b50610c3f60048036036020811015610c2957600080fd5b8101908080359060200190929190505050612e80565b005b348015610c4d57600080fd5b50610c9a60048036036040811015610c6457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050612fd3565b005b348015610ca857600080fd5b50610d0b60048036036040811015610cbf57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050613204565b6040518082815260200191505060405180910390f35b60016020528060005260406000206000915054906101000a900460ff1681565b60056020528160005260406000206020528060005260406000206000915091509054906101000a900460ff1681565b600b602052816000526040600020602052806000526040600020600091509150508054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610e255780601f10610dfa57610100808354040283529160200191610e25565b820191906000526020600020905b815481529060010190602001808311610e0857829003601f168201915b505050505081565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1610610ece576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f696e76616c69642074726164696e67207061697200000000000000000000000081525060200191505060405180910390fd5b600082118015610ede5750600081115b610f50576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f696e76616c696420616d6f756e7400000000000000000000000000000000000081525060200191505060405180910390fd5b60003390507ffaaebcb30b1b421f4f2ca7f2620e5add6a64532c087ee0646fd665a33d36fdf53386868686604051808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018381526020018281526020019550505050505060405180910390a15050505050565b8473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff16106110d8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f696e76616c69642074726164696e67207061697200000000000000000000000081525060200191505060405180910390fd5b60003390508073ffffffffffffffffffffffffffffffffffffffff166009600087815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146111b1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f756e617574686f72697a656420696e7075746d61736b0000000000000000000081525060200191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166009600086815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611285576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f756e617574686f72697a656420696e7075746d61736b0000000000000000000081525060200191505060405180910390fd5b7fec7d4752dd44bf7fc59045c9d80163de2a1b9dbd9032d11cb1156f7f867c641181888888888888604051808873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200185815260200184815260200183815260200182815260200197505050505050505060405180910390a150505050505050565b60035481565b6000339050600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16611443576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f6e6f7420612076616c696420736572766572000000000000000000000000000081525060200191505060405180910390fd5b600e60008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600084815260200190815260200160002060008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615806115275750600083145b156119ac576001600e60008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600085815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506001600d60008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000858152602001908152602001600020836040518082805190602001908083835b602083106116d257805182526020820191506020810190506020830392506116af565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060008282540192505081905550600054600d60008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000858152602001908152602001600020836040518082805190602001908083835b602083106117d657805182526020820191506020810190506020830392506117b3565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020541180156118935750600c60008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548310155b156119ab5781600b60008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209080519060200190611928929190613696565b5082600c60008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b5b5050505050565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1610611a54576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f696e76616c69642074726164696e67207061697200000000000000000000000081525060200191505060405180910390fd5b60003390508073ffffffffffffffffffffffffffffffffffffffff166009600085815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611b2d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f756e617574686f72697a656420696e7075746d61736b0000000000000000000081525060200191505060405180910390fd5b7fa8dbaaebbb025c88e9e34c84635cd8238043556e9af43fb161508c898a8e1ef98186868686604051808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018381526020018281526020019550505050505060405180910390a15050505050565b60005481565b60046020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020154905083565b60085481565b600033905060008211611d00576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f696e76616c696420616d6f756e7400000000000000000000000000000000000081525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415611db7576001820262010000340214611db2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f616d6f756e7473206e6f74206d6174636800000000000000000000000000000081525060200191505060405180910390fd5b611dfa565b6000839050611df882306001620100008781611dcf57fe5b04028473ffffffffffffffffffffffffffffffffffffffff16613229909392919063ffffffff16565b505b81600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282540192505081905550505050565b60405180602001604052806000815250600b60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209080519060200190611f28929190613696565b506000600c60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555060008090505b81518110156120bd576000600e60008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000808152602001908152602001600020600084848151811061205757fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508080600101915050611fb1565b506000600d60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600080815260200190815260200160002060405180807f302e3000000000000000000000000000000000000000000000000000000000008152506003019050908152602001604051809103902081905550505050565b6000339050600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16612256576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f6e6f7420612076616c696420736572766572000000000000000000000000000081525060200191505060405180910390fd5b6005600083815260200190815260200160002060008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615612327576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600f8152602001807f616c726561647920636f6e73656e74000000000000000000000000000000000081525060200191505060405180910390fd5b60016005600084815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506001600660008481526020019081526020016000206000828254019250508190555060005460066000848152602001908152602001600020541180156123f457506007600083815260200190815260200160002060009054906101000a900460ff16155b1561259b57612401613716565b600460008481526020019081526020016000206040518060600160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820154815250509050806040015160026000836000015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000836020015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254019250508190555060016007600085815260200190815260200160002060006101000a81548160ff021916908315150217905550505b5050565b6201000081565b600c602052816000526040600020602052806000526040600020600091509150505481565b600033905060008211801561265c575081600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410155b6126ce576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f696e76616c696420616d6f756e7400000000000000000000000000000000000081525060200191505060405180910390fd5b81600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825403925050819055507f07c06144435b7d2bdccf9ee7e5a7022c63382ac7c3a0e14ed08b5969dedf0ecf838284604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060405180910390a1505050565b600a5481565b8473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff16106128a3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f696e76616c69642074726164696e67207061697200000000000000000000000081525060200191505060405180910390fd5b60003390508073ffffffffffffffffffffffffffffffffffffffff166009600087815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461297c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f756e617574686f72697a656420696e7075746d61736b0000000000000000000081525060200191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166009600086815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614612a50576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f756e617574686f72697a656420696e7075746d61736b0000000000000000000081525060200191505060405180910390fd5b6001600a600082825401925050819055507f2b4d91cd20cc8800407e3614b8466a6f0729ac3b1fa43d4e2b059ff5593cbae6600a5482898989898989604051808981526020018873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018581526020018481526020018381526020018281526020019850505050505050505060405180910390a150505050505050565b60096020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000339050600082118015612c23575081600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410155b612c95576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f696e76616c696420616d6f756e7400000000000000000000000000000000000081525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415612d25578073ffffffffffffffffffffffffffffffffffffffff166108fc6001620100008581612cf257fe5b04029081150290604051600060405180830381858888f19350505050158015612d1f573d6000803e3d6000fd5b50612d66565b6000839050612d64826001620100008681612d3c57fe5b04028373ffffffffffffffffffffffffffffffffffffffff1661332f9092919063ffffffff16565b505b81600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282540392505081905550505050565b600181565b6000600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050565b6000339050606082604051908082528060200260200182016040528015612eb65781602001602082028038833980820191505090505b50905060008090505b83811015612f5557600060086000815480929190600101919050559050836009600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080838381518110612f3b57fe5b602002602001018181525050508080600101915050612ebf565b507fff742b54a62e49c065680fc16aa208f37f60ded27e03a3d19e6088ba0c7b41d4816040518080602001828103825283818151815260200191508051906020019060200280838360005b83811015612fbb578082015181840152602081019050612fa0565b505050509050019250505060405180910390a1505050565b60003390506000821161304e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f696e76616c696420616d6f756e7400000000000000000000000000000000000081525060200191505060405180910390fd5b600160036000828254019250508190555060405180606001604052808473ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1681526020018381525060046000600354815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550604082015181600201559050507f4ef3cc4825a92c3b6922acc8a45152cc96ef48463e8ed500dacd5df9e29a67f3600354848385604051808581526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200194505050505060405180910390a1505050565b6002602052816000526040600020602052806000526040600020600091509150505481565b613329848573ffffffffffffffffffffffffffffffffffffffff166323b872dd905060e01b858585604051602401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050613400565b50505050565b6133fb838473ffffffffffffffffffffffffffffffffffffffff1663a9059cbb905060e01b8484604051602401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050613400565b505050565b61341f8273ffffffffffffffffffffffffffffffffffffffff1661364b565b613491576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f5361666545524332303a2063616c6c20746f206e6f6e2d636f6e74726163740081525060200191505060405180910390fd5b600060608373ffffffffffffffffffffffffffffffffffffffff16836040518082805190602001908083835b602083106134e057805182526020820191506020810190506020830392506134bd565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114613542576040519150601f19603f3d011682016040523d82523d6000602084013e613547565b606091505b5091509150816135bf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c656481525060200191505060405180910390fd5b600081511115613645578080602001905160208110156135de57600080fd5b8101908080519060200190929190505050613644576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a815260200180613789602a913960400191505060405180910390fd5b5b50505050565b60008060007fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47060001b9050833f915080821415801561368d57506000801b8214155b92505050919050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106136d757805160ff1916838001178555613705565b82800160010185558215613705579182015b828111156137045782518255916020019190600101906136e9565b5b5090506137129190613763565b5090565b6040518060600160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b61378591905b80821115613781576000816000905550600101613769565b5090565b9056fe5361666545524332303a204552433230206f7065726174696f6e20646964206e6f742073756363656564a265627a7a72315820822d2d89c121eafc100a03f84ca1756f66a19afcc16c7fe2907c725b81421d2564736f6c63430005110032"

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

// Decimals is a free data retrieval call binding the contract method 0xc6ceb50b.
//
// Solidity: function Decimals() constant returns(uint256)
func (_HbSwap *HbSwapCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HbSwap.contract.Call(opts, out, "Decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0xc6ceb50b.
//
// Solidity: function Decimals() constant returns(uint256)
func (_HbSwap *HbSwapSession) Decimals() (*big.Int, error) {
	return _HbSwap.Contract.Decimals(&_HbSwap.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0xc6ceb50b.
//
// Solidity: function Decimals() constant returns(uint256)
func (_HbSwap *HbSwapCallerSession) Decimals() (*big.Int, error) {
	return _HbSwap.Contract.Decimals(&_HbSwap.CallOpts)
}

// Fp is a free data retrieval call binding the contract method 0xa45b6ce6.
//
// Solidity: function Fp() constant returns(uint256)
func (_HbSwap *HbSwapCaller) Fp(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HbSwap.contract.Call(opts, out, "Fp")
	return *ret0, err
}

// Fp is a free data retrieval call binding the contract method 0xa45b6ce6.
//
// Solidity: function Fp() constant returns(uint256)
func (_HbSwap *HbSwapSession) Fp() (*big.Int, error) {
	return _HbSwap.Contract.Fp(&_HbSwap.CallOpts)
}

// Fp is a free data retrieval call binding the contract method 0xa45b6ce6.
//
// Solidity: function Fp() constant returns(uint256)
func (_HbSwap *HbSwapCallerSession) Fp() (*big.Int, error) {
	return _HbSwap.Contract.Fp(&_HbSwap.CallOpts)
}

// ConsentRecord is a free data retrieval call binding the contract method 0x1768af96.
//
// Solidity: function consentRecord(uint256 , address ) constant returns(bool)
func (_HbSwap *HbSwapCaller) ConsentRecord(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _HbSwap.contract.Call(opts, out, "consentRecord", arg0, arg1)
	return *ret0, err
}

// ConsentRecord is a free data retrieval call binding the contract method 0x1768af96.
//
// Solidity: function consentRecord(uint256 , address ) constant returns(bool)
func (_HbSwap *HbSwapSession) ConsentRecord(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _HbSwap.Contract.ConsentRecord(&_HbSwap.CallOpts, arg0, arg1)
}

// ConsentRecord is a free data retrieval call binding the contract method 0x1768af96.
//
// Solidity: function consentRecord(uint256 , address ) constant returns(bool)
func (_HbSwap *HbSwapCallerSession) ConsentRecord(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _HbSwap.Contract.ConsentRecord(&_HbSwap.CallOpts, arg0, arg1)
}

// InputMaskOwner is a free data retrieval call binding the contract method 0xb7deef82.
//
// Solidity: function inputMaskOwner(uint256 ) constant returns(address)
func (_HbSwap *HbSwapCaller) InputMaskOwner(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HbSwap.contract.Call(opts, out, "inputMaskOwner", arg0)
	return *ret0, err
}

// InputMaskOwner is a free data retrieval call binding the contract method 0xb7deef82.
//
// Solidity: function inputMaskOwner(uint256 ) constant returns(address)
func (_HbSwap *HbSwapSession) InputMaskOwner(arg0 *big.Int) (common.Address, error) {
	return _HbSwap.Contract.InputMaskOwner(&_HbSwap.CallOpts, arg0)
}

// InputMaskOwner is a free data retrieval call binding the contract method 0xb7deef82.
//
// Solidity: function inputMaskOwner(uint256 ) constant returns(address)
func (_HbSwap *HbSwapCallerSession) InputMaskOwner(arg0 *big.Int) (common.Address, error) {
	return _HbSwap.Contract.InputMaskOwner(&_HbSwap.CallOpts, arg0)
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

// LastUpdateSeq is a free data retrieval call binding the contract method 0xa4bea6c8.
//
// Solidity: function lastUpdateSeq(address , address ) constant returns(uint256)
func (_HbSwap *HbSwapCaller) LastUpdateSeq(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HbSwap.contract.Call(opts, out, "lastUpdateSeq", arg0, arg1)
	return *ret0, err
}

// LastUpdateSeq is a free data retrieval call binding the contract method 0xa4bea6c8.
//
// Solidity: function lastUpdateSeq(address , address ) constant returns(uint256)
func (_HbSwap *HbSwapSession) LastUpdateSeq(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _HbSwap.Contract.LastUpdateSeq(&_HbSwap.CallOpts, arg0, arg1)
}

// LastUpdateSeq is a free data retrieval call binding the contract method 0xa4bea6c8.
//
// Solidity: function lastUpdateSeq(address , address ) constant returns(uint256)
func (_HbSwap *HbSwapCallerSession) LastUpdateSeq(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _HbSwap.Contract.LastUpdateSeq(&_HbSwap.CallOpts, arg0, arg1)
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

// PublicBalance is a free data retrieval call binding the contract method 0xe7e89276.
//
// Solidity: function publicBalance(address , address ) constant returns(uint256)
func (_HbSwap *HbSwapCaller) PublicBalance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HbSwap.contract.Call(opts, out, "publicBalance", arg0, arg1)
	return *ret0, err
}

// PublicBalance is a free data retrieval call binding the contract method 0xe7e89276.
//
// Solidity: function publicBalance(address , address ) constant returns(uint256)
func (_HbSwap *HbSwapSession) PublicBalance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _HbSwap.Contract.PublicBalance(&_HbSwap.CallOpts, arg0, arg1)
}

// PublicBalance is a free data retrieval call binding the contract method 0xe7e89276.
//
// Solidity: function publicBalance(address , address ) constant returns(uint256)
func (_HbSwap *HbSwapCallerSession) PublicBalance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _HbSwap.Contract.PublicBalance(&_HbSwap.CallOpts, arg0, arg1)
}

// SecretWithdrawCnt is a free data retrieval call binding the contract method 0x3394dc6f.
//
// Solidity: function secretWithdrawCnt() constant returns(uint256)
func (_HbSwap *HbSwapCaller) SecretWithdrawCnt(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HbSwap.contract.Call(opts, out, "secretWithdrawCnt")
	return *ret0, err
}

// SecretWithdrawCnt is a free data retrieval call binding the contract method 0x3394dc6f.
//
// Solidity: function secretWithdrawCnt() constant returns(uint256)
func (_HbSwap *HbSwapSession) SecretWithdrawCnt() (*big.Int, error) {
	return _HbSwap.Contract.SecretWithdrawCnt(&_HbSwap.CallOpts)
}

// SecretWithdrawCnt is a free data retrieval call binding the contract method 0x3394dc6f.
//
// Solidity: function secretWithdrawCnt() constant returns(uint256)
func (_HbSwap *HbSwapCallerSession) SecretWithdrawCnt() (*big.Int, error) {
	return _HbSwap.Contract.SecretWithdrawCnt(&_HbSwap.CallOpts)
}

// SecretWithdrawMap is a free data retrieval call binding the contract method 0x7aa6fd65.
//
// Solidity: function secretWithdrawMap(uint256 ) constant returns(address token, address user, uint256 amt)
func (_HbSwap *HbSwapCaller) SecretWithdrawMap(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Token common.Address
	User  common.Address
	Amt   *big.Int
}, error) {
	ret := new(struct {
		Token common.Address
		User  common.Address
		Amt   *big.Int
	})
	out := ret
	err := _HbSwap.contract.Call(opts, out, "secretWithdrawMap", arg0)
	return *ret, err
}

// SecretWithdrawMap is a free data retrieval call binding the contract method 0x7aa6fd65.
//
// Solidity: function secretWithdrawMap(uint256 ) constant returns(address token, address user, uint256 amt)
func (_HbSwap *HbSwapSession) SecretWithdrawMap(arg0 *big.Int) (struct {
	Token common.Address
	User  common.Address
	Amt   *big.Int
}, error) {
	return _HbSwap.Contract.SecretWithdrawMap(&_HbSwap.CallOpts, arg0)
}

// SecretWithdrawMap is a free data retrieval call binding the contract method 0x7aa6fd65.
//
// Solidity: function secretWithdrawMap(uint256 ) constant returns(address token, address user, uint256 amt)
func (_HbSwap *HbSwapCallerSession) SecretWithdrawMap(arg0 *big.Int) (struct {
	Token common.Address
	User  common.Address
	Amt   *big.Int
}, error) {
	return _HbSwap.Contract.SecretWithdrawMap(&_HbSwap.CallOpts, arg0)
}

// Servers is a free data retrieval call binding the contract method 0x12ada8de.
//
// Solidity: function servers(address ) constant returns(bool)
func (_HbSwap *HbSwapCaller) Servers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _HbSwap.contract.Call(opts, out, "servers", arg0)
	return *ret0, err
}

// Servers is a free data retrieval call binding the contract method 0x12ada8de.
//
// Solidity: function servers(address ) constant returns(bool)
func (_HbSwap *HbSwapSession) Servers(arg0 common.Address) (bool, error) {
	return _HbSwap.Contract.Servers(&_HbSwap.CallOpts, arg0)
}

// Servers is a free data retrieval call binding the contract method 0x12ada8de.
//
// Solidity: function servers(address ) constant returns(bool)
func (_HbSwap *HbSwapCallerSession) Servers(arg0 common.Address) (bool, error) {
	return _HbSwap.Contract.Servers(&_HbSwap.CallOpts, arg0)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() constant returns(uint256)
func (_HbSwap *HbSwapCaller) Threshold(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HbSwap.contract.Call(opts, out, "threshold")
	return *ret0, err
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() constant returns(uint256)
func (_HbSwap *HbSwapSession) Threshold() (*big.Int, error) {
	return _HbSwap.Contract.Threshold(&_HbSwap.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() constant returns(uint256)
func (_HbSwap *HbSwapCallerSession) Threshold() (*big.Int, error) {
	return _HbSwap.Contract.Threshold(&_HbSwap.CallOpts)
}

// TradeCnt is a free data retrieval call binding the contract method 0xaf4170c4.
//
// Solidity: function tradeCnt() constant returns(uint256)
func (_HbSwap *HbSwapCaller) TradeCnt(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HbSwap.contract.Call(opts, out, "tradeCnt")
	return *ret0, err
}

// TradeCnt is a free data retrieval call binding the contract method 0xaf4170c4.
//
// Solidity: function tradeCnt() constant returns(uint256)
func (_HbSwap *HbSwapSession) TradeCnt() (*big.Int, error) {
	return _HbSwap.Contract.TradeCnt(&_HbSwap.CallOpts)
}

// TradeCnt is a free data retrieval call binding the contract method 0xaf4170c4.
//
// Solidity: function tradeCnt() constant returns(uint256)
func (_HbSwap *HbSwapCallerSession) TradeCnt() (*big.Int, error) {
	return _HbSwap.Contract.TradeCnt(&_HbSwap.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x3351733f.
//
// Solidity: function addLiquidity(address _tokenA, address _tokenB, uint256 _idxA, uint256 _idxB, uint256 _maskedAmtA, uint256 _maskedAmtB) returns()
func (_HbSwap *HbSwapTransactor) AddLiquidity(opts *bind.TransactOpts, _tokenA common.Address, _tokenB common.Address, _idxA *big.Int, _idxB *big.Int, _maskedAmtA *big.Int, _maskedAmtB *big.Int) (*types.Transaction, error) {
	return _HbSwap.contract.Transact(opts, "addLiquidity", _tokenA, _tokenB, _idxA, _idxB, _maskedAmtA, _maskedAmtB)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x3351733f.
//
// Solidity: function addLiquidity(address _tokenA, address _tokenB, uint256 _idxA, uint256 _idxB, uint256 _maskedAmtA, uint256 _maskedAmtB) returns()
func (_HbSwap *HbSwapSession) AddLiquidity(_tokenA common.Address, _tokenB common.Address, _idxA *big.Int, _idxB *big.Int, _maskedAmtA *big.Int, _maskedAmtB *big.Int) (*types.Transaction, error) {
	return _HbSwap.Contract.AddLiquidity(&_HbSwap.TransactOpts, _tokenA, _tokenB, _idxA, _idxB, _maskedAmtA, _maskedAmtB)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x3351733f.
//
// Solidity: function addLiquidity(address _tokenA, address _tokenB, uint256 _idxA, uint256 _idxB, uint256 _maskedAmtA, uint256 _maskedAmtB) returns()
func (_HbSwap *HbSwapTransactorSession) AddLiquidity(_tokenA common.Address, _tokenB common.Address, _idxA *big.Int, _idxB *big.Int, _maskedAmtA *big.Int, _maskedAmtB *big.Int) (*types.Transaction, error) {
	return _HbSwap.Contract.AddLiquidity(&_HbSwap.TransactOpts, _tokenA, _tokenB, _idxA, _idxB, _maskedAmtA, _maskedAmtB)
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

// PublicDeposit is a paid mutator transaction binding the contract method 0x93f54644.
//
// Solidity: function publicDeposit(address _token, uint256 _amt) returns()
func (_HbSwap *HbSwapTransactor) PublicDeposit(opts *bind.TransactOpts, _token common.Address, _amt *big.Int) (*types.Transaction, error) {
	return _HbSwap.contract.Transact(opts, "publicDeposit", _token, _amt)
}

// PublicDeposit is a paid mutator transaction binding the contract method 0x93f54644.
//
// Solidity: function publicDeposit(address _token, uint256 _amt) returns()
func (_HbSwap *HbSwapSession) PublicDeposit(_token common.Address, _amt *big.Int) (*types.Transaction, error) {
	return _HbSwap.Contract.PublicDeposit(&_HbSwap.TransactOpts, _token, _amt)
}

// PublicDeposit is a paid mutator transaction binding the contract method 0x93f54644.
//
// Solidity: function publicDeposit(address _token, uint256 _amt) returns()
func (_HbSwap *HbSwapTransactorSession) PublicDeposit(_token common.Address, _amt *big.Int) (*types.Transaction, error) {
	return _HbSwap.Contract.PublicDeposit(&_HbSwap.TransactOpts, _token, _amt)
}

// PublicWithdraw is a paid mutator transaction binding the contract method 0xbc94e7cf.
//
// Solidity: function publicWithdraw(address _token, uint256 _amt) returns()
func (_HbSwap *HbSwapTransactor) PublicWithdraw(opts *bind.TransactOpts, _token common.Address, _amt *big.Int) (*types.Transaction, error) {
	return _HbSwap.contract.Transact(opts, "publicWithdraw", _token, _amt)
}

// PublicWithdraw is a paid mutator transaction binding the contract method 0xbc94e7cf.
//
// Solidity: function publicWithdraw(address _token, uint256 _amt) returns()
func (_HbSwap *HbSwapSession) PublicWithdraw(_token common.Address, _amt *big.Int) (*types.Transaction, error) {
	return _HbSwap.Contract.PublicWithdraw(&_HbSwap.TransactOpts, _token, _amt)
}

// PublicWithdraw is a paid mutator transaction binding the contract method 0xbc94e7cf.
//
// Solidity: function publicWithdraw(address _token, uint256 _amt) returns()
func (_HbSwap *HbSwapTransactorSession) PublicWithdraw(_token common.Address, _amt *big.Int) (*types.Transaction, error) {
	return _HbSwap.Contract.PublicWithdraw(&_HbSwap.TransactOpts, _token, _amt)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x3ee7f9d3.
//
// Solidity: function removeLiquidity(address _tokenA, address _tokenB, uint256 _idx, uint256 _maskedAmt) returns()
func (_HbSwap *HbSwapTransactor) RemoveLiquidity(opts *bind.TransactOpts, _tokenA common.Address, _tokenB common.Address, _idx *big.Int, _maskedAmt *big.Int) (*types.Transaction, error) {
	return _HbSwap.contract.Transact(opts, "removeLiquidity", _tokenA, _tokenB, _idx, _maskedAmt)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x3ee7f9d3.
//
// Solidity: function removeLiquidity(address _tokenA, address _tokenB, uint256 _idx, uint256 _maskedAmt) returns()
func (_HbSwap *HbSwapSession) RemoveLiquidity(_tokenA common.Address, _tokenB common.Address, _idx *big.Int, _maskedAmt *big.Int) (*types.Transaction, error) {
	return _HbSwap.Contract.RemoveLiquidity(&_HbSwap.TransactOpts, _tokenA, _tokenB, _idx, _maskedAmt)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x3ee7f9d3.
//
// Solidity: function removeLiquidity(address _tokenA, address _tokenB, uint256 _idx, uint256 _maskedAmt) returns()
func (_HbSwap *HbSwapTransactorSession) RemoveLiquidity(_tokenA common.Address, _tokenB common.Address, _idx *big.Int, _maskedAmt *big.Int) (*types.Transaction, error) {
	return _HbSwap.Contract.RemoveLiquidity(&_HbSwap.TransactOpts, _tokenA, _tokenB, _idx, _maskedAmt)
}

// ReserveInput is a paid mutator transaction binding the contract method 0xcc5bd7fa.
//
// Solidity: function reserveInput(uint256 _num) returns()
func (_HbSwap *HbSwapTransactor) ReserveInput(opts *bind.TransactOpts, _num *big.Int) (*types.Transaction, error) {
	return _HbSwap.contract.Transact(opts, "reserveInput", _num)
}

// ReserveInput is a paid mutator transaction binding the contract method 0xcc5bd7fa.
//
// Solidity: function reserveInput(uint256 _num) returns()
func (_HbSwap *HbSwapSession) ReserveInput(_num *big.Int) (*types.Transaction, error) {
	return _HbSwap.Contract.ReserveInput(&_HbSwap.TransactOpts, _num)
}

// ReserveInput is a paid mutator transaction binding the contract method 0xcc5bd7fa.
//
// Solidity: function reserveInput(uint256 _num) returns()
func (_HbSwap *HbSwapTransactorSession) ReserveInput(_num *big.Int) (*types.Transaction, error) {
	return _HbSwap.Contract.ReserveInput(&_HbSwap.TransactOpts, _num)
}

// ResetBalance is a paid mutator transaction binding the contract method 0xca7b58ad.
//
// Solidity: function resetBalance(address _token, address _user) returns()
func (_HbSwap *HbSwapTransactor) ResetBalance(opts *bind.TransactOpts, _token common.Address, _user common.Address) (*types.Transaction, error) {
	return _HbSwap.contract.Transact(opts, "resetBalance", _token, _user)
}

// ResetBalance is a paid mutator transaction binding the contract method 0xca7b58ad.
//
// Solidity: function resetBalance(address _token, address _user) returns()
func (_HbSwap *HbSwapSession) ResetBalance(_token common.Address, _user common.Address) (*types.Transaction, error) {
	return _HbSwap.Contract.ResetBalance(&_HbSwap.TransactOpts, _token, _user)
}

// ResetBalance is a paid mutator transaction binding the contract method 0xca7b58ad.
//
// Solidity: function resetBalance(address _token, address _user) returns()
func (_HbSwap *HbSwapTransactorSession) ResetBalance(_token common.Address, _user common.Address) (*types.Transaction, error) {
	return _HbSwap.Contract.ResetBalance(&_HbSwap.TransactOpts, _token, _user)
}

// ResetPrice is a paid mutator transaction binding the contract method 0x9450fb26.
//
// Solidity: function resetPrice(address _tokenA, address _tokenB, address[] _servers) returns()
func (_HbSwap *HbSwapTransactor) ResetPrice(opts *bind.TransactOpts, _tokenA common.Address, _tokenB common.Address, _servers []common.Address) (*types.Transaction, error) {
	return _HbSwap.contract.Transact(opts, "resetPrice", _tokenA, _tokenB, _servers)
}

// ResetPrice is a paid mutator transaction binding the contract method 0x9450fb26.
//
// Solidity: function resetPrice(address _tokenA, address _tokenB, address[] _servers) returns()
func (_HbSwap *HbSwapSession) ResetPrice(_tokenA common.Address, _tokenB common.Address, _servers []common.Address) (*types.Transaction, error) {
	return _HbSwap.Contract.ResetPrice(&_HbSwap.TransactOpts, _tokenA, _tokenB, _servers)
}

// ResetPrice is a paid mutator transaction binding the contract method 0x9450fb26.
//
// Solidity: function resetPrice(address _tokenA, address _tokenB, address[] _servers) returns()
func (_HbSwap *HbSwapTransactorSession) ResetPrice(_tokenA common.Address, _tokenB common.Address, _servers []common.Address) (*types.Transaction, error) {
	return _HbSwap.Contract.ResetPrice(&_HbSwap.TransactOpts, _tokenA, _tokenB, _servers)
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
	User       common.Address
	TokenA     common.Address
	TokenB     common.Address
	IdxA       *big.Int
	IdxB       *big.Int
	MaskedAmtA *big.Int
	MaskedAmtB *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddLiquidity is a free log retrieval operation binding the contract event 0xec7d4752dd44bf7fc59045c9d80163de2a1b9dbd9032d11cb1156f7f867c6411.
//
// Solidity: event AddLiquidity(address user, address tokenA, address tokenB, uint256 idxA, uint256 idxB, uint256 maskedAmtA, uint256 maskedAmtB)
func (_HbSwap *HbSwapFilterer) FilterAddLiquidity(opts *bind.FilterOpts) (*HbSwapAddLiquidityIterator, error) {

	logs, sub, err := _HbSwap.contract.FilterLogs(opts, "AddLiquidity")
	if err != nil {
		return nil, err
	}
	return &HbSwapAddLiquidityIterator{contract: _HbSwap.contract, event: "AddLiquidity", logs: logs, sub: sub}, nil
}

// WatchAddLiquidity is a free log subscription operation binding the contract event 0xec7d4752dd44bf7fc59045c9d80163de2a1b9dbd9032d11cb1156f7f867c6411.
//
// Solidity: event AddLiquidity(address user, address tokenA, address tokenB, uint256 idxA, uint256 idxB, uint256 maskedAmtA, uint256 maskedAmtB)
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

// ParseAddLiquidity is a log parse operation binding the contract event 0xec7d4752dd44bf7fc59045c9d80163de2a1b9dbd9032d11cb1156f7f867c6411.
//
// Solidity: event AddLiquidity(address user, address tokenA, address tokenB, uint256 idxA, uint256 idxB, uint256 maskedAmtA, uint256 maskedAmtB)
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

// HbSwapInputMaskIterator is returned from FilterInputMask and is used to iterate over the raw logs and unpacked data for InputMask events raised by the HbSwap contract.
type HbSwapInputMaskIterator struct {
	Event *HbSwapInputMask // Event containing the contract specifics and raw log

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
func (it *HbSwapInputMaskIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HbSwapInputMask)
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
		it.Event = new(HbSwapInputMask)
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
func (it *HbSwapInputMaskIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HbSwapInputMaskIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HbSwapInputMask represents a InputMask event raised by the HbSwap contract.
type HbSwapInputMask struct {
	InpusMaskIndexes []*big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterInputMask is a free log retrieval operation binding the contract event 0xff742b54a62e49c065680fc16aa208f37f60ded27e03a3d19e6088ba0c7b41d4.
//
// Solidity: event InputMask(uint256[] inpusMaskIndexes)
func (_HbSwap *HbSwapFilterer) FilterInputMask(opts *bind.FilterOpts) (*HbSwapInputMaskIterator, error) {

	logs, sub, err := _HbSwap.contract.FilterLogs(opts, "InputMask")
	if err != nil {
		return nil, err
	}
	return &HbSwapInputMaskIterator{contract: _HbSwap.contract, event: "InputMask", logs: logs, sub: sub}, nil
}

// WatchInputMask is a free log subscription operation binding the contract event 0xff742b54a62e49c065680fc16aa208f37f60ded27e03a3d19e6088ba0c7b41d4.
//
// Solidity: event InputMask(uint256[] inpusMaskIndexes)
func (_HbSwap *HbSwapFilterer) WatchInputMask(opts *bind.WatchOpts, sink chan<- *HbSwapInputMask) (event.Subscription, error) {

	logs, sub, err := _HbSwap.contract.WatchLogs(opts, "InputMask")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HbSwapInputMask)
				if err := _HbSwap.contract.UnpackLog(event, "InputMask", log); err != nil {
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

// ParseInputMask is a log parse operation binding the contract event 0xff742b54a62e49c065680fc16aa208f37f60ded27e03a3d19e6088ba0c7b41d4.
//
// Solidity: event InputMask(uint256[] inpusMaskIndexes)
func (_HbSwap *HbSwapFilterer) ParseInputMask(log types.Log) (*HbSwapInputMask, error) {
	event := new(HbSwapInputMask)
	if err := _HbSwap.contract.UnpackLog(event, "InputMask", log); err != nil {
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
	User      common.Address
	TokenA    common.Address
	TokenB    common.Address
	Idx       *big.Int
	MaskedAmt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidity is a free log retrieval operation binding the contract event 0xa8dbaaebbb025c88e9e34c84635cd8238043556e9af43fb161508c898a8e1ef9.
//
// Solidity: event RemoveLiquidity(address user, address tokenA, address tokenB, uint256 idx, uint256 maskedAmt)
func (_HbSwap *HbSwapFilterer) FilterRemoveLiquidity(opts *bind.FilterOpts) (*HbSwapRemoveLiquidityIterator, error) {

	logs, sub, err := _HbSwap.contract.FilterLogs(opts, "RemoveLiquidity")
	if err != nil {
		return nil, err
	}
	return &HbSwapRemoveLiquidityIterator{contract: _HbSwap.contract, event: "RemoveLiquidity", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidity is a free log subscription operation binding the contract event 0xa8dbaaebbb025c88e9e34c84635cd8238043556e9af43fb161508c898a8e1ef9.
//
// Solidity: event RemoveLiquidity(address user, address tokenA, address tokenB, uint256 idx, uint256 maskedAmt)
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

// ParseRemoveLiquidity is a log parse operation binding the contract event 0xa8dbaaebbb025c88e9e34c84635cd8238043556e9af43fb161508c898a8e1ef9.
//
// Solidity: event RemoveLiquidity(address user, address tokenA, address tokenB, uint256 idx, uint256 maskedAmt)
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
	TradeSeq   *big.Int
	User       common.Address
	TokenA     common.Address
	TokenB     common.Address
	IdxA       *big.Int
	IdxB       *big.Int
	MaskedAmtA *big.Int
	MaskedAmtB *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTrade is a free log retrieval operation binding the contract event 0x2b4d91cd20cc8800407e3614b8466a6f0729ac3b1fa43d4e2b059ff5593cbae6.
//
// Solidity: event Trade(uint256 tradeSeq, address user, address tokenA, address tokenB, uint256 idxA, uint256 idxB, uint256 maskedAmtA, uint256 maskedAmtB)
func (_HbSwap *HbSwapFilterer) FilterTrade(opts *bind.FilterOpts) (*HbSwapTradeIterator, error) {

	logs, sub, err := _HbSwap.contract.FilterLogs(opts, "Trade")
	if err != nil {
		return nil, err
	}
	return &HbSwapTradeIterator{contract: _HbSwap.contract, event: "Trade", logs: logs, sub: sub}, nil
}

// WatchTrade is a free log subscription operation binding the contract event 0x2b4d91cd20cc8800407e3614b8466a6f0729ac3b1fa43d4e2b059ff5593cbae6.
//
// Solidity: event Trade(uint256 tradeSeq, address user, address tokenA, address tokenB, uint256 idxA, uint256 idxB, uint256 maskedAmtA, uint256 maskedAmtB)
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
// Solidity: event Trade(uint256 tradeSeq, address user, address tokenA, address tokenB, uint256 idxA, uint256 idxB, uint256 maskedAmtA, uint256 maskedAmtB)
func (_HbSwap *HbSwapFilterer) ParseTrade(log types.Log) (*HbSwapTrade, error) {
	event := new(HbSwapTrade)
	if err := _HbSwap.contract.UnpackLog(event, "Trade", log); err != nil {
		return nil, err
	}
	return event, nil
}
