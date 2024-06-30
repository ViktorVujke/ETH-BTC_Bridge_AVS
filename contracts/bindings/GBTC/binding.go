// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractGBTC

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

// ContractGBTCMetaData contains all meta data concerning the ContractGBTC contract.
var ContractGBTCMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_taskManagerAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getTaskManagerAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockGBTC\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"btcTxHash\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unlockGBTC\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x60a060405234801561001057600080fd5b50604051610ab3380380610ab383398101604081905261002f91610040565b6001600160a01b0316608052610070565b60006020828403121561005257600080fd5b81516001600160a01b038116811461006957600080fd5b9392505050565b608051610a136100a060003960008181610120015281816102030152818161032401526104220152610a136000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c806370a082311161007157806370a082311461014a5780639c2a807d146101735780639dc29fac14610186578063a9059cbb14610199578063d3fc9864146101ac578063dd62ed3e146101bf57600080fd5b8063095ea7b3146100ae57806318160ddd146100d95780631c86b91f146100eb57806323b872dd1461010057806365ce8df714610113575b600080fd5b6100c46100bc366004610846565b600192915050565b60405190151581526020015b60405180910390f35b6004545b6040519081526020016100d0565b6100fe6100f9366004610846565b6101f8565b005b6100c461010e366004610870565b610302565b6040516001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001681526020016100d0565b6100dd6101583660046108ac565b6001600160a01b031660009081526020819052604090205490565b6100fe610181366004610846565b610319565b6100fe610194366004610846565b610417565b6100c46101a7366004610846565b61045a565b6100fe6101ba3660046108ce565b610468565b6100dd6101cd366004610955565b6001600160a01b03918216600090815260026020908152604080832093909416825291909152205490565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461022d57600080fd5b6001600160a01b0382166000908152602081905260409020548111156102a45760405162461bcd60e51b815260206004820152602160248201527f43616e206e6f74206c6f636b206d6f7265207468616e20696e2062616c616e636044820152606560f81b60648201526084015b60405180910390fd5b6001600160a01b038216600090815260208190526040812080548392906102cc90849061099e565b90915550506001600160a01b038216600090815260016020526040812080548392906102f99084906109b5565b90915550505050565b600061030f84848461047a565b5060019392505050565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461034e57600080fd5b6001600160a01b0382166000908152600160205260409020548111156103c25760405162461bcd60e51b815260206004820152602360248201527f43616e206e6f7420756e6c6f636b206d6f7265207468616e20696e2062616c616044820152626e636560e81b606482015260840161029b565b6001600160a01b038216600090815260016020526040812080548392906103ea90849061099e565b90915550506001600160a01b038216600090815260208190526040812080548392906102f99084906109b5565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461044c57600080fd5b6104568282610620565b5050565b60003361030f81858561047a565b61047484848484610724565b50505050565b6001600160a01b0383166104de5760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b606482015260840161029b565b6001600160a01b0382166105405760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b606482015260840161029b565b6001600160a01b0383166000908152602081905260409020548111156105b75760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b606482015260840161029b565b6001600160a01b038381166000818152602081815260408083208054879003905593861680835291849020805486019055925184815290927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35b505050565b6001600160a01b0382166106805760405162461bcd60e51b815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f206164647265736044820152607360f81b606482015260840161029b565b6001600160a01b038216600090815260016020526040902054818110156106f45760405162461bcd60e51b815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e604482015261636560f01b606482015260840161029b565b6001600160a01b0383166000908152600160205260408120838303905560048054849003905561061b9084908483565b6001600160a01b03841661077a5760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015260640161029b565b836003838360405161078d9291906109cd565b908152602001604051809103902060006101000a8154816001600160a01b0302191690836001600160a01b0316021790555082600460008282546107d191906109b5565b90915550506001600160a01b038416600081815260208181526040808320805488019055518681527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a350505050565b80356001600160a01b038116811461084157600080fd5b919050565b6000806040838503121561085957600080fd5b6108628361082a565b946020939093013593505050565b60008060006060848603121561088557600080fd5b61088e8461082a565b925061089c6020850161082a565b9150604084013590509250925092565b6000602082840312156108be57600080fd5b6108c78261082a565b9392505050565b600080600080606085870312156108e457600080fd5b6108ed8561082a565b935060208501359250604085013567ffffffffffffffff8082111561091157600080fd5b818701915087601f83011261092557600080fd5b81358181111561093457600080fd5b88602082850101111561094657600080fd5b95989497505060200194505050565b6000806040838503121561096857600080fd5b6109718361082a565b915061097f6020840161082a565b90509250929050565b634e487b7160e01b600052601160045260246000fd5b6000828210156109b0576109b0610988565b500390565b600082198211156109c8576109c8610988565b500190565b818382376000910190815291905056fea264697066735822122053b0253a115f341e3d9063a157050973cc8213e23c71eaee8206e00e9db06ffb64736f6c634300080c0033",
}

// ContractGBTCABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractGBTCMetaData.ABI instead.
var ContractGBTCABI = ContractGBTCMetaData.ABI

// ContractGBTCBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractGBTCMetaData.Bin instead.
var ContractGBTCBin = ContractGBTCMetaData.Bin

// DeployContractGBTC deploys a new Ethereum contract, binding an instance of ContractGBTC to it.
func DeployContractGBTC(auth *bind.TransactOpts, backend bind.ContractBackend, _taskManagerAddress common.Address) (common.Address, *types.Transaction, *ContractGBTC, error) {
	parsed, err := ContractGBTCMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractGBTCBin), backend, _taskManagerAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractGBTC{ContractGBTCCaller: ContractGBTCCaller{contract: contract}, ContractGBTCTransactor: ContractGBTCTransactor{contract: contract}, ContractGBTCFilterer: ContractGBTCFilterer{contract: contract}}, nil
}

// ContractGBTC is an auto generated Go binding around an Ethereum contract.
type ContractGBTC struct {
	ContractGBTCCaller     // Read-only binding to the contract
	ContractGBTCTransactor // Write-only binding to the contract
	ContractGBTCFilterer   // Log filterer for contract events
}

// ContractGBTCCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractGBTCCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractGBTCTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractGBTCTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractGBTCFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractGBTCFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractGBTCSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractGBTCSession struct {
	Contract     *ContractGBTC     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractGBTCCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractGBTCCallerSession struct {
	Contract *ContractGBTCCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ContractGBTCTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractGBTCTransactorSession struct {
	Contract     *ContractGBTCTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ContractGBTCRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractGBTCRaw struct {
	Contract *ContractGBTC // Generic contract binding to access the raw methods on
}

// ContractGBTCCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractGBTCCallerRaw struct {
	Contract *ContractGBTCCaller // Generic read-only contract binding to access the raw methods on
}

// ContractGBTCTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractGBTCTransactorRaw struct {
	Contract *ContractGBTCTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractGBTC creates a new instance of ContractGBTC, bound to a specific deployed contract.
func NewContractGBTC(address common.Address, backend bind.ContractBackend) (*ContractGBTC, error) {
	contract, err := bindContractGBTC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractGBTC{ContractGBTCCaller: ContractGBTCCaller{contract: contract}, ContractGBTCTransactor: ContractGBTCTransactor{contract: contract}, ContractGBTCFilterer: ContractGBTCFilterer{contract: contract}}, nil
}

// NewContractGBTCCaller creates a new read-only instance of ContractGBTC, bound to a specific deployed contract.
func NewContractGBTCCaller(address common.Address, caller bind.ContractCaller) (*ContractGBTCCaller, error) {
	contract, err := bindContractGBTC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractGBTCCaller{contract: contract}, nil
}

// NewContractGBTCTransactor creates a new write-only instance of ContractGBTC, bound to a specific deployed contract.
func NewContractGBTCTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractGBTCTransactor, error) {
	contract, err := bindContractGBTC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractGBTCTransactor{contract: contract}, nil
}

// NewContractGBTCFilterer creates a new log filterer instance of ContractGBTC, bound to a specific deployed contract.
func NewContractGBTCFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractGBTCFilterer, error) {
	contract, err := bindContractGBTC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractGBTCFilterer{contract: contract}, nil
}

