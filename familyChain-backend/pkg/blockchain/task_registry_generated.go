// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package blockchain

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

// BlockchainMetaData contains all meta data concerning the Blockchain contract.
var BlockchainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approvedBy\",\"type\":\"address\"}],\"name\":\"TaskApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assignedTo\",\"type\":\"address\"}],\"name\":\"TaskAssigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"completedBy\",\"type\":\"address\"}],\"name\":\"TaskCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"TaskCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rejectedBy\",\"type\":\"address\"}],\"name\":\"TaskRejected\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"}],\"name\":\"approveTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"childAddress\",\"type\":\"address\"}],\"name\":\"assignTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"}],\"name\":\"completeTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"createTask\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"}],\"name\":\"getTask\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"}],\"name\":\"rejectTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"taskCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tasks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"assignedTo\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"completed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"rejected\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// BlockchainABI is the input ABI used to generate the binding from.
// Deprecated: Use BlockchainMetaData.ABI instead.
var BlockchainABI = BlockchainMetaData.ABI

// Blockchain is an auto generated Go binding around an Ethereum contract.
type Blockchain struct {
	BlockchainCaller     // Read-only binding to the contract
	BlockchainTransactor // Write-only binding to the contract
	BlockchainFilterer   // Log filterer for contract events
}

// BlockchainCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlockchainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockchainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlockchainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockchainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlockchainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockchainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlockchainSession struct {
	Contract     *Blockchain       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BlockchainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlockchainCallerSession struct {
	Contract *BlockchainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BlockchainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlockchainTransactorSession struct {
	Contract     *BlockchainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BlockchainRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlockchainRaw struct {
	Contract *Blockchain // Generic contract binding to access the raw methods on
}

// BlockchainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlockchainCallerRaw struct {
	Contract *BlockchainCaller // Generic read-only contract binding to access the raw methods on
}

// BlockchainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlockchainTransactorRaw struct {
	Contract *BlockchainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlockchain creates a new instance of Blockchain, bound to a specific deployed contract.
func NewBlockchain(address common.Address, backend bind.ContractBackend) (*Blockchain, error) {
	contract, err := bindBlockchain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Blockchain{BlockchainCaller: BlockchainCaller{contract: contract}, BlockchainTransactor: BlockchainTransactor{contract: contract}, BlockchainFilterer: BlockchainFilterer{contract: contract}}, nil
}

// NewBlockchainCaller creates a new read-only instance of Blockchain, bound to a specific deployed contract.
func NewBlockchainCaller(address common.Address, caller bind.ContractCaller) (*BlockchainCaller, error) {
	contract, err := bindBlockchain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlockchainCaller{contract: contract}, nil
}

// NewBlockchainTransactor creates a new write-only instance of Blockchain, bound to a specific deployed contract.
func NewBlockchainTransactor(address common.Address, transactor bind.ContractTransactor) (*BlockchainTransactor, error) {
	contract, err := bindBlockchain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlockchainTransactor{contract: contract}, nil
}

// NewBlockchainFilterer creates a new log filterer instance of Blockchain, bound to a specific deployed contract.
func NewBlockchainFilterer(address common.Address, filterer bind.ContractFilterer) (*BlockchainFilterer, error) {
	contract, err := bindBlockchain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlockchainFilterer{contract: contract}, nil
}

// bindBlockchain binds a generic wrapper to an already deployed contract.
func bindBlockchain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BlockchainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Blockchain *BlockchainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Blockchain.Contract.BlockchainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Blockchain *BlockchainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blockchain.Contract.BlockchainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Blockchain *BlockchainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Blockchain.Contract.BlockchainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Blockchain *BlockchainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Blockchain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Blockchain *BlockchainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blockchain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Blockchain *BlockchainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Blockchain.Contract.contract.Transact(opts, method, params...)
}

