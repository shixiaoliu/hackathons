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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"exchangeId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"parent\",\"type\":\"address\"}],\"name\":\"ExchangeFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"rewardId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"familyId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenPrice\",\"type\":\"uint256\"}],\"name\":\"RewardCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"exchangeId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"rewardId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"child\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"RewardExchanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"rewardId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"name\":\"RewardUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"childExchanges\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_familyId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_imageURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_tokenPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stock\",\"type\":\"uint256\"}],\"name\":\"createReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardId\",\"type\":\"uint256\"}],\"name\":\"exchangeReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"exchanges\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"child\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exchangeDate\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"fulfilled\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"familyRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_exchangeId\",\"type\":\"uint256\"}],\"name\":\"fulfillExchange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_child\",\"type\":\"address\"}],\"name\":\"getChildExchangeCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_child\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getChildExchangeId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_exchangeId\",\"type\":\"uint256\"}],\"name\":\"getExchange\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_familyId\",\"type\":\"uint256\"}],\"name\":\"getFamilyRewardCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_familyId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getFamilyRewardId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardId\",\"type\":\"uint256\"}],\"name\":\"getReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"familyId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"imageURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"tokenPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stock\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenContract\",\"outputs\":[{\"internalType\":\"contractRewardToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_imageURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_tokenPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stock\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_active\",\"type\":\"bool\"}],\"name\":\"updateReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// ChildExchanges is a free data retrieval call binding the contract method 0x83dbdb64.
//
// Solidity: function childExchanges(address , uint256 ) view returns(uint256)
func (_Blockchain *BlockchainCaller) ChildExchanges(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "childExchanges", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChildExchanges is a free data retrieval call binding the contract method 0x83dbdb64.
//
// Solidity: function childExchanges(address , uint256 ) view returns(uint256)
func (_Blockchain *BlockchainSession) ChildExchanges(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Blockchain.Contract.ChildExchanges(&_Blockchain.CallOpts, arg0, arg1)
}

// ChildExchanges is a free data retrieval call binding the contract method 0x83dbdb64.
//
// Solidity: function childExchanges(address , uint256 ) view returns(uint256)
func (_Blockchain *BlockchainCallerSession) ChildExchanges(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Blockchain.Contract.ChildExchanges(&_Blockchain.CallOpts, arg0, arg1)
}

// ExchangeCount is a free data retrieval call binding the contract method 0x68972e50.
//
// Solidity: function exchangeCount() view returns(uint256)
func (_Blockchain *BlockchainCaller) ExchangeCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "exchangeCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExchangeCount is a free data retrieval call binding the contract method 0x68972e50.
//
// Solidity: function exchangeCount() view returns(uint256)
func (_Blockchain *BlockchainSession) ExchangeCount() (*big.Int, error) {
	return _Blockchain.Contract.ExchangeCount(&_Blockchain.CallOpts)
}

// ExchangeCount is a free data retrieval call binding the contract method 0x68972e50.
//
// Solidity: function exchangeCount() view returns(uint256)
func (_Blockchain *BlockchainCallerSession) ExchangeCount() (*big.Int, error) {
	return _Blockchain.Contract.ExchangeCount(&_Blockchain.CallOpts)
}

// Exchanges is a free data retrieval call binding the contract method 0x2839fc29.
//
// Solidity: function exchanges(uint256 ) view returns(uint256 id, uint256 rewardId, address child, uint256 tokenAmount, uint256 exchangeDate, bool fulfilled)
func (_Blockchain *BlockchainCaller) Exchanges(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id           *big.Int
	RewardId     *big.Int
	Child        common.Address
	TokenAmount  *big.Int
	ExchangeDate *big.Int
	Fulfilled    bool
}, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "exchanges", arg0)

	outstruct := new(struct {
		Id           *big.Int
		RewardId     *big.Int
		Child        common.Address
		TokenAmount  *big.Int
		ExchangeDate *big.Int
		Fulfilled    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RewardId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Child = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.TokenAmount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ExchangeDate = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Fulfilled = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// Exchanges is a free data retrieval call binding the contract method 0x2839fc29.
//
// Solidity: function exchanges(uint256 ) view returns(uint256 id, uint256 rewardId, address child, uint256 tokenAmount, uint256 exchangeDate, bool fulfilled)
func (_Blockchain *BlockchainSession) Exchanges(arg0 *big.Int) (struct {
	Id           *big.Int
	RewardId     *big.Int
	Child        common.Address
	TokenAmount  *big.Int
	ExchangeDate *big.Int
	Fulfilled    bool
}, error) {
	return _Blockchain.Contract.Exchanges(&_Blockchain.CallOpts, arg0)
}

// Exchanges is a free data retrieval call binding the contract method 0x2839fc29.
//
// Solidity: function exchanges(uint256 ) view returns(uint256 id, uint256 rewardId, address child, uint256 tokenAmount, uint256 exchangeDate, bool fulfilled)
func (_Blockchain *BlockchainCallerSession) Exchanges(arg0 *big.Int) (struct {
	Id           *big.Int
	RewardId     *big.Int
	Child        common.Address
	TokenAmount  *big.Int
	ExchangeDate *big.Int
	Fulfilled    bool
}, error) {
	return _Blockchain.Contract.Exchanges(&_Blockchain.CallOpts, arg0)
}

// FamilyRewards is a free data retrieval call binding the contract method 0x9b676dc7.
//
// Solidity: function familyRewards(uint256 , uint256 ) view returns(uint256)
func (_Blockchain *BlockchainCaller) FamilyRewards(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "familyRewards", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FamilyRewards is a free data retrieval call binding the contract method 0x9b676dc7.
//
// Solidity: function familyRewards(uint256 , uint256 ) view returns(uint256)
func (_Blockchain *BlockchainSession) FamilyRewards(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Blockchain.Contract.FamilyRewards(&_Blockchain.CallOpts, arg0, arg1)
}

// FamilyRewards is a free data retrieval call binding the contract method 0x9b676dc7.
//
// Solidity: function familyRewards(uint256 , uint256 ) view returns(uint256)
func (_Blockchain *BlockchainCallerSession) FamilyRewards(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Blockchain.Contract.FamilyRewards(&_Blockchain.CallOpts, arg0, arg1)
}

// GetChildExchangeCount is a free data retrieval call binding the contract method 0x83e8ef54.
//
// Solidity: function getChildExchangeCount(address _child) view returns(uint256)
func (_Blockchain *BlockchainCaller) GetChildExchangeCount(opts *bind.CallOpts, _child common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "getChildExchangeCount", _child)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChildExchangeCount is a free data retrieval call binding the contract method 0x83e8ef54.
//
// Solidity: function getChildExchangeCount(address _child) view returns(uint256)
func (_Blockchain *BlockchainSession) GetChildExchangeCount(_child common.Address) (*big.Int, error) {
	return _Blockchain.Contract.GetChildExchangeCount(&_Blockchain.CallOpts, _child)
}

// GetChildExchangeCount is a free data retrieval call binding the contract method 0x83e8ef54.
//
// Solidity: function getChildExchangeCount(address _child) view returns(uint256)
func (_Blockchain *BlockchainCallerSession) GetChildExchangeCount(_child common.Address) (*big.Int, error) {
	return _Blockchain.Contract.GetChildExchangeCount(&_Blockchain.CallOpts, _child)
}

// GetChildExchangeId is a free data retrieval call binding the contract method 0xf59cc39a.
//
// Solidity: function getChildExchangeId(address _child, uint256 _index) view returns(uint256)
func (_Blockchain *BlockchainCaller) GetChildExchangeId(opts *bind.CallOpts, _child common.Address, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "getChildExchangeId", _child, _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChildExchangeId is a free data retrieval call binding the contract method 0xf59cc39a.
//
// Solidity: function getChildExchangeId(address _child, uint256 _index) view returns(uint256)
func (_Blockchain *BlockchainSession) GetChildExchangeId(_child common.Address, _index *big.Int) (*big.Int, error) {
	return _Blockchain.Contract.GetChildExchangeId(&_Blockchain.CallOpts, _child, _index)
}

// GetChildExchangeId is a free data retrieval call binding the contract method 0xf59cc39a.
//
// Solidity: function getChildExchangeId(address _child, uint256 _index) view returns(uint256)
func (_Blockchain *BlockchainCallerSession) GetChildExchangeId(_child common.Address, _index *big.Int) (*big.Int, error) {
	return _Blockchain.Contract.GetChildExchangeId(&_Blockchain.CallOpts, _child, _index)
}

// GetExchange is a free data retrieval call binding the contract method 0x0b9d5847.
//
// Solidity: function getExchange(uint256 _exchangeId) view returns(uint256, uint256, address, uint256, uint256, bool)
func (_Blockchain *BlockchainCaller) GetExchange(opts *bind.CallOpts, _exchangeId *big.Int) (*big.Int, *big.Int, common.Address, *big.Int, *big.Int, bool, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "getExchange", _exchangeId)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(common.Address), *new(*big.Int), *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(bool)).(*bool)

	return out0, out1, out2, out3, out4, out5, err

}

// GetExchange is a free data retrieval call binding the contract method 0x0b9d5847.
//
// Solidity: function getExchange(uint256 _exchangeId) view returns(uint256, uint256, address, uint256, uint256, bool)
func (_Blockchain *BlockchainSession) GetExchange(_exchangeId *big.Int) (*big.Int, *big.Int, common.Address, *big.Int, *big.Int, bool, error) {
	return _Blockchain.Contract.GetExchange(&_Blockchain.CallOpts, _exchangeId)
}

// GetExchange is a free data retrieval call binding the contract method 0x0b9d5847.
//
// Solidity: function getExchange(uint256 _exchangeId) view returns(uint256, uint256, address, uint256, uint256, bool)
func (_Blockchain *BlockchainCallerSession) GetExchange(_exchangeId *big.Int) (*big.Int, *big.Int, common.Address, *big.Int, *big.Int, bool, error) {
	return _Blockchain.Contract.GetExchange(&_Blockchain.CallOpts, _exchangeId)
}

// GetFamilyRewardCount is a free data retrieval call binding the contract method 0xe72f1b00.
//
// Solidity: function getFamilyRewardCount(uint256 _familyId) view returns(uint256)
func (_Blockchain *BlockchainCaller) GetFamilyRewardCount(opts *bind.CallOpts, _familyId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "getFamilyRewardCount", _familyId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFamilyRewardCount is a free data retrieval call binding the contract method 0xe72f1b00.
//
// Solidity: function getFamilyRewardCount(uint256 _familyId) view returns(uint256)
func (_Blockchain *BlockchainSession) GetFamilyRewardCount(_familyId *big.Int) (*big.Int, error) {
	return _Blockchain.Contract.GetFamilyRewardCount(&_Blockchain.CallOpts, _familyId)
}

// GetFamilyRewardCount is a free data retrieval call binding the contract method 0xe72f1b00.
//
// Solidity: function getFamilyRewardCount(uint256 _familyId) view returns(uint256)
func (_Blockchain *BlockchainCallerSession) GetFamilyRewardCount(_familyId *big.Int) (*big.Int, error) {
	return _Blockchain.Contract.GetFamilyRewardCount(&_Blockchain.CallOpts, _familyId)
}

// GetFamilyRewardId is a free data retrieval call binding the contract method 0x77936d1a.
//
// Solidity: function getFamilyRewardId(uint256 _familyId, uint256 _index) view returns(uint256)
func (_Blockchain *BlockchainCaller) GetFamilyRewardId(opts *bind.CallOpts, _familyId *big.Int, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "getFamilyRewardId", _familyId, _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFamilyRewardId is a free data retrieval call binding the contract method 0x77936d1a.
//
// Solidity: function getFamilyRewardId(uint256 _familyId, uint256 _index) view returns(uint256)
func (_Blockchain *BlockchainSession) GetFamilyRewardId(_familyId *big.Int, _index *big.Int) (*big.Int, error) {
	return _Blockchain.Contract.GetFamilyRewardId(&_Blockchain.CallOpts, _familyId, _index)
}

// GetFamilyRewardId is a free data retrieval call binding the contract method 0x77936d1a.
//
// Solidity: function getFamilyRewardId(uint256 _familyId, uint256 _index) view returns(uint256)
func (_Blockchain *BlockchainCallerSession) GetFamilyRewardId(_familyId *big.Int, _index *big.Int) (*big.Int, error) {
	return _Blockchain.Contract.GetFamilyRewardId(&_Blockchain.CallOpts, _familyId, _index)
}

// GetReward is a free data retrieval call binding the contract method 0x1c4b774b.
//
// Solidity: function getReward(uint256 _rewardId) view returns(uint256, address, uint256, string, string, string, uint256, uint256, bool)
func (_Blockchain *BlockchainCaller) GetReward(opts *bind.CallOpts, _rewardId *big.Int) (*big.Int, common.Address, *big.Int, string, string, string, *big.Int, *big.Int, bool, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "getReward", _rewardId)

	if err != nil {
		return *new(*big.Int), *new(common.Address), *new(*big.Int), *new(string), *new(string), *new(string), *new(*big.Int), *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(string)).(*string)
	out4 := *abi.ConvertType(out[4], new(string)).(*string)
	out5 := *abi.ConvertType(out[5], new(string)).(*string)
	out6 := *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	out7 := *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	out8 := *abi.ConvertType(out[8], new(bool)).(*bool)

	return out0, out1, out2, out3, out4, out5, out6, out7, out8, err

}

// GetReward is a free data retrieval call binding the contract method 0x1c4b774b.
//
// Solidity: function getReward(uint256 _rewardId) view returns(uint256, address, uint256, string, string, string, uint256, uint256, bool)
func (_Blockchain *BlockchainSession) GetReward(_rewardId *big.Int) (*big.Int, common.Address, *big.Int, string, string, string, *big.Int, *big.Int, bool, error) {
	return _Blockchain.Contract.GetReward(&_Blockchain.CallOpts, _rewardId)
}

// GetReward is a free data retrieval call binding the contract method 0x1c4b774b.
//
// Solidity: function getReward(uint256 _rewardId) view returns(uint256, address, uint256, string, string, string, uint256, uint256, bool)
func (_Blockchain *BlockchainCallerSession) GetReward(_rewardId *big.Int) (*big.Int, common.Address, *big.Int, string, string, string, *big.Int, *big.Int, bool, error) {
	return _Blockchain.Contract.GetReward(&_Blockchain.CallOpts, _rewardId)
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

// RewardCount is a free data retrieval call binding the contract method 0x79085425.
//
// Solidity: function rewardCount() view returns(uint256)
func (_Blockchain *BlockchainCaller) RewardCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "rewardCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardCount is a free data retrieval call binding the contract method 0x79085425.
//
// Solidity: function rewardCount() view returns(uint256)
func (_Blockchain *BlockchainSession) RewardCount() (*big.Int, error) {
	return _Blockchain.Contract.RewardCount(&_Blockchain.CallOpts)
}

// RewardCount is a free data retrieval call binding the contract method 0x79085425.
//
// Solidity: function rewardCount() view returns(uint256)
func (_Blockchain *BlockchainCallerSession) RewardCount() (*big.Int, error) {
	return _Blockchain.Contract.RewardCount(&_Blockchain.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0xf301af42.
//
// Solidity: function rewards(uint256 ) view returns(uint256 id, address creator, uint256 familyId, string name, string description, string imageURI, uint256 tokenPrice, uint256 stock, bool active)
func (_Blockchain *BlockchainCaller) Rewards(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id          *big.Int
	Creator     common.Address
	FamilyId    *big.Int
	Name        string
	Description string
	ImageURI    string
	TokenPrice  *big.Int
	Stock       *big.Int
	Active      bool
}, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "rewards", arg0)

	outstruct := new(struct {
		Id          *big.Int
		Creator     common.Address
		FamilyId    *big.Int
		Name        string
		Description string
		ImageURI    string
		TokenPrice  *big.Int
		Stock       *big.Int
		Active      bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Creator = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.FamilyId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Name = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Description = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.ImageURI = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.TokenPrice = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Stock = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.Active = *abi.ConvertType(out[8], new(bool)).(*bool)

	return *outstruct, err

}

// Rewards is a free data retrieval call binding the contract method 0xf301af42.
//
// Solidity: function rewards(uint256 ) view returns(uint256 id, address creator, uint256 familyId, string name, string description, string imageURI, uint256 tokenPrice, uint256 stock, bool active)
func (_Blockchain *BlockchainSession) Rewards(arg0 *big.Int) (struct {
	Id          *big.Int
	Creator     common.Address
	FamilyId    *big.Int
	Name        string
	Description string
	ImageURI    string
	TokenPrice  *big.Int
	Stock       *big.Int
	Active      bool
}, error) {
	return _Blockchain.Contract.Rewards(&_Blockchain.CallOpts, arg0)
}

// Rewards is a free data retrieval call binding the contract method 0xf301af42.
//
// Solidity: function rewards(uint256 ) view returns(uint256 id, address creator, uint256 familyId, string name, string description, string imageURI, uint256 tokenPrice, uint256 stock, bool active)
func (_Blockchain *BlockchainCallerSession) Rewards(arg0 *big.Int) (struct {
	Id          *big.Int
	Creator     common.Address
	FamilyId    *big.Int
	Name        string
	Description string
	ImageURI    string
	TokenPrice  *big.Int
	Stock       *big.Int
	Active      bool
}, error) {
	return _Blockchain.Contract.Rewards(&_Blockchain.CallOpts, arg0)
}

// TokenContract is a free data retrieval call binding the contract method 0x55a373d6.
//
// Solidity: function tokenContract() view returns(address)
func (_Blockchain *BlockchainCaller) TokenContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "tokenContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenContract is a free data retrieval call binding the contract method 0x55a373d6.
//
// Solidity: function tokenContract() view returns(address)
func (_Blockchain *BlockchainSession) TokenContract() (common.Address, error) {
	return _Blockchain.Contract.TokenContract(&_Blockchain.CallOpts)
}

// TokenContract is a free data retrieval call binding the contract method 0x55a373d6.
//
// Solidity: function tokenContract() view returns(address)
func (_Blockchain *BlockchainCallerSession) TokenContract() (common.Address, error) {
	return _Blockchain.Contract.TokenContract(&_Blockchain.CallOpts)
}

// CreateReward is a paid mutator transaction binding the contract method 0x80421276.
//
// Solidity: function createReward(uint256 _familyId, string _name, string _description, string _imageURI, uint256 _tokenPrice, uint256 _stock) returns(uint256)
func (_Blockchain *BlockchainTransactor) CreateReward(opts *bind.TransactOpts, _familyId *big.Int, _name string, _description string, _imageURI string, _tokenPrice *big.Int, _stock *big.Int) (*types.Transaction, error) {
	return _Blockchain.contract.Transact(opts, "createReward", _familyId, _name, _description, _imageURI, _tokenPrice, _stock)
}

// CreateReward is a paid mutator transaction binding the contract method 0x80421276.
//
// Solidity: function createReward(uint256 _familyId, string _name, string _description, string _imageURI, uint256 _tokenPrice, uint256 _stock) returns(uint256)
func (_Blockchain *BlockchainSession) CreateReward(_familyId *big.Int, _name string, _description string, _imageURI string, _tokenPrice *big.Int, _stock *big.Int) (*types.Transaction, error) {
	return _Blockchain.Contract.CreateReward(&_Blockchain.TransactOpts, _familyId, _name, _description, _imageURI, _tokenPrice, _stock)
}

// CreateReward is a paid mutator transaction binding the contract method 0x80421276.
//
// Solidity: function createReward(uint256 _familyId, string _name, string _description, string _imageURI, uint256 _tokenPrice, uint256 _stock) returns(uint256)
func (_Blockchain *BlockchainTransactorSession) CreateReward(_familyId *big.Int, _name string, _description string, _imageURI string, _tokenPrice *big.Int, _stock *big.Int) (*types.Transaction, error) {
	return _Blockchain.Contract.CreateReward(&_Blockchain.TransactOpts, _familyId, _name, _description, _imageURI, _tokenPrice, _stock)
}

// ExchangeReward is a paid mutator transaction binding the contract method 0x6a60630d.
//
// Solidity: function exchangeReward(uint256 _rewardId) returns(uint256)
func (_Blockchain *BlockchainTransactor) ExchangeReward(opts *bind.TransactOpts, _rewardId *big.Int) (*types.Transaction, error) {
	return _Blockchain.contract.Transact(opts, "exchangeReward", _rewardId)
}

// ExchangeReward is a paid mutator transaction binding the contract method 0x6a60630d.
//
// Solidity: function exchangeReward(uint256 _rewardId) returns(uint256)
func (_Blockchain *BlockchainSession) ExchangeReward(_rewardId *big.Int) (*types.Transaction, error) {
	return _Blockchain.Contract.ExchangeReward(&_Blockchain.TransactOpts, _rewardId)
}

// ExchangeReward is a paid mutator transaction binding the contract method 0x6a60630d.
//
// Solidity: function exchangeReward(uint256 _rewardId) returns(uint256)
func (_Blockchain *BlockchainTransactorSession) ExchangeReward(_rewardId *big.Int) (*types.Transaction, error) {
	return _Blockchain.Contract.ExchangeReward(&_Blockchain.TransactOpts, _rewardId)
}

// FulfillExchange is a paid mutator transaction binding the contract method 0x0670ea71.
//
// Solidity: function fulfillExchange(uint256 _exchangeId) returns()
func (_Blockchain *BlockchainTransactor) FulfillExchange(opts *bind.TransactOpts, _exchangeId *big.Int) (*types.Transaction, error) {
	return _Blockchain.contract.Transact(opts, "fulfillExchange", _exchangeId)
}

// FulfillExchange is a paid mutator transaction binding the contract method 0x0670ea71.
//
// Solidity: function fulfillExchange(uint256 _exchangeId) returns()
func (_Blockchain *BlockchainSession) FulfillExchange(_exchangeId *big.Int) (*types.Transaction, error) {
	return _Blockchain.Contract.FulfillExchange(&_Blockchain.TransactOpts, _exchangeId)
}

// FulfillExchange is a paid mutator transaction binding the contract method 0x0670ea71.
//
// Solidity: function fulfillExchange(uint256 _exchangeId) returns()
func (_Blockchain *BlockchainTransactorSession) FulfillExchange(_exchangeId *big.Int) (*types.Transaction, error) {
	return _Blockchain.Contract.FulfillExchange(&_Blockchain.TransactOpts, _exchangeId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Blockchain *BlockchainTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blockchain.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Blockchain *BlockchainSession) RenounceOwnership() (*types.Transaction, error) {
	return _Blockchain.Contract.RenounceOwnership(&_Blockchain.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Blockchain *BlockchainTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Blockchain.Contract.RenounceOwnership(&_Blockchain.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Blockchain *BlockchainTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Blockchain.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Blockchain *BlockchainSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Blockchain.Contract.TransferOwnership(&_Blockchain.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Blockchain *BlockchainTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Blockchain.Contract.TransferOwnership(&_Blockchain.TransactOpts, newOwner)
}

// UpdateReward is a paid mutator transaction binding the contract method 0x1da30cc5.
//
// Solidity: function updateReward(uint256 _rewardId, string _name, string _description, string _imageURI, uint256 _tokenPrice, uint256 _stock, bool _active) returns()
func (_Blockchain *BlockchainTransactor) UpdateReward(opts *bind.TransactOpts, _rewardId *big.Int, _name string, _description string, _imageURI string, _tokenPrice *big.Int, _stock *big.Int, _active bool) (*types.Transaction, error) {
	return _Blockchain.contract.Transact(opts, "updateReward", _rewardId, _name, _description, _imageURI, _tokenPrice, _stock, _active)
}

// UpdateReward is a paid mutator transaction binding the contract method 0x1da30cc5.
//
// Solidity: function updateReward(uint256 _rewardId, string _name, string _description, string _imageURI, uint256 _tokenPrice, uint256 _stock, bool _active) returns()
func (_Blockchain *BlockchainSession) UpdateReward(_rewardId *big.Int, _name string, _description string, _imageURI string, _tokenPrice *big.Int, _stock *big.Int, _active bool) (*types.Transaction, error) {
	return _Blockchain.Contract.UpdateReward(&_Blockchain.TransactOpts, _rewardId, _name, _description, _imageURI, _tokenPrice, _stock, _active)
}

// UpdateReward is a paid mutator transaction binding the contract method 0x1da30cc5.
//
// Solidity: function updateReward(uint256 _rewardId, string _name, string _description, string _imageURI, uint256 _tokenPrice, uint256 _stock, bool _active) returns()
func (_Blockchain *BlockchainTransactorSession) UpdateReward(_rewardId *big.Int, _name string, _description string, _imageURI string, _tokenPrice *big.Int, _stock *big.Int, _active bool) (*types.Transaction, error) {
	return _Blockchain.Contract.UpdateReward(&_Blockchain.TransactOpts, _rewardId, _name, _description, _imageURI, _tokenPrice, _stock, _active)
}

// BlockchainExchangeFulfilledIterator is returned from FilterExchangeFulfilled and is used to iterate over the raw logs and unpacked data for ExchangeFulfilled events raised by the Blockchain contract.
type BlockchainExchangeFulfilledIterator struct {
	Event *BlockchainExchangeFulfilled // Event containing the contract specifics and raw log

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
func (it *BlockchainExchangeFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockchainExchangeFulfilled)
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
		it.Event = new(BlockchainExchangeFulfilled)
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
func (it *BlockchainExchangeFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockchainExchangeFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockchainExchangeFulfilled represents a ExchangeFulfilled event raised by the Blockchain contract.
type BlockchainExchangeFulfilled struct {
	ExchangeId *big.Int
	Parent     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterExchangeFulfilled is a free log retrieval operation binding the contract event 0x362de7542ac4a995332d8029cd2fcb9c46c9668b23c962e0dba8c1a7f28a28e6.
//
// Solidity: event ExchangeFulfilled(uint256 indexed exchangeId, address indexed parent)
func (_Blockchain *BlockchainFilterer) FilterExchangeFulfilled(opts *bind.FilterOpts, exchangeId []*big.Int, parent []common.Address) (*BlockchainExchangeFulfilledIterator, error) {

	var exchangeIdRule []interface{}
	for _, exchangeIdItem := range exchangeId {
		exchangeIdRule = append(exchangeIdRule, exchangeIdItem)
	}
	var parentRule []interface{}
	for _, parentItem := range parent {
		parentRule = append(parentRule, parentItem)
	}

	logs, sub, err := _Blockchain.contract.FilterLogs(opts, "ExchangeFulfilled", exchangeIdRule, parentRule)
	if err != nil {
		return nil, err
	}
	return &BlockchainExchangeFulfilledIterator{contract: _Blockchain.contract, event: "ExchangeFulfilled", logs: logs, sub: sub}, nil
}

// WatchExchangeFulfilled is a free log subscription operation binding the contract event 0x362de7542ac4a995332d8029cd2fcb9c46c9668b23c962e0dba8c1a7f28a28e6.
//
// Solidity: event ExchangeFulfilled(uint256 indexed exchangeId, address indexed parent)
func (_Blockchain *BlockchainFilterer) WatchExchangeFulfilled(opts *bind.WatchOpts, sink chan<- *BlockchainExchangeFulfilled, exchangeId []*big.Int, parent []common.Address) (event.Subscription, error) {

	var exchangeIdRule []interface{}
	for _, exchangeIdItem := range exchangeId {
		exchangeIdRule = append(exchangeIdRule, exchangeIdItem)
	}
	var parentRule []interface{}
	for _, parentItem := range parent {
		parentRule = append(parentRule, parentItem)
	}

	logs, sub, err := _Blockchain.contract.WatchLogs(opts, "ExchangeFulfilled", exchangeIdRule, parentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockchainExchangeFulfilled)
				if err := _Blockchain.contract.UnpackLog(event, "ExchangeFulfilled", log); err != nil {
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

// ParseExchangeFulfilled is a log parse operation binding the contract event 0x362de7542ac4a995332d8029cd2fcb9c46c9668b23c962e0dba8c1a7f28a28e6.
//
// Solidity: event ExchangeFulfilled(uint256 indexed exchangeId, address indexed parent)
func (_Blockchain *BlockchainFilterer) ParseExchangeFulfilled(log types.Log) (*BlockchainExchangeFulfilled, error) {
	event := new(BlockchainExchangeFulfilled)
	if err := _Blockchain.contract.UnpackLog(event, "ExchangeFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockchainOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Blockchain contract.
type BlockchainOwnershipTransferredIterator struct {
	Event *BlockchainOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BlockchainOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockchainOwnershipTransferred)
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
		it.Event = new(BlockchainOwnershipTransferred)
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
func (it *BlockchainOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockchainOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockchainOwnershipTransferred represents a OwnershipTransferred event raised by the Blockchain contract.
type BlockchainOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Blockchain *BlockchainFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BlockchainOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Blockchain.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BlockchainOwnershipTransferredIterator{contract: _Blockchain.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Blockchain *BlockchainFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BlockchainOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Blockchain.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockchainOwnershipTransferred)
				if err := _Blockchain.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Blockchain *BlockchainFilterer) ParseOwnershipTransferred(log types.Log) (*BlockchainOwnershipTransferred, error) {
	event := new(BlockchainOwnershipTransferred)
	if err := _Blockchain.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockchainRewardCreatedIterator is returned from FilterRewardCreated and is used to iterate over the raw logs and unpacked data for RewardCreated events raised by the Blockchain contract.
type BlockchainRewardCreatedIterator struct {
	Event *BlockchainRewardCreated // Event containing the contract specifics and raw log

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
func (it *BlockchainRewardCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockchainRewardCreated)
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
		it.Event = new(BlockchainRewardCreated)
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
func (it *BlockchainRewardCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockchainRewardCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockchainRewardCreated represents a RewardCreated event raised by the Blockchain contract.
type BlockchainRewardCreated struct {
	RewardId   *big.Int
	Creator    common.Address
	FamilyId   *big.Int
	Name       string
	TokenPrice *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRewardCreated is a free log retrieval operation binding the contract event 0xb96d3f67d5ecf915732544d895500cf70d6ab68bccb6e156ce171f20019ec116.
//
// Solidity: event RewardCreated(uint256 indexed rewardId, address indexed creator, uint256 indexed familyId, string name, uint256 tokenPrice)
func (_Blockchain *BlockchainFilterer) FilterRewardCreated(opts *bind.FilterOpts, rewardId []*big.Int, creator []common.Address, familyId []*big.Int) (*BlockchainRewardCreatedIterator, error) {

	var rewardIdRule []interface{}
	for _, rewardIdItem := range rewardId {
		rewardIdRule = append(rewardIdRule, rewardIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var familyIdRule []interface{}
	for _, familyIdItem := range familyId {
		familyIdRule = append(familyIdRule, familyIdItem)
	}

	logs, sub, err := _Blockchain.contract.FilterLogs(opts, "RewardCreated", rewardIdRule, creatorRule, familyIdRule)
	if err != nil {
		return nil, err
	}
	return &BlockchainRewardCreatedIterator{contract: _Blockchain.contract, event: "RewardCreated", logs: logs, sub: sub}, nil
}

// WatchRewardCreated is a free log subscription operation binding the contract event 0xb96d3f67d5ecf915732544d895500cf70d6ab68bccb6e156ce171f20019ec116.
//
// Solidity: event RewardCreated(uint256 indexed rewardId, address indexed creator, uint256 indexed familyId, string name, uint256 tokenPrice)
func (_Blockchain *BlockchainFilterer) WatchRewardCreated(opts *bind.WatchOpts, sink chan<- *BlockchainRewardCreated, rewardId []*big.Int, creator []common.Address, familyId []*big.Int) (event.Subscription, error) {

	var rewardIdRule []interface{}
	for _, rewardIdItem := range rewardId {
		rewardIdRule = append(rewardIdRule, rewardIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var familyIdRule []interface{}
	for _, familyIdItem := range familyId {
		familyIdRule = append(familyIdRule, familyIdItem)
	}

	logs, sub, err := _Blockchain.contract.WatchLogs(opts, "RewardCreated", rewardIdRule, creatorRule, familyIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockchainRewardCreated)
				if err := _Blockchain.contract.UnpackLog(event, "RewardCreated", log); err != nil {
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

// ParseRewardCreated is a log parse operation binding the contract event 0xb96d3f67d5ecf915732544d895500cf70d6ab68bccb6e156ce171f20019ec116.
//
// Solidity: event RewardCreated(uint256 indexed rewardId, address indexed creator, uint256 indexed familyId, string name, uint256 tokenPrice)
func (_Blockchain *BlockchainFilterer) ParseRewardCreated(log types.Log) (*BlockchainRewardCreated, error) {
	event := new(BlockchainRewardCreated)
	if err := _Blockchain.contract.UnpackLog(event, "RewardCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockchainRewardExchangedIterator is returned from FilterRewardExchanged and is used to iterate over the raw logs and unpacked data for RewardExchanged events raised by the Blockchain contract.
type BlockchainRewardExchangedIterator struct {
	Event *BlockchainRewardExchanged // Event containing the contract specifics and raw log

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
func (it *BlockchainRewardExchangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockchainRewardExchanged)
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
		it.Event = new(BlockchainRewardExchanged)
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
func (it *BlockchainRewardExchangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockchainRewardExchangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockchainRewardExchanged represents a RewardExchanged event raised by the Blockchain contract.
type BlockchainRewardExchanged struct {
	ExchangeId  *big.Int
	RewardId    *big.Int
	Child       common.Address
	TokenAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewardExchanged is a free log retrieval operation binding the contract event 0xe150ee44da4a37592e6368bbdc0d8901353c1bfb770394b3268e70c9b5055050.
//
// Solidity: event RewardExchanged(uint256 indexed exchangeId, uint256 indexed rewardId, address indexed child, uint256 tokenAmount)
func (_Blockchain *BlockchainFilterer) FilterRewardExchanged(opts *bind.FilterOpts, exchangeId []*big.Int, rewardId []*big.Int, child []common.Address) (*BlockchainRewardExchangedIterator, error) {

	var exchangeIdRule []interface{}
	for _, exchangeIdItem := range exchangeId {
		exchangeIdRule = append(exchangeIdRule, exchangeIdItem)
	}
	var rewardIdRule []interface{}
	for _, rewardIdItem := range rewardId {
		rewardIdRule = append(rewardIdRule, rewardIdItem)
	}
	var childRule []interface{}
	for _, childItem := range child {
		childRule = append(childRule, childItem)
	}

	logs, sub, err := _Blockchain.contract.FilterLogs(opts, "RewardExchanged", exchangeIdRule, rewardIdRule, childRule)
	if err != nil {
		return nil, err
	}
	return &BlockchainRewardExchangedIterator{contract: _Blockchain.contract, event: "RewardExchanged", logs: logs, sub: sub}, nil
}

// WatchRewardExchanged is a free log subscription operation binding the contract event 0xe150ee44da4a37592e6368bbdc0d8901353c1bfb770394b3268e70c9b5055050.
//
// Solidity: event RewardExchanged(uint256 indexed exchangeId, uint256 indexed rewardId, address indexed child, uint256 tokenAmount)
func (_Blockchain *BlockchainFilterer) WatchRewardExchanged(opts *bind.WatchOpts, sink chan<- *BlockchainRewardExchanged, exchangeId []*big.Int, rewardId []*big.Int, child []common.Address) (event.Subscription, error) {

	var exchangeIdRule []interface{}
	for _, exchangeIdItem := range exchangeId {
		exchangeIdRule = append(exchangeIdRule, exchangeIdItem)
	}
	var rewardIdRule []interface{}
	for _, rewardIdItem := range rewardId {
		rewardIdRule = append(rewardIdRule, rewardIdItem)
	}
	var childRule []interface{}
	for _, childItem := range child {
		childRule = append(childRule, childItem)
	}

	logs, sub, err := _Blockchain.contract.WatchLogs(opts, "RewardExchanged", exchangeIdRule, rewardIdRule, childRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockchainRewardExchanged)
				if err := _Blockchain.contract.UnpackLog(event, "RewardExchanged", log); err != nil {
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

// ParseRewardExchanged is a log parse operation binding the contract event 0xe150ee44da4a37592e6368bbdc0d8901353c1bfb770394b3268e70c9b5055050.
//
// Solidity: event RewardExchanged(uint256 indexed exchangeId, uint256 indexed rewardId, address indexed child, uint256 tokenAmount)
func (_Blockchain *BlockchainFilterer) ParseRewardExchanged(log types.Log) (*BlockchainRewardExchanged, error) {
	event := new(BlockchainRewardExchanged)
	if err := _Blockchain.contract.UnpackLog(event, "RewardExchanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockchainRewardUpdatedIterator is returned from FilterRewardUpdated and is used to iterate over the raw logs and unpacked data for RewardUpdated events raised by the Blockchain contract.
type BlockchainRewardUpdatedIterator struct {
	Event *BlockchainRewardUpdated // Event containing the contract specifics and raw log

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
func (it *BlockchainRewardUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockchainRewardUpdated)
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
		it.Event = new(BlockchainRewardUpdated)
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
func (it *BlockchainRewardUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockchainRewardUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockchainRewardUpdated represents a RewardUpdated event raised by the Blockchain contract.
type BlockchainRewardUpdated struct {
	RewardId   *big.Int
	Name       string
	TokenPrice *big.Int
	Active     bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRewardUpdated is a free log retrieval operation binding the contract event 0x22cedbe07fd5a5dca8cc6b034fc2cd0f0b9802543d3efe5c4b705ac3dcbcf522.
//
// Solidity: event RewardUpdated(uint256 indexed rewardId, string name, uint256 tokenPrice, bool active)
func (_Blockchain *BlockchainFilterer) FilterRewardUpdated(opts *bind.FilterOpts, rewardId []*big.Int) (*BlockchainRewardUpdatedIterator, error) {

	var rewardIdRule []interface{}
	for _, rewardIdItem := range rewardId {
		rewardIdRule = append(rewardIdRule, rewardIdItem)
	}

	logs, sub, err := _Blockchain.contract.FilterLogs(opts, "RewardUpdated", rewardIdRule)
	if err != nil {
		return nil, err
	}
	return &BlockchainRewardUpdatedIterator{contract: _Blockchain.contract, event: "RewardUpdated", logs: logs, sub: sub}, nil
}

// WatchRewardUpdated is a free log subscription operation binding the contract event 0x22cedbe07fd5a5dca8cc6b034fc2cd0f0b9802543d3efe5c4b705ac3dcbcf522.
//
// Solidity: event RewardUpdated(uint256 indexed rewardId, string name, uint256 tokenPrice, bool active)
func (_Blockchain *BlockchainFilterer) WatchRewardUpdated(opts *bind.WatchOpts, sink chan<- *BlockchainRewardUpdated, rewardId []*big.Int) (event.Subscription, error) {

	var rewardIdRule []interface{}
	for _, rewardIdItem := range rewardId {
		rewardIdRule = append(rewardIdRule, rewardIdItem)
	}

	logs, sub, err := _Blockchain.contract.WatchLogs(opts, "RewardUpdated", rewardIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockchainRewardUpdated)
				if err := _Blockchain.contract.UnpackLog(event, "RewardUpdated", log); err != nil {
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

// ParseRewardUpdated is a log parse operation binding the contract event 0x22cedbe07fd5a5dca8cc6b034fc2cd0f0b9802543d3efe5c4b705ac3dcbcf522.
//
// Solidity: event RewardUpdated(uint256 indexed rewardId, string name, uint256 tokenPrice, bool active)
func (_Blockchain *BlockchainFilterer) ParseRewardUpdated(log types.Log) (*BlockchainRewardUpdated, error) {
	event := new(BlockchainRewardUpdated)
	if err := _Blockchain.contract.UnpackLog(event, "RewardUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