// bindContractGBTC binds a generic wrapper to an already deployed contract.
func bindContractGBTC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractGBTCMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractGBTC *ContractGBTCRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractGBTC.Contract.ContractGBTCCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractGBTC *ContractGBTCRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractGBTC.Contract.ContractGBTCTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractGBTC *ContractGBTCRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractGBTC.Contract.ContractGBTCTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractGBTC *ContractGBTCCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractGBTC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractGBTC *ContractGBTCTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractGBTC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractGBTC *ContractGBTCTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractGBTC.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ContractGBTC *ContractGBTCCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractGBTC.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ContractGBTC *ContractGBTCSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ContractGBTC.Contract.Allowance(&_ContractGBTC.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ContractGBTC *ContractGBTCCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ContractGBTC.Contract.Allowance(&_ContractGBTC.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ContractGBTC *ContractGBTCCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractGBTC.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ContractGBTC *ContractGBTCSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ContractGBTC.Contract.BalanceOf(&_ContractGBTC.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ContractGBTC *ContractGBTCCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ContractGBTC.Contract.BalanceOf(&_ContractGBTC.CallOpts, account)
}

// GetTaskManagerAddress is a free data retrieval call binding the contract method 0x65ce8df7.
//
// Solidity: function getTaskManagerAddress() view returns(address)
func (_ContractGBTC *ContractGBTCCaller) GetTaskManagerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractGBTC.contract.Call(opts, &out, "getTaskManagerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTaskManagerAddress is a free data retrieval call binding the contract method 0x65ce8df7.
//
// Solidity: function getTaskManagerAddress() view returns(address)
func (_ContractGBTC *ContractGBTCSession) GetTaskManagerAddress() (common.Address, error) {
	return _ContractGBTC.Contract.GetTaskManagerAddress(&_ContractGBTC.CallOpts)
}

// GetTaskManagerAddress is a free data retrieval call binding the contract method 0x65ce8df7.
//
// Solidity: function getTaskManagerAddress() view returns(address)
func (_ContractGBTC *ContractGBTCCallerSession) GetTaskManagerAddress() (common.Address, error) {
	return _ContractGBTC.Contract.GetTaskManagerAddress(&_ContractGBTC.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ContractGBTC *ContractGBTCCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractGBTC.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ContractGBTC *ContractGBTCSession) TotalSupply() (*big.Int, error) {
	return _ContractGBTC.Contract.TotalSupply(&_ContractGBTC.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ContractGBTC *ContractGBTCCallerSession) TotalSupply() (*big.Int, error) {
	return _ContractGBTC.Contract.TotalSupply(&_ContractGBTC.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) returns(bool)
func (_ContractGBTC *ContractGBTCTransactor) Approve(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _ContractGBTC.contract.Transact(opts, "approve", arg0, arg1)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) returns(bool)
func (_ContractGBTC *ContractGBTCSession) Approve(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _ContractGBTC.Contract.Approve(&_ContractGBTC.TransactOpts, arg0, arg1)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) returns(bool)
func (_ContractGBTC *ContractGBTCTransactorSession) Approve(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _ContractGBTC.Contract.Approve(&_ContractGBTC.TransactOpts, arg0, arg1)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_ContractGBTC *ContractGBTCTransactor) Burn(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractGBTC.contract.Transact(opts, "burn", account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_ContractGBTC *ContractGBTCSession) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractGBTC.Contract.Burn(&_ContractGBTC.TransactOpts, account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_ContractGBTC *ContractGBTCTransactorSession) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractGBTC.Contract.Burn(&_ContractGBTC.TransactOpts, account, amount)
}

// LockGBTC is a paid mutator transaction binding the contract method 0x1c86b91f.
//
// Solidity: function lockGBTC(address account, uint256 amount) returns()
func (_ContractGBTC *ContractGBTCTransactor) LockGBTC(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractGBTC.contract.Transact(opts, "lockGBTC", account, amount)
}

// LockGBTC is a paid mutator transaction binding the contract method 0x1c86b91f.
//
// Solidity: function lockGBTC(address account, uint256 amount) returns()
func (_ContractGBTC *ContractGBTCSession) LockGBTC(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractGBTC.Contract.LockGBTC(&_ContractGBTC.TransactOpts, account, amount)
}

// LockGBTC is a paid mutator transaction binding the contract method 0x1c86b91f.
//
// Solidity: function lockGBTC(address account, uint256 amount) returns()
func (_ContractGBTC *ContractGBTCTransactorSession) LockGBTC(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractGBTC.Contract.LockGBTC(&_ContractGBTC.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xd3fc9864.
//
// Solidity: function mint(address account, uint256 amount, string btcTxHash) returns()
func (_ContractGBTC *ContractGBTCTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int, btcTxHash string) (*types.Transaction, error) {
	return _ContractGBTC.contract.Transact(opts, "mint", account, amount, btcTxHash)
}

// Mint is a paid mutator transaction binding the contract method 0xd3fc9864.
//
// Solidity: function mint(address account, uint256 amount, string btcTxHash) returns()
func (_ContractGBTC *ContractGBTCSession) Mint(account common.Address, amount *big.Int, btcTxHash string) (*types.Transaction, error) {
	return _ContractGBTC.Contract.Mint(&_ContractGBTC.TransactOpts, account, amount, btcTxHash)
}

// Mint is a paid mutator transaction binding the contract method 0xd3fc9864.
//
// Solidity: function mint(address account, uint256 amount, string btcTxHash) returns()
func (_ContractGBTC *ContractGBTCTransactorSession) Mint(account common.Address, amount *big.Int, btcTxHash string) (*types.Transaction, error) {
	return _ContractGBTC.Contract.Mint(&_ContractGBTC.TransactOpts, account, amount, btcTxHash)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ContractGBTC *ContractGBTCTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractGBTC.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ContractGBTC *ContractGBTCSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractGBTC.Contract.Transfer(&_ContractGBTC.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ContractGBTC *ContractGBTCTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractGBTC.Contract.Transfer(&_ContractGBTC.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ContractGBTC *ContractGBTCTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractGBTC.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ContractGBTC *ContractGBTCSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractGBTC.Contract.TransferFrom(&_ContractGBTC.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ContractGBTC *ContractGBTCTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractGBTC.Contract.TransferFrom(&_ContractGBTC.TransactOpts, from, to, amount)
}

// UnlockGBTC is a paid mutator transaction binding the contract method 0x9c2a807d.
//
// Solidity: function unlockGBTC(address account, uint256 amount) returns()
func (_ContractGBTC *ContractGBTCTransactor) UnlockGBTC(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractGBTC.contract.Transact(opts, "unlockGBTC", account, amount)
}

// UnlockGBTC is a paid mutator transaction binding the contract method 0x9c2a807d.
//
// Solidity: function unlockGBTC(address account, uint256 amount) returns()
func (_ContractGBTC *ContractGBTCSession) UnlockGBTC(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractGBTC.Contract.UnlockGBTC(&_ContractGBTC.TransactOpts, account, amount)
}

// UnlockGBTC is a paid mutator transaction binding the contract method 0x9c2a807d.
//
// Solidity: function unlockGBTC(address account, uint256 amount) returns()
func (_ContractGBTC *ContractGBTCTransactorSession) UnlockGBTC(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractGBTC.Contract.UnlockGBTC(&_ContractGBTC.TransactOpts, account, amount)
}

// ContractGBTCApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ContractGBTC contract.
type ContractGBTCApprovalIterator struct {
	Event *ContractGBTCApproval // Event containing the contract specifics and raw log

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
func (it *ContractGBTCApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractGBTCApproval)
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
		it.Event = new(ContractGBTCApproval)
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
func (it *ContractGBTCApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractGBTCApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractGBTCApproval represents a Approval event raised by the ContractGBTC contract.
type ContractGBTCApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ContractGBTC *ContractGBTCFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ContractGBTCApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ContractGBTC.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ContractGBTCApprovalIterator{contract: _ContractGBTC.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ContractGBTC *ContractGBTCFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ContractGBTCApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ContractGBTC.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractGBTCApproval)
				if err := _ContractGBTC.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ContractGBTC *ContractGBTCFilterer) ParseApproval(log types.Log) (*ContractGBTCApproval, error) {
	event := new(ContractGBTCApproval)
	if err := _ContractGBTC.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractGBTCTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ContractGBTC contract.
type ContractGBTCTransferIterator struct {
	Event *ContractGBTCTransfer // Event containing the contract specifics and raw log

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
func (it *ContractGBTCTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractGBTCTransfer)
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
		it.Event = new(ContractGBTCTransfer)
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
func (it *ContractGBTCTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractGBTCTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractGBTCTransfer represents a Transfer event raised by the ContractGBTC contract.
type ContractGBTCTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ContractGBTC *ContractGBTCFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ContractGBTCTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ContractGBTC.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ContractGBTCTransferIterator{contract: _ContractGBTC.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ContractGBTC *ContractGBTCFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ContractGBTCTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ContractGBTC.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractGBTCTransfer)
				if err := _ContractGBTC.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ContractGBTC *ContractGBTCFilterer) ParseTransfer(log types.Log) (*ContractGBTCTransfer, error) {
	event := new(ContractGBTCTransfer)
	if err := _ContractGBTC.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