// GetTask is a free data retrieval call binding the contract method 0x1d65e77e.
//
// Solidity: function getTask(uint256 taskId) view returns(uint256, address, address, string, string, uint256, bool, bool, bool)
func (_Blockchain *BlockchainCaller) GetTask(opts *bind.CallOpts, taskId *big.Int) (*big.Int, common.Address, common.Address, string, string, *big.Int, bool, bool, bool, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "getTask", taskId)

	if err != nil {
		return *new(*big.Int), *new(common.Address), *new(common.Address), *new(string), *new(string), *new(*big.Int), *new(bool), *new(bool), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	out2 := *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	out3 := *abi.ConvertType(out[3], new(string)).(*string)
	out4 := *abi.ConvertType(out[4], new(string)).(*string)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	out6 := *abi.ConvertType(out[6], new(bool)).(*bool)
	out7 := *abi.ConvertType(out[7], new(bool)).(*bool)
	out8 := *abi.ConvertType(out[8], new(bool)).(*bool)

	return out0, out1, out2, out3, out4, out5, out6, out7, out8, err

}

// GetTask is a free data retrieval call binding the contract method 0x1d65e77e.
//
// Solidity: function getTask(uint256 taskId) view returns(uint256, address, address, string, string, uint256, bool, bool, bool)
func (_Blockchain *BlockchainSession) GetTask(taskId *big.Int) (*big.Int, common.Address, common.Address, string, string, *big.Int, bool, bool, bool, error) {
	return _Blockchain.Contract.GetTask(&_Blockchain.CallOpts, taskId)
}

