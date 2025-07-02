// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package familyregistry

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

// FamilyRegistryMetaData contains all meta data concerning the FamilyRegistry contract.
var FamilyRegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"familyId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"childAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"age\",\"type\":\"uint8\"}],\"name\":\"ChildAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"familyId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"childAddress\",\"type\":\"address\"}],\"name\":\"ChildRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"familyId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"parent\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"FamilyCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"familyId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"FamilyUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"familyId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"childAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"age\",\"type\":\"uint8\"}],\"name\":\"addChild\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"childToFamily\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"createFamily\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"families\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"parent\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"familyCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"familyId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"childAddress\",\"type\":\"address\"}],\"name\":\"getChild\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"familyId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getChildAddressByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"familyId\",\"type\":\"uint256\"}],\"name\":\"getChildCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isChild\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isParent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"parentFamilies\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"familyId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"childAddress\",\"type\":\"address\"}],\"name\":\"removeChild\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"familyId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"updateFamily\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FamilyRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use FamilyRegistryMetaData.ABI instead.
var FamilyRegistryABI = FamilyRegistryMetaData.ABI

// FamilyRegistry is an auto generated Go binding around an Ethereum contract.
type FamilyRegistry struct {
	FamilyRegistryCaller     // Read-only binding to the contract
	FamilyRegistryTransactor // Write-only binding to the contract
	FamilyRegistryFilterer   // Log filterer for contract events
}

// FamilyRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type FamilyRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FamilyRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FamilyRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FamilyRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FamilyRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FamilyRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FamilyRegistrySession struct {
	Contract     *FamilyRegistry   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FamilyRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FamilyRegistryCallerSession struct {
	Contract *FamilyRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// FamilyRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FamilyRegistryTransactorSession struct {
	Contract     *FamilyRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// FamilyRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type FamilyRegistryRaw struct {
	Contract *FamilyRegistry // Generic contract binding to access the raw methods on
}

// FamilyRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FamilyRegistryCallerRaw struct {
	Contract *FamilyRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// FamilyRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FamilyRegistryTransactorRaw struct {
	Contract *FamilyRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFamilyRegistry creates a new instance of FamilyRegistry, bound to a specific deployed contract.
func NewFamilyRegistry(address common.Address, backend bind.ContractBackend) (*FamilyRegistry, error) {
	contract, err := bindFamilyRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FamilyRegistry{FamilyRegistryCaller: FamilyRegistryCaller{contract: contract}, FamilyRegistryTransactor: FamilyRegistryTransactor{contract: contract}, FamilyRegistryFilterer: FamilyRegistryFilterer{contract: contract}}, nil
}

// NewFamilyRegistryCaller creates a new read-only instance of FamilyRegistry, bound to a specific deployed contract.
func NewFamilyRegistryCaller(address common.Address, caller bind.ContractCaller) (*FamilyRegistryCaller, error) {
	contract, err := bindFamilyRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FamilyRegistryCaller{contract: contract}, nil
}

// NewFamilyRegistryTransactor creates a new write-only instance of FamilyRegistry, bound to a specific deployed contract.
func NewFamilyRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*FamilyRegistryTransactor, error) {
	contract, err := bindFamilyRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FamilyRegistryTransactor{contract: contract}, nil
}

// NewFamilyRegistryFilterer creates a new log filterer instance of FamilyRegistry, bound to a specific deployed contract.
func NewFamilyRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*FamilyRegistryFilterer, error) {
	contract, err := bindFamilyRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FamilyRegistryFilterer{contract: contract}, nil
}

// bindFamilyRegistry binds a generic wrapper to an already deployed contract.
func bindFamilyRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FamilyRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FamilyRegistry *FamilyRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FamilyRegistry.Contract.FamilyRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FamilyRegistry *FamilyRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FamilyRegistry.Contract.FamilyRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FamilyRegistry *FamilyRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FamilyRegistry.Contract.FamilyRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FamilyRegistry *FamilyRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FamilyRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FamilyRegistry *FamilyRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FamilyRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FamilyRegistry *FamilyRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FamilyRegistry.Contract.contract.Transact(opts, method, params...)
}

