// 这是修复后的 RewardRegistry 绑定
// 请用这个文件替换 familyChain-backend/pkg/blockchain/reward_registry_generated.go

package blockchain

import (
	"errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// RewardRegistryABI 是 RewardRegistry 合约的完整 ABI
const RewardRegistryABI = `[
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "_tokenAddress",
          "type": "address"
        }
      ],
      "stateMutability": "nonpayable",
      "type": "constructor"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "owner",
          "type": "address"
        }
      ],
      "name": "OwnableInvalidOwner",
      "type": "error"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "account",
          "type": "address"
        }
      ],
      "name": "OwnableUnauthorizedAccount",
      "type": "error"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "uint256",
          "name": "exchangeId",
          "type": "uint256"
        },
        {
          "indexed": true,
          "internalType": "address",
          "name": "parent",
          "type": "address"
        }
      ],
      "name": "ExchangeFulfilled",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "address",
          "name": "previousOwner",
          "type": "address"
        },
        {
          "indexed": true,
          "internalType": "address",
          "name": "newOwner",
          "type": "address"
        }
      ],
      "name": "OwnershipTransferred",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "uint256",
          "name": "rewardId",
          "type": "uint256"
        },
        {
          "indexed": true,
          "internalType": "address",
          "name": "creator",
          "type": "address"
        },
        {
          "indexed": true,
          "internalType": "uint256",
          "name": "familyId",
          "type": "uint256"
        },
        {
          "indexed": false,
          "internalType": "string",
          "name": "name",
          "type": "string"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "tokenPrice",
          "type": "uint256"
        }
      ],
      "name": "RewardCreated",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "uint256",
          "name": "exchangeId",
          "type": "uint256"
        },
        {
          "indexed": true,
          "internalType": "uint256",
          "name": "rewardId",
          "type": "uint256"
        },
        {
          "indexed": true,
          "internalType": "address",
          "name": "child",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "tokenAmount",
          "type": "uint256"
        }
      ],
      "name": "RewardExchanged",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "uint256",
          "name": "rewardId",
          "type": "uint256"
        },
        {
          "indexed": false,
          "internalType": "string",
          "name": "name",
          "type": "string"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "tokenPrice",
          "type": "uint256"
        },
        {
          "indexed": false,
          "internalType": "bool",
          "name": "active",
          "type": "bool"
        }
      ],
      "name": "RewardUpdated",
      "type": "event"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "_child",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "_index",
          "type": "uint256"
        }
      ],
      "name": "getChildExchangeId",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "_familyId",
          "type": "uint256"
        },
        {
          "internalType": "string",
          "name": "_name",
          "type": "string"
        },
        {
          "internalType": "string",
          "name": "_description",
          "type": "string"
        },
        {
          "internalType": "string",
          "name": "_imageURI",
          "type": "string"
        },
        {
          "internalType": "uint256",
          "name": "_tokenPrice",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "_stock",
          "type": "uint256"
        }
      ],
      "name": "createReward",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "_rewardId",
          "type": "uint256"
        }
      ],
      "name": "exchangeReward",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "_exchangeId",
          "type": "uint256"
        }
      ],
      "name": "fulfillExchange",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "_child",
          "type": "address"
        }
      ],
      "name": "getChildExchangeCount",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "_exchangeId",
          "type": "uint256"
        }
      ],
      "name": "getExchange",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        },
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        },
        {
          "internalType": "bool",
          "name": "",
          "type": "bool"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "_familyId",
          "type": "uint256"
        }
      ],
      "name": "getFamilyRewardCount",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "_familyId",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "_index",
          "type": "uint256"
        }
      ],
      "name": "getFamilyRewardId",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "_rewardId",
          "type": "uint256"
        }
      ],
      "name": "getReward",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        },
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        },
        {
          "internalType": "string",
          "name": "",
          "type": "string"
        },
        {
          "internalType": "string",
          "name": "",
          "type": "string"
        },
        {
          "internalType": "string",
          "name": "",
          "type": "string"
        },
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        },
        {
          "internalType": "bool",
          "name": "",
          "type": "bool"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "owner",
      "outputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "renounceOwnership",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "newOwner",
          "type": "address"
        }
      ],
      "name": "transferOwnership",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "_rewardId",
          "type": "uint256"
        },
        {
          "internalType": "string",
          "name": "_name",
          "type": "string"
        },
        {
          "internalType": "string",
          "name": "_description",
          "type": "string"
        },
        {
          "internalType": "string",
          "name": "_imageURI",
          "type": "string"
        },
        {
          "internalType": "uint256",
          "name": "_tokenPrice",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "_stock",
          "type": "uint256"
        },
        {
          "internalType": "bool",
          "name": "_active",
          "type": "bool"
        }
      ],
      "name": "updateReward",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    }
]`

// RewardRegistry 是与 RewardRegistry 合约交互的绑定
type RewardRegistry struct {
	RewardRegistryCaller     // 只读绑定
	RewardRegistryTransactor // 写入绑定
	RewardRegistryFilterer   // 事件过滤器
}

// RewardRegistryCaller 是用于合约只读调用的绑定
type RewardRegistryCaller struct {
	contract *bind.BoundContract
}

// RewardRegistryTransactor 是用于合约写入操作的绑定
type RewardRegistryTransactor struct {
	contract *bind.BoundContract
}

// RewardRegistryFilterer 是用于合约事件过滤的绑定
type RewardRegistryFilterer struct {
	contract *bind.BoundContract
}

// NewRewardRegistry 创建一个新的 RewardRegistry 合约绑定实例
func NewRewardRegistry(address common.Address, backend bind.ContractBackend) (*RewardRegistry, error) {
	contract, err := bindRewardRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RewardRegistry{
		RewardRegistryCaller:     RewardRegistryCaller{contract: contract},
		RewardRegistryTransactor: RewardRegistryTransactor{contract: contract},
		RewardRegistryFilterer:   RewardRegistryFilterer{contract: contract},
	}, nil
}

// bindRewardRegistry 绑定一个合约实例
func bindRewardRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RewardRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// RewardInfo 奖励信息结构体
type RewardInfo struct {
	Id          *big.Int
	Creator     common.Address
	FamilyId    *big.Int
	Name        string
	Description string
	ImageURI    string
	TokenPrice  *big.Int
	Stock       *big.Int
	Active      bool
}

// ExchangeInfo 兑换记录信息结构体
type ExchangeInfo struct {
	Id           *big.Int
	RewardId     *big.Int
	Child        common.Address
	TokenAmount  *big.Int
	ExchangeDate *big.Int
	Fulfilled    bool
}

// CreateReward 创建新奖励
func (r *RewardRegistryTransactor) CreateReward(opts *bind.TransactOpts, familyId *big.Int, name string, description string, imageURI string, tokenPrice *big.Int, stock *big.Int) (*types.Transaction, error) {
	return r.contract.Transact(opts, "createReward", familyId, name, description, imageURI, tokenPrice, stock)
}

// UpdateReward 更新奖励信息
func (r *RewardRegistryTransactor) UpdateReward(opts *bind.TransactOpts, rewardId *big.Int, name string, description string, imageURI string, tokenPrice *big.Int, stock *big.Int, active bool) (*types.Transaction, error) {
	return r.contract.Transact(opts, "updateReward", rewardId, name, description, imageURI, tokenPrice, stock, active)
}

// ExchangeReward 兑换奖励
func (r *RewardRegistryTransactor) ExchangeReward(opts *bind.TransactOpts, rewardId *big.Int) (*types.Transaction, error) {
	return r.contract.Transact(opts, "exchangeReward", rewardId)
}

// FulfillExchange 标记兑换为已完成
func (r *RewardRegistryTransactor) FulfillExchange(opts *bind.TransactOpts, exchangeId *big.Int) (*types.Transaction, error) {
	return r.contract.Transact(opts, "fulfillExchange", exchangeId)
}

// GetReward 获取奖励信息
func (r *RewardRegistryCaller) GetReward(opts *bind.CallOpts, rewardId *big.Int) (RewardInfo, error) {
	var out []interface{}
	err := r.contract.Call(opts, &out, "getReward", rewardId)

	if err != nil {
		return RewardInfo{}, err
	}

	if len(out) == 0 {
		return RewardInfo{}, errors.New("reward not found")
	}

	result := RewardInfo{
		Id:          out[0].(*big.Int),
		Creator:     out[1].(common.Address),
		FamilyId:    out[2].(*big.Int),
		Name:        out[3].(string),
		Description: out[4].(string),
		ImageURI:    out[5].(string),
		TokenPrice:  out[6].(*big.Int),
		Stock:       out[7].(*big.Int),
		Active:      out[8].(bool),
	}

	return result, nil
}

// GetExchange 获取兑换记录信息
func (r *RewardRegistryCaller) GetExchange(opts *bind.CallOpts, exchangeId *big.Int) (ExchangeInfo, error) {
	var out []interface{}
	err := r.contract.Call(opts, &out, "getExchange", exchangeId)

	if err != nil {
		return ExchangeInfo{}, err
	}

	if len(out) == 0 {
		return ExchangeInfo{}, errors.New("exchange not found")
	}

	result := ExchangeInfo{
		Id:           out[0].(*big.Int),
		RewardId:     out[1].(*big.Int),
		Child:        out[2].(common.Address),
		TokenAmount:  out[3].(*big.Int),
		ExchangeDate: out[4].(*big.Int),
		Fulfilled:    out[5].(bool),
	}

	return result, nil
}

// GetFamilyRewardCount 获取家庭奖励数量
func (r *RewardRegistryCaller) GetFamilyRewardCount(opts *bind.CallOpts, familyId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := r.contract.Call(opts, &out, "getFamilyRewardCount", familyId)

	if err != nil {
		return nil, err
	}

	return out[0].(*big.Int), nil
}

// GetFamilyRewardId 获取家庭奖励ID
func (r *RewardRegistryCaller) GetFamilyRewardId(opts *bind.CallOpts, familyId *big.Int, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := r.contract.Call(opts, &out, "getFamilyRewardId", familyId, index)

	if err != nil {
		return nil, err
	}

	return out[0].(*big.Int), nil
}

// GetChildExchangeCount 获取孩子兑换记录数量
func (r *RewardRegistryCaller) GetChildExchangeCount(opts *bind.CallOpts, child common.Address) (*big.Int, error) {
	var out []interface{}
	err := r.contract.Call(opts, &out, "getChildExchangeCount", child)

	if err != nil {
		return nil, err
	}

	return out[0].(*big.Int), nil
}

// GetChildExchangeId 获取孩子兑换记录ID
func (r *RewardRegistryCaller) GetChildExchangeId(opts *bind.CallOpts, child common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := r.contract.Call(opts, &out, "getChildExchangeId", child, index)

	if err != nil {
		return nil, err
	}

	return out[0].(*big.Int), nil
}

// RewardCount 获取奖励总数
func (r *RewardRegistryCaller) RewardCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := r.contract.Call(opts, &out, "rewardCount")

	if err != nil {
		return nil, err
	}

	return out[0].(*big.Int), nil
}

// ExchangeCount 获取兑换总数
func (r *RewardRegistryCaller) ExchangeCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := r.contract.Call(opts, &out, "exchangeCount")

	if err != nil {
		return nil, err
	}

	return out[0].(*big.Int), nil
}
