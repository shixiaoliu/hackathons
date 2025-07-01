// Code generated - DO NOT EDIT.
// This file is a generated binding for RewardToken contract

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

// RewardToken is an auto generated Go binding around an Ethereum contract.
type RewardToken struct {
	RewardTokenCaller     // Read-only binding to the contract
	RewardTokenTransactor // Write-only binding to the contract
	RewardTokenFilterer   // Log filterer for contract events
}

// RewardTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type RewardTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RewardTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RewardTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
// Solidity: function balanceOf(address account) view returns(uint256)
func (c *RewardTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := c.contract.Call(opts, &out, "balanceOf", account)
	if err != nil {
		return nil, err
	}
	return *abi.ConvertType(out[0], new(*big.Int)).(**big.Int), nil
}