// ChildToFamily is a free data retrieval call binding the contract method 0xffa65a20.
//
// Solidity: function childToFamily(address ) view returns(uint256)
func (_FamilyRegistry *FamilyRegistryCaller) ChildToFamily(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FamilyRegistry.contract.Call(opts, &out, "childToFamily", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChildToFamily is a free data retrieval call binding the contract method 0xffa65a20.
//
// Solidity: function childToFamily(address ) view returns(uint256)
func (_FamilyRegistry *FamilyRegistrySession) ChildToFamily(arg0 common.Address) (*big.Int, error) {
	return _FamilyRegistry.Contract.ChildToFamily(&_FamilyRegistry.CallOpts, arg0)
}

// ChildToFamily is a free data retrieval call binding the contract method 0xffa65a20.
//
// Solidity: function childToFamily(address ) view returns(uint256)
func (_FamilyRegistry *FamilyRegistryCallerSession) ChildToFamily(arg0 common.Address) (*big.Int, error) {
	return _FamilyRegistry.Contract.ChildToFamily(&_FamilyRegistry.CallOpts, arg0)
}

// Families is a free data retrieval call binding the contract method 0x078beae5.
//
// Solidity: function families(uint256 ) view returns(uint256 id, address parent, string name, bool active)
func (_FamilyRegistry *FamilyRegistryCaller) Families(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id     *big.Int
	Parent common.Address
	Name   string
	Active bool
}, error) {
	var out []interface{}
	err := _FamilyRegistry.contract.Call(opts, &out, "families", arg0)

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
func (_FamilyRegistry *FamilyRegistrySession) Families(arg0 *big.Int) (struct {
	Id     *big.Int
	Parent common.Address
	Name   string
	Active bool
}, error) {
	return _FamilyRegistry.Contract.Families(&_FamilyRegistry.CallOpts, arg0)
}

// Families is a free data retrieval call binding the contract method 0x078beae5.
//
// Solidity: function families(uint256 ) view returns(uint256 id, address parent, string name, bool active)
func (_FamilyRegistry *FamilyRegistryCallerSession) Families(arg0 *big.Int) (struct {
	Id     *big.Int
	Parent common.Address
	Name   string
	Active bool
}, error) {
	return _FamilyRegistry.Contract.Families(&_FamilyRegistry.CallOpts, arg0)
}

// FamilyCount is a free data retrieval call binding the contract method 0x784f7a93.
//
// Solidity: function familyCount() view returns(uint256)
func (_FamilyRegistry *FamilyRegistryCaller) FamilyCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FamilyRegistry.contract.Call(opts, &out, "familyCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FamilyCount is a free data retrieval call binding the contract method 0x784f7a93.
//
// Solidity: function familyCount() view returns(uint256)
func (_FamilyRegistry *FamilyRegistrySession) FamilyCount() (*big.Int, error) {
	return _FamilyRegistry.Contract.FamilyCount(&_FamilyRegistry.CallOpts)
}

// FamilyCount is a free data retrieval call binding the contract method 0x784f7a93.
//
// Solidity: function familyCount() view returns(uint256)
func (_FamilyRegistry *FamilyRegistryCallerSession) FamilyCount() (*big.Int, error) {
	return _FamilyRegistry.Contract.FamilyCount(&_FamilyRegistry.CallOpts)
}

// GetChild is a free data retrieval call binding the contract method 0x9fa1140c.
//
// Solidity: function getChild(uint256 familyId, address childAddress) view returns(address, string, uint8, bool)
func (_FamilyRegistry *FamilyRegistryCaller) GetChild(opts *bind.CallOpts, familyId *big.Int, childAddress common.Address) (common.Address, string, uint8, bool, error) {
	var out []interface{}
	err := _FamilyRegistry.contract.Call(opts, &out, "getChild", familyId, childAddress)

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
func (_FamilyRegistry *FamilyRegistrySession) GetChild(familyId *big.Int, childAddress common.Address) (common.Address, string, uint8, bool, error) {
	return _FamilyRegistry.Contract.GetChild(&_FamilyRegistry.CallOpts, familyId, childAddress)
}

// GetChild is a free data retrieval call binding the contract method 0x9fa1140c.
//
// Solidity: function getChild(uint256 familyId, address childAddress) view returns(address, string, uint8, bool)
func (_FamilyRegistry *FamilyRegistryCallerSession) GetChild(familyId *big.Int, childAddress common.Address) (common.Address, string, uint8, bool, error) {
	return _FamilyRegistry.Contract.GetChild(&_FamilyRegistry.CallOpts, familyId, childAddress)
}

// GetChildAddressByIndex is a free data retrieval call binding the contract method 0x82ef100f.
//
// Solidity: function getChildAddressByIndex(uint256 familyId, uint256 index) view returns(address)
func (_FamilyRegistry *FamilyRegistryCaller) GetChildAddressByIndex(opts *bind.CallOpts, familyId *big.Int, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _FamilyRegistry.contract.Call(opts, &out, "getChildAddressByIndex", familyId, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetChildAddressByIndex is a free data retrieval call binding the contract method 0x82ef100f.
//
// Solidity: function getChildAddressByIndex(uint256 familyId, uint256 index) view returns(address)
func (_FamilyRegistry *FamilyRegistrySession) GetChildAddressByIndex(familyId *big.Int, index *big.Int) (common.Address, error) {
	return _FamilyRegistry.Contract.GetChildAddressByIndex(&_FamilyRegistry.CallOpts, familyId, index)
}

// GetChildAddressByIndex is a free data retrieval call binding the contract method 0x82ef100f.
//
// Solidity: function getChildAddressByIndex(uint256 familyId, uint256 index) view returns(address)
func (_FamilyRegistry *FamilyRegistryCallerSession) GetChildAddressByIndex(familyId *big.Int, index *big.Int) (common.Address, error) {
	return _FamilyRegistry.Contract.GetChildAddressByIndex(&_FamilyRegistry.CallOpts, familyId, index)
}

// GetChildCount is a free data retrieval call binding the contract method 0x7a77af5c.
//
// Solidity: function getChildCount(uint256 familyId) view returns(uint256)
func (_FamilyRegistry *FamilyRegistryCaller) GetChildCount(opts *bind.CallOpts, familyId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FamilyRegistry.contract.Call(opts, &out, "getChildCount", familyId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChildCount is a free data retrieval call binding the contract method 0x7a77af5c.
//
// Solidity: function getChildCount(uint256 familyId) view returns(uint256)
func (_FamilyRegistry *FamilyRegistrySession) GetChildCount(familyId *big.Int) (*big.Int, error) {
	return _FamilyRegistry.Contract.GetChildCount(&_FamilyRegistry.CallOpts, familyId)
}

// GetChildCount is a free data retrieval call binding the contract method 0x7a77af5c.
//
// Solidity: function getChildCount(uint256 familyId) view returns(uint256)
func (_FamilyRegistry *FamilyRegistryCallerSession) GetChildCount(familyId *big.Int) (*big.Int, error) {
	return _FamilyRegistry.Contract.GetChildCount(&_FamilyRegistry.CallOpts, familyId)
}

// IsChild is a free data retrieval call binding the contract method 0xfc91a897.
//
// Solidity: function isChild(address account) view returns(bool)
func (_FamilyRegistry *FamilyRegistryCaller) IsChild(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _FamilyRegistry.contract.Call(opts, &out, "isChild", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsChild is a free data retrieval call binding the contract method 0xfc91a897.
//
// Solidity: function isChild(address account) view returns(bool)
func (_FamilyRegistry *FamilyRegistrySession) IsChild(account common.Address) (bool, error) {
	return _FamilyRegistry.Contract.IsChild(&_FamilyRegistry.CallOpts, account)
}

// IsChild is a free data retrieval call binding the contract method 0xfc91a897.
//
// Solidity: function isChild(address account) view returns(bool)
func (_FamilyRegistry *FamilyRegistryCallerSession) IsChild(account common.Address) (bool, error) {
	return _FamilyRegistry.Contract.IsChild(&_FamilyRegistry.CallOpts, account)
}

// IsParent is a free data retrieval call binding the contract method 0x0814ff13.
//
// Solidity: function isParent(address account) view returns(bool)
func (_FamilyRegistry *FamilyRegistryCaller) IsParent(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _FamilyRegistry.contract.Call(opts, &out, "isParent", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsParent is a free data retrieval call binding the contract method 0x0814ff13.
//
// Solidity: function isParent(address account) view returns(bool)
func (_FamilyRegistry *FamilyRegistrySession) IsParent(account common.Address) (bool, error) {
	return _FamilyRegistry.Contract.IsParent(&_FamilyRegistry.CallOpts, account)
}

// IsParent is a free data retrieval call binding the contract method 0x0814ff13.
//
// Solidity: function isParent(address account) view returns(bool)
func (_FamilyRegistry *FamilyRegistryCallerSession) IsParent(account common.Address) (bool, error) {
	return _FamilyRegistry.Contract.IsParent(&_FamilyRegistry.CallOpts, account)
}

// ParentFamilies is a free data retrieval call binding the contract method 0x5e2ad09b.
//
// Solidity: function parentFamilies(address , uint256 ) view returns(uint256)
func (_FamilyRegistry *FamilyRegistryCaller) ParentFamilies(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FamilyRegistry.contract.Call(opts, &out, "parentFamilies", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ParentFamilies is a free data retrieval call binding the contract method 0x5e2ad09b.
//
// Solidity: function parentFamilies(address , uint256 ) view returns(uint256)
func (_FamilyRegistry *FamilyRegistrySession) ParentFamilies(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _FamilyRegistry.Contract.ParentFamilies(&_FamilyRegistry.CallOpts, arg0, arg1)
}

// ParentFamilies is a free data retrieval call binding the contract method 0x5e2ad09b.
//
// Solidity: function parentFamilies(address , uint256 ) view returns(uint256)
func (_FamilyRegistry *FamilyRegistryCallerSession) ParentFamilies(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _FamilyRegistry.Contract.ParentFamilies(&_FamilyRegistry.CallOpts, arg0, arg1)
}

// AddChild is a paid mutator transaction binding the contract method 0x74956e78.
//
// Solidity: function addChild(uint256 familyId, address childAddress, string name, uint8 age) returns()
func (_FamilyRegistry *FamilyRegistryTransactor) AddChild(opts *bind.TransactOpts, familyId *big.Int, childAddress common.Address, name string, age uint8) (*types.Transaction, error) {
	return _FamilyRegistry.contract.Transact(opts, "addChild", familyId, childAddress, name, age)
}

// AddChild is a paid mutator transaction binding the contract method 0x74956e78.
//
// Solidity: function addChild(uint256 familyId, address childAddress, string name, uint8 age) returns()
func (_FamilyRegistry *FamilyRegistrySession) AddChild(familyId *big.Int, childAddress common.Address, name string, age uint8) (*types.Transaction, error) {
	return _FamilyRegistry.Contract.AddChild(&_FamilyRegistry.TransactOpts, familyId, childAddress, name, age)
}

// AddChild is a paid mutator transaction binding the contract method 0x74956e78.
//
// Solidity: function addChild(uint256 familyId, address childAddress, string name, uint8 age) returns()
func (_FamilyRegistry *FamilyRegistryTransactorSession) AddChild(familyId *big.Int, childAddress common.Address, name string, age uint8) (*types.Transaction, error) {
	return _FamilyRegistry.Contract.AddChild(&_FamilyRegistry.TransactOpts, familyId, childAddress, name, age)
}

// CreateFamily is a paid mutator transaction binding the contract method 0x9f9cc6a3.
//
// Solidity: function createFamily(string name) returns(uint256)
func (_FamilyRegistry *FamilyRegistryTransactor) CreateFamily(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _FamilyRegistry.contract.Transact(opts, "createFamily", name)
}

// CreateFamily is a paid mutator transaction binding the contract method 0x9f9cc6a3.
//
// Solidity: function createFamily(string name) returns(uint256)
func (_FamilyRegistry *FamilyRegistrySession) CreateFamily(name string) (*types.Transaction, error) {
	return _FamilyRegistry.Contract.CreateFamily(&_FamilyRegistry.TransactOpts, name)
}

// CreateFamily is a paid mutator transaction binding the contract method 0x9f9cc6a3.
//
// Solidity: function createFamily(string name) returns(uint256)
func (_FamilyRegistry *FamilyRegistryTransactorSession) CreateFamily(name string) (*types.Transaction, error) {
	return _FamilyRegistry.Contract.CreateFamily(&_FamilyRegistry.TransactOpts, name)
}

// RemoveChild is a paid mutator transaction binding the contract method 0x87a69a7c.
//
// Solidity: function removeChild(uint256 familyId, address childAddress) returns()
func (_FamilyRegistry *FamilyRegistryTransactor) RemoveChild(opts *bind.TransactOpts, familyId *big.Int, childAddress common.Address) (*types.Transaction, error) {
	return _FamilyRegistry.contract.Transact(opts, "removeChild", familyId, childAddress)
}

// RemoveChild is a paid mutator transaction binding the contract method 0x87a69a7c.
//
// Solidity: function removeChild(uint256 familyId, address childAddress) returns()
func (_FamilyRegistry *FamilyRegistrySession) RemoveChild(familyId *big.Int, childAddress common.Address) (*types.Transaction, error) {
	return _FamilyRegistry.Contract.RemoveChild(&_FamilyRegistry.TransactOpts, familyId, childAddress)
}

// RemoveChild is a paid mutator transaction binding the contract method 0x87a69a7c.
//
// Solidity: function removeChild(uint256 familyId, address childAddress) returns()
func (_FamilyRegistry *FamilyRegistryTransactorSession) RemoveChild(familyId *big.Int, childAddress common.Address) (*types.Transaction, error) {
	return _FamilyRegistry.Contract.RemoveChild(&_FamilyRegistry.TransactOpts, familyId, childAddress)
}

// UpdateFamily is a paid mutator transaction binding the contract method 0xd083e409.
//
// Solidity: function updateFamily(uint256 familyId, string name) returns()
func (_FamilyRegistry *FamilyRegistryTransactor) UpdateFamily(opts *bind.TransactOpts, familyId *big.Int, name string) (*types.Transaction, error) {
	return _FamilyRegistry.contract.Transact(opts, "updateFamily", familyId, name)
}

// UpdateFamily is a paid mutator transaction binding the contract method 0xd083e409.
//
// Solidity: function updateFamily(uint256 familyId, string name) returns()
func (_FamilyRegistry *FamilyRegistrySession) UpdateFamily(familyId *big.Int, name string) (*types.Transaction, error) {
	return _FamilyRegistry.Contract.UpdateFamily(&_FamilyRegistry.TransactOpts, familyId, name)
}

// UpdateFamily is a paid mutator transaction binding the contract method 0xd083e409.
//
// Solidity: function updateFamily(uint256 familyId, string name) returns()
func (_FamilyRegistry *FamilyRegistryTransactorSession) UpdateFamily(familyId *big.Int, name string) (*types.Transaction, error) {
	return _FamilyRegistry.Contract.UpdateFamily(&_FamilyRegistry.TransactOpts, familyId, name)
}

// FamilyRegistryChildAddedIterator is returned from FilterChildAdded and is used to iterate over the raw logs and unpacked data for ChildAdded events raised by the FamilyRegistry contract.
type FamilyRegistryChildAddedIterator struct {
	Event *FamilyRegistryChildAdded // Event containing the contract specifics and raw log

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
func (it *FamilyRegistryChildAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FamilyRegistryChildAdded)
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
		it.Event = new(FamilyRegistryChildAdded)
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
func (it *FamilyRegistryChildAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FamilyRegistryChildAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FamilyRegistryChildAdded represents a ChildAdded event raised by the FamilyRegistry contract.
type FamilyRegistryChildAdded struct {
	FamilyId     *big.Int
	ChildAddress common.Address
	Name         string
	Age          uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterChildAdded is a free log retrieval operation binding the contract event 0xd9873979ce5fb9532e10a8a7b8abff463d0ec206b4de1c4bd377527f47725600.
//
// Solidity: event ChildAdded(uint256 indexed familyId, address indexed childAddress, string name, uint8 age)
func (_FamilyRegistry *FamilyRegistryFilterer) FilterChildAdded(opts *bind.FilterOpts, familyId []*big.Int, childAddress []common.Address) (*FamilyRegistryChildAddedIterator, error) {

	var familyIdRule []interface{}
	for _, familyIdItem := range familyId {
		familyIdRule = append(familyIdRule, familyIdItem)
	}
	var childAddressRule []interface{}
	for _, childAddressItem := range childAddress {
		childAddressRule = append(childAddressRule, childAddressItem)
	}

	logs, sub, err := _FamilyRegistry.contract.FilterLogs(opts, "ChildAdded", familyIdRule, childAddressRule)
	if err != nil {
		return nil, err
	}
	return &FamilyRegistryChildAddedIterator{contract: _FamilyRegistry.contract, event: "ChildAdded", logs: logs, sub: sub}, nil
}

// WatchChildAdded is a free log subscription operation binding the contract event 0xd9873979ce5fb9532e10a8a7b8abff463d0ec206b4de1c4bd377527f47725600.
//
// Solidity: event ChildAdded(uint256 indexed familyId, address indexed childAddress, string name, uint8 age)
func (_FamilyRegistry *FamilyRegistryFilterer) WatchChildAdded(opts *bind.WatchOpts, sink chan<- *FamilyRegistryChildAdded, familyId []*big.Int, childAddress []common.Address) (event.Subscription, error) {

	var familyIdRule []interface{}
	for _, familyIdItem := range familyId {
		familyIdRule = append(familyIdRule, familyIdItem)
	}
	var childAddressRule []interface{}
	for _, childAddressItem := range childAddress {
		childAddressRule = append(childAddressRule, childAddressItem)
	}

	logs, sub, err := _FamilyRegistry.contract.WatchLogs(opts, "ChildAdded", familyIdRule, childAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FamilyRegistryChildAdded)
				if err := _FamilyRegistry.contract.UnpackLog(event, "ChildAdded", log); err != nil {
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
func (_FamilyRegistry *FamilyRegistryFilterer) ParseChildAdded(log types.Log) (*FamilyRegistryChildAdded, error) {
	event := new(FamilyRegistryChildAdded)
	if err := _FamilyRegistry.contract.UnpackLog(event, "ChildAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FamilyRegistryChildRemovedIterator is returned from FilterChildRemoved and is used to iterate over the raw logs and unpacked data for ChildRemoved events raised by the FamilyRegistry contract.
type FamilyRegistryChildRemovedIterator struct {
	Event *FamilyRegistryChildRemoved // Event containing the contract specifics and raw log

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
func (it *FamilyRegistryChildRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FamilyRegistryChildRemoved)
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
		it.Event = new(FamilyRegistryChildRemoved)
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
func (it *FamilyRegistryChildRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FamilyRegistryChildRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FamilyRegistryChildRemoved represents a ChildRemoved event raised by the FamilyRegistry contract.
type FamilyRegistryChildRemoved struct {
	FamilyId     *big.Int
	ChildAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterChildRemoved is a free log retrieval operation binding the contract event 0xba3641578d2900570c01f205ed35c07faca7ec9dac0e9e4f69b7d43fd3034668.
//
// Solidity: event ChildRemoved(uint256 indexed familyId, address indexed childAddress)
func (_FamilyRegistry *FamilyRegistryFilterer) FilterChildRemoved(opts *bind.FilterOpts, familyId []*big.Int, childAddress []common.Address) (*FamilyRegistryChildRemovedIterator, error) {

	var familyIdRule []interface{}
	for _, familyIdItem := range familyId {
		familyIdRule = append(familyIdRule, familyIdItem)
	}
	var childAddressRule []interface{}
	for _, childAddressItem := range childAddress {
		childAddressRule = append(childAddressRule, childAddressItem)
	}

	logs, sub, err := _FamilyRegistry.contract.FilterLogs(opts, "ChildRemoved", familyIdRule, childAddressRule)
	if err != nil {
		return nil, err
	}
	return &FamilyRegistryChildRemovedIterator{contract: _FamilyRegistry.contract, event: "ChildRemoved", logs: logs, sub: sub}, nil
}

// WatchChildRemoved is a free log subscription operation binding the contract event 0xba3641578d2900570c01f205ed35c07faca7ec9dac0e9e4f69b7d43fd3034668.
//
// Solidity: event ChildRemoved(uint256 indexed familyId, address indexed childAddress)
func (_FamilyRegistry *FamilyRegistryFilterer) WatchChildRemoved(opts *bind.WatchOpts, sink chan<- *FamilyRegistryChildRemoved, familyId []*big.Int, childAddress []common.Address) (event.Subscription, error) {

	var familyIdRule []interface{}
	for _, familyIdItem := range familyId {
		familyIdRule = append(familyIdRule, familyIdItem)
	}
	var childAddressRule []interface{}
	for _, childAddressItem := range childAddress {
		childAddressRule = append(childAddressRule, childAddressItem)
	}

	logs, sub, err := _FamilyRegistry.contract.WatchLogs(opts, "ChildRemoved", familyIdRule, childAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FamilyRegistryChildRemoved)
				if err := _FamilyRegistry.contract.UnpackLog(event, "ChildRemoved", log); err != nil {
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
func (_FamilyRegistry *FamilyRegistryFilterer) ParseChildRemoved(log types.Log) (*FamilyRegistryChildRemoved, error) {
	event := new(FamilyRegistryChildRemoved)
	if err := _FamilyRegistry.contract.UnpackLog(event, "ChildRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FamilyRegistryFamilyCreatedIterator is returned from FilterFamilyCreated and is used to iterate over the raw logs and unpacked data for FamilyCreated events raised by the FamilyRegistry contract.
type FamilyRegistryFamilyCreatedIterator struct {
	Event *FamilyRegistryFamilyCreated // Event containing the contract specifics and raw log

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
func (it *FamilyRegistryFamilyCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FamilyRegistryFamilyCreated)
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
		it.Event = new(FamilyRegistryFamilyCreated)
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
func (it *FamilyRegistryFamilyCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FamilyRegistryFamilyCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FamilyRegistryFamilyCreated represents a FamilyCreated event raised by the FamilyRegistry contract.
type FamilyRegistryFamilyCreated struct {
	FamilyId *big.Int
	Parent   common.Address
	Name     string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFamilyCreated is a free log retrieval operation binding the contract event 0xd7bd94563e5293bf92f0e48f697f09a7113742a3c029967ccbcd686cfadd0cbc.
//
// Solidity: event FamilyCreated(uint256 indexed familyId, address indexed parent, string name)
func (_FamilyRegistry *FamilyRegistryFilterer) FilterFamilyCreated(opts *bind.FilterOpts, familyId []*big.Int, parent []common.Address) (*FamilyRegistryFamilyCreatedIterator, error) {

	var familyIdRule []interface{}
	for _, familyIdItem := range familyId {
		familyIdRule = append(familyIdRule, familyIdItem)
	}
	var parentRule []interface{}
	for _, parentItem := range parent {
		parentRule = append(parentRule, parentItem)
	}

	logs, sub, err := _FamilyRegistry.contract.FilterLogs(opts, "FamilyCreated", familyIdRule, parentRule)
	if err != nil {
		return nil, err
	}
	return &FamilyRegistryFamilyCreatedIterator{contract: _FamilyRegistry.contract, event: "FamilyCreated", logs: logs, sub: sub}, nil
}

// WatchFamilyCreated is a free log subscription operation binding the contract event 0xd7bd94563e5293bf92f0e48f697f09a7113742a3c029967ccbcd686cfadd0cbc.
//
// Solidity: event FamilyCreated(uint256 indexed familyId, address indexed parent, string name)
func (_FamilyRegistry *FamilyRegistryFilterer) WatchFamilyCreated(opts *bind.WatchOpts, sink chan<- *FamilyRegistryFamilyCreated, familyId []*big.Int, parent []common.Address) (event.Subscription, error) {

	var familyIdRule []interface{}
	for _, familyIdItem := range familyId {
		familyIdRule = append(familyIdRule, familyIdItem)
	}
	var parentRule []interface{}
	for _, parentItem := range parent {
		parentRule = append(parentRule, parentItem)
	}

	logs, sub, err := _FamilyRegistry.contract.WatchLogs(opts, "FamilyCreated", familyIdRule, parentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FamilyRegistryFamilyCreated)
				if err := _FamilyRegistry.contract.UnpackLog(event, "FamilyCreated", log); err != nil {
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
func (_FamilyRegistry *FamilyRegistryFilterer) ParseFamilyCreated(log types.Log) (*FamilyRegistryFamilyCreated, error) {
	event := new(FamilyRegistryFamilyCreated)
	if err := _FamilyRegistry.contract.UnpackLog(event, "FamilyCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FamilyRegistryFamilyUpdatedIterator is returned from FilterFamilyUpdated and is used to iterate over the raw logs and unpacked data for FamilyUpdated events raised by the FamilyRegistry contract.
type FamilyRegistryFamilyUpdatedIterator struct {
	Event *FamilyRegistryFamilyUpdated // Event containing the contract specifics and raw log

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
func (it *FamilyRegistryFamilyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FamilyRegistryFamilyUpdated)
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
		it.Event = new(FamilyRegistryFamilyUpdated)
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
func (it *FamilyRegistryFamilyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FamilyRegistryFamilyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FamilyRegistryFamilyUpdated represents a FamilyUpdated event raised by the FamilyRegistry contract.
type FamilyRegistryFamilyUpdated struct {
	FamilyId *big.Int
	Name     string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFamilyUpdated is a free log retrieval operation binding the contract event 0x0111fc39a759d49b242d0f2a9db604740d45eae5a48cdb344cdff4ec3237ed0a.
//
// Solidity: event FamilyUpdated(uint256 indexed familyId, string name)
func (_FamilyRegistry *FamilyRegistryFilterer) FilterFamilyUpdated(opts *bind.FilterOpts, familyId []*big.Int) (*FamilyRegistryFamilyUpdatedIterator, error) {

	var familyIdRule []interface{}
	for _, familyIdItem := range familyId {
		familyIdRule = append(familyIdRule, familyIdItem)
	}

	logs, sub, err := _FamilyRegistry.contract.FilterLogs(opts, "FamilyUpdated", familyIdRule)
	if err != nil {
		return nil, err
	}
	return &FamilyRegistryFamilyUpdatedIterator{contract: _FamilyRegistry.contract, event: "FamilyUpdated", logs: logs, sub: sub}, nil
}

// WatchFamilyUpdated is a free log subscription operation binding the contract event 0x0111fc39a759d49b242d0f2a9db604740d45eae5a48cdb344cdff4ec3237ed0a.
//
// Solidity: event FamilyUpdated(uint256 indexed familyId, string name)
func (_FamilyRegistry *FamilyRegistryFilterer) WatchFamilyUpdated(opts *bind.WatchOpts, sink chan<- *FamilyRegistryFamilyUpdated, familyId []*big.Int) (event.Subscription, error) {

	var familyIdRule []interface{}
	for _, familyIdItem := range familyId {
		familyIdRule = append(familyIdRule, familyIdItem)
	}

	logs, sub, err := _FamilyRegistry.contract.WatchLogs(opts, "FamilyUpdated", familyIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FamilyRegistryFamilyUpdated)
				if err := _FamilyRegistry.contract.UnpackLog(event, "FamilyUpdated", log); err != nil {
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
func (_FamilyRegistry *FamilyRegistryFilterer) ParseFamilyUpdated(log types.Log) (*FamilyRegistryFamilyUpdated, error) {
	event := new(FamilyRegistryFamilyUpdated)
	if err := _FamilyRegistry.contract.UnpackLog(event, "FamilyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