// GetTask is a free data retrieval call binding the contract method 0x1d65e77e.
//
// Solidity: function getTask(uint256 taskId) view returns(uint256, address, address, string, string, uint256, bool, bool, bool)
func (_Blockchain *BlockchainCallerSession) GetTask(taskId *big.Int) (*big.Int, common.Address, common.Address, string, string, *big.Int, bool, bool, bool, error) {
	return _Blockchain.Contract.GetTask(&_Blockchain.CallOpts, taskId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Blockchain *BlockchainCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Blockchain *BlockchainSession) Owner() (common.Address, error) {
	return _Blockchain.Contract.Owner(&_Blockchain.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Blockchain *BlockchainCallerSession) Owner() (common.Address, error) {
	return _Blockchain.Contract.Owner(&_Blockchain.CallOpts)
}

// TaskCount is a free data retrieval call binding the contract method 0xb6cb58a5.
//
// Solidity: function taskCount() view returns(uint256)
func (_Blockchain *BlockchainCaller) TaskCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "taskCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TaskCount is a free data retrieval call binding the contract method 0xb6cb58a5.
//
// Solidity: function taskCount() view returns(uint256)
func (_Blockchain *BlockchainSession) TaskCount() (*big.Int, error) {
	return _Blockchain.Contract.TaskCount(&_Blockchain.CallOpts)
}

// TaskCount is a free data retrieval call binding the contract method 0xb6cb58a5.
//
// Solidity: function taskCount() view returns(uint256)
func (_Blockchain *BlockchainCallerSession) TaskCount() (*big.Int, error) {
	return _Blockchain.Contract.TaskCount(&_Blockchain.CallOpts)
}

// Tasks is a free data retrieval call binding the contract method 0x8d977672.
//
// Solidity: function tasks(uint256 ) view returns(uint256 id, address creator, address assignedTo, string title, string description, uint256 reward, bool completed, bool approved, bool rejected)
func (_Blockchain *BlockchainCaller) Tasks(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id          *big.Int
	Creator     common.Address
	AssignedTo  common.Address
	Title       string
	Description string
	Reward      *big.Int
	Completed   bool
	Approved    bool
	Rejected    bool
}, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "tasks", arg0)

	outstruct := new(struct {
		Id          *big.Int
		Creator     common.Address
		AssignedTo  common.Address
		Title       string
		Description string
		Reward      *big.Int
		Completed   bool
		Approved    bool
		Rejected    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Creator = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.AssignedTo = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Title = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Description = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.Reward = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Completed = *abi.ConvertType(out[6], new(bool)).(*bool)
	outstruct.Approved = *abi.ConvertType(out[7], new(bool)).(*bool)
	outstruct.Rejected = *abi.ConvertType(out[8], new(bool)).(*bool)

	return *outstruct, err

}

// Tasks is a free data retrieval call binding the contract method 0x8d977672.
//
// Solidity: function tasks(uint256 ) view returns(uint256 id, address creator, address assignedTo, string title, string description, uint256 reward, bool completed, bool approved, bool rejected)
func (_Blockchain *BlockchainSession) Tasks(arg0 *big.Int) (struct {
	Id          *big.Int
	Creator     common.Address
	AssignedTo  common.Address
	Title       string
	Description string
	Reward      *big.Int
	Completed   bool
	Approved    bool
	Rejected    bool
}, error) {
	return _Blockchain.Contract.Tasks(&_Blockchain.CallOpts, arg0)
}

// Tasks is a free data retrieval call binding the contract method 0x8d977672.
//
// Solidity: function tasks(uint256 ) view returns(uint256 id, address creator, address assignedTo, string title, string description, uint256 reward, bool completed, bool approved, bool rejected)
func (_Blockchain *BlockchainCallerSession) Tasks(arg0 *big.Int) (struct {
	Id          *big.Int
	Creator     common.Address
	AssignedTo  common.Address
	Title       string
	Description string
	Reward      *big.Int
	Completed   bool
	Approved    bool
	Rejected    bool
}, error) {
	return _Blockchain.Contract.Tasks(&_Blockchain.CallOpts, arg0)
}

// ApproveTask is a paid mutator transaction binding the contract method 0x0a07fae6.
//
// Solidity: function approveTask(uint256 taskId) returns()
func (_Blockchain *BlockchainTransactor) ApproveTask(opts *bind.TransactOpts, taskId *big.Int) (*types.Transaction, error) {
	return _Blockchain.contract.Transact(opts, "approveTask", taskId)
}

// ApproveTask is a paid mutator transaction binding the contract method 0x0a07fae6.
//
// Solidity: function approveTask(uint256 taskId) returns()
func (_Blockchain *BlockchainSession) ApproveTask(taskId *big.Int) (*types.Transaction, error) {
	return _Blockchain.Contract.ApproveTask(&_Blockchain.TransactOpts, taskId)
}

// ApproveTask is a paid mutator transaction binding the contract method 0x0a07fae6.
//
// Solidity: function approveTask(uint256 taskId) returns()
func (_Blockchain *BlockchainTransactorSession) ApproveTask(taskId *big.Int) (*types.Transaction, error) {
	return _Blockchain.Contract.ApproveTask(&_Blockchain.TransactOpts, taskId)
}

// AssignTask is a paid mutator transaction binding the contract method 0x5293ee81.
//
// Solidity: function assignTask(uint256 taskId, address childAddress) returns()
func (_Blockchain *BlockchainTransactor) AssignTask(opts *bind.TransactOpts, taskId *big.Int, childAddress common.Address) (*types.Transaction, error) {
	return _Blockchain.contract.Transact(opts, "assignTask", taskId, childAddress)
}

// AssignTask is a paid mutator transaction binding the contract method 0x5293ee81.
//
// Solidity: function assignTask(uint256 taskId, address childAddress) returns()
func (_Blockchain *BlockchainSession) AssignTask(taskId *big.Int, childAddress common.Address) (*types.Transaction, error) {
	return _Blockchain.Contract.AssignTask(&_Blockchain.TransactOpts, taskId, childAddress)
}

// AssignTask is a paid mutator transaction binding the contract method 0x5293ee81.
//
// Solidity: function assignTask(uint256 taskId, address childAddress) returns()
func (_Blockchain *BlockchainTransactorSession) AssignTask(taskId *big.Int, childAddress common.Address) (*types.Transaction, error) {
	return _Blockchain.Contract.AssignTask(&_Blockchain.TransactOpts, taskId, childAddress)
}

// CompleteTask is a paid mutator transaction binding the contract method 0xe1e29558.
//
// Solidity: function completeTask(uint256 taskId) returns()
func (_Blockchain *BlockchainTransactor) CompleteTask(opts *bind.TransactOpts, taskId *big.Int) (*types.Transaction, error) {
	return _Blockchain.contract.Transact(opts, "completeTask", taskId)
}

// CompleteTask is a paid mutator transaction binding the contract method 0xe1e29558.
//
// Solidity: function completeTask(uint256 taskId) returns()
func (_Blockchain *BlockchainSession) CompleteTask(taskId *big.Int) (*types.Transaction, error) {
	return _Blockchain.Contract.CompleteTask(&_Blockchain.TransactOpts, taskId)
}

// CompleteTask is a paid mutator transaction binding the contract method 0xe1e29558.
//
// Solidity: function completeTask(uint256 taskId) returns()
func (_Blockchain *BlockchainTransactorSession) CompleteTask(taskId *big.Int) (*types.Transaction, error) {
	return _Blockchain.Contract.CompleteTask(&_Blockchain.TransactOpts, taskId)
}

// CreateTask is a paid mutator transaction binding the contract method 0x41a4e30a.
//
// Solidity: function createTask(string title, string description, uint256 reward) payable returns(uint256)
func (_Blockchain *BlockchainTransactor) CreateTask(opts *bind.TransactOpts, title string, description string, reward *big.Int) (*types.Transaction, error) {
	return _Blockchain.contract.Transact(opts, "createTask", title, description, reward)
}

// CreateTask is a paid mutator transaction binding the contract method 0x41a4e30a.
//
// Solidity: function createTask(string title, string description, uint256 reward) payable returns(uint256)
func (_Blockchain *BlockchainSession) CreateTask(title string, description string, reward *big.Int) (*types.Transaction, error) {
	return _Blockchain.Contract.CreateTask(&_Blockchain.TransactOpts, title, description, reward)
}

// CreateTask is a paid mutator transaction binding the contract method 0x41a4e30a.
//
// Solidity: function createTask(string title, string description, uint256 reward) payable returns(uint256)
func (_Blockchain *BlockchainTransactorSession) CreateTask(title string, description string, reward *big.Int) (*types.Transaction, error) {
	return _Blockchain.Contract.CreateTask(&_Blockchain.TransactOpts, title, description, reward)
}

// RejectTask is a paid mutator transaction binding the contract method 0x7d81b40b.
//
// Solidity: function rejectTask(uint256 taskId) returns()
func (_Blockchain *BlockchainTransactor) RejectTask(opts *bind.TransactOpts, taskId *big.Int) (*types.Transaction, error) {
	return _Blockchain.contract.Transact(opts, "rejectTask", taskId)
}

// RejectTask is a paid mutator transaction binding the contract method 0x7d81b40b.
//
// Solidity: function rejectTask(uint256 taskId) returns()
func (_Blockchain *BlockchainSession) RejectTask(taskId *big.Int) (*types.Transaction, error) {
	return _Blockchain.Contract.RejectTask(&_Blockchain.TransactOpts, taskId)
}

// RejectTask is a paid mutator transaction binding the contract method 0x7d81b40b.
//
// Solidity: function rejectTask(uint256 taskId) returns()
func (_Blockchain *BlockchainTransactorSession) RejectTask(taskId *big.Int) (*types.Transaction, error) {
	return _Blockchain.Contract.RejectTask(&_Blockchain.TransactOpts, taskId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Blockchain *BlockchainTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blockchain.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Blockchain *BlockchainSession) Withdraw() (*types.Transaction, error) {
	return _Blockchain.Contract.Withdraw(&_Blockchain.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Blockchain *BlockchainTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Blockchain.Contract.Withdraw(&_Blockchain.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Blockchain *BlockchainTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blockchain.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Blockchain *BlockchainSession) Receive() (*types.Transaction, error) {
	return _Blockchain.Contract.Receive(&_Blockchain.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Blockchain *BlockchainTransactorSession) Receive() (*types.Transaction, error) {
	return _Blockchain.Contract.Receive(&_Blockchain.TransactOpts)
}

// BlockchainRewardTransferredIterator is returned from FilterRewardTransferred and is used to iterate over the raw logs and unpacked data for RewardTransferred events raised by the Blockchain contract.
type BlockchainRewardTransferredIterator struct {
	Event *BlockchainRewardTransferred // Event containing the contract specifics and raw log

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
func (it *BlockchainRewardTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockchainRewardTransferred)
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
		it.Event = new(BlockchainRewardTransferred)
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
func (it *BlockchainRewardTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockchainRewardTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockchainRewardTransferred represents a RewardTransferred event raised by the Blockchain contract.
type BlockchainRewardTransferred struct {
	TaskId    *big.Int
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRewardTransferred is a free log retrieval operation binding the contract event 0x99f2f3325e35634c09536d4983b1b2b0c9313e74f2b74508c7a1a521f6fce9ab.
//
// Solidity: event RewardTransferred(uint256 indexed taskId, address indexed recipient, uint256 amount)
func (_Blockchain *BlockchainFilterer) FilterRewardTransferred(opts *bind.FilterOpts, taskId []*big.Int, recipient []common.Address) (*BlockchainRewardTransferredIterator, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Blockchain.contract.FilterLogs(opts, "RewardTransferred", taskIdRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &BlockchainRewardTransferredIterator{contract: _Blockchain.contract, event: "RewardTransferred", logs: logs, sub: sub}, nil
}

// WatchRewardTransferred is a free log subscription operation binding the contract event 0x99f2f3325e35634c09536d4983b1b2b0c9313e74f2b74508c7a1a521f6fce9ab.
//
// Solidity: event RewardTransferred(uint256 indexed taskId, address indexed recipient, uint256 amount)
func (_Blockchain *BlockchainFilterer) WatchRewardTransferred(opts *bind.WatchOpts, sink chan<- *BlockchainRewardTransferred, taskId []*big.Int, recipient []common.Address) (event.Subscription, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Blockchain.contract.WatchLogs(opts, "RewardTransferred", taskIdRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockchainRewardTransferred)
				if err := _Blockchain.contract.UnpackLog(event, "RewardTransferred", log); err != nil {
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

// ParseRewardTransferred is a log parse operation binding the contract event 0x99f2f3325e35634c09536d4983b1b2b0c9313e74f2b74508c7a1a521f6fce9ab.
//
// Solidity: event RewardTransferred(uint256 indexed taskId, address indexed recipient, uint256 amount)
func (_Blockchain *BlockchainFilterer) ParseRewardTransferred(log types.Log) (*BlockchainRewardTransferred, error) {
	event := new(BlockchainRewardTransferred)
	if err := _Blockchain.contract.UnpackLog(event, "RewardTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockchainTaskApprovedIterator is returned from FilterTaskApproved and is used to iterate over the raw logs and unpacked data for TaskApproved events raised by the Blockchain contract.
type BlockchainTaskApprovedIterator struct {
	Event *BlockchainTaskApproved // Event containing the contract specifics and raw log

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
func (it *BlockchainTaskApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockchainTaskApproved)
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
		it.Event = new(BlockchainTaskApproved)
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
func (it *BlockchainTaskApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockchainTaskApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockchainTaskApproved represents a TaskApproved event raised by the Blockchain contract.
type BlockchainTaskApproved struct {
	TaskId     *big.Int
	ApprovedBy common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTaskApproved is a free log retrieval operation binding the contract event 0xb664fa2a6a4e034515db6ca96da0095c8d853e3349904bd555c50d1ac9a63595.
//
// Solidity: event TaskApproved(uint256 indexed taskId, address indexed approvedBy)
func (_Blockchain *BlockchainFilterer) FilterTaskApproved(opts *bind.FilterOpts, taskId []*big.Int, approvedBy []common.Address) (*BlockchainTaskApprovedIterator, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var approvedByRule []interface{}
	for _, approvedByItem := range approvedBy {
		approvedByRule = append(approvedByRule, approvedByItem)
	}

	logs, sub, err := _Blockchain.contract.FilterLogs(opts, "TaskApproved", taskIdRule, approvedByRule)
	if err != nil {
		return nil, err
	}
	return &BlockchainTaskApprovedIterator{contract: _Blockchain.contract, event: "TaskApproved", logs: logs, sub: sub}, nil
}

// WatchTaskApproved is a free log subscription operation binding the contract event 0xb664fa2a6a4e034515db6ca96da0095c8d853e3349904bd555c50d1ac9a63595.
//
// Solidity: event TaskApproved(uint256 indexed taskId, address indexed approvedBy)
func (_Blockchain *BlockchainFilterer) WatchTaskApproved(opts *bind.WatchOpts, sink chan<- *BlockchainTaskApproved, taskId []*big.Int, approvedBy []common.Address) (event.Subscription, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var approvedByRule []interface{}
	for _, approvedByItem := range approvedBy {
		approvedByRule = append(approvedByRule, approvedByItem)
	}

	logs, sub, err := _Blockchain.contract.WatchLogs(opts, "TaskApproved", taskIdRule, approvedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockchainTaskApproved)
				if err := _Blockchain.contract.UnpackLog(event, "TaskApproved", log); err != nil {
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

// ParseTaskApproved is a log parse operation binding the contract event 0xb664fa2a6a4e034515db6ca96da0095c8d853e3349904bd555c50d1ac9a63595.
//
// Solidity: event TaskApproved(uint256 indexed taskId, address indexed approvedBy)
func (_Blockchain *BlockchainFilterer) ParseTaskApproved(log types.Log) (*BlockchainTaskApproved, error) {
	event := new(BlockchainTaskApproved)
	if err := _Blockchain.contract.UnpackLog(event, "TaskApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockchainTaskAssignedIterator is returned from FilterTaskAssigned and is used to iterate over the raw logs and unpacked data for TaskAssigned events raised by the Blockchain contract.
type BlockchainTaskAssignedIterator struct {
	Event *BlockchainTaskAssigned // Event containing the contract specifics and raw log

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
func (it *BlockchainTaskAssignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockchainTaskAssigned)
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
		it.Event = new(BlockchainTaskAssigned)
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
func (it *BlockchainTaskAssignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockchainTaskAssignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockchainTaskAssigned represents a TaskAssigned event raised by the Blockchain contract.
type BlockchainTaskAssigned struct {
	TaskId     *big.Int
	AssignedTo common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTaskAssigned is a free log retrieval operation binding the contract event 0x52476d55ecef5cf13caa64038f297fe6bbf865d9584a98b8722a15a6d5db128f.
//
// Solidity: event TaskAssigned(uint256 indexed taskId, address indexed assignedTo)
func (_Blockchain *BlockchainFilterer) FilterTaskAssigned(opts *bind.FilterOpts, taskId []*big.Int, assignedTo []common.Address) (*BlockchainTaskAssignedIterator, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var assignedToRule []interface{}
	for _, assignedToItem := range assignedTo {
		assignedToRule = append(assignedToRule, assignedToItem)
	}

	logs, sub, err := _Blockchain.contract.FilterLogs(opts, "TaskAssigned", taskIdRule, assignedToRule)
	if err != nil {
		return nil, err
	}
	return &BlockchainTaskAssignedIterator{contract: _Blockchain.contract, event: "TaskAssigned", logs: logs, sub: sub}, nil
}

// WatchTaskAssigned is a free log subscription operation binding the contract event 0x52476d55ecef5cf13caa64038f297fe6bbf865d9584a98b8722a15a6d5db128f.
//
// Solidity: event TaskAssigned(uint256 indexed taskId, address indexed assignedTo)
func (_Blockchain *BlockchainFilterer) WatchTaskAssigned(opts *bind.WatchOpts, sink chan<- *BlockchainTaskAssigned, taskId []*big.Int, assignedTo []common.Address) (event.Subscription, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var assignedToRule []interface{}
	for _, assignedToItem := range assignedTo {
		assignedToRule = append(assignedToRule, assignedToItem)
	}

	logs, sub, err := _Blockchain.contract.WatchLogs(opts, "TaskAssigned", taskIdRule, assignedToRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockchainTaskAssigned)
				if err := _Blockchain.contract.UnpackLog(event, "TaskAssigned", log); err != nil {
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

// ParseTaskAssigned is a log parse operation binding the contract event 0x52476d55ecef5cf13caa64038f297fe6bbf865d9584a98b8722a15a6d5db128f.
//
// Solidity: event TaskAssigned(uint256 indexed taskId, address indexed assignedTo)
func (_Blockchain *BlockchainFilterer) ParseTaskAssigned(log types.Log) (*BlockchainTaskAssigned, error) {
	event := new(BlockchainTaskAssigned)
	if err := _Blockchain.contract.UnpackLog(event, "TaskAssigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockchainTaskCompletedIterator is returned from FilterTaskCompleted and is used to iterate over the raw logs and unpacked data for TaskCompleted events raised by the Blockchain contract.
type BlockchainTaskCompletedIterator struct {
	Event *BlockchainTaskCompleted // Event containing the contract specifics and raw log

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
func (it *BlockchainTaskCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockchainTaskCompleted)
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
		it.Event = new(BlockchainTaskCompleted)
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
func (it *BlockchainTaskCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockchainTaskCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockchainTaskCompleted represents a TaskCompleted event raised by the Blockchain contract.
type BlockchainTaskCompleted struct {
	TaskId      *big.Int
	CompletedBy common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTaskCompleted is a free log retrieval operation binding the contract event 0xbb5889c77948badf90e8a5c73d55265e5f5d6e4837a79a78c5669691b897faed.
//
// Solidity: event TaskCompleted(uint256 indexed taskId, address indexed completedBy)
func (_Blockchain *BlockchainFilterer) FilterTaskCompleted(opts *bind.FilterOpts, taskId []*big.Int, completedBy []common.Address) (*BlockchainTaskCompletedIterator, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var completedByRule []interface{}
	for _, completedByItem := range completedBy {
		completedByRule = append(completedByRule, completedByItem)
	}

	logs, sub, err := _Blockchain.contract.FilterLogs(opts, "TaskCompleted", taskIdRule, completedByRule)
	if err != nil {
		return nil, err
	}
	return &BlockchainTaskCompletedIterator{contract: _Blockchain.contract, event: "TaskCompleted", logs: logs, sub: sub}, nil
}

// WatchTaskCompleted is a free log subscription operation binding the contract event 0xbb5889c77948badf90e8a5c73d55265e5f5d6e4837a79a78c5669691b897faed.
//
// Solidity: event TaskCompleted(uint256 indexed taskId, address indexed completedBy)
func (_Blockchain *BlockchainFilterer) WatchTaskCompleted(opts *bind.WatchOpts, sink chan<- *BlockchainTaskCompleted, taskId []*big.Int, completedBy []common.Address) (event.Subscription, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var completedByRule []interface{}
	for _, completedByItem := range completedBy {
		completedByRule = append(completedByRule, completedByItem)
	}

	logs, sub, err := _Blockchain.contract.WatchLogs(opts, "TaskCompleted", taskIdRule, completedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockchainTaskCompleted)
				if err := _Blockchain.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
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

// ParseTaskCompleted is a log parse operation binding the contract event 0xbb5889c77948badf90e8a5c73d55265e5f5d6e4837a79a78c5669691b897faed.
//
// Solidity: event TaskCompleted(uint256 indexed taskId, address indexed completedBy)
func (_Blockchain *BlockchainFilterer) ParseTaskCompleted(log types.Log) (*BlockchainTaskCompleted, error) {
	event := new(BlockchainTaskCompleted)
	if err := _Blockchain.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockchainTaskCreatedIterator is returned from FilterTaskCreated and is used to iterate over the raw logs and unpacked data for TaskCreated events raised by the Blockchain contract.
type BlockchainTaskCreatedIterator struct {
	Event *BlockchainTaskCreated // Event containing the contract specifics and raw log

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
func (it *BlockchainTaskCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockchainTaskCreated)
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
		it.Event = new(BlockchainTaskCreated)
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
func (it *BlockchainTaskCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockchainTaskCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockchainTaskCreated represents a TaskCreated event raised by the Blockchain contract.
type BlockchainTaskCreated struct {
	TaskId  *big.Int
	Creator common.Address
	Title   string
	Reward  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTaskCreated is a free log retrieval operation binding the contract event 0x9e3c9757475c00b86393e183b68033bb76e48fa164849c7428ad17f62e1954a7.
//
// Solidity: event TaskCreated(uint256 indexed taskId, address indexed creator, string title, uint256 reward)
func (_Blockchain *BlockchainFilterer) FilterTaskCreated(opts *bind.FilterOpts, taskId []*big.Int, creator []common.Address) (*BlockchainTaskCreatedIterator, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _Blockchain.contract.FilterLogs(opts, "TaskCreated", taskIdRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &BlockchainTaskCreatedIterator{contract: _Blockchain.contract, event: "TaskCreated", logs: logs, sub: sub}, nil
}

// WatchTaskCreated is a free log subscription operation binding the contract event 0x9e3c9757475c00b86393e183b68033bb76e48fa164849c7428ad17f62e1954a7.
//
// Solidity: event TaskCreated(uint256 indexed taskId, address indexed creator, string title, uint256 reward)
func (_Blockchain *BlockchainFilterer) WatchTaskCreated(opts *bind.WatchOpts, sink chan<- *BlockchainTaskCreated, taskId []*big.Int, creator []common.Address) (event.Subscription, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _Blockchain.contract.WatchLogs(opts, "TaskCreated", taskIdRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockchainTaskCreated)
				if err := _Blockchain.contract.UnpackLog(event, "TaskCreated", log); err != nil {
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

// ParseTaskCreated is a log parse operation binding the contract event 0x9e3c9757475c00b86393e183b68033bb76e48fa164849c7428ad17f62e1954a7.
//
// Solidity: event TaskCreated(uint256 indexed taskId, address indexed creator, string title, uint256 reward)
func (_Blockchain *BlockchainFilterer) ParseTaskCreated(log types.Log) (*BlockchainTaskCreated, error) {
	event := new(BlockchainTaskCreated)
	if err := _Blockchain.contract.UnpackLog(event, "TaskCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockchainTaskRejectedIterator is returned from FilterTaskRejected and is used to iterate over the raw logs and unpacked data for TaskRejected events raised by the Blockchain contract.
type BlockchainTaskRejectedIterator struct {
	Event *BlockchainTaskRejected // Event containing the contract specifics and raw log

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
func (it *BlockchainTaskRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockchainTaskRejected)
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
		it.Event = new(BlockchainTaskRejected)
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
func (it *BlockchainTaskRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockchainTaskRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockchainTaskRejected represents a TaskRejected event raised by the Blockchain contract.
type BlockchainTaskRejected struct {
	TaskId     *big.Int
	RejectedBy common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTaskRejected is a free log retrieval operation binding the contract event 0xae1c93c6393c5abab2eaef81c691beeb381335c7913caabf8feaa0b24bd03cc0.
//
// Solidity: event TaskRejected(uint256 indexed taskId, address indexed rejectedBy)
func (_Blockchain *BlockchainFilterer) FilterTaskRejected(opts *bind.FilterOpts, taskId []*big.Int, rejectedBy []common.Address) (*BlockchainTaskRejectedIterator, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var rejectedByRule []interface{}
	for _, rejectedByItem := range rejectedBy {
		rejectedByRule = append(rejectedByRule, rejectedByItem)
	}

	logs, sub, err := _Blockchain.contract.FilterLogs(opts, "TaskRejected", taskIdRule, rejectedByRule)
	if err != nil {
		return nil, err
	}
	return &BlockchainTaskRejectedIterator{contract: _Blockchain.contract, event: "TaskRejected", logs: logs, sub: sub}, nil
}

// WatchTaskRejected is a free log subscription operation binding the contract event 0xae1c93c6393c5abab2eaef81c691beeb381335c7913caabf8feaa0b24bd03cc0.
//
// Solidity: event TaskRejected(uint256 indexed taskId, address indexed rejectedBy)
func (_Blockchain *BlockchainFilterer) WatchTaskRejected(opts *bind.WatchOpts, sink chan<- *BlockchainTaskRejected, taskId []*big.Int, rejectedBy []common.Address) (event.Subscription, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var rejectedByRule []interface{}
	for _, rejectedByItem := range rejectedBy {
		rejectedByRule = append(rejectedByRule, rejectedByItem)
	}

	logs, sub, err := _Blockchain.contract.WatchLogs(opts, "TaskRejected", taskIdRule, rejectedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockchainTaskRejected)
				if err := _Blockchain.contract.UnpackLog(event, "TaskRejected", log); err != nil {
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

// ParseTaskRejected is a log parse operation binding the contract event 0xae1c93c6393c5abab2eaef81c691beeb381335c7913caabf8feaa0b24bd03cc0.
//
// Solidity: event TaskRejected(uint256 indexed taskId, address indexed rejectedBy)
func (_Blockchain *BlockchainFilterer) ParseTaskRejected(log types.Log) (*BlockchainTaskRejected, error) {
	event := new(BlockchainTaskRejected)
	if err := _Blockchain.contract.UnpackLog(event, "TaskRejected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
