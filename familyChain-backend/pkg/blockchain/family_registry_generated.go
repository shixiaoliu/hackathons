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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"familyId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"childAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"age\",\"type\":\"uint8\"}],\"name\":\"ChildAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"familyId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"childAddress\",\"type\":\"address\"}],\"name\":\"ChildRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"familyId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"parent\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"FamilyCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"familyId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"FamilyUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"familyId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"childAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"age\",\"type\":\"uint8\"}],\"name\":\"addChild\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"childToFamily\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"createFamily\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"families\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"parent\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"familyCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"familyId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"childAddress\",\"type\":\"address\"}],\"name\":\"getChild\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"familyId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getChildAddressByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"familyId\",\"type\":\"uint256\"}],\"name\":\"getChildCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isChild\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isParent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"parentFamilies\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"familyId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"childAddress\",\"type\":\"address\"}],\"name\":\"removeChild\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"familyId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"updateFamily\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// ChildToFamily is a free data retrieval call binding the contract method 0xffa65a20.
//
// Solidity: function childToFamily(address ) view returns(uint256)
func (_Blockchain *BlockchainCaller) ChildToFamily(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "childToFamily", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChildToFamily is a free data retrieval call binding the contract method 0xffa65a20.
//
// Solidity: function childToFamily(address ) view returns(uint256)
func (_Blockchain *BlockchainSession) ChildToFamily(arg0 common.Address) (*big.Int, error) {
	return _Blockchain.Contract.ChildToFamily(&_Blockchain.CallOpts, arg0)
}

// ChildToFamily is a free data retrieval call binding the contract method 0xffa65a20.
//
// Solidity: function childToFamily(address ) view returns(uint256)
func (_Blockchain *BlockchainCallerSession) ChildToFamily(arg0 common.Address) (*big.Int, error) {
	return _Blockchain.Contract.ChildToFamily(&_Blockchain.CallOpts, arg0)
}

// Families is a free data retrieval call binding the contract method 0x078beae5.
//
// Solidity: function families(uint256 ) view returns(uint256 id, address parent, string name, bool active)
func (_Blockchain *BlockchainCaller) Families(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id     *big.Int
	Parent common.Address
	Name   string
	Active bool
}, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "families", arg0)

	outstruct := new(struct {
		Id     *big.Int
		Parent common.Address
		Name   string
		Active bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Parent = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Name = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Active = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// Families is a free data retrieval call binding the contract method 0x078beae5.
//
// Solidity: function families(uint256 ) view returns(uint256 id, address parent, string name, bool active)
func (_Blockchain *BlockchainSession) Families(arg0 *big.Int) (struct {
	Id     *big.Int
	Parent common.Address
	Name   string
	Active bool
}, error) {
	return _Blockchain.Contract.Families(&_Blockchain.CallOpts, arg0)
}

// Families is a free data retrieval call binding the contract method 0x078beae5.
//
// Solidity: function families(uint256 ) view returns(uint256 id, address parent, string name, bool active)
func (_Blockchain *BlockchainCallerSession) Families(arg0 *big.Int) (struct {
	Id     *big.Int
	Parent common.Address
	Name   string
	Active bool
}, error) {
	return _Blockchain.Contract.Families(&_Blockchain.CallOpts, arg0)
}

// FamilyCount is a free data retrieval call binding the contract method 0x784f7a93.
//
// Solidity: function familyCount() view returns(uint256)
func (_Blockchain *BlockchainCaller) FamilyCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "familyCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FamilyCount is a free data retrieval call binding the contract method 0x784f7a93.
//
// Solidity: function familyCount() view returns(uint256)
func (_Blockchain *BlockchainSession) FamilyCount() (*big.Int, error) {
	return _Blockchain.Contract.FamilyCount(&_Blockchain.CallOpts)
}

// FamilyCount is a free data retrieval call binding the contract method 0x784f7a93.
//
// Solidity: function familyCount() view returns(uint256)
func (_Blockchain *BlockchainCallerSession) FamilyCount() (*big.Int, error) {
	return _Blockchain.Contract.FamilyCount(&_Blockchain.CallOpts)
}

// GetChild is a free data retrieval call binding the contract method 0x9fa1140c.
//
// Solidity: function getChild(uint256 familyId, address childAddress) view returns(address, string, uint8, bool)
func (_Blockchain *BlockchainCaller) GetChild(opts *bind.CallOpts, familyId *big.Int, childAddress common.Address) (common.Address, string, uint8, bool, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "getChild", familyId, childAddress)

	if err != nil {
		return *new(common.Address), *new(string), *new(uint8), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(uint8)).(*uint8)
	out3 := *abi.ConvertType(out[3], new(bool)).(*bool)

	return out0, out1, out2, out3, err

}

// GetChild is a free data retrieval call binding the contract method 0x9fa1140c.
//
// Solidity: function getChild(uint256 familyId, address childAddress) view returns(address, string, uint8, bool)
func (_Blockchain *BlockchainSession) GetChild(familyId *big.Int, childAddress common.Address) (common.Address, string, uint8, bool, error) {
	return _Blockchain.Contract.GetChild(&_Blockchain.CallOpts, familyId, childAddress)
}

// GetChild is a free data retrieval call binding the contract method 0x9fa1140c.
//
// Solidity: function getChild(uint256 familyId, address childAddress) view returns(address, string, uint8, bool)
func (_Blockchain *BlockchainCallerSession) GetChild(familyId *big.Int, childAddress common.Address) (common.Address, string, uint8, bool, error) {
	return _Blockchain.Contract.GetChild(&_Blockchain.CallOpts, familyId, childAddress)
}

// GetChildAddressByIndex is a free data retrieval call binding the contract method 0x82ef100f.
//
// Solidity: function getChildAddressByIndex(uint256 familyId, uint256 index) view returns(address)
func (_Blockchain *BlockchainCaller) GetChildAddressByIndex(opts *bind.CallOpts, familyId *big.Int, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "getChildAddressByIndex", familyId, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetChildAddressByIndex is a free data retrieval call binding the contract method 0x82ef100f.
//
// Solidity: function getChildAddressByIndex(uint256 familyId, uint256 index) view returns(address)
func (_Blockchain *BlockchainSession) GetChildAddressByIndex(familyId *big.Int, index *big.Int) (common.Address, error) {
	return _Blockchain.Contract.GetChildAddressByIndex(&_Blockchain.CallOpts, familyId, index)
}

// GetChildAddressByIndex is a free data retrieval call binding the contract method 0x82ef100f.
//
// Solidity: function getChildAddressByIndex(uint256 familyId, uint256 index) view returns(address)
func (_Blockchain *BlockchainCallerSession) GetChildAddressByIndex(familyId *big.Int, index *big.Int) (common.Address, error) {
	return _Blockchain.Contract.GetChildAddressByIndex(&_Blockchain.CallOpts, familyId, index)
}

// GetChildCount is a free data retrieval call binding the contract method 0x7a77af5c.
//
// Solidity: function getChildCount(uint256 familyId) view returns(uint256)
func (_Blockchain *BlockchainCaller) GetChildCount(opts *bind.CallOpts, familyId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "getChildCount", familyId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChildCount is a free data retrieval call binding the contract method 0x7a77af5c.
//
// Solidity: function getChildCount(uint256 familyId) view returns(uint256)
func (_Blockchain *BlockchainSession) GetChildCount(familyId *big.Int) (*big.Int, error) {
	return _Blockchain.Contract.GetChildCount(&_Blockchain.CallOpts, familyId)
}

// GetChildCount is a free data retrieval call binding the contract method 0x7a77af5c.
//
// Solidity: function getChildCount(uint256 familyId) view returns(uint256)
func (_Blockchain *BlockchainCallerSession) GetChildCount(familyId *big.Int) (*big.Int, error) {
	return _Blockchain.Contract.GetChildCount(&_Blockchain.CallOpts, familyId)
}

// IsChild is a free data retrieval call binding the contract method 0xfc91a897.
//
// Solidity: function isChild(address account) view returns(bool)
func (_Blockchain *BlockchainCaller) IsChild(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "isChild", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsChild is a free data retrieval call binding the contract method 0xfc91a897.
//
// Solidity: function isChild(address account) view returns(bool)
func (_Blockchain *BlockchainSession) IsChild(account common.Address) (bool, error) {
	return _Blockchain.Contract.IsChild(&_Blockchain.CallOpts, account)
}

// IsChild is a free data retrieval call binding the contract method 0xfc91a897.
//
// Solidity: function isChild(address account) view returns(bool)
func (_Blockchain *BlockchainCallerSession) IsChild(account common.Address) (bool, error) {
	return _Blockchain.Contract.IsChild(&_Blockchain.CallOpts, account)
}

// IsParent is a free data retrieval call binding the contract method 0x0814ff13.
//
// Solidity: function isParent(address account) view returns(bool)
func (_Blockchain *BlockchainCaller) IsParent(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "isParent", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsParent is a free data retrieval call binding the contract method 0x0814ff13.
//
// Solidity: function isParent(address account) view returns(bool)
func (_Blockchain *BlockchainSession) IsParent(account common.Address) (bool, error) {
	return _Blockchain.Contract.IsParent(&_Blockchain.CallOpts, account)
}

// IsParent is a free data retrieval call binding the contract method 0x0814ff13.
//
// Solidity: function isParent(address account) view returns(bool)
func (_Blockchain *BlockchainCallerSession) IsParent(account common.Address) (bool, error) {
	return _Blockchain.Contract.IsParent(&_Blockchain.CallOpts, account)
}

// ParentFamilies is a free data retrieval call binding the contract method 0x5e2ad09b.
//
// Solidity: function parentFamilies(address , uint256 ) view returns(uint256)
func (_Blockchain *BlockchainCaller) ParentFamilies(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "parentFamilies", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ParentFamilies is a free data retrieval call binding the contract method 0x5e2ad09b.
//
// Solidity: function parentFamilies(address , uint256 ) view returns(uint256)
func (_Blockchain *BlockchainSession) ParentFamilies(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Blockchain.Contract.ParentFamilies(&_Blockchain.CallOpts, arg0, arg1)
}

// ParentFamilies is a free data retrieval call binding the contract method 0x5e2ad09b.
//
// Solidity: function parentFamilies(address , uint256 ) view returns(uint256)
func (_Blockchain *BlockchainCallerSession) ParentFamilies(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Blockchain.Contract.ParentFamilies(&_Blockchain.CallOpts, arg0, arg1)
}

// AddChild is a paid mutator transaction binding the contract method 0x74956e78.
//
// Solidity: function addChild(uint256 familyId, address childAddress, string name, uint8 age) returns()
func (_Blockchain *BlockchainTransactor) AddChild(opts *bind.TransactOpts, familyId *big.Int, childAddress common.Address, name string, age uint8) (*types.Transaction, error) {
	return _Blockchain.contract.Transact(opts, "addChild", familyId, childAddress, name, age)
}

// AddChild is a paid mutator transaction binding the contract method 0x74956e78.
//
// Solidity: function addChild(uint256 familyId, address childAddress, string name, uint8 age) returns()
func (_Blockchain *BlockchainSession) AddChild(familyId *big.Int, childAddress common.Address, name string, age uint8) (*types.Transaction, error) {
	return _Blockchain.Contract.AddChild(&_Blockchain.TransactOpts, familyId, childAddress, name, age)
}

// AddChild is a paid mutator transaction binding the contract method 0x74956e78.
//
// Solidity: function addChild(uint256 familyId, address childAddress, string name, uint8 age) returns()
func (_Blockchain *BlockchainTransactorSession) AddChild(familyId *big.Int, childAddress common.Address, name string, age uint8) (*types.Transaction, error) {
	return _Blockchain.Contract.AddChild(&_Blockchain.TransactOpts, familyId, childAddress, name, age)
}

// CreateFamily is a paid mutator transaction binding the contract method 0x9f9cc6a3.
//
// Solidity: function createFamily(string name) returns(uint256)
func (_Blockchain *BlockchainTransactor) CreateFamily(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _Blockchain.contract.Transact(opts, "createFamily", name)
}

// CreateFamily is a paid mutator transaction binding the contract method 0x9f9cc6a3.
//
// Solidity: function createFamily(string name) returns(uint256)
func (_Blockchain *BlockchainSession) CreateFamily(name string) (*types.Transaction, error) {
	return _Blockchain.Contract.CreateFamily(&_Blockchain.TransactOpts, name)
}

// CreateFamily is a paid mutator transaction binding the contract method 0x9f9cc6a3.
//
// Solidity: function createFamily(string name) returns(uint256)
func (_Blockchain *BlockchainTransactorSession) CreateFamily(name string) (*types.Transaction, error) {
	return _Blockchain.Contract.CreateFamily(&_Blockchain.TransactOpts, name)
}

// RemoveChild is a paid mutator transaction binding the contract method 0x87a69a7c.
//
// Solidity: function removeChild(uint256 familyId, address childAddress) returns()
func (_Blockchain *BlockchainTransactor) RemoveChild(opts *bind.TransactOpts, familyId *big.Int, childAddress common.Address) (*types.Transaction, error) {
	return _Blockchain.contract.Transact(opts, "removeChild", familyId, childAddress)
}

// RemoveChild is a paid mutator transaction binding the contract method 0x87a69a7c.
//
// Solidity: function removeChild(uint256 familyId, address childAddress) returns()
func (_Blockchain *BlockchainSession) RemoveChild(familyId *big.Int, childAddress common.Address) (*types.Transaction, error) {
	return _Blockchain.Contract.RemoveChild(&_Blockchain.TransactOpts, familyId, childAddress)
}

// RemoveChild is a paid mutator transaction binding the contract method 0x87a69a7c.
//
// Solidity: function removeChild(uint256 familyId, address childAddress) returns()
func (_Blockchain *BlockchainTransactorSession) RemoveChild(familyId *big.Int, childAddress common.Address) (*types.Transaction, error) {
	return _Blockchain.Contract.RemoveChild(&_Blockchain.TransactOpts, familyId, childAddress)
}

// UpdateFamily is a paid mutator transaction binding the contract method 0xd083e409.
//
// Solidity: function updateFamily(uint256 familyId, string name) returns()
func (_Blockchain *BlockchainTransactor) UpdateFamily(opts *bind.TransactOpts, familyId *big.Int, name string) (*types.Transaction, error) {
	return _Blockchain.contract.Transact(opts, "updateFamily", familyId, name)
}

// UpdateFamily is a paid mutator transaction binding the contract method 0xd083e409.
//
// Solidity: function updateFamily(uint256 familyId, string name) returns()
func (_Blockchain *BlockchainSession) UpdateFamily(familyId *big.Int, name string) (*types.Transaction, error) {
	return _Blockchain.Contract.UpdateFamily(&_Blockchain.TransactOpts, familyId, name)
}

// UpdateFamily is a paid mutator transaction binding the contract method 0xd083e409.
//
// Solidity: function updateFamily(uint256 familyId, string name) returns()
func (_Blockchain *BlockchainTransactorSession) UpdateFamily(familyId *big.Int, name string) (*types.Transaction, error) {
	return _Blockchain.Contract.UpdateFamily(&_Blockchain.TransactOpts, familyId, name)
}

// BlockchainChildAddedIterator is returned from FilterChildAdded and is used to iterate over the raw logs and unpacked data for ChildAdded events raised by the Blockchain contract.
type BlockchainChildAddedIterator struct {
	Event *BlockchainChildAdded // Event containing the contract specifics and raw log

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
func (it *BlockchainChildAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockchainChildAdded)
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
		it.Event = new(BlockchainChildAdded)
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
func (it *BlockchainChildAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockchainChildAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockchainChildAdded represents a ChildAdded event raised by the Blockchain contract.
type BlockchainChildAdded struct {
	FamilyId     *big.Int
	ChildAddress common.Address
	Name         string
	Age          uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterChildAdded is a free log retrieval operation binding the contract event 0xd9873979ce5fb9532e10a8a7b8abff463d0ec206b4de1c4bd377527f47725600.
//
// Solidity: event ChildAdded(uint256 indexed familyId, address indexed childAddress, string name, uint8 age)
func (_Blockchain *BlockchainFilterer) FilterChildAdded(opts *bind.FilterOpts, familyId []*big.Int, childAddress []common.Address) (*BlockchainChildAddedIterator, error) {

	var familyIdRule []interface{}
	for _, familyIdItem := range familyId {
		familyIdRule = append(familyIdRule, familyIdItem)
	}
	var childAddressRule []interface{}
	for _, childAddressItem := range childAddress {
		childAddressRule = append(childAddressRule, childAddressItem)
	}

	logs, sub, err := _Blockchain.contract.FilterLogs(opts, "ChildAdded", familyIdRule, childAddressRule)
	if err != nil {
		return nil, err
	}
	return &BlockchainChildAddedIterator{contract: _Blockchain.contract, event: "ChildAdded", logs: logs, sub: sub}, nil
}

// WatchChildAdded is a free log subscription operation binding the contract event 0xd9873979ce5fb9532e10a8a7b8abff463d0ec206b4de1c4bd377527f47725600.
//
// Solidity: event ChildAdded(uint256 indexed familyId, address indexed childAddress, string name, uint8 age)
func (_Blockchain *BlockchainFilterer) WatchChildAdded(opts *bind.WatchOpts, sink chan<- *BlockchainChildAdded, familyId []*big.Int, childAddress []common.Address) (event.Subscription, error) {

	var familyIdRule []interface{}
	for _, familyIdItem := range familyId {
		familyIdRule = append(familyIdRule, familyIdItem)
	}
	var childAddressRule []interface{}
	for _, childAddressItem := range childAddress {
		childAddressRule = append(childAddressRule, childAddressItem)
	}

	logs, sub, err := _Blockchain.contract.WatchLogs(opts, "ChildAdded", familyIdRule, childAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockchainChildAdded)
				if err := _Blockchain.contract.UnpackLog(event, "ChildAdded", log); err != nil {
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

// ParseChildAdded is a log parse operation binding the contract event 0xd9873979ce5fb9532e10a8a7b8abff463d0ec206b4de1c4bd377527f47725600.
//
// Solidity: event ChildAdded(uint256 indexed familyId, address indexed childAddress, string name, uint8 age)
func (_Blockchain *BlockchainFilterer) ParseChildAdded(log types.Log) (*BlockchainChildAdded, error) {
	event := new(BlockchainChildAdded)
	if err := _Blockchain.contract.UnpackLog(event, "ChildAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockchainChildRemovedIterator is returned from FilterChildRemoved and is used to iterate over the raw logs and unpacked data for ChildRemoved events raised by the Blockchain contract.
type BlockchainChildRemovedIterator struct {
	Event *BlockchainChildRemoved // Event containing the contract specifics and raw log

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
func (it *BlockchainChildRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockchainChildRemoved)
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
		it.Event = new(BlockchainChildRemoved)
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
func (it *BlockchainChildRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockchainChildRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockchainChildRemoved represents a ChildRemoved event raised by the Blockchain contract.
type BlockchainChildRemoved struct {
	FamilyId     *big.Int
	ChildAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterChildRemoved is a free log retrieval operation binding the contract event 0xba3641578d2900570c01f205ed35c07faca7ec9dac0e9e4f69b7d43fd3034668.
//
// Solidity: event ChildRemoved(uint256 indexed familyId, address indexed childAddress)
func (_Blockchain *BlockchainFilterer) FilterChildRemoved(opts *bind.FilterOpts, familyId []*big.Int, childAddress []common.Address) (*BlockchainChildRemovedIterator, error) {

	var familyIdRule []interface{}
	for _, familyIdItem := range familyId {
		familyIdRule = append(familyIdRule, familyIdItem)
	}
	var childAddressRule []interface{}
	for _, childAddressItem := range childAddress {
		childAddressRule = append(childAddressRule, childAddressItem)
	}

	logs, sub, err := _Blockchain.contract.FilterLogs(opts, "ChildRemoved", familyIdRule, childAddressRule)
	if err != nil {
		return nil, err
	}
	return &BlockchainChildRemovedIterator{contract: _Blockchain.contract, event: "ChildRemoved", logs: logs, sub: sub}, nil
}

// WatchChildRemoved is a free log subscription operation binding the contract event 0xba3641578d2900570c01f205ed35c07faca7ec9dac0e9e4f69b7d43fd3034668.
//
// Solidity: event ChildRemoved(uint256 indexed familyId, address indexed childAddress)
func (_Blockchain *BlockchainFilterer) WatchChildRemoved(opts *bind.WatchOpts, sink chan<- *BlockchainChildRemoved, familyId []*big.Int, childAddress []common.Address) (event.Subscription, error) {

	var familyIdRule []interface{}
	for _, familyIdItem := range familyId {
		familyIdRule = append(familyIdRule, familyIdItem)
	}
	var childAddressRule []interface{}
	for _, childAddressItem := range childAddress {
		childAddressRule = append(childAddressRule, childAddressItem)
	}

	logs, sub, err := _Blockchain.contract.WatchLogs(opts, "ChildRemoved", familyIdRule, childAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockchainChildRemoved)
				if err := _Blockchain.contract.UnpackLog(event, "ChildRemoved", log); err != nil {
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

// ParseChildRemoved is a log parse operation binding the contract event 0xba3641578d2900570c01f205ed35c07faca7ec9dac0e9e4f69b7d43fd3034668.
//
// Solidity: event ChildRemoved(uint256 indexed familyId, address indexed childAddress)
func (_Blockchain *BlockchainFilterer) ParseChildRemoved(log types.Log) (*BlockchainChildRemoved, error) {
	event := new(BlockchainChildRemoved)
	if err := _Blockchain.contract.UnpackLog(event, "ChildRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockchainFamilyCreatedIterator is returned from FilterFamilyCreated and is used to iterate over the raw logs and unpacked data for FamilyCreated events raised by the Blockchain contract.
type BlockchainFamilyCreatedIterator struct {
	Event *BlockchainFamilyCreated // Event containing the contract specifics and raw log

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
func (it *BlockchainFamilyCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockchainFamilyCreated)
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
		it.Event = new(BlockchainFamilyCreated)
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
func (it *BlockchainFamilyCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockchainFamilyCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockchainFamilyCreated represents a FamilyCreated event raised by the Blockchain contract.
type BlockchainFamilyCreated struct {
	FamilyId *big.Int
	Parent   common.Address
	Name     string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFamilyCreated is a free log retrieval operation binding the contract event 0xd7bd94563e5293bf92f0e48f697f09a7113742a3c029967ccbcd686cfadd0cbc.
//
// Solidity: event FamilyCreated(uint256 indexed familyId, address indexed parent, string name)
func (_Blockchain *BlockchainFilterer) FilterFamilyCreated(opts *bind.FilterOpts, familyId []*big.Int, parent []common.Address) (*BlockchainFamilyCreatedIterator, error) {

	var familyIdRule []interface{}
	for _, familyIdItem := range familyId {
		familyIdRule = append(familyIdRule, familyIdItem)
	}
	var parentRule []interface{}
	for _, parentItem := range parent {
		parentRule = append(parentRule, parentItem)
	}

	logs, sub, err := _Blockchain.contract.FilterLogs(opts, "FamilyCreated", familyIdRule, parentRule)
	if err != nil {
		return nil, err
	}
	return &BlockchainFamilyCreatedIterator{contract: _Blockchain.contract, event: "FamilyCreated", logs: logs, sub: sub}, nil
}

// WatchFamilyCreated is a free log subscription operation binding the contract event 0xd7bd94563e5293bf92f0e48f697f09a7113742a3c029967ccbcd686cfadd0cbc.
//
// Solidity: event FamilyCreated(uint256 indexed familyId, address indexed parent, string name)
func (_Blockchain *BlockchainFilterer) WatchFamilyCreated(opts *bind.WatchOpts, sink chan<- *BlockchainFamilyCreated, familyId []*big.Int, parent []common.Address) (event.Subscription, error) {

	var familyIdRule []interface{}
	for _, familyIdItem := range familyId {
		familyIdRule = append(familyIdRule, familyIdItem)
	}
	var parentRule []interface{}
	for _, parentItem := range parent {
		parentRule = append(parentRule, parentItem)
	}

	logs, sub, err := _Blockchain.contract.WatchLogs(opts, "FamilyCreated", familyIdRule, parentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockchainFamilyCreated)
				if err := _Blockchain.contract.UnpackLog(event, "FamilyCreated", log); err != nil {
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

// ParseFamilyCreated is a log parse operation binding the contract event 0xd7bd94563e5293bf92f0e48f697f09a7113742a3c029967ccbcd686cfadd0cbc.
//
// Solidity: event FamilyCreated(uint256 indexed familyId, address indexed parent, string name)
func (_Blockchain *BlockchainFilterer) ParseFamilyCreated(log types.Log) (*BlockchainFamilyCreated, error) {
	event := new(BlockchainFamilyCreated)
	if err := _Blockchain.contract.UnpackLog(event, "FamilyCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockchainFamilyUpdatedIterator is returned from FilterFamilyUpdated and is used to iterate over the raw logs and unpacked data for FamilyUpdated events raised by the Blockchain contract.
type BlockchainFamilyUpdatedIterator struct {
	Event *BlockchainFamilyUpdated // Event containing the contract specifics and raw log

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
func (it *BlockchainFamilyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockchainFamilyUpdated)
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
		it.Event = new(BlockchainFamilyUpdated)
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
func (it *BlockchainFamilyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockchainFamilyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockchainFamilyUpdated represents a FamilyUpdated event raised by the Blockchain contract.
type BlockchainFamilyUpdated struct {
	FamilyId *big.Int
	Name     string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFamilyUpdated is a free log retrieval operation binding the contract event 0x0111fc39a759d49b242d0f2a9db604740d45eae5a48cdb344cdff4ec3237ed0a.
//
// Solidity: event FamilyUpdated(uint256 indexed familyId, string name)
func (_Blockchain *BlockchainFilterer) FilterFamilyUpdated(opts *bind.FilterOpts, familyId []*big.Int) (*BlockchainFamilyUpdatedIterator, error) {

	var familyIdRule []interface{}
	for _, familyIdItem := range familyId {
		familyIdRule = append(familyIdRule, familyIdItem)
	}

	logs, sub, err := _Blockchain.contract.FilterLogs(opts, "FamilyUpdated", familyIdRule)
	if err != nil {
		return nil, err
	}
	return &BlockchainFamilyUpdatedIterator{contract: _Blockchain.contract, event: "FamilyUpdated", logs: logs, sub: sub}, nil
}

// WatchFamilyUpdated is a free log subscription operation binding the contract event 0x0111fc39a759d49b242d0f2a9db604740d45eae5a48cdb344cdff4ec3237ed0a.
//
// Solidity: event FamilyUpdated(uint256 indexed familyId, string name)
func (_Blockchain *BlockchainFilterer) WatchFamilyUpdated(opts *bind.WatchOpts, sink chan<- *BlockchainFamilyUpdated, familyId []*big.Int) (event.Subscription, error) {

	var familyIdRule []interface{}
	for _, familyIdItem := range familyId {
		familyIdRule = append(familyIdRule, familyIdItem)
	}

	logs, sub, err := _Blockchain.contract.WatchLogs(opts, "FamilyUpdated", familyIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockchainFamilyUpdated)
				if err := _Blockchain.contract.UnpackLog(event, "FamilyUpdated", log); err != nil {
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

// ParseFamilyUpdated is a log parse operation binding the contract event 0x0111fc39a759d49b242d0f2a9db604740d45eae5a48cdb344cdff4ec3237ed0a.
//
// Solidity: event FamilyUpdated(uint256 indexed familyId, string name)
func (_Blockchain *BlockchainFilterer) ParseFamilyUpdated(log types.Log) (*BlockchainFamilyUpdated, error) {
	event := new(BlockchainFamilyUpdated)
	if err := _Blockchain.contract.UnpackLog(event, "FamilyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
